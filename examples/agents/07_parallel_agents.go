package agents

import (
	"context"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// ParallelAgents is the Python SDK's examples/agents/07_parallel_agents.py:
// three analysts examine the same topic at once.
func ParallelAgents(model string) *ai.Agent {
	marketAnalyst := &ai.Agent{
		Name:         "market_analyst",
		Model:        model,
		Instructions: "You are a market analyst. Analyze the given topic from a market perspective: market size, growth trends, key players, and opportunities.",
	}
	riskAnalyst := &ai.Agent{
		Name:         "risk_analyst",
		Model:        model,
		Instructions: "You are a risk analyst. Analyze the given topic for risks: regulatory risks, technical risks, competitive threats, and mitigation strategies.",
	}
	complianceChecker := &ai.Agent{
		Name:         "compliance",
		Model:        model,
		Instructions: "You are a compliance specialist. Check the given topic for compliance considerations: data privacy, regulatory requirements, and industry standards.",
	}
	return &ai.Agent{
		Name:     "analysis",
		Model:    model,
		Agents:   []*ai.Agent{marketAnalyst, riskAnalyst, complianceChecker},
		Strategy: ai.StrategyParallel,
	}
}

func runParallelAgents(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, ParallelAgents(Model()), "Launching an AI-powered healthcare diagnostic tool in the US market", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
