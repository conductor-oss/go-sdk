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
	"net/url"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/log"
)

// ============= Proxy Options =============

// WithProxyURL sets proxy URL from string.
func WithProxyURL(urlStr string) Option {
	return func(s *ClientSettings) {
		proxyURL, err := url.Parse(urlStr)
		if err != nil {
			log.Error("WithProxyURL: invalid URL", "error", err, "url", urlStr)
			return
		}
		if s.Proxy == nil {
			s.Proxy = newProxySettings()
		}
		s.Proxy.URL = proxyURL
	}
}

// WithProxy sets proxy from parsed *url.URL
func WithProxy(proxyURL *url.URL) Option {
	return func(s *ClientSettings) {
		if s.Proxy == nil {
			s.Proxy = newProxySettings()
		}
		s.Proxy.URL = proxyURL
	}
}

// WithProxyCredentials sets proxy authentication credentials
func WithProxyCredentials(username, password string) Option {
	return func(s *ClientSettings) {
		if s.Proxy == nil {
			s.Proxy = newProxySettings()
		}
		s.Proxy.Username = username
		s.Proxy.Password = password
	}
}

// WithProxySettings sets complete proxy settings
func WithProxySettings(proxySettings *ProxySettings) Option {
	return func(s *ClientSettings) {
		s.Proxy = proxySettings
	}
}

// ============= Auth Options =============

// WithAuthKey overrides authentication key
func WithAuthCredentials(key string, secret string) Option {
	return func(s *ClientSettings) {
		s.Authentication = NewAuthenticationSettings(key, secret)
	}
}

// ============= HTTP Options =============

// WithServerURL overrides server URL
func WithServerURL(serverURL string) Option {
	return func(s *ClientSettings) {
		if s.HTTP == nil {
			s.HTTP = NewHttpDefaultSettings()
		}
		s.HTTP.BaseUrl = serverURL
	}
}

// WithHTTPTimeout sets HTTP timeout
func WithHTTPTimeout(timeout time.Duration) Option {
	return func(s *ClientSettings) {
		if s.HTTP == nil {
			s.HTTP = NewHttpDefaultSettings()
		}
		if timeout > 0 {
			s.HTTP.Timeout = timeout
		}
	}
}

// WithHTTPHeaders sets HTTP headers
func WithHTTPHeaders(headers map[string]string) Option {
	return func(s *ClientSettings) {
		if s.HTTP == nil {
			s.HTTP = NewHttpDefaultSettings()
		}
		for key, value := range headers {
			s.HTTP.Headers[key] = value
		}
	}
}

// ============= Token Options =============

// WithTokenExpiration sets token expiration settings
func WithTokenExpiration(tokenExpiration TokenExpirationInterface) Option {
	return func(s *ClientSettings) {
		s.TokenExpiration = tokenExpiration
	}
}

// WithTokenManager sets token manager
func WithTokenManager(tokenManager TokenManagerInterface) Option {
	return func(s *ClientSettings) {
		s.TokenManager = tokenManager
	}
}

// ============= TLS Options =============

// WithTLSSettings sets complete TLS settings.
func WithTLSSettings(tlsSettings *TLSSettings) Option {
	return func(s *ClientSettings) {
		s.TLS = tlsSettings
	}
}

// WithInsecureSkipVerify disables SSL certificate verification (INSECURE!)
// It is recommended to use this option only for testing or development purposes.
func WithInsecureSkipVerify(skip bool) Option {
	return func(s *ClientSettings) {
		if s.TLS == nil {
			s.TLS = NewTLSDefaultSettings()
		}
		s.TLS.InsecureSkipVerify = skip

		if skip {
			log.Warn("TLS certificate verification disabled (InsecureSkipVerify=true).")
		}
	}
}

// WithCACertFromFile loads a CA certificate from a file to trust self-signed certificates.
// This is the recommended way to connect to servers with self-signed certificates.
// The file must be in PEM format.
func WithCACertFromFile(path string) Option {
	return func(s *ClientSettings) {
		if s.TLS == nil {
			s.TLS = NewTLSDefaultSettings()
		}
		if err := s.TLS.loadCACertFromFile(path); err != nil {
			log.Error("WithCACertFromFile: failed to load cert for file", "path", path, "error", err)
			return
		}
	}
}

// WithCACertFromPEM loads a CA certificate from PEM-encoded bytes.
// Use this when you have the certificate data in memory (e.g., from a secrets manager).
func WithCACertFromPEM(pemCert []byte) Option {
	return func(s *ClientSettings) {
		if s.TLS == nil {
			s.TLS = NewTLSDefaultSettings()
		}
		if err := s.TLS.loadCACertFromPEM(pemCert); err != nil {
			log.Error("WithCACertFromPEM: failed to load cert from PEM", "error", err)
			return
		}
	}
}

// WithClientCertFromPEM loads a client certificate from PEM-encoded bytes.
// Use this when you have the certificate data in memory (e.g., from a secrets manager).
// Both certificate and private key must be in PEM format.
func WithClientCertFromPEM(certPEM, keyPEM []byte) Option {
	return func(s *ClientSettings) {
		if s.TLS == nil {
			s.TLS = NewTLSDefaultSettings()
		}
		if err := s.TLS.loadClientCertFromPEM(certPEM, keyPEM); err != nil {
			log.Error("WithClientCertFromPEM: failed to load client cert from PEM", "error", err)
			return
		}
	}
}

// WithClientCert loads a client certificate for mutual TLS authentication.
// Both certificate and private key files must be in PEM format.
func WithClientCert(certPath, keyPath string) Option {
	return func(s *ClientSettings) {
		if s.TLS == nil {
			s.TLS = NewTLSDefaultSettings()
		}
		if err := s.TLS.loadClientCertFromFiles(certPath, keyPath); err != nil {
			log.Error("WithClientCert: failed to load client cert from files", "certPath", certPath, "keyPath", keyPath, "error", err)
			return
		}
	}
}

// WithTLSServerName sets the server name for TLS SNI and certificate verification.
// Use this when connecting via IP address or when the certificate uses a different hostname.
func WithTLSServerName(serverName string) Option {
	return func(s *ClientSettings) {
		if s.TLS == nil {
			s.TLS = NewTLSDefaultSettings()
		}
		s.TLS.ServerName = serverName
	}
}

// WithSelfSignedCert is a convenience function for the common case of trusting
// a self-signed certificate while keeping system certificates.
// Use this to trust both public CAs and your custom CA (hybrid mode).
func WithSelfSignedCert(certPath string) Option {
	return func(s *ClientSettings) {
		if s.TLS == nil {
			s.TLS = NewTLSDefaultSettings()
		}
		if err := s.TLS.loadCACertFromFileWithSystemCerts(certPath); err != nil {
			log.Error("WithSelfSignedCert: failed to load cert from file", "path", certPath, "error", err)
			return
		}
	}
}
