# \WorkflowResourceAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Decide**](WorkflowResourceAPI.md#Decide) | **Put** /workflow/decide/{workflowId} | Starts the decision task for a workflow
[**Delete**](WorkflowResourceAPI.md#Delete) | **Delete** /workflow/{workflowId}/remove | Removes the workflow from the system
[**GetExecutionStatus**](WorkflowResourceAPI.md#GetExecutionStatus) | **Get** /workflow/{workflowId} | Gets the workflow by workflow id
[**GetExternalStorageLocation**](WorkflowResourceAPI.md#GetExternalStorageLocation) | **Get** /workflow/external-storage-location | Get the uri and path of the external storage where the workflow payload is to be stored
[**GetExternalStorageLocation1**](WorkflowResourceAPI.md#GetExternalStorageLocation1) | **Get** /workflow/externalstoragelocation | Get the uri and path of the external storage where the workflow payload is to be stored
[**GetRunningWorkflow**](WorkflowResourceAPI.md#GetRunningWorkflow) | **Get** /workflow/running/{name} | Retrieve all the running workflows
[**GetWorkflows**](WorkflowResourceAPI.md#GetWorkflows) | **Post** /workflow/{name}/correlated | Lists workflows for the given correlation id list
[**GetWorkflows1**](WorkflowResourceAPI.md#GetWorkflows1) | **Get** /workflow/{name}/correlated/{correlationId} | Lists workflows for the given correlation id
[**PauseWorkflow**](WorkflowResourceAPI.md#PauseWorkflow) | **Put** /workflow/{workflowId}/pause | Pauses the workflow
[**Rerun**](WorkflowResourceAPI.md#Rerun) | **Post** /workflow/{workflowId}/rerun | Reruns the workflow from a specific task
[**ResetWorkflow**](WorkflowResourceAPI.md#ResetWorkflow) | **Post** /workflow/{workflowId}/resetcallbacks | Resets callback times of all non-terminal SIMPLE tasks to 0
[**Restart**](WorkflowResourceAPI.md#Restart) | **Post** /workflow/{workflowId}/restart | Restarts a completed workflow
[**ResumeWorkflow**](WorkflowResourceAPI.md#ResumeWorkflow) | **Put** /workflow/{workflowId}/resume | Resumes the workflow
[**Retry**](WorkflowResourceAPI.md#Retry) | **Post** /workflow/{workflowId}/retry | Retries the last failed task
[**Search**](WorkflowResourceAPI.md#Search) | **Get** /workflow/search | Search for workflows based on payload and other parameters
[**SearchV2**](WorkflowResourceAPI.md#SearchV2) | **Get** /workflow/search-v2 | Search for workflows based on payload and other parameters
[**SearchWorkflowsByTasks**](WorkflowResourceAPI.md#SearchWorkflowsByTasks) | **Get** /workflow/search-by-tasks | Search for workflows based on task parameters
[**SearchWorkflowsByTasksV2**](WorkflowResourceAPI.md#SearchWorkflowsByTasksV2) | **Get** /workflow/search-by-tasks-v2 | Search for workflows based on task parameters
[**SkipTaskFromWorkflow**](WorkflowResourceAPI.md#SkipTaskFromWorkflow) | **Put** /workflow/{workflowId}/skiptask/{taskReferenceName} | Skips a given task from a current running workflow
[**StartWorkflow**](WorkflowResourceAPI.md#StartWorkflow) | **Post** /workflow | Start a new workflow with StartWorkflowRequest, which allows task to be executed in a domain
[**StartWorkflow1**](WorkflowResourceAPI.md#StartWorkflow1) | **Post** /workflow/{name} | Start a new workflow. Returns the ID of the workflow instance that can be later used for tracking
[**Terminate1**](WorkflowResourceAPI.md#Terminate1) | **Delete** /workflow/{workflowId} | Terminate workflow execution
[**TerminateRemove**](WorkflowResourceAPI.md#TerminateRemove) | **Delete** /workflow/{workflowId}/terminate-remove | Terminate workflow execution and remove the workflow from the system
[**TestWorkflow**](WorkflowResourceAPI.md#TestWorkflow) | **Post** /workflow/test | Test workflow execution using mock data



## Decide

> Decide(ctx, workflowId).Execute()

Starts the decision task for a workflow

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
    r, err := apiClient.WorkflowResourceAPI.Decide(context.Background(), workflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.Decide``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, workflowId).ArchiveWorkflow(archiveWorkflow).Execute()

Removes the workflow from the system

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
    archiveWorkflow := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkflowResourceAPI.Delete(context.Background(), workflowId).ArchiveWorkflow(archiveWorkflow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archiveWorkflow** | **bool** |  | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecutionStatus

> Workflow GetExecutionStatus(ctx, workflowId).IncludeTasks(includeTasks).Execute()

Gets the workflow by workflow id

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
    includeTasks := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.GetExecutionStatus(context.Background(), workflowId).IncludeTasks(includeTasks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.GetExecutionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExecutionStatus`: Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.GetExecutionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTasks** | **bool** |  | [default to true]

### Return type

[**Workflow**](Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalStorageLocation

> ExternalStorageLocation GetExternalStorageLocation(ctx).Path(path).Operation(operation).PayloadType(payloadType).Execute()

Get the uri and path of the external storage where the workflow payload is to be stored

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
    path := "path_example" // string | 
    operation := "operation_example" // string | 
    payloadType := "payloadType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.GetExternalStorageLocation(context.Background()).Path(path).Operation(operation).PayloadType(payloadType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.GetExternalStorageLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalStorageLocation`: ExternalStorageLocation
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.GetExternalStorageLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalStorageLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 
 **operation** | **string** |  | 
 **payloadType** | **string** |  | 

### Return type

[**ExternalStorageLocation**](ExternalStorageLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalStorageLocation1

> ExternalStorageLocation GetExternalStorageLocation1(ctx).Path(path).Operation(operation).PayloadType(payloadType).Execute()

Get the uri and path of the external storage where the workflow payload is to be stored

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
    path := "path_example" // string | 
    operation := "operation_example" // string | 
    payloadType := "payloadType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.GetExternalStorageLocation1(context.Background()).Path(path).Operation(operation).PayloadType(payloadType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.GetExternalStorageLocation1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalStorageLocation1`: ExternalStorageLocation
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.GetExternalStorageLocation1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalStorageLocation1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 
 **operation** | **string** |  | 
 **payloadType** | **string** |  | 

### Return type

[**ExternalStorageLocation**](ExternalStorageLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunningWorkflow

> []string GetRunningWorkflow(ctx, name).Version(version).StartTime(startTime).EndTime(endTime).Execute()

Retrieve all the running workflows

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
    name := "name_example" // string | 
    version := int32(56) // int32 |  (optional) (default to 1)
    startTime := int64(789) // int64 |  (optional)
    endTime := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.GetRunningWorkflow(context.Background(), name).Version(version).StartTime(startTime).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.GetRunningWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRunningWorkflow`: []string
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.GetRunningWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunningWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** |  | [default to 1]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflows

> map[string][]Workflow GetWorkflows(ctx, name).RequestBody(requestBody).IncludeClosed(includeClosed).IncludeTasks(includeTasks).Execute()

Lists workflows for the given correlation id list

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
    name := "name_example" // string | 
    requestBody := []string{"Property_example"} // []string | 
    includeClosed := true // bool |  (optional) (default to false)
    includeTasks := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.GetWorkflows(context.Background(), name).RequestBody(requestBody).IncludeClosed(includeClosed).IncludeTasks(includeTasks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.GetWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflows`: map[string][]Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.GetWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 
 **includeClosed** | **bool** |  | [default to false]
 **includeTasks** | **bool** |  | [default to false]

### Return type

[**map[string][]Workflow**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflows1

> []Workflow GetWorkflows1(ctx, name, correlationId).IncludeClosed(includeClosed).IncludeTasks(includeTasks).Execute()

Lists workflows for the given correlation id

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
    name := "name_example" // string | 
    correlationId := "correlationId_example" // string | 
    includeClosed := true // bool |  (optional) (default to false)
    includeTasks := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.GetWorkflows1(context.Background(), name, correlationId).IncludeClosed(includeClosed).IncludeTasks(includeTasks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.GetWorkflows1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflows1`: []Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.GetWorkflows1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**correlationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflows1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeClosed** | **bool** |  | [default to false]
 **includeTasks** | **bool** |  | [default to false]

### Return type

[**[]Workflow**](Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseWorkflow

> PauseWorkflow(ctx, workflowId).Execute()

Pauses the workflow

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
    r, err := apiClient.WorkflowResourceAPI.PauseWorkflow(context.Background(), workflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.PauseWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Rerun

> string Rerun(ctx, workflowId).RerunWorkflowRequest(rerunWorkflowRequest).Execute()

Reruns the workflow from a specific task

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
    rerunWorkflowRequest := *openapiclient.NewRerunWorkflowRequest() // RerunWorkflowRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.Rerun(context.Background(), workflowId).RerunWorkflowRequest(rerunWorkflowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.Rerun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Rerun`: string
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.Rerun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRerunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rerunWorkflowRequest** | [**RerunWorkflowRequest**](RerunWorkflowRequest.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetWorkflow

> ResetWorkflow(ctx, workflowId).Execute()

Resets callback times of all non-terminal SIMPLE tasks to 0

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
    r, err := apiClient.WorkflowResourceAPI.ResetWorkflow(context.Background(), workflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.ResetWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Restart

> Restart(ctx, workflowId).UseLatestDefinitions(useLatestDefinitions).Execute()

Restarts a completed workflow

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
    useLatestDefinitions := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkflowResourceAPI.Restart(context.Background(), workflowId).UseLatestDefinitions(useLatestDefinitions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.Restart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **useLatestDefinitions** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeWorkflow

> ResumeWorkflow(ctx, workflowId).Execute()

Resumes the workflow

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
    r, err := apiClient.WorkflowResourceAPI.ResumeWorkflow(context.Background(), workflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.ResumeWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retry

> Retry(ctx, workflowId).ResumeSubworkflowTasks(resumeSubworkflowTasks).Execute()

Retries the last failed task

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
    resumeSubworkflowTasks := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkflowResourceAPI.Retry(context.Background(), workflowId).ResumeSubworkflowTasks(resumeSubworkflowTasks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.Retry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resumeSubworkflowTasks** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> SearchResultWorkflowSummary Search(ctx).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()

Search for workflows based on payload and other parameters



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
    start := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)
    sort := "sort_example" // string |  (optional)
    freeText := "freeText_example" // string |  (optional) (default to "*")
    query := "query_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.Search(context.Background()).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: SearchResultWorkflowSummary
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **sort** | **string** |  | 
 **freeText** | **string** |  | [default to &quot;*&quot;]
 **query** | **string** |  | 

### Return type

[**SearchResultWorkflowSummary**](SearchResultWorkflowSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchV2

> SearchResultWorkflow SearchV2(ctx).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()

Search for workflows based on payload and other parameters



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
    start := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)
    sort := "sort_example" // string |  (optional)
    freeText := "freeText_example" // string |  (optional) (default to "*")
    query := "query_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.SearchV2(context.Background()).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.SearchV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchV2`: SearchResultWorkflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.SearchV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **sort** | **string** |  | 
 **freeText** | **string** |  | [default to &quot;*&quot;]
 **query** | **string** |  | 

### Return type

[**SearchResultWorkflow**](SearchResultWorkflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchWorkflowsByTasks

> SearchResultWorkflowSummary SearchWorkflowsByTasks(ctx).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()

Search for workflows based on task parameters



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
    start := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)
    sort := "sort_example" // string |  (optional)
    freeText := "freeText_example" // string |  (optional) (default to "*")
    query := "query_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.SearchWorkflowsByTasks(context.Background()).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.SearchWorkflowsByTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchWorkflowsByTasks`: SearchResultWorkflowSummary
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.SearchWorkflowsByTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchWorkflowsByTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **sort** | **string** |  | 
 **freeText** | **string** |  | [default to &quot;*&quot;]
 **query** | **string** |  | 

### Return type

[**SearchResultWorkflowSummary**](SearchResultWorkflowSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchWorkflowsByTasksV2

> SearchResultWorkflow SearchWorkflowsByTasksV2(ctx).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()

Search for workflows based on task parameters



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
    start := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)
    sort := "sort_example" // string |  (optional)
    freeText := "freeText_example" // string |  (optional) (default to "*")
    query := "query_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.SearchWorkflowsByTasksV2(context.Background()).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.SearchWorkflowsByTasksV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchWorkflowsByTasksV2`: SearchResultWorkflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.SearchWorkflowsByTasksV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchWorkflowsByTasksV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **sort** | **string** |  | 
 **freeText** | **string** |  | [default to &quot;*&quot;]
 **query** | **string** |  | 

### Return type

[**SearchResultWorkflow**](SearchResultWorkflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkipTaskFromWorkflow

> SkipTaskFromWorkflow(ctx, workflowId, taskReferenceName).Arg2(arg2).Execute()

Skips a given task from a current running workflow

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
    taskReferenceName := "taskReferenceName_example" // string | 
    arg2 := *openapiclient.NewSkipTaskRequest() // SkipTaskRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkflowResourceAPI.SkipTaskFromWorkflow(context.Background(), workflowId, taskReferenceName).Arg2(arg2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.SkipTaskFromWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 
**taskReferenceName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSkipTaskFromWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **arg2** | [**SkipTaskRequest**](SkipTaskRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartWorkflow

> string StartWorkflow(ctx).StartWorkflowRequest(startWorkflowRequest).Execute()

Start a new workflow with StartWorkflowRequest, which allows task to be executed in a domain

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
    startWorkflowRequest := *openapiclient.NewStartWorkflowRequest("Name_example") // StartWorkflowRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.StartWorkflow(context.Background()).StartWorkflowRequest(startWorkflowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.StartWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartWorkflow`: string
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.StartWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startWorkflowRequest** | [**StartWorkflowRequest**](StartWorkflowRequest.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartWorkflow1

> string StartWorkflow1(ctx, name).RequestBody(requestBody).Version(version).CorrelationId(correlationId).Priority(priority).Execute()

Start a new workflow. Returns the ID of the workflow instance that can be later used for tracking

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
    name := "name_example" // string | 
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 
    version := int32(56) // int32 |  (optional)
    correlationId := "correlationId_example" // string |  (optional)
    priority := int32(56) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.StartWorkflow1(context.Background(), name).RequestBody(requestBody).Version(version).CorrelationId(correlationId).Priority(priority).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.StartWorkflow1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartWorkflow1`: string
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.StartWorkflow1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartWorkflow1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** |  | 
 **version** | **int32** |  | 
 **correlationId** | **string** |  | 
 **priority** | **int32** |  | [default to 0]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Terminate1

> Terminate1(ctx, workflowId).Reason(reason).Execute()

Terminate workflow execution

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
    reason := "reason_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkflowResourceAPI.Terminate1(context.Background(), workflowId).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.Terminate1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminate1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateRemove

> TerminateRemove(ctx, workflowId).Reason(reason).ArchiveWorkflow(archiveWorkflow).Execute()

Terminate workflow execution and remove the workflow from the system

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
    reason := "reason_example" // string |  (optional)
    archiveWorkflow := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkflowResourceAPI.TerminateRemove(context.Background(), workflowId).Reason(reason).ArchiveWorkflow(archiveWorkflow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.TerminateRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** |  | 
 **archiveWorkflow** | **bool** |  | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestWorkflow

> Workflow TestWorkflow(ctx).WorkflowTestRequest(workflowTestRequest).Execute()

Test workflow execution using mock data

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
    workflowTestRequest := *openapiclient.NewWorkflowTestRequest("Name_example") // WorkflowTestRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.TestWorkflow(context.Background()).WorkflowTestRequest(workflowTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.TestWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestWorkflow`: Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.TestWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowTestRequest** | [**WorkflowTestRequest**](WorkflowTestRequest.md) |  | 

### Return type

[**Workflow**](Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

