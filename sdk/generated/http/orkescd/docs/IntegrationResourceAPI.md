# \IntegrationResourceAPI

All URIs are relative to *https://sdkdev.orkesconductor.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociatePromptWithIntegration**](IntegrationResourceAPI.md#AssociatePromptWithIntegration) | **Post** /integrations/provider/{integration_provider}/integration/{integration_name}/prompt/{prompt_name} | Associate a Prompt Template with an Integration
[**DeleteIntegrationApi**](IntegrationResourceAPI.md#DeleteIntegrationApi) | **Delete** /integrations/provider/{name}/integration/{integration_name} | Delete an Integration
[**DeleteIntegrationProvider**](IntegrationResourceAPI.md#DeleteIntegrationProvider) | **Delete** /integrations/provider/{name} | Delete an Integration Provider
[**DeleteTagForIntegration**](IntegrationResourceAPI.md#DeleteTagForIntegration) | **Delete** /integrations/provider/{name}/integration/{integration_name}/tags | Delete a tag for Integration
[**DeleteTagForIntegrationProvider**](IntegrationResourceAPI.md#DeleteTagForIntegrationProvider) | **Delete** /integrations/provider/{name}/tags | Delete a tag for Integration Provider
[**GetAllIntegrations**](IntegrationResourceAPI.md#GetAllIntegrations) | **Get** /integrations/ | Get all Integrations
[**GetIntegrationApi**](IntegrationResourceAPI.md#GetIntegrationApi) | **Get** /integrations/provider/{name}/integration/{integration_name} | Get Integration details
[**GetIntegrationApis**](IntegrationResourceAPI.md#GetIntegrationApis) | **Get** /integrations/provider/{name}/integration | Get Integrations of an Integration Provider
[**GetIntegrationAvailableApis**](IntegrationResourceAPI.md#GetIntegrationAvailableApis) | **Get** /integrations/provider/{name}/integration/all | Get Integrations Available for an Integration Provider
[**GetIntegrationDef**](IntegrationResourceAPI.md#GetIntegrationDef) | **Get** /integrations/def/{name} | Get an integration definition
[**GetIntegrationProvider**](IntegrationResourceAPI.md#GetIntegrationProvider) | **Get** /integrations/provider/{name} | Get Integration provider
[**GetIntegrationProviderDefs**](IntegrationResourceAPI.md#GetIntegrationProviderDefs) | **Get** /integrations/def | Get Integration provider definitions
[**GetIntegrationProviders**](IntegrationResourceAPI.md#GetIntegrationProviders) | **Get** /integrations/provider | Get all Integrations Providers
[**GetPromptsWithIntegration**](IntegrationResourceAPI.md#GetPromptsWithIntegration) | **Get** /integrations/provider/{integration_provider}/integration/{integration_name}/prompt | Get the list of prompt templates associated with an integration
[**GetProvidersAndIntegrations**](IntegrationResourceAPI.md#GetProvidersAndIntegrations) | **Get** /integrations/all | Get Integrations Providers and Integrations combo
[**GetTagsForIntegration**](IntegrationResourceAPI.md#GetTagsForIntegration) | **Get** /integrations/provider/{name}/integration/{integration_name}/tags | Get tags by Integration
[**GetTagsForIntegrationProvider**](IntegrationResourceAPI.md#GetTagsForIntegrationProvider) | **Get** /integrations/provider/{name}/tags | Get tags by Integration Provider
[**PutTagForIntegration**](IntegrationResourceAPI.md#PutTagForIntegration) | **Put** /integrations/provider/{name}/integration/{integration_name}/tags | Put a tag to Integration
[**PutTagForIntegrationProvider**](IntegrationResourceAPI.md#PutTagForIntegrationProvider) | **Put** /integrations/provider/{name}/tags | Put a tag to Integration Provider
[**RecordEventStats**](IntegrationResourceAPI.md#RecordEventStats) | **Post** /integrations/eventStats/{type} | Record Event Stats
[**RegisterIntegration**](IntegrationResourceAPI.md#RegisterIntegration) | **Post** /integrations/def/register | upsert an integration definition
[**SaveAllIntegrations**](IntegrationResourceAPI.md#SaveAllIntegrations) | **Post** /integrations/ | Save all Integrations
[**SaveIntegrationApi**](IntegrationResourceAPI.md#SaveIntegrationApi) | **Post** /integrations/provider/{name}/integration/{integration_name} | Create or Update Integration
[**SaveIntegrationProvider**](IntegrationResourceAPI.md#SaveIntegrationProvider) | **Post** /integrations/provider/{name} | Create or Update Integration provider



## AssociatePromptWithIntegration

> AssociatePromptWithIntegration(ctx, integrationProvider, integrationName, promptName).Execute()

Associate a Prompt Template with an Integration

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
    integrationProvider := "integrationProvider_example" // string | 
    integrationName := "integrationName_example" // string | 
    promptName := "promptName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationResourceAPI.AssociatePromptWithIntegration(context.Background(), integrationProvider, integrationName, promptName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.AssociatePromptWithIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationProvider** | **string** |  | 
**integrationName** | **string** |  | 
**promptName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociatePromptWithIntegrationRequest struct via the builder pattern


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


## DeleteIntegrationApi

> DeleteIntegrationApi(ctx, name, integrationName).Execute()

Delete an Integration

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
    integrationName := "integrationName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationResourceAPI.DeleteIntegrationApi(context.Background(), name, integrationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.DeleteIntegrationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**integrationName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationApiRequest struct via the builder pattern


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


## DeleteIntegrationProvider

> DeleteIntegrationProvider(ctx, name).Execute()

Delete an Integration Provider

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
    r, err := apiClient.IntegrationResourceAPI.DeleteIntegrationProvider(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.DeleteIntegrationProvider``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIntegrationProviderRequest struct via the builder pattern


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


## DeleteTagForIntegration

> DeleteTagForIntegration(ctx, name, integrationName).Tag(tag).Execute()

Delete a tag for Integration

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
    integrationName := "integrationName_example" // string | 
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationResourceAPI.DeleteTagForIntegration(context.Background(), name, integrationName).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.DeleteTagForIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**integrationName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagForIntegrationRequest struct via the builder pattern


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


## DeleteTagForIntegrationProvider

> DeleteTagForIntegrationProvider(ctx, name).Tag(tag).Execute()

Delete a tag for Integration Provider

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
    r, err := apiClient.IntegrationResourceAPI.DeleteTagForIntegrationProvider(context.Background(), name).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.DeleteTagForIntegrationProvider``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTagForIntegrationProviderRequest struct via the builder pattern


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


## GetAllIntegrations

> []Integration GetAllIntegrations(ctx).Category(category).ActiveOnly(activeOnly).Execute()

Get all Integrations

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
    category := "category_example" // string |  (optional)
    activeOnly := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationResourceAPI.GetAllIntegrations(context.Background()).Category(category).ActiveOnly(activeOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.GetAllIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllIntegrations`: []Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationResourceAPI.GetAllIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** |  | 
 **activeOnly** | **bool** |  | [default to true]

### Return type

[**[]Integration**](Integration.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationApi

> IntegrationApi GetIntegrationApi(ctx, name, integrationName).Execute()

Get Integration details

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
    integrationName := "integrationName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationResourceAPI.GetIntegrationApi(context.Background(), name, integrationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.GetIntegrationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationApi`: IntegrationApi
    fmt.Fprintf(os.Stdout, "Response from `IntegrationResourceAPI.GetIntegrationApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**integrationName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IntegrationApi**](IntegrationApi.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationApis

> []IntegrationApi GetIntegrationApis(ctx, name).ActiveOnly(activeOnly).Execute()

Get Integrations of an Integration Provider

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
    activeOnly := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationResourceAPI.GetIntegrationApis(context.Background(), name).ActiveOnly(activeOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.GetIntegrationApis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationApis`: []IntegrationApi
    fmt.Fprintf(os.Stdout, "Response from `IntegrationResourceAPI.GetIntegrationApis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationApisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeOnly** | **bool** |  | [default to true]

### Return type

[**[]IntegrationApi**](IntegrationApi.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationAvailableApis

> []string GetIntegrationAvailableApis(ctx, name).Execute()

Get Integrations Available for an Integration Provider

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
    resp, r, err := apiClient.IntegrationResourceAPI.GetIntegrationAvailableApis(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.GetIntegrationAvailableApis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationAvailableApis`: []string
    fmt.Fprintf(os.Stdout, "Response from `IntegrationResourceAPI.GetIntegrationAvailableApis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationAvailableApisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetIntegrationDef

> IntegrationDef GetIntegrationDef(ctx, name).Execute()

Get an integration definition

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
    resp, r, err := apiClient.IntegrationResourceAPI.GetIntegrationDef(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.GetIntegrationDef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationDef`: IntegrationDef
    fmt.Fprintf(os.Stdout, "Response from `IntegrationResourceAPI.GetIntegrationDef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationDefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationDef**](IntegrationDef.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationProvider

> Integration GetIntegrationProvider(ctx, name).Execute()

Get Integration provider

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
    resp, r, err := apiClient.IntegrationResourceAPI.GetIntegrationProvider(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.GetIntegrationProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationProvider`: Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationResourceAPI.GetIntegrationProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Integration**](Integration.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationProviderDefs

> []IntegrationDef GetIntegrationProviderDefs(ctx).Execute()

Get Integration provider definitions

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
    resp, r, err := apiClient.IntegrationResourceAPI.GetIntegrationProviderDefs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.GetIntegrationProviderDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationProviderDefs`: []IntegrationDef
    fmt.Fprintf(os.Stdout, "Response from `IntegrationResourceAPI.GetIntegrationProviderDefs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationProviderDefsRequest struct via the builder pattern


### Return type

[**[]IntegrationDef**](IntegrationDef.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationProviders

> []Integration GetIntegrationProviders(ctx).Category(category).ActiveOnly(activeOnly).Execute()

Get all Integrations Providers

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
    category := "category_example" // string |  (optional)
    activeOnly := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationResourceAPI.GetIntegrationProviders(context.Background()).Category(category).ActiveOnly(activeOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.GetIntegrationProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationProviders`: []Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationResourceAPI.GetIntegrationProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** |  | 
 **activeOnly** | **bool** |  | [default to true]

### Return type

[**[]Integration**](Integration.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromptsWithIntegration

> []MessageTemplate GetPromptsWithIntegration(ctx, integrationProvider, integrationName).Execute()

Get the list of prompt templates associated with an integration

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
    integrationProvider := "integrationProvider_example" // string | 
    integrationName := "integrationName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationResourceAPI.GetPromptsWithIntegration(context.Background(), integrationProvider, integrationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.GetPromptsWithIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPromptsWithIntegration`: []MessageTemplate
    fmt.Fprintf(os.Stdout, "Response from `IntegrationResourceAPI.GetPromptsWithIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationProvider** | **string** |  | 
**integrationName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptsWithIntegrationRequest struct via the builder pattern


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


## GetProvidersAndIntegrations

> []string GetProvidersAndIntegrations(ctx).Type_(type_).ActiveOnly(activeOnly).Execute()

Get Integrations Providers and Integrations combo

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
    type_ := "type__example" // string |  (optional)
    activeOnly := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationResourceAPI.GetProvidersAndIntegrations(context.Background()).Type_(type_).ActiveOnly(activeOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.GetProvidersAndIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProvidersAndIntegrations`: []string
    fmt.Fprintf(os.Stdout, "Response from `IntegrationResourceAPI.GetProvidersAndIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProvidersAndIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **activeOnly** | **bool** |  | [default to true]

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


## GetTagsForIntegration

> []Tag GetTagsForIntegration(ctx, name, integrationName).Execute()

Get tags by Integration

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
    integrationName := "integrationName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationResourceAPI.GetTagsForIntegration(context.Background(), name, integrationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.GetTagsForIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagsForIntegration`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IntegrationResourceAPI.GetTagsForIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**integrationName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsForIntegrationRequest struct via the builder pattern


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


## GetTagsForIntegrationProvider

> []Tag GetTagsForIntegrationProvider(ctx, name).Execute()

Get tags by Integration Provider

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
    resp, r, err := apiClient.IntegrationResourceAPI.GetTagsForIntegrationProvider(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.GetTagsForIntegrationProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagsForIntegrationProvider`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `IntegrationResourceAPI.GetTagsForIntegrationProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsForIntegrationProviderRequest struct via the builder pattern


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


## PutTagForIntegration

> PutTagForIntegration(ctx, name, integrationName).Tag(tag).Execute()

Put a tag to Integration

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
    integrationName := "integrationName_example" // string | 
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationResourceAPI.PutTagForIntegration(context.Background(), name, integrationName).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.PutTagForIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**integrationName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTagForIntegrationRequest struct via the builder pattern


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


## PutTagForIntegrationProvider

> PutTagForIntegrationProvider(ctx, name).Tag(tag).Execute()

Put a tag to Integration Provider

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
    r, err := apiClient.IntegrationResourceAPI.PutTagForIntegrationProvider(context.Background(), name).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.PutTagForIntegrationProvider``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutTagForIntegrationProviderRequest struct via the builder pattern


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


## RecordEventStats

> RecordEventStats(ctx, type_).EventLog(eventLog).Execute()

Record Event Stats

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
    type_ := "type__example" // string | 
    eventLog := []openapiclient.EventLog{*openapiclient.NewEventLog()} // []EventLog | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationResourceAPI.RecordEventStats(context.Background(), type_).EventLog(eventLog).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.RecordEventStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordEventStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventLog** | [**[]EventLog**](EventLog.md) |  | 

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


## RegisterIntegration

> RegisterIntegration(ctx).IntegrationDef(integrationDef).Execute()

upsert an integration definition

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
    integrationDef := []openapiclient.IntegrationDef{*openapiclient.NewIntegrationDef()} // []IntegrationDef | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationResourceAPI.RegisterIntegration(context.Background()).IntegrationDef(integrationDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.RegisterIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationDef** | [**[]IntegrationDef**](IntegrationDef.md) |  | 

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


## SaveAllIntegrations

> SaveAllIntegrations(ctx).Integration(integration).Execute()

Save all Integrations

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
    integration := []openapiclient.Integration{*openapiclient.NewIntegration()} // []Integration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationResourceAPI.SaveAllIntegrations(context.Background()).Integration(integration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.SaveAllIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveAllIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integration** | [**[]Integration**](Integration.md) |  | 

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


## SaveIntegrationApi

> SaveIntegrationApi(ctx, name, integrationName).IntegrationApiUpdate(integrationApiUpdate).Execute()

Create or Update Integration

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
    integrationName := "integrationName_example" // string | 
    integrationApiUpdate := *openapiclient.NewIntegrationApiUpdate() // IntegrationApiUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationResourceAPI.SaveIntegrationApi(context.Background(), name, integrationName).IntegrationApiUpdate(integrationApiUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.SaveIntegrationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**integrationName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveIntegrationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **integrationApiUpdate** | [**IntegrationApiUpdate**](IntegrationApiUpdate.md) |  | 

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


## SaveIntegrationProvider

> SaveIntegrationProvider(ctx, name).IntegrationUpdate(integrationUpdate).Execute()

Create or Update Integration provider

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
    integrationUpdate := *openapiclient.NewIntegrationUpdate() // IntegrationUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationResourceAPI.SaveIntegrationProvider(context.Background(), name).IntegrationUpdate(integrationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationResourceAPI.SaveIntegrationProvider``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSaveIntegrationProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationUpdate** | [**IntegrationUpdate**](IntegrationUpdate.md) |  | 

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

