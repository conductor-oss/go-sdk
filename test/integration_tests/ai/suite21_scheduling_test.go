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
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// The Python SDK's e2e/test_suite21_scheduling.py as Go tests, one per Python
// test and under the same names.
//
// Declaring an agent's schedules is a reconcile: the list given is the list
// that ends up on the server, with anything missing from it removed. The rest
// is the lifecycle of one schedule: pause, resume, delete, and reading it back.
//
// Two of the Python tests are not here. Both need a call the Go SDK does not
// have: previewing a cron's next fire times, and running a schedule at once.
// See the README, "Suite 2-26 tests".

// s21Agent is a name no other test uses, so these can run beside anything else.
func s21Agent(t *testing.T) string {
	t.Helper()
	return fmt.Sprintf("e2e_s21_%s", strings.ToLower(strings.ReplaceAll(t.Name(), "/", "_")))
}

// s21Runtime gives a runtime and an agent whose schedules are cleared when the
// test ends, however it ends.
func s21Runtime(t *testing.T) (*ai.Runtime, string, context.Context) {
	t.Helper()
	rt := newRuntime(t)
	requireScheduler(t)
	agent := s21Agent(t)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	t.Cleanup(func() {
		defer cancel()
		if err := rt.ReconcileSchedules(context.Background(), agent, []ai.Schedule{}); err != nil {
			t.Logf("cleanup: %v", err)
		}
	})
	return rt, agent, ctx
}

// byName is the agent's schedules, keyed by their short name.
func byName(t *testing.T, rt *ai.Runtime, ctx context.Context, agent string) map[string]ai.Schedule {
	t.Helper()
	got, err := rt.ListSchedules(ctx, agent)
	if err != nil {
		t.Fatalf("list schedules: %v", err)
	}
	out := map[string]ai.Schedule{}
	for _, s := range got {
		out[s.Name] = s
	}
	return out
}

// Declaring two schedules creates both, with the cron and the input they were
// given.
func TestCreatesSchedules(t *testing.T) {
	rt, agent, ctx := s21Runtime(t)
	desired := []ai.Schedule{
		{Name: "daily", Cron: "0 0 9 * * ?", Input: map[string]any{"k": 1}},
		{Name: "weekly", Cron: "0 0 9 * * MON"},
	}
	if err := rt.ReconcileSchedules(ctx, agent, desired); err != nil {
		t.Fatalf("reconcile: %v", err)
	}

	got := byName(t, rt, ctx, agent)
	if len(got) != 2 || got["daily"].Name == "" || got["weekly"].Name == "" {
		t.Fatalf("schedules = %v, want daily and weekly", got)
	}
	if got["daily"].Cron != "0 0 9 * * ?" {
		t.Errorf("daily cron = %q", got["daily"].Cron)
	}
	// The input the schedule fires with survives the round trip.
	if v := got["daily"].Input["k"]; asInt(v) != 1 {
		t.Errorf("daily input = %v, want k=1", got["daily"].Input)
	}
}

// Declaring a different list updates what stayed, adds what is new, and
// removes what is gone.
func TestUpsertAndPrune(t *testing.T) {
	rt, agent, ctx := s21Runtime(t)
	first := []ai.Schedule{{Name: "a", Cron: "0 0 1 * * ?"}, {Name: "b", Cron: "0 0 2 * * ?"}}
	if err := rt.ReconcileSchedules(ctx, agent, first); err != nil {
		t.Fatalf("first reconcile: %v", err)
	}
	second := []ai.Schedule{{Name: "a", Cron: "0 0 9 * * ?"}, {Name: "c", Cron: "0 0 17 * * ?"}}
	if err := rt.ReconcileSchedules(ctx, agent, second); err != nil {
		t.Fatalf("second reconcile: %v", err)
	}

	got := byName(t, rt, ctx, agent)
	if len(got) != 2 || got["a"].Name == "" || got["c"].Name == "" {
		t.Fatalf("schedules = %v, want a and c", got)
	}
	if got["a"].Cron != "0 0 9 * * ?" {
		t.Errorf("a's cron = %q, want the updated one", got["a"].Cron)
	}
}

