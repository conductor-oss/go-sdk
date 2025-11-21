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
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithProxyURL(t *testing.T) {
	tests := []struct {
		name        string
		urlStr      string
		expectError bool
		expectedURL string
	}{
		{
			name:        "valid HTTP proxy URL",
			urlStr:      "http://proxy.example.com:8080",
			expectError: false,
			expectedURL: "http://proxy.example.com:8080",
		},
		{
			name:        "valid HTTPS proxy URL",
			urlStr:      "https://proxy.example.com:8080",
			expectError: false,
			expectedURL: "https://proxy.example.com:8080",
		},
		{
			name:        "valid proxy URL with credentials",
			urlStr:      "http://user:pass@proxy.example.com:8080",
			expectError: false,
			expectedURL: "http://user:pass@proxy.example.com:8080",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := NewClientSettings(WithProxyURL(tt.urlStr))

			require.NotNil(t, settings.Proxy, "Proxy settings should be created")
			require.NotNil(t, settings.Proxy.URL, "Proxy URL should be set")
			assert.Equal(t, tt.expectedURL, settings.Proxy.URL.String())
		})
	}
}

func TestWithProxy(t *testing.T) {
	tests := []struct {
		name        string
		proxyURL    *url.URL
		expectedURL string
	}{
		{
			name:        "valid HTTP proxy URL",
			proxyURL:    &url.URL{Scheme: "http", Host: "proxy.example.com:8080"},
			expectedURL: "http://proxy.example.com:8080",
		},
		{
			name:        "valid HTTPS proxy URL",
			proxyURL:    &url.URL{Scheme: "https", Host: "proxy.example.com:8080"},
			expectedURL: "https://proxy.example.com:8080",
		},
		{
			name:        "proxy URL with credentials",
			proxyURL:    &url.URL{Scheme: "http", Host: "proxy.example.com:8080", User: url.UserPassword("user", "pass")},
			expectedURL: "http://user:pass@proxy.example.com:8080",
		},
		{
			name:        "nil proxy URL",
			proxyURL:    nil,
			expectedURL: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := NewClientSettings(WithProxy(tt.proxyURL))

			require.NotNil(t, settings.Proxy, "Proxy settings should be created")
			if tt.proxyURL != nil {
				require.NotNil(t, settings.Proxy.URL, "Proxy URL should be set")
				assert.Equal(t, tt.expectedURL, settings.Proxy.URL.String())
			} else {
				assert.Nil(t, settings.Proxy.URL, "Proxy URL should be nil")
			}
		})
	}
}

func TestWithProxyCredentials(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
	}{
		{
			name:     "valid credentials",
			username: "testuser",
			password: "testpass",
		},
		{
			name:     "empty credentials",
			username: "",
			password: "",
		},
		{
			name:     "username only",
			username: "testuser",
			password: "",
		},
		{
			name:     "special characters in credentials",
			username: "user@domain",
			password: "p@ss:w0rd!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := NewClientSettings(WithProxyCredentials(tt.username, tt.password))

			require.NotNil(t, settings.Proxy, "Proxy settings should be created")
			assert.Equal(t, tt.username, settings.Proxy.Username)
			assert.Equal(t, tt.password, settings.Proxy.Password)
		})
	}
}

func TestWithAuthCredentials(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		secret string
	}{
		{
			name:   "valid credentials",
			key:    "testkey",
			secret: "testsecret",
		},
		{
			name:   "empty credentials",
			key:    "",
			secret: "",
		},
		{
			name:   "key only",
			key:    "testkey",
			secret: "",
		},
		{
			name:   "special characters in credentials",
			key:    "key@domain",
			secret: "s3cr3t!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := NewClientSettings(WithAuthCredentials(tt.key, tt.secret))

			require.NotNil(t, settings.Authentication, "Authentication settings should be created")
			assert.Equal(t, tt.key, settings.Authentication.GetBody()["keyId"])
			assert.Equal(t, tt.secret, settings.Authentication.GetBody()["keySecret"])
		})
	}
}

func TestWithServerURL(t *testing.T) {
	tests := []struct {
		name        string
		serverURL   string
		expectedURL string
	}{
		{
			name:        "valid HTTP URL",
			serverURL:   "http://api.example.com",
			expectedURL: "http://api.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := NewClientSettings(WithServerURL(tt.serverURL))

			require.NotNil(t, settings.HTTP, "HTTP settings should be created")
			assert.Equal(t, tt.expectedURL, settings.HTTP.BaseUrl)
		})
	}
}

