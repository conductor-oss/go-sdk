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

// Tool types the server dispatches itself; their schemas describe what the server accepts, not a handler.

// Defaults for an HTTP tool, applied before options so an option can override
// them. The other SDKs send these keys whether or not the caller mentions them.
const (
	defaultHTTPMethod      = "GET"
	defaultHTTPContentType = "application/json"
)

// HTTP builds a tool the Conductor server calls over HTTP itself; no worker
// runs, and a ${NAME} credential in a header, declared too with
// WithCredentials, is resolved server-side without passing through your process.
func HTTP(name, description, url string, opts ...Option) ai.ToolDef {
	td := ai.ToolDef{
		Name:        name,
		Description: description,
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

// Human pauses the run for a person to answer, as a Conductor HUMAN task. The
// default schema asks one question string; WithInputSchema collects structure.
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

// Agent exposes another agent as a tool for a parent to delegate to. The
// sub-agent is serialized into this tool's config, needing no registration of
// its own; an empty name takes the sub-agent's, an empty description is generated.
func Agent(agent *ai.Agent, name, description string, opts ...Option) ai.ToolDef {
	if agent == nil {
		// Otherwise the failure surfaces only as a server-side compile error.
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
		// Stored under "agent", translated to "agentConfig" by the parent's own serializer.
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

// MCP exposes the tools of an MCP server. No worker is involved: at compile
// time the server lists what the MCP server offers and expands this definition
// into a tool per entry, each call running as a CALL_MCP_TOOL task. The input
// schema is empty because the model calls those discovered tools, which carry
// their own. A ${NAME} credential in a header is resolved server-side, and
// needs the same name in WithCredentials or Validate rejects the tool. Empty
// name or description: Python's "mcp_tools", "MCP tools from <serverURL>".
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

// WithMaxTools caps discovered tools before the server has the model pick a subset each turn. Default 64.
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

// WithInputSchema replaces the schema the model is shown, for a tool type whose default does not fit.
func WithInputSchema(doc map[string]any) Option {
	return func(t *ai.ToolDef) { t.InputSchema = doc }
}

func setConfig(t *ai.ToolDef, key string, value any) {
	if t.Config == nil {
		t.Config = map[string]any{}
	}
	t.Config[key] = value
}
