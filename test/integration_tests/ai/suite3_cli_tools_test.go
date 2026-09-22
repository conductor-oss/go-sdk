//go:build integration

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package integration

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
	taskmodel "github.com/conductor-sdk/conductor-go/sdk/model"
)

// The Python SDK's e2e/test_suite3_cli_tools.py as a Go test, under the same
// name. Tools that wrap command-line programs get their credential only from
// the server's store, so a token sitting in the worker's environment is not
// used until it is stored; and the CLI tool's command allow-list is compiled
// into its description and enforced when the model tries something else.
//
// Live only: cli_mktemp returns a fresh temp path on every run, so no
// recording can match, and cli_gh needs a real GitHub token.

type s3PathIn struct {
	Path string `json:"path,omitempty"`
}

type s3GhIn struct {
	Subcommand string
	Args       string `json:"args,omitempty"`
}

func clip(s string, n int) string {
	s = strings.TrimSpace(s)
	if len(s) > n {
		return s[:n]
	}
	return s
}

func runCLI(timeout time.Duration, env []string, name string, args ...string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, name, args...)
	if env != nil {
		cmd.Env = env
	}
	var out, errOut strings.Builder
	cmd.Stdout, cmd.Stderr = &out, &errOut
	err := cmd.Run()
	return out.String(), errOut.String(), err
}

