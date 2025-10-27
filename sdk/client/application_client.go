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
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
)

type ApplicationClient interface {
	// AddRoleToApplicationUser adds a role to an application user.
	AddRoleToApplicationUser(ctx context.Context, applicationId string, role string) (*http.Response, error)

	// CreateAccessKey creates an access key for an application.
	CreateAccessKey(ctx context.Context, id string) (*rbac.CreateAccessKeyResponse, *http.Response, error)

	// CreateApplication creates a new application.
	CreateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest) (*rbac.ConductorApplication, *http.Response, error)

	// DeleteAccessKey deletes an access key for an application.
	DeleteAccessKey(ctx context.Context, applicationId string, keyId string) (*http.Response, error)

	// DeleteApplication deletes an application.
	DeleteApplication(ctx context.Context, id string) (*rbac.DeleteApplicationResponse, *http.Response, error)

	// DeleteTagForApplication deletes tags from an application.
	DeleteTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error)

	// GetAccessKeys retrieves all access keys for an application.
	GetAccessKeys(ctx context.Context, id string) ([]rbac.AccessKeyResponse, *http.Response, error)

	// GetAppByAccessKeyId retrieves an application by its access key ID.
	GetAppByAccessKeyId(ctx context.Context, accessKeyId string) (*rbac.ConductorApplication, *http.Response, error)

	// GetApplication retrieves an application by its ID.
	GetApplication(ctx context.Context, id string) (*rbac.ConductorApplication, *http.Response, error)

	// GetTagsForApplication retrieves all tags associated with an application.
	GetTagsForApplication(ctx context.Context, id string) ([]model.Tag, *http.Response, error)

	// ListApplications retrieves all applications in the system.
	ListApplications(ctx context.Context) ([]rbac.ConductorApplication, *http.Response, error)

	// PutTagForApplication adds or updates tags for an application.
	PutTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error)

	// RemoveRoleFromApplicationUser removes a role from an application user.
	RemoveRoleFromApplicationUser(ctx context.Context, applicationId string, role string) (*http.Response, error)

	// ToggleAccessKeyStatus toggles the status of an access key.
	ToggleAccessKeyStatus(ctx context.Context, applicationId string, keyId string) (*rbac.AccessKeyResponse, *http.Response, error)

	// UpdateApplication updates an existing application.
	UpdateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest, id string) (*rbac.ConductorApplication, *http.Response, error)
}

// NewApplicationClient creates a new ApplicationClient.
func NewApplicationClient(apiClient *APIClient) ApplicationClient {
	return &ApplicationResourceApiService{apiClient}
}
