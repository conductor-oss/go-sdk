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

// validName mirrors _VALID_NAME_RE in the Python SDK. The name becomes a
// Conductor workflow name, so it must be a legal identifier.
var validName = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_-]*$`)

// defaultMaxTurns matches the Python and Java SDKs. It is always sent.
const defaultMaxTurns = 25

// PromptTemplate references a prompt template stored on the Conductor server.
//
// The SDK never creates templates; they are managed through the UI, the API,
// or PromptClient. Values in Variables may be Conductor expressions such as
// "${workflow.input.user_tier}".
type PromptTemplate struct {
	Name      string
	Variables map[string]any
	// Version selects a template version. Nil means latest.
	Version *int
}

// ReasoningEffort selects how much reasoning an OpenAI reasoning model does
// before answering. Other models ignore it.
//
// The values come from the "reasoningEffort" enum in agent-schema.json, the
// contract shared by every SDK. Note that it includes "minimal", which the
// Python docstring and the java-sdk field reference both omit.
//
// Validate rejects any other value, as it does for Strategy. A named string
// type alone would not: Go converts untyped constants implicitly, so
// ReasoningEffort: "hgih" compiles. Enforcing the schema's enum here turns a
// server-side rejection mid-run into an error at definition time.
type ReasoningEffort string

const (
	// ReasoningEffortMinimal does the least reasoning before answering.
	ReasoningEffortMinimal ReasoningEffort = "minimal"
	ReasoningEffortLow     ReasoningEffort = "low"
	ReasoningEffortMedium  ReasoningEffort = "medium"
	ReasoningEffortHigh    ReasoningEffort = "high"
)

// validReasoningEfforts mirrors the schema's enum. Kept as a set so Validate
// can reject unknown values the same way it does for Strategy.
var validReasoningEfforts = map[ReasoningEffort]struct{}{
	ReasoningEffortMinimal: {},
	ReasoningEffortLow:     {},
	ReasoningEffortMedium:  {},
	ReasoningEffortHigh:    {},
}

// ConversationMemory seeds an agent with prior messages and bounds how many
// it retains.
//
// Both fields are omitted from the wire when empty, so an entirely empty
// ConversationMemory still sends "memory": {} — non-nil memory means the
// agent is stateful, which is distinct from having no memory at all.
type ConversationMemory struct {
	// Messages are prior turns, each a role/content map.
	Messages []map[string]any
	// MaxMessages caps retained history. Zero means no explicit cap.
	MaxMessages int
}

// Agent is a declarative description of an agent, sent to the server for
// compilation into a Conductor workflow.
//
// It is a configuration document rather than a constructor call: set only the
// fields you need and leave the rest at their zero value, which is what
// "unset" means on the wire.
//
//	agent := &ai.Agent{
//	    Name:         "weather_bot",
//	    Model:        "openai/gpt-4o",
//	    Instructions: "Use tools to answer questions.",
//	}
//
// Sub-agents nest as an ordinary slice literal, so a tree reads as a tree:
//
//	analysis := &ai.Agent{
//	    Name:     "analysis",
//	    Model:    "openai/gpt-4o",
//	    Strategy: ai.StrategyParallel,
//	    Agents: []*ai.Agent{
//	        {Name: "market", Model: m, Instructions: "..."},
//	        {Name: "risk",   Model: m, Instructions: "..."},
//	    },
//	}
//
// Optional numeric settings are pointers because zero is a meaningful value:
// a Temperature of 0 must reach the server, and an unset Temperature must not.
// Use the Ptr helper to set them.
//
// Fields are validated together rather than individually, because most rules
// are cross-field (a router strategy needs a Router, plan-execute needs a
// Planner). Runtime methods call Validate before serializing, so an invalid
// agent cannot reach the server whether or not you call it yourself.
type Agent struct {
	// Name becomes the Conductor workflow name. Required.
	Name string
	// Model is a "provider/model" identifier, e.g. "openai/gpt-4o". An empty
	// value on a sub-agent inherits the parent's model at compile time.
	Model string
	// Instructions is the system prompt. Mutually exclusive with
	// InstructionsTemplate.
	Instructions string
	// InstructionsTemplate uses a named server-side prompt template instead of
	// a literal Instructions string.
	InstructionsTemplate *PromptTemplate

	// Tools the agent may call. Build them with the tool package.
	Tools []ToolDef
	// RequiredTools names tools the run should call before it finishes. Matches
	// Python's required_tools and Java's requiredTools. Not usable on current
	// servers: the compiled check nests a DO_WHILE inside a DO_WHILE, which
	// Conductor does not support, and the run deadlocks.
	RequiredTools []string

	// Planner and Fallback are the named slots StrategyPlanExecute uses. The
	// planner emits a plan; the parent's Tools become the tools that plan may
	// name. Fallback runs if the plan cannot be carried out.
	Planner  *Agent
	Fallback *Agent
	// FallbackMaxTurns bounds the fallback agent. Omitted when zero, which
	// leaves the limit to the server.
	FallbackMaxTurns int
	// EnablePlanning asks the model to plan before acting. It is unrelated to
	// Planner: this is a preamble on a single agent, that is a sub-agent slot.
	EnablePlanning bool

	// Handoffs move control between sub-agents, usually with StrategySwarm.
	Handoffs []HandoffCondition
	// AllowedTransitions restricts which sub-agent may hand off to which, as
	// from-name to permitted target names. Empty means no restriction.
	AllowedTransitions map[string][]string

	// Guardrails check this agent's input or output. Tools carry their own in
	// ToolDef.Guardrails.
	Guardrails []Guardrail

	// OutputType constrains the final answer to a struct's shape. Pass a zero
	// value of the type, as in OutputType: Ticket{}; the schema is derived from
	// its json tags the same way a tool's input schema is.
	//
	// The wire schema leaves the inner document open (additionalProperties is
	// true), so what the SDKs send inside it differs: Python carries Pydantic's
	// per-field titles and Java sends types only, as Go does here.
	OutputType any

	// CodeExecution and CLI give the agent derived tools for running code and
	// shell commands. Both execute server side; see CodeExecutionConfig.
	CodeExecution *CodeExecutionConfig
	CLI           *CLIConfig

	// Credentials are secret names available to every tool on this agent, as a
	// fallback for tools whose own declaration cannot name them. Prefer
	// tool.WithCredentials, which scopes a secret to the one tool that reads it.
	Credentials []string

	// MaskedFields names input and output fields the server redacts from
	// execution history and the UI.
	MaskedFields []string

	// Introduction is the agent's opening message to a user.
	Introduction string

	// Metadata is arbitrary data carried with the agent definition.
	Metadata map[string]any

	// Stateful routes every tool on this agent to a per-execution worker
	// domain, so calls in one run reach the same worker process. It is not a
	// field of its own on the wire: it stamps stateful onto each tool.
	Stateful bool

	// Agents are sub-agents; a non-empty value makes this a multi-agent system
	// and causes Strategy to be sent.
	Agents []*Agent
	// Strategy orchestrates Agents. Defaults to StrategyHandoff.
	Strategy Strategy
	// Router selects one sub-agent using an LLM. Required by StrategyRouter
	// unless RouterFunc is set; the two are mutually exclusive.
	Router *Agent
	// RouterFunc selects one sub-agent with Go code, run as a worker.
	RouterFunc RouterFunc

	// MaxTurns bounds the agent loop. Zero means the default of 25.
	MaxTurns int
	// TimeoutSeconds bounds the whole execution. Zero means no explicit limit.
	TimeoutSeconds int

	// MaxTokens caps tokens generated per LLM call. Nil leaves it to the
	// model's own default.
	MaxTokens *int
	// Temperature is the sampling temperature. Nil is unset; Ptr(0) is a
	// deliberate request for deterministic output and is sent as 0.
	Temperature *float64
	// ReasoningEffort applies to OpenAI reasoning models; others ignore it.
	// Empty means unset.
	ReasoningEffort ReasoningEffort
	// ThinkingBudgetTokens enables extended thinking with the given budget.
	// Serialized as thinkingConfig, not as a bare number.
	ThinkingBudgetTokens *int
	// ContextWindowBudget is the token threshold at which the server starts
	// condensing context proactively.
	ContextWindowBudget *int
	// IncludeContents controls context inheritance. "none" starts the agent
	// with fresh context; empty inherits the parent's.
	IncludeContents string
	// Memory seeds the agent with prior messages and bounds how many it keeps.
	Memory *ConversationMemory

	// Termination stops the loop on a server-evaluated condition tree.
	Termination TerminationCondition
	// StopWhen stops the loop on arbitrary Go logic, run as a worker. It is
	// independent of Termination; both may be set.
	StopWhen StopWhenFunc

	// External marks the agent as served elsewhere; no workers are started
	// for it locally.
	External bool
}

// Ptr returns a pointer to v, for setting optional fields inline:
//
//	Temperature: ai.Ptr(0.0)   // sent as 0
//	Temperature: nil           // not sent
func Ptr[T any](v T) *T { return &v }

// Validate reports the first configuration error in the agent tree.
//
// The rules are cross-field, so no type signature can express them; this is
// where they live. Runtime methods call it before serializing.
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

// validateIdentity checks the name and the closed enums.
//
// Closed enums (strategy, reasoningEffort) are declared in the shared
// agent-schema.json, so checking them enforces the contract rather than
// inventing a rule, and turns a mid-run server rejection into an error at
// definition time. Open ranges (timeouts, temperature bounds, model names)
// are the server's to police, because they can change without an SDK
// release; Python checks only max_turns and we match that. Do not add
// client-side range checks here without making the same change in the
// Python and Java SDKs.
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
	// Python rejects max_turns < 1. Go cannot match that exactly, because 0
	// is the zero value and has to mean "unset" so it can default to 25;
	// rejecting negatives keeps this a subset of Python's rule rather than
	// a stricter one.
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

// validateRouting checks the router slots, which are mutually exclusive and
// required by StrategyRouter.
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

// validateComposition checks the multi-agent slots: guardrails, handoffs,
// and the planner and fallback agents.
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

// hasSubAgents reports whether the agent declares sub-agents in any slot.
// Strategy is only emitted when this is true, matching the Python serializer.
func (a *Agent) hasSubAgents() bool {
	// The plan-execute slots count: they are sub-agents held in named fields
	// rather than in the list, and a strategy without them would be dropped.
	return len(a.Agents) > 0 || a.Planner != nil || a.Fallback != nil
}
