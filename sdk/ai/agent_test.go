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
	"reflect"
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	cases := []struct {
		name    string
		agent   *Agent
		wantErr string // substring; empty means the agent must validate
	}{
		{
			name:  "minimal is valid",
			agent: &Agent{Name: "ok", Model: testModel},
		},
		{
			name:  "underscore start is valid",
			agent: &Agent{Name: "_private", Model: testModel},
		},
		{
			name:  "hyphens and digits are valid",
			agent: &Agent{Name: "agent-2_b", Model: testModel},
		},
		{
			name:    "empty name",
			agent:   &Agent{Model: testModel},
			wantErr: "non-empty string",
		},
		{
			name:    "leading digit",
			agent:   &Agent{Name: "2fast", Model: testModel},
			wantErr: "invalid agent name",
		},
		{
			name:    "space in name",
			agent:   &Agent{Name: "my agent", Model: testModel},
			wantErr: "invalid agent name",
		},
		{
			name:    "unknown strategy",
			agent:   &Agent{Name: "a", Model: testModel, Strategy: Strategy("teamwork")},
			wantErr: `invalid strategy "teamwork"`,
		},
		{
			name:  "empty strategy is allowed and defaults later",
			agent: &Agent{Name: "a", Model: testModel},
		},
		{
			name:    "negative maxTurns",
			agent:   &Agent{Name: "a", Model: testModel, MaxTurns: -1},
			wantErr: "maxTurns must be >= 0",
		},
		{
			// The Python SDK does not validate timeouts, so neither do we;
			// the server owns range checks. This case pins that choice so a
			// future range check has to delete a test rather than slip in.
			name:  "timeouts are not range-checked, matching Python",
			agent: &Agent{Name: "a", Model: testModel, TimeoutSeconds: -1},
		},
		{
			name:  "empty reasoningEffort is unset, not invalid",
			agent: &Agent{Name: "a", Model: testModel, ReasoningEffort: ""},
		},
		{
			name:  "minimal is a valid reasoningEffort",
			agent: &Agent{Name: "a", Model: testModel, ReasoningEffort: ReasoningEffortMinimal},
		},
		{
			// A named string type does not stop this compiling, which is why
			// the check has to exist at all.
			name:    "typo in reasoningEffort",
			agent:   &Agent{Name: "a", Model: testModel, ReasoningEffort: "hgih"},
			wantErr: `invalid reasoningEffort "hgih"`,
		},
		{
			name:    "reasoningEffort is case sensitive",
			agent:   &Agent{Name: "a", Model: testModel, ReasoningEffort: "High"},
			wantErr: `invalid reasoningEffort "High"`,
		},
		{
			name: "valid termination tree",
			agent: &Agent{Name: "a", Model: testModel, Termination: AndTermination(
				OrTermination(
					TextMentionTermination{Text: "DONE"},
					MaxMessageTermination{MaxMessages: 20},
				),
				TokenUsageTermination{MaxTotalTokens: Ptr(8000)},
			)},
		},
		{
			// Python's constructor requires the argument; a Go struct literal
			// does not, so the zero value has to be rejected here instead.
			name:    "zero-value MaxMessageTermination",
			agent:   &Agent{Name: "a", Model: testModel, Termination: MaxMessageTermination{}},
			wantErr: "maxMessages must be >= 1",
		},
		{
			name:    "TokenUsageTermination with no limits",
			agent:   &Agent{Name: "a", Model: testModel, Termination: TokenUsageTermination{}},
			wantErr: "at least one token limit must be set",
		},
		{
			name:    "empty composite",
			agent:   &Agent{Name: "a", Model: testModel, Termination: AndTermination()},
			wantErr: "AndTermination: at least one condition is required",
		},
		{
			name: "invalid condition nested inside a composite",
			agent: &Agent{Name: "a", Model: testModel, Termination: OrTermination(
				TextMentionTermination{Text: "DONE"},
				AndTermination(TokenUsageTermination{}),
			)},
			wantErr: "at least one token limit must be set",
		},
		{
			name: "instructions and template together",
			agent: &Agent{
				Name:                 "a",
				Model:                testModel,
				Instructions:         "hi",
				InstructionsTemplate: &PromptTemplate{Name: "tpl"},
			},
			wantErr: "not both",
		},
		{
			name: "template without a name",
			agent: &Agent{
				Name:                 "a",
				Model:                testModel,
				InstructionsTemplate: &PromptTemplate{},
			},
			wantErr: "InstructionsTemplate.Name is required",
		},
		{
			name: "invalid sub-agent is reported with the parent's name",
			agent: &Agent{
				Name:   "parent",
				Model:  testModel,
				Agents: []*Agent{{Name: "2bad", Model: testModel}},
			},
			wantErr: `agent "parent": invalid agent name "2bad"`,
		},
		{
			name: "invalid grandchild is reported through the chain",
			agent: &Agent{
				Name:  "root",
				Model: testModel,
				Agents: []*Agent{{
					Name:   "mid",
					Model:  testModel,
					Agents: []*Agent{{Name: "", Model: testModel}},
				}},
			},
			wantErr: `agent "root": agent "mid": agent name must be a non-empty string`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.agent.Validate()
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("want valid, got error: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("want error containing %q, got nil", tc.wantErr)
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("want error containing %q, got %q", tc.wantErr, err.Error())
			}
		})
	}
}

