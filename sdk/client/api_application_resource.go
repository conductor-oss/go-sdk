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
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
)

// ApplicationResourceApiService wraps the generated client to maintain backward compatibility
type ApplicationResourceApiService struct {
	*APIClient
}

// NewApplicationResourceApiService creates a service from existing APIClient
func NewApplicationResourceApiService(apiClient *APIClient) *ApplicationResourceApiService {
	if apiClient == nil {
		return nil
	}

	return &ApplicationResourceApiService{
		APIClient: apiClient,
	}
}

// AddRoleToApplicationUser adds role to application user
func (a *ApplicationResourceApiService) AddRoleToApplicationUser(ctx context.Context, applicationId string, role string) (interface{}, *http.Response, error) {
	req := a.http_orkes.ApplicationResourceAPI.AddRoleToApplicationUser(ctx, applicationId, role)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// CreateAccessKey creates an access key for an application
func (a *ApplicationResourceApiService) CreateAccessKey(ctx context.Context, id string) (*rbac.ConductorApplication, *http.Response, error) {
	req := a.http_orkes.ApplicationResourceAPI.CreateAccessKey(ctx, id)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert result using mapper
	app := toDomainConductorApplication(result)
	return app, resp, nil
}

// CreateApplication creates an application
func (a *ApplicationResourceApiService) CreateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest) (*rbac.ConductorApplication, *http.Response, error) {
	// Convert to generated model
	genBody := orkes.CreateOrUpdateApplicationRequest{
		Name: body.Name,
	}

	req := a.http_orkes.ApplicationResourceAPI.CreateApplication(ctx).CreateOrUpdateApplicationRequest(genBody)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert result from map[string]interface{}
	app := toDomainConductorApplication(result)
	return app, resp, nil
}

// DeleteAccessKey deletes an access key
func (a *ApplicationResourceApiService) DeleteAccessKey(ctx context.Context, applicationId string, keyId string) (*http.Response, error) {
	req := a.http_orkes.ApplicationResourceAPI.DeleteAccessKey(ctx, applicationId, keyId)
	_, resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DeleteApplication deletes an application
func (a *ApplicationResourceApiService) DeleteApplication(ctx context.Context, id string) (interface{}, *http.Response, error) {
	req := a.http_orkes.ApplicationResourceAPI.DeleteApplication(ctx, id)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// DeleteTagForApplication removes tags from the application
func (a *ApplicationResourceApiService) DeleteTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error) {
	// Convert tags
	genTags := toGeneratedTags(body)

	req := a.http_orkes.ApplicationResourceAPI.DeleteTagForApplication(ctx, id).Tag(genTags)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// GetAccessKeys gets access keys for an application
func (a *ApplicationResourceApiService) GetAccessKeys(ctx context.Context, id string) ([]rbac.AccessKeyResponse, *http.Response, error) {
	// // Prefer direct HTTP because server returns an array and generator expects a map
	// if keys, resp, err := a.getAccessKeysDirect(ctx, id); err == nil {
	// 	return keys, resp, nil
	// }

	// Fallback to generated client in case server behavior changes
	req := a.http_orkes.ApplicationResourceAPI.GetAccessKeys(ctx, id)
	genKeys, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	// Convert the generic map into domain slice
	mapped := toDomainAccessKeysResponseFromGenerated(genKeys)
	if len(mapped) > 0 {
		return mapped, resp, nil
	}
	return nil, resp, nil
}

// GetApplication gets an application by ID
func (a *ApplicationResourceApiService) GetApplication(ctx context.Context, id string) (*rbac.ConductorApplication, *http.Response, error) {
	req := a.http_orkes.ApplicationResourceAPI.GetApplication(ctx, id)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	app := toDomainConductorApplication(result)
	return app, resp, nil
}

// GetApplicationByAccessKeyId gets an application by access key ID
func (a *ApplicationResourceApiService) GetApplicationByAccessKeyId(ctx context.Context, accessKeyId string) (interface{}, *http.Response, error) {
	req := a.http_orkes.ApplicationResourceAPI.GetAppByAccessKeyId(ctx, accessKeyId)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// GetAppByAccessKeyId gets an application by access key ID
func (a *ApplicationResourceApiService) GetAppByAccessKeyId(ctx context.Context, accessKeyId string) (interface{}, *http.Response, error) {
	return a.GetApplicationByAccessKeyId(ctx, accessKeyId)
}

// GetTagsForApplication gets tags for an application
func (a *ApplicationResourceApiService) GetTagsForApplication(ctx context.Context, id string) ([]model.Tag, *http.Response, error) {
	req := a.http_orkes.ApplicationResourceAPI.GetTagsForApplication(ctx, id)
	genTags, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert tags back to model
	tags := toDomainTags(genTags)
	return tags, resp, nil
}

// ListApplications lists all applications
func (a *ApplicationResourceApiService) ListApplications(ctx context.Context) ([]rbac.ConductorApplication, *http.Response, error) {
	req := a.http_orkes.ApplicationResourceAPI.ListApplications(ctx)
	genApps, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert to []rbac.ConductorApplication
	apps := make([]rbac.ConductorApplication, len(genApps))
	for i, genApp := range genApps {
		apps[i] = toConductorApplicationFromExtendedConductorApplication(&genApp)
	}

	return apps, resp, nil
}

// PutTagForApplication adds tags to the application
func (a *ApplicationResourceApiService) PutTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error) {
	// Convert tags
	genTags := toGeneratedTags(body)

	req := a.http_orkes.ApplicationResourceAPI.PutTagForApplication(ctx, id).Tag(genTags)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// RemoveRoleFromApplicationUser removes a role from the application user
func (a *ApplicationResourceApiService) RemoveRoleFromApplicationUser(ctx context.Context, applicationId string, role string) (interface{}, *http.Response, error) {
	req := a.http_orkes.ApplicationResourceAPI.RemoveRoleFromApplicationUser(ctx, applicationId, role)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// ToggleAccessKeyStatus toggles the status of an access key
func (a *ApplicationResourceApiService) ToggleAccessKeyStatus(ctx context.Context, applicationId string, keyId string) (interface{}, *http.Response, error) {
	req := a.http_orkes.ApplicationResourceAPI.ToggleAccessKeyStatus(ctx, applicationId, keyId)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// UpdateApplication updates an application
func (a *ApplicationResourceApiService) UpdateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest, id string) (*rbac.ConductorApplication, *http.Response, error) {
	// Convert to generated model
	genBody := orkes.CreateOrUpdateApplicationRequest{
		Name: body.Name,
	}

	req := a.http_orkes.ApplicationResourceAPI.UpdateApplication(ctx, id).CreateOrUpdateApplicationRequest(genBody)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	app := toDomainConductorApplication(result)
	return app, resp, nil
}