func TestWithHTTPTimeout(t *testing.T) {
	tests := []struct {
		name            string
		timeout         time.Duration
		expectedTimeout time.Duration
	}{
		{
			name:            "30 seconds",
			timeout:         30 * time.Second,
			expectedTimeout: 30 * time.Second,
		},
		{
			name:            "zero timeout",
			timeout:         0,
			expectedTimeout: 30 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := NewClientSettings(WithHTTPTimeout(tt.timeout))

			require.NotNil(t, settings.HTTP, "HTTP settings should be created")
			assert.Equal(t, tt.expectedTimeout, settings.HTTP.Timeout)
		})
	}
}

func TestWithHTTPHeaders(t *testing.T) {
	tests := []struct {
		name            string
		headers         map[string]string
		expectedHeaders map[string]string
	}{
		{
			name: "single header",
			headers: map[string]string{
				"Authorization": "Bearer token123",
			},
			expectedHeaders: map[string]string{
				"Content-Type":    "application/json",
				"Accept":          "application/json",
				"Accept-Encoding": "gzip",
				"Authorization":   "Bearer token123",
			},
		},
		{
			name: "multiple headers",
			headers: map[string]string{
				"Authorization": "Bearer token123",
				"X-Custom":      "value",
			},
			expectedHeaders: map[string]string{
				"Content-Type":    "application/json",
				"Accept":          "application/json",
				"Accept-Encoding": "gzip",
				"Authorization":   "Bearer token123",
				"X-Custom":        "value",
			},
		},
		{
			name: "override default header",
			headers: map[string]string{
				"Content-Type": "application/xml",
			},
			expectedHeaders: map[string]string{
				"Content-Type":    "application/xml",
				"Accept":          "application/json",
				"Accept-Encoding": "gzip",
			},
		},
		{
			name:    "empty headers",
			headers: map[string]string{},
			expectedHeaders: map[string]string{
				"Content-Type":    "application/json",
				"Accept":          "application/json",
				"Accept-Encoding": "gzip",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := NewClientSettings(WithHTTPHeaders(tt.headers))

			require.NotNil(t, settings.HTTP, "HTTP settings should be created")
			assert.Equal(t, tt.expectedHeaders, settings.HTTP.Headers)
		})
	}
}

func TestMultipleOptions(t *testing.T) {
	t.Run("combine multiple options", func(t *testing.T) {
		proxyURL, _ := url.Parse("http://proxy.example.com:8080")
		settings := NewClientSettings(
			WithProxy(proxyURL),
			WithProxyCredentials("user", "pass"),
			WithAuthCredentials("key", "secret"),
			WithServerURL("https://api.example.com"),
			WithHTTPTimeout(60*time.Second),
			WithHTTPHeaders(map[string]string{"X-Custom": "value"}),
		)

		// Check proxy settings
		require.NotNil(t, settings.Proxy)
		assert.Equal(t, "http://proxy.example.com:8080", settings.Proxy.URL.String())
		assert.Equal(t, "user", settings.Proxy.Username)
		assert.Equal(t, "pass", settings.Proxy.Password)

		// Check authentication settings
		require.NotNil(t, settings.Authentication)
		assert.Equal(t, "key", settings.Authentication.GetBody()["keyId"])
		assert.Equal(t, "secret", settings.Authentication.GetBody()["keySecret"])

		// Check HTTP settings
		require.NotNil(t, settings.HTTP)
		assert.Equal(t, "https://api.example.com", settings.HTTP.BaseUrl)
		assert.Equal(t, 60*time.Second, settings.HTTP.Timeout)
		assert.Equal(t, "value", settings.HTTP.Headers["X-Custom"])
	})
}

func TestApplyOptions(t *testing.T) {
	t.Run("apply options to existing settings", func(t *testing.T) {
		settings := NewClientSettings()

		// Apply options after creation
		proxyURL, _ := url.Parse("http://proxy.example.com:8080")
		settings.ApplyOptions(
			WithProxy(proxyURL),
			WithAuthCredentials("key", "secret"),
		)

		// Check that options were applied
		require.NotNil(t, settings.Proxy)
		assert.Equal(t, "http://proxy.example.com:8080", settings.Proxy.URL.String())

		require.NotNil(t, settings.Authentication)
		assert.Equal(t, "key", settings.Authentication.GetBody()["keyId"])
		assert.Equal(t, "secret", settings.Authentication.GetBody()["keySecret"])
	})
}

