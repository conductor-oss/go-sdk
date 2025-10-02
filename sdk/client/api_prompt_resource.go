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

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
)

// PromptResourceApiService is the service for the prompt resource.
type PromptResourceApiService struct {
	*APIClient
}

// CreateMessageTemplates creates message templates in bulk.
func (a *PromptResourceApiService) CreateMessageTemplates(ctx context.Context, templates []integration.PromptTemplate) (*http.Response, error) {
	genTemplates := toGeneratedPromptTemplates(templates)
	req := a.http_orkes.PromptResourceAPI.CreateMessageTemplates(ctx).MessageTemplate(genTemplates)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DeleteMessageTemplate deletes a message template.
func (a *PromptResourceApiService) DeleteMessageTemplate(ctx context.Context, name string) (*http.Response, error) {
	req := a.http_orkes.PromptResourceAPI.DeleteMessageTemplate(ctx, name)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DeleteTagForPromptTemplate deletes tags for a prompt template.
func (a *PromptResourceApiService) DeleteTagForPromptTemplate(ctx context.Context, tags []model.Tag, name string) (*http.Response, error) {
	// Convert domain tags to generated tags
	genTags := toGeneratedTags(tags)

	req := a.http_orkes.PromptResourceAPI.DeleteTagForPromptTemplate(ctx, name).Tag(genTags)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// GetMessageTemplate gets a message template by name.
func (a *PromptResourceApiService) GetMessageTemplate(ctx context.Context, name string) (*integration.PromptTemplate, *http.Response, error) {
	req := a.http_orkes.PromptResourceAPI.GetMessageTemplate(ctx, name)

	genTemplate, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainPromptTemplate(genTemplate)
	return &result, resp, nil
}

// GetMessageTemplates gets all message templates.
func (a *PromptResourceApiService) GetMessageTemplates(ctx context.Context) ([]integration.PromptTemplate, *http.Response, error) {
	req := a.http_orkes.PromptResourceAPI.GetMessageTemplates(ctx)

	genTemplates, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainPromptTemplates(genTemplates)
	return result, resp, nil
}

// GetTagsForPromptTemplate gets tags for a prompt template.
func (a *PromptResourceApiService) GetTagsForPromptTemplate(ctx context.Context, name string) ([]model.Tag, *http.Response, error) {
	req := a.http_orkes.PromptResourceAPI.GetTagsForPromptTemplate(ctx, name)

	genTags, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert generated tags to domain tags
	result := toDomainTags(genTags)

	return result, resp, nil
}

// PutTagForPromptTemplate updates tags for a prompt template.
func (a *PromptResourceApiService) PutTagForPromptTemplate(ctx context.Context, tags []model.Tag, name string) (*http.Response, error) {
	// Convert domain tags to generated tags
	genTags := toGeneratedTags(tags)

	req := a.http_orkes.PromptResourceAPI.PutTagForPromptTemplate(ctx, name).Tag(genTags)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// PromptResourceApiSaveMessageTemplateOpts - Optional parameters for SaveMessageTemplate.
type PromptResourceApiSaveMessageTemplateOpts struct {
	Models []string
}

// SaveMessageTemplate saves a message template.
func (a *PromptResourceApiService) SaveMessageTemplate(ctx context.Context, templateText string, description string, name string, optionals *PromptResourceApiSaveMessageTemplateOpts) (*http.Response, error) {
	// Use SaveMessageTemplate which has the correct method signature
	req := a.http_orkes.PromptResourceAPI.SaveMessageTemplate(ctx, name).
		Description(description).
		Body(templateText)

	// Apply optional parameters if provided
	if optionals != nil && len(optionals.Models) > 0 {
		req = req.Models(optionals.Models)
	}

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// TestMessageTemplate tests a message template.
func (a *PromptResourceApiService) TestMessageTemplate(ctx context.Context, request model.PromptTemplateTestRequest) (string, *http.Response, error) {
	// Convert domain model to generated model
	genRequest := toGeneratedPromptTemplateTestRequest(&request)

	req := a.http_orkes.PromptResourceAPI.TestMessageTemplate(ctx).PromptTemplateTestRequest(*genRequest)

	result, resp, err := req.Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}

	return result, resp, nil
}
