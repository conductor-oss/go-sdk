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

// HandoffCondition decides when one sub-agent passes control to another.
//
// Handoffs belong to a parent agent, usually with StrategySwarm: the parent
// holds the sub-agents and the rules for moving between them. The interface is
// sealed the same way TerminationCondition is — the server understands a fixed
// set of types, so a caller-defined one could not be compiled.
type HandoffCondition interface {
	// handoffConfig needs the parent's name because an on_condition handoff
	// references a worker whose task name is derived from both the parent and
	// the target.
	handoffConfig(agentName string) map[string]any
	validateHandoff() error
	// target is read by Agent.Validate to check the handoff points at a real
	// sub-agent.
	handoffTarget() string
}

// handoffSuffix builds the derived task name for an OnCondition handoff:
// "<parent>_handoff_<target>". The runtime must register the predicate under
// exactly this name.
func handoffTaskName(agentName, target string) string {
	return agentName + "_handoff_" + target
}

// OnToolResult hands off after a named tool returns.
//
// With ResultContains set, only a result containing that substring triggers
// the handoff; without it, any successful call does.
type OnToolResult struct {
	// Target is the sub-agent to hand control to. Required.
	Target string
	// ToolName is the tool whose result is watched. Required.
	ToolName string
	// ResultContains narrows the trigger to results containing this text.
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

// OnCondition hands off when a Go predicate says so.
//
// The predicate runs as a worker, registered under "<parent>_handoff_<target>",
// so one target may have only one condition per parent — the task name is
// derived from the pair and a second would collide.
type OnCondition struct {
	// Target is the sub-agent to hand control to. Required.
	Target string
	// Condition is evaluated after each turn. Required.
	Condition HandoffFunc
}

// HandoffFunc decides whether to hand off. Returning true moves control to
// the handoff's Target.
type HandoffFunc func(ctx context.Context, state HandoffState) (bool, error)

// HandoffState is what the server hands the predicate, normalised so a check
// can be written against plain values.
//
// The server's own inputs are looser than they look. result is the active
// agent's last response: usually a string, but a list when that agent
// produced no text — and every turn that ends in a transfer-tool call leaves
// it empty. active_agent is an *index*, as a string, into the server's list of
// [parent, sub-agents...], so 0 is the parent itself. Python resolves the index
// to a name through an idx_to_name map; this does the same. Result is
// flattened to text for either shape.
type HandoffState struct {
	// Result is the active agent's last response, as text.
	Result string
	// Conversation is the transcript so far.
	Conversation string
	// Context is the agent state shared across turns.
	Context map[string]any
	// ActiveAgent is the name of the agent currently holding control — the
	// parent's own name when control has returned to it. Falls back to the raw
	// value if the server sent an index that cannot be resolved.
	ActiveAgent string
	// ToolResults are the results of the last turn's tool calls.
	ToolResults []any
}

// handoffIn is the task input, keyed as the server sends it. Result,
// conversation and active_agent are bound loosely because their JSON type is
// not fixed — see HandoffState.
type handoffIn struct {
	Result       any            `json:"result"`
	Conversation any            `json:"conversation"`
	Context      map[string]any `json:"context"`
	ActiveAgent  any            `json:"active_agent"`
	ToolResults  []any          `json:"tool_results"`
}

// handoffOut is the task output. The swarm resolver reads the handoff key
// and throws unless it is a boolean, so it is never omitted.
type handoffOut struct {
	Handoff bool `json:"handoff"`
}

// textOf flattens a value the server may send as a string, a list of strings,
// or a list of content blocks with text or content fields, into one string.
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

// resolveAgent turns the server's active_agent index into a name. Anything
// that is not a valid index is returned as text, so a caller still sees what
// the server sent.
func resolveAgent(v any, agents []string) string {
	raw := textOf(v)
	if i, err := strconv.Atoi(strings.TrimSpace(raw)); err == nil && i >= 0 && i < len(agents) {
		return agents[i]
	}
	return raw
}

// handoffHandler adapts a HandoffFunc to the worker contract. agents is the
// server's index space — the parent first, then its sub-agents in declaration
// order — used to resolve active_agent.
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
			// A predicate that errors does not hand off, matching Python's
			// should_handoff, which swallows exceptions and returns False.
			return handoffOut{Handoff: false}, nil
		}
		return handoffOut{Handoff: hit}, nil
	}
}

// onConditions returns the agent's worker-backed handoffs, so each can be
// registered under its derived task name.
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

// validateHandoffs checks each handoff and that it points at a real sub-agent.
//
// A target that names nothing would compile server side and then fail at the
// moment the handoff fires, which is the worst time to find out.
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
		// Two on_condition handoffs to the same target would derive the same
		// worker task name, so the second would silently never run.
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

	// AllowedTransitions restricts movement between sub-agents, so every name
	// in it has to be one.
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
