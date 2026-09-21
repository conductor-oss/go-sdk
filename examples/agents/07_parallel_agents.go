//go:build ignore

// Parallel Agents — several agents analyze the same input at once.
//
// Run with:  go run agents/07_parallel_agents.go
//
// A market analyst, a risk analyst and a compliance specialist each receive
// the topic; the parallel agent runs them together and combines their output.
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
)

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	marketAnalyst := &ai.Agent{
		Name:  "market_analyst",
		Model: model,
		Instructions: "You are a market analyst. Analyze the given topic from a market perspective: " +
			"market size, growth trends, key players, and opportunities.",
	}
	riskAnalyst := &ai.Agent{
		Name:  "risk_analyst",
		Model: model,
		Instructions: "You are a risk analyst. Analyze the given topic for risks: " +
			"regulatory risks, technical risks, competitive threats, and mitigation strategies.",
	}
	complianceChecker := &ai.Agent{
		Name:  "compliance",
		Model: model,
		Instructions: "You are a compliance specialist. Check the given topic for compliance considerations: " +
			"data privacy, regulatory requirements, and industry standards.",
	}

	analysis := &ai.Agent{
		Name:     "analysis",
		Model:    model,
		Agents:   []*ai.Agent{marketAnalyst, riskAnalyst, complianceChecker},
		Strategy: ai.StrategyParallel,
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	result, err := runtime.Run(context.Background(), analysis, "Launching an AI-powered healthcare diagnostic tool in the US market")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()
}
