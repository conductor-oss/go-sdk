package integration_tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrUpdateEnvVariable(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	testdata.SkipIfOSS(t, ossGapEnvVarWrites)

	ctx := context.Background()
	envClient := NewEnvironmentClient()

	resp, err := envClient.CreateOrUpdateEnvVariable(ctx, "test value", "testKey")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	resp, err = envClient.CreateOrUpdateEnvVariable(ctx, "", "") // Edge case with empty values
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestDeleteEnvVariable(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	testdata.SkipIfOSS(t, ossGapEnvVarWrites)

	TestCreateOrUpdateEnvVariable(t)
	ctx := context.Background()
	envClient := NewEnvironmentClient()

	message, resp, err := envClient.DeleteEnvVariable(ctx, "testKey")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "test value", message)

}

func TestDeleteTagForEnvVar(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	testdata.SkipIfOSS(t, ossGapEnvVarTags)

	TestCreateOrUpdateEnvVariable(t)
	ctx := context.Background()
	envClient := NewEnvironmentClient()
	tags := []model.Tag{{Key: "tag1", Value: "value1"}}

	resp, err := envClient.DeleteTagForEnvVar(ctx, tags, "envVarName")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func TestGetEnvVariable(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	testdata.SkipIfOSS(t, ossGapEnvVarWrites)

	ctx := context.Background()
	envClient := NewEnvironmentClient()

	value, resp, err := envClient.Get(ctx, "testKey")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, value)

}

func TestGetAllEnvVariables(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	ctx := context.Background()
	envClient := NewEnvironmentClient()

	if testdata.OSSGapSkipped() {
		// OSS serves GET /environment but has no way to seed a variable
		// (create/update is Orkes-Enterprise-only, see ossGapEnvVarWrites),
		// so assert the list call itself works rather than skipping the whole
		// test. An OSS image predating EnvironmentResource 404s here.
		variables, resp, err := envClient.GetAll(ctx)
		if err != nil {
			t.Skipf("skip: environment API unavailable on this OSS server (GetAll: %v)", err)
		}
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotNil(t, variables)
		return
	}

	TestCreateOrUpdateEnvVariable(t)

	variables, resp, err := envClient.GetAll(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Greater(t, len(variables), 0) // Expecting at least one variable
}

func TestGetTagsForEnvVar(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	testdata.SkipIfOSS(t, ossGapEnvVarTags)

	TestUpsertUser(t)
	ctx := context.Background()
	envClient := NewEnvironmentClient()
	tags := []model.Tag{{Key: "tag1", Value: "value1"}}
	envClient.PutTagForEnvVar(ctx, tags, "envVarName")
	tags, resp, err := envClient.GetTagsForEnvVar(ctx, "envVarName")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Greater(t, len(tags), 0)

}

func TestPutTagForEnvVar(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	testdata.SkipIfOSS(t, ossGapEnvVarTags)

	ctx := context.Background()
	envClient := NewEnvironmentClient()
	tags := []model.Tag{{Key: "tag1", Value: "value1"}}

	resp, err := envClient.PutTagForEnvVar(ctx, tags, "envVarName")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
func NewEnvironmentClient() client.EnvironmentClient {
	return testdata.EnvironmentClient
}