func TestValidateNilAgent(t *testing.T) {
	var a *Agent
	if err := a.Validate(); err == nil {
		t.Fatal("want an error for a nil agent, got nil")
	}
}

// A leaf agent must not send a strategy even though Strategy defaults to
// handoff, because the server dispatches on its presence.
func TestStrategyOmittedWithoutSubAgents(t *testing.T) {
	leaf := (&Agent{Name: "leaf", Model: testModel}).toConfig()
	if _, ok := leaf["strategy"]; ok {
		t.Errorf("leaf agent must not send strategy, got %v", leaf["strategy"])
	}

	explicit := (&Agent{Name: "leaf", Model: testModel, Strategy: StrategyParallel}).toConfig()
	if _, ok := explicit["strategy"]; ok {
		t.Errorf("strategy without sub-agents must not be sent, got %v", explicit["strategy"])
	}

	parent := (&Agent{
		Name:   "parent",
		Model:  testModel,
		Agents: []*Agent{{Name: "child", Model: testModel}},
	}).toConfig()
	if got := parent["strategy"]; got != string(StrategyHandoff) {
		t.Errorf("parent with sub-agents must default to handoff, got %v", got)
	}
}

// maxTurns must always reach the wire with the SDK default substituted, never
// omitted and never sent as 0.
//
// The schema's default differs by side: the SDK default is 25, the server
// default is 100. Omitting the key therefore does not mean "use 25", it means
// "use 100" — a silent 4x change to the agent loop budget with no error
// anywhere. Go is the only SDK exposed to this, because 0 is its zero value
// and so is indistinguishable from unset; Python and Java carry 25 as a real
// parameter default and cannot land here.
//
// If a future change makes the serializer omit zero-valued fields, this test
// is what catches it.
func TestMaxTurnsAlwaysSentWithDefault(t *testing.T) {
	unset := (&Agent{Name: "x", Model: testModel}).toConfig()
	got, ok := unset["maxTurns"]
	if !ok {
		t.Fatal("maxTurns was omitted; the server would apply 100 instead of 25")
	}
	if got != defaultMaxTurns {
		t.Errorf("unset maxTurns must serialize as %d, got %v", defaultMaxTurns, got)
	}

	explicit := (&Agent{Name: "x", Model: testModel, MaxTurns: 7}).toConfig()
	if explicit["maxTurns"] != 7 {
		t.Errorf("explicit maxTurns must be preserved, got %v", explicit["maxTurns"])
	}
}

// timeoutSeconds is the mirror image of maxTurns: 0 is a meaningful wire
// value, not an absence.
//
// The schema defines "0 → server default", and all three SDKs default the
// field to 0 and emit it unconditionally (Python `timeout_seconds: int = 0`,
// Java `private int timeoutSeconds = 0`). So Go must send 0 through untouched.
// Substituting a client-side default here — the helpful-looking fix that is
// correct for maxTurns — is precisely the bug: it would take the choice away
// from the server and away from every other SDK.
//
// The one place a non-zero default is right is the scatter_gather helper,
// which sets 300s because it waits on N parallel sub-agents. That belongs to
// the helper, not to Agent, and must stay there when it is ported.
func TestTimeoutSecondsSentVerbatim(t *testing.T) {
	unset := (&Agent{Name: "x", Model: testModel}).toConfig()
	got, ok := unset["timeoutSeconds"]
	if !ok {
		t.Fatal("timeoutSeconds must always be sent; Python and Java both emit it unconditionally")
	}
	if got != 0 {
		t.Errorf("unset timeoutSeconds must serialize as 0, meaning server default; got %v", got)
	}

	explicit := (&Agent{Name: "x", Model: testModel, TimeoutSeconds: 600}).toConfig()
	if explicit["timeoutSeconds"] != 600 {
		t.Errorf("explicit timeoutSeconds must be preserved, got %v", explicit["timeoutSeconds"])
	}
}

