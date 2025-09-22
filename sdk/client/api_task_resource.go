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
	"net/http"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// TaskResourceApiService
type TaskResourceApiService struct {
	*APIClient
}

// NewTaskApiService creates a new TaskResourceApiService instance
func NewTaskApiService(client *APIClient) *TaskResourceApiService {
	return &TaskResourceApiService{APIClient: client}
}

// All gets all task queue sizes
func (a *TaskResourceApiService) All(ctx context.Context) (map[string]int64, *http.Response, error) {
	// Use Orkes client
	result, resp, err := a.http_orkes.TaskResourceAPI.All(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// AllVerbose gets all tasks with details
func (a *TaskResourceApiService) AllVerbose(ctx context.Context) (map[string]map[string]map[string]int64, *http.Response, error) {
	// Use Orkes client
	result, resp, err := a.http_orkes.TaskResourceAPI.AllVerbose(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	if result != nil {
		return *result, resp, nil
	}
	return nil, resp, nil
}

// TaskResourceApiBatchPollOpts are the options for the BatchPoll method
type TaskResourceApiBatchPollOpts struct {
	Workerid optional.String
	Domain   optional.String
	Count    optional.Int32
	Timeout  optional.Int32
}

// BatchPoll batch poll for tasks
func (a *TaskResourceApiService) BatchPoll(ctx context.Context, tasktype string, opts *TaskResourceApiBatchPollOpts) ([]model.Task, *http.Response, error) {
	// Use Orkes client
	req := a.http_orkes.TaskResourceAPI.BatchPoll(ctx, tasktype)

	if opts != nil {
		if opts.Workerid.IsSet() {
			req = req.Workerid(opts.Workerid.Value())
		}
		if opts.Count.IsSet() {
			req = req.Count(opts.Count.Value())
		}
		if opts.Timeout.IsSet() {
			req = req.Timeout(opts.Timeout.Value())
		}
		if opts.Domain.IsSet() {
			req = req.Domain(opts.Domain.Value())
		}
	}

	genTasks, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert tasks using mapper
	tasks := toDomainTasksFromGenerated(genTasks)
	return tasks, resp, nil
}

// GetAllPollData gets poll data for all task types
func (a *TaskResourceApiService) GetAllPollData(ctx context.Context) ([]model.PollData, *http.Response, error) {
	// Use Orkes client
	result, resp, err := a.http_orkes.TaskResourceAPI.GetAllPollData(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert map result to []model.PollData using mapper
	pollData := make([]model.PollData, 0)
	for _, data := range result {
		pd := toDomainPollDataFromMap(data)
		pollData = append(pollData, pd)
	}

	return pollData, resp, nil
}

// GetPollData gets poll data for a task type
func (a *TaskResourceApiService) GetPollData(ctx context.Context, taskType string) ([]model.PollData, *http.Response, error) {
	// Use Orkes client
	req := a.http_orkes.TaskResourceAPI.GetPollData(ctx)
	if taskType != "" {
		req = req.TaskType(taskType)
	}
	genPollData, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert poll data using mapper
	pollData := toDomainPollDataListFromGenerated(genPollData)
	return pollData, resp, nil
}

// GetExternalStorageLocation1 gets external storage location
func (a *TaskResourceApiService) GetExternalStorageLocation1(ctx context.Context, path string, operation string, payloadType string) (model.ExternalStorageLocation, *http.Response, error) {
	// Use Conductor client as this method may not be available in Orkes
	result, resp, err := a.http_conductor.TaskResourceAPI.GetExternalStorageLocation3(ctx).
		Path(path).
		Operation(operation).
		PayloadType(payloadType).
		Execute()
	if err != nil {
		return model.ExternalStorageLocation{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert using existing mapper
	domainLoc := toDomainExternalStorageLocation(result)
	return domainLoc, resp, nil
}

// GetTask gets task details by task ID
func (a *TaskResourceApiService) GetTask(ctx context.Context, taskId string) (model.Task, *http.Response, error) {
	// Use Orkes client
	genTask, resp, err := a.http_orkes.TaskResourceAPI.GetTask(ctx, taskId).Execute()
	if err != nil {
		return model.Task{}, resp, wrapGeneratedError(err, resp)
	}

	task := toDomainTask(genTask)
	return task, resp, nil
}

// Log adds a log to a task
func (a *TaskResourceApiService) Log(ctx context.Context, body string, taskId string) (*http.Response, error) {
	// Use Orkes client
	resp, err := a.http_orkes.TaskResourceAPI.Log(ctx, taskId).Body(body).Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// GetTaskLogs gets task execution logs
func (a *TaskResourceApiService) GetTaskLogs(ctx context.Context, taskId string) ([]model.TaskExecLog, *http.Response, error) {
	// Use Orkes client
	genLogs, resp, err := a.http_orkes.TaskResourceAPI.GetTaskLogs(ctx, taskId).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert logs using mapper
	logs := toDomainTaskExecLogsFromGenerated(genLogs)
	return logs, resp, nil
}

// TaskResourceApiPollOpts are the options for the Poll method
type TaskResourceApiPollOpts struct {
	Workerid optional.String
	Domain   optional.String
}

// Poll polls for a task
func (a *TaskResourceApiService) Poll(ctx context.Context, tasktype string, opts *TaskResourceApiPollOpts) (model.Task, *http.Response, error) {
	// Use Orkes client
	req := a.http_orkes.TaskResourceAPI.Poll(ctx, tasktype)

	if opts != nil {
		if opts.Workerid.IsSet() {
			req = req.Workerid(opts.Workerid.Value())
		}
		if opts.Domain.IsSet() {
			req = req.Domain(opts.Domain.Value())
		}
	}

	genTask, resp, err := req.Execute()
	if err != nil {
		return model.Task{}, resp, wrapGeneratedError(err, resp)
	}

	task := toDomainTask(genTask)
	return task, resp, nil
}

// AddTaskLog adds a task log
func (a *TaskResourceApiService) AddTaskLog(ctx context.Context, body string, taskId string) (*http.Response, error) {
	// Use Orkes client
	resp, err := a.http_orkes.TaskResourceAPI.Log(ctx, taskId).Body(body).Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// RequeuePendingTask requeues a pending task by task type
func (a *TaskResourceApiService) RequeuePendingTask(ctx context.Context, taskType string) (string, *http.Response, error) {
	// Use Orkes client
	result, resp, err := a.http_orkes.TaskResourceAPI.RequeuePendingTask(ctx, taskType).Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// RequeuePendingTasksByTaskType requeues pending tasks by task type
func (a *TaskResourceApiService) RequeuePendingTasksByTaskType(ctx context.Context, taskType string) (string, *http.Response, error) {
	// Use Orkes client
	result, resp, err := a.http_orkes.TaskResourceAPI.RequeuePendingTask(ctx, taskType).Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// TaskResourceApiSearch1Opts are the options for the Search method
type TaskResourceApiSearch1Opts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

// Search searches for tasks
func (a *TaskResourceApiService) Search(ctx context.Context, opts *TaskResourceApiSearch1Opts) (model.SearchResultTaskSummary, *http.Response, error) {
	// Use Orkes client
	req := a.http_orkes.TaskResourceAPI.Search2(ctx)

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
		return model.SearchResultTaskSummary{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert generated result to domain model using mapper
	result := toDomainSearchResultTaskSummaryFromGenerated(genResult)
	return result, resp, nil
}

// TaskResourceApiSearchV21Opts are the options for the SearchV2 method
type TaskResourceApiSearchV2Opts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

// SearchV2 searches for tasks with V2 API
func (a *TaskResourceApiService) SearchV2(ctx context.Context, opts *TaskResourceApiSearchV2Opts) (model.SearchResultTask, *http.Response, error) {
	// Use Conductor client as this method may not be available in Orkes
	req := a.http_conductor.TaskResourceAPI.SearchV21(ctx)

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
		return model.SearchResultTask{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert generated result to domain model using mapper
	result := toDomainSearchResultTask(genResult)
	return result, resp, nil
}

type TaskResourceApiSizeOpts struct {
	TaskType optional.Interface
}

// Size gets the size of the pending queue for a task type
func (a *TaskResourceApiService) Size(ctx context.Context, opts *TaskResourceApiSizeOpts) (map[string]int32, *http.Response, error) {
	// Extract task types from options
	taskTypes := []string{}
	if opts != nil && opts.TaskType.IsSet() {
		if types, ok := opts.TaskType.Value().([]string); ok {
			taskTypes = types
		}
	}

	// Call GetTaskQueueSizes and convert result
	sizes64, resp, err := a.http_orkes.TaskResourceAPI.Size(ctx).TaskType(taskTypes).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert map[string]int64 to map[string]int32
	sizes32 := make(map[string]int32)
	for k, v := range sizes64 {
		sizes32[k] = int32(v)
	}

	return sizes32, resp, nil
}

// UpdateTask updates a task
func (a *TaskResourceApiService) UpdateTask(ctx context.Context, body *model.TaskResult) (string, *http.Response, error) {
	// Convert to generated model using mapper
	genBody := toGeneratedTaskResult(body)

	// Use Orkes client
	result, resp, err := a.http_orkes.TaskResourceAPI.UpdateTask(ctx).TaskResult(*genBody).Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}

	return result, resp, nil
}

// UpdateTaskByRefName updates a task by reference name
func (a *TaskResourceApiService) UpdateTaskByRefName(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, status string) (string, *http.Response, error) {
	// Use Orkes client
	result, resp, err := a.http_orkes.TaskResourceAPI.UpdateTask1(ctx, workflowId, taskRefName, status).
		RequestBody(body).
		Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// TaskResourceApiUpdateTaskSyncOpts are the options for the UpdateTaskSync method
type TaskResourceApiUpdateTaskSyncOpts struct {
	Workerid optional.String
}

// UpdateTaskByRefNameWithWorkerId updates a task by reference name with worker ID
func (a *TaskResourceApiService) UpdateTaskByRefNameWithWorkerId(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, status string, workerId optional.String) (string, *http.Response, error) {
	// Use Orkes client
	req := a.http_orkes.TaskResourceAPI.UpdateTask1(ctx, workflowId, taskRefName, status).
		RequestBody(body)

	if workerId.IsSet() {
		req = req.Workerid(workerId.Value())
	}

	result, resp, err := req.Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// SignalTaskOpts contains options for the Signal method
type SignalTaskOpts struct {
	ReturnStrategy model.ReturnStrategy
}

// DefaultSignalTaskOpts returns the default options for Signal
func DefaultSignalTaskOpts() SignalTaskOpts {
	return SignalTaskOpts{
		ReturnStrategy: model.ReturnTargetWorkflow, // Set TARGET_WORKFLOW as default
	}
}

// Signal signals a task
func (a *TaskResourceApiService) Signal(ctx context.Context, body map[string]interface{}, workflowID string, status model.WorkflowStatus, opts ...SignalTaskOpts) (*model.SignalResponse, error) {

	// Get options with defaults
	options := DefaultSignalTaskOpts()
	if len(opts) > 0 {
		options = opts[0]
	}

	// Use the generated method directly
	response, _, err := a.http_orkes.TaskResourceAPI.SignalWorkflowTaskSync(ctx, workflowID, string(status)).
		RequestBody(body).
		ReturnStrategy(string(options.ReturnStrategy)).
		Execute()
	if err != nil {
		return nil, wrapGeneratedError(err, nil)
	}

	// Convert generated SignalResponse to domain model
	domainResponse := toDomainSignalResponseFromGenerated(response)
	return &domainResponse, nil
}

// SignalAsync signals a task asynchronously
func (a *TaskResourceApiService) SignalAsync(ctx context.Context, body map[string]interface{}, workflowId string, status string) (*http.Response, error) {
	// Use Orkes generated client SignalWorkflowTaskASync
	resp, err := a.http_orkes.TaskResourceAPI.SignalWorkflowTaskASync(ctx, workflowId, status).
		RequestBody(body).
		Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// UpdateTaskSync updates a task synchronously
func (a *TaskResourceApiService) UpdateTaskSync(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, status string) (model.Workflow, *http.Response, error) {
	// Use Orkes client
	result, resp, err := a.http_orkes.TaskResourceAPI.UpdateTaskSync(ctx, workflowId, taskRefName, status).
		RequestBody(body).
		Execute()
	if err != nil {
		return model.Workflow{}, resp, wrapGeneratedError(err, resp)
	}

	// Create minimal task response
	workflow := toDomainWorkflow(result)
	return workflow, resp, nil
}
