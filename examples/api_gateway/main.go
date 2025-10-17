package main

import (
	"context"
	"fmt"
	"log"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	// Uncomment when using the commented examples below:
	// "github.com/conductor-sdk/conductor-go/sdk/model/gateway"
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

	// Example 2: Get a specific service
	if len(services) > 0 {
		serviceId := services[0].Id
		fmt.Printf("\n=== Getting Service Details for '%s' ===\n", serviceId)
		service, _, err := gatewayClient.GetService(ctx, serviceId)
		if err != nil {
			log.Printf("Error getting service: %v\n", err)
		} else {
			fmt.Printf("Service Name: %s\n", service.Name)
			fmt.Printf("Service Path: %s\n", service.Path)
			fmt.Printf("Auth Config ID: %s\n", service.AuthConfigId)
			fmt.Printf("Number of Routes: %d\n", len(service.Routes))
		}
	}

	// Example 3: List all authentication configs
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

	// Example 4: Get routes for a service
	if len(services) > 0 {
		serviceId := services[0].Id
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
			}
		}
	}

	// Example 5: Create a new auth config (commented out to avoid side effects)
	// Uncomment to test auth config creation
	/*
		fmt.Println("\n=== Creating New Auth Config ===")
		newAuthConfig := gateway.ApiGatewayAuthConfig{
			Id:            "example-auth-config",
			AuthType:      gateway.AuthTypeApiKey,
			ApplicationId: "example-application",
			ApiKeys:       []string{"example-key-1", "example-key-2"},
		}
		_, err = gatewayClient.CreateAuthConfig(ctx, newAuthConfig)
		if err != nil {
			log.Printf("Error creating auth config: %v\n", err)
		} else {
			fmt.Println("Auth config created successfully!")
		}
	*/

	// Example 6: Create a new service (commented out to avoid side effects)
	// Uncomment to test service creation
	/*
		fmt.Println("\n=== Creating New Service ===")
		newService := gateway.ApiGatewayService{
			Id:          "example-service",
			Name:        "Example Service",
			Path:        "/example",
			Description: "Example API Gateway Service",
			Enabled:     true,
			AuthConfigId: "api-key-config",
			CorsConfig: &gateway.CorsConfig{
				AccessControlAllowOrigin:  []string{"*"},
				AccessControlAllowMethods: []string{"GET", "POST"},
				AccessControlAllowHeaders: []string{"*"},
			},
		}
		_, err = gatewayClient.CreateService(ctx, newService)
		if err != nil {
			log.Printf("Error creating service: %v\n", err)
		} else {
			fmt.Println("Service created successfully!")
		}
	*/

	// Example 7: Create a new route (commented out to avoid side effects)
	// Uncomment to test route creation
	/*
		fmt.Println("\n=== Creating New Route ===")
		newRoute := gateway.ApiGatewayRoute{
			HttpMethod:  "GET",
			Path:        "/test",
			Description: "Example route",
			ServiceId:   "example-service",
			MappedWorkflow: &gateway.MappedWorkflow{
				Name:                   "hello_world",
				Version:                1,
				RequestMetadataAsInput: true,
			},
			WorkflowExecutionMode: "SYNCHRONOUS",
		}
		_, err = gatewayClient.CreateRoute(ctx, "example-service", newRoute)
		if err != nil {
			log.Printf("Error creating route: %v\n", err)
		} else {
			fmt.Println("Route created successfully!")
		}
	*/
}
