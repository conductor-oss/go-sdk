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
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

// weatherPlan is the two-step chain the e2e suite runs: the second step takes
// the first step's whole output through a Ref.
func weatherPlan() *Plan {
	return &Plan{Steps: []Step{
		{ID: "weather", Operations: []Op{
			{Tool: "get_weather", Args: map[string]any{"city": "Oslo"}},
		}},
		{ID: "packing", DependsOn: []string{"weather"}, Operations: []Op{
			{Tool: "packing_advice", Args: map[string]any{"weather": Ref{StepID: "weather"}}},
		}},
	}}
}

func canonical(t *testing.T, v any) string {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	return string(b)
}

// The document for the weather plan, byte for byte what Java's Plan.toJson()
// and Python's Plan.to_dict() produce: no depends_on on a step without one,
// no parallel flag unless set, and a Ref as {"$ref": id}. json.Marshal sorts
// map keys, so the comparison is order-independent.
func TestPlanPayloadMatchesSiblingSDKs(t *testing.T) {
	want := `{"steps":[` +
		`{"id":"weather","operations":[{"args":{"city":"Oslo"},"tool":"get_weather"}]},` +
		`{"depends_on":["weather"],"id":"packing","operations":[` +
		`{"args":{"weather":{"$ref":"weather"}},"tool":"packing_advice"}]}]}`
	if got := canonical(t, weatherPlan().toPayload()); got != want {
		t.Errorf("payload mismatch\n got: %s\nwant: %s", got, want)
	}
}

// Every optional field, present: Refs nested in maps and slices, generate
// with a token cap and a Ref context, parallel, validation, on_success and
// on_failure. Each is spelled the way the server (and the other SDKs) expect.
func TestPlanPayloadOptionalFields(t *testing.T) {
	p := &Plan{
		Steps: []Step{
			{ID: "a", Operations: []Op{{Tool: "t1", Args: map[string]any{}}}},
			{ID: "b", DependsOn: []string{"a"}, Parallel: true, Operations: []Op{
				{Tool: "t2", Args: map[string]any{
					"nested": map[string]any{"in": Ref{StepID: "a"}},
					"list":   []any{"x", Ref{StepID: "a"}},
					"ptr":    &Ref{StepID: "a"},
				}},
				{Tool: "t3", Generate: &Generate{
					Instructions: "Pick the temperature.",
					OutputSchema: `{"temp_f": 0}`,
					MaxTokens:    64,
					Context:      Ref{StepID: "a"},
				}},
			}},
		},
		Validation: []Validation{{Tool: "check", Args: map[string]any{"r": Ref{StepID: "b"}}, SuccessCondition: "$.ok == true"}},
		OnSuccess:  []Action{{Tool: "notify"}},
		OnFailure:  []Action{{Tool: "alert", Args: map[string]any{"why": "failed"}}},
	}
	want := `{"on_failure":[{"args":{"why":"failed"},"tool":"alert"}],` +
		`"on_success":[{"tool":"notify"}],` +
		`"steps":[` +
		`{"id":"a","operations":[{"args":{},"tool":"t1"}]},` +
		`{"depends_on":["a"],"id":"b","operations":[` +
		`{"args":{"list":["x",{"$ref":"a"}],"nested":{"in":{"$ref":"a"}},"ptr":{"$ref":"a"}},"tool":"t2"},` +
		`{"generate":{"context":{"$ref":"a"},"instructions":"Pick the temperature.","max_tokens":64,"output_schema":"{\"temp_f\": 0}"},"tool":"t3"}],` +
		`"parallel":true}],` +
		`"validation":[{"args":{"r":{"$ref":"b"}},"success_condition":"$.ok == true","tool":"check"}]}`
	if got := canonical(t, p.toPayload()); got != want {
		t.Errorf("payload mismatch\n got: %s\nwant: %s", got, want)
	}
	if err := p.Validate(); err != nil {
		t.Errorf("Validate: %v", err)
	}
}

