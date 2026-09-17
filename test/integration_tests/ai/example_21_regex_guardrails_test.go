//go:build integration

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package integration

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type userProfileIn struct {
	UserID string `json:"user_id"`
}

// Regex guardrails — the Python SDK's examples/agents/21_regex_guardrails.py
// as a test.
//
// Two server-side regex guardrails block email addresses (retrying the
// model) and SSNs (failing the run). Scenario 1 asks for a profile holding
// both; scenario 2 asks a clean question. This is that flow, copied, with the
// prints replaced by validation: both runs complete with their recorded
// answers, and the first answer carries neither the email nor the SSN.
func TestExample21RegexGuardrails(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "21_regex_guardrails")
	if len(recorded) != 4 {
		t.Fatalf("expected four recorded calls (three for scenario 1, one for scenario 2), found %d", len(recorded))
	}

	getUserProfile := func(ctx context.Context, in userProfileIn) (map[string]any, error) {
		return map[string]any{
			"name":       "Alice Johnson",
			"email":      "alice.johnson@example.com",
			"ssn":        "123-45-6789",
			"department": "Engineering",
			"role":       "Senior Developer",
		}, nil
	}

	noEmails := &ai.RegexGuardrail{
		Patterns: []string{`[\w.+-]+@[\w-]+\.[\w.-]+`},
		Mode:     "block",
		Message:  "Response must not contain email addresses. Redact them.",
	}
	noEmails.Name = "no_email_addresses"
	noEmails.Position = ai.PositionOutput
	noEmails.OnFail = ai.OnFailRetry
	noSSN := &ai.RegexGuardrail{
		Patterns: []string{`\b\d{3}-\d{2}-\d{4}\b`},
		Mode:     "block",
		Message:  "Response must not contain Social Security Numbers.",
	}
	noSSN.Name = "no_ssn"
	noSSN.Position = ai.PositionOutput
	noSSN.OnFail = ai.OnFailRaise

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	// Scenario 1: the profile holds PII; the guardrails make the model redact it.
	agent := &ai.Agent{
		Name:  "hr_assistant",
		Model: mockModel,
		Tools: []ai.ToolDef{tool.Func("get_user_profile", "Retrieve a user's profile from the database.", getUserProfile)},
		Instructions: "You are an HR assistant. When asked about employees, look up their " +
			"profile and share ALL the details you find.",
		Guardrails: []ai.Guardrail{noEmails, noSSN},
	}
	result, err := runtime.Run(ctx, agent, "Tell me everything about user U-001.")
	if err != nil {
		t.Fatalf("scenario 1 run failed: %v", err)
	}
	if result.Status != ai.StatusCompleted {
		t.Fatalf("scenario 1 status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	if strings.Contains(result.Output, "alice.johnson@example.com") {
		t.Error("scenario 1: email leaked past the regex guardrail")
	}
	if strings.Contains(result.Output, "123-45-6789") {
		t.Error("scenario 1: SSN leaked past the regex guardrail")
	}
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[2]); got != want {
		t.Errorf("scenario 1 output is not the recorded answer\n--- got ---\n%.400s\n--- recorded ---\n%.400s", got, want)
	}

	// Scenario 2: nothing to block; the guardrails pass.
	cleanAgent := &ai.Agent{
		Name:         "dept_assistant",
		Model:        mockModel,
		Instructions: "You are an HR assistant. Answer questions about departments.",
		Guardrails:   []ai.Guardrail{noEmails, noSSN},
	}
	result2, err := runtime.Run(ctx, cleanAgent, "What departments exist at the company?")
	if err != nil {
		t.Fatalf("scenario 2 run failed: %v", err)
	}
	if result2.Status != ai.StatusCompleted {
		t.Fatalf("scenario 2 status = %q, want %q (error=%q)", result2.Status, ai.StatusCompleted, result2.Error)
	}
	if got, want := strings.TrimSpace(result2.Output), strings.TrimSpace(recorded[3]); got != want {
		t.Errorf("scenario 2 output is not the recorded answer\n--- got ---\n%.300s\n--- recorded ---\n%.300s", got, want)
	}
}
