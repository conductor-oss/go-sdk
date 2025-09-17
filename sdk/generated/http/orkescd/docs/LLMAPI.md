# \LLMAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTokenLimit**](LLMAPI.md#GetTokenLimit) | **Get** /integrations/llm/{name}/token | Get the Token Limit for an integration
[**GetTokenUsage**](LLMAPI.md#GetTokenUsage) | **Get** /integrations/llm/{name}/token/history | Get Token Usage by Integration provider
[**UpdateTokenLimit**](LLMAPI.md#UpdateTokenLimit) | **Post** /integrations/llm/{name}/token | Register Token Limit for an integration



## GetTokenLimit

> TokenLimit GetTokenLimit(ctx, name).Execute()

Get the Token Limit for an integration

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
    resp, r, err := apiClient.LLMAPI.GetTokenLimit(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LLMAPI.GetTokenLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenLimit`: TokenLimit
    fmt.Fprintf(os.Stdout, "Response from `LLMAPI.GetTokenLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenLimit**](TokenLimit.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenUsage

> []TokenUsageLog GetTokenUsage(ctx, name).Model(model).LookUpWindow(lookUpWindow).Execute()

Get Token Usage by Integration provider

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
    model := "model_example" // string |  (optional) (default to "*")
    lookUpWindow := int32(56) // int32 |  (optional) (default to 31)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LLMAPI.GetTokenUsage(context.Background(), name).Model(model).LookUpWindow(lookUpWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LLMAPI.GetTokenUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenUsage`: []TokenUsageLog
    fmt.Fprintf(os.Stdout, "Response from `LLMAPI.GetTokenUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | **string** |  | [default to &quot;*&quot;]
 **lookUpWindow** | **int32** |  | [default to 31]

### Return type

[**[]TokenUsageLog**](TokenUsageLog.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTokenLimit

> UpdateTokenLimit(ctx, name).TokenLimit(tokenLimit).Execute()

Register Token Limit for an integration

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
    tokenLimit := *openapiclient.NewTokenLimit() // TokenLimit | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LLMAPI.UpdateTokenLimit(context.Background(), name).TokenLimit(tokenLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LLMAPI.UpdateTokenLimit``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateTokenLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenLimit** | [**TokenLimit**](TokenLimit.md) |  | 

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

