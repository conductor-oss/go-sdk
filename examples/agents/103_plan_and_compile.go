package agents

import (
	"context"
	"fmt"
	"io"
	"unicode/utf8"

	"github.com/antihax/optional"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
	"github.com/conductor-sdk/conductor-go/sdk/client"
)

// The tool descriptions are the Python functions' docstrings, verbatim: the
// planner's prompt is built from them.
const (
	factorialDoc    = "Compute n! and return it as a string.\n\nArgs:\n    n: Non-negative integer. Capped at 20 to keep things sane."
	writeSummaryDoc = "Persist a short summary string. Returns it back for the validator."
	checkSummaryDoc = "Return JSON ``{passed, length, min_chars}`` for the validator.\n\nArgs:\n    text: The summary to check.\n    min_chars: Minimum acceptable length in characters."

	plannerInstructions = "You are a math-explainer planner. Plan a workflow that:\n\n1. Computes factorials of 1, 2, 3, 4, 5 in PARALLEL using ``factorial`` (static args).\n2. Writes a short prose summary about factorial growth using ``write_summary``\n   (use a ``generate`` block — the LLM produces the ``text`` arg at run time).\n3. Validates the summary is at least 30 characters via ``check_summary``,\n   with ``success_condition: \"$.passed === true\"``.\n"
)

func factorial(_ context.Context, in factorialIn) (string, error) {
	if in.N < 0 || in.N > 20 {
		return fmt.Sprintf("ERROR: n must be in [0, 20], got %d", in.N), nil
	}
	result := 1
	for i := 2; i <= in.N; i++ {
		result *= i
	}
	return fmt.Sprint(result), nil
}

func writeSummary(_ context.Context, in summaryIn) (string, error) {
	return in.Text, nil
}

func checkSummary(_ context.Context, in checkSummaryIn) (string, error) {
	length := utf8.RuneCountInString(in.Text)
	return fmt.Sprintf(`{"passed": %t, "length": %d, "min_chars": %d}`, length >= in.MinChars, length, in.MinChars), nil
}

// PlanAndCompile is the Python SDK's examples/agents/103_plan_and_compile.py:
// a planner writes a plan over three tools and the server compiles and runs
// it. Python builds it with plan_execute(), which names the planner and
// fallback after the harness.
func PlanAndCompile(model string) *ai.Agent {
	tools := ai.Tools(
		tool.Func("factorial", factorial, factorialDoc),
		tool.Func("write_summary", writeSummary, writeSummaryDoc),
		tool.Func("check_summary", checkSummary, checkSummaryDoc),
	)
	return &ai.Agent{
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
}

func runPlanAndCompile(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, PlanAndCompile(Model()), "Topic: factorials", out)
	if err != nil {
		return nil, err
	}
	compiled, err := findPlanAndCompileOutput(ctx, res.ExecutionID)
	if err != nil {
		return nil, err
	}
	if compiled == nil {
		return nil, fmt.Errorf("no PLAN_AND_COMPILE task found in the workflow tree")
	}
	if compiled["error"] != nil {
		return nil, fmt.Errorf("plan compilation failed: %v", compiled["error"])
	}
	fmt.Fprintf(out, "Compiled workflow: %v\n", compiled["workflowName"])
	fmt.Fprintf(out, "Stats: %v\n", compiled["stats"])
	if def, ok := compiled["workflowDef"].(map[string]any); ok {
		tasks, _ := def["tasks"].([]any)
		for _, t := range tasks {
			task, _ := t.(map[string]any)
			fmt.Fprintf(out, "%v: %v\n", task["type"], task["taskReferenceName"])
		}
	}
	return []*ai.AgentResult{res}, nil
}

// findPlanAndCompileOutput walks the execution and its sub-workflows for the
// PLAN_AND_COMPILE task's output.
func findPlanAndCompileOutput(ctx context.Context, executionID string) (map[string]any, error) {
	svc := &client.WorkflowResourceApiService{APIClient: client.NewAPIClientFromEnv()}
	visited := map[string]bool{}
	pending := []string{executionID}
	for len(pending) > 0 {
		id := pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		if visited[id] {
			continue
		}
		visited[id] = true
		wf, _, err := svc.GetExecutionStatus(ctx, id,
			&client.WorkflowResourceApiGetExecutionStatusOpts{IncludeTasks: optional.NewBool(true)})
		if err != nil {
			return nil, fmt.Errorf("get workflow %s: %w", id, err)
		}
		for _, task := range wf.Tasks {
			if task.TaskType == "PLAN_AND_COMPILE" {
				return task.OutputData, nil
			}
			if task.SubWorkflowId != "" {
				pending = append(pending, task.SubWorkflowId)
			}
		}
	}
	return nil, nil
}
