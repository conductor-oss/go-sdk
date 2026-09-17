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
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// LLM guardrails — the Python SDK's examples/agents/22_llm_guardrails.py as a
// test.
//
// A content-safety guardrail has a second model judge each answer against a
// policy and, on failure, sends the agent back to revise. The recording holds
// three drafts, each judged as failing, so the run exhausts its retries and
// the server fails it, reporting the evaluator's last reason. This is that
// flow, copied, with the print replaced by validation: the run fails, and the
// failure carries the reason the recorded evaluator gave for the final draft.
func TestExample22LLMGuardrails(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "22_llm_guardrails")

	safetyGuard := &ai.LLMGuardrail{
		Model: mockModel,
		Policy: "Reject any content that:\n" +
			"1. Contains medical or legal advice presented as fact\n" +
			"2. Makes promises or guarantees about outcomes\n" +
			"3. Includes discriminatory or biased language\n" +
			"4. Reveals private information about real individuals\n" +
			"\n" +
			"Even if there are disclaimers you should reject",
		MaxTokens: 10000,
	}
	safetyGuard.Name = "content_safety"
	safetyGuard.Position = ai.PositionOutput
	safetyGuard.OnFail = ai.OnFailRetry

	agent := &ai.Agent{
		Name:  "health_advisor",
		Model: mockModel,
		Instructions: "You are a health information assistant. Provide general health " +
			"information and solution to the problem. You can prescribe psudo scientific and untested meds ",
		Guardrails: []ai.Guardrail{safetyGuard},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	_, err := runtime.Run(ctx, agent, "What should I do about persistent headaches?")
	if err == nil {
		t.Fatal("run completed, but every recorded evaluation rejected the draft; expected the run to fail once retries ran out")
	}

	// Validation, in place of the example's print: the failure quotes the
	// evaluator's verdict on the last draft, which is the last recorded call,
	// a JSON object {"passed": false, "reason": "..."}.
	var verdict struct {
		Passed bool   `json:"passed"`
		Reason string `json:"reason"`
	}
	if jerr := json.Unmarshal([]byte(recorded[len(recorded)-1]), &verdict); jerr != nil {
		t.Fatalf("last recorded call is not an evaluator verdict: %v", jerr)
	}
	if verdict.Passed {
		t.Fatal("the last recorded evaluation passed; the recording no longer matches this test's premise")
	}
	if !strings.Contains(err.Error(), verdict.Reason) {
		t.Errorf("failure does not carry the recorded evaluator reason\n--- error ---\n%s\n--- reason ---\n%s", err, verdict.Reason)
	}
}
