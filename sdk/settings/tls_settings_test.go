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
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// generateTestCert generates a self-signed certificate for testing
func generateTestCert() (certPEM, keyPEM []byte, err error) {
	// Generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Create certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "test-cert",
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

func TestNewTLSDefaultSettings(t *testing.T) {
	settings := NewTLSDefaultSettings()

	require.NotNil(t, settings)
	assert.False(t, settings.InsecureSkipVerify)
	assert.Nil(t, settings.RootCAs)
	assert.Empty(t, settings.ServerName)
	assert.Empty(t, settings.Certificates)
}

func TestLoadCACert(t *testing.T) {
	// Generate test certificate for valid cases
	certPEM, _, err := generateTestCert()
	require.NoError(t, err, "Failed to generate test cert")

	// Create temporary cert file
	tmpFile, err := os.CreateTemp("", "test-cert-*.pem")
	require.NoError(t, err)
	t.Cleanup(func() { os.Remove(tmpFile.Name()) })

	_, err = tmpFile.Write(certPEM)
	require.NoError(t, err)
	tmpFile.Close()

	tests := []struct {
		name         string
		testFunc     func(*TLSSettings) error
		shouldError  bool
		shouldHaveCA bool
	}{
		{
			name: "loadCACertFromFile_valid",
			testFunc: func(s *TLSSettings) error {
				return s.loadCACertFromFile(tmpFile.Name())
			},
			shouldError:  false,
			shouldHaveCA: true,
		},
		{
			name: "loadCACertFromFile_nonexistent",
			testFunc: func(s *TLSSettings) error {
				return s.loadCACertFromFile("/non/existent/path.pem")
			},
			shouldError:  true,
			shouldHaveCA: false,
		},
		{
			name: "loadCACertFromPEM_valid",
			testFunc: func(s *TLSSettings) error {
				return s.loadCACertFromPEM(certPEM)
			},
			shouldError:  false,
			shouldHaveCA: true,
		},
		{
			name: "loadCACertFromPEM_invalid",
			testFunc: func(s *TLSSettings) error {
				return s.loadCACertFromPEM([]byte("invalid pem data"))
			},
			shouldError:  true,
			shouldHaveCA: false,
		},
		{
			name: "loadCACertFromFileWithSystemCerts_valid",
			testFunc: func(s *TLSSettings) error {
				return s.loadCACertFromFileWithSystemCerts(tmpFile.Name())
			},
			shouldError:  false,
			shouldHaveCA: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := NewTLSDefaultSettings()
			err := tt.testFunc(settings)

			if tt.shouldError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.shouldHaveCA {
				assert.NotNil(t, settings.RootCAs)
			} else {
				assert.Nil(t, settings.RootCAs)
			}
		})
	}
}

func TestLoadClientCertFromPEM(t *testing.T) {
	// Generate test certificate and key
	certPEM, keyPEM, err := generateTestCert()
	require.NoError(t, err, "Failed to generate test cert")

	tests := []struct {
		name        string
		certPEM     []byte
		keyPEM      []byte
		shouldError bool
		certCount   int
	}{
		{
			name:        "valid_cert_and_key",
			certPEM:     certPEM,
			keyPEM:      keyPEM,
			shouldError: false,
			certCount:   1,
		},
		{
			name:        "invalid_cert_pem",
			certPEM:     []byte("invalid pem data"),
			keyPEM:      keyPEM,
			shouldError: true,
			certCount:   0,
		},
		{
			name:        "invalid_key_pem",
			certPEM:     certPEM,
			keyPEM:      []byte("invalid pem data"),
			shouldError: true,
			certCount:   0,
		},
		{
			name:        "empty_cert",
			certPEM:     []byte(""),
			keyPEM:      keyPEM,
			shouldError: true,
			certCount:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := NewTLSDefaultSettings()
			err := settings.loadClientCertFromPEM(tt.certPEM, tt.keyPEM)

			if tt.shouldError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, settings.Certificates, tt.certCount)
			}
		})
	}
}

func TestLoadClientCertFromFiles(t *testing.T) {
	// Generate test certificate and key
	certPEM, keyPEM, err := generateTestCert()
	require.NoError(t, err, "Failed to generate test cert")

	// Create temporary cert file
	certFile, err := os.CreateTemp("", "test-cert-*.pem")
	require.NoError(t, err)
	t.Cleanup(func() { os.Remove(certFile.Name()) })
	_, err = certFile.Write(certPEM)
	require.NoError(t, err)
	certFile.Close()

	// Create temporary key file
	keyFile, err := os.CreateTemp("", "test-key-*.pem")
	require.NoError(t, err)
	t.Cleanup(func() { os.Remove(keyFile.Name()) })
	_, err = keyFile.Write(keyPEM)
	require.NoError(t, err)
	keyFile.Close()

	tests := []struct {
		name        string
		certPath    string
		keyPath     string
		shouldError bool
		certCount   int
	}{
		{
			name:        "valid_cert_and_key",
			certPath:    certFile.Name(),
			keyPath:     keyFile.Name(),
			shouldError: false,
			certCount:   1,
		},
		{
			name:        "nonexistent_files",
			certPath:    "/non/existent/cert.pem",
			keyPath:     "/non/existent/key.pem",
			shouldError: true,
			certCount:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := NewTLSDefaultSettings()
			err := settings.loadClientCertFromFiles(tt.certPath, tt.keyPath)

			if tt.shouldError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, settings.Certificates, tt.certCount)
			}
		})
	}
}

