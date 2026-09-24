package agents

import (
	"context"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// HandoffToParallel is the Python SDK's examples/agents/66_handoff_to_parallel.py:
// a coordinator hands off to a single agent or to a parallel group.
func HandoffToParallel(model string) *ai.Agent {
	quickCheck := &ai.Agent{
		Name:         "quick_check",
		Model:        model,
		Instructions: "You provide quick, 1-sentence assessments. Be brief and direct.",
	}
	marketAnalyst := &ai.Agent{
		Name:         "market_analyst_66",
		Model:        model,
		Instructions: "You are a market analyst. Analyze the market opportunity: size, growth rate, key players. 3-4 bullet points.",
	}
	riskAnalyst := &ai.Agent{
		Name:         "risk_analyst_66",
		Model:        model,
		Instructions: "You are a risk analyst. Identify the top 3 risks: regulatory, technical, and competitive. 3-4 bullet points.",
	}
	deepAnalysis := &ai.Agent{
		Name:     "deep_analysis",
		Model:    model,
		Agents:   []*ai.Agent{marketAnalyst, riskAnalyst},
		Strategy: ai.StrategyParallel,
	}
	return &ai.Agent{
		Name:         "coordinator_66",
		Model:        model,
		Instructions: "You are a business strategist. Route requests to the right team:\n- quick_check for simple yes/no questions or quick assessments\n- deep_analysis for comprehensive analysis requiring multiple perspectives",
		Agents:       []*ai.Agent{quickCheck, deepAnalysis},
		Strategy:     ai.StrategyHandoff,
	}
}

func runHandoffToParallel(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	coordinator := HandoffToParallel(Model())
	first, err := run(ctx, rt, coordinator, "Provide a deep analysis of entering the AI healthcare market.", out)
	if err != nil {
		return nil, err
	}
	second, err := run(ctx, rt, coordinator, "Is the mobile app market still growing?", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{first, second}, nil
}
