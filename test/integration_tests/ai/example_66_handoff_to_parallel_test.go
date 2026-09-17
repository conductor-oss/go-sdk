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
)

// Handoff to parallel — the Python SDK's examples/agents/66_handoff_to_parallel.py
// as a test.
//
// A coordinator hands a deep-analysis request to a parallel group of two
// analysts and a quick question to a single agent. This is that flow, copied,
// with the prints replaced by validation: each run completes and the
// coordinator's final answer is the recorded one, which requires the handoff
// to the right target, the group's fan-out and fan-in, and the synthesis to
// have happened as recorded.
func TestExample66HandoffToParallel(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "66_handoff_to_parallel")
	if len(recorded) != 9 {
		t.Fatalf("expected nine recorded model calls, found %d", len(recorded))
	}

	quickCheck := &ai.Agent{
		Name:         "quick_check",
		Model:        mockModel,
		Instructions: "You provide quick, 1-sentence assessments. Be brief and direct.",
	}
	marketAnalyst := &ai.Agent{
		Name:  "market_analyst_66",
		Model: mockModel,
		Instructions: "You are a market analyst. Analyze the market opportunity: " +
			"size, growth rate, key players. 3-4 bullet points.",
	}
	riskAnalyst := &ai.Agent{
		Name:  "risk_analyst_66",
		Model: mockModel,
		Instructions: "You are a risk analyst. Identify the top 3 risks: " +
			"regulatory, technical, and competitive. 3-4 bullet points.",
	}
	deepAnalysis := &ai.Agent{
		Name:     "deep_analysis",
		Model:    mockModel,
		Agents:   []*ai.Agent{marketAnalyst, riskAnalyst},
		Strategy: ai.StrategyParallel,
	}
	coordinator := &ai.Agent{
		Name:  "coordinator_66",
		Model: mockModel,
		Instructions: "You are a business strategist. Route requests to the right team:\n" +
			"- quick_check for simple yes/no questions or quick assessments\n" +
			"- deep_analysis for comprehensive analysis requiring multiple perspectives",
		Agents:   []*ai.Agent{quickCheck, deepAnalysis},
		Strategy: ai.StrategyHandoff,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Minute)
	defer cancel()

	// Scenario 1: deep analysis (handoff to the parallel group).
	result, err := runtime.Run(ctx, coordinator,
		"Provide a deep analysis of entering the AI healthcare market.")
	if err != nil {
		t.Fatalf("scenario 1 failed: %v", err)
	}
	if result.Status != ai.StatusCompleted {
		t.Fatalf("scenario 1 status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	// The coordinator's synthesis of the two analyses is the fifth recorded call.
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[4]); got != want {
		t.Errorf("scenario 1 output is not the recorded answer\n--- got ---\n%.400s\n--- recorded ---\n%.400s", got, want)
	}

	// Scenario 2: quick check (handoff to a single agent).
	result2, err := runtime.Run(ctx, coordinator, "Is the mobile app market still growing?")
	if err != nil {
		t.Fatalf("scenario 2 failed: %v", err)
	}
	if result2.Status != ai.StatusCompleted {
		t.Fatalf("scenario 2 status = %q, want %q (error=%q)", result2.Status, ai.StatusCompleted, result2.Error)
	}
	// The coordinator's answer after the quick check is the last recorded call.
	if got, want := strings.TrimSpace(result2.Output), strings.TrimSpace(recorded[8]); got != want {
		t.Errorf("scenario 2 output is not the recorded answer\n--- got ---\n%.400s\n--- recorded ---\n%.400s", got, want)
	}
}
