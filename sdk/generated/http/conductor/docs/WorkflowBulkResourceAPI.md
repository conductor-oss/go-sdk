# \WorkflowBulkResourceAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWorkflow**](WorkflowBulkResourceAPI.md#DeleteWorkflow) | **Delete** /workflow/bulk/remove | 
[**PauseWorkflow1**](WorkflowBulkResourceAPI.md#PauseWorkflow1) | **Put** /workflow/bulk/pause | Pause the list of workflows
[**Restart1**](WorkflowBulkResourceAPI.md#Restart1) | **Post** /workflow/bulk/restart | Restart the list of completed workflow
[**ResumeWorkflow1**](WorkflowBulkResourceAPI.md#ResumeWorkflow1) | **Put** /workflow/bulk/resume | Resume the list of workflows
[**Retry1**](WorkflowBulkResourceAPI.md#Retry1) | **Post** /workflow/bulk/retry | Retry the last failed task for each workflow from the list
[**SearchWorkflow**](WorkflowBulkResourceAPI.md#SearchWorkflow) | **Post** /workflow/bulk/search | 
[**Terminate**](WorkflowBulkResourceAPI.md#Terminate) | **Post** /workflow/bulk/terminate | Terminate workflows execution
[**TerminateRemove1**](WorkflowBulkResourceAPI.md#TerminateRemove1) | **Delete** /workflow/bulk/terminate-remove | 



## DeleteWorkflow

> BulkResponseString DeleteWorkflow(ctx).RequestBody(requestBody).ArchiveWorkflow(archiveWorkflow).Execute()



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
    requestBody := []string{"Property_example"} // []string | 
    archiveWorkflow := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.DeleteWorkflow(context.Background()).RequestBody(requestBody).ArchiveWorkflow(archiveWorkflow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.DeleteWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkflow`: BulkResponseString
    fmt.Fprintf(os.Stdout, "Response from `WorkflowBulkResourceAPI.DeleteWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 
 **archiveWorkflow** | **bool** |  | [default to true]

### Return type

[**BulkResponseString**](BulkResponseString.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseWorkflow1

> BulkResponseString PauseWorkflow1(ctx).RequestBody(requestBody).Execute()

Pause the list of workflows

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.PauseWorkflow1(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.PauseWorkflow1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseWorkflow1`: BulkResponseString
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

[**BulkResponseString**](BulkResponseString.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Restart1

> BulkResponseString Restart1(ctx).RequestBody(requestBody).UseLatestDefinitions(useLatestDefinitions).Execute()

Restart the list of completed workflow

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
    requestBody := []string{"Property_example"} // []string | 
    useLatestDefinitions := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.Restart1(context.Background()).RequestBody(requestBody).UseLatestDefinitions(useLatestDefinitions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.Restart1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Restart1`: BulkResponseString
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

[**BulkResponseString**](BulkResponseString.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeWorkflow1

> BulkResponseString ResumeWorkflow1(ctx).RequestBody(requestBody).Execute()

Resume the list of workflows

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.ResumeWorkflow1(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.ResumeWorkflow1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeWorkflow1`: BulkResponseString
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

[**BulkResponseString**](BulkResponseString.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retry1

> BulkResponseString Retry1(ctx).RequestBody(requestBody).Execute()

Retry the last failed task for each workflow from the list

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.Retry1(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.Retry1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Retry1`: BulkResponseString
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

[**BulkResponseString**](BulkResponseString.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchWorkflow

> BulkResponseWorkflowModel SearchWorkflow(ctx).RequestBody(requestBody).IncludeTasks(includeTasks).Execute()



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
    requestBody := []string{"Property_example"} // []string | 
    includeTasks := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.SearchWorkflow(context.Background()).RequestBody(requestBody).IncludeTasks(includeTasks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.SearchWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchWorkflow`: BulkResponseWorkflowModel
    fmt.Fprintf(os.Stdout, "Response from `WorkflowBulkResourceAPI.SearchWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 
 **includeTasks** | **bool** |  | [default to true]

### Return type

[**BulkResponseWorkflowModel**](BulkResponseWorkflowModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Terminate

> BulkResponseString Terminate(ctx).RequestBody(requestBody).Reason(reason).Execute()

Terminate workflows execution

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
    requestBody := []string{"Property_example"} // []string | 
    reason := "reason_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.Terminate(context.Background()).RequestBody(requestBody).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.Terminate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Terminate`: BulkResponseString
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

### Return type

[**BulkResponseString**](BulkResponseString.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateRemove1

> BulkResponseString TerminateRemove1(ctx).RequestBody(requestBody).ArchiveWorkflow(archiveWorkflow).Reason(reason).Execute()



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
    requestBody := []string{"Property_example"} // []string | 
    archiveWorkflow := true // bool |  (optional) (default to true)
    reason := "reason_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowBulkResourceAPI.TerminateRemove1(context.Background()).RequestBody(requestBody).ArchiveWorkflow(archiveWorkflow).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowBulkResourceAPI.TerminateRemove1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminateRemove1`: BulkResponseString
    fmt.Fprintf(os.Stdout, "Response from `WorkflowBulkResourceAPI.TerminateRemove1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTerminateRemove1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 
 **archiveWorkflow** | **bool** |  | [default to true]
 **reason** | **string** |  | 

### Return type

[**BulkResponseString**](BulkResponseString.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

