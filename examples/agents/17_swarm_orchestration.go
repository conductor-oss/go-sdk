package agents

import (
	"context"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// SwarmOrchestration is the Python SDK's examples/agents/17_swarm_orchestration.py:
// front-line support transfers the conversation to a specialist.
func SwarmOrchestration(model string) *ai.Agent {
	refundAgent := &ai.Agent{
		Name:         "refund_specialist",
		Model:        model,
		Instructions: "You are a refund specialist. Process the customer's refund request. Check eligibility, confirm the refund amount, and let them know the timeline. Be empathetic and clear. Do NOT ask follow-up questions — just process the refund based on what the customer told you.",
	}
	techAgent := &ai.Agent{
		Name:         "tech_support",
		Model:        model,
		Instructions: "You are a technical support specialist. Diagnose the customer's technical issue and provide clear troubleshooting steps.",
	}
	return &ai.Agent{
		Name:         "support",
		Model:        model,
		Instructions: "You are the front-line customer support agent. Triage customer requests. If the customer needs a refund, transfer to the refund specialist. If they have a technical issue, transfer to tech support. Use the transfer tools available to you to hand off the conversation.",
		Agents:       []*ai.Agent{refundAgent, techAgent},
		Strategy:     ai.StrategySwarm,
		Handoffs: []ai.HandoffCondition{
			&ai.OnTextMention{Text: "refund", Target: "refund_specialist"},
			&ai.OnTextMention{Text: "technical", Target: "tech_support"},
		},
		MaxTurns: 3,
	}
}

func runSwarmOrchestration(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, SwarmOrchestration(Model()),
		"I bought a product last week and it arrived damaged. I want my money back.", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
