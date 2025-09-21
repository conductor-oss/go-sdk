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

	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
)

// AuthorizationResourceApiService
type AuthorizationResourceApiService struct {
	*APIClient
}

// NewAuthorizationResourceApiService creates a new AuthorizationResourceApiService instance
func NewAuthorizationResourceApiService(client *APIClient) *AuthorizationResourceApiService {
	return &AuthorizationResourceApiService{APIClient: client}
}

// GetPermissions Get the access that have been granted over the given object
func (a *AuthorizationResourceApiService) GetPermissions(ctx context.Context, type_ string, id string) (interface{}, *http.Response, error) {
	result, resp, err := a.http_orkes.AuthorizationResourceAPI.GetPermissions(ctx, type_, id).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// GrantPermissions Grant access to a user over the target
func (a *AuthorizationResourceApiService) GrantPermissions(ctx context.Context, body rbac.AuthorizationRequest) (*http.Response, error) {
	genRequest := toGeneratedAuthorizationRequest(body)
	_, resp, err := a.http_orkes.AuthorizationResourceAPI.GrantPermissions(ctx).AuthorizationRequest(genRequest).Execute()
	return resp, wrapGeneratedError(err, resp)
}

// RemovePermissions Remove user's access over the target
func (a *AuthorizationResourceApiService) RemovePermissions(ctx context.Context, body rbac.AuthorizationRequest) (*http.Response, error) {
	genRequest := toGeneratedAuthorizationRequest(body)
	_, resp, err := a.http_orkes.AuthorizationResourceAPI.RemovePermissions(ctx).AuthorizationRequest(genRequest).Execute()
	return resp, wrapGeneratedError(err, resp)
}
