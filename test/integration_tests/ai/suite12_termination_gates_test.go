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
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	taskmodel "github.com/conductor-sdk/conductor-go/sdk/model"
)

// The Python SDK's e2e/test_suite12_termination_gates.py as Go tests, one per
// Python test and under the same names. A termination condition stops an
// agent's loop before its turn limit, and a gate on a pipeline stage compiles
// into the switch that decides whether the pipeline goes on.

// loopIterations is the suite's _get_loop_iterations: how many times the
// agent's DO_WHILE ran.
func loopIterations(wf taskmodel.Workflow) int {
	for _, task := range wf.Tasks {
		if task.TaskType == "DO_WHILE" {
			if n, ok := task.OutputData["iteration"].(float64); ok {
				return int(n)
			}
		}
	}
	return 0
}

// taskByRef is the suite's _find_task_by_ref: an exact reference, or one the
// server suffixed with "__<n>" for a loop iteration.
func taskByRef(wf taskmodel.Workflow, ref string) *taskmodel.Task {
	for i := range wf.Tasks {
		got := wf.Tasks[i].ReferenceTaskName
		if got == ref || strings.HasPrefix(got, ref+"__") {
			return &wf.Tasks[i]
		}
	}
	return nil
}

// resultOutput is the suite's _task_output: a task's output, unwrapped from
// the "result" key the worker nests it under.
func resultOutput(task *taskmodel.Task) map[string]any {
	if task == nil {
		return map[string]any{}
	}
	if inner, ok := task.OutputData["result"].(map[string]any); ok {
		return inner
	}
	return task.OutputData
}

func s12EchoTool() ai.ToolDef { return echoTool() }

// A sentinel in the reply ends the loop, so the agent stops before its turn
// limit rather than running to it.
func TestTextMentionTerminatesEarly(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name: "e2e_s12_text_term", Model: model(t), MaxTurns: 3,
		Instructions: "You MUST include the exact text TASK_COMPLETE in every response. " +
			"Answer the user's question and always end with TASK_COMPLETE.",
		Tools:       []ai.ToolDef{s12EchoTool()},
		Termination: ai.TextMentionTermination{Text: "TASK_COMPLETE"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, "Say hello.")

	if res.ExecutionID == "" {
		t.Fatal("no execution id")
	}
	if res.Status != ai.StatusCompleted && res.Status != ai.StatusTerminated {
		t.Fatalf("status = %q, want COMPLETED or TERMINATED (error=%q)", res.Status, res.Error)
	}
	if n := loopIterations(getWorkflow(t, res.ExecutionID)); n > 3 {
		t.Errorf("the loop ran %d times, past the turn limit of 3", n)
	}
}

// A one-message limit makes the server's termination task stop the loop, and
// the task says so.
func TestMaxMessageTerminatesAtLimit(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name: "e2e_s12_max_msg", Model: model(t), MaxTurns: 25,
		Instructions: "You are a counting assistant. You MUST use the echo_tool for every " +
			"step — never answer directly. Call echo_tool once per number with " +
			"{text: \"<number>\"}. After each tool result, call echo_tool again " +
			"for the next number. Continue until told to stop.",
		Tools:       []ai.ToolDef{s12EchoTool()},
		Termination: ai.MaxMessageTermination{MaxMessages: 1},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, "Say hello.")

	if res.ExecutionID == "" {
		t.Fatal("no execution id")
	}
	if res.Status != ai.StatusCompleted && res.Status != ai.StatusTerminated {
		t.Fatalf("status = %q, want COMPLETED or TERMINATED (error=%q)", res.Status, res.Error)
	}
	wf := getWorkflow(t, res.ExecutionID)
	task := taskByRef(wf, "e2e_s12_max_msg_termination")
	if task == nil {
		var refs []string
		for _, tk := range wf.Tasks {
			refs = append(refs, tk.ReferenceTaskName)
		}
		t.Fatalf("no termination task; references: %v", refs)
	}
	if taskStatus(*task) != "COMPLETED" {
		t.Fatalf("termination task status = %q (reason=%q)", taskStatus(*task), task.ReasonForIncompletion)
	}
	out := resultOutput(task)
	if out["should_continue"] != false {
		t.Errorf("termination task should_continue = %v, want false: %v", out["should_continue"], out)
	}
	if fmt.Sprint(out["reason"]) == "" || out["reason"] == nil {
		t.Errorf("termination task gave no reason: %v", out)
	}
	if n := loopIterations(wf); n >= 25 {
		t.Errorf("the loop ran %d times; the termination did not stop it short of the turn limit", n)
	}
}

