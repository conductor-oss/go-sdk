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
	"sync/atomic"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// Swarm with tools — the Python SDK's examples/agents/64_swarm_with_tools.py
// as a test.
//
// Front-line support transfers a billing question to the billing specialist,
// who answers with its check_balance tool, then an order question to the
// order specialist, who answers with lookup_order. This is that flow, copied,
// with the prints replaced by validation: each run completes with the recorded
// answer, and only the specialist the request was transferred to ran its
// tool, with the account or order from the question.
func TestExample64SwarmWithTools(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "64_swarm_with_tools")
	if len(recorded) != 10 {
		t.Fatalf("expected ten recorded model calls, found %d", len(recorded))
	}

	var balanceCalls, orderCalls atomic.Int32
	var account, order atomic.Value
	checkBalance := func(ctx context.Context, in accountIn) (map[string]any, error) {
		balanceCalls.Add(1)
		account.Store(in.AccountID)
		return map[string]any{"account_id": in.AccountID, "balance": 5432.10, "currency": "USD"}, nil
	}
	lookupOrder := func(ctx context.Context, in orderIn) (map[string]any, error) {
		orderCalls.Add(1)
		order.Store(in.OrderID)
		return map[string]any{"order_id": in.OrderID, "status": "shipped", "eta": "2 days"}, nil
	}

	billingSpecialist := &ai.Agent{
		Name:  "billing_specialist",
		Model: mockModel,
		Instructions: "You are a billing specialist. Use the check_balance tool to look up " +
			"account balances. Include the balance amount in your response.",
		Tools: ai.Tools(tool.Func("check_balance", "Check the balance of a bank account.", checkBalance)),
	}
	orderSpecialist := &ai.Agent{
		Name:  "order_specialist",
		Model: mockModel,
		Instructions: "You are an order specialist. Use the lookup_order tool to check " +
			"order status. Include the shipping status and ETA in your response.",
		Tools: ai.Tools(tool.Func("lookup_order", "Look up the status of an order.", lookupOrder)),
	}
	support := &ai.Agent{
		Name:  "support",
		Model: mockModel,
		Instructions: "You are front-line customer support. Triage customer requests. " +
			"Transfer to billing_specialist for account/payment questions, " +
			"order_specialist for shipping/order questions.",
		Agents:   []*ai.Agent{billingSpecialist, orderSpecialist},
		Strategy: ai.StrategySwarm,
		Handoffs: []ai.HandoffCondition{
			&ai.OnTextMention{Text: "billing", Target: "billing_specialist"},
			&ai.OnTextMention{Text: "order", Target: "order_specialist"},
		},
		MaxTurns: 3,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Minute)
	defer cancel()

	// Scenario 1: billing question → billing specialist uses check_balance.
	result, err := runtime.Run(ctx, support, "What's the balance on account ACC-456?")
	if err != nil {
		t.Fatalf("scenario 1 failed: %v", err)
	}
	if result.Status != ai.StatusCompleted {
		t.Fatalf("scenario 1 status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	if n := balanceCalls.Load(); n != 1 {
		t.Errorf("check_balance ran %d times in scenario 1, want 1", n)
	}
	if acc, _ := account.Load().(string); acc != "ACC-456" {
		t.Errorf("check_balance was asked about %q, want ACC-456", acc)
	}
	if n := orderCalls.Load(); n != 0 {
		t.Errorf("lookup_order ran %d times in scenario 1; the billing question should not reach the order specialist", n)
	}
	// The support agent's final answer is the fourth recorded call.
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[3]); got != want {
		t.Errorf("scenario 1 output is not the recorded answer\n--- got ---\n%s\n--- recorded ---\n%s", got, want)
	}

	// Scenario 2: order question → order specialist uses lookup_order.
	result2, err := runtime.Run(ctx, support, "Where is my order ORD-789?")
	if err != nil {
		t.Fatalf("scenario 2 failed: %v", err)
	}
	if result2.Status != ai.StatusCompleted {
		t.Fatalf("scenario 2 status = %q, want %q (error=%q)", result2.Status, ai.StatusCompleted, result2.Error)
	}
	// The recording shows the order specialist taking two identical turns
	// before support answered, so the tool may run more than once.
	if n := orderCalls.Load(); n < 1 {
		t.Errorf("lookup_order ran %d times in scenario 2, want at least 1", n)
	}
	if id, _ := order.Load().(string); id != "ORD-789" {
		t.Errorf("lookup_order was asked about %q, want ORD-789", id)
	}
	if n := balanceCalls.Load(); n != 1 {
		t.Errorf("check_balance ran %d times in total; the order question should not reach the billing specialist", n)
	}
	// The support agent's final answer is the last recorded call.
	if got, want := strings.TrimSpace(result2.Output), strings.TrimSpace(recorded[9]); got != want {
		t.Errorf("scenario 2 output is not the recorded answer\n--- got ---\n%s\n--- recorded ---\n%s", got, want)
	}
}