// An empty list removes every schedule the agent has.
func TestEmptyListPurges(t *testing.T) {
	rt, agent, ctx := s21Runtime(t)
	if err := rt.ReconcileSchedules(ctx, agent, []ai.Schedule{{Name: "x", Cron: "0 * * * * ?"}}); err != nil {
		t.Fatalf("reconcile: %v", err)
	}
	if n := len(byName(t, rt, ctx, agent)); n != 1 {
		t.Fatalf("%d schedules after declaring one", n)
	}
	if err := rt.ReconcileSchedules(ctx, agent, []ai.Schedule{}); err != nil {
		t.Fatalf("purge: %v", err)
	}
	if got := byName(t, rt, ctx, agent); len(got) != 0 {
		t.Errorf("schedules remain after the purge: %v", got)
	}
}

// No list at all leaves the schedules alone, which is what makes it safe to
// call from code that does not manage them.
func TestNonePreserves(t *testing.T) {
	rt, agent, ctx := s21Runtime(t)
	if err := rt.ReconcileSchedules(ctx, agent, []ai.Schedule{{Name: "x", Cron: "0 * * * * ?"}}); err != nil {
		t.Fatalf("reconcile: %v", err)
	}
	if err := rt.ReconcileSchedules(ctx, agent, nil); err != nil {
		t.Fatalf("reconcile with none: %v", err)
	}
	got := byName(t, rt, ctx, agent)
	if len(got) != 1 || got["x"].Name == "" {
		t.Errorf("schedules = %v, want the one already there", got)
	}
}

// Two schedules of one name are refused before anything is written, so a
// mistake cannot half-apply.
func TestDuplicateNameRaisesBeforeIO(t *testing.T) {
	rt, agent, ctx := s21Runtime(t)
	err := rt.ReconcileSchedules(ctx, agent, []ai.Schedule{
		{Name: "dup", Cron: "0 * * * * ?"},
		{Name: "dup", Cron: "0 0 9 * * ?"},
	})
	if err == nil {
		t.Fatal("two schedules of one name were accepted")
	}
	if !strings.Contains(strings.ToLower(err.Error()), "duplicate") {
		t.Errorf("error = %v, want it to name the duplicate", err)
	}
	if got := byName(t, rt, ctx, agent); len(got) != 0 {
		t.Errorf("the refused list still wrote something: %v", got)
	}
}

// A schedule can be paused and started again, and says which it is.
func TestPauseThenResume(t *testing.T) {
	rt, agent, ctx := s21Runtime(t)
	if err := rt.ReconcileSchedules(ctx, agent, []ai.Schedule{{Name: "p", Cron: "0 0 9 * * ?"}}); err != nil {
		t.Fatalf("reconcile: %v", err)
	}
	read := func(step string) ai.Schedule {
		s, err := rt.GetSchedule(ctx, agent, "p")
		if err != nil {
			t.Fatalf("%s: get: %v", step, err)
		}
		return *s
	}
	if read("after create").Paused {
		t.Fatal("a new schedule is paused")
	}
	if err := rt.PauseSchedule(ctx, agent, "p"); err != nil {
		t.Fatalf("pause: %v", err)
	}
	if !read("after pause").Paused {
		t.Error("the schedule is not paused after pausing it")
	}
	if err := rt.ResumeSchedule(ctx, agent, "p"); err != nil {
		t.Fatalf("resume: %v", err)
	}
	if read("after resume").Paused {
		t.Error("the schedule is still paused after resuming it")
	}
}

// A schedule declared paused is stored that way rather than starting to fire.
func TestPausedOnCreatePreservesState(t *testing.T) {
	rt, agent, ctx := s21Runtime(t)
	if err := rt.ReconcileSchedules(ctx, agent, []ai.Schedule{
		{Name: "silent", Cron: "0 0 9 * * ?", Paused: true}}); err != nil {
		t.Fatalf("reconcile: %v", err)
	}
	s, err := rt.GetSchedule(ctx, agent, "silent")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if !s.Paused {
		t.Error("a schedule declared paused was created running")
	}
}

// Deleting one removes it from the agent's schedules.
func TestDeleteRemoves(t *testing.T) {
	rt, agent, ctx := s21Runtime(t)
	if err := rt.ReconcileSchedules(ctx, agent, []ai.Schedule{{Name: "d", Cron: "0 * * * * ?"}}); err != nil {
		t.Fatalf("reconcile: %v", err)
	}
	if err := rt.DeleteSchedule(ctx, agent, "d"); err != nil {
		t.Fatalf("delete: %v", err)
	}
	if got := byName(t, rt, ctx, agent); len(got) != 0 {
		t.Errorf("the schedule is still listed after deleting it: %v", got)
	}
}
