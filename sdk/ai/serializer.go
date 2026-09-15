//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

// toConfig serializes the agent tree into the agentConfig document the server
// compiles.
//
// This is the cross-SDK wire contract: the same agent must produce the same
// document here, in the Python SDK and in the Java SDK. Two rules govern it,
// and both are load-bearing:
//
//  1. A field is emitted only when the Python serializer emits it. Sending a
//     field Python omits is as wrong as omitting one Python sends.
//  2. Callables are not sent. They are registered as Conductor workers and
//     referenced by a derived task name, e.g. "<agent>_stop_when".
//
// The golden-file tests in serializer_golden_test.go hold this honest against
// documents captured from the Python SDK.
func (a *Agent) toConfig() map[string]any {
	cfg := map[string]any{
		"name":           a.Name,
		"maxTurns":       a.maxTurnsOrDefault(),
		"timeoutSeconds": a.TimeoutSeconds,
		"external":       a.External,
	}
	// Empty model is omitted rather than sent as "", so a sub-agent can
	// inherit its parent's model at compile time.
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
	a.addComposition(cfg)
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

// addLLMKnobs emits the optional model parameters.
//
// Optional scalars are emitted whenever they are set, including when set to
// zero. Python's guard is `is not None`, so temperature=0.0 and maxTokens=0
// both reach the server; a plain Go float64 or int could not express that,
// which is why these fields are pointers.
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

// addMemory emits conversation memory. Its inner fields use emptiness rather
// than nil-ness, matching Python's truthiness checks: MaxMessages of 0 is
// omitted, not sent.
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

// addTools emits the tool list and the required-tool names.
func (a *Agent) addTools(cfg map[string]any) {
	if len(a.Tools) > 0 {
		tools := make([]any, 0, len(a.Tools))
		for _, t := range a.Tools {
			tools = append(tools, t.toolConfig())
		}
		cfg["tools"] = tools
	}
	if len(a.RequiredTools) > 0 {
		cfg["requiredTools"] = a.RequiredTools
	}
}

// addComposition emits the router, loop control, and the sub-agent tree.
func (a *Agent) addComposition(cfg map[string]any) {
	// A router is either a nested agent the server runs, or a reference to a
	// worker. Both land on the same "router" key, so the two forms are
	// mutually exclusive and Validate rejects setting both.
	switch {
	case a.Router != nil:
		cfg["router"] = a.Router.toConfig()
	case a.RouterFunc != nil:
		cfg["router"] = workerRef(a.workerTaskName(routerSuffix))
	}
	if a.Termination != nil {
		cfg["termination"] = a.Termination.terminationConfig()
	}
	// Callables are never serialized. They are registered as workers and
	// referenced by a derived task name, which is why the agent's name is
	// part of the wire output here.
	if a.StopWhen != nil {
		cfg["stopWhen"] = workerRef(a.workerTaskName(stopWhenSuffix))
	}
	// Strategy rides on the presence of sub-agents, not on the field itself:
	// a leaf agent sends no strategy even though Strategy defaults to handoff.
	if a.hasSubAgents() {
		cfg["strategy"] = string(a.strategyOrDefault())
		subs := make([]any, 0, len(a.Agents))
		for _, sub := range a.Agents {
			subs = append(subs, sub.toConfig())
		}
		cfg["agents"] = subs
	}
}
