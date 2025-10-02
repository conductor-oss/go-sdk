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

	"github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// TagsApiService is the service for the tags resource.
type TagsApiService struct {
	*APIClient
}

// AddTaskTag adds the tag to the task.
func (a *TagsApiService) AddTaskTag(ctx context.Context, body model.TagObject, taskName string) (interface{}, *http.Response, error) {
	genTag := toGeneratedTagFromTagObject(body)
	result, resp, err := a.http_orkes.TagsAPI.AddTaskTag(ctx, taskName).Tag(genTag).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// AddWorkflowTag adds the tag to the workflow.
func (a *TagsApiService) AddWorkflowTag(ctx context.Context, body model.TagObject, name string) (interface{}, *http.Response, error) {
	genTag := toGeneratedTagFromTagObject(body)
	result, resp, err := a.http_orkes.TagsAPI.AddWorkflowTag(ctx, name).Tag(genTag).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// DeleteTaskTag removes the tag of the task.
func (a *TagsApiService) DeleteTaskTag(ctx context.Context, body model.TagString, taskName string) (interface{}, *http.Response, error) {
	genTag := toGeneratedTagFromTagString(body)
	result, resp, err := a.http_orkes.TagsAPI.DeleteTaskTag(ctx, taskName).Tag(genTag).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// DeleteWorkflowTag removes the tag of the workflow.
func (a *TagsApiService) DeleteWorkflowTag(ctx context.Context, body model.TagObject, name string) (interface{}, *http.Response, error) {
	genTag := toGeneratedTagFromTagObject(body)
	result, resp, err := a.http_orkes.TagsAPI.DeleteWorkflowTag(ctx, name).Tag(genTag).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// GetTags1 list all tags.
func (a *TagsApiService) GetTags(ctx context.Context) ([]model.TagObject, *http.Response, error) {
	genResult, resp, err := a.http_orkes.TagsAPI.GetTags1(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	result := toDomainTagObjectsFromGenerated(genResult)
	return result, resp, nil
}

// GetTaskTags returns all the tags of the task.
func (a *TagsApiService) GetTaskTags(ctx context.Context, taskName string) ([]model.TagObject, *http.Response, error) {
	genResult, resp, err := a.http_orkes.TagsAPI.GetTaskTags(ctx, taskName).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	result := toDomainTagObjectsFromGenerated(genResult)
	return result, resp, nil
}

// GetWorkflowTags Returns all the tags of the workflow.
func (a *TagsApiService) GetWorkflowTags(ctx context.Context, name string) ([]model.TagObject, *http.Response, error) {
	genResult, resp, err := a.http_orkes.TagsAPI.GetWorkflowTags(ctx, name).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	result := toDomainTagObjectsFromGenerated(genResult)
	return result, resp, nil
}

// SetTaskTags Set the tags of the task.
func (a *TagsApiService) SetTaskTags(ctx context.Context, body []model.TagObject, taskName string) (interface{}, *http.Response, error) {
	genTags := make([]orkes.Tag, len(body))
	for i, tag := range body {
		genTags[i] = toGeneratedTagFromTagObject(tag)
	}

	result, resp, err := a.http_orkes.TagsAPI.SetTaskTags(ctx, taskName).Tag(genTags).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// SetWorkflowTags Set the tags of the workflow.
func (a *TagsApiService) SetWorkflowTags(ctx context.Context, body []model.TagObject, name string) (interface{}, *http.Response, error) {
	genTags := make([]orkes.Tag, len(body))
	for i, tag := range body {
		genTags[i] = toGeneratedTagFromTagObject(tag)
	}

	result, resp, err := a.http_orkes.TagsAPI.SetWorkflowTags(ctx, name).Tag(genTags).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}
