//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

type Workflow struct {
	CorrelationId                    string                 `json:"correlationId,omitempty"`
	CreateTime                       int64                  `json:"createTime,omitempty"`
	CreatedBy                        string                 `json:"createdBy,omitempty"`
	EndTime                          int64                  `json:"endTime,omitempty"`
	Event                            string                 `json:"event,omitempty"`
	ExternalInputPayloadStoragePath  string                 `json:"externalInputPayloadStoragePath,omitempty"`
	ExternalOutputPayloadStoragePath string                 `json:"externalOutputPayloadStoragePath,omitempty"`
	FailedReferenceTaskNames         []string               `json:"failedReferenceTaskNames,omitempty"`
	FailedTaskNames                  []string               `json:"failedTaskNames,omitempty"`
	History                          []Workflow             `json:"history,omitempty"`
	IdempotencyKey                   string                 `json:"idempotencyKey,omitempty"`
	Input                            map[string]interface{} `json:"input,omitempty"`
	LastRetriedTime                  int64                  `json:"lastRetriedTime,omitempty"`
	Output                           map[string]interface{} `json:"output,omitempty"`
	OwnerApp                         string                 `json:"ownerApp,omitempty"`
	ParentWorkflowId                 string                 `json:"parentWorkflowId,omitempty"`
	ParentWorkflowTaskId             string                 `json:"parentWorkflowTaskId,omitempty"`
	Priority                         int32                  `json:"priority,omitempty"`
	RateLimitKey                     string                 `json:"rateLimitKey,omitempty"`
	RateLimited                      bool                   `json:"rateLimited,omitempty"`
	ReRunFromWorkflowId              string                 `json:"reRunFromWorkflowId,omitempty"`
	ReasonForIncompletion            string                 `json:"reasonForIncompletion,omitempty"`
	StartTime                        int64                  `json:"startTime,omitempty"`
	Status                           WorkflowStatus         `json:"status,omitempty"`
	TaskToDomain                     map[string]string      `json:"taskToDomain,omitempty"`
	Tasks                            []Task                 `json:"tasks,omitempty"`
	UpdateTime                       int64                  `json:"updateTime,omitempty"`
	UpdatedBy                        string                 `json:"updatedBy,omitempty"`
	Variables                        map[string]interface{} `json:"variables,omitempty"`
	WorkflowDefinition               *WorkflowDef           `json:"workflowDefinition,omitempty"`
	WorkflowId                       string                 `json:"workflowId,omitempty"`
	WorkflowName                     string                 `json:"workflowName,omitempty"`
	WorkflowVersion                  int32                  `json:"workflowVersion,omitempty"`
}

// IsRunning returns true if the workflow is currently running.
// Status is RUNNING
func (w *Workflow) IsRunning() bool {
	return w.Status == RunningWorkflow
}

// IsPaused returns true if the workflow is currently paused.
// Status is PAUSED
func (w *Workflow) IsPaused() bool {
	return w.Status == PausedWorkflow
}

// IsFailed returns true if the workflow has failed.
// Status is FAILED
func (w *Workflow) IsFailed() bool {
	return w.Status == FailedWorkflow
}

// IsCompleted returns true if the workflow has completed successfully.
// Status is COMPLETED
func (w *Workflow) IsCompleted() bool {
	return w.Status == CompletedWorkflow
}

// IsTimedOut returns true if the workflow has timed out.
// Status is TIMED_OUT
func (w *Workflow) IsTimedOut() bool {
	return w.Status == TimedOutWorkflow
}

// IsTerminated returns true if the workflow has been terminated.
// Status is TERMINATED
func (w *Workflow) IsTerminated() bool {
	return w.Status == TerminatedWorkflow
}

// GetInProgressTasks returns all tasks that are currently in progress.
// Status is IN_PROGRESS
func (w *Workflow) GetInProgressTasks() []Task {
	return w.GetTasksByStatus(InProgressTask)
}

// GetFailedTasks returns all tasks that have failed.
// Status is FAILED or FAILED_WITH_TERMINAL_ERROR
func (w *Workflow) GetFailedTasks() []Task {
	return w.GetTasksByStatus(FailedTask, FailedWithTerminalErrorTask)
}

// GetCompletedTasks returns all tasks that have completed successfully.
// Status is COMPLETED
func (w *Workflow) GetCompletedTasks() []Task {
	return w.GetTasksByStatus(CompletedTask)
}

// GetScheduledTasks returns all tasks that are scheduled but not yet started.
// Status is SCHEDULED
func (w *Workflow) GetScheduledTasks() []Task {
	return w.GetTasksByStatus(ScheduledTask)
}

// GetTasksByStatus returns all tasks with the specified status(es).
func (w *Workflow) GetTasksByStatus(statuses ...TaskResultStatus) []Task {
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
func (w *Workflow) GetTaskByReferenceName(referenceTaskName string) *Task {
	for _, task := range w.Tasks {
		if task.ReferenceTaskName == referenceTaskName {
			return &task
		}
	}
	return nil
}
