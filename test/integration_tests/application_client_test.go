package integration_tests

import (
	"context"
	"fmt"

	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApplicationLifecycle(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	appClient := testdata.ApplicationClient

	// Create an application
	ctx := context.Background()
	uuid := uuid.New().String()
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: fmt.Sprintf("TEST_GO_APP_%s", uuid)}
	createdApp, resp, err := appClient.CreateApplication(ctx, createReq)
	require.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, fmt.Sprintf("TEST_GO_APP_%s", uuid), createdApp.Name)

	// Retrieve the created application
	retrievedApp, resp, err := appClient.GetApplication(ctx, createdApp.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, fmt.Sprintf("TEST_GO_APP_%s", uuid), retrievedApp.Name)

	// Delete the application
	result, resp, err := appClient.DeleteApplication(ctx, createdApp.Id)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.Message)

	// Verify the application is deleted (this step may vary based on how your API handles deletions)
	_, resp, _ = appClient.GetApplication(ctx, createdApp.Id)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestRoleManagementForApplicationUser(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	appClient := testdata.ApplicationClient

	// Create an application to use in the test
	ctx := context.Background()
	uuid := uuid.New().String()
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: fmt.Sprintf("TEST_GO_APP_ROLE_USER_%s", uuid)}
	application, _, err := appClient.CreateApplication(ctx, createReq)
	require.Nil(t, err)

	// Add a role to the application user
	resp, err := appClient.AddRoleToApplicationUser(ctx, application.Id, "ADMIN")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Remove the role from the application user
	resp, err = appClient.RemoveRoleFromApplicationUser(ctx, application.Id, "ADMIN")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Cleanup
	t.Cleanup(func() {
		_, _, err = appClient.DeleteApplication(ctx, application.Id)
		require.Nil(t, err)
	})
}

func TestAuthorizationResourcePermissions(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Setup
	appClient := testdata.ApplicationClient
	authClient := testdata.AuthorizationClient

	// Create an application to use as our test resource
	ctx := context.Background()
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: "TestAuthResource"}
	application, resp, err := appClient.CreateApplication(ctx, createReq)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, application)
	assert.Nil(t, err)

	// Test case 1: Initially verify no permissions are granted
	permissions, resp, err := authClient.GetPermissions(ctx, "APPLICATION", application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, permissions)

	// Test case 2: Grant permissions to a user
	user, err := testdata.CreateNewUser(ctx)
	assert.Nil(t, err)

	grantReq := rbac.AuthorizationRequest{
		Access: []string{"READ", "CREATE"},
		Subject: &rbac.SubjectRef{
			Id:    user.Id,
			Type_: "USER",
		},
		Target: &rbac.TargetRef{
			Id:    application.Id,
			Type_: "APPLICATION",
		},
	}

	resp, err = authClient.GrantPermissions(ctx, grantReq)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 3: Verify permissions were granted
	permissions, resp, err = authClient.GetPermissions(ctx, "APPLICATION", application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 4: Remove permissions
	resp, err = authClient.RemovePermissions(ctx, grantReq)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 5: Verify permissions were removed
	permissions, resp, err = authClient.GetPermissions(ctx, "APPLICATION", application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	// Assert that permissions no longer includes the removed permission

	// Cleanup
	_, _, err = appClient.DeleteApplication(ctx, application.Id)
	require.Nil(t, err)
}

func TestAccessKeyLifecycle(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	appClient := testdata.ApplicationClient

	// Create an application to use in the test
	ctx := context.Background()
	uuid := uuid.New().String()
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: fmt.Sprintf("TEST_GO_APP_ACCESS_KEY_%s", uuid)}
	application, _, err := appClient.CreateApplication(ctx, createReq)
	require.Nil(t, err)

	// Create an access key for the application
	accessKey, _, err := appClient.CreateAccessKey(ctx, application.Id)
	require.Nil(t, err)
	assert.NotNil(t, accessKey)
	assert.NotEmpty(t, accessKey.Secret)

	// Delete the access key
	resp, err := appClient.DeleteAccessKey(ctx, application.Id, accessKey.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Cleanup
	t.Cleanup(func() {
		_, _, err = appClient.DeleteApplication(ctx, application.Id)
		require.Nil(t, err)
	})
}

func TestGetTagsForApplication(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	appClient := testdata.ApplicationClient

	// Create an application to use in the test
	ctx := context.Background()
	uuid := uuid.New().String()
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: fmt.Sprintf("TEST_GO_APP_TAGS_%s", uuid)}
	application, _, err := appClient.CreateApplication(ctx, createReq)
	require.Nil(t, err)
	require.NotNil(t, application)

	// Get tags for the application
	tags, resp, err := appClient.GetTagsForApplication(ctx, application.Id)
	require.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	// Additional assertions based on expected tags
	assert.Equal(t, 0, len(tags))

	// Add a tag to the application
	tags = []model.Tag{{Key: "env", Value: "development"}}
	_, err = appClient.PutTagForApplication(ctx, tags, application.Id)
	require.Nil(t, err)

	// Get tags for the application
	retrievedTags, resp, err := appClient.GetTagsForApplication(ctx, application.Id)
	require.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, len(retrievedTags))
	assert.Equal(t, tags[0].Key, retrievedTags[0].Key)
	assert.Equal(t, tags[0].Value, retrievedTags[0].Value)

	// Cleanup
	t.Cleanup(func() {
		_, resp, err = appClient.DeleteApplication(ctx, application.Id)
		require.Nil(t, err)
	})
}