func TestCliCredentialLifecycle(t *testing.T) {
	skipInPlayback(t, "cli_mktemp returns a fresh temp path every run and cli_gh needs a real GitHub token")
	realToken := os.Getenv("GITHUB_TOKEN")
	if realToken == "" {
		t.Skip("GITHUB_TOKEN not set in environment — required for Suite 3 CLI tools test")
	}
	if _, err := exec.LookPath("gh"); err != nil {
		t.Skip("gh CLI not installed — required for Suite 3 CLI tools test")
	}
	requireRuntimeMetadata(t)
	rt := newRuntime(t)
	store := newSecretStore(t)
	const cred = "GITHUB_TOKEN"

	cliLs := tool.Func("cli_ls", func(_ context.Context, in s3PathIn) (string, error) {
		path := in.Path
		if path == "" {
			path = "."
		}
		out, errOut, err := runCLI(15*time.Second, nil, "ls", path)
		if err != nil {
			return "ls_error:" + clip(errOut, 200), nil
		}
		return "ls_ok:" + clip(out, 200), nil
	},
		"List directory contents using the ls command.")
	cliMktemp := tool.Func("cli_mktemp", func(context.Context, struct{}) (string, error) {
		out, errOut, err := runCLI(15*time.Second, nil, "mktemp")
		if err != nil {
			return "mktemp_error:" + clip(errOut, 200), nil
		}
		return "mktemp_ok:" + strings.TrimSpace(out), nil
	},
		"Create a temporary file and return its path.")
	cliGh := tool.Func("cli_gh", func(ctx context.Context, in s3GhIn) (string, error) {
		// gh reads the token from its environment; it comes from the
		// task, never from this process.
		env, err := ai.SecretsEnv(ctx, cred)
		if err != nil {
			return "", taskmodel.NewNonRetryableError(err)
		}
		args := strings.Fields(in.Subcommand)
		args = append(args, strings.Fields(in.Args)...)
		out, errOut, err := runCLI(30*time.Second, append(os.Environ(), env...), "gh", args...)
		if err != nil {
			return "gh_error:" + clip(errOut, 200), nil
		}
		return "gh_ok:" + clip(out, 200), nil
	},
		"Run a gh CLI command. Requires GITHUB_TOKEN credential.\nExample: subcommand=\"repo list\", args=\"--limit 3\"", tool.WithCredentials(cred))

	agent := &ai.Agent{Name: "e2e_cli_tools", Model: model(t),
		Instructions: "You have three tools: cli_ls, cli_mktemp, and cli_gh.\n" +
			"You MUST call each tool exactly once as directed and report the output verbatim.\n" +
			"Do not skip any tool. Do not add commentary beyond the results.\n",
		Tools: ai.Tools(cliLs, cliMktemp, cliGh)}
	const promptAllThree = "Call all three tools:\n" +
		"1. cli_ls with path=\"/tmp\"\n" +
		"2. cli_mktemp (no arguments)\n" +
		"3. cli_gh with subcommand=\"repo list\" and args=\"--limit 3\"\n" +
		"Report each result in this format:\n" +
		"  cli_ls: <output>\n  cli_mktemp: <output>\n  cli_gh: <output>\n"

	t.Cleanup(func() { store.delete(cred) })
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	// Step 1: no token in the store. Step 2: a real token in this process only.
	store.delete(cred)
	os.Setenv(cred, realToken)

	// Step 3: ls and mktemp work; gh must not, since its credential was never delivered.
	res := runTolerant(t, rt, ctx, agent, promptAllThree)
	assertTerminal(t, res, "Step 3: Token in env only")
	if !strings.Contains(res.Output, "ls_ok") || !strings.Contains(res.Output, "mktemp_ok") {
		t.Errorf("[Step 3] ls and mktemp should have run: %.400s", res.Output)
	}
	if strings.Contains(res.Output, "gh_ok") {
		t.Errorf("[Step 3] SECURITY: cli_gh ran with a token that was only in the environment: %.400s", res.Output)
	}

	// Step 4: the CLI tool's allow-list, in the plan and at run time.
	whitelist := &ai.Agent{Name: "e2e_cli_whitelist", Model: model(t),
		Instructions: "You have a run_command tool that executes CLI commands. Always call the tool as instructed and report the exact output.",
		CLI:          &ai.CLIConfig{AllowedCommands: []string{"ls", "mktemp", "gh"}, TimeoutSeconds: 30}}
	ad := agentDef(t, planAgent(t, rt, whitelist))
	var runCommand map[string]any
	for _, tl := range asList(ad["tools"]) {
		if m, ok := tl.(map[string]any); ok && strings.Contains(fmt.Sprint(m["name"]), "run_command") {
			runCommand = m
		}
	}
	if runCommand == nil {
		t.Fatalf("[Step 4a] no run_command tool in the plan: %v", ad["tools"])
	}
	m := regexp.MustCompile(`Allowed commands:\s*(.+?)\.`).FindStringSubmatch(fmt.Sprint(runCommand["description"]))
	if m == nil {
		t.Fatalf("[Step 4a] run_command description lacks the allow-list: %v", runCommand["description"])
	}
	var allowed []string
	for _, c := range strings.Split(m[1], ",") {
		allowed = append(allowed, strings.TrimSpace(c))
	}
	sort.Strings(allowed)
	if strings.Join(allowed, ",") != "gh,ls,mktemp" {
		t.Errorf("[Step 4a] allow-list in the description = %v, want gh, ls, mktemp", allowed)
	}
	// Step 4b, the validator itself, is a unit test of the SDK: see cli_runner_test.go.
	resCd := runTolerant(t, rt, ctx, whitelist,
		"You MUST call the run_command tool with command=\"cd\" and args=[\"/etc\"].\n"+
			"Report the exact output or error message verbatim.\n")
	assertTerminal(t, resCd, "Step 4c: Disallowed command")

	// Step 5: store the token. A read-only store ends the test here as a skip.
	putSecretOrSkip(t, store, cred, realToken)

	// Step 6: with the credential stored, gh runs too.
	res = runTolerant(t, rt, ctx, agent, promptAllThree)
	assertRunCompleted(t, res, "Step 5: With credential")
	for _, want := range []string{"ls_ok", "mktemp_ok", "gh_ok"} {
		if !strings.Contains(res.Output, want) {
			t.Errorf("[Step 6] answer lacks %s: %.400s", want, res.Output)
		}
	}
}
