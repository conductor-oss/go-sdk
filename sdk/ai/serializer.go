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
	"reflect"

	"github.com/conductor-sdk/conductor-go/sdk/ai/internal/schema"
)

// toConfig serializes the agent tree into the agentConfig document the server
// compiles. It is the cross-SDK wire contract, identical in the Python and Java
// SDKs: a field is emitted only when the Python serializer emits it, and
// callables are never sent but registered as Conductor workers referenced by a
// derived task name, e.g. "<agent>_stop_when". The golden-file tests in
// serializer_golden_test.go check it against documents captured from Python.
func (a *Agent) toConfig() map[string]any {
	// A skill's document is the raw skill directory, marked with _framework so
	// the server normalizes it; golden fixture 18_skill pins the shape.
	if a.skill != nil {
		return a.skill.wireConfig(a.Name)
	}
	cfg := map[string]any{
		"name":           a.Name,
		"maxTurns":       a.maxTurnsOrDefault(),
		"timeoutSeconds": a.TimeoutSeconds,
		"external":       a.External,
	}
	// Empty model is omitted rather than sent as "", so a sub-agent inherits
	// its parent's model at compile time.
	if a.Model != "" {
		cfg["model"] = a.Model
	}
	if a.BaseURL != "" {
		cfg["baseUrl"] = a.BaseURL
	}
	a.addInstructions(cfg)
	a.addLLMKnobs(cfg)
	a.addMemory(cfg)
	a.addTools(cfg)
	a.addDefinition(cfg)
	a.addComposition(cfg)
	a.addSubAgents(cfg)
	return cfg
}

// addInstructions emits either the literal instructions or the template
// reference; Validate guarantees at most one is set.
func (a *Agent) addInstructions(cfg map[string]any) {
	switch {
	case a.InstructionsTemplate != nil:
		tpl := map[string]any{
			"type": "prompt_template",
			"name": a.InstructionsTemplate.Name,
		}
		if len(a.InstructionsTemplate.Variables) > 0 {
			tpl["variables"] = a.InstructionsTemplate.Variables
		}
		if a.InstructionsTemplate.Version != nil {
			tpl["version"] = *a.InstructionsTemplate.Version
		}
		cfg["instructions"] = tpl
	case a.Instructions != "":
		cfg["instructions"] = a.Instructions
	}
}

// addLLMKnobs emits the optional model parameters. Python's guard is
// `is not None`, so temperature=0.0 and maxTokens=0 must both reach the
// server, which is why these fields are pointers.
func (a *Agent) addLLMKnobs(cfg map[string]any) {
	if a.MaxTokens != nil {
		cfg["maxTokens"] = *a.MaxTokens
	}
	if a.Temperature != nil {
		cfg["temperature"] = *a.Temperature
	}
	if a.ContextWindowBudget != nil {
		cfg["contextWindowBudget"] = *a.ContextWindowBudget
	}
	if a.ReasoningEffort != "" {
		cfg["reasoningEffort"] = string(a.ReasoningEffort)
	}
	if a.IncludeContents != "" {
		cfg["includeContents"] = a.IncludeContents
	}
	// A thinking budget is not sent as a bare number: it expands into a
	// thinkingConfig object, and setting it is what enables thinking.
	if a.ThinkingBudgetTokens != nil {
		cfg["thinkingConfig"] = map[string]any{
			"enabled":      true,
			"budgetTokens": *a.ThinkingBudgetTokens,
		}
	}
}

// addMemory emits conversation memory. Its inner fields use emptiness,
// matching Python's truthiness checks: MaxMessages of 0 is omitted.
func (a *Agent) addMemory(cfg map[string]any) {
	if a.Memory == nil {
		return
	}
	mem := map[string]any{}
	if len(a.Memory.Messages) > 0 {
		mem["messages"] = a.Memory.Messages
	}
	if a.Memory.MaxMessages != 0 {
		mem["maxMessages"] = a.Memory.MaxMessages
	}
	cfg["memory"] = mem
}

func (a *Agent) addTools(cfg map[string]any) {
	declared := a.derivedTools()
	if len(a.Tools) > 0 || len(declared) > 0 {
		tools := make([]any, 0, len(a.Tools)+len(declared))
		for _, t := range a.Tools {
			tools = append(tools, t.toolConfig(a.Stateful))
		}
		for _, t := range declared {
			tools = append(tools, t.toolConfig(a.Stateful))
		}
		cfg["tools"] = tools
	}
	if len(a.RequiredTools) > 0 {
		cfg["requiredTools"] = a.RequiredTools
	}
}

