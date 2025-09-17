# ServiceMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExampleInput** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InputType** | Pointer to **string** |  | [optional] 
**MethodName** | Pointer to **string** |  | [optional] 
**MethodType** | Pointer to **string** |  | [optional] 
**OperationName** | Pointer to **string** |  | [optional] 
**OutputType** | Pointer to **string** |  | [optional] 
**RequestParams** | Pointer to [**[]RequestParam**](RequestParam.md) |  | [optional] 

## Methods

### NewServiceMethod

`func NewServiceMethod() *ServiceMethod`

NewServiceMethod instantiates a new ServiceMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMethodWithDefaults

`func NewServiceMethodWithDefaults() *ServiceMethod`

NewServiceMethodWithDefaults instantiates a new ServiceMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExampleInput

`func (o *ServiceMethod) GetExampleInput() map[string]map[string]interface{}`

GetExampleInput returns the ExampleInput field if non-nil, zero value otherwise.

### GetExampleInputOk

`func (o *ServiceMethod) GetExampleInputOk() (*map[string]map[string]interface{}, bool)`

GetExampleInputOk returns a tuple with the ExampleInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleInput

`func (o *ServiceMethod) SetExampleInput(v map[string]map[string]interface{})`

SetExampleInput sets ExampleInput field to given value.

### HasExampleInput

`func (o *ServiceMethod) HasExampleInput() bool`

HasExampleInput returns a boolean if a field has been set.

### GetId

`func (o *ServiceMethod) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceMethod) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceMethod) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputType

`func (o *ServiceMethod) GetInputType() string`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *ServiceMethod) GetInputTypeOk() (*string, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *ServiceMethod) SetInputType(v string)`

SetInputType sets InputType field to given value.

### HasInputType

`func (o *ServiceMethod) HasInputType() bool`

HasInputType returns a boolean if a field has been set.

### GetMethodName

`func (o *ServiceMethod) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *ServiceMethod) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *ServiceMethod) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.

### HasMethodName

`func (o *ServiceMethod) HasMethodName() bool`

HasMethodName returns a boolean if a field has been set.

### GetMethodType

`func (o *ServiceMethod) GetMethodType() string`

GetMethodType returns the MethodType field if non-nil, zero value otherwise.

### GetMethodTypeOk

`func (o *ServiceMethod) GetMethodTypeOk() (*string, bool)`

GetMethodTypeOk returns a tuple with the MethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodType

`func (o *ServiceMethod) SetMethodType(v string)`

SetMethodType sets MethodType field to given value.

### HasMethodType

`func (o *ServiceMethod) HasMethodType() bool`

HasMethodType returns a boolean if a field has been set.

### GetOperationName

`func (o *ServiceMethod) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *ServiceMethod) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *ServiceMethod) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.

### HasOperationName

`func (o *ServiceMethod) HasOperationName() bool`

HasOperationName returns a boolean if a field has been set.

### GetOutputType

`func (o *ServiceMethod) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *ServiceMethod) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *ServiceMethod) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.

### HasOutputType

`func (o *ServiceMethod) HasOutputType() bool`

HasOutputType returns a boolean if a field has been set.

### GetRequestParams

`func (o *ServiceMethod) GetRequestParams() []RequestParam`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *ServiceMethod) GetRequestParamsOk() (*[]RequestParam, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *ServiceMethod) SetRequestParams(v []RequestParam)`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *ServiceMethod) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


