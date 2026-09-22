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

type accountIn struct {
	AccountID string
}

type orderIn struct {
	OrderID string
}

type productIn struct {
	Product string
}

// Handoffs — the Python SDK's examples/agents/05_handoffs.py as a test.
//
// A support agent routes a request to one of three specialists, each with a
// tool. This is that flow, copied, with the print replaced by validation: the
// run completes with the recorded answer, the billing specialist's tool ran
// once for the account asked about, and the other specialists' tools never
// ran — the routing happened, and to the right agent.
func TestExample05Handoffs(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "05_handoffs")

	var balanceCalls, orderCalls, pricingCalls atomic.Int32
	var account atomic.Value
	checkBalance := func(ctx context.Context, in accountIn) (map[string]any, error) {
		balanceCalls.Add(1)
		account.Store(in.AccountID)
		return map[string]any{"account_id": in.AccountID, "balance": 5432.10, "currency": "USD"}, nil
	}
	lookupOrder := func(ctx context.Context, in orderIn) (map[string]any, error) {
		orderCalls.Add(1)
		return map[string]any{"order_id": in.OrderID, "status": "shipped", "eta": "2 days"}, nil
	}
	getPricing := func(ctx context.Context, in productIn) (map[string]any, error) {
		pricingCalls.Add(1)
		return map[string]any{"product": in.Product, "price": 99.99, "discount": "10% off"}, nil
	}

	billing := &ai.Agent{
		Name:         "billing",
		Model:        mockModel,
		Instructions: "You handle billing questions: balances, payments, invoices.",
		Tools:        ai.Tools(tool.Func("check_balance", "Check the balance of a bank account.", checkBalance)),
	}
	technical := &ai.Agent{
		Name:         "technical",
		Model:        mockModel,
		Instructions: "You handle technical questions: order status, shipping, returns.",
		Tools:        ai.Tools(tool.Func("lookup_order", "Look up the status of an order.", lookupOrder)),
	}
	sales := &ai.Agent{
		Name:         "sales",
		Model:        mockModel,
		Instructions: "You handle sales questions: pricing, products, promotions.",
		Tools:        ai.Tools(tool.Func("get_pricing", "Get pricing information for a product.", getPricing)),
	}
	support := &ai.Agent{
		Name:         "support",
		Model:        mockModel,
		Instructions: "Route customer requests to the right specialist: billing, technical, or sales.",
		Agents:       []*ai.Agent{billing, technical, sales},
		Strategy:     ai.StrategyHandoff,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	result, err := runtime.Run(ctx, support, "What's the balance on account ACC-123?")
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}

	// Validation, in place of the example's result.print_result().
	if result.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	if n := balanceCalls.Load(); n != 1 {
		t.Errorf("check_balance ran %d times, want 1", n)
	}
	if acc, _ := account.Load().(string); acc != "ACC-123" {
		t.Errorf("check_balance was asked about %q, want ACC-123", acc)
	}
	if n := orderCalls.Load() + pricingCalls.Load(); n != 0 {
		t.Errorf("technical/sales tools ran %d times; the request should have gone to billing only", n)
	}
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[len(recorded)-1]); got != want {
		t.Errorf("output is not the recorded answer\n--- got ---\n%s\n--- recorded ---\n%s", got, want)
	}
}
