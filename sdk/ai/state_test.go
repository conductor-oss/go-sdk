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
	"sync"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// run drives a handler through the dispatcher the way a worker would, and
// returns the output the server would receive.
func runTool(t *testing.T, td ToolDef, input map[string]any) map[string]any {
	t.Helper()
	fn, err := toolExecutor(td)
	if err != nil {
		t.Fatal(err)
	}
	v, err := fn(&model.Task{InputData: input})
	if err != nil {
		t.Fatal(err)
	}
	res, err := model.GetTaskResultFromTaskExecutionOutput(&model.Task{}, v)
	if err != nil {
		t.Fatal(err)
	}
	return res.OutputData
}

// The server delivers the execution's state under _agent_state. It is not a
// tool argument, so the handler's input must not see it and the tool must read
// it through the state API instead.
func TestAgentStateIsReadableAndNotAnArgument(t *testing.T) {
	var gotCity string
	var gotTenant any
	var gotAll map[string]any
	td := ToolDef{Name: "x", Handler: func(ctx context.Context, in echoIn) (string, error) {
		gotCity = in.City
		gotTenant, _ = StateValue(ctx, "tenant")
		gotAll = State(ctx)
		return "ok", nil
	}}
	runTool(t, td, map[string]any{
		"city":          "Paris",
		agentStateKey:   map[string]any{"tenant": "acme", "region": "eu"},
		"unrelated_arg": 1,
	})

	if gotCity != "Paris" {
		t.Errorf("the tool's own argument did not bind: %q", gotCity)
	}
	if gotTenant != "acme" {
		t.Errorf("StateValue(tenant) = %v, want acme", gotTenant)
	}
	if !reflect.DeepEqual(gotAll, map[string]any{"tenant": "acme", "region": "eu"}) {
		t.Errorf("State() = %v", gotAll)
	}
}

// What the handler writes travels back under _state_updates, and only what it
// wrote: the state it merely read must not be echoed.
func TestSetStateTravelsBackAsUpdates(t *testing.T) {
	td := ToolDef{Name: "x", Handler: func(ctx context.Context, in echoIn) (map[string]any, error) {
		if _, ok := StateValue(ctx, "tenant"); !ok {
			t.Error("delivered state not readable")
		}
		if err := SetState(ctx, "ticket_id", "T-1"); err != nil {
			t.Fatal(err)
		}
		return map[string]any{"done": true}, nil
	}}
	out := runTool(t, td, map[string]any{"city": "x", agentStateKey: map[string]any{"tenant": "acme"}})

	if out["done"] != true {
		t.Errorf("the handler's own output was lost: %v", out)
	}
	want := map[string]any{"ticket_id": "T-1"}
	if !reflect.DeepEqual(out[stateUpdatesKey], want) {
		t.Errorf("%s = %v, want %v (only what was written)", stateUpdatesKey, out[stateUpdatesKey], want)
	}
}

// A handler that writes nothing must leave its output byte for byte as it was,
// so adding the state API cannot change any existing tool's result.
func TestNoWritesLeavesOutputUntouched(t *testing.T) {
	td := ToolDef{Name: "x", Handler: func(ctx context.Context, in echoIn) (string, error) {
		State(ctx)                // reading is not writing
		StateValue(ctx, "tenant") //nolint:errcheck // no error to check
		return "plain", nil
	}}
	out := runTool(t, td, map[string]any{"city": "x", agentStateKey: map[string]any{"tenant": "acme"}})
	if !reflect.DeepEqual(out, map[string]any{"result": "plain"}) {
		t.Errorf("output = %v, want the untouched scalar wrapper", out)
	}
}

// A struct return still carries its fields when state is written; the output is
// widened to an object rather than replaced.
//
// The json tags are load-bearing, not decoration: a tagged struct needs no
// snake-casing, so it reaches the merge as a struct and exercises the JSON
// round-trip that widens it. Drop them and the value arrives already converted
// to a map, taking the other branch — which TestStateUpdatesOnAnUntaggedStruct
// covers.
func TestStateUpdatesPreserveAStructResult(t *testing.T) {
	type out struct {
		City string `json:"city"`
		Hits int    `json:"hits"`
	}
	td := ToolDef{Name: "x", Handler: func(ctx context.Context, in echoIn) (out, error) {
		if err := SetState(ctx, "seen", true); err != nil {
			t.Fatal(err)
		}
		return out{City: in.City, Hits: 2}, nil
	}}
	got := runTool(t, td, map[string]any{"city": "Paris"})

	if got["city"] != "Paris" || got["hits"] != float64(2) {
		t.Errorf("struct fields lost: %v", got)
	}
	if !reflect.DeepEqual(got[stateUpdatesKey], map[string]any{"seen": true}) {
		t.Errorf("%s = %v", stateUpdatesKey, got[stateUpdatesKey])
	}
}

