//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package gateway

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

type WorkflowExecutionMode string

const (
	// ExecutionModeSynchronous waits for workflow completion and returns the result
	ExecutionModeSynchronous WorkflowExecutionMode = "SYNCHRONOUS"
	// ExecutionModeAsynchronous starts workflow and immediately returns workflow ID
	ExecutionModeAsynchronous WorkflowExecutionMode = "ASYNCHRONOUS"
)

// ServiceRouteDetails wraps a route with its metrics
type ServiceRouteDetails struct {
	// The route configuration
	Route *ApiGatewayRoute `json:"route,omitempty"`
	// Service route metrics
	ServiceRouteMetrics map[string]interface{} `json:"serviceRouteMetrics,omitempty"`
}

// ApiGatewayRoute represents a route configuration within an API Gateway service
type ApiGatewayRoute struct {
	// HTTP method for the route (GET, POST, PUT, DELETE, etc.)
	HttpMethod string `json:"httpMethod,omitempty"`
	// Path for the route
	Path string `json:"path,omitempty"`
	// Description of the route
	Description string `json:"description,omitempty"`
	// Service ID this route belongs to
	ServiceId string `json:"serviceId,omitempty"`
	// Mapped workflow configuration
	MappedWorkflow *MappedWorkflow `json:"mappedWorkflow,omitempty"`
	// Whether to include workflow metadata in output
	WorkflowMetadataInOutput bool `json:"workflowMetadataInOutput,omitempty"`
	// Workflow execution mode
	WorkflowExecutionMode WorkflowExecutionMode `json:"workflowExecutionMode,omitempty"`
	// Wait until tasks (comma-separated task reference names)
	WaitUntilTasks string `json:"waitUntilTasks,omitempty"`
	// Creation timestamp
	CreateTime int64 `json:"createTime,omitempty"`
	// Update timestamp
	UpdateTime int64 `json:"updateTime,omitempty"`
	// Tags
	Tags []model.Tag `json:"tags,omitempty"`
}

// MappedWorkflow represents the workflow mapping configuration for a route
type MappedWorkflow struct {
	// Workflow name
	Name string `json:"name,omitempty"`
	// Workflow version
	Version int32 `json:"version,omitempty"`
	// Whether to pass request metadata as input
	RequestMetadataAsInput bool `json:"requestMetadataAsInput,omitempty"`
}
