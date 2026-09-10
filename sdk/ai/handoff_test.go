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
	"errors"
	"strings"
	"testing"
)

func noHandoff(context.Context, HandoffState) (bool, error) { return false, nil }

// The derived task name is the contract between the serializer and the worker
// registry: the runtime has to register the predicate under exactly this name,
// so a change here is a change to the wire format.
func TestOnConditionTaskName(t *testing.T) {
	h := &OnCondition{Target: "billing", Condition: noHandoff}
	cfg := h.handoffConfig("swarm")

	if got := cfg["taskName"]; got != "swarm_handoff_billing" {
		t.Errorf("taskName = %v, want swarm_handoff_billing", got)
	}
	if got := cfg["type"]; got != "on_condition" {
		t.Errorf("type = %v, want on_condition", got)
	}
	// The predicate itself is never serialized.
	if _, ok := cfg["condition"]; ok {
		t.Error("the condition function must not reach the wire")
	}
}

// ResultContains is the one conditional key; the fixture sets it, so this
// covers its absence.
func TestOnToolResultOmitsEmptyResultContains(t *testing.T) {
	cfg := (&OnToolResult{Target: "refunds", ToolName: "refund"}).handoffConfig("swarm")
	if _, ok := cfg["resultContains"]; ok {
		t.Errorf("resultContains should be omitted when unset, got %v", cfg["resultContains"])
	}
	if got := cfg["toolName"]; got != "refund" {
		t.Errorf("toolName = %v", got)
	}
}

func TestHandoffValidation(t *testing.T) {
	// Every case is a swarm of two sub-agents, so a bad target is the only
	// thing under test in the target cases.
	subs := []*Agent{
		{Name: "billing", Model: testModel},
		{Name: "refunds", Model: testModel},
	}

	cases := []struct {
		name        string
		handoffs    []HandoffCondition
		transitions map[string][]string
		want        string
	}{
		{
			name:     "tool result without a tool name",
			handoffs: []HandoffCondition{&OnToolResult{Target: "refunds"}},
			want:     "has no toolName",
		},
		{
			name:     "text mention without text",
			handoffs: []HandoffCondition{&OnTextMention{Target: "refunds"}},
			want:     "has no text",
		},
		{
			name:     "condition without a function",
			handoffs: []HandoffCondition{&OnCondition{Target: "refunds"}},
			want:     "has no Condition function",
		},
		{
			name:     "no target",
			handoffs: []HandoffCondition{&OnTextMention{Text: "x"}},
			want:     "has no target",
		},
		{
			// The failure this catches would otherwise surface only when the
			// handoff fired, mid-run.
			name:     "target is not a sub-agent",
			handoffs: []HandoffCondition{&OnTextMention{Target: "shipping", Text: "x"}},
			want:     "not one of its sub-agents",
		},
		{
			// Both would derive swarm_handoff_refunds, so the second would
			// never be dispatched.
			name: "two conditions to the same target",
			handoffs: []HandoffCondition{
				&OnCondition{Target: "refunds", Condition: noHandoff},
				&OnCondition{Target: "refunds", Condition: noHandoff},
			},
			want: "more than one on_condition handoff",
		},
		{
			name:        "transition from an unknown agent",
			handoffs:    []HandoffCondition{&OnTextMention{Target: "refunds", Text: "x"}},
			transitions: map[string][]string{"shipping": {"refunds"}},
			want:        "allows transitions from",
		},
		{
			name:        "transition to an unknown agent",
			handoffs:    []HandoffCondition{&OnTextMention{Target: "refunds", Text: "x"}},
			transitions: map[string][]string{"billing": {"shipping"}},
			want:        "to \"shipping\"",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := &Agent{
				Name:               "swarm",
				Model:              testModel,
				Strategy:           StrategySwarm,
				Agents:             subs,
				Handoffs:           tc.handoffs,
				AllowedTransitions: tc.transitions,
			}
			err := a.Validate()
			if err == nil {
				t.Fatalf("want an error mentioning %q, got nil", tc.want)
			}
			if !strings.Contains(err.Error(), tc.want) {
				t.Errorf("error = %q, want it to mention %q", err, tc.want)
			}
		})
	}
}

