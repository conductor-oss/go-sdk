# \EventResourceAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEventHandler**](EventResourceAPI.md#AddEventHandler) | **Post** /event | Add a new event handler.
[**DeleteQueueConfig**](EventResourceAPI.md#DeleteQueueConfig) | **Delete** /event/queue/config/{queueType}/{queueName} | Delete queue config by name
[**DeleteTagForEventHandler**](EventResourceAPI.md#DeleteTagForEventHandler) | **Delete** /event/{name}/tags | Delete a tag for event handler
[**GetEventHandlerByName**](EventResourceAPI.md#GetEventHandlerByName) | **Get** /event/handler/{name} | Get event handler by name
[**GetEventHandlers**](EventResourceAPI.md#GetEventHandlers) | **Get** /event | Get all the event handlers
[**GetEventHandlersForEvent**](EventResourceAPI.md#GetEventHandlersForEvent) | **Get** /event/{event} | Get event handlers for a given event
[**GetQueueConfig**](EventResourceAPI.md#GetQueueConfig) | **Get** /event/queue/config/{queueType}/{queueName} | Get queue config by name
[**GetQueueNames**](EventResourceAPI.md#GetQueueNames) | **Get** /event/queue/config | Get all queue configs
[**GetTagsForEventHandler**](EventResourceAPI.md#GetTagsForEventHandler) | **Get** /event/{name}/tags | Get tags by event handler
[**HandleIncomingEvent**](EventResourceAPI.md#HandleIncomingEvent) | **Post** /event/handleIncomingEvent | Handle an incoming event
[**PutQueueConfig**](EventResourceAPI.md#PutQueueConfig) | **Put** /event/queue/config/{queueType}/{queueName} | Create or update queue config by name
[**PutTagForEventHandler**](EventResourceAPI.md#PutTagForEventHandler) | **Put** /event/{name}/tags | Put a tag to event handler
[**RemoveEventHandlerStatus**](EventResourceAPI.md#RemoveEventHandlerStatus) | **Delete** /event/{name} | Remove an event handler
[**Test**](EventResourceAPI.md#Test) | **Get** /event/handler/ | Get event handler by name
[**TestConnectivity**](EventResourceAPI.md#TestConnectivity) | **Post** /event/queue/connectivity | Test connectivity for a given queue using a workflow with EVENT task and an EventHandler
[**UpdateEventHandler**](EventResourceAPI.md#UpdateEventHandler) | **Put** /event | Update an existing event handler.



## AddEventHandler

> AddEventHandler(ctx).EventHandler(eventHandler).Execute()

Add a new event handler.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {
    eventHandler := []openapiclient.EventHandler{*openapiclient.NewEventHandler()} // []EventHandler | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventResourceAPI.AddEventHandler(context.Background()).EventHandler(eventHandler).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.AddEventHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddEventHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventHandler** | [**[]EventHandler**](EventHandler.md) |  | 

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


## DeleteQueueConfig

> DeleteQueueConfig(ctx, queueType, queueName).Execute()

Delete queue config by name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {
    queueType := "queueType_example" // string | 
    queueName := "queueName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventResourceAPI.DeleteQueueConfig(context.Background(), queueType, queueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.DeleteQueueConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueType** | **string** |  | 
**queueName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueueConfigRequest struct via the builder pattern


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


## DeleteTagForEventHandler

> DeleteTagForEventHandler(ctx, name).Tag(tag).Execute()

Delete a tag for event handler

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {
    name := "name_example" // string | 
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventResourceAPI.DeleteTagForEventHandler(context.Background(), name).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.DeleteTagForEventHandler``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTagForEventHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | [**[]Tag**](Tag.md) |  | 

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


## GetEventHandlerByName

> EventHandler GetEventHandlerByName(ctx, name).Execute()

Get event handler by name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventResourceAPI.GetEventHandlerByName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.GetEventHandlerByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventHandlerByName`: EventHandler
    fmt.Fprintf(os.Stdout, "Response from `EventResourceAPI.GetEventHandlerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventHandlerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventHandler**](EventHandler.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventHandlers

> []EventHandler GetEventHandlers(ctx).Execute()

Get all the event handlers

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventResourceAPI.GetEventHandlers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.GetEventHandlers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventHandlers`: []EventHandler
    fmt.Fprintf(os.Stdout, "Response from `EventResourceAPI.GetEventHandlers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventHandlersRequest struct via the builder pattern


### Return type

[**[]EventHandler**](EventHandler.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventHandlersForEvent

> []EventHandler GetEventHandlersForEvent(ctx, event).ActiveOnly(activeOnly).Execute()

Get event handlers for a given event

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {
    event := "event_example" // string | 
    activeOnly := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventResourceAPI.GetEventHandlersForEvent(context.Background(), event).ActiveOnly(activeOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.GetEventHandlersForEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventHandlersForEvent`: []EventHandler
    fmt.Fprintf(os.Stdout, "Response from `EventResourceAPI.GetEventHandlersForEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**event** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventHandlersForEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeOnly** | **bool** |  | [default to true]

### Return type

[**[]EventHandler**](EventHandler.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueueConfig

> map[string]map[string]interface{} GetQueueConfig(ctx, queueType, queueName).Execute()

Get queue config by name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {
    queueType := "queueType_example" // string | 
    queueName := "queueName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventResourceAPI.GetQueueConfig(context.Background(), queueType, queueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.GetQueueConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueueConfig`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EventResourceAPI.GetQueueConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueType** | **string** |  | 
**queueName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueueNames

> map[string]string GetQueueNames(ctx).Execute()

Get all queue configs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventResourceAPI.GetQueueNames(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.GetQueueNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueueNames`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `EventResourceAPI.GetQueueNames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueNamesRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagsForEventHandler

> []Tag GetTagsForEventHandler(ctx, name).Execute()

Get tags by event handler

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventResourceAPI.GetTagsForEventHandler(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.GetTagsForEventHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagsForEventHandler`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `EventResourceAPI.GetTagsForEventHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsForEventHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Tag**](Tag.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleIncomingEvent

> HandleIncomingEvent(ctx).RequestBody(requestBody).Execute()

Handle an incoming event

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventResourceAPI.HandleIncomingEvent(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.HandleIncomingEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleIncomingEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]interface{}** |  | 

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


## PutQueueConfig

> PutQueueConfig(ctx, queueType, queueName).Body(body).Execute()

Create or update queue config by name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {
    queueType := "queueType_example" // string | 
    queueName := "queueName_example" // string | 
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventResourceAPI.PutQueueConfig(context.Background(), queueType, queueName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.PutQueueConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueType** | **string** |  | 
**queueName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutQueueConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

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


## PutTagForEventHandler

> PutTagForEventHandler(ctx, name).Tag(tag).Execute()

Put a tag to event handler

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {
    name := "name_example" // string | 
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventResourceAPI.PutTagForEventHandler(context.Background(), name).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.PutTagForEventHandler``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutTagForEventHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | [**[]Tag**](Tag.md) |  | 

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


## RemoveEventHandlerStatus

> RemoveEventHandlerStatus(ctx, name).Execute()

Remove an event handler

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventResourceAPI.RemoveEventHandlerStatus(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.RemoveEventHandlerStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveEventHandlerStatusRequest struct via the builder pattern


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


## Test

> EventHandler Test(ctx).Execute()

Get event handler by name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventResourceAPI.Test(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.Test``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Test`: EventHandler
    fmt.Fprintf(os.Stdout, "Response from `EventResourceAPI.Test`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestRequest struct via the builder pattern


### Return type

[**EventHandler**](EventHandler.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConnectivity

> ConnectivityTestResult TestConnectivity(ctx).ConnectivityTestInput(connectivityTestInput).Execute()

Test connectivity for a given queue using a workflow with EVENT task and an EventHandler

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {
    connectivityTestInput := *openapiclient.NewConnectivityTestInput("Sink_example") // ConnectivityTestInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventResourceAPI.TestConnectivity(context.Background()).ConnectivityTestInput(connectivityTestInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.TestConnectivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestConnectivity`: ConnectivityTestResult
    fmt.Fprintf(os.Stdout, "Response from `EventResourceAPI.TestConnectivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestConnectivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectivityTestInput** | [**ConnectivityTestInput**](ConnectivityTestInput.md) |  | 

### Return type

[**ConnectivityTestResult**](ConnectivityTestResult.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEventHandler

> UpdateEventHandler(ctx).EventHandler(eventHandler).Execute()

Update an existing event handler.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
)

func main() {
    eventHandler := *openapiclient.NewEventHandler() // EventHandler | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventResourceAPI.UpdateEventHandler(context.Background()).EventHandler(eventHandler).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventResourceAPI.UpdateEventHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEventHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventHandler** | [**EventHandler**](EventHandler.md) |  | 

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

