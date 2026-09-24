package agents

import (
	"context"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// SwarmWithTools is the Python SDK's examples/agents/64_swarm_with_tools.py:
// swarm specialists that each carry their own domain tool.
func SwarmWithTools(model string) *ai.Agent {
	billingSpecialist := &ai.Agent{
		Name:         "billing_specialist",
		Model:        model,
		Instructions: "You are a billing specialist. Use the check_balance tool to look up account balances. Include the balance amount in your response.",
		Tools:        ai.Tools(tool.Func("check_balance", checkBalance, "Check the balance of a bank account.")),
	}
	orderSpecialist := &ai.Agent{
		Name:         "order_specialist",
		Model:        model,
		Instructions: "You are an order specialist. Use the lookup_order tool to check order status. Include the shipping status and ETA in your response.",
		Tools:        ai.Tools(tool.Func("lookup_order", lookupOrder, "Look up the status of an order.")),
	}
	return &ai.Agent{
		Name:         "support",
		Model:        model,
		Instructions: "You are front-line customer support. Triage customer requests. Transfer to billing_specialist for account/payment questions, order_specialist for shipping/order questions.",
		Agents:       []*ai.Agent{billingSpecialist, orderSpecialist},
		Strategy:     ai.StrategySwarm,
		Handoffs: []ai.HandoffCondition{
			&ai.OnTextMention{Text: "billing", Target: "billing_specialist"},
			&ai.OnTextMention{Text: "order", Target: "order_specialist"},
		},
		MaxTurns: 3,
	}
}

func runSwarmWithTools(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	support := SwarmWithTools(Model())
	first, err := run(ctx, rt, support, "What's the balance on account ACC-456?", out)
	if err != nil {
		return nil, err
	}
	second, err := run(ctx, rt, support, "Where is my order ORD-789?", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{first, second}, nil
}
