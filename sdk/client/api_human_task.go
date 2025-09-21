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
	"github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/human"
)

// HumanTaskApiService
type HumanTaskApiService struct {
	*APIClient
}

// NewHumanTaskApiService creates a new HumanTaskApiService instance
func NewHumanTaskApiService(client *APIClient) *HumanTaskApiService {
	return &HumanTaskApiService{APIClient: client}
}

// HumanTaskApiAssignAndClaimOpts is the optional parameters for the AssignAndClaim method
type HumanTaskApiAssignAndClaimOpts struct {
	OverrideAssignment optional.Bool
	WithTemplate       optional.Bool
}

// HumanTaskApiGetAllTemplatesOpts is the optional parameters for the GetAllTemplates method
type HumanTaskApiGetAllTemplatesOpts struct {
	Name    optional.String
	Version optional.Int32
}

// HumanTaskApiGetTaskOpts is the optional parameters for the GetTask method
type HumanTaskApiGetTaskOpts struct {
	WithTemplate optional.Bool
}

// HumanTaskApiSaveTemplateOpts is the optional parameters for the SaveTemplate method
type HumanTaskApiSaveTemplateOpts struct {
	NewVersion optional.Bool
}

// HumanTaskApiSaveTemplatesOpts is the optional parameters for the SaveTemplates method
type HumanTaskApiSaveTemplatesOpts struct {
	NewVersion optional.Bool
}

// HumanTaskApiSkipTaskOpts is the optional parameters for the SkipTask method
type HumanTaskApiSkipTaskOpts struct {
	Reason optional.String
}

// HumanTaskApiUpdateTaskOutputOpts is the optional parameters for the UpdateTaskOutput method
type HumanTaskApiUpdateTaskOutputOpts struct {
	Complete optional.Bool
}

// HumanTaskApiUpdateTaskOutputByRefOpts is the optional parameters for the UpdateTaskOutputByRef method
type HumanTaskApiUpdateTaskOutputByRefOpts struct {
	Complete  optional.Bool
	Iteration optional.Interface
}

