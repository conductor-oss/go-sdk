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
// suites can be read side by side. Each compiles an agent with Plan and
// asserts on the workflow the way the Python test does.
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
}

// suite1MCPURL is the MCP test server the Python suite's kitchen sink points
// at. The URL is part of the compiled workflow, so it is the same literal as
// rather than the mcp_test.go default.
const suite1MCPURL = "http://localhost:3001"

func passingGuardrail(name string, position ai.Position) *ai.CustomGuardrail {
	g := ai.NewCustomGuardrail(name, func(ctx context.Context, in ai.GuardrailInput) (ai.GuardrailResult, error) {
		return ai.GuardrailResult{Passed: true}, nil
	})
	g.Position = position
	g.OnFail = ai.OnFailRetry
	return g
}

func regexGuardrail(name, pattern, message string) *ai.RegexGuardrail {
	g := &ai.RegexGuardrail{Patterns: []string{pattern}, Message: message}
	g.Name = name
	g.OnFail = ai.OnFailRetry
	return g
}

func TestPlanReflectsGuardrails(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name: "e2e_guardrails", Model: suite1Model, Instructions: "Answer questions.",
		Tools: []ai.ToolDef{greetTool},
		Guardrails: []ai.Guardrail{
			passingGuardrail("check_input", ai.PositionInput),
			passingGuardrail("no_pii", ai.PositionOutput),
			regexGuardrail("no_ssn", `\b\d{3}-\d{2}-\d{4}\b`, "No SSNs allowed."),
		},
	}
	plan := planAgent(t, rt, agent)
	ad := agentDef(t, plan)
	guardrails := asList(ad["guardrails"])
	if len(guardrails) != 3 {
		t.Fatalf("agentDef.guardrails has %d entries, want 3", len(guardrails))
	}
	byName := map[string]map[string]any{}
	for _, raw := range guardrails {
		g, _ := raw.(map[string]any)
		byName[fmt.Sprint(g["name"])] = g
		if g["onFail"] != "retry" {
			t.Errorf("guardrail %v onFail = %v, want retry", g["name"], g["onFail"])
		}
	}
	for _, name := range []string{"check_input", "no_pii", "no_ssn"} {
		if _, ok := byName[name]; !ok {
			t.Errorf("guardrail %s missing; have %v", name, keys(toAny(byName)))
		}
	}
	if byName["check_input"]["position"] != "input" || byName["no_pii"]["position"] != "output" {
		t.Errorf("positions: check_input=%v no_pii=%v", byName["check_input"]["position"], byName["no_pii"]["position"])
	}
	if byName["no_ssn"]["guardrailType"] != "regex" {
		t.Errorf("no_ssn guardrailType = %v, want regex", byName["no_ssn"]["guardrailType"])
	}
	if patterns := asList(byName["no_ssn"]["patterns"]); len(patterns) != 1 || patterns[0] != `\b\d{3}-\d{2}-\d{4}\b` {
		t.Errorf("no_ssn patterns = %v; the pattern must survive verbatim", patterns)
	}
}

// kitchenSink is the suite's _make_kitchen_sink_agent: every tool type,
// three guardrails, a credential, and all eight sub-agent strategies.
func kitchenSink() *ai.Agent {
	leaf := func(name, instructions string) *ai.Agent {
		return &ai.Agent{Name: name, Model: suite1Model, Instructions: instructions}
	}
	team := func(name string, strategy ai.Strategy, a, b *ai.Agent) *ai.Agent {
		return &ai.Agent{Name: name, Model: suite1Model, Strategy: strategy, Agents: []*ai.Agent{a, b}}
	}
	handoff := team("ks_handoff", ai.StrategyHandoff, leaf("ks_h1", "H1."), leaf("ks_h2", "H2."))
	handoff.Instructions = "Route tasks."
	router := team("ks_router", ai.StrategyRouter, leaf("ks_r1", "R1."), leaf("ks_r2", "R2."))
	router.Router = leaf("ks_router_lead", "Route to correct agent.")
	swarm := team("ks_swarm", ai.StrategySwarm, leaf("ks_sw1", "SW1."), leaf("ks_sw2", "SW2."))
	swarm.Handoffs = []ai.HandoffCondition{
		&ai.OnTextMention{Text: "GOTO_SW2", Target: "ks_sw2"},
		&ai.OnTextMention{Text: "GOTO_SW1", Target: "ks_sw1"},
	}

	return &ai.Agent{
		Name: "e2e_kitchen_sink", Model: suite1Model, Instructions: "You are the kitchen sink agent.",
		Tools: []ai.ToolDef{
			tool.Func("local_tool", "A local worker tool.",
				func(ctx context.Context, in xIn) (string, error) { return in.X, nil }),
			tool.Func("cred_local_tool", "Worker tool with credentials.",
				func(ctx context.Context, in xIn) (string, error) { return in.X, nil },
				tool.WithCredentials("KS_SECRET")),
			tool.HTTP("ks_http", "HTTP endpoint", suite1MCPURL+"/echo", tool.WithMethod("POST")),
			tool.MCP("ks_mcp", "MCP tools", suite1MCPURL),
			tool.Image("ks_image", "Generate image", "openai", "dall-e-3"),
			tool.Audio("ks_audio", "Generate audio", "openai", "tts-1"),
			tool.Video("ks_video", "Generate video", "openai", "sora"),
			tool.PDF("ks_pdf", "Generate PDF"),
		},
		Guardrails: []ai.Guardrail{
			passingGuardrail("check_input", ai.PositionInput),
			passingGuardrail("no_pii", ai.PositionOutput),
			regexGuardrail("no_password", "password", "No passwords in output."),
		},
		Strategy: ai.StrategyHandoff,
		Agents: []*ai.Agent{
			handoff,
			team("ks_sequential", ai.StrategySequential, leaf("ks_seq1", "Seq1."), leaf("ks_seq2", "Seq2.")),
			team("ks_parallel", ai.StrategyParallel, leaf("ks_p1", "P1."), leaf("ks_p2", "P2.")),
			router,
			team("ks_round_robin", ai.StrategyRoundRobin, leaf("ks_rr1", "RR1."), leaf("ks_rr2", "RR2.")),
			team("ks_random", ai.StrategyRandom, leaf("ks_rand1", "Rand1."), leaf("ks_rand2", "Rand2.")),
			swarm,
			team("ks_manual", ai.StrategyManual, leaf("ks_m1", "M1."), leaf("ks_m2", "M2.")),
		},
	}
}

