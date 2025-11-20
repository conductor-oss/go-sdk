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

func TestSchemaLifecycle(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	schemaClient := testdata.SchemaClient

	// Create a schema
	ctx := context.Background()
	uuid := uuid.New().String()
	schemaName := fmt.Sprintf("TEST_GO_SCHEMA_%s", uuid)
	schema := model.SchemaDefinition{
		Name: schemaName,
		Type: model.SchemaTypeJSON,
		Data: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{
					"type": "string",
				},
				"age": map[string]interface{}{
					"type": "integer",
				},
			},
		},
	}

	resp, err := schemaClient.CreateSchema(ctx, []model.SchemaDefinition{schema}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	t.Cleanup(func() {
		_, _ = schemaClient.DeleteSchema(ctx, schemaName)
	})

	// Retrieve the created schema
	retrievedSchema, resp, err := schemaClient.GetSchema(ctx, schemaName)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, schemaName, retrievedSchema.Name)
	assert.Equal(t, model.SchemaTypeJSON, retrievedSchema.Type)

	// Delete the schema
	resp, err = schemaClient.DeleteSchema(ctx, schemaName)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify the schema is deleted
	_, resp, _ = schemaClient.GetSchema(ctx, schemaName)
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
	schemaV1 := model.SchemaDefinition{
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
	schemaV2 := model.SchemaDefinition{
		Name: schemaName,
		Type: model.SchemaTypeJSON,
		Data: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{
					"type": "string",
				},
				"age": map[string]interface{}{
					"type": "integer",
				},
			},
		},
	}

	opts := &client.SchemaResourceApiCreateSchemaOpts{
		NewVersion: optional.NewBool(true),
	}
	resp, err = schemaClient.CreateSchema(ctx, []model.SchemaDefinition{schemaV2}, opts)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Get the second version
	schema2, resp, err := schemaClient.GetSchemaVersion(ctx, schemaName, 2)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, int32(2), schema2.Version)

	// Get latest version (should be version 2)
	latestSchema, resp, err := schemaClient.GetSchema(ctx, schemaName)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, int32(2), latestSchema.Version)

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

	// Cleanup - delete all versions
	t.Cleanup(func() {
		_, _ = schemaClient.DeleteSchema(ctx, schemaName)
	})
}

func TestGetAllSchemas(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	schemaClient := testdata.SchemaClient

	ctx := context.Background()
	uuid := uuid.New().String()
	schemaName := fmt.Sprintf("TEST_GO_SCHEMA_GETALL_%s", uuid)

	// Create a schema
	schema := model.SchemaDefinition{
		Name: schemaName,
		Type: model.SchemaTypeJSON,
		Data: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"test": map[string]interface{}{
					"type": "string",
				},
			},
		},
	}

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
			break
		}
	}
	assert.True(t, found, "Created schema should be in the list")

	// Get all schemas with short option
	opts := &client.SchemaResourceApiGetAllOpts{
		Short: optional.NewBool(true),
	}
	shortSchemas, resp, err := schemaClient.GetAll(ctx, opts)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, shortSchemas)

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

	schemas := []model.SchemaDefinition{
		{
			Name: schema1Name,
			Type: model.SchemaTypeJSON,
			Data: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"field1": map[string]interface{}{
						"type": "string",
					},
				},
			},
		},
		{
			Name: schema2Name,
			Type: model.SchemaTypeJSON,
			Data: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"field2": map[string]interface{}{
						"type": "integer",
					},
				},
			},
		},
	}

	resp, err := schemaClient.CreateSchema(ctx, schemas, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	t.Cleanup(func() {
		_, _ = schemaClient.DeleteSchema(ctx, schema1Name)
		_, _ = schemaClient.DeleteSchema(ctx, schema2Name)
	})

	// Verify both schemas were created
	schema1, resp, err := schemaClient.GetSchema(ctx, schema1Name)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, schema1Name, schema1.Name)

	schema2, resp, err := schemaClient.GetSchema(ctx, schema2Name)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, schema2Name, schema2.Name)

	// Cleanup

}

