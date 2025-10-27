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
	"os"
	"strconv"
	"time"
)

// HttpSettings configures HTTP settings for the client.
type HttpSettings struct {
	BaseUrl string
	Headers map[string]string
	Timeout time.Duration
}

// NewHttpDefaultSettings creates default HTTP settings.
func NewHttpDefaultSettings() *HttpSettings {
	return NewHttpSettings(
		"http://localhost:8080/api",
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
		serverURL = "http://localhost:8080/api"
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
