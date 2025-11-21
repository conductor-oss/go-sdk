package integration_tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createTestSchema creates a test schema definition with the given name.
// Returns a schema with a simple default property.
func createTestSchema(schemaName string) model.SchemaDefinition {
	return model.SchemaDefinition{
		Name: schemaName,
		Type: model.SchemaTypeJSON,
		Data: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{
					"type": "string",
				},
			},
		},
	}
}

func TestSchemaLifecycle(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	schemaClient := testdata.SchemaClient

	// Create a schema
	ctx := context.Background()
	uuid := uuid.New().String()
	schemaName := fmt.Sprintf("TEST_GO_SCHEMA_%s", uuid)
	schema := createTestSchema(schemaName)

	resp, err := schemaClient.CreateSchema(ctx, []model.SchemaDefinition{schema}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	t.Cleanup(func() {
		_, _ = schemaClient.DeleteSchema(ctx, schemaName)
	})

	// Retrieve the created schema (version 1)
	retrievedSchema, resp, err := schemaClient.GetSchemaVersion(ctx, schemaName, 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, schema.Name, retrievedSchema.Name)
	assert.Equal(t, schema.Type, retrievedSchema.Type)
	assert.Equal(t, schema.Data, retrievedSchema.Data)

	// Delete the schema
	resp, err = schemaClient.DeleteSchema(ctx, schemaName)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify the schema is deleted
	_, resp, _ = schemaClient.GetSchemaVersion(ctx, schemaName, 1)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestSchemaVersioning(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	schemaClient := testdata.SchemaClient

	ctx := context.Background()
	uuid := uuid.New().String()
	schemaName := fmt.Sprintf("TEST_GO_SCHEMA_VERSION_%s", uuid)

	// Create initial schema version
	schemaV1 := createTestSchema(schemaName)

	resp, err := schemaClient.CreateSchema(ctx, []model.SchemaDefinition{schemaV1}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Get the first version
	schema1, resp, err := schemaClient.GetSchemaVersion(ctx, schemaName, 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, int32(1), schema1.Version)

	// Create a new version
	schemaV2 := createTestSchema(schemaName)
	// Add age field to show version difference
	propertiesV2 := schemaV2.Data["properties"].(map[string]interface{})
	propertiesV2["age"] = map[string]interface{}{
		"type": "integer",
	}

	opts := &client.SchemaResourceApiCreateSchemaOpts{
		NewVersion: optional.NewBool(true),
	}
	resp, err = schemaClient.CreateSchema(ctx, []model.SchemaDefinition{schemaV2}, opts)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	t.Cleanup(func() {
		_, _ = schemaClient.DeleteSchema(ctx, schemaName)
	})

	// Get version 2
	schema2, resp, err := schemaClient.GetSchemaVersion(ctx, schemaName, 2)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, int32(2), schema2.Version)

	// Delete specific version
	resp, err = schemaClient.DeleteSchemaVersion(ctx, schemaName, 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify version 1 is deleted
	_, resp, _ = schemaClient.GetSchemaVersion(ctx, schemaName, 1)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)

	// Verify version 2 still exists
	schema2After, resp, err := schemaClient.GetSchemaVersion(ctx, schemaName, 2)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, int32(2), schema2After.Version)
}

func TestGetAllSchemas(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	schemaClient := testdata.SchemaClient

	ctx := context.Background()
	uuid := uuid.New().String()
	schemaName := fmt.Sprintf("TEST_GO_SCHEMA_GETALL_%s", uuid)

	// Create a schema
	schema := createTestSchema(schemaName)

	resp, err := schemaClient.CreateSchema(ctx, []model.SchemaDefinition{schema}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Get all schemas
	allSchemas, resp, err := schemaClient.GetAll(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, allSchemas)

	// Verify our schema is in the list
	found := false
	for _, s := range allSchemas {
		if s.Name == schemaName {
			found = true
			assert.Equal(t, model.SchemaTypeJSON, s.Type)
			assert.Equal(t, schema.Data, s.Data)
			break
		}
	}
	assert.True(t, found, "Created schema should be in the list")

	// Cleanup
	t.Cleanup(func() {
		_, _ = schemaClient.DeleteSchema(ctx, schemaName)
	})
}

func TestCreateMultipleSchemas(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	schemaClient := testdata.SchemaClient

	ctx := context.Background()
	uuid := uuid.New().String()

	schema1Name := fmt.Sprintf("TEST_GO_SCHEMA_MULTI_1_%s", uuid)
	schema2Name := fmt.Sprintf("TEST_GO_SCHEMA_MULTI_2_%s", uuid)

	schema1 := createTestSchema(schema1Name)
	schema2 := createTestSchema(schema2Name)

	schemas := []model.SchemaDefinition{schema1, schema2}

	resp, err := schemaClient.CreateSchema(ctx, schemas, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	t.Cleanup(func() {
		_, _ = schemaClient.DeleteSchema(ctx, schema1Name)
		_, _ = schemaClient.DeleteSchema(ctx, schema2Name)
	})

	// Verify both schemas were created (version 1)
	retrievedSchema1, resp, err := schemaClient.GetSchemaVersion(ctx, schema1Name, 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, schema1Name, retrievedSchema1.Name)

	retrievedSchema2, resp, err := schemaClient.GetSchemaVersion(ctx, schema2Name, 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, schema2Name, retrievedSchema2.Name)

}

// Tests using GetSchema method (requires version 5.0+)
func TestSchemaLifecycleWithGetSchema(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV52)

	schemaClient := testdata.SchemaClient

	// Create a schema
	ctx := context.Background()
	uuid := uuid.New().String()
	schemaName := fmt.Sprintf("TEST_GO_SCHEMA_GET_%s", uuid)
	schema := createTestSchema(schemaName)

	resp, err := schemaClient.CreateSchema(ctx, []model.SchemaDefinition{schema}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	t.Cleanup(func() {
		_, _ = schemaClient.DeleteSchema(ctx, schemaName)
	})

	// Retrieve the created schema using GetSchema (latest version)
	retrievedSchema, resp, err := schemaClient.GetSchema(ctx, schemaName)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, schema.Name, retrievedSchema.Name)
	assert.Equal(t, schema.Type, retrievedSchema.Type)
	assert.Equal(t, schema.Data, retrievedSchema.Data)
	assert.Equal(t, int32(1), retrievedSchema.Version)

	// Delete the schema
	resp, err = schemaClient.DeleteSchema(ctx, schemaName)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify the schema is deleted using GetSchema
	_, resp, _ = schemaClient.GetSchema(ctx, schemaName)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}
