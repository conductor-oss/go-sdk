# \EventResourceAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEventHandler**](EventResourceAPI.md#AddEventHandler) | **Post** /event | Add a new event handler.
[**GetEventHandlers**](EventResourceAPI.md#GetEventHandlers) | **Get** /event | Get all the event handlers
[**GetEventHandlersForEvent**](EventResourceAPI.md#GetEventHandlersForEvent) | **Get** /event/{event} | Get event handlers for a given event
[**RemoveEventHandlerStatus**](EventResourceAPI.md#RemoveEventHandlerStatus) | **Delete** /event/{name} | Remove an event handler
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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
)

func main() {
    eventHandler := *openapiclient.NewEventHandler("Name_example", "Event_example", []openapiclient.Action{*openapiclient.NewAction()}) // EventHandler | 

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
 **eventHandler** | [**EventHandler**](EventHandler.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
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

No authorization required

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
)

func main() {
    eventHandler := *openapiclient.NewEventHandler("Name_example", "Event_example", []openapiclient.Action{*openapiclient.NewAction()}) // EventHandler | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

