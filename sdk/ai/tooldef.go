//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

import "fmt"

// ToolType selects how the server dispatches a tool call.
//
// The values come from the "toolType" field in agent-schema.json. Only
// ToolTypeWorker is dispatched back to this SDK; the rest the server handles
// itself, which is why they need no Go handler.
type ToolType string

const (
	// ToolTypeWorker dispatches to a Conductor worker — a Go function.
	ToolTypeWorker ToolType = "worker"
	ToolTypeHTTP   ToolType = "http"
	ToolTypeAPI    ToolType = "api"
	ToolTypeMCP    ToolType = "mcp"
	ToolTypeHuman  ToolType = "human"
	ToolTypeAgent  ToolType = "agent_tool"

	ToolTypeGenerateImage ToolType = "generate_image"
	ToolTypeGenerateAudio ToolType = "generate_audio"
	ToolTypeGenerateVideo ToolType = "generate_video"
	ToolTypeGeneratePDF   ToolType = "generate_pdf"

	ToolTypeRAGIndex  ToolType = "rag_index"
	ToolTypeRAGSearch ToolType = "rag_search"

	ToolTypePullWorkflowMessages ToolType = "pull_workflow_messages"
)

var validToolTypes = map[ToolType]struct{}{
	ToolTypeWorker: {}, ToolTypeHTTP: {}, ToolTypeAPI: {}, ToolTypeMCP: {},
	ToolTypeHuman: {}, ToolTypeAgent: {},
	ToolTypeGenerateImage: {}, ToolTypeGenerateAudio: {},
	ToolTypeGenerateVideo: {}, ToolTypeGeneratePDF: {},
	ToolTypeRAGIndex: {}, ToolTypeRAGSearch: {},
	ToolTypePullWorkflowMessages: {},
}

// ToolDef is a tool as the server sees it. Build one with the constructors in
// the tool package rather than by hand: they derive InputSchema and
// OutputSchema by reflection, which is what keeps the wire format identical to
// the other SDKs.
type ToolDef struct {
	// Name is the tool name the model calls, and the Conductor task name a
	// worker is registered under. Required.
	Name string
	// Description is what the model reads to decide whether to call it.
	Description string
	// InputSchema and OutputSchema are JSON Schema documents derived from the
	// handler's argument and return types.
	InputSchema  map[string]any
	OutputSchema map[string]any
	// ToolType defaults to ToolTypeWorker.
	ToolType ToolType

	// Config carries type-specific settings: a URL for http, a server URL for
	// mcp, and always the declared credential names.
	Config map[string]any
	// Credentials are the secret names this tool may read. They land under
	// Config on the wire, which is where the server's compiler looks for them.
	Credentials []string

	// ApprovalRequired pauses the run for a human before the call is dispatched.
	ApprovalRequired bool
	// Stateful routes the tool to a per-execution worker domain.
	Stateful bool
	// TimeoutSeconds and MaxCalls bound one tool. Nil leaves them to the server.
	TimeoutSeconds *int
	MaxCalls       *int

	// Guardrails run against this tool's input or output.
	Guardrails []Guardrail

	// Handler is the Go function to register as a worker, set by the tool
	// constructors. Nil for tool types the server dispatches itself, and for
	// external workers served from another process.
	Handler any
}

// Guardrail is a placeholder until guardrails land; it keeps ToolDef's shape
// stable so adding them is not a breaking change.
type Guardrail interface {
	guardrailConfig() map[string]any
}

// Validate reports a malformed tool.
func (t ToolDef) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("tool name must be a non-empty string")
	}
	if !validName.MatchString(t.Name) {
		return fmt.Errorf("invalid tool name %q: must start with a letter or "+
			"underscore and contain only letters, digits, underscores and hyphens", t.Name)
	}
	if t.ToolType != "" {
		if _, ok := validToolTypes[t.ToolType]; !ok {
			return fmt.Errorf("invalid toolType %q for tool %q", t.ToolType, t.Name)
		}
	}
	return nil
}

// toolConfig serializes one tool. Optional fields are emitted only when set,
// matching the Python serializer field for field.
func (t ToolDef) toolConfig() map[string]any {
	tt := t.ToolType
	if tt == "" {
		tt = ToolTypeWorker
	}
	cfg := map[string]any{
		"name":        t.Name,
		"description": t.Description,
		"inputSchema": t.InputSchema,
		"toolType":    string(tt),
	}
	if len(t.OutputSchema) > 0 {
		cfg["outputSchema"] = t.OutputSchema
	}
	if t.ApprovalRequired {
		cfg["approvalRequired"] = true
	}
	if t.Stateful {
		cfg["stateful"] = true
	}
	if t.TimeoutSeconds != nil {
		cfg["timeoutSeconds"] = *t.TimeoutSeconds
	}
	if t.MaxCalls != nil {
		cfg["maxCalls"] = *t.MaxCalls
	}

	// Credentials ride inside config, not at the top level: the server's
	// compiler reads tool.config["credentials"] when collecting what a task
	// definition may resolve.
	var conf map[string]any
	if len(t.Config) > 0 {
		conf = map[string]any{}
		for k, v := range t.Config {
			conf[k] = v
		}
	}
	if len(t.Credentials) > 0 {
		if conf == nil {
			conf = map[string]any{}
		}
		conf["credentials"] = t.Credentials
	}
	if conf != nil {
		cfg["config"] = conf
	}

	if len(t.Guardrails) > 0 {
		gs := make([]any, 0, len(t.Guardrails))
		for _, g := range t.Guardrails {
			gs = append(gs, g.guardrailConfig())
		}
		cfg["guardrails"] = gs
	}
	return cfg
}