// Search searches human tasks
func (a *HumanTaskApiService) Search(ctx context.Context, body human.HumanTaskSearch) (human.HumanTaskSearchResult, *http.Response, error) {
	// Convert domain HumanTaskSearch to generated HumanTaskSearch
	searchRequest := toGeneratedHumanTaskSearch(body)

	search := a.http_orkes.HumanTaskAPI.Search(ctx).HumanTaskSearch(searchRequest)

	genResult, resp, err := search.Execute()
	if err != nil {
		return human.HumanTaskSearchResult{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert to domain HumanTaskSearchResult
	result := toDomainHumanTaskSearchResult(genResult)
	return result, resp, nil
}

// GetTaskDisplayNames gets list of task display names applicable for the user
func (a *HumanTaskApiService) GetTaskDisplayNames(ctx context.Context, searchType string) ([]string, *http.Response, error) {
	// Use Orkes generated client
	result, resp, err := a.http_orkes.HumanTaskAPI.GetTaskDisplayNames(ctx).
		SearchType(searchType).
		Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// AssignAndClaim assigns and claims a task
func (a *HumanTaskApiService) AssignAndClaim(ctx context.Context, taskId string, userId string, optionals *HumanTaskApiAssignAndClaimOpts) (human.HumanTaskEntry, *http.Response, error) {
	// Use Orkes generated client
	req := a.http_orkes.HumanTaskAPI.AssignAndClaim(ctx, taskId, userId)

	if optionals != nil {
		if optionals.OverrideAssignment.IsSet() {
			req = req.OverrideAssignment(optionals.OverrideAssignment.Value())
		}
		if optionals.WithTemplate.IsSet() {
			req = req.WithTemplate(optionals.WithTemplate.Value())
		}
	}

	genResult, resp, err := req.Execute()
	if err != nil {
		return human.HumanTaskEntry{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert to domain HumanTaskEntry
	result := toDomainHumanTaskEntry(genResult)
	return result, resp, nil
}

// HumanTaskApiClaimTaskOpts is the optional parameters for the ClaimTask method
type HumanTaskApiClaimTaskOpts struct {
	OverrideAssignment optional.Bool
	WithTemplate       optional.Bool
}

// ClaimTask claims a task by authenticated Conductor user
func (a *HumanTaskApiService) ClaimTask(ctx context.Context, taskId string, optionals *HumanTaskApiClaimTaskOpts) (human.HumanTaskEntry, *http.Response, error) {
	// Use Orkes generated client
	req := a.http_orkes.HumanTaskAPI.ClaimTask(ctx, taskId)

	if optionals != nil {
		if optionals.OverrideAssignment.IsSet() {
			req = req.OverrideAssignment(optionals.OverrideAssignment.Value())
		}
		if optionals.WithTemplate.IsSet() {
			req = req.WithTemplate(optionals.WithTemplate.Value())
		}
	}

	genResult, resp, err := req.Execute()
	if err != nil {
		return human.HumanTaskEntry{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert to domain HumanTaskEntry
	result := toDomainHumanTaskEntry(genResult)
	return result, resp, nil
}

// GetTask gets a task
func (a *HumanTaskApiService) GetTask(ctx context.Context, taskId string, optionals *HumanTaskApiGetTaskOpts) (human.HumanTaskEntry, *http.Response, error) {
	// Use Orkes generated client (method is called GetTask1 in generated client)
	req := a.http_orkes.HumanTaskAPI.GetTask1(ctx, taskId)
	if optionals != nil && optionals.WithTemplate.IsSet() {
		req = req.WithTemplate(optionals.WithTemplate.Value())
	}
	genResult, resp, err := req.Execute()
	if err != nil {
		return human.HumanTaskEntry{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert to domain HumanTaskEntry
	result := toDomainHumanTaskEntry(genResult)
	return result, resp, nil
}

// GetConductorTaskById gets Conductor task by id (for human tasks only)
func (a *HumanTaskApiService) GetConductorTaskById(ctx context.Context, taskId string) (model.Task, *http.Response, error) {
	// Use Orkes generated client
	genResult, resp, err := a.http_orkes.HumanTaskResourceAPI.GetConductorTaskById(ctx, taskId).Execute()
	if err != nil {
		return model.Task{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert using existing mapper
	result := toDomainTask(genResult)
	return result, resp, nil
}

// UpdateTaskOutput updates task output, optionally complete
func (a *HumanTaskApiService) UpdateTaskOutput(ctx context.Context, body map[string]interface{}, taskId string, optionals *HumanTaskApiUpdateTaskOutputOpts) (*http.Response, error) {
	// Use Orkes generated client
	req := a.http_orkes.HumanTaskAPI.UpdateTaskOutput(ctx, taskId).RequestBody(body)

	if optionals != nil {
		if optionals.Complete.IsSet() {
			req = req.Complete(optionals.Complete.Value())
		}
	}

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// UpdateTaskOutputByRef updates task output by reference, optionally complete
func (a *HumanTaskApiService) UpdateTaskOutputByRef(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, optionals *HumanTaskApiUpdateTaskOutputByRefOpts) (*http.Response, error) {
	// Use Orkes generated client
	req := a.http_orkes.HumanTaskAPI.UpdateTaskOutputByRef(ctx).
		WorkflowId(workflowId).
		TaskRefName(taskRefName).
		RequestBody(body)

	if optionals != nil {
		if optionals.Complete.IsSet() {
			req = req.Complete(optionals.Complete.Value())
		}
		if optionals.Iteration.IsSet() {
			if iteration, ok := optionals.Iteration.Value().([]int32); ok {
				req = req.Iteration(iteration)
			}
		}
	}

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// ReleaseTask releases a task without completing it
func (a *HumanTaskApiService) ReleaseTask(ctx context.Context, taskId string) (*http.Response, error) {
	resp, err := a.http_orkes.HumanTaskAPI.ReleaseTask(ctx, taskId).Execute()
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ReassignTask reassigns a task without completing it
func (a *HumanTaskApiService) ReassignTask(ctx context.Context, body []human.HumanTaskAssignment, taskId string) (*http.Response, error) {
	// Convert domain HumanTaskAssignment slice to generated slice
	assignments := make([]orkes.HumanTaskAssignment, 0, len(body))
	for _, assignment := range body {
		assignments = append(assignments, orkes.HumanTaskAssignment{
			Assignee: &orkes.HumanTaskUser{
				User:     ToPointer(assignment.Assignee.User),
				UserType: ToPointer(assignment.Assignee.UserType),
			},
		})
	}

	// Use Orkes generated client
	req := a.http_orkes.HumanTaskAPI.ReassignTask(ctx, taskId).HumanTaskAssignment(assignments)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// SkipTask skips assignment and moves to the next assignee
func (a *HumanTaskApiService) SkipTask(ctx context.Context, taskId string, optionals *HumanTaskApiSkipTaskOpts) (*http.Response, error) {
	// Use Orkes generated client
	req := a.http_orkes.HumanTaskAPI.SkipTask(ctx, taskId)

	if optionals != nil {
		if optionals.Reason.IsSet() {
			req = req.Reason(optionals.Reason.Value())
		}
	}

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DeleteTaskFromHumanTaskRecords deletes task from human task records
func (a *HumanTaskApiService) DeleteTaskFromHumanTaskRecords(ctx context.Context, taskIds []string) (*http.Response, error) {
	req := a.http_orkes.HumanTaskAPI.DeleteTaskFromHumanTaskRecords(ctx).RequestBody(taskIds)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DeleteTaskFromHumanTaskRecordsByTaskId deletes single task from human task records
func (a *HumanTaskApiService) DeleteTaskFromHumanTaskRecordsByTaskId(ctx context.Context, taskId string) (*http.Response, error) {
	resp, err := a.http_orkes.HumanTaskAPI.DeleteTaskFromHumanTaskRecords1(ctx, taskId).Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DeleteTaskFromHumanTaskRecords deletes single task from human task records
func (a *HumanTaskApiService) DeleteTaskFromHumanTaskRecord(ctx context.Context, taskId string) (*http.Response, error) {
	return a.DeleteTaskFromHumanTaskRecordsByTaskId(ctx, taskId)
}

// DeleteTemplateByName deletes template by name
func (a *HumanTaskApiService) DeleteTemplateByName(ctx context.Context, name string) (*http.Response, error) {
	resp, err := a.http_orkes.HumanTaskAPI.DeleteTemplateByName(ctx, name).Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DeleteTemplatesByNameAndVersion deletes templates by name and version
func (a *HumanTaskApiService) DeleteTemplatesByNameAndVersion(ctx context.Context, name string, version int32) (*http.Response, error) {
	resp, err := a.http_orkes.HumanTaskAPI.DeleteTemplatesByNameAndVersion(ctx, name, version).Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// GetAllTemplates gets all templates
func (a *HumanTaskApiService) GetAllTemplates(ctx context.Context, optionals *HumanTaskApiGetAllTemplatesOpts) ([]human.HumanTaskSearch, *http.Response, error) {
	// Use Orkes generated client
	req := a.http_orkes.HumanTaskAPI.GetAllTemplates(ctx)

	if optionals != nil {
		if optionals.Name.IsSet() {
			req = req.Name(optionals.Name.Value())
		}
		if optionals.Version.IsSet() {
			req = req.Version(optionals.Version.Value())
		}
	}

	genResult, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert to domain HumanTaskSearch slice using mapper
	result := toDomainHumanTaskTemplates(genResult)
	return result, resp, nil
}

// GetTemplateByNameAndVersion gets template by name and version
func (a *HumanTaskApiService) GetTemplateByNameAndVersion(ctx context.Context, name string, version int32) (human.HumanTaskSearch, *http.Response, error) {
	// Use Orkes generated client
	genResult, resp, err := a.http_orkes.HumanTaskAPI.GetTemplateByNameAndVersion(ctx, name, version).Execute()
	if err != nil {
		return human.HumanTaskSearch{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert to domain HumanTaskSearch using proper mapper
	result := toDomainHumanTaskSearchFromTemplate(genResult)
	return result, resp, nil
}

// GetTemplateByTaskId gets template by task id
func (a *HumanTaskApiService) GetTemplateByTaskId(ctx context.Context, humanTaskId string) (human.HumanTaskSearch, *http.Response, error) {
	// Use Orkes generated client
	genResult, resp, err := a.http_orkes.HumanTaskAPI.GetTemplateByTaskId(ctx, humanTaskId).Execute()
	if err != nil {
		return human.HumanTaskSearch{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert to domain HumanTaskSearch using proper mapper
	result := toDomainHumanTaskSearchFromTemplate(genResult)
	return result, resp, nil
}

// SaveTemplate saves a template
func (a *HumanTaskApiService) SaveTemplate(ctx context.Context, body human.HumanTaskSearch, optionals *HumanTaskApiSaveTemplateOpts) (human.HumanTaskSearch, *http.Response, error) {
	// Convert domain HumanTaskSearch to generated HumanTaskTemplate using proper mapper
	templateRequest := toGeneratedHumanTaskTemplateFromSearch(body)

	// Use Orkes generated client
	req := a.http_orkes.HumanTaskAPI.SaveTemplate(ctx).HumanTaskTemplate(templateRequest)

	if optionals != nil {
		if optionals.NewVersion.IsSet() {
			req = req.NewVersion(optionals.NewVersion.Value())
		}
	}

	genResult, resp, err := req.Execute()
	if err != nil {
		return human.HumanTaskSearch{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert result back to domain HumanTaskSearch using proper mapper
	result := toDomainHumanTaskSearchFromTemplate(genResult)
	return result, resp, nil
}

// SaveTemplates saves multiple templates
func (a *HumanTaskApiService) SaveTemplates(ctx context.Context, body []human.HumanTaskSearch, optionals *HumanTaskApiSaveTemplatesOpts) ([]human.HumanTaskSearch, *http.Response, error) {
	// Convert domain HumanTaskSearch slice to generated HumanTaskTemplate slice using proper mapper
	templateRequests := make([]orkes.HumanTaskTemplate, 0, len(body))
	for _, search := range body {
		templateRequests = append(templateRequests, toGeneratedHumanTaskTemplateFromSearch(search))
	}

	// Use Orkes generated client
	req := a.http_orkes.HumanTaskAPI.SaveTemplates(ctx).HumanTaskTemplate(templateRequests)

	if optionals != nil {
		if optionals.NewVersion.IsSet() {
			req = req.NewVersion(optionals.NewVersion.Value())
		}
	}

	genResult, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert result back to domain HumanTaskSearch slice using proper mapper
	result := toDomainHumanTaskTemplates(genResult)
	return result, resp, nil
}

// BackPopulateFullTextIndex back populates full text index
func (a *HumanTaskApiService) BackPopulateFullTextIndex(ctx context.Context, var100 int32) (map[string]interface{}, *http.Response, error) {
	// Use Orkes generated client
	genResult, resp, err := a.http_orkes.HumanTaskAPI.BackPopulateFullTextIndex(ctx).Var100(var100).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	return genResult, resp, nil
}
