# \MetricsTokenResourceAPI

All URIs are relative to *https://siliconmint-dev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Token**](MetricsTokenResourceAPI.md#Token) | **Get** /metrics/token | 



## Token

> MetricsToken Token(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsTokenResourceAPI.Token(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsTokenResourceAPI.Token``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Token`: MetricsToken
    fmt.Fprintf(os.Stdout, "Response from `MetricsTokenResourceAPI.Token`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTokenRequest struct via the builder pattern


### Return type

[**MetricsToken**](MetricsToken.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

