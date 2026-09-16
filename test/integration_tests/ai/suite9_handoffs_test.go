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
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
	taskmodel "github.com/conductor-sdk/conductor-go/sdk/model"
)

// The Python SDK's e2e/test_suite9_handoffs.py as Go tests, one per Python
// test and under the same name: every multi-agent strategy compiles, and the
// sequential, parallel, handoff, router and swarm strategies run their
// children as sub-workflows the way the Python suite checks.

type s9ExprIn struct {
	Expr string `json:"expr"`
}

type s9TextIn struct {
	Text string `json:"text"`
}

type s9QueryIn struct {
	Query string `json:"query"`
}

// evalExpr stands in for the Python tool's eval() on the small arithmetic
// the suite asks for: two integers and one of + - * /.
func evalExpr(expr string) (string, error) {
	s := strings.ReplaceAll(expr, " ", "")
	for _, op := range []string{"+", "-", "*", "/"} {
		if i := strings.LastIndex(s, op); i > 0 {
			a, errA := strconv.Atoi(s[:i])
			b, errB := strconv.Atoi(s[i+1:])
			if errA != nil || errB != nil {
				break
			}
			switch op {
			case "+":
				return strconv.Itoa(a + b), nil
			case "-":
				return strconv.Itoa(a - b), nil
			case "*":
				return strconv.Itoa(a * b), nil
			case "/":
				if b == 0 {
					return "", fmt.Errorf("division by zero")
				}
				if a%b == 0 {
					return strconv.Itoa(a / b), nil
				}
				return strconv.FormatFloat(float64(a)/float64(b), 'g', -1, 64), nil
			}
		}
	}
	return "", fmt.Errorf("cannot evaluate %q", expr)
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

var (
	doMath = tool.Func("do_math", "Evaluate a math expression.",
		func(_ context.Context, in s9ExprIn) (string, error) {
			v, err := evalExpr(in.Expr)
			if err != nil {
				return "", err
			}
			return "math_result:" + in.Expr + "=" + v, nil
		})
	doText = tool.Func("do_text", "Reverse a string.",
		func(_ context.Context, in s9TextIn) (string, error) {
			return "text_result:" + reverseString(in.Text), nil
		})
	// do_data and the data agent exist in the Python suite but no test uses them.
	doData = tool.Func("do_data", "Echo a data query.",
		func(_ context.Context, in s9QueryIn) (string, error) { return "data_result:" + in.Query, nil })
)

func s9MathAgent(m string) *ai.Agent {
	return &ai.Agent{Name: "math_agent", Model: m, Tools: []ai.ToolDef{doMath}, MaxTurns: 3,
		Instructions: "You are a math agent. When asked to compute something, call do_math with the expression. " +
			"For example, for \"3+4\" call do_math with expr=\"3+4\". Only handle math operations — ignore non-math requests. " +
			"If there is nothing to compute, just respond with a summary."}
}

func s9TextAgent(m string) *ai.Agent {
	return &ai.Agent{Name: "text_agent", Model: m, Tools: []ai.ToolDef{doText}, MaxTurns: 3,
		Instructions: "You are a text agent. When asked to reverse text, call do_text with the text. " +
			"For example, for \"hello\" call do_text with text=\"hello\". " +
			"If there is nothing to reverse, just respond with a summary of what you received."}
}

var _ = doData

// subWorkflowRefs returns the references of the SUB_WORKFLOW tasks, and of
// the completed ones.
func subWorkflowRefs(wf taskmodel.Workflow) (all, completed []string) {
	for _, task := range tasksOfType(wf, "SUB_WORKFLOW") {
		all = append(all, task.ReferenceTaskName)
		if taskStatus(task) == "COMPLETED" {
			completed = append(completed, task.ReferenceTaskName)
		}
	}
	return all, completed
}

func anyContains(refs []string, sub string, fold bool) bool {
	for _, r := range refs {
		if fold {
			r, sub = strings.ToLower(r), strings.ToLower(sub)
		}
		if strings.Contains(r, sub) {
			return true
		}
	}
	return false
}

func TestAllStrategiesCompile(t *testing.T) {
	rt := newRuntime(t)
	m := model(t)
	childA := &ai.Agent{Name: "child_a", Model: m, Instructions: "Child A."}
	childB := &ai.Agent{Name: "child_b", Model: m, Instructions: "Child B."}
	routerLead := &ai.Agent{Name: "router_lead", Model: m, Instructions: "Route tasks."}

	cases := []struct {
		name     string
		strategy ai.Strategy
		router   *ai.Agent
	}{
		{"handoff", ai.StrategyHandoff, nil}, {"sequential", ai.StrategySequential, nil},
		{"parallel", ai.StrategyParallel, nil}, {"router", ai.StrategyRouter, routerLead},
		{"round_robin", ai.StrategyRoundRobin, nil}, {"random", ai.StrategyRandom, nil},
		{"swarm", ai.StrategySwarm, nil}, {"manual", ai.StrategyManual, nil},
	}
	for _, c := range cases {
		agent := &ai.Agent{Name: "e2e_s9_" + c.name, Model: m,
			Instructions: fmt.Sprintf("Parent with %s strategy.", c.name),
			Agents:       []*ai.Agent{childA, childB}, Strategy: c.strategy, Router: c.router}
		plan := planAgent(t, rt, agent)
		if _, ok := plan["workflowDef"]; !ok {
			t.Errorf("[%s] plan has no workflowDef", c.name)
		}
		if _, ok := plan["requiredWorkers"]; !ok {
			t.Errorf("[%s] plan has no requiredWorkers", c.name)
		}
		ad := agentDef(t, plan)
		if ad["strategy"] != c.name {
			t.Errorf("[%s] strategy = %v", c.name, ad["strategy"])
		}
		names := subAgentNames(ad)
		if !contains(names, "child_a") || !contains(names, "child_b") {
			t.Errorf("[%s] sub-agents = %v", c.name, names)
		}
	}
}

// The Python constructor raises for a router strategy without a router; the
// Go SDK reports it from Validate, which Plan and Run call first.
func TestRouterRequiresRouterArgument(t *testing.T) {
	agent := &ai.Agent{Name: "e2e_s9_router_no_arg", Model: "anthropic/claude-sonnet-4-6",
		Instructions: "This should fail.",
		Agents:       []*ai.Agent{{Name: "dummy", Model: "anthropic/claude-sonnet-4-6", Instructions: "X."}},
		Strategy:     ai.StrategyRouter}
	err := agent.Validate()
	if err == nil || !strings.Contains(strings.ToLower(err.Error()), "router") {
		t.Fatalf("Validate() = %v, want an error naming the router", err)
	}
}

// Every child of a sequential parent runs as a sub-workflow, in order, and
// receives the original prompt rather than only the previous output.
func TestSequentialExecution(t *testing.T) {
	rt := newRuntime(t)
	m := model(t)
	parent := &ai.Agent{Name: "e2e_s9_seq_run", Model: m,
		Instructions: "You orchestrate two agents sequentially. First delegate math to math_agent, then text to text_agent.",
		Agents:       []*ai.Agent{s9MathAgent(m), s9TextAgent(m)}, Strategy: ai.StrategySequential}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, parent, "First compute 3+4, then reverse the word hello")
	assertRunCompleted(t, res, "Sequential execution")

	wf := getWorkflow(t, res.ExecutionID)
	all, completed := subWorkflowRefs(wf)
	if len(all) < 2 {
		t.Fatalf("expected at least 2 sub-workflows, found %d: %v", len(all), all)
	}
	if !anyContains(completed, "math", true) || !anyContains(completed, "text", true) {
		t.Errorf("completed sub-workflows %v should include the math and text agents", completed)
	}
	for _, task := range tasksOfType(wf, "SUB_WORKFLOW") {
		if task.SubWorkflowId == "" {
			continue
		}
		child := getWorkflow(t, task.SubWorkflowId)
		prompt := fmt.Sprint(child.Input["prompt"])
		if !strings.Contains(strings.ToLower(prompt), "reverse") && !strings.Contains(prompt, "3+4") {
			t.Errorf("child %s did not receive the original prompt: %q", task.ReferenceTaskName, prompt)
		}
	}
}

