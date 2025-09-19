package client

import (
	"context"
	"net/http"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

type WorkflowClient interface {
	Decide(ctx context.Context, workflowId string) (*http.Response, error)
	Delete(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiDeleteOpts) (*http.Response, error)
	GetExecutionStatus(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiGetExecutionStatusOpts) (model.Workflow, *http.Response, error)
	GetWorkflowState(ctx context.Context, workflowId string, includeOutput bool, includeVariables bool) (model.WorkflowState, *http.Response, error)
	GetExternalStorageLocation(ctx context.Context, path string, operation string, payloadType string) (model.ExternalStorageLocation, *http.Response, error)
	GetRunningWorkflow(ctx context.Context, name string, localVarOptionals *WorkflowResourceApiGetRunningWorkflowOpts) ([]string, *http.Response, error)
	GetWorkflows(ctx context.Context, body []string, name string, localVarOptionals *WorkflowResourceApiGetWorkflowsOpts) (map[string][]model.Workflow, *http.Response, error)
	GetWorkflowsBatch(ctx context.Context, body map[string][]string, localVarOptionals *WorkflowResourceApiGetWorkflowsOpts) (map[string][]model.Workflow, *http.Response, error)
	GetWorkflowsByCorrelationId(ctx context.Context, name string, correlationId string, localVarOptionals *WorkflowResourceApiGetWorkflowsOpts) ([]model.Workflow, *http.Response, error)
	PauseWorkflow(ctx context.Context, workflowId string) (*http.Response, error)
	Rerun(ctx context.Context, body model.RerunWorkflowRequest, workflowId string) (string, *http.Response, error)
	ResetWorkflow(ctx context.Context, workflowId string) (*http.Response, error)
	Restart(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiRestartOpts) (*http.Response, error)
	ResumeWorkflow(ctx context.Context, workflowId string) (*http.Response, error)
	Retry(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiRetryOpts) (*http.Response, error)
	Search(ctx context.Context, localVarOptionals *WorkflowResourceApiSearchOpts) (model.SearchResultWorkflowSummary, *http.Response, error)
	SearchWorkflowsByTasks(ctx context.Context, localVarOptionals *WorkflowResourceApiSearchWorkflowsByTasksOpts) (model.SearchResultWorkflowSummary, *http.Response, error)
	SearchWorkflowsByTasksV2(ctx context.Context, localVarOptionals *WorkflowResourceApiSearchWorkflowsByTasksV2Opts) (model.SearchResultWorkflow, *http.Response, error)
	SkipTaskFromWorkflow(ctx context.Context, workflowId string, taskReferenceName string, skipTaskRequest model.SkipTaskRequest) (*http.Response, error)
	StartWorkflow(ctx context.Context, body map[string]interface{}, name string, localVarOptionals *WorkflowResourceApiStartWorkflowOpts) (string, *http.Response, error)
	ExecuteWorkflow(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask string) (model.WorkflowRun, *http.Response, error)
	StartWorkflowWithRequest(ctx context.Context, body model.StartWorkflowRequest) (string, *http.Response, error)
	ExecuteWorkflowWithReturnStrategy(ctx context.Context, body model.StartWorkflowRequest, opts ExecuteWorkflowOpts) (*model.SignalResponse, error)
	ExecuteAndGetTarget(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask []string, waitForSeconds int, consistency string) (model.WorkflowRun, *http.Response, error)
	ExecuteAndGetBlockingWorkflow(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask []string, waitForSeconds int, consistency string) (model.WorkflowRun, *http.Response, error)
	ExecuteAndGetBlockingTask(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask []string, waitForSeconds int, consistency string) (model.TaskRun, *http.Response, error)
	ExecuteAndGetBlockingTaskInput(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask []string, waitForSeconds int, consistency string) (model.TaskRun, *http.Response, error)
	JumpToTask(ctx context.Context, body map[string]interface{}, workflowID string, optionals *WorkflowResourceApiJumpToTaskOpts) (*http.Response, error)
	UpdateWorkflowAndTaskState(ctx context.Context, body model.WorkflowStateUpdate, requestID string, workflowID string, optionals *WorkflowResourceApiUpdateWorkflowAndTaskStateOpts) (model.WorkflowRun, *http.Response, error)
	UpgradeRunningWorkflowToVersion(ctx context.Context, body model.UpgradeWorkflowRequest, workflowID string) (*http.Response, error)
	TestWorkflow(ctx context.Context, body model.WorkflowTestRequest) (model.Workflow, *http.Response, error)
	GetExecutionStatusTaskList(ctx context.Context, workflowID string, opts *WorkflowResourceAPIGetExecutionStatusTaskListOpts) (model.TaskListSearchResultSummary, *http.Response, error)
	UpdateWorkflowState(ctx context.Context, body map[string]interface{}, workflowID string) (model.Workflow, *http.Response, error)
	Terminate(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiTerminateOpts) (*http.Response, error)
}

func NewWorkflowClient(apiClient *APIClient) WorkflowClient {
	return NewWorkflowResourceApiService(apiClient)
}
