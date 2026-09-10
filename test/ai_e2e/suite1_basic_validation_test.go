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
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// These are the compile-only tests of the Python SDK's
// e2e/test_suite1_basic_validation.py, for the agents this SDK can express:
// worker tools, credentials, and sub-agents. Each Go test keeps its Python
// test's name (test_plan_reflects_tools is TestPlanReflectsTools) so the two
// suites can be read side by side. Each compiles an agent with
// Plan, asserts on the workflow the way the Python test does, and then goes
// one step further: it compares the whole compiled workflow with the one the
// server returned for the Python agent, captured under testdata/compiled by
// generate_compiled.py. The server compiles deterministically, so any
// difference is a difference in the agentConfig the two SDKs sent.
//
// No model is involved. Compilation never calls an LLM, which is why these
// need no recordings and the model name is the fixed one the Python suite
// uses rather than model(t).
const suite1Model = "anthropic/claude-sonnet-4-6"

type addIn struct {
	A int `json:"a"`
	B int `json:"b"`
}

type multiplyIn struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type greetIn struct {
	Name string `json:"name"`
}

type queryIn struct {
	Query string `json:"query"`
}

type dataIn struct {
	Data string `json:"data"`
}

// The suite's tools, signature for signature with the Python @tool functions.
var (
	addTool = tool.Func("add", "Add two numbers.",
		func(ctx context.Context, in addIn) (int, error) { return in.A + in.B, nil })
	multiplyTool = tool.Func("multiply", "Multiply two numbers.",
		func(ctx context.Context, in multiplyIn) (int, error) { return in.X * in.Y, nil })
	greetTool = tool.Func("greet", "Greet someone.",
		func(ctx context.Context, in greetIn) (string, error) { return "Hello " + in.Name, nil })
	credentialedTool = tool.Func("credentialed_tool", "A tool that needs credentials.",
		func(ctx context.Context, in queryIn) (string, error) { return in.Query, nil },
		tool.WithCredentials("API_KEY_1"))
	multiCredTool = tool.Func("multi_cred_tool", "A tool needing multiple credentials.",
		func(ctx context.Context, in dataIn) (string, error) { return in.Data, nil },
		tool.WithCredentials("SECRET_A", "SECRET_B"))
)

func TestSmokeSimpleAgentPlan(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name: "e2e_smoke", Model: suite1Model, Instructions: "You are a calculator.",
		Tools: []ai.ToolDef{addTool, multiplyTool},
	}
	plan := planAgent(t, rt, agent)
	assertPlanStructure(t, plan, "e2e_smoke")
	ad := agentDef(t, plan)
	assertToolType(t, ad, "add", "worker")
	assertToolType(t, ad, "multiply", "worker")
	assertMatchesPython(t, "e2e_smoke", plan)
}

func TestPlanReflectsTools(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name: "e2e_tools", Model: suite1Model, Instructions: "Use tools.",
		Tools: []ai.ToolDef{addTool, multiplyTool, greetTool},
	}
	plan := planAgent(t, rt, agent)
	ad := agentDef(t, plan)
	for _, name := range []string{"add", "multiply", "greet"} {
		assertToolType(t, ad, name, "worker")
	}
	assertMatchesPython(t, "e2e_tools", plan)
}

func TestPlanReflectsCredentials(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name: "e2e_creds", Model: suite1Model, Instructions: "Use tools.",
		Tools: []ai.ToolDef{credentialedTool, multiCredTool},
	}
	plan := planAgent(t, rt, agent)
	creds := toolCredentials(agentDef(t, plan))
	if got := creds["credentialed_tool"]; !reflect.DeepEqual(got, []string{"API_KEY_1"}) {
		t.Errorf("credentialed_tool credentials = %v, want [API_KEY_1]; all: %v", got, creds)
	}
	got := append([]string(nil), creds["multi_cred_tool"]...)
	sort.Strings(got)
	if !reflect.DeepEqual(got, []string{"SECRET_A", "SECRET_B"}) {
		t.Errorf("multi_cred_tool credentials = %v, want [SECRET_A SECRET_B]", got)
	}
	assertMatchesPython(t, "e2e_creds", plan)
}

func TestPlanSubAgentProducesSubWorkflow(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name: "e2e_parent", Model: suite1Model, Instructions: "Delegate to child.",
		Strategy: ai.StrategyHandoff,
		Agents: []*ai.Agent{
			{Name: "e2e_child", Model: suite1Model, Instructions: "You are a helper."},
		},
	}
	plan := planAgent(t, rt, agent)
	ad := agentDef(t, plan)
	if names := subAgentNames(ad); !contains(names, "e2e_child") {
		t.Errorf("agentDef.agents = %v, want e2e_child", names)
	}
	if ad["strategy"] != "handoff" {
		t.Errorf("agentDef.strategy = %v, want handoff", ad["strategy"])
	}
	if types := taskTypes(allTasks(workflowDef(t, plan))); !types["SUB_WORKFLOW"] {
		t.Errorf("no SUB_WORKFLOW task; types = %v", types)
	}
	assertMatchesPython(t, "e2e_parent", plan)
}

