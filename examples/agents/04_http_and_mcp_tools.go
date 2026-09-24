package agents

import (
	"context"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

func formatReport(_ context.Context, in reportIn) (map[string]any, error) {
	return map[string]any{"report": fmt.Sprintf("=== %s ===\n%s\n%s", in.Title, in.Body,
		strings.Repeat("=", utf8.RuneCountInString(in.Title)+8))}, nil
}

// HTTPAndMCPTools is the Python SDK's examples/agents/04_http_and_mcp_tools.py:
// server-side HTTP and MCP tools next to a worker tool. It needs mcp-testkit
// on port 3001 and the two credentials on the server.
func HTTPAndMCPTools(model string) *ai.Agent {
	reverseAPI := tool.HTTP("reverse_string", "http://localhost:3001/api/string/reverse",
		"Reverse a string using the HTTP API",
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
	mcpTestTools := tool.MCP("mcp_test_tools", "http://localhost:3001/mcp",
		"Deterministic test tools via MCP — math, string, collection, encoding, hash, datetime, validation, and conversion operations.",
		tool.WithHeaders(map[string]string{"Authorization": "Bearer ${MCP_TEST_API_KEY}"}),
		tool.WithCredentials("MCP_TEST_API_KEY"),
	)
	return &ai.Agent{
		Name:  "http_tools_demo",
		Model: model,
		Tools: ai.Tools(
			tool.Func("format_report", formatReport, "Format a title and body into a structured report."),
			reverseAPI,
			mcpTestTools,
		),
		Instructions: "You can reverse strings and format reports. When asked to reverse a string, use reverse_string first, then format_report with the result.",
	}
}

func runHTTPAndMCPTools(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, HTTPAndMCPTools(Model()),
		"Reverse the string 'hello world' and add 33 and 21 append the result to that string, then write a report with the result.", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
