//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

import (
	"fmt"
	"math"
	"regexp"
	"slices"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// ToolType selects how the server dispatches a tool call; the values are the "toolType"
// field in agent-schema.json. Only ToolTypeWorker is dispatched back to this SDK, so the
// rest need no Go handler.
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

// RetryPolicy is how a failed worker call is retried. It lives on the task definition the
// runtime registers, not in agentConfig; the names are Python's retry_policy values.
type RetryPolicy string

const (
	RetryFixed              RetryPolicy = "fixed"
	RetryLinearBackoff      RetryPolicy = "linear_backoff"
	RetryExponentialBackoff RetryPolicy = "exponential_backoff"
)

// retryLogic maps a policy to Conductor's TaskDef.retryLogic constant.
var retryLogic = map[RetryPolicy]string{
	RetryFixed:              "FIXED",
	RetryLinearBackoff:      "LINEAR_BACKOFF",
	RetryExponentialBackoff: "EXPONENTIAL_BACKOFF",
}

// Task definition defaults, from the Python SDK's _default_task_def. Timeout is
// 0 because the agent controls execution duration; the short response timeout
// detects a dead worker quickly, with lease extension keeping live ones alive.
const (
	defaultRetryCount             = 2
	defaultRetryDelaySeconds      = 2
	defaultResponseTimeoutSeconds = 10
)

// ToolDef is a tool as the server sees it. Prefer the tool package constructors, which derive
// InputSchema and OutputSchema by reflection, keeping the wire format identical to other SDKs.
type ToolDef struct {
	// Name is the tool name the model calls and the task name a worker registers under. Required.
	Name string
	// Description is what the model reads to decide whether to call it.
	Description string
	// InputSchema and OutputSchema are JSON Schema for the handler's argument and return types.
	InputSchema  map[string]any
	OutputSchema map[string]any
	// ToolType defaults to ToolTypeWorker.
	ToolType ToolType

	// Config carries type-specific settings: an http url, an mcp server url, and always the credentials.
	Config map[string]any
	// Credentials are the secret names this tool may read; they land under Config on the wire.
	Credentials []string

	// ApprovalRequired pauses the run for a human before the call is dispatched.
	ApprovalRequired bool
	// Stateful routes the tool to a per-execution worker domain.
	Stateful bool
	// TimeoutSeconds and MaxCalls bound one tool. Nil leaves them to the server.
	TimeoutSeconds *int
	MaxCalls       *int

	// RetryCount, RetryDelaySeconds and RetryPolicy configure the task definition the runtime
	// registers, not agentConfig; unset takes the defaults above plus linear backoff. See
	// tool.WithRetry.
	RetryCount        *int
	RetryDelaySeconds *int
	RetryPolicy       RetryPolicy

	// Guardrails run against this tool's input or output.
	Guardrails []Guardrail

	// Handler is the worker function, set by the tool constructors. Nil for server-dispatched
	// tool types and for workers served from another process.
	Handler any
}

// Guardrail is one of RegexGuardrail, LLMGuardrail or CustomGuardrail.
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
	if t.RetryPolicy != "" {
		if _, ok := retryLogic[t.RetryPolicy]; !ok {
			return fmt.Errorf("tool %q: invalid retry policy %q", t.Name, t.RetryPolicy)
		}
	}
	switch t.ToolType {
	case ToolTypeMCP:
		if u, ok := t.Config["server_url"].(string); !ok || u == "" {
			return fmt.Errorf("mcp tool %q has no server_url", t.Name)
		}
	case ToolTypeAPI:
		if u, ok := t.Config["url"].(string); !ok || u == "" {
			return fmt.Errorf("api tool %q has no url", t.Name)
		}
	default:
		return nil
	}
	// The server resolves a ${NAME} header from the credentials the tool declares, so an
	// undeclared one would reach the remote server as literal text. Python refuses too.
	for _, ref := range credentialRefs(t.Config["headers"]) {
		if !slices.Contains(t.Credentials, ref) {
			return fmt.Errorf("%s tool %q: header placeholder ${%s} is not declared "+
				"in credentials %v", t.ToolType, t.Name, ref, t.Credentials)
		}
	}
	return nil
}

// taskDef is the task definition the runtime registers for this worker. The tool's credential
// names ride as runtimeMetadata so the registration does not wipe what the server compiled.
func (t ToolDef) taskDef() model.TaskDef {
	retries, delay := defaultRetryCount, defaultRetryDelaySeconds
	if t.RetryCount != nil {
		retries = *t.RetryCount
	}
	if t.RetryDelaySeconds != nil {
		delay = *t.RetryDelaySeconds
	}
	policy := t.RetryPolicy
	if policy == "" {
		policy = RetryLinearBackoff
	}
	return model.TaskDef{
		Name:                   t.Name,
		RetryCount:             toInt32(retries),
		RetryLogic:             retryLogic[policy],
		RetryDelaySeconds:      toInt32(delay),
		TimeoutSeconds:         0,
		ResponseTimeoutSeconds: defaultResponseTimeoutSeconds,
		TimeoutPolicy:          "RETRY",
		RuntimeMetadata:        t.Credentials,
	}
}

var credentialRef = regexp.MustCompile(`\$\{(\w+)\}`)

func credentialRefs(headers any) []string {
	var refs []string
	visit := func(v string) {
		for _, m := range credentialRef.FindAllStringSubmatch(v, -1) {
			refs = append(refs, m[1])
		}
	}
	switch h := headers.(type) {
	case map[string]string:
		for _, v := range h {
			visit(v)
		}
	case map[string]any:
		for _, v := range h {
			if s, ok := v.(string); ok {
				visit(s)
			}
		}
	}
	return refs
}

// toolConfig serializes one tool, emitting optional fields only when set, as the
// Python serializer does. agentStateful is the agent's own Stateful flag.
func (t ToolDef) toolConfig(agentStateful bool) map[string]any {
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
	if t.Stateful || agentStateful {
		cfg["stateful"] = true
	}
	if t.TimeoutSeconds != nil {
		cfg["timeoutSeconds"] = *t.TimeoutSeconds
	}
	if t.MaxCalls != nil {
		cfg["maxCalls"] = *t.MaxCalls
	}

	if conf := t.configMap(); conf != nil {
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

// configMap assembles the wire config. Credentials ride inside config, not at the top level:
// the server's compiler reads tool.config["credentials"] to collect what a task may resolve.
// definition may resolve.
func (t ToolDef) configMap() map[string]any {
	var conf map[string]any
	if len(t.Config) > 0 {
		conf = map[string]any{}
		for k, v := range t.Config {
			conf[k] = v
		}
		// An agent-as-tool's nested document must come from this serializer, not the constructor.
		if sub, ok := conf["agent"].(*Agent); ok {
			delete(conf, "agent")
			conf["agentConfig"] = sub.toConfig()
		}
	}
	if len(t.Credentials) > 0 {
		if conf == nil {
			conf = map[string]any{}
		}
		conf["credentials"] = t.Credentials
	}
	return conf
}

// toInt32 narrows a retry setting to the task definition's int32, clamping rather than wrapping.
func toInt32(v int) int32 {
	switch {
	case v < 0:
		return 0
	case v > math.MaxInt32:
		return math.MaxInt32
	}
	return int32(v)
}
