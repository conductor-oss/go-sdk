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
	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
)

// IntegrationResourceApiService wraps the generated client to maintain backward compatibility
type IntegrationResourceApiService struct {
	// Embedded for backward compatibility with helper methods
	*APIClient
}

// NewIntegrationResourceApiServicecreates a service
func NewIntegrationResourceApiService(apiClient *APIClient) *IntegrationResourceApiService {
	if apiClient == nil {
		return nil
	}

	return &IntegrationResourceApiService{
		APIClient: apiClient,
	}
}

// GetIntegrationProvidersOpts - Optional parameters for GetIntegrationProviders
type GetIntegrationProvidersOpts struct {
	Category   optional.String
	ActiveOnly optional.Bool
}

// GetIntegrationProviders - Get all integration providers
func (a *IntegrationResourceApiService) GetIntegrationProviders(ctx context.Context, opts *GetIntegrationProvidersOpts) ([]integration.Integration, *http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.GetIntegrationProviders(ctx)
	if opts != nil {
		if opts.Category.IsSet() {
			req = req.Category(opts.Category.Value())
		}
		if opts.ActiveOnly.IsSet() {
			req = req.ActiveOnly(opts.ActiveOnly.Value())
		}
	}
	genIntegrations, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert generated response to domain models
	result := make([]integration.Integration, len(genIntegrations))
	for i, genIntegration := range genIntegrations {
		result[i] = toDomainIntegration(genIntegration)
	}

	return result, resp, nil
}

