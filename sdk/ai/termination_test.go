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
	"reflect"
	"testing"
)

// An empty StopMessage means the default, not an empty stop token.
//
// Python carries "TERMINATE" as a constructor default, so Go has to substitute
// it at serialization for the same reason maxTurns does: the zero value cannot
// be told apart from unset. Sending "" would produce a stop token that never
// matches, silently disabling the condition.
func TestStopMessageDefaultSubstituted(t *testing.T) {
	got := StopMessageTermination{}.terminationConfig()
	if got["stopMessage"] != defaultStopMessage {
		t.Errorf("empty StopMessage must serialize as %q, got %v",
			defaultStopMessage, got["stopMessage"])
	}

	explicit := StopMessageTermination{StopMessage: "HALT"}.terminationConfig()
	if explicit["stopMessage"] != "HALT" {
		t.Errorf("explicit StopMessage must be preserved, got %v", explicit["stopMessage"])
	}
}

// Unset token limits are omitted rather than sent as zero, which the server
// would read as a budget of nothing.
func TestTokenUsageOmitsUnsetLimits(t *testing.T) {
	got := TokenUsageTermination{MaxTotalTokens: Ptr(8000)}.terminationConfig()
	want := map[string]any{"type": "token_usage", "maxTotalTokens": 8000}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	// A limit of zero is a real request and must survive.
	zero := TokenUsageTermination{MaxPromptTokens: Ptr(0)}.terminationConfig()
	if v, ok := zero["maxPromptTokens"]; !ok || v != 0 {
		t.Errorf("a zero limit must be sent, got %#v", zero["maxPromptTokens"])
	}
}

// The composite types are unexported, so a caller cannot build one with a nil
// conditions slice; the constructors are the only way in.
func TestCompositesNest(t *testing.T) {
	got := AndTermination(
		OrTermination(
			TextMentionTermination{Text: "DONE", CaseSensitive: true},
			MaxMessageTermination{MaxMessages: 20},
		),
		TokenUsageTermination{MaxTotalTokens: Ptr(8000)},
	).terminationConfig()

	want := map[string]any{
		"type": "and",
		"conditions": []any{
			map[string]any{
				"type": "or",
				"conditions": []any{
					map[string]any{"type": "text_mention", "text": "DONE", "caseSensitive": true},
					map[string]any{"type": "max_message", "maxMessages": 20},
				},
			},
			map[string]any{"type": "token_usage", "maxTotalTokens": 8000},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v\nwant %v", got, want)
	}
}

// Go functions never reach the wire. StopWhen becomes a task name the runtime
// must later register a worker under, so the derived name is a contract and
// not a display string.
func TestStopWhenSerializesAsDerivedTaskName(t *testing.T) {
	cfg := (&Agent{
		Name:     "terminating",
		Model:    testModel,
		StopWhen: func(context.Context, StopWhenState) (bool, error) { return false, nil },
	}).toConfig()

	want := map[string]any{"taskName": "terminating_stop_when"}
	if !reflect.DeepEqual(cfg["stopWhen"], want) {
		t.Errorf("stopWhen = %v, want %v", cfg["stopWhen"], want)
	}

	unset := (&Agent{Name: "terminating", Model: testModel}).toConfig()
	if _, ok := unset["stopWhen"]; ok {
		t.Errorf("a nil StopWhen must be omitted, got %v", unset["stopWhen"])
	}
}

// Termination and StopWhen are independent: one is evaluated server-side, the
// other runs as a worker, and setting both is legitimate.
func TestTerminationAndStopWhenCoexist(t *testing.T) {
	cfg := (&Agent{
		Name:        "both",
		Model:       testModel,
		Termination: MaxMessageTermination{MaxMessages: 5},
		StopWhen:    func(context.Context, StopWhenState) (bool, error) { return false, nil },
	}).toConfig()

	if _, ok := cfg["termination"]; !ok {
		t.Error("termination must be sent alongside stopWhen")
	}
	if _, ok := cfg["stopWhen"]; !ok {
		t.Error("stopWhen must be sent alongside termination")
	}
}
