package workflow

import (
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

// WorkflowDefinition holds a workflow and its metadata
type WorkflowDefinition struct {
	Workflow *workflow.ConductorWorkflow
	Name     string
}

const (
	SimpleJumpWorkflowName    = "jump_to_task_demo"
	WaitBasedJumpWorkflowName = "wait_jump_demo"
)

// CreateJumpDemoWorkflow creates a workflow with multiple tasks for jump demonstration
func CreateSimpleJumpWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	wf := workflow.NewConductorWorkflow(executor).
		Name(SimpleJumpWorkflowName).
		Version(1).
		Description("Workflow demonstrating JumpToTask API functionality")

	// Task 1: Initial data processing
	task1 := workflow.NewSimpleTask("process_initial_data", "task1_ref").
		Input("step", "1").
		Input("message", "Processing initial data").
		Description("Process initial workflow data")

	// Task 2: Business data validation
	task2 := workflow.NewSimpleTask("validate_business_data", "task2_ref").
		Input("step", "2").
		Input("data", "${task1_ref.output.processedData}").
		Description("Validate business data integrity")

	// Task 3: Apply business logic rules
	task3 := workflow.NewSimpleTask("apply_business_logic", "task3_ref").
		Input("step", "3").
		Input("validatedData", "${task2_ref.output.validatedData}").
		Description("Apply business logic rules")

	// Task 4: Call external API (jump target)
	httpTask := workflow.NewHttpTask("call_external_api", &workflow.HttpInput{
		Uri:    "https://jsonplaceholder.typicode.com/posts/1",
		Method: "GET",
	}).Description("Call external API service")

	// Task 5: Complete final processing
	task5 := workflow.NewSimpleTask("complete_final_process", "task5_ref").
		Input("step", "5").
		Input("apiResponse", "${call_external_api.output.response}").
		Description("Complete final workflow processing")

	// Chain all tasks together
	wf.Add(task1).Add(task2).Add(task3).Add(httpTask).Add(task5)

	return wf
}

// CreateWaitBasedJumpWorkflow creates a workflow with wait tasks for controlled jumping
func CreateWaitBasedJumpWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	wf := workflow.NewConductorWorkflow(executor).
		Name(WaitBasedJumpWorkflowName).
		Version(1).
		Description("Workflow with wait tasks for controlled jump demonstration")

	// Task 1: Initialize workflow state
	task1 := workflow.NewSimpleTask("initialize_workflow_state", "init_ref").
		Input("step", "initialization").
		Description("Initialize workflow state")

	// Task 2: Wait for 5 seconds (will be skipped)
	waitTask := workflow.NewWaitForDurationTask("wait_for_delay", 5*time.Second).
		Description("Wait for time delay - will be skipped via jump")

	// Task 3: Process middle step (will be skipped)
	task3 := workflow.NewSimpleTask("process_middle_step", "middle_ref").
		Input("step", "middle_processing").
		Description("Process middle workflow step")

	// Task 4: Execute API call (jump target)
	httpTask := workflow.NewHttpTask("execute_api_call", &workflow.HttpInput{
		Uri:    "https://httpbin.org/get",
		Method: "GET",
	}).Description("Execute external API call")

	// Task 5: Store final results
	setVarTask := workflow.NewSetVariableTask("store_final_results").
		Input("jump_completed", true).
		Input("api_result", "${execute_api_call.output.response}").
		Description("Store final workflow results")

	wf.Add(task1).Add(waitTask).Add(task3).Add(httpTask).Add(setVarTask)

	return wf
}
