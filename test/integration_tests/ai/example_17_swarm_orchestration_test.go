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

// Swarm orchestration — the Python SDK's examples/agents/17_swarm_orchestration.py
// as a test.
//
// A support agent transfers a damaged-goods complaint to the refund
// specialist, who processes it and hands back. This is that flow, copied,
// with the print replaced by validation: the run completes and the support
// agent's final answer is the recorded one, which requires the transfer to
// the specialist and the hand-back to have happened as recorded.
func TestExample17SwarmOrchestration(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "17_swarm_orchestration")

	refundAgent := &ai.Agent{
		Name:  "refund_specialist",
		Model: mockModel,
		Instructions: "You are a refund specialist. Process the customer's refund request. " +
			"Check eligibility, confirm the refund amount, and let them know the " +
			"timeline. Be empathetic and clear. Do NOT ask follow-up questions — " +
			"just process the refund based on what the customer told you.",
	}
	techAgent := &ai.Agent{
		Name:  "tech_support",
		Model: mockModel,
		Instructions: "You are a technical support specialist. Diagnose the customer's " +
			"technical issue and provide clear troubleshooting steps.",
	}
	support := &ai.Agent{
		Name:  "support",
		Model: mockModel,
		Instructions: "You are the front-line customer support agent. Triage customer requests. " +
			"If the customer needs a refund, transfer to the refund specialist. " +
			"If they have a technical issue, transfer to tech support. " +
			"Use the transfer tools available to you to hand off the conversation.",
		Agents:   []*ai.Agent{refundAgent, techAgent},
		Strategy: ai.StrategySwarm,
		Handoffs: []ai.HandoffCondition{
			&ai.OnTextMention{Text: "refund", Target: "refund_specialist"},
			&ai.OnTextMention{Text: "technical", Target: "tech_support"},
		},
		MaxTurns: 3,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	result, err := runtime.Run(ctx, support,
		"I bought a product last week and it arrived damaged. I want my money back.")
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}

	// Validation, in place of the example's result.print_result().
	if result.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[len(recorded)-1]); got != want {
		t.Errorf("output is not the recorded answer\n--- got ---\n%.300s\n--- recorded ---\n%.300s", got, want)
	}
}