// Optional scalars must distinguish "set to zero" from "not set". Python
// guards them with `is not None`, so temperature=0.0 and maxTokens=0 both
// reach the server — verified against the Python serializer directly. A plain
// float64 or int in Go would collapse those two states and silently drop a
// deliberate request for deterministic output.
func TestOptionalScalarsDistinguishZeroFromUnset(t *testing.T) {
	zero := (&Agent{
		Name: "z", Model: testModel,
		Temperature: Ptr(0.0), MaxTokens: Ptr(0), ContextWindowBudget: Ptr(0),
	}).toConfig()

	for _, k := range []string{"temperature", "maxTokens", "contextWindowBudget"} {
		v, ok := zero[k]
		if !ok {
			t.Errorf("%s set to zero must still be sent; it was omitted", k)
			continue
		}
		if fmt.Sprint(v) != "0" {
			t.Errorf("%s set to zero must serialize as 0, got %v", k, v)
		}
	}

	unset := (&Agent{Name: "u", Model: testModel}).toConfig()
	for _, k := range []string{
		"temperature", "maxTokens", "contextWindowBudget",
		"reasoningEffort", "includeContents", "thinkingConfig", "memory",
	} {
		if _, ok := unset[k]; ok {
			t.Errorf("unset %s must be omitted, got %v", k, unset[k])
		}
	}
}

// A thinking budget expands into an object; setting it is what enables
// thinking, so "enabled" is always true when the key is present.
func TestThinkingBudgetExpandsToConfig(t *testing.T) {
	cfg := (&Agent{Name: "t", Model: testModel, ThinkingBudgetTokens: Ptr(2048)}).toConfig()
	want := map[string]any{"enabled": true, "budgetTokens": 2048}
	if !reflect.DeepEqual(cfg["thinkingConfig"], want) {
		t.Errorf("thinkingConfig = %v, want %v", cfg["thinkingConfig"], want)
	}
}

// Memory's inner fields use emptiness, not nil-ness, matching Python's
// truthiness checks. A non-nil but empty memory still sends "memory": {},
// because having memory is distinct from having none.
func TestMemoryEmissionMatchesPythonTruthiness(t *testing.T) {
	empty := (&Agent{Name: "m", Model: testModel, Memory: &ConversationMemory{}}).toConfig()
	if got, ok := empty["memory"]; !ok {
		t.Error("non-nil empty memory must still be sent")
	} else if !reflect.DeepEqual(got, map[string]any{}) {
		t.Errorf("empty memory must serialize as {}, got %v", got)
	}

	// MaxMessages of 0 is dropped, exactly as Python's truthiness drops it.
	zeroMax := (&Agent{Name: "m", Model: testModel,
		Memory: &ConversationMemory{MaxMessages: 0}}).toConfig()
	if !reflect.DeepEqual(zeroMax["memory"], map[string]any{}) {
		t.Errorf("maxMessages of 0 must be omitted, got %v", zeroMax["memory"])
	}

	msgs := []map[string]any{{"role": "user", "content": "hi"}}
	withMsgs := (&Agent{Name: "m", Model: testModel,
		Memory: &ConversationMemory{Messages: msgs, MaxMessages: 25}}).toConfig()
	mem, _ := withMsgs["memory"].(map[string]any)
	if mem["maxMessages"] != 25 {
		t.Errorf("maxMessages = %v, want 25", mem["maxMessages"])
	}
	if !reflect.DeepEqual(mem["messages"], msgs) {
		t.Errorf("messages = %v, want %v", mem["messages"], msgs)
	}
}

// The ReasoningEffort constants must stay exactly the "reasoningEffort" enum
// in agent-schema.json. "minimal" is easy to drop, because both the Python
// docstring and the java-sdk field reference list only low/medium/high.
//
// The enum cannot be checked against the schema from here without a cross-repo
// dependency at test time; that belongs in java-sdk's verify.py. This test
// pins the values so a change is at least deliberate.
func TestReasoningEffortValues(t *testing.T) {
	want := []ReasoningEffort{"minimal", "low", "medium", "high"}
	got := []ReasoningEffort{
		ReasoningEffortMinimal, ReasoningEffortLow,
		ReasoningEffortMedium, ReasoningEffortHigh,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("reasoning effort constants = %v, want %v", got, want)
	}

	// The named type must reach the wire as a plain JSON string, not as a
	// wrapper the server would fail to parse.
	cfg := (&Agent{Name: "r", Model: testModel, ReasoningEffort: ReasoningEffortHigh}).toConfig()
	if v, ok := cfg["reasoningEffort"].(string); !ok || v != "high" {
		t.Errorf("reasoningEffort must serialize as the string \"high\", got %#v", cfg["reasoningEffort"])
	}
}

// An empty testModel is omitted so a sub-agent can inherit the parent's at
// compile time; sending "" would override that.
func TestEmptyModelOmitted(t *testing.T) {
	cfg := (&Agent{Name: "inherits"}).toConfig()
	if _, ok := cfg["testModel"]; ok {
		t.Errorf("empty testModel must be omitted, got %v", cfg["testModel"])
	}
}
