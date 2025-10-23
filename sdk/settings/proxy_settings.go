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
	"fmt"
	"log"

	"net/http"
	"net/url"
	"os"
)

// ProxySettings configures HTTP proxy for the client
type ProxySettings struct {
	URL      *url.URL
	Username string
	Password string
}

// newProxySettings creates default proxy settings (no proxy)
func newProxySettings() *ProxySettings {
	return &ProxySettings{}
}

// NewProxySettingsFromEnv creates ProxySettings from environment variables
func NewProxySettingsFromEnv() *ProxySettings {
	s := newProxySettings()
	if err := loadProxyFromEnv(s); err != nil {
		log.Fatalf("failed to load proxy from environment: %v", err)
	}
	return s
}

// loadProxyFromEnv loads proxy configuration from environment into ProxySettings
func loadProxyFromEnv(s *ProxySettings) error {
	if proxyStr := os.Getenv(EnvProxy); proxyStr != "" {
		proxyURL, err := url.Parse(proxyStr)
		if err != nil {
			return fmt.Errorf("invalid %s: %w", EnvProxy, err)
		}

		s.URL = proxyURL

		// Extract embedded credentials (e.g., http://user:pass@proxy.com:8080)
		if proxyURL.User != nil {
			s.Username = proxyURL.User.Username()
			s.Password, _ = proxyURL.User.Password()
			// Remove credentials from URL to prevent exposure in logs
			proxyURL.User = nil
		}
	}
	return nil
}

// BuildProxyFunc creates an HTTP proxy function from ProxySettings
func (s *ProxySettings) BuildProxyFunc() func(*http.Request) (*url.URL, error) {
	if s == nil || s.URL == nil {
		return http.ProxyFromEnvironment
	}

	return s.createProxyFunc()
}

// createProxyFunc creates the actual proxy function with bypass and credentials
func (s *ProxySettings) createProxyFunc() func(*http.Request) (*url.URL, error) {
	return func(_ *http.Request) (*url.URL, error) {
		// Clone URL to avoid modifying original
		proxyURL := *s.URL

		// Add credentials if configured
		if s.Username != "" {
			if s.Password != "" {
				proxyURL.User = url.UserPassword(s.Username, s.Password)
			} else {
				proxyURL.User = url.User(s.Username)
			}
		}

		return &proxyURL, nil
	}
}
