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
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

// capturing fake server: records the path and body of every request so the
// operational-surface tests can assert on the exact wire calls.
type capture struct {
	mu         sync.Mutex
	hits       []hit
	methodHits []methodHit
}

type hit struct {
	path string
	body map[string]any
}

func (c *capture) add(path string, body map[string]any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.hits = append(c.hits, hit{path, body})
}

func (c *capture) find(path string) (map[string]any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, h := range c.hits {
		if h.path == path {
			return h.body, true
		}
	}
	return nil, false
}

func captureServer(t *testing.T, rec *capture, handlers map[string]func() map[string]any) *Runtime {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/token", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"token": "fake-token"})
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		json.NewDecoder(r.Body).Decode(&body)
		rec.add(r.URL.Path, body)
		rec.mu.Lock()
		rec.methodHits = append(rec.methodHits, methodHit{r.Method, r.URL.Path})
		rec.mu.Unlock()
		if h, ok := handlers[r.URL.Path]; ok {
			json.NewEncoder(w).Encode(h())
			return
		}
		w.Write([]byte("{}"))
	})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		mux.ServeHTTP(w, r)
	}))
	t.Cleanup(srv.Close)
	api := client.NewAPIClient(
		settings.NewAuthenticationSettings("key", "secret"),
		settings.NewHttpSettings(srv.URL+"/api"),
	)
	rt := NewRuntimeWithClient(api, Config{})
	t.Cleanup(rt.Shutdown)
	return rt
}

func TestDeployReturnsRegisteredName(t *testing.T) {
	rec := &capture{}
	rt := captureServer(t, rec, map[string]func() map[string]any{
		"/api/agent/deploy": func() map[string]any { return map[string]any{"agentName": "greeter_v3"} },
	})
	name, err := rt.Deploy(context.Background(), &Agent{Name: "greeter", Model: testModel, Instructions: "Hi."})
	if err != nil {
		t.Fatal(err)
	}
	if name != "greeter_v3" {
		t.Errorf("registered name = %q, want greeter_v3", name)
	}
	body, ok := rec.find("/api/agent/deploy")
	if !ok {
		t.Fatal("no request to /agent/deploy")
	}
	cfg, ok := body["agentConfig"].(map[string]any)
	if !ok || cfg["name"] != "greeter" {
		t.Errorf("deploy payload = %v, want agentConfig for greeter", body)
	}
	if _, ok := body["prompt"]; ok {
		t.Error("deploy must not carry a prompt")
	}
	// A server that echoes no name falls back to the agent's own name.
	rec2 := &capture{}
	rt2 := captureServer(t, rec2, nil)
	name, _ = rt2.Deploy(context.Background(), &Agent{Name: "solo", Model: testModel, Instructions: "Hi."})
	if name != "solo" {
		t.Errorf("fallback name = %q, want solo", name)
	}
}

func TestRunSettingsAndMediaReachThePayload(t *testing.T) {
	rec := &capture{}
	rt := captureServer(t, rec, map[string]func() map[string]any{
		"/api/agent/start": func() map[string]any { return map[string]any{"executionId": "e1"} },
		"/api/agent/e1/status": func() map[string]any {
			return map[string]any{"status": "COMPLETED", "output": map[string]any{"result": "done"}}
		},
	})
	agent := &Agent{Name: "vision", Model: testModel, Instructions: "Read."}
	_, err := rt.Run(context.Background(), agent, "what is in the image?",
		WithMedia("/srv/media/a.png", "https://example.test/b.png"),
		WithRunSettings(RunSettings{
			Model: "openai/gpt-4o", Temperature: Ptr(0.2), MaxTokens: Ptr(1024),
			ReasoningEffort: ReasoningEffortHigh, ThinkingBudgetTokens: Ptr(2048),
		}))
	if err != nil {
		t.Fatal(err)
	}
	body, _ := rec.find("/api/agent/start")
	media, _ := body["media"].([]any)
	if !reflect.DeepEqual(media, []any{"/srv/media/a.png", "https://example.test/b.png"}) {
		t.Errorf("media = %v", body["media"])
	}
	cfg, _ := body["agentConfig"].(map[string]any)
	if cfg["model"] != "openai/gpt-4o" || cfg["temperature"] != 0.2 || cfg["maxTokens"] != float64(1024) ||
		cfg["reasoningEffort"] != "high" {
		t.Errorf("run settings not merged into agentConfig: model=%v temp=%v max=%v effort=%v",
			cfg["model"], cfg["temperature"], cfg["maxTokens"], cfg["reasoningEffort"])
	}
	tc, ok := cfg["thinkingConfig"].(map[string]any)
	if !ok || tc["enabled"] != true || tc["budgetTokens"] != float64(2048) {
		t.Errorf("thinkingConfig = %v", cfg["thinkingConfig"])
	}
	// The stored agent is untouched: its model is still the original.
	if agent.Model != testModel {
		t.Errorf("run settings mutated the agent: model = %q", agent.Model)
	}
}

