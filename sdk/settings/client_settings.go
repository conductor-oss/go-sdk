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
	"time"
)

// TokenExpirationInterface defines the interface for token expiration settings
type TokenExpirationInterface interface {
	GetDefaultExpiration() time.Duration
	GetCleanupInterval() time.Duration
}

// TokenManagerInterface defines the interface for token management
type TokenManagerInterface interface {
	RefreshToken(httpSettings *HttpSettings, httpClient *http.Client) (string, error)
}

// ClientSettings aggregates ALL settings for APIClient.
type ClientSettings struct {
	Authentication  *AuthenticationSettings
	HTTP            *HttpSettings
	Proxy           *ProxySettings
	TLS             *TLSSettings
	TokenExpiration TokenExpirationInterface
	TokenManager    TokenManagerInterface
}

// Option configures ClientSettings
type Option func(*ClientSettings)

// NewClientSettings creates ClientSettings with defaults and options
func NewClientSettings(opts ...Option) *ClientSettings {
	s := &ClientSettings{
		Authentication: &AuthenticationSettings{},
		HTTP:           NewHttpDefaultSettings(),
		Proxy:          newProxySettings(),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// NewClientSettingsFromEnv creates ClientSettings from environment variables
// and applies optional overrides
func NewClientSettingsFromEnv(opts ...Option) *ClientSettings {
	s := &ClientSettings{
		Authentication: NewAuthenticationSettingsFromEnv(),
		HTTP:           NewHttpSettingsFromEnv(),
		Proxy:          NewProxySettingsFromEnv(),
		TLS:            NewTLSSettingsFromEnv(),
	}

	// Apply options (override environment values)
	for _, opt := range opts {
		opt(s)
	}

	return s
}

// GetAuthentication returns authentication settings
func (s *ClientSettings) GetAuthentication() *AuthenticationSettings {
	return s.Authentication
}

// GetHTTP returns HTTP settings
func (s *ClientSettings) GetHTTP() *HttpSettings {
	return s.HTTP
}

// GetProxy returns proxy settings
func (s *ClientSettings) GetProxy() *ProxySettings {
	return s.Proxy
}

// GetTokenExpiration returns token expiration settings
func (s *ClientSettings) GetTokenExpiration() TokenExpirationInterface {
	return s.TokenExpiration
}

// GetTokenManager returns token manager
func (s *ClientSettings) GetTokenManager() TokenManagerInterface {
	return s.TokenManager
}

// GetTLS returns TLS settings
func (s *ClientSettings) GetTLS() *TLSSettings {
	return s.TLS
}

// ApplyOptions applies the given options to the ClientSettings
func (s *ClientSettings) ApplyOptions(opts ...Option) {
	for _, opt := range opts {
		opt(s)
	}
}
