# \AdminResourceAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllConfig**](AdminResourceAPI.md#GetAllConfig) | **Get** /admin/config | Get all the configuration parameters
[**GetEventQueues**](AdminResourceAPI.md#GetEventQueues) | **Get** /admin/queues | Get registered queues
[**RequeueSweep**](AdminResourceAPI.md#RequeueSweep) | **Post** /admin/sweep/requeue/{workflowId} | Queue up all the running workflows for sweep
[**VerifyAndRepairWorkflowConsistency**](AdminResourceAPI.md#VerifyAndRepairWorkflowConsistency) | **Post** /admin/consistency/verifyAndRepair/{workflowId} | Verify and repair workflow consistency
[**View**](AdminResourceAPI.md#View) | **Get** /admin/task/{tasktype} | Get the list of pending tasks for a given task type



## GetAllConfig

> map[string]map[string]interface{} GetAllConfig(ctx).Execute()

Get all the configuration parameters

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
    resp, r, err := apiClient.AdminResourceAPI.GetAllConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminResourceAPI.GetAllConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllConfig`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdminResourceAPI.GetAllConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConfigRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventQueues

> map[string]map[string]interface{} GetEventQueues(ctx).Verbose(verbose).Execute()

Get registered queues

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
    verbose := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminResourceAPI.GetEventQueues(context.Background()).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminResourceAPI.GetEventQueues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventQueues`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdminResourceAPI.GetEventQueues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventQueuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verbose** | **bool** |  | [default to false]

### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequeueSweep

> string RequeueSweep(ctx, workflowId).Execute()

Queue up all the running workflows for sweep

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
    workflowId := "workflowId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminResourceAPI.RequeueSweep(context.Background(), workflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminResourceAPI.RequeueSweep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequeueSweep`: string
    fmt.Fprintf(os.Stdout, "Response from `AdminResourceAPI.RequeueSweep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequeueSweepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyAndRepairWorkflowConsistency

> string VerifyAndRepairWorkflowConsistency(ctx, workflowId).Execute()

Verify and repair workflow consistency

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
    workflowId := "workflowId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminResourceAPI.VerifyAndRepairWorkflowConsistency(context.Background(), workflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminResourceAPI.VerifyAndRepairWorkflowConsistency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyAndRepairWorkflowConsistency`: string
    fmt.Fprintf(os.Stdout, "Response from `AdminResourceAPI.VerifyAndRepairWorkflowConsistency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyAndRepairWorkflowConsistencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## View

> []Task View(ctx, tasktype).Start(start).Count(count).Execute()

Get the list of pending tasks for a given task type

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
    tasktype := "tasktype_example" // string | 
    start := int32(56) // int32 |  (optional) (default to 0)
    count := int32(56) // int32 |  (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminResourceAPI.View(context.Background(), tasktype).Start(start).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminResourceAPI.View``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `View`: []Task
    fmt.Fprintf(os.Stdout, "Response from `AdminResourceAPI.View`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tasktype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** |  | [default to 0]
 **count** | **int32** |  | [default to 100]

### Return type

[**[]Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

