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
	"reflect"
	"strings"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// shellSplit stands in for Python's shlex.split, so the same command string
// has to produce the same argv in both SDKs.
func TestShellSplit(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{`gh repo list --limit 5`, []string{"gh", "repo", "list", "--limit", "5"}},
		{`echo "hello world"`, []string{"echo", "hello world"}},
		{`echo 'it''s'`, []string{"echo", "its"}},
		{`printf "a\"b"`, []string{"printf", `a"b`}},
		{`ls  -la   /tmp`, []string{"ls", "-la", "/tmp"}},
		{`echo a\ b`, []string{"echo", "a b"}},
		{``, nil},
		{`   `, nil},
	}
	for _, tc := range cases {
		got, err := shellSplit(tc.in)
		if err != nil {
			t.Errorf("shellSplit(%q): unexpected error %v", tc.in, err)
			continue
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("shellSplit(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}

	// An unterminated quote is a parse error, not a best guess.
	if _, err := shellSplit(`echo "unterminated`); err == nil {
		t.Error("shellSplit accepted an unterminated quote")
	}
}

func TestShellQuote(t *testing.T) {
	cases := map[string]string{
		"simple":     "simple",
		"/usr/bin":   "/usr/bin",
		"has space":  "'has space'",
		"it's":       `'it'"'"'s'`,
		"":           "''",
		"$HOME":      "'$HOME'",
		"a;rm -rf /": "'a;rm -rf /'",
	}
	for in, want := range cases {
		if got := shellQuote(in); got != want {
			t.Errorf("shellQuote(%q) = %s, want %s", in, got, want)
		}
	}
}

// The allow-list keys off the executable's base name, so a full path and a
// bare name validate the same way; an empty list permits everything.
func TestValidateCLICommand(t *testing.T) {
	if msg := validateCLICommand("rm", nil); msg != "" {
		t.Errorf("empty allow-list should permit everything, got %q", msg)
	}
	if msg := validateCLICommand("/usr/bin/git", []string{"git", "ls"}); msg != "" {
		t.Errorf("path prefix should be stripped before checking, got %q", msg)
	}
	msg := validateCLICommand("rm", []string{"ls", "git"})
	if !strings.Contains(msg, "'rm' is not allowed") || !strings.Contains(msg, "git, ls") {
		t.Errorf("rejection message = %q, want the command and the sorted list", msg)
	}
}

// A command that runs and exits non-zero is a result the model reads. A command
// that cannot run is a terminal failure. That split is Python's, and the
// difference matters: the first lets the model try again, the second stops a
// retry loop that would never succeed.
func TestRunCommandOutcomes(t *testing.T) {
	ctx := context.Background()
	c := &CLIConfig{TimeoutSeconds: 10}

	// Success, with args embedded in the command line and in the list, in
	// that order.
	out, err := c.runCommand(ctx, cliIn{Command: "echo one", Args: []any{"two", 3}})
	if err != nil {
		t.Fatalf("echo: %v", err)
	}
	if out.Status != "success" || strings.TrimSpace(out.Stdout) != "one two 3" {
		t.Errorf("echo = %+v", out)
	}

	// Non-zero exit is a result, not an error.
	out, err = c.runCommand(ctx, cliIn{Command: "sh", Args: []any{"-c", "echo oops >&2; exit 3"}})
	if err != nil {
		t.Fatalf("non-zero exit should not be an error: %v", err)
	}
	if out.Status != "error" || out.ExitCode != 3 || !strings.Contains(out.Stderr, "oops") {
		t.Errorf("non-zero exit = %+v", out)
	}

	// Not found is terminal.
	_, err = c.runCommand(ctx, cliIn{Command: "definitely-not-a-real-command-xyz"})
	var nre *model.NonRetryableError
	if !errors.As(err, &nre) {
		t.Errorf("command not found should be a NonRetryableError, got %T: %v", err, err)
	}

	// Timeout is terminal too.
	short := &CLIConfig{TimeoutSeconds: 1}
	_, err = short.runCommand(ctx, cliIn{Command: "sleep", Args: []any{"5"}})
	if !errors.As(err, &nre) || !strings.Contains(err.Error(), "timed out") {
		t.Errorf("timeout should be a NonRetryableError mentioning the timeout, got %v", err)
	}

	// An empty command is a result the model can read.
	out, err = c.runCommand(ctx, cliIn{Command: "   "})
	if err != nil || out.Status != "error" || !strings.Contains(out.Stderr, "No command") {
		t.Errorf("empty command = %+v, %v", out, err)
	}
}

// Config violations fail the task as ordinary errors, as Python's ValueError
// does — retryable, and not shown to the model.
func TestRunCommandEnforcesConfig(t *testing.T) {
	ctx := context.Background()

	restricted := &CLIConfig{AllowedCommands: []string{"ls"}}
	_, err := restricted.runCommand(ctx, cliIn{Command: "rm -rf /tmp/x"})
	if err == nil || !strings.Contains(err.Error(), "'rm' is not allowed") {
		t.Errorf("disallowed command should error, got %v", err)
	}
	var nre *model.NonRetryableError
	if errors.As(err, &nre) {
		t.Error("a disallowed command is an ordinary error in Python, not terminal")
	}

	noShell := &CLIConfig{}
	_, err = noShell.runCommand(ctx, cliIn{Command: "echo hi", Shell: true})
	if err == nil || !strings.Contains(err.Error(), "shell mode is disabled") {
		t.Errorf("shell request against AllowShell=false should error, got %v", err)
	}

	// With shell allowed, shell features work and quoting keeps argv intact.
	withShell := &CLIConfig{AllowShell: true, TimeoutSeconds: 10}
	out, err := withShell.runCommand(ctx, cliIn{Command: "echo", Args: []any{"a b", "$HOME"}, Shell: true})
	if err != nil {
		t.Fatalf("shell echo: %v", err)
	}
	// $HOME is quoted, so it arrives literally rather than expanded.
	if strings.TrimSpace(out.Stdout) != "a b $HOME" {
		t.Errorf("shell echo stdout = %q, want %q", out.Stdout, "a b $HOME")
	}
}

// Cwd is honoured: the same command sees a different directory.
func TestRunCommandCwd(t *testing.T) {
	c := &CLIConfig{TimeoutSeconds: 10}
	out, err := c.runCommand(context.Background(), cliIn{Command: "pwd", Cwd: "/"})
	if err != nil {
		t.Fatalf("pwd: %v", err)
	}
	if strings.TrimSpace(out.Stdout) != "/" {
		t.Errorf("pwd in / = %q", out.Stdout)
	}
}
