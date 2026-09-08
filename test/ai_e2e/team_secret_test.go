//go:build e2e

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai_e2e

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// The value the server must have in its store for the credential assertions to
// mean anything. conductor-oss reads secrets from its own process environment
// and rejects writes, so it has to be provisioned at server start:
//
//	CONDUCTOR_SECRET_GH_TOKEN=... java -jar conductor-server.jar
const (
	credentialName  = "GH_TOKEN"
	expectedSecret  = "ghp_fake_e2e_token_value_12345"
	secretEnvForSrv = "CONDUCTOR_SECRET_GH_TOKEN"
)

type prIn struct {
	Title string `json:"title"`
}

type prResult struct {
	URL   string `json:"url,omitempty"`
	Error string `json:"error,omitempty"`
}

// A team of sub-agents where one tool needs a secret.
//
// Two things are under test and they are independent: that a sequential team
// runs its sub-agents in order, and that a declared credential is delivered to
// the tool that declared it. The credential half is the sharper one — it is the
// only path in the SDK where a value the code never sees has to arrive intact.
func TestTeamWithSecret(t *testing.T) {
	rt := newRuntime(t)

	var toolRan atomic.Int32
	var gotSecret atomic.Value
	gotSecret.Store("")
	var secretErr atomic.Value
	secretErr.Store("")

	openPR := func(ctx context.Context, in prIn) (prResult, error) {
		toolRan.Add(1)
		token, err := ai.Secret(ctx, credentialName)
		if err != nil {
			// Report to the model rather than failing the task, so the run
			// finishes and the test can assert on what went wrong.
			secretErr.Store(err.Error())
			if !errors.Is(err, ai.ErrCredentialNotFound) {
				t.Errorf("want ErrCredentialNotFound, got %v", err)
			}
			return prResult{Error: err.Error()}, nil
		}
		gotSecret.Store(token)
		return prResult{URL: "https://github.com/example/repo/pull/42"}, nil
	}

	publisher := &ai.Agent{
		Name:         "publisher",
		Model:        model(t),
		Instructions: "Open a pull request for the reviewed change using the open_pr tool.",
		Tools: []ai.ToolDef{
			tool.Func("open_pr", "Open a pull request", openPR,
				tool.WithCredentials(credentialName)),
		},
	}

	team := &ai.Agent{
		Name:     "go_e2e_release_team",
		Model:    model(t),
		Strategy: ai.StrategySequential,
		Agents: []*ai.Agent{
			{
				Name:         "reviewer",
				Model:        model(t),
				Instructions: "Review the described change and summarise the risk in one sentence.",
			},
			publisher,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Minute)
	defer cancel()

	res, err := rt.Run(ctx, team, "The change adds a retry to the payments client. Review it, then open a pull request titled 'Add retry to payments client'.")
	if err != nil {
		t.Fatalf("Run: %v", err)
	}

	if res.Status != ai.StatusCompleted {
		t.Errorf("status = %q, want %q", res.Status, ai.StatusCompleted)
	}
	if toolRan.Load() == 0 {
		t.Fatal("open_pr never ran; the publisher sub-agent did not reach its tool")
	}

	// The credential assertions. A skip here rather than a failure when the
	// server has no secret provisioned: the SDK behaved correctly, the
	// environment simply cannot exercise this half.
	if msg, _ := secretErr.Load().(string); msg != "" {
		if os.Getenv("CONDUCTOR_E2E_SECRET_PROVISIONED") == "" {
			t.Skipf("credential not delivered, and %s was not declared as provisioned. "+
				"Start the server with %s=%s to exercise this. Tool reported: %s",
				"CONDUCTOR_E2E_SECRET_PROVISIONED", secretEnvForSrv, expectedSecret, msg)
		}
		t.Fatalf("credential %q was declared but not delivered: %s", credentialName, msg)
	}

	got, _ := gotSecret.Load().(string)
	if got == "" {
		t.Fatal("tool ran and reported no error, but read an empty credential")
	}
	if got != expectedSecret {
		t.Errorf("credential value mismatch: the server delivered something other than "+
			"the provisioned test value (got %d chars)", len(got))
	}

	t.Logf("execution %s: team ran, open_pr received a %d-character credential",
		res.ExecutionID, len(got))
	if !strings.Contains(strings.ToLower(res.Output), "pull request") &&
		!strings.Contains(res.Output, "42") {
		t.Logf("note: final output does not mention the PR: %q", res.Output)
	}
}

// A tool that declares nothing must not receive credentials, and Secret must
// say so clearly rather than returning an empty string.
func TestSecretRequiresDeclaration(t *testing.T) {
	rt := newRuntime(t)

	var errMsg atomic.Value
	errMsg.Store("")

	peek := func(ctx context.Context, in prIn) (prResult, error) {
		_, err := ai.Secret(ctx, credentialName)
		if err != nil {
			errMsg.Store(err.Error())
			return prResult{Error: "no credential"}, nil
		}
		return prResult{URL: "leaked"}, nil
	}

	agent := &ai.Agent{
		Name:         "go_e2e_undeclared",
		Model:        model(t),
		Instructions: "Call the peek tool once with any title, then stop.",
		Tools: []ai.ToolDef{
			// Deliberately no WithCredentials.
			tool.Func("peek", "Peek at configuration", peek),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	if _, err := rt.Run(ctx, agent, "Call peek with the title 'test'."); err != nil {
		t.Fatalf("Run: %v", err)
	}

	msg, _ := errMsg.Load().(string)
	if msg == "" {
		t.Fatal("an undeclared credential was delivered to the tool, or the tool never ran")
	}
	if !strings.Contains(msg, credentialName) {
		t.Errorf("the error should name the missing credential, got: %s", msg)
	}
	t.Logf("undeclared credential correctly withheld: %s", msg)
}

// The shell-out form from the Secrets section: credentials handed to a child
// process via cmd.Env rather than the parent's environment.
func TestSecretsEnvForSubprocess(t *testing.T) {
	rt := newRuntime(t)

	var envLine atomic.Value
	envLine.Store("")
	var callErr atomic.Value
	callErr.Store("")

	// Instead of shelling out to gh, run `sh -c 'echo $GH_TOKEN'` so the test
	// proves the value actually reached a child process's environment.
	echoToken := func(ctx context.Context, in prIn) (prResult, error) {
		env, err := ai.SecretsEnv(ctx, credentialName)
		if err != nil {
			callErr.Store(err.Error())
			return prResult{Error: err.Error()}, nil
		}
		cmd := exec.CommandContext(ctx, "sh", "-c", "printf %s \"$"+credentialName+"\"")
		cmd.Env = append(os.Environ(), env...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			callErr.Store(err.Error())
			return prResult{Error: err.Error()}, nil
		}
		envLine.Store(string(out))
		return prResult{URL: "ok"}, nil
	}

	agent := &ai.Agent{
		Name:         "go_e2e_shellout",
		Model:        model(t),
		Instructions: "Call the echo_token tool once with the title 'test', then stop.",
		Tools: []ai.ToolDef{
			tool.Func("echo_token", "Echo a configured token", echoToken,
				tool.WithCredentials(credentialName)),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	if _, err := rt.Run(ctx, agent, "Call echo_token with the title 'test'."); err != nil {
		t.Fatalf("Run: %v", err)
	}

	if msg, _ := callErr.Load().(string); msg != "" {
		if os.Getenv("CONDUCTOR_E2E_SECRET_PROVISIONED") == "" {
			t.Skipf("credential not provisioned on the server: %s", msg)
		}
		t.Fatalf("SecretsEnv failed: %s", msg)
	}

	got, _ := envLine.Load().(string)
	if got == "" {
		t.Fatal("the child process saw an empty value; SecretsEnv did not reach cmd.Env")
	}
	if got != expectedSecret {
		t.Errorf("child process saw a different value than provisioned (%d chars)", len(got))
	}

	// The parent must be untouched: that is the whole point of cmd.Env.
	if os.Getenv(credentialName) != "" {
		t.Errorf("%s leaked into the parent process environment", credentialName)
	}
	t.Logf("child process received the credential (%d chars); parent environment clean", len(got))
}
