# \WorkflowBulkResourceAPI

All URIs are relative to *https://siliconmint-dev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](WorkflowBulkResourceAPI.md#Delete) | **Post** /workflow/bulk/delete | Permanently remove workflows from the system
[**PauseWorkflow1**](WorkflowBulkResourceAPI.md#PauseWorkflow1) | **Put** /workflow/bulk/pause | Pause the list of workflows
[**Restart1**](WorkflowBulkResourceAPI.md#Restart1) | **Post** /workflow/bulk/restart | Restart the list of completed workflow
[**ResumeWorkflow1**](WorkflowBulkResourceAPI.md#ResumeWorkflow1) | **Put** /workflow/bulk/resume | Resume the list of workflows
[**Retry1**](WorkflowBulkResourceAPI.md#Retry1) | **Post** /workflow/bulk/retry | Retry the last failed task for each workflow from the list
[**Terminate**](WorkflowBulkResourceAPI.md#Terminate) | **Post** /workflow/bulk/terminate | Terminate workflows execution



## Delete

> BulkResponse Delete(ctx).RequestBody(requestBody).Execute()

Permanently remove workflows from the system

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.Delete(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: BulkResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowBulkResourceAPI.Delete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseWorkflow1

> BulkResponse PauseWorkflow1(ctx).RequestBody(requestBody).Execute()

Pause the list of workflows

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.PauseWorkflow1(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.PauseWorkflow1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseWorkflow1`: BulkResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowBulkResourceAPI.PauseWorkflow1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPauseWorkflow1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Restart1

> BulkResponse Restart1(ctx).RequestBody(requestBody).UseLatestDefinitions(useLatestDefinitions).Execute()

Restart the list of completed workflow

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
    requestBody := []string{"Property_example"} // []string | 
    useLatestDefinitions := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.Restart1(context.Background()).RequestBody(requestBody).UseLatestDefinitions(useLatestDefinitions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.Restart1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Restart1`: BulkResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowBulkResourceAPI.Restart1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestart1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 
 **useLatestDefinitions** | **bool** |  | [default to false]

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeWorkflow1

> BulkResponse ResumeWorkflow1(ctx).RequestBody(requestBody).Execute()

Resume the list of workflows

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.ResumeWorkflow1(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.ResumeWorkflow1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeWorkflow1`: BulkResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowBulkResourceAPI.ResumeWorkflow1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResumeWorkflow1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retry1

> BulkResponse Retry1(ctx).RequestBody(requestBody).Execute()

Retry the last failed task for each workflow from the list

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.Retry1(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.Retry1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Retry1`: BulkResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowBulkResourceAPI.Retry1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetry1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Terminate

> BulkResponse Terminate(ctx).RequestBody(requestBody).Reason(reason).TriggerFailureWorkflow(triggerFailureWorkflow).Execute()

Terminate workflows execution

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
    requestBody := []string{"Property_example"} // []string | 
    reason := "reason_example" // string |  (optional)
    triggerFailureWorkflow := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.Terminate(context.Background()).RequestBody(requestBody).Reason(reason).TriggerFailureWorkflow(triggerFailureWorkflow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.Terminate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Terminate`: BulkResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowBulkResourceAPI.Terminate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTerminateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 
 **reason** | **string** |  | 
 **triggerFailureWorkflow** | **bool** |  | [default to false]

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

