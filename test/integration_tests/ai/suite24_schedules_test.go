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
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// The schedule portion of the Python SDK's e2e/test_suite24_agent_client.py:
// declare an agent's schedules, read them back, and purge them, through the
// runtime's schedule surface. It calls no model, so it runs live only, with no
// recording, and skips when the server has no scheduler, matching the Python
// suite's _scheduler_available gate.

func requireScheduler(t *testing.T) {
	t.Helper()
	base := os.Getenv("CONDUCTOR_SERVER_URL")
	req, _ := http.NewRequest(http.MethodGet, strings.TrimRight(base, "/")+"/scheduler/schedules", nil)
	resp, err := (&http.Client{Timeout: 5 * time.Second}).Do(req)
	if err != nil {
		t.Skipf("scheduler not reachable: %v", err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Skipf("scheduler not available (HTTP %d)", resp.StatusCode)
	}
}

func TestScheduleReconcileListPurge(t *testing.T) {
	rt := newRuntime(t)
	requireScheduler(t)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	const agent = "e2e_s24_sched"
	// Always leave the server clean, even on failure.
	defer rt.ReconcileSchedules(ctx, agent, []ai.Schedule{})

	// Declare two schedules.
	desired := []ai.Schedule{
		{Name: "nightly", Cron: "0 0 * * *", Input: map[string]any{"prompt": "nightly"}},
		{Name: "hourly", Cron: "0 * * * *", Timezone: "America/New_York"},
	}
	if err := rt.ReconcileSchedules(ctx, agent, desired); err != nil {
		t.Fatalf("reconcile: %v", err)
	}

	got, err := rt.ListSchedules(ctx, agent)
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	names := map[string]ai.Schedule{}
	for _, s := range got {
		names[s.Name] = s
	}
	for _, want := range []string{"nightly", "hourly"} {
		if _, ok := names[want]; !ok {
			t.Errorf("schedule %q not listed; got %v", want, keysOf(names))
		}
	}
	// The short name round-trips (the agent prefix is stripped on read).
	if s := names["hourly"]; s.Cron != "0 * * * *" || s.Timezone != "America/New_York" {
		t.Errorf("hourly read back as cron=%q tz=%q", s.Cron, s.Timezone)
	}

	// Reconcile to a subset: hourly is pruned, nightly stays.
	if err := rt.ReconcileSchedules(ctx, agent, []ai.Schedule{desired[0]}); err != nil {
		t.Fatalf("reconcile subset: %v", err)
	}
	got, _ = rt.ListSchedules(ctx, agent)
	names = map[string]ai.Schedule{}
	for _, s := range got {
		names[s.Name] = s
	}
	if _, ok := names["hourly"]; ok {
		t.Error("hourly should have been pruned")
	}
	if _, ok := names["nightly"]; !ok {
		t.Error("nightly should remain")
	}

	// Purge and confirm none remain.
	if err := rt.ReconcileSchedules(ctx, agent, []ai.Schedule{}); err != nil {
		t.Fatalf("purge: %v", err)
	}
	if got, _ = rt.ListSchedules(ctx, agent); len(got) != 0 {
		t.Errorf("after purge, %d schedules remain", len(got))
	}
}

func keysOf(m map[string]ai.Schedule) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
