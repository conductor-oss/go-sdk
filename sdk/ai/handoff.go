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
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// HandoffCondition decides when one sub-agent passes control to another. It is
// set on the parent agent, usually with StrategySwarm. The interface is sealed:
// the server knows a fixed set of types, so a caller-defined one cannot compile.
type HandoffCondition interface {
	handoffConfig(agentName string) map[string]any
	validateHandoff() error
	handoffTarget() string
}

// handoffTaskName is the name an OnCondition predicate must be registered under.
func handoffTaskName(agentName, target string) string {
	return agentName + "_handoff_" + target
}

// OnToolResult hands off after a named tool returns.
type OnToolResult struct {
	// Target is the sub-agent to hand control to. Required.
	Target string
	// ToolName is the tool whose result is watched. Required.
	ToolName string
	// ResultContains narrows the trigger to results containing this text; empty means any successful call.
	ResultContains string
}

func (h *OnToolResult) handoffConfig(string) map[string]any {
	cfg := map[string]any{
		"target":   h.Target,
		"type":     "on_tool_result",
		"toolName": h.ToolName,
	}
	if h.ResultContains != "" {
		cfg["resultContains"] = h.ResultContains
	}
	return cfg
}

func (h *OnToolResult) handoffTarget() string { return h.Target }

func (h *OnToolResult) validateHandoff() error {
	if h.Target == "" {
		return fmt.Errorf("on_tool_result handoff has no target")
	}
	if h.ToolName == "" {
		return fmt.Errorf("on_tool_result handoff to %q has no toolName", h.Target)
	}
	return nil
}

// OnTextMention hands off when the agent's own text mentions something.
type OnTextMention struct {
	// Target is the sub-agent to hand control to. Required.
	Target string
	// Text is the substring to watch for. Required.
	Text string
}

func (h *OnTextMention) handoffConfig(string) map[string]any {
	return map[string]any{
		"target": h.Target,
		"type":   "on_text_mention",
		"text":   h.Text,
	}
}

func (h *OnTextMention) handoffTarget() string { return h.Target }

func (h *OnTextMention) validateHandoff() error {
	if h.Target == "" {
		return fmt.Errorf("on_text_mention handoff has no target")
	}
	if h.Text == "" {
		return fmt.Errorf("on_text_mention handoff to %q has no text", h.Target)
	}
	return nil
}

// OnCondition hands off when a Go predicate says so. The predicate runs as a
// worker named after the parent and target pair, so a parent may have only one
// condition per target; a second would collide on the task name.
type OnCondition struct {
	// Target is the sub-agent to hand control to. Required.
	Target string
	// Condition is evaluated after each turn. Required.
	Condition HandoffFunc
}

// HandoffFunc decides whether to hand off; true moves control to the Target.
type HandoffFunc func(ctx context.Context, state HandoffState) (bool, error)

// HandoffState is the server's input to the predicate, normalised. result is a
// string, or a list when the agent produced no text, and is empty on any turn ending
// in a transfer-tool call, so it is flattened to text; active_agent is a string index
// into [parent, sub-agents...] resolved to a name, as Python's idx_to_name map does.
type HandoffState struct {
	// Result is the active agent's last response, as text.
	Result string
	// Conversation is the transcript so far.
	Conversation string
	// Context is the agent state shared across turns.
	Context map[string]any
	// ActiveAgent is the name of the agent currently holding control, or the
	// raw value if the index cannot be resolved.
	ActiveAgent string
	// ToolResults are the results of the last turn's tool calls.
	ToolResults []any
}

// handoffIn is the task input, keyed as the server sends it; the loosely typed
// fields have no fixed JSON type (see HandoffState).
type handoffIn struct {
	Result       any            `json:"result"`
	Conversation any            `json:"conversation"`
	Context      map[string]any `json:"context"`
	ActiveAgent  any            `json:"active_agent"`
	ToolResults  []any          `json:"tool_results"`
}

// handoffOut is the task output; the swarm resolver throws unless handoff is a boolean.
type handoffOut struct {
	Handoff bool `json:"handoff"`
}