// GetIntegrationProvider - Get integration provider by name
func (a *IntegrationResourceApiService) GetIntegrationProvider(ctx context.Context, name string) (integration.Integration, *http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.GetIntegrationProvider(ctx, name)

	genIntegration, resp, err := req.Execute()
	if err != nil {
		return integration.Integration{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert raw response to domain model
	result := toDomainIntegration(*genIntegration)
	return result, resp, nil
}

// SaveIntegrationProvider - Save integration provider
func (a *IntegrationResourceApiService) SaveIntegrationProvider(ctx context.Context, update integration.IntegrationUpdate, name string) (*http.Response, error) {
	// Convert domain model to generated model using mapper
	genUpdate := toGeneratedIntegrationUpdate(update)

	req := a.APIClient.http_orkes.IntegrationResourceAPI.SaveIntegrationProvider(ctx, name).IntegrationUpdate(genUpdate)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DeleteIntegrationProvider - Delete integration provider
func (a *IntegrationResourceApiService) DeleteIntegrationProvider(ctx context.Context, name string) (*http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.DeleteIntegrationProvider(ctx, name)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// GetIntegrationProviderDefs - Get all integration provider definitions
func (a *IntegrationResourceApiService) GetIntegrationProviderDefs(ctx context.Context) ([]model.IntegrationDef, *http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.GetIntegrationProviderDefs(ctx)

	genDefs, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert generated models to domain models
	result := make([]model.IntegrationDef, len(genDefs))
	for i, genDef := range genDefs {
		domainDef := toDomainIntegrationDef(&genDef)
		result[i] = domainDef
	}

	return result, resp, nil
}

// GetTagsForIntegrationProvider - Get tags for integration provider
func (a *IntegrationResourceApiService) GetTagsForIntegrationProvider(ctx context.Context, name string) ([]model.TagObject, *http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.GetTagsForIntegrationProvider(ctx, name)

	tags, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert tags to TagObject
	result := make([]model.TagObject, len(tags))
	for i, tag := range tags {
		result[i] = model.TagObject{
			Key:   GetPointerValue(tag.Key, ""),
			Value: GetPointerValue(tag.Value, ""),
			Type:  GetPointerValue(tag.Type, ""),
		}
	}

	return result, resp, nil
}

// UpdateTagForIntegrationProvider - Update tags for integration provider
func (a *IntegrationResourceApiService) UpdateTagForIntegrationProvider(ctx context.Context, tags []model.TagObject, name string) (*http.Response, error) {
	// Convert domain tags to generated tags
	genTags := make([]orkes.Tag, len(tags))
	for i, tag := range tags {
		genTags[i] = orkes.Tag{
			Key:   ToPointer(tag.Key),
			Value: ToPointer(tag.Value),
		}
	}

	req := a.APIClient.http_orkes.IntegrationResourceAPI.PutTagForIntegrationProvider(ctx, name).Tag(genTags)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DeleteTagForIntegrationProvider - Delete tags for integration provider
func (a *IntegrationResourceApiService) DeleteTagForIntegrationProvider(ctx context.Context, tags []model.TagObject, name string) (*http.Response, error) {
	// Convert domain tags to generated tags
	genTags := make([]orkes.Tag, len(tags))
	for i, tag := range tags {
		genTags[i] = orkes.Tag{
			Key:   ToPointer(tag.Key),
			Value: ToPointer(tag.Value),
			Type:  ToPointer(tag.Type),
		}
	}

	req := a.APIClient.http_orkes.IntegrationResourceAPI.DeleteTagForIntegrationProvider(ctx, name).Tag(genTags)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// GetTagsForIntegration - Get tags for integration
func (a *IntegrationResourceApiService) GetTagsForIntegration(ctx context.Context, name string, modelName string) ([]model.TagObject, *http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.GetTagsForIntegration(ctx, name, modelName)

	tags, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert tags to TagObject
	result := make([]model.TagObject, len(tags))
	for i, tag := range tags {
		result[i] = model.TagObject{
			Key:   GetPointerValue(tag.Key, ""),
			Value: GetPointerValue(tag.Value, ""),
		}
	}

	return result, resp, nil
}

// UpdateTagForIntegration - Update tags for integration
func (a *IntegrationResourceApiService) UpdateTagForIntegration(ctx context.Context, tags []model.TagObject, name string, modelName string) (*http.Response, error) {
	// Convert domain tags to generated tags
	genTags := make([]orkes.Tag, len(tags))
	for i, tag := range tags {
		genTags[i] = orkes.Tag{
			Key:   ToPointer(tag.Key),
			Value: ToPointer(tag.Value),
		}
	}

	req := a.APIClient.http_orkes.IntegrationResourceAPI.PutTagForIntegration(ctx, name, modelName).Tag(genTags)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DeleteTagForIntegration - Delete tags for integration
func (a *IntegrationResourceApiService) DeleteTagForIntegration(ctx context.Context, tags []model.TagObject, name string, modelName string) (*http.Response, error) {
	// Convert domain tags to generated tags
	genTags := make([]orkes.Tag, len(tags))
	for i, tag := range tags {
		genTags[i] = orkes.Tag{
			Key:   ToPointer(tag.Key),
			Value: ToPointer(tag.Value),
			Type:  ToPointer(tag.Type),
		}
	}

	req := a.APIClient.http_orkes.IntegrationResourceAPI.DeleteTagForIntegration(ctx, name, modelName).Tag(genTags)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// GetIntegrationApis - Get integration APIs
func (a *IntegrationResourceApiService) GetIntegrationApis(ctx context.Context, name string, activeOnly optional.Bool) ([]integration.IntegrationApi, *http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.GetIntegrationApis(ctx, name)

	if activeOnly.IsSet() {
		req = req.ActiveOnly(activeOnly.Value())
	}

	genApis, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainIntegrationApis(genApis)
	return result, resp, nil
}

// GetIntegrationApi - Get integration API by name and model
func (a *IntegrationResourceApiService) GetIntegrationApi(ctx context.Context, name string, modelName string) (integration.IntegrationApi, *http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.GetIntegrationApi(ctx, name, modelName)

	genApi, resp, err := req.Execute()
	if err != nil {
		return integration.IntegrationApi{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainIntegrationApi(*genApi)
	return result, resp, nil
}

// SaveIntegrationApi - Save integration API
func (a *IntegrationResourceApiService) SaveIntegrationApi(ctx context.Context, update integration.IntegrationApiUpdate, name string, modelName string) (*http.Response, error) {
	// Convert domain model to generated model using mapper
	genUpdate := toGeneratedIntegrationApiUpdate(update)

	req := a.APIClient.http_orkes.IntegrationResourceAPI.SaveIntegrationApi(ctx, name, modelName).IntegrationApiUpdate(genUpdate)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DeleteIntegrationApi - Delete integration API
func (a *IntegrationResourceApiService) DeleteIntegrationApi(ctx context.Context, name string, modelName string) (*http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.DeleteIntegrationApi(ctx, name, modelName)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// GetPromptsWithIntegration - Get prompts associated with integration
func (a *IntegrationResourceApiService) GetPromptsWithIntegration(ctx context.Context, integrationName string, modelName string) ([]integration.PromptTemplate, *http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.GetPromptsWithIntegration(ctx, integrationName, modelName)

	genTemplates, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainPromptTemplates(genTemplates)
	return result, resp, nil
}

// AssociatePromptWithIntegration - Associate prompt with integration
func (a *IntegrationResourceApiService) AssociatePromptWithIntegration(ctx context.Context, integrationName string, modelName string, promptName string) (*http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.AssociatePromptWithIntegration(ctx, integrationName, modelName, promptName)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// GetTokenUsageForIntegration - Get token usage for integration
func (a *IntegrationResourceApiService) GetTokenUsageForIntegration(ctx context.Context, integrationName string, modelName string) (int32, *http.Response, error) {
	req := a.APIClient.http_orkes_v4.IntegrationResourceAPI.GetTokenUsageForIntegration(ctx, integrationName, modelName)

	usage, resp, err := req.Execute()
	if err != nil {
		return 0, resp, wrapGeneratedError(err, resp)
	}

	return usage, resp, nil
}

// GetTokenUsageForIntegrationProvider - Get token usage for integration provider
func (a *IntegrationResourceApiService) GetTokenUsageForIntegrationProvider(ctx context.Context, name string) (map[string]string, *http.Response, error) {
	req := a.APIClient.http_orkes_v4.IntegrationResourceAPI.GetTokenUsageForIntegrationProvider(ctx, name)

	usage, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	return usage, resp, nil
}

// IntegrationResourceApiGetAllIntegrationsOpts - Optional parameters for GetAllIntegrations
type IntegrationResourceApiGetAllIntegrationsOpts struct {
	Category   optional.String
	ActiveOnly optional.Bool
}

// GetAllIntegrations - Get all integrations
func (a *IntegrationResourceApiService) GetAllIntegrations(ctx context.Context, optionals *IntegrationResourceApiGetAllIntegrationsOpts) ([]model.Integration, *http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.GetAllIntegrations(ctx)

	// Apply optional parameters if provided
	if optionals != nil {
		if optionals.Category.IsSet() {
			req = req.Category(optionals.Category.Value())
		}
		if optionals.ActiveOnly.IsSet() {
			req = req.ActiveOnly(optionals.ActiveOnly.Value())
		}
	}

	genIntegrations, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert generated models to domain models using mapper
	result := toDomainIntegrationsFromGenerated(genIntegrations)
	return result, resp, nil
}

// RecordEventStats - Record event statistics
func (a *IntegrationResourceApiService) RecordEventStats(ctx context.Context, body []model.EventLog, eventType string) (*http.Response, error) {
	// Convert domain EventLog to generated EventLog using mapper
	genEventLogs := toGeneratedEventLogs(body)

	req := a.APIClient.http_orkes.IntegrationResourceAPI.RecordEventStats(ctx, eventType).EventLog(genEventLogs)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// GetIntegrationAvailableApis - Get available APIs for integration
func (a *IntegrationResourceApiService) GetIntegrationAvailableApis(ctx context.Context, integrationProvider string) ([]string, *http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.GetIntegrationAvailableApis(ctx, integrationProvider)

	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	return result, resp, nil
}

// GetProvidersAndIntegrations - Get all providers and their integrations
func (a *IntegrationResourceApiService) GetProvidersAndIntegrations(ctx context.Context) ([]string, *http.Response, error) {
	req := a.APIClient.http_orkes.IntegrationResourceAPI.GetProvidersAndIntegrations(ctx)

	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	return result, resp, nil
}

// RegisterTokenUsage - Register token usage for integration
func (a *IntegrationResourceApiService) RegisterTokenUsage(ctx context.Context, integrationName string, modelName string, body interface{}) (*http.Response, error) {
	// Convert body to int32 for token usage
	var tokenUsage int32
	if body != nil {
		if usage, ok := body.(int32); ok {
			tokenUsage = usage
		} else if usage, ok := body.(int); ok {
			tokenUsage = int32(usage)
		}
	}

	req := a.APIClient.http_orkes_v4.IntegrationResourceAPI.RegisterTokenUsage(ctx, integrationName, modelName).Body(tokenUsage)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// SaveAllIntegrations - Save all integrations
func (a *IntegrationResourceApiService) SaveAllIntegrations(ctx context.Context, integrations []integration.IntegrationUpdate) (*http.Response, error) {
	// Convert domain models to generated models
	genIntegrations := make([]orkes.Integration, len(integrations))
	for i, update := range integrations {
		// Convert IntegrationUpdate to Integration for the API
		genIntegrations[i] = orkes.Integration{
			Category:    ToPointer(update.Category),
			Description: ToPointer(update.Description),
			Type:        ToPointer(update.Type_),
			Enabled:     ToPointer(update.Enabled),
		}

		// Convert configuration if present
		if update.Configuration != nil {
			config := make(map[string]interface{})
			for k, v := range update.Configuration {
				config[string(k)] = v

			}
			genIntegrations[i].Configuration = config
		}
	}

	req := a.APIClient.http_orkes.IntegrationResourceAPI.SaveAllIntegrations(ctx).Integration(genIntegrations)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}
