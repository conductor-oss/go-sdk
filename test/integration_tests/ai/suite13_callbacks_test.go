//go:build integration

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package integration

import (
	"context"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// The Python SDK's e2e/test_suite13_callbacks.py, test for test and under the
// same names: two that compile an agent and read its callbacks out of the
// plan, and three that run one. Each runs an agent that
// must call a tool, with lifecycle callbacks attached, and proves the
// callbacks fired without blocking the run. Rather than inspect the workflow's
// tasks, each callback increments an in-process counter, which directly shows
// the callback worker was dispatched and ran here — the same kind of proof the
// skill script tests use, and deterministic under replay.

type echoIn struct {
	Text string
}

func echoTool() ai.ToolDef {
	return tool.Func("echo_tool", func(ctx context.Context, in echoIn) (string, error) { return "echo:" + in.Text, nil },
		"Echo the input text back.")
}

const echoInstructions = "You are a helpful assistant. You MUST call the echo_tool " +
	"with text='hello' to answer the user. Always use the tool."

// callbackEntries reads an agent's compiled callbacks: one {position,
// taskName} per hook the agent set.
func callbackEntries(t *testing.T, rt *ai.Runtime, agent *ai.Agent) []map[string]any {
	t.Helper()
	var out []map[string]any
	for _, raw := range asList(agentDef(t, planAgent(t, rt, agent))["callbacks"]) {
		if cb, ok := raw.(map[string]any); ok {
			out = append(out, cb)
		}
	}
	return out
}

// assertCallbackTask fails unless the compiled callbacks hold the position
// with the task name the server will dispatch for it.
func assertCallbackTask(t *testing.T, entries []map[string]any, position, taskName string) {
	t.Helper()
	for _, cb := range entries {
		if cb["position"] == position && cb["taskName"] == taskName {
			return
		}
	}
	t.Errorf("no callback at %q named %q; compiled: %v", position, taskName, entries)
}

// An agent that sets only the two tool hooks compiles to exactly those two
// callbacks, each named "<agent>_<position>".
func TestToolCallbacksCompile(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name: "e2e_s13_tool_cb", Model: model(t), MaxTurns: 3,
		Instructions: "You are a helpful assistant. Use the echo tool.",
		Tools:        ai.Tools(echoTool()),
		Callbacks: &ai.Callbacks{
			OnToolStart: func(context.Context, ai.CallbackInput) (map[string]any, error) { return nil, nil },
			OnToolEnd:   func(context.Context, ai.CallbackInput) (map[string]any, error) { return nil, nil },
		},
	}
	entries := callbackEntries(t, rt, agent)
	if len(entries) < 2 {
		t.Fatalf("compiled %d callbacks, want at least 2: %v", len(entries), entries)
	}
	assertCallbackTask(t, entries, "before_tool", "e2e_s13_tool_cb_before_tool")
	assertCallbackTask(t, entries, "after_tool", "e2e_s13_tool_cb_after_tool")
}

// The same for the two model hooks, on an agent with no tools at all.
func TestModelCallbacksCompile(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name: "e2e_s13_model_cb", Model: model(t), MaxTurns: 3,
		Instructions: "You are a helpful assistant.",
		Callbacks: &ai.Callbacks{
			OnModelStart: func(context.Context, ai.CallbackInput) (map[string]any, error) { return nil, nil },
			OnModelEnd:   func(context.Context, ai.CallbackInput) (map[string]any, error) { return nil, nil },
		},
	}
	entries := callbackEntries(t, rt, agent)
	if len(entries) < 2 {
		t.Fatalf("compiled %d callbacks, want at least 2: %v", len(entries), entries)
	}
	assertCallbackTask(t, entries, "before_model", "e2e_s13_model_cb_before_model")
	assertCallbackTask(t, entries, "after_model", "e2e_s13_model_cb_after_model")
}

func TestBeforeToolCallbackExecutes(t *testing.T) {
	rt := newRuntime(t)
	var calls atomic.Int32
	agent := &ai.Agent{
		Name: "e2e_s13_before_tool", Model: model(t), MaxTurns: 3,
		Instructions: echoInstructions,
		Tools:        ai.Tools(echoTool()),
		Callbacks: &ai.Callbacks{
			OnToolStart: func(context.Context, ai.CallbackInput) (map[string]any, error) {
				calls.Add(1)
				return nil, nil
			},
		},
	}
	res := runCallbackAgent(t, rt, agent, "Say hello using the echo tool.")
	if calls.Load() == 0 {
		t.Errorf("before_tool callback never ran; status=%q", res.Status)
	}
}

func TestAfterToolCallbackExecutes(t *testing.T) {
	rt := newRuntime(t)
	var calls atomic.Int32
	agent := &ai.Agent{
		Name: "e2e_s13_after_tool", Model: model(t), MaxTurns: 3,
		Instructions: echoInstructions,
		Tools:        ai.Tools(echoTool()),
		Callbacks: &ai.Callbacks{
			OnToolEnd: func(context.Context, ai.CallbackInput) (map[string]any, error) {
				calls.Add(1)
				return nil, nil
			},
		},
	}
	res := runCallbackAgent(t, rt, agent, "Say hello using the echo tool.")
	if calls.Load() == 0 {
		t.Errorf("after_tool callback never ran; status=%q", res.Status)
	}
}

func TestAllCallbacksDontBlockExecution(t *testing.T) {
	rt := newRuntime(t)
	var before, after atomic.Int32
	count := func(c *atomic.Int32) ai.CallbackFunc {
		return func(context.Context, ai.CallbackInput) (map[string]any, error) {
			c.Add(1)
			return nil, nil
		}
	}
	agent := &ai.Agent{
		Name: "e2e_s13_all_cb", Model: model(t), MaxTurns: 3,
		Instructions: echoInstructions,
		Tools:        ai.Tools(echoTool()),
		Callbacks: &ai.Callbacks{
			OnAgentStart: count(&before), OnAgentEnd: count(&after),
			OnModelStart: count(&before), OnModelEnd: count(&after),
			OnToolStart: count(&before), OnToolEnd: count(&after),
		},
	}
	res := runCallbackAgent(t, rt, agent, "Say hello using the echo tool.")
	if res.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q; all six callbacks must not block the run", res.Status, ai.StatusCompleted)
	}
	// At least the model-lifecycle hooks fire on any run; if the model called
	// the tool, the tool hooks fire too. Requiring both sides proves the run
	// advanced through its lifecycle with every hook attached.
	if before.Load() == 0 || after.Load() == 0 {
		t.Errorf("callbacks did not run: before=%d after=%d", before.Load(), after.Load())
	}
	if !strings.Contains(res.Output, "hello") {
		t.Logf("note: final output does not mention hello: %q", res.Output)
	}
}

func runCallbackAgent(t *testing.T, rt *ai.Runtime, agent *ai.Agent, prompt string) *ai.AgentResult {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	res, err := rt.Run(ctx, agent, prompt)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if res.Status != ai.StatusCompleted && res.Status != ai.StatusTerminated {
		t.Fatalf("status = %q, want COMPLETED or TERMINATED", res.Status)
	}
	return res
}
