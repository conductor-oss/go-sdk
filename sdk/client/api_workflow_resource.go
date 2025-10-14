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
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// Linger please
var (
	_ context.Context
)

type WorkflowResourceApiService struct {
	*APIClient
}

// Decide starts the decision task for a workflow.
func (a *WorkflowResourceApiService) Decide(ctx context.Context, workflowId string) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/decide/%s", workflowId)

	resp, err := a.Put(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// WorkflowResourceApiDeleteOpts contains optional parameters for Delete
type WorkflowResourceApiDeleteOpts struct {
	// Deprecated: There is no effect when configured.
	ArchiveWorkflow optional.Bool
}

// Delete deletes the workflow from the system
func (a *WorkflowResourceApiService) Delete(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiDeleteOpts) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/remove", workflowId)

	queryParams := url.Values{}
	if localVarOptionals != nil && localVarOptionals.ArchiveWorkflow.IsSet() {
		queryParams.Add("archiveWorkflow", parameterToString(localVarOptionals.ArchiveWorkflow.Value(), ""))
	}

	resp, err := a.APIClient.Delete(ctx, path, queryParams, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// WorkflowResourceApiGetExecutionStatusOpts contains optional parameters for GetExecutionStatus
type WorkflowResourceApiGetExecutionStatusOpts struct {
	// IncludeTasks if set to true, all task execution details will be fetched in a tasks array
	IncludeTasks optional.Bool
}

// GetExecutionStatus gets the workflow by workflow id
func (a *WorkflowResourceApiService) GetExecutionStatus(ctx context.Context, workflowId string, opts *WorkflowResourceApiGetExecutionStatusOpts) (model.Workflow, *http.Response, error) {
	var result model.Workflow

	path := fmt.Sprintf("/workflow/%s", workflowId)

	queryParams := url.Values{}
	if opts != nil && opts.IncludeTasks.IsSet() {
		queryParams.Add("includeTasks", parameterToString(opts.IncludeTasks.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.Workflow{}, resp, err
	}

	return result, resp, nil
}

// GetWorkflowState gets the workflow state
func (a *WorkflowResourceApiService) GetWorkflowState(ctx context.Context, workflowId string, includeOutput bool, includeVariables bool) (model.WorkflowState, *http.Response, error) {
	var result model.WorkflowState

	path := fmt.Sprintf("/workflow/%s/status", workflowId)

	queryParams := url.Values{}
	queryParams.Add("includeOutput", parameterToString(includeOutput, ""))
	queryParams.Add("includeVariables", parameterToString(includeVariables, ""))

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.WorkflowState{}, resp, err
	}
	return result, resp, nil
}

// GetExternalStorageLocation gets external storage location.
func (a *WorkflowResourceApiService) GetExternalStorageLocation(ctx context.Context, path string, operation string, payloadType string) (model.ExternalStorageLocation, *http.Response, error) {
	var result model.ExternalStorageLocation

	urlPath := "/workflow/externalstoragelocation"

	queryParams := url.Values{}
	queryParams.Add("path", parameterToString(path, ""))
	queryParams.Add("operation", parameterToString(operation, ""))
	queryParams.Add("payloadType", parameterToString(payloadType, ""))

	resp, err := a.Get(ctx, urlPath, queryParams, &result)
	if err != nil {
		return model.ExternalStorageLocation{}, resp, err
	}
	return result, resp, nil
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

// GetRunningWorkflow gets running workflows.
func (a *WorkflowResourceApiService) GetRunningWorkflow(ctx context.Context, name string, opts *WorkflowResourceApiGetRunningWorkflowOpts) ([]string, *http.Response, error) {
	var result []string

	path := fmt.Sprintf("/workflow/running/%s", name)

	queryParams := url.Values{}
	if opts != nil && opts.Version.IsSet() {
		queryParams.Add("version", parameterToString(opts.Version.Value(), ""))
	}
	if opts != nil && opts.StartTime.IsSet() {
		queryParams.Add("startTime", parameterToString(opts.StartTime.Value(), ""))
	}
	if opts != nil && opts.EndTime.IsSet() {
		queryParams.Add("endTime", parameterToString(opts.EndTime.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetWorkflows gets workflows by correlation IDs.
func (a *WorkflowResourceApiService) GetWorkflows(ctx context.Context, body []string, name string, opts *WorkflowResourceApiGetWorkflowsOpts) (map[string][]model.Workflow, *http.Response, error) {
	var result map[string][]model.Workflow

	path := fmt.Sprintf("/workflow/%s/correlated", name)

	queryParams := url.Values{}
	if opts != nil && opts.IncludeClosed.IsSet() {
		queryParams.Add("includeClosed", parameterToString(opts.IncludeClosed.Value(), ""))
	}
	if opts != nil && opts.IncludeTasks.IsSet() {
		queryParams.Add("includeTasks", parameterToString(opts.IncludeTasks.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetWorkflowsBatch gets workflows by correlation IDs.
func (a *WorkflowResourceApiService) GetWorkflowsBatch(ctx context.Context, body map[string][]string, localVarOptionals *WorkflowResourceApiGetWorkflowsOpts) (map[string][]model.Workflow, *http.Response, error) {
	var result map[string][]model.Workflow

	path := "/workflow/correlated/batch"

	queryParams := url.Values{}
	if localVarOptionals != nil && localVarOptionals.IncludeClosed.IsSet() {
		queryParams.Add("includeClosed", parameterToString(localVarOptionals.IncludeClosed.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeTasks.IsSet() {
		queryParams.Add("includeTasks", parameterToString(localVarOptionals.IncludeTasks.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, &result)
	if err != nil {
		return nil, resp, err
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

// GetWorkflowsByCorrelationId gets workflows by correlation ID.
func (a *WorkflowResourceApiService) GetWorkflowsByCorrelationId(ctx context.Context, name string, correlationId string, opts *WorkflowResourceApiGetWorkflowsOpts) ([]model.Workflow, *http.Response, error) {
	return a.GetWorkflows1(ctx, name, correlationId, opts)
}

// Deprecated: Use GetWorkflowsByCorrelationId instead.
func (a *WorkflowResourceApiService) GetWorkflows1(ctx context.Context, name string, correlationId string, opts *WorkflowResourceApiGetWorkflowsOpts) ([]model.Workflow, *http.Response, error) {
	var result []model.Workflow

	localVarPath := fmt.Sprintf("/workflow/%s/correlated/%s", name, correlationId)

	queryParams := url.Values{}
	if opts != nil && opts.IncludeClosed.IsSet() {
		queryParams.Add("includeClosed", parameterToString(opts.IncludeClosed.Value(), ""))
	}
	if opts != nil && opts.IncludeTasks.IsSet() {
		queryParams.Add("includeTasks", parameterToString(opts.IncludeTasks.Value(), ""))
	}

	resp, err := a.Get(ctx, localVarPath, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// Deprecated: Use Pause instead.
func (a *WorkflowResourceApiService) PauseWorkflow(ctx context.Context, workflowId string) (*http.Response, error) {
	return a.Pause(ctx, workflowId)
}

// Pause pauses an ongoing workflow execution.
func (a *WorkflowResourceApiService) Pause(ctx context.Context, workflowId string) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/pause", workflowId)

	resp, err := a.Put(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Rerun reruns the workflow from a specific task.
func (a *WorkflowResourceApiService) Rerun(ctx context.Context, body model.RerunWorkflowRequest, workflowId string) (string, *http.Response, error) {
	var result string

	path := fmt.Sprintf("/workflow/%s/rerun", workflowId)

	resp, err := a.Post(ctx, path, body, &result)
	if err != nil {
		return "", resp, err
	}
	return result, resp, nil
}

// Deprecated: Use Reset instead.
func (a *WorkflowResourceApiService) ResetWorkflow(ctx context.Context, workflowId string) (*http.Response, error) {
	return a.Reset(ctx, workflowId)
}

// Reset resets the callback times of all IN_PROGRESS tasks to 0 for the given workflow.
func (a *WorkflowResourceApiService) Reset(ctx context.Context, workflowId string) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/resetcallbacks", workflowId)

	resp, err := a.Post(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// WorkflowResourceApiRestartOpts contains optional parameters for Restart
type WorkflowResourceApiRestartOpts struct {
	// UseLatestDefinitions if set to true, the restarted workflow will use the latest definition from the metadata store.
	UseLatestDefinitions optional.Bool
}

// Restart restarts a completed workflow.
func (a *WorkflowResourceApiService) Restart(ctx context.Context, workflowId string, opts *WorkflowResourceApiRestartOpts) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/restart", workflowId)

	queryParams := url.Values{}
	if opts != nil && opts.UseLatestDefinitions.IsSet() {
		queryParams.Add("useLatestDefinitions", parameterToString(opts.UseLatestDefinitions.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Deprecated: Use Resume instead.
func (a *WorkflowResourceApiService) ResumeWorkflow(ctx context.Context, workflowId string) (*http.Response, error) {
	return a.Resume(ctx, workflowId)
}

// Resume resumes a paused workflow execution.
func (a *WorkflowResourceApiService) Resume(ctx context.Context, workflowId string) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/resume", workflowId)

	resp, err := a.Put(ctx, path, nil, nil)
	if err != nil {
		return resp, err
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

// Retry retries the last failed task.
func (a *WorkflowResourceApiService) Retry(ctx context.Context, workflowId string, opts *WorkflowResourceApiRetryOpts) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/retry", workflowId)

	queryParams := url.Values{}
	if opts != nil && opts.ResumeSubworkflowTasks.IsSet() {
		queryParams.Add("resumeSubworkflowTasks", parameterToString(opts.ResumeSubworkflowTasks.Value(), ""))
	}
	if opts != nil && opts.RetryIfRetriedByParent.IsSet() {
		queryParams.Add("retryIfRetriedByParent", parameterToString(opts.RetryIfRetriedByParent.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, nil, nil)
	if err != nil {
		return resp, err
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

// Search searches for workflows.

func (a *WorkflowResourceApiService) Search(ctx context.Context, opts *WorkflowResourceApiSearchOpts) (model.SearchResultWorkflowSummary, *http.Response, error) {
	var result model.SearchResultWorkflowSummary

	path := "/workflow/search"

	queryParams := url.Values{}
	if opts != nil && opts.Start.IsSet() {
		queryParams.Add("start", parameterToString(opts.Start.Value(), ""))
	}
	if opts != nil && opts.Size.IsSet() {
		queryParams.Add("size", parameterToString(opts.Size.Value(), ""))
	}
	if opts != nil && opts.Sort.IsSet() {
		queryParams.Add("sort", parameterToString(opts.Sort.Value(), ""))
	}
	if opts != nil && opts.FreeText.IsSet() {
		queryParams.Add("freeText", parameterToString(opts.FreeText.Value(), ""))
	}
	if opts != nil && opts.Query.IsSet() {
		queryParams.Add("query", parameterToString(opts.Query.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.SearchResultWorkflowSummary{}, resp, err
	}
	return result, resp, nil
}

/*
WorkflowResourceApiService Search for workflows based on payload and other parameters
use sort options as sort&#x3D;&lt;field&gt;:ASC|DESC e.g. sort&#x3D;name&amp;sort&#x3D;workflowId:DESC. If order is not specified, defaults to ASC.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *WorkflowResourceApiSearchV2Opts - Optional Parameters:
     * @param "Start" (optional.Int32) -
     * @param "Size" (optional.Int32) -
     * @param "Sort" (optional.String) -
     * @param "FreeText" (optional.String) -
     * @param "Query" (optional.String) -
@return http_model.SearchResultWorkflow
*/

type WorkflowResourceApiSearchV2Opts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

func (a *WorkflowResourceApiService) SearchV2(ctx context.Context, opts *WorkflowResourceApiSearchV2Opts) (model.SearchResultWorkflow, *http.Response, error) {
	var result model.SearchResultWorkflow

	path := "/workflow/search-v2"

	queryParams := url.Values{}
	if opts != nil && opts.Start.IsSet() {
		queryParams.Add("start", parameterToString(opts.Start.Value(), ""))
	}
	if opts != nil && opts.Size.IsSet() {
		queryParams.Add("size", parameterToString(opts.Size.Value(), ""))
	}
	if opts != nil && opts.Sort.IsSet() {
		queryParams.Add("sort", parameterToString(opts.Sort.Value(), ""))
	}
	if opts != nil && opts.FreeText.IsSet() {
		queryParams.Add("freeText", parameterToString(opts.FreeText.Value(), ""))
	}
	if opts != nil && opts.Query.IsSet() {
		queryParams.Add("query", parameterToString(opts.Query.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.SearchResultWorkflow{}, resp, err
	}
	return result, resp, nil
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

// SearchWorkflowsByTasks search workflows by tasks.
func (a *WorkflowResourceApiService) SearchWorkflowsByTasks(ctx context.Context, opts *WorkflowResourceApiSearchWorkflowsByTasksOpts) (model.SearchResultWorkflowSummary, *http.Response, error) {
	var result model.SearchResultWorkflowSummary

	localVarPath := "/workflow/search-by-tasks"

	queryParams := url.Values{}
	if opts != nil && opts.Start.IsSet() {
		queryParams.Add("start", parameterToString(opts.Start.Value(), ""))
	}
	if opts != nil && opts.Size.IsSet() {
		queryParams.Add("size", parameterToString(opts.Size.Value(), ""))
	}
	if opts != nil && opts.Sort.IsSet() {
		queryParams.Add("sort", parameterToString(opts.Sort.Value(), ""))
	}
	if opts != nil && opts.FreeText.IsSet() {
		queryParams.Add("freeText", parameterToString(opts.FreeText.Value(), ""))
	}
	if opts != nil && opts.Query.IsSet() {
		queryParams.Add("query", parameterToString(opts.Query.Value(), ""))
	}

	resp, err := a.Get(ctx, localVarPath, queryParams, &result)
	if err != nil {
		return model.SearchResultWorkflowSummary{}, resp, err
	}
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

// SearchWorkflowsByTasksV2 search workflows by tasks V2.
func (a *WorkflowResourceApiService) SearchWorkflowsByTasksV2(ctx context.Context, opts *WorkflowResourceApiSearchWorkflowsByTasksV2Opts) (model.SearchResultWorkflow, *http.Response, error) {
	var result model.SearchResultWorkflow

	localVarPath := "/workflow/search-by-tasks-v2"

	queryParams := url.Values{}
	if opts != nil && opts.Start.IsSet() {
		queryParams.Add("start", parameterToString(opts.Start.Value(), ""))
	}
	if opts != nil && opts.Size.IsSet() {
		queryParams.Add("size", parameterToString(opts.Size.Value(), ""))
	}
	if opts != nil && opts.Sort.IsSet() {
		queryParams.Add("sort", parameterToString(opts.Sort.Value(), ""))
	}
	if opts != nil && opts.FreeText.IsSet() {
		queryParams.Add("freeText", parameterToString(opts.FreeText.Value(), ""))
	}
	if opts != nil && opts.Query.IsSet() {
		queryParams.Add("query", parameterToString(opts.Query.Value(), ""))
	}

	resp, err := a.Get(ctx, localVarPath, queryParams, &result)
	if err != nil {
		return model.SearchResultWorkflow{}, resp, err
	}
	return result, resp, nil
}

// SkipTaskFromWorkflow skip task from workflow.
func (a *WorkflowResourceApiService) SkipTaskFromWorkflow(ctx context.Context, workflowId string, taskReferenceName string, skipTaskRequest model.SkipTaskRequest) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/skiptask/%s", workflowId, taskReferenceName)

	queryParams := url.Values{}

	resp, err := a.PutWithParams(ctx, path, queryParams, skipTaskRequest, &model.SkipTaskRequest{})
	if err != nil {
		return resp, err
	}
	return resp, nil
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

// StartWorkflow starts a workflow execution.
func (a *WorkflowResourceApiService) StartWorkflow(ctx context.Context, body map[string]interface{}, name string, opts *WorkflowResourceApiStartWorkflowOpts) (string, *http.Response, error) {
	var result string

	path := fmt.Sprintf("/workflow/%s", name)

	queryParams := url.Values{}
	if opts != nil && opts.Version.IsSet() {
		queryParams.Add("version", parameterToString(opts.Version.Value(), ""))
	}
	if opts != nil && opts.CorrelationId.IsSet() {
		queryParams.Add("correlationId", parameterToString(opts.CorrelationId.Value(), ""))
	}
	if opts != nil && opts.Priority.IsSet() {
		queryParams.Add("priority", parameterToString(opts.Priority.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, &result)
	if err != nil {
		return "", resp, err
	}
	return result, resp, nil
}

func (a *WorkflowResourceApiService) executeWorkflowImpl(
	ctx context.Context,
	body model.StartWorkflowRequest,
	requestId string,
	name string,
	version int32,
	waitUntilTask []string,
	waitForSeconds int,
	consistency string,
	returnStrategy string) (model.SignalResponse, *http.Response, error) {

	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue model.SignalResponse
	)

	path := fmt.Sprintf("/workflow/execute/%s/%d", name, version)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "application/json"
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if requestId != "" {
		localVarQueryParams.Add("requestId", parameterToString(requestId, ""))
	}

	if len(waitUntilTask) > 0 {
		localVarQueryParams.Add("waitUntilTaskRef", strings.Join(waitUntilTask, ","))
	}

	if waitForSeconds > 0 {
		localVarQueryParams.Add("waitForSeconds", parameterToString(waitForSeconds, ""))
	}

	if consistency != "" {
		localVarQueryParams.Add("consistency", parameterToString(consistency, ""))
	}

	if returnStrategy != "" {
		localVarQueryParams.Add("returnStrategy", parameterToString(returnStrategy, ""))
	}

	localVarPostBody = &body
	r, err := a.prepareRequest(ctx, path, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return model.SignalResponse{}, nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return model.SignalResponse{}, localVarHttpResponse, err
	}

	localVarBody, err := getDecompressedBody(localVarHttpResponse)

	localVarHttpResponse.Body.Close()
	if err != nil {
		return model.SignalResponse{}, localVarHttpResponse, err
	}

	if isSuccessfulStatus(localVarHttpResponse.StatusCode) {
		// Decode directly into SignalResponse since API returns unified format
		var signalResponse model.SignalResponse
		err = a.decode(&signalResponse, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		localVarReturnValue = signalResponse
	} else {
		newErr := NewGenericSwaggerError(localVarBody, localVarHttpResponse.Status, nil, localVarHttpResponse.StatusCode)
		return model.SignalResponse{}, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
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

// ExecuteWorkflowWithReturnStrategy executes a workflow with the specified return strategy
func (a *WorkflowResourceApiService) ExecuteWorkflowWithReturnStrategy(ctx context.Context, body model.StartWorkflowRequest, opts ExecuteWorkflowOpts) (*model.SignalResponse, error) {
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

	// Validate required fields
	if body.Name == "" {
		return nil, fmt.Errorf("workflow name is required")
	}
	if body.Version <= 0 {
		return nil, fmt.Errorf("workflow version must be greater than 0")
	}

	// Create a new context with the same timeout as waitForSeconds
	var cancelFunc context.CancelFunc
	var effectiveCtx context.Context
	if opts.WaitForSeconds > 0 {
		// Add buffer time: 5 seconds for HTTP overhead + API processing
		// This ensures the context doesn't timeout before the API can respond
		bufferSeconds := 10
		totalTimeout := time.Duration(opts.WaitForSeconds+bufferSeconds) * time.Second

		effectiveCtx, cancelFunc = context.WithTimeout(ctx, totalTimeout)
		defer cancelFunc()

	}

	// Call the existing internal method
	response, _, err := a.executeWorkflowImpl(
		effectiveCtx,
		body,
		opts.RequestID,
		body.Name,
		body.Version,
		opts.WaitUntilTaskRef,
		opts.WaitForSeconds,
		string(opts.Consistency),
		string(opts.ReturnStrategy),
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ExecuteWorkflow execute workflow synchronously.
func (a *WorkflowResourceApiService) ExecuteWorkflow(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask string) (model.WorkflowRun, *http.Response, error) {
	var result model.WorkflowRun

	path := fmt.Sprintf("/workflow/execute/%s/%d", name, version)

	queryParams := url.Values{}
	queryParams.Add("requestId", parameterToString(requestId, ""))
	if len(waitUntilTask) > 0 {
		queryParams.Add("waitUntilTaskRef", parameterToString(waitUntilTask, ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, &result)
	if err != nil {
		return model.WorkflowRun{}, resp, err
	}
	return result, resp, nil
}

// Enterprise: This feature requires Orkes Conductor Enterprise license, NOT AVAILABLE in OSS.
func (a *WorkflowResourceApiService) ExecuteAndGetBlockingTask(
	ctx context.Context,
	body model.StartWorkflowRequest,
	requestId string,
	name string,
	version int32,
	waitUntilTask []string,
	waitForSeconds int,
	consistency string) (model.TaskRun, *http.Response, error) {

	returnStrategy := "BLOCKING_TASK"

	response, httpResponse, err := a.executeWorkflowImpl(
		ctx,
		body,
		requestId,
		name,
		version,
		waitUntilTask,
		waitForSeconds,
		consistency,
		returnStrategy,
	)

	if err != nil {
		return model.TaskRun{}, httpResponse, err
	}

	return response.GetTaskRun(), httpResponse, nil
}

// Enterprise: This feature requires Orkes Conductor Enterprise license, NOT AVAILABLE in OSS.
func (a *WorkflowResourceApiService) ExecuteAndGetBlockingTaskInput(
	ctx context.Context,
	body model.StartWorkflowRequest,
	requestId string,
	name string,
	version int32,
	waitUntilTask []string,
	waitForSeconds int,
	consistency string) (model.TaskRun, *http.Response, error) {

	returnStrategy := "BLOCKING_TASK_INPUT"

	response, httpResponse, err := a.executeWorkflowImpl(
		ctx,
		body,
		requestId,
		name,
		version,
		waitUntilTask,
		waitForSeconds,
		consistency,
		returnStrategy,
	)

	if err != nil {
		return model.TaskRun{}, httpResponse, err
	}

	return response.GetTaskRun(), httpResponse, nil
}

// Enterprise: This feature requires Orkes Conductor Enterprise license, NOT AVAILABLE in OSS.
func (a *WorkflowResourceApiService) ExecuteAndGetBlockingWorkflow(
	ctx context.Context,
	body model.StartWorkflowRequest,
	requestId string,
	name string,
	version int32,
	waitUntilTask []string,
	waitForSeconds int,
	consistency string) (model.WorkflowRun, *http.Response, error) {

	returnStrategy := "BLOCKING_WORKFLOW"

	response, httpResponse, err := a.executeWorkflowImpl(
		ctx,
		body,
		requestId,
		name,
		version,
		waitUntilTask,
		waitForSeconds,
		consistency,
		returnStrategy,
	)

	if err != nil {
		return model.WorkflowRun{}, httpResponse, err
	}

	return response.GetWorkflowRun(), httpResponse, nil
}

// Enterprise: This feature requires Orkes Conductor Enterprise license, NOT AVAILABLE in OSS.
func (a *WorkflowResourceApiService) ExecuteAndGetTarget(
	ctx context.Context,
	body model.StartWorkflowRequest,
	requestId string,
	name string,
	version int32,
	waitUntilTask []string,
	waitForSeconds int,
	consistency string) (model.WorkflowRun, *http.Response, error) {

	returnStrategy := "TARGET_WORKFLOW"

	response, httpResponse, err := a.executeWorkflowImpl(
		ctx,
		body,
		requestId,
		name,
		version,
		waitUntilTask,
		waitForSeconds,
		consistency,
		returnStrategy,
	)

	if err != nil {
		return model.WorkflowRun{}, httpResponse, err
	}

	return response.GetWorkflowRun(), httpResponse, nil
}

// StartWorkflowWithRequest starts a workflow with request
func (a *WorkflowResourceApiService) StartWorkflowWithRequest(ctx context.Context, body model.StartWorkflowRequest) (string, *http.Response, error) {
	var result string

	path := "/workflow"

	resp, err := a.Post(ctx, path, body, &result)
	if err != nil {
		return "", resp, err
	}
	return result, resp, nil
}

// WorkflowResourceApiTerminateOpts contains optional parameters for Terminate
type WorkflowResourceApiTerminateOpts struct {
	// Reason a reason for termination.
	Reason optional.String
	// TriggerFailureWorkflow if set to true,  the associated compensation flow (if any) will be triggered.
	TriggerFailureWorkflow optional.Bool
}

// Terminate terminates a workflow execution.
func (a *WorkflowResourceApiService) Terminate(ctx context.Context, workflowId string, opts *WorkflowResourceApiTerminateOpts) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s", workflowId)

	queryParams := url.Values{}
	if opts != nil && opts.Reason.IsSet() {
		queryParams.Add("reason", parameterToString(opts.Reason.Value(), ""))
	}
	if opts != nil && opts.TriggerFailureWorkflow.IsSet() {
		queryParams.Add("triggerFailureWorkflow", parameterToString(opts.TriggerFailureWorkflow.Value(), ""))
	}

	resp, err := a.APIClient.Delete(ctx, path, queryParams, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// WorkflowResourceApiJumpToTaskOpts contains optional parameters for JumpToTask
type WorkflowResourceApiJumpToTaskOpts struct {
	// TaskReferenceName the reference name of the task to jump to.
	TaskReferenceName optional.String
}

// JumpToTask jumps to a specific task in a running workflow.
func (a *WorkflowResourceApiService) JumpToTask(ctx context.Context, body map[string]interface{}, workflowId string, optionals *WorkflowResourceApiJumpToTaskOpts) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/jump/{taskReferenceName}", workflowId)

	queryParams := url.Values{}
	if optionals != nil && optionals.TaskReferenceName.IsSet() {
		queryParams.Add("taskReferenceName", parameterToString(optionals.TaskReferenceName.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, nil)
	if err != nil {
		return resp, err
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

// UpdateWorkflowAndTaskState update workflow and task state.
func (a *WorkflowResourceApiService) UpdateWorkflowAndTaskState(ctx context.Context, body model.WorkflowStateUpdate, requestId string, workflowId string, optionals *WorkflowResourceApiUpdateWorkflowAndTaskStateOpts) (model.WorkflowRun, *http.Response, error) {
	var result model.WorkflowRun

	// create path and map variables
	path := fmt.Sprintf("/workflow/%s/state", workflowId)

	queryParams := url.Values{}
	queryParams.Add("requestId", parameterToString(requestId, ""))
	if optionals != nil && optionals.WaitUntilTaskRef.IsSet() {
		queryParams.Add("waitUntilTaskRef", parameterToString(optionals.WaitUntilTaskRef.Value(), ""))
	}
	if optionals != nil && optionals.WaitForSeconds.IsSet() {
		queryParams.Add("waitForSeconds", parameterToString(optionals.WaitForSeconds.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil
}

// UpgradeRunningWorkflowToVersion upgrade running workflow to newer version.
func (a *WorkflowResourceApiService) UpgradeRunningWorkflowToVersion(ctx context.Context, body model.UpgradeWorkflowRequest, workflowId string) (*http.Response, error) {
	// create path and map variables
	path := fmt.Sprintf("/workflow/%s/upgrade", workflowId)

	resp, err := a.Post(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// TestWorkflow tests workflow execution using mock data.
func (a *WorkflowResourceApiService) TestWorkflow(ctx context.Context, body model.WorkflowTestRequest) (model.Workflow, *http.Response, error) {
	var result model.Workflow

	// create path and map variables
	path := "/workflow/test"

	resp, err := a.Post(ctx, path, body, &result)
	if err != nil {
		return model.Workflow{}, resp, err
	}
	return result, resp, nil
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

// GetExecutionStatusTaskList gets execution task list.
func (a *WorkflowResourceApiService) GetExecutionStatusTaskList(ctx context.Context, workflowID string, opts *WorkflowResourceAPIGetExecutionStatusTaskListOpts) (model.TaskListSearchResultSummary, *http.Response, error) {
	var result model.TaskListSearchResultSummary

	path := fmt.Sprintf("/workflow/%s/tasks", workflowID)

	queryParams := url.Values{}
	if opts != nil && opts.Start.IsSet() {
		queryParams.Add("start", parameterToString(opts.Start.Value(), ""))
	}
	if opts != nil && opts.Count.IsSet() {
		queryParams.Add("count", parameterToString(opts.Count.Value(), ""))
	}
	if opts != nil && opts.Status.IsSet() {
		queryParams.Add("status", parameterToString(opts.Status.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.TaskListSearchResultSummary{}, resp, err
	}
	return result, resp, nil
}

// UpdateWorkflowState updates workflow variables for a running workflow.
//
// This method is similar to the Set Variable task, except the variables can be updated anytime in real time.
func (a *WorkflowResourceApiService) UpdateWorkflowState(ctx context.Context, body map[string]interface{}, workflowID string) (model.Workflow, *http.Response, error) {
	var result model.Workflow

	path := fmt.Sprintf("/workflow/%s/variables", workflowID)

	resp, err := a.Post(ctx, path, body, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil
}