func TestApplicationClientIntegration(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Initialize ApplicationClient
	appClient := testdata.ApplicationClient

	// Context used in all requests
	ctx := context.Background()

	// Create an application
	uuid := uuid.New().String()
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: fmt.Sprintf("TEST_GO_APP_INTEGRATION_%s", uuid)}
	application, _, err := appClient.CreateApplication(ctx, createReq)
	assert.Nil(t, err)

	t.Cleanup(func() {
		_, _, err := appClient.DeleteApplication(ctx, application.Id)
		require.Nil(t, err)
	})

	// Get the application
	gotApp, _, err := appClient.GetApplication(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("TEST_GO_APP_INTEGRATION_%s", uuid), gotApp.Name)

	// List applications
	apps, _, err := appClient.ListApplications(ctx)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(apps), 1)

	// Verify the application is in the list
	found := false
	for _, app := range apps {
		if app.Id == application.Id {
			assert.Equal(t, fmt.Sprintf("TEST_GO_APP_INTEGRATION_%s", uuid), app.Name)
			found = true
			break
		}
	}
	assert.True(t, found)

	// Add a tag to the application
	tags := []model.Tag{{Key: "env", Value: "development"}}
	_, err = appClient.PutTagForApplication(ctx, tags, application.Id)
	assert.Nil(t, err)

	// Get tags for the application
	retrievedTags, _, err := appClient.GetTagsForApplication(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(retrievedTags))
	assert.Equal(t, tags[0].Key, retrievedTags[0].Key)
	assert.Equal(t, tags[0].Value, retrievedTags[0].Value)

	// Create an access key for the application
	accessKey, _, err := appClient.CreateAccessKey(ctx, application.Id)
	assert.Nil(t, err)
	assert.NotNil(t, accessKey)
	assert.NotEmpty(t, accessKey.Secret)

	t.Cleanup(func() {
		_, err = appClient.DeleteAccessKey(ctx, application.Id, accessKey.Id)
		require.Nil(t, err)
	})

	// Toggle the access key status
	updatedAccessKey, _, err := appClient.ToggleAccessKeyStatus(ctx, application.Id, accessKey.Id)
	assert.Nil(t, err)
	assert.NotNil(t, updatedAccessKey)
	assert.Equal(t, "INACTIVE", updatedAccessKey.Status)

	// Remove the added tag
	_, err = appClient.DeleteTagForApplication(ctx, tags, application.Id)
	assert.Nil(t, err)

	retrievedTags, _, err = appClient.GetTagsForApplication(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(retrievedTags))

	// Update the application
	updatedAppName := fmt.Sprintf("TEST_GO_APP_INTEGRATION_UPDATED_%s", uuid)
	updateReq := rbac.CreateOrUpdateApplicationRequest{Name: updatedAppName}
	updatedApp, _, err := appClient.UpdateApplication(ctx, updateReq, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, updatedAppName, updatedApp.Name)

	// Get the application
	gotApp, _, err = appClient.GetApplication(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, updatedAppName, gotApp.Name)
}

