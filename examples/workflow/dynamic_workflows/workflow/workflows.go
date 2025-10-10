package workflow

import (
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

// CreateEmptyDynamicWorkflow creates an empty workflow that can have tasks added dynamically
func CreateEmptyDynamicWorkflow(executor *executor.WorkflowExecutor, workflowName string) *workflow.ConductorWorkflow {
	// Create empty workflow - tasks will be added in main function
	return workflow.NewConductorWorkflow(executor).
		Name(workflowName).
		Version(1)
}
