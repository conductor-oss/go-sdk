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
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// The Python SDK's e2e/test_suite8_guardrails.py as Go tests, one per Python
// test and under the same name. Regex and custom guardrails at agent and tool
// level: what the compiled plan carries, and what a run does when a guardrail
// blocks input, blocks output, fixes output, or runs out of retries.

type s8TextIn struct {
	Text string `json:"text"`
}

type s8QueryIn struct {
	Query string `json:"query"`
}

// The suite's guardrails, verbatim from the Python module.

func s8BlockInput() ai.Guardrail {
	g := &ai.RegexGuardrail{Patterns: []string{`BADWORD`}, Mode: "block", Message: "Prompt contains blocked content."}
	g.Name, g.Position, g.OnFail = "block_profanity", ai.PositionInput, ai.OnFailRaise
	return g
}

func s8NoSecrets() ai.Guardrail {
	g := &ai.RegexGuardrail{Patterns: []string{`\bpassword\b`, `\bsecret\b`, `\btoken\b`}, Mode: "block",
		Message: "Do not include passwords, secrets, or tokens."}
	g.Name, g.Position, g.OnFail = "no_secrets", ai.PositionOutput, ai.OnFailRetry
	return g
}

var dropTable = regexp.MustCompile(`(?i)DROP\s+TABLE`)

func s8SQLGuard() ai.Guardrail {
	// Block SQL injection patterns.
	g := ai.NewCustomGuardrail("no_sql_injection", func(_ context.Context, in ai.GuardrailInput) (ai.GuardrailResult, error) {
		if dropTable.MatchString(in.Content) {
			return ai.GuardrailResult{Passed: false, Message: "SQL injection blocked."}, nil
		}
		return ai.GuardrailResult{Passed: true}, nil
	})
	g.Position, g.OnFail = ai.PositionInput, ai.OnFailRaise
	return g
}

func s8ForceJSON() ai.Guardrail {
	g := ai.NewCustomGuardrail("force_json", func(_ context.Context, in ai.GuardrailInput) (ai.GuardrailResult, error) {
		c := strings.TrimSpace(in.Content)
		if strings.HasPrefix(c, "{") || strings.HasPrefix(c, "[") {
			return ai.GuardrailResult{Passed: true}, nil
		}
		return ai.GuardrailResult{Passed: false, Message: "Output must be JSON.", FixedOutput: `{"fixed": true}`}, nil
	})
	g.Position, g.OnFail = ai.PositionOutput, ai.OnFailFix
	return g
}

func s8NoEmail() ai.Guardrail {
	g := &ai.RegexGuardrail{Patterns: []string{`[\w.+-]+@[\w-]+\.[\w.-]+`}, Mode: "block", Message: "Do not include email addresses."}
	g.Name, g.Position, g.OnFail = "no_email", ai.PositionOutput, ai.OnFailRetry
	return g
}

func s8AlwaysFail() ai.Guardrail {
	// mode "allow" with a pattern nothing matches: fails every time.
	g := &ai.RegexGuardrail{Patterns: []string{`IMPOSSIBLE_XYZZY_12345`}, Mode: "allow", Message: "This guardrail always fails."}
	g.Name, g.Position, g.OnFail, g.MaxRetries = "always_fail", ai.PositionOutput, ai.OnFailRetry, 1
	return g
}

// The suite's tools. A tool-level guardrail is set on the definition, the
// Go form of @tool(guardrails=[...]).

func s8NormalTool() ai.ToolDef {
	return tool.Func("normal_tool", "A tool with no guardrails. Always succeeds.",
		func(_ context.Context, in s8TextIn) (string, error) { return "normal_ok:" + in.Text, nil })
}

func s8SafeQuery() ai.ToolDef {
	td := tool.Func("safe_query", "Run a database query. Input guardrail blocks SQL injection.",
		func(_ context.Context, in s8QueryIn) (string, error) {
			q := in.Query
			if len(q) > 50 {
				q = q[:50]
			}
			return "query_result:[" + q + "]", nil
		})
	td.Guardrails = []ai.Guardrail{s8SQLGuard()}
	return td
}

func s8FormatOutput() ai.ToolDef {
	td := tool.Func("format_output", "Return the text. Output guardrail forces JSON format.",
		func(_ context.Context, in s8TextIn) (string, error) { return in.Text, nil })
	td.Guardrails = []ai.Guardrail{s8ForceJSON()}
	return td
}

