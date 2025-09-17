# \EventExecutionResourceAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventHandlersForEvent1**](EventExecutionResourceAPI.md#GetEventHandlersForEvent1) | **Get** /event/execution | Get All active Event Handlers
[**GetEventHandlersForEvent2**](EventExecutionResourceAPI.md#GetEventHandlersForEvent2) | **Get** /event/execution/{eventHandlerName} | Get event handlers for a given event



## GetEventHandlersForEvent1

> SearchResultHandledEventResponse GetEventHandlersForEvent1(ctx).Execute()

Get All active Event Handlers

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
    resp, r, err := apiClient.EventExecutionResourceAPI.GetEventHandlersForEvent1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventExecutionResourceAPI.GetEventHandlersForEvent1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventHandlersForEvent1`: SearchResultHandledEventResponse
    fmt.Fprintf(os.Stdout, "Response from `EventExecutionResourceAPI.GetEventHandlersForEvent1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventHandlersForEvent1Request struct via the builder pattern


### Return type

[**SearchResultHandledEventResponse**](SearchResultHandledEventResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventHandlersForEvent2

> []ExtendedEventExecution GetEventHandlersForEvent2(ctx, eventHandlerName).From(from).Execute()

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
    eventHandlerName := "eventHandlerName_example" // string | 
    from := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventExecutionResourceAPI.GetEventHandlersForEvent2(context.Background(), eventHandlerName).From(from).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventExecutionResourceAPI.GetEventHandlersForEvent2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventHandlersForEvent2`: []ExtendedEventExecution
    fmt.Fprintf(os.Stdout, "Response from `EventExecutionResourceAPI.GetEventHandlersForEvent2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventHandlerName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventHandlersForEvent2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int64** |  | 

### Return type

[**[]ExtendedEventExecution**](ExtendedEventExecution.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

