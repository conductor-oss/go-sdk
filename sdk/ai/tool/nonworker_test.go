//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package tool_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/internal/schema"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// These constructors are the half the golden fixture cannot check. That test
// lives in package ai, which cannot import this package, so it builds the same
// tool definitions as literals. What is verified here is that the constructors
// produce those literals — if the two drift, one of the two tests fails.

func TestHTTPDefaults(t *testing.T) {
	td := tool.HTTP("lookup", "Look up a record.", "https://example.test/api/{id}")

	if td.ToolType != ai.ToolTypeHTTP {
		t.Errorf("toolType = %q, want %q", td.ToolType, ai.ToolTypeHTTP)
	}
	want := map[string]any{
		"url":         "https://example.test/api/{id}",
		"method":      "GET",
		"headers":     map[string]string{},
		"accept":      []string{"application/json"},
		"contentType": "application/json",
	}
	if !reflect.DeepEqual(td.Config, want) {
		t.Errorf("config = %#v\nwant %#v", td.Config, want)
	}
	if !reflect.DeepEqual(td.InputSchema, schema.EmptyObject()) {
		t.Errorf("inputSchema = %#v, want an empty object schema", td.InputSchema)
	}
}

func TestHTTPOptionsOverrideDefaults(t *testing.T) {
	td := tool.HTTP("post", "Send it", "https://example.test/x",
		// Lower case on purpose: the method is upper-cased, as in the other SDKs.
		tool.WithMethod("post"),
		tool.WithHeaders(map[string]string{"X-Api-Version": "2"}),
		tool.WithAccept("text/plain"),
		tool.WithContentType("text/plain"),
		tool.WithCredentials("API_TOKEN"),
	)

	if got := td.Config["method"]; got != "POST" {
		t.Errorf("method = %v, want POST", got)
	}
	if got, ok := td.Config["headers"].(map[string]string); !ok || got["X-Api-Version"] != "2" {
		t.Errorf("headers = %#v", td.Config["headers"])
	}
	if got, ok := td.Config["accept"].([]string); !ok || len(got) != 1 || got[0] != "text/plain" {
		t.Errorf("accept = %#v", td.Config["accept"])
	}
	if got := td.Config["contentType"]; got != "text/plain" {
		t.Errorf("contentType = %v", got)
	}
	// Credentials stay on the field; toolConfig moves them under config.
	if len(td.Credentials) != 1 || td.Credentials[0] != "API_TOKEN" {
		t.Errorf("credentials = %v", td.Credentials)
	}
}

func TestHumanUsesTheFixedSchema(t *testing.T) {
	td := tool.Human("ask_human", "Ask a person to decide.")

	if td.ToolType != ai.ToolTypeHuman {
		t.Errorf("toolType = %q, want %q", td.ToolType, ai.ToolTypeHuman)
	}
	if !reflect.DeepEqual(td.InputSchema, schema.HumanInput()) {
		t.Errorf("inputSchema = %#v", td.InputSchema)
	}
	// A human tool has nothing to configure; an empty config map would still
	// serialize as a "config" key the other SDKs do not send.
	if td.Config != nil {
		t.Errorf("config = %#v, want nil", td.Config)
	}
	if td.Handler != nil {
		t.Error("a human tool must register no worker")
	}
}

func TestAgentToolCarriesTheSubAgent(t *testing.T) {
	billing := &ai.Agent{Name: "billing", Model: "openai/gpt-4o"}
	td := tool.Agent(billing, "delegate_billing", "Delegate.")

	if td.ToolType != ai.ToolTypeAgent {
		t.Errorf("toolType = %q, want %q", td.ToolType, ai.ToolTypeAgent)
	}
	if got := td.Config["agent"]; got != billing {
		t.Errorf("config[agent] = %#v, want the sub-agent itself", got)
	}
	if !reflect.DeepEqual(td.InputSchema, schema.AgentRequest()) {
		t.Errorf("inputSchema = %#v", td.InputSchema)
	}
}

