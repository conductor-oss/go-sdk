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
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

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

	return tlsConfig
}
