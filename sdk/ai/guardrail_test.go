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

// The 09_guardrails fixture sets most fields, so it proves the emitted shape
// but leaves these paths uncovered: the mode default, the two conditional
// keys, and the external variant.
func TestGuardrailDefaults(t *testing.T) {
	// Only Patterns set: everything else comes from a default.
	cfg := (&RegexGuardrail{Patterns: []string{`\d{16}`}}).guardrailConfig()

	if got := cfg["name"]; got != defaultRegexName {
		t.Errorf("name = %v, want %q", got, defaultRegexName)
	}
	if got := cfg["mode"]; got != "block" {
		t.Errorf("mode = %v, want block", got)
	}
	if got := cfg["position"]; got != PositionOutput {
		t.Errorf("position = %v, want output", got)
	}
	if got := cfg["onFail"]; got != OnFailRaise {
		t.Errorf("onFail = %v, want raise", got)
	}
	if got := cfg["maxRetries"]; got != defaultMaxRetries {
		t.Errorf("maxRetries = %v, want %d", got, defaultMaxRetries)
	}
	// Absent rather than empty: an empty message would be a key the other
	// SDKs do not send.
	if _, ok := cfg["message"]; ok {
		t.Errorf("message should be omitted when unset, got %v", cfg["message"])
	}

	llm := (&LLMGuardrail{Model: testModel, Policy: "No advice."}).guardrailConfig()
	if got := llm["name"]; got != defaultLLMName {
		t.Errorf("llm name = %v, want %q", got, defaultLLMName)
	}
	if _, ok := llm["maxTokens"]; ok {
		t.Errorf("maxTokens should be omitted when zero, got %v", llm["maxTokens"])
	}
}

// A custom guardrail is worker-backed, so its name is also the task name the
// server dispatches to.
func TestCustomGuardrailTaskName(t *testing.T) {
	custom := &CustomGuardrail{
		guardrailBase: guardrailBase{Name: "no_secrets"},
		Check: func(context.Context, GuardrailInput) (GuardrailResult, error) {
			return GuardrailResult{Passed: true}, nil
		},
	}
	cfg := custom.guardrailConfig()
	if cfg["guardrailType"] != "custom" {
		t.Errorf("guardrailType = %v, want custom", cfg["guardrailType"])
	}
	if cfg["taskName"] != "no_secrets" || cfg["name"] != "no_secrets" {
		t.Errorf("taskName and name must agree: %v / %v", cfg["taskName"], cfg["name"])
	}

	// External changes only the type: the task name still identifies it, but
	// nothing here registers a worker for it.
	ext := &CustomGuardrail{guardrailBase: guardrailBase{Name: "remote"}, External: true}
	if got := ext.guardrailConfig()["guardrailType"]; got != "external" {
		t.Errorf("guardrailType = %v, want external", got)
	}
}