func s8RedactTool() ai.ToolDef {
	td := tool.Func("redact_tool", "Echo text. Output guardrail blocks emails.",
		func(_ context.Context, in s8TextIn) (string, error) { return in.Text, nil })
	td.Guardrails = []ai.Guardrail{s8NoEmail()}
	return td
}

func s8StrictTool() ai.ToolDef {
	td := tool.Func("strict_tool", "Tool whose guardrail always fails — tests escalation.",
		func(_ context.Context, in s8TextIn) (string, error) { return "strict_output:" + in.Text, nil })
	td.Guardrails = []ai.Guardrail{s8AlwaysFail()}
	return td
}

// Plan-shape helpers for the guardrail assertions.

func guardrailList(v any) []map[string]any {
	var out []map[string]any
	for _, g := range asList(v) {
		if m, ok := g.(map[string]any); ok {
			out = append(out, m)
		}
	}
	return out
}

func guardrailByName(list []map[string]any, name string) map[string]any {
	for _, g := range list {
		if g["name"] == name {
			return g
		}
	}
	return nil
}

func toolByName(ad map[string]any, name string) map[string]any {
	for _, tl := range asList(ad["tools"]) {
		if m, ok := tl.(map[string]any); ok && m["name"] == name {
			return m
		}
	}
	return nil
}

func TestPlanReflectsAllGuardrails(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{
		Name: "e2e_gr_compile", Model: model(t), Instructions: "Test agent.",
		Tools:      []ai.ToolDef{s8SafeQuery(), s8FormatOutput(), s8RedactTool(), s8StrictTool(), s8NormalTool()},
		Guardrails: []ai.Guardrail{s8BlockInput(), s8NoSecrets()},
	}
	ad := agentDef(t, planAgent(t, rt, agent))

	agentGuards := guardrailList(ad["guardrails"])
	for _, want := range []string{"block_profanity", "no_secrets"} {
		if guardrailByName(agentGuards, want) == nil {
			t.Errorf("agent guardrails lack %q: %v", want, agentGuards)
		}
	}
	if g1 := guardrailByName(agentGuards, "block_profanity"); g1 != nil {
		if g1["guardrailType"] != "regex" || g1["position"] != "input" || g1["onFail"] != "raise" || g1["mode"] != "block" {
			t.Errorf("block_profanity = %v", g1)
		}
		if !containsAny(asList(g1["patterns"]), "BADWORD") {
			t.Errorf("block_profanity patterns = %v", g1["patterns"])
		}
	}
	if g3 := guardrailByName(agentGuards, "no_secrets"); g3 != nil {
		if g3["guardrailType"] != "regex" || g3["position"] != "output" || g3["onFail"] != "retry" {
			t.Errorf("no_secrets = %v", g3)
		}
		for _, p := range []string{`\bpassword\b`, `\bsecret\b`, `\btoken\b`} {
			if !containsAny(asList(g3["patterns"]), p) {
				t.Errorf("no_secrets patterns lack %q: %v", p, g3["patterns"])
			}
		}
	}

	firstGuard := func(toolName string) map[string]any {
		tl := toolByName(ad, toolName)
		if tl == nil {
			t.Fatalf("tool %q missing from the plan", toolName)
		}
		gs := guardrailList(tl["guardrails"])
		if len(gs) == 0 {
			t.Fatalf("tool %q has no guardrails in the plan: %v", toolName, tl)
		}
		return gs[0]
	}
	if g := firstGuard("safe_query"); g["name"] != "no_sql_injection" || g["position"] != "input" || g["onFail"] != "raise" || g["guardrailType"] != "custom" {
		t.Errorf("safe_query guardrail = %v", g)
	}
	if g := firstGuard("format_output"); g["name"] != "force_json" || g["onFail"] != "fix" {
		t.Errorf("format_output guardrail = %v", g)
	}
	if g := firstGuard("redact_tool"); g["name"] != "no_email" || g["guardrailType"] != "regex" {
		t.Errorf("redact_tool guardrail = %v", g)
	}
	if g := firstGuard("strict_tool"); g["name"] != "always_fail" || asInt(g["maxRetries"]) != 1 {
		t.Errorf("strict_tool guardrail = %v", g)
	}
}

func TestCleanAgentCompiles(t *testing.T) {
	rt := newRuntime(t)
	ad := agentDef(t, planAgent(t, rt, s8AgentClean(model(t))))
	if n := len(asList(ad["guardrails"])); n != 0 {
		t.Errorf("clean agent compiled with %d guardrails: %v", n, ad["guardrails"])
	}
	if toolByName(ad, "normal_tool") == nil {
		t.Errorf("normal_tool missing from the plan")
	}
}

