//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package gateway

// ApiGatewayService represents an API Gateway service configuration
type ApiGatewayService struct {
	// Unique identifier for the service
	Id string `json:"id,omitempty"`
	// Name of the service
	Name string `json:"name,omitempty"`
	// Base path for the service
	Path string `json:"path,omitempty"`
	// Description of the service
	Description string `json:"description,omitempty"`
	// Whether the service is enabled
	Enabled bool `json:"enabled,omitempty"`
	// Creation timestamp
	CreateTime int64 `json:"createTime,omitempty"`
	// Update timestamp
	UpdateTime int64 `json:"updateTime,omitempty"`
	// Tags associated with the service
	Tags []interface{} `json:"tags,omitempty"`
	// Routes defined for this service
	Routes []ApiGatewayRoute `json:"routes,omitempty"`
	// Authentication config ID
	AuthConfigId string `json:"authConfigId,omitempty"`
	// CORS configuration
	CorsConfig *CorsConfig `json:"corsConfig,omitempty"`
	// Whether MCP is enabled
	McpEnabled bool `json:"mcpEnabled,omitempty"`
}

// CorsConfig represents CORS configuration for a service
type CorsConfig struct {
	// Allowed origins
	AccessControlAllowOrigin []string `json:"accessControlAllowOrigin,omitempty"`
	// Allowed HTTP methods
	AccessControlAllowMethods []string `json:"accessControlAllowMethods,omitempty"`
	// Allowed headers
	AccessControlAllowHeaders []string `json:"accessControlAllowHeaders,omitempty"`
}