func TestGuardrailValidation(t *testing.T) {
	cases := []struct {
		name string
		g    Guardrail
		want string
	}{
		{
			name: "regex with no patterns",
			g:    &RegexGuardrail{},
			want: "no patterns",
		},
		{
			name: "invalid mode",
			g:    &RegexGuardrail{Patterns: []string{"x"}, Mode: "permit"},
			want: "must be block or allow",
		},
		{
			name: "invalid position",
			g: &RegexGuardrail{
				guardrailBase: guardrailBase{Position: "middle"},
				Patterns:      []string{"x"},
			},
			want: "invalid position",
		},
		{
			name: "invalid onFail",
			g: &RegexGuardrail{
				guardrailBase: guardrailBase{OnFail: "shrug"},
				Patterns:      []string{"x"},
			},
			want: "invalid onFail",
		},
		{
			// The server rejects this pairing: there is no output to escalate
			// when the guardrail runs on the way in.
			name: "human escalation on input",
			g: &RegexGuardrail{
				guardrailBase: guardrailBase{Position: PositionInput, OnFail: OnFailHuman},
				Patterns:      []string{"x"},
			},
			want: "only valid at position=output",
		},
		{
			name: "llm with no model",
			g:    &LLMGuardrail{Policy: "p"},
			want: "no model",
		},
		{
			name: "llm with no policy",
			g:    &LLMGuardrail{Model: testModel},
			want: "no policy",
		},
		{
			name: "custom with no name",
			g:    &CustomGuardrail{Check: func(context.Context, GuardrailInput) (GuardrailResult, error) { return GuardrailResult{}, nil }},
			want: "needs a Name",
		},
		{
			name: "custom with no check",
			g:    &CustomGuardrail{guardrailBase: guardrailBase{Name: "x"}},
			want: "no Check function",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := &Agent{Name: "a", Model: testModel, Guardrails: []Guardrail{tc.g}}
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

// A valid set must pass, or the checks above would be worthless.
func TestValidGuardrailsPass(t *testing.T) {
	if err := guardedAgent().Validate(); err != nil {
		t.Errorf("the fixture agent must validate: %v", err)
	}
}

// The worker contract, as Python's GuardrailEntry returns it and the server's
// normalize step reads it. The iteration arrives as JSON does it: a float64.
func TestGuardrailHandlerMapsTheContract(t *testing.T) {
	verdict := GuardrailResult{Passed: true}
	var gotContent string
	g := &CustomGuardrail{
		guardrailBase: guardrailBase{Name: "no_secrets", OnFail: OnFailRetry, MaxRetries: 2},
		Check: func(_ context.Context, in GuardrailInput) (GuardrailResult, error) {
			gotContent = in.Content
			return verdict, nil
		},
	}
	h := g.guardrailHandler()
	ctx := context.Background()

	out, err := h(ctx, guardrailIn{Content: "hello", Iteration: float64(1)})
	if err != nil {
		t.Fatal(err)
	}
	if gotContent != "hello" {
		t.Errorf("content = %q, want hello", gotContent)
	}
	if !out.Passed || out.OnFail != "pass" || out.ShouldContinue || out.FixedOutput != nil {
		t.Errorf("pass verdict = %+v, want passed with on_fail=pass", out)
	}

	// Under the retry budget the failure asks for another turn.
	verdict = GuardrailResult{Passed: false, Message: "contains a key"}
	out, _ = h(ctx, guardrailIn{Content: "x", Iteration: float64(1)})
	if out.Passed || out.OnFail != "retry" || !out.ShouldContinue ||
		out.Message != "contains a key" || out.GuardrailName != "no_secrets" {
		t.Errorf("retry verdict = %+v", out)
	}

	// At the budget it escalates to raise, as Python does client-side too.
	out, _ = h(ctx, guardrailIn{Content: "x", Iteration: float64(2)})
	if out.OnFail != "raise" || out.ShouldContinue {
		t.Errorf("escalated verdict = %+v, want on_fail=raise", out)
	}

	// Non-text content is handed over as JSON, not Go's %v rendering.
	h(ctx, guardrailIn{Content: map[string]any{"a": 1}})
	if gotContent != `{"a":1}` {
		t.Errorf("object content = %q, want JSON", gotContent)
	}
}

// A fix needs something to substitute; without it the guardrail raises.
func TestGuardrailHandlerFixNeedsOutput(t *testing.T) {
	verdict := GuardrailResult{Passed: false, Message: "redacted", FixedOutput: "[redacted]"}
	g := &CustomGuardrail{
		guardrailBase: guardrailBase{Name: "redact", OnFail: OnFailFix},
		Check:         func(context.Context, GuardrailInput) (GuardrailResult, error) { return verdict, nil },
	}
	h := g.guardrailHandler()

	out, _ := h(context.Background(), guardrailIn{Content: "secret"})
	if out.OnFail != "fix" || out.FixedOutput == nil || *out.FixedOutput != "[redacted]" {
		t.Errorf("fix verdict = %+v, want on_fail=fix with the fixed output", out)
	}

	verdict.FixedOutput = ""
	out, _ = h(context.Background(), guardrailIn{Content: "secret"})
	if out.OnFail != "raise" || out.FixedOutput != nil {
		t.Errorf("fix without output = %+v, want on_fail=raise", out)
	}
}

// A Check that errors fails the content rather than passing it, and the
// verdict carries the error so the failure reason says what broke.
func TestGuardrailHandlerErrorFails(t *testing.T) {
	g := &CustomGuardrail{
		guardrailBase: guardrailBase{Name: "flaky"},
		Check: func(context.Context, GuardrailInput) (GuardrailResult, error) {
			return GuardrailResult{Passed: true}, errors.New("boom")
		},
	}
	out, err := g.guardrailHandler()(context.Background(), guardrailIn{Content: "x"})
	if err != nil {
		t.Fatalf("the worker must not fail the task: %v", err)
	}
	if out.Passed || out.OnFail != "raise" || !strings.Contains(out.Message, "boom") {
		t.Errorf("error verdict = %+v", out)
	}
}

func TestIterationOf(t *testing.T) {
	cases := map[string]struct {
		in   any
		want int
	}{
		"json number": {float64(3), 3},
		"string":      {" 4 ", 4},
		"bool":        {true, 0},
		"nil":         {nil, 0},
	}
	for name, tc := range cases {
		if got := iterationOf(tc.in); got != tc.want {
			t.Errorf("%s: iterationOf(%v) = %d, want %d", name, tc.in, got, tc.want)
		}
	}
}

// Only guardrails this process can serve get a worker: custom ones with a
// Check, from the agent and from its tools, in declaration order.
func TestCustomGuardrailsCollector(t *testing.T) {
	ok := func(context.Context, GuardrailInput) (GuardrailResult, error) {
		return GuardrailResult{Passed: true}, nil
	}
	a := &Agent{
		Name: "a", Model: testModel,
		Guardrails: []Guardrail{
			&RegexGuardrail{Patterns: []string{"x"}},
			&CustomGuardrail{guardrailBase: guardrailBase{Name: "first"}, Check: ok},
			&CustomGuardrail{guardrailBase: guardrailBase{Name: "remote"}, External: true},
		},
		Tools: []ToolDef{{
			Name: "lookup",
			Guardrails: []Guardrail{
				&LLMGuardrail{Model: testModel, Policy: "p"},
				&CustomGuardrail{guardrailBase: guardrailBase{Name: "second"}, Check: ok},
			},
		}},
	}
	var names []string
	for _, g := range a.customGuardrails() {
		names = append(names, g.Name)
	}
	if got := strings.Join(names, ","); got != "first,second" {
		t.Errorf("customGuardrails = %s, want first,second", got)
	}
}

// The server sends more than the content; a check must be able to see it.
func TestGuardrailInputCarriesIterationAndToolCalls(t *testing.T) {
	var got GuardrailInput
	g := &CustomGuardrail{
		guardrailBase: guardrailBase{Name: "sees_all"},
		Check: func(_ context.Context, in GuardrailInput) (GuardrailResult, error) {
			got = in
			return GuardrailResult{Passed: true}, nil
		},
	}
	calls := []any{map[string]any{"name": "lookup", "arguments": map[string]any{"q": "x"}}}
	if _, err := g.guardrailHandler()(context.Background(), guardrailIn{
		Content: map[string]any{"k": "v"}, Iteration: "2", ToolCalls: calls,
	}); err != nil {
		t.Fatal(err)
	}
	if got.Content != `{"k":"v"}` {
		t.Errorf("Content = %q, want non-text content rendered as JSON", got.Content)
	}
	if got.Iteration != 2 {
		t.Errorf("Iteration = %d, want 2 parsed from a numeric string", got.Iteration)
	}
	if len(got.ToolCalls) != 1 {
		t.Errorf("ToolCalls = %v, want the one call passed through", got.ToolCalls)
	}
	// Absent tool calls stay nil rather than becoming an empty slice, so a
	// check can tell "none sent" from "sent, none made".
	g.guardrailHandler()(context.Background(), guardrailIn{Content: "x"})
	if got.ToolCalls != nil {
		t.Errorf("ToolCalls = %#v, want nil when the server sends none", got.ToolCalls)
	}
}

func TestNewCustomGuardrailSetsNameAndDefaults(t *testing.T) {
	check := func(context.Context, GuardrailInput) (GuardrailResult, error) {
		return GuardrailResult{Passed: true}, nil
	}
	g := NewCustomGuardrail("no_secrets", check)
	if g.Name != "no_secrets" || g.Check == nil {
		t.Fatalf("constructor did not set Name/Check: %+v", g)
	}
	if err := g.validateGuardrail(); err != nil {
		t.Errorf("a constructed guardrail must validate as-is, got %v", err)
	}
	cfg := g.guardrailConfig()
	// position and onFail are emitted as their typed constants, not strings.
	if cfg["taskName"] != "no_secrets" || cfg["position"] != PositionOutput || cfg["onFail"] != OnFailRaise {
		t.Errorf("defaults not applied on the wire: %v", cfg)
	}
}