func (a *Agent) addDefinition(cfg map[string]any) {
	if a.OutputType != nil {
		t := reflect.TypeOf(a.OutputType)
		for t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		cfg["outputType"] = map[string]any{
			"schema":    schema.Of(t),
			"className": t.Name(),
		}
	}

	if a.CodeExecution != nil {
		cfg["codeExecution"] = a.CodeExecution.config()
	}
	if a.CLI != nil {
		cfg["cliConfig"] = a.CLI.config()
	}
	if len(a.Credentials) > 0 {
		cfg["credentials"] = a.Credentials
	}
	if len(a.MaskedFields) > 0 {
		cfg["maskedFields"] = a.MaskedFields
	}
	if a.Introduction != "" {
		cfg["introduction"] = a.Introduction
	}
	if len(a.Metadata) > 0 {
		cfg["metadata"] = a.Metadata
	}
}

func (a *Agent) addComposition(cfg map[string]any) {
	// A router is either a nested agent the server runs or a worker reference,
	// both on the same "router" key, so Validate rejects setting both.
	switch {
	case a.Router != nil:
		cfg["router"] = a.Router.toConfig()
	case a.RouterFunc != nil:
		cfg["router"] = workerRef(a.workerTaskName(routerSuffix))
	}
	if a.Termination != nil {
		cfg["termination"] = a.Termination.terminationConfig()
	}
	if a.StopWhen != nil {
		cfg["stopWhen"] = workerRef(a.workerTaskName(stopWhenSuffix))
	}
	if len(a.Guardrails) > 0 {
		gs := make([]any, 0, len(a.Guardrails))
		for _, g := range a.Guardrails {
			gs = append(gs, g.guardrailConfig())
		}
		cfg["guardrails"] = gs
	}
	if cbs := a.callbackConfigs(); len(cbs) > 0 {
		cfg["callbacks"] = cbs
	}
	if a.Gate != nil {
		cfg["gate"] = a.Gate.gateConfig(a.Name)
	}
	a.addPlanning(cfg)
}

func (a *Agent) addPlanning(cfg map[string]any) {
	if a.EnablePlanning {
		cfg["enablePlanning"] = true
	}
	// Both slots are nested agent documents, so a planner may itself have
	// tools or sub-agents.
	if a.Planner != nil {
		cfg["planner"] = a.Planner.toConfig()
	}
	if a.Fallback != nil {
		cfg["fallback"] = a.Fallback.toConfig()
	}
	if a.FallbackMaxTurns > 0 {
		cfg["fallbackMaxTurns"] = a.FallbackMaxTurns
	}
	if len(a.PrefillTools) > 0 {
		calls := make([]any, 0, len(a.PrefillTools))
		for _, p := range a.PrefillTools {
			calls = append(calls, p.config())
		}
		cfg["prefillTools"] = calls
	}
	// Python emits planSource whenever it is not nil, an empty map included.
	if a.PlanSource != nil {
		cfg["planSource"] = a.PlanSource
	}
	if len(a.PlannerContext) > 0 {
		entries := make([]any, 0, len(a.PlannerContext))
		for _, c := range a.PlannerContext {
			entries = append(entries, c.config())
		}
		cfg["plannerContext"] = entries
	}
	// Synthesis is on by default; only the opt-out travels.
	if a.Synthesize != nil && !*a.Synthesize {
		cfg["synthesize"] = false
	}
}

func (a *Agent) addSubAgents(cfg map[string]any) {
	if len(a.Handoffs) > 0 {
		hs := make([]any, 0, len(a.Handoffs))
		for _, h := range a.Handoffs {
			hs = append(hs, h.handoffConfig(a.Name))
		}
		cfg["handoffs"] = hs
	}
	if len(a.AllowedTransitions) > 0 {
		cfg["allowedTransitions"] = a.AllowedTransitions
	}
	// Strategy rides on the presence of sub-agents: a leaf agent sends none
	// even though Strategy defaults to handoff.
	if a.hasSubAgents() {
		cfg["strategy"] = string(a.strategyOrDefault())
	}
	if len(a.Agents) > 0 {
		subs := make([]any, 0, len(a.Agents))
		for _, sub := range a.Agents {
			subs = append(subs, sub.toConfig())
		}
		cfg["agents"] = subs
	}
}