func TestParallelExecution(t *testing.T) {
	rt := newRuntime(t)
	m := model(t)
	parent := &ai.Agent{Name: "e2e_s9_par_run", Model: m,
		Instructions: "You orchestrate two agents in parallel. Delegate math to math_agent and text to text_agent simultaneously.",
		Agents:       []*ai.Agent{s9MathAgent(m), s9TextAgent(m)}, Strategy: ai.StrategyParallel}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, parent, "Compute 3+4 AND reverse the word hello")
	assertRunCompleted(t, res, "Parallel execution")

	wf := getWorkflow(t, res.ExecutionID)
	if len(tasksOfType(wf, "FORK", "FORK_JOIN")) < 1 {
		t.Errorf("expected a FORK task in a parallel run")
	}
	_, completed := subWorkflowRefs(wf)
	if !anyContains(completed, "math", true) || !anyContains(completed, "text", true) {
		t.Errorf("completed sub-workflows %v should include the math and text agents", completed)
	}
}

func TestHandoffExecution(t *testing.T) {
	rt := newRuntime(t)
	m := model(t)
	parent := &ai.Agent{Name: "e2e_s9_handoff_run", Model: m,
		Instructions: "You route requests. If the user needs math, delegate to math_agent. If the user needs text manipulation, delegate to text_agent.",
		Agents:       []*ai.Agent{s9MathAgent(m), s9TextAgent(m)}, Strategy: ai.StrategyHandoff}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, parent, "I need to reverse the word hello")
	assertTerminal(t, res, "Handoff execution")
	_, completed := subWorkflowRefs(getWorkflow(t, res.ExecutionID))
	if len(completed) < 1 {
		t.Errorf("no sub-workflow completed; the handoff never reached a child")
	}
}

