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
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/ai/internal/schema"
)

// goldenDir holds agentConfig documents captured from the Python SDK's
// serializer. Regenerate them with testdata/agent_config/generate_fixtures.py.
const goldenDir = "testdata/agent_config"

const testModel = "openai/gpt-4o"

// The three leaf agents the Python fixtures share. Kept as constructors so
// each fixture gets its own values and nothing is aliased across cases.
func billing() *Agent {
	return &Agent{Name: "billing", Model: testModel, Instructions: "Handle billing."}
}
func refunds() *Agent {
	return &Agent{Name: "refunds", Model: testModel, Instructions: "Handle refunds."}
}
func tech() *Agent {
	return &Agent{Name: "tech", Model: testModel, Instructions: "Handle tech support."}
}

// goldenFixtures maps each golden file to the Go agent that must reproduce it.
//
// A nil builder means the wire surface that fixture covers is not implemented
// yet; the test reports it as pending rather than failing, so the count of
// green fixtures is the honest progress measure for the port. Replacing a nil
// with a builder is what "adding a feature" means here.
var goldenFixtures = map[string]func() *Agent{
	"01_minimal": func() *Agent {
		return &Agent{Name: "minimal", Model: testModel, Instructions: "You are a helpful assistant."}
	},
	"04_strategy_sequential": func() *Agent {
		return &Agent{
			Name:         "pipeline",
			Model:        testModel,
			Instructions: "Run in order.",
			Strategy:     StrategySequential,
			Agents:       []*Agent{billing(), refunds()},
		}
	},
	"05_strategy_parallel": func() *Agent {
		return &Agent{
			Name:     "fanout",
			Model:    testModel,
			Strategy: StrategyParallel,
			Agents:   []*Agent{billing(), refunds(), tech()},
		}
	},
	"16_nested_tree": func() *Agent {
		return &Agent{
			Name:     "root",
			Model:    testModel,
			Strategy: StrategySequential,
			Agents: []*Agent{
				{Name: "mid", Model: testModel, Strategy: StrategyParallel,
					Agents: []*Agent{billing(), tech()}},
				refunds(),
			},
		}
	},

	"12_llm_knobs": func() *Agent {
		return &Agent{
			Name:  "knobs",
			Model: testModel,
			InstructionsTemplate: &PromptTemplate{
				Name:      "support_prompt",
				Variables: map[string]any{"tier": "${workflow.input.tier}"},
				Version:   Ptr(2),
			},
			MaxTurns:             7,
			TimeoutSeconds:       600,
			MaxTokens:            Ptr(4096),
			Temperature:          Ptr(0.2),
			ReasoningEffort:      ReasoningEffortHigh,
			ThinkingBudgetTokens: Ptr(2048),
			ContextWindowBudget:  Ptr(100000),
			IncludeContents:      "none",
			Memory:               &ConversationMemory{MaxMessages: 25},
		}
	},

	"10_termination": func() *Agent {
		return &Agent{
			Name:  "terminating",
			Model: testModel,
			Termination: AndTermination(
				OrTermination(
					TextMentionTermination{Text: "DONE", CaseSensitive: true},
					MaxMessageTermination{MaxMessages: 20},
				),
				TokenUsageTermination{
					MaxTotalTokens:      Ptr(8000),
					MaxCompletionTokens: Ptr(2000),
				},
			),
			StopWhen: func(context.Context, StopWhenState) (bool, error) { return false, nil },
		}
	},

	"06_strategy_router_agent": func() *Agent {
		return &Agent{
			Name:     "router_agent",
			Model:    testModel,
			Strategy: StrategyRouter,
			Agents:   []*Agent{billing(), tech()},
			Router:   &Agent{Name: "classifier", Model: testModel, Instructions: "Classify."},
		}
	},
	"07_strategy_router_fn": func() *Agent {
		return &Agent{
			Name:       "router_fn",
			Model:      testModel,
			Strategy:   StrategyRouter,
			Agents:     []*Agent{billing(), tech()},
			RouterFunc: func(context.Context, string) (string, error) { return "billing", nil },
		}
	},

	"02_tools_worker": func() *Agent {
		return &Agent{
			Name: "tools_worker", Model: testModel, Instructions: "Use tools.",
			Tools: []ToolDef{
				mkTool("get_weather", "Get the current weather for a city.",
					weatherIn{}, map[string]any{}),
				func() ToolDef {
					t := mkTool("refund", "Issue a refund. Carries per-tool wire options.",
						orderIn{}, "")
					t.ApprovalRequired = true
					t.TimeoutSeconds = Ptr(45)
					t.MaxCalls = Ptr(2)
					return t
				}(),
				func() ToolDef {
					t := mkTool("push_branch",
						"Declares a credential, which lands under config.credentials.",
						branchIn{}, "")
					t.Credentials = []string{"GITHUB_TOKEN"}
					return t
				}(),
			},
		}
	},
	"03_tools_schema_types": func() *Agent {
		return &Agent{
			Name: "tools_schema_types", Model: testModel,
			Instructions: "Pin the JSON Schema mapping.",
			Tools: []ToolDef{
				mkTool("scalar_kinds", "Every scalar type the schema generator maps.",
					scalarKindsIn{}, ""),
				mkTool("container_kinds",
					"List and dict parameters, to pin items/additionalProperties.",
					containerKindsIn{}, map[string]any{}),
			},
		}
	},

	// Pending. Each needs the wire surface named below.
	"08_handoffs":     swarmAgent,
	"09_guardrails":   guardedAgent,
	"11_output_type":  structuredAgent,
	"13_plan_execute": planExecuteAgent,

	"14_tools_nonworker":     nonWorkerToolsAgent,
	"15_execution_and_creds": executorAgent,
	"17_tools_mcp":           mcpToolsAgent,

	"18_skill":         skillAgent,
	"19_skill_as_tool": skillAsToolAgent,
}

