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
	"fmt"
	"time"
)

// AgentHandle controls a run that was started without blocking.
type AgentHandle struct {
	ExecutionID string

	rt *Runtime
}

// Start begins a run and returns at once.
//
// Workers for the agent's tools are registered before the run starts, so a tool
// call cannot arrive before something is polling for it.
func (r *Runtime) Start(ctx context.Context, agent *Agent, prompt string) (*AgentHandle, error) {
	if err := agent.Validate(); err != nil {
		return nil, err
	}
	if err := r.registerWorkers(agent); err != nil {
		return nil, err
	}

	started, err := r.agents.Start(ctx, map[string]any{
		"agentConfig": agent.toConfig(),
		"prompt":      prompt,
		"sessionId":   "",
		"media":       []any{},
	})
	if err != nil {
		return nil, fmt.Errorf("start agent %q: %w", agent.Name, err)
	}
	executionID, ok := started["executionId"].(string)
	if !ok || executionID == "" {
		return nil, fmt.Errorf("start agent %q: server returned no executionId", agent.Name)
	}
	return &AgentHandle{ExecutionID: executionID, rt: r}, nil
}

// Status reports the current state without waiting.
func (h *AgentHandle) Status(ctx context.Context) (*AgentResult, error) {
	status, err := h.rt.agents.Status(ctx, h.ExecutionID)
	if err != nil {
		return nil, err
	}
	return resultFrom(h.ExecutionID, status), nil
}

// Waiting reports whether the run is blocked on a human.
func (h *AgentHandle) Waiting(ctx context.Context) (bool, error) {
	status, err := h.rt.agents.Status(ctx, h.ExecutionID)
	if err != nil {
		return false, err
	}
	if w, ok := status["isWaiting"].(bool); ok {
		return w, nil
	}
	return resultFrom(h.ExecutionID, status).Status == StatusWaiting, nil
}

// Respond answers a run waiting on a human.
//
// It first waits for the server to report the run as waiting. The stream's
// waiting event announces the pause slightly before the server will accept a
// response, so posting on the event alone races it: an early response is
// dropped and the run hangs, and retrying until one lands injects spurious
// turns that re-run the pending tool. Confirming the state first makes
// answering from the event stream safe, which is the way callers naturally
// write it.
//
// A run that never reaches a waiting state returns an error rather than
// blocking forever.
func (h *AgentHandle) Respond(ctx context.Context, output map[string]any) error {
	if err := h.awaitWaiting(ctx); err != nil {
		return err
	}
	return h.rt.agents.Respond(ctx, h.ExecutionID, output)
}

// awaitWaiting blocks until the server reports the run waiting on a human, the
// run reaches a terminal state, or ctx ends.
func (h *AgentHandle) awaitWaiting(ctx context.Context) error {
	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()
	for {
		status, err := h.rt.agents.Status(ctx, h.ExecutionID)
		if err != nil {
			return err
		}
		if w, ok := status["isWaiting"].(bool); ok && w {
			return nil
		}
		if resultFrom(h.ExecutionID, status).Status.Terminal() {
			return fmt.Errorf(
				"execution %s is no longer running, so there is nothing to respond to",
				h.ExecutionID)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
		}
	}
}

// Approve releases a tool call that is waiting for approval.
func (h *AgentHandle) Approve(ctx context.Context) error {
	return h.Respond(ctx, map[string]any{"approved": true})
}

// Reject refuses a tool call that is waiting for approval.
func (h *AgentHandle) Reject(ctx context.Context, reason string) error {
	return h.Respond(ctx, map[string]any{"approved": false, "reason": reason})
}

// Stop terminates the run.
func (h *AgentHandle) Stop(ctx context.Context) error {
	return h.rt.agents.Stop(ctx, h.ExecutionID)
}

// Result blocks until the run reaches a terminal state.
func (h *AgentHandle) Result(ctx context.Context) (*AgentResult, error) {
	return h.rt.awaitResult(ctx, h.ExecutionID)
}
