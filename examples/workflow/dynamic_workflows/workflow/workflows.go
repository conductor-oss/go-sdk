package workflow

import (
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

const (
	DynamicWorkflowName = "dynamic_workflow"
)

// CreateEmptyDynamicWorkflow creates an empty workflow that can have tasks added dynamically
func CreateEmptyDynamicWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	// Create empty workflow - tasks will be added in main function
	return workflow.NewConductorWorkflow(executor).
		Name(DynamicWorkflowName).
		Version(1)
}
