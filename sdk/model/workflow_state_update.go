package model

// WorkflowStateUpdate is the request body for the UpdateWorkflowAndTaskState endpoint.
type WorkflowStateUpdate struct {
	// TaskReferenceName the reference name of the task to update.
	TaskReferenceName string `json:"taskReferenceName,omitempty"`
	// TaskResult the result of the task to update.
	TaskResult *TaskResult `json:"taskResult,omitempty"`
	// Variables the variables to update.
	Variables map[string]interface{} `json:"variables,omitempty"`
}
