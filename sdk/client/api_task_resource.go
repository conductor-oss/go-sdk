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
	"os"
	"sync"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/log"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

var (
	hostname string
	once     sync.Once
)

// TaskResourceApiService is the service for the task resource.
type TaskResourceApiService struct {
	*APIClient
}

// All gets all task queue sizes.
func (a *TaskResourceApiService) All(ctx context.Context) (map[string]int64, *http.Response, error) {
	// Use Orkes client
	result, resp, err := a.http_orkes.TaskResourceAPI.All(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// AllVerbose gets all tasks with details.
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

// BatchPoll batch poll for tasks.
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

// GetAllPollData gets poll data for all task types.
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

// GetPollData gets poll data for a task type.
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

func (a *TaskResourceApiService) GetExternalStorageLocation(ctx context.Context, path string, operation string, payloadType string) (model.ExternalStorageLocation, *http.Response, error) {
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

// Deprecated: Use GetExternalStorageLocation instead.
func (a *TaskResourceApiService) GetExternalStorageLocation1(ctx context.Context, path string, operation string, payloadType string) (model.ExternalStorageLocation, *http.Response, error) {
	return a.GetExternalStorageLocation(ctx, path, operation, payloadType)
}

// GetTask gets task details by task ID.
func (a *TaskResourceApiService) GetTask(ctx context.Context, taskId string) (model.Task, *http.Response, error) {
	// Use Orkes client
	genTask, resp, err := a.http_orkes.TaskResourceAPI.GetTask(ctx, taskId).Execute()
	if err != nil {
		return model.Task{}, resp, wrapGeneratedError(err, resp)
	}

	task := toDomainTask(genTask)
	return task, resp, nil
}

// Log adds a log to a task.
func (a *TaskResourceApiService) Log(ctx context.Context, body string, taskId string) (*http.Response, error) {
	// Use Orkes client
	resp, err := a.http_orkes.TaskResourceAPI.Log(ctx, taskId).Body(body).Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// GetTaskLogs gets task execution logs.
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

// TaskResourceApiPollOpts are the options for the Poll method.
type TaskResourceApiPollOpts struct {
	Workerid optional.String
	Domain   optional.String
}

// Poll polls for a task.
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

// RequeuePendingTask requeues a pending task by task type.
func (a *TaskResourceApiService) RequeuePendingTask(ctx context.Context, taskType string) (string, *http.Response, error) {
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

// Search searches for tasks.
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
type TaskResourceApiSearchV21Opts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

// SearchV2 searches for tasks.
func (a *TaskResourceApiService) SearchV2(ctx context.Context, opts *TaskResourceApiSearchV21Opts) (model.SearchResultTask, *http.Response, error) {
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

// TaskResourceApiSizeOpts are the options for the Size method
type TaskResourceApiSizeOpts struct {
	TaskType optional.Interface
}

// Size gets the size of the task type.
func (a *TaskResourceApiService) Size(ctx context.Context, opts *TaskResourceApiSizeOpts) (map[string]int32, *http.Response, error) {
	req := a.http_orkes.TaskResourceAPI.Size(ctx)

	if opts != nil && opts.TaskType.IsSet() {
		if taskType, ok := opts.TaskType.Value().([]string); ok {
			req = req.TaskType(taskType)
		}
	}

	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	return result, resp, nil
}

// UpdateTask updates a task.
func (a *TaskResourceApiService) UpdateTask(ctx context.Context, taskResult *model.TaskResult) (string, *http.Response, error) {
	// Convert to generated model using mapper
	genBody := toGeneratedTaskResult(taskResult)

	// Use Orkes client
	result, resp, err := a.http_orkes.TaskResourceAPI.UpdateTask(ctx).TaskResult(*genBody).Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}

	return result, resp, nil
}

// TaskResourceApiUpdateTaskSyncOpts are the options for the UpdateTaskSync method.
type TaskResourceApiUpdateTaskSyncOpts struct {
	Workerid optional.String
}

// UpdateTaskSync updates a task synchronously.
func (a *TaskResourceApiService) UpdateTaskSync(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, status string, localVarOptionals *TaskResourceApiUpdateTaskSyncOpts) (model.Workflow, *http.Response, error) {
	req := a.http_orkes.TaskResourceAPI.UpdateTaskSync(ctx, workflowId, taskRefName, status).
		RequestBody(body)

	if localVarOptionals != nil {
		if localVarOptionals.Workerid.IsSet() {
			req = req.Workerid(localVarOptionals.Workerid.Value())
		}
	}

	result, resp, err := req.Execute()
	if err != nil {
		return model.Workflow{}, resp, wrapGeneratedError(err, resp)
	}

	// Create minimal task response
	workflow := toDomainWorkflow(result)
	return workflow, resp, nil
}

// SignalAsync signals a task asynchronously.
func (a *TaskResourceApiService) SignalAsync(ctx context.Context, body map[string]interface{}, workflowId string, status string) (*http.Response, error) {
	// create path and map variables
	resp, err := a.http_orkes.TaskResourceAPI.SignalWorkflowTaskASync(ctx, workflowId, status).
		RequestBody(body).
		Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// SignalTaskOpts contains options for the Signal method.
type SignalTaskOpts struct {
	ReturnStrategy model.ReturnStrategy
}

// DefaultSignalTaskOpts returns the default options for Signal.
func DefaultSignalTaskOpts() SignalTaskOpts {
	return SignalTaskOpts{
		ReturnStrategy: model.ReturnTargetWorkflow, // Set TARGET_WORKFLOW as default
	}
}

// SignalTask signals a task in a workflow synchronously with the specified return strategy.
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

// UpdateTaskByRefName updates a task by reference name.
func (a *TaskResourceApiService) UpdateTaskByRefName(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, status string) (string, *http.Response, error) {
	return a.UpdateTaskByRefNameWithWorkerId(ctx, body, workflowId, taskRefName, status, optional.EmptyString())
}

// UpdateTaskByRefNameWithWorkerId updates a task by reference name with worker ID.
func (a *TaskResourceApiService) UpdateTaskByRefNameWithWorkerId(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, status string, workerId optional.String) (string, *http.Response, error) {
	req := a.http_orkes.TaskResourceAPI.UpdateTask1(ctx, workflowId, taskRefName, status).
		RequestBody(body)

	if workerId.IsSet() {
		req = req.Workerid(workerId.Value())
	} else {
		req = req.Workerid(getHostname())
	}

	result, resp, err := req.Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

func getHostname() string {
	once.Do(updateHostname)
	return hostname
}

func updateHostname() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		log.Warn("Failed to get hostname", "error", err)
	}
}
