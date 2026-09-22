//go:build ignore

// Handoffs — agent delegating to sub-agents.
//
// Run with:  go run agents/05_handoffs.go
//
// Demonstrates the handoff strategy where the parent agent's LLM decides
// which sub-agent to delegate to. Sub-agents appear as callable tools.
//
// Requirements:
//   - Conductor server with LLM support
//   - CONDUCTOR_SERVER_URL=http://localhost:8080/api in the environment
//   - CONDUCTOR_AGENT_LLM_MODEL in the environment (optional)
package main

import (
	"context"
	"fmt"
	"os"

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

func checkBalance(ctx context.Context, in accountIn) (map[string]any, error) {
	return map[string]any{"account_id": in.AccountID, "balance": 5432.10, "currency": "USD"}, nil
}

func lookupOrder(ctx context.Context, in orderIn) (map[string]any, error) {
	return map[string]any{"order_id": in.OrderID, "status": "shipped", "eta": "2 days"}, nil
}

func getPricing(ctx context.Context, in productIn) (map[string]any, error) {
	return map[string]any{"product": in.Product, "price": 99.99, "discount": "10% off"}, nil
}

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	billing := &ai.Agent{
		Name:         "billing",
		Model:        model,
		Instructions: "You handle billing questions: balances, payments, invoices.",
		Tools:        ai.Tools(tool.Func("check_balance", "Check the balance of a bank account.", checkBalance)),
	}
	technical := &ai.Agent{
		Name:         "technical",
		Model:        model,
		Instructions: "You handle technical questions: order status, shipping, returns.",
		Tools:        ai.Tools(tool.Func("lookup_order", "Look up the status of an order.", lookupOrder)),
	}
	sales := &ai.Agent{
		Name:         "sales",
		Model:        model,
		Instructions: "You handle sales questions: pricing, products, promotions.",
		Tools:        ai.Tools(tool.Func("get_pricing", "Get pricing information for a product.", getPricing)),
	}
	support := &ai.Agent{
		Name:         "support",
		Model:        model,
		Instructions: "Route customer requests to the right specialist: billing, technical, or sales.",
		Agents:       []*ai.Agent{billing, technical, sales},
		Strategy:     ai.StrategyHandoff,
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	result, err := runtime.Run(context.Background(), support, "What's the balance on account ACC-123?")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()
}