func TestPlanSubAgentReferencesCorrectNames(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name: "e2e_manager", Model: suite1Model,
		Instructions: "Delegate analysis to analyst and writing to writer.",
		Strategy:     ai.StrategyHandoff,
		Agents: []*ai.Agent{
			{Name: "e2e_analyst", Model: suite1Model, Instructions: "You analyze data."},
			{Name: "e2e_writer", Model: suite1Model, Instructions: "You write reports."},
		},
	}
	plan := planAgent(t, rt, agent)
	names := subAgentNames(agentDef(t, plan))
	for _, want := range []string{"e2e_analyst", "e2e_writer"} {
		if !contains(names, want) {
			t.Errorf("agentDef.agents = %v, want %s", names, want)
		}
	}
	sw := subWorkflowNames(allTasks(workflowDef(t, plan)))
	for _, part := range []string{"analyst", "writer"} {
		found := false
		for _, n := range sw {
			if strings.Contains(n, part) {
				found = true
			}
		}
		if !found {
			t.Errorf("no SUB_WORKFLOW references %q; subWorkflowParam names = %v", part, sw)
		}
	}
	assertMatchesPython(t, "e2e_manager", plan)
}

// The suite's TestBaseUrl class: base_url reaches the LLM task's input
// parameters when set, and leaves no key behind when not.
func TestBaseUrlInCompiledWorkflow(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name: "e2e_base_url", Model: suite1Model, Instructions: "Say hello.",
		BaseURL: "https://my-custom-proxy.example.com/v1",
	}
	plan := planAgent(t, rt, agent)
	assertPlanStructure(t, plan, "e2e_base_url")
	llm := llmTasks(plan)
	if len(llm) == 0 {
		t.Fatal("no LLM_CHAT_COMPLETE task in workflow")
	}
	params, _ := llm[0]["inputParameters"].(map[string]any)
	if params["baseUrl"] != "https://my-custom-proxy.example.com/v1" {
		t.Errorf("LLM task baseUrl = %v, want the agent's BaseURL", params["baseUrl"])
	}
	assertMatchesPython(t, "e2e_base_url", plan)
}

func TestNoBaseUrlWhenOmitted(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_no_base_url", Model: suite1Model, Instructions: "Say hello."}
	plan := planAgent(t, rt, agent)
	assertPlanStructure(t, plan, "e2e_no_base_url")
	llm := llmTasks(plan)
	if len(llm) == 0 {
		t.Fatal("no LLM_CHAT_COMPLETE task in workflow")
	}
	params, _ := llm[0]["inputParameters"].(map[string]any)
	if _, has := params["baseUrl"]; has {
		t.Errorf("LLM task carries baseUrl %v although the agent set none", params["baseUrl"])
	}
	assertMatchesPython(t, "e2e_no_base_url", plan)
}

// ── helpers ─────────────────────────────────────────────────────────

func planAgent(t *testing.T, rt *ai.Runtime, agent *ai.Agent) map[string]any {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	plan, err := rt.Plan(ctx, agent)
	if err != nil {
		t.Fatalf("Plan(%s): %v", agent.Name, err)
	}
	return plan
}

func workflowDef(t *testing.T, plan map[string]any) map[string]any {
	t.Helper()
	wf, ok := plan["workflowDef"].(map[string]any)
	if !ok {
		t.Fatalf("plan result has no workflowDef; keys = %v", keys(plan))
	}
	return wf
}

func agentDef(t *testing.T, plan map[string]any) map[string]any {
	t.Helper()
	meta, ok := workflowDef(t, plan)["metadata"].(map[string]any)
	if !ok {
		t.Fatalf("workflowDef has no metadata")
	}
	ad, ok := meta["agentDef"].(map[string]any)
	if !ok {
		t.Fatalf("workflowDef.metadata has no agentDef; keys = %v", keys(meta))
	}
	return ad
}

func assertPlanStructure(t *testing.T, plan map[string]any, name string) {
	t.Helper()
	if _, ok := plan["requiredWorkers"]; !ok {
		t.Errorf("plan result has no requiredWorkers; keys = %v", keys(plan))
	}
	wf := workflowDef(t, plan)
	if wf["name"] != name {
		t.Errorf("workflowDef.name = %v, want %s", wf["name"], name)
	}
	if tasks, _ := wf["tasks"].([]any); len(tasks) == 0 {
		t.Errorf("workflowDef.tasks is empty")
	}
}

func assertToolType(t *testing.T, ad map[string]any, name, want string) {
	t.Helper()
	for _, raw := range asList(ad["tools"]) {
		tl, _ := raw.(map[string]any)
		if tl["name"] == name {
			if tl["toolType"] != want {
				t.Errorf("tool %s has toolType %v, want %s", name, tl["toolType"], want)
			}
			return
		}
	}
	t.Errorf("tool %s not in agentDef.tools", name)
}

func toolCredentials(ad map[string]any) map[string][]string {
	out := map[string][]string{}
	for _, raw := range asList(ad["tools"]) {
		tl, _ := raw.(map[string]any)
		cfg, _ := tl["config"].(map[string]any)
		var creds []string
		for _, c := range asList(cfg["credentials"]) {
			creds = append(creds, fmt.Sprint(c))
		}
		if len(creds) > 0 {
			out[fmt.Sprint(tl["name"])] = creds
		}
	}
	return out
}

