//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package tool

import (
	"fmt"
	"strings"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/internal/schema"
)

// Tool types the server dispatches itself. None of them registers a worker, so
// none takes a Go function: the schemas below describe what the server accepts,
// not what a local handler expects.

// Defaults for an HTTP tool, applied before options so an option can override
// them. They match the other SDKs, which send these keys whether or not the
// caller mentioned them.
const (
	defaultHTTPMethod      = "GET"
	defaultHTTPContentType = "application/json"
)

// HTTP builds a tool the Conductor server calls over HTTP. No worker is
// involved — the server makes the request itself.
//
// Headers may reference a credential as ${NAME}; the server resolves it at
// execution time, so the value never passes through your process. Declare the
// same names with WithCredentials.
//
//	tool.HTTP("lookup", "Look up a record", "https://example.test/api/{id}",
//	    tool.WithHeaders(map[string]string{"Authorization": "Bearer ${API_TOKEN}"}),
//	    tool.WithCredentials("API_TOKEN"))
func HTTP(name, description, url string, opts ...Option) ai.ToolDef {
	td := ai.ToolDef{
		Name:        name,
		Description: description,
		// An HTTP tool takes no model-supplied arguments unless the caller
		// describes some with WithInputSchema.
		InputSchema: schema.EmptyObject(),
		ToolType:    ai.ToolTypeHTTP,
		Config: map[string]any{
			"url":         url,
			"method":      defaultHTTPMethod,
			"headers":     map[string]string{},
			"accept":      []string{defaultHTTPContentType},
			"contentType": defaultHTTPContentType,
		},
	}
	for _, o := range opts {
		o(&td)
	}
	return td
}

// Human pauses the run for a person to answer, as a Conductor HUMAN task.
//
// The default schema asks for a single question string. Pass WithInputSchema to
// collect something structured instead.
func Human(name, description string, opts ...Option) ai.ToolDef {
	td := ai.ToolDef{
		Name:        name,
		Description: description,
		InputSchema: schema.HumanInput(),
		ToolType:    ai.ToolTypeHuman,
	}
	for _, o := range opts {
		o(&td)
	}
	return td
}

// Agent exposes another agent as a tool, so a parent can delegate to it.
//
// The sub-agent is serialized into this tool's config, which is why it needs no
// separate registration. An empty name or description falls back to the
// sub-agent's own name and a generated description.
func Agent(agent *ai.Agent, name, description string, opts ...Option) ai.ToolDef {
	if agent == nil {
		// A nil sub-agent would serialize to nothing useful and the failure
		// would surface as a server-side compile error; name it here instead.
		return ai.ToolDef{Name: name, Description: description, ToolType: ai.ToolTypeAgent}
	}
	if name == "" {
		name = agent.Name
	}
	if description == "" {
		description = fmt.Sprintf("Invoke the %s agent", agent.Name)
	}

	td := ai.ToolDef{
		Name:        name,
		Description: description,
		InputSchema: schema.AgentRequest(),
		ToolType:    ai.ToolTypeAgent,
		// Stored under "agent" and translated to "agentConfig" when the parent
		// is serialized: the sub-agent's document has to be built by the same
		// serializer that builds the parent's.
		Config: map[string]any{"agent": agent},
	}
	for _, o := range opts {
		o(&td)
	}
	return td
}

// Defaults for an MCP tool, matching Python's mcp_tool.
const (
	defaultMCPName     = "mcp_tools"
	defaultMCPMaxTools = 64
)

// MCP exposes the tools of an MCP server. No worker is involved: when the
// agent is compiled the Conductor server lists what the MCP server offers and
// expands this one definition into a tool per entry, and each call then runs
// as a CALL_MCP_TOOL task on the server.
//
// The input schema is empty because the model never calls this definition
// itself; the discovered tools carry their own. Headers may reference a
// credential as ${NAME}, which the server resolves at execution time; declare
// the same names with WithCredentials or Validate rejects the tool.
//
// An empty name or description takes Python's default: "mcp_tools" and
// "MCP tools from <serverURL>".
//
//	tool.MCP("weather_mcp", "Weather tools", "http://localhost:3001/mcp",
//	    tool.WithHeaders(map[string]string{"Authorization": "Bearer ${MCP_KEY}"}),
//	    tool.WithCredentials("MCP_KEY"))
func MCP(name, description, serverURL string, opts ...Option) ai.ToolDef {
	if name == "" {
		name = defaultMCPName
	}
	if description == "" {
		description = "MCP tools from " + serverURL
	}
	td := ai.ToolDef{
		Name:        name,
		Description: description,
		InputSchema: map[string]any{},
		ToolType:    ai.ToolTypeMCP,
		// snake_case on purpose: these are the keys the server's compiler reads.
		Config: map[string]any{
			"server_url": serverURL,
			"max_tools":  defaultMCPMaxTools,
		},
	}
	for _, o := range opts {
		o(&td)
	}
	return td
}

// WithToolNames limits an MCP tool to the named tools on the server. Sent for
// parity with Python; the current server does not filter on it.
func WithToolNames(names ...string) Option {
	return func(t *ai.ToolDef) { setConfig(t, "tool_names", names) }
}

// WithMaxTools sets how many discovered tools an MCP server may expose before
// the server asks the model to pick a relevant subset each turn. Default 64.
func WithMaxTools(n int) Option {
	return func(t *ai.ToolDef) { setConfig(t, "max_tools", n) }
}

// WithMethod sets the HTTP method. It is upper-cased, as in the other SDKs.
func WithMethod(method string) Option {
	return func(t *ai.ToolDef) { setConfig(t, "method", strings.ToUpper(method)) }
}

// WithHeaders sets the HTTP headers, replacing any already set.
func WithHeaders(headers map[string]string) Option {
	return func(t *ai.ToolDef) { setConfig(t, "headers", headers) }
}

// WithAccept sets the Accept header values.
func WithAccept(types ...string) Option {
	return func(t *ai.ToolDef) { setConfig(t, "accept", types) }
}

// WithContentType sets the Content-Type header.
func WithContentType(contentType string) Option {
	return func(t *ai.ToolDef) { setConfig(t, "contentType", contentType) }
}

// WithInputSchema replaces the schema the model is shown. Use it when a tool
// type's default schema does not describe the arguments you want.
func WithInputSchema(doc map[string]any) Option {
	return func(t *ai.ToolDef) { t.InputSchema = doc }
}

func setConfig(t *ai.ToolDef, key string, value any) {
	if t.Config == nil {
		t.Config = map[string]any{}
	}
	t.Config[key] = value
}
