package integration_tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Name/value seeded via CONDUCTOR_SECRET_GO_SDK_INTEGRATION_TEST in
// scripts/docker-compose-oss.yaml -- keep these in sync with it.
const (
	ossSeededSecretName  = "GO_SDK_INTEGRATION_TEST"
	ossSeededSecretValue = "go-sdk-oss-secret-value"
)

// TestSecretReads covers the read half of the secrets API, which plain OSS
// Conductor does serve (list / get / exists), unlike the writes and tags that
// TestSecretResourceApiService below exercises.
//
// This is OSS-only: on Orkes the same three reads are already covered by
// TestSecretResourceApiService, so running here too would only duplicate them
// and write another secret to a shared cluster.
//
// The reads run against the secret seeded into the compose stack, since the
// only bundled SecretsDAO backends (env-var, noop) are read-only; a write is
// asserted to fail with 501 rather than persisting. An OSS image old enough to
// predate the secrets controller entirely 404s on every call, so a failed first
// read skips with a clear message instead of being reported as an SDK bug.
func TestSecretReads(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	if !testdata.IsOSS() {
		t.Skip("skip: read coverage on Orkes is already provided by TestSecretResourceApiService")
	}

	secretClient := testdata.SecretClient
	ctx := context.Background()

	retrieved, resp, err := secretClient.GetSecret(ctx, ossSeededSecretName)
	if err != nil {
		t.Skipf("skip: secrets API unavailable on this OSS server (GetSecret(%q): %v)", ossSeededSecretName, err)
	}
	require.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, ossSeededSecretValue, retrieved)

	existsResult, resp, err := secretClient.SecretExists(ctx, ossSeededSecretName)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, existsResult)

	allSecretNames, resp, err := secretClient.ListAllSecretNames(ctx)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Contains(t, allSecretNames, ossSeededSecretName)

	// Read-only backend: a write must fail rather than silently no-op. Success
	// is accepted too, in case a future OSS release ships a writable backend --
	// that is a signal to revisit ossGapSecretWrites, not a failure.
	throwaway := fmt.Sprintf("test-secret-write-%d", time.Now().UnixNano())
	_, resp, err = secretClient.PutSecret(ctx, "value", throwaway)
	switch {
	case err == nil:
		_, _, _ = secretClient.DeleteSecret(ctx, throwaway)
	case resp != nil:
		assert.Equal(t, 501, resp.StatusCode,
			"expected PutSecret to fail with 501 on OSS's read-only SecretsDAO")
	default:
		t.Errorf("PutSecret failed with no HTTP response: %v", err)
	}
}

func TestSecretResourceApiService(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	testdata.SkipIfOSS(t, ossGapSecretWrites)

	// Setup
	secretClient := testdata.SecretClient // Assuming this exists in your testdata package
	ctx := context.Background()

	// Generate a unique secret key for testing
	secretKey := fmt.Sprintf("test-secret-%d", time.Now().UnixNano())
	secretValue := "this-is-a-test-secret-value"

	// Test case 1: Put a new secret
	_, resp, err := secretClient.PutSecret(ctx, secretValue, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 2: Check if secret exists
	existsResult, resp, err := secretClient.SecretExists(ctx, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, existsResult)

	// Test case 3: Get the secret value
	secretRetrieved, resp, err := secretClient.GetSecret(ctx, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, secretValue, secretRetrieved)

	// Test case 4: Add tags to the secret
	tags := []model.Tag{
		{
			Key:   "environment",
			Value: "test",
			Type_: "metadata",
		},
		{
			Key:   "owner",
			Value: "integration-test",
			Type_: "ownership",
		},
	}

	resp, err = secretClient.PutTagForSecret(ctx, tags, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 5: Get tags for the secret
	retrievedTags, resp, err := secretClient.GetTags(ctx, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify tags were added correctly
	assert.GreaterOrEqual(t, len(retrievedTags), 2)
	var foundEnvironmentTag, foundOwnerTag bool
	for _, tag := range retrievedTags {
		if tag.Key == "environment" && tag.Value == "test" {
			foundEnvironmentTag = true
		}
		if tag.Key == "owner" && tag.Value == "integration-test" {
			foundOwnerTag = true
		}
	}
	assert.True(t, foundEnvironmentTag, "Environment tag not found")
	assert.True(t, foundOwnerTag, "Owner tag not found")

	// Test case 6: List all secret names
	allSecretNames, resp, err := secretClient.ListAllSecretNames(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Check if our secret is in the list
	var foundSecret bool
	for _, name := range allSecretNames {
		if name == secretKey {
			foundSecret = true
			break
		}
	}
	assert.True(t, foundSecret, "Created secret not found in list of all secrets")

	// Test case 7: List secrets that user can grant access to
	accessibleSecrets, resp, err := secretClient.ListSecretsThatUserCanGrantAccessTo(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, accessibleSecrets)
	assert.Equal(t, len(allSecretNames), len(accessibleSecrets))

	// Test case 8: List secrets with tags that user can grant access to
	secretsWithTags, resp, err := secretClient.ListSecretsWithTagsThatUserCanGrantAccessTo(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Check if our secret is in the list with correct tags
	foundSecretWithTags := false
	for _, secret := range secretsWithTags {
		if secret.Name == secretKey {
			foundSecretWithTags = true

			// Verify tags are present
			foundEnvironmentTag = false
			foundOwnerTag = false
			for _, tag := range secret.Tags {
				if tag.Key == "environment" && tag.Value == "test" {
					foundEnvironmentTag = true
				}
				if tag.Key == "owner" && tag.Value == "integration-test" {
					foundOwnerTag = true
				}
			}
			assert.True(t, foundEnvironmentTag, "Environment tag not found in secret with tags")
			assert.True(t, foundOwnerTag, "Owner tag not found in secret with tags")

			break
		}
	}
	assert.True(t, foundSecretWithTags, "Created secret not found in list of secrets with tags")

	// Test case 9: Delete a tag from the secret
	tagToDelete := []model.Tag{
		{
			Key:   "environment",
			Value: "test",
			Type_: "metadata",
		},
	}

	resp, err = secretClient.DeleteTagForSecret(ctx, tagToDelete, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 10: Verify tag was deleted
	updatedTags, resp, err := secretClient.GetTags(ctx, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	environmentTagFound := false
	for _, tag := range updatedTags {
		if tag.Key == "environment" && tag.Value == "test" {
			environmentTagFound = true
			break
		}
	}
	assert.False(t, environmentTagFound, "Environment tag should have been deleted")

	// Test case 11: Test cache clearing operations (these usually don't return meaningful results in tests)
	_, resp, err = secretClient.ClearLocalCache(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	_, resp, err = secretClient.ClearRedisCache(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 12: Delete the secret
	_, resp, err = secretClient.DeleteSecret(ctx, secretKey)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 13: Verify the secret was deleted
	_, resp, err = secretClient.GetSecret(ctx, secretKey)
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode) // Assuming 404 is returned for non-existent secret
}
