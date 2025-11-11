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
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// generateSelfSignedCert generates a self-signed certificate for testing
func generateSelfSignedCert(t *testing.T) (*x509.Certificate, []byte, *ecdsa.PrivateKey) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Test Org"},
			CommonName:   "test.example.com",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(24 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	require.NoError(t, err)

	cert, err := x509.ParseCertificate(certDER)
	require.NoError(t, err)

	return cert, certDER, privateKey
}

func TestTLSSettings_AllowSelfSigned_Default(t *testing.T) {
	settings := NewTLSDefaultSettings()
	assert.False(t, settings.AllowSelfSigned, "AllowSelfSigned should be false by default")
	assert.Empty(t, settings.PinnedThumbprints, "PinnedThumbprints should be empty by default")
}

func TestWithTlsAllowSelfSigned(t *testing.T) {
	clientSettings := NewClientSettings(
		WithTlsAllowSelfSigned(true),
	)

	assert.NotNil(t, clientSettings.TLS)
	assert.True(t, clientSettings.TLS.AllowSelfSigned)
}

func TestWithTlsPinnedThumbprints(t *testing.T) {
	thumbprints := []string{"abc123", "def456"}
	clientSettings := NewClientSettings(
		WithTlsPinnedThumbprints(thumbprints),
	)

	assert.NotNil(t, clientSettings.TLS)
	assert.Equal(t, thumbprints, clientSettings.TLS.PinnedThumbprints)
}

