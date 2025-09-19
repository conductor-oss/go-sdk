package model

type WorkflowRun struct {
	CorrelationId        string                 `json:"correlationId,omitempty"`
	CreateTime           int64                  `json:"createTime,omitempty"`
	CreatedBy            string                 `json:"createdBy,omitempty"`
	Input                map[string]interface{} `json:"input,omitempty"`
	Output               map[string]interface{} `json:"output,omitempty"`
	Priority             int32                  `json:"priority,omitempty"`
	RequestId            string                 `json:"requestId,omitempty"`
	ResponseType         ReturnStrategy         `json:"responseType,omitempty"`
	Status               WorkflowStatus         `json:"status,omitempty"`
	TargetWorkflowId     string                 `json:"targetWorkflowId,omitempty"`
	TargetWorkflowStatus WorkflowStatus         `json:"targetWorkflowStatus,omitempty"`
	Tasks                []Task                 `json:"tasks,omitempty"`
	UpdateTime           int64                  `json:"updateTime,omitempty"`
	Variables            map[string]interface{} `json:"variables,omitempty"`
	WorkflowId           string                 `json:"workflowId,omitempty"`
}

// IsRunning returns true if the workflow is currently running.
// status is RUNNING
func (w *WorkflowRun) IsRunning() bool {
	return w.Status == RunningWorkflow
}

// IsPaused returns true if the workflow is currently paused.
// status is PAUSED
func (w *WorkflowRun) IsPaused() bool {
	return w.Status == PausedWorkflow
}

// IsFailed returns true if the workflow has failed.
// status is FAILED
func (w *WorkflowRun) IsFailed() bool {
	return w.Status == FailedWorkflow
}

// IsCompleted returns true if the workflow has completed successfully.
// status is COMPLETED
func (w *WorkflowRun) IsCompleted() bool {
	return w.Status == CompletedWorkflow
}

// IsTimedOut returns true if the workflow has timed out.
// status is TIMED_OUT
func (w *WorkflowRun) IsTimedOut() bool {
	return w.Status == TimedOutWorkflow
}

// IsTerminated returns true if the workflow has been terminated.
// status is TERMINATED
func (w *WorkflowRun) IsTerminated() bool {
	return w.Status == TerminatedWorkflow
}
