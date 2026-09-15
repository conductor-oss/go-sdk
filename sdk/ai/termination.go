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

// TerminationCondition decides when an agent's loop should stop.
//
// Conditions compose into a tree with AndTermination and OrTermination, which
// the server evaluates after each turn:
//
//	Termination: ai.AndTermination(
//	    ai.OrTermination(
//	        ai.TextMentionTermination{Text: "DONE", CaseSensitive: true},
//	        ai.MaxMessageTermination{MaxMessages: 20},
//	    ),
//	    ai.TokenUsageTermination{MaxTotalTokens: ai.Ptr(8000)},
//	)
//
// The interface is sealed by an unexported method: the server understands a
// fixed set of condition types, so a caller-defined implementation could not
// be serialized. Use StopWhen for arbitrary Go logic instead — that runs as a
// worker rather than as a server-side condition.
type TerminationCondition interface {
	terminationConfig() map[string]any
	// validateTermination mirrors the constructor checks the Python SDK
	// performs, which Go cannot do at construction time because a struct
	// literal has no constructor to run. Agent.Validate walks the tree.
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

// defaultStopMessage matches the Python SDK's StopMessageTermination default.
//
// As with maxTurns, Go cannot carry this as a parameter default: an empty
// StopMessage is the zero value and so is indistinguishable from unset, and
// an empty stop token would never match anything. It is substituted at
// serialization instead, so StopMessageTermination{} behaves like Python's
// StopMessageTermination().
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

// MaxMessageTermination stops after the conversation reaches this many
// messages.
type MaxMessageTermination struct {
	MaxMessages int
}

// validateTermination rejects a limit below 1, matching Python. Go needs the
// check more than Python does: MaxMessageTermination{} is constructible and
// leaves MaxMessages at 0, whereas Python's constructor requires the argument.
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

// TokenUsageTermination stops once a token budget is exhausted.
//
// The three limits are independent and all optional, so they are pointers:
// Python guards each with `is not None` and omits the ones left unset, rather
// than sending a zero the server would read as "no tokens allowed".
type TokenUsageTermination struct {
	MaxTotalTokens      *int
	MaxPromptTokens     *int
	MaxCompletionTokens *int
}

// validateTermination rejects a condition with no limits at all, matching
// Python. Without it the wire carries {"type": "token_usage"} and no budget,
// which can never fire.
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

// andTermination and orTermination are unexported so the only way to build a
// composite is through the constructors below, which keeps a composite from
// being created with a nil Conditions slice.
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

// AndTermination stops only when every condition holds.
//
// Python spells this with the & operator; Go has no operator overloading, so
// composites are built with variadic constructors instead.
func AndTermination(conditions ...TerminationCondition) TerminationCondition {
	return andTermination{conditions: conditions}
}

// OrTermination stops as soon as any condition holds.
func OrTermination(conditions ...TerminationCondition) TerminationCondition {
	return orTermination{conditions: conditions}
}

// validateConditions rejects an empty composite, which would serialize to a
// condition that can never fire, and recurses into the children.
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

// conditionConfigs serializes a composite's children, skipping nils so a
// stray nil in a literal cannot panic during serialization.
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
