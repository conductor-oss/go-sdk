//go:build e2e

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai_e2e

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type svcIn struct {
	ServiceName string `json:"service_name"`
}

type svcOut struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

type deleteIn struct {
	ServiceName string `json:"service_name"`
	DataType    string `json:"data_type"`
}

// recorder collects what happened, in order, so the test can assert on the
// sequence rather than only the final state.
type recorder struct {
	mu     sync.Mutex
	events []string
}

func (r *recorder) add(s string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.events = append(r.events, s)
}

func (r *recorder) snapshot() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]string(nil), r.events...)
}

func (r *recorder) indexOf(s string) int {
	for i, e := range r.snapshot() {
		if e == s {
			return i
		}
	}
	return -1
}

// Streaming with a human approval gate.
//
// The point of the test is the gate: a tool marked RequiresApproval must not be
// dispatched to a worker until a human responds, even though the worker is
// polling the whole time. Asserting only on the final output would pass whether
// or not the gate held.
func TestStreamingWithApproval(t *testing.T) {
	// The model sometimes stops after restart_service, deciding the job is
	// done, so the guarded tool is never requested and the gate is never
	// exercised. That is a model decision, not an SDK defect — retry it rather
	// than reporting a failure or silently skipping. Three attempts takes a
	// ~20% per-run miss to under 1%.
	const attempts = 3
	for attempt := 1; attempt <= attempts; attempt++ {
		if runApprovalAttempt(t, attempt, attempt == attempts) {
			return
		}
		t.Logf("attempt %d did not exercise the gate; retrying", attempt)
	}
}

