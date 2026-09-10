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
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
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
	"08_handoffs":            nil, // handoff conditions + allowedTransitions
	"09_guardrails":          nil, // regex / llm / custom guardrails
	"11_output_type":         nil, // struct to JSON Schema
	"13_plan_execute":        nil, // planner and fallback slots
	"14_tools_nonworker":     nil, // http / human / agent tools
	"15_execution_and_creds": nil, // code execution, CLI, credentials
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

			got := normalize(t, agent.toConfig())
			want := readGolden(t, name)

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
