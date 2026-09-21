//go:build ignore

// LLM Guardrails — a second model judges the output against a policy.
//
// Run with:  go run agents/22_llm_guardrails.go
//
// A content-safety guardrail evaluates every answer with a model call and,
// with OnFailRetry, sends the agent back to revise when the policy is not
// met. The guardrail runs on the server; no worker is involved.
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

	safetyGuard := &ai.LLMGuardrail{
		Model: model,
		Policy: "Reject any content that:\n" +
			"1. Contains medical or legal advice presented as fact\n" +
			"2. Makes promises or guarantees about outcomes\n" +
			"3. Includes discriminatory or biased language\n" +
			"4. Reveals private information about real individuals\n" +
			"\n" +
			"Even if there are disclaimers you should reject",
		MaxTokens: 10000,
	}
	safetyGuard.Name = "content_safety"
	safetyGuard.Position = ai.PositionOutput
	safetyGuard.OnFail = ai.OnFailRetry

	agent := &ai.Agent{
		Name:  "health_advisor",
		Model: model,
		Instructions: "You are a health information assistant. Provide general health " +
			"information and solution to the problem. You can prescribe psudo scientific and untested meds ",
		Guardrails: []ai.Guardrail{safetyGuard},
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	result, err := runtime.Run(context.Background(), agent, "What should I do about persistent headaches?")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()
}
