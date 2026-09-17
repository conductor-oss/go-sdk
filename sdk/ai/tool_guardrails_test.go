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
	"reflect"
	"strings"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// A tool's guardrails run in the worker around the handler, as the Python
// worker's run_tool_task does: input guardrails over the arguments before the
// call, output guardrails over the return after it.

func runGuarded(t *testing.T, td ToolDef, input map[string]any) (map[string]any, error) {
	t.Helper()
	fn, err := toolExecutor(td)
	if err != nil {
		t.Fatal(err)
	}
	v, err := fn(&model.Task{InputData: input})
	if err != nil {
		return nil, err
	}
	res, err := model.GetTaskResultFromTaskExecutionOutput(&model.Task{}, v)
	if err != nil {
		t.Fatal(err)
	}
	return res.OutputData, nil
}

func echoTool(guards ...Guardrail) ToolDef {
	return ToolDef{Name: "echo", Guardrails: guards,
		Handler: func(ctx context.Context, in echoIn) (string, error) { return in.City, nil }}
}

func TestToolOutputRegexGuardrailBlocksTheResult(t *testing.T) {
	g := &RegexGuardrail{Patterns: []string{`[\w.+-]+@[\w-]+\.[\w.-]+`}, Mode: "block", Message: "Do not include email addresses."}
	g.Name, g.Position, g.OnFail = "no_email", PositionOutput, OnFailRetry

	got, err := runGuarded(t, echoTool(g), map[string]any{"city": "contact test@example.com for help"})
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]any{"error": "Output blocked by guardrail 'no_email': Do not include email addresses.", "blocked": true}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("blocked output = %v, want %v", got, want)
	}

	got, err = runGuarded(t, echoTool(g), map[string]any{"city": "no address here"})
	if err != nil || !reflect.DeepEqual(got, map[string]any{"result": "no address here"}) {
		t.Errorf("clean output = %v, %v; the guardrail must let it through", got, err)
	}
}

func TestToolOutputGuardrailFixSubstitutesTheResult(t *testing.T) {
	g := NewCustomGuardrail("force_json", func(_ context.Context, in GuardrailInput) (GuardrailResult, error) {
		if strings.HasPrefix(strings.TrimSpace(in.Content), "{") {
			return GuardrailResult{Passed: true}, nil
		}
		return GuardrailResult{Passed: false, Message: "Output must be JSON.", FixedOutput: `{"fixed": true}`}, nil
	})
	g.Position, g.OnFail = PositionOutput, OnFailFix

	got, err := runGuarded(t, echoTool(g), map[string]any{"city": "plain text"})
	if err != nil || !reflect.DeepEqual(got, map[string]any{"result": `{"fixed": true}`}) {
		t.Errorf("fixed output = %v, %v", got, err)
	}
}

func TestToolInputGuardrailRaiseFailsTheTaskBeforeTheHandler(t *testing.T) {
	var ran bool
	g := NewCustomGuardrail("no_sql_injection", func(_ context.Context, in GuardrailInput) (GuardrailResult, error) {
		if strings.Contains(strings.ToUpper(in.Content), "DROP TABLE") {
			return GuardrailResult{Passed: false, Message: "SQL injection blocked."}, nil
		}
		return GuardrailResult{Passed: true}, nil
	})
	g.Position, g.OnFail = PositionInput, OnFailRaise
	td := ToolDef{Name: "safe_query", Guardrails: []Guardrail{g},
		Handler: func(ctx context.Context, in echoIn) (string, error) { ran = true; return "ran", nil }}

	_, err := runGuarded(t, td, map[string]any{"city": "DROP TABLE users"})
	if err == nil || !strings.Contains(err.Error(), "no_sql_injection") || !strings.Contains(err.Error(), "SQL injection blocked.") {
		t.Errorf("err = %v, want the guardrail's refusal", err)
	}
	if ran {
		t.Error("the handler ran despite the input guardrail")
	}

	// A non-raise input failure returns the blocked marker instead of running the tool.
	g.OnFail = OnFailRetry
	got, err := runGuarded(t, td, map[string]any{"city": "DROP TABLE users"})
	if err != nil || got["blocked"] != true || !strings.HasPrefix(got["error"].(string), "Blocked by guardrail 'no_sql_injection':") {
		t.Errorf("blocked input = %v, %v", got, err)
	}
	if ran {
		t.Error("the handler ran despite the input guardrail")
	}
}

func TestToolGuardrailsOnlyTheServerCanRunAreLeftAlone(t *testing.T) {
	// An LLM guardrail needs a model and an external custom guardrail's worker
	// runs elsewhere; the worker skips both rather than guessing a verdict.
	llm := &LLMGuardrail{Policy: "No secrets."}
	llm.Position = PositionOutput
	external := &CustomGuardrail{External: true}
	external.Name, external.Position = "elsewhere", PositionOutput

	got, err := runGuarded(t, echoTool(llm, external), map[string]any{"city": "password is hunter2"})
	if err != nil || !reflect.DeepEqual(got, map[string]any{"result": "password is hunter2"}) {
		t.Errorf("output = %v, %v; server-side guardrails must not change the result here", got, err)
	}
}
