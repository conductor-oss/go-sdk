//go:build ignore

// Handoff to Parallel — delegate to a multi-agent group.
//
// Run with:  go run agents/66_handoff_to_parallel.go
//
// A parent agent hands off either to a single agent (for quick checks) or to
// a parallel multi-agent group (for deep analysis). The parallel sub-agent
// runs its own fan-out/fan-in internally.
//
// Architecture:
//
//	coordinator (HANDOFF)
//	├── quick_check           (single agent, fast)
//	└── deep_analysis         (PARALLEL group)
//	    ├── market_analyst
//	    └── risk_analyst
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
	"strings"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	// Quick check: a single agent.
	quickCheck := &ai.Agent{
		Name:         "quick_check",
		Model:        model,
		Instructions: "You provide quick, 1-sentence assessments. Be brief and direct.",
	}

	// Deep analysis: a parallel group.
	marketAnalyst := &ai.Agent{
		Name:  "market_analyst_66",
		Model: model,
		Instructions: "You are a market analyst. Analyze the market opportunity: " +
			"size, growth rate, key players. 3-4 bullet points.",
	}
	riskAnalyst := &ai.Agent{
		Name:  "risk_analyst_66",
		Model: model,
		Instructions: "You are a risk analyst. Identify the top 3 risks: " +
			"regulatory, technical, and competitive. 3-4 bullet points.",
	}
	deepAnalysis := &ai.Agent{
		Name:     "deep_analysis",
		Model:    model,
		Agents:   []*ai.Agent{marketAnalyst, riskAnalyst},
		Strategy: ai.StrategyParallel,
	}

	// Coordinator with handoff.
	coordinator := &ai.Agent{
		Name:  "coordinator_66",
		Model: model,
		Instructions: "You are a business strategist. Route requests to the right team:\n" +
			"- quick_check for simple yes/no questions or quick assessments\n" +
			"- deep_analysis for comprehensive analysis requiring multiple perspectives",
		Agents:   []*ai.Agent{quickCheck, deepAnalysis},
		Strategy: ai.StrategyHandoff,
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()
	ctx := context.Background()

	// Scenario 1: deep analysis (handoff to the parallel group).
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  Scenario 1: Deep analysis (handoff → parallel group)")
	fmt.Println(strings.Repeat("=", 60))
	result, err := runtime.Run(ctx, coordinator,
		"Provide a deep analysis of entering the AI healthcare market.")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	printResult(result)
	if result.Status == ai.StatusCompleted {
		fmt.Println("[OK] Handoff to parallel group completed successfully")
	} else {
		fmt.Printf("[WARN] Unexpected status: %s\n", result.Status)
	}

	// Scenario 2: quick check (handoff to a single agent).
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("  Scenario 2: Quick check (handoff → single agent)")
	fmt.Println(strings.Repeat("=", 60))
	result2, err := runtime.Run(ctx, coordinator, "Is the mobile app market still growing?")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	printResult(result2)
	if result2.Status == ai.StatusCompleted {
		fmt.Println("[OK] Quick check completed successfully")
	} else {
		fmt.Printf("[WARN] Unexpected status: %s\n", result2.Status)
	}

	// Production pattern:
	// 1. Deploy once during CI/CD: runtime.Deploy(ctx, coordinator)
	// 2. In a separate long-lived worker process: runtime.Serve(ctx, coordinator)
}

// printResult mirrors the Python AgentResult.print_result helper.
func printResult(r *ai.AgentResult) {
	const width = 50
	line := ""
	for i := 0; i < width; i++ {
		line += "═"
	}
	fmt.Printf("\n╒%s╕\n", line)
	fmt.Printf("│ %-*s│\n", width-1, "Agent Output")
	fmt.Printf("╘%s╛\n\n", line)

	if r.Status == ai.StatusFailed && r.Error != "" {
		fmt.Println("ERROR:", r.Error)
	} else {
		fmt.Println(r.Output)
	}
	fmt.Println()
	fmt.Println("Status:", r.Status)
	fmt.Println("Execution ID:", r.ExecutionID)
}
