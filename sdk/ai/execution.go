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
	"fmt"
	"sort"
	"strings"
)

// Python applies these defaults when a field is left out, so Go substitutes
// them for its zero values rather than sending a 0 the server would enforce.
const (
	defaultExecutionTimeout = 30
	defaultLanguage         = "python"
)

// CodeExecutionConfig lets the agent run code the model writes, through a
// derived tool named "{agent}_execute_code" whose config the server enforces.
type CodeExecutionConfig struct {
	// Enabled defaults to true when nil; false keeps the config on the wire but offers no tool, as in Python.
	Enabled *bool
	// AllowedLanguages defaults to ["python"] when empty.
	AllowedLanguages []string
	// AllowedCommands restricts shell commands the code may invoke; empty means no restriction.
	AllowedCommands []string
	// TimeoutSeconds defaults to 30 when zero.
	TimeoutSeconds int
	// Executor runs the code. Nil means a LocalExecutor, a subprocess on the worker host with the
	// language the model asked for; Docker, Jupyter and Serverless executors ignore that language
	// and run what they were built for, as in Python. The executor never travels to the server,
	// so agentConfig is the same whichever is set.
	Executor CodeExecutor
}

// CLIConfig lets the agent run shell commands directly. Attaching it gives
// the agent a derived tool named "{agent}_run_command".
type CLIConfig struct {
	// Enabled defaults to true when nil, as in CodeExecutionConfig.
	Enabled *bool
	// AllowedCommands restricts what may be run. Empty means no restriction.
	AllowedCommands []string
	// TimeoutSeconds defaults to 30 when zero.
	TimeoutSeconds int
	// AllowShell permits shell=true on a call; false by default, and the derived tool's description says so.
	AllowShell bool
}

func enabledOrDefault(v *bool) bool { return v == nil || *v }

func timeoutOrDefault(v int) int {
	if v <= 0 {
		return defaultExecutionTimeout
	}
	return v
}

func languagesOrDefault(v []string) []string {
	if len(v) == 0 {
		return []string{defaultLanguage}
	}
	return v
}

// nonNil avoids a JSON null: both configs always send allowedCommands.
func nonNil(v []string) []string {
	if v == nil {
		return []string{}
	}
	return v
}

// config emits the four keys Python emits unconditionally, with no omission rules.
func (c *CodeExecutionConfig) config() map[string]any {
	return map[string]any{
		"enabled":          enabledOrDefault(c.Enabled),
		"allowedLanguages": languagesOrDefault(c.AllowedLanguages),
		"allowedCommands":  nonNil(c.AllowedCommands),
		"timeout":          timeoutOrDefault(c.TimeoutSeconds),
	}
}

func (c *CLIConfig) config() map[string]any {
	return map[string]any{
		"enabled":         enabledOrDefault(c.Enabled),
		"allowedCommands": nonNil(c.AllowedCommands),
		"timeout":         timeoutOrDefault(c.TimeoutSeconds),
		"allowShell":      c.AllowShell,
	}
}

// The derived tools' schemas are fixed rather than reflected: they describe the
// server's own executors, and the other SDKs send exactly these shapes.
func dictOutputSchema() map[string]any {
	return map[string]any{"type": "object", "additionalProperties": map[string]any{}}
}

// codeTool builds the derived execute_code tool; its description carries the
// config values because that is what the model reads.
func (c *CodeExecutionConfig) codeTool(agentName string) ToolDef {
	langs := languagesOrDefault(c.AllowedLanguages)
	timeout := timeoutOrDefault(c.TimeoutSeconds)

	desc := fmt.Sprintf(
		"Execute code in a sandboxed environment. Supported languages: %s. Timeout: %ds.",
		strings.Join(langs, ", "), timeout)
	if len(c.AllowedCommands) > 0 {
		desc += fmt.Sprintf(" Allowed shell commands: %s.", strings.Join(c.AllowedCommands, ", "))
	}

	return ToolDef{
		Name:        agentName + "_execute_code",
		Description: desc,
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"code":     map[string]any{"type": "string"},
				"language": map[string]any{"type": "string"},
			},
			"required": []string{"code"},
		},
		OutputSchema: dictOutputSchema(),
		ToolType:     ToolTypeWorker,
	}
}

func (c *CLIConfig) cliTool(agentName string) ToolDef {
	timeout := timeoutOrDefault(c.TimeoutSeconds)

	desc := fmt.Sprintf("Run a CLI command directly. Timeout: %ds.", timeout)
	if len(c.AllowedCommands) > 0 {
		// Sorted on a copy: a stable description, and the caller's slice untouched.
		sorted := append([]string(nil), c.AllowedCommands...)
		sort.Strings(sorted)
		desc += fmt.Sprintf(" Allowed commands: %s.", strings.Join(sorted, ", "))
	}
	if !c.AllowShell {
		desc += " Shell mode is disabled — do not set shell=True."
	}
	desc += " If you need to save a command's output for later pipeline steps," +
		" set context_key. Well-known keys: repo, branch, working_dir," +
		" issue_number, pr_url, commit_sha."

	return ToolDef{
		Name:        agentName + "_run_command",
		Description: desc,
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"command":     map[string]any{"type": "string"},
				"args":        map[string]any{"type": "array", "items": map[string]any{}},
				"cwd":         map[string]any{"type": "string"},
				"shell":       map[string]any{"type": "boolean"},
				"context_key": map[string]any{"type": "string"},
			},
			"required": []string{"command"},
		},
		OutputSchema: dictOutputSchema(),
		ToolType:     ToolTypeWorker,
	}
}

// derivedTools returns the tools the execution configs imply, in the order the
// other SDKs append them. A disabled config still serializes but adds no tool.
func (a *Agent) derivedTools() []ToolDef {
	var out []ToolDef
	if a.CodeExecution != nil && enabledOrDefault(a.CodeExecution.Enabled) {
		out = append(out, a.CodeExecution.codeTool(a.Name))
	}
	if a.CLI != nil && enabledOrDefault(a.CLI.Enabled) {
		out = append(out, a.CLI.cliTool(a.Name))
	}
	return out
}
