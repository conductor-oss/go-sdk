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
	"sync/atomic"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// A custom guardrail is a Go function the server calls as a worker after
// each final reply, under a SIMPLE task named after the guardrail. The
// fixture proves it serializes; these ask whether the server dispatches to
// it, hands it the reply, and acts on the verdict.
//
// Both verdicts are fixed rather than judged, so the outcome does not depend
// on what the model writes: a guardrail that always passes must leave the
// run COMPLETED, and one that always raises must fail it with its message.

func TestCustomGuardrailPasses(t *testing.T) {
	rt := newRuntime(t)
	defer rt.Shutdown()

	var calls atomic.Int32
	var sawContent atomic.Bool
	agent := &ai.Agent{
		Name:         "go_e2e_guarded",
		Model:        model(t),
		Instructions: "Answer in one short sentence.",
		Guardrails: []ai.Guardrail{
			ai.NewCustomGuardrail("go_e2e_allow_all", func(_ context.Context, in ai.GuardrailInput) (ai.GuardrailResult, error) {
				calls.Add(1)
				if strings.TrimSpace(in.Content) != "" {
					sawContent.Store(true)
				}
				return ai.GuardrailResult{Passed: true}, nil
			}),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	res, err := rt.Run(ctx, agent, "What colour is the sky on a clear day?")
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	t.Logf("status=%s calls=%d sawContent=%v output=%.100q",
		res.Status, calls.Load(), sawContent.Load(), res.Output)

	if res.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q", res.Status, ai.StatusCompleted)
	}
	if calls.Load() == 0 {
		t.Fatal("the guardrail worker was never dispatched: the server did not route " +
			"to the task named after the guardrail")
	}
	if !sawContent.Load() {
		t.Error("the guardrail never received the reply text: content binding is wrong")
	}
}

func TestCustomGuardrailRaises(t *testing.T) {
	rt := newRuntime(t)
	defer rt.Shutdown()

	const reason = "GO-E2E-GUARDRAIL-REJECTED"
	var calls atomic.Int32
	deny := ai.NewCustomGuardrail("go_e2e_deny_all", func(context.Context, ai.GuardrailInput) (ai.GuardrailResult, error) {
		calls.Add(1)
		return ai.GuardrailResult{Passed: false, Message: reason}, nil
	})
	deny.OnFail = ai.OnFailRaise

	agent := &ai.Agent{
		Name:         "go_e2e_guarded_deny",
		Model:        model(t),
		Instructions: "Answer in one short sentence.",
		Guardrails:   []ai.Guardrail{deny},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	// A different question from TestCustomGuardrailPasses on purpose: two tests
	// that send the same request record two answers, and the LLM recorder
	// refuses to replay a request that has more than one.
	res, err := rt.Run(ctx, agent, "What colour is grass in spring?")
	if res == nil {
		t.Fatalf("Run returned no result: %v", err)
	}
	t.Logf("status=%s calls=%d err=%v", res.Status, calls.Load(), err)

	// A raise verdict ends the run FAILED; Run reports that as an error and
	// still returns the result, so both are checked.
	if err == nil || res.Status != ai.StatusFailed {
		t.Fatalf("status = %q, err = %v; want %q with an error", res.Status, err, ai.StatusFailed)
	}
	if calls.Load() == 0 {
		t.Fatal("the guardrail worker was never dispatched")
	}
	// The verdict's message is the run's failure reason.
	if !strings.Contains(res.Error, reason) {
		t.Errorf("failure reason %q does not carry the guardrail message %q", res.Error, reason)
	}
}
