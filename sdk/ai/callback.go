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
// run: before and after the agent, each LLM call, and each tool call. Each is
// optional; a set hook runs as a worker named "<agent>_<position>". They are
// the counterpart of the Python SDK's CallbackHandler, whose six overridable
// methods map one to one to these fields.
//
// A hook may observe the run or steer it: return a non-empty map to override
// what happens next (the server merges it), or nil to let the run continue
// unchanged. An error is treated as nil, so a broken hook never blocks the
// run, matching the Python worker.
type Callbacks struct {
	OnAgentStart CallbackFunc // before the agent starts
	OnAgentEnd   CallbackFunc // after the agent finishes
	OnModelStart CallbackFunc // before each LLM call
	OnModelEnd   CallbackFunc // after each LLM call
	OnToolStart  CallbackFunc // before each tool call
	OnToolEnd    CallbackFunc // after each tool call
}

// CallbackFunc is one lifecycle hook. Returning a non-empty map overrides what
// the run does next; nil continues unchanged.
type CallbackFunc func(ctx context.Context, in CallbackInput) (map[string]any, error)

// CallbackInput is what the server passes a hook. The server sends the
// position, the agent name, the model's result and its tool calls; a field is
// empty or nil at positions that do not carry it (there is no LLM result
// before the first model call, and no tool calls when the model requested
// none). LLMResult and ToolCalls are left as their decoded JSON — the model
// result's shape varies by turn — so a hook type-asserts what it needs.
type CallbackInput struct {
	Position  string
	AgentName string
	LLMResult any
	ToolCalls any
}

// callbackPosition pairs a wire position name with the field that holds its
// hook. The order is the Python SDK's POSITION_TO_METHOD order, so the
// serialized callback list matches field for field.
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

// callbackConfigs is the wire list: one {position, taskName} per set hook, in
// position order.
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

// callbackTools returns a worker tool for each set hook, named
// "<agent>_<position>" to match the serialized task name.
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

// The worker contract, from the server's buildCallbackTask: it sends
// callback_position, agent_name, llm_result (the model output) and tool_calls,
// and expects the override map back, or {} to continue. The varying fields are
// any so binding never fails on their shape.
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
