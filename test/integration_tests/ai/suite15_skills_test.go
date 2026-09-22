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
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// The server-side half of the Python SDK's e2e/test_suite15_skills.py: what
// the server makes of a skill document, and how a skill nested under another
// agent behaves at run time.
//
// The suite's other tests load and serialize a skill without a server, and
// those are unit tests of this SDK already: see sdk/ai/skill_test.go, which
// covers discovery, the wire document, script languages, params and the
// script and file-reading workers. The two standalone run tests are
// TestSkillScriptRunsAsWorker and TestSkillAsAgentTool in skill_test.go here.

const suite15Skill = "test_skill"

// writeSuite15Skill lays out the Python fixture's skill: two sub-agents, one
// script and one resource file, so the compiled workflow can be checked for
// all four names the server derives from them.
func writeSuite15Skill(t *testing.T) (dir, markerFile string) {
	t.Helper()
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("bash not on PATH; the skill script needs it")
	}
	root := t.TempDir()
	dir = filepath.Join(root, suite15Skill)
	markerFile = filepath.Join(root, "echo_args.calls")

	script := "#!/bin/bash\n" +
		"args=\"$*\"\n" +
		"[ -z \"$args\" ] && args=no-args\n" +
		"printf '%s\\n' \"$args\" >> '" + markerFile + "'\n" +
		"printf 'ECHO_ARGS_RESULT:%s\\n' \"$args\"\n"

	files := map[string]string{
		"SKILL.md": "---\n" +
			"name: " + suite15Skill + "\n" +
			"description: End-to-end test skill. Use when asked to echo something.\n" +
			"params:\n  mode:\n    default: fast\n" +
			"---\n" +
			"## Overview\n" +
			"A test skill with two sub-agents and a script tool.\n\n" +
			"## Workflow\n" +
			"1. If no prior tool result is available, call the " + suite15Skill + "__echo_args tool exactly once.\n" +
			"2. Pass the original user's input as the argument.\n" +
			"3. After a tool result containing ECHO_ARGS_RESULT: is available, return that exact line as the final answer.\n" +
			"4. If asked to continue, do not call any tool. Return the most recent ECHO_ARGS_RESULT: line exactly.\n",
		"alpha-agent.md":       "# Alpha Agent\nYou analyze the input.\n",
		"beta-agent.md":        "# Beta Agent\nYou summarize the analysis.\n",
		"references/guide.md":  "# REFERENCE_GUIDE\nUse this deterministic guide.\n",
		"scripts/echo_args.sh": script,
	}
	for rel, content := range files {
		p := filepath.Join(dir, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(content), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	return dir, markerFile
}

func loadSuite15Skill(t *testing.T, opts ...ai.SkillOption) *ai.Agent {
	t.Helper()
	dir, _ := writeSuite15Skill(t)
	skill, err := ai.LoadSkill(dir, append([]ai.SkillOption{ai.WithSkillModel(model(t))}, opts...)...)
	if err != nil {
		t.Fatalf("LoadSkill: %v", err)
	}
	return skill
}

// The server turns a skill document into an agent-loop workflow named after
// the skill.
func TestSkillPlanCompilation(t *testing.T) {
	rt := newRuntime(t)
	def := workflowDef(t, planAgent(t, rt, loadSuite15Skill(t)))

	if def["name"] != suite15Skill {
		t.Errorf("compiled workflow name = %v, want %s", def["name"], suite15Skill)
	}
	types := taskTypes(allTasks(def))
	if !types["LLM_CHAT_COMPLETE"] {
		t.Errorf("the compiled skill has no model call: %v", types)
	}
	if !types["DO_WHILE"] && !types["FORK_JOIN_DYNAMIC"] {
		t.Errorf("the compiled skill has no agent loop: %v", types)
	}
}

// Compilation derives a name per part of the document: a sub-workflow for
// each sub-agent, and a worker task for the script and for reading resources.
// Those are the names this process must serve, so they are the contract
// between the document and the workers.
func TestSkillPlanExposesMultiAgentScriptAndResourceTools(t *testing.T) {
	rt := newRuntime(t)
	def := workflowDef(t, planAgent(t, rt, loadSuite15Skill(t)))
	raw, err := json.Marshal(def)
	if err != nil {
		t.Fatal(err)
	}
	compiled := string(raw)

	for _, want := range []string{
		suite15Skill + "__alpha",
		suite15Skill + "__beta",
		suite15Skill + "__echo_args",
		suite15Skill + "__read_skill_file",
		"references/guide.md",
		"SUB_WORKFLOW",
		"SIMPLE",
	} {
		if !strings.Contains(compiled, want) {
			t.Errorf("the compiled workflow does not mention %q", want)
		}
	}

	// The script is a worker task, not something the server runs itself. The
	// spec for it can sit anywhere in the document, since the tool tasks
	// themselves are created while the run is going, so look everywhere.
	var specs, simple int
	walkObjects(def, func(m map[string]any) {
		if m["name"] != suite15Skill+"__echo_args" {
			return
		}
		specs++
		if m["type"] == "SIMPLE" {
			simple++
		}
	})
	if specs == 0 {
		t.Errorf("the compiled workflow has no spec named %s__echo_args", suite15Skill)
	} else if simple == 0 {
		t.Errorf("no spec named %s__echo_args is a SIMPLE task, so the script is not a worker tool", suite15Skill)
	}
}

// walkObjects visits every JSON object in a compiled document, the Python
// suite's _all_dicts.
func walkObjects(v any, visit func(map[string]any)) {
	switch value := v.(type) {
	case map[string]any:
		visit(value)
		for _, inner := range value {
			walkObjects(inner, visit)
		}
	case []any:
		for _, inner := range value {
			walkObjects(inner, visit)
		}
	}
}

// Parameters given at load time reach the document the server compiles.
func TestSkillParamsInCompiledWorkflow(t *testing.T) {
	rt := newRuntime(t)
	skill := loadSuite15Skill(t, ai.WithSkillParams(map[string]any{"mode": "turbo", "rounds": 1}))
	raw, err := json.Marshal(workflowDef(t, planAgent(t, rt, skill)))
	if err != nil {
		t.Fatal(err)
	}
	compiled := string(raw)
	// Tighter than the Python assertion, which accepts the word "mode" alone.
	for _, want := range []string{"[Skill Parameters]", "mode: turbo", "rounds: 1"} {
		if !strings.Contains(compiled, want) {
			t.Errorf("the compiled workflow does not carry %q", want)
		}
	}
}

// A skill nested under a stateful parent: the skill's own workers have to
// poll the parent run's domain, not the shared queue, or its script task is
// scheduled where nothing is listening and the run stalls.
func TestAgentToolSkillWorkersWithDomain(t *testing.T) {
	rt := newRuntime(t)
	if model(t) == mockModel {
		// As with TestSkillAsAgentTool: the parent's next request carries the
		// sub-workflow's id, which is new every run, so no recording matches.
		t.Skip("agent tool results include a per-run subWorkflowId; the recorder cannot replay this")
	}
	dir, marker := writeSuite15Skill(t)
	skill, err := ai.LoadSkill(dir, ai.WithSkillModel(model(t)))
	if err != nil {
		t.Fatalf("LoadSkill: %v", err)
	}

	parent := &ai.Agent{
		Name: "e2e_skill_at_domain", Model: model(t), Stateful: true, MaxTurns: 3,
		Instructions: "You have one tool: " + suite15Skill + ". " +
			"Call it once with the user's request, then return the result.",
		Tools: ai.Tools(tool.Agent(skill, "", "Run test skill with echo_args")),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res, err := rt.Run(ctx, parent, "Echo 'domain_proof'")
	if err != nil {
		t.Fatalf("Run: %v (a stalled run here means the nested skill's workers polled the wrong domain)", err)
	}
	if res.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", res.Status, ai.StatusCompleted, res.Error)
	}

	// The parent is stateful, so its run has a domain of its own.
	if len(domainsOf(t, res.ExecutionID)) == 0 {
		t.Error("the stateful parent recorded no task domain")
	}
	// Nothing anywhere, the skill's sub-workflow included, went unpolled.
	assertNothingScheduled(t, allTasksDeep(t, res.ExecutionID))
	// And the skill's script really ran in this process.
	if calls := markerCalls(t, marker); len(calls) == 0 {
		t.Fatal("echo_args never ran; the nested skill's workers did not reach the parent run's domain")
	}
}
