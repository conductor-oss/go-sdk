//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package integration_tests

import (
	"context"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/authentication"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/require"
)

// TestNewClientCreationMethods tests the new client creation methods
func TestNewClientCreationMethods(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	t.Parallel()

	// Test 1: NewAPIClientFromEnv with overrides
	t.Run("NewAPIClientFromEnv", func(t *testing.T) {
		apiClient := client.NewAPIClientFromEnv(
			settings.WithHTTPTimeout(30*time.Second),
			settings.WithHTTPHeaders(map[string]string{
				"User-Agent": "NewClientTest/1.0",
			}),
		)
		require.NotNil(t, apiClient, "API client should not be nil")

		// Test basic functionality - get version
		versionClient := client.NewVersionResourceClient(apiClient)
		version, resp, err := versionClient.GetVersion(context.Background())
		require.NoError(t, err, "Should be able to get version")
		require.Equal(t, 200, resp.StatusCode, "Version request should succeed")
		require.NotEmpty(t, version, "Version should not be empty")

		t.Logf("Successfully created client with NewAPIClientFromEnv, version: %s", version)
	})

	// Test 2: NewAPIClientFromSettings with token management
	t.Run("NewAPIClientWithTokenManagement", func(t *testing.T) {
		authSettings := settings.NewAuthenticationSettings(testdata.GetAuthKey(), testdata.GetAuthSecret())
		tokenExpiration := authentication.NewTokenExpiration(
			45*time.Minute, // default expiration
			3*time.Hour,    // cleanup interval
		)
		tokenManager := authentication.NewTokenManager(*authSettings, tokenExpiration)

		clientSettings := settings.NewClientSettings(
			settings.WithServerURL(testdata.GetServerURL()),
			settings.WithHTTPTimeout(30*time.Second),
			settings.WithAuthCredentials(testdata.GetAuthKey(), testdata.GetAuthSecret()),
			settings.WithTokenExpiration(tokenExpiration),
			settings.WithTokenManager(tokenManager),
		)

		apiClient := client.NewAPIClientFromSettings(clientSettings)
		require.NotNil(t, apiClient, "API client should not be nil")

		// Test basic functionality - get version
		versionClient := client.NewVersionResourceClient(apiClient)
		version, resp, err := versionClient.GetVersion(context.Background())
		require.NoError(t, err, "Should be able to get version")
		require.Equal(t, 200, resp.StatusCode, "Version request should succeed")
		require.NotEmpty(t, version, "Version should not be empty")

		t.Logf("Successfully created client with token management, version: %s", version)
	})
}