// s12Pipeline is Python's `checker >> fixer`: a sequential parent named after
// its two children. Go has no operator for it, so the test builds it by hand.
func s12Pipeline(t *testing.T, suffix string) *ai.Agent {
	t.Helper()
	m := model(t)
	checker := &ai.Agent{Name: "e2e_s12_checker_" + suffix, Model: m, MaxTurns: 2,
		Instructions: "Check for issues.", Gate: ai.TextGate{Text: "STOP"}}
	fixer := &ai.Agent{Name: "e2e_s12_fixer_" + suffix, Model: m, MaxTurns: 2,
		Instructions: "Fix any issues found.", Tools: []ai.ToolDef{s12EchoTool()}}
	return &ai.Agent{Name: checker.Name + "_" + fixer.Name, Model: m,
		Agents: []*ai.Agent{checker, fixer}, Strategy: ai.StrategySequential}
}

// A gate on a pipeline stage compiles into a gate task and the switch that
// reads it. Compile only, no model involved.
func TestTextGateStopsPipeline(t *testing.T) {
	rt := newRuntime(t)
	tasks := allTasks(workflowDef(t, planAgent(t, rt, s12Pipeline(t, "stop"))))

	var gates int
	for _, task := range tasks {
		if strings.Contains(strings.ToLower(fmt.Sprint(task["taskReferenceName"])), "gate") {
			gates++
		}
	}
	if gates == 0 {
		var refs []string
		for _, task := range tasks {
			refs = append(refs, fmt.Sprint(task["taskReferenceName"]))
		}
		t.Errorf("no gate task in the compiled pipeline; references: %v", refs)
	}
	if types := taskTypes(tasks); !types["SWITCH"] {
		var names []string
		for typ := range types {
			names = append(names, typ)
		}
		sort.Strings(names)
		t.Errorf("no SWITCH task in the compiled pipeline; types: %v", names)
	}
}

// The gate's switch carries a populated continue branch, so the stage after
// the gate is reachable.
func TestTextGateSwitchHasContinueAndStop(t *testing.T) {
	rt := newRuntime(t)
	def := workflowDef(t, planAgent(t, rt, s12Pipeline(t, "pass")))

	var switches []map[string]any
	for _, raw := range asList(def["tasks"]) {
		if task, ok := raw.(map[string]any); ok && task["type"] == "SWITCH" {
			switches = append(switches, task)
		}
	}
	if len(switches) == 0 {
		t.Fatal("no top-level SWITCH task in the compiled pipeline")
	}
	cases, _ := switches[0]["decisionCases"].(map[string]any)
	if cases == nil {
		t.Fatalf("the SWITCH has no decision cases: %v", switches[0])
	}
	branch, ok := cases["continue"]
	if !ok {
		t.Fatalf("the SWITCH has no continue branch; cases: %v", keys(cases))
	}
	if len(asList(branch)) == 0 {
		t.Errorf("the continue branch is empty, so the next stage is unreachable")
	}
}

// A model the server does not know ends the run, rather than passing silently.
func TestInvalidModelFails(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_s12_bad_model", Model: "nonexistent/xyz-model-does-not-exist",
		Instructions: "This agent should never execute successfully.",
		Tools:        []ai.ToolDef{s12EchoTool()}}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, "Hello.")

	if res.Status != ai.StatusFailed && res.Status != ai.StatusTerminated {
		t.Errorf("status = %q, want FAILED or TERMINATED for an unknown model", res.Status)
	}
}
