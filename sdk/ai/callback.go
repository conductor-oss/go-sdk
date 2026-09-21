//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

import "context"

// Callbacks are lifecycle hooks the server calls at fixed points in an agent
// run; each is optional, and a set hook runs as a worker named
// "<agent>_<position>", mirroring the Python SDK's CallbackHandler methods. A
// hook returns a non-empty map to override what happens next (the server
// merges it), or nil to continue; an error counts as nil, so a broken hook
// never blocks the run, as in the Python worker.
type Callbacks struct {
	OnAgentStart CallbackFunc // before the agent starts
	OnAgentEnd   CallbackFunc // after the agent finishes
	OnModelStart CallbackFunc // before each LLM call
	OnModelEnd   CallbackFunc // after each LLM call
	OnToolStart  CallbackFunc // before each tool call
	OnToolEnd    CallbackFunc // after each tool call
}

// CallbackFunc is one lifecycle hook.
type CallbackFunc func(ctx context.Context, in CallbackInput) (map[string]any, error)

// CallbackInput is what the server passes a hook; a field is empty or nil at
// positions that do not carry it. LLMResult and ToolCalls stay as decoded JSON
// because the model result's shape varies by turn.
type CallbackInput struct {
	Position  string
	AgentName string
	LLMResult any
	ToolCalls any
}

// callbackPosition pairs a wire position name with the field that holds its
// hook, in the Python SDK's POSITION_TO_METHOD order.
type callbackPosition struct {
	name string
	fn   func(c *Callbacks) CallbackFunc
}

var callbackPositions = []callbackPosition{
	{"before_agent", func(c *Callbacks) CallbackFunc { return c.OnAgentStart }},
	{"after_agent", func(c *Callbacks) CallbackFunc { return c.OnAgentEnd }},
	{"before_model", func(c *Callbacks) CallbackFunc { return c.OnModelStart }},
	{"after_model", func(c *Callbacks) CallbackFunc { return c.OnModelEnd }},
	{"before_tool", func(c *Callbacks) CallbackFunc { return c.OnToolStart }},
	{"after_tool", func(c *Callbacks) CallbackFunc { return c.OnToolEnd }},
}

func (a *Agent) callbackConfigs() []any {
	if a.Callbacks == nil {
		return nil
	}
	var out []any
	for _, p := range callbackPositions {
		if p.fn(a.Callbacks) != nil {
			out = append(out, map[string]any{
				"position": p.name,
				"taskName": a.Name + "_" + p.name,
			})
		}
	}
	return out
}

// callbackTools returns a worker tool for each set hook, named to match its
// serialized task name.
func (a *Agent) callbackTools() []ToolDef {
	if a.Callbacks == nil {
		return nil
	}
	var out []ToolDef
	for _, p := range callbackPositions {
		if fn := p.fn(a.Callbacks); fn != nil {
			out = append(out, ToolDef{
				Name:    a.Name + "_" + p.name,
				Handler: callbackHandler(fn),
			})
		}
	}
	return out
}

// The worker contract, from the server's buildCallbackTask; the varying fields
// are any so binding never fails on their shape.
type callbackIn struct {
	Position  string `json:"callback_position"`
	AgentName string `json:"agent_name"`
	LLMResult any    `json:"llm_result"`
	ToolCalls any    `json:"tool_calls"`
}

func callbackHandler(fn CallbackFunc) func(context.Context, callbackIn) (map[string]any, error) {
	return func(ctx context.Context, in callbackIn) (map[string]any, error) {
		out, err := fn(ctx, CallbackInput(in))
		if err != nil || out == nil {
			return map[string]any{}, nil
		}
		return out, nil
	}
}
