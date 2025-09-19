package model

// UpgradeWorkflowRequest is the request body for the UpgradeRunningWorkflowToVersionå.
type UpgradeWorkflowRequest struct {
	// Name the name of the workflow definition.
	Name string `json:"name"`
	// TaskOutput a map of task outputs for any skipped tasks,
	// with the key as the task reference name, and the value as the task output object.
	TaskOutput map[string]interface{} `json:"taskOutput,omitempty"`
	// Version the version to which the workflow is to be updated.
	Version int32 `json:"version,omitempty"`
	// WorkflowInput a map of inputs for the upgraded workflow execution,
	// with the parameter name as the key and its input value as the value.
	WorkflowInput map[string]interface{} `json:"workflowInput,omitempty"`
}
