package client

import (
	"context"
	"net/http"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// WorkflowClient is the client for the workflow API
type WorkflowClient interface {
	// Decide starts the decision task for a workflow.
	Decide(ctx context.Context, workflowId string) (*http.Response, error)
	// Delete deletes a workflow execution from the server using the workflow ID.
	Delete(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiDeleteOpts) (*http.Response, error)
	// GetExecutionStatus gets a workflow’s execution details using its workflow ID.
	GetExecutionStatus(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiGetExecutionStatusOpts) (model.Workflow, *http.Response, error)
	// GetWorkflowState gets a workflow’s state using its workflow ID.
	GetWorkflowState(ctx context.Context, workflowId string, includeOutput bool, includeVariables bool) (model.WorkflowState, *http.Response, error)
	// GetExternalStorageLocation gets the external storage location for the given path, operation, and payload type.
	GetExternalStorageLocation(ctx context.Context, path string, operation string, payloadType string) (model.ExternalStorageLocation, *http.Response, error)
	// GetRunningWorkflow gets a list of ids running workflows for the given name.
	GetRunningWorkflow(ctx context.Context, name string, localVarOptionals *WorkflowResourceApiGetRunningWorkflowOpts) ([]string, *http.Response, error)
	// GetWorkflows gets a list of workflows for the given name and correlation IDs.
	GetWorkflows(ctx context.Context, body []string, name string, localVarOptionals *WorkflowResourceApiGetWorkflowsOpts) (map[string][]model.Workflow, *http.Response, error)
	// GetWorkflowsBatch gets a list of workflows for the given names.
	GetWorkflowsBatch(ctx context.Context, body map[string][]string, localVarOptionals *WorkflowResourceApiGetWorkflowsOpts) (map[string][]model.Workflow, *http.Response, error)
	// GetWorkflowsByCorrelationId gets a list of execution details for a specific workflow definition based on a list of correlation IDs.
	GetWorkflowsByCorrelationId(ctx context.Context, name string, correlationId string, localVarOptionals *WorkflowResourceApiGetWorkflowsOpts) ([]model.Workflow, *http.Response, error)
	// Deprecated: Use Pause instead.
	PauseWorkflow(ctx context.Context, workflowId string) (*http.Response, error)
	// Pause pauses an ongoing workflow execution.
	// Any currently running tasks will be completed, but no new tasks will be scheduled until the workflow is resumed.
	Pause(ctx context.Context, workflowId string) (*http.Response, error)
	// Rerun reruns an ongoing or terminal workflow from a specific reRunFromTaskId task,
	// or reruns a terminal workflow execution from the start, with the option to supply updated inputs.
	Rerun(ctx context.Context, body model.RerunWorkflowRequest, workflowId string) (string, *http.Response, error)
	// Deprecated: Use Reset instead.
	ResetWorkflow(ctx context.Context, workflowId string) (*http.Response, error)
	// Reset resets the callback times of all IN_PROGRESS tasks to 0 for the given workflow
	Reset(ctx context.Context, workflowId string) (*http.Response, error)
	// Restart restarts a terminal workflow execution from the beginning using the same input and workflow definition.
	//
	// You can optionally restart it using the latest workflow definition. This method has no effect on ongoing workflows.
	Restart(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiRestartOpts) (*http.Response, error)
	// Deprecated: Use Resume instead.
	ResumeWorkflow(ctx context.Context, workflowId string) (*http.Response, error)
	// Resume resumes a paused workflow execution.
	//
	// This method has no effect if the workflow is not paused.
	Resume(ctx context.Context, workflowId string) (*http.Response, error)
	// Retries a failed workflow execution from the last failed task.
	//
	// When invoked, the failed task is scheduled again, and the workflow moves to RUNNING status.
	Retry(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiRetryOpts) (*http.Response, error)
	// Searches retrieves a list of workflow executions based on the provided search criteria.
	Search(ctx context.Context, localVarOptionals *WorkflowResourceApiSearchOpts) (model.SearchResultWorkflowSummary, *http.Response, error)
	// SearchWorkflowsByTasks searches for workflows based on task parameters.
	SearchWorkflowsByTasks(ctx context.Context, localVarOptionals *WorkflowResourceApiSearchWorkflowsByTasksOpts) (model.SearchResultWorkflowSummary, *http.Response, error)
	// SearchWorkflowsByTasksV2 searches for workflows based on task parameters.
	SearchWorkflowsByTasksV2(ctx context.Context, localVarOptionals *WorkflowResourceApiSearchWorkflowsByTasksV2Opts) (model.SearchResultWorkflow, *http.Response, error)
	// SkipTaskFromWorkflow skips a scheduled task in an ongoing workflow using the task reference name.
	//
	// The skipped task’s inputs and outputs can be updated using the query parameters.
	SkipTaskFromWorkflow(ctx context.Context, workflowId string, taskReferenceName string, skipTaskRequest model.SkipTaskRequest) (*http.Response, error)
	// StartWorkflow starts a workflow execution asynchronously with workflow name and input.
	//
	// This method returns immediately with the workflow ID without waiting for the workflow to be completed.
	StartWorkflow(ctx context.Context, body map[string]interface{}, name string, localVarOptionals *WorkflowResourceApiStartWorkflowOpts) (string, *http.Response, error)
	// ExecuteWorkflow starts a workflow execution synchronously.
	ExecuteWorkflow(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask string) (model.WorkflowRun, *http.Response, error)
	// StartWorkflowWithRequest starts a workflow execution asynchronously with a StartWorkflowRequest.
	//
	// This method returns immediately with the workflow ID without waiting for the workflow to be completed.
	StartWorkflowWithRequest(ctx context.Context, body model.StartWorkflowRequest) (string, *http.Response, error)
	// ExecuteWorkflowWithReturnStrategy starts a workflow execution synchronously with a return strategy.
	//
	// Set consistency to SYNCHRONOUS to execute the workflow directly from memory without persisting the request.
	ExecuteWorkflowWithReturnStrategy(ctx context.Context, body model.StartWorkflowRequest, opts ExecuteWorkflowOpts) (*model.SignalResponse, error)
	// ExecuteAndGetTarget starts a workflow execution synchronously and uses TARGET_WORKFLOW return strategy.
	ExecuteAndGetTarget(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask []string, waitForSeconds int, consistency string) (model.WorkflowRun, *http.Response, error)
	// ExecuteAndGetBlockingWorkflow starts a workflow execution synchronously and uses BLOCKING_WORKFLOW return strategy.
	ExecuteAndGetBlockingWorkflow(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask []string, waitForSeconds int, consistency string) (model.WorkflowRun, *http.Response, error)
	// ExecuteAndGetBlockingTask starts a workflow execution synchronously and uses BLOCKING_TASK return strategy.
	ExecuteAndGetBlockingTask(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask []string, waitForSeconds int, consistency string) (model.TaskRun, *http.Response, error)
	// ExecuteAndGetBlockingTaskInput starts a workflow execution synchronously and uses BLOCKING_TASK_INPUT return strategy.
	ExecuteAndGetBlockingTaskInput(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask []string, waitForSeconds int, consistency string) (model.TaskRun, *http.Response, error)
	// Terminate terminates a running workflow, with the option to provide a reason for termination and/or trigger a failure workflow.
	Terminate(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiTerminateOpts) (*http.Response, error)
	// JumpToTask jumps to a specific task in a running workflow.
	JumpToTask(ctx context.Context, body map[string]interface{}, workflowID string, optionals *WorkflowResourceApiJumpToTaskOpts) (*http.Response, error)
	// UpdateWorkflowAndTaskState updates the state of a workflow and its tasks.
	UpdateWorkflowAndTaskState(ctx context.Context, body model.WorkflowStateUpdate, requestID string, workflowID string, optionals *WorkflowResourceApiUpdateWorkflowAndTaskStateOpts) (model.WorkflowRun, *http.Response, error)
	// UpgradeRunningWorkflowToVersion upgrades a running workflow to a different version.
	//
	// The workflow execution will continue from its last running task, even when it has been upgraded.
	// In other words, all the tasks in the upgraded definition prior to the currently running task will be marked as skipped.
	UpgradeRunningWorkflowToVersion(ctx context.Context, body model.UpgradeWorkflowRequest, workflowID string) (*http.Response, error)
	// TestWorkflow tests a workflow execution using mock data.
	TestWorkflow(ctx context.Context, body model.WorkflowTestRequest) (model.Workflow, *http.Response, error)
	// GetExecutionStatusTaskList gets the workflow tasks by workflow id.
	GetExecutionStatusTaskList(ctx context.Context, workflowID string, opts *WorkflowResourceAPIGetExecutionStatusTaskListOpts) (model.TaskListSearchResultSummary, *http.Response, error)
	// UpdateWorkflowStateUpdates updates workflow variables for a running workflow.
	//
	// This method is similar to the Set Variable task, except the variables can be updated anytime in real time.
	UpdateWorkflowState(ctx context.Context, body map[string]interface{}, workflowID string) (model.Workflow, *http.Response, error)
}

func NewWorkflowClient(apiClient *APIClient) WorkflowClient {
	return &WorkflowResourceApiService{apiClient}
}
