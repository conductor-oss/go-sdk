//go:build integration

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package integration

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type reportIn struct {
	Title string
	Body  string
}

// HTTP and MCP tools — the Python SDK's examples/agents/04_http_and_mcp_tools.py
// as a test.
//
// The example mixes a worker tool with two server-side tools, an HTTP
// endpoint and an MCP server, and asks for a reversed string, a sum and a
// report. This is that flow, copied, with the print replaced by validation:
// the run completes with the recorded answer and the worker tool
// format_report ran once with the reversed string and the sum in its body,
// which shows the server-side tools' results flowed back into the model's
// next step. It needs the MCP test server and the two credentials the tools
// name on the server (CONDUCTOR_SECRET_HTTP_TEST_API_KEY and
// CONDUCTOR_SECRET_MCP_TEST_API_KEY at server start). The credentials are
// consumed server-side, so nothing here can check for them up front; as the
// credential tests do, the run goes first and a failure is a skip unless
// CONDUCTOR_E2E_SECRET_PROVISIONED says they are there.
func TestExample04HTTPAndMCPTools(t *testing.T) {
	runtime := newRuntime(t)
	requireMCPTestkit(t)
	recorded := recordedAnswers(t, "04_http_and_mcp_tools")

	var reports atomic.Int32
	var lastBody atomic.Value
	formatReport := func(ctx context.Context, in reportIn) (map[string]any, error) {
		reports.Add(1)
		lastBody.Store(in.Body)
		return map[string]any{"report": fmt.Sprintf("=== %s ===\n%s\n%s", in.Title, in.Body,
			strings.Repeat("=", utf8.RuneCountInString(in.Title)+8))}, nil
	}

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
	mcpTestTools := tool.MCP("mcp_test_tools",
		"http://localhost:3001/mcp",
		"Deterministic test tools via MCP — math, string, collection, encoding, hash, datetime, validation, and conversion operations.",
		tool.WithHeaders(map[string]string{"Authorization": "Bearer ${MCP_TEST_API_KEY}"}),
		tool.WithCredentials("MCP_TEST_API_KEY"),
	)

	agent := &ai.Agent{
		Name:  "http_tools_demo",
		Model: mockModel,
		Tools: ai.Tools(
			tool.Func("format_report", formatReport, "Format a title and body into a structured report."),
			reverseAPI,
			mcpTestTools,
		),
		Instructions: "You can reverse strings and format reports. " +
			"When asked to reverse a string, use reverse_string first, then format_report with the result.",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	result, err := runtime.Run(ctx, agent,
		"Reverse the string 'hello world' and add 33 and 21 append the result to that string, then write a report with the result.")
	if err != nil || result.Status != ai.StatusCompleted {
		msg := fmt.Sprintf("err=%v", err)
		if result != nil {
			msg = fmt.Sprintf("status=%q error=%q", result.Status, result.Error)
		}
		if os.Getenv("CONDUCTOR_E2E_SECRET_PROVISIONED") == "" {
			t.Skipf("run did not complete (%s); the tools need HTTP_TEST_API_KEY and MCP_TEST_API_KEY on the server — "+
				"start it with CONDUCTOR_SECRET_HTTP_TEST_API_KEY and CONDUCTOR_SECRET_MCP_TEST_API_KEY set, "+
				"and set CONDUCTOR_E2E_SECRET_PROVISIONED to make this a failure", msg)
		}
		t.Fatalf("run did not complete: %s", msg)
	}

	// Validation, in place of the example's result.print_result().
	if n := reports.Load(); n != 1 {
		t.Errorf("format_report ran %d times, want 1", n)
	}
	if body, _ := lastBody.Load().(string); !strings.Contains(body, "dlrow olleh") || !strings.Contains(body, "54") {
		t.Errorf("format_report body %q should carry the reversed string and the sum from the server-side tools", body)
	}
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[len(recorded)-1]); got != want {
		t.Errorf("output is not the recorded answer\n--- got ---\n%s\n--- recorded ---\n%s", got, want)
	}
}

// requireMCPTestkit skips unless the MCP test server the example points at is
// reachable, as the MCP feature test does.
func requireMCPTestkit(t *testing.T) {
	t.Helper()
	url := os.Getenv("CONDUCTOR_E2E_MCP_URL")
	if url == "" {
		url = "http://localhost:3001/mcp"
	}
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/event-stream")
	resp, err := (&http.Client{Timeout: 3 * time.Second}).Do(req)
	if err != nil {
		t.Skipf("no MCP server at %s; start one with `mcp-testkit --transport http --port 3001`", url)
	}
	resp.Body.Close()
}
