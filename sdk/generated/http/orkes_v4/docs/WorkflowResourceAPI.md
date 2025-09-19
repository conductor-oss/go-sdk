# \WorkflowResourceAPI

All URIs are relative to *https://siliconmint-dev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Decide**](WorkflowResourceAPI.md#Decide) | **Put** /workflow/decide/{workflowId} | Starts the decision task for a workflow
[**Delete1**](WorkflowResourceAPI.md#Delete1) | **Delete** /workflow/{workflowId}/remove | Removes the workflow from the system
[**ExecuteWorkflow**](WorkflowResourceAPI.md#ExecuteWorkflow) | **Post** /workflow/execute/{name}/{version} | Execute a workflow synchronously
[**ExecuteWorkflowAsAPI**](WorkflowResourceAPI.md#ExecuteWorkflowAsAPI) | **Post** /workflow/execute/{name} | Execute a workflow synchronously with input and outputs
[**ExecuteWorkflowAsGetAPI**](WorkflowResourceAPI.md#ExecuteWorkflowAsGetAPI) | **Get** /workflow/execute/{name} | Execute a workflow synchronously with input and outputs using get api
[**GetExecutionStatus**](WorkflowResourceAPI.md#GetExecutionStatus) | **Get** /workflow/{workflowId} | Gets the workflow by workflow id
[**GetExecutionStatusTaskList**](WorkflowResourceAPI.md#GetExecutionStatusTaskList) | **Get** /workflow/{workflowId}/tasks | Gets the workflow tasks by workflow id
[**GetRunningWorkflow**](WorkflowResourceAPI.md#GetRunningWorkflow) | **Get** /workflow/running/{name} | Retrieve all the running workflows
[**GetWorkflowStatusSummary**](WorkflowResourceAPI.md#GetWorkflowStatusSummary) | **Get** /workflow/{workflowId}/status | Gets the workflow by workflow id
[**GetWorkflows**](WorkflowResourceAPI.md#GetWorkflows) | **Post** /workflow/{name}/correlated | Lists workflows for the given correlation id list
[**GetWorkflows1**](WorkflowResourceAPI.md#GetWorkflows1) | **Post** /workflow/correlated/batch | Lists workflows for the given correlation id list and workflow name list
[**GetWorkflows2**](WorkflowResourceAPI.md#GetWorkflows2) | **Get** /workflow/{name}/correlated/{correlationId} | Lists workflows for the given correlation id
[**JumpToTask**](WorkflowResourceAPI.md#JumpToTask) | **Post** /workflow/{workflowId}/jump/{taskReferenceName} | Jump workflow execution to given task
[**PauseWorkflow**](WorkflowResourceAPI.md#PauseWorkflow) | **Put** /workflow/{workflowId}/pause | Pauses the workflow
[**Rerun**](WorkflowResourceAPI.md#Rerun) | **Post** /workflow/{workflowId}/rerun | Reruns the workflow from a specific task
[**ResetWorkflow**](WorkflowResourceAPI.md#ResetWorkflow) | **Post** /workflow/{workflowId}/resetcallbacks | Resets callback times of all non-terminal SIMPLE tasks to 0
[**Restart**](WorkflowResourceAPI.md#Restart) | **Post** /workflow/{workflowId}/restart | Restarts a completed workflow
[**ResumeWorkflow**](WorkflowResourceAPI.md#ResumeWorkflow) | **Put** /workflow/{workflowId}/resume | Resumes the workflow
[**Retry**](WorkflowResourceAPI.md#Retry) | **Post** /workflow/{workflowId}/retry | Retries the last failed task
[**Search**](WorkflowResourceAPI.md#Search) | **Get** /workflow/search | Search for workflows based on payload and other parameters
[**SkipTaskFromWorkflow**](WorkflowResourceAPI.md#SkipTaskFromWorkflow) | **Put** /workflow/{workflowId}/skiptask/{taskReferenceName} | Skips a given task from a current running workflow
[**StartWorkflow**](WorkflowResourceAPI.md#StartWorkflow) | **Post** /workflow | Start a new workflow with StartWorkflowRequest, which allows task to be executed in a domain
[**StartWorkflow1**](WorkflowResourceAPI.md#StartWorkflow1) | **Post** /workflow/{name} | Start a new workflow. Returns the ID of the workflow instance that can be later used for tracking
[**Terminate1**](WorkflowResourceAPI.md#Terminate1) | **Delete** /workflow/{workflowId} | Terminate workflow execution
[**TestWorkflow**](WorkflowResourceAPI.md#TestWorkflow) | **Post** /workflow/test | Test workflow execution using mock data
[**UpdateWorkflowAndTaskState**](WorkflowResourceAPI.md#UpdateWorkflowAndTaskState) | **Post** /workflow/{workflowId}/state | Update a workflow state by updating variables or in progress task
[**UpdateWorkflowState**](WorkflowResourceAPI.md#UpdateWorkflowState) | **Post** /workflow/{workflowId}/variables | Update workflow variables
[**UpgradeRunningWorkflowToVersion**](WorkflowResourceAPI.md#UpgradeRunningWorkflowToVersion) | **Post** /workflow/{workflowId}/upgrade | Upgrade running workflow to newer version



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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes_v4"
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

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete1

> Delete1(ctx, workflowId).ArchiveWorkflow(archiveWorkflow).Execute()

Removes the workflow from the system

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
    workflowId := "workflowId_example" // string | 
    archiveWorkflow := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkflowResourceAPI.Delete1(context.Background(), workflowId).ArchiveWorkflow(archiveWorkflow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.Delete1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDelete1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archiveWorkflow** | **bool** |  | [default to true]

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteWorkflow

> WorkflowRun ExecuteWorkflow(ctx, name, version).RequestId(requestId).StartWorkflowRequest(startWorkflowRequest).WaitUntilTaskRef(waitUntilTaskRef).WaitForSeconds(waitForSeconds).Execute()

Execute a workflow synchronously

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
    name := "name_example" // string | 
    version := int32(56) // int32 | 
    requestId := "requestId_example" // string | 
    startWorkflowRequest := *openapiclient.NewStartWorkflowRequest("Name_example") // StartWorkflowRequest | 
    waitUntilTaskRef := "waitUntilTaskRef_example" // string |  (optional)
    waitForSeconds := int32(56) // int32 |  (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.ExecuteWorkflow(context.Background(), name, version).RequestId(requestId).StartWorkflowRequest(startWorkflowRequest).WaitUntilTaskRef(waitUntilTaskRef).WaitForSeconds(waitForSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.ExecuteWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteWorkflow`: WorkflowRun
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.ExecuteWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**version** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **startWorkflowRequest** | [**StartWorkflowRequest**](StartWorkflowRequest.md) |  | 
 **waitUntilTaskRef** | **string** |  | 
 **waitForSeconds** | **int32** |  | [default to 10]

### Return type

[**WorkflowRun**](WorkflowRun.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteWorkflowAsAPI

> map[string]map[string]interface{} ExecuteWorkflowAsAPI(ctx, name).RequestBody(requestBody).Version(version).RequestId(requestId).WaitUntilTaskRef(waitUntilTaskRef).WaitForSeconds(waitForSeconds).XIdempotencyKey(xIdempotencyKey).XOnConflict(xOnConflict).Execute()

Execute a workflow synchronously with input and outputs

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
    name := "name_example" // string | 
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 
    version := int32(56) // int32 |  (optional)
    requestId := "requestId_example" // string |  (optional)
    waitUntilTaskRef := "waitUntilTaskRef_example" // string |  (optional)
    waitForSeconds := int32(56) // int32 |  (optional) (default to 10)
    xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)
    xOnConflict := "xOnConflict_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.ExecuteWorkflowAsAPI(context.Background(), name).RequestBody(requestBody).Version(version).RequestId(requestId).WaitUntilTaskRef(waitUntilTaskRef).WaitForSeconds(waitForSeconds).XIdempotencyKey(xIdempotencyKey).XOnConflict(xOnConflict).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.ExecuteWorkflowAsAPI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteWorkflowAsAPI`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.ExecuteWorkflowAsAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteWorkflowAsAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** |  | 
 **version** | **int32** |  | 
 **requestId** | **string** |  | 
 **waitUntilTaskRef** | **string** |  | 
 **waitForSeconds** | **int32** |  | [default to 10]
 **xIdempotencyKey** | **string** |  | 
 **xOnConflict** | **string** |  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteWorkflowAsGetAPI

> map[string]map[string]interface{} ExecuteWorkflowAsGetAPI(ctx, name).Version(version).RequestId(requestId).WaitUntilTaskRef(waitUntilTaskRef).WaitForSeconds(waitForSeconds).XIdempotencyKey(xIdempotencyKey).XOnConflict(xOnConflict).Execute()

Execute a workflow synchronously with input and outputs using get api

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
    name := "name_example" // string | 
    version := int32(56) // int32 |  (optional)
    requestId := "requestId_example" // string |  (optional)
    waitUntilTaskRef := "waitUntilTaskRef_example" // string |  (optional)
    waitForSeconds := int32(56) // int32 |  (optional) (default to 10)
    xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)
    xOnConflict := "xOnConflict_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.ExecuteWorkflowAsGetAPI(context.Background(), name).Version(version).RequestId(requestId).WaitUntilTaskRef(waitUntilTaskRef).WaitForSeconds(waitForSeconds).XIdempotencyKey(xIdempotencyKey).XOnConflict(xOnConflict).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.ExecuteWorkflowAsGetAPI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteWorkflowAsGetAPI`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.ExecuteWorkflowAsGetAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteWorkflowAsGetAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** |  | 
 **requestId** | **string** |  | 
 **waitUntilTaskRef** | **string** |  | 
 **waitForSeconds** | **int32** |  | [default to 10]
 **xIdempotencyKey** | **string** |  | 
 **xOnConflict** | **string** |  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecutionStatus

> Workflow GetExecutionStatus(ctx, workflowId).IncludeTasks(includeTasks).Summarize(summarize).Execute()

Gets the workflow by workflow id

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
    workflowId := "workflowId_example" // string | 
    includeTasks := true // bool |  (optional) (default to true)
    summarize := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.GetExecutionStatus(context.Background(), workflowId).IncludeTasks(includeTasks).Summarize(summarize).Execute()
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
 **summarize** | **bool** |  | [default to false]

### Return type

[**Workflow**](Workflow.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecutionStatusTaskList

> TaskListSearchResultSummary GetExecutionStatusTaskList(ctx, workflowId).Start(start).Count(count).Status(status).Execute()

Gets the workflow tasks by workflow id

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
    workflowId := "workflowId_example" // string | 
    start := int32(56) // int32 |  (optional) (default to 0)
    count := int32(56) // int32 |  (optional) (default to 15)
    status := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.GetExecutionStatusTaskList(context.Background(), workflowId).Start(start).Count(count).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.GetExecutionStatusTaskList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExecutionStatusTaskList`: TaskListSearchResultSummary
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.GetExecutionStatusTaskList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionStatusTaskListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** |  | [default to 0]
 **count** | **int32** |  | [default to 15]
 **status** | **[]string** |  | 

### Return type

[**TaskListSearchResultSummary**](TaskListSearchResultSummary.md)

### Authorization

[api_key](../README.md#api_key)

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes_v4"
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

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowStatusSummary

> WorkflowStatus GetWorkflowStatusSummary(ctx, workflowId).IncludeOutput(includeOutput).IncludeVariables(includeVariables).Execute()

Gets the workflow by workflow id

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
    workflowId := "workflowId_example" // string | 
    includeOutput := true // bool |  (optional) (default to false)
    includeVariables := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.GetWorkflowStatusSummary(context.Background(), workflowId).IncludeOutput(includeOutput).IncludeVariables(includeVariables).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.GetWorkflowStatusSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowStatusSummary`: WorkflowStatus
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.GetWorkflowStatusSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowStatusSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeOutput** | **bool** |  | [default to false]
 **includeVariables** | **bool** |  | [default to false]

### Return type

[**WorkflowStatus**](WorkflowStatus.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflows

> map[string]interface{} GetWorkflows(ctx, name).RequestBody(requestBody).IncludeClosed(includeClosed).IncludeTasks(includeTasks).Execute()

Lists workflows for the given correlation id list

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
    // response from `GetWorkflows`: map[string]interface{}
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

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflows1

> map[string]interface{} GetWorkflows1(ctx).CorrelationIdsSearchRequest(correlationIdsSearchRequest).IncludeClosed(includeClosed).IncludeTasks(includeTasks).Execute()

Lists workflows for the given correlation id list and workflow name list

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
    correlationIdsSearchRequest := *openapiclient.NewCorrelationIdsSearchRequest() // CorrelationIdsSearchRequest | 
    includeClosed := true // bool |  (optional) (default to false)
    includeTasks := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.GetWorkflows1(context.Background()).CorrelationIdsSearchRequest(correlationIdsSearchRequest).IncludeClosed(includeClosed).IncludeTasks(includeTasks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.GetWorkflows1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflows1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.GetWorkflows1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflows1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **correlationIdsSearchRequest** | [**CorrelationIdsSearchRequest**](CorrelationIdsSearchRequest.md) |  | 
 **includeClosed** | **bool** |  | [default to false]
 **includeTasks** | **bool** |  | [default to false]

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflows2

> []Workflow GetWorkflows2(ctx, name, correlationId).IncludeClosed(includeClosed).IncludeTasks(includeTasks).Execute()

Lists workflows for the given correlation id

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
    name := "name_example" // string | 
    correlationId := "correlationId_example" // string | 
    includeClosed := true // bool |  (optional) (default to false)
    includeTasks := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.GetWorkflows2(context.Background(), name, correlationId).IncludeClosed(includeClosed).IncludeTasks(includeTasks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.GetWorkflows2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflows2`: []Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.GetWorkflows2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**correlationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflows2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeClosed** | **bool** |  | [default to false]
 **includeTasks** | **bool** |  | [default to false]

### Return type

[**[]Workflow**](Workflow.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JumpToTask

> JumpToTask(ctx, workflowId).TaskReferenceName(taskReferenceName).RequestBody(requestBody).Execute()

Jump workflow execution to given task



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
    workflowId := "workflowId_example" // string | 
    taskReferenceName := "taskReferenceName_example" // string | 
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkflowResourceAPI.JumpToTask(context.Background(), workflowId).TaskReferenceName(taskReferenceName).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.JumpToTask``: %v\n", err)
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

Other parameters are passed through a pointer to a apiJumpToTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskReferenceName** | **string** |  | 
 **requestBody** | **map[string]map[string]interface{}** |  | 

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes_v4"
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

[api_key](../README.md#api_key)

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes_v4"
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

[api_key](../README.md#api_key)

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes_v4"
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

[api_key](../README.md#api_key)

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes_v4"
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

[api_key](../README.md#api_key)

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes_v4"
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

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retry

> Retry(ctx, workflowId).ResumeSubworkflowTasks(resumeSubworkflowTasks).RetryIfRetriedByParent(retryIfRetriedByParent).Execute()

Retries the last failed task

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
    workflowId := "workflowId_example" // string | 
    resumeSubworkflowTasks := true // bool |  (optional) (default to false)
    retryIfRetriedByParent := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkflowResourceAPI.Retry(context.Background(), workflowId).ResumeSubworkflowTasks(resumeSubworkflowTasks).RetryIfRetriedByParent(retryIfRetriedByParent).Execute()
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
 **retryIfRetriedByParent** | **bool** |  | [default to true]

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> ScrollableSearchResultWorkflowSummary Search(ctx).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).SkipCache(skipCache).Execute()

Search for workflows based on payload and other parameters



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
    start := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)
    sort := "sort_example" // string |  (optional)
    freeText := "freeText_example" // string |  (optional) (default to "*")
    query := "query_example" // string |  (optional)
    skipCache := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.Search(context.Background()).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).SkipCache(skipCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: ScrollableSearchResultWorkflowSummary
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
 **skipCache** | **bool** |  | [default to false]

### Return type

[**ScrollableSearchResultWorkflowSummary**](ScrollableSearchResultWorkflowSummary.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkipTaskFromWorkflow

> SkipTaskFromWorkflow(ctx, workflowId, taskReferenceName).SkipTaskRequest(skipTaskRequest).Execute()

Skips a given task from a current running workflow

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
    workflowId := "workflowId_example" // string | 
    taskReferenceName := "taskReferenceName_example" // string | 
    skipTaskRequest := *openapiclient.NewSkipTaskRequest() // SkipTaskRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkflowResourceAPI.SkipTaskFromWorkflow(context.Background(), workflowId, taskReferenceName).SkipTaskRequest(skipTaskRequest).Execute()
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


 **skipTaskRequest** | [**SkipTaskRequest**](SkipTaskRequest.md) |  | 

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes_v4"
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

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartWorkflow1

> string StartWorkflow1(ctx, name).RequestBody(requestBody).Version(version).CorrelationId(correlationId).Priority(priority).XIdempotencyKey(xIdempotencyKey).XOnConflict(xOnConflict).Execute()

Start a new workflow. Returns the ID of the workflow instance that can be later used for tracking

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
    name := "name_example" // string | 
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 
    version := int32(56) // int32 |  (optional)
    correlationId := "correlationId_example" // string |  (optional)
    priority := int32(56) // int32 |  (optional) (default to 0)
    xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)
    xOnConflict := "xOnConflict_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.StartWorkflow1(context.Background(), name).RequestBody(requestBody).Version(version).CorrelationId(correlationId).Priority(priority).XIdempotencyKey(xIdempotencyKey).XOnConflict(xOnConflict).Execute()
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
 **xIdempotencyKey** | **string** |  | 
 **xOnConflict** | **string** |  | 

### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Terminate1

> Terminate1(ctx, workflowId).Reason(reason).TriggerFailureWorkflow(triggerFailureWorkflow).Execute()

Terminate workflow execution

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
    workflowId := "workflowId_example" // string | 
    reason := "reason_example" // string |  (optional)
    triggerFailureWorkflow := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkflowResourceAPI.Terminate1(context.Background(), workflowId).Reason(reason).TriggerFailureWorkflow(triggerFailureWorkflow).Execute()
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
 **triggerFailureWorkflow** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes_v4"
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

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowAndTaskState

> WorkflowRun UpdateWorkflowAndTaskState(ctx, workflowId).RequestId(requestId).WorkflowStateUpdate(workflowStateUpdate).WaitUntilTaskRef(waitUntilTaskRef).WaitForSeconds(waitForSeconds).Execute()

Update a workflow state by updating variables or in progress task



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
    workflowId := "workflowId_example" // string | 
    requestId := "requestId_example" // string | 
    workflowStateUpdate := *openapiclient.NewWorkflowStateUpdate() // WorkflowStateUpdate | 
    waitUntilTaskRef := "waitUntilTaskRef_example" // string |  (optional)
    waitForSeconds := int32(56) // int32 |  (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.UpdateWorkflowAndTaskState(context.Background(), workflowId).RequestId(requestId).WorkflowStateUpdate(workflowStateUpdate).WaitUntilTaskRef(waitUntilTaskRef).WaitForSeconds(waitForSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.UpdateWorkflowAndTaskState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowAndTaskState`: WorkflowRun
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.UpdateWorkflowAndTaskState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowAndTaskStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestId** | **string** |  | 
 **workflowStateUpdate** | [**WorkflowStateUpdate**](WorkflowStateUpdate.md) |  | 
 **waitUntilTaskRef** | **string** |  | 
 **waitForSeconds** | **int32** |  | [default to 10]

### Return type

[**WorkflowRun**](WorkflowRun.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowState

> Workflow UpdateWorkflowState(ctx, workflowId).RequestBody(requestBody).Execute()

Update workflow variables



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
    workflowId := "workflowId_example" // string | 
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowResourceAPI.UpdateWorkflowState(context.Background(), workflowId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.UpdateWorkflowState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowState`: Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowResourceAPI.UpdateWorkflowState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** |  | 

### Return type

[**Workflow**](Workflow.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeRunningWorkflowToVersion

> UpgradeRunningWorkflowToVersion(ctx, workflowId).UpgradeWorkflowRequest(upgradeWorkflowRequest).Execute()

Upgrade running workflow to newer version



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
    workflowId := "workflowId_example" // string | 
    upgradeWorkflowRequest := *openapiclient.NewUpgradeWorkflowRequest("Name_example") // UpgradeWorkflowRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkflowResourceAPI.UpgradeRunningWorkflowToVersion(context.Background(), workflowId).UpgradeWorkflowRequest(upgradeWorkflowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowResourceAPI.UpgradeRunningWorkflowToVersion``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpgradeRunningWorkflowToVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upgradeWorkflowRequest** | [**UpgradeWorkflowRequest**](UpgradeWorkflowRequest.md) |  | 

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

