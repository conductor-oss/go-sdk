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
	"fmt"
	"time"
)

// EventType classifies a streamed event.
//
// The server's event names vary by build, so Event keeps the raw name too:
// Type is a direct conversion of the wire name, so a name not listed here
// still arrives intact.
//
// The set mirrors the Java and Python SDKs' EventType, in their declaration
// order, so the three can be diffed against each other.
type EventType string

const (
	EventThinking      EventType = "thinking"
	EventToolCall      EventType = "tool_call"
	EventToolResult    EventType = "tool_result"
	EventHandoff       EventType = "handoff"
	EventWaiting       EventType = "waiting"
	EventMessage       EventType = "message"
	EventError         EventType = "error"
	EventDone          EventType = "done"
	EventGuardrailPass EventType = "guardrail_pass"
	EventGuardrailFail EventType = "guardrail_fail"
)

// Event is one update from a running agent.
type Event struct {
	// Type is the normalized event kind.
	Type EventType
	// Name is the server's raw event name, kept because the taxonomy is still
	// settling and an unmapped name would otherwise be lost.
	Name string
	// Text carries streamed output or a message, when the event has one.
	Text string
	// Data is the decoded payload, for fields Type and Text do not cover.
	Data map[string]any
}

// AgentHandle controls a run that was started without blocking.
type AgentHandle struct {
	ExecutionID string

	rt *Runtime
}

// Events streams updates until the run ends or ctx is cancelled.
//
// The channel closes when the stream does. A run that is already finished
// yields no events, so callers that need the outcome should use Result rather
// than inferring it from the stream ending.
func (h *AgentHandle) Events(ctx context.Context) (<-chan Event, error) {
	raw, stream, err := h.rt.api.StreamSSE(ctx, "/agent/stream/"+h.ExecutionID, "")
	if err != nil {
		return nil, err
	}

	out := make(chan Event)
	go func() {
		defer close(out)
		// Closing unblocks the reader on early return. Its error has no consumer
		// here: the read loop records body errors in stream.Err() itself.
		defer stream.Close() //nolint:errcheck // see above
		for ev := range raw {
			select {
			case out <- decodeEvent(ev.Event, ev.Data):
			case <-ctx.Done():
				return
			}
		}
	}()
	return out, nil
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

// decodeEvent normalizes one SSE frame.
//
// The event name may arrive as the SSE "event:" field or inside the JSON
// payload, depending on the server build, so both are consulted.
func decodeEvent(name, data string) Event {
	ev := Event{Name: name, Type: EventType(name)}

	var payload map[string]any
	if data != "" && json.Unmarshal([]byte(data), &payload) == nil {
		ev.Data = payload
		if ev.Name == "" {
			if t, ok := payload["type"].(string); ok {
				ev.Name, ev.Type = t, EventType(t)
			}
		}
		ev.Text = firstString(payload, "text", "content", "message", "result", "delta")
	} else if data != "" {
		ev.Text = data
	}
	return ev
}
