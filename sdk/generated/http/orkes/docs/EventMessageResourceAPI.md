# \EventMessageResourceAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvents**](EventMessageResourceAPI.md#GetEvents) | **Get** /event/message | Get all event handlers with statistics
[**GetMessages**](EventMessageResourceAPI.md#GetMessages) | **Get** /event/message/{event} | Get event messages for a given event



## GetEvents

> SearchResultHandledEventResponse GetEvents(ctx).From(from).Execute()

Get all event handlers with statistics

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
    from := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventMessageResourceAPI.GetEvents(context.Background()).From(from).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventMessageResourceAPI.GetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvents`: SearchResultHandledEventResponse
    fmt.Fprintf(os.Stdout, "Response from `EventMessageResourceAPI.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** |  | 

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


## GetMessages

> []EventMessage GetMessages(ctx, event).From(from).Execute()

Get event messages for a given event

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
    from := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventMessageResourceAPI.GetMessages(context.Background(), event).From(from).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventMessageResourceAPI.GetMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessages`: []EventMessage
    fmt.Fprintf(os.Stdout, "Response from `EventMessageResourceAPI.GetMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**event** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int64** |  | 

### Return type

[**[]EventMessage**](EventMessage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

