# ServiceRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CircuitBreakerEnabled** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**Config**](Config.md) |  | [optional] 
**Methods** | Pointer to [**[]ServiceMethod**](ServiceMethod.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RequestParams** | Pointer to [**[]RequestParam**](RequestParam.md) |  | [optional] 
**ServiceURI** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceRegistry

`func NewServiceRegistry() *ServiceRegistry`

NewServiceRegistry instantiates a new ServiceRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRegistryWithDefaults

`func NewServiceRegistryWithDefaults() *ServiceRegistry`

NewServiceRegistryWithDefaults instantiates a new ServiceRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuitBreakerEnabled

`func (o *ServiceRegistry) GetCircuitBreakerEnabled() bool`

GetCircuitBreakerEnabled returns the CircuitBreakerEnabled field if non-nil, zero value otherwise.

### GetCircuitBreakerEnabledOk

`func (o *ServiceRegistry) GetCircuitBreakerEnabledOk() (*bool, bool)`

GetCircuitBreakerEnabledOk returns a tuple with the CircuitBreakerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreakerEnabled

`func (o *ServiceRegistry) SetCircuitBreakerEnabled(v bool)`

SetCircuitBreakerEnabled sets CircuitBreakerEnabled field to given value.

### HasCircuitBreakerEnabled

`func (o *ServiceRegistry) HasCircuitBreakerEnabled() bool`

HasCircuitBreakerEnabled returns a boolean if a field has been set.

### GetConfig

`func (o *ServiceRegistry) GetConfig() Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServiceRegistry) GetConfigOk() (*Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServiceRegistry) SetConfig(v Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ServiceRegistry) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetMethods

`func (o *ServiceRegistry) GetMethods() []ServiceMethod`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *ServiceRegistry) GetMethodsOk() (*[]ServiceMethod, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *ServiceRegistry) SetMethods(v []ServiceMethod)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *ServiceRegistry) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetName

`func (o *ServiceRegistry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceRegistry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceRegistry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceRegistry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequestParams

`func (o *ServiceRegistry) GetRequestParams() []RequestParam`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *ServiceRegistry) GetRequestParamsOk() (*[]RequestParam, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *ServiceRegistry) SetRequestParams(v []RequestParam)`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *ServiceRegistry) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### GetServiceURI

`func (o *ServiceRegistry) GetServiceURI() string`

GetServiceURI returns the ServiceURI field if non-nil, zero value otherwise.

### GetServiceURIOk

`func (o *ServiceRegistry) GetServiceURIOk() (*string, bool)`

GetServiceURIOk returns a tuple with the ServiceURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceURI

`func (o *ServiceRegistry) SetServiceURI(v string)`

SetServiceURI sets ServiceURI field to given value.

### HasServiceURI

`func (o *ServiceRegistry) HasServiceURI() bool`

HasServiceURI returns a boolean if a field has been set.

### GetType

`func (o *ServiceRegistry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceRegistry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceRegistry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceRegistry) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