// The skill fixtures read testdata/agent_config/skills/review-skill, the same
// directory generate_fixtures.py read, so the embedded file contents match by
// construction and the comparison is about what the loader does with them.
func mustLoadSkill(name string, opts ...SkillOption) *Agent {
	agent, err := LoadSkill(filepath.Join(goldenDir, "skills", name), opts...)
	if err != nil {
		panic(fmt.Sprintf("load fixture skill %s: %v", name, err))
	}
	return agent
}

// Every convention at once: sub-agent files, three scripts (two by extension,
// one by shebang), a loose resource and a references/ file, a cross-skill
// reference to the sibling cleanup-skill, and params merged from frontmatter
// defaults with an override and an addition.
func skillAgent() *Agent {
	return mustLoadSkill("review-skill",
		WithSkillModel(testModel),
		WithAgentModels(map[string]string{"critic": "openai/gpt-4o-mini"}),
		WithSkillParams(map[string]any{"rounds": 1, "style": "terse"}))
}

// The skill nested under an agent tool, with only its frontmatter defaults.
func skillAsToolAgent() *Agent {
	return &Agent{
		Name:         "lead",
		Model:        testModel,
		Instructions: "Delegate reviews.",
		Tools: []ToolDef{{
			Name:        "review",
			Description: "Run the review skill.",
			InputSchema: schema.AgentRequest(),
			ToolType:    ToolTypeAgent,
			Config:      map[string]any{"agent": mustLoadSkill("review-skill", WithSkillModel(testModel))},
		}},
	}
}

// MCP tools: a bare one and one with every option, pinning both the defaults
// and the snake_case config keys the server's compiler reads.
func mcpToolsAgent() *Agent {
	return &Agent{
		Name:         "tools_mcp",
		Model:        testModel,
		Instructions: "MCP tool types.",
		Tools: []ToolDef{
			{
				Name:        "mcp_tools",
				Description: "MCP tools from http://localhost:3001/mcp",
				InputSchema: map[string]any{},
				ToolType:    ToolTypeMCP,
				Config:      map[string]any{"server_url": "http://localhost:3001/mcp", "max_tools": 64},
			},
			{
				Name:        "secured_mcp",
				Description: "Authenticated MCP tools.",
				InputSchema: map[string]any{},
				ToolType:    ToolTypeMCP,
				Config: map[string]any{
					"server_url": "http://localhost:3002/mcp",
					"max_tools":  16,
					"headers":    map[string]string{"Authorization": "Bearer ${MCP_AUTH_KEY}"},
					"tool_names": []string{"get_weather", "math_add"},
				},
				Credentials: []string{"MCP_AUTH_KEY"},
			},
		},
	}
}

