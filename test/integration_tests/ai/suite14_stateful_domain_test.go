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
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
	taskmodel "github.com/conductor-sdk/conductor-go/sdk/model"
)

// The Python SDK's e2e/test_suite14_stateful_domain.py as Go tests, one per
// Python test and under the same names. A stateful run routes its tasks to a
// queue of its own and this process polls that queue, so the run's calls reach
// its own workers and two runs never take each other's tasks. A run that is
// not stateful uses the shared queue and records no domain at all.

type s14MessageIn struct {
	Message string
}

type s14TaskIn struct {
	Task string
}

func s14EchoTool() ai.ToolDef {
	return tool.Func("echo_tool", "Return the message with a deterministic prefix.",
		func(_ context.Context, in s14MessageIn) (string, error) { return "ECHO:" + in.Message, nil })
}

func s14StatefulEcho() ai.ToolDef {
	return tool.Func("stateful_echo", "A stateful tool that echoes with a prefix.",
		func(_ context.Context, in s14MessageIn) (string, error) { return "STATEFUL_ECHO:" + in.Message, nil },
		tool.Stateful())
}

func s14SwarmTool() ai.ToolDef {
	return tool.Func("swarm_tool", "Perform a task and return a marker.",
		func(_ context.Context, in s14TaskIn) (string, error) { return "SWARM_RESULT:" + in.Task, nil },
		tool.Stateful())
}

// shouldStopOnEcho is the suite's stop_when predicate: stop once the tool's
// prefix appears in the turn's result.
func shouldStopOnEcho(_ context.Context, state ai.StopWhenState) (bool, error) {
	return strings.Contains(state.Result, "ECHO:"), nil
}

// allTasksDeep is the suite's _get_all_tasks: the run's tasks plus those of
// every sub-workflow that completed.
func allTasksDeep(t *testing.T, executionID string) []taskmodel.Task {
	t.Helper()
	wf := getWorkflow(t, executionID)
	out := append([]taskmodel.Task(nil), wf.Tasks...)
	for _, task := range wf.Tasks {
		if task.TaskType != "SUB_WORKFLOW" || taskStatus(task) != "COMPLETED" {
			continue
		}
		sub := task.SubWorkflowId
		if sub == "" {
			sub, _ = task.OutputData["subWorkflowId"].(string)
		}
		if sub != "" {
			out = append(out, allTasksDeep(t, sub)...)
		}
	}
	return out
}

// tasksNamed is the suite's _find_tasks_by_type: a substring match on the
// task's definition name.
func tasksNamed(tasks []taskmodel.Task, name string) []taskmodel.Task {
	var out []taskmodel.Task
	for _, task := range tasks {
		if strings.Contains(task.TaskDefName, name) {
			out = append(out, task)
		}
	}
	return out
}

// scheduledTasks are the tasks nothing has polled. One left behind means a
// worker is listening on the wrong queue.
func scheduledTasks(tasks []taskmodel.Task) []taskmodel.Task {
	var out []taskmodel.Task
	for _, task := range tasks {
		if taskStatus(task) == "SCHEDULED" {
			out = append(out, task)
		}
	}
	return out
}

func assertNothingScheduled(t *testing.T, tasks []taskmodel.Task) {
	t.Helper()
	for _, task := range scheduledTasks(tasks) {
		t.Errorf("task %s (%s) was never polled: status=%s pollCount=%d domain=%q",
			task.ReferenceTaskName, task.TaskDefName, task.Status, task.PollCount, task.Domain)
	}
}

// domainsOf is the distinct set of domains a workflow routed its tasks to.
func domainsOf(t *testing.T, executionID string) map[string]bool {
	t.Helper()
	out := map[string]bool{}
	for _, domain := range getWorkflow(t, executionID).TaskToDomain {
		if domain != "" {
			out[domain] = true
		}
	}
	return out
}

// A stateful agent's tool task is picked up and completed by a worker polling
// this run's own domain, rather than sitting in a queue nothing listens to.
func TestStatefulToolCompletes(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_s14_stateful_tool", Model: model(t), Stateful: true, MaxTurns: 3,
		Instructions: "You have an echo_tool. Call echo_tool with message='hello'. " +
			"Then respond with what the tool returned.",
		Tools: ai.Tools(s14EchoTool())}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, "Call the echo tool with hello")
	assertRunCompleted(t, res, "Stateful tool")

	wf := getWorkflow(t, res.ExecutionID)
	if len(wf.TaskToDomain) == 0 {
		t.Fatal("the run recorded no task domain, so nothing routed its tasks to this process")
	}
	tasks := allTasksDeep(t, res.ExecutionID)
	echoes := tasksNamed(tasks, "echo_tool")
	if len(echoes) == 0 {
		t.Fatal("echo_tool never ran")
	}
	for _, task := range echoes {
		if taskStatus(task) != "COMPLETED" {
			t.Errorf("echo_tool status = %q (domain=%q pollCount=%d)", taskStatus(task), task.Domain, task.PollCount)
		}
		if want := wf.TaskToDomain["echo_tool"]; want != "" && task.Domain != want {
			t.Errorf("echo_tool ran in domain %q, want the run's own %q", task.Domain, want)
		}
	}
	assertNothingScheduled(t, tasks)
}

