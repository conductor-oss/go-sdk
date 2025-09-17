# \TaskResourceAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**All**](TaskResourceAPI.md#All) | **Get** /tasks/queue/all | Get the details about each queue
[**AllVerbose**](TaskResourceAPI.md#AllVerbose) | **Get** /tasks/queue/all/verbose | Get the details about each queue
[**BatchPoll**](TaskResourceAPI.md#BatchPoll) | **Get** /tasks/poll/batch/{tasktype} | Batch poll for a task of a certain type
[**GetAllPollData**](TaskResourceAPI.md#GetAllPollData) | **Get** /tasks/queue/polldata/all | Get the last poll data for all task types
[**GetExternalStorageLocation2**](TaskResourceAPI.md#GetExternalStorageLocation2) | **Get** /tasks/external-storage-location | Get the external uri where the task payload is to be stored
[**GetExternalStorageLocation3**](TaskResourceAPI.md#GetExternalStorageLocation3) | **Get** /tasks/externalstoragelocation | Get the external uri where the task payload is to be stored
[**GetPollData**](TaskResourceAPI.md#GetPollData) | **Get** /tasks/queue/polldata | Get the last poll data for a given task type
[**GetTask**](TaskResourceAPI.md#GetTask) | **Get** /tasks/{taskId} | Get task by Id
[**GetTaskLogs**](TaskResourceAPI.md#GetTaskLogs) | **Get** /tasks/{taskId}/log | Get Task Execution Logs
[**Log**](TaskResourceAPI.md#Log) | **Post** /tasks/{taskId}/log | Log Task Execution Details
[**Poll**](TaskResourceAPI.md#Poll) | **Get** /tasks/poll/{tasktype} | Poll for a task of a certain type
[**RequeuePendingTask**](TaskResourceAPI.md#RequeuePendingTask) | **Post** /tasks/queue/requeue/{taskType} | Requeue pending tasks
[**Search1**](TaskResourceAPI.md#Search1) | **Get** /tasks/search | Search for tasks based in payload and other parameters
[**SearchV21**](TaskResourceAPI.md#SearchV21) | **Get** /tasks/search-v2 | Search for tasks based in payload and other parameters
[**Size**](TaskResourceAPI.md#Size) | **Get** /tasks/queue/sizes | Deprecated. Please use /tasks/queue/size endpoint
[**TaskDepth**](TaskResourceAPI.md#TaskDepth) | **Get** /tasks/queue/size | Get queue size for a task type.
[**UpdateTask**](TaskResourceAPI.md#UpdateTask) | **Post** /tasks | Update a task
[**UpdateTask1**](TaskResourceAPI.md#UpdateTask1) | **Post** /tasks/{workflowId}/{taskRefName}/{status} | Update a task By Ref Name
[**UpdateTaskSync**](TaskResourceAPI.md#UpdateTaskSync) | **Post** /tasks/{workflowId}/{taskRefName}/{status}/sync | Update a task By Ref Name synchronously and return the updated workflow
[**UpdateTaskV2**](TaskResourceAPI.md#UpdateTaskV2) | **Post** /tasks/update-v2 | Update a task and return the next available task to be processed



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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
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

No authorization required

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
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

No authorization required

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPollData

> []PollData GetAllPollData(ctx).Execute()

Get the last poll data for all task types

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
    resp, r, err := apiClient.TaskResourceAPI.GetAllPollData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.GetAllPollData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPollData`: []PollData
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.GetAllPollData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPollDataRequest struct via the builder pattern


### Return type

[**[]PollData**](PollData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalStorageLocation2

> ExternalStorageLocation GetExternalStorageLocation2(ctx).Path(path).Operation(operation).PayloadType(payloadType).Execute()

Get the external uri where the task payload is to be stored

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
    resp, r, err := apiClient.TaskResourceAPI.GetExternalStorageLocation2(context.Background()).Path(path).Operation(operation).PayloadType(payloadType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.GetExternalStorageLocation2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalStorageLocation2`: ExternalStorageLocation
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.GetExternalStorageLocation2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalStorageLocation2Request struct via the builder pattern


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


## GetExternalStorageLocation3

> ExternalStorageLocation GetExternalStorageLocation3(ctx).Path(path).Operation(operation).PayloadType(payloadType).Execute()

Get the external uri where the task payload is to be stored

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
    resp, r, err := apiClient.TaskResourceAPI.GetExternalStorageLocation3(context.Background()).Path(path).Operation(operation).PayloadType(payloadType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.GetExternalStorageLocation3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalStorageLocation3`: ExternalStorageLocation
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.GetExternalStorageLocation3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalStorageLocation3Request struct via the builder pattern


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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
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

No authorization required

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
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

No authorization required

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
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

No authorization required

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
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

No authorization required

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
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

No authorization required

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search1

> SearchResultTaskSummary Search1(ctx).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()

Search for tasks based in payload and other parameters



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
    resp, r, err := apiClient.TaskResourceAPI.Search1(context.Background()).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.Search1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search1`: SearchResultTaskSummary
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.Search1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearch1Request struct via the builder pattern


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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchV21

> SearchResultTask SearchV21(ctx).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()

Search for tasks based in payload and other parameters



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
    resp, r, err := apiClient.TaskResourceAPI.SearchV21(context.Background()).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.SearchV21``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchV21`: SearchResultTask
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.SearchV21`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchV21Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **sort** | **string** |  | 
 **freeText** | **string** |  | [default to &quot;*&quot;]
 **query** | **string** |  | 

### Return type

[**SearchResultTask**](SearchResultTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Size

> map[string]int32 Size(ctx).TaskType(taskType).Execute()

Deprecated. Please use /tasks/queue/size endpoint

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskDepth

> int32 TaskDepth(ctx).TaskType(taskType).Domain(domain).IsolationGroupId(isolationGroupId).ExecutionNamespace(executionNamespace).Execute()

Get queue size for a task type.

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
    taskType := "taskType_example" // string | 
    domain := "domain_example" // string |  (optional)
    isolationGroupId := "isolationGroupId_example" // string |  (optional)
    executionNamespace := "executionNamespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResourceAPI.TaskDepth(context.Background()).TaskType(taskType).Domain(domain).IsolationGroupId(isolationGroupId).ExecutionNamespace(executionNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResourceAPI.TaskDepth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskDepth`: int32
    fmt.Fprintf(os.Stdout, "Response from `TaskResourceAPI.TaskDepth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskDepthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskType** | **string** |  | 
 **domain** | **string** |  | 
 **isolationGroupId** | **string** |  | 
 **executionNamespace** | **string** |  | 

### Return type

**int32**

### Authorization

No authorization required

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
)

func main() {
    taskResult := *openapiclient.NewTaskResult("WorkflowInstanceId_example", "TaskId_example") // TaskResult | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask1

> string UpdateTask1(ctx, workflowId, taskRefName, status).RequestBody(requestBody).Workerid(workerid).Execute()

Update a task By Ref Name

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskSync

> Workflow UpdateTaskSync(ctx, workflowId, taskRefName, status).RequestBody(requestBody).Workerid(workerid).Execute()

Update a task By Ref Name synchronously and return the updated workflow

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskV2

> Task UpdateTaskV2(ctx).TaskResult(taskResult).Execute()

Update a task and return the next available task to be processed

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
    taskResult := *openapiclient.NewTaskResult("WorkflowInstanceId_example", "TaskId_example") // TaskResult | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

