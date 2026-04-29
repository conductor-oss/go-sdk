//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package client

import (
	"net/http"
	"strconv"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/metrics"
)

// metricsRoundTripper wraps an http.RoundTripper to record every HTTP request
// into the canonical http_api_client_request_seconds metric. When the legacy
// collector is active the call is a noop.
type metricsRoundTripper struct {
	next http.RoundTripper
}

// NewMetricsRoundTripper wraps an existing http.RoundTripper. If next is nil,
// http.DefaultTransport is used.
func NewMetricsRoundTripper(next http.RoundTripper) http.RoundTripper {
	if next == nil {
		next = http.DefaultTransport
	}
	return &metricsRoundTripper{next: next}
}

func (m *metricsRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := m.next.RoundTrip(req)
	elapsed := time.Since(start).Seconds()

	method := req.Method
	uri := ""
	if req.URL != nil {
		uri = req.URL.Path
	}

	// Network/transport failures (no HTTP response) are labeled "0",
	// matching the cross-SDK convention.
	status := "0"
	if err == nil && resp != nil {
		status = strconv.Itoa(resp.StatusCode)
	}

	metrics.RecordHTTPRequestTime(method, uri, status, elapsed)
	return resp, err
}
