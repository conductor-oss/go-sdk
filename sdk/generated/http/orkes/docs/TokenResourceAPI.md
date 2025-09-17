# \TokenResourceAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateToken**](TokenResourceAPI.md#GenerateToken) | **Post** /token | Generate JWT with the given access key
[**GetUserInfo**](TokenResourceAPI.md#GetUserInfo) | **Get** /token/userInfo | Get the user info from the token



## GenerateToken

> map[string]interface{} GenerateToken(ctx).GenerateTokenRequest(generateTokenRequest).Execute()

Generate JWT with the given access key

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
    generateTokenRequest := *openapiclient.NewGenerateTokenRequest("KeyId_example", "KeySecret_example") // GenerateTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenResourceAPI.GenerateToken(context.Background()).GenerateTokenRequest(generateTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenResourceAPI.GenerateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateToken`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TokenResourceAPI.GenerateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateTokenRequest** | [**GenerateTokenRequest**](GenerateTokenRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserInfo

> map[string]interface{} GetUserInfo(ctx).Claims(claims).Execute()

Get the user info from the token

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
    claims := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenResourceAPI.GetUserInfo(context.Background()).Claims(claims).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenResourceAPI.GetUserInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TokenResourceAPI.GetUserInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **claims** | **bool** |  | [default to false]

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

