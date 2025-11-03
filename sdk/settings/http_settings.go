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
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/log"
)

var defaultBaseUrl = "http://localhost:8080/api"

// HttpSettings configures HTTP settings for the client.
type HttpSettings struct {
	BaseUrl string
	Headers map[string]string
	Timeout time.Duration
}

// NewHttpDefaultSettings creates default HTTP settings.
func NewHttpDefaultSettings() *HttpSettings {
	return NewHttpSettings(
		defaultBaseUrl,
	)
}

// NewHttpSettings creates HTTP settings with the given base URL.
func NewHttpSettings(baseUrl string) *HttpSettings {
	return &HttpSettings{
		BaseUrl: baseUrl,
		Headers: map[string]string{
			"Content-Type":    "application/json",
			"Accept":          "application/json",
			"Accept-Encoding": "gzip",
		},
		Timeout: 30 * time.Second,
	}
}

// NewHttpSettingsFromEnv creates HTTP settings from environment variables.
func NewHttpSettingsFromEnv() *HttpSettings {
	serverURL := os.Getenv(EnvServerURL)
	if serverURL == "" {
		serverURL = defaultBaseUrl
	}

	httpSettings := NewHttpSettings(serverURL)

	// Load timeout from environment
	if timeoutStr := os.Getenv(EnvTimeout); timeoutStr != "" {
		if timeoutInt, err := strconv.Atoi(timeoutStr); err == nil {
			httpSettings.Timeout = time.Duration(timeoutInt) * time.Second
		}
	}

	return httpSettings
}

// normalizeBaseURL makes BaseUrl end with exactly one "/api" (no trailing slash).
// No errors returned per contract: on irrecoverable input we log and return unchanged.
// Special case: if BaseUrl is empty -> default to "http://localhost:8080/api".
func (h *HttpSettings) normalizeBaseURL() {
	raw := strings.TrimSpace(h.BaseUrl)

	// Default for empty
	if raw == "" {
		h.BaseUrl = defaultBaseUrl
		return
	}

	u, err := url.Parse(raw)
	if err != nil {
		log.Warn("normalizeBaseURL: cannot parse base URL", "url", raw, "error", err)
		return
	}
	if u.Scheme == "" || u.Host == "" {
		log.Error("normalizeBaseURL: base URL must include scheme and host", "url", raw)
		return
	}

	// If someone passed query/fragment on base — drop them.
	if u.RawQuery != "" || u.Fragment != "" {
		log.Warn("normalizeBaseURL: dropping query/fragment from base URL", "url", raw)
		u.RawQuery = ""
		u.Fragment = ""
	}

	// Clean and normalize path to ensure exactly one "/api" suffix.
	p := path.Clean(u.Path)
	if p == "." {
		p = "/"
	}
	p = strings.TrimRight(p, "/")

	switch {
	case p == "" || p == "/":
		p = "/api"
	case strings.HasSuffix(p, "/api"):
		// already ends with /api -> keep as-is
	default:
		p = p + "/api"
	}

	u.Path = p
	u.RawPath = ""

	h.BaseUrl = u.String()
}