func TestSignalAndSendMessage(t *testing.T) {
	rec := &capture{}
	rt := captureServer(t, rec, nil)
	ctx := context.Background()

	if err := rt.Signal(ctx, "e1", "be brief"); err != nil {
		t.Fatal(err)
	}
	if body, ok := rec.find("/api/agent/e1/signal"); !ok || body["message"] != "be brief" {
		t.Errorf("signal request = %v (found=%v)", body, ok)
	}

	// A non-map message is wrapped under "message"; a map is sent as is.
	if err := rt.SendMessage(ctx, "e2", "hello"); err != nil {
		t.Fatal(err)
	}
	if body, ok := rec.find("/api/workflow/e2/messages"); !ok || body["message"] != "hello" {
		t.Errorf("string send-message = %v (found=%v)", body, ok)
	}
	rec.hits = nil
	if err := rt.SendMessage(ctx, "e3", map[string]any{"topic": "sales", "n": 3}); err != nil {
		t.Fatal(err)
	}
	body, ok := rec.find("/api/workflow/e3/messages")
	if !ok || body["topic"] != "sales" || body["n"] != float64(3) {
		t.Errorf("map send-message = %v (found=%v)", body, ok)
	}
	if _, wrapped := body["message"]; wrapped {
		t.Error("a map message must not be wrapped under 'message'")
	}

	// The handle delegates to the runtime.
	h := &AgentHandle{ExecutionID: "e4", rt: rt}
	if err := h.Signal(ctx, "stop soon"); err != nil {
		t.Fatal(err)
	}
	if body, ok := rec.find("/api/agent/e4/signal"); !ok || body["message"] != "stop soon" {
		t.Errorf("handle signal = %v (found=%v)", body, ok)
	}
}

