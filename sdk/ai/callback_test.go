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
	"reflect"
	"testing"
)

// Only the set hooks serialize, in position order, each as {position,taskName}.
func TestCallbackConfigsOnlySetPositions(t *testing.T) {
	noop := func(context.Context, CallbackInput) (map[string]any, error) { return nil, nil }
	a := &Agent{Name: "obs", Model: testModel, Callbacks: &Callbacks{
		OnModelEnd: noop, OnToolStart: noop, // deliberately out of field order
	}}
	got := a.callbackConfigs()
	want := []any{
		map[string]any{"position": "after_model", "taskName": "obs_after_model"},
		map[string]any{"position": "before_tool", "taskName": "obs_before_tool"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("callbackConfigs = %v, want %v", got, want)
	}
	if (&Agent{Name: "x"}).callbackConfigs() != nil {
		t.Error("no callbacks should serialize to nil")
	}
}

// Each set hook becomes a worker named "<agent>_<position>".
func TestCallbackToolsNaming(t *testing.T) {
	noop := func(context.Context, CallbackInput) (map[string]any, error) { return nil, nil }
	a := &Agent{Name: "obs", Model: testModel, Callbacks: &Callbacks{OnAgentStart: noop, OnToolEnd: noop}}
	var names []string
	for _, td := range a.callbackTools() {
		names = append(names, td.Name)
		if td.Handler == nil {
			t.Errorf("callback tool %s has no handler", td.Name)
		}
	}
	if !reflect.DeepEqual(names, []string{"obs_before_agent", "obs_after_tool"}) {
		t.Errorf("callback tool names = %v", names)
	}
}

// The worker contract: input carries messages and llm_result; a nil return or
// an error becomes {} (continue); a non-empty map passes through as the
// override.
func TestCallbackHandlerContract(t *testing.T) {
	var seen CallbackInput
	observe := func(_ context.Context, in CallbackInput) (map[string]any, error) {
		seen = in
		return nil, nil
	}
	h := callbackHandler(observe)
	out, err := h(context.Background(), callbackIn{
		Position:  "after_model",
		AgentName: "obs",
		LLMResult: []any{map[string]any{"text": "the answer"}},
		ToolCalls: []any{},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(out, map[string]any{}) {
		t.Errorf("nil return should map to {}, got %v", out)
	}
	if seen.Position != "after_model" || seen.AgentName != "obs" {
		t.Errorf("hook did not receive position/agent: %+v", seen)
	}
	if _, ok := seen.LLMResult.([]any); !ok {
		t.Errorf("llm_result should bind as decoded JSON, got %T", seen.LLMResult)
	}

	override := callbackHandler(func(context.Context, CallbackInput) (map[string]any, error) {
		return map[string]any{"stop": true}, nil
	})
	out, _ = override(context.Background(), callbackIn{})
	if !reflect.DeepEqual(out, map[string]any{"stop": true}) {
		t.Errorf("override = %v", out)
	}

	// An error never blocks the run: it becomes {}.
	broken := callbackHandler(func(context.Context, CallbackInput) (map[string]any, error) {
		return nil, errors.New("boom")
	})
	if out, err := broken(context.Background(), callbackIn{}); err != nil || len(out) != 0 {
		t.Errorf("errored hook = %v, %v; want {}, nil", out, err)
	}
}

// registerWorkers starts a worker for each set callback.
func TestCallbackWorkersRegister(t *testing.T) {
	rt := NewRuntimeWithClient(nil, Config{})
	noop := func(context.Context, CallbackInput) (map[string]any, error) { return nil, nil }
	agent := &Agent{
		Name: "obs", Model: testModel, Instructions: "Go.",
		Callbacks: &Callbacks{OnModelStart: noop, OnModelEnd: noop},
	}
	if err := rt.registerWorkers(agent); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"obs_before_model", "obs_after_model"} {
		if !rt.started[name] {
			t.Errorf("no worker started for %s; started=%v", name, rt.started)
		}
	}
}
