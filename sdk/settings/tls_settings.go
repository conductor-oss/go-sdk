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
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/log"
)

// TLSSettings configures TLS/SSL settings for the Conductor client.
// These settings control how the client verifies server certificates and
// handles client authentication.
type TLSSettings struct {
	// InsecureSkipVerify controls whether the client verifies the server's
	// certificate chain and hostname.
	// WARNING: Setting this to true is insecure. Recommended to use this for testing/development.
	InsecureSkipVerify bool

	// AllowSelfSigned allows the client to trust self-signed server certificates.
	// When enabled, the client will accept self-signed certificates without requiring
	// them to be added to the trust store.
	// Consider using PinnedThumbprints for additional security.
	AllowSelfSigned bool

	// PinnedThumbprints contains SHA-256 thumbprints of certificates to pin.
	// When AllowSelfSigned is true and this list is not empty, only certificates
	// matching one of these thumbprints will be accepted.
	// Format: lowercase hex string (e.g., "abc123def456...")
	PinnedThumbprints []string

	// RootCAs defines the set of root certificate authorities that the client
	// uses when verifying server certificates.
	// If nil, the system's default root CA set will be used.
	RootCAs *x509.CertPool

	// Certificates contains the client certificates for mutual TLS authentication.
	// Most servers don't require client certificates.
	Certificates []tls.Certificate

	// ServerName is used to verify the hostname on the returned certificates.
	// If empty, the hostname from the server URL will be used.
	// This is used when connecting via IP address or when the server certificate
	// uses a different hostname.
	ServerName string
}

// NewTLSDefaultSettings creates default TLS settings with secure defaults.
// By default:
// - Certificate verification is enabled
// - System's default root CAs are used
// - Hostname verification is enabled
func NewTLSDefaultSettings() *TLSSettings {
	return &TLSSettings{
		InsecureSkipVerify: false,
		RootCAs:            nil, // Use system defaults
		Certificates:       nil,
		ServerName:         "",
	}
}

// NewTLSSettingsFromEnv creates TLS settings from environment variables.
func NewTLSSettingsFromEnv() *TLSSettings {
	settings := NewTLSDefaultSettings()

	// Check for insecure mode
	if os.Getenv(EnvTLSInsecureSkipVerify) == "true" {
		settings.InsecureSkipVerify = true
		log.Warn("TLS certificate verification is disabled via environment variable. This is insecure!")
	}

	// Check for self-signed mode
	if os.Getenv(EnvTLSAllowSelfSigned) == "true" {
		settings.AllowSelfSigned = true
	}

	// Load pinned thumbprints if provided
	pinnedThumbprintsStr := os.Getenv(EnvTLSPinnedThumbprints)
	if pinnedThumbprintsStr != "" {
		thumbprints := parseThumbprints(pinnedThumbprintsStr)
		if len(thumbprints) > 0 {
			settings.PinnedThumbprints = thumbprints
			log.With("count", len(thumbprints)).Info("Loaded pinned certificate thumbprints from environment variable")
		}
	}

	// Load custom CA certificate
	caCertPath := os.Getenv(EnvTLSCACert)
	if caCertPath != "" {
		if err := settings.loadCACertFromFile(caCertPath); err != nil {
			// Log warning but don't fail - allow fallback to system certs
			log.With("path", caCertPath, "error", err).Warn("Failed to load CA certificate from environment variable")
		} else {
			log.With("path", caCertPath).Info("Loaded custom CA certificate")
		}
	}

	// Load client certificate for mutual TLS
	certPath := os.Getenv(EnvTLSClientCert)
	keyPath := os.Getenv(EnvTLSClientKey)
	if certPath != "" && keyPath != "" {
		if err := settings.loadClientCertFromFiles(certPath, keyPath); err != nil {
			log.With("certPath", certPath, "keyPath", keyPath, "error", err).Warn("Failed to load client certificate from environment variables")
		} else {
			log.Info("Loaded client certificate for mutual TLS")
		}
	}

	return settings
}

