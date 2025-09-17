# \SchemaResourceAPI

All URIs are relative to *https://siliconmint-dev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSchemaByName**](SchemaResourceAPI.md#DeleteSchemaByName) | **Delete** /schema/{name} | Delete all versions of schema by name
[**DeleteSchemaByNameAndVersion**](SchemaResourceAPI.md#DeleteSchemaByNameAndVersion) | **Delete** /schema/{name}/{version} | Delete a version of schema by name
[**GetAllSchemas**](SchemaResourceAPI.md#GetAllSchemas) | **Get** /schema | Get all schemas
[**GetSchemaByNameAndVersion**](SchemaResourceAPI.md#GetSchemaByNameAndVersion) | **Get** /schema/{name}/{version} | Get schema by name and version
[**Save**](SchemaResourceAPI.md#Save) | **Post** /schema | Save schema



## DeleteSchemaByName

> DeleteSchemaByName(ctx, name).Execute()

Delete all versions of schema by name

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SchemaResourceAPI.DeleteSchemaByName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaResourceAPI.DeleteSchemaByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSchemaByNameRequest struct via the builder pattern


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


## DeleteSchemaByNameAndVersion

> DeleteSchemaByNameAndVersion(ctx, name, version).Execute()

Delete a version of schema by name

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SchemaResourceAPI.DeleteSchemaByNameAndVersion(context.Background(), name, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaResourceAPI.DeleteSchemaByNameAndVersion``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSchemaByNameAndVersionRequest struct via the builder pattern


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


## GetAllSchemas

> []SchemaDef GetAllSchemas(ctx).Execute()

Get all schemas

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaResourceAPI.GetAllSchemas(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaResourceAPI.GetAllSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSchemas`: []SchemaDef
    fmt.Fprintf(os.Stdout, "Response from `SchemaResourceAPI.GetAllSchemas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSchemasRequest struct via the builder pattern


### Return type

[**[]SchemaDef**](SchemaDef.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaByNameAndVersion

> SchemaDef GetSchemaByNameAndVersion(ctx, name, version).Execute()

Get schema by name and version

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemaResourceAPI.GetSchemaByNameAndVersion(context.Background(), name, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaResourceAPI.GetSchemaByNameAndVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaByNameAndVersion`: SchemaDef
    fmt.Fprintf(os.Stdout, "Response from `SchemaResourceAPI.GetSchemaByNameAndVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**version** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaByNameAndVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SchemaDef**](SchemaDef.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Save

> Save(ctx).SchemaDef(schemaDef).NewVersion(newVersion).Execute()

Save schema

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
    schemaDef := []openapiclient.SchemaDef{*openapiclient.NewSchemaDef("Name_example", "Type_example", int32(123))} // []SchemaDef | 
    newVersion := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SchemaResourceAPI.Save(context.Background()).SchemaDef(schemaDef).NewVersion(newVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaResourceAPI.Save``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaDef** | [**[]SchemaDef**](SchemaDef.md) |  | 
 **newVersion** | **bool** |  | [default to false]

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