// A tool call with no state delivered still works, and a write from it is still
// reported, so a tool need not know whether the server sent anything.
func TestStateAbsentFromTheTask(t *testing.T) {
	td := ToolDef{Name: "x", Handler: func(ctx context.Context, in echoIn) (map[string]any, error) {
		if v, ok := StateValue(ctx, "anything"); ok {
			t.Errorf("found %v in empty state", v)
		}
		return map[string]any{"ok": true}, SetState(ctx, "first", 1)
	}}
	out := runTool(t, td, map[string]any{"city": "x"})
	if !reflect.DeepEqual(out[stateUpdatesKey], map[string]any{"first": float64(1)}) {
		t.Errorf("%s = %v", stateUpdatesKey, out[stateUpdatesKey])
	}
}

// Outside a tool handler the calls must say so rather than discard the write,
// which is the failure the CLI runner's ContextKey used to have.
func TestStateOutsideAToolHandler(t *testing.T) {
	ctx := context.Background()
	if _, ok := StateValue(ctx, "k"); ok {
		t.Error("StateValue reported a value with no task")
	}
	if got := State(ctx); len(got) != 0 {
		t.Errorf("State() = %v, want empty", got)
	}
	if err := SetState(ctx, "k", "v"); !errors.Is(err, ErrNoTaskContext) {
		t.Errorf("SetState err = %v, want ErrNoTaskContext", err)
	}
}

// A handler may use goroutines, so concurrent writes must not race.
func TestStateIsSafeForConcurrentWrites(t *testing.T) {
	td := ToolDef{Name: "x", Handler: func(ctx context.Context, in echoIn) (string, error) {
		var wg sync.WaitGroup
		for i := 0; i < 32; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				_ = SetState(ctx, "k", i)
				_ = State(ctx)
			}(i)
		}
		wg.Wait()
		return "ok", nil
	}}
	out := runTool(t, td, map[string]any{"city": "x"})
	updates, ok := out[stateUpdatesKey].(map[string]any)
	if !ok || len(updates) != 1 {
		t.Errorf("%s = %v", stateUpdatesKey, out[stateUpdatesKey])
	}
}

// CLIConfig.ContextKey was serialized and ignored before the state API existed,
// which is what made the gap concrete. It must now save the command's output
// where a later tool can read it, as Python's runner does.
func TestCLIContextKeySavesOutputToState(t *testing.T) {
	cfg := &CLIConfig{AllowedCommands: []string{"echo"}}
	td := ToolDef{Name: "run_command", Handler: cfg.runCommand}
	out := runTool(t, td, map[string]any{
		"command":     "echo",
		"args":        []any{"hello"},
		"context_key": "greeting",
	})
	if !reflect.DeepEqual(out[stateUpdatesKey], map[string]any{"greeting": "hello"}) {
		t.Errorf("%s = %v, want the trimmed stdout under greeting", stateUpdatesKey, out[stateUpdatesKey])
	}
}

// Without the key the command still runs and writes no state.
func TestCLIWithoutContextKeyWritesNoState(t *testing.T) {
	cfg := &CLIConfig{AllowedCommands: []string{"echo"}}
	td := ToolDef{Name: "run_command", Handler: cfg.runCommand}
	out := runTool(t, td, map[string]any{"command": "echo", "args": []any{"hi"}})
	if _, wrote := out[stateUpdatesKey]; wrote {
		t.Errorf("wrote state without a context key: %v", out[stateUpdatesKey])
	}
}

// The other branch: an untagged struct is snake-cased into a map before the
// merge, so the state updates are added to a map that already exists.
func TestStateUpdatesOnAnUntaggedStruct(t *testing.T) {
	type out struct {
		City  string
		TempF int
	}
	td := ToolDef{Name: "x", Handler: func(ctx context.Context, in echoIn) (out, error) {
		return out{City: in.City, TempF: 72}, SetState(ctx, "seen", true)
	}}
	got := runTool(t, td, map[string]any{"city": "Paris"})

	if got["city"] != "Paris" || got["temp_f"] != float64(72) {
		t.Errorf("snake-cased fields lost or misnamed: %v", got)
	}
	if !reflect.DeepEqual(got[stateUpdatesKey], map[string]any{"seen": true}) {
		t.Errorf("%s = %v", stateUpdatesKey, got[stateUpdatesKey])
	}
}
