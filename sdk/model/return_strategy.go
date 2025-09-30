package model

type ReturnStrategy string

const (
	// ReturnTargetWorkflow - Returns the state of the originally triggered workflow
	ReturnTargetWorkflow ReturnStrategy = "TARGET_WORKFLOW" // Default
	// ReturnBlockingWorkflow - Returns the state of the workflow that is currently blocking the execution, which may be a sub-workflow.
	ReturnBlockingWorkflow ReturnStrategy = "BLOCKING_WORKFLOW"
	// ReturnBlockingTask - Returns the execution status of the task that is currently blocking workflow execution.
	ReturnBlockingTask ReturnStrategy = "BLOCKING_TASK"
	// ReturnBlockingTaskInput - Returns the input of the task that is currently blocking workflow execution.
	ReturnBlockingTaskInput ReturnStrategy = "BLOCKING_TASK_INPUT"
)

func (r ReturnStrategy) String() string {
	return string(r)
}
