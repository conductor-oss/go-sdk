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
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/gateway"
)

// ApiGatewayResourceApiService provides methods to manage API Gateway services, routes, and authentication configs
type ApiGatewayResourceApiService struct {
	*APIClient
}

// ==================== Service Operations ====================

// CreateService creates a new API Gateway service
func (a *ApiGatewayResourceApiService) CreateService(ctx context.Context, service gateway.ApiGatewayService) (*http.Response, error) {
	path := "/gateway/config/services"
	return a.Post(ctx, path, service, nil)
}

// GetService retrieves an API Gateway service by ID
func (a *ApiGatewayResourceApiService) GetService(ctx context.Context, serviceId string) (gateway.ApiGatewayService, *http.Response, error) {
	var result gateway.ApiGatewayService
	path := fmt.Sprintf("/gateway/config/services/%s", serviceId)
	resp, err := a.Get(ctx, path, nil, &result)
	return result, resp, err
}

// GetAllServices retrieves all API Gateway services
func (a *ApiGatewayResourceApiService) GetAllServices(ctx context.Context) ([]gateway.ApiGatewayService, *http.Response, error) {
	var result []gateway.ApiGatewayService
	path := "/gateway/config/services"
	resp, err := a.Get(ctx, path, nil, &result)
	return result, resp, err
}

// UpdateService updates an existing API Gateway service
func (a *ApiGatewayResourceApiService) UpdateService(ctx context.Context, serviceId string, service gateway.ApiGatewayService) (*http.Response, error) {
	path := fmt.Sprintf("/gateway/config/services/%s", serviceId)
	return a.Put(ctx, path, service, nil)
}

// DeleteService deletes an API Gateway service by ID
func (a *ApiGatewayResourceApiService) DeleteService(ctx context.Context, serviceId string) (*http.Response, error) {
	path := fmt.Sprintf("/gateway/config/services/%s", serviceId)
	return a.Delete(ctx, path, nil, nil)
}

// GetTagsForService retrieves tags for a service
func (a *ApiGatewayResourceApiService) GetTagsForService(ctx context.Context, serviceId string) ([]model.Tag, *http.Response, error) {
	var result []model.Tag
	path := fmt.Sprintf("/gateway/config/services/%s/tags", serviceId)
	resp, err := a.Get(ctx, path, nil, &result)
	return result, resp, err
}

// DeleteTagsForService removes tags from a service
func (a *ApiGatewayResourceApiService) DeleteTagsForService(ctx context.Context, serviceId string, tags []model.Tag) (*http.Response, error) {
	path := fmt.Sprintf("/gateway/config/services/%s/tags", serviceId)
	return a.DeleteWithBody(ctx, path, tags, nil)
}

// ==================== Authentication Config Operations ====================

// CreateAuthConfig creates a new authentication configuration
func (a *ApiGatewayResourceApiService) CreateAuthConfig(ctx context.Context, authConfig gateway.ApiGatewayAuthConfig) (*http.Response, error) {
	path := "/gateway/config/auth"
	return a.Post(ctx, path, authConfig, nil)
}

// GetAuthConfig retrieves an authentication configuration by ID
func (a *ApiGatewayResourceApiService) GetAuthConfig(ctx context.Context, authConfigId string) (gateway.ApiGatewayAuthConfig, *http.Response, error) {
	var result gateway.ApiGatewayAuthConfig
	path := fmt.Sprintf("/gateway/config/auth/%s", authConfigId)
	resp, err := a.Get(ctx, path, nil, &result)
	return result, resp, err
}

// GetAllAuthConfigs retrieves all authentication configurations
func (a *ApiGatewayResourceApiService) GetAllAuthConfigs(ctx context.Context) ([]gateway.ApiGatewayAuthConfig, *http.Response, error) {
	var result []gateway.ApiGatewayAuthConfig
	path := "/gateway/config/auth"
	resp, err := a.Get(ctx, path, nil, &result)
	return result, resp, err
}

// UpdateAuthConfig updates an existing authentication configuration
func (a *ApiGatewayResourceApiService) UpdateAuthConfig(ctx context.Context, authConfigId string, authConfig gateway.ApiGatewayAuthConfig) (*http.Response, error) {
	path := fmt.Sprintf("/gateway/config/auth/%s", authConfigId)
	return a.Put(ctx, path, authConfig, nil)
}

// DeleteAuthConfig deletes an authentication configuration by ID
func (a *ApiGatewayResourceApiService) DeleteAuthConfig(ctx context.Context, authConfigId string) (*http.Response, error) {
	path := fmt.Sprintf("/gateway/config/auth/%s", authConfigId)
	return a.Delete(ctx, path, nil, nil)
}

// ==================== Route Operations ====================

// CreateRoute creates a new route for a service
func (a *ApiGatewayResourceApiService) CreateRoute(ctx context.Context, serviceId string, route gateway.ApiGatewayRoute) (*http.Response, error) {
	path := fmt.Sprintf("/gateway/config/services/%s/routes", serviceId)
	return a.Post(ctx, path, route, nil)
}

// GetRoutes retrieves all routes for a service
func (a *ApiGatewayResourceApiService) GetRoutes(ctx context.Context, serviceId string) ([]gateway.ApiGatewayRoute, *http.Response, error) {
	var detailsList []gateway.ServiceRouteDetails
	path := fmt.Sprintf("/gateway/config/services/%s/routes", serviceId)
	resp, err := a.Get(ctx, path, nil, &detailsList)
	if err != nil {
		return nil, resp, err
	}

	// Extract routes from the details wrapper
	routes := make([]gateway.ApiGatewayRoute, 0, len(detailsList))
	for _, details := range detailsList {
		if details.Route != nil {
			routes = append(routes, *details.Route)
		}
	}

	return routes, resp, nil
}

// UpdateRoute updates an existing route
func (a *ApiGatewayResourceApiService) UpdateRoute(ctx context.Context, serviceId string, routePath string, route gateway.ApiGatewayRoute) (*http.Response, error) {
	path := fmt.Sprintf("/gateway/config/services/%s/routes", serviceId)
	queryParams := url.Values{}
	queryParams.Add("path", routePath)
	return a.PutWithParams(ctx, path, queryParams, route, nil)
}

// DeleteRoute deletes a route from a service
func (a *ApiGatewayResourceApiService) DeleteRoute(ctx context.Context, serviceId string, httpMethod string, routePath string) (*http.Response, error) {
	path := fmt.Sprintf("/gateway/config/services/%s/routes", serviceId)
	queryParams := url.Values{}
	queryParams.Add("method", httpMethod)
	queryParams.Add("path", routePath)
	return a.Delete(ctx, path, queryParams, nil)
}

// PutTagsForRoute replaces tags for a specific route
func (a *ApiGatewayResourceApiService) PutTagsForRoute(ctx context.Context, serviceId string, httpMethod string, routePath string, tags []model.Tag) (*http.Response, error) {
	path := fmt.Sprintf("/gateway/config/services/%s/routes/tags", serviceId)
	queryParams := url.Values{}
	queryParams.Add("method", httpMethod)
	queryParams.Add("path", routePath)
	return a.PutWithParams(ctx, path, queryParams, tags, nil)
}
