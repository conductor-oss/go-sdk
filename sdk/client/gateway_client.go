package client

import (
	"context"
	"net/http"

	"github.com/conductor-sdk/conductor-go/sdk/model/gateway"
)

// ApiGatewayClient provides methods to manage API Gateway services, routes, and authentication configs
type ApiGatewayClient interface {
	// Service operations
	CreateService(ctx context.Context, service gateway.ApiGatewayService) (*http.Response, error)
	GetService(ctx context.Context, serviceId string) (gateway.ApiGatewayService, *http.Response, error)
	GetAllServices(ctx context.Context) ([]gateway.ApiGatewayService, *http.Response, error)
	UpdateService(ctx context.Context, serviceId string, service gateway.ApiGatewayService) (*http.Response, error)
	DeleteService(ctx context.Context, serviceId string) (*http.Response, error)

	// Authentication config operations
	CreateAuthConfig(ctx context.Context, authConfig gateway.ApiGatewayAuthConfig) (*http.Response, error)
	GetAuthConfig(ctx context.Context, authConfigId string) (gateway.ApiGatewayAuthConfig, *http.Response, error)
	GetAllAuthConfigs(ctx context.Context) ([]gateway.ApiGatewayAuthConfig, *http.Response, error)
	UpdateAuthConfig(ctx context.Context, authConfigId string, authConfig gateway.ApiGatewayAuthConfig) (*http.Response, error)
	DeleteAuthConfig(ctx context.Context, authConfigId string) (*http.Response, error)

	// Route operations
	CreateRoute(ctx context.Context, serviceId string, route gateway.ApiGatewayRoute) (*http.Response, error)
	GetRoutes(ctx context.Context, serviceId string) ([]gateway.ApiGatewayRoute, *http.Response, error)
	UpdateRoute(ctx context.Context, serviceId string, routePath string, route gateway.ApiGatewayRoute) (*http.Response, error)
	DeleteRoute(ctx context.Context, serviceId string, httpMethod string, routePath string) (*http.Response, error)
}

// NewApiGatewayClient creates a new API Gateway client
func NewApiGatewayClient(apiClient *APIClient) ApiGatewayClient {
	return &ApiGatewayResourceApiService{apiClient}
}
