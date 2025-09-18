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

// SchedulerResourceApiService wraps the generated client to maintain backward compatibility
type SchedulerResourceApiService struct {
	*APIClient
}

// NewSchedulerResourceApiService creates a service from existing APIClient
func NewSchedulerResourceApiService(apiClient *APIClient) *SchedulerResourceApiService {
	if apiClient == nil {
		return nil
	}

	return &SchedulerResourceApiService{
		APIClient: apiClient,
	}
}

// DeleteSchedule deletes a schedule
func (a *SchedulerResourceApiService) DeleteSchedule(ctx context.Context, name string) (interface{}, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.DeleteSchedule(ctx, name)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// DeleteTagForSchedule removes tag from schedule
func (a *SchedulerResourceApiService) DeleteTagForSchedule(ctx context.Context, body []model.Tag, name string) (*http.Response, error) {
	// Convert tags
	genTags := toGeneratedTags(body)

	req := a.http_orkes.SchedulerResourceAPI.DeleteTagForSchedule(ctx, name).Tag(genTags)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// getAllSchedulesInternal gets all schedules (internal)
func (a *SchedulerResourceApiService) getAllSchedulesInternal(ctx context.Context, workflowName string) ([]model.SaveScheduleRequest, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.GetAllSchedules(ctx)
	if workflowName != "" {
		req = req.WorkflowName(workflowName)
	}

	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert result - result is []WorkflowScheduleModel
	schedules := make([]model.SaveScheduleRequest, 0)
	for _, scheduleModel := range result {
		schedule := toDomainSaveScheduleRequestFromModel(&scheduleModel)
		schedules = append(schedules, schedule)
	}

	return schedules, resp, nil
}

// getNextFewSchedulesInternal gets next few schedules (internal)
func (a *SchedulerResourceApiService) getNextFewSchedulesInternal(ctx context.Context, cronExpression string, scheduleEndTime int64, scheduleStartTime int64, limit int32) ([]int64, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.GetNextFewSchedules(ctx)
	req = req.CronExpression(cronExpression)
	if scheduleEndTime > 0 {
		req = req.ScheduleEndTime(scheduleEndTime)
	}
	if scheduleStartTime > 0 {
		req = req.ScheduleStartTime(scheduleStartTime)
	}
	if limit > 0 {
		req = req.Limit(limit)
	}

	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	return result, resp, nil
}

// getScheduleInternal gets a specific schedule (internal)
func (a *SchedulerResourceApiService) getScheduleInternal(ctx context.Context, name string) (*model.SaveScheduleRequest, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.GetSchedule(ctx, name)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert result - result is *WorkflowSchedule
	schedule := toDomainSaveScheduleRequestFromWorkflowSchedule(result)

	return &schedule, resp, nil
}

// GetSchedulesByTag - Get schedules by tag
func (a *SchedulerResourceApiService) GetSchedulesByTag(ctx context.Context, tag string) ([]model.WorkflowScheduleModel, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.GetSchedulesByTag(ctx).Tag(tag)
	genSchedules, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert generated models to domain models
	result := toDomainWorkflowScheduleModelsFromGenerated(genSchedules)
	return result, resp, nil
}

// GetTagsForSchedule gets tags for schedule
func (a *SchedulerResourceApiService) GetTagsForSchedule(ctx context.Context, name string) ([]model.Tag, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.GetTagsForSchedule(ctx, name)
	genTags, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert tags back to model
	tags := toDomainTags(genTags)
	return tags, resp, nil
}

// PauseAllSchedules pauses all schedules
func (a *SchedulerResourceApiService) PauseAllSchedules(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.PauseAllSchedules(ctx)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert map[string]map[string]interface{} to map[string]interface{}
	flattened := make(map[string]interface{})
	for k, v := range result {
		flattened[k] = v
	}
	return flattened, resp, nil
}

// PauseSchedule pauses a schedule
func (a *SchedulerResourceApiService) PauseSchedule(ctx context.Context, name string) (interface{}, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.PauseSchedule(ctx, name)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// PutTagForSchedule adds tag to schedule
func (a *SchedulerResourceApiService) PutTagForSchedule(ctx context.Context, body []model.Tag, name string) (*http.Response, error) {
	// Convert tags
	genTags := toGeneratedTags(body)

	req := a.http_orkes.SchedulerResourceAPI.PutTagForSchedule(ctx, name).Tag(genTags)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// ResumeAllSchedules resumes all schedules
func (a *SchedulerResourceApiService) ResumeAllSchedules(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.ResumeAllSchedules(ctx)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert map[string]map[string]interface{} to map[string]interface{}
	flattened := make(map[string]interface{})
	for k, v := range result {
		flattened[k] = v
	}
	return flattened, resp, nil
}

// ResumeSchedule resumes a schedule
func (a *SchedulerResourceApiService) ResumeSchedule(ctx context.Context, name string) (interface{}, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.ResumeSchedule(ctx, name)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// SaveSchedule saves a schedule
func (a *SchedulerResourceApiService) SaveSchedule(ctx context.Context, body model.SaveScheduleRequest) (interface{}, *http.Response, error) {
	// Convert to generated model
	genSchedule := toGeneratedSaveScheduleRequest(&body)

	req := a.http_orkes.SchedulerResourceAPI.SaveSchedule(ctx).SaveScheduleRequest(genSchedule)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// SchedulerSearchOpts defines the options for searching schedules
type SchedulerSearchOpts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

// SearchV2 performs a V2 search
func (a *SchedulerResourceApiService) SearchV2(ctx context.Context, opts *SchedulerSearchOpts) (model.SearchResultWorkflowSchedule, *http.Response, error) {
	// Use the generated client method
	req := a.http_orkes.SchedulerResourceAPI.SearchV2(ctx)

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

	result, resp, err := req.Execute()
	if err != nil {
		return model.SearchResultWorkflowSchedule{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert generated model to domain model
	domainResult := toDomainSearchResultWorkflowScheduleFromGenerated(result)

	return domainResult, resp, nil
}

// GetAllSchedulesByWorkflowName gets schedules by workflow name
func (a *SchedulerResourceApiService) GetAllSchedulesByWorkflowName(ctx context.Context, workflowName string) ([]model.SaveScheduleRequest, *http.Response, error) {
	// Use getAllSchedulesInternal with workflow name filter
	return a.getAllSchedulesInternal(ctx, workflowName)
}

type SchedulerResourceApiGetAllSchedulesOpts struct {
	WorkflowName optional.String
}

// GetAllSchedules with options
func (a *SchedulerResourceApiService) GetAllSchedules(ctx context.Context, optionals *SchedulerResourceApiGetAllSchedulesOpts) ([]model.WorkflowScheduleModel, *http.Response, error) {
	workflowName := ""
	if optionals != nil && optionals.WorkflowName.IsSet() {
		workflowName = optionals.WorkflowName.Value()
	}
	schedules, resp, err := a.getAllSchedulesInternal(ctx, workflowName)
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert SaveScheduleRequest to WorkflowScheduleModel
	result := toDomainWorkflowScheduleModelsFromSaveRequests(schedules)

	return result, resp, nil
}

type SchedulerResourceApiGetNextFewSchedulesOpts struct {
	ScheduleStartTime optional.Int64
	ScheduleEndTime   optional.Int64
	Limit             optional.Int32
}

// GetNextFewSchedules with options
func (a *SchedulerResourceApiService) GetNextFewSchedules(ctx context.Context, cronExpression string, optionals *SchedulerResourceApiGetNextFewSchedulesOpts) ([]int64, *http.Response, error) {
	var scheduleEndTime, scheduleStartTime int64
	var limit int32

	if optionals != nil {
		if optionals.ScheduleEndTime.IsSet() {
			scheduleEndTime = optionals.ScheduleEndTime.Value()
		}
		if optionals.ScheduleStartTime.IsSet() {
			scheduleStartTime = optionals.ScheduleStartTime.Value()
		}
		if optionals.Limit.IsSet() {
			limit = optionals.Limit.Value()
		}
	}

	return a.getNextFewSchedulesInternal(ctx, cronExpression, scheduleEndTime, scheduleStartTime, limit)
}

// GetSchedule interface method
func (a *SchedulerResourceApiService) GetSchedule(ctx context.Context, name string) (model.WorkflowSchedule, *http.Response, error) {
	schedule, resp, err := a.getScheduleInternal(ctx, name)
	if err != nil {
		return model.WorkflowSchedule{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert SaveScheduleRequest to WorkflowSchedule
	result := toDomainWorkflowScheduleFromSaveRequest(schedule)

	return result, resp, nil
}

// RequeueAllExecutionRecords requeues all execution records
func (a *SchedulerResourceApiService) RequeueAllExecutionRecords(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	// Use Orkes generated client RequeueAllExecutionRecords method
	genResult, resp, err := a.http_orkes.SchedulerResourceAPI.RequeueAllExecutionRecords(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert map[string]map[string]interface{} to map[string]interface{}
	convertedResult := make(map[string]interface{})
	for k, v := range genResult {
		convertedResult[k] = v
	}

	return convertedResult, resp, nil
}
