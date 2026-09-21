//go:build ignore

// Sequential Pipeline — agents run one after another.
//
// Run with:  go run agents/06_sequential_pipeline.go
//
// A researcher, a writer and an editor each receive the topic and the
// previous agent's output, in order. The pipeline is a sequential agent whose
// name joins its members', as the Python SDK's `researcher >> writer >> editor`
// builds it.
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

	researcher := &ai.Agent{
		Name:  "researcher",
		Model: model,
		Instructions: "You are a researcher. Given a topic, provide key facts and data points. " +
			"Be thorough but concise. Output raw research findings.",
	}
	writer := &ai.Agent{
		Name:  "writer",
		Model: model,
		Instructions: "You are a writer. Take research findings and write a clear, engaging " +
			"article. Use headers and bullet points where appropriate.",
	}
	editor := &ai.Agent{
		Name:  "editor",
		Model: model,
		Instructions: "You are an editor. Review the article for clarity, grammar, and tone. " +
			"Make improvements and output the final polished version.",
	}

	// researcher >> writer >> editor
	pipeline := &ai.Agent{
		Name:     "researcher_writer_editor",
		Model:    model,
		Agents:   []*ai.Agent{researcher, writer, editor},
		Strategy: ai.StrategySequential,
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	result, err := runtime.Run(context.Background(), pipeline, "The impact of AI agents on software development in 2025")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()
}