func TestPlanValidate(t *testing.T) {
	step := func(id string, ops ...Op) Step { return Step{ID: id, Operations: ops} }
	args := func(tool string, a map[string]any) Op { return Op{Tool: tool, Args: a} }

	cases := []struct {
		name string
		plan *Plan
		want string // substring of the error; "" means valid
	}{
		{"nil", nil, "plan is nil"},
		{"no steps", &Plan{}, "no steps"},
		{"missing id", &Plan{Steps: []Step{step("", args("t", nil))}}, "has no id"},
		{"duplicate id", &Plan{Steps: []Step{step("a", args("t", map[string]any{})), step("a", args("t", map[string]any{}))}}, "used twice"},
		{"unknown dependency", &Plan{Steps: []Step{{ID: "a", DependsOn: []string{"zz"}, Operations: []Op{args("t", map[string]any{})}}}}, `unknown step "zz"`},
		{"op without tool", &Plan{Steps: []Step{step("a", args("", map[string]any{}))}}, "op has no tool"},
		{"op neither args nor generate", &Plan{Steps: []Step{step("a", Op{Tool: "t"})}}, "exactly one of Args or Generate"},
		{"op both args and generate", &Plan{Steps: []Step{step("a", Op{Tool: "t", Args: map[string]any{}, Generate: &Generate{Instructions: "i", OutputSchema: "{}"}})}}, "exactly one of Args or Generate"},
		{"generate missing schema", &Plan{Steps: []Step{step("a", Op{Tool: "t", Generate: &Generate{Instructions: "i"}})}}, "needs Instructions and OutputSchema"},
		{"ref to unknown step", &Plan{Steps: []Step{step("a", args("t", map[string]any{"x": Ref{StepID: "nope"}}))}}, `Ref to unknown step "nope"`},
		{"ref nested in slice", &Plan{Steps: []Step{step("a", args("t", map[string]any{"x": []any{Ref{StepID: "nope"}}}))}}, `Ref to unknown step "nope"`},
		{"empty ref", &Plan{Steps: []Step{step("a", args("t", map[string]any{"x": Ref{}}))}}, "Ref has no step id"},
		{"generate context ref unknown", &Plan{Steps: []Step{step("a", Op{Tool: "t", Generate: &Generate{Instructions: "i", OutputSchema: "{}", Context: Ref{StepID: "nope"}}})}}, `unknown step "nope"`},
		{"validation without tool", &Plan{Steps: []Step{step("a", args("t", map[string]any{}))}, Validation: []Validation{{}}}, "validation 0 has no tool"},
		{"on_failure without tool", &Plan{Steps: []Step{step("a", args("t", map[string]any{}))}, OnFailure: []Action{{}}}, "on_failure action 0 has no tool"},
		{"valid weather plan", weatherPlan(), ""},
		{"forward ref is fine", &Plan{Steps: []Step{
			step("a", args("t", map[string]any{"later": Ref{StepID: "b"}})),
			step("b", args("t", map[string]any{})),
		}}, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.plan.Validate()
			switch {
			case tc.want == "" && err != nil:
				t.Errorf("want valid, got %v", err)
			case tc.want != "" && err == nil:
				t.Errorf("want error containing %q, got nil", tc.want)
			case tc.want != "" && !strings.Contains(err.Error(), tc.want):
				t.Errorf("want error containing %q, got %v", tc.want, err)
			}
		})
	}
}

// startRecorder answers /agent/start and keeps the body it was sent.
func startRecorder(t *testing.T) (*httptest.Server, func() map[string]any) {
	t.Helper()
	var mu sync.Mutex
	var body map[string]any
	mux := http.NewServeMux()
	mux.HandleFunc("/api/token", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"token": "fake-token"})
	})
	mux.HandleFunc("/api/agent/start", func(w http.ResponseWriter, r *http.Request) {
		var b map[string]any
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			t.Errorf("decode start payload: %v", err)
		}
		mu.Lock()
		body = b
		mu.Unlock()
		json.NewEncoder(w).Encode(map[string]any{"executionId": "exec-plan"})
	})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		mux.ServeHTTP(w, r)
	}))
	return srv, func() map[string]any {
		mu.Lock()
		defer mu.Unlock()
		return body
	}
}

