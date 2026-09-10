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
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// Attaching CodeExecutionConfig gives the agent a derived tool named
// "<agent>_execute_code". The config serializes correctly — the golden fixture
// proves that — but this asks the sharper question: does a run using it work?
func TestCodeExecution(t *testing.T) {
	rt := newRuntime(t)
	defer rt.Shutdown()

	agent := &ai.Agent{
		Name:  "go_e2e_executor",
		Model: model(t),
		// The tool returns only what the program writes to stdout: it runs the
		// snippet as a script, not in a REPL, so a bare trailing expression
		// produces nothing. Without this instruction the model writes
		// REPL-style code, sees an empty result and retries until the run
		// times out. Python's executor has the same shape, so the same
		// instruction is needed there.
		Instructions: "You must run Python with the execute_code tool. Use it rather than " +
			"working answers out yourself, and always print results with print() — " +
			"the tool returns only what the program writes to stdout.",
		CodeExecution: &ai.CodeExecutionConfig{
			AllowedLanguages: []string{"python"},
			TimeoutSeconds:   30,
		},
	}

	// Short on purpose: if nothing serves the derived tool, the run stalls and
	// there is no reason to wait four minutes to learn that.
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	res, err := rt.Run(ctx, agent, "Compute the 20th Fibonacci number using Python. Report the number.")
	if err != nil {
		t.Fatalf("Run: %v", err)
	}

	t.Logf("status=%s output=%q", res.Status, res.Output)
	if res.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q", res.Status, ai.StatusCompleted)
	}
	// fib(20) = 6765.
	if !strings.Contains(res.Output, "6765") {
		t.Errorf("output does not contain the answer 6765: %q", res.Output)
	}
}
