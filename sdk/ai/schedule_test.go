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
)

func TestScheduleValidate(t *testing.T) {
	cases := []struct {
		name    string
		s       Schedule
		wantErr string
	}{
		{"ok", Schedule{Name: "nightly", Cron: "0 0 * * *"}, ""},
		{"no name", Schedule{Cron: "0 0 * * *"}, "name is required"},
		{"no cron", Schedule{Name: "n"}, "cron is required"},
		{"start after end", Schedule{Name: "n", Cron: "0 0 * * *", StartAt: 10, EndAt: 5}, "StartAt must be before EndAt"},
	}
	for _, c := range cases {
		err := c.s.Validate()
		if c.wantErr == "" && err != nil {
			t.Errorf("%s: unexpected error %v", c.name, err)
		}
		if c.wantErr != "" && (err == nil || !strings.Contains(err.Error(), c.wantErr)) {
			t.Errorf("%s: err = %v, want %q", c.name, err, c.wantErr)
		}
	}
}

func TestSaveScheduleWire(t *testing.T) {
	rec := &capture{}
	rt := captureServer(t, rec, nil)
	err := rt.SaveSchedule(context.Background(), "greeter", Schedule{
		Name: "nightly", Cron: "0 0 * * *", Timezone: "America/New_York",
		Input: map[string]any{"prompt": "run the report"}, Catchup: true, Description: "nightly report",
	})
	if err != nil {
		t.Fatal(err)
	}
	body, ok := rec.find("/api/scheduler/schedules")
	if !ok {
		t.Fatal("no POST to /scheduler/schedules")
	}
	if body["name"] != "greeter-nightly" {
		t.Errorf("wire name = %v, want greeter-nightly", body["name"])
	}
	if body["cronExpression"] != "0 0 * * *" || body["zoneId"] != "America/New_York" {
		t.Errorf("cron/zone = %v / %v", body["cronExpression"], body["zoneId"])
	}
	if body["runCatchupScheduleInstances"] != true || body["description"] != "nightly report" {
		t.Errorf("catchup/description = %v / %v", body["runCatchupScheduleInstances"], body["description"])
	}
	swr, ok := body["startWorkflowRequest"].(map[string]any)
	if !ok || swr["name"] != "greeter" {
		t.Fatalf("startWorkflowRequest = %v, want the agent name", body["startWorkflowRequest"])
	}
	if in, _ := swr["input"].(map[string]any); in["prompt"] != "run the report" {
		t.Errorf("scheduled input = %v", swr["input"])
	}
	// Empty timezone defaults to UTC.
	rec2 := &capture{}
	rt2 := captureServer(t, rec2, nil)
	_ = rt2.SaveSchedule(context.Background(), "a", Schedule{Name: "n", Cron: "* * * * *"})
	if b, _ := rec2.find("/api/scheduler/schedules"); b["zoneId"] != "UTC" {
		t.Errorf("default zone = %v, want UTC", b["zoneId"])
	}
}

func TestReconcileSchedules(t *testing.T) {
	rec := &capture{}
	// The server already holds two schedules for this agent and one for another.
	rt := captureServer(t, rec, map[string]func() map[string]any{
		"/api/scheduler/schedules": func() map[string]any { return nil }, // POST save -> {}
	})
	// GetAllSchedules is a GET returning a list; captureServer's per-path
	// handler returns an object, so serve the list from a dedicated server.
	rt = schedulerListServer(t, rec, []string{"greeter-old", "greeter-keep", "other-x"})

	err := rt.ReconcileSchedules(context.Background(), "greeter", []Schedule{
		{Name: "keep", Cron: "0 0 * * *"},
		{Name: "new", Cron: "0 12 * * *"},
	})
	if err != nil {
		t.Fatal(err)
	}
	// greeter-old is pruned; other-x (different agent) is left alone.
	if !rec.sawDelete("/api/scheduler/schedules/greeter-old") {
		t.Error("did not delete the removed schedule greeter-old")
	}
	if rec.sawDelete("/api/scheduler/schedules/other-x") {
		t.Error("deleted another agent's schedule other-x")
	}
	if rec.sawDelete("/api/scheduler/schedules/greeter-keep") {
		t.Error("deleted a schedule that is still desired")
	}
	// keep and new are upserted (two POSTs to /scheduler/schedules).
	saves := rec.countPost("/api/scheduler/schedules")
	if saves != 2 {
		t.Errorf("saved %d schedules, want 2 (keep, new)", saves)
	}
	// A nil desired is a no-op: nothing deleted or saved.
	rec2 := &capture{}
	rt2 := schedulerListServer(t, rec2, []string{"greeter-old"})
	if err := rt2.ReconcileSchedules(context.Background(), "greeter", nil); err != nil {
		t.Fatal(err)
	}
	if rec2.sawDelete("/api/scheduler/schedules/greeter-old") || rec2.countPost("/api/scheduler/schedules") != 0 {
		t.Error("nil desired should be a no-op")
	}
}