func TestRouterSelectsCorrectAgent(t *testing.T) {
	rt := newRuntime(t)
	m := model(t)
	router := &ai.Agent{Name: "e2e_s9_router_lead", Model: m,
		Instructions: "You are a router. Route math requests to math_agent and text requests to text_agent. Pick the best agent."}
	parent := &ai.Agent{Name: "e2e_s9_router_run", Model: m, Instructions: "You coordinate agents via a router.",
		Agents: []*ai.Agent{s9MathAgent(m), s9TextAgent(m)}, Strategy: ai.StrategyRouter, Router: router}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, parent, "Compute 7 times 8")
	assertRunCompleted(t, res, "Router selects correct agent")
	all, _ := subWorkflowRefs(getWorkflow(t, res.ExecutionID))
	if !anyContains(all, "math_agent", false) {
		t.Errorf("the router did not pick math_agent; sub-workflows: %v", all)
	}
}

func TestSwarmWithTextMention(t *testing.T) {
	rt := newRuntime(t)
	m := model(t)
	parent := &ai.Agent{Name: "e2e_s9_swarm_run", Model: m,
		Instructions: "You are a swarm coordinator. Handle requests by delegating to the appropriate agent.",
		Agents:       []*ai.Agent{s9MathAgent(m), s9TextAgent(m)}, Strategy: ai.StrategySwarm, MaxTurns: 5,
		Handoffs: []ai.HandoffCondition{
			&ai.OnTextMention{Text: "reverse", Target: "text_agent"},
			&ai.OnTextMention{Text: "compute", Target: "math_agent"},
		}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, parent, "Please reverse the word hello")
	assertTerminal(t, res, "Swarm with text mention")
	all, _ := subWorkflowRefs(getWorkflow(t, res.ExecutionID))
	if !anyContains(all, "text_agent", false) {
		t.Errorf("the swarm never reached text_agent; sub-workflows: %v", all)
	}
}

// Python's `math >> text` builds a sequential parent named after its
// children. Go has no operator for it, so the test builds that parent by
// hand and checks it behaves the same way.
func TestPipeOperatorSequential(t *testing.T) {
	rt := newRuntime(t)
	m := model(t)
	pipeline := &ai.Agent{Name: "math_agent_text_agent", Model: m,
		Agents: []*ai.Agent{s9MathAgent(m), s9TextAgent(m)}, Strategy: ai.StrategySequential}
	if pipeline.Strategy != ai.StrategySequential {
		t.Fatalf("strategy = %q", pipeline.Strategy)
	}
	ad := agentDef(t, planAgent(t, rt, pipeline))
	if ad["strategy"] != "sequential" {
		t.Errorf("compiled strategy = %v", ad["strategy"])
	}
	names := subAgentNames(ad)
	if !contains(names, "math_agent") || !contains(names, "text_agent") {
		t.Errorf("sub-agents = %v", names)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, pipeline, "Compute 2+3 then reverse the word hello")
	assertRunCompleted(t, res, "Pipe operator sequential")
	all, completed := subWorkflowRefs(getWorkflow(t, res.ExecutionID))
	if len(all) < 2 {
		t.Fatalf("expected at least 2 sub-workflows, found %d: %v", len(all), all)
	}
	if !anyContains(completed, "math", true) || !anyContains(completed, "text", true) {
		t.Errorf("completed sub-workflows %v should include the math and text agents", completed)
	}
}
