//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package transport

import (
	"net/http"
	"strings"

	"github.com/conductor-sdk/conductor-go/sdk/authentication"
	"github.com/conductor-sdk/conductor-go/sdk/log"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

// AuthTransport adds authentication headers to requests
type AuthTransport struct {
	TokenManager  authentication.TokenManager
	HttpSettings  *settings.HttpSettings
	HttpClient    *http.Client // Client for token refresh only
	BaseTransport http.RoundTripper
}

// RoundTrip implements http.RoundTripper
func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Skip auth for token endpoints to avoid recursion
	if strings.Contains(req.URL.Path, "/token") || strings.Contains(req.URL.Path, "/login") {
		return t.getBaseTransport().RoundTrip(req)
	}

	// Get token if we have token manager
	if t.TokenManager != nil {
		token, err := t.TokenManager.RefreshToken(t.HttpSettings, t.HttpClient)
		if err != nil {
			log.Info("Failed to refresh token", "error", err)
		}

		if err == nil && token != "" {
			// Clone request to avoid modifying original
			newReq := req.Clone(req.Context())
			newReq.Header.Set("X-Authorization", token)
			req = newReq
		} else {
			log.Warn("Not adding X-Authorization header",
				"has_error", err != nil,
				"token_empty", token == "",
			)
		}
	}

	resp, err := t.getBaseTransport().RoundTrip(req)
	return resp, err
}

func (t *AuthTransport) getBaseTransport() http.RoundTripper {
	if t.BaseTransport != nil {
		return t.BaseTransport
	}
	return http.DefaultTransport
}
