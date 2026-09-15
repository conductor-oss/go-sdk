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

// Parallel agents — the Python SDK's examples/agents/07_parallel_agents.py
// as a test.
//
// Three analysts examine the same topic at once. This is that flow, copied,
// with the print replaced by validation: the run completes and the combined
// output carries each analyst's recorded answer, identified by its first
// line, so all three ran and all three results were kept.
func TestExample07ParallelAgents(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "07_parallel_agents")

	marketAnalyst := &ai.Agent{
		Name:  "market_analyst",
		Model: mockModel,
		Instructions: "You are a market analyst. Analyze the given topic from a market perspective: " +
			"market size, growth trends, key players, and opportunities.",
	}
	riskAnalyst := &ai.Agent{
		Name:  "risk_analyst",
		Model: mockModel,
		Instructions: "You are a risk analyst. Analyze the given topic for risks: " +
			"regulatory risks, technical risks, competitive threats, and mitigation strategies.",
	}
	complianceChecker := &ai.Agent{
		Name:  "compliance",
		Model: mockModel,
		Instructions: "You are a compliance specialist. Check the given topic for compliance considerations: " +
			"data privacy, regulatory requirements, and industry standards.",
	}
	analysis := &ai.Agent{
		Name:     "analysis",
		Model:    mockModel,
		Agents:   []*ai.Agent{marketAnalyst, riskAnalyst, complianceChecker},
		Strategy: ai.StrategyParallel,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	result, err := runtime.Run(ctx, analysis, "Launching an AI-powered healthcare diagnostic tool in the US market")
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}

	// Validation, in place of the example's result.print_result().
	if result.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	if len(recorded) != 3 {
		t.Fatalf("expected three recorded analyses, found %d", len(recorded))
	}
	for _, answer := range recorded {
		first := strings.TrimSpace(strings.SplitN(strings.TrimSpace(answer), "\n", 2)[0])
		if !strings.Contains(result.Output, first) {
			t.Errorf("combined output lacks an analyst's answer starting %q\n--- output ---\n%.400s", first, result.Output)
		}
	}
}
