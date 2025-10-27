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
	"fmt"
	"net/http"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
)

// ApplicationResourceApiService is the service for the application resource.
type ApplicationResourceApiService struct {
	*APIClient
}

// AddRoleToApplicationUser adds a role to an application user
func (a *ApplicationResourceApiService) AddRoleToApplicationUser(ctx context.Context, applicationId string, role string) (*http.Response, error) {
	path := fmt.Sprintf("/applications/%s/roles/%s", applicationId, role)
	resp, err := a.Post(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// CreateAccessKey creates an access key for an application.
func (a *ApplicationResourceApiService) CreateAccessKey(ctx context.Context, id string) (*rbac.CreateAccessKeyResponse, *http.Response, error) {
	var result rbac.CreateAccessKeyResponse
	path := fmt.Sprintf("/applications/%s/accessKeys", id)
	resp, err := a.Post(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateApplication creates an application.
func (a *ApplicationResourceApiService) CreateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest) (*rbac.ConductorApplication, *http.Response, error) {
	var result rbac.ConductorApplication
	resp, err := a.Post(ctx, "/applications", body, &result)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteAccessKey deletes an access key for an application.
func (a *ApplicationResourceApiService) DeleteAccessKey(ctx context.Context, applicationId string, keyId string) (*http.Response, error) {
	path := fmt.Sprintf("/applications/%s/accessKeys/%s", applicationId, keyId)
	resp, err := a.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteApplication deletes an application.
func (a *ApplicationResourceApiService) DeleteApplication(ctx context.Context, id string) (*rbac.DeleteApplicationResponse, *http.Response, error) {
	var result rbac.DeleteApplicationResponse
	path := fmt.Sprintf("/applications/%s", id)
	resp, err := a.Delete(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DeleteTagForApplication deletes a tag for an application.
func (a *ApplicationResourceApiService) DeleteTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error) {
	path := fmt.Sprintf("/applications/%s/tags", id)
	resp, err := a.DeleteWithBody(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetAccessKeys gets all access keys for an application
func (a *ApplicationResourceApiService) GetAccessKeys(ctx context.Context, id string) ([]rbac.AccessKeyResponse, *http.Response, error) {
	var result []rbac.AccessKeyResponse
	path := fmt.Sprintf("/applications/%s/accessKeys", id)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetAppByAccessKeyId gets an application by access key ID
func (a *ApplicationResourceApiService) GetAppByAccessKeyId(ctx context.Context, accessKeyId string) (*rbac.ConductorApplication, *http.Response, error) {
	var result rbac.ConductorApplication
	path := fmt.Sprintf("/applications/key/%s", accessKeyId)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetApplication gets an application by ID.
func (a *ApplicationResourceApiService) GetApplication(ctx context.Context, id string) (*rbac.ConductorApplication, *http.Response, error) {
	var result rbac.ConductorApplication
	path := fmt.Sprintf("/applications/%s", id)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, err
}

// GetTagsForApplication gets all tags for an application
func (a *ApplicationResourceApiService) GetTagsForApplication(ctx context.Context, id string) ([]model.Tag, *http.Response, error) {
	var result []model.Tag
	path := fmt.Sprintf("/applications/%s/tags", id)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// ListApplications lists all applications
func (a *ApplicationResourceApiService) ListApplications(ctx context.Context) ([]rbac.ConductorApplication, *http.Response, error) {
	var result []rbac.ConductorApplication
	resp, err := a.Get(ctx, "/applications", nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// PutTagForApplication adds tags to an application
func (a *ApplicationResourceApiService) PutTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error) {
	path := fmt.Sprintf("/applications/%s/tags", id)
	resp, err := a.Put(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// RemoveRoleFromApplicationUser removes a role from an application user
func (a *ApplicationResourceApiService) RemoveRoleFromApplicationUser(ctx context.Context, applicationId string, role string) (*http.Response, error) {
	path := fmt.Sprintf("/applications/%s/roles/%s", applicationId, role)
	resp, err := a.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ToggleAccessKeyStatus toggles the status of an access key
func (a *ApplicationResourceApiService) ToggleAccessKeyStatus(ctx context.Context, applicationId string, keyId string) (*rbac.AccessKeyResponse, *http.Response, error) {
	var result rbac.AccessKeyResponse
	path := fmt.Sprintf("/applications/%s/accessKeys/%s/status", applicationId, keyId)
	resp, err := a.Post(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateApplication updates an application
func (a *ApplicationResourceApiService) UpdateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest, id string) (*rbac.ConductorApplication, *http.Response, error) {
	var result rbac.ConductorApplication
	path := fmt.Sprintf("/applications/%s", id)
	resp, err := a.Put(ctx, path, body, &result)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
