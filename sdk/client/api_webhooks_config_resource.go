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
	"errors"
	"net/http"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// WebhooksConfigResourceApiService
type WebhooksConfigResourceApiService struct {
	*APIClient
}

// DeleteTagForWebhook implements WebhooksConfigClient.
func (a *WebhooksConfigResourceApiService) DeleteTagForWebhook(ctx context.Context, id string, body []model.Tag) (*http.Response, error) {
	genTags := toGeneratedTags(body)

	req := a.http_orkes.WebhooksConfigResourceAPI.DeleteTagForWebhook(ctx, id).Tag(genTags)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// GetAllWebhook implements WebhooksConfigClient.
func (a *WebhooksConfigResourceApiService) GetAllWebhook(ctx context.Context) ([]model.WebhookConfig, *http.Response, error) {
	req := a.http_orkes.WebhooksConfigResourceAPI.GetAllWebhook(ctx)

	genWebhooks, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainWebhookConfigsFromGenerated(genWebhooks)
	return result, resp, nil
}

// GetTagsForWebhook implements WebhooksConfigClient.
func (a *WebhooksConfigResourceApiService) GetTagsForWebhook(ctx context.Context, id string) ([]model.Tag, *http.Response, error) {
	req := a.http_orkes.WebhooksConfigResourceAPI.GetTagsForWebhook(ctx, id)

	genTags, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainTagsFromGenerated(genTags)
	return result, resp, nil
}

// NewWebhooksConfigResourceApiService creates a new WebhooksConfigResourceApiService instance
func NewWebhooksConfigResourceApiService(client *APIClient) *WebhooksConfigResourceApiService {
	return &WebhooksConfigResourceApiService{client}
}

// CreateWebhook Create webhook
func (a *WebhooksConfigResourceApiService) CreateWebhook(ctx context.Context, body model.WebhookConfig) (model.WebhookConfig, *http.Response, error) {
	genBody := toGeneratedWebhookConfig(body)

	req := a.http_orkes.WebhooksConfigResourceAPI.CreateWebhook(ctx).WebhookConfig(genBody)

	genWebhook, resp, err := req.Execute()
	if err != nil {
		return model.WebhookConfig{}, resp, wrapGeneratedError(err, resp)
	}

	if genWebhook == nil {
		return model.WebhookConfig{}, resp, wrapGeneratedError(errors.New("webhook not found"), resp)
	}

	// Convert using mapper
	result := toDomainWebhookConfigFromGenerated(*genWebhook)
	return result, resp, nil
}

// DeleteWebhook Delete webhook
func (a *WebhooksConfigResourceApiService) DeleteWebhook(ctx context.Context, id string) (*http.Response, error) {
	req := a.http_orkes.WebhooksConfigResourceAPI.DeleteWebhook(ctx, id)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// GetWebhook Get webhook
func (a *WebhooksConfigResourceApiService) GetWebhook(ctx context.Context, id string) (model.WebhookConfig, *http.Response, error) {
	req := a.http_orkes.WebhooksConfigResourceAPI.GetWebhook(ctx, id)

	genWebhook, resp, err := req.Execute()
	if err != nil {
		return model.WebhookConfig{}, resp, wrapGeneratedError(err, resp)
	}

	if genWebhook == nil {
		return model.WebhookConfig{}, resp, wrapGeneratedError(errors.New("webhook not found"), resp)
	}

	// Convert using mapper
	result := toDomainWebhookConfigFromGenerated(*genWebhook)
	return result, resp, nil
}

// ListWebhooks List webhooks
func (a *WebhooksConfigResourceApiService) ListWebhooks(ctx context.Context) ([]model.WebhookConfig, *http.Response, error) {
	req := a.http_orkes.WebhooksConfigResourceAPI.GetAllWebhook(ctx)

	genWebhooks, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainWebhookConfigsFromGenerated(genWebhooks)
	return result, resp, nil
}

// UpdateWebhook Update webhook
func (a *WebhooksConfigResourceApiService) UpdateWebhook(ctx context.Context, body model.WebhookConfig, name string) (model.WebhookConfig, *http.Response, error) {
	// Convert domain model to generated model
	genBody := toGeneratedWebhookConfig(body)

	req := a.http_orkes.WebhooksConfigResourceAPI.UpdateWebhook(ctx, name).WebhookConfig(genBody)

	genWebhook, resp, err := req.Execute()
	if err != nil {
		return model.WebhookConfig{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainWebhookConfigFromGenerated(*genWebhook)
	return result, resp, nil
}

// PutTagForWebhook Add Tag to webhook
func (a *WebhooksConfigResourceApiService) PutTagForWebhook(ctx context.Context, body []model.Tag, id string) (*http.Response, error) {
	genBody := toGeneratedTags(body)

	resp, err := a.http_orkes.WebhooksConfigResourceAPI.PutTagForWebhook(ctx, id).Tag(genBody).Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}
