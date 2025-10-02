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

// GroupResourceApiService is the service for the group resource.
type GroupResourceApiService struct {
	*APIClient
}

// AddUserToGroup adds user to group.
func (a *GroupResourceApiService) AddUserToGroup(ctx context.Context, groupId string, userId string) (interface{}, *http.Response, error) {
	result, resp, err := a.http_orkes.GroupResourceAPI.AddUserToGroup(ctx, groupId, userId).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// AddUsersToGroup adds users to group.
func (a *GroupResourceApiService) AddUsersToGroup(ctx context.Context, body []string, groupId string) (*http.Response, error) {
	resp, err := a.http_orkes.GroupResourceAPI.AddUsersToGroup(ctx, groupId).RequestBody(body).Execute()
	return resp, wrapGeneratedError(err, resp)
}

// DeleteGroup Delete a group.
func (a *GroupResourceApiService) DeleteGroup(ctx context.Context, id string) (*http.Response, error) {
	_, resp, err := a.http_orkes.GroupResourceAPI.DeleteGroup(ctx, id).Execute()
	return resp, wrapGeneratedError(err, resp)
}

// Deprecated: Use GetGrantedPermissions instead.
func (a *GroupResourceApiService) GetGrantedPermissions1(ctx context.Context, groupId string) (rbac.GrantedAccessResponse, *http.Response, error) {
	return a.GetGrantedPermissions(ctx, groupId)
}

// GetGrantedPermissions gets the permissions this group has over workflows and tasks.
func (a *GroupResourceApiService) GetGrantedPermissions(ctx context.Context, groupId string) (rbac.GrantedAccessResponse, *http.Response, error) {
	genResult, resp, err := a.http_orkes.GroupResourceAPI.GetGrantedPermissions1(ctx, groupId).Execute()
	if err != nil {
		return rbac.GrantedAccessResponse{}, resp, wrapGeneratedError(err, resp)
	}

	result := toDomainGrantedAccessResponseFromGenerated(genResult)
	return result, resp, nil
}

// GetGroup gets a group by id.
func (a *GroupResourceApiService) GetGroup(ctx context.Context, id string) (interface{}, *http.Response, error) {
	result, resp, err := a.http_orkes.GroupResourceAPI.GetGroup(ctx, id).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// GetUsersInGroup gets all users in group.
func (a *GroupResourceApiService) GetUsersInGroup(ctx context.Context, id string) (interface{}, *http.Response, error) {
	res, resp, err := a.http_orkes.GroupResourceAPI.GetUsersInGroup(ctx, id).Execute()

	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	return res, resp, nil
}

// ListGroups gets all groups.
func (a *GroupResourceApiService) ListGroups(ctx context.Context) ([]rbac.Group, *http.Response, error) {
	genResult, resp, err := a.http_orkes.GroupResourceAPI.ListGroups(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	result := toDomainGroupsFromGenerated(genResult)
	return result, resp, nil
}

// RemoveUserFromGroup removes user from group.
func (a *GroupResourceApiService) RemoveUserFromGroup(ctx context.Context, groupId string, userId string) (interface{}, *http.Response, error) {
	result, resp, err := a.http_orkes.GroupResourceAPI.RemoveUserFromGroup(ctx, groupId, userId).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// RemoveUsersFromGroup removes users from group.
func (a *GroupResourceApiService) RemoveUsersFromGroup(ctx context.Context, body []string, groupId string) (*http.Response, error) {
	resp, err := a.http_orkes.GroupResourceAPI.RemoveUsersFromGroup(ctx, groupId).RequestBody(body).Execute()
	return resp, wrapGeneratedError(err, resp)
}

// UpsertGroup Create or update a group.
func (a *GroupResourceApiService) UpsertGroup(ctx context.Context, body rbac.UpsertGroupRequest, id string) (interface{}, *http.Response, error) {
	genRequest := toGeneratedUpsertGroupRequest(body)
	result, resp, err := a.http_orkes.GroupResourceAPI.UpsertGroup(ctx, id).UpsertGroupRequest(genRequest).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}