func TestCalculateThumbprint(t *testing.T) {
	_, certDER, _ := generateSelfSignedCert(t)

	thumbprint := calculateThumbprint(certDER)

	assert.NotEmpty(t, thumbprint)
	assert.Len(t, thumbprint, 64, "SHA-256 thumbprint should be 64 hex characters")

	// Verify it's hex
	for _, c := range thumbprint {
		assert.True(t, (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f'), "thumbprint should be lowercase hex")
	}
}

func TestVerifySelfSignedCertificate_NoPin(t *testing.T) {
	cert, certDER, _ := generateSelfSignedCert(t)

	settings := &TLSSettings{
		AllowSelfSigned:   true,
		PinnedThumbprints: nil,
		ServerName:        "", // No hostname check
	}

	err := settings.verifySelfSignedCertificate([][]byte{certDER}, nil)
	assert.NoError(t, err, "Should accept self-signed cert when AllowSelfSigned=true and no pins")
	assert.NotNil(t, cert)
}

func TestVerifySelfSignedCertificate_WithMatchingPin(t *testing.T) {
	_, certDER, _ := generateSelfSignedCert(t)
	thumbprint := calculateThumbprint(certDER)

	settings := &TLSSettings{
		AllowSelfSigned:   true,
		PinnedThumbprints: []string{thumbprint},
		ServerName:        "",
	}

	err := settings.verifySelfSignedCertificate([][]byte{certDER}, nil)
	assert.NoError(t, err, "Should accept self-signed cert when thumbprint matches")
}

func TestVerifySelfSignedCertificate_WithMatchingPinCaseInsensitive(t *testing.T) {
	_, certDER, _ := generateSelfSignedCert(t)
	thumbprint := calculateThumbprint(certDER)

	// Test with uppercase thumbprint
	settings := &TLSSettings{
		AllowSelfSigned:   true,
		PinnedThumbprints: []string{strings.ToUpper(thumbprint)},
		ServerName:        "",
	}

	err := settings.verifySelfSignedCertificate([][]byte{certDER}, nil)
	assert.NoError(t, err, "Should accept self-signed cert when thumbprint matches (case insensitive)")
}

func TestVerifySelfSignedCertificate_WithNonMatchingPin(t *testing.T) {
	cert, certDER, _ := generateSelfSignedCert(t)

	settings := &TLSSettings{
		AllowSelfSigned:   true,
		PinnedThumbprints: []string{"wrongthumbprint123"},
		ServerName:        "",
	}

	err := settings.verifySelfSignedCertificate([][]byte{certDER}, nil)
	require.Error(t, err, "Should reject self-signed cert when thumbprint doesn't match")

	// Check error message contains expected information
	errorMessage := err.Error()
	assert.Contains(t, errorMessage, "pin mismatch")
	assert.Contains(t, errorMessage, cert.Subject.CommonName)
	assert.Contains(t, errorMessage, "wrongthumbprint123")
}

func TestVerifySelfSignedCertificate_NoCerts(t *testing.T) {
	settings := &TLSSettings{
		AllowSelfSigned: true,
	}

	err := settings.verifySelfSignedCertificate([][]byte{}, nil)
	require.Error(t, err, "Should fail when no certificates provided")
	assert.Contains(t, err.Error(), "no certificates")
}

func TestValidateSelfSignedCert_Expired(t *testing.T) {
	cert, _, _ := generateSelfSignedCert(t)
	// Manually set expired time
	cert.NotAfter = time.Now().Add(-24 * time.Hour)

	settings := &TLSSettings{
		AllowSelfSigned: true,
	}

	err := settings.validateSelfSignedCert(cert)
	require.Error(t, err, "Should reject expired certificate")
	assert.Contains(t, err.Error(), "expired")
}

func TestValidateSelfSignedCert_NotYetValid(t *testing.T) {
	cert, _, _ := generateSelfSignedCert(t)
	// Manually set future start time
	cert.NotBefore = time.Now().Add(24 * time.Hour)

	settings := &TLSSettings{
		AllowSelfSigned: true,
	}

	err := settings.validateSelfSignedCert(cert)
	require.Error(t, err, "Should reject certificate that is not yet valid")
	assert.Contains(t, err.Error(), "not yet valid")
}

func TestValidateSelfSignedCert_HostnameMismatch(t *testing.T) {
	cert, _, _ := generateSelfSignedCert(t)

	settings := &TLSSettings{
		AllowSelfSigned: true,
		ServerName:      "wrong.example.com",
	}

	err := settings.validateSelfSignedCert(cert)
	require.Error(t, err, "Should reject certificate with hostname mismatch")
	assert.Contains(t, err.Error(), "hostname mismatch")
}

func TestValidateSelfSignedCert_HostnameMatch(t *testing.T) {
	cert, _, _ := generateSelfSignedCert(t)

	settings := &TLSSettings{
		AllowSelfSigned: true,
		ServerName:      "test.example.com", // Matches CommonName from generateSelfSignedCert
	}

	err := settings.validateSelfSignedCert(cert)
	assert.NoError(t, err, "Should accept certificate with matching hostname")
}

func TestBuildTLSConfig_WithAllowSelfSigned(t *testing.T) {
	settings := &TLSSettings{
		AllowSelfSigned: true,
	}

	tlsConfig := settings.BuildTLSConfig()

	assert.NotNil(t, tlsConfig)
	assert.True(t, tlsConfig.InsecureSkipVerify, "InsecureSkipVerify should be true when using custom verification")
	assert.NotNil(t, tlsConfig.VerifyPeerCertificate, "Should have custom verification function")
	assert.Equal(t, uint16(tls.VersionTLS12), tlsConfig.MinVersion)
}

func TestParseThumbprints(t *testing.T) {
	// Test single thumbprint
	result := parseThumbprints("abc123def456")
	assert.Equal(t, []string{"abc123def456"}, result)

	// Test multiple thumbprints
	result = parseThumbprints("abc123,def456,ghi789")
	assert.Equal(t, []string{"abc123", "def456", "ghi789"}, result)

	// Test with whitespace
	result = parseThumbprints("abc123 , def456 , ghi789")
	assert.Equal(t, []string{"abc123", "def456", "ghi789"}, result)

	// Test empty string
	result = parseThumbprints("")
	assert.Nil(t, result)

	// Test with empty parts
	result = parseThumbprints("abc123,,def456")
	assert.Equal(t, []string{"abc123", "def456"}, result)

	// Test with only whitespace
	result = parseThumbprints("   ,  ,  ")
	assert.Empty(t, result)
}