// Two conditions to *different* targets is fine — the task names differ.
func TestTwoConditionsToDifferentTargetsAreFine(t *testing.T) {
	a := &Agent{
		Name:     "swarm",
		Model:    testModel,
		Strategy: StrategySwarm,
		Agents: []*Agent{
			{Name: "billing", Model: testModel},
			{Name: "refunds", Model: testModel},
		},
		Handoffs: []HandoffCondition{
			&OnCondition{Target: "billing", Condition: noHandoff},
			&OnCondition{Target: "refunds", Condition: noHandoff},
		},
	}
	if err := a.Validate(); err != nil {
		t.Errorf("distinct targets must validate: %v", err)
	}
}

func TestValidHandoffsPass(t *testing.T) {
	if err := swarmAgent().Validate(); err != nil {
		t.Errorf("the fixture agent must validate: %v", err)
	}
}

// The plan-execute slots hold sub-agents, so they get the same validation the
// list and the router already had. Without this a bad planner would serialize
// and fail server side.
func TestPlanExecuteSlotsAreValidated(t *testing.T) {
	cases := []struct {
		name  string
		agent *Agent
		want  string
	}{
		{
			name: "planner with an invalid name",
			agent: &Agent{
				Name: "root", Model: testModel, Strategy: StrategyPlanExecute,
				Planner: &Agent{Name: "9bad", Model: testModel},
			},
			want: "planner: invalid agent name",
		},
		{
			name: "fallback with no name",
			agent: &Agent{
				Name: "root", Model: testModel, Strategy: StrategyPlanExecute,
				Planner:  &Agent{Name: "ok", Model: testModel},
				Fallback: &Agent{Model: testModel},
			},
			want: "fallback: agent name must be",
		},
		{
			name: "planner carrying a malformed tool",
			agent: &Agent{
				Name: "root", Model: testModel, Strategy: StrategyPlanExecute,
				Planner: &Agent{
					Name: "p", Model: testModel,
					Tools: []ToolDef{{Name: "", Description: "no name"}},
				},
			},
			want: "planner:",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.agent.Validate()
			if err == nil {
				t.Fatalf("want an error mentioning %q, got nil", tc.want)
			}
			if !strings.Contains(err.Error(), tc.want) {
				t.Errorf("error = %q, want it to mention %q", err, tc.want)
			}
		})
	}
}

// A plan-execute parent has a strategy but no agents list. Both halves of that
// were wrong before: strategy was dropped because the slots did not count as
// sub-agents, and fixing that made an empty agents list appear.
func TestPlanExecuteEmitsStrategyWithoutAgentsList(t *testing.T) {
	cfg := planExecuteAgent().toConfig()

	if got := cfg["strategy"]; got != string(StrategyPlanExecute) {
		t.Errorf("strategy = %v, want %q", got, StrategyPlanExecute)
	}
	if _, ok := cfg["agents"]; ok {
		t.Errorf("agents must be absent when the parent has none, got %v", cfg["agents"])
	}
	if _, ok := cfg["planner"]; !ok {
		t.Error("planner slot missing")
	}
}