func TestBuildTLSConfig(t *testing.T) {
	// Generate test certificate for tests that need it
	certPEM, keyPEM, err := generateTestCert()
	require.NoError(t, err, "Failed to generate test cert")

	// Create temporary files for client cert test
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

	tests := []struct {
		name               string
		setupSettings      func(t *testing.T) *TLSSettings
		shouldReturnConfig bool
		validateConfig     func(*testing.T, *tls.Config)
	}{
		{
			name: "nil_settings",
			setupSettings: func(t *testing.T) *TLSSettings {
				return nil
			},
			shouldReturnConfig: false,
		},
		{
			name: "default_settings",
			setupSettings: func(t *testing.T) *TLSSettings {
				return NewTLSDefaultSettings()
			},
			shouldReturnConfig: false,
		},
		{
			name: "insecure_skip_verify",
			setupSettings: func(t *testing.T) *TLSSettings {
				s := NewTLSDefaultSettings()
				s.InsecureSkipVerify = true
				return s
			},
			shouldReturnConfig: true,
			validateConfig: func(t *testing.T, config *tls.Config) {
				assert.True(t, config.InsecureSkipVerify)
				assert.Equal(t, uint16(tls.VersionTLS12), config.MinVersion)
			},
		},
		{
			name: "with_root_cas",
			setupSettings: func(t *testing.T) *TLSSettings {
				s := NewTLSDefaultSettings()
				err := s.loadCACertFromPEM(certPEM)
				require.NoError(t, err)
				return s
			},
			shouldReturnConfig: true,
			validateConfig: func(t *testing.T, config *tls.Config) {
				assert.NotNil(t, config.RootCAs)
				assert.False(t, config.InsecureSkipVerify)
			},
		},
		{
			name: "with_server_name",
			setupSettings: func(t *testing.T) *TLSSettings {
				s := NewTLSDefaultSettings()
				s.ServerName = "example.com"
				return s
			},
			shouldReturnConfig: true,
			validateConfig: func(t *testing.T, config *tls.Config) {
				assert.Equal(t, "example.com", config.ServerName)
			},
		},
		{
			name: "with_client_cert",
			setupSettings: func(t *testing.T) *TLSSettings {
				s := NewTLSDefaultSettings()
				err := s.loadClientCertFromFiles(certFile.Name(), keyFile.Name())
				require.NoError(t, err)
				return s
			},
			shouldReturnConfig: true,
			validateConfig: func(t *testing.T, config *tls.Config) {
				assert.Len(t, config.Certificates, 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := tt.setupSettings(t)
			config := settings.BuildTLSConfig()

			if tt.shouldReturnConfig {
				require.NotNil(t, config)
				if tt.validateConfig != nil {
					tt.validateConfig(t, config)
				}
			} else {
				assert.Nil(t, config)
			}
		})
	}
}

func TestNewTLSSettingsFromEnv(t *testing.T) {
	t.Run("insecure_skip_verify_true", func(t *testing.T) {
		t.Setenv(EnvTLSInsecureSkipVerify, "true")
		t.Setenv(EnvTLSCACert, "")

		settings := NewTLSSettingsFromEnv()

		assert.True(t, settings.InsecureSkipVerify)
	})

	t.Run("insecure_skip_verify_false", func(t *testing.T) {
		t.Setenv(EnvTLSInsecureSkipVerify, "false")
		t.Setenv(EnvTLSCACert, "")

		settings := NewTLSSettingsFromEnv()

		assert.False(t, settings.InsecureSkipVerify)
	})

	t.Run("no_env_vars", func(t *testing.T) {
		t.Setenv(EnvTLSInsecureSkipVerify, "")
		t.Setenv(EnvTLSCACert, "")
		t.Setenv(EnvTLSClientCert, "")
		t.Setenv(EnvTLSClientKey, "")

		settings := NewTLSSettingsFromEnv()

		assert.False(t, settings.InsecureSkipVerify)
		assert.Nil(t, settings.RootCAs)
	})

	t.Run("with_ca_cert_file", func(t *testing.T) {
		// Generate test certificate
		certPEM, _, err := generateTestCert()
		require.NoError(t, err)

		// Create temporary cert file
		tmpFile, err := os.CreateTemp("", "test-ca-*.pem")
		require.NoError(t, err)
		t.Cleanup(func() { os.Remove(tmpFile.Name()) })
		_, err = tmpFile.Write(certPEM)
		require.NoError(t, err)
		tmpFile.Close()

		t.Setenv(EnvTLSCACert, tmpFile.Name())

		settings := NewTLSSettingsFromEnv()

		assert.NotNil(t, settings.RootCAs)
	})

	t.Run("with_client_cert_files", func(t *testing.T) {
		// Generate test certificate and key
		certPEM, keyPEM, err := generateTestCert()
		require.NoError(t, err)

		// Create temporary cert file
		certFile, err := os.CreateTemp("", "test-client-cert-*.pem")
		require.NoError(t, err)
		t.Cleanup(func() { os.Remove(certFile.Name()) })
		_, err = certFile.Write(certPEM)
		require.NoError(t, err)
		certFile.Close()

		// Create temporary key file
		keyFile, err := os.CreateTemp("", "test-client-key-*.pem")
		require.NoError(t, err)
		t.Cleanup(func() { os.Remove(keyFile.Name()) })
		_, err = keyFile.Write(keyPEM)
		require.NoError(t, err)
		keyFile.Close()

		t.Setenv(EnvTLSClientCert, certFile.Name())
		t.Setenv(EnvTLSClientKey, keyFile.Name())

		settings := NewTLSSettingsFromEnv()

		assert.Len(t, settings.Certificates, 1)
	})
}
