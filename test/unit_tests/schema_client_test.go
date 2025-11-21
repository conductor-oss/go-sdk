package unit_tests

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchemaClientValidationErrors(t *testing.T) {
	apiClient := client.NewAPIClientFromEnv()
	schemaClient := client.NewSchemaClient(apiClient)
	ctx := context.Background()

	t.Run("CreateSchema - empty schemas slice", func(t *testing.T) {
		resp, err := schemaClient.CreateSchema(ctx, []model.SchemaDefinition{}, nil)
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "schemas")
	})

	t.Run("CreateSchema - nil schemas", func(t *testing.T) {
		resp, err := schemaClient.CreateSchema(ctx, nil, nil)
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "schemas")
	})

	t.Run("DeleteSchema - empty name", func(t *testing.T) {
		resp, err := schemaClient.DeleteSchema(ctx, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "name")
	})

	t.Run("GetSchema - empty name", func(t *testing.T) {
		_, resp, err := schemaClient.GetSchema(ctx, "")
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "name")
	})

	t.Run("DeleteSchemaVersion - empty name", func(t *testing.T) {
		resp, err := schemaClient.DeleteSchemaVersion(ctx, "", 1)
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "name")
	})

	t.Run("GetSchemaVersion - empty name", func(t *testing.T) {
		_, resp, err := schemaClient.GetSchemaVersion(ctx, "", 1)
		require.Error(t, err)
		assert.Nil(t, resp)
		assertValidationError(t, err, "name")
	})
}
