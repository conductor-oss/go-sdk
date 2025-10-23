package unit_tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/authentication"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_Proxy(t *testing.T) {
	originalConductorProxy := os.Getenv(settings.EnvProxy)
	originalServerURL := os.Getenv(settings.EnvServerURL)
	defer func() {
		// Restore original env values
		if originalConductorProxy != "" {
			os.Setenv(settings.EnvProxy, originalConductorProxy)
		} else {
			os.Unsetenv(settings.EnvProxy)
		}
		if originalServerURL != "" {
			os.Setenv("CONDUCTOR_SERVER_URL", originalServerURL)
		} else {
			os.Unsetenv("CONDUCTOR_SERVER_URL")
		}
	}()

	versionDirect := "3.0.0-test"
	versionViaProxy := "3.0.0-test-via-proxy"

	tests := []struct {
		name             string
		customProxy      bool
		expectViaProxy   bool // Whether request is expected to go via proxy
		expectDirectCall bool // Whether direct connection is expected
	}{
		{
			name:             "no proxy set - direct connection",
			customProxy:      false,
			expectViaProxy:   false,
			expectDirectCall: true,
		},
		{
			name:             "custom CONDUCTOR_PROXY - goes through proxy",
			customProxy:      true,
			expectViaProxy:   true,
			expectDirectCall: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear env vars
			os.Unsetenv(settings.EnvProxy)

			// Create TARGET server (main API) and optional PROXY server first
			var targetCalled bool
			var proxyCalled bool

			targetServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				targetCalled = true

				// Handle authentication token endpoint
				if r.URL.Path == "/token" {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"access_token":"mock-token","token_type":"Bearer","expires_in":3600}`))
					return
				}

				// Handle version endpoint
				if r.URL.Path == "/version" {
					w.Header().Set("Content-Type", "text/plain")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(versionDirect))
					return
				}

				// Default response
				w.WriteHeader(http.StatusNotFound)
			}))
			defer targetServer.Close()

			var proxyServer *httptest.Server
			if tt.customProxy {
				proxyServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					proxyCalled = true

					// Handle authentication token endpoint
					if r.URL.Path == "/token" {
						w.Header().Set("Content-Type", "application/json")
						w.Header().Set("X-Via-Proxy", "true")
						w.WriteHeader(http.StatusOK)
						w.Write([]byte(`{"access_token":"mock-token","token_type":"Bearer","expires_in":3600}`))
						return
					}

					// Handle version endpoint
					if r.URL.Path == "/version" {
						w.Header().Set("Content-Type", "text/plain")
						w.Header().Set("X-Via-Proxy", "true")
						w.WriteHeader(http.StatusOK)
						w.Write([]byte(versionViaProxy))
						return
					}

					// Default response
					w.WriteHeader(http.StatusNotFound)
				}))
				defer proxyServer.Close()
			}

			// Configure server URL and proxy
			os.Setenv("CONDUCTOR_SERVER_URL", targetServer.URL)
			if tt.customProxy {
				// For custom proxy: CONDUCTOR_PROXY points to proxy server
				os.Setenv(settings.EnvProxy, proxyServer.URL)
			}

			// Create API client
			apiClient := client.NewAPIClientFromEnv()
			require.NotNil(t, apiClient)

			// Create VersionResourceClient
			versionClient := client.NewVersionResourceClient(apiClient)
			require.NotNil(t, versionClient)

			// Make the request
			ctx := context.Background()
			version, resp, err := versionClient.GetVersion(ctx)

			// Verify result
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			// Verify path used
			if tt.expectDirectCall {
				assert.True(t, targetCalled, "Target server should have received the request")
				assert.False(t, proxyCalled, "Proxy server should NOT have received the request")
			}
			if tt.expectViaProxy {
				assert.True(t, proxyCalled, "Proxy server should have received the request")
			}

			// Check version string directly
			expectedVersion := versionDirect
			if tt.expectViaProxy {
				expectedVersion = versionViaProxy
			}
			assert.Equal(t, expectedVersion, version, "Version should match expected value")
		})
	}
}

func TestTypeAssertions_TokenManagerInterface(t *testing.T) {
	t.Run("authentication.TokenManager can be passed to NewAPIClientWithTokenManager", func(t *testing.T) {
		// Test that authentication.TokenManager can be passed to NewAPIClientWithTokenManager
		authSettings := settings.NewAuthenticationSettings("test-key", "test-secret")
		httpSettings := settings.NewHttpDefaultSettings()
		tokenExpiration := authentication.NewDefaultTokenExpiration()
		authTokenManager := authentication.NewTokenManager(*authSettings, tokenExpiration)

		// This should compile and work
		apiClient := client.NewAPIClientWithTokenManager(
			authSettings,
			httpSettings,
			tokenExpiration,
			authTokenManager,
		)

		require.NotNil(t, apiClient)
		assert.NotNil(t, apiClient)
	})

	t.Run("NewAPIClient with TokenExpiration but without TokenManager", func(t *testing.T) {
		// Create authentication settings and token expiration
		authSettings := settings.NewAuthenticationSettings("test-key", "test-secret")
		httpSettings := settings.NewHttpDefaultSettings()
		tokenExpiration := authentication.NewDefaultTokenExpiration()

		// Create client settings with TokenExpiration but WITHOUT TokenManager
		clientSettings := &settings.ClientSettings{
			Authentication:  authSettings,
			HTTP:            httpSettings,
			TokenExpiration: tokenExpiration,
			TokenManager:    nil, // Explicitly nil - should auto-create
		}

		// Create API client - should auto-create TokenManager with our TokenExpiration
		apiClient := client.NewAPIClientFromSettings(clientSettings)
		require.NotNil(t, apiClient)
		assert.NotNil(t, apiClient)
	})

	t.Run("NewAPIClientWithTokenExpiration creates TokenManager automatically", func(t *testing.T) {
		// Test the specific constructor that takes TokenExpiration
		authSettings := settings.NewAuthenticationSettings("test-key", "test-secret")
		httpSettings := settings.NewHttpDefaultSettings()
		tokenExpiration := authentication.NewDefaultTokenExpiration()

		// This should create a TokenManager automatically using the provided TokenExpiration
		apiClient := client.NewAPIClientWithTokenExpiration(
			authSettings,
			httpSettings,
			tokenExpiration,
		)

		require.NotNil(t, apiClient)
		assert.NotNil(t, apiClient)
	})
}
