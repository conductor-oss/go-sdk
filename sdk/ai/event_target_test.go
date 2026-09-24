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
	"strings"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/client"
)

// The event-targeted half of the Python SDK's
// e2e/test_suite23_from_instance_and_event_hitl.py, under the same names.
//
// When a nested agent asks for human input, the waiting event names that
// agent's own execution, not the run the caller started. Answering the run
// would approve something nobody was shown, so an answer has to go to the
// execution the event came from. The other half of that Python suite builds
// agents by reflecting over a Python class and has no counterpart here.

// respondSpy records where each answer was sent, in place of a server.
type respondSpy struct {
	client.AgentClient
	calls []respondCall
}

type respondCall struct {
	executionID string
	body        map[string]any
}

func (s *respondSpy) Respond(_ context.Context, executionID string, output map[string]any) error {
	s.calls = append(s.calls, respondCall{executionID: executionID, body: output})
	return nil
}

// Status answers that the execution is waiting, so the handle does not block.
func (s *respondSpy) Status(context.Context, string) (map[string]any, error) {
	return map[string]any{"isWaiting": true}, nil
}

func spyHandle(executionID string) (*AgentHandle, *respondSpy) {
	spy := &respondSpy{}
	rt := &Runtime{agents: spy}
	return &AgentHandle{ExecutionID: executionID, rt: rt}, spy
}

const (
	topLevelExecution = "root-exec-111"
	subExecution      = "sub-exec-999"
)

func waitingEvent() Event {
	return Event{Type: EventType("waiting"), Name: "waiting", Text: "Waiting for human input",
		ExecutionID: subExecution}
}

func TestWaitingEventExposesExecutionID(t *testing.T) {
	if got := waitingEvent().ExecutionID; got != subExecution {
		t.Errorf("the event names execution %q, want the one that is waiting", got)
	}
}

// An event names its own execution when the server gives one.
func TestSseEventInheritsServerExecutionID(t *testing.T) {
	ev := decodeEvent("waiting", `{"type":"waiting","executionId":"`+subExecution+`"}`, topLevelExecution)
	if ev.ExecutionID != subExecution {
		t.Errorf("event execution = %q, want the sub-execution the server named", ev.ExecutionID)
	}
}

// An event that names none belongs to the execution being streamed.
func TestSseEventFallsBackToStreamID(t *testing.T) {
	ev := decodeEvent("thinking", `{"type":"thinking"}`, topLevelExecution)
	if ev.ExecutionID != topLevelExecution {
		t.Errorf("event execution = %q, want the stream's own", ev.ExecutionID)
	}
}

func TestApproveEventTargetsSubExecution(t *testing.T) {
	handle, spy := spyHandle(topLevelExecution)
	sub, err := handle.For(waitingEvent())
	if err != nil {
		t.Fatalf("For: %v", err)
	}
	if err := sub.Approve(context.Background()); err != nil {
		t.Fatalf("Approve: %v", err)
	}
	if len(spy.calls) != 1 {
		t.Fatalf("%d answers sent, want 1", len(spy.calls))
	}
	if spy.calls[0].executionID != subExecution {
		t.Errorf("answered execution %q, want the one that was waiting", spy.calls[0].executionID)
	}
	if spy.calls[0].body["approved"] != true {
		t.Errorf("body = %v, want an approval", spy.calls[0].body)
	}
}

// Approving without an event still answers the run the handle started.
func TestApproveNoEventTargetsTopLevel(t *testing.T) {
	handle, spy := spyHandle(topLevelExecution)
	if err := handle.Approve(context.Background()); err != nil {
		t.Fatalf("Approve: %v", err)
	}
	if len(spy.calls) != 1 || spy.calls[0].executionID != topLevelExecution {
		t.Fatalf("answers = %v, want one to the run itself", spy.calls)
	}
	if spy.calls[0].body["approved"] != true {
		t.Errorf("body = %v", spy.calls[0].body)
	}
}

func TestRejectEventTargetsSubExecution(t *testing.T) {
	handle, spy := spyHandle(topLevelExecution)
	sub, err := handle.For(waitingEvent())
	if err != nil {
		t.Fatalf("For: %v", err)
	}
	if err := sub.Reject(context.Background(), "not allowed"); err != nil {
		t.Fatalf("Reject: %v", err)
	}
	if len(spy.calls) != 1 || spy.calls[0].executionID != subExecution {
		t.Fatalf("answers = %v, want one to the waiting execution", spy.calls)
	}
	body := spy.calls[0].body
	if body["approved"] != false || body["reason"] != "not allowed" {
		t.Errorf("body = %v, want a refusal carrying the reason", body)
	}
}

func TestRespondEventTargetsSubExecution(t *testing.T) {
	handle, spy := spyHandle(topLevelExecution)
	sub, err := handle.For(waitingEvent())
	if err != nil {
		t.Fatalf("For: %v", err)
	}
	if err := sub.Respond(context.Background(), map[string]any{"selected": "writer"}); err != nil {
		t.Fatalf("Respond: %v", err)
	}
	if len(spy.calls) != 1 || spy.calls[0].executionID != subExecution {
		t.Fatalf("answers = %v, want one to the waiting execution", spy.calls)
	}
	if spy.calls[0].body["selected"] != "writer" {
		t.Errorf("body = %v", spy.calls[0].body)
	}
}

// An event naming no execution is refused rather than answered somewhere.
func TestEventWithoutExecutionIDRaises(t *testing.T) {
	handle, spy := spyHandle(topLevelExecution)
	_, err := handle.For(Event{Type: EventType("waiting"), Name: "waiting"})
	if err == nil {
		t.Fatal("an event naming no execution was accepted")
	}
	if !strings.Contains(err.Error(), "names no execution") {
		t.Errorf("error = %v, want it to say the event names no execution", err)
	}
	if len(spy.calls) != 0 {
		t.Errorf("%d answers were sent anyway: %v", len(spy.calls), spy.calls)
	}
}
