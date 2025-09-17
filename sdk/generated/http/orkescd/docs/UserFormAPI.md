# \UserFormAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTemplateByName**](UserFormAPI.md#DeleteTemplateByName) | **Delete** /human/template/{name} | Delete all versions of user form template by name
[**GetAllTemplates**](UserFormAPI.md#GetAllTemplates) | **Get** /human/template | List all user form templates or get templates by name, or a template by name and version
[**GetTemplateByNameAndVersion**](UserFormAPI.md#GetTemplateByNameAndVersion) | **Get** /human/template/{name}/{version} | Get user form template by name and version
[**GetTemplateByTaskId**](UserFormAPI.md#GetTemplateByTaskId) | **Get** /human/template/{humanTaskId} | Get user form by human task id
[**SaveTemplate**](UserFormAPI.md#SaveTemplate) | **Post** /human/template | Save user form template
[**SaveTemplates**](UserFormAPI.md#SaveTemplates) | **Post** /human/template/bulk | Save user form template



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
    r, err := apiClient.UserFormAPI.DeleteTemplateByName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFormAPI.DeleteTemplateByName``: %v\n", err)
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
    resp, r, err := apiClient.UserFormAPI.GetAllTemplates(context.Background()).Name(name).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFormAPI.GetAllTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTemplates`: []HumanTaskTemplate
    fmt.Fprintf(os.Stdout, "Response from `UserFormAPI.GetAllTemplates`: %v\n", resp)
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
    resp, r, err := apiClient.UserFormAPI.GetTemplateByNameAndVersion(context.Background(), name, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFormAPI.GetTemplateByNameAndVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateByNameAndVersion`: HumanTaskTemplate
    fmt.Fprintf(os.Stdout, "Response from `UserFormAPI.GetTemplateByNameAndVersion`: %v\n", resp)
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
    resp, r, err := apiClient.UserFormAPI.GetTemplateByTaskId(context.Background(), humanTaskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFormAPI.GetTemplateByTaskId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateByTaskId`: HumanTaskTemplate
    fmt.Fprintf(os.Stdout, "Response from `UserFormAPI.GetTemplateByTaskId`: %v\n", resp)
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
    resp, r, err := apiClient.UserFormAPI.SaveTemplate(context.Background()).HumanTaskTemplate(humanTaskTemplate).NewVersion(newVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFormAPI.SaveTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveTemplate`: HumanTaskTemplate
    fmt.Fprintf(os.Stdout, "Response from `UserFormAPI.SaveTemplate`: %v\n", resp)
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
    resp, r, err := apiClient.UserFormAPI.SaveTemplates(context.Background()).HumanTaskTemplate(humanTaskTemplate).NewVersion(newVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFormAPI.SaveTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveTemplates`: []HumanTaskTemplate
    fmt.Fprintf(os.Stdout, "Response from `UserFormAPI.SaveTemplates`: %v\n", resp)
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

