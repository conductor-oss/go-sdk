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
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

func TestPlannerContextValidation(t *testing.T) {
	planner := &Agent{Name: "p", Model: testModel, Instructions: "Plan."}
	cases := []struct {
		name    string
		agent   *Agent
		wantErr string
	}{
		{
			name: "requires plan-execute",
			agent: &Agent{Name: "a", Model: testModel,
				PlannerContext: []PlanContext{{Text: "x"}}},
			wantErr: "requires StrategyPlanExecute",
		},
		{
			name: "text and url together",
			agent: &Agent{Name: "a", Model: testModel, Strategy: StrategyPlanExecute, Planner: planner,
				PlannerContext: []PlanContext{{Text: "x", URL: "https://x"}}},
			wantErr: "exactly one of Text or URL",
		},
		{
			name: "neither text nor url",
			agent: &Agent{Name: "a", Model: testModel, Strategy: StrategyPlanExecute, Planner: planner,
				PlannerContext: []PlanContext{{}}},
			wantErr: "exactly one of Text or URL",
		},
		{
			name: "valid",
			agent: &Agent{Name: "a", Model: testModel, Strategy: StrategyPlanExecute, Planner: planner,
				PlannerContext: []PlanContext{{Text: "x"}, {URL: "https://x"}}},
		},
	}
	for _, c := range cases {
		err := c.agent.Validate()
		if c.wantErr == "" && err != nil {
			t.Errorf("%s: unexpected error %v", c.name, err)
		}
		if c.wantErr != "" && (err == nil || !strings.Contains(err.Error(), c.wantErr)) {
			t.Errorf("%s: err = %v, want %q", c.name, err, c.wantErr)
		}
	}
}

