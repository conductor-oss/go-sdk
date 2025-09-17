# \MetadataResourceAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](MetadataResourceAPI.md#Create) | **Post** /metadata/workflow | Create a new workflow definition
[**Get1**](MetadataResourceAPI.md#Get1) | **Get** /metadata/workflow/{name} | Retrieves workflow definition along with blueprint
[**GetTaskDef**](MetadataResourceAPI.md#GetTaskDef) | **Get** /metadata/taskdefs/{tasktype} | Gets the task definition
[**GetTaskDefs**](MetadataResourceAPI.md#GetTaskDefs) | **Get** /metadata/taskdefs | Gets all task definition
[**GetWorkflowDefs**](MetadataResourceAPI.md#GetWorkflowDefs) | **Get** /metadata/workflow | Retrieves all workflow definition along with blueprint
[**RegisterTaskDef**](MetadataResourceAPI.md#RegisterTaskDef) | **Post** /metadata/taskdefs | Create or update task definition(s)
[**UnregisterTaskDef**](MetadataResourceAPI.md#UnregisterTaskDef) | **Delete** /metadata/taskdefs/{tasktype} | Remove a task definition
[**UnregisterWorkflowDef**](MetadataResourceAPI.md#UnregisterWorkflowDef) | **Delete** /metadata/workflow/{name}/{version} | Removes workflow definition. It does not remove workflows associated with the definition.
[**Update**](MetadataResourceAPI.md#Update) | **Put** /metadata/workflow | Create or update workflow definition(s)
[**UpdateTaskDef**](MetadataResourceAPI.md#UpdateTaskDef) | **Put** /metadata/taskdefs | Update an existing task
[**UploadBpmnFile**](MetadataResourceAPI.md#UploadBpmnFile) | **Post** /metadata/workflow-importer/import-bpm | Imports bpmn workflow
[**UploadWorkflowsAndTasksDefinitionsToS3**](MetadataResourceAPI.md#UploadWorkflowsAndTasksDefinitionsToS3) | **Post** /metadata/workflow-task-defs/upload | Upload all workflows and tasks definitions to Object storage if configured



## Create

> map[string]interface{} Create(ctx).ExtendedWorkflowDef(extendedWorkflowDef).Overwrite(overwrite).NewVersion(newVersion).Execute()

Create a new workflow definition

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
    extendedWorkflowDef := *openapiclient.NewExtendedWorkflowDef("Name_example", []openapiclient.WorkflowTask{*openapiclient.NewWorkflowTask("Name_example", "TaskReferenceName_example")}, int64(123)) // ExtendedWorkflowDef | 
    overwrite := true // bool |  (optional) (default to false)
    newVersion := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataResourceAPI.Create(context.Background()).ExtendedWorkflowDef(extendedWorkflowDef).Overwrite(overwrite).NewVersion(newVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extendedWorkflowDef** | [**ExtendedWorkflowDef**](ExtendedWorkflowDef.md) |  | 
 **overwrite** | **bool** |  | [default to false]
 **newVersion** | **bool** |  | [default to false]

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


## Get1

> WorkflowDef Get1(ctx, name).Version(version).Metadata(metadata).Execute()

Retrieves workflow definition along with blueprint

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
    version := int32(56) // int32 |  (optional)
    metadata := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataResourceAPI.Get1(context.Background(), name).Version(version).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.Get1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get1`: WorkflowDef
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.Get1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** |  | 
 **metadata** | **bool** |  | [default to false]

### Return type

[**WorkflowDef**](WorkflowDef.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskDef

> map[string]interface{} GetTaskDef(ctx, tasktype).Metadata(metadata).Execute()

Gets the task definition

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
    metadata := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataResourceAPI.GetTaskDef(context.Background(), tasktype).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.GetTaskDef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskDef`: map[string]interface{}
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

 **metadata** | **bool** |  | [default to false]

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskDefs

> []TaskDef GetTaskDefs(ctx).Access(access).Metadata(metadata).TagKey(tagKey).TagValue(tagValue).Execute()

Gets all task definition

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
    access := "access_example" // string |  (optional) (default to "READ")
    metadata := true // bool |  (optional) (default to false)
    tagKey := "tagKey_example" // string |  (optional)
    tagValue := "tagValue_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataResourceAPI.GetTaskDefs(context.Background()).Access(access).Metadata(metadata).TagKey(tagKey).TagValue(tagValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.GetTaskDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskDefs`: []TaskDef
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.GetTaskDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **access** | **string** |  | [default to &quot;READ&quot;]
 **metadata** | **bool** |  | [default to false]
 **tagKey** | **string** |  | 
 **tagValue** | **string** |  | 

### Return type

[**[]TaskDef**](TaskDef.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowDefs

> []WorkflowDef GetWorkflowDefs(ctx).Access(access).Metadata(metadata).TagKey(tagKey).TagValue(tagValue).Name(name).Short(short).Execute()

Retrieves all workflow definition along with blueprint

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
    access := "access_example" // string |  (optional) (default to "READ")
    metadata := true // bool |  (optional) (default to false)
    tagKey := "tagKey_example" // string |  (optional)
    tagValue := "tagValue_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    short := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataResourceAPI.GetWorkflowDefs(context.Background()).Access(access).Metadata(metadata).TagKey(tagKey).TagValue(tagValue).Name(name).Short(short).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.GetWorkflowDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowDefs`: []WorkflowDef
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.GetWorkflowDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **access** | **string** |  | [default to &quot;READ&quot;]
 **metadata** | **bool** |  | [default to false]
 **tagKey** | **string** |  | 
 **tagValue** | **string** |  | 
 **name** | **string** |  | 
 **short** | **bool** |  | [default to false]

### Return type

[**[]WorkflowDef**](WorkflowDef.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterTaskDef

> map[string]interface{} RegisterTaskDef(ctx).ExtendedTaskDef(extendedTaskDef).Execute()

Create or update task definition(s)

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
    extendedTaskDef := []openapiclient.ExtendedTaskDef{*openapiclient.NewExtendedTaskDef("Name_example", int64(123), int64(123))} // []ExtendedTaskDef | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataResourceAPI.RegisterTaskDef(context.Background()).ExtendedTaskDef(extendedTaskDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.RegisterTaskDef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterTaskDef`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.RegisterTaskDef`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterTaskDefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extendedTaskDef** | [**[]ExtendedTaskDef**](ExtendedTaskDef.md) |  | 

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
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

[api_key](../README.md#api_key)

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
    openapiclient "github.com/conductor-sdk/conductor-go/sdk/generated/http/orkescd"
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

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> map[string]interface{} Update(ctx).ExtendedWorkflowDef(extendedWorkflowDef).Overwrite(overwrite).NewVersion(newVersion).Execute()

Create or update workflow definition(s)

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
    extendedWorkflowDef := []openapiclient.ExtendedWorkflowDef{*openapiclient.NewExtendedWorkflowDef("Name_example", []openapiclient.WorkflowTask{*openapiclient.NewWorkflowTask("Name_example", "TaskReferenceName_example")}, int64(123))} // []ExtendedWorkflowDef | 
    overwrite := true // bool |  (optional) (default to true)
    newVersion := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataResourceAPI.Update(context.Background()).ExtendedWorkflowDef(extendedWorkflowDef).Overwrite(overwrite).NewVersion(newVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.Update`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extendedWorkflowDef** | [**[]ExtendedWorkflowDef**](ExtendedWorkflowDef.md) |  | 
 **overwrite** | **bool** |  | [default to true]
 **newVersion** | **bool** |  | [default to false]

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


## UpdateTaskDef

> map[string]interface{} UpdateTaskDef(ctx).ExtendedTaskDef(extendedTaskDef).Execute()

Update an existing task

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
    extendedTaskDef := *openapiclient.NewExtendedTaskDef("Name_example", int64(123), int64(123)) // ExtendedTaskDef | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataResourceAPI.UpdateTaskDef(context.Background()).ExtendedTaskDef(extendedTaskDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.UpdateTaskDef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTaskDef`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.UpdateTaskDef`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskDefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extendedTaskDef** | [**ExtendedTaskDef**](ExtendedTaskDef.md) |  | 

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


## UploadBpmnFile

> []ExtendedWorkflowDef UploadBpmnFile(ctx).IncomingBpmnFile(incomingBpmnFile).Overwrite(overwrite).Execute()

Imports bpmn workflow

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
    incomingBpmnFile := *openapiclient.NewIncomingBpmnFile("FileContent_example", "FileName_example") // IncomingBpmnFile | 
    overwrite := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataResourceAPI.UploadBpmnFile(context.Background()).IncomingBpmnFile(incomingBpmnFile).Overwrite(overwrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.UploadBpmnFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBpmnFile`: []ExtendedWorkflowDef
    fmt.Fprintf(os.Stdout, "Response from `MetadataResourceAPI.UploadBpmnFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadBpmnFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incomingBpmnFile** | [**IncomingBpmnFile**](IncomingBpmnFile.md) |  | 
 **overwrite** | **bool** |  | [default to true]

### Return type

[**[]ExtendedWorkflowDef**](ExtendedWorkflowDef.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadWorkflowsAndTasksDefinitionsToS3

> UploadWorkflowsAndTasksDefinitionsToS3(ctx).Execute()

Upload all workflows and tasks definitions to Object storage if configured

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
    r, err := apiClient.MetadataResourceAPI.UploadWorkflowsAndTasksDefinitionsToS3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataResourceAPI.UploadWorkflowsAndTasksDefinitionsToS3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUploadWorkflowsAndTasksDefinitionsToS3Request struct via the builder pattern


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

