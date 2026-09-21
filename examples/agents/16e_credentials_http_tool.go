//go:build ignore

// Credentials — HTTP tool with server-side credential resolution.
//
// Run with:  go run agents/16e_credentials_http_tool.go
//
// Demonstrates:
//   - tool.HTTP with tool.WithCredentials("GITHUB_TOKEN")
//   - ${GITHUB_TOKEN} in headers resolved server-side (not in Go)
//   - No worker process needed — Conductor makes the HTTP call directly
//
// The ${NAME} syntax in headers tells the server to substitute the credential
// value from the store at execution time. The plaintext value never appears
// in the workflow definition.
//
// Requirements:
//   - Conductor server running at CONDUCTOR_SERVER_URL
//   - CONDUCTOR_AGENT_LLM_MODEL in the environment (optional)
//   - GITHUB_TOKEN stored in the Conductor server credential store
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	// HTTP tool with credential-bearing headers.
	// ${GITHUB_TOKEN} is resolved server-side from the credential store.
	listRepos := tool.HTTP("list_github_repos",
		"List public GitHub repositories for a user. Returns JSON array with name, url, and stars.",
		"https://api.github.com/users/Conductor/repos?per_page=5&sort=updated",
		tool.WithHeaders(map[string]string{
			"Authorization": "Bearer ${GITHUB_TOKEN}",
			"Accept":        "application/vnd.github.v3+json",
		}),
		tool.WithCredentials("GITHUB_TOKEN"),
	)

	agent := &ai.Agent{
		Name:         "github_http_agent",
		Model:        model,
		Tools:        []ai.ToolDef{listRepos},
		Instructions: "You list GitHub repos using the list_github_repos tool. Summarize the results.",
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	result, err := runtime.Run(context.Background(), agent, "List the repos for Conductor")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()

	// Production pattern:
	// 1. Deploy once during CI/CD: runtime.Deploy(ctx, agent)
	// 2. In a separate long-lived worker process: runtime.Serve(ctx, agent)
}