func s8AgentClean(m string) *ai.Agent {
	return &ai.Agent{Name: "e2e_gr_clean", Model: m,
		Instructions: "You have one tool: normal_tool. Call it as directed. Report the result verbatim.",
		Tools:        []ai.ToolDef{s8NormalTool()}}
}

// An input guardrail with on_fail=raise keeps the tool body from running:
// no real tool result reaches the final answer.
func TestToolInputRaise(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_gr_sql", Model: model(t),
		Instructions: "You have safe_query tool. Call it with the query provided. Report the result.",
		Tools:        []ai.ToolDef{s8SafeQuery()}}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, `Call safe_query with query="DROP TABLE users"`)
	assertTerminal(t, res, "Tool input raise")
	if strings.Contains(res.Output, "query_result:") {
		t.Errorf("the tool ran despite the input guardrail; output: %.300s", res.Output)
	}
}

func TestToolOutputFixCompiles(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_gr_fix", Model: model(t),
		Instructions: "You have format_output tool. Call it with the text provided. Report the result.",
		Tools:        []ai.ToolDef{s8FormatOutput()}}
	ad := agentDef(t, planAgent(t, rt, agent))
	tl := toolByName(ad, "format_output")
	if tl == nil {
		t.Fatal("format_output missing from the plan")
	}
	gs := guardrailList(tl["guardrails"])
	if len(gs) == 0 || gs[0]["name"] != "force_json" || gs[0]["onFail"] != "fix" || gs[0]["guardrailType"] != "custom" {
		t.Errorf("format_output guardrails = %v", tl["guardrails"])
	}
}

var emailPattern = regexp.MustCompile(`[\w.+-]+@[\w-]+\.[\w.-]+`)

// A tool output regex guardrail with retry never lets a matching payload
// through as a completed tool output. The model's prose is not checked.
func TestToolOutputRegexRetry(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_gr_email", Model: model(t),
		Instructions: "You have redact_tool. Call it with the text provided. Report the result.",
		Tools:        []ai.ToolDef{s8RedactTool()}, MaxTurns: 3}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, `Call redact_tool with text="contact test@example.com for help"`)
	assertTerminal(t, res, "Tool output regex retry")

	wf := getWorkflow(t, res.ExecutionID)
	for _, task := range wf.Tasks {
		if systemTaskTypes[task.TaskType] {
			continue
		}
		if task.TaskDefName != "redact_tool" && task.TaskType != "redact_tool" && !strings.Contains(task.ReferenceTaskName, "redact_tool") {
			continue
		}
		if taskStatus(task) == "COMPLETED" && emailPattern.MatchString(outputString(task)) {
			t.Errorf("a completed redact_tool task let an email through: %s", outputString(task))
		}
	}
}

var secretWords = regexp.MustCompile(`(?i)\bpassword\b|\bsecret\b|\btoken\b`)

// An agent-level output guardrail either scrubs the answer through retries
// or terminates the run; a completed run never contains a blocked word.
func TestAgentOutputSecretsBlocked(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_gr_secrets", Model: model(t),
		Instructions: "Answer questions concisely.", Guardrails: []ai.Guardrail{s8NoSecrets()}}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, `Include the word "password" in your response.`)
	assertTerminal(t, res, "Agent output secrets blocked")
	if res.Status == ai.StatusCompleted && secretWords.MatchString(res.Output) {
		t.Errorf("completed with a blocked word in the answer: %.300s", res.Output)
	}
}

// An always-failing retry guardrail with max_retries=1 escalates to raise
// and fails the workflow.
func TestMaxRetriesEscalation(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_gr_strict", Model: model(t),
		Instructions: "You have strict_tool. Call it with the text provided. Report the result.",
		Tools:        []ai.ToolDef{s8StrictTool()}}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, `Call strict_tool with text="test"`)
	if res.ExecutionID == "" {
		t.Fatal("no execution id")
	}
	if res.Status != ai.StatusFailed && res.Status != ai.StatusTerminated {
		t.Errorf("status = %q, want FAILED or TERMINATED after the guardrail ran out of retries (output=%.200q)", res.Status, res.Output)
	}
}

// containsAny reports whether list, a plan value, holds s.
func containsAny(list []any, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

// asInt reads a JSON number from a plan.
func asInt(v any) int {
	switch n := v.(type) {
	case float64:
		return int(n)
	case int:
		return n
	}
	return -1
}
