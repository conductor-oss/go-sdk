# \MetricsResourceAPI

All URIs are relative to *https://siliconmint-dev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrometheusTaskMetrics**](MetricsResourceAPI.md#PrometheusTaskMetrics) | **Get** /metrics/task/{taskName} | Returns prometheus task metrics



## PrometheusTaskMetrics

> map[string]map[string]interface{} PrometheusTaskMetrics(ctx, taskName).Start(start).End(end).Step(step).Execute()

Returns prometheus task metrics



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
    taskName := "taskName_example" // string | 
    start := "start_example" // string | 
    end := "end_example" // string | 
    step := "step_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsResourceAPI.PrometheusTaskMetrics(context.Background(), taskName).Start(start).End(end).Step(step).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsResourceAPI.PrometheusTaskMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrometheusTaskMetrics`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetricsResourceAPI.PrometheusTaskMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrometheusTaskMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **string** |  | 
 **end** | **string** |  | 
 **step** | **string** |  | 

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

