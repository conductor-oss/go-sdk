package agents

import (
	"context"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

func checkBalance(_ context.Context, in accountIn) (map[string]any, error) {
	return map[string]any{"account_id": in.AccountID, "balance": 5432.10, "currency": "USD"}, nil
}

func lookupOrder(_ context.Context, in orderIn) (map[string]any, error) {
	return map[string]any{"order_id": in.OrderID, "status": "shipped", "eta": "2 days"}, nil
}

func getPricing(_ context.Context, in productIn) (map[string]any, error) {
	return map[string]any{"product": in.Product, "price": 99.99, "discount": "10% off"}, nil
}

// Handoffs is the Python SDK's examples/agents/05_handoffs.py: a support agent
// hands off to a billing, technical or sales specialist.
func Handoffs(model string) *ai.Agent {
	billing := &ai.Agent{
		Name:         "billing",
		Model:        model,
		Instructions: "You handle billing questions: balances, payments, invoices.",
		Tools:        ai.Tools(tool.Func("check_balance", checkBalance, "Check the balance of a bank account.")),
	}
	technical := &ai.Agent{
		Name:         "technical",
		Model:        model,
		Instructions: "You handle technical questions: order status, shipping, returns.",
		Tools:        ai.Tools(tool.Func("lookup_order", lookupOrder, "Look up the status of an order.")),
	}
	sales := &ai.Agent{
		Name:         "sales",
		Model:        model,
		Instructions: "You handle sales questions: pricing, products, promotions.",
		Tools:        ai.Tools(tool.Func("get_pricing", getPricing, "Get pricing information for a product.")),
	}
	return &ai.Agent{
		Name:         "support",
		Model:        model,
		Instructions: "Route customer requests to the right specialist: billing, technical, or sales.",
		Agents:       []*ai.Agent{billing, technical, sales},
		Strategy:     ai.StrategyHandoff,
	}
}

func runHandoffs(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, Handoffs(Model()), "What's the balance on account ACC-123?", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
