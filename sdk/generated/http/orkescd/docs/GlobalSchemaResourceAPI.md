# \GlobalSchemaResourceAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSchemaByNameWithLatestVersion1**](GlobalSchemaResourceAPI.md#GetSchemaByNameWithLatestVersion1) | **Get** /global_schema/{name} | Get schema by name with latest version
[**Save1**](GlobalSchemaResourceAPI.md#Save1) | **Post** /global_schema | Save schema



## GetSchemaByNameWithLatestVersion1

> SchemaDef GetSchemaByNameWithLatestVersion1(ctx, name).Execute()

Get schema by name with latest version

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
    resp, r, err := apiClient.GlobalSchemaResourceAPI.GetSchemaByNameWithLatestVersion1(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalSchemaResourceAPI.GetSchemaByNameWithLatestVersion1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaByNameWithLatestVersion1`: SchemaDef
    fmt.Fprintf(os.Stdout, "Response from `GlobalSchemaResourceAPI.GetSchemaByNameWithLatestVersion1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaByNameWithLatestVersion1Request struct via the builder pattern


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


## Save1

> Save1(ctx).SchemaDef(schemaDef).NewVersion(newVersion).Execute()

Save schema

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
    schemaDef := []openapiclient.SchemaDef{*openapiclient.NewSchemaDef("Name_example", "Type_example", int32(123))} // []SchemaDef | 
    newVersion := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GlobalSchemaResourceAPI.Save1(context.Background()).SchemaDef(schemaDef).NewVersion(newVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalSchemaResourceAPI.Save1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSave1Request struct via the builder pattern


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

