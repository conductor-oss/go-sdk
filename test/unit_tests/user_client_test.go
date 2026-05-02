package unit_tests

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"github.com/conductor-sdk/conductor-go/sdk/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserClientValidationErrors(t *testing.T) {
	apiClient := client.NewAPIClientFromEnv()
	userClient := client.NewUserClient(apiClient)
	ctx := context.Background()

	t.Run("CheckPermissions - empty userId", func(t *testing.T) {
		result, resp, err := userClient.CheckPermissions(ctx, "", "test", "test")
		require.Error(t, err)
		assert.Nil(t, resp)
		assert.Nil(t, result)
		assertValidationError(t, err, "userId")
	})

	t.Run("GetGrantedPermissions - empty userId", func(t *testing.T) {
		result, resp, err := userClient.GetGrantedPermissions(ctx, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, rbac.GrantedAccessResponse{}, result)
		assertValidationError(t, err, "userId")
	})

	t.Run("DeleteUser - empty userId", func(t *testing.T) {
		resp, err := userClient.DeleteUser(ctx, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "userId")
	})

	t.Run("GetUser - empty userId", func(t *testing.T) {
		result, resp, err := userClient.GetUser(ctx, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assert.Nil(t, result)
		assertValidationError(t, err, "userId")
	})

	t.Run("UpsertUser - empty userId", func(t *testing.T) {
		body := rbac.UpsertUserRequest{}
		result, resp, err := userClient.UpsertUser(ctx, body, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assert.Nil(t, result)
		assertValidationError(t, err, "userId")
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
