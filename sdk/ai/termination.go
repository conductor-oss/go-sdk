//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

import "fmt"

// TerminationCondition decides when an agent's loop should stop. Conditions
// compose into a tree with AndTermination and OrTermination, which the server
// evaluates after each turn. The interface is sealed: the server knows a fixed
// set of condition types, so a caller-defined one could not be serialized. Use
// StopWhen for arbitrary Go logic, which runs as a worker instead.
type TerminationCondition interface {
	terminationConfig() map[string]any
	// validateTermination mirrors Python's constructor checks, which a Go struct
	// literal cannot run; Agent.Validate walks the tree.
	validateTermination() error
}

// TextMentionTermination stops when the given text appears in the output.
type TextMentionTermination struct {
	Text          string
	CaseSensitive bool
}

func (t TextMentionTermination) validateTermination() error { return nil }

func (t TextMentionTermination) terminationConfig() map[string]any {
	return map[string]any{
		"type":          "text_mention",
		"text":          t.Text,
		"caseSensitive": t.CaseSensitive,
	}
}

// defaultStopMessage matches the Python SDK's StopMessageTermination default. An
// empty StopMessage is indistinguishable from unset, so it is substituted at
// serialization instead.
const defaultStopMessage = "TERMINATE"

// StopMessageTermination stops when the model emits exactly this message.
type StopMessageTermination struct {
	// StopMessage is the stop token. Empty means the default, "TERMINATE".
	StopMessage string
}

func (t StopMessageTermination) validateTermination() error { return nil }

func (t StopMessageTermination) terminationConfig() map[string]any {
	msg := t.StopMessage
	if msg == "" {
		msg = defaultStopMessage
	}
	return map[string]any{
		"type":        "stop_message",
		"stopMessage": msg,
	}
}

// MaxMessageTermination stops after the conversation reaches this many messages.
type MaxMessageTermination struct {
	MaxMessages int
}

// validateTermination rejects a limit below 1, matching Python; Go needs it
// because MaxMessageTermination{} leaves MaxMessages at 0.
func (t MaxMessageTermination) validateTermination() error {
	if t.MaxMessages < 1 {
		return fmt.Errorf("MaxMessageTermination: maxMessages must be >= 1, got %d", t.MaxMessages)
	}
	return nil
}

func (t MaxMessageTermination) terminationConfig() map[string]any {
	return map[string]any{
		"type":        "max_message",
		"maxMessages": t.MaxMessages,
	}
}

// TokenUsageTermination stops once a token budget is exhausted. The three limits
// are independent and optional, hence pointers: an unset one is omitted, not
// sent as a zero the server would read as "no tokens allowed".
type TokenUsageTermination struct {
	MaxTotalTokens      *int
	MaxPromptTokens     *int
	MaxCompletionTokens *int
}

// validateTermination rejects a condition with no limits, matching Python:
// a token_usage with no budget can never fire.
func (t TokenUsageTermination) validateTermination() error {
	if t.MaxTotalTokens == nil && t.MaxPromptTokens == nil && t.MaxCompletionTokens == nil {
		return fmt.Errorf("TokenUsageTermination: at least one token limit must be set")
	}
	return nil
}

func (t TokenUsageTermination) terminationConfig() map[string]any {
	cfg := map[string]any{"type": "token_usage"}
	if t.MaxTotalTokens != nil {
		cfg["maxTotalTokens"] = *t.MaxTotalTokens
	}
	if t.MaxPromptTokens != nil {
		cfg["maxPromptTokens"] = *t.MaxPromptTokens
	}
	if t.MaxCompletionTokens != nil {
		cfg["maxCompletionTokens"] = *t.MaxCompletionTokens
	}
	return cfg
}

// andTermination and orTermination are unexported so composites come only from
// the constructors below, never with a nil conditions slice.
type andTermination struct{ conditions []TerminationCondition }

type orTermination struct{ conditions []TerminationCondition }

func (t andTermination) validateTermination() error {
	return validateConditions("AndTermination", t.conditions)
}

func (t orTermination) validateTermination() error {
	return validateConditions("OrTermination", t.conditions)
}

func (t andTermination) terminationConfig() map[string]any {
	return map[string]any{"type": "and", "conditions": conditionConfigs(t.conditions)}
}

func (t orTermination) terminationConfig() map[string]any {
	return map[string]any{"type": "or", "conditions": conditionConfigs(t.conditions)}
}

// AndTermination stops only when every condition holds. Python spells this with
// the & operator, which Go cannot overload, so composites take variadic args.
func AndTermination(conditions ...TerminationCondition) TerminationCondition {
	return andTermination{conditions: conditions}
}

// OrTermination stops as soon as any condition holds.
func OrTermination(conditions ...TerminationCondition) TerminationCondition {
	return orTermination{conditions: conditions}
}

// validateConditions rejects an empty composite, which could never fire.
func validateConditions(kind string, conditions []TerminationCondition) error {
	if len(conditions) == 0 {
		return fmt.Errorf("%s: at least one condition is required", kind)
	}
	for _, c := range conditions {
		if c == nil {
			return fmt.Errorf("%s: nil condition", kind)
		}
		if err := c.validateTermination(); err != nil {
			return err
		}
	}
	return nil
}

// conditionConfigs serializes a composite's children, skipping nils that would otherwise panic.
func conditionConfigs(conditions []TerminationCondition) []any {
	out := make([]any, 0, len(conditions))
	for _, c := range conditions {
		if c == nil {
			continue
		}
		out = append(out, c.terminationConfig())
	}
	return out
}