func TestWithTokenExpiration(t *testing.T) {
	tests := []struct {
		name               string
		tokenExpiration    TokenExpirationInterface
		expectedExpiration time.Duration
		expectedCleanup    time.Duration
	}{
		{
			name: "custom token expiration",
			tokenExpiration: &mockTokenExpiration{
				defaultExpiration: 45 * time.Minute,
				cleanupInterval:   3 * time.Hour,
			},
			expectedExpiration: 45 * time.Minute,
			expectedCleanup:    3 * time.Hour,
		},
		{
			name:               "nil token expiration",
			tokenExpiration:    nil,
			expectedExpiration: 0,
			expectedCleanup:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := NewClientSettings(WithTokenExpiration(tt.tokenExpiration))

			if tt.tokenExpiration != nil {
				require.NotNil(t, settings.TokenExpiration, "Token expiration should be set")
				assert.Equal(t, tt.expectedExpiration, settings.TokenExpiration.GetDefaultExpiration())
				assert.Equal(t, tt.expectedCleanup, settings.TokenExpiration.GetCleanupInterval())
			} else {
				assert.Nil(t, settings.TokenExpiration, "Token expiration should be nil")
			}
		})
	}
}

func TestWithTokenManager(t *testing.T) {
	tests := []struct {
		name         string
		tokenManager TokenManagerInterface
		expectNil    bool
	}{
		{
			name:         "custom token manager",
			tokenManager: &mockTokenManager{},
			expectNil:    false,
		},
		{
			name:         "nil token manager",
			tokenManager: nil,
			expectNil:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := NewClientSettings(WithTokenManager(tt.tokenManager))

			if tt.expectNil {
				assert.Nil(t, settings.TokenManager, "Token manager should be nil")
			} else {
				require.NotNil(t, settings.TokenManager, "Token manager should be set")
				assert.Equal(t, tt.tokenManager, settings.TokenManager)
			}
		})
	}
}

// Mock implementations for testing
type mockTokenExpiration struct {
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
}

func (m *mockTokenExpiration) GetDefaultExpiration() time.Duration {
	return m.defaultExpiration
}

func (m *mockTokenExpiration) GetCleanupInterval() time.Duration {
	return m.cleanupInterval
}

type mockTokenManager struct{}

func (m *mockTokenManager) RefreshToken(httpSettings *HttpSettings, httpClient *http.Client) (string, error) {
	return "mock-token", nil
}

// ============= TLS Options Tests =============

// generateTestCertForOptions generates a self-signed certificate for testing TLS options
func generateTestCertForOptions() (certPEM, keyPEM []byte, err error) {
	// Generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Create certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "test-cert-options",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// Create self-signed certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, err
	}

	// Encode certificate to PEM
	certPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	})

	// Encode private key to PEM
	keyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	return certPEM, keyPEM, nil
}

