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
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model/gateway"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

// TestApiGatewayClientCreation tests that the API Gateway client can be created
func TestApiGatewayClientCreation(t *testing.T) {
	authSettings := settings.NewAuthenticationSettings("key", "secret")
	httpSettings := settings.NewHttpSettings("http://localhost:8080/api")
	apiClient := NewAPIClient(authSettings, httpSettings)

	gatewayClient := NewApiGatewayClient(apiClient)

	if gatewayClient == nil {
		t.Fatal("Expected gateway client to be created, got nil")
	}
}

// TestServiceModelCreation tests that the service model can be created
func TestServiceModelCreation(t *testing.T) {
	service := gateway.ApiGatewayService{
		Id:          "test-service",
		Name:        "Test Service",
		Path:        "/test",
		Description: "Test Description",
		Enabled:     true,
		CorsConfig: &gateway.CorsConfig{
			AccessControlAllowOrigin:  []string{"*"},
			AccessControlAllowMethods: []string{"GET", "POST"},
			AccessControlAllowHeaders: []string{"*"},
		},
	}

	if service.Id != "test-service" {
		t.Errorf("Expected service ID to be 'test-service', got '%s'", service.Id)
	}

	if service.Name != "Test Service" {
		t.Errorf("Expected service name to be 'Test Service', got '%s'", service.Name)
	}

	if !service.Enabled {
		t.Error("Expected service to be enabled")
	}

	if service.CorsConfig == nil {
		t.Fatal("Expected CORS config to be set")
	}

	if len(service.CorsConfig.AccessControlAllowMethods) != 2 {
		t.Errorf("Expected 2 allowed methods, got %d", len(service.CorsConfig.AccessControlAllowMethods))
	}
}

// TestRouteModelCreation tests that the route model can be created
func TestRouteModelCreation(t *testing.T) {
	route := gateway.ApiGatewayRoute{
		HttpMethod:  "GET",
		Path:        "/users",
		Description: "Get users",
		ServiceId:   "test-service",
		MappedWorkflow: &gateway.MappedWorkflow{
			Name:                   "get_users",
			Version:                1,
			RequestMetadataAsInput: true,
		},
		WorkflowExecutionMode: "SYNCHRONOUS",
	}

	if route.HttpMethod != "GET" {
		t.Errorf("Expected HTTP method to be 'GET', got '%s'", route.HttpMethod)
	}

	if route.Path != "/users" {
		t.Errorf("Expected path to be '/users', got '%s'", route.Path)
	}

	if route.MappedWorkflow == nil {
		t.Fatal("Expected mapped workflow to be set")
	}

	if route.MappedWorkflow.Name != "get_users" {
		t.Errorf("Expected workflow name to be 'get_users', got '%s'", route.MappedWorkflow.Name)
	}

	if route.WorkflowExecutionMode != "SYNCHRONOUS" {
		t.Errorf("Expected workflow execution mode to be 'SYNCHRONOUS', got '%s'", route.WorkflowExecutionMode)
	}
}

// TestAuthConfigModelCreation tests that the auth config model can be created
func TestAuthConfigModelCreation(t *testing.T) {
	authConfig := gateway.ApiGatewayAuthConfig{
		Id:            "api-key-config",
		AuthType:      gateway.AuthTypeApiKey,
		ApplicationId: "test-application",
		ApiKeys:       []string{"test-key-1", "test-key-2"},
	}

	if authConfig.Id != "api-key-config" {
		t.Errorf("Expected auth config ID to be 'api-key-config', got '%s'", authConfig.Id)
	}

	if authConfig.AuthType != gateway.AuthTypeApiKey {
		t.Errorf("Expected auth type to be 'API_KEY', got '%s'", authConfig.AuthType)
	}

	if authConfig.ApplicationId != "test-application" {
		t.Errorf("Expected application ID to be 'test-application', got '%s'", authConfig.ApplicationId)
	}

	if authConfig.ApiKeys == nil {
		t.Fatal("Expected API keys to be set")
	}

	if len(authConfig.ApiKeys) != 2 {
		t.Errorf("Expected 2 API keys, got %d", len(authConfig.ApiKeys))
	}

	if authConfig.ApiKeys[0] != "test-key-1" {
		t.Errorf("Expected first API key to be 'test-key-1', got '%s'", authConfig.ApiKeys[0])
	}
}

// TestAuthConfigModelCreationNoAuth tests that the auth config model can be created with no authentication
func TestAuthConfigModelCreationNoAuth(t *testing.T) {
	authConfig := gateway.ApiGatewayAuthConfig{
		Id:            "no-auth-config",
		AuthType:      gateway.AuthTypeNone,
		ApplicationId: "test-application",
	}

	if authConfig.Id != "no-auth-config" {
		t.Errorf("Expected auth config ID to be 'no-auth-config', got '%s'", authConfig.Id)
	}

	if authConfig.AuthType != gateway.AuthTypeNone {
		t.Errorf("Expected auth type to be 'NONE', got '%s'", authConfig.AuthType)
	}

	if authConfig.ApplicationId != "test-application" {
		t.Errorf("Expected application ID to be 'test-application', got '%s'", authConfig.ApplicationId)
	}
}
