# \IncomingWebhookResourceAPI

All URIs are relative to *https://siliconmint-dev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleWebhook**](IncomingWebhookResourceAPI.md#HandleWebhook) | **Get** /webhook/{id} | 
[**HandleWebhook1**](IncomingWebhookResourceAPI.md#HandleWebhook1) | **Post** /webhook/{id} | 



## HandleWebhook

> string HandleWebhook(ctx, id).RequestParams(requestParams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes_v4"
)

func main() {
    id := "id_example" // string | 
    requestParams := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IncomingWebhookResourceAPI.HandleWebhook(context.Background(), id).RequestParams(requestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncomingWebhookResourceAPI.HandleWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HandleWebhook`: string
    fmt.Fprintf(os.Stdout, "Response from `IncomingWebhookResourceAPI.HandleWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestParams** | **map[string]map[string]interface{}** |  | 

### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleWebhook1

> string HandleWebhook1(ctx, id).RequestParams(requestParams).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes_v4"
)

func main() {
    id := "id_example" // string | 
    requestParams := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IncomingWebhookResourceAPI.HandleWebhook1(context.Background(), id).RequestParams(requestParams).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncomingWebhookResourceAPI.HandleWebhook1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HandleWebhook1`: string
    fmt.Fprintf(os.Stdout, "Response from `IncomingWebhookResourceAPI.HandleWebhook1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleWebhook1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestParams** | **map[string]map[string]interface{}** |  | 
 **body** | **string** |  | 

### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

