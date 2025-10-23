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
	"log"
	"net/url"
	"time"
)

// ============= Proxy Options =============

// WithProxyURL sets proxy URL from string
// returns fatal error if URL is invalid
func WithProxyURL(urlStr string) Option {
	return func(s *ClientSettings) {
		proxyURL, err := url.Parse(urlStr)
		if err != nil {
			log.Fatalf("WithProxyURL: invalid URL %q: %v", urlStr, err)
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