// An omitted name or description falls back to the sub-agent's own.
func TestAgentToolDefaultsToTheSubAgentName(t *testing.T) {
	td := tool.Agent(&ai.Agent{Name: "billing"}, "", "")

	if td.Name != "billing" {
		t.Errorf("name = %q, want %q", td.Name, "billing")
	}
	if td.Description != "Invoke the billing agent" {
		t.Errorf("description = %q", td.Description)
	}
}

// The fixed schemas must be fresh each call: a caller adding a property to one
// tool's schema must not change another's.
func TestFixedSchemasAreNotShared(t *testing.T) {
	a := tool.Human("a", "first")
	b := tool.Human("b", "second")

	props, _ := a.InputSchema["properties"].(map[string]any)
	props["extra"] = map[string]any{"type": "string"}

	bProps, _ := b.InputSchema["properties"].(map[string]any)
	if _, leaked := bProps["extra"]; leaked {
		t.Error("editing one tool's schema changed another's")
	}
}

// The bare constructor emits exactly what Python's mcp_tool(server_url=...)
// does: default name and description, an empty schema, and the two config
// keys the server reads.
func TestMCPDefaults(t *testing.T) {
	td := tool.MCP("", "", "http://localhost:3001/mcp")

	if td.ToolType != ai.ToolTypeMCP {
		t.Errorf("toolType = %q, want %q", td.ToolType, ai.ToolTypeMCP)
	}
	if td.Name != "mcp_tools" || td.Description != "MCP tools from http://localhost:3001/mcp" {
		t.Errorf("name/description = %q / %q, want Python's defaults", td.Name, td.Description)
	}
	want := map[string]any{"server_url": "http://localhost:3001/mcp", "max_tools": 64}
	if !reflect.DeepEqual(td.Config, want) {
		t.Errorf("config = %#v\nwant %#v", td.Config, want)
	}
	// Empty, not an empty object schema: the model never calls this
	// definition, and Python sends {} here.
	if !reflect.DeepEqual(td.InputSchema, map[string]any{}) {
		t.Errorf("inputSchema = %#v, want {}", td.InputSchema)
	}
	if err := td.Validate(); err != nil {
		t.Errorf("Validate: %v", err)
	}
}

func TestMCPOptions(t *testing.T) {
	td := tool.MCP("secured_mcp", "Authenticated MCP tools.", "http://localhost:3002/mcp",
		tool.WithHeaders(map[string]string{"Authorization": "Bearer ${MCP_AUTH_KEY}"}),
		tool.WithToolNames("get_weather", "math_add"),
		tool.WithMaxTools(16),
		tool.WithCredentials("MCP_AUTH_KEY"),
	)
	want := map[string]any{
		"server_url": "http://localhost:3002/mcp",
		"max_tools":  16,
		"headers":    map[string]string{"Authorization": "Bearer ${MCP_AUTH_KEY}"},
		"tool_names": []string{"get_weather", "math_add"},
	}
	if !reflect.DeepEqual(td.Config, want) {
		t.Errorf("config = %#v\nwant %#v", td.Config, want)
	}
	if err := td.Validate(); err != nil {
		t.Errorf("Validate: %v", err)
	}
}

// Python's mcp_tool raises on these two at construction; Go reports them from
// Validate, which Agent.Validate and Runtime.Run call.
func TestMCPValidation(t *testing.T) {
	if err := tool.MCP("x", "", "").Validate(); err == nil || !strings.Contains(err.Error(), "server_url") {
		t.Errorf("empty server_url: err = %v, want a server_url error", err)
	}
	undeclared := tool.MCP("x", "", "http://h/mcp",
		tool.WithHeaders(map[string]string{"Authorization": "Bearer ${MCP_AUTH_KEY}"}))
	if err := undeclared.Validate(); err == nil || !strings.Contains(err.Error(), "MCP_AUTH_KEY") {
		t.Errorf("undeclared placeholder: err = %v, want it named", err)
	}
}