// The handoff worker contract, as MultiAgentCompiler dispatches and the swarm
// resolver reads it: five inputs in, {handoff: bool} out. The resolver throws
// unless handoff is a boolean, so the key is always present. The server sends
// active_agent as an index into [parent, sub-agents...] and result as either
// text or a list, so both are normalised before the predicate sees them.
func TestHandoffHandlerMapsTheContract(t *testing.T) {
	var seen HandoffState
	c := &OnCondition{
		Target: "billing",
		Condition: func(_ context.Context, s HandoffState) (bool, error) {
			seen = s
			return s.ActiveAgent == "triage" && s.Result != "", nil
		},
	}
	h := c.handoffHandler([]string{"swarm", "triage", "billing"})

	in := handoffIn{
		// A list, as the server sends when the last response was not plain
		// text; the predicate must still see a string.
		Result:       []any{"please", "bill me"},
		Conversation: "u: hi\na: hello",
		Context:      map[string]any{"k": "v"},
		ActiveAgent:  "1", // the server's index for "triage" (0 is the parent)
		ToolResults:  []any{"r1"},
	}
	out, err := h(context.Background(), in)
	if err != nil || !out.Handoff {
		t.Fatalf("expected handoff=true, got %+v %v", out, err)
	}
	if seen.Result != "please\nbill me" {
		t.Errorf("Result = %q, want the list flattened to text", seen.Result)
	}
	if seen.ActiveAgent != "triage" {
		t.Errorf("ActiveAgent = %q, want index 1 resolved to triage", seen.ActiveAgent)
	}
	if seen.Conversation != in.Conversation || seen.Context["k"] != "v" || len(seen.ToolResults) != 1 {
		t.Errorf("predicate saw %+v", seen)
	}

	out, _ = h(context.Background(), handoffIn{ActiveAgent: "2", Result: "x"})
	if out.Handoff {
		t.Error("expected handoff=false when the predicate declines")
	}
}

// Index 0 is the parent, and an index the list cannot resolve is passed
// through as text rather than silently becoming empty.
func TestResolveAgent(t *testing.T) {
	agents := []string{"swarm", "a", "b"}
	if got := resolveAgent("0", agents); got != "swarm" {
		t.Errorf("index 0 = %q, want the parent", got)
	}
	if got := resolveAgent(" 2 ", agents); got != "b" {
		t.Errorf("padded index = %q, want b", got)
	}
	if got := resolveAgent("7", agents); got != "7" {
		t.Errorf("out-of-range index = %q, want raw \"7\"", got)
	}
	if got := resolveAgent("billing", agents); got != "billing" {
		t.Errorf("non-index = %q, want passthrough", got)
	}
}

// textOf covers the shapes the server has been seen to send for result.
func TestTextOf(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{nil, ""},
		{"plain", "plain"},
		{[]any{}, ""},
		{[]any{"a", "b"}, "a\nb"},
		{[]any{map[string]any{"type": "text", "text": "hello"}}, "hello"},
		{[]any{map[string]any{"content": "c"}, "d"}, "c\nd"},
		{map[string]any{"text": "t"}, "t"},
		{42, "42"},
	}
	for _, tc := range cases {
		if got := textOf(tc.in); got != tc.want {
			t.Errorf("textOf(%#v) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

// An erroring predicate does not hand off, as Python's should_handoff swallows
// exceptions and returns False.
func TestHandoffHandlerErrorMeansNoHandoff(t *testing.T) {
	c := &OnCondition{Target: "x", Condition: func(context.Context, HandoffState) (bool, error) {
		return true, errors.New("broken")
	}}
	out, err := c.handoffHandler(nil)(context.Background(), handoffIn{})
	if err != nil || out.Handoff {
		t.Errorf("out=%+v err=%v; want handoff=false and no task error", out, err)
	}
}

// Registration collects every on_condition and nothing else.
func TestOnConditionsCollector(t *testing.T) {
	a := &Agent{
		Name: "a", Model: testModel,
		Handoffs: []HandoffCondition{
			&OnCondition{Target: "b", Condition: noHandoff},
			&OnTextMention{Target: "b", Text: "x"},
			&OnToolResult{Target: "b", ToolName: "t"},
		},
	}
	if got := a.onConditions(); len(got) != 1 || got[0].Target != "b" {
		t.Errorf("onConditions = %v, want the one OnCondition", got)
	}
}