// The plan-execute slots. The parent's tools are what the plan may name, so
// they are part of the fixture rather than incidental.
func planExecuteAgent() *Agent {
	return &Agent{
		Name:         "planner_root",
		Model:        testModel,
		Instructions: "Plan then execute.",
		Strategy:     StrategyPlanExecute,
		Tools: []ToolDef{
			mkTool("get_weather", "Get the current weather for a city.",
				weatherIn{}, map[string]any{}),
			mkTool("container_kinds",
				"List and dict parameters, to pin items/additionalProperties.",
				containerKindsIn{}, map[string]any{}),
		},
		Planner: &Agent{
			Name: "the_planner", Model: testModel, Instructions: "Emit JSON plan.",
		},
		Fallback: &Agent{
			Name: "the_fallback", Model: testModel, Instructions: "Best effort.",
		},
		FallbackMaxTurns: 4,
	}
}

// A swarm with all three handoff kinds and a transition allow-list.
func swarmAgent() *Agent {
	sub := func(name, instructions string) *Agent {
		return &Agent{Name: name, Model: testModel, Instructions: instructions}
	}
	return &Agent{
		Name:     "swarm",
		Model:    testModel,
		Strategy: StrategySwarm,
		Agents: []*Agent{
			sub("billing", "Handle billing."),
			sub("refunds", "Handle refunds."),
			sub("tech", "Handle tech support."),
		},
		Handoffs: []HandoffCondition{
			&OnToolResult{Target: "refunds", ToolName: "refund", ResultContains: "ok"},
			&OnTextMention{Target: "tech", Text: "broken"},
			&OnCondition{Target: "billing", Condition: func(context.Context, HandoffState) (bool, error) {
				return false, nil
			}},
		},
		AllowedTransitions: map[string][]string{
			"billing": {"refunds"},
			"refunds": {"tech"},
		},
	}
}

// All three guardrail kinds on one agent: two the server evaluates itself and
// one backed by a worker.
func guardedAgent() *Agent {
	return &Agent{
		Name:         "guarded",
		Model:        testModel,
		Instructions: "Be careful.",
		Guardrails: []Guardrail{
			&RegexGuardrail{
				Patterns: []string{`\d{16}`, `sk-\w+`},
				Mode:     "block",
				Message:  "no cards",
			},
			&LLMGuardrail{
				Model:     testModel,
				Policy:    "No medical advice.",
				MaxTokens: 256,
			},
			&CustomGuardrail{
				guardrailBase: guardrailBase{
					Name:     "no_secrets",
					Position: PositionOutput,
					OnFail:   OnFailRetry,
				},
				Check: func(context.Context, GuardrailInput) (GuardrailResult, error) {
					return GuardrailResult{Passed: true}, nil
				},
			},
		},
	}
}

// A ticket, as the structured-output fixture's answer type.
type Ticket struct {
	Summary  string   `json:"summary"`
	Priority int      `json:"priority"`
	Tags     []string `json:"tags"`
}

func structuredAgent() *Agent {
	return &Agent{
		Name:         "structured",
		Model:        testModel,
		Instructions: "Return a ticket.",
		OutputType:   Ticket{},
	}
}

// The non-worker tool types: the server dispatches all three itself, so none
// of them registers a Go function.
func nonWorkerToolsAgent() *Agent {
	return &Agent{
		Name:         "tools_nonworker",
		Model:        testModel,
		Instructions: "Non-worker tool types.",
		Tools: []ToolDef{
			{
				Name:        "lookup",
				Description: "Look up a record.",
				InputSchema: schema.EmptyObject(),
				ToolType:    ToolTypeHTTP,
				Config: map[string]any{
					"url":         "https://example.test/api/{id}",
					"method":      "GET",
					"headers":     map[string]string{"X-Api-Version": "2"},
					"accept":      []string{"application/json"},
					"contentType": "application/json",
				},
			},
			{
				Name:        "ask_human",
				Description: "Ask a person to decide.",
				InputSchema: schema.HumanInput(),
				ToolType:    ToolTypeHuman,
			},
			{
				Name:        "delegate_billing",
				Description: "Delegate.",
				InputSchema: schema.AgentRequest(),
				ToolType:    ToolTypeAgent,
				Config: map[string]any{"agent": &Agent{
					Name:         "billing",
					Model:        testModel,
					Instructions: "Handle billing.",
				}},
			},
		},
	}
}

