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
	"strings"
	"testing"
)

// The golden fixture sets every field explicitly, so it proves the emitted
// shape but never exercises a default. These are the substitutions that keep Go
// aligned with Python: a Go zero value must not reach the server as 0 where
// Python would have sent 30.
func TestExecutionConfigDefaults(t *testing.T) {
	code := (&CodeExecutionConfig{}).config()
	if got := code["timeout"]; got != defaultExecutionTimeout {
		t.Errorf("code timeout = %v, want %d", got, defaultExecutionTimeout)
	}
	if got, ok := code["allowedLanguages"].([]string); !ok ||
		len(got) != 1 || got[0] != defaultLanguage {
		t.Errorf("allowedLanguages = %v, want [%q]", code["allowedLanguages"], defaultLanguage)
	}
	if got := code["enabled"]; got != true {
		t.Errorf("enabled = %v, want true when nil", got)
	}
	// An empty list must stay a list: a JSON null here would be a different
	// document from the one the other SDKs send.
	if got, ok := code["allowedCommands"].([]string); !ok || got == nil || len(got) != 0 {
		t.Errorf("allowedCommands = %#v, want an empty slice", code["allowedCommands"])
	}

	cli := (&CLIConfig{}).config()
	if got := cli["timeout"]; got != defaultExecutionTimeout {
		t.Errorf("cli timeout = %v, want %d", got, defaultExecutionTimeout)
	}
	if got := cli["enabled"]; got != true {
		t.Errorf("cli enabled = %v, want true when nil", got)
	}
	if got := cli["allowShell"]; got != false {
		t.Errorf("allowShell = %v, want false by default", got)
	}
}

// Enabled=false is the one case where config and tools disagree: the config
// still reaches the server, but the model is not offered the tool.
func TestDisabledExecutionSerializesButAddsNoTool(t *testing.T) {
	off := false
	a := &Agent{
		Name:          "executor",
		Model:         testModel,
		CodeExecution: &CodeExecutionConfig{Enabled: &off},
		CLI:           &CLIConfig{Enabled: &off},
	}

	cfg := a.toConfig()
	if _, ok := cfg["codeExecution"]; !ok {
		t.Error("codeExecution must still be sent when disabled")
	}
	if _, ok := cfg["cliConfig"]; !ok {
		t.Error("cliConfig must still be sent when disabled")
	}
	if tools, ok := cfg["tools"]; ok {
		t.Errorf("a disabled config must contribute no tool, got %v", tools)
	}
	if got := len(a.derivedTools()); got != 0 {
		t.Errorf("derivedTools = %d, want 0", got)
	}
}

// The description is what the model reads, so its optional clauses are
// behaviour rather than cosmetics. The golden fixture covers the inverse of
// both branches below.
func TestDerivedToolDescriptions(t *testing.T) {
	code := (&CodeExecutionConfig{
		AllowedLanguages: []string{"python"},
		AllowedCommands:  []string{"pip", "ls"},
		TimeoutSeconds:   10,
	}).codeTool("a")
	if !strings.Contains(code.Description, "Allowed shell commands: pip, ls.") {
		t.Errorf("code description missing the allowed-commands clause: %q", code.Description)
	}

	// Commands are sorted for a stable description, and the caller's slice must
	// not be reordered underneath them.
	given := []string{"ls", "git"}
	cli := (&CLIConfig{AllowedCommands: given, TimeoutSeconds: 5}).cliTool("a")
	if !strings.Contains(cli.Description, "Allowed commands: git, ls.") {
		t.Errorf("cli description not sorted: %q", cli.Description)
	}
	if given[0] != "ls" {
		t.Errorf("cliTool reordered the caller's slice: %v", given)
	}
	if !strings.Contains(cli.Description, "Shell mode is disabled") {
		t.Errorf("cli description must warn when shell is off: %q", cli.Description)
	}

	// And the opposite: with shell allowed, the warning must be absent.
	allowed := (&CLIConfig{AllowShell: true}).cliTool("a")
	if strings.Contains(allowed.Description, "Shell mode is disabled") {
		t.Errorf("shell warning present although shell is allowed: %q", allowed.Description)
	}
}

// Agent-level Stateful has no key of its own: it marks every tool, including
// ones the caller declared. The fixture only has derived tools, so this is the
// half it cannot check.
func TestAgentStatefulMarksDeclaredTools(t *testing.T) {
	a := &Agent{
		Name:     "s",
		Model:    testModel,
		Stateful: true,
		Tools: []ToolDef{{
			Name:        "plain",
			Description: "No stateful flag of its own",
			InputSchema: map[string]any{"type": "object"},
			ToolType:    ToolTypeWorker,
		}},
	}

	cfg := a.toConfig()
	if _, ok := cfg["stateful"]; ok {
		t.Error("stateful must not be a top-level key")
	}
	tools, ok := cfg["tools"].([]any)
	if !ok || len(tools) != 1 {
		t.Fatalf("tools = %v, want one entry", cfg["tools"])
	}
	tc, _ := tools[0].(map[string]any)
	if tc["stateful"] != true {
		t.Errorf("declared tool not marked stateful: %v", tc)
	}

	// Without the agent flag the key is absent rather than false.
	a.Stateful = false
	tools, _ = a.toConfig()["tools"].([]any)
	tc, _ = tools[0].(map[string]any)
	if _, ok := tc["stateful"]; ok {
		t.Errorf("stateful should be omitted, not false: %v", tc)
	}
}

// Declared tools come before derived ones, and JSON arrays compare by
// position, so the order is part of the contract.
func TestDerivedToolsFollowDeclaredTools(t *testing.T) {
	a := &Agent{
		Name:  "executor",
		Model: testModel,
		Tools: []ToolDef{{
			Name: "mine", InputSchema: map[string]any{"type": "object"},
			ToolType: ToolTypeWorker,
		}},
		CodeExecution: &CodeExecutionConfig{},
		CLI:           &CLIConfig{},
	}

	tools, _ := a.toConfig()["tools"].([]any)
	var names []string
	for _, t := range tools {
		if m, ok := t.(map[string]any); ok {
			names = append(names, m["name"].(string))
		}
	}
	want := []string{"mine", "executor_execute_code", "executor_run_command"}
	if len(names) != len(want) {
		t.Fatalf("tool names = %v, want %v", names, want)
	}
	for i := range want {
		if names[i] != want[i] {
			t.Errorf("tool %d = %q, want %q", i, names[i], want[i])
		}
	}
}
