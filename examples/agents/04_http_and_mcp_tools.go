//go:build ignore

// HTTP and MCP Tools — server-side tools (no workers needed).
//
// Run with:  go run agents/04_http_and_mcp_tools.go
//
// Demonstrates:
//   - tool.HTTP: HTTP endpoints as tools (Conductor HttpTask)
//   - tool.MCP: MCP server tools (Conductor ListMcpTools + CallMcpTool)
//   - mixing Go tools with server-side tools
//
// These tools execute entirely server-side — no worker process needed for
// them; only format_report runs here.
//
// MCP test server setup (mcp-testkit):
//
//	pip install mcp-testkit
//	mcp-testkit --transport http
//
// The two credentials, HTTP_TEST_API_KEY and MCP_TEST_API_KEY, must exist in
// the Conductor server's credential store.
//
// Requirements:
//   - Conductor server with LLM support
//   - mcp-testkit running on http://localhost:3001
//   - CONDUCTOR_SERVER_URL=http://localhost:8080/api in the environment
//   - CONDUCTOR_AGENT_LLM_MODEL in the environment (optional)
package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type reportIn struct {
	Title string
	Body  string
}

func formatReport(ctx context.Context, in reportIn) (map[string]any, error) {
	return map[string]any{"report": fmt.Sprintf("=== %s ===\n%s\n%s", in.Title, in.Body,
		strings.Repeat("=", utf8.RuneCountInString(in.Title)+8))}, nil
}

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	reverseAPI := tool.HTTP("reverse_string", "Reverse a string using the HTTP API",
		"http://localhost:3001/api/string/reverse",
		tool.WithMethod("POST"),
		tool.WithHeaders(map[string]string{"Authorization": "Bearer ${HTTP_TEST_API_KEY}"}),
		tool.WithCredentials("HTTP_TEST_API_KEY"),
		tool.WithInputSchema(map[string]any{
			"type": "object",
			"properties": map[string]any{
				"text": map[string]any{"type": "string", "description": "Text to reverse"},
			},
			"required": []string{"text"},
		}),
	)

	mcpTestTools := tool.MCP("mcp_test_tools",
		"Deterministic test tools via MCP — math, string, collection, encoding, hash, datetime, validation, and conversion operations.",
		"http://localhost:3001/mcp",
		tool.WithHeaders(map[string]string{"Authorization": "Bearer ${MCP_TEST_API_KEY}"}),
		tool.WithCredentials("MCP_TEST_API_KEY"),
	)

	agent := &ai.Agent{
		Name:  "http_tools_demo",
		Model: model,
		Tools: ai.Tools(
			tool.Func("format_report", "Format a title and body into a structured report.", formatReport),
			reverseAPI,
			mcpTestTools,
		),
		Instructions: "You can reverse strings and format reports. " +
			"When asked to reverse a string, use reverse_string first, then format_report with the result.",
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	result, err := runtime.Run(context.Background(), agent,
		"Reverse the string 'hello world' and add 33 and 21 append the result to that string, then write a report with the result.")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()
}
