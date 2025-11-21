//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package settings

// Environment variable constants
const (
	// Authentication key for the client.
	EnvAuthKey = "CONDUCTOR_AUTH_KEY"
	// Authentication secret for the client.
	EnvAuthSecret = "CONDUCTOR_AUTH_SECRET" //nolint:gosec

	// Server URL for the client.
	EnvServerURL = "CONDUCTOR_SERVER_URL"
	// Timeout for the client HTTP requests.
	EnvTimeout = "CONDUCTOR_CLIENT_HTTP_TIMEOUT"

	// Proxy URL for the client HTTP requests.
	EnvProxy = "CONDUCTOR_PROXY"

	// EnvTLSInsecureSkipVerify disables SSL certificate verification (INSECURE!)
	// Set to "true" to disable verification.
	EnvTLSInsecureSkipVerify = "CONDUCTOR_TLS_INSECURE_SKIP_VERIFY"

	// EnvTLSAllowSelfSigned enables acceptance of self-signed certificates.
	// Set to "true" to allow self-signed certificates.
	// Consider using CONDUCTOR_TLS_PINNED_THUMBPRINTS for additional security.
	EnvTLSAllowSelfSigned = "CONDUCTOR_TLS_ALLOW_SELF_SIGNED"

	// EnvTLSPinnedThumbprints specifies SHA-256 thumbprints for certificate pinning.
	// Comma-separated list of thumbprints (e.g., "abc123...,def456...").
	// Only used when CONDUCTOR_TLS_ALLOW_SELF_SIGNED is enabled.
	EnvTLSPinnedThumbprints = "CONDUCTOR_TLS_PINNED_THUMBPRINTS"

	// EnvTLSCACert specifies the path to a CA certificate file in PEM format.
	EnvTLSCACert = "CONDUCTOR_TLS_CA_CERT"

	// EnvTLSClientCert specifies the path to a client certificate file for mutual TLS.
	// Must be used together with CONDUCTOR_TLS_CLIENT_KEY.
	EnvTLSClientCert = "CONDUCTOR_TLS_CLIENT_CERT"

	// EnvTLSClientKey specifies the path to a client private key file for mutual TLS.
	// Must be used together with CONDUCTOR_TLS_CLIENT_CERT.
	EnvTLSClientKey = "CONDUCTOR_TLS_CLIENT_KEY"
)
