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
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// The skill tests mirror the execution half of the Python SDK's
// test_suite15_skills.py. The golden fixtures prove the skill document
// serializes identically; these ask whether a run built from it works: the
// server normalizes the document, the model calls the script tool, and
// Conductor dispatches that call back to a worker in this process.
//
// Proof that the worker ran here, rather than the model answering from
// memory, is a marker file the script appends to. The file lives in the test's
// temp dir and the script is generated with its path, so it cannot be written
// by anything but this process's worker.
const skillName = "go_e2e_skill"

// writeE2ESkill lays out a skill with one bash script, echo_args, which prints
// a deterministic marker and records its arguments. The instructions name the
// tool exactly, the way the Python fixture does, so the run is short and the
// assertion is on the tool path rather than on the model's judgement.
func writeE2ESkill(t *testing.T) (dir, markerFile string) {
	t.Helper()
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("bash not on PATH; the skill script needs it")
	}
	root := t.TempDir()
	dir = filepath.Join(root, skillName)
	markerFile = filepath.Join(root, "echo_args.calls")

	skillMd := "---\n" +
		"name: " + skillName + "\n" +
		"description: End-to-end test skill. Use when asked to echo something.\n" +
		"params:\n  mode:\n    default: fast\n" +
		"---\n" +
		"## Overview\n" +
		"A test skill with a script tool.\n\n" +
		"## Workflow\n" +
		"1. If no prior tool result is available, call the " + skillName + "__echo_args tool exactly once.\n" +
		"2. Pass the original user's input as the argument.\n" +
		"3. After a tool result containing ECHO_ARGS_RESULT: is available, return that exact line as the final answer.\n" +
		"4. If asked to continue, do not call any tool. Return the most recent ECHO_ARGS_RESULT: line exactly.\n"

	script := "#!/bin/bash\n" +
		"args=\"$*\"\n" +
		"[ -z \"$args\" ] && args=no-args\n" +
		"printf '%s\\n' \"$args\" >> '" + markerFile + "'\n" +
		"printf 'ECHO_ARGS_RESULT:%s\\n' \"$args\"\n"

	files := map[string]string{
		"SKILL.md":             skillMd,
		"scripts/echo_args.sh": script,
		"references/guide.md":  "# REFERENCE_GUIDE\nUse this deterministic guide.\n",
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

// markerCalls reads what the script recorded, one call per line.
func markerCalls(t *testing.T, markerFile string) []string {
	t.Helper()
	data, err := os.ReadFile(markerFile)
	if err != nil {
		return nil
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

// A standalone skill: the run goes to /agent/start as framework plus rawConfig,
// and the script tool the server compiles from the document is served by the
// worker this runtime registered under "<skill>__echo_args".
//
// A run that times out is the signature of the worker not polling: the server
// schedules the tool task and nothing picks it up.
func TestSkillScriptRunsAsWorker(t *testing.T) {
	rt := newRuntime(t)
	dir, marker := writeE2ESkill(t)

	skill, err := ai.LoadSkill(dir, ai.WithSkillModel(model(t)))
	if err != nil {
		t.Fatalf("LoadSkill: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	res, err := rt.Run(ctx, skill,
		"tool_parity_proof. Call "+skillName+"__echo_args exactly once with "+
			"tool_parity_proof as the command argument, then return the tool output.")
	if err != nil {
		t.Fatalf("Run: %v (a timeout here means the skill's script worker did not poll)", err)
	}
	t.Logf("execution %s status=%s output=%q", res.ExecutionID, res.Status, res.Output)

	if res.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q", res.Status, ai.StatusCompleted)
	}

	calls := markerCalls(t, marker)
	if len(calls) == 0 {
		t.Fatal("echo_args never ran in this process; the script worker was not dispatched to")
	}
	if !strings.Contains(strings.Join(calls, "\n"), "tool_parity_proof") {
		t.Errorf("echo_args ran with %q, want the argument tool_parity_proof", calls)
	}
	// The instructions say to return the tool's line verbatim, so the answer
	// must carry what the script printed.
	if !strings.Contains(res.Output, "ECHO_ARGS_RESULT:tool_parity_proof") {
		t.Errorf("output does not carry the script's result: %q", res.Output)
	}
}

// A skill as a tool on another agent: the skill document nests under the
// tool's agentConfig with its _framework marker, the server compiles it into a
// sub-workflow, and this runtime must register the skill's workers even though
// the parent agent has no worker tools of its own.
//
// This is the regression the Python suite calls out — skill workers registered
// but never polled when the parent had nothing else to poll for — and the case
// that needs registerWorkers to walk into agents nested under agent tools.
func TestSkillAsAgentTool(t *testing.T) {
	rt := newRuntime(t)
	dir, marker := writeE2ESkill(t)

	skill, err := ai.LoadSkill(dir, ai.WithSkillModel(model(t)))
	if err != nil {
		t.Fatalf("LoadSkill: %v", err)
	}

	parent := &ai.Agent{
		Name:  "go_e2e_skill_as_tool",
		Model: model(t),
		Instructions: "You have one tool: " + skillName + ". " +
			"Call it once with the user's request, then return the result.",
		Tools: []ai.ToolDef{
			tool.Agent(skill, "", "Run the test skill, which echoes its input with echo_args"),
		},
		MaxTurns: 3,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	res, err := rt.Run(ctx, parent, "Echo 'proof42'")
	if err != nil {
		t.Fatalf("Run: %v (a timeout here means the nested skill's workers were not registered)", err)
	}
	t.Logf("execution %s status=%s output=%q", res.ExecutionID, res.Status, res.Output)

	if res.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q", res.Status, ai.StatusCompleted)
	}

	// The script ran in this process, inside the skill's sub-workflow. The
	// exact argument is the skill's model's choice, so only the call is
	// asserted, as the Python suite does.
	calls := markerCalls(t, marker)
	if len(calls) == 0 {
		t.Fatal("echo_args never ran; the skill's workers were not registered for the nested skill")
	}
	t.Logf("echo_args ran %d time(s) with %q", len(calls), calls)
}
