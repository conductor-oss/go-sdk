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
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type orderStatusIn struct {
	OrderID string `json:"order_id"`
}

type customerInfoIn struct {
	CustomerID string `json:"customer_id"`
}

// Guardrails — the Python SDK's examples/agents/10_guardrails.py as a test.
//
// A support agent's tools return a card number; a custom output guardrail
// rejects any answer that repeats it and makes the model revise. This is that
// flow, copied, with the print replaced by validation: the run completes with
// the recorded, redacted answer; the guardrail worker rejected the draft that
// carried the card number and passed the revision; and the card number does
// not appear in the output. The rejection message is the Python guardrail's,
// verbatim, because the server quotes it to the model in the retry prompt.
// (The guardrail also runs on the tool-calling turn, whose empty text
// passes, so the verdicts read pass, fail, pass.)
func TestExample10Guardrails(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "10_guardrails")

	getOrderStatus := func(ctx context.Context, in orderStatusIn) (map[string]any, error) {
		return map[string]any{
			"order_id":           in.OrderID,
			"status":             "shipped",
			"tracking":           "1Z999AA10123456784",
			"estimated_delivery": "2026-02-22",
		}, nil
	}
	getCustomerInfo := func(ctx context.Context, in customerInfoIn) (map[string]any, error) {
		return map[string]any{
			"customer_id":  in.CustomerID,
			"name":         "Alice Johnson",
			"email":        "alice@example.com",
			"card_on_file": "4532-0150-1234-5678",
			"membership":   "gold",
		}, nil
	}

	ccPattern := regexp.MustCompile(`\b\d{4}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}\b`)
	ssnPattern := regexp.MustCompile(`\b\d{3}-\d{2}-\d{4}\b`)
	var mu sync.Mutex
	var verdicts []bool
	noPII := func(ctx context.Context, in ai.GuardrailInput) (ai.GuardrailResult, error) {
		passed := !ccPattern.MatchString(in.Content) && !ssnPattern.MatchString(in.Content)
		mu.Lock()
		verdicts = append(verdicts, passed)
		mu.Unlock()
		if !passed {
			return ai.GuardrailResult{
				Passed: false,
				Message: "Your response contains PII (credit card or SSN). " +
					"Redact all card numbers and SSNs before responding.",
			}, nil
		}
		return ai.GuardrailResult{Passed: true}, nil
	}
	guard := ai.NewCustomGuardrail("no_pii", noPII)
	guard.Position = ai.PositionOutput
	guard.OnFail = ai.OnFailRetry

	agent := &ai.Agent{
		Name:  "support_agent",
		Model: mockModel,
		Tools: []ai.ToolDef{
			tool.Func("get_order_status", "Look up the current status of an order.", getOrderStatus),
			tool.Func("get_customer_info", "Retrieve customer details including payment info on file.", getCustomerInfo),
		},
		Instructions: "You are a customer support assistant. Use the available tools to " +
			"answer questions about orders and customers. Always include all " +
			"details from the tool results in your response.",
		Guardrails: []ai.Guardrail{guard},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	result, err := runtime.Run(ctx, agent,
		"I need a full summary: What's the status of order ORD-42, "+
			"and what's the profile for customer CUST-7?")
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}

	// Validation, in place of the example's print and PII check.
	if result.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	if strings.Contains(result.Output, "4532-0150-1234-5678") {
		t.Error("PII leaked through the guardrail")
	}
	mu.Lock()
	v := append([]bool(nil), verdicts...)
	mu.Unlock()
	rejected := false
	for _, ok := range v {
		if !ok {
			rejected = true
		}
	}
	if !rejected || len(v) == 0 || !v[len(v)-1] {
		t.Errorf("guardrail verdicts = %v, want a rejected draft and a passing final revision", v)
	}
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[len(recorded)-1]); got != want {
		t.Errorf("output is not the recorded answer\n--- got ---\n%.400s\n--- recorded ---\n%.400s", got, want)
	}
}
