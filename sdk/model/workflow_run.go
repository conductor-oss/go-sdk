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
// Status is RUNNING
func (w *WorkflowRun) IsRunning() bool {
	return w.Status == RunningWorkflow
}

// IsPaused returns true if the workflow is currently paused.
// Status is PAUSED
func (w *WorkflowRun) IsPaused() bool {
	return w.Status == PausedWorkflow
}

// IsFailed returns true if the workflow has failed.
// Status is FAILED
func (w *WorkflowRun) IsFailed() bool {
	return w.Status == FailedWorkflow
}

// IsCompleted returns true if the workflow has completed successfully.
// Status is COMPLETED
func (w *WorkflowRun) IsCompleted() bool {
	return w.Status == CompletedWorkflow
}

// IsTimedOut returns true if the workflow has timed out.
// Status is TIMED_OUT
func (w *WorkflowRun) IsTimedOut() bool {
	return w.Status == TimedOutWorkflow
}

// IsTerminated returns true if the workflow has been terminated.
// Status is TERMINATED
func (w *WorkflowRun) IsTerminated() bool {
	return w.Status == TerminatedWorkflow
}

// GetFailedTasks returns all tasks that have failed.
// Status is FAILED or FAILED_WITH_TERMINAL_ERROR
func (w *WorkflowRun) GetFailedTasks() []Task {
	return w.GetTasksByStatus(FailedTask, FailedWithTerminalErrorTask)
}

// GetCompletedTasks returns all tasks that have completed successfully.
// Status is COMPLETED
func (w *WorkflowRun) GetCompletedTasks() []Task {
	return w.GetTasksByStatus(CompletedTask)
}

// GetTasksByStatus returns all tasks with the specified status(es).
func (w *WorkflowRun) GetTasksByStatus(statuses ...TaskResultStatus) []Task {
	var filteredTasks []Task
	for _, task := range w.Tasks {
		for _, status := range statuses {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
				break
			}
		}
	}
	return filteredTasks
}

// GetTaskByReferenceName returns the task with the specified reference name.
func (w *WorkflowRun) GetTaskByReferenceName(referenceTaskName string) *Task {
	for _, task := range w.Tasks {
		if task.ReferenceTaskName == referenceTaskName {
			return &task
		}
	}
	return nil
}
