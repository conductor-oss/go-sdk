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
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// The derived run_command tool runs on the worker host with the worker's
// privileges, limited only by AllowedCommands: as with code execution, do not
// attach CLIConfig to an agent handling untrusted input unless it is isolated.
//
// Ported from Python's _CliCommandRunner, including how failures are reported:
// a non-zero exit is a result the model can act on; a command that cannot run
// at all — not found, timed out — is a terminal task failure, since a retry
// will not help; a disallowed command or a forbidden shell request is an
// ordinary error, as Python's ValueError.

// cliIn is what the model sends; field names match the schema in
// CLIConfig.cliTool, and Args is []any because the model may not send strings.
type cliIn struct {
	Command    string `json:"command"`
	Args       []any  `json:"args"`
	Cwd        string `json:"cwd"`
	Shell      bool   `json:"shell"`
	ContextKey string `json:"context_key"`
}

type cliOut struct {
	Status   string `json:"status"`
	ExitCode int    `json:"exit_code"`
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
}

func (c *CLIConfig) runCommand(ctx context.Context, in cliIn) (cliOut, error) {
	if strings.TrimSpace(in.Command) == "" {
		return cliOut{Status: "error", Stderr: "No command provided."}, nil
	}

	// Models frequently pass the whole command line as Command, so tokenize: the
	// allow-list keys off the executable and execution gets a proper argv.
	tokens, err := shellSplit(in.Command)
	if err != nil {
		return cliOut{Status: "error", Stderr: "Could not parse command: " + err.Error()}, nil
	}
	if len(tokens) == 0 {
		return cliOut{Status: "error", Stderr: "No command provided."}, nil
	}
	executable := tokens[0]

	if msg := validateCLICommand(executable, c.AllowedCommands); msg != "" {
		return cliOut{}, errors.New(msg)
	}
	if in.Shell && !c.AllowShell {
		return cliOut{}, errors.New(
			"shell mode is disabled for this agent; do not set shell=true")
	}

	// Args embedded in the command line come first, then the explicit list.
	argv := append([]string{}, tokens[1:]...)
	for _, a := range in.Args {
		argv = append(argv, fmt.Sprint(a))
	}

	timeout := timeoutOrDefault(c.TimeoutSeconds)
	runCtx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()

	cmd := buildCommand(runCtx, executable, argv, in)

	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	out := cliOut{Stdout: stdout.String(), Stderr: stderr.String()}

	switch {
	case err == nil:
		out.Status = "success"
		if saveErr := saveToState(ctx, in.ContextKey, out); saveErr != nil {
			return cliOut{}, saveErr
		}
		return out, nil
	case runCtx.Err() != nil:
		return cliOut{}, model.NewNonRetryableError(
			fmt.Errorf("command timed out after %ds", timeout))
	default:
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			out.Status = "error"
			out.ExitCode = ee.ExitCode()
			return out, nil
		}
		if errors.Is(err, exec.ErrNotFound) {
			return cliOut{}, model.NewNonRetryableError(
				fmt.Errorf("command not found: %s", in.Command))
		}
		return cliOut{}, model.NewNonRetryableError(err)
	}
}

// validateCLICommand checks the executable against the allow-list, stripping
// any path prefix so /usr/bin/git and git validate alike; empty permits all.
// buildCommand assembles the process to run. Shell mode joins a fully quoted
// line so the shell sees exactly the argv the model built while still granting
// the shell features it asked for; direct mode execs the executable itself.
func buildCommand(ctx context.Context, executable string, argv []string, in cliIn) *exec.Cmd {
	var cmd *exec.Cmd
	if in.Shell {
		quoted := make([]string, 0, 1+len(argv))
		for _, a := range append([]string{executable}, argv...) {
			quoted = append(quoted, shellQuote(a))
		}
		// The executable passed validateCLICommand, shell mode is gated by
		// AllowShell, and every token is shell-quoted above.
		cmd = exec.CommandContext(ctx, "sh", "-c", strings.Join(quoted, " ")) //nolint:gosec // see above
	} else {
		cmd = exec.CommandContext(ctx, executable, argv...) //nolint:gosec // executable is allow-listed by validateCLICommand
	}
	if in.Cwd != "" {
		cmd.Dir = in.Cwd
	}
	return cmd
}

