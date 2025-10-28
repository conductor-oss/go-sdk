package integration_tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/require"
)

// TestCheckPermissions checks if permissions for a user can be retrieved correctly.
func TestCheckPermissions(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	TestUpsertUser(t)
	client := NewUserClient()
	ctx := context.Background()
	userId := "testuser"
	type_ := "WORKFLOW_DEF"
	id := "kitchen_sink"

	permissions, resp, err := client.CheckPermissions(ctx, userId, type_, id)
	require.NoError(t, err, "CheckPermissions failed")
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200, got %d", resp.StatusCode)
	_, ok := permissions["CREATE"]
	require.True(t, ok, "Expected 'allowed' field in the response, but found %s", permissions)
}

// TestDeleteUser verifies that a user can be successfully deleted.
func TestDeleteUser(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	TestUpsertUser(t)
	client := NewUserClient()
	ctx := context.Background()
	id := "testuser"

	resp, err := client.DeleteUser(ctx, id)
	require.NoError(t, err, "DeleteUser failed")
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200, got %d", resp.StatusCode)
}

// TestGetGrantedPermissions checks if granted permissions can be fetched for a user.
func TestGetGrantedPermissions(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	TestUpsertUser(t)
	client := NewUserClient()
	ctx := context.Background()
	userId := "testuser"

	permissions, resp, err := client.GetGrantedPermissions(ctx, userId)
	require.NoError(t, err, "GetGrantedPermissions failed")
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200, got %d", resp.StatusCode)
	require.Equal(t, 0, len(permissions.GrantedAccess), "Expected non-empty permissions")
}

// TestGetUser checks fetching a specific user's details.
func TestGetUser(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	TestUpsertUser(t)
	client := NewUserClient()
	ctx := context.Background()
	id := "testuser"

	user, resp, err := client.GetUser(ctx, id)
	require.NoError(t, err, "GetUser failed")
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200, got %d", resp.StatusCode)
	require.Equal(t, id, user.Id, "Expected user ID %v, got %v", id, user.Id)
}

func TestGetUserNotFound(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	TestUpsertUser(t)
	client := NewUserClient()
	ctx := context.Background()
	id := "testuserxxx_doesnot_exist"

	user, resp, _ := client.GetUser(ctx, id)

	require.Equal(t, http.StatusNotFound, resp.StatusCode, "Expected status code 404, got %d", resp.StatusCode)
	require.Nil(t, user)
}

// TestListUsers checks listing users with optional parameters.
func TestListUsers(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	user_client := NewUserClient()
	ctx := context.Background()
	options := client.UserResourceApiListUsersOpts{Apps: optional.NewBool(true)}

	users, resp, err := user_client.ListUsers(ctx, &options)
	require.NoError(t, err, "ListUsers failed")
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200, got %d", resp.StatusCode)
	require.Greater(t, len(users), 0, "Expected non-empty user list")
}

// TestUpsertUser verifies that a user can be updated or inserted.
func TestUpsertUser(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	client := NewUserClient()
	ctx := context.Background()
	body := rbac.UpsertUserRequest{
		Name:  "testuser",
		Roles: []string{"ADMIN", "USER"},
	}
	id := "testUser"

	user, resp, err := client.UpsertUser(ctx, body, id)
	require.NoError(t, err, "UpsertUser failed")
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200, got %d", resp.StatusCode)
	require.Equal(t, body.Name, user.Name, "Expected username %v, got %v", body.Name, user.Name)
}

func NewUserClient() client.UserClient {
	return testdata.UserClient
}