func TestSchemaClientErrorHandling(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	schemaClient := testdata.SchemaClient

	ctx := context.Background()
	uuid := uuid.New().String()
	invalidSchemaName := fmt.Sprintf("nonexistent-schema-%s", uuid)

	// Try to get a non-existent schema
	_, resp, err := schemaClient.GetSchema(ctx, invalidSchemaName)
	require.NotNil(t, err)
	require.NotNil(t, resp)
	// API may return 404 or 500 for non-existent schemas
	assert.GreaterOrEqual(t, resp.StatusCode, 400, "Expected error status code (>= 400)")

	// Try to get a non-existent schema version
	_, resp, err = schemaClient.GetSchemaVersion(ctx, invalidSchemaName, 1)
	require.NotNil(t, err)
	require.NotNil(t, resp)
	// API may return 404 or 500 for non-existent schema versions
	assert.GreaterOrEqual(t, resp.StatusCode, 400, "Expected error status code (>= 400)")

	// Try to delete a non-existent schema
	resp, err = schemaClient.DeleteSchema(ctx, invalidSchemaName)
	require.NotNil(t, err)
	require.NotNil(t, resp)
	// API may return 404 or 500 for non-existent schemas
	assert.GreaterOrEqual(t, resp.StatusCode, 400, "Expected error status code (>= 400)")

	// Try to delete a non-existent schema version
	resp, err = schemaClient.DeleteSchemaVersion(ctx, invalidSchemaName, 1)
	require.NotNil(t, err)
	require.NotNil(t, resp)
	// API may return 404 or 500 for non-existent schema versions
	assert.GreaterOrEqual(t, resp.StatusCode, 400, "Expected error status code (>= 400)")
}

func TestSchemaClientIntegration(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	schemaClient := testdata.SchemaClient

	ctx := context.Background()
	uuid := uuid.New().String()
	schemaName := fmt.Sprintf("TEST_GO_SCHEMA_INTEGRATION_%s", uuid)

	// Create a schema
	schema := model.SchemaDefinition{
		Name: schemaName,
		Type: model.SchemaTypeJSON,
		Data: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{
					"type": "string",
				},
				"email": map[string]interface{}{
					"type": "string",
				},
			},
		},
	}

	resp, err := schemaClient.CreateSchema(ctx, []model.SchemaDefinition{schema}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	t.Cleanup(func() {
		_, _ = schemaClient.DeleteSchema(ctx, schemaName)
	})

	// Get the schema
	gotSchema, resp, err := schemaClient.GetSchema(ctx, schemaName)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, schemaName, gotSchema.Name)
	assert.Equal(t, model.SchemaTypeJSON, gotSchema.Type)
	assert.Equal(t, int32(1), gotSchema.Version)

	// Get all schemas and verify our schema is in the list
	allSchemas, resp, err := schemaClient.GetAll(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	found := false
	for _, s := range allSchemas {
		if s.Name == schemaName {
			found = true
			break
		}
	}
	assert.True(t, found, "Created schema should be in the list")

	// Create a new version
	schemaV2 := model.SchemaDefinition{
		Name: schemaName,
		Type: model.SchemaTypeJSON,
		Data: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{
					"type": "string",
				},
				"email": map[string]interface{}{
					"type": "string",
				},
				"age": map[string]interface{}{
					"type": "integer",
				},
			},
		},
	}

	opts := &client.SchemaResourceApiCreateSchemaOpts{
		NewVersion: optional.NewBool(true),
	}
	resp, err = schemaClient.CreateSchema(ctx, []model.SchemaDefinition{schemaV2}, opts)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Get version 2
	gotSchemaV2, resp, err := schemaClient.GetSchemaVersion(ctx, schemaName, 2)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, int32(2), gotSchemaV2.Version)

	// Get latest version (should be version 2)
	latestSchema, resp, err := schemaClient.GetSchema(ctx, schemaName)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, int32(2), latestSchema.Version)

	// Delete version 1
	resp, err = schemaClient.DeleteSchemaVersion(ctx, schemaName, 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify version 1 is deleted
	_, resp, _ = schemaClient.GetSchemaVersion(ctx, schemaName, 1)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)

	// Verify version 2 still exists
	gotSchemaV2After, resp, err := schemaClient.GetSchemaVersion(ctx, schemaName, 2)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, int32(2), gotSchemaV2After.Version)
}

func TestSchemaWithDifferentTypes(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	schemaClient := testdata.SchemaClient

	ctx := context.Background()
	uuid := uuid.New().String()

	// Test JSON schema
	jsonSchemaName := fmt.Sprintf("TEST_GO_SCHEMA_JSON_%s", uuid)
	jsonSchema := model.SchemaDefinition{
		Name: jsonSchemaName,
		Type: model.SchemaTypeJSON,
		Data: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"test": map[string]interface{}{
					"type": "string",
				},
			},
		},
	}

	resp, err := schemaClient.CreateSchema(ctx, []model.SchemaDefinition{jsonSchema}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify JSON schema
	retrievedJSON, resp, err := schemaClient.GetSchema(ctx, jsonSchemaName)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, model.SchemaTypeJSON, retrievedJSON.Type)

	// Cleanup
	t.Cleanup(func() {
		_, _ = schemaClient.DeleteSchema(ctx, jsonSchemaName)
	})
}
