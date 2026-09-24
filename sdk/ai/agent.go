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
	"regexp"
)

// validName mirrors Python's _VALID_NAME_RE: the name becomes a Conductor workflow name.
var validName = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_-]*$`)

// defaultMaxTurns matches the Python and Java SDKs. It is always sent.
const defaultMaxTurns = 25

// PromptTemplate references a prompt template stored on the Conductor server; the SDK
// never creates them. Variables may hold expressions like "${workflow.input.user_tier}".
type PromptTemplate struct {
	Name      string
	Variables map[string]any
	// Version selects a template version. Nil means latest.
	Version *int
}

// ReasoningEffort selects how much reasoning an OpenAI reasoning model does; other models
// ignore it. The values are the "reasoningEffort" enum of the shared agent-schema.json, which
// includes "minimal" although the Python and Java references omit it. Validate rejects the
// rest, which a named string type alone would not: an untyped constant typo still compiles.
type ReasoningEffort string

const (
	// ReasoningEffortMinimal does the least reasoning before answering.
	ReasoningEffortMinimal ReasoningEffort = "minimal"
	ReasoningEffortLow     ReasoningEffort = "low"
	ReasoningEffortMedium  ReasoningEffort = "medium"
	ReasoningEffortHigh    ReasoningEffort = "high"
)

// validReasoningEfforts mirrors the schema's enum.
var validReasoningEfforts = map[ReasoningEffort]struct{}{
	ReasoningEffortMinimal: {},
	ReasoningEffortLow:     {},
	ReasoningEffortMedium:  {},
	ReasoningEffortHigh:    {},
}

// ConversationMemory seeds an agent with prior messages and bounds how many it retains.
// Both fields are omitted from the wire when empty, so an empty ConversationMemory still
// sends "memory": {}: non-nil memory means stateful, distinct from having no memory.
type ConversationMemory struct {
	// Messages are prior turns, each a map of "role" to a role name and
	// "message" to the text, as the Python SDK's memory builds them.
	Messages []map[string]any
	// MaxMessages caps retained history. Zero means no explicit cap.
	MaxMessages int
}

// Agent is a declarative description of an agent, sent to the server for compilation
// into a Conductor workflow. It is a configuration document rather than a constructor
// call: set only the fields you need and leave the rest at their zero value, which is
// what "unset" means on the wire. Optional numeric settings are pointers because zero
// is a meaningful value; set them with Ptr.
//
//	agent := &ai.Agent{
//	    Name:         "weather_bot",
//	    Model:        "openai/gpt-4o",
//	    Instructions: "Use tools to answer questions.",
//	}
type Agent struct {
	// Name becomes the Conductor workflow name. Required.
	Name string
	// Model is a "provider/model" identifier, e.g. "openai/gpt-4o". Empty on a
	// sub-agent inherits the parent's model at compile time.
	Model string
	// BaseURL points this agent's provider at another endpoint, such as a proxy;
	// empty uses the server's configured URL.
	BaseURL string
	// Instructions is the system prompt. Mutually exclusive with InstructionsTemplate.
	Instructions string
	// InstructionsTemplate uses a named server-side prompt template instead of a
	// literal Instructions string.
	InstructionsTemplate *PromptTemplate

	// Tools the agent may call. Build them with the tool package.
	Tools []ToolDef
	// RequiredTools names tools the run should call before it finishes (Python's required_tools).
	// Unusable on current servers: it compiles to nested DO_WHILEs, unsupported, and deadlocks.
	RequiredTools []string

	// Planner and Fallback are the StrategyPlanExecute slots: the planner emits a
	// plan over the parent's Tools, and Fallback runs if the plan cannot be carried out.
	Planner  *Agent
	Fallback *Agent
	// FallbackMaxTurns bounds the fallback agent; zero omits it, leaving the limit to the server.
	FallbackMaxTurns int
	// EnablePlanning makes the server append a fixed "plan first, then execute
	// step by step" paragraph to this agent's instructions. Prompt text only,
	// and unrelated to Planner.
	EnablePlanning bool
	// PlannerContext is extra text or server-fetched URLs for the Planner. StrategyPlanExecute only.
	PlannerContext []PlanContext
	// PlanSource (Python's plan_source) supplies the plan from an expression the server
	// evaluates instead of the Planner; sent as written, see the server's plan-source docs.
	PlanSource map[string]any
	// Synthesize controls the final LLM step that combines the specialists'
	// results. Nil keeps the server's default of true; Ptr(false) skips it.
	Synthesize *bool
	// PrefillTools run with fixed arguments before the first LLM turn and their results open
	// the conversation. Build them with Prefill.
	PrefillTools []PrefillToolCall
	// Gate decides, after this agent finishes inside a sequential pipeline, whether
	// the pipeline continues to the next agent. See TextGate and GateFunc.
	Gate GateCondition

	// Handoffs move control between sub-agents, usually with StrategySwarm.
	Handoffs []HandoffCondition
	// AllowedTransitions maps a sub-agent name to its permitted targets; empty means no restriction.
	AllowedTransitions map[string][]string

	// Guardrails check this agent's input or output; tools carry their own.
	Guardrails []Guardrail

	// Callbacks are hooks the server calls around the agent, each LLM call and each tool
	// call. Each set hook runs as a worker.
	Callbacks *Callbacks

	// OutputType constrains the final answer to a struct's shape: pass a zero value, as in
	// OutputType: Ticket{}, and the schema comes from its json tags as a tool's input schema
	// does. The wire schema leaves the inner document open, so Go and Java send types only,
	// while Python also carries Pydantic's per-field titles.
	OutputType any

	// CodeExecution and CLI add derived tools for running code and shell commands, both
	// server side; see CodeExecutionConfig.
	CodeExecution *CodeExecutionConfig
	CLI           *CLIConfig

	// Credentials are secret names available to every tool on this agent, as a
	// fallback for tools that cannot name their own. Prefer tool.WithCredentials.
	Credentials []string

	// MaskedFields names input and output fields the server redacts from execution history and the UI.
	MaskedFields []string

	// Introduction is the agent's opening message to a user.
	Introduction string

	// Metadata is arbitrary data carried with the agent definition.
	Metadata map[string]any

	// Stateful routes every tool on this agent to a per-execution worker domain, so one run's
	// calls reach the same process. It has no wire key: it stamps stateful onto each tool.
	Stateful bool

	// Agents are sub-agents; a non-empty value makes this a multi-agent system and sends Strategy.
	Agents []*Agent
	// Strategy orchestrates Agents. Defaults to StrategyHandoff.
	Strategy Strategy
	// Router selects one sub-agent using an LLM; required by StrategyRouter unless RouterFunc is set.
	Router *Agent
	// RouterFunc selects one sub-agent with Go code, run as a worker.
	RouterFunc RouterFunc

	// MaxTurns bounds the agent loop. Zero means the default of 25.
	MaxTurns int
	// TimeoutSeconds bounds the whole execution. Zero means no explicit limit.
	TimeoutSeconds int

	// MaxTokens caps tokens generated per LLM call; nil leaves it to the model.
	MaxTokens *int
	// Temperature is the sampling temperature. Nil is unset; Ptr(0) is sent as 0.
	Temperature *float64
	// ReasoningEffort applies to OpenAI reasoning models. Empty means unset.
	ReasoningEffort ReasoningEffort
	// ThinkingBudgetTokens enables extended thinking, serialized as thinkingConfig.
	ThinkingBudgetTokens *int
	// ContextWindowBudget is the token threshold for proactive context condensing.
	ContextWindowBudget *int
	// IncludeContents controls context inheritance: "none" starts fresh, empty inherits.
	IncludeContents string
	// Memory seeds the agent with prior conversation.
	Memory *ConversationMemory

	// Termination stops the loop on a server-evaluated condition tree.
	Termination TerminationCondition
	// StopWhen stops the loop on Go logic run as a worker, independently of Termination.
	StopWhen StopWhenFunc

	// External marks the agent as served elsewhere; no workers are started locally.
	External bool

	// skill is set by LoadSkill: the agent serializes as the raw skill document the
	// server's SkillNormalizer compiles, not agentConfig, so only Name and Model apply.
	skill *skillConfig
}

// Ptr returns a pointer to v, for setting optional fields inline.
func Ptr[T any](v T) *T { return &v }

// Validate reports the first configuration error in the agent tree. The rules are cross-field,
// so no type signature can express them; runtime methods call Validate before serializing.
func (a *Agent) Validate() error {
	if a == nil {
		return fmt.Errorf("agent is nil")
	}
	if err := a.validateIdentity(); err != nil {
		return err
	}
	if err := a.validateInstructions(); err != nil {
		return err
	}
	for _, t := range a.Tools {
		if err := t.Validate(); err != nil {
			return fmt.Errorf("agent %q: %w", a.Name, err)
		}
	}
	for _, p := range a.PrefillTools {
		if err := p.Tool.Validate(); err != nil {
			return fmt.Errorf("agent %q prefill: %w", a.Name, err)
		}
	}
	if err := a.validateRouting(); err != nil {
		return err
	}
	if err := a.validateComposition(); err != nil {
		return err
	}
	if a.Termination != nil {
		if err := a.Termination.validateTermination(); err != nil {
			return fmt.Errorf("agent %q: %w", a.Name, err)
		}
	}
	for _, sub := range a.Agents {
		if err := sub.Validate(); err != nil {
			return fmt.Errorf("agent %q: %w", a.Name, err)
		}
	}
	return nil
}

// validateIdentity checks the name and the closed enums (strategy, reasoningEffort) that the
// shared agent-schema.json declares, turning a mid-run server rejection into an error at
// definition time. Open ranges — timeouts, temperature bounds, model names — stay the
// server's, since they can change without an SDK release; Python checks only max_turns. Do
// not add client-side range checks without the same change in the Python and Java SDKs.
func (a *Agent) validateIdentity() error {
	if a.Name == "" {
		return fmt.Errorf("agent name must be a non-empty string")
	}
	if !validName.MatchString(a.Name) {
		return fmt.Errorf(
			"invalid agent name %q: must start with a letter or underscore and "+
				"contain only letters, digits, underscores and hyphens", a.Name)
	}
	if a.Strategy != "" {
		if _, ok := validStrategies[a.Strategy]; !ok {
			return fmt.Errorf("invalid strategy %q for agent %q", a.Strategy, a.Name)
		}
	}
	if a.ReasoningEffort != "" {
		if _, ok := validReasoningEfforts[a.ReasoningEffort]; !ok {
			return fmt.Errorf(
				"invalid reasoningEffort %q for agent %q", a.ReasoningEffort, a.Name)
		}
	}
	// Python rejects max_turns < 1; here 0 must mean "unset" so it can default to
	// 25, and rejecting only negatives keeps this a subset of Python's rule.
	if a.MaxTurns < 0 {
		return fmt.Errorf("agent %q: maxTurns must be >= 0, got %d", a.Name, a.MaxTurns)
	}
	return nil
}

// validateInstructions enforces Instructions xor InstructionsTemplate.
func (a *Agent) validateInstructions() error {
	if a.Instructions != "" && a.InstructionsTemplate != nil {
		return fmt.Errorf(
			"agent %q: set either Instructions or InstructionsTemplate, not both", a.Name)
	}
	if a.InstructionsTemplate != nil && a.InstructionsTemplate.Name == "" {
		return fmt.Errorf("agent %q: InstructionsTemplate.Name is required", a.Name)
	}
	return nil
}

// validateRouting checks the router slots: mutually exclusive, and required by StrategyRouter.
func (a *Agent) validateRouting() error {
	if a.Router != nil && a.RouterFunc != nil {
		return fmt.Errorf(
			"agent %q: set either Router or RouterFunc, not both", a.Name)
	}
	if a.Strategy == StrategyRouter && a.Router == nil && a.RouterFunc == nil {
		return fmt.Errorf(
			"agent %q: StrategyRouter requires Router or RouterFunc", a.Name)
	}
	if a.Router != nil {
		if err := a.Router.Validate(); err != nil {
			return fmt.Errorf("agent %q router: %w", a.Name, err)
		}
	}
	return nil
}

// validateComposition checks guardrails, handoffs, planner, fallback and gate.
func (a *Agent) validateComposition() error {
	if err := validateGuardrails("agent "+a.Name, a.Guardrails); err != nil {
		return err
	}
	if err := a.validateHandoffs(); err != nil {
		return err
	}
	if a.Planner != nil {
		if err := a.Planner.Validate(); err != nil {
			return fmt.Errorf("agent %q planner: %w", a.Name, err)
		}
	}
	if a.Fallback != nil {
		if err := a.Fallback.Validate(); err != nil {
			return fmt.Errorf("agent %q fallback: %w", a.Name, err)
		}
	}
	// The planner prompt exists only under plan-execute, so Python rejects PlannerContext elsewhere.
	if len(a.PlannerContext) > 0 && a.Strategy != StrategyPlanExecute {
		return fmt.Errorf("agent %q: PlannerContext requires StrategyPlanExecute", a.Name)
	}
	for i, c := range a.PlannerContext {
		if err := c.validate(); err != nil {
			return fmt.Errorf("agent %q: PlannerContext[%d]: %w", a.Name, i, err)
		}
	}
	if a.Gate != nil {
		if err := a.Gate.validateGate(); err != nil {
			return fmt.Errorf("agent %q: %w", a.Name, err)
		}
	}
	return nil
}

// maxTurnsOrDefault reports the value actually sent on the wire.
func (a *Agent) maxTurnsOrDefault() int {
	if a.MaxTurns == 0 {
		return defaultMaxTurns
	}
	return a.MaxTurns
}

// strategyOrDefault reports the strategy actually sent on the wire.
func (a *Agent) strategyOrDefault() Strategy {
	if a.Strategy == "" {
		return StrategyHandoff
	}
	return a.Strategy
}

// hasSubAgents reports whether the agent declares sub-agents in any slot, including the
// plan-execute slots held in named fields rather than in the list. Strategy is emitted
// only when this is true, matching the Python serializer.
func (a *Agent) hasSubAgents() bool {
	return len(a.Agents) > 0 || a.Planner != nil || a.Fallback != nil
}
