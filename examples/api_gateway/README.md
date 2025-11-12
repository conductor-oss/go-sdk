# API Gateway Management Example

This example demonstrates how to manage API Gateway services, routes, and authentication configurations.

## Overview

The API Gateway functionality allows you to:
- Create, read, update, and delete API Gateway services
- Manage authentication configurations
- Define and manage routes that map HTTP endpoints to Conductor workflows

# Running the Example

## Prerequisites

- A running Conductor server with API Gateway enabled
- Authentication credentials (Key/Secret)

##  Setup and Run

```bash
export CONDUCTOR_SERVER_URL="http://localhost:8080/api"
export CONDUCTOR_AUTH_KEY="your-key"
export CONDUCTOR_AUTH_SECRET="your-secret"
```

```bash
go run main.go
```

## Usage

### Initialize the Client

```go
apiClient := client.NewAPIClientFromEnv()

// Create the API Gateway client
gatewayClient := client.NewApiGatewayClient(apiClient)
```

### Service Management

#### List All Services

```go
services, _, err := gatewayClient.GetAllServices(ctx)
```

#### Get a Specific Service

```go
service, _, err := gatewayClient.GetService(ctx, "service-id")
```

#### Create a Service

```go
newService := gateway.ApiGatewayService{
    Id:          "my-service",
    Name:        "My Service",
    Path:        "/my-service/v1",
    Description: "My API Gateway Service",
    Enabled:     true,
    AuthConfigId: "api-key-config",
    CorsConfig: &gateway.CorsConfig{
        AccessControlAllowOrigin:  []string{"*"},
        AccessControlAllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AccessControlAllowHeaders: []string{"*"},
    },
}
_, err := gatewayClient.CreateService(ctx, newService)
```

#### Update a Service

```go
_, err := gatewayClient.UpdateService(ctx, "service-id", updatedService)
```

#### Delete a Service

```go
_, err := gatewayClient.DeleteService(ctx, "service-id")
```

### Authentication Config Management

#### List All Auth Configs

```go
authConfigs, _, err := gatewayClient.GetAllAuthConfigs(ctx)
```

#### Get a Specific Auth Config

```go
authConfig, _, err := gatewayClient.GetAuthConfig(ctx, "auth-config-id")
```

#### Create an Auth Config

**API Key Authentication:**
```go
apiKeyAuthConfig := gateway.ApiGatewayAuthConfig{
    Id:            "api-key-config",
    AuthType:      gateway.AuthTypeApiKey,
    ApplicationId: "my-application",
    ApiKeys:       []string{"your-api-key-here"},
}
_, err := gatewayClient.CreateAuthConfig(ctx, apiKeyAuthConfig)
```

**No Authentication:**
```go
noAuthConfig := gateway.ApiGatewayAuthConfig{
    Id:            "no-auth-config",
    AuthType:      gateway.AuthTypeNone,
    ApplicationId: "my-application",
}
_, err := gatewayClient.CreateAuthConfig(ctx, noAuthConfig)
```

#### Update an Auth Config

```go
_, err := gatewayClient.UpdateAuthConfig(ctx, "auth-config-id", updatedAuthConfig)
```

#### Delete an Auth Config

```go
_, err := gatewayClient.DeleteAuthConfig(ctx, "auth-config-id")
```

### Route Management

#### List Routes for a Service

```go
routes, _, err := gatewayClient.GetRoutes(ctx, "service-id")
```

#### Create a Route

```go
newRoute := gateway.ApiGatewayRoute{
    HttpMethod:  "GET",
    Path:        "/users",
    Description: "Get users",
    ServiceId:   "my-service",
    MappedWorkflow: &gateway.MappedWorkflow{
        Name:                   "get_users_workflow",
        Version:                1,
        RequestMetadataAsInput: true,
    },
    WorkflowExecutionMode: "SYNCHRONOUS",
}
_, err := gatewayClient.CreateRoute(ctx, "my-service", newRoute)
```

#### Update a Route

```go
_, err := gatewayClient.UpdateRoute(ctx, "service-id", "/users", updatedRoute)
```

#### Delete a Route

```go
_, err := gatewayClient.DeleteRoute(ctx, "service-id", "GET", "/users")
```

## API Gateway Concepts

### Service
A service represents a collection of routes that share common configuration like base path, authentication, and CORS settings.

### Route
A route maps an HTTP method and path to a Conductor workflow. Routes belong to a service.

### Authentication Config
An authentication configuration defines how API Gateway authenticates incoming requests.

### Workflow Execution Modes
- `SYNCHRONOUS`: Wait for workflow completion and return the result
- `ASYNCHRONOUS`: Start workflow and immediately return workflow ID

## Architecture

The API Gateway implementation follows the Conductor SDK patterns:

```
sdk/
├── client/
│   ├── api_gateway_resource.go      # API Gateway resource service
│   └── gateway_client.go            # Client interface and factory
└── model/
    └── gateway/
        ├── api_gateway_service.go    # Service model
        ├── api_gateway_route.go      # Route model
        └── api_gateway_auth_config.go # Auth config model
```
