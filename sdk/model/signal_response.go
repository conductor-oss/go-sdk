package model

import (
	"fmt"

	"github.com/conductor-sdk/conductor-go/sdk/log"
)

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
	Tasks      []Task `json:"tasks,omitempty"`
	CreatedBy  string `json:"createdBy,omitempty"`
	CreateTime int64  `json:"createTime,omitempty"`
	Status     string `json:"status,omitempty"`
	UpdateTime int64  `json:"updateTime,omitempty"`

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
	workflowStatus, err := ParseWorkflowStatus(r.Status)
	if err != nil {
		log.Error("failed to parse workflow status", "error", err)
		workflowStatus = WorkflowStatus(r.Status)
	}

	return &Workflow{
		WorkflowId: r.WorkflowId,
		Status:     workflowStatus,
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
func (r *SignalResponse) GetBlockingTask() (*Task, error) {
	if r.ResponseType != ReturnBlockingTask && r.ResponseType != ReturnBlockingTaskInput {
		return nil, fmt.Errorf("response type %s does not contain task details", r.ResponseType)
	}

	taskStatus, err := ParseTaskResultStatus(r.Status)
	if err != nil {
		log.Error("failed to parse task result status", "error", err)
		taskStatus = TaskResultStatus(r.Status)
	}

	return &Task{
		TaskId:             r.TaskId,
		TaskType:           r.TaskType,
		TaskDefName:        r.TaskDefName,
		WorkflowType:       r.WorkflowType,
		ReferenceTaskName:  r.ReferenceTaskName,
		RetryCount:         r.RetryCount,
		Status:             taskStatus,
		WorkflowInstanceId: r.WorkflowId,
		InputData:          r.Input,
		OutputData:         r.Output,
	}, nil
}

// GetTaskInput extracts task input from a SignalResponse
func (r *SignalResponse) GetTaskInput() (map[string]interface{}, error) {
	if r.ResponseType != ReturnBlockingTaskInput {
		return nil, fmt.Errorf("response type %s does not contain task input details", r.ResponseType)
	}

	return r.Input, nil
}

// GetTaskRun extracts task run details from a SignalResponse
func (r *SignalResponse) GetTaskRun() TaskRun {
	taskStatus, err := ParseTaskResultStatus(r.Status)
	if err != nil {
		log.Error("failed to parse task result status", "error", err)
		taskStatus = TaskResultStatus(r.Status)
	}

	// Comprehensive field-by-field mapping from SignalResponse to TaskRun
	return TaskRun{
		// Core task identification fields
		TaskId:            r.TaskId,
		TaskType:          r.TaskType,
		TaskDefName:       r.TaskDefName,
		WorkflowType:      r.WorkflowType,
		WorkflowId:        r.WorkflowId,
		ReferenceTaskName: r.ReferenceTaskName,
		RetryCount:        r.RetryCount,

		// Status and execution fields
		Status:               taskStatus,
		TargetWorkflowStatus: r.TargetWorkflowStatus,

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

// GetWorkflowRun extracts workflow run details from a SignalResponse
func (r *SignalResponse) GetWorkflowRun() WorkflowRun {
	workflowStatus, err := ParseWorkflowStatus(r.Status)
	if err != nil {
		log.Error("failed to parse workflow status", "error", err)
		workflowStatus = WorkflowStatus(r.Status)
	}

	return WorkflowRun{
		WorkflowId:           r.WorkflowId,
		CorrelationId:        r.CorrelationID,
		Priority:             r.Priority,
		Status:               workflowStatus,
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
