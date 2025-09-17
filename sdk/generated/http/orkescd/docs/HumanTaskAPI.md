# \HumanTaskAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignAndClaim**](HumanTaskAPI.md#AssignAndClaim) | **Post** /human/tasks/{taskId}/externalUser/{userId} | Claim a task to an external user
[**BackPopulateFullTextIndex**](HumanTaskAPI.md#BackPopulateFullTextIndex) | **Get** /human/tasks/backPopulateFullTextIndex | API for backpopulating index data
[**ClaimTask**](HumanTaskAPI.md#ClaimTask) | **Post** /human/tasks/{taskId}/claim | Claim a task by authenticated Conductor user
[**DeleteTaskFromHumanTaskRecords**](HumanTaskAPI.md#DeleteTaskFromHumanTaskRecords) | **Delete** /human/tasks/delete | If the workflow is disconnected from tasks, this API can be used to clean up (in bulk)
[**DeleteTaskFromHumanTaskRecords1**](HumanTaskAPI.md#DeleteTaskFromHumanTaskRecords1) | **Delete** /human/tasks/delete/{taskId} | If the workflow is disconnected from tasks, this API can be used to clean up
[**DeleteTemplateByName**](HumanTaskAPI.md#DeleteTemplateByName) | **Delete** /human/template/{name} | Delete all versions of user form template by name
[**DeleteTemplatesByNameAndVersion**](HumanTaskAPI.md#DeleteTemplatesByNameAndVersion) | **Delete** /human/template/{name}/{version} | Delete a version of form template by name
[**GetAllTemplates**](HumanTaskAPI.md#GetAllTemplates) | **Get** /human/template | List all user form templates or get templates by name, or a template by name and version
[**GetTask1**](HumanTaskAPI.md#GetTask1) | **Get** /human/tasks/{taskId} | Get a task
[**GetTaskDisplayNames**](HumanTaskAPI.md#GetTaskDisplayNames) | **Get** /human/tasks/getTaskDisplayNames | Get list of task display names applicable for the user
[**GetTemplateByNameAndVersion**](HumanTaskAPI.md#GetTemplateByNameAndVersion) | **Get** /human/template/{name}/{version} | Get user form template by name and version
[**GetTemplateByTaskId**](HumanTaskAPI.md#GetTemplateByTaskId) | **Get** /human/template/{humanTaskId} | Get user form by human task id
[**ReassignTask**](HumanTaskAPI.md#ReassignTask) | **Post** /human/tasks/{taskId}/reassign | Reassign a task without completing it
[**ReleaseTask**](HumanTaskAPI.md#ReleaseTask) | **Post** /human/tasks/{taskId}/release | Release a task without completing it
[**SaveTemplate**](HumanTaskAPI.md#SaveTemplate) | **Post** /human/template | Save user form template
[**SaveTemplates**](HumanTaskAPI.md#SaveTemplates) | **Post** /human/template/bulk | Save user form template
[**Search**](HumanTaskAPI.md#Search) | **Post** /human/tasks/search | Search human tasks
[**SkipTask**](HumanTaskAPI.md#SkipTask) | **Post** /human/tasks/{taskId}/skip | If a task is assigned to a user, this API can be used to skip that assignment and move to the next assignee
[**UpdateTaskOutput**](HumanTaskAPI.md#UpdateTaskOutput) | **Post** /human/tasks/{taskId}/update | Update task output, optionally complete
[**UpdateTaskOutputByRef**](HumanTaskAPI.md#UpdateTaskOutputByRef) | **Post** /human/tasks/update/taskRef | Update task output, optionally complete



## AssignAndClaim

> HumanTaskEntry AssignAndClaim(ctx, taskId, userId).OverrideAssignment(overrideAssignment).WithTemplate(withTemplate).Execute()

Claim a task to an external user

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
    userId := "userId_example" // string | 
    overrideAssignment := true // bool |  (optional) (default to false)
    withTemplate := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HumanTaskAPI.AssignAndClaim(context.Background(), taskId, userId).OverrideAssignment(overrideAssignment).WithTemplate(withTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.AssignAndClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignAndClaim`: HumanTaskEntry
    fmt.Fprintf(os.Stdout, "Response from `HumanTaskAPI.AssignAndClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignAndClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **overrideAssignment** | **bool** |  | [default to false]
 **withTemplate** | **bool** |  | [default to false]

### Return type

[**HumanTaskEntry**](HumanTaskEntry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackPopulateFullTextIndex

> map[string]map[string]interface{} BackPopulateFullTextIndex(ctx).Var100(var100).Execute()

API for backpopulating index data

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
    var100 := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HumanTaskAPI.BackPopulateFullTextIndex(context.Background()).Var100(var100).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.BackPopulateFullTextIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackPopulateFullTextIndex`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HumanTaskAPI.BackPopulateFullTextIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackPopulateFullTextIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **var100** | **int32** |  | 

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


## ClaimTask

> HumanTaskEntry ClaimTask(ctx, taskId).OverrideAssignment(overrideAssignment).WithTemplate(withTemplate).Execute()

Claim a task by authenticated Conductor user

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
    overrideAssignment := true // bool |  (optional) (default to false)
    withTemplate := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HumanTaskAPI.ClaimTask(context.Background(), taskId).OverrideAssignment(overrideAssignment).WithTemplate(withTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.ClaimTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClaimTask`: HumanTaskEntry
    fmt.Fprintf(os.Stdout, "Response from `HumanTaskAPI.ClaimTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **overrideAssignment** | **bool** |  | [default to false]
 **withTemplate** | **bool** |  | [default to false]

### Return type

[**HumanTaskEntry**](HumanTaskEntry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTaskFromHumanTaskRecords

> DeleteTaskFromHumanTaskRecords(ctx).RequestBody(requestBody).Execute()

If the workflow is disconnected from tasks, this API can be used to clean up (in bulk)

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HumanTaskAPI.DeleteTaskFromHumanTaskRecords(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.DeleteTaskFromHumanTaskRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskFromHumanTaskRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

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


## DeleteTaskFromHumanTaskRecords1

> DeleteTaskFromHumanTaskRecords1(ctx, taskId).Execute()

If the workflow is disconnected from tasks, this API can be used to clean up

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
    r, err := apiClient.HumanTaskAPI.DeleteTaskFromHumanTaskRecords1(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.DeleteTaskFromHumanTaskRecords1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTaskFromHumanTaskRecords1Request struct via the builder pattern


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


## DeleteTemplateByName

> DeleteTemplateByName(ctx, name).Execute()

Delete all versions of user form template by name

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HumanTaskAPI.DeleteTemplateByName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.DeleteTemplateByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTemplateByNameRequest struct via the builder pattern


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


## DeleteTemplatesByNameAndVersion

> DeleteTemplatesByNameAndVersion(ctx, name, version).Execute()

Delete a version of form template by name

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
    name := "name_example" // string | 
    version := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HumanTaskAPI.DeleteTemplatesByNameAndVersion(context.Background(), name, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.DeleteTemplatesByNameAndVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**version** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplatesByNameAndVersionRequest struct via the builder pattern


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


## GetAllTemplates

> []HumanTaskTemplate GetAllTemplates(ctx).Name(name).Version(version).Execute()

List all user form templates or get templates by name, or a template by name and version

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
    name := "name_example" // string |  (optional)
    version := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HumanTaskAPI.GetAllTemplates(context.Background()).Name(name).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.GetAllTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTemplates`: []HumanTaskTemplate
    fmt.Fprintf(os.Stdout, "Response from `HumanTaskAPI.GetAllTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **version** | **int32** |  | 

### Return type

[**[]HumanTaskTemplate**](HumanTaskTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask1

> HumanTaskEntry GetTask1(ctx, taskId).WithTemplate(withTemplate).Execute()

Get a task

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
    withTemplate := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HumanTaskAPI.GetTask1(context.Background(), taskId).WithTemplate(withTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.GetTask1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTask1`: HumanTaskEntry
    fmt.Fprintf(os.Stdout, "Response from `HumanTaskAPI.GetTask1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTask1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withTemplate** | **bool** |  | [default to false]

### Return type

[**HumanTaskEntry**](HumanTaskEntry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskDisplayNames

> []string GetTaskDisplayNames(ctx).SearchType(searchType).Execute()

Get list of task display names applicable for the user

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
    searchType := "searchType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HumanTaskAPI.GetTaskDisplayNames(context.Background()).SearchType(searchType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.GetTaskDisplayNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskDisplayNames`: []string
    fmt.Fprintf(os.Stdout, "Response from `HumanTaskAPI.GetTaskDisplayNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskDisplayNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchType** | **string** |  | 

### Return type

**[]string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateByNameAndVersion

> HumanTaskTemplate GetTemplateByNameAndVersion(ctx, name, version).Execute()

Get user form template by name and version

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
    name := "name_example" // string | 
    version := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HumanTaskAPI.GetTemplateByNameAndVersion(context.Background(), name, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.GetTemplateByNameAndVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateByNameAndVersion`: HumanTaskTemplate
    fmt.Fprintf(os.Stdout, "Response from `HumanTaskAPI.GetTemplateByNameAndVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**version** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateByNameAndVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HumanTaskTemplate**](HumanTaskTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateByTaskId

> HumanTaskTemplate GetTemplateByTaskId(ctx, humanTaskId).Execute()

Get user form by human task id

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
    humanTaskId := "humanTaskId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HumanTaskAPI.GetTemplateByTaskId(context.Background(), humanTaskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.GetTemplateByTaskId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateByTaskId`: HumanTaskTemplate
    fmt.Fprintf(os.Stdout, "Response from `HumanTaskAPI.GetTemplateByTaskId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**humanTaskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateByTaskIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HumanTaskTemplate**](HumanTaskTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReassignTask

> ReassignTask(ctx, taskId).HumanTaskAssignment(humanTaskAssignment).Execute()

Reassign a task without completing it

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
    humanTaskAssignment := []openapiclient.HumanTaskAssignment{*openapiclient.NewHumanTaskAssignment()} // []HumanTaskAssignment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HumanTaskAPI.ReassignTask(context.Background(), taskId).HumanTaskAssignment(humanTaskAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.ReassignTask``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReassignTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **humanTaskAssignment** | [**[]HumanTaskAssignment**](HumanTaskAssignment.md) |  | 

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


## ReleaseTask

> ReleaseTask(ctx, taskId).Execute()

Release a task without completing it

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
    r, err := apiClient.HumanTaskAPI.ReleaseTask(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.ReleaseTask``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReleaseTaskRequest struct via the builder pattern


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


## SaveTemplate

> HumanTaskTemplate SaveTemplate(ctx).HumanTaskTemplate(humanTaskTemplate).NewVersion(newVersion).Execute()

Save user form template

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
    humanTaskTemplate := *openapiclient.NewHumanTaskTemplate(map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "Name_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int32(123)) // HumanTaskTemplate | 
    newVersion := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HumanTaskAPI.SaveTemplate(context.Background()).HumanTaskTemplate(humanTaskTemplate).NewVersion(newVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.SaveTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveTemplate`: HumanTaskTemplate
    fmt.Fprintf(os.Stdout, "Response from `HumanTaskAPI.SaveTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **humanTaskTemplate** | [**HumanTaskTemplate**](HumanTaskTemplate.md) |  | 
 **newVersion** | **bool** |  | [default to false]

### Return type

[**HumanTaskTemplate**](HumanTaskTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveTemplates

> []HumanTaskTemplate SaveTemplates(ctx).HumanTaskTemplate(humanTaskTemplate).NewVersion(newVersion).Execute()

Save user form template

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
    humanTaskTemplate := []openapiclient.HumanTaskTemplate{*openapiclient.NewHumanTaskTemplate(map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "Name_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int32(123))} // []HumanTaskTemplate | 
    newVersion := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HumanTaskAPI.SaveTemplates(context.Background()).HumanTaskTemplate(humanTaskTemplate).NewVersion(newVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.SaveTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveTemplates`: []HumanTaskTemplate
    fmt.Fprintf(os.Stdout, "Response from `HumanTaskAPI.SaveTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **humanTaskTemplate** | [**[]HumanTaskTemplate**](HumanTaskTemplate.md) |  | 
 **newVersion** | **bool** |  | [default to false]

### Return type

[**[]HumanTaskTemplate**](HumanTaskTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> HumanTaskSearchResult Search(ctx).HumanTaskSearch(humanTaskSearch).Execute()

Search human tasks

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
    humanTaskSearch := *openapiclient.NewHumanTaskSearch() // HumanTaskSearch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HumanTaskAPI.Search(context.Background()).HumanTaskSearch(humanTaskSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: HumanTaskSearchResult
    fmt.Fprintf(os.Stdout, "Response from `HumanTaskAPI.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **humanTaskSearch** | [**HumanTaskSearch**](HumanTaskSearch.md) |  | 

### Return type

[**HumanTaskSearchResult**](HumanTaskSearchResult.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkipTask

> SkipTask(ctx, taskId).Reason(reason).Execute()

If a task is assigned to a user, this API can be used to skip that assignment and move to the next assignee

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
    reason := "reason_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HumanTaskAPI.SkipTask(context.Background(), taskId).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.SkipTask``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSkipTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** |  | 

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


## UpdateTaskOutput

> UpdateTaskOutput(ctx, taskId).RequestBody(requestBody).Complete(complete).Execute()

Update task output, optionally complete

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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 
    complete := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HumanTaskAPI.UpdateTaskOutput(context.Background(), taskId).RequestBody(requestBody).Complete(complete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.UpdateTaskOutput``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateTaskOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** |  | 
 **complete** | **bool** |  | [default to false]

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


## UpdateTaskOutputByRef

> UpdateTaskOutputByRef(ctx).WorkflowId(workflowId).TaskRefName(taskRefName).RequestBody(requestBody).Complete(complete).Iteration(iteration).Execute()

Update task output, optionally complete

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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 
    complete := true // bool |  (optional) (default to false)
    iteration := []int32{int32(123)} // []int32 | Populate this value if your task is in a loop and you want to update a specific iteration. If its not in a loop OR if you want to just update the latest iteration, leave this as empty (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HumanTaskAPI.UpdateTaskOutputByRef(context.Background()).WorkflowId(workflowId).TaskRefName(taskRefName).RequestBody(requestBody).Complete(complete).Iteration(iteration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HumanTaskAPI.UpdateTaskOutputByRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskOutputByRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowId** | **string** |  | 
 **taskRefName** | **string** |  | 
 **requestBody** | **map[string]map[string]interface{}** |  | 
 **complete** | **bool** |  | [default to false]
 **iteration** | **[]int32** | Populate this value if your task is in a loop and you want to update a specific iteration. If its not in a loop OR if you want to just update the latest iteration, leave this as empty | 

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

