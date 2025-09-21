//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package client

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// WorkflowResourceApiService
type WorkflowResourceApiService struct {
	// Embedded for helper methods
	*APIClient
}

// NewWorkflowResourceApiService creates a service from existing APIClient
func NewWorkflowResourceApiService(apiClient *APIClient) *WorkflowResourceApiService {
	if apiClient == nil {
		return nil
	}

	return &WorkflowResourceApiService{
		APIClient: apiClient,
	}
}

// Decide - Starts the decision task for a workflow
func (a *WorkflowResourceApiService) Decide(ctx context.Context, workflowId string) (*http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.Decide(ctx, workflowId)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// WorkflowResourceApiDeleteOpts contains optional parameters for Delete
type WorkflowResourceApiDeleteOpts struct {
	// Deprecated: There is no effect when configured.
	ArchiveWorkflow optional.Bool
}

// Delete - Removes the workflow from the system
func (a *WorkflowResourceApiService) Delete(ctx context.Context, workflowId string, opts *WorkflowResourceApiDeleteOpts) (*http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.Delete1(ctx, workflowId)

	if opts != nil && opts.ArchiveWorkflow.IsSet() {
		req = req.ArchiveWorkflow(opts.ArchiveWorkflow.Value())
	}

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// WorkflowResourceApiGetExecutionStatusOpts contains optional parameters for GetExecutionStatus
type WorkflowResourceApiGetExecutionStatusOpts struct {
	// IncludeTasks if set to true, all task execution details will be fetched in a tasks array
	IncludeTasks optional.Bool
}

// GetExecutionStatus - Gets the workflow by workflow id
func (a *WorkflowResourceApiService) GetExecutionStatus(ctx context.Context, workflowId string, opts *WorkflowResourceApiGetExecutionStatusOpts) (model.Workflow, *http.Response, error) {

	// Use conductor client instead of orkes client to avoid JSON unmarshaling issues
	req := a.http_orkes.WorkflowResourceAPI.GetExecutionStatus(ctx, workflowId)

	if opts != nil && opts.IncludeTasks.IsSet() {
		req = req.IncludeTasks(opts.IncludeTasks.Value())
	}

	genWorkflow, resp, err := req.Execute()
	if err != nil {
		return model.Workflow{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert generated model to our model using mapper
	workflow := toDomainWorkflow(genWorkflow)
	return workflow, resp, nil
}

// WorkflowResourceApiStartWorkflowOpts contains optional parameters for StartWorkflow
type WorkflowResourceApiStartWorkflowOpts struct {
	// Version the workflow version. If unspecified, the latest version will be used.
	Version optional.Int32
	// CorrelationId A unique identifier used to correlate the current workflow execution with other executions of the same workflow.
	CorrelationId optional.String
	// Priority Priority of the workflow execution. Supported values: 0-99.
	// Default is 0, which means workflows are completed in a first-in-first-out order.
	Priority optional.Int32
}

// StartWorkflow - Start a new workflow
func (a *WorkflowResourceApiService) StartWorkflow(ctx context.Context, body map[string]interface{}, name string, opts *WorkflowResourceApiStartWorkflowOpts) (string, *http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.StartWorkflow1(ctx, name).RequestBody(body)

	if opts != nil {
		if opts.Version.IsSet() {
			req = req.Version(opts.Version.Value())
		}
		if opts.CorrelationId.IsSet() {
			req = req.CorrelationId(opts.CorrelationId.Value())
		}
		if opts.Priority.IsSet() {
			req = req.Priority(opts.Priority.Value())
		}
	}

	result, resp, err := req.Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// StartWorkflowWithRequest - Start workflow with StartWorkflowRequest
func (a *WorkflowResourceApiService) StartWorkflowWithRequest(ctx context.Context, body model.StartWorkflowRequest) (string, *http.Response, error) {
	// Convert model to generated model using mapper
	genBody := toGeneratedStartWorkflowRequestForExecute(&body)

	req := a.http_orkes.WorkflowResourceAPI.StartWorkflow(ctx).StartWorkflowRequest(genBody)
	result, resp, err := req.Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// ExecuteWorkflow - Execute workflow synchronously
func (a *WorkflowResourceApiService) ExecuteWorkflow(
	ctx context.Context,
	body model.StartWorkflowRequest,
	requestId string,
	name string,
	version int32,
	waitUntilTask string,
) (model.WorkflowRun, *http.Response, error) {

	// Convert model using mapper
	genBody := toGeneratedStartWorkflowRequestForExecute(&body)

	// Build request
	req := a.http_orkes.WorkflowResourceAPI.ExecuteWorkflow(ctx, name, version).
		StartWorkflowRequest(genBody).
		RequestId(requestId)

	if waitUntilTask != "" {
		req = req.WaitUntilTaskRef(waitUntilTask)
	}

	// Execute
	genResult, resp, err := req.Execute()
	if err != nil {
		return model.WorkflowRun{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert result using mapper
	signalResponse := toDomainSignalResponseFromGenerated(genResult)

	return signalResponse.GetWorkflowRun(), resp, nil
}

// ExecuteWorkflowOpts contains optional parameters for ExecuteWorkflow
type ExecuteWorkflowOpts struct {
	// RequestID a user-generated request ID, which can be used to track the API request.
	RequestID string
	// WaitUntilTaskRef the reference name of the task to wait for before returning a response
	WaitUntilTaskRef []string
	// WaitForSeconds the duration in seconds to wait before returning a response. Default is 10.
	WaitForSeconds int
	// Input the input for the workflow. If unspecified, the workflow will use the input from the StartWorkflowRequest.
	Input map[string]interface{}
	// Specifies how the request persists and is replicated. Supported values: DURABLE, REGION_DURABLE.
	Consistency model.WorkflowConsistency
	// ReturnStrategy this parameter defines the strategy for when the API returns a response.
	// Supported values: TARGET_WORKFLOW, BLOCKING_WORKFLOW, BLOCKING_TASK, BLOCKING_TASK_INPUT.
	ReturnStrategy model.ReturnStrategy
}

// DefaultExecuteWorkflowOpts returns the default options
func DefaultExecuteWorkflowOpts() ExecuteWorkflowOpts {
	return ExecuteWorkflowOpts{
		Consistency:    model.DurableConsistency,   // Default is "DURABLE"
		ReturnStrategy: model.ReturnTargetWorkflow, // Default is TARGET_WORKFLOW
		WaitForSeconds: 10,                         // Default is 10 seconds
	}
}

// ExecuteWorkflowWithReturnStrategy - Execute workflow using orkes generated client
func (a *WorkflowResourceApiService) ExecuteWorkflowWithReturnStrategy(
	ctx context.Context,
	body model.StartWorkflowRequest,
	opts ExecuteWorkflowOpts,
) (*model.SignalResponse, error) {

	signalResponse, _, err := a.executeWorkflowWithReturnStrategy(ctx, body, opts)
	if err != nil {
		return nil, err
	}
	return signalResponse, nil
}

func (a *WorkflowResourceApiService) executeWorkflowWithReturnStrategy(
	ctx context.Context,
	body model.StartWorkflowRequest,
	opts ExecuteWorkflowOpts) (
	*model.SignalResponse, *http.Response, error) {
	// Apply defaults if not specified
	if opts.Consistency == "" {
		opts.Consistency = model.DurableConsistency
	}
	if opts.ReturnStrategy == "" {
		opts.ReturnStrategy = model.ReturnTargetWorkflow
	}
	if opts.WaitForSeconds <= 0 {
		opts.WaitForSeconds = 10
	}

	// Create a new context with timeout
	var cancelFunc context.CancelFunc
	effectiveCtx := ctx
	if opts.WaitForSeconds > 0 {
		// Add buffer time for HTTP overhead
		bufferSeconds := 10
		totalTimeout := time.Duration(opts.WaitForSeconds+bufferSeconds) * time.Second
		effectiveCtx, cancelFunc = context.WithTimeout(ctx, totalTimeout)
		defer cancelFunc()
	}

	// Use orkes generated client
	req := a.http_orkes.WorkflowResourceAPI.ExecuteWorkflow(effectiveCtx, body.Name, int32(body.Version))

	// Set required requestId
	requestId := opts.RequestID
	if requestId == "" {
		// Generate a unique request ID if not provided
		requestId = fmt.Sprintf("%s-%d-%d", body.Name, body.Version, time.Now().UnixNano())
	}
	req = req.RequestId(requestId)

	// Convert domain model to orkes model using mapper
	orkesRequest := toGeneratedStartWorkflowRequestForExecute(&body)
	req = req.StartWorkflowRequest(orkesRequest)

	// Set optional parameters
	if len(opts.WaitUntilTaskRef) > 0 {
		req = req.WaitUntilTaskRef(strings.Join(opts.WaitUntilTaskRef, ","))
	}
	if opts.WaitForSeconds > 0 {
		// Check for overflow before conversion
		if opts.WaitForSeconds <= math.MaxInt32 {
			req = req.WaitForSeconds(int32(opts.WaitForSeconds))
		}
	}

	// Execute the request
	signalResponseGenerated, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert orkes.WorkflowRun to model.SignalResponse using mapper
	signalResponse := toDomainSignalResponseFromGenerated(signalResponseGenerated)

	return &signalResponse, resp, nil
}

// PauseWorkflow - Pauses the workflow
func (a *WorkflowResourceApiService) PauseWorkflow(ctx context.Context, workflowId string) (*http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.PauseWorkflow(ctx, workflowId)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// ResumeWorkflow - Resumes the workflow
func (a *WorkflowResourceApiService) ResumeWorkflow(ctx context.Context, workflowId string) (*http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.ResumeWorkflow(ctx, workflowId)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// WorkflowResourceApiTerminateOpts contains optional parameters for Terminate
type WorkflowResourceApiTerminateOpts struct {
	// Reason a reason for termination.
	Reason optional.String
	// TriggerFailureWorkflow if set to true,  the associated compensation flow (if any) will be triggered.
	TriggerFailureWorkflow optional.Bool
}

// Terminate - Terminate workflow execution
func (a *WorkflowResourceApiService) Terminate(ctx context.Context, workflowId string, opts *WorkflowResourceApiTerminateOpts) (*http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.Terminate1(ctx, workflowId)

	if opts != nil {
		if opts.Reason.IsSet() {
			req = req.Reason(opts.Reason.Value())
		}
		if opts.TriggerFailureWorkflow.IsSet() {
			req = req.TriggerFailureWorkflow(opts.TriggerFailureWorkflow.Value())
		}
	}

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// WorkflowResourceApiRetryOpts contains optional parameters for Retry
type WorkflowResourceApiRetryOpts struct {
	// ResumeSubworkflowTasks If set to true, the parent workflow is restarted from the sub-workflow’s last failed task.
	// If set to false, a new sub-workflow execution is created
	ResumeSubworkflowTasks optional.Bool
	// RetryIfRetriedByParent if set to false, the sub-workflow will be prohibited from retrying if its parent workflow has been retried before
	RetryIfRetriedByParent optional.Bool
}

// Retry - Retries the last failed task
func (a *WorkflowResourceApiService) Retry(ctx context.Context, workflowId string, opts *WorkflowResourceApiRetryOpts) (*http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.Retry(ctx, workflowId)

	if opts != nil && opts.ResumeSubworkflowTasks.IsSet() {
		req = req.ResumeSubworkflowTasks(opts.ResumeSubworkflowTasks.Value())
	}
	if opts != nil && opts.RetryIfRetriedByParent.IsSet() {
		req = req.RetryIfRetriedByParent(opts.RetryIfRetriedByParent.Value())
	}

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// WorkflowResourceApiRestartOpts contains optional parameters for Restart
type WorkflowResourceApiRestartOpts struct {
	// UseLatestDefinitions if set to true, the restarted workflow will use the latest definition from the metadata store.
	UseLatestDefinitions optional.Bool
}

// Restart - Restarts a completed workflow
func (a *WorkflowResourceApiService) Restart(ctx context.Context, workflowId string, opts *WorkflowResourceApiRestartOpts) (*http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.Restart(ctx, workflowId)

	if opts != nil && opts.UseLatestDefinitions.IsSet() {
		req = req.UseLatestDefinitions(opts.UseLatestDefinitions.Value())
	}

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// WorkflowResourceApiSearchOpts contains optional parameters for Search
type WorkflowResourceApiSearchOpts struct {
	// Start starts of the search results list, which is used for pagination. Default is 0.
	Start optional.Int32
	// Size the number of workflows to return. Default is 100.
	Size optional.Int32
	// Sort the field to sort the results by. Format "FIELD:ASC|DESC". For example, "workflowId:DESC".
	Sort optional.String
	// FreeText the free text associated with the workflow execution
	// (workflow input values, workflow output values, workflow variable values, task output values, correlation ID, and reason for incompletion).
	FreeText optional.String
	// Query the query expression in the format FIELD = VALUE or FIELD IN (value1, value2).
	// Supported fields for querying:
	// workflowId, correlationId, workflowType, status, startTime, modifiedTime.
	//
	// Example queries:
	// workflowType = your_workflow_name
	// status IN (PAUSED, RUNNING)
	// startTime >1726655978410
	// startTime < 1696143600000
	// workflowType = your_workflow_name AND status = PAUSED
	// workflowId IN (3434546, 45365767, 20984885) AND workflowType = test_workflow
	Query optional.String
}

// Search - Search for workflows
func (a *WorkflowResourceApiService) Search(ctx context.Context, opts *WorkflowResourceApiSearchOpts) (model.SearchResultWorkflowSummary, *http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.Search1(ctx)

	if opts != nil {
		if opts.Start.IsSet() {
			req = req.Start(opts.Start.Value())
		}
		if opts.Size.IsSet() {
			req = req.Size(opts.Size.Value())
		}
		if opts.Sort.IsSet() {
			req = req.Sort(opts.Sort.Value())
		}
		if opts.FreeText.IsSet() {
			req = req.FreeText(opts.FreeText.Value())
		}
		if opts.Query.IsSet() {
			req = req.Query(opts.Query.Value())
		}
	}

	genResult, resp, err := req.Execute()
	if err != nil {
		return model.SearchResultWorkflowSummary{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert result using mapper
	result := toDomainSearchResultWorkflowSummary(genResult)
	return result, resp, nil
}

// GetWorkflowState - Get workflow state
func (a *WorkflowResourceApiService) GetWorkflowState(ctx context.Context, workflowId string, includeOutput bool, includeVariables bool) (model.WorkflowState, *http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.GetWorkflowStatusSummary(ctx, workflowId).
		IncludeOutput(includeOutput).
		IncludeVariables(includeVariables)

	genResult, resp, err := req.Execute()
	if err != nil {
		return model.WorkflowState{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert WorkflowStatus to WorkflowState using comprehensive mapper
	result := toDomainWorkflowStateFromGenerated(*genResult)
	return result, resp, nil
}

// GetExternalStorageLocation - Get external storage location // TODO: check if this endpoint exists
func (a *WorkflowResourceApiService) GetExternalStorageLocation(ctx context.Context, path string, operation string, payloadType string) (model.ExternalStorageLocation, *http.Response, error) {
	// Use conductor generated client
	req := a.http_conductor.WorkflowResourceAPI.GetExternalStorageLocation1(ctx)
	req = req.Path(path).Operation(operation).PayloadType(payloadType)

	conductorLoc, resp, err := req.Execute()
	if err != nil {
		return model.ExternalStorageLocation{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert conductor model to domain model using mapper
	domainLoc := toDomainExternalStorageLocation(conductorLoc)
	return domainLoc, resp, nil
}

// WorkflowResourceApiGetRunningWorkflowOpts contains optional parameters for GetRunningWorkflow
type WorkflowResourceApiGetRunningWorkflowOpts struct {
	// Version the version of the workflow.
	Version optional.Int32
	// StartTime the start time of the workflow.
	StartTime optional.Int64
	// EndTime the end time of the workflow.
	EndTime optional.Int64
}

// GetRunningWorkflow - Get running workflows
func (a *WorkflowResourceApiService) GetRunningWorkflow(ctx context.Context, name string, opts *WorkflowResourceApiGetRunningWorkflowOpts) ([]string, *http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.GetRunningWorkflow(ctx, name)

	if opts != nil {
		if opts.Version.IsSet() {
			req = req.Version(opts.Version.Value())
		}
		if opts.StartTime.IsSet() {
			req = req.StartTime(opts.StartTime.Value())
		}
		if opts.EndTime.IsSet() {
			req = req.EndTime(opts.EndTime.Value())
		}
	}

	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// GetWorkflows - Get workflows by correlation IDs
func (a *WorkflowResourceApiService) GetWorkflows(ctx context.Context, body []string, name string, opts *WorkflowResourceApiGetWorkflowsOpts) (map[string][]model.Workflow, *http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.GetWorkflows(ctx, name).RequestBody(body)

	if opts != nil {
		if opts.IncludeClosed.IsSet() {
			req = req.IncludeClosed(opts.IncludeClosed.Value())
		}
		if opts.IncludeTasks.IsSet() {
			req = req.IncludeTasks(opts.IncludeTasks.Value())
		}
	}

	genResult, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	if genResult == nil {
		return nil, resp, fmt.Errorf("nil result from GetWorkflows")
	}

	// Convert result - GetWorkflows returns *map[string][]orkes.Workflow
	result := make(map[string][]model.Workflow)

	// The generated API returns a map of workflow arrays, convert each orkes.Workflow to model.Workflow
	for key, workflows := range *genResult {
		converted := make([]model.Workflow, len(workflows))
		for i, w := range workflows {
			converted[i] = toDomainWorkflow(&w)
		}
		result[key] = converted
	}

	return result, resp, nil
}

// GetWorkflowsBatch - Get workflows batch
func (a *WorkflowResourceApiService) GetWorkflowsBatch(ctx context.Context, body map[string][]string, opts *WorkflowResourceApiGetWorkflowsOpts) (map[string][]model.Workflow, *http.Response, error) {
	// Convert body to CorrelationIdsSearchRequest
	genBody := orkes.CorrelationIdsSearchRequest{
		CorrelationIds: body["correlationIds"],
		WorkflowNames:  body["workflowNames"],
	}

	req := a.http_orkes.WorkflowResourceAPI.GetWorkflows1(ctx).CorrelationIdsSearchRequest(genBody)

	if opts != nil {
		if opts.IncludeClosed.IsSet() {
			req = req.IncludeClosed(opts.IncludeClosed.Value())
		}
		if opts.IncludeTasks.IsSet() {
			req = req.IncludeTasks(opts.IncludeTasks.Value())
		}
	}

	genResult, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	if genResult == nil {
		return nil, resp, fmt.Errorf("nil result from GetWorkflowsBatch")
	}

	// Convert result - GetWorkflows returns *map[string][]orkes.Workflow
	result := make(map[string][]model.Workflow)

	// The generated API returns a map of workflow arrays, convert each orkes.Workflow to model.Workflow
	for key, workflows := range *genResult {
		converted := make([]model.Workflow, len(workflows))
		for i, w := range workflows {
			converted[i] = toDomainWorkflow(&w)
		}
		result[key] = converted
	}

	return result, resp, nil
}

// WorkflowResourceApiGetWorkflowsOpts contains optional parameters for GetWorkflows
type WorkflowResourceApiGetWorkflowsOpts struct {
	// IncludeClosed if set to true, the response will also include workflows in a terminal state.
	IncludeClosed optional.Bool
	// IncludeTasks if set to true, all task execution details will be fetched in a tasks array.
	IncludeTasks optional.Bool
}

// GetWorkflowsByCorrelationId - Get workflows by correlation ID
func (a *WorkflowResourceApiService) GetWorkflowsByCorrelationId(ctx context.Context, name string, correlationId string, opts *WorkflowResourceApiGetWorkflowsOpts) ([]model.Workflow, *http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.GetWorkflows2(ctx, name, correlationId)

	if opts != nil {
		if opts.IncludeClosed.IsSet() {
			req = req.IncludeClosed(opts.IncludeClosed.Value())
		}
		if opts.IncludeTasks.IsSet() {
			req = req.IncludeTasks(opts.IncludeTasks.Value())
		}
	}

	genResult, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert result using mapper
	result := make([]model.Workflow, len(genResult))
	for i, gen := range genResult {
		result[i] = toDomainWorkflow(&gen)
	}
	return result, resp, nil
}

// Rerun - Rerun workflow from specific task
func (a *WorkflowResourceApiService) Rerun(ctx context.Context, body model.RerunWorkflowRequest, workflowId string) (string, *http.Response, error) {
	// Convert model using mapper
	genBody := toGeneratedRerunWorkflowRequest(&body)

	req := a.http_orkes.WorkflowResourceAPI.Rerun(ctx, workflowId).RerunWorkflowRequest(genBody)
	result, resp, err := req.Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// ResetWorkflow - Reset callback times
func (a *WorkflowResourceApiService) ResetWorkflow(ctx context.Context, workflowId string) (*http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.ResetWorkflow(ctx, workflowId)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// WorkflowResourceApiSearchWorkflowsByTasksOpts contains optional parameters for SearchWorkflowsByTasks
type WorkflowResourceApiSearchWorkflowsByTasksOpts struct {
	// Start starts of the search results list, which is used for pagination. Default is 0.
	Start optional.Int32
	// Size the number of workflows to return. Default is 100.
	Size optional.Int32
	// Sort the field to sort the results by.
	Sort optional.String
	// FreeText the free text associated with the workflow execution
	// (workflow input values, workflow output values, workflow variable values, task output values, correlation ID, and reason for incompletion).
	FreeText optional.String
	// Query the query expression in the format FIELD = VALUE or FIELD IN (value1, value2).
	// Supported fields for querying:
	// workflowId, correlationId, workflowType, status, startTime, modifiedTime.
	Query optional.String
}

// SearchWorkflowsByTasks - Search workflows by tasks
func (a *WorkflowResourceApiService) SearchWorkflowsByTasks(ctx context.Context, opts *WorkflowResourceApiSearchWorkflowsByTasksOpts) (model.SearchResultWorkflowSummary, *http.Response, error) {
	req := a.http_conductor.WorkflowResourceAPI.SearchWorkflowsByTasks(ctx)

	if opts != nil {
		req = req.Start(opts.Start.Value())
		req = req.Size(opts.Size.Value())
		req = req.Sort(opts.Sort.Value())
		req = req.FreeText(opts.FreeText.Value())
		req = req.Query(opts.Query.Value())
	}

	genResult, resp, err := req.Execute()
	if err != nil {
		return model.SearchResultWorkflowSummary{}, resp, wrapGeneratedError(err, resp)
	}

	result := toDomainSearchResultWorkflowSummaryFromConductorGenerated(genResult)
	return result, resp, nil
}

// WorkflowResourceApiSearchWorkflowsByTasksV2Opts contains optional parameters for SearchWorkflowsByTasksV2.
type WorkflowResourceApiSearchWorkflowsByTasksV2Opts struct {
	// Start starts of the search results list, which is used for pagination. Default is 0.
	Start optional.Int32
	// Size the number of workflows to return. Default is 100.
	Size optional.Int32
	// Sort the field to sort the results by.
	Sort optional.String
	// FreeText the free text associated with the workflow execution
	// (workflow input values, workflow output values, workflow variable values, task output values, correlation ID, and reason for incompletion).
	FreeText optional.String
	// Query the query expression in the format FIELD = VALUE or FIELD IN (value1, value2).
	// Supported fields for querying:
	// workflowId, correlationId, workflowType, status, startTime, modifiedTime.
	Query optional.String
}

// SearchWorkflowsByTasksV2 - Search workflows by tasks V2
func (a *WorkflowResourceApiService) SearchWorkflowsByTasksV2(ctx context.Context, opts *WorkflowResourceApiSearchWorkflowsByTasksV2Opts) (model.SearchResultWorkflow, *http.Response, error) {
	req := a.http_conductor.WorkflowResourceAPI.SearchWorkflowsByTasksV2(ctx)

	if opts != nil {
		req = req.Start(opts.Start.Value())
		req = req.Size(opts.Size.Value())
		req = req.Sort(opts.Sort.Value())
		req = req.FreeText(opts.FreeText.Value())
		req = req.Query(opts.Query.Value())
	}

	genResult, resp, err := req.Execute()
	if err != nil {
		return model.SearchResultWorkflow{}, resp, wrapGeneratedError(err, resp)
	}

	result := toDomainSearchResultWorkflowFromConductorGenerated(genResult)
	return result, resp, nil
}

// SkipTaskFromWorkflow - Skip task from workflow
func (a *WorkflowResourceApiService) SkipTaskFromWorkflow(ctx context.Context, workflowId string, taskReferenceName string, skipTaskRequest model.SkipTaskRequest) (*http.Response, error) {
	// Convert model using mapper
	genBody := toGeneratedSkipTaskRequest(&skipTaskRequest)

	req := a.http_orkes.WorkflowResourceAPI.SkipTaskFromWorkflow(ctx, workflowId, taskReferenceName).SkipTaskRequest(genBody)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// ExecuteAndGetTarget - Execute and get target workflow
func (a *WorkflowResourceApiService) ExecuteAndGetTarget(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask []string, waitForSeconds int, consistency string) (model.WorkflowRun, *http.Response, error) {
	// Use ExecuteWorkflowWithReturnStrategy with TARGET_WORKFLOW
	opts := ExecuteWorkflowOpts{
		RequestID:        requestId,
		WaitUntilTaskRef: waitUntilTask,
		WaitForSeconds:   waitForSeconds,
		Consistency:      model.WorkflowConsistency(consistency),
		ReturnStrategy:   model.ReturnTargetWorkflow,
	}

	signal, resp, err := a.executeWorkflowWithReturnStrategy(ctx, body, opts)
	if err != nil {
		return model.WorkflowRun{}, nil, err
	}

	return signal.GetWorkflowRun(), resp, nil
}

// ExecuteAndGetBlockingWorkflow - Execute and get blocking workflow
func (a *WorkflowResourceApiService) ExecuteAndGetBlockingWorkflow(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask []string, waitForSeconds int, consistency string) (model.WorkflowRun, *http.Response, error) {
	// Use ExecuteWorkflowWithReturnStrategy with BLOCKING_WORKFLOW
	opts := ExecuteWorkflowOpts{
		RequestID:        requestId,
		WaitUntilTaskRef: waitUntilTask,
		WaitForSeconds:   waitForSeconds,
		Consistency:      model.WorkflowConsistency(consistency),
		ReturnStrategy:   model.ReturnBlockingWorkflow,
	}

	signal, resp, err := a.executeWorkflowWithReturnStrategy(ctx, body, opts)
	if err != nil {
		return model.WorkflowRun{}, nil, err
	}

	return signal.GetWorkflowRun(), resp, nil
}

// ExecuteAndGetBlockingTask - Execute and get blocking task
func (a *WorkflowResourceApiService) ExecuteAndGetBlockingTask(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask []string, waitForSeconds int, consistency string) (model.TaskRun, *http.Response, error) {
	// Use ExecuteWorkflowWithReturnStrategy with BLOCKING_TASK
	opts := ExecuteWorkflowOpts{
		RequestID:        requestId,
		WaitUntilTaskRef: waitUntilTask,
		WaitForSeconds:   waitForSeconds,
		Consistency:      model.WorkflowConsistency(consistency),
		ReturnStrategy:   model.ReturnBlockingTask,
	}

	signal, resp, err := a.executeWorkflowWithReturnStrategy(ctx, body, opts)
	if err != nil {
		return model.TaskRun{}, nil, err
	}

	return signal.GetTaskRun(), resp, nil
}

// ExecuteAndGetBlockingTaskInput - Execute and get blocking task input
func (a *WorkflowResourceApiService) ExecuteAndGetBlockingTaskInput(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask []string, waitForSeconds int, consistency string) (model.TaskRun, *http.Response, error) {
	// Use ExecuteWorkflowWithReturnStrategy with BLOCKING_TASK_INPUT
	opts := ExecuteWorkflowOpts{
		RequestID:        requestId,
		WaitUntilTaskRef: waitUntilTask,
		WaitForSeconds:   waitForSeconds,
		Consistency:      model.WorkflowConsistency(consistency),
		ReturnStrategy:   model.ReturnBlockingTaskInput,
	}

	signal, resp, err := a.executeWorkflowWithReturnStrategy(ctx, body, opts)
	if err != nil {
		return model.TaskRun{}, nil, err
	}

	return signal.GetTaskRun(), resp, nil
}

// WorkflowResourceApiJumpToTaskOpts contains optional parameters for JumpToTask
type WorkflowResourceApiJumpToTaskOpts struct {
	// TaskReferenceName the reference name of the task to jump to.
	TaskReferenceName optional.String
}

// JumpToTask jumps to a specific task in a running workflow.
func (a *WorkflowResourceApiService) JumpToTask(ctx context.Context, body map[string]interface{}, workflowId string, optionals *WorkflowResourceApiJumpToTaskOpts) (*http.Response, error) {
	// Build request with required body and path parameter
	req := a.http_orkes.WorkflowResourceAPI.
		JumpToTask(ctx, workflowId).
		RequestBody(body)

	if optionals != nil && optionals.TaskReferenceName.IsSet() {
		req = req.TaskReferenceName(optionals.TaskReferenceName.Value())
	}

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// WorkflowResourceApiUpdateWorkflowAndTaskStateOpts contains optional parameters for UpdateWorkflowAndTaskState
type WorkflowResourceApiUpdateWorkflowAndTaskStateOpts struct {
	// WaitUntilTaskRef the reference name of the task to wait for before returning a response.
	WaitUntilTaskRef optional.String
	// WaitForSeconds the duration in seconds to wait before returning a response.
	WaitForSeconds optional.Int32
}

// UpdateWorkflowAndTaskState - Update workflow and task state
func (a *WorkflowResourceApiService) UpdateWorkflowAndTaskState(ctx context.Context, body model.WorkflowStateUpdate, requestId string, workflowId string, optionals *WorkflowResourceApiUpdateWorkflowAndTaskStateOpts) (model.WorkflowRun, *http.Response, error) {
	genBody := toGeneratedWorkflowStateUpdate(&body)

	req := a.http_orkes.WorkflowResourceAPI.UpdateWorkflowAndTaskState(ctx, workflowId)
	req = req.WorkflowStateUpdate(genBody)
	req = req.RequestId(requestId)

	// Handle optional parameters
	if optionals != nil && optionals.WaitUntilTaskRef.IsSet() {
		req = req.WaitUntilTaskRef(optionals.WaitUntilTaskRef.Value())
	}
	if optionals != nil && optionals.WaitForSeconds.IsSet() {
		req = req.WaitForSeconds(optionals.WaitForSeconds.Value())
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return model.WorkflowRun{}, httpResp, wrapGeneratedError(err, httpResp)
	}

	// Convert the response to domain model
	domainResp := toDomainWorkflowRun(resp)
	return domainResp, httpResp, nil
}

// UpgradeRunningWorkflowToVersion - Upgrade running workflow to newer version
func (a *WorkflowResourceApiService) UpgradeRunningWorkflowToVersion(ctx context.Context, body model.UpgradeWorkflowRequest, workflowID string) (*http.Response, error) {
	genBody := toGeneratedUpgradeWorkflowRequest(&body)

	resp, err := a.http_orkes.WorkflowResourceAPI.UpgradeRunningWorkflowToVersion(ctx, workflowID).UpgradeWorkflowRequest(genBody).Execute()
	if err != nil {
		return nil, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// TestWorkflow tests workflow execution using mock data
func (a *WorkflowResourceApiService) TestWorkflow(ctx context.Context, body model.WorkflowTestRequest) (model.Workflow, *http.Response, error) {
	genBody := toGeneratedWorkflowTestRequest(&body)

	resp, httpResp, err := a.http_orkes.WorkflowResourceAPI.TestWorkflow(ctx).WorkflowTestRequest(genBody).Execute()
	if err != nil {
		return model.Workflow{}, httpResp, wrapGeneratedError(err, httpResp)
	}
	return toDomainWorkflow(resp), httpResp, nil
}

// WorkflowResourceAPIGetExecutionStatusTaskListOpts contains optional parameters for GetExecutionStatusTaskList
type WorkflowResourceAPIGetExecutionStatusTaskListOpts struct {
	// Start the start of the task list.
	Start optional.Int32
	// Count the count of the task list.
	Count optional.Int32
	// Status the status of the task list.
	Status optional.Interface
}

// GetExecutionStatusTaskList - Get execution status task list
func (a *WorkflowResourceApiService) GetExecutionStatusTaskList(ctx context.Context, workflowID string, opts *WorkflowResourceAPIGetExecutionStatusTaskListOpts) (model.TaskListSearchResultSummary, *http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.GetExecutionStatusTaskList(ctx, workflowID)

	if opts != nil && opts.Start.IsSet() {
		req = req.Start(opts.Start.Value())
	}
	if opts != nil && opts.Count.IsSet() {
		req = req.Count(opts.Count.Value())
	}

	if opts != nil && opts.Status.IsSet() {
		// Accept string or []string and convert to []string for the generated client
		switch v := opts.Status.Value().(type) {
		case []string:
			req = req.Status(v)
		case string:
			req = req.Status([]string{v})
		case []interface{}:
			var ss []string
			for _, it := range v {
				if s, ok := it.(string); ok {
					ss = append(ss, s)
				}
			}
			if len(ss) > 0 {
				req = req.Status(ss)
			}
		default:
			// ignore unsupported types
		}
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return model.TaskListSearchResultSummary{}, nil, wrapGeneratedError(err, httpResp)
	}

	return toDomainTaskListSearchResultSummaryFromConductorGenerated(resp), httpResp, nil
}

// UpdateWorkflowStateUpdates updates workflow variables for a running workflow.
//
// This method is similar to the Set Variable task, except the variables can be updated anytime in real time.
func (a *WorkflowResourceApiService) UpdateWorkflowState(ctx context.Context, body map[string]interface{}, workflowID string) (model.Workflow, *http.Response, error) {
	req := a.http_orkes.WorkflowResourceAPI.UpdateWorkflowState(ctx, workflowID)
	req = req.RequestBody(body)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return model.Workflow{}, nil, wrapGeneratedError(err, httpResp)
	}

	return toDomainWorkflow(resp), httpResp, nil
}
