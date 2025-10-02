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
)

// MetadataResourceApiService is the service for the metadata resource.
type MetadataResourceApiService struct {
	*APIClient
}

// RegisterWorkflowDef registers workflow definition.
func (a *MetadataResourceApiService) RegisterWorkflowDef(ctx context.Context, overwrite bool, body model.WorkflowDef) (*http.Response, error) {
	extDef := toExtendedWorkflowDef(&body)

	req := a.http_orkes.MetadataResourceAPI.Create(ctx).ExtendedWorkflowDef(extDef)
	if overwrite {
		req = req.Overwrite(overwrite)
	}

	_, resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// RegisterWorkflowDefWithTags registers workflow definition with tags.
func (a *MetadataResourceApiService) RegisterWorkflowDefWithTags(ctx context.Context, overwrite bool, body model.WorkflowDef, tags []model.MetadataTag) (*http.Response, error) {
	// Convert domain model to extended workflow def with tags
	extDef := toExtendedWorkflowDefWithTags(&body, tags, true)

	req := a.http_orkes.MetadataResourceAPI.Create(ctx).Overwrite(overwrite).ExtendedWorkflowDef(extDef)

	_, resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// RegisterTaskDefWithTags registers task definition with tags.
func (a *MetadataResourceApiService) RegisterTaskDefWithTags(ctx context.Context, body model.TaskDef, tags []model.MetadataTag) (*http.Response, error) {
	// Convert domain model to extended task def with tags
	extDef := toExtendedTaskDefWithTags(&body, tags, true)

	req := a.http_orkes.MetadataResourceAPI.RegisterTaskDef(ctx).ExtendedTaskDef([]orkes.ExtendedTaskDef{extDef})

	_, resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// UpdateTaskDefWithTags updates task definition with tags.
func (a *MetadataResourceApiService) UpdateTaskDefWithTags(ctx context.Context, body model.TaskDef, tags []model.MetadataTag, overwriteTags bool) (*http.Response, error) {
	// Convert domain model to extended task def with tags
	extDef := toExtendedTaskDefWithTags(&body, tags, overwriteTags)

	req := a.http_orkes.MetadataResourceAPI.UpdateTaskDef(ctx).ExtendedTaskDef(extDef)

	_, resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// UpdateWorkflowDefWithTags updates workflow definition with tags.
func (a *MetadataResourceApiService) UpdateWorkflowDefWithTags(ctx context.Context, body model.WorkflowDef, tags []model.MetadataTag, overwriteTags bool) (*http.Response, error) {
	// Convert domain model to extended workflow def with tags
	extDef := toExtendedWorkflowDefWithTags(&body, tags, overwriteTags)

	req := a.http_orkes.MetadataResourceAPI.Update(ctx).ExtendedWorkflowDef([]orkes.ExtendedWorkflowDef{extDef})

	_, resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// GetTagsForWorkflowDef gets tags for workflow definition.
func (a *MetadataResourceApiService) GetTagsForWorkflowDef(ctx context.Context, name string) ([]model.MetadataTag, error) {
	req := a.http_orkes.TagsAPI.GetWorkflowTags(ctx, name)

	genTags, resp, err := req.Execute()
	if err != nil {
		return nil, wrapGeneratedError(err, resp)
	}

	// Convert generated tags to domain metadata tags
	result := toDomainMetadataTagsFromGenerated(genTags)
	return result, nil
}

// GetTagsForTaskDef gets tags for task definition.
func (a *MetadataResourceApiService) GetTagsForTaskDef(ctx context.Context, tasktype string) ([]model.MetadataTag, error) {
	req := a.http_orkes.TagsAPI.GetTaskTags(ctx, tasktype)

	genTags, resp, err := req.Execute()
	if err != nil {
		return nil, wrapGeneratedError(err, resp)
	}

	// Convert generated tags to domain metadata tags
	result := toDomainMetadataTagsFromGenerated(genTags)
	return result, nil
}

// SetTaskTags sets all tags for task definition.
func (a *MetadataResourceApiService) SetTaskTags(ctx context.Context, taskName string, tags []model.MetadataTag) (*http.Response, error) {
	genTags := toGeneratedTagsFromMetadataTags(tags)

	req := a.http_orkes.TagsAPI.SetTaskTags(ctx, taskName).Tag(genTags)

	_, resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// MetadataResourceApiGetOpts optional parameters.
type MetadataResourceApiGetOpts struct {
	Version optional.Int32
}

// Get gets workflow definition.
func (a *MetadataResourceApiService) Get(ctx context.Context, name string, localVarOptionals *MetadataResourceApiGetOpts) (model.WorkflowDef, *http.Response, error) {
	req := a.http_orkes.MetadataResourceAPI.Get1(ctx, name)

	// Apply optional parameters
	if localVarOptionals != nil && localVarOptionals.Version.IsSet() {
		req = req.Version(localVarOptionals.Version.Value())
	}

	genDef, resp, err := req.Execute()
	if err != nil {
		return model.WorkflowDef{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert to model
	def := toDomainWorkflowDefFromGenerated(genDef)
	return def, resp, nil
}

// GetAll gets all workflow definitions.
func (a *MetadataResourceApiService) GetAll(ctx context.Context) ([]model.WorkflowDef, *http.Response, error) {
	req := a.http_orkes.MetadataResourceAPI.GetWorkflowDefs(ctx)
	genDefs, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert all definitions
	defs := make([]model.WorkflowDef, len(genDefs))
	for i, genDef := range genDefs {
		converted := toDomainWorkflowDefFromGenerated(&genDef)
		defs[i] = converted
	}

	return defs, resp, nil
}

// GetTaskDef gets task definition.
func (a *MetadataResourceApiService) GetTaskDef(ctx context.Context, tasktype string) (model.TaskDef, *http.Response, error) {
	req := a.http_orkes.MetadataResourceAPI.GetTaskDef(ctx, tasktype)
	result, resp, err := req.Execute()
	if err != nil {
		return model.TaskDef{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	def := toDomainTaskDef(result)
	return def, resp, nil
}

// GetTaskDefs gets all task definitions.
func (a *MetadataResourceApiService) GetTaskDefs(ctx context.Context) ([]model.TaskDef, *http.Response, error) {
	req := a.http_orkes.MetadataResourceAPI.GetTaskDefs(ctx)
	genDefs, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert all definitions
	defs := make([]model.TaskDef, len(genDefs))
	for i, genDef := range genDefs {
		defs[i] = toDomainTaskDefPtr(&genDef)
	}

	return defs, resp, nil
}

// RegisterTaskDef registers task definitions.
func (a *MetadataResourceApiService) RegisterTaskDef(ctx context.Context, body []model.TaskDef) (*http.Response, error) {
	// Convert to ExtendedTaskDef models
	genDefs := make([]orkes.ExtendedTaskDef, len(body))
	for i, def := range body {
		extTaskDef := toExtendedTaskDef(&def)
		genDefs[i] = extTaskDef
	}

	req := a.http_orkes.MetadataResourceAPI.RegisterTaskDef(ctx).ExtendedTaskDef(genDefs)
	_, resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// UnregisterTaskDef unregisters task definition.
func (a *MetadataResourceApiService) UnregisterTaskDef(ctx context.Context, tasktype string) (*http.Response, error) {
	req := a.http_orkes.MetadataResourceAPI.UnregisterTaskDef(ctx, tasktype)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// UnregisterWorkflowDef unregisters workflow definition.
func (a *MetadataResourceApiService) UnregisterWorkflowDef(ctx context.Context, name string, version int32) (*http.Response, error) {
	req := a.http_orkes.MetadataResourceAPI.UnregisterWorkflowDef(ctx, name, version)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// Update updates an existing workflow.
func (a *MetadataResourceApiService) Update(ctx context.Context, body []model.WorkflowDef) (*http.Response, error) {
	// Convert to ExtendedWorkflowDef models
	genDefs := make([]orkes.ExtendedWorkflowDef, len(body))
	for i, def := range body {
		extDef := toExtendedWorkflowDef(&def)
		genDefs[i] = extDef
	}

	req := a.http_orkes.MetadataResourceAPI.Update(ctx).ExtendedWorkflowDef(genDefs)
	_, resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// UpdateTaskDef updates an existing task.
func (a *MetadataResourceApiService) UpdateTaskDef(ctx context.Context, body model.TaskDef) (*http.Response, error) {
	// Convert to ExtendedTaskDef
	extDef := toExtendedTaskDef(&body)

	req := a.http_orkes.MetadataResourceAPI.UpdateTaskDef(ctx).ExtendedTaskDef(extDef)
	_, resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}
