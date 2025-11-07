package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/gateway"
)

func main() {
	apiClient := client.NewAPIClientFromEnv()
	gatewayClient := client.NewApiGatewayClient(apiClient)

	ctx := context.Background()

	// Example 1: List all services
	fmt.Println("=== Listing all API Gateway Services ===")
	services, _, err := gatewayClient.GetAllServices(ctx)
	if err != nil {
		log.Printf("Error getting services: %v\n", err)
	} else {
		fmt.Printf("Found %d services:\n", len(services))
		for _, svc := range services {
			fmt.Printf("  - %s (ID: %s, Path: %s, Enabled: %v)\n",
				svc.Name, svc.Id, svc.Path, svc.Enabled)
		}
	}

	// Example 2: List all authentication configs
	fmt.Println("\n=== Listing all Authentication Configs ===")
	authConfigs, _, err := gatewayClient.GetAllAuthConfigs(ctx)
	if err != nil {
		log.Printf("Error getting auth configs: %v\n", err)
	} else {
		fmt.Printf("Found %d auth configs:\n", len(authConfigs))
		for _, auth := range authConfigs {
			fmt.Printf("  - Application: %s (ID: %s, Type: %s)\n",
				auth.ApplicationId, auth.Id, auth.AuthType)
		}
	}

	// Example 3: Create a new auth config
	fmt.Println("\n=== Creating New Auth Config ===")
	authConfigId := fmt.Sprintf("example-auth-config-%d", time.Now().UnixNano())
	newAuthConfig := gateway.ApiGatewayAuthConfig{
		Id:            authConfigId,
		AuthType:      gateway.AuthTypeApiKey,
		ApplicationId: "example-application",
		ApiKeys:       []string{"example-key-1", "example-key-2"},
	}
	_, err = gatewayClient.CreateAuthConfig(ctx, newAuthConfig)
	if err != nil {
		log.Fatalf("Error creating auth config: %v\n", err)
	}
	fmt.Println("Auth config created successfully!")

	// Setup cleanup for auth config
	defer func() {
		fmt.Printf("\n=== Cleaning up: Deleting Auth Config '%s' ===\n", authConfigId)
		_, err := gatewayClient.DeleteAuthConfig(ctx, authConfigId)
		if err != nil {
			log.Printf("Error deleting auth config: %v\n", err)
		} else {
			fmt.Println("Auth config deleted successfully!")
		}
	}()

	// Example 4: Get the created auth config
	fmt.Printf("\n=== Getting Auth Config '%s' ===\n", authConfigId)
	retrievedAuthConfig, _, err := gatewayClient.GetAuthConfig(ctx, authConfigId)
	if err != nil {
		log.Printf("Error getting auth config: %v\n", err)
	} else {
		fmt.Printf("Auth Config ID: %s\n", retrievedAuthConfig.Id)
		fmt.Printf("Auth Type: %s\n", retrievedAuthConfig.AuthType)
		fmt.Printf("Application ID: %s\n", retrievedAuthConfig.ApplicationId)
		fmt.Printf("API Keys: %v\n", retrievedAuthConfig.ApiKeys)
	}

	// Example 5: Create a new service
	fmt.Println("\n=== Creating New Service ===")
	serviceId := fmt.Sprintf("example-service-%d", time.Now().UnixNano())
	newService := gateway.ApiGatewayService{
		Id:           serviceId,
		Name:         "Example Service",
		Path:         fmt.Sprintf("/example-%d", time.Now().UnixNano()),
		Description:  "Example API Gateway Service",
		Enabled:      true,
		AuthConfigId: authConfigId,
		CorsConfig: &gateway.CorsConfig{
			AccessControlAllowOrigin:  []string{"*"},
			AccessControlAllowMethods: []string{"GET", "POST"},
			AccessControlAllowHeaders: []string{"*"},
		},
	}
	_, err = gatewayClient.CreateService(ctx, newService)
	if err != nil {
		log.Fatalf("Error creating service: %v\n", err)
	}
	fmt.Println("Service created successfully!")

	// Setup cleanup for service (will execute before auth config cleanup due to defer order)
	defer func() {
		fmt.Printf("\n=== Cleaning up: Deleting Service '%s' ===\n", serviceId)
		_, err := gatewayClient.DeleteService(ctx, serviceId)
		if err != nil {
			log.Printf("Error deleting service: %v\n", err)
		} else {
			fmt.Println("Service deleted successfully!")
		}
	}()

	// Example 6: Get the created service
	fmt.Printf("\n=== Getting Service '%s' ===\n", serviceId)
	service, _, err := gatewayClient.GetService(ctx, serviceId)
	if err != nil {
		log.Printf("Error getting service: %v\n", err)
	} else {
		fmt.Printf("Service Name: %s\n", service.Name)
		fmt.Printf("Service Path: %s\n", service.Path)
		fmt.Printf("Auth Config ID: %s\n", service.AuthConfigId)
		fmt.Printf("Enabled: %v\n", service.Enabled)
	}

	// Example 7: Create a new route
	fmt.Println("\n=== Creating New Route ===")
	newRoute := gateway.ApiGatewayRoute{
		HttpMethod:  "GET",
		Path:        "/test",
		Description: "Example route",
		ServiceId:   serviceId,
		MappedWorkflow: &gateway.MappedWorkflow{
			Name:                   "hello_world",
			Version:                1,
			RequestMetadataAsInput: true,
		},
		WorkflowExecutionMode:    gateway.ExecutionModeSynchronous,
		WorkflowMetadataInOutput: true,
		Tags: []model.Tag{
			{
				Key:   "environment",
				Value: "example",
				Type_: "metadata",
			},
		},
	}
	_, err = gatewayClient.CreateRoute(ctx, serviceId, newRoute)
	if err != nil {
		log.Fatalf("Error creating route: %v\n", err)
	}
	fmt.Println("Route created successfully!")

	// Setup cleanup for route (will execute before service cleanup due to defer order)
	defer func() {
		fmt.Printf("\n=== Cleaning up: Deleting Route 'GET /test' from Service '%s' ===\n", serviceId)
		_, err := gatewayClient.DeleteRoute(ctx, serviceId, "GET", "/test")
		if err != nil {
			log.Printf("Error deleting route: %v\n", err)
		} else {
			fmt.Println("Route deleted successfully!")
		}
	}()

	// Example 8: Get routes for the service
	fmt.Printf("\n=== Getting Routes for Service '%s' ===\n", serviceId)
	routes, _, err := gatewayClient.GetRoutes(ctx, serviceId)
	if err != nil {
		log.Printf("Error getting routes: %v\n", err)
	} else {
		fmt.Printf("Found %d routes:\n", len(routes))
		for _, route := range routes {
			fmt.Printf("  - %s %s -> Workflow: %s (v%d)\n",
				route.HttpMethod, route.Path,
				route.MappedWorkflow.Name, route.MappedWorkflow.Version)
			if len(route.Tags) > 0 {
				fmt.Printf("    Tags: ")
				for i, tag := range route.Tags {
					if i > 0 {
						fmt.Printf(", ")
					}
					fmt.Printf("%s=%s", tag.Key, tag.Value)
				}
				fmt.Println()
			}
		}
	}

	// Example 9: Update the route
	fmt.Println("\n=== Updating Route ===")
	updatedRoute := newRoute
	updatedRoute.Description = "Updated example route"
	updatedRoute.Tags = append(updatedRoute.Tags, model.Tag{
		Key:   "updated",
		Value: "true",
		Type_: "metadata",
	})
	_, err = gatewayClient.UpdateRoute(ctx, serviceId, "/test", updatedRoute)
	if err != nil {
		log.Printf("Error updating route: %v\n", err)
	} else {
		fmt.Println("Route updated successfully!")
	}

	// Example 10: Update the service
	fmt.Println("\n=== Updating Service ===")
	updatedService := service
	updatedService.Description = "Updated example service"
	updatedService.Enabled = false
	_, err = gatewayClient.UpdateService(ctx, serviceId, updatedService)
	if err != nil {
		log.Printf("Error updating service: %v\n", err)
	} else {
		fmt.Println("Service updated successfully!")
	}

	// Example 11: Verify the updates
	fmt.Printf("\n=== Verifying Updates ===\n")
	verifyService, _, err := gatewayClient.GetService(ctx, serviceId)
	if err != nil {
		log.Printf("Error getting updated service: %v\n", err)
	} else {
		fmt.Printf("Service Description: %s\n", verifyService.Description)
		fmt.Printf("Service Enabled: %v\n", verifyService.Enabled)
	}

	fmt.Println("\n=== Example Complete ===")
	fmt.Println("Cleanup will now execute (in reverse order: route -> service -> auth config)...")
}