type xIn struct {
	X string `json:"x"`
}

func TestKitchenSinkCompiles(t *testing.T) {
	rt := newRuntime(t)
	plan := planAgent(t, rt, kitchenSink())
	wf := workflowDef(t, plan)
	assertPlanStructure(t, plan, "e2e_kitchen_sink")
	ad := agentDef(t, plan)

	for name, want := range map[string]string{
		"local_tool": "worker", "cred_local_tool": "worker", "ks_http": "http", "ks_mcp": "mcp",
		"ks_image": "generate_image", "ks_audio": "generate_audio", "ks_video": "generate_video", "ks_pdf": "generate_pdf",
	} {
		assertToolType(t, ad, name, want)
	}
	if creds := toolCredentials(ad); !reflect.DeepEqual(creds["cred_local_tool"], []string{"KS_SECRET"}) {
		t.Errorf("cred_local_tool credentials = %v, want [KS_SECRET]", creds["cred_local_tool"])
	}

	guardrails := asList(ad["guardrails"])
	if len(guardrails) != 3 {
		t.Errorf("agentDef.guardrails has %d entries, want 3", len(guardrails))
	}
	for _, raw := range guardrails {
		g, _ := raw.(map[string]any)
		if g["name"] == "no_password" {
			if g["guardrailType"] != "regex" || !contains(toStrings(asList(g["patterns"])), "password") {
				t.Errorf("no_password = %v", g)
			}
		}
	}

	wantStrategies := map[string]string{
		"ks_handoff": "handoff", "ks_sequential": "sequential", "ks_parallel": "parallel", "ks_router": "router",
		"ks_round_robin": "round_robin", "ks_random": "random", "ks_swarm": "swarm", "ks_manual": "manual",
	}
	subs := map[string]map[string]any{}
	for _, raw := range asList(ad["agents"]) {
		a, _ := raw.(map[string]any)
		subs[fmt.Sprint(a["name"])] = a
	}
	for name, want := range wantStrategies {
		sub, ok := subs[name]
		if !ok {
			t.Errorf("sub-agent %s missing; have %v", name, keys(toAny(subs)))
			continue
		}
		if sub["strategy"] != want {
			t.Errorf("sub-agent %s strategy = %v, want %s", name, sub["strategy"], want)
		}
	}
	if ad["strategy"] != "handoff" {
		t.Errorf("parent strategy = %v, want handoff", ad["strategy"])
	}
	if types := taskTypes(allTasks(wf)); !types["SUB_WORKFLOW"] {
		t.Errorf("no SUB_WORKFLOW task; types = %v", types)
	}
}

func toAny[V any](m map[string]V) map[string]any {
	out := make(map[string]any, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}

func toStrings(list []any) []string {
	out := make([]string, 0, len(list))
	for _, v := range list {
		out = append(out, fmt.Sprint(v))
	}
	return out
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

func sortedTokens(script string) string {
	tokens := strings.FieldsFunc(script, func(r rune) bool {
		return strings.ContainsRune("{}[],: \n\t", r)
	})
	sort.Strings(tokens)
	return strings.Join(tokens, " ")
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