func subAgentNames(ad map[string]any) []string {
	var names []string
	for _, raw := range asList(ad["agents"]) {
		a, _ := raw.(map[string]any)
		names = append(names, fmt.Sprint(a["name"]))
	}
	return names
}

// allTasks flattens the workflow's tasks through DO_WHILE, SWITCH and
// FORK_JOIN nesting, as the Python suite's _all_tasks_flat does.
func allTasks(wf map[string]any) []map[string]any {
	var out []map[string]any
	var walk func(raw any)
	walk = func(raw any) {
		task, ok := raw.(map[string]any)
		if !ok {
			return
		}
		out = append(out, task)
		for _, n := range asList(task["loopOver"]) {
			walk(n)
		}
		if cases, ok := task["decisionCases"].(map[string]any); ok {
			for _, c := range cases {
				for _, n := range asList(c) {
					walk(n)
				}
			}
		}
		for _, n := range asList(task["defaultCase"]) {
			walk(n)
		}
		for _, branch := range asList(task["forkTasks"]) {
			for _, n := range asList(branch) {
				walk(n)
			}
		}
	}
	for _, raw := range asList(wf["tasks"]) {
		walk(raw)
	}
	return out
}

func taskTypes(tasks []map[string]any) map[string]bool {
	out := map[string]bool{}
	for _, task := range tasks {
		out[fmt.Sprint(task["type"])] = true
	}
	return out
}

func subWorkflowNames(tasks []map[string]any) []string {
	var names []string
	for _, task := range tasks {
		if task["type"] != "SUB_WORKFLOW" {
			continue
		}
		params, _ := task["subWorkflowParam"].(map[string]any)
		if params == nil {
			params, _ = task["subWorkflowParams"].(map[string]any)
		}
		if n, ok := params["name"].(string); ok && n != "" {
			names = append(names, n)
		}
	}
	return names
}

// assertMatchesPython compares the plan with the one the Python SDK got for
// the same agent, from testdata/compiled. Both sides are JSON round-tripped
// first so the comparison is about content, not Go types.
func assertMatchesPython(t *testing.T, name string, plan map[string]any) {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("testdata", "compiled", name+".json"))
	if err != nil {
		t.Fatalf("read Python fixture: %v", err)
	}
	var want map[string]any
	if err := json.Unmarshal(raw, &want); err != nil {
		t.Fatalf("parse Python fixture: %v", err)
	}
	var got map[string]any
	buf, _ := json.Marshal(plan)
	if err := json.Unmarshal(buf, &got); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("compiled workflow differs from the Python SDK's for %s\n--- go ---\n%s\n--- python ---\n%s",
			name, indent(got), indent(want))
	}
	assertSameLLMCalls(t, name, got, want)
}

// assertSameLLMCalls is the point of the comparison, stated on its own: the
// LLM_CHAT_COMPLETE tasks are the requests the server will send to the model
// when the agent runs, with the model, the instructions, the message template
// and the tool schemas. Identical tasks mean identical LLM calls, which is
// all a compile-only suite can show about them.
func assertSameLLMCalls(t *testing.T, name string, got, want map[string]any) {
	t.Helper()
	gotLLM := llmTasks(got)
	wantLLM := llmTasks(want)
	if len(gotLLM) == 0 || len(gotLLM) != len(wantLLM) {
		t.Fatalf("%s: Go compiled %d LLM_CHAT_COMPLETE tasks, Python %d", name, len(gotLLM), len(wantLLM))
	}
	for i := range gotLLM {
		if !reflect.DeepEqual(gotLLM[i]["inputParameters"], wantLLM[i]["inputParameters"]) {
			t.Errorf("%s: LLM task %v would call the model differently from Python\n--- go ---\n%s\n--- python ---\n%s",
				name, gotLLM[i]["taskReferenceName"], indent(gotLLM[i]["inputParameters"]), indent(wantLLM[i]["inputParameters"]))
		}
	}
	params, _ := gotLLM[0]["inputParameters"].(map[string]any)
	t.Logf("%s: %d LLM_CHAT_COMPLETE task(s) identical to Python's; inputs carry %v",
		name, len(gotLLM), keys(params))
}

// llmTasks returns the LLM_CHAT_COMPLETE tasks of a plan result, in workflow order.
func llmTasks(plan map[string]any) []map[string]any {
	wf, _ := plan["workflowDef"].(map[string]any)
	var out []map[string]any
	for _, task := range allTasks(wf) {
		if task["type"] == "LLM_CHAT_COMPLETE" {
			out = append(out, task)
		}
	}
	return out
}

func indent(v any) string {
	raw, _ := json.MarshalIndent(v, "", "  ")
	return string(raw)
}

func asList(v any) []any {
	l, _ := v.([]any)
	return l
}

func contains(list []string, s string) bool {
	for _, x := range list {
		if x == s {
			return true
		}
	}
	return false
}

func keys(m map[string]any) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
