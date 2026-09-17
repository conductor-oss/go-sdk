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
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type factorialIn struct {
	N int `json:"n"`
}

type summaryIn struct {
	Text string `json:"text"`
}

type checkIn struct {
	Text     string `json:"text"`
	MinChars int    `json:"min_chars"`
}

const (
	factorialDoc = "Compute n! and return it as a string.\n\nArgs:\n    n: Non-negative integer. Capped at 20 to keep things sane."
	summaryDoc   = "Persist a short summary string. Returns it back for the validator."
	checkDoc     = "Return JSON ``{passed, length, min_chars}`` for the validator.\n\nArgs:\n    text: The summary to check.\n    min_chars: Minimum acceptable length in characters."

	plannerInstructions = "You are a math-explainer planner. Plan a workflow that:\n\n" +
		"1. Computes factorials of 1, 2, 3, 4, 5 in PARALLEL using ``factorial`` (static args).\n" +
		"2. Writes a short prose summary about factorial growth using ``write_summary``\n" +
		"   (use a ``generate`` block — the LLM produces the ``text`` arg at run time).\n" +
		"3. Validates the summary is at least 30 characters via ``check_summary``,\n" +
		"   with ``success_condition: \"$.passed === true\"``.\n"
)

// Plan and compile — the Python SDK's examples/agents/103_plan_and_compile.py
// as a test.
//
// A plan-execute agent has its planner write a plan over three tools; the
// server compiles the plan into a workflow and runs it. This is that flow,
// copied, with the print replaced by validation: the run completes, the
// factorial tool ran once for each of 1..5, the summary was written and then
// checked, the check passed, and the workflow tree holds a PLAN_AND_COMPILE
// task that compiled without error.
func TestExample103PlanAndCompile(t *testing.T) {
	runtime := newRuntime(t)
	recordedAnswers(t, "103_plan_and_compile") // skips when the recordings are not configured

	var mu sync.Mutex
	var factorials []int
	var summaries, checks []string
	factorial := func(ctx context.Context, in factorialIn) (string, error) {
		mu.Lock()
		factorials = append(factorials, in.N)
		mu.Unlock()
		if in.N < 0 || in.N > 20 {
			return fmt.Sprintf("ERROR: n must be in [0, 20], got %d", in.N), nil
		}
		result := 1
		for i := 2; i <= in.N; i++ {
			result *= i
		}
		return fmt.Sprint(result), nil
	}
	writeSummary := func(ctx context.Context, in summaryIn) (string, error) {
		mu.Lock()
		summaries = append(summaries, in.Text)
		mu.Unlock()
		return in.Text, nil
	}
	checkSummary := func(ctx context.Context, in checkIn) (string, error) {
		length := utf8.RuneCountInString(in.Text)
		out := fmt.Sprintf(`{"passed": %t, "length": %d, "min_chars": %d}`, length >= in.MinChars, length, in.MinChars)
		mu.Lock()
		checks = append(checks, out)
		mu.Unlock()
		return out, nil
	}

	tools := []ai.ToolDef{
		tool.Func("factorial", factorialDoc, factorial),
		tool.Func("write_summary", summaryDoc, writeSummary),
		tool.Func("check_summary", checkDoc, checkSummary),
	}
	harness := &ai.Agent{
		Name:     "plan_and_compile_demo",
		Model:    mockModel,
		Strategy: ai.StrategyPlanExecute,
		Tools:    tools,
		Planner: &ai.Agent{
			Name:         "plan_and_compile_demo_planner",
			Model:        mockModel,
			Instructions: plannerInstructions,
		},
		Fallback: &ai.Agent{
			Name:         "plan_and_compile_demo_fallback",
			Model:        mockModel,
			Instructions: "The plan failed. Use the available tools to recover.",
			Tools:        tools,
		},
		FallbackMaxTurns: 4,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	result, err := runtime.Run(ctx, harness, "Topic: factorials")
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}

	// Validation, in place of the example's printing.
	if result.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	mu.Lock()
	defer mu.Unlock()
	sort.Ints(factorials)
	if fmt.Sprint(factorials) != "[1 2 3 4 5]" {
		t.Errorf("factorial ran for %v, want 1..5 once each", factorials)
	}
	if len(summaries) != 1 || len(checks) != 1 {
		t.Errorf("write_summary ran %d time(s) and check_summary %d, want 1 each", len(summaries), len(checks))
	}
	if len(checks) == 1 && !strings.Contains(checks[0], `"passed": true`) {
		t.Errorf("the summary did not pass its check: %s (summary %q)", checks[0], summaries)
	}

	// The compiled plan: a PLAN_AND_COMPILE task, somewhere in the tree, with no error.
	pac := findPlanAndCompile(t, result.ExecutionID)
	if pac == nil {
		t.Fatal("no PLAN_AND_COMPILE task in the workflow tree")
	}
	if e, _ := pac["error"].(string); e != "" {
		t.Errorf("plan compilation reported an error: %s", e)
	}
	if stats, ok := pac["stats"].(map[string]any); ok {
		t.Logf("compiled plan: stepCount=%v taskCount=%v", stats["stepCount"], stats["taskCount"])
	}
}

func findPlanAndCompile(t *testing.T, executionID string) map[string]any {
	t.Helper()
	base := strings.TrimRight(os.Getenv("CONDUCTOR_SERVER_URL"), "/")
	seen := map[string]bool{}
	pending := []string{executionID}
	for len(pending) > 0 {
		id := pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		if seen[id] {
			continue
		}
		seen[id] = true
		resp, err := http.Get(base + "/workflow/" + id + "?includeTasks=true")
		if err != nil {
			t.Fatalf("read workflow %s: %v", id, err)
		}
		var wf struct {
			Tasks []map[string]any `json:"tasks"`
		}
		derr := json.NewDecoder(resp.Body).Decode(&wf)
		resp.Body.Close()
		if derr != nil {
			t.Fatalf("decode workflow %s: %v", id, derr)
		}
		for _, task := range wf.Tasks {
			if task["taskType"] == "PLAN_AND_COMPILE" {
				out, _ := task["outputData"].(map[string]any)
				return out
			}
			if sub, ok := task["subWorkflowId"].(string); ok && sub != "" && !seen[sub] {
				pending = append(pending, sub)
			}
		}
	}
	return nil
}
