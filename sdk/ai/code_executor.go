//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// Code the model writes runs here, in a subprocess on the worker host. There
// is no sandbox: AllowedLanguages and AllowedCommands are the only limits, and
// AllowedCommands is a best-effort scan of the source rather than a guarantee.
// Do not attach CodeExecutionConfig to an agent handling untrusted input
// unless the worker itself is isolated.
//
// This mirrors Python's LocalCodeExecutor and CommandValidator so the same
// agent behaves the same way in both SDKs.

// interpreters maps a language to the command that runs it.
var interpreters = map[string][]string{
	"python":     {"python3"},
	"python3":    {"python3"},
	"bash":       {"bash"},
	"sh":         {"sh"},
	"node":       {"node"},
	"javascript": {"node"},
	"ruby":       {"ruby"},
}

var fileExtensions = map[string]string{
	"python":     ".py",
	"python3":    ".py",
	"bash":       ".sh",
	"sh":         ".sh",
	"node":       ".js",
	"javascript": ".js",
	"ruby":       ".rb",
}

// codeExecIn is what the model sends the derived execute_code tool. The field
// names match the schema in CodeExecutionConfig.codeTool.
type codeExecIn struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

// codeExecOut is what the tool returns. Errors in the executed code are
// reported here rather than failing the task, so the model can read the
// stderr and try again.
type codeExecOut struct {
	Status string `json:"status"`
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

// executeCode runs one snippet and reports the outcome.
func (c *CodeExecutionConfig) executeCode(ctx context.Context, in codeExecIn) (codeExecOut, error) {
	// The model sometimes omits arguments. Python answers with a success and
	// an explanatory stdout rather than an error, so the run continues.
	if in.Code == "" {
		return codeExecOut{
			Status: "success",
			Stdout: "No code provided. Nothing to execute.",
		}, nil
	}
	language := in.Language
	if language == "" {
		language = defaultLanguage
	}

	allowed := languagesOrDefault(c.AllowedLanguages)
	if !contains(allowed, language) {
		return codeExecOut{
			Status: "error",
			Stderr: fmt.Sprintf("Language %q is not allowed. Allowed: %s",
				language, strings.Join(allowed, ", ")),
		}, nil
	}

	if msg := validateCommands(in.Code, language, c.AllowedCommands); msg != "" {
		return codeExecOut{Status: "error", Stderr: msg}, nil
	}

	interpreter, ok := interpreters[language]
	if !ok {
		return codeExecOut{
			Status: "error",
			Stderr: fmt.Sprintf("Unsupported language: %s", language),
		}, nil
	}

	return runInterpreter(ctx, interpreter, in.Code,
		fileExtensions[language], timeoutOrDefault(c.TimeoutSeconds))
}

// runInterpreter writes the code to a temporary file and runs it.
func runInterpreter(ctx context.Context, interpreter []string, code, ext string,
	timeoutSeconds int) (codeExecOut, error) {

	f, err := os.CreateTemp("", "conductor_code_*"+ext)
	if err != nil {
		return codeExecOut{Status: "error", Stderr: err.Error()}, nil
	}
	path := f.Name()
	// Removed however this returns: a failed write still leaves a file behind.
	// A leftover temp file is not an execution failure, so its removal error
	// has no consumer.
	defer os.Remove(path) //nolint:errcheck // see above

	if _, werr := f.WriteString(code); werr != nil {
		return codeExecOut{Status: "error", Stderr: errors.Join(werr, f.Close()).Error()}, nil
	}
	if cerr := f.Close(); cerr != nil {
		return codeExecOut{Status: "error", Stderr: cerr.Error()}, nil
	}

	// The timeout is the config's, but the task's context still wins if it
	// expires first — a worker should not outlive the run it belongs to.
	runCtx, cancel := context.WithTimeout(ctx, time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	args := append(append([]string{}, interpreter[1:]...), path)
	// The interpreter is operator configuration; the model's code reaches it
	// as a file path, never as arguments.
	cmd := exec.CommandContext(runCtx, interpreter[0], args...) //nolint:gosec // see above
	cmd.Dir = filepath.Dir(path)

	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	out := codeExecOut{Status: "success", Stdout: stdout.String(), Stderr: stderr.String()}

	switch {
	case err == nil:
		return out, nil
	case runCtx.Err() != nil:
		// Distinguished from an ordinary failure because the model can act on
		// it: shorter code, or a smaller problem.
		out.Status = "error"
		out.Stderr = strings.TrimRight(out.Stderr, "\n")
		if out.Stderr != "" {
			out.Stderr += "\n"
		}
		out.Stderr += fmt.Sprintf("TIMED OUT after %ds", timeoutSeconds)
		return out, nil
	default:
		out.Status = "error"
		parts := []string{}
		if s := strings.TrimRight(out.Stderr, "\n"); s != "" {
			parts = append(parts, s)
		}
		// 127 is the shell's "command not found": interpreter missing,
		// permission denied, and so on. Then the error itself is the only
		// useful detail.
		exitCode := 127
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			exitCode = ee.ExitCode()
		} else {
			parts = append(parts, err.Error())
		}
		parts = append(parts, fmt.Sprintf("Exit code: %d", exitCode))
		out.Stderr = strings.Join(parts, "\n")
		return out, nil
	}
}

// Command validation. A best-effort scan for commands the code shells out to,
// ported from Python's CommandValidator. It reads source text, so it can be
// evaded; it exists to catch mistakes, not to contain an adversary.

var pythonShellPatterns = []*regexp.Regexp{
	// subprocess.run(["cmd", ...]) and friends
	regexp.MustCompile(`subprocess\.\w+\(\s*\[?\s*["'](\S+?)["']`),
	// os.system("cmd ...") / os.popen("cmd ...")
	regexp.MustCompile(`os\.(?:system|popen)\(\s*["'](\S+)`),
	// Jupyter ! syntax
	regexp.MustCompile(`(?m)^\s*!(\S+)`),
}

var (
	bashCommandRe = regexp.MustCompile(`(?m)(?:^|[|;&]\s*|` + "`" + `|\$\(\s*)(\w[\w.+-]*)`)
	heredocRe     = regexp.MustCompile(`<<-?\s*'?(\w+)'?`)
)

// bashBuiltins are shell keywords rather than commands, so they are never
// checked against the allow-list.
var bashBuiltins = map[string]struct{}{}

func init() {
	for _, b := range strings.Fields(
		"if then else elif fi for while do done case esac in function select until " +
			"echo printf read local export unset set shift return exit true false test " +
			"declare typeset readonly source eval exec trap wait break continue cd " +
			"pushd popd pwd dirs hash type command builtin enable let shopt complete compgen") {
		bashBuiltins[b] = struct{}{}
	}
}

// validateCommands returns an error message, or "" when the code passes. An
// empty allow-list means no restriction, matching Python.
func validateCommands(code, language string, allowedCommands []string) string {
	if len(allowedCommands) == 0 {
		return ""
	}
	allowed := make(map[string]struct{}, len(allowedCommands))
	for _, c := range allowedCommands {
		allowed[c] = struct{}{}
	}

	switch language {
	case "python", "python3":
		return validatePython(code, allowed, allowedCommands)
	case "bash", "sh":
		return validateBash(code, allowed, allowedCommands)
	default:
		// Other languages are not scanned, as in Python.
		return ""
	}
}

func validatePython(code string, allowed map[string]struct{}, list []string) string {
	for _, re := range pythonShellPatterns {
		for _, m := range re.FindAllStringSubmatch(code, -1) {
			// Handle an absolute path: /usr/bin/git is git.
			cmd := m[1]
			if i := strings.LastIndex(cmd, "/"); i >= 0 {
				cmd = cmd[i+1:]
			}
			if _, ok := allowed[cmd]; !ok {
				return notAllowed(cmd, list)
			}
		}
	}
	return ""
}

func validateBash(code string, allowed map[string]struct{}, list []string) string {
	// A heredoc delimiter looks like a command to the scanner, so collect the
	// delimiters first and skip them.
	delimiters := map[string]struct{}{}
	for _, m := range heredocRe.FindAllStringSubmatch(code, -1) {
		delimiters[m[1]] = struct{}{}
	}

	lines := strings.Split(code, "\n")
	cleaned := make([]string, 0, len(lines))
	for _, line := range lines {
		if strings.HasPrefix(strings.TrimLeft(line, " \t"), "#") {
			continue
		}
		// Naive inline-comment strip, as in Python: a # inside quotes is not
		// handled, which can only make the scan stricter, never looser.
		if i := strings.Index(line, " #"); i >= 0 {
			line = line[:i]
		}
		cleaned = append(cleaned, line)
	}

	for _, m := range bashCommandRe.FindAllStringSubmatch(strings.Join(cleaned, "\n"), -1) {
		cmd := m[1]
		if _, ok := bashBuiltins[cmd]; ok {
			continue
		}
		if _, ok := delimiters[cmd]; ok {
			continue
		}
		if _, ok := allowed[cmd]; !ok {
			return notAllowed(cmd, list)
		}
	}
	return ""
}

func notAllowed(cmd string, list []string) string {
	sorted := append([]string(nil), list...)
	sort.Strings(sorted)
	return fmt.Sprintf("Command '%s' is not allowed. Allowed commands: %s",
		cmd, strings.Join(sorted, ", "))
}

func contains(haystack []string, needle string) bool {
	for _, h := range haystack {
		if h == needle {
			return true
		}
	}
	return false
}