func planRuntime(t *testing.T, url string) *Runtime {
	t.Helper()
	rt := NewRuntimeWithClient(client.NewAPIClient(
		settings.NewAuthenticationSettings("key", "secret"),
		settings.NewHttpSettings(url+"/api"),
	), Config{})
	t.Cleanup(rt.Shutdown)
	return rt
}

func planAgent(strategy Strategy) *Agent {
	return &Agent{
		Name: "trip", Model: testModel, Strategy: strategy,
		Instructions: "Carry out the plan.",
		Planner:      &Agent{Name: "planner", Model: testModel, Instructions: "unused"},
	}
}

// WithPlan lands on the start request as static_plan — the key the server
// reads ahead of the planner's output, and the one Java and Python send.
func TestStartWithPlanSendsStaticPlan(t *testing.T) {
	srv, sent := startRecorder(t)
	defer srv.Close()
	rt := planRuntime(t, srv.URL)

	h, err := rt.Start(context.Background(), planAgent(StrategyPlanExecute), "pack for Oslo", WithPlan(weatherPlan()))
	if err != nil {
		t.Fatalf("Start: %v", err)
	}
	if h.ExecutionID != "exec-plan" {
		t.Errorf("executionID = %q", h.ExecutionID)
	}
	body := sent()
	if body == nil {
		t.Fatal("server saw no start request")
	}
	// The plan travels as JSON either way, so compare through the wire form.
	if got, want := canonical(t, body["static_plan"]), canonical(t, weatherPlan().toPayload()); got != want {
		t.Errorf("static_plan\n got: %s\nwant: %s", got, want)
	}
	if p, _ := body["prompt"].(string); p != "pack for Oslo" {
		t.Errorf("prompt = %q", p)
	}
}

// Without WithPlan the request is exactly what it was before the option
// existed: no static_plan key, not even a null one.
func TestStartWithoutPlanOmitsStaticPlan(t *testing.T) {
	srv, sent := startRecorder(t)
	defer srv.Close()
	rt := planRuntime(t, srv.URL)

	if _, err := rt.Start(context.Background(), planAgent(StrategyPlanExecute), "hi"); err != nil {
		t.Fatalf("Start: %v", err)
	}
	if _, present := sent()["static_plan"]; present {
		t.Error("static_plan was sent without WithPlan")
	}
}

// A plan is only meaningful to StrategyPlanExecute; on any other strategy the
// server would ignore it and run the agent as if none were given. That is the
// silent failure this option exists to avoid, so it is an error up front.
func TestWithPlanRequiresPlanExecute(t *testing.T) {
	srv, sent := startRecorder(t)
	defer srv.Close()
	rt := planRuntime(t, srv.URL)

	_, err := rt.Run(context.Background(), planAgent(StrategyHandoff), "hi", WithPlan(weatherPlan()))
	if err == nil || !strings.Contains(err.Error(), "WithPlan requires StrategyPlanExecute") {
		t.Fatalf("want a strategy error, got %v", err)
	}
	if sent() != nil {
		t.Error("an invalid combination still reached the server")
	}
}

// An invalid plan is rejected before anything is sent.
func TestWithPlanValidatesBeforeStart(t *testing.T) {
	srv, sent := startRecorder(t)
	defer srv.Close()
	rt := planRuntime(t, srv.URL)

	bad := &Plan{Steps: []Step{{ID: "a", Operations: []Op{{Tool: "t", Args: map[string]any{"x": Ref{StepID: "missing"}}}}}}}
	_, err := rt.Start(context.Background(), planAgent(StrategyPlanExecute), "hi", WithPlan(bad))
	if err == nil || !strings.Contains(err.Error(), `Ref to unknown step "missing"`) {
		t.Fatalf("want a validation error, got %v", err)
	}
	if sent() != nil {
		t.Error("an invalid plan still reached the server")
	}
}