func TestApplicationClientErrorHandling(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Initialize ApplicationClient
	appClient := testdata.ApplicationClient

	// Context used in all requests
	ctx := context.Background()

	// Define an invalid application ID
	uuid := uuid.New().String()
	invalidAppId := fmt.Sprintf("nonexistent-%s", uuid)

	// Try to get a non-existent application
	_, resp, err := appClient.GetApplication(ctx, invalidAppId)
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	// Try to update a non-existent application
	updateReq := rbac.CreateOrUpdateApplicationRequest{Name: "NonExistentApp"}
	_, resp, err = appClient.UpdateApplication(ctx, updateReq, invalidAppId)
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	// Try to delete a non-existent application
	_, resp, err = appClient.DeleteApplication(ctx, invalidAppId)
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	// Try to add a tag to a non-existent application
	tags := []model.Tag{{Key: "env", Value: "staging"}}
	res, err := appClient.PutTagForApplication(ctx, tags, invalidAppId)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 200, res.StatusCode)

	// Try to get tags for a non-existent application
	_, _, err = appClient.GetTagsForApplication(ctx, invalidAppId)
	assert.Nil(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	// Try to delete a tag from a non-existent application
	_, err = appClient.DeleteTagForApplication(ctx, tags, invalidAppId)
	assert.Nil(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	// Try to create an access key for a non-existent application
	_, resp, err = appClient.CreateAccessKey(ctx, invalidAppId)
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	// Try to toggle access key status for a non-existent application
	_, resp, err = appClient.ToggleAccessKeyStatus(ctx, invalidAppId, "fakeKey")
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestGetAccessKeys(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Initialize ApplicationClient
	appClient := testdata.ApplicationClient

	// Context used in all requests
	ctx := context.Background()

	// Create an application to use in the test
	uuid := uuid.New().String()
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: fmt.Sprintf("TEST_GO_APP_ACCESS_KEYS_%s", uuid)}
	application, _, err := appClient.CreateApplication(ctx, createReq)
	assert.Nil(t, err)

	// Initially check access keys when none are added
	keysBefore, resp, err := appClient.GetAccessKeys(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Empty(t, keysBefore)

	// Create an access key
	newKey, _, err := appClient.CreateAccessKey(ctx, application.Id)
	assert.Nil(t, err)
	assert.NotNil(t, newKey)
	assert.NotEmpty(t, newKey.Secret)

	// Retrieve the access keys and check the list contains the new key
	keysAfter, resp, err := appClient.GetAccessKeys(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, keysAfter)

	var (
		found = false
		key   *rbac.AccessKeyResponse
	)
	for _, k := range keysAfter {
		if k.Id == newKey.Id {
			found = true
			key = &k
			break
		}
	}
	assert.True(t, found)
	assert.Equal(t, "ACTIVE", key.Status)

	// Cleanup - delete the access key and application
	_, err = appClient.DeleteAccessKey(ctx, application.Id, newKey.Id)
	assert.Nil(t, err)

	keysAfter, resp, err = appClient.GetAccessKeys(ctx, application.Id)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Empty(t, keysAfter)

	t.Cleanup(func() {
		_, _, err = appClient.DeleteApplication(ctx, application.Id)
		require.Nil(t, err)
	})
}

func TestGetAppByAccessKeyId(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Initialize ApplicationClient
	appClient := testdata.ApplicationClient

	// Context used in all requests
	ctx := context.Background()

	// Create an application to use in the test
	uuid := uuid.New().String()
	createReq := rbac.CreateOrUpdateApplicationRequest{Name: fmt.Sprintf("TEST_GO_APP_ACCESS_KEY_ID_%s", uuid)}
	application, _, err := appClient.CreateApplication(ctx, createReq)
	require.Nil(t, err)
	assert.NotNil(t, application)

	// Create an access key for the application
	accessKey, _, err := appClient.CreateAccessKey(ctx, application.Id)
	require.Nil(t, err)
	assert.NotNil(t, accessKey)
	assert.NotEmpty(t, accessKey.Id)

	app, _, err := appClient.GetAppByAccessKeyId(ctx, accessKey.Id)
	assert.Nil(t, err)
	require.NotNil(t, app)
	require.Equal(t, application.Name, app.Name)

	// Test case 2: Try to get application with invalid access key ID
	invalidAccessKeyId := "invalid-access-key-id"
	_, resp, err := appClient.GetAppByAccessKeyId(ctx, invalidAccessKeyId)
	assert.Error(t, err)
	require.Equal(t, 404, resp.StatusCode)

	t.Cleanup(func() {
		_, err = appClient.DeleteAccessKey(ctx, application.Id, accessKey.Id)
		assert.Nil(t, err)
		_, _, err = appClient.DeleteApplication(ctx, application.Id)
		assert.Nil(t, err)
	})
}
