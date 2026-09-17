//go:build ignore

// Basic Agent — the smallest possible agent.
//
// Run with:  go run agents/01_basic_agent.go
// (Each example here is a standalone file; the build tag keeps them out of
// the module build so several can sit side by side in this directory.)
//
// Demonstrates the simplest agent: define an agent, call Run, and print the
// result.
//
// Requirements:
//   - Conductor server with LLM support
//   - CONDUCTOR_SERVER_URL=http://localhost:8080/api in the environment
//   - CONDUCTOR_AGENT_LLM_MODEL in the environment (optional; defaults to
//     openai/gpt-4o, as the Python example's settings do)
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

	agent := &ai.Agent{
		Name:         "greeter",
		Model:        model,
		Instructions: "You are a friendly assistant. Keep responses brief.",
	}

	prompt := "Say hello and tell me a fun fact about Go."

	// NewRuntime connects to the server named by CONDUCTOR_SERVER_URL.
	//
	// Open-source Conductor has no authentication, so that one variable is
	// all it needs.
	//
	// Orkes Conductor requires authentication. In the Orkes UI, create an
	// application and generate an access key for it; the key ID goes in
	// CONDUCTOR_AUTH_KEY and the key secret in CONDUCTOR_AUTH_SECRET. The
	// client then fetches a token and sends it with every request.
	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	result, err := runtime.Run(context.Background(), agent, prompt)
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	printResult(result)

	// Production pattern:
	// 1. Deploy once during CI/CD, from a release script:
	//    runtime.Deploy(ctx, agent)
	//
	// 2. In a separate long-lived worker process:
	//    runtime.Serve(ctx, agent)
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
