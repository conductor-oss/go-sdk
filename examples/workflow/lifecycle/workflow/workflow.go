package workflow

import (
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

const (
	FailureWorkflowName   = "api_showcase_failure_handler"
	LifecycleWorkflowName = "lifecycle_demo_workflow"
)

func CreateFailureWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	wf := workflow.NewConductorWorkflow(executor).
		Name(FailureWorkflowName).
		Version(1).
		Description("Handles failures from terminated workflows")

	// Log failure details
	logFailure := workflow.NewSetVariableTask("log_failure").
		Input("failure_reason", "${workflow.input.failure_reason}").
		Input("failed_workflow_id", "${workflow.input.failed_workflow_id}").
		Input("failure_logged", true)

	// Send notification (simulated with set variable)
	notifyFailure := workflow.NewSetVariableTask("notify_failure").
		Input("notification_sent", true).
		Input("notification_message", "Workflow failed: ${workflow.input.failed_workflow_id}")

	// Add a wait to make the failure workflow observable
	// This gives time to detect it as "running"
	observableWait := workflow.NewWaitForDurationTask("failure_handler_wait", 5*time.Second)

	wf.Add(logFailure).Add(notifyFailure).Add(observableWait).OwnerEmail("owner@example.com")

	return wf
}

func CreateLifecycleWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	// Create comprehensive workflow for lifecycle demonstration
	wf := workflow.NewConductorWorkflow(executor).
		Name(LifecycleWorkflowName).
		Version(1).
		Description("Comprehensive workflow for lifecycle API demonstration")

	// Task 1: Initial wait task (2 seconds)
	initialWait := workflow.NewWaitForDurationTask("wait_for_2_sec", 2*time.Second)

	// Task 2: Set variable task to store workflow state
	setVariables := workflow.NewSetVariableTask("set_workflow_variables").
		Input("workflow_input", "${workflow.input}").
		Input("workflow_name", "${workflow.workflowType}").
		Input("correlation_id", "${workflow.correlationId}")

	// Task 3: Wait task that can be completed manually
	waitForSignal := workflow.NewWaitTask("wait_for_signal").
		Description("This task waits indefinitely until manually completed")

	// Task 4: HTTP task (similar to Python example)
	httpCall := workflow.NewHttpTask("call_remote_api", &workflow.HttpInput{
		Uri: "https://orkes-api-tester.orkesconductor.com/api",
	}).Optional(true) // Make optional so workflow can continue even if it fails

	// Task 5: Another wait task for demonstrating pause/resume
	pauseCheckpoint := workflow.NewWaitForDurationTask("pause_checkpoint", 3*time.Second)

	// Task 6: Final task to mark completion
	finalTask := workflow.NewSetVariableTask("mark_completion").
		Input("workflow_id", "${workflow.workflowId}").
		Input("status", "completed")

	// Chain all tasks together
	wf.Add(initialWait).
		Add(setVariables).
		Add(waitForSignal).
		Add(httpCall).
		Add(pauseCheckpoint).
		Add(finalTask).
		OwnerEmail("owner@example.com")

	return wf
}
