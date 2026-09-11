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
	mu   sync.Mutex
	hits []hit
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
		if _, ok := rec.find("/api/agent/deploy"); ok && rt.started["ping"] {
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
