package unit_tests

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/sdk/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApplicationClientValidationErrors(t *testing.T) {
	apiClient := client.NewAPIClientFromEnv()
	appClient := client.NewApplicationClient(apiClient)
	ctx := context.Background()

	t.Run("AddRoleToApplicationUser - empty applicationId", func(t *testing.T) {
		resp, err := appClient.AddRoleToApplicationUser(ctx, "", "ADMIN")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId")
	})

	t.Run("AddRoleToApplicationUser - empty role", func(t *testing.T) {
		resp, err := appClient.AddRoleToApplicationUser(ctx, "app-id", "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "role")
	})

	t.Run("AddRoleToApplicationUser - both empty", func(t *testing.T) {
		resp, err := appClient.AddRoleToApplicationUser(ctx, "", "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId", "role")
	})

	t.Run("CreateAccessKey - empty applicationId", func(t *testing.T) {
		_, resp, err := appClient.CreateAccessKey(ctx, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationID")
	})

	t.Run("DeleteAccessKey - empty applicationId", func(t *testing.T) {
		resp, err := appClient.DeleteAccessKey(ctx, "", "key-id")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId")
	})

	t.Run("DeleteAccessKey - empty keyId", func(t *testing.T) {
		resp, err := appClient.DeleteAccessKey(ctx, "app-id", "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "accessKeyId")
	})

	t.Run("DeleteAccessKey - both empty", func(t *testing.T) {
		resp, err := appClient.DeleteAccessKey(ctx, "", "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId", "accessKeyId")
	})

	t.Run("DeleteApplication - empty applicationId", func(t *testing.T) {
		_, resp, err := appClient.DeleteApplication(ctx, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId")
	})

	t.Run("DeleteTagForApplication - empty applicationId", func(t *testing.T) {
		tags := []model.Tag{{Key: "env", Value: "test"}}
		resp, err := appClient.DeleteTagForApplication(ctx, tags, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId")
	})

	t.Run("DeleteTagForApplication - empty tags slice", func(t *testing.T) {
		resp, err := appClient.DeleteTagForApplication(ctx, []model.Tag{}, "app-id")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "tags")
	})

	t.Run("DeleteTagForApplication - nil tags", func(t *testing.T) {
		resp, err := appClient.DeleteTagForApplication(ctx, nil, "app-id")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "tags")
	})

	t.Run("DeleteTagForApplication - both invalid", func(t *testing.T) {
		resp, err := appClient.DeleteTagForApplication(ctx, []model.Tag{}, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId", "tags")
	})

	t.Run("GetAccessKeys - empty applicationId", func(t *testing.T) {
		_, resp, err := appClient.GetAccessKeys(ctx, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId")
	})

	t.Run("GetAppByAccessKeyId - empty accessKeyId", func(t *testing.T) {
		_, resp, err := appClient.GetAppByAccessKeyId(ctx, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "accessKeyId")
	})

	t.Run("GetApplication - empty applicationId", func(t *testing.T) {
		_, resp, err := appClient.GetApplication(ctx, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId")
	})

	t.Run("GetTagsForApplication - empty applicationId", func(t *testing.T) {
		_, resp, err := appClient.GetTagsForApplication(ctx, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId")
	})

	t.Run("PutTagForApplication - empty applicationId", func(t *testing.T) {
		tags := []model.Tag{{Key: "env", Value: "test"}}
		resp, err := appClient.PutTagForApplication(ctx, tags, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId")
	})

	t.Run("PutTagForApplication - empty tags slice", func(t *testing.T) {
		resp, err := appClient.PutTagForApplication(ctx, []model.Tag{}, "app-id")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "tags")
	})

	t.Run("PutTagForApplication - nil tags", func(t *testing.T) {
		resp, err := appClient.PutTagForApplication(ctx, nil, "app-id")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "tags")
	})

	t.Run("PutTagForApplication - both invalid", func(t *testing.T) {
		resp, err := appClient.PutTagForApplication(ctx, []model.Tag{}, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId", "tags")
	})

	t.Run("RemoveRoleFromApplicationUser - empty applicationId", func(t *testing.T) {
		resp, err := appClient.RemoveRoleFromApplicationUser(ctx, "", "ADMIN")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId")
	})

	t.Run("RemoveRoleFromApplicationUser - empty role", func(t *testing.T) {
		resp, err := appClient.RemoveRoleFromApplicationUser(ctx, "app-id", "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "role")
	})

	t.Run("RemoveRoleFromApplicationUser - both empty", func(t *testing.T) {
		resp, err := appClient.RemoveRoleFromApplicationUser(ctx, "", "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId", "role")
	})

	t.Run("ToggleAccessKeyStatus - empty applicationId", func(t *testing.T) {
		_, resp, err := appClient.ToggleAccessKeyStatus(ctx, "", "key-id")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId")
	})

	t.Run("ToggleAccessKeyStatus - empty keyId", func(t *testing.T) {
		_, resp, err := appClient.ToggleAccessKeyStatus(ctx, "app-id", "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "accessKeyId")
	})

	t.Run("ToggleAccessKeyStatus - both empty", func(t *testing.T) {
		_, resp, err := appClient.ToggleAccessKeyStatus(ctx, "", "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationId", "accessKeyId")
	})

	t.Run("UpdateApplication - empty applicationId", func(t *testing.T) {
		updateReq := rbac.CreateOrUpdateApplicationRequest{Name: "TestApp"}
		_, resp, err := appClient.UpdateApplication(ctx, updateReq, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "applicationID")
	})
}

// assertValidationError verifies that the error is a validation error and contains the expected field names
func assertValidationError(t *testing.T, err error, expectedFields ...string) {
	t.Helper()
	require.NotNil(t, err, "error should not be nil")

	// Check if it's a MultiValidationError
	multiErr, ok := err.(*validation.MultiValidationError)
	require.True(t, ok, "error should be a MultiValidationError, got: %T", err)
	require.NotEmpty(t, multiErr.Errors, "MultiValidationError should contain errors")

	errorFields := make(map[string]bool)
	for _, validationErr := range multiErr.Errors {
		errorFields[validationErr.FieldPath] = true
		assert.NotEmpty(t, validationErr.Message, "validation error message should not be empty")
	}
	for _, field := range expectedFields {
		assert.True(t, errorFields[field], "expected validation error for field: %s", field)
	}
}
