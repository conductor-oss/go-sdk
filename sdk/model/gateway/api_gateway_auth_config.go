//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package gateway

type AuthType string

const (
	// AuthTypeNone no authentication, or Public
	AuthTypeNone AuthType = "NONE"
	// AuthTypeApiKey API key based authentication
	AuthTypeApiKey AuthType = "API_KEY"
)

// ApiGatewayAuthConfig represents authentication configuration for API Gateway
type ApiGatewayAuthConfig struct {
	Id string `json:"id,omitempty"`

	AuthType AuthType `json:"authenticationType,omitempty"`

	ApplicationId string `json:"applicationId,omitempty"`

	ApiKeys []string `json:"apiKeys,omitempty"`
}
