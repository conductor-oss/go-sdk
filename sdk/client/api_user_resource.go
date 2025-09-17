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
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
)

// UserResourceApiService wraps the generated client to maintain backward compatibility
type UserResourceApiService struct {
	// Embedded for backward compatibility with helper methods
	*APIClient
}

// NewUserResourceApiService creates service from existing APIClient
func NewUserResourceApiService(apiClient *APIClient) *UserResourceApiService {
	if apiClient == nil {
		return nil
	}

	return &UserResourceApiService{
		APIClient: apiClient,
	}
}

// UserResourceApiListUsersOpts contains optional parameters for ListUsers
type UserResourceApiListUsersOpts struct {
	Apps optional.Bool
}

// CheckPermissions - Get the permissions this user has over workflows and tasks
func (a *UserResourceApiService) CheckPermissions(ctx context.Context, userId string, type_ string, id string) (map[string]interface{}, *http.Response, error) {
	req := a.APIClient.http_orkes.UserResourceAPI.CheckPermissions(ctx, userId).
		Type_(type_).
		Id(id)

	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// DeleteUser - Delete a user
func (a *UserResourceApiService) DeleteUser(ctx context.Context, id string) (*http.Response, error) {
	req := a.APIClient.http_orkes.UserResourceAPI.DeleteUser(ctx, id)

	_, resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// GetGrantedPermissions - Get the list of access which was granted over workflows and tasks
func (a *UserResourceApiService) GetGrantedPermissions(ctx context.Context, userId string) (rbac.GrantedAccessResponse, *http.Response, error) {
	req := a.APIClient.http_orkes.UserResourceAPI.GetGrantedPermissions(ctx, userId)

	genResult, resp, err := req.Execute()
	if err != nil {
		return rbac.GrantedAccessResponse{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert map[string]interface{} to rbac.GrantedAccessResponse using comprehensive mapper
	result := toDomainGrantedAccessResponseFromMap(genResult)
	return result, resp, nil
}

// GetUser - Get the user by id
func (a *UserResourceApiService) GetUser(ctx context.Context, id string) (*rbac.ConductorUser, *http.Response, error) {
	req := a.APIClient.http_orkes.UserResourceAPI.GetUser(ctx, id)

	genUser, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainConductorUser(genUser)
	return result, resp, nil
}

// ListUsers - Get all users
func (a *UserResourceApiService) ListUsers(ctx context.Context, optionals *UserResourceApiListUsersOpts) ([]rbac.ConductorUser, *http.Response, error) {
	req := a.APIClient.http_orkes.UserResourceAPI.ListUsers(ctx)

	if optionals != nil && optionals.Apps.IsSet() {
		req = req.Apps(optionals.Apps.Value())
	}

	genUsers, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainConductorUsers(genUsers)
	return result, resp, nil
}

// UpsertUser - Create or update a user
func (a *UserResourceApiService) UpsertUser(ctx context.Context, body rbac.UpsertUserRequest, id string) (*rbac.ConductorUser, *http.Response, error) {
	// Convert domain model to generated model
	genBody := toGeneratedUpsertUserRequest(body)

	req := a.APIClient.http_orkes.UserResourceAPI.UpsertUser(ctx, id).UpsertUserRequest(genBody)

	genUser, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainConductorUser(genUser)
	return result, resp, nil
}
