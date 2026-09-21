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

// Code the model writes runs here, in a subprocess on the worker host. There is
// no sandbox: AllowedLanguages and AllowedCommands are the only limits, and
// AllowedCommands is only a best-effort scan of the source. Do not attach
// CodeExecutionConfig to an agent handling untrusted input unless the worker
// itself is isolated. Mirrors Python's LocalCodeExecutor and CommandValidator.

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

// codeExecIn is what the model sends the derived execute_code tool; its fields
// match the schema in CodeExecutionConfig.codeTool.
type codeExecIn struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

// codeExecOut is what the tool returns; errors in the executed code are
// reported here rather than failing the task, so the model can try again.
type codeExecOut struct {
	Status string `json:"status"`
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

func (c *CodeExecutionConfig) executeCode(ctx context.Context, in codeExecIn) (codeExecOut, error) {
	// Python answers a missing snippet with a success and an explanatory
	// stdout rather than an error, so the run continues.
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

	timeout := timeoutOrDefault(c.TimeoutSeconds)
	// A local executor is language-specific, so one is built per call; any
	// other executor runs as configured, whatever the language argument says.
	executor := c.Executor
	switch l := executor.(type) {
	case nil:
		executor = LocalExecutor{Language: language, TimeoutSeconds: timeout}
	case LocalExecutor:
		l.Language = language
		if l.TimeoutSeconds == 0 {
			l.TimeoutSeconds = timeout
		}
		executor = l
	}
	return executor.Execute(ctx, in.Code).toolOutput(timeout), nil
}

// runInterpreter is the LocalExecutor's engine. workingDir empty means the
// temporary file's own directory.
func runInterpreter(ctx context.Context, interpreter []string, code, ext string,
	timeoutSeconds int, workingDir string) ExecutionResult {

	f, err := os.CreateTemp("", "conductor_code_*"+ext)
	if err != nil {
		return ExecutionResult{Error: err.Error(), ExitCode: 1}
	}
	path := f.Name()
	// Removed however this returns, since a failed write still leaves a file
	// behind; a leftover temp file is not a failure, so the error has no consumer.
	defer os.Remove(path) //nolint:errcheck // see above

	if _, werr := f.WriteString(code); werr != nil {
		return ExecutionResult{Error: errors.Join(werr, f.Close()).Error(), ExitCode: 1}
	}
	if cerr := f.Close(); cerr != nil {
		return ExecutionResult{Error: cerr.Error(), ExitCode: 1}
	}

	// The config's timeout applies, but the task's context wins if it expires
	// first: a worker should not outlive the run it belongs to.
	runCtx, cancel := context.WithTimeout(ctx, time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	args := append(append([]string{}, interpreter[1:]...), path)
	// The interpreter is operator configuration; the model's code reaches it as a
	// file path, never as arguments.
	cmd := exec.CommandContext(runCtx, interpreter[0], args...) //nolint:gosec // see above
	cmd.Dir = filepath.Dir(path)
	if workingDir != "" {
		cmd.Dir = workingDir
	}

	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	res := ExecutionResult{Output: stdout.String(), Error: stderr.String()}

	switch {
	case err == nil:
		return res
	case runCtx.Err() != nil:
		res.ExitCode, res.TimedOut = -1, true
		return res
	default:
		// 127 is the shell's "command not found" (interpreter missing, permission
		// denied), where the error itself is the only useful detail.
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			res.ExitCode = ee.ExitCode()
		} else {
			res.ExitCode = 127
			res.Error = strings.TrimRight(res.Error, "\n")
			if res.Error != "" {
				res.Error += "\n"
			}
			res.Error += err.Error()
		}
		return res
	}
}

// Command validation, ported from Python's CommandValidator.

// Matched in order: subprocess.*, os.system/os.popen, and Jupyter's ! syntax.
var pythonShellPatterns = []*regexp.Regexp{
	regexp.MustCompile(`subprocess\.\w+\(\s*\[?\s*["'](\S+?)["']`),
	regexp.MustCompile(`os\.(?:system|popen)\(\s*["'](\S+)`),
	regexp.MustCompile(`(?m)^\s*!(\S+)`),
}

var (
	bashCommandRe = regexp.MustCompile(`(?m)(?:^|[|;&]\s*|` + "`" + `|\$\(\s*)(\w[\w.+-]*)`)
	heredocRe     = regexp.MustCompile(`<<-?\s*'?(\w+)'?`)
)

// bashBuiltins are shell keywords, never checked against the allow-list.
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
	// A heredoc delimiter looks like a command to the scanner, so skip them.
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
		// handled, which can only make the scan stricter.
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
