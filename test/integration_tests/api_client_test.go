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
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/authentication"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/require"
)

func TestAPIClientCreationMethods(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

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
	})

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
		)

		apiClient := client.NewAPIClientFromSettings(clientSettings, settings.WithTokenManager(tokenManager))
		require.NotNil(t, apiClient, "API client should not be nil")

		// Test basic functionality - get version
		versionClient := client.NewVersionResourceClient(apiClient)
		version, resp, err := versionClient.GetVersion(context.Background())
		require.NoError(t, err, "Should be able to get version")
		require.Equal(t, 200, resp.StatusCode, "Version request should succeed")
		require.NotEmpty(t, version, "Version should not be empty")
	})
}

func TestAPIClientServerURLNormalization(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	base := testdata.GetServerURL() // usually ends with /api
	root := strings.TrimRight(strings.TrimSuffix(base, "/api"), "/")

	type testCase struct {
		name string
		in   string
	}

	cases := []testCase{
		{name: "root", in: root},
		{name: "root-slash", in: root + "/"},
		{name: "root-double-slash", in: root + "//"},
		{name: "with-api", in: root + "/api"},
		{name: "with-api-trailing", in: root + "/api/"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			clientSettings := settings.NewClientSettings(
				settings.WithServerURL(tc.in),
				settings.WithHTTPTimeout(30*time.Second),
				settings.WithAuthCredentials(testdata.GetAuthKey(), testdata.GetAuthSecret()),
			)
			apiClient := client.NewAPIClientFromSettings(clientSettings)
			require.NotNil(t, apiClient)

			versionClient := client.NewVersionResourceClient(apiClient)
			version, resp, err := versionClient.GetVersion(context.Background())
			require.NoError(t, err)
			require.Equal(t, 200, resp.StatusCode)
			require.NotEmpty(t, version)
		})
	}
}
