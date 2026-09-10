//go:build e2e

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

// Package ai_e2e exercises the agent surface against a live Conductor
// server. These tests call a real LLM, so they are behind the `e2e` build tag
// and are not part of `go test ./...`.
//
//	go test -tags e2e ./test/ai_e2e/ -v
//
// Requirements:
//   - CONDUCTOR_SERVER_URL pointing at a server with the agent runtime, which
//     conductor-oss ships from 3.32.0-rc.8 onward. Earlier builds answer 404 on
//     every /agent path.
//   - An LLM provider key in the *server* process environment. conductor-oss
//     reads provider credentials from its own environment; the SDK never sees
//     them.
package ai_e2e

import (
	"context"
	"os"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

const defaultModel = "openai/gpt-4o-mini"

// mockModel is the server's playback provider: responses come from recorded
// files instead of a model. See README, "Recording and replaying".
const mockModel = "mock/mockLLM"

// newRuntime builds a Runtime against the server under test, or skips.
//
// Skipping rather than failing on a missing URL keeps `go test -tags e2e ./...`
// usable without a server; a wrong URL still fails loudly on first use.
func newRuntime(t *testing.T) *ai.Runtime {
	t.Helper()
	url := os.Getenv("CONDUCTOR_SERVER_URL")
	if url == "" {
		t.Skip("CONDUCTOR_SERVER_URL is not set")
	}
	api := client.NewAPIClient(
		settings.NewAuthenticationSettings(
			os.Getenv("CONDUCTOR_AUTH_KEY"), os.Getenv("CONDUCTOR_AUTH_SECRET")),
		settings.NewHttpSettings(url),
	)
	rt := ai.NewRuntimeWithClient(api, ai.Config{
		WorkerPollInterval: 100 * time.Millisecond,
		StatusPollInterval: 500 * time.Millisecond,
	})
	t.Cleanup(rt.Shutdown)
	return rt
}

func model(t *testing.T) string {
	if m := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL"); m != "" {
		return m
	}
	return defaultModel
}

type weatherIn struct {
	City string `json:"city"`
}

type weatherOut struct {
	City      string `json:"city"`
	TempF     int    `json:"temp_f"`
	Condition string `json:"condition"`
}

// A tool-using agent, end to end: the SDK serializes the definition, the server
// compiles and runs it, the model chooses the tool, and Conductor dispatches
// that call back to a worker in this process.
func TestToolCall(t *testing.T) {
	rt := newRuntime(t)

	var calls atomic.Int32
	var gotCity atomic.Value
	gotCity.Store("")

	getWeather := func(ctx context.Context, in weatherIn) (weatherOut, error) {
		calls.Add(1)
		gotCity.Store(in.City)
		return weatherOut{City: in.City, TempF: 72, Condition: "Sunny"}, nil
	}

	agent := &ai.Agent{
		Name:         "go_e2e_weather_bot",
		Model:        model(t),
		Instructions: "You are a helpful assistant. Use the tools to answer questions.",
		Tools: []ai.ToolDef{
			tool.Func("get_weather", "Get the current weather for a city", getWeather),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	res, err := rt.Run(ctx, agent, "What is the weather like in San Francisco?")
	if err != nil {
		t.Fatalf("Run: %v", err)
	}

	if res.Status != ai.StatusCompleted {
		t.Errorf("status = %q, want %q", res.Status, ai.StatusCompleted)
	}
	if res.ExecutionID == "" {
		t.Error("no executionId returned")
	}

	// The tool must have run in this process. Without this the test would pass
	// on a model that answered from memory and never called the tool.
	if n := calls.Load(); n != 1 {
		t.Errorf("get_weather ran %d times, want exactly 1", n)
	}
	if city, _ := gotCity.Load().(string); !strings.Contains(strings.ToLower(city), "san francisco") {
		t.Errorf("tool received city %q, want San Francisco", city)
	}

	// The answer must reflect what the tool returned, not just be non-empty.
	if res.Output == "" {
		t.Fatal("no output; check that resultFrom still matches the server's status shape")
	}
	if !strings.Contains(res.Output, "72") {
		t.Errorf("output does not carry the tool's value 72: %q", res.Output)
	}
	if res.FinishReason == "" {
		t.Error("no finishReason reported")
	}

	t.Logf("execution %s finished %s: %s", res.ExecutionID, res.FinishReason, res.Output)
}
