package integration_tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/gateway"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestApiGatewayClient(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Setup
	gatewayClient := testdata.GatewayClient
	ctx := context.Background()

	// Generate unique identifiers for testing
	serviceId := fmt.Sprintf("test-service-%d", time.Now().UnixNano())
	authConfigId := fmt.Sprintf("test-auth-%d", time.Now().UnixNano())

	// Test case 1: Create a new authentication config
	authConfig := gateway.ApiGatewayAuthConfig{
		Id:            authConfigId,
		AuthType:      gateway.AuthTypeApiKey,
		ApplicationId: "test-application",
		ApiKeys:       []string{"test-key-1", "test-key-2"},
	}

	resp, err := gatewayClient.CreateAuthConfig(ctx, authConfig)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Add cleanup to delete the auth config at the end
	defer func() {
		_, _ = gatewayClient.DeleteAuthConfig(ctx, authConfigId)
	}()

	// Test case 2: Get the auth config by ID
	retrievedAuthConfig, resp, err := gatewayClient.GetAuthConfig(ctx, authConfigId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, authConfigId, retrievedAuthConfig.Id)
	assert.Equal(t, gateway.AuthTypeApiKey, retrievedAuthConfig.AuthType)
	assert.Equal(t, "test-application", retrievedAuthConfig.ApplicationId)

	// Test case 3: Get all auth configs and verify our config exists
	allAuthConfigs, resp, err := gatewayClient.GetAllAuthConfigs(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var foundAuthConfig bool
	for _, auth := range allAuthConfigs {
		if auth.Id == authConfigId {
			foundAuthConfig = true
			assert.Equal(t, gateway.AuthTypeApiKey, auth.AuthType)
			break
		}
	}
	assert.True(t, foundAuthConfig, "Created auth config not found in list of all auth configs")

	// Test case 4: Create a new service
	service := gateway.ApiGatewayService{
		Id:           serviceId,
		Name:         fmt.Sprintf("Test Service %d", time.Now().UnixNano()),
		Path:         fmt.Sprintf("/test-%d", time.Now().UnixNano()),
		Description:  "Test service for integration testing",
		Enabled:      true,
		AuthConfigId: authConfigId,
		CorsConfig: &gateway.CorsConfig{
			AccessControlAllowOrigin:  []string{"*"},
			AccessControlAllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AccessControlAllowHeaders: []string{"*"},
		},
	}

	resp, err = gatewayClient.CreateService(ctx, service)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Add cleanup to delete the service at the end
	defer func() {
		_, _ = gatewayClient.DeleteService(ctx, serviceId)
	}()

	// Test case 5: Get the service by ID
	retrievedService, resp, err := gatewayClient.GetService(ctx, serviceId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, serviceId, retrievedService.Id)
	assert.Equal(t, service.Name, retrievedService.Name)
	assert.Equal(t, service.Path, retrievedService.Path)
	assert.True(t, retrievedService.Enabled)

	// Test case 6: Get all services and verify our service exists
	allServices, resp, err := gatewayClient.GetAllServices(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var foundService bool
	for _, svc := range allServices {
		if svc.Id == serviceId {
			foundService = true
			assert.Equal(t, service.Name, svc.Name)
			break
		}
	}
	assert.True(t, foundService, "Created service not found in list of all services")

	// Test case 7: Create a route for the service
	route := gateway.ApiGatewayRoute{
		HttpMethod:  "GET",
		Path:        "/test-route",
		Description: "Test route for integration testing",
		ServiceId:   serviceId,
		MappedWorkflow: &gateway.MappedWorkflow{
			Name:                   "test_workflow",
			Version:                1,
			RequestMetadataAsInput: true,
		},
		WorkflowExecutionMode:    gateway.ExecutionModeSynchronous,
		WorkflowMetadataInOutput: true,
	}

	resp, err = gatewayClient.CreateRoute(ctx, serviceId, route)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 8: Get routes for the service
	routes, resp, err := gatewayClient.GetRoutes(ctx, serviceId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, len(routes), 1)

	var foundRoute bool
	for _, r := range routes {
		if r.HttpMethod == "GET" && r.Path == "/test-route" {
			foundRoute = true
			assert.Equal(t, "test_workflow", r.MappedWorkflow.Name)
			assert.Equal(t, int32(1), r.MappedWorkflow.Version)
			assert.True(t, r.MappedWorkflow.RequestMetadataAsInput)
			break
		}
	}
	assert.True(t, foundRoute, "Created route not found in list of routes")

	// Test case 9: Update the route with tags
	updatedRoute := route
	updatedRoute.Tags = []model.Tag{
		{
			Key:   "environment",
			Value: "test",
			Type_: "metadata",
		},
		{
			Key:   "owner",
			Value: "integration-test",
			Type_: "ownership",
		},
	}
	updatedRoute.Description = "Updated test route"

	resp, err = gatewayClient.UpdateRoute(ctx, serviceId, "/test-route", updatedRoute)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 10: Verify route update
	updatedRoutes, resp, err := gatewayClient.GetRoutes(ctx, serviceId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var foundUpdatedRoute bool
	for _, r := range updatedRoutes {
		if r.HttpMethod == "GET" && r.Path == "/test-route" {
			foundUpdatedRoute = true
			assert.Equal(t, "Updated test route", r.Description)
			assert.GreaterOrEqual(t, len(r.Tags), 2)
			break
		}
	}
	assert.True(t, foundUpdatedRoute, "Updated route not found")

	// Test case 11: Get tags for service (should be empty initially)
	_, resp, err = gatewayClient.GetTagsForService(ctx, serviceId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 12: Update the service
	updatedService := retrievedService
	updatedService.Description = "Updated test service"
	updatedService.Enabled = false

	resp, err = gatewayClient.UpdateService(ctx, serviceId, updatedService)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 13: Verify service update
	verifyService, resp, err := gatewayClient.GetService(ctx, serviceId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Updated test service", verifyService.Description)
	assert.False(t, verifyService.Enabled)

	// Test case 14: Update the auth config
	updatedAuthConfig := retrievedAuthConfig
	updatedAuthConfig.ApiKeys = []string{"test-key-1", "test-key-2", "test-key-3"}

	resp, err = gatewayClient.UpdateAuthConfig(ctx, authConfigId, updatedAuthConfig)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 15: Verify auth config update
	verifyAuthConfig, resp, err := gatewayClient.GetAuthConfig(ctx, authConfigId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 3, len(verifyAuthConfig.ApiKeys))

	// Test case 16: Put tags for route using PutTagsForRoute
	routeTags := []model.Tag{
		{
			Key:   "department",
			Value: "engineering",
			Type_: "metadata",
		},
		{
			Key:   "team",
			Value: "platform",
			Type_: "ownership",
		},
		{
			Key:   "version",
			Value: "v1",
			Type_: "metadata",
		},
	}

	resp, err = gatewayClient.PutTagsForRoute(ctx, serviceId, "GET", "/test-route", routeTags)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 17: Verify tags were set on the route
	routesWithTags, resp, err := gatewayClient.GetRoutes(ctx, serviceId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var foundRouteWithTags bool
	for _, r := range routesWithTags {
		if r.HttpMethod == "GET" && r.Path == "/test-route" {
			foundRouteWithTags = true
			assert.GreaterOrEqual(t, len(r.Tags), 3, "Route should have at least 3 tags")

			// Verify specific tags exist
			tagMap := make(map[string]string)
			for _, tag := range r.Tags {
				tagMap[tag.Key] = tag.Value
			}
			assert.Equal(t, "engineering", tagMap["department"])
			assert.Equal(t, "platform", tagMap["team"])
			assert.Equal(t, "v1", tagMap["version"])
			break
		}
	}
	assert.True(t, foundRouteWithTags, "Route with tags not found")

	// Test case 18: Delete specific tags from service
	tagsToDelete := []model.Tag{
		{
			Key:   "department",
			Value: "engineering",
			Type_: "metadata",
		},
	}

	// First, verify we can get tags for the service again
	serviceTagsBeforeDelete, resp, err := gatewayClient.GetTagsForService(ctx, serviceId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Delete tags from service (if any exist)
	if len(serviceTagsBeforeDelete) > 0 {
		resp, err = gatewayClient.DeleteTagsForService(ctx, serviceId, tagsToDelete)
		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		// Verify tags were deleted
		serviceTagsAfterDelete, resp, err := gatewayClient.GetTagsForService(ctx, serviceId)
		assert.Nil(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.LessOrEqual(t, len(serviceTagsAfterDelete), len(serviceTagsBeforeDelete))
	}

	// Test case 19: Delete the route
	resp, err = gatewayClient.DeleteRoute(ctx, serviceId, "GET", "/test-route")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 20: Verify the route was deleted
	routesAfterDelete, resp, err := gatewayClient.GetRoutes(ctx, serviceId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	routeFoundAfterDelete := false
	for _, r := range routesAfterDelete {
		if r.HttpMethod == "GET" && r.Path == "/test-route" {
			routeFoundAfterDelete = true
			break
		}
	}
	assert.False(t, routeFoundAfterDelete, "Route should have been deleted")

	// Test case 21: Delete the service
	resp, err = gatewayClient.DeleteService(ctx, serviceId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 22: Verify the service was deleted
	allServicesAfterDelete, resp, err := gatewayClient.GetAllServices(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	serviceFoundAfterDelete := false
	for _, svc := range allServicesAfterDelete {
		if svc.Id == serviceId {
			serviceFoundAfterDelete = true
			break
		}
	}
	assert.False(t, serviceFoundAfterDelete, "Service should have been deleted")

	// Test case 23: Delete the auth config
	resp, err = gatewayClient.DeleteAuthConfig(ctx, authConfigId)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test case 24: Verify the auth config was deleted
	allAuthConfigsAfterDelete, resp, err := gatewayClient.GetAllAuthConfigs(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	authConfigFoundAfterDelete := false
	for _, auth := range allAuthConfigsAfterDelete {
		if auth.Id == authConfigId {
			authConfigFoundAfterDelete = true
			break
		}
	}
	assert.False(t, authConfigFoundAfterDelete, "Auth config should have been deleted")
}
