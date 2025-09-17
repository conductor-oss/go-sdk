# \PromptResourceAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessageTemplates**](PromptResourceAPI.md#CreateMessageTemplates) | **Post** /prompts/ | Create message templates in bulk
[**DeleteMessageTemplate**](PromptResourceAPI.md#DeleteMessageTemplate) | **Delete** /prompts/{name} | Delete Template
[**DeleteMessageTemplate1**](PromptResourceAPI.md#DeleteMessageTemplate1) | **Delete** /prompts/{name}/versions/{version} | Delete Template
[**DeleteTagForPromptTemplate**](PromptResourceAPI.md#DeleteTagForPromptTemplate) | **Delete** /prompts/{name}/tags | Delete a tag for Prompt Template
[**GetMessageTemplate**](PromptResourceAPI.md#GetMessageTemplate) | **Get** /prompts/{name} | Get Template
[**GetMessageTemplateVersions**](PromptResourceAPI.md#GetMessageTemplateVersions) | **Get** /prompts/{name}/versions | Get all versions of a Template
[**GetMessageTemplates**](PromptResourceAPI.md#GetMessageTemplates) | **Get** /prompts | Get Templates
[**GetTagsForPromptTemplate**](PromptResourceAPI.md#GetTagsForPromptTemplate) | **Get** /prompts/{name}/tags | Get tags by Prompt Template
[**PutTagForPromptTemplate**](PromptResourceAPI.md#PutTagForPromptTemplate) | **Put** /prompts/{name}/tags | Put a tag to Prompt Template
[**SaveMessageTemplate**](PromptResourceAPI.md#SaveMessageTemplate) | **Post** /prompts/{name} | Create or Update a template
[**TestMessageTemplate**](PromptResourceAPI.md#TestMessageTemplate) | **Post** /prompts/test | Test Prompt Template
[**UpdateMessageTemplate**](PromptResourceAPI.md#UpdateMessageTemplate) | **Put** /prompts/{name} | Create a template



## CreateMessageTemplates

> CreateMessageTemplates(ctx).MessageTemplate(messageTemplate).Execute()

Create message templates in bulk

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
    messageTemplate := []openapiclient.MessageTemplate{*openapiclient.NewMessageTemplate()} // []MessageTemplate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PromptResourceAPI.CreateMessageTemplates(context.Background()).MessageTemplate(messageTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromptResourceAPI.CreateMessageTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageTemplate** | [**[]MessageTemplate**](MessageTemplate.md) |  | 

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


## DeleteMessageTemplate

> DeleteMessageTemplate(ctx, name).Execute()

Delete Template

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
    r, err := apiClient.PromptResourceAPI.DeleteMessageTemplate(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromptResourceAPI.DeleteMessageTemplate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMessageTemplateRequest struct via the builder pattern


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


## DeleteMessageTemplate1

> DeleteMessageTemplate1(ctx, name, version).Execute()

Delete Template

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
    r, err := apiClient.PromptResourceAPI.DeleteMessageTemplate1(context.Background(), name, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromptResourceAPI.DeleteMessageTemplate1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMessageTemplate1Request struct via the builder pattern


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


## DeleteTagForPromptTemplate

> DeleteTagForPromptTemplate(ctx, name).Tag(tag).Execute()

Delete a tag for Prompt Template

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
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PromptResourceAPI.DeleteTagForPromptTemplate(context.Background(), name).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromptResourceAPI.DeleteTagForPromptTemplate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTagForPromptTemplateRequest struct via the builder pattern


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


## GetMessageTemplate

> MessageTemplate GetMessageTemplate(ctx, name).Version(version).Execute()

Get Template

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromptResourceAPI.GetMessageTemplate(context.Background(), name).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromptResourceAPI.GetMessageTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageTemplate`: MessageTemplate
    fmt.Fprintf(os.Stdout, "Response from `PromptResourceAPI.GetMessageTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** |  | 

### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageTemplateVersions

> []MessageTemplate GetMessageTemplateVersions(ctx, name).Execute()

Get all versions of a Template

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
    resp, r, err := apiClient.PromptResourceAPI.GetMessageTemplateVersions(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromptResourceAPI.GetMessageTemplateVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageTemplateVersions`: []MessageTemplate
    fmt.Fprintf(os.Stdout, "Response from `PromptResourceAPI.GetMessageTemplateVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageTemplateVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MessageTemplate**](MessageTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageTemplates

> []MessageTemplate GetMessageTemplates(ctx).Execute()

Get Templates

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
    resp, r, err := apiClient.PromptResourceAPI.GetMessageTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromptResourceAPI.GetMessageTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageTemplates`: []MessageTemplate
    fmt.Fprintf(os.Stdout, "Response from `PromptResourceAPI.GetMessageTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageTemplatesRequest struct via the builder pattern


### Return type

[**[]MessageTemplate**](MessageTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagsForPromptTemplate

> []Tag GetTagsForPromptTemplate(ctx, name).Execute()

Get tags by Prompt Template

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
    resp, r, err := apiClient.PromptResourceAPI.GetTagsForPromptTemplate(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromptResourceAPI.GetTagsForPromptTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagsForPromptTemplate`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `PromptResourceAPI.GetTagsForPromptTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsForPromptTemplateRequest struct via the builder pattern


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


## PutTagForPromptTemplate

> PutTagForPromptTemplate(ctx, name).Tag(tag).Execute()

Put a tag to Prompt Template

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
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PromptResourceAPI.PutTagForPromptTemplate(context.Background(), name).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromptResourceAPI.PutTagForPromptTemplate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutTagForPromptTemplateRequest struct via the builder pattern


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


## SaveMessageTemplate

> SaveMessageTemplate(ctx, name).Description(description).Body(body).Models(models).Version(version).AutoIncrement(autoIncrement).Execute()

Create or Update a template

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
    description := "description_example" // string | 
    body := "body_example" // string | 
    models := []string{"Inner_example"} // []string |  (optional)
    version := int32(56) // int32 |  (optional)
    autoIncrement := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PromptResourceAPI.SaveMessageTemplate(context.Background(), name).Description(description).Body(body).Models(models).Version(version).AutoIncrement(autoIncrement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromptResourceAPI.SaveMessageTemplate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSaveMessageTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **string** |  | 
 **body** | **string** |  | 
 **models** | **[]string** |  | 
 **version** | **int32** |  | 
 **autoIncrement** | **bool** |  | [default to false]

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


## TestMessageTemplate

> string TestMessageTemplate(ctx).PromptTemplateTestRequest(promptTemplateTestRequest).Execute()

Test Prompt Template

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
    promptTemplateTestRequest := *openapiclient.NewPromptTemplateTestRequest() // PromptTemplateTestRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromptResourceAPI.TestMessageTemplate(context.Background()).PromptTemplateTestRequest(promptTemplateTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromptResourceAPI.TestMessageTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestMessageTemplate`: string
    fmt.Fprintf(os.Stdout, "Response from `PromptResourceAPI.TestMessageTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestMessageTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promptTemplateTestRequest** | [**PromptTemplateTestRequest**](PromptTemplateTestRequest.md) |  | 

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


## UpdateMessageTemplate

> UpdateMessageTemplate(ctx, name).Description(description).Body(body).Version(version).Models(models).Execute()

Create a template

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
    description := "description_example" // string | 
    body := "body_example" // string | 
    version := int32(56) // int32 |  (optional)
    models := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PromptResourceAPI.UpdateMessageTemplate(context.Background(), name).Description(description).Body(body).Version(version).Models(models).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromptResourceAPI.UpdateMessageTemplate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateMessageTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **string** |  | 
 **body** | **string** |  | 
 **version** | **int32** |  | 
 **models** | **[]string** |  | 

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

