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

// EventType classifies a streamed event. Server event names vary by build, so Type is a
// direct conversion of the wire name and an unlisted name still arrives intact. The set
// mirrors the Java and Python SDKs' EventType, in declaration order, so the three diff.
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
	// Name is the server's raw event name, kept so an unmapped name is not lost.
	Name string
	// Text carries streamed output or a message, when the event has one.
	Text string
	// Data is the decoded payload, for fields Type and Text do not cover.
	Data map[string]any
	// ExecutionID is the event's own execution: a nested agent's is a sub-execution; see For.
	ExecutionID string
}

// AgentHandle controls a run that was started without blocking.
type AgentHandle struct {
	ExecutionID string

	rt *Runtime
}

// Events streams updates until the run ends or ctx is cancelled, closing the channel with
// the stream. An already finished run yields no events, so take the outcome from Result.
func (h *AgentHandle) Events(ctx context.Context) (<-chan Event, error) {
	raw, stream, err := h.rt.api.StreamSSE(ctx, "/agent/stream/"+h.ExecutionID, "")
	if err != nil {
		return nil, err
	}

	out := make(chan Event)
	go func() {
		defer close(out)
		// Closing unblocks the reader on early return; body errors land in stream.Err().
		defer stream.Close() //nolint:errcheck // see above
		for ev := range raw {
			select {
			case out <- decodeEvent(ev.Event, ev.Data, h.ExecutionID):
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

// For returns a handle to the execution an event came from, so a nested agent's request
// for human input is answered on its own execution, not the run this handle started. An
// event naming no execution is an error: answering the wrong one approves it unseen.
func (h *AgentHandle) For(ev Event) (*AgentHandle, error) {
	if ev.ExecutionID == "" {
		return nil, fmt.Errorf(
			"cannot answer this %s event: it names no execution. Use an event from Events, "+
				"which carries the execution that is waiting", ev.Name)
	}
	if ev.ExecutionID == h.ExecutionID {
		return h, nil
	}
	return &AgentHandle{ExecutionID: ev.ExecutionID, rt: h.rt}, nil
}

// Respond answers a run waiting on a human, first confirming the server reports it
// waiting: the stream's waiting event announces the pause slightly before the server
// accepts a response, so answering on the event alone is dropped and the run hangs, while
// retrying injects spurious turns that re-run the pending tool. A run that never waits errors.
func (h *AgentHandle) Respond(ctx context.Context, output map[string]any) error {
	if err := h.awaitWaiting(ctx); err != nil {
		return err
	}
	return h.rt.agents.Respond(ctx, h.ExecutionID, output)
}

// awaitWaiting blocks until the server reports the run waiting, the run ends, or ctx does.
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

// Signal injects a persistent signal into this run's context; see Runtime.Signal.
func (h *AgentHandle) Signal(ctx context.Context, message string) error {
	return h.rt.Signal(ctx, h.ExecutionID, message)
}

// SendMessage pushes a message into this run's message queue; see Runtime.SendMessage.
func (h *AgentHandle) SendMessage(ctx context.Context, message any) error {
	return h.rt.SendMessage(ctx, h.ExecutionID, message)
}

// Pause suspends this run; see Runtime.Pause.
func (h *AgentHandle) Pause(ctx context.Context) error {
	return h.rt.Pause(ctx, h.ExecutionID)
}

// Resume continues this run after Pause; see Runtime.Resume.
func (h *AgentHandle) Resume(ctx context.Context) error {
	return h.rt.Resume(ctx, h.ExecutionID)
}

// decodeEvent normalizes one SSE frame. The event name may arrive as the SSE "event:"
// field or inside the JSON payload, depending on the server build, so both are consulted.
func decodeEvent(name, data, streamExecutionID string) Event {
	ev := Event{Name: name, Type: EventType(name), ExecutionID: streamExecutionID}

	var payload map[string]any
	if data != "" && json.Unmarshal([]byte(data), &payload) == nil {
		ev.Data = payload
		if ev.Name == "" {
			if t, ok := payload["type"].(string); ok {
				ev.Name, ev.Type = t, EventType(t)
			}
		}
		ev.Text = firstString(payload, "text", "content", "message", "result", "delta")
		// A nested agent's event names its own execution; without one it is the streamed one.
		if id, ok := payload["executionId"].(string); ok && id != "" {
			ev.ExecutionID = id
		}
	} else if data != "" {
		ev.Text = data
	}
	return ev
}
