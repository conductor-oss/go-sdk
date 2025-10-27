//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package settings

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProxyConfigFromEnvironment(t *testing.T) {
	// Cleanup
	originalProxy := os.Getenv(EnvProxy)
	defer func() {
		if originalProxy != "" {
			os.Setenv(EnvProxy, originalProxy)
		} else {
			os.Unsetenv(EnvProxy)
		}
	}()

	tests := []struct {
		name        string
		proxyEnv    string
		expectProxy bool
		expectError bool
	}{
		{
			name:        "no proxy configured",
			proxyEnv:    "",
			expectProxy: false,
			expectError: false,
		},
		{
			name:        "simple proxy URL",
			proxyEnv:    "http://proxy.example.com:8080",
			expectProxy: true,
			expectError: false,
		},
		{
			name:        "proxy with credentials",
			proxyEnv:    "http://user:pass@proxy.example.com:8080",
			expectProxy: true,
			expectError: false,
		},
		{
			name:        "invalid proxy URL",
			proxyEnv:    "://invalid",
			expectProxy: false,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup environment
			if tt.proxyEnv != "" {
				os.Setenv(EnvProxy, tt.proxyEnv)
			} else {
				os.Unsetenv(EnvProxy)
			}

			// Load proxy settings
			proxySettings := newProxySettings()
			err := loadProxyFromEnv(proxySettings)

			// Check error expectation
			if tt.expectError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)

			// Check proxy configuration
			if tt.expectProxy {
				assert.NotNil(t, proxySettings.URL, "Proxy URL should be set")
			} else {
				assert.Nil(t, proxySettings.URL, "Proxy URL should not be set")
			}
		})
	}
}

func TestProxyConfigCredentials(t *testing.T) {
	originalProxy := os.Getenv(EnvProxy)
	defer func() {
		if originalProxy != "" {
			os.Setenv(EnvProxy, originalProxy)
		} else {
			os.Unsetenv(EnvProxy)
		}
	}()

	tests := []struct {
		name             string
		proxyURL         string
		expectedUsername string
		expectedPassword string
	}{
		{
			name:             "no credentials",
			proxyURL:         "http://proxy.example.com:8080",
			expectedUsername: "",
			expectedPassword: "",
		},
		{
			name:             "username only",
			proxyURL:         "http://admin@proxy.example.com:8080",
			expectedUsername: "admin",
			expectedPassword: "",
		},
		{
			name:             "username and password",
			proxyURL:         "http://admin:secret@proxy.example.com:8080",
			expectedUsername: "admin",
			expectedPassword: "secret",
		},
		{
			name:             "special characters in password",
			proxyURL:         "http://user:p@ss:w0rd!@proxy.example.com:8080",
			expectedUsername: "user",
			expectedPassword: "p@ss:w0rd!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(EnvProxy, tt.proxyURL)

			proxySettings := NewProxySettingsFromEnv()

			assert.Equal(t, tt.expectedUsername, proxySettings.Username)
			assert.Equal(t, tt.expectedPassword, proxySettings.Password)

			if proxySettings.URL != nil {
				assert.Nil(t, proxySettings.URL.User, "URL should not contain credentials")
			}
		})
	}
}

func TestBuildProxyFuncUsesSystemProxy(t *testing.T) {
	// Clean up all proxy environment variables at the end
	originalProxy := os.Getenv(EnvProxy)
	originalHTTPProxy := os.Getenv("HTTP_PROXY")
	defer func() {
		if originalProxy != "" {
			os.Setenv(EnvProxy, originalProxy)
		} else {
			os.Unsetenv(EnvProxy)
		}
		if originalHTTPProxy != "" {
			os.Setenv("HTTP_PROXY", originalHTTPProxy)
		} else {
			os.Unsetenv("HTTP_PROXY")
		}

	}()

	t.Run("no custom proxy uses system proxy", func(t *testing.T) {
		// Clear all proxy settings first
		os.Unsetenv(EnvProxy)
		os.Unsetenv("HTTP_PROXY")

		// Set system proxy
		os.Setenv("HTTP_PROXY", "http://system-proxy.example.com:8080")

		proxySettings := NewProxySettingsFromEnv()
		proxyFunc := proxySettings.BuildProxyFunc()

		require.NotNil(t, proxyFunc, "Proxy function should always be created")

		// Test that it uses system proxy
		req, _ := http.NewRequest("GET", "http://example.com", nil)
		resultURL, err := proxyFunc(req)

		require.NoError(t, err)
		assert.Equal(t, "http://system-proxy.example.com:8080", resultURL.String())

	})

	t.Run("custom proxy overrides system proxy", func(t *testing.T) {
		// Clear all proxy settings first
		os.Unsetenv(EnvProxy)
		os.Unsetenv("HTTP_PROXY")

		// Set custom proxy
		os.Setenv(EnvProxy, "http://custom-proxy.example.com:8080")

		// Set system proxy (should be ignored)
		os.Setenv("HTTP_PROXY", "http://system-proxy.example.com:8080")

		// Clean up after this test
		t.Cleanup(func() {
			os.Unsetenv(EnvProxy)
			os.Unsetenv("HTTP_PROXY")
		})

		proxySettings := NewProxySettingsFromEnv()
		proxyFunc := proxySettings.BuildProxyFunc()

		require.NotNil(t, proxyFunc, "Proxy function should always be created")

		// Test that it uses custom proxy, not system proxy
		req, _ := http.NewRequest("GET", "http://example.com", nil)
		resultURL, err := proxyFunc(req)

		require.NoError(t, err)

		// Should return custom proxy, not system proxy
		assert.NotNil(t, resultURL, "Should return custom proxy URL")
		assert.Equal(t, "http://custom-proxy.example.com:8080", resultURL.String())
	})

}
