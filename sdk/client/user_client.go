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

	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
)

type UserClient interface {
	// CheckPermissions checks the permissions this user has over workflows and tasks.
	CheckPermissions(ctx context.Context, userId string, type_ string, id string) (map[string]interface{}, *http.Response, error)

	// DeleteUser deletes a user.
	DeleteUser(ctx context.Context, id string) (*http.Response, error)

	// GetGrantedPermissions gets the permissions this user has over workflows and tasks.
	GetGrantedPermissions(ctx context.Context, userId string) (rbac.GrantedAccessResponse, *http.Response, error)

	// GetUser retrieves a user by its ID.
	GetUser(ctx context.Context, id string) (*rbac.ConductorUser, *http.Response, error)

	// ListUsers retrieves all users in the system.
	ListUsers(ctx context.Context, optionals *UserResourceApiListUsersOpts) ([]rbac.ConductorUser, *http.Response, error)

	// UpsertUser creates or updates a user.
	UpsertUser(ctx context.Context, body rbac.UpsertUserRequest, id string) (*rbac.ConductorUser, *http.Response, error)
}

// NewUserClient creates a new UserClient.
func NewUserClient(apiClient *APIClient) UserClient {
	return &UserResourceApiService{apiClient}
}
