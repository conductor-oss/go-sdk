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
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

type echoIn struct {
	City string `json:"city"`
}

// steps records what the fake server observed, so the test asserts on the whole
// exchange rather than only the final value.
type steps struct {
	mu sync.Mutex
	s  []string
}

func (r *steps) add(v string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.s = append(r.s, v)
}

func (r *steps) all() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return strings.Join(r.s, " ")
}

// fakeConductor accepts /agent/start, hands the worker exactly one tool task,
// then reports COMPLETED. That is enough to close the loop the runtime depends
// on: serialize, start, dispatch a task into the Go handler, report its output.
//
// Every response declares application/json. Without it Go sniffs the body as
// text/plain and the SDK's decoder tries to assign a string to a struct.
func fakeConductor(t *testing.T, rec *steps) *httptest.Server {
	t.Helper()
	// dispatched: the task has been handed to a worker.
	// resultSeen: the worker has posted its result back.
	//
	// The run only completes once the result is in, mirroring a real server.
	// Completing at dispatch instead makes Run race the worker: it can return
	// before the tool has finished, which passes or fails depending on load.
	var mu sync.Mutex
	dispatched, resultSeen := false, false

	mux := http.NewServeMux()

	mux.HandleFunc("/api/token", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"token": "fake-token"})
	})

	mux.HandleFunc("/api/agent/start", func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("decode start payload: %v", err)
		}
		cfg, _ := body["agentConfig"].(map[string]any)
		name, _ := cfg["name"].(string)
		tools, _ := cfg["tools"].([]any)
		rec.add("start:" + name)
		if len(tools) != 1 {
			t.Errorf("server saw %d tools, want 1", len(tools))
		}
		if p, _ := body["prompt"].(string); p == "" {
			t.Error("start payload carried no prompt")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"executionId":     "exec-1",
			"requiredWorkers": []string{"get_weather"},
		})
	})

	mux.HandleFunc("/api/agent/exec-1/status", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		done := resultSeen
		mu.Unlock()
		status := "RUNNING"
		if done {
			status = "COMPLETED"
		}
		json.NewEncoder(w).Encode(map[string]any{
			"status": status,
			// Nested, as the real server returns it — see resultFrom.
			"output": map[string]any{
				"result":       "72F and sunny in San Francisco",
				"finishReason": "STOP",
			},
		})
	})

	mux.HandleFunc("/api/tasks/poll/batch/get_weather", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		if dispatched {
			json.NewEncoder(w).Encode([]any{})
			return
		}
		dispatched = true
		rec.add("poll:get_weather")
		json.NewEncoder(w).Encode([]map[string]any{{
			"taskId":             "t1",
			"taskDefName":        "get_weather",
			"workflowInstanceId": "exec-1",
			"inputData":          map[string]any{"city": "San Francisco"},
		}})
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// The task update lands here; capture what the worker produced.
		var body map[string]any
		json.NewDecoder(r.Body).Decode(&body)
		if out, ok := body["outputData"].(map[string]any); ok {
			if city, ok := out["city"].(string); ok {
				rec.add("result:" + city)
				mu.Lock()
				resultSeen = true
				mu.Unlock()
			}
		}
		w.Write([]byte("{}"))
	})

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		mux.ServeHTTP(w, r)
	}))
}

