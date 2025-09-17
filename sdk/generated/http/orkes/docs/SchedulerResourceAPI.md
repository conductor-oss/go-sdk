# \SchedulerResourceAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSchedule**](SchedulerResourceAPI.md#DeleteSchedule) | **Delete** /scheduler/schedules/{name} | Deletes an existing workflow schedule by name
[**DeleteTagForSchedule**](SchedulerResourceAPI.md#DeleteTagForSchedule) | **Delete** /scheduler/schedules/{name}/tags | Delete a tag for schedule
[**GetAllSchedules**](SchedulerResourceAPI.md#GetAllSchedules) | **Get** /scheduler/schedules | Get all existing workflow schedules and optionally filter by workflow name
[**GetNextFewSchedules**](SchedulerResourceAPI.md#GetNextFewSchedules) | **Get** /scheduler/nextFewSchedules | Get list of the next x (default 3, max 5) execution times for a scheduler
[**GetSchedule**](SchedulerResourceAPI.md#GetSchedule) | **Get** /scheduler/schedules/{name} | Get an existing workflow schedule by name
[**GetSchedulesByTag**](SchedulerResourceAPI.md#GetSchedulesByTag) | **Get** /scheduler/schedules/tags | Get schedules by tag
[**GetTagsForSchedule**](SchedulerResourceAPI.md#GetTagsForSchedule) | **Get** /scheduler/schedules/{name}/tags | Get tags by schedule
[**PauseAllSchedules**](SchedulerResourceAPI.md#PauseAllSchedules) | **Get** /scheduler/admin/pause | Pause all scheduling in a single conductor server instance (for debugging only)
[**PauseSchedule**](SchedulerResourceAPI.md#PauseSchedule) | **Get** /scheduler/schedules/{name}/pause | Pauses an existing schedule by name
[**PutTagForSchedule**](SchedulerResourceAPI.md#PutTagForSchedule) | **Put** /scheduler/schedules/{name}/tags | Put a tag to schedule
[**RequeueAllExecutionRecords**](SchedulerResourceAPI.md#RequeueAllExecutionRecords) | **Get** /scheduler/admin/requeue | Requeue all execution records
[**ResumeAllSchedules**](SchedulerResourceAPI.md#ResumeAllSchedules) | **Get** /scheduler/admin/resume | Resume all scheduling
[**ResumeSchedule**](SchedulerResourceAPI.md#ResumeSchedule) | **Get** /scheduler/schedules/{name}/resume | Resume a paused schedule by name
[**SaveSchedule**](SchedulerResourceAPI.md#SaveSchedule) | **Post** /scheduler/schedules | Create or update a schedule for a specified workflow with a corresponding start workflow request
[**SearchV2**](SchedulerResourceAPI.md#SearchV2) | **Get** /scheduler/search/executions | Search for workflow executions based on payload and other parameters



## DeleteSchedule

> map[string]interface{} DeleteSchedule(ctx, name).Execute()

Deletes an existing workflow schedule by name

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerResourceAPI.DeleteSchedule(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.DeleteSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSchedule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SchedulerResourceAPI.DeleteSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteTagForSchedule

> DeleteTagForSchedule(ctx, name).Tag(tag).Execute()

Delete a tag for schedule

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
    name := "name_example" // string | 
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SchedulerResourceAPI.DeleteTagForSchedule(context.Background(), name).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.DeleteTagForSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagForScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | [**[]Tag**](Tag.md) |  | 

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


## GetAllSchedules

> []WorkflowScheduleModel GetAllSchedules(ctx).WorkflowName(workflowName).Execute()

Get all existing workflow schedules and optionally filter by workflow name

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
    workflowName := "workflowName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerResourceAPI.GetAllSchedules(context.Background()).WorkflowName(workflowName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.GetAllSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSchedules`: []WorkflowScheduleModel
    fmt.Fprintf(os.Stdout, "Response from `SchedulerResourceAPI.GetAllSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowName** | **string** |  | 

### Return type

[**[]WorkflowScheduleModel**](WorkflowScheduleModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNextFewSchedules

> []int64 GetNextFewSchedules(ctx).CronExpression(cronExpression).ScheduleStartTime(scheduleStartTime).ScheduleEndTime(scheduleEndTime).Limit(limit).Execute()

Get list of the next x (default 3, max 5) execution times for a scheduler

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
    cronExpression := "cronExpression_example" // string | 
    scheduleStartTime := int64(789) // int64 |  (optional)
    scheduleEndTime := int64(789) // int64 |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 3)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerResourceAPI.GetNextFewSchedules(context.Background()).CronExpression(cronExpression).ScheduleStartTime(scheduleStartTime).ScheduleEndTime(scheduleEndTime).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.GetNextFewSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNextFewSchedules`: []int64
    fmt.Fprintf(os.Stdout, "Response from `SchedulerResourceAPI.GetNextFewSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNextFewSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cronExpression** | **string** |  | 
 **scheduleStartTime** | **int64** |  | 
 **scheduleEndTime** | **int64** |  | 
 **limit** | **int32** |  | [default to 3]

### Return type

**[]int64**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchedule

> WorkflowSchedule GetSchedule(ctx, name).Execute()

Get an existing workflow schedule by name

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerResourceAPI.GetSchedule(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.GetSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchedule`: WorkflowSchedule
    fmt.Fprintf(os.Stdout, "Response from `SchedulerResourceAPI.GetSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowSchedule**](WorkflowSchedule.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchedulesByTag

> []WorkflowScheduleModel GetSchedulesByTag(ctx).Tag(tag).Execute()

Get schedules by tag

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
    tag := "tag_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerResourceAPI.GetSchedulesByTag(context.Background()).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.GetSchedulesByTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchedulesByTag`: []WorkflowScheduleModel
    fmt.Fprintf(os.Stdout, "Response from `SchedulerResourceAPI.GetSchedulesByTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSchedulesByTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **string** |  | 

### Return type

[**[]WorkflowScheduleModel**](WorkflowScheduleModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagsForSchedule

> []Tag GetTagsForSchedule(ctx, name).Execute()

Get tags by schedule

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerResourceAPI.GetTagsForSchedule(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.GetTagsForSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagsForSchedule`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `SchedulerResourceAPI.GetTagsForSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsForScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Tag**](Tag.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseAllSchedules

> map[string]map[string]interface{} PauseAllSchedules(ctx).Execute()

Pause all scheduling in a single conductor server instance (for debugging only)

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerResourceAPI.PauseAllSchedules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.PauseAllSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseAllSchedules`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SchedulerResourceAPI.PauseAllSchedules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPauseAllSchedulesRequest struct via the builder pattern


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


## PauseSchedule

> map[string]interface{} PauseSchedule(ctx, name).Execute()

Pauses an existing schedule by name

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerResourceAPI.PauseSchedule(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.PauseSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseSchedule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SchedulerResourceAPI.PauseSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PutTagForSchedule

> PutTagForSchedule(ctx, name).Tag(tag).Execute()

Put a tag to schedule

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
    name := "name_example" // string | 
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SchedulerResourceAPI.PutTagForSchedule(context.Background(), name).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.PutTagForSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTagForScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | [**[]Tag**](Tag.md) |  | 

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


## RequeueAllExecutionRecords

> map[string]map[string]interface{} RequeueAllExecutionRecords(ctx).Execute()

Requeue all execution records

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerResourceAPI.RequeueAllExecutionRecords(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.RequeueAllExecutionRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequeueAllExecutionRecords`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SchedulerResourceAPI.RequeueAllExecutionRecords`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRequeueAllExecutionRecordsRequest struct via the builder pattern


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


## ResumeAllSchedules

> map[string]map[string]interface{} ResumeAllSchedules(ctx).Execute()

Resume all scheduling

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerResourceAPI.ResumeAllSchedules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.ResumeAllSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeAllSchedules`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SchedulerResourceAPI.ResumeAllSchedules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResumeAllSchedulesRequest struct via the builder pattern


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


## ResumeSchedule

> map[string]interface{} ResumeSchedule(ctx, name).Execute()

Resume a paused schedule by name

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerResourceAPI.ResumeSchedule(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.ResumeSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeSchedule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SchedulerResourceAPI.ResumeSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## SaveSchedule

> map[string]interface{} SaveSchedule(ctx).SaveScheduleRequest(saveScheduleRequest).Execute()

Create or update a schedule for a specified workflow with a corresponding start workflow request

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
    saveScheduleRequest := *openapiclient.NewSaveScheduleRequest("CronExpression_example", "Name_example", *openapiclient.NewStartWorkflowRequest("Name_example")) // SaveScheduleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerResourceAPI.SaveSchedule(context.Background()).SaveScheduleRequest(saveScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.SaveSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveSchedule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SchedulerResourceAPI.SaveSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveScheduleRequest** | [**SaveScheduleRequest**](SaveScheduleRequest.md) |  | 

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


## SearchV2

> SearchResultWorkflowScheduleExecutionModel SearchV2(ctx).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()

Search for workflow executions based on payload and other parameters



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
    start := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)
    sort := "sort_example" // string |  (optional)
    freeText := "freeText_example" // string |  (optional) (default to "*")
    query := "query_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerResourceAPI.SearchV2(context.Background()).Start(start).Size(size).Sort(sort).FreeText(freeText).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerResourceAPI.SearchV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchV2`: SearchResultWorkflowScheduleExecutionModel
    fmt.Fprintf(os.Stdout, "Response from `SchedulerResourceAPI.SearchV2`: %v\n", resp)
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

[**SearchResultWorkflowScheduleExecutionModel**](SearchResultWorkflowScheduleExecutionModel.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