// textOf flattens a string, a list of strings, or content blocks with a text
// or content field into one string.
func textOf(v any) string {
	switch x := v.(type) {
	case nil:
		return ""
	case string:
		return x
	case []any:
		var parts []string
		for _, e := range x {
			if t := textOf(e); t != "" {
				parts = append(parts, t)
			}
		}
		return strings.Join(parts, "\n")
	case map[string]any:
		if t, ok := x["text"].(string); ok {
			return t
		}
		if t, ok := x["content"].(string); ok {
			return t
		}
		b, err := json.Marshal(x)
		if err != nil {
			return fmt.Sprint(x)
		}
		return string(b)
	default:
		return fmt.Sprint(x)
	}
}

// resolveAgent turns the active_agent index into a name, or returns it as text.
func resolveAgent(v any, agents []string) string {
	raw := textOf(v)
	if i, err := strconv.Atoi(strings.TrimSpace(raw)); err == nil && i >= 0 && i < len(agents) {
		return agents[i]
	}
	return raw
}

// handoffHandler adapts a HandoffFunc to the worker contract. agents is the
// server's index space: the parent, then its sub-agents in declaration order.
func (h *OnCondition) handoffHandler(agents []string) func(context.Context, handoffIn) (handoffOut, error) {
	return func(ctx context.Context, in handoffIn) (handoffOut, error) {
		hit, err := h.Condition(ctx, HandoffState{
			Result:       textOf(in.Result),
			Conversation: textOf(in.Conversation),
			Context:      in.Context,
			ActiveAgent:  resolveAgent(in.ActiveAgent, agents),
			ToolResults:  in.ToolResults,
		})
		if err != nil {
			// Python's should_handoff swallows exceptions and returns False.
			return handoffOut{Handoff: false}, nil
		}
		return handoffOut{Handoff: hit}, nil
	}
}

func (a *Agent) onConditions() []*OnCondition {
	var out []*OnCondition
	for _, h := range a.Handoffs {
		if c, ok := h.(*OnCondition); ok && c.Condition != nil {
			out = append(out, c)
		}
	}
	return out
}

func (h *OnCondition) handoffConfig(agentName string) map[string]any {
	return map[string]any{
		"target":   h.Target,
		"type":     "on_condition",
		"taskName": handoffTaskName(agentName, h.Target),
	}
}

func (h *OnCondition) handoffTarget() string { return h.Target }

func (h *OnCondition) validateHandoff() error {
	if h.Target == "" {
		return fmt.Errorf("on_condition handoff has no target")
	}
	if h.Condition == nil {
		return fmt.Errorf("on_condition handoff to %q has no Condition function", h.Target)
	}
	return nil
}

// validateHandoffs checks each handoff points at a real sub-agent: a target
// that names nothing compiles server side and fails only when it fires.
func (a *Agent) validateHandoffs() error {
	if len(a.Handoffs) == 0 {
		return nil
	}

	known := make(map[string]struct{}, len(a.Agents))
	for _, sub := range a.Agents {
		if sub != nil {
			known[sub.Name] = struct{}{}
		}
	}

	seen := make(map[string]struct{}, len(a.Handoffs))
	for _, h := range a.Handoffs {
		if err := h.validateHandoff(); err != nil {
			return fmt.Errorf("agent %q: %w", a.Name, err)
		}
		target := h.handoffTarget()
		if _, ok := known[target]; !ok {
			return fmt.Errorf(
				"agent %q hands off to %q, which is not one of its sub-agents",
				a.Name, target)
		}
		if _, isCond := h.(*OnCondition); isCond {
			if _, dup := seen[target]; dup {
				return fmt.Errorf(
					"agent %q has more than one on_condition handoff to %q; "+
						"they would share the task name %q",
					a.Name, target, handoffTaskName(a.Name, target))
			}
			seen[target] = struct{}{}
		}
	}

	for from, tos := range a.AllowedTransitions {
		if _, ok := known[from]; !ok {
			return fmt.Errorf(
				"agent %q allows transitions from %q, which is not one of its sub-agents",
				a.Name, from)
		}
		for _, to := range tos {
			if _, ok := known[to]; !ok {
				return fmt.Errorf(
					"agent %q allows a transition from %q to %q, which is not one of "+
						"its sub-agents", a.Name, from, to)
			}
		}
	}
	return nil
}
