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
	"net/http"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// ServiceRegistryResourceApiService is the service for the service registry resource.
type ServiceRegistryResourceApiService struct {
	*APIClient
}

// AddOrUpdateService adds or updates a service.
func (a *ServiceRegistryResourceApiService) AddOrUpdateService(ctx context.Context, serviceRegistry model.ServiceRegistry) (*http.Response, error) {
	genRequest := toGeneratedServiceRegistry(&serviceRegistry)
	resp, err := a.http_orkes.ServiceRegistryResourceAPI.AddOrUpdateService(ctx).
		ServiceRegistry(*genRequest).
		Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	return resp, nil
}

// AddOrUpdateMethod adds or updates a method.
func (a *ServiceRegistryResourceApiService) AddOrUpdateMethod(ctx context.Context, body model.ServiceMethod, registryName string) (*http.Response, error) {
	genRequest := toGeneratedServiceMethod(&body)
	resp, err := a.http_orkes.ServiceRegistryResourceAPI.AddOrUpdateMethod(ctx, registryName).
		ServiceMethod(genRequest).
		Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// GetService gets a service by name and version.
func (a *ServiceRegistryResourceApiService) GetService(ctx context.Context, serviceName string) (model.ServiceRegistry, *http.Response, error) {
	genResult, resp, err := a.http_orkes.ServiceRegistryResourceAPI.GetService(ctx, serviceName).Execute()
	if err != nil {
		return model.ServiceRegistry{}, resp, wrapGeneratedError(err, resp)
	}

	result := toDomainServiceRegistryFromGenerated(genResult)
	return result, resp, nil
}

// GetServices gets all services.
func (a *ServiceRegistryResourceApiService) GetServices(ctx context.Context) ([]model.ServiceRegistry, *http.Response, error) {
	genResult, resp, err := a.http_orkes.ServiceRegistryResourceAPI.GetRegisteredServices(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	result := toDomainServiceRegistriesFromGenerated(genResult)
	return result, resp, nil
}

func (a *ServiceRegistryResourceApiService) CloseCircuitBreaker(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error) {
	result, resp, err := a.http_orkes.ServiceRegistryResourceAPI.CloseCircuitBreaker(ctx, name).Execute()
	if err != nil {
		return model.CircuitBreakerTransitionResponse{}, resp, wrapGeneratedError(err, resp)
	}
	return toDomainCircuitBreakerTransitionResponseFromGenerated(result), resp, nil
}

// DeleteProto deletes a proto by registry name and filename.
func (a *ServiceRegistryResourceApiService) DeleteProto(ctx context.Context, registryName string, filename string) (*http.Response, error) {
	resp, err := a.http_orkes.ServiceRegistryResourceAPI.DeleteProto(ctx, registryName, filename).Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DiscoverOpts are the options for the Discover method
type ServiceRegistryResourceApiDiscoverOpts struct {
	Create optional.Bool
}

// Discover discovers a service by name.
func (a *ServiceRegistryResourceApiService) Discover(ctx context.Context, name string, optionals *ServiceRegistryResourceApiDiscoverOpts) ([]model.ServiceMethod, *http.Response, error) {
	req := a.http_orkes.ServiceRegistryResourceAPI.Discover(ctx, name)

	if optionals != nil && optionals.Create.IsSet() {
		req = req.Create(optionals.Create.Value())
	}

	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return toDomainServiceMethodsFromGenerated(result), resp, nil
}

// GetAllProtos gets all protos by registry name.
func (a *ServiceRegistryResourceApiService) GetAllProtos(ctx context.Context, registryName string) ([]model.ProtoRegistryEntry, *http.Response, error) {
	req := a.http_orkes.ServiceRegistryResourceAPI.GetAllProtos(ctx, registryName)

	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return toDomainProtoRegistryEntriesFromGenerated(result), resp, nil
}

// GetCircuitBreakerStatus gets the circuit breaker status by name.
func (a *ServiceRegistryResourceApiService) GetCircuitBreakerStatus(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error) {
	req := a.http_orkes.ServiceRegistryResourceAPI.GetCircuitBreakerStatus(ctx, name)

	result, resp, err := req.Execute()
	if err != nil {
		return model.CircuitBreakerTransitionResponse{}, resp, wrapGeneratedError(err, resp)
	}
	return toDomainCircuitBreakerTransitionResponseFromGenerated(result), resp, nil
}

// GetProtoData gets the proto data by registry name and filename.
func (a *ServiceRegistryResourceApiService) GetProtoData(ctx context.Context, registryName string, filename string) (string, *http.Response, error) {
	req := a.http_orkes.ServiceRegistryResourceAPI.GetProtoData(ctx, registryName, filename)

	result, resp, err := req.Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// GetRegisteredServices gets all services.
func (a *ServiceRegistryResourceApiService) GetRegisteredServices(ctx context.Context) ([]model.ServiceRegistry, *http.Response, error) {
	req := a.http_orkes.ServiceRegistryResourceAPI.GetRegisteredServices(ctx)

	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return toDomainServiceRegistriesFromGenerated(result), resp, nil
}

// OpenCircuitBreaker opens a circuit breaker by name.
func (a *ServiceRegistryResourceApiService) OpenCircuitBreaker(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error) {
	req := a.http_orkes.ServiceRegistryResourceAPI.OpenCircuitBreaker(ctx, name)

	result, resp, err := req.Execute()
	if err != nil {
		return model.CircuitBreakerTransitionResponse{}, resp, wrapGeneratedError(err, resp)
	}
	return toDomainCircuitBreakerTransitionResponseFromGenerated(result), resp, nil
}

// RemoveMethod removes a method by registry name, service name, method name and method type.
func (a *ServiceRegistryResourceApiService) RemoveMethod(ctx context.Context, registryName string, serviceName string, method string, methodType string) (*http.Response, error) {
	req := a.http_orkes.ServiceRegistryResourceAPI.RemoveMethod(ctx, registryName)

	req = req.ServiceName(serviceName).Method(method).MethodType(methodType)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// RemoveService removes a service by name.
func (a *ServiceRegistryResourceApiService) RemoveService(ctx context.Context, name string) (*http.Response, error) {
	req := a.http_orkes.ServiceRegistryResourceAPI.RemoveService(ctx, name)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// SetProtoData sets the proto data by registry name and filename.
func (a *ServiceRegistryResourceApiService) SetProtoData(ctx context.Context, body string, registryName string, filename string) (*http.Response, error) {
	req := a.http_orkes.ServiceRegistryResourceAPI.SetProtoData(ctx, registryName, filename)

	req = req.Body(body)

	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}
