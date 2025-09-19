package model

import "fmt"

// SignalResponse represents a unified response from the signal API
// It directly maps to the JSON response from the API
type SignalResponse struct {
	// Common fields in all responses
	ResponseType         ReturnStrategy         `json:"responseType"`
	TargetWorkflowId     string                 `json:"targetWorkflowId"`
	TargetWorkflowStatus WorkflowStatus         `json:"targetWorkflowStatus"`
	WorkflowId           string                 `json:"workflowId"`
	Input                map[string]interface{} `json:"input"`
	Output               map[string]interface{} `json:"output"`
	Priority             int32                  `json:"priority,omitempty"`
	Variables            map[string]interface{} `json:"variables,omitempty"`
	CorrelationID        string                 `json:"correlationId,omitempty"`
	RequestID            string                 `json:"requestId,omitempty"`

	// Fields specific to TARGET_WORKFLOW & BLOCKING_WORKFLOW
	Tasks      []Task         `json:"tasks,omitempty"`
	CreatedBy  string         `json:"createdBy,omitempty"`
	CreateTime int64          `json:"createTime,omitempty"`
	Status     WorkflowStatus `json:"status,omitempty"`
	UpdateTime int64          `json:"updateTime,omitempty"`

	// Fields specific to BLOCKING_TASK & BLOCKING_TASK_INPUT
	TaskType          string `json:"taskType,omitempty"`
	TaskId            string `json:"taskId,omitempty"`
	ReferenceTaskName string `json:"referenceTaskName,omitempty"`
	RetryCount        int32  `json:"retryCount,omitempty"`
	TaskDefName       string `json:"taskDefName,omitempty"`
	WorkflowType      string `json:"workflowType,omitempty"`
}

// Type check methods
func (r *SignalResponse) IsTargetWorkflow() bool {
	return r.ResponseType == ReturnTargetWorkflow
}

func (r *SignalResponse) IsBlockingWorkflow() bool {
	return r.ResponseType == ReturnBlockingWorkflow
}

func (r *SignalResponse) IsBlockingTask() bool {
	return r.ResponseType == ReturnBlockingTask
}

func (r *SignalResponse) IsBlockingTaskInput() bool {
	return r.ResponseType == ReturnBlockingTaskInput
}

// GetWorkflow extracts workflow details from a SignalResponse
func (r *SignalResponse) GetWorkflow() (*Workflow, error) {
	if r.ResponseType != ReturnTargetWorkflow && r.ResponseType != ReturnBlockingWorkflow {
		return nil, fmt.Errorf("response type %s does not contain workflow details", r.ResponseType)
	}

	return &Workflow{
		WorkflowId: r.WorkflowId,
		Status:     r.Status,
		Tasks:      r.Tasks,
		CreatedBy:  r.CreatedBy,
		CreateTime: r.CreateTime,
		UpdateTime: r.UpdateTime,
		Input:      r.Input,
		Output:     r.Output,
		Variables:  r.Variables,
		Priority:   r.Priority,
	}, nil
}