// saveToState puts a successful command's output into the agent state under
// key, as Python's runner does: trimmed stdout, or stderr when stdout is empty,
// and nothing at all when there is no key or nothing to save.
func saveToState(ctx context.Context, key string, out cliOut) error {
	if key == "" {
		return nil
	}
	value := strings.TrimSpace(out.Stdout)
	if value == "" {
		value = strings.TrimSpace(out.Stderr)
	}
	if value == "" {
		return nil
	}
	if err := SetState(ctx, key, value); err != nil {
		return fmt.Errorf("save output to state key %q: %w", key, err)
	}
	return nil
}

func validateCLICommand(executable string, allowed []string) string {
	if len(allowed) == 0 {
		return ""
	}
	base := filepath.Base(executable)
	for _, a := range allowed {
		if a == base {
			return ""
		}
	}
	sorted := append([]string(nil), allowed...)
	sort.Strings(sorted)
	return fmt.Sprintf("Command '%s' is not allowed. Allowed commands: %s",
		base, strings.Join(sorted, ", "))
}

// shellSplit tokenizes a command line as a POSIX shell would (quotes, backslash
// escapes): a stand-in for shlex.split, so both SDKs yield the same argv.
func shellSplit(s string) ([]string, error) {
	var tokens []string
	var cur strings.Builder
	inToken := false
	quote := rune(0)

	flush := func() {
		if inToken {
			tokens = append(tokens, cur.String())
			cur.Reset()
			inToken = false
		}
	}

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		switch {
		case quote == '\'':
			if r == '\'' {
				quote = 0
			} else {
				cur.WriteRune(r)
			}
		case quote == '"':
			var closed bool
			i, closed = readDoubleQuoted(runes, i, &cur)
			if closed {
				quote = 0
			}
		case r == '\'' || r == '"':
			quote = r
			inToken = true
		case r == '\\':
			if i+1 >= len(runes) {
				return nil, errors.New("no escaped character")
			}
			i++
			cur.WriteRune(runes[i])
			inToken = true
		case r == ' ' || r == '\t' || r == '\n':
			flush()
		default:
			cur.WriteRune(r)
			inToken = true
		}
	}
	if quote != 0 {
		return nil, errors.New("no closing quotation")
	}
	flush()
	return tokens, nil
}

// readDoubleQuoted consumes one rune of a double-quoted span at runes[i] and
// reports whether it was the closing quote, returning the last index consumed.
func readDoubleQuoted(runes []rune, i int, cur *strings.Builder) (int, bool) {
	r := runes[i]
	switch r {
	case '"':
		return i, true
	case '\\':
		// Inside double quotes only a few characters are escapable.
		if i+1 < len(runes) && strings.ContainsRune(`"\$`+"`", runes[i+1]) {
			i++
			cur.WriteRune(runes[i])
			return i, false
		}
	}
	cur.WriteRune(r)
	return i, false
}

// shellQuote wraps s in single quotes, matching Python's shlex.quote.
func shellQuote(s string) string {
	if s == "" {
		return "''"
	}
	if strings.IndexFunc(s, func(r rune) bool { return !shellSafeRune(r) }) < 0 {
		return s
	}
	return "'" + strings.ReplaceAll(s, "'", `'"'"'`) + "'"
}

// shellSafeRune reports whether r needs no quoting in a POSIX shell word.
func shellSafeRune(r rune) bool {
	switch {
	case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9':
		return true
	}
	return strings.ContainsRune("_-./=:@", r)
}
