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

// The executable callback tests of the Python SDK's
// e2e/test_suite13_callbacks.py, under the same names. Each runs an agent that
// must call a tool, with lifecycle callbacks attached, and proves the
// callbacks fired without blocking the run. Rather than inspect the workflow's
// tasks, each callback increments an in-process counter, which directly shows
// the callback worker was dispatched and ran here — the same kind of proof the
// skill script tests use, and deterministic under replay.

type echoIn struct {
	Text string `json:"text"`
}

func echoTool() ai.ToolDef {
	return tool.Func("echo_tool", "Echo the input text back.",
		func(ctx context.Context, in echoIn) (string, error) { return "echo:" + in.Text, nil })
}

const echoInstructions = "You are a helpful assistant. You MUST call the echo_tool " +
	"with text='hello' to answer the user. Always use the tool."

func TestBeforeToolCallbackExecutes(t *testing.T) {
	rt := newRuntime(t)
	var calls atomic.Int32
	agent := &ai.Agent{
		Name: "e2e_s13_before_tool", Model: model(t), MaxTurns: 3,
		Instructions: echoInstructions,
		Tools:        []ai.ToolDef{echoTool()},
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
		Tools:        []ai.ToolDef{echoTool()},
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
		Tools:        []ai.ToolDef{echoTool()},
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
