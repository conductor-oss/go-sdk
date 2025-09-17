# \TaskResourceAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**All**](TaskResourceAPI.md#All) | **Get** /tasks/queue/all | Get the details about each queue
[**AllVerbose**](TaskResourceAPI.md#AllVerbose) | **Get** /tasks/queue/all/verbose | Get the details about each queue
[**BatchPoll**](TaskResourceAPI.md#BatchPoll) | **Get** /tasks/poll/batch/{tasktype} | Batch poll for a task of a certain type
[**GetAllPollData**](TaskResourceAPI.md#GetAllPollData) | **Get** /tasks/queue/polldata/all | Get the last poll data for all task types
[**GetPollData**](TaskResourceAPI.md#GetPollData) | **Get** /tasks/queue/polldata | Get the last poll data for a given task type
[**GetTask**](TaskResourceAPI.md#GetTask) | **Get** /tasks/{taskId} | Get task by Id
[**GetTaskLogs**](TaskResourceAPI.md#GetTaskLogs) | **Get** /tasks/{taskId}/log | Get Task Execution Logs
[**Log**](TaskResourceAPI.md#Log) | **Post** /tasks/{taskId}/log | Log Task Execution Details
[**Poll**](TaskResourceAPI.md#Poll) | **Get** /tasks/poll/{tasktype} | Poll for a task of a certain type
[**RequeuePendingTask**](TaskResourceAPI.md#RequeuePendingTask) | **Post** /tasks/queue/requeue/{taskType} | Requeue pending tasks
[**Search2**](TaskResourceAPI.md#Search2) | **Get** /tasks/search | Search for tasks based in payload and other parameters
[**SignalWorkflowTaskASync**](TaskResourceAPI.md#SignalWorkflowTaskASync) | **Post** /tasks/{workflowId}/{status}/signal | Update running task in the workflow with given status and output asynchronously
[**SignalWorkflowTaskSync**](TaskResourceAPI.md#SignalWorkflowTaskSync) | **Post** /tasks/{workflowId}/{status}/signal/sync | Update running task in the workflow with given status and output synchronously and return back updated workflow
[**Size**](TaskResourceAPI.md#Size) | **Get** /tasks/queue/sizes | Get Task type queue sizes
[**UpdateTask**](TaskResourceAPI.md#UpdateTask) | **Post** /tasks | Update a task
[**UpdateTask1**](TaskResourceAPI.md#UpdateTask1) | **Post** /tasks/{workflowId}/{taskRefName}/{status} | Update a task By Ref Name. The output data is merged if data from a previous API call already exists.
[**UpdateTaskSync**](TaskResourceAPI.md#UpdateTaskSync) | **Post** /tasks/{workflowId}/{taskRefName}/{status}/sync | Update a task By Ref Name synchronously. The output data is merged if data from a previous API call already exists.
[**UpdateTaskV2**](TaskResourceAPI.md#UpdateTaskV2) | **Post** /tasks/update-v2 | Update a task



## All

> map[string]int64 All(ctx).Execute()

Get the details about each queue

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.All(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.All``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `All`: map[string]int64
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.All`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAllRequest struct via the builder pattern


### Return type

**map[string]int64**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllVerbose

> map[string]map[string]map[string]int64 AllVerbose(ctx).Execute()

Get the details about each queue

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.AllVerbose(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.AllVerbose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllVerbose`: map[string]map[string]map[string]int64
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.AllVerbose`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAllVerboseRequest struct via the builder pattern


### Return type

[**map[string]map[string]map[string]int64**](map.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchPoll

> []Task BatchPoll(ctx, tasktype).Workerid(workerid).Domain(domain).Count(count).Timeout(timeout).Execute()

Batch poll for a task of a certain type

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
    tasktype := "tasktype_example" // string | 
    workerid := "workerid_example" // string |  (optional)
    domain := "domain_example" // string |  (optional)
    count := int32(56) // int32 |  (optional) (default to 1)
    timeout := int32(56) // int32 |  (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.BatchPoll(context.Background(), tasktype).Workerid(workerid).Domain(domain).Count(count).Timeout(timeout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.BatchPoll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchPoll`: []Task
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.BatchPoll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tasktype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchPollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workerid** | **string** |  | 
 **domain** | **string** |  | 
 **count** | **int32** |  | [default to 1]
 **timeout** | **int32** |  | [default to 100]

### Return type

[**[]Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPollData

> map[string]map[string]interface{} GetAllPollData(ctx).WorkerSize(workerSize).WorkerOpt(workerOpt).QueueSize(queueSize).QueueOpt(queueOpt).LastPollTimeSize(lastPollTimeSize).LastPollTimeOpt(lastPollTimeOpt).Execute()

Get the last poll data for all task types

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
    workerSize := int64(789) // int64 |  (optional)
    workerOpt := "workerOpt_example" // string |  (optional)
    queueSize := int64(789) // int64 |  (optional)
    queueOpt := "queueOpt_example" // string |  (optional)
    lastPollTimeSize := int64(789) // int64 |  (optional)
    lastPollTimeOpt := "lastPollTimeOpt_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.GetAllPollData(context.Background()).WorkerSize(workerSize).WorkerOpt(workerOpt).QueueSize(queueSize).QueueOpt(queueOpt).LastPollTimeSize(lastPollTimeSize).LastPollTimeOpt(lastPollTimeOpt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.GetAllPollData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPollData`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.GetAllPollData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPollDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workerSize** | **int64** |  | 
 **workerOpt** | **string** |  | 
 **queueSize** | **int64** |  | 
 **queueOpt** | **string** |  | 
 **lastPollTimeSize** | **int64** |  | 
 **lastPollTimeOpt** | **string** |  | 

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


## GetPollData

> []PollData GetPollData(ctx).TaskType(taskType).Execute()

Get the last poll data for a given task type

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
    taskType := "taskType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.GetPollData(context.Background()).TaskType(taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.GetPollData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPollData`: []PollData
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.GetPollData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPollDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskType** | **string** |  | 

### Return type

[**[]PollData**](PollData.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> Task GetTask(ctx, taskId).Execute()

Get task by Id

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
    taskId := "taskId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.GetTask(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.GetTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTask`: Task
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskLogs

> []TaskExecLog GetTaskLogs(ctx, taskId).Execute()

Get Task Execution Logs

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
    taskId := "taskId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.GetTaskLogs(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.GetTaskLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskLogs`: []TaskExecLog
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.GetTaskLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TaskExecLog**](TaskExecLog.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Log

> Log(ctx, taskId).Body(body).Execute()

Log Task Execution Details

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
    taskId := "taskId_example" // string | 
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TaskResourceAPI.Log(context.Background(), taskId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.Log``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

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


## Poll

> Task Poll(ctx, tasktype).Workerid(workerid).Domain(domain).Execute()

Poll for a task of a certain type

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
    tasktype := "tasktype_example" // string | 
    workerid := "workerid_example" // string |  (optional)
    domain := "domain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.Poll(context.Background(), tasktype).Workerid(workerid).Domain(domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.Poll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Poll`: Task
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.Poll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tasktype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workerid** | **string** |  | 
 **domain** | **string** |  | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequeuePendingTask

> string RequeuePendingTask(ctx, taskType).Execute()

Requeue pending tasks

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
    taskType := "taskType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.RequeuePendingTask(context.Background(), taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.RequeuePendingTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequeuePendingTask`: string
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.RequeuePendingTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequeuePendingTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search2

> SearchResultTaskSummary Search2(ctx).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()

Search for tasks based in payload and other parameters



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
    start := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)
    sort := "sort_example" // string |  (optional)
    freeText := "freeText_example" // string |  (optional) (default to "*")
    query := "query_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.Search2(context.Background()).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.Search2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search2`: SearchResultTaskSummary
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.Search2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearch2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **sort** | **string** |  | 
 **freeText** | **string** |  | [default to &quot;*&quot;]
 **query** | **string** |  | 

### Return type

[**SearchResultTaskSummary**](SearchResultTaskSummary.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignalWorkflowTaskASync

> SignalWorkflowTaskASync(ctx, workflowId, status).RequestBody(requestBody).Execute()

Update running task in the workflow with given status and output asynchronously

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
    workflowId := "workflowId_example" // string | 
    status := "status_example" // string | 
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TaskResourceAPI.SignalWorkflowTaskASync(context.Background(), workflowId, status).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.SignalWorkflowTaskASync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 
**status** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignalWorkflowTaskASyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## SignalWorkflowTaskSync

> SignalResponse SignalWorkflowTaskSync(ctx, workflowId, status).RequestBody(requestBody).ReturnStrategy(returnStrategy).Execute()

Update running task in the workflow with given status and output synchronously and return back updated workflow

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
    workflowId := "workflowId_example" // string | 
    status := "status_example" // string | 
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 
    returnStrategy := "returnStrategy_example" // string |  (optional) (default to "TARGET_WORKFLOW")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.SignalWorkflowTaskSync(context.Background(), workflowId, status).RequestBody(requestBody).ReturnStrategy(returnStrategy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.SignalWorkflowTaskSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignalWorkflowTaskSync`: SignalResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.SignalWorkflowTaskSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 
**status** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignalWorkflowTaskSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]map[string]interface{}** |  | 
 **returnStrategy** | **string** |  | [default to &quot;TARGET_WORKFLOW&quot;]

### Return type

[**SignalResponse**](SignalResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Size

> map[string]int32 Size(ctx).TaskType(taskType).Execute()

Get Task type queue sizes

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
    taskType := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.Size(context.Background()).TaskType(taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.Size``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Size`: map[string]int32
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.Size`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskType** | **[]string** |  | 

### Return type

**map[string]int32**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> string UpdateTask(ctx).TaskResult(taskResult).Execute()

Update a task

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
    taskResult := *openapiclient.NewTaskResult("TaskId_example", "WorkflowInstanceId_example") // TaskResult | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.UpdateTask(context.Background()).TaskResult(taskResult).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.UpdateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTask`: string
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.UpdateTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskResult** | [**TaskResult**](TaskResult.md) |  | 

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


## UpdateTask1

> string UpdateTask1(ctx, workflowId, taskRefName, status).RequestBody(requestBody).Workerid(workerid).Execute()

Update a task By Ref Name. The output data is merged if data from a previous API call already exists.

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
    workflowId := "workflowId_example" // string | 
    taskRefName := "taskRefName_example" // string | 
    status := "status_example" // string | 
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 
    workerid := "workerid_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.UpdateTask1(context.Background(), workflowId, taskRefName, status).RequestBody(requestBody).Workerid(workerid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.UpdateTask1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTask1`: string
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.UpdateTask1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 
**taskRefName** | **string** |  | 
**status** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTask1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requestBody** | **map[string]map[string]interface{}** |  | 
 **workerid** | **string** |  | 

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


## UpdateTaskSync

> Workflow UpdateTaskSync(ctx, workflowId, taskRefName, status).RequestBody(requestBody).Workerid(workerid).Execute()

Update a task By Ref Name synchronously. The output data is merged if data from a previous API call already exists.

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
    workflowId := "workflowId_example" // string | 
    taskRefName := "taskRefName_example" // string | 
    status := "status_example" // string | 
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 
    workerid := "workerid_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.UpdateTaskSync(context.Background(), workflowId, taskRefName, status).RequestBody(requestBody).Workerid(workerid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.UpdateTaskSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTaskSync`: Workflow
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.UpdateTaskSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 
**taskRefName** | **string** |  | 
**status** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requestBody** | **map[string]map[string]interface{}** |  | 
 **workerid** | **string** |  | 

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


## UpdateTaskV2

> Task UpdateTaskV2(ctx).TaskResult(taskResult).Execute()

Update a task

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
    taskResult := *openapiclient.NewTaskResult("TaskId_example", "WorkflowInstanceId_example") // TaskResult | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.UpdateTaskV2(context.Background()).TaskResult(taskResult).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.UpdateTaskV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTaskV2`: Task
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.UpdateTaskV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskResult** | [**TaskResult**](TaskResult.md) |  | 

### Return type

[**Task**](Task.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

