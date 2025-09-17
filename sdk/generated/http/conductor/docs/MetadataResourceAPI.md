# \MetadataResourceAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](MetadataResourceAPI.md#Create) | **Post** /metadata/workflow | Create a new workflow definition
[**Get**](MetadataResourceAPI.md#Get) | **Get** /metadata/workflow/{name} | Retrieves workflow definition along with blueprint
[**GetAll**](MetadataResourceAPI.md#GetAll) | **Get** /metadata/workflow | Retrieves all workflow definition along with blueprint
[**GetAllWorkflowsWithLatestVersions**](MetadataResourceAPI.md#GetAllWorkflowsWithLatestVersions) | **Get** /metadata/workflow/latest-versions | Returns only the latest version of all workflow definitions
[**GetTaskDef**](MetadataResourceAPI.md#GetTaskDef) | **Get** /metadata/taskdefs/{tasktype} | Gets the task definition
[**GetTaskDefs**](MetadataResourceAPI.md#GetTaskDefs) | **Get** /metadata/taskdefs | Gets all task definition
[**GetWorkflowNamesAndVersions**](MetadataResourceAPI.md#GetWorkflowNamesAndVersions) | **Get** /metadata/workflow/names-and-versions | Returns workflow names and versions only (no definition bodies)
[**RegisterTaskDef**](MetadataResourceAPI.md#RegisterTaskDef) | **Put** /metadata/taskdefs | Update an existing task
[**RegisterTaskDef1**](MetadataResourceAPI.md#RegisterTaskDef1) | **Post** /metadata/taskdefs | Create new task definition(s)
[**UnregisterTaskDef**](MetadataResourceAPI.md#UnregisterTaskDef) | **Delete** /metadata/taskdefs/{tasktype} | Remove a task definition
[**UnregisterWorkflowDef**](MetadataResourceAPI.md#UnregisterWorkflowDef) | **Delete** /metadata/workflow/{name}/{version} | Removes workflow definition. It does not remove workflows associated with the definition.
[**Update**](MetadataResourceAPI.md#Update) | **Put** /metadata/workflow | Create or update workflow definition
[**Validate**](MetadataResourceAPI.md#Validate) | **Post** /metadata/workflow/validate | Validates a new workflow definition



## Create

> Create(ctx).WorkflowDef(workflowDef).Execute()

Create a new workflow definition

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
    workflowDef := *openapiclient.NewWorkflowDef("Name_example", []openapiclient.WorkflowTask{*openapiclient.NewWorkflowTask("Name_example", "TaskReferenceName_example")}, int64(123)) // WorkflowDef | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetadataResourceAPI.Create(context.Background()).WorkflowDef(workflowDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowDef** | [**WorkflowDef**](WorkflowDef.md) |  | 

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


## Get

> WorkflowDef Get(ctx, name).Version(version).Execute()

Retrieves workflow definition along with blueprint

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
    version := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataResourceAPI.Get(context.Background(), name).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: WorkflowDef
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** |  | 

### Return type

[**WorkflowDef**](WorkflowDef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> []WorkflowDef GetAll(ctx).Execute()

Retrieves all workflow definition along with blueprint

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
    resp, r, err := apiClient.MetadataResourceAPI.GetAll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.GetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAll`: []WorkflowDef
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.GetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


### Return type

[**[]WorkflowDef**](WorkflowDef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllWorkflowsWithLatestVersions

> []WorkflowDef GetAllWorkflowsWithLatestVersions(ctx).Execute()

Returns only the latest version of all workflow definitions

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
    resp, r, err := apiClient.MetadataResourceAPI.GetAllWorkflowsWithLatestVersions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.GetAllWorkflowsWithLatestVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllWorkflowsWithLatestVersions`: []WorkflowDef
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.GetAllWorkflowsWithLatestVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllWorkflowsWithLatestVersionsRequest struct via the builder pattern


### Return type

[**[]WorkflowDef**](WorkflowDef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskDef

> TaskDef GetTaskDef(ctx, tasktype).Execute()

Gets the task definition

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataResourceAPI.GetTaskDef(context.Background(), tasktype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.GetTaskDef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskDef`: TaskDef
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.GetTaskDef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tasktype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskDefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskDef**](TaskDef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskDefs

> []TaskDef GetTaskDefs(ctx).Execute()

Gets all task definition

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
    resp, r, err := apiClient.MetadataResourceAPI.GetTaskDefs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.GetTaskDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskDefs`: []TaskDef
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.GetTaskDefs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskDefsRequest struct via the builder pattern


### Return type

[**[]TaskDef**](TaskDef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowNamesAndVersions

> map[string]map[string]interface{} GetWorkflowNamesAndVersions(ctx).Execute()

Returns workflow names and versions only (no definition bodies)

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
    resp, r, err := apiClient.MetadataResourceAPI.GetWorkflowNamesAndVersions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.GetWorkflowNamesAndVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowNamesAndVersions`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.GetWorkflowNamesAndVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowNamesAndVersionsRequest struct via the builder pattern


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


## RegisterTaskDef

> RegisterTaskDef(ctx).TaskDef(taskDef).Execute()

Update an existing task

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
    taskDef := *openapiclient.NewTaskDef("Name_example", int64(123), int64(123)) // TaskDef | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetadataResourceAPI.RegisterTaskDef(context.Background()).TaskDef(taskDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.RegisterTaskDef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterTaskDefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskDef** | [**TaskDef**](TaskDef.md) |  | 

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


## RegisterTaskDef1

> RegisterTaskDef1(ctx).TaskDef(taskDef).Execute()

Create new task definition(s)

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
    taskDef := []openapiclient.TaskDef{*openapiclient.NewTaskDef("Name_example", int64(123), int64(123))} // []TaskDef | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetadataResourceAPI.RegisterTaskDef1(context.Background()).TaskDef(taskDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.RegisterTaskDef1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterTaskDef1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskDef** | [**[]TaskDef**](TaskDef.md) |  | 

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


## UnregisterTaskDef

> UnregisterTaskDef(ctx, tasktype).Execute()

Remove a task definition

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetadataResourceAPI.UnregisterTaskDef(context.Background(), tasktype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.UnregisterTaskDef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tasktype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterTaskDefRequest struct via the builder pattern


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


## UnregisterWorkflowDef

> UnregisterWorkflowDef(ctx, name, version).Execute()

Removes workflow definition. It does not remove workflows associated with the definition.

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
    version := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetadataResourceAPI.UnregisterWorkflowDef(context.Background(), name, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.UnregisterWorkflowDef``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnregisterWorkflowDefRequest struct via the builder pattern


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


## Update

> BulkResponseString Update(ctx).WorkflowDef(workflowDef).Execute()

Create or update workflow definition

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
    workflowDef := []openapiclient.WorkflowDef{*openapiclient.NewWorkflowDef("Name_example", []openapiclient.WorkflowTask{*openapiclient.NewWorkflowTask("Name_example", "TaskReferenceName_example")}, int64(123))} // []WorkflowDef | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataResourceAPI.Update(context.Background()).WorkflowDef(workflowDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: BulkResponseString
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.Update`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowDef** | [**[]WorkflowDef**](WorkflowDef.md) |  | 

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


## Validate

> Validate(ctx).WorkflowDef(workflowDef).Execute()

Validates a new workflow definition

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
    workflowDef := *openapiclient.NewWorkflowDef("Name_example", []openapiclient.WorkflowTask{*openapiclient.NewWorkflowTask("Name_example", "TaskReferenceName_example")}, int64(123)) // WorkflowDef | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetadataResourceAPI.Validate(context.Background()).WorkflowDef(workflowDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.Validate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowDef** | [**WorkflowDef**](WorkflowDef.md) |  | 

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

