//go:build ignore

// Plan and compile — a planner writes a plan, the server compiles and runs it.
//
// Run with:  go run agents/103_plan_and_compile.go [topic]
//
// A plan-execute agent asks its planner for a plan over three tools, then the
// server compiles that plan into a workflow and executes it: five factorials
// in parallel, a generated summary, and a validation step with a success
// condition. Afterwards the compiled workflow is fetched and summarized.
//
// Requirements:
//   - Conductor server with LLM support
//   - CONDUCTOR_SERVER_URL=http://localhost:8080/api in the environment
//   - CONDUCTOR_AGENT_LLM_MODEL in the environment (optional)
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
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

// The tool descriptions are the Python functions' docstrings, verbatim: the
// planner's prompt is built from them.
const (
	factorialDoc = "Compute n! and return it as a string.\n\nArgs:\n    n: Non-negative integer. Capped at 20 to keep things sane."
	summaryDoc   = "Persist a short summary string. Returns it back for the validator."
	checkDoc     = "Return JSON ``{passed, length, min_chars}`` for the validator.\n\nArgs:\n    text: The summary to check.\n    min_chars: Minimum acceptable length in characters."
)

func factorial(ctx context.Context, in factorialIn) (string, error) {
	if in.N < 0 || in.N > 20 {
		return fmt.Sprintf("ERROR: n must be in [0, 20], got %d", in.N), nil
	}
	result := 1
	for i := 2; i <= in.N; i++ {
		result *= i
	}
	return fmt.Sprint(result), nil
}

func writeSummary(ctx context.Context, in summaryIn) (string, error) {
	fmt.Printf("[write_summary] %s\n", in.Text)
	return in.Text, nil
}

func checkSummary(ctx context.Context, in checkIn) (string, error) {
	length := utf8.RuneCountInString(in.Text)
	return fmt.Sprintf(`{"passed": %t, "length": %d, "min_chars": %d}`, length >= in.MinChars, length, in.MinChars), nil
}

const plannerInstructions = "You are a math-explainer planner. Plan a workflow that:\n\n" +
	"1. Computes factorials of 1, 2, 3, 4, 5 in PARALLEL using ``factorial`` (static args).\n" +
	"2. Writes a short prose summary about factorial growth using ``write_summary``\n" +
	"   (use a ``generate`` block — the LLM produces the ``text`` arg at run time).\n" +
	"3. Validates the summary is at least 30 characters via ``check_summary``,\n" +
	"   with ``success_condition: \"$.passed === true\"``.\n"

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}
	topic := "factorials"
	if len(os.Args) > 1 {
		topic = strings.Join(os.Args[1:], " ")
	}

	tools := []ai.ToolDef{
		tool.Func("factorial", factorialDoc, factorial),
		tool.Func("write_summary", summaryDoc, writeSummary),
		tool.Func("check_summary", checkDoc, checkSummary),
	}
	// The Python example builds this with plan_execute(...): a planner and a
	// fallback sub-agent named after the harness.
	harness := &ai.Agent{
		Name:     "plan_and_compile_demo",
		Model:    model,
		Strategy: ai.StrategyPlanExecute,
		Tools:    tools,
		Planner: &ai.Agent{
			Name:         "plan_and_compile_demo_planner",
			Model:        model,
			Instructions: plannerInstructions,
		},
		Fallback: &ai.Agent{
			Name:         "plan_and_compile_demo_fallback",
			Model:        model,
			Instructions: "The plan failed. Use the available tools to recover.",
			Tools:        tools,
		},
		FallbackMaxTurns: 4,
	}

	fmt.Printf("\n=== PLAN_AND_COMPILE demo ===\nTopic: %s\nModel: %s\n\n", topic, model)
	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	result, err := runtime.Run(context.Background(), harness, "Topic: "+topic)
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()

	// The compiled plan lives in a PLAN_AND_COMPILE task somewhere in the
	// workflow tree; walk parent and sub-workflows to find it.
	pac := findPlanAndCompileOutput(result.ExecutionID)
	if pac == nil {
		fmt.Println("(!) No PLAN_AND_COMPILE task found in workflow tree")
		os.Exit(1)
	}
	fmt.Println("--- PLAN_AND_COMPILE output ---")
	fmt.Printf("error:         %v\n", pac["error"])
	fmt.Printf("workflowName:  %v\n", pac["workflowName"])
	if stats, ok := pac["stats"].(map[string]any); ok {
		fmt.Printf("stats:         stepCount=%v, taskCount=%v\n", stats["stepCount"], stats["taskCount"])
	}
	if wfDef, ok := pac["workflowDef"].(map[string]any); ok {
		tasks, _ := wfDef["tasks"].([]any)
		fmt.Printf("\ntop-level tasks in compiled WorkflowDef (%d):\n", len(tasks))
		for _, t := range tasks {
			task, _ := t.(map[string]any)
			fmt.Printf("  - %-12v ref=%v\n", task["type"], task["taskReferenceName"])
		}
	}
}

func findPlanAndCompileOutput(executionID string) map[string]any {
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
			continue
		}
		var wf struct {
			Tasks []map[string]any `json:"tasks"`
		}
		json.NewDecoder(resp.Body).Decode(&wf)
		resp.Body.Close()
		for _, t := range wf.Tasks {
			if t["taskType"] == "PLAN_AND_COMPILE" {
				out, _ := t["outputData"].(map[string]any)
				return out
			}
			if sub, ok := t["subWorkflowId"].(string); ok && sub != "" && !seen[sub] {
				pending = append(pending, sub)
			}
		}
	}
	return nil
}
