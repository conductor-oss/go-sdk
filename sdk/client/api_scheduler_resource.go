// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package client

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// SchedulerResourceApiService is the service for the scheduler resource.
type SchedulerResourceApiService struct {
	*APIClient
}

// DeleteSchedule deletes a schedule.
func (a *SchedulerResourceApiService) DeleteSchedule(ctx context.Context, name string) (interface{}, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.DeleteSchedule(ctx, name)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// DeleteTagForSchedule removes tag from schedule.
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

// SchedulerResourceApiGetAllSchedulesOpts - Optional parameters for GetAllSchedules.
type SchedulerResourceApiGetAllSchedulesOpts struct {
	WorkflowName optional.String
}

// GetAllSchedules with options.
func (a *SchedulerResourceApiService) GetAllSchedules(ctx context.Context, optionals *SchedulerResourceApiGetAllSchedulesOpts) ([]model.WorkflowScheduleModel, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.GetAllSchedules(ctx)
	if optionals != nil && optionals.WorkflowName.IsSet() {
		req = req.WorkflowName(optionals.WorkflowName.Value())
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

	// Convert SaveSchedule	Request to WorkflowScheduleModel
	workflowScheduleModels := toDomainWorkflowScheduleModelsFromSaveRequests(schedules)

	return workflowScheduleModels, resp, nil
}

// SchedulerResourceApiGetNextFewSchedulesOpts - Optional parameters for GetNextFewSchedules.
type SchedulerResourceApiGetNextFewSchedulesOpts struct {
	ScheduleStartTime optional.Int64
	ScheduleEndTime   optional.Int64
	Limit             optional.Int32
}

// GetNextFewSchedules with options.
func (a *SchedulerResourceApiService) GetNextFewSchedules(ctx context.Context, cronExpression string, optionals *SchedulerResourceApiGetNextFewSchedulesOpts) ([]int64, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.GetNextFewSchedules(ctx)
	req = req.CronExpression(cronExpression)

	if optionals != nil {
		if optionals.ScheduleEndTime.IsSet() {
			req = req.ScheduleEndTime(optionals.ScheduleEndTime.Value())
		}
		if optionals.ScheduleStartTime.IsSet() {
			req = req.ScheduleStartTime(optionals.ScheduleStartTime.Value())
		}
		if optionals.Limit.IsSet() {
			req = req.Limit(optionals.Limit.Value())
		}
	}

	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	return result, resp, nil
}

// GetSchedule interface method.
func (a *SchedulerResourceApiService) GetSchedule(ctx context.Context, name string) (model.WorkflowSchedule, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.GetSchedule(ctx, name)
	result, resp, err := req.Execute()
	if err != nil {
		return model.WorkflowSchedule{}, resp, wrapGeneratedError(err, resp)
	}

	schedule := toDomainWorkflowScheduleFromWorkflowSchedule(result)

	return schedule, resp, nil
}

// GetTagsForSchedule gets tags for schedule.
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

// PauseAllSchedules pauses all schedules.
func (a *SchedulerResourceApiService) PauseAllSchedules(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.PauseAllSchedules(ctx)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	return result, resp, nil
}

// PauseSchedule pauses a schedule.
func (a *SchedulerResourceApiService) PauseSchedule(ctx context.Context, name string) (interface{}, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.PauseSchedule(ctx, name)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// PutTagForSchedule adds tag to schedule.
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

// RequeueAllExecutionRecords requeues all execution records.
func (a *SchedulerResourceApiService) RequeueAllExecutionRecords(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	genResult, resp, err := a.http_orkes.SchedulerResourceAPI.RequeueAllExecutionRecords(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	return genResult, resp, nil
}

// ResumeAllSchedules resumes all schedules.
func (a *SchedulerResourceApiService) ResumeAllSchedules(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.ResumeAllSchedules(ctx)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	return result, resp, nil
}

// ResumeSchedule resumes a schedule.
func (a *SchedulerResourceApiService) ResumeSchedule(ctx context.Context, name string) (interface{}, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.ResumeSchedule(ctx, name)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// SaveSchedule saves a schedule.
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

// SearchV2 performs a V2 search.
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

	domainResult := toDomainSearchResultWorkflowScheduleFromGenerated(result)
	return domainResult, resp, nil
}

// GetSchedulesByTag gets schedules by tag.
func (a *SchedulerResourceApiService) GetSchedulesByTag(ctx context.Context, tag string) ([]model.WorkflowScheduleModel, *http.Response, error) {
	req := a.http_orkes.SchedulerResourceAPI.GetSchedulesByTag(ctx).Tag(tag)
	genSchedules, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	result := toDomainWorkflowScheduleModelsFromGenerated(genSchedules)
	return result, resp, nil
}