// The executor fixture: both execution configs, their derived tools, and the
// agent-level fields that ride alongside them.
func executorAgent() *Agent {
	return &Agent{
		Name:         "executor",
		Model:        testModel,
		Instructions: "Run code and commands.",
		CodeExecution: &CodeExecutionConfig{
			AllowedLanguages: []string{"python", "bash"},
			TimeoutSeconds:   60,
		},
		CLI: &CLIConfig{
			AllowedCommands: []string{"git", "ls"},
			TimeoutSeconds:  45,
			AllowShell:      true,
		},
		Credentials:   []string{"GITHUB_TOKEN", "OPENAI_API_KEY"},
		RequiredTools: []string{"get_weather"},
		MaskedFields:  []string{"ssn", "card"},
		Introduction:  "Hi, I run code.",
		Metadata:      map[string]any{"team": "platform", "tier": 2},
		Stateful:      true,
	}
}

// TestGoldenAgentConfig checks the Go serializer against documents captured
// from the Python SDK. Comparison is semantic rather than byte-for-byte,
// because key ordering legitimately differs between languages.
func TestGoldenAgentConfig(t *testing.T) {
	entries, err := os.ReadDir(goldenDir)
	if err != nil {
		t.Fatalf("read golden dir: %v", err)
	}

	var names []string
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".json") {
			names = append(names, strings.TrimSuffix(e.Name(), ".json"))
		}
	}
	if len(names) == 0 {
		t.Fatalf("no golden files in %s", goldenDir)
	}
	sort.Strings(names)

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			build, declared := goldenFixtures[name]
			if !declared {
				t.Fatalf("golden file %s.json has no entry in goldenFixtures; "+
					"add a builder or an explicit nil", name)
			}
			if build == nil {
				t.Skip("pending: wire surface not implemented yet")
			}

			agent := build()
			if err := agent.Validate(); err != nil {
				t.Fatalf("fixture does not validate: %v", err)
			}

			got := dropSchemaTitles(normalize(t, agent.toConfig()))
			want := dropSchemaTitles(readGolden(t, name))

			if !reflect.DeepEqual(got, want) {
				t.Errorf("agentConfig mismatch for %s\n--- got ---\n%s\n--- want ---\n%s",
					name, mustIndent(t, got), mustIndent(t, want))
			}
		})
	}
}

// TestGoldenCoverage reports how much of the wire surface is implemented.
// It never fails; it exists so the ratchet is visible in test output.
func TestGoldenCoverage(t *testing.T) {
	var done, pending []string
	for name, build := range goldenFixtures {
		if build == nil {
			pending = append(pending, name)
		} else {
			done = append(done, name)
		}
	}
	sort.Strings(done)
	sort.Strings(pending)
	t.Logf("golden fixtures implemented: %d/%d", len(done), len(done)+len(pending))
	if len(pending) > 0 {
		t.Logf("pending: %s", strings.Join(pending, ", "))
	}
}

// normalize round-trips a config through JSON so both sides of the comparison
// use the same Go types (all numbers become float64, all maps map[string]any).
// dropSchemaTitles removes "title" from the outputType schema on both sides of
// the comparison.
//
// The wire schema declares outputType.schema with additionalProperties true —
// an opaque JSON Schema document it deliberately does not constrain. Python
// builds that document with Pydantic, which adds a title per field and one for
// the model; Java generates it from declared fields and adds none. Both are
// conformant, so Go cannot match one without diverging from the other, and
// matching Python here would mean carrying a Pydantic artifact that means
// nothing in Go.
//
// Everything else in the fixture is still compared exactly, including the
// property types, the required list and its order, and className.
func dropSchemaTitles(cfg map[string]any) map[string]any {
	ot, ok := cfg["outputType"].(map[string]any)
	if !ok {
		return cfg
	}
	sch, ok := ot["schema"].(map[string]any)
	if !ok {
		return cfg
	}
	stripTitles(sch)
	return cfg
}

func stripTitles(node any) {
	switch v := node.(type) {
	case map[string]any:
		delete(v, "title")
		for _, child := range v {
			stripTitles(child)
		}
	case []any:
		for _, child := range v {
			stripTitles(child)
		}
	}
}

func normalize(t *testing.T, cfg map[string]any) map[string]any {
	t.Helper()
	raw, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal config: %v", err)
	}
	var out map[string]any
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatalf("unmarshal config: %v", err)
	}
	return out
}

func readGolden(t *testing.T, name string) map[string]any {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(goldenDir, name+".json"))
	if err != nil {
		t.Fatalf("read golden %s: %v", name, err)
	}
	var out map[string]any
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatalf("parse golden %s: %v", name, err)
	}
	return out
}

func mustIndent(t *testing.T, v any) string {
	t.Helper()
	raw, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatalf("indent: %v", err)
	}
	return string(raw)
}
