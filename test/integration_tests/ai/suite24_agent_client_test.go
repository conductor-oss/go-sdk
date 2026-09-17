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
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// The control-plane run tests of the Python SDK's
// e2e/test_suite24_agent_client.py, under the same names: an LLM-only agent
// (no local tools, so no workers) reaches COMPLETED through Run, and Start
// returns a handle that resolves to a COMPLETED result. The schedule tests in
// that suite belong to step 5 and are not ported here.

func TestRunLLMOnlyAgentCompletes(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name:         "e2e_client_run",
		Model:        model(t),
		Instructions: "You are a calculator. Reply with only the number.",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	res, err := rt.Run(ctx, agent, "What is 2 + 2? Reply with only the number.")
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if res.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", res.Status, ai.StatusCompleted, res.Error)
	}
	if res.ExecutionID == "" {
		t.Error("no executionId returned")
	}
	if !strings.Contains(res.Output, "4") {
		t.Errorf("answer omits 4: %q", res.Output)
	}
}

func TestStartReturnsHandleThenJoins(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name:         "e2e_client_start",
		Model:        model(t),
		Instructions: "Reply with the single word: ok",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	handle, err := rt.Start(ctx, agent, "Say ok")
	if err != nil {
		t.Fatalf("Start: %v", err)
	}
	if handle.ExecutionID == "" {
		t.Fatal("handle has no executionId")
	}
	res, err := handle.Result(ctx)
	if err != nil {
		t.Fatalf("Result: %v", err)
	}
	if res.Status != ai.StatusCompleted {
		t.Errorf("status = %q, want %q", res.Status, ai.StatusCompleted)
	}
}
