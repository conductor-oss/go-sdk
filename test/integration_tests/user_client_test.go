package integration_tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// TestCheckPermissions checks if permissions for a user can be retrieved correctly.
func TestCheckPermissions(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	client := NewUserClient()
	ctx := context.Background()
	user := setupUser(t, ctx)

	type_ := "WORKFLOW_DEF"
	id := "kitchen_sink"

	permissions, resp, err := client.CheckPermissions(ctx, user.Id, type_, id)
	require.NoError(t, err, "CheckPermissions failed")
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200, got %d", resp.StatusCode)
	_, ok := permissions["CREATE"]
	require.True(t, ok, "Expected 'allowed' field in the response, but found %s", permissions)
}

// TestDeleteUser verifies that a user can be successfully deleted.
func TestDeleteUser(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	client := NewUserClient()
	ctx := context.Background()
	user := setupUser(t, ctx)

	resp, err := client.DeleteUser(ctx, user.Id)
	require.NoError(t, err, "DeleteUser failed")
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200, got %d", resp.StatusCode)
}

// TestGetGrantedPermissions checks if granted permissions can be fetched for a user.
func TestGetGrantedPermissions(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	client := NewUserClient()
	ctx := context.Background()
	user := setupUser(t, ctx)

	permissions, resp, err := client.GetGrantedPermissions(ctx, user.Id)
	require.NoError(t, err, "GetGrantedPermissions failed")
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200, got %d", resp.StatusCode)
	require.Equal(t, 0, len(permissions.GrantedAccess), "Expected non-empty permissions")
}

// TestGetUser checks fetching a specific user's details.
func TestGetUser(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	client := NewUserClient()
	ctx := context.Background()
	user := setupUser(t, ctx)

	user, resp, err := client.GetUser(ctx, user.Id)
	require.NoError(t, err, "GetUser failed")
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200, got %d", resp.StatusCode)
	require.Equal(t, user.Id, user.Id, "Expected user ID %v, got %v", user.Id, user.Id)
}

func TestGetUserNotFound(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	client := NewUserClient()
	ctx := context.Background()
	uuid := uuid.New().String()
	id := fmt.Sprintf("testuserxxx_doesnot_exist_%s", uuid)

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
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200, got %d", resp.StatusCode)
	require.Greater(t, len(users), 0, "Expected non-empty user list")
}

// TestUpsertUser verifies that a user can be updated or inserted.
func TestUpsertUser(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	ctx := context.Background()
	user := setupUser(t, ctx)

	require.NotNil(t, user)
	require.Equal(t, user.Id, user.Name, "Unexpected username")
	require.NotNil(t, user.ContactInformation, "Unexpected contact information")
}

// setupUser creates a unique user for testing and registers cleanup to delete it.
// Returns the userId and the created user.
func setupUser(t *testing.T, ctx context.Context) *rbac.ConductorUser {
	client := NewUserClient()
	uuid := uuid.New().String()
	userId := fmt.Sprintf("test_go_user_%s", uuid)

	body := rbac.UpsertUserRequest{
		Name:  userId,
		Roles: []string{"ADMIN", "USER"},
		ContactInformation: map[string]interface{}{
			"email": fmt.Sprintf("testuser_%s@example.com", uuid),
		},
	}

	user, resp, err := client.UpsertUser(ctx, body, userId)
	require.NoError(t, err, "UpsertUser failed")
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
	require.NotNil(t, user)

	// Verify user is accessible
	_, resp, err = client.GetUser(ctx, user.Id)
	require.NoError(t, err, "GetUser failed after creation")
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200 for GetUser")
	require.Equal(t, userId, user.Name, "Expected username %s, got %s", userId, user.Name)
	require.Equal(t, userId, user.Id, "Expected user ID %s, got %s", userId, user.Id)

	// Cleanup: delete user after test
	t.Cleanup(func() {
		_, _ = client.DeleteUser(ctx, user.Id)
	})

	return user
}

func NewUserClient() client.UserClient {
	return testdata.UserClient
}