// loadCACertFromPEM loads a CA certificate from PEM-encoded bytes.
func (s *TLSSettings) loadCACertFromPEM(pemCert []byte) error {
	pool := s.RootCAs
	if pool == nil {
		pool = x509.NewCertPool()
	}

	if !pool.AppendCertsFromPEM(pemCert) {
		return fmt.Errorf("failed to parse CA certificate (invalid PEM format)")
	}

	s.RootCAs = pool
	return nil
}

// loadCACertFromFile loads a CA certificate from file, creating an EMPTY pool
// (system certificates are NOT included).
func (s *TLSSettings) loadCACertFromFile(path string) error {
	//nolint:gosec // G304: File path is user-provided configuration, not user input from untrusted source
	pemCert, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read CA cert file: %w", err)
	}

	return s.loadCACertFromPEM(pemCert)
}

// loadCACertFromFileWithSystemCerts loads a CA certificate from file and
// adds it to the system certificate pool.
func (s *TLSSettings) loadCACertFromFileWithSystemCerts(path string) error {
	// Load system cert pool first
	systemPool, err := x509.SystemCertPool()
	if err != nil {
		log.With("error", err).Warn("Failed to load system certificate pool, using empty pool")
		systemPool = x509.NewCertPool()
	}
	s.RootCAs = systemPool

	return s.loadCACertFromFile(path)
}

// loadClientCertFromPEM loads a client certificate and private key from PEM bytes
// for mutual TLS authentication.
func (s *TLSSettings) loadClientCertFromPEM(certPEM, keyPEM []byte) error {
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return fmt.Errorf("failed to load client certificate from PEM: %w", err)
	}

	s.Certificates = append(s.Certificates, cert)
	return nil
}

// loadClientCertFromFiles loads a client certificate and private key for
// mutual TLS authentication.
func (s *TLSSettings) loadClientCertFromFiles(certPath, keyPath string) error {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return fmt.Errorf("failed to load client certificate: %w", err)
	}

	s.Certificates = append(s.Certificates, cert)
	return nil
}

// BuildTLSConfig creates a *tls.Config from the TLSSettings.
// Returns nil if no custom TLS configuration is needed.
func (s *TLSSettings) BuildTLSConfig() *tls.Config {
	if s == nil {
		return nil // Use default TLS config
	}

	// Check if we have any custom settings
	hasCustomSettings := s.InsecureSkipVerify ||
		s.AllowSelfSigned ||
		s.RootCAs != nil ||
		len(s.Certificates) > 0 ||
		s.ServerName != ""

	if !hasCustomSettings {
		return nil // Use default TLS config
	}

	tlsConfig := &tls.Config{
		//nolint:gosec // G402: InsecureSkipVerify is intentionally configurable for development/testing, documented with security warnings
		InsecureSkipVerify: s.InsecureSkipVerify,
		RootCAs:            s.RootCAs,
		Certificates:       s.Certificates,
		ServerName:         s.ServerName,
		MinVersion:         tls.VersionTLS12, // Enforce minimum TLS 1.2
	}

	// If AllowSelfSigned is enabled, add custom verification logic
	if s.AllowSelfSigned {
		tlsConfig.VerifyPeerCertificate = s.verifySelfSignedCertificate
		// We need to skip the default verification to use our custom one
		tlsConfig.InsecureSkipVerify = true
	}

	return tlsConfig
}

