//go:build e2e

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai_e2e

import (
	"context"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// mcpServerURL returns the MCP server under test, or skips. Python's e2e uses
// mcp-testkit, which the Conductor server must be able to reach, so the
// default is the port Python's orchestrator runs it on.
func mcpServerURL(t *testing.T) string {
	url := os.Getenv("CONDUCTOR_E2E_MCP_URL")
	if url == "" {
		url = "http://localhost:3001/mcp"
	}
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-03-26",`+
			`"capabilities":{},"clientInfo":{"name":"go-e2e","version":"0"}}}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/event-stream")
	resp, err := (&http.Client{Timeout: 5 * time.Second}).Do(req)
	if err == nil {
		resp.Body.Close()
	}
	if err != nil || resp.StatusCode >= 300 {
		t.Skipf("no MCP server at %s; start one with `mcp-testkit --transport http --port 3001` "+
			"or set CONDUCTOR_E2E_MCP_URL", url)
	}
	return url
}

// An MCP tool is entirely server-side: the server discovers the MCP server's
// tools when it compiles the agent and calls them itself, so nothing here
// registers a worker. The fixture proves the definition serializes; this asks
// whether a run using it works.
//
// The question is one the model cannot answer from memory, and mcp-testkit's
// get_weather returns the same fixed values for any city, so an answer carrying
// them proves the tool was discovered, called, and its result reached the model.
// Same prompt and expectations as Python's suite 4.
func TestMCPToolResultReachesTheAnswer(t *testing.T) {
	url := mcpServerURL(t)
	rt := newRuntime(t)
	defer rt.Shutdown()

	agent := &ai.Agent{
		Name:  "go_e2e_mcp_weather",
		Model: model(t),
		Instructions: "You are a weather assistant. Use the available MCP tools to answer " +
			"questions about weather conditions.",
		Tools: []ai.ToolDef{
			tool.MCP("weather_mcp", "Weather tools via MCP: current conditions for a city", url),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	res, err := rt.Run(ctx, agent, "What is the weather in San Francisco right now?")
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	t.Logf("status=%s output=%.200q", res.Status, res.Output)

	if res.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q", res.Status, ai.StatusCompleted)
	}
	// 77°F and 45% humidity, mcp-testkit's fixed get_weather result.
	for _, want := range []string{"77", "45"} {
		if !strings.Contains(res.Output, want) {
			t.Errorf("answer omits %q from get_weather's result, so the tool's output "+
				"never reached the model: %q", want, res.Output)
		}
	}
}
