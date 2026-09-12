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
	"strings"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// Schedule runs an agent on a cron cadence. It is scoped to one agent: its
// Name is unique among that agent's schedules, and the wire name the server
// stores is "<agent>-<Name>", so two agents may reuse the same short name.
// It is the counterpart of the Python SDK's Schedule.
type Schedule struct {
	// Name identifies the schedule within its agent. Required.
	Name string
	// Cron is the schedule expression, for example "0 0 * * *". Required.
	Cron string
	// Timezone is the zone the cron runs in. Empty means UTC.
	Timezone string
	// Input is the prompt/input the scheduled run starts the agent with.
	Input map[string]any
	// Catchup runs missed occurrences when the scheduler was down. Off by default.
	Catchup bool
	// Paused creates the schedule paused.
	Paused bool
	// StartAt and EndAt bound when the schedule is active, epoch millis. Zero
	// means unbounded.
	StartAt int64
	EndAt   int64
	// Description is optional free text.
	Description string
}

// Validate reports the first problem with the schedule.
func (s Schedule) Validate() error {
	if strings.TrimSpace(s.Name) == "" {
		return fmt.Errorf("schedule name is required")
	}
	if strings.TrimSpace(s.Cron) == "" {
		return fmt.Errorf("schedule %q: cron is required", s.Name)
	}
	if s.StartAt != 0 && s.EndAt != 0 && s.StartAt >= s.EndAt {
		return fmt.Errorf("schedule %q: StartAt must be before EndAt", s.Name)
	}
	return nil
}

// scheduleWireName is the server-stored name for a schedule: agent-prefixed so
// schedules of different agents never collide.
func scheduleWireName(agentName, short string) string {
	return agentName + "-" + short
}

// schedulePrefix is the prefix every one of an agent's schedules carries.
func schedulePrefix(agentName string) string {
	return agentName + "-"
}

// saveRequest builds the wire request: the agent-prefixed name, the schedule
// settings, and a start-workflow request that runs the agent with this
// schedule's input. Matches the Python SDK's _to_save_request.
func (s Schedule) saveRequest(agentName string) model.SaveScheduleRequest {
	input := s.Input
	if input == nil {
		input = map[string]any{}
	}
	tz := s.Timezone
	if tz == "" {
		tz = "UTC"
	}
	return model.SaveScheduleRequest{
		Name:                        scheduleWireName(agentName, s.Name),
		CronExpression:              s.Cron,
		ZoneId:                      tz,
		RunCatchupScheduleInstances: s.Catchup,
		Paused:                      s.Paused,
		ScheduleStartTime:           s.StartAt,
		ScheduleEndTime:             s.EndAt,
		Description:                 s.Description,
		StartWorkflowRequest: &model.StartWorkflowRequest{
			Name:  agentName,
			Input: input,
		},
	}
}

// SaveSchedule creates or updates one schedule for an agent.
func (r *Runtime) SaveSchedule(ctx context.Context, agentName string, s Schedule) error {
	if err := s.Validate(); err != nil {
		return err
	}
	if _, _, err := r.scheduler.SaveSchedule(ctx, s.saveRequest(agentName)); err != nil {
		return fmt.Errorf("save schedule %q: %w", s.Name, err)
	}
	return nil
}

// GetSchedule reads one of an agent's schedules by its short name.
func (r *Runtime) GetSchedule(ctx context.Context, agentName, name string) (*Schedule, error) {
	ws, _, err := r.scheduler.GetSchedule(ctx, scheduleWireName(agentName, name))
	if err != nil {
		return nil, fmt.Errorf("get schedule %q: %w", name, err)
	}
	s := scheduleFromWire(agentName, ws)
	return &s, nil
}

// ListSchedules returns all of an agent's schedules.
func (r *Runtime) ListSchedules(ctx context.Context, agentName string) ([]Schedule, error) {
	all, _, err := r.scheduler.GetAllSchedules(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("list schedules for %q: %w", agentName, err)
	}
	prefix := schedulePrefix(agentName)
	var out []Schedule
	for _, ws := range all {
		if !strings.HasPrefix(ws.Name, prefix) {
			continue
		}
		out = append(out, Schedule{
			Name:        strings.TrimPrefix(ws.Name, prefix),
			Cron:        ws.CronExpression,
			Timezone:    ws.ZoneId,
			Catchup:     ws.RunCatchupScheduleInstances,
			Paused:      ws.Paused,
			StartAt:     ws.ScheduleStartTime,
			EndAt:       ws.ScheduleEndTime,
			Description: ws.Description,
		})
	}
	return out, nil
}

// DeleteSchedule removes one of an agent's schedules by its short name.
func (r *Runtime) DeleteSchedule(ctx context.Context, agentName, name string) error {
	if _, _, err := r.scheduler.DeleteSchedule(ctx, scheduleWireName(agentName, name)); err != nil {
		return fmt.Errorf("delete schedule %q: %w", name, err)
	}
	return nil
}

// PauseSchedule pauses one of an agent's schedules by its short name.
func (r *Runtime) PauseSchedule(ctx context.Context, agentName, name string) error {
	if _, _, err := r.scheduler.PauseSchedule(ctx, scheduleWireName(agentName, name)); err != nil {
		return fmt.Errorf("pause schedule %q: %w", name, err)
	}
	return nil
}

// ResumeSchedule resumes one of an agent's paused schedules by its short name.
func (r *Runtime) ResumeSchedule(ctx context.Context, agentName, name string) error {
	if _, _, err := r.scheduler.ResumeSchedule(ctx, scheduleWireName(agentName, name)); err != nil {
		return fmt.Errorf("resume schedule %q: %w", name, err)
	}
	return nil
}

// ReconcileSchedules makes an agent's schedules match desired exactly: it
// creates or updates every schedule in desired and deletes any other schedule
// that belongs to this agent. It is the counterpart of the Python SDK's
// scheduler reconcile, and the way to declare an agent's schedules from a
// deploy step. A nil desired is a no-op; an empty desired deletes them all.
func (r *Runtime) ReconcileSchedules(ctx context.Context, agentName string, desired []Schedule) error {
	if desired == nil {
		return nil
	}
	seen := map[string]bool{}
	for _, s := range desired {
		if err := s.Validate(); err != nil {
			return err
		}
		if seen[s.Name] {
			return fmt.Errorf("duplicate schedule name %q", s.Name)
		}
		seen[s.Name] = true
	}
	existing, err := r.ListSchedules(ctx, agentName)
	if err != nil {
		return err
	}
	for _, e := range existing {
		if !seen[e.Name] {
			if err := r.DeleteSchedule(ctx, agentName, e.Name); err != nil {
				return err
			}
		}
	}
	for _, s := range desired {
		if err := r.SaveSchedule(ctx, agentName, s); err != nil {
			return err
		}
	}
	return nil
}

func scheduleFromWire(agentName string, ws model.WorkflowSchedule) Schedule {
	return Schedule{
		Name:        strings.TrimPrefix(ws.Name, schedulePrefix(agentName)),
		Cron:        ws.CronExpression,
		Timezone:    ws.ZoneId,
		Catchup:     ws.RunCatchupScheduleInstances,
		Paused:      ws.Paused,
		StartAt:     ws.ScheduleStartTime,
		EndAt:       ws.ScheduleEndTime,
		Description: ws.Description,
	}
}