// verifySelfSignedCertificate is a custom certificate verification function
// that allows self-signed certificates with optional thumbprint pinning.
func (s *TLSSettings) verifySelfSignedCertificate(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
	if len(rawCerts) == 0 {
		return fmt.Errorf("no certificates presented by server")
	}

	// Parse the server certificate
	cert, err := x509.ParseCertificate(rawCerts[0])
	if err != nil {
		return fmt.Errorf("failed to parse server certificate: %w", err)
	}

	// Check if certificate is self-signed
	isSelfSigned := cert.Issuer.String() == cert.Subject.String()

	if !isSelfSigned {
		// For non-self-signed certificates, use standard verification
		opts := x509.VerifyOptions{
			Roots:         s.RootCAs,
			Intermediates: x509.NewCertPool(),
			DNSName:       s.ServerName, // Check hostname for non-self-signed certificates
		}

		// Add intermediate certificates if present
		for i := 1; i < len(rawCerts); i++ {
			intermediateCert, err := x509.ParseCertificate(rawCerts[i])
			if err != nil {
				continue
			}
			opts.Intermediates.AddCert(intermediateCert)
		}

		if _, err := cert.Verify(opts); err != nil {
			return fmt.Errorf("certificate verification failed: %w", err)
		}
		return nil
	}

	// Self-signed certificate - perform basic validation
	if err := s.validateSelfSignedCert(cert); err != nil {
		return err
	}

	// Check if pinning is required
	if len(s.PinnedThumbprints) > 0 {
		// Calculate SHA-256 thumbprint of the certificate
		thumbprint := calculateThumbprint(rawCerts[0])

		// Normalize thumbprint to lowercase for comparison
		thumbprint = strings.ToLower(thumbprint)

		// Check if thumbprint matches any of the pinned thumbprints
		for _, pinnedThumbprint := range s.PinnedThumbprints {
			// Normalize pinned thumbprint to lowercase for comparison
			if thumbprint == strings.ToLower(pinnedThumbprint) {
				// Match found - certificate is accepted
				return nil
			}
		}

		// No match found - reject the certificate
		return fmt.Errorf(
			"self-signed certificate pin mismatch for server '%s': observed thumbprint '%s' does not match any pinned thumbprints %v",
			cert.Subject.CommonName,
			thumbprint,
			s.PinnedThumbprints,
		)
	}

	// No pinning required - accept the self-signed certificate
	return nil
}

// validateSelfSignedCert performs basic validation on a self-signed certificate.
// Checks: validity period, KeyUsage, and hostname (if ServerName is set).
func (s *TLSSettings) validateSelfSignedCert(cert *x509.Certificate) error {
	// Check validity period
	now := time.Now()
	if now.Before(cert.NotBefore) {
		return fmt.Errorf("self-signed certificate is not yet valid (starts at %v)", cert.NotBefore)
	}
	if now.After(cert.NotAfter) {
		return fmt.Errorf("self-signed certificate has expired (expired at %v)", cert.NotAfter)
	}

	// Check KeyUsage - must support server authentication
	if cert.ExtKeyUsage != nil {
		hasServerAuth := false
		for _, usage := range cert.ExtKeyUsage {
			if usage == x509.ExtKeyUsageServerAuth {
				hasServerAuth = true
				break
			}
		}
		if !hasServerAuth {
			return fmt.Errorf("self-signed certificate does not have ExtKeyUsageServerAuth")
		}
	}

	// Check hostname if ServerName is configured
	if s.ServerName != "" {
		if err := s.verifyHostname(cert, s.ServerName); err != nil {
			return fmt.Errorf("self-signed certificate hostname mismatch: %w", err)
		}
	}

	return nil
}

// verifyHostname checks if the certificate is valid for the given hostname.
// Checks both SANs (preferred) and CommonName (legacy fallback).
func (s *TLSSettings) verifyHostname(cert *x509.Certificate, hostname string) error {
	// First try standard verification (works with SANs)
	if err := cert.VerifyHostname(hostname); err == nil {
		return nil
	}

	// Fallback to CommonName check for legacy certificates
	if cert.Subject.CommonName == hostname {
		return nil
	}

	// Also check DNSNames and IPAddresses manually as additional fallback
	for _, dnsName := range cert.DNSNames {
		if dnsName == hostname {
			return nil
		}
	}

	return fmt.Errorf("hostname '%s' does not match certificate (CN: %s, SANs: %v)", hostname, cert.Subject.CommonName, cert.DNSNames)
}

// calculateThumbprint calculates the SHA-256 thumbprint of a certificate.
func calculateThumbprint(certDER []byte) string {
	hash := sha256.Sum256(certDER)
	return hex.EncodeToString(hash[:])
}

// parseThumbprints parses a comma-separated list of thumbprints from a string.
// Trims whitespace from each thumbprint, converts to lowercase, and filters out empty strings.
func parseThumbprints(thumbprintsStr string) []string {
	if thumbprintsStr == "" {
		return nil
	}

	parts := strings.Split(thumbprintsStr, ",")
	thumbprints := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			// Normalize to lowercase for consistent comparison
			thumbprints = append(thumbprints, strings.ToLower(trimmed))
		}
	}

	return thumbprints
}