func TestTLSOptions(t *testing.T) {
	// Generate test certificate and key for tests that need them
	certPEM, keyPEM, err := generateTestCertForOptions()
	require.NoError(t, err)

	// Create temporary files
	certFile, err := os.CreateTemp("", "test-cert-*.pem")
	require.NoError(t, err)
	t.Cleanup(func() { os.Remove(certFile.Name()) })
	_, err = certFile.Write(certPEM)
	require.NoError(t, err)
	certFile.Close()

	keyFile, err := os.CreateTemp("", "test-key-*.pem")
	require.NoError(t, err)
	t.Cleanup(func() { os.Remove(keyFile.Name()) })
	_, err = keyFile.Write(keyPEM)
	require.NoError(t, err)
	keyFile.Close()

	t.Run("WithInsecureSkipVerify", func(t *testing.T) {
		settings := NewClientSettings(WithInsecureSkipVerify(true))

		require.NotNil(t, settings.TLS)
		assert.True(t, settings.TLS.InsecureSkipVerify)
	})

	t.Run("WithCACertFromFile_valid", func(t *testing.T) {
		settings := NewClientSettings(WithCACertFromFile(certFile.Name()))

		require.NotNil(t, settings.TLS)
		assert.NotNil(t, settings.TLS.RootCAs)
	})

	t.Run("WithCACertFromFile_invalid", func(t *testing.T) {
		settings := NewClientSettings(WithCACertFromFile("/non/existent/path.pem"))

		// Should not panic, just log error and skip
		require.NotNil(t, settings.TLS)
	})

	t.Run("WithCACertFromPEM_valid", func(t *testing.T) {
		settings := NewClientSettings(WithCACertFromPEM(certPEM))

		require.NotNil(t, settings.TLS)
		assert.NotNil(t, settings.TLS.RootCAs)
	})

	t.Run("WithCACertFromPEM_invalid", func(t *testing.T) {
		settings := NewClientSettings(WithCACertFromPEM([]byte("invalid pem")))

		// Should not panic, just log error and skip
		require.NotNil(t, settings.TLS)
		assert.Nil(t, settings.TLS.RootCAs)
	})

	t.Run("WithClientCert_valid", func(t *testing.T) {
		settings := NewClientSettings(WithClientCert(certFile.Name(), keyFile.Name()))

		require.NotNil(t, settings.TLS)
		assert.Len(t, settings.TLS.Certificates, 1)
	})

	t.Run("WithClientCert_invalid", func(t *testing.T) {
		settings := NewClientSettings(WithClientCert("/non/existent/cert.pem", "/non/existent/key.pem"))

		// Should not panic, just log error and skip
		require.NotNil(t, settings.TLS)
		assert.Empty(t, settings.TLS.Certificates)
	})

	t.Run("WithClientCertFromPEM_valid", func(t *testing.T) {
		settings := NewClientSettings(WithClientCertFromPEM(certPEM, keyPEM))

		require.NotNil(t, settings.TLS)
		assert.Len(t, settings.TLS.Certificates, 1)
	})

	t.Run("WithClientCertFromPEM_invalid_cert", func(t *testing.T) {
		settings := NewClientSettings(WithClientCertFromPEM([]byte("invalid cert"), keyPEM))

		// Should not panic, just log error and skip
		require.NotNil(t, settings.TLS)
		assert.Empty(t, settings.TLS.Certificates)
	})

	t.Run("WithClientCertFromPEM_invalid_key", func(t *testing.T) {
		settings := NewClientSettings(WithClientCertFromPEM(certPEM, []byte("invalid key")))

		// Should not panic, just log error and skip
		require.NotNil(t, settings.TLS)
		assert.Empty(t, settings.TLS.Certificates)
	})

	t.Run("WithTLSServerName", func(t *testing.T) {
		settings := NewClientSettings(WithTLSServerName("example.com"))

		require.NotNil(t, settings.TLS)
		assert.Equal(t, "example.com", settings.TLS.ServerName)
	})

	t.Run("WithSelfSignedCert_valid", func(t *testing.T) {
		settings := NewClientSettings(WithSelfSignedCert(certFile.Name()))

		require.NotNil(t, settings.TLS)
		assert.NotNil(t, settings.TLS.RootCAs)
	})

	t.Run("WithSelfSignedCert_invalid", func(t *testing.T) {
		settings := NewClientSettings(WithSelfSignedCert("/non/existent/path.pem"))

		// Should not panic, just log error and skip
		require.NotNil(t, settings.TLS)
	})

	t.Run("WithTLSSettings", func(t *testing.T) {
		customTLS := &TLSSettings{
			InsecureSkipVerify: true,
			ServerName:         "custom.example.com",
		}
		settings := NewClientSettings(WithTLSSettings(customTLS))

		require.NotNil(t, settings.TLS)
		assert.True(t, settings.TLS.InsecureSkipVerify)
		assert.Equal(t, "custom.example.com", settings.TLS.ServerName)
	})

	t.Run("multiple_options", func(t *testing.T) {
		settings := NewClientSettings(
			WithTLSServerName("multi.example.com"),
			WithCACertFromPEM(certPEM),
			WithClientCert(certFile.Name(), keyFile.Name()),
		)

		require.NotNil(t, settings.TLS)
		assert.Equal(t, "multi.example.com", settings.TLS.ServerName)
		assert.NotNil(t, settings.TLS.RootCAs)
		assert.Len(t, settings.TLS.Certificates, 1)
	})

	t.Run("multiple_options_with_pem", func(t *testing.T) {
		settings := NewClientSettings(
			WithTLSServerName("multi.example.com"),
			WithCACertFromPEM(certPEM),
			WithClientCertFromPEM(certPEM, keyPEM),
		)

		require.NotNil(t, settings.TLS)
		assert.Equal(t, "multi.example.com", settings.TLS.ServerName)
		assert.NotNil(t, settings.TLS.RootCAs)
		assert.Len(t, settings.TLS.Certificates, 1)
	})
}
