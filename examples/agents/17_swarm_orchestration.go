//go:build ignore

// Swarm Orchestration — a front-line agent transfers the conversation.
//
// Run with:  go run agents/17_swarm_orchestration.go
//
// The support agent triages the request and hands off to the refund
// specialist or tech support through transfer tools; a specialist can hand
// back. Text-mention handoffs route on keywords as well.
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

	refundAgent := &ai.Agent{
		Name:  "refund_specialist",
		Model: model,
		Instructions: "You are a refund specialist. Process the customer's refund request. " +
			"Check eligibility, confirm the refund amount, and let them know the " +
			"timeline. Be empathetic and clear. Do NOT ask follow-up questions — " +
			"just process the refund based on what the customer told you.",
	}
	techAgent := &ai.Agent{
		Name:  "tech_support",
		Model: model,
		Instructions: "You are a technical support specialist. Diagnose the customer's " +
			"technical issue and provide clear troubleshooting steps.",
	}
	support := &ai.Agent{
		Name:  "support",
		Model: model,
		Instructions: "You are the front-line customer support agent. Triage customer requests. " +
			"If the customer needs a refund, transfer to the refund specialist. " +
			"If they have a technical issue, transfer to tech support. " +
			"Use the transfer tools available to you to hand off the conversation.",
		Agents:   []*ai.Agent{refundAgent, techAgent},
		Strategy: ai.StrategySwarm,
		Handoffs: []ai.HandoffCondition{
			&ai.OnTextMention{Text: "refund", Target: "refund_specialist"},
			&ai.OnTextMention{Text: "technical", Target: "tech_support"},
		},
		MaxTurns: 3,
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	fmt.Println("--- Refund scenario ---")
	result, err := runtime.Run(context.Background(), support,
		"I bought a product last week and it arrived damaged. I want my money back.")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	printResult(result)
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