// runApprovalAttempt runs one attempt, reporting whether the approval gate was
// actually exercised. Assertions about SDK behaviour fail immediately on any
// attempt; only "the model did not cooperate" is retryable.
func runApprovalAttempt(t *testing.T, attempt int, last bool) bool {
	rt := newRuntime(t)
	// Per attempt, not per test. newRuntime registers Shutdown with t.Cleanup,
	// which does not run until the whole test ends — so a retry would leave the
	// previous attempt's workers polling the same task names, and they would
	// serve this attempt's tool calls into a recorder nobody reads. Shutdown is
	// idempotent, so the later cleanup call is harmless.
	defer rt.Shutdown()
	rec := &recorder{}

	var deleteCalls atomic.Int32
	var approvedAt atomic.Int64 // unix nanos, 0 until approval is sent

	checkService := func(ctx context.Context, in svcIn) (svcOut, error) {
		rec.add("tool:check_service")
		return svcOut{Service: in.ServiceName, Status: "unhealthy"}, nil
	}
	restartService := func(ctx context.Context, in svcIn) (svcOut, error) {
		rec.add("tool:restart_service")
		return svcOut{Service: in.ServiceName, Status: "restarted"}, nil
	}
	deleteServiceData := func(ctx context.Context, in deleteIn) (svcOut, error) {
		deleteCalls.Add(1)
		rec.add("tool:delete_service_data")
		if approvedAt.Load() == 0 {
			t.Error("delete_service_data ran before any approval was sent")
		}
		return svcOut{Service: in.ServiceName, Status: "deleted"}, nil
	}

	ops := &ai.Agent{
		Name:  "go_e2e_ops_agent",
		Model: model(t),
		// Temperature 0 is the most reproducible the API offers. It narrows
		// but does not remove the chance that the model declines to call the
		// guarded tool, which is why the gate assertion below distinguishes
		// "the model did not cooperate" from "the SDK's gate failed".
		Temperature: ai.Ptr(0.0),
		Instructions: "You are an operations assistant. Work one tool call at a time: " +
			"first check_service, then restart_service if unhealthy, then " +
			"delete_service_data if asked to clear data. A human approves the " +
			"deletion, not you — never ask for approval in your reply.",
		Tools: []ai.ToolDef{
			tool.Func("check_service", "Check the health of a service", checkService),
			tool.Func("restart_service", "Restart a service", restartService),
			tool.Func("delete_service_data", "Delete service data. Destructive.",
				deleteServiceData, tool.RequiresApproval()),
		},
	}

	// 90s is generous: a healthy run finishes in 10-25s. Keeping it short is
	// what makes three attempts fit inside a sane wall clock.
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	h, err := rt.Start(ctx, ops,
		"The payments service is down. Do all three of these, in order: check it, "+
			"restart it, and then clear its stale cache data. The cache deletion is "+
			"required — do not finish until you have called delete_service_data.")
	if err != nil {
		t.Fatalf("Start: %v", err)
	}
	if h.ExecutionID == "" {
		t.Fatal("Start returned no executionId")
	}

	// Consume the stream concurrently. Its events are the evidence that
	// streaming works at all; the control flow below uses status polling so the
	// test does not hang if a build stops emitting a particular event.
	events, err := h.Events(ctx)
	if err != nil {
		t.Fatalf("Events: %v", err)
	}
	var streamDone sync.WaitGroup
	streamDone.Add(1)
	go func() {
		defer streamDone.Done()
		for ev := range events {
			rec.add("event:" + ev.Name)
		}
	}()

	// Approve once the run blocks on a human.
	approved := false
	for i := 0; i < 90; i++ {
		if !approved {
			if waiting, _ := h.Waiting(ctx); waiting {
				approvedAt.Store(time.Now().UnixNano())
				rec.add("approve")
				if err := h.Approve(ctx); err != nil {
					t.Fatalf("Approve: %v", err)
				}
				approved = true
			}
		}
		st, err := h.Status(ctx)
		if err == nil && st.Status.Terminal() {
			break
		}
		time.Sleep(time.Second)
	}

	res, err := h.Result(ctx)
	if err != nil {
		streamDone.Wait()
		// Two stall modes were measured over twenty runs, each about once: the
		// run produces no turns at all, and the model runs away instead of
		// finishing (141 events against the usual 12). Nothing the SDK does
		// causes either, so both belong with "the environment did not
		// cooperate" rather than with a real failure — unless the tool ran
		// unapproved, which is never excusable.
		if deleteCalls.Load() > 0 && approvedAt.Load() == 0 {
			t.Fatalf("delete_service_data ran without approval, then the run stalled: %v; "+
				"sequence: %v", err, rec.snapshot())
		}
		if last {
			t.Fatalf("the run stalled on every attempt; last error: %v, sequence: %v",
				err, rec.snapshot())
		}
		t.Logf("attempt %d: the run stalled (%v, %d events); retrying",
			attempt, err, len(rec.snapshot()))
		return false
	}
	streamDone.Wait()

	if !approved {
		// The gate not holding is always a failure; the model not asking for
		// the tool is only a reason to try again.
		if deleteCalls.Load() > 0 {
			t.Fatalf("delete_service_data ran without the run ever waiting: "+
				"the approval gate did not engage; sequence: %v", rec.snapshot())
		}
		if last {
			t.Skipf("the model never requested the guarded tool in %d attempts, so "+
				"the approval gate was not exercised; last sequence: %v",
				attempt, rec.snapshot())
		}
		return false
	}
	if res.Status != ai.StatusCompleted {
		t.Errorf("status = %q, want %q", res.Status, ai.StatusCompleted)
	}
	if n := deleteCalls.Load(); n != 1 {
		t.Errorf("delete_service_data ran %d times, want exactly 1", n)
	}

	seq := rec.snapshot()

	// The gate: approval must precede the guarded tool call.
	iApprove := rec.indexOf("approve")
	iDelete := rec.indexOf("tool:delete_service_data")
	if iApprove < 0 || iDelete < 0 || iApprove > iDelete {
		t.Errorf("approval must come before the guarded tool ran; sequence: %v", seq)
	}

	// The unguarded tools must have run without waiting for anyone.
	if rec.indexOf("tool:check_service") < 0 {
		t.Errorf("check_service never ran; sequence: %v", seq)
	}

	// Streaming must have produced events, not just an open connection.
	var streamed int
	for _, e := range seq {
		if len(e) > 6 && e[:6] == "event:" {
			streamed++
		}
	}
	if streamed == 0 {
		t.Error("no SSE events arrived; the stream connected but delivered nothing")
	}
	if rec.indexOf("event:waiting") < 0 {
		t.Errorf("no waiting event was streamed; sequence: %v", seq)
	}

	t.Logf("attempt %d: %d events, sequence: %v", attempt, streamed, seq)
	return true
}
