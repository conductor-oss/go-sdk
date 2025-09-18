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

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// EnvironmentResourceApiService wraps the generated client to maintain backward compatibility
type EnvironmentResourceApiService struct {
	// Embedded for backward compatibility with helper methods
	*APIClient
}

// NewEnvironmentResourceApiService creates a service from existing APIClient
func NewEnvironmentResourceApiService(apiClient *APIClient) *EnvironmentResourceApiService {
	if apiClient == nil {
		return nil
	}

	return &EnvironmentResourceApiService{
		APIClient: apiClient,
	}
}

// CreateOrUpdateEnvVariable - Create or update an environment variable
func (a *EnvironmentResourceApiService) CreateOrUpdateEnvVariable(ctx context.Context, body string, key string) (*http.Response, error) {
	req := a.http_orkes.EnvironmentResourceAPI.CreateOrUpdateEnvVariable(ctx, key).Body(body)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DeleteEnvVariable - Delete an environment variable
func (a *EnvironmentResourceApiService) DeleteEnvVariable(ctx context.Context, key string) (string, *http.Response, error) {
	req := a.http_orkes.EnvironmentResourceAPI.DeleteEnvVariable(ctx, key)

	result, resp, err := req.Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// DeleteTagForEnvVar - Delete tags for environment variable
func (a *EnvironmentResourceApiService) DeleteTagForEnvVar(ctx context.Context, body []model.Tag, name string) (*http.Response, error) {
	// Convert domain tags to generated tags
	genTags := toGeneratedTags(body)

	req := a.http_orkes.EnvironmentResourceAPI.DeleteTagForEnvVar(ctx, name).Tag(genTags)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// Get - Get the environment value by key
func (a *EnvironmentResourceApiService) Get(ctx context.Context, key string) (string, *http.Response, error) {
	req := a.http_orkes.EnvironmentResourceAPI.Get3(ctx, key)

	result, resp, err := req.Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// GetAll - List all the environment variables
func (a *EnvironmentResourceApiService) GetAll(ctx context.Context) ([]model.EnvironmentVariable, *http.Response, error) {
	req := a.http_orkes.EnvironmentResourceAPI.GetAll(ctx)

	genEnvVars, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainEnvironmentVariables(genEnvVars)
	return result, resp, nil
}

// GetTagsForEnvVar - Get tags by environment variable name
func (a *EnvironmentResourceApiService) GetTagsForEnvVar(ctx context.Context, name string) ([]model.Tag, *http.Response, error) {
	req := a.http_orkes.EnvironmentResourceAPI.GetTagsForEnvVar(ctx, name)

	genTags, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert generated tags to domain tags
	result := toDomainTags(genTags)

	return result, resp, nil
}

// PutTagForEnvVar - Update tags for environment variable
func (a *EnvironmentResourceApiService) PutTagForEnvVar(ctx context.Context, body []model.Tag, name string) (*http.Response, error) {
	// Convert domain tags to generated tags
	genTags := toGeneratedTags(body)

	req := a.http_orkes.EnvironmentResourceAPI.PutTagForEnvVar(ctx, name).Tag(genTags)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}