// A stop-when predicate is a task of its own, and it too must be polled from
// the run's domain.
func TestStatefulStopWhenCompletes(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_s14_stateful_stop", Model: model(t), Stateful: true, MaxTurns: 5,
		Instructions: "Call echo_tool with message='stop_test'. Then report the tool's response.",
		Tools:        ai.Tools(s14EchoTool()), StopWhen: shouldStopOnEcho}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, "Call echo_tool with stop_test")
	assertRunCompleted(t, res, "Stateful stop_when")

	tasks := allTasksDeep(t, res.ExecutionID)
	stops := tasksNamed(tasks, "stop_when")
	if len(stops) == 0 {
		t.Fatal("the stop_when task never ran")
	}
	var completed int
	for _, task := range stops {
		if taskStatus(task) == "COMPLETED" {
			completed++
		}
	}
	if completed == 0 {
		t.Errorf("no stop_when task completed: %v", stops)
	}
	if len(domainsOf(t, res.ExecutionID)) == 0 {
		t.Error("the run recorded no task domain")
	}
	assertNothingScheduled(t, tasks)
}

// The Python suite disables this one: a stateful swarm handoff does not
// reliably complete in its domain. Kept for parity, skipped for the same
// reason.
func TestStatefulSwarmHandoffCompletes(t *testing.T) {
	t.Skip("disabled in the Python suite too: a stateful swarm handoff does not reliably complete in domain")
	_ = s14SwarmTool
}

// An agent's own statefulness reaches every tool it has, flagged or not, and
// they all run in one domain.
func TestStatefulMixedTools(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_s14_mixed_tools", Model: model(t), Stateful: true, MaxTurns: 5,
		Instructions: "You have two tools. First call echo_tool with message='regular'. " +
			"Then call stateful_echo with message='stateful'. Report both results.",
		Tools: ai.Tools(s14EchoTool(), s14StatefulEcho())}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, "Call both tools")
	assertRunCompleted(t, res, "Stateful mixed tools")

	if len(domainsOf(t, res.ExecutionID)) == 0 {
		t.Fatal("the run recorded no task domain")
	}
	tasks := allTasksDeep(t, res.ExecutionID)
	plain, flagged := tasksNamed(tasks, "echo_tool"), tasksNamed(tasks, "stateful_echo")
	if len(plain) == 0 || len(flagged) == 0 {
		t.Fatalf("both tools should have run; echo_tool=%d stateful_echo=%d", len(plain), len(flagged))
	}
	domains := func(tasks []taskmodel.Task) map[string]bool {
		out := map[string]bool{}
		for _, task := range tasks {
			if taskStatus(task) != "COMPLETED" {
				t.Errorf("%s status = %q", task.TaskDefName, taskStatus(task))
			}
			if task.Domain != "" {
				out[task.Domain] = true
			}
		}
		return out
	}
	plainDomains, flaggedDomains := domains(plain), domains(flagged)
	if len(plainDomains) > 0 && len(flaggedDomains) > 0 {
		for domain := range plainDomains {
			if !flaggedDomains[domain] {
				t.Errorf("the unflagged tool ran in %q, which the flagged one did not use: %v vs %v",
					domain, plainDomains, flaggedDomains)
			}
		}
	}
	assertNothingScheduled(t, tasks)
}

// Each stateful run mints its own domain, so a second run neither reuses nor
// steals the first's. The two runs are sequential, as in the Python suite.
func TestConcurrentStatefulIsolation(t *testing.T) {
	makeAgent := func(suffix string) *ai.Agent {
		return &ai.Agent{Name: "e2e_s14_concurrent_" + suffix, Model: model(t), Stateful: true, MaxTurns: 3,
			Instructions: "Call echo_tool with message='concurrent_test'. Respond with the tool result.",
			Tools:        ai.Tools(s14EchoTool())}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// One runtime per run, the first shut down before the second starts, as
	// the Python suite does: its workers are bound to the first run's domain.
	first := newRuntime(t)
	one := runTolerant(t, first, ctx, makeAgent("a"), "Run 1: call echo_tool")
	assertRunCompleted(t, one, "Run 1")
	first.Shutdown()

	second := newRuntime(t)
	two := runTolerant(t, second, ctx, makeAgent("b"), "Run 2: call echo_tool")
	assertRunCompleted(t, two, "Run 2")

	if one.ExecutionID == two.ExecutionID {
		t.Fatal("both runs reported the same execution")
	}
	oneDomains, twoDomains := domainsOf(t, one.ExecutionID), domainsOf(t, two.ExecutionID)
	if len(oneDomains) == 0 || len(twoDomains) == 0 {
		t.Fatalf("a run recorded no domain: run 1 %v, run 2 %v", oneDomains, twoDomains)
	}
	for domain := range oneDomains {
		if twoDomains[domain] {
			t.Errorf("both runs used domain %q; each run must have its own: %v vs %v",
				domain, oneDomains, twoDomains)
		}
	}
	assertNothingScheduled(t, allTasksDeep(t, one.ExecutionID))
	assertNothingScheduled(t, allTasksDeep(t, two.ExecutionID))
}

// The domain plumbing is opt-in: an ordinary agent records none and its tools
// still run, from the shared queue.
func TestNonStatefulNoDomain(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_s14_non_stateful", Model: model(t), MaxTurns: 3,
		Instructions: "Call echo_tool with message='non_stateful'. Respond with the result.",
		Tools:        ai.Tools(s14EchoTool())}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, "Call echo_tool")
	assertRunCompleted(t, res, "Non-stateful")

	if domains := domainsOf(t, res.ExecutionID); len(domains) != 0 {
		t.Errorf("an ordinary agent routed its tasks to %v; it should use the shared queue", domains)
	}
	echoes := tasksNamed(allTasksDeep(t, res.ExecutionID), "echo_tool")
	if len(echoes) == 0 {
		t.Fatal("echo_tool never ran")
	}
	for _, task := range echoes {
		if taskStatus(task) != "COMPLETED" {
			t.Errorf("echo_tool status = %q", taskStatus(task))
		}
		if task.Domain != "" {
			t.Errorf("echo_tool ran in domain %q, want none", task.Domain)
		}
	}
}