// The whole path a tools example depends on: define an agent with a Go tool,
// Run it, and have the server dispatch that tool back to this process.
func TestRunEndToEnd(t *testing.T) {
	rec := &steps{}
	srv := fakeConductor(t, rec)
	defer srv.Close()

	api := client.NewAPIClient(
		settings.NewAuthenticationSettings("key", "secret"),
		settings.NewHttpSettings(srv.URL+"/api"),
	)
	rt := NewRuntimeWithClient(api, Config{
		WorkerPollInterval: 20 * time.Millisecond,
		StatusPollInterval: 20 * time.Millisecond,
	})
	defer rt.Shutdown()

	agent := &Agent{
		Name: "weather_bot", Model: testModel, Instructions: "Use tools.",
		Tools: []ToolDef{{
			Name:        "get_weather",
			Description: "Get the current weather for a city",
			InputSchema: map[string]any{"type": "object"},
			ToolType:    ToolTypeWorker,
			Handler: func(ctx context.Context, in echoIn) (map[string]any, error) {
				return map[string]any{"city": in.City, "tempF": 72}, nil
			},
		}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	res, err := rt.Run(ctx, agent, "What is the weather in San Francisco?")
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if res.Status != StatusCompleted {
		t.Errorf("status = %v, want %v", res.Status, StatusCompleted)
	}
	if res.Output != "72F and sunny in San Francisco" {
		t.Errorf("output = %q", res.Output)
	}
	if res.ExecutionID != "exec-1" {
		t.Errorf("executionID = %q", res.ExecutionID)
	}

	got := rec.all()
	for _, want := range []string{"start:weather_bot", "poll:get_weather", "result:San Francisco"} {
		if !strings.Contains(got, want) {
			t.Errorf("missing step %q; saw: %s", want, got)
		}
	}
}

// A tool whose handler has the wrong shape is a programming error and must be
// reported at registration, not when the first task arrives.
func TestRunRejectsMalformedHandler(t *testing.T) {
	rt := NewRuntimeWithClient(
		client.NewAPIClient(
			settings.NewAuthenticationSettings("k", "s"),
			settings.NewHttpSettings("http://127.0.0.1:1/api"),
		), Config{})

	agent := &Agent{
		Name: "bad", Model: testModel,
		Tools: []ToolDef{{
			Name: "oops", ToolType: ToolTypeWorker,
			Handler: func(a, b, c int) error { return nil },
		}},
	}
	_, err := rt.Run(context.Background(), agent, "hi")
	if err == nil || !strings.Contains(err.Error(), "func(context.Context, In) (Out, error)") {
		t.Fatalf("want a handler-shape error, got %v", err)
	}
}

// Plan is compile-only: one request to /agent/compile carrying agentConfig,
// no start, no workers. The fake records what arrived so the test can hold
// the payload to the same shape the Python SDK's plan() sends.
func TestPlanCompilesWithoutStarting(t *testing.T) {
	var got map[string]any
	var hits []string
	mux := http.NewServeMux()
	mux.HandleFunc("/api/token", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"token": "fake-token"})
	})
	mux.HandleFunc("/api/agent/compile", func(w http.ResponseWriter, r *http.Request) {
		hits = append(hits, r.URL.Path)
		json.NewDecoder(r.Body).Decode(&got)
		json.NewEncoder(w).Encode(map[string]any{
			"workflowDef":     map[string]any{"name": "planner", "tasks": []any{map[string]any{"type": "LLM_CHAT_COMPLETE"}}},
			"requiredWorkers": []string{"ping"},
		})
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hits = append(hits, r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
	})
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

	agent := &Agent{
		Name: "planner", Model: testModel, Instructions: "Plan.",
		Tools: []ToolDef{{
			Name: "ping", Description: "Ping.", InputSchema: map[string]any{"type": "object"},
			Handler: func(ctx context.Context, in echoIn) (string, error) { return "pong", nil },
		}},
	}
	plan, err := rt.Plan(context.Background(), agent)
	if err != nil {
		t.Fatalf("Plan: %v", err)
	}

	cfg, ok := got["agentConfig"].(map[string]any)
	if !ok || cfg["name"] != "planner" {
		t.Fatalf("compile payload = %v, want agentConfig for planner", got)
	}
	if _, ok := got["prompt"]; ok {
		t.Errorf("compile payload carries a prompt; Plan must not look like a start")
	}
	for _, h := range hits {
		if h != "/api/agent/compile" {
			t.Errorf("unexpected request %s; Plan must only compile", h)
		}
	}
	if wf, ok := plan["workflowDef"].(map[string]any); !ok || wf["name"] != "planner" {
		t.Errorf("plan result = %v, want the server's workflowDef", plan)
	}
	if _, ok := plan["requiredWorkers"]; !ok {
		t.Errorf("plan result lacks requiredWorkers")
	}

	// Validation runs first, so a broken agent never reaches the server.
	hits = nil
	if _, err := rt.Plan(context.Background(), &Agent{Model: testModel}); err == nil {
		t.Fatal("Plan accepted an agent with no name")
	}
	if len(hits) != 0 {
		t.Errorf("invalid agent still produced requests: %v", hits)
	}
}