func TestServeDeploysRegistersAndBlocksUntilCancel(t *testing.T) {
	rec := &capture{}
	rt := captureServer(t, rec, map[string]func() map[string]any{
		"/api/agent/deploy": func() map[string]any { return map[string]any{"agentName": "worker_bot"} },
	})
	agent := &Agent{
		Name: "worker_bot", Model: testModel, Instructions: "Use the tool.",
		Tools: []ToolDef{{Name: "ping", InputSchema: map[string]any{"type": "object"},
			Handler: func(ctx context.Context, in echoIn) (string, error) { return "pong", nil }}},
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- rt.Serve(ctx, agent) }()

	// Serve blocks; wait for it to have deployed and registered before cancelling.
	deadline := time.After(2 * time.Second)
	for {
		if _, ok := rec.find("/api/agent/deploy"); ok && rt.started[workerKey{name: "ping"}] {
			break
		}
		select {
		case <-deadline:
			t.Fatal("Serve did not deploy and register within 2s")
		case err := <-done:
			t.Fatalf("Serve returned early: %v", err)
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}

	cancel()
	select {
	case err := <-done:
		if err != context.Canceled {
			t.Errorf("Serve returned %v, want context.Canceled", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Serve did not return after cancel")
	}
	if _, ok := rec.find("/api/metadata/taskdefs"); !ok {
		t.Error("Serve did not register a task definition")
	}
	// Serve needs at least one agent.
	if err := rt.Serve(context.Background()); err == nil {
		t.Error("Serve with no agents should error")
	}
}

// captureAll is captureServer with the request method recorded too, for the
// pause/resume tests where the path alone is not enough.
type methodHit struct {
	method, path string
}

func TestPauseResumeAndHandleDelegation(t *testing.T) {
	var mu sync.Mutex
	var hits []methodHit
	mux := http.NewServeMux()
	mux.HandleFunc("/api/token", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"token": "fake-token"})
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		hits = append(hits, methodHit{r.Method, r.URL.Path})
		mu.Unlock()
		w.Write([]byte("{}"))
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
	ctx := context.Background()

	seen := func(method, path string) bool {
		mu.Lock()
		defer mu.Unlock()
		for _, h := range hits {
			if h.method == method && h.path == path {
				return true
			}
		}
		return false
	}

	if err := rt.Pause(ctx, "e1"); err != nil {
		t.Fatal(err)
	}
	if !seen("PUT", "/api/workflow/e1/pause") {
		t.Errorf("Pause did not PUT /workflow/e1/pause; hits=%v", hits)
	}
	if err := rt.Resume(ctx, "e1"); err != nil {
		t.Fatal(err)
	}
	if !seen("PUT", "/api/workflow/e1/resume") {
		t.Errorf("Resume did not PUT /workflow/e1/resume; hits=%v", hits)
	}

	// The handle delegates to the runtime.
	h := &AgentHandle{ExecutionID: "e2", rt: rt}
	if err := h.Pause(ctx); err != nil {
		t.Fatal(err)
	}
	if err := h.Resume(ctx); err != nil {
		t.Fatal(err)
	}
	if !seen("PUT", "/api/workflow/e2/pause") || !seen("PUT", "/api/workflow/e2/resume") {
		t.Errorf("handle pause/resume did not reach the workflow endpoints; hits=%v", hits)
	}
}

// sawDelete reports whether a DELETE reached path.
func (c *capture) sawDelete(path string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, h := range c.methodHits {
		if h.method == "DELETE" && h.path == path {
			return true
		}
	}
	return false
}

// countPost counts POSTs to path.
func (c *capture) countPost(path string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	n := 0
	for _, h := range c.methodHits {
		if h.method == "POST" && h.path == path {
			n++
		}
	}
	return n
}

// schedulerListServer serves GET /scheduler/schedules as a list of the given
// wire names (so ReconcileSchedules can read the existing set), records every
// request's method and path, and answers save/delete/pause/resume with {}.
func schedulerListServer(t *testing.T, rec *capture, names []string) *Runtime {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/token", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"token": "fake-token"})
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		json.NewDecoder(r.Body).Decode(&body)
		rec.add(r.URL.Path, body)
		rec.mu.Lock()
		rec.methodHits = append(rec.methodHits, methodHit{r.Method, r.URL.Path})
		rec.mu.Unlock()
		if r.Method == "GET" && r.URL.Path == "/api/scheduler/schedules" {
			list := make([]map[string]any, 0, len(names))
			for _, n := range names {
				list = append(list, map[string]any{"name": n})
			}
			json.NewEncoder(w).Encode(list)
			return
		}
		w.Write([]byte("{}"))
	})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		mux.ServeHTTP(w, r)
	}))
	t.Cleanup(srv.Close)
	api := client.NewAPIClient(
		settings.NewAuthenticationSettings("key", "secret"),
		settings.NewHttpSettings(srv.URL+"/api"),
	)
	rt := NewRuntimeWithClient(api, Config{})
	t.Cleanup(rt.Shutdown)
	return rt
}

// The session id groups runs into one conversation, so it must reach the start
// request verbatim on both the native and the skill payload.
func TestWithSessionReachesTheStartRequest(t *testing.T) {
	rec := &capture{}
	rt := captureServer(t, rec, map[string]func() map[string]any{
		"/api/agent/start": func() map[string]any { return map[string]any{"executionId": "e1"} },
		"/api/agent/e1/status": func() map[string]any {
			return map[string]any{"status": "COMPLETED", "output": map[string]any{"result": "done"}}
		},
	})
	agent := &Agent{Name: "assistant", Model: testModel, Instructions: "Be brief."}
	if _, err := rt.Run(context.Background(), agent, "hi", WithSession("user-42")); err != nil {
		t.Fatal(err)
	}
	body, _ := rec.find("/api/agent/start")
	if body["sessionId"] != "user-42" {
		t.Errorf("sessionId = %v, want user-42", body["sessionId"])
	}
}

// Without the option the key is still sent, empty, which is the shape the server
// and the Python SDK expect; a missing key is not the same as an empty one.
func TestSessionDefaultsToEmpty(t *testing.T) {
	rt := &Runtime{}
	agent := &Agent{Name: "assistant", Model: testModel, Instructions: "Be brief."}
	payload, err := rt.startPayload(agent, "hi", nil, "")
	if err != nil {
		t.Fatal(err)
	}
	got, present := payload["sessionId"]
	if !present || got != "" {
		t.Errorf("sessionId = %v (present %v), want an empty string", got, present)
	}
}
