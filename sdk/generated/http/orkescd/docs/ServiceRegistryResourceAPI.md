# \ServiceRegistryResourceAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrUpdateMethod**](ServiceRegistryResourceAPI.md#AddOrUpdateMethod) | **Post** /registry/service/{registryName}/methods | 
[**AddOrUpdateService**](ServiceRegistryResourceAPI.md#AddOrUpdateService) | **Post** /registry/service | 
[**CloseCircuitBreaker**](ServiceRegistryResourceAPI.md#CloseCircuitBreaker) | **Post** /registry/service/{name}/circuit-breaker/close | 
[**DeleteProto**](ServiceRegistryResourceAPI.md#DeleteProto) | **Delete** /registry/service/protos/{registryName}/{filename} | 
[**Discover**](ServiceRegistryResourceAPI.md#Discover) | **Get** /registry/service/{name}/discover | 
[**GetAllProtos**](ServiceRegistryResourceAPI.md#GetAllProtos) | **Get** /registry/service/protos/{registryName} | 
[**GetCircuitBreakerStatus**](ServiceRegistryResourceAPI.md#GetCircuitBreakerStatus) | **Get** /registry/service/{name}/circuit-breaker/status | 
[**GetProtoData**](ServiceRegistryResourceAPI.md#GetProtoData) | **Get** /registry/service/protos/{registryName}/{filename} | 
[**GetRegisteredServices**](ServiceRegistryResourceAPI.md#GetRegisteredServices) | **Get** /registry/service | 
[**GetService**](ServiceRegistryResourceAPI.md#GetService) | **Get** /registry/service/{name} | 
[**OpenCircuitBreaker**](ServiceRegistryResourceAPI.md#OpenCircuitBreaker) | **Post** /registry/service/{name}/circuit-breaker/open | 
[**RemoveMethod**](ServiceRegistryResourceAPI.md#RemoveMethod) | **Delete** /registry/service/{registryName}/methods | 
[**RemoveService**](ServiceRegistryResourceAPI.md#RemoveService) | **Delete** /registry/service/{name} | 
[**SetProtoData**](ServiceRegistryResourceAPI.md#SetProtoData) | **Post** /registry/service/protos/{registryName}/{filename} | 



## AddOrUpdateMethod

> AddOrUpdateMethod(ctx, registryName).ServiceMethod(serviceMethod).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {
    registryName := "registryName_example" // string | 
    serviceMethod := *openapiclient.NewServiceMethod() // ServiceMethod | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRegistryResourceAPI.AddOrUpdateMethod(context.Background(), registryName).ServiceMethod(serviceMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.AddOrUpdateMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrUpdateMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceMethod** | [**ServiceMethod**](ServiceMethod.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddOrUpdateService

> AddOrUpdateService(ctx).ServiceRegistry(serviceRegistry).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {
    serviceRegistry := *openapiclient.NewServiceRegistry() // ServiceRegistry | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRegistryResourceAPI.AddOrUpdateService(context.Background()).ServiceRegistry(serviceRegistry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.AddOrUpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddOrUpdateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceRegistry** | [**ServiceRegistry**](ServiceRegistry.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloseCircuitBreaker

> CircuitBreakerTransitionResponse CloseCircuitBreaker(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRegistryResourceAPI.CloseCircuitBreaker(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.CloseCircuitBreaker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseCircuitBreaker`: CircuitBreakerTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceRegistryResourceAPI.CloseCircuitBreaker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseCircuitBreakerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CircuitBreakerTransitionResponse**](CircuitBreakerTransitionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProto

> DeleteProto(ctx, registryName, filename).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {
    registryName := "registryName_example" // string | 
    filename := "filename_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRegistryResourceAPI.DeleteProto(context.Background(), registryName, filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.DeleteProto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryName** | **string** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProtoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Discover

> []ServiceMethod Discover(ctx, name).Create(create).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {
    name := "name_example" // string | 
    create := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRegistryResourceAPI.Discover(context.Background(), name).Create(create).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.Discover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Discover`: []ServiceMethod
    fmt.Fprintf(os.Stdout, "Response from `ServiceRegistryResourceAPI.Discover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **create** | **bool** |  | [default to false]

### Return type

[**[]ServiceMethod**](ServiceMethod.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllProtos

> []ProtoRegistryEntry GetAllProtos(ctx, registryName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {
    registryName := "registryName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRegistryResourceAPI.GetAllProtos(context.Background(), registryName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.GetAllProtos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllProtos`: []ProtoRegistryEntry
    fmt.Fprintf(os.Stdout, "Response from `ServiceRegistryResourceAPI.GetAllProtos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProtosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProtoRegistryEntry**](ProtoRegistryEntry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCircuitBreakerStatus

> CircuitBreakerTransitionResponse GetCircuitBreakerStatus(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRegistryResourceAPI.GetCircuitBreakerStatus(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.GetCircuitBreakerStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCircuitBreakerStatus`: CircuitBreakerTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceRegistryResourceAPI.GetCircuitBreakerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCircuitBreakerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CircuitBreakerTransitionResponse**](CircuitBreakerTransitionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtoData

> string GetProtoData(ctx, registryName, filename).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {
    registryName := "registryName_example" // string | 
    filename := "filename_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRegistryResourceAPI.GetProtoData(context.Background(), registryName, filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.GetProtoData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtoData`: string
    fmt.Fprintf(os.Stdout, "Response from `ServiceRegistryResourceAPI.GetProtoData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryName** | **string** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtoDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisteredServices

> []ServiceRegistry GetRegisteredServices(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRegistryResourceAPI.GetRegisteredServices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.GetRegisteredServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegisteredServices`: []ServiceRegistry
    fmt.Fprintf(os.Stdout, "Response from `ServiceRegistryResourceAPI.GetRegisteredServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredServicesRequest struct via the builder pattern


### Return type

[**[]ServiceRegistry**](ServiceRegistry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> ServiceRegistry GetService(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRegistryResourceAPI.GetService(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.GetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetService`: ServiceRegistry
    fmt.Fprintf(os.Stdout, "Response from `ServiceRegistryResourceAPI.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceRegistry**](ServiceRegistry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenCircuitBreaker

> CircuitBreakerTransitionResponse OpenCircuitBreaker(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRegistryResourceAPI.OpenCircuitBreaker(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.OpenCircuitBreaker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenCircuitBreaker`: CircuitBreakerTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceRegistryResourceAPI.OpenCircuitBreaker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenCircuitBreakerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CircuitBreakerTransitionResponse**](CircuitBreakerTransitionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMethod

> RemoveMethod(ctx, registryName).ServiceName(serviceName).Method(method).MethodType(methodType).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {
    registryName := "registryName_example" // string | 
    serviceName := "serviceName_example" // string | 
    method := "method_example" // string | 
    methodType := "methodType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRegistryResourceAPI.RemoveMethod(context.Background(), registryName).ServiceName(serviceName).Method(method).MethodType(methodType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.RemoveMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceName** | **string** |  | 
 **method** | **string** |  | 
 **methodType** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveService

> RemoveService(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRegistryResourceAPI.RemoveService(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.RemoveService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetProtoData

> SetProtoData(ctx, registryName, filename).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
)

func main() {
    registryName := "registryName_example" // string | 
    filename := "filename_example" // string | 
    body := string(BYTE_ARRAY_DATA_HERE) // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRegistryResourceAPI.SetProtoData(context.Background(), registryName, filename).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryResourceAPI.SetProtoData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryName** | **string** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProtoDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