// Only what differs from the defaults travels, as Context.to_dict does.
func TestPlanContextConfig(t *testing.T) {
	if got := (PlanContext{URL: "https://x", MaxBytes: defaultContextMaxBytes}).config(); !reflect.DeepEqual(got, map[string]any{"url": "https://x"}) {
		t.Errorf("default max bytes must be omitted: %v", got)
	}
	got := (PlanContext{URL: "https://x", Optional: true, MaxBytes: 10, Headers: map[string]string{"A": "b"}}).config()
	want := map[string]any{"url": "https://x", "required": false, "maxBytes": 10, "headers": map[string]string{"A": "b"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("config = %v, want %v", got, want)
	}
}

func TestGateValidationAndHandler(t *testing.T) {
	if err := (&Agent{Name: "a", Model: testModel, Gate: TextGate{}}).Validate(); err == nil {
		t.Error("an empty TextGate validated")
	}
	if err := (&Agent{Name: "a", Model: testModel, Gate: GateFunc(nil)}).Validate(); err == nil {
		t.Error("a nil GateFunc validated")
	}

	// The worker contract: {"result"} in, {"decision"} out, and an error
	// means continue.
	stopOnDone := GateFunc(func(ctx context.Context, s GateState) (bool, error) {
		if s.Result == "boom" {
			return false, errors.New("gate broke")
		}
		return !strings.Contains(s.Result, "DONE"), nil
	})
	h := stopOnDone.gateHandler()
	for in, want := range map[string]string{"keep going": "continue", "all DONE": "stop", "boom": "continue"} {
		out, err := h(context.Background(), gateIn{Result: in})
		if err != nil || out.Decision != want {
			t.Errorf("gate(%q) = %v, %v; want %s", in, out, err, want)
		}
	}
	if got := stopOnDone.gateConfig("review"); got["taskName"] != "review_gate" {
		t.Errorf("gate document = %v", got)
	}
}

// The task definition mirrors the Python SDK's _default_task_def, and
// WithRetry-style settings replace the defaults.
func TestTaskDefDefaultsAndRetry(t *testing.T) {
	td := ToolDef{Name: "ping", Credentials: []string{"KEY"}}.taskDef()
	want := model.TaskDef{Name: "ping", RetryCount: 2, RetryLogic: "LINEAR_BACKOFF", RetryDelaySeconds: 2,
		TimeoutSeconds: 0, ResponseTimeoutSeconds: 10, TimeoutPolicy: "RETRY", RuntimeMetadata: []string{"KEY"}}
	if !reflect.DeepEqual(td, want) {
		t.Errorf("default task def = %+v, want %+v", td, want)
	}
	custom := ToolDef{Name: "flaky", RetryCount: Ptr(5), RetryDelaySeconds: Ptr(30), RetryPolicy: RetryExponentialBackoff}.taskDef()
	if custom.RetryCount != 5 || custom.RetryDelaySeconds != 30 || custom.RetryLogic != "EXPONENTIAL_BACKOFF" {
		t.Errorf("custom task def = %+v", custom)
	}
	if err := (ToolDef{Name: "x", RetryPolicy: "sometimes"}).Validate(); err == nil {
		t.Error("an unknown retry policy validated")
	}
}

// registerWorkers starts workers for the gate and for prefill tools that are
// not in Tools, and collects a task definition for each; registerTaskDefs then
// sends them, after the run has started.
func TestRegisterWorkersRegistersTaskDefsGateAndPrefill(t *testing.T) {
	var mu sync.Mutex
	registered := map[string]model.TaskDef{}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/token", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"token": "fake-token"})
	})
	mux.HandleFunc("/api/metadata/taskdefs", func(w http.ResponseWriter, r *http.Request) {
		var def model.TaskDef
		json.NewDecoder(r.Body).Decode(&def)
		mu.Lock()
		registered[def.Name] = def
		mu.Unlock()
		w.Write([]byte("{}"))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("[]")) })
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		mux.ServeHTTP(w, r)
	}))
	defer srv.Close()

	api := client.NewAPIClient(
		settings.NewAuthenticationSettings("key", "secret"),
		settings.NewHttpSettings(srv.URL+"/api"),
	)
	rt := NewRuntimeWithClient(api, Config{})
	defer rt.Shutdown()

	prefill := ToolDef{Name: "warm_up", InputSchema: map[string]any{"type": "object"},
		Handler: func(ctx context.Context, in echoIn) (string, error) { return "ok", nil }}
	agent := &Agent{
		Name: "gated", Model: testModel, Instructions: "Go.",
		Tools: []ToolDef{{Name: "ping", InputSchema: map[string]any{"type": "object"}, Credentials: []string{"KEY"},
			RetryCount: Ptr(1), RetryPolicy: RetryFixed,
			Handler: func(ctx context.Context, in echoIn) (string, error) { return "pong", nil }}},
		PrefillTools: []PrefillToolCall{Prefill(prefill, map[string]any{"city": "x"})},
		Gate:         GateFunc(func(context.Context, GateState) (bool, error) { return true, nil }),
	}
	if err := rt.registerWorkers(agent, nil); err != nil {
		t.Fatalf("registerWorkers: %v", err)
	}
	if len(registered) != 0 {
		t.Errorf("task definitions were registered before the run started: %v", registered)
	}
	if err := rt.registerTaskDefs(context.Background()); err != nil {
		t.Fatalf("registerTaskDefs: %v", err)
	}

	for _, name := range []string{"ping", "warm_up", "gated_gate"} {
		if !rt.started[workerKey{name: name}] {
			t.Errorf("no worker started for %s; started = %v", name, rt.started)
		}
	}
	mu.Lock()
	defer mu.Unlock()
	ping, ok := registered["ping"]
	if !ok {
		t.Fatalf("no task definition registered for ping; got %v", registered)
	}
	if ping.RetryCount != 1 || ping.RetryLogic != "FIXED" || !reflect.DeepEqual(ping.RuntimeMetadata, []string{"KEY"}) {
		t.Errorf("ping task def = %+v", ping)
	}
	if _, ok := registered["gated_gate"]; !ok {
		t.Errorf("no task definition for the gate worker")
	}
}
