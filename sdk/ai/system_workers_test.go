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
	"errors"
	"testing"
)

// The server compiles a termination condition, a stop-when predicate and a
// router function into tasks of their own, so each needs a worker. Without
// one the task stays SCHEDULED and the run never finishes.

func workerNames(a *Agent) map[string]bool {
	out := map[string]bool{}
	for _, t := range a.workerTools() {
		out[t.Name] = true
	}
	return out
}

func TestAgentRegistersAWorkerForEverySystemTask(t *testing.T) {
	agent := &Agent{
		Name:        "acme",
		Termination: TextMentionTermination{Text: "DONE"},
		StopWhen:    func(context.Context, StopWhenState) (bool, error) { return false, nil },
		RouterFunc:  func(context.Context, string) (string, error) { return "b", nil },
		Agents:      []*Agent{{Name: "b"}},
	}
	got := workerNames(agent)
	for _, want := range []string{"acme_termination", "acme_stop_when", "acme_router_fn"} {
		if !got[want] {
			t.Errorf("no worker registered for %q; registered: %v", want, got)
		}
	}
	// An agent that sets none of them gets none of these workers.
	if n := len(workerNames(&Agent{Name: "plain"})); n != 0 {
		t.Errorf("a plain agent registered %d system workers", n)
	}
}

func runTermination(t *testing.T, cond TerminationCondition, in loopStateIn) shouldContinueOut {
	t.Helper()
	out, err := terminationHandler(cond)(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func TestTerminationWorkerVerdicts(t *testing.T) {
	// A sentinel in the reply stops the loop, and the reason says which.
	out := runTermination(t, TextMentionTermination{Text: "TASK_COMPLETE"},
		loopStateIn{Result: "all done, task_complete"})
	if out.ShouldContinue || out.Reason != "Text 'TASK_COMPLETE' found in output" {
		t.Errorf("case-insensitive mention = %+v", out)
	}
	out = runTermination(t, TextMentionTermination{Text: "TASK_COMPLETE", CaseSensitive: true},
		loopStateIn{Result: "all done, task_complete"})
	if !out.ShouldContinue {
		t.Errorf("a case-sensitive mention must not match a different case: %+v", out)
	}

	// A message limit counts the loop's iterations when no history is sent.
	out = runTermination(t, MaxMessageTermination{MaxMessages: 1}, loopStateIn{Iteration: 1})
	if out.ShouldContinue || out.Reason != "Message count (1) >= limit (1)" {
		t.Errorf("max message at the limit = %+v", out)
	}
	if out := runTermination(t, MaxMessageTermination{MaxMessages: 3}, loopStateIn{Iteration: 1}); !out.ShouldContinue {
		t.Errorf("max message below the limit = %+v", out)
	}

	// An empty stop message means the default.
	out = runTermination(t, StopMessageTermination{}, loopStateIn{Result: "  TERMINATE \n"})
	if out.ShouldContinue || out.Reason != "Stop message 'TERMINATE' received" {
		t.Errorf("default stop message = %+v", out)
	}

	// And needs every condition; Or needs one.
	both := AndTermination(TextMentionTermination{Text: "a"}, TextMentionTermination{Text: "b"})
	if out := runTermination(t, both, loopStateIn{Result: "only a"}); !out.ShouldContinue {
		t.Errorf("and with one condition unmet = %+v", out)
	}
	if out := runTermination(t, both, loopStateIn{Result: "a and b"}); out.ShouldContinue ||
		out.Reason != "Text 'a' found in output AND Text 'b' found in output" {
		t.Errorf("and with both met = %+v", out)
	}
	either := OrTermination(TextMentionTermination{Text: "x"}, TextMentionTermination{Text: "y"})
	if out := runTermination(t, either, loopStateIn{Result: "has y"}); out.ShouldContinue ||
		out.Reason != "Text 'y' found in output" {
		t.Errorf("or with one met = %+v", out)
	}
	if out := runTermination(t, either, loopStateIn{Result: "neither"}); !out.ShouldContinue {
		t.Errorf("or with none met = %+v", out)
	}
}

func TestStopWhenWorkerInvertsAndFailsOpen(t *testing.T) {
	var seen StopWhenState
	stop := StopWhenFunc(func(_ context.Context, s StopWhenState) (bool, error) {
		seen = s
		return s.Result == "halt", nil
	})
	out, _ := stop.stopWhenHandler()(context.Background(),
		loopStateIn{Result: "halt", Iteration: 2, Messages: []map[string]any{{"role": "user"}}})
	if out.ShouldContinue {
		t.Errorf("a true predicate must stop the loop: %+v", out)
	}
	if seen.Result != "halt" || seen.Iteration != 2 || len(seen.Messages) != 1 {
		t.Errorf("the predicate saw %+v, not the turn the server sent", seen)
	}
	if out, _ := stop.stopWhenHandler()(context.Background(), loopStateIn{Result: "go on"}); !out.ShouldContinue {
		t.Errorf("a false predicate must continue the loop: %+v", out)
	}

	// An error continues the loop rather than ending the run.
	failing := StopWhenFunc(func(context.Context, StopWhenState) (bool, error) { return true, errors.New("boom") })
	if out, _ := failing.stopWhenHandler()(context.Background(), loopStateIn{}); !out.ShouldContinue {
		t.Errorf("a failing predicate must fail open: %+v", out)
	}
}

func TestRouterWorkerPicksAnAgentAndFallsBack(t *testing.T) {
	pick := RouterFunc(func(_ context.Context, prompt string) (string, error) {
		if prompt == "math" {
			return "calculator", nil
		}
		return "", errors.New("no idea")
	})
	out, _ := pick.routerHandler("first_agent")(context.Background(), routerIn{Prompt: "math"})
	if out.SelectedAgent != "calculator" {
		t.Errorf("selected %q, want the function's choice", out.SelectedAgent)
	}
	out, _ = pick.routerHandler("first_agent")(context.Background(), routerIn{Prompt: "other"})
	if out.SelectedAgent != "first_agent" {
		t.Errorf("selected %q, want the fallback after an error", out.SelectedAgent)
	}
}