// GetBlockingTask extracts task details from a SignalResponse
// This method provides comprehensive field-by-field mapping from SignalResponse to Task
func (r *SignalResponse) GetBlockingTask() (*Task, error) {
	if r.ResponseType != ReturnBlockingTask && r.ResponseType != ReturnBlockingTaskInput {
		return nil, fmt.Errorf("response type %s does not contain task details", r.ResponseType)
	}

	// Convert WorkflowStatus to TaskResultStatus with comprehensive mapping
	var taskStatus TaskResultStatus
	switch r.Status {
	case RunningWorkflow:
		taskStatus = InProgressTask
	case CompletedWorkflow:
		taskStatus = CompletedTask
	case FailedWorkflow:
		taskStatus = FailedTask
	case TimedOutWorkflow:
		taskStatus = FailedTask // Map TIMED_OUT to FAILED for task status
	case TerminatedWorkflow:
		taskStatus = FailedTask // Map TERMINATED to FAILED for task status
	case PausedWorkflow:
		taskStatus = InProgressTask // Map PAUSED to IN_PROGRESS for task status
	default:
		// For unknown statuses, try direct conversion or default to IN_PROGRESS
		if string(r.Status) != "" {
			taskStatus = TaskResultStatus(r.Status)
		} else {
			taskStatus = InProgressTask
		}
	}

	// Comprehensive field-by-field mapping from SignalResponse to Task
	return &Task{
		// Core task identification fields
		TaskId:            r.TaskId,
		TaskType:          r.TaskType,
		TaskDefName:       r.TaskDefName,
		WorkflowType:      r.WorkflowType,
		ReferenceTaskName: r.ReferenceTaskName,
		RetryCount:        r.RetryCount,
		// Status and execution fields
		Status: taskStatus,
		// Data fields
		InputData:  r.Input,
		OutputData: r.Output,
		// Workflow context fields
		WorkflowInstanceId: r.WorkflowId,
	}, nil
}

// GetTaskRun extracts task run details from a SignalResponse
func (r *SignalResponse) GetTaskRun() TaskRun {
	// Convert WorkflowStatus to TaskResultStatus with comprehensive mapping
	var taskStatus TaskResultStatus
	switch r.Status {
	case RunningWorkflow:
		taskStatus = InProgressTask
	case CompletedWorkflow:
		taskStatus = CompletedTask
	case FailedWorkflow:
		taskStatus = FailedTask
	case TimedOutWorkflow:
		taskStatus = FailedTask // Map TIMED_OUT to FAILED for task status
	case TerminatedWorkflow:
		taskStatus = FailedTask // Map TERMINATED to FAILED for task status
	case PausedWorkflow:
		taskStatus = InProgressTask // Map PAUSED to IN_PROGRESS for task status
	default:
		if string(r.Status) != "" {
			taskStatus = TaskResultStatus(r.Status)
		} else {
			taskStatus = InProgressTask
		}
	}

	// Comprehensive field-by-field mapping from SignalResponse to TaskRun
	return TaskRun{
		// Core task identification fields
		TaskId:            r.TaskId,
		TaskType:          r.TaskType,
		TaskDefName:       r.TaskDefName,
		WorkflowType:      r.WorkflowType,
		ReferenceTaskName: r.ReferenceTaskName,
		RetryCount:        r.RetryCount,

		// Status and execution fields
		Status: taskStatus,

		// Data fields
		InputData:  r.Input,
		OutputData: r.Output,
		// Workflow context fields
		Variables: r.Variables,
		Priority:  int(r.Priority), // Convert int32 to int
		// Timing fields
		CreateTime: r.CreateTime,
		UpdateTime: r.UpdateTime,
		// Metadata fields
		CreatedBy: r.CreatedBy,
	}
}

func (r *SignalResponse) GetWorkflowRun() WorkflowRun {
	return WorkflowRun{
		WorkflowId:           r.WorkflowId,
		CorrelationId:        r.CorrelationID,
		Priority:             r.Priority,
		Status:               r.Status,
		Input:                r.Input,
		Output:               r.Output,
		Tasks:                r.Tasks,
		CreatedBy:            r.CreatedBy,
		CreateTime:           r.CreateTime,
		UpdateTime:           r.UpdateTime,
		Variables:            r.Variables,
		ResponseType:         r.ResponseType,
		TargetWorkflowId:     r.TargetWorkflowId,
		TargetWorkflowStatus: r.TargetWorkflowStatus,
	}
}

// GetTaskInput extracts task input from a SignalResponse
func (r *SignalResponse) GetTaskInput() (map[string]interface{}, error) {
	if r.ResponseType != ReturnBlockingTaskInput {
		return nil, fmt.Errorf("response type %s does not contain task input details", r.ResponseType)
	}

	return r.Input, nil
}
