# SubWorkflowParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | Pointer to **string** |  | [optional] 
**IdempotencyStrategy** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **map[string]interface{}** |  | [optional] 
**TaskToDomain** | Pointer to **map[string]string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**WorkflowDefinition** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSubWorkflowParams

`func NewSubWorkflowParams() *SubWorkflowParams`

NewSubWorkflowParams instantiates a new SubWorkflowParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubWorkflowParamsWithDefaults

`func NewSubWorkflowParamsWithDefaults() *SubWorkflowParams`

NewSubWorkflowParamsWithDefaults instantiates a new SubWorkflowParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdempotencyKey

`func (o *SubWorkflowParams) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *SubWorkflowParams) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *SubWorkflowParams) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *SubWorkflowParams) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetIdempotencyStrategy

`func (o *SubWorkflowParams) GetIdempotencyStrategy() string`

GetIdempotencyStrategy returns the IdempotencyStrategy field if non-nil, zero value otherwise.

### GetIdempotencyStrategyOk

`func (o *SubWorkflowParams) GetIdempotencyStrategyOk() (*string, bool)`

GetIdempotencyStrategyOk returns a tuple with the IdempotencyStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyStrategy

`func (o *SubWorkflowParams) SetIdempotencyStrategy(v string)`

SetIdempotencyStrategy sets IdempotencyStrategy field to given value.

### HasIdempotencyStrategy

`func (o *SubWorkflowParams) HasIdempotencyStrategy() bool`

HasIdempotencyStrategy returns a boolean if a field has been set.

### GetName

`func (o *SubWorkflowParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubWorkflowParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubWorkflowParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubWorkflowParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *SubWorkflowParams) GetPriority() map[string]interface{}`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SubWorkflowParams) GetPriorityOk() (*map[string]interface{}, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SubWorkflowParams) SetPriority(v map[string]interface{})`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SubWorkflowParams) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTaskToDomain

`func (o *SubWorkflowParams) GetTaskToDomain() map[string]string`

GetTaskToDomain returns the TaskToDomain field if non-nil, zero value otherwise.

### GetTaskToDomainOk

`func (o *SubWorkflowParams) GetTaskToDomainOk() (*map[string]string, bool)`

GetTaskToDomainOk returns a tuple with the TaskToDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskToDomain

`func (o *SubWorkflowParams) SetTaskToDomain(v map[string]string)`

SetTaskToDomain sets TaskToDomain field to given value.

### HasTaskToDomain

`func (o *SubWorkflowParams) HasTaskToDomain() bool`

HasTaskToDomain returns a boolean if a field has been set.

### GetVersion

`func (o *SubWorkflowParams) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SubWorkflowParams) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SubWorkflowParams) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SubWorkflowParams) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkflowDefinition

`func (o *SubWorkflowParams) GetWorkflowDefinition() map[string]interface{}`

GetWorkflowDefinition returns the WorkflowDefinition field if non-nil, zero value otherwise.

### GetWorkflowDefinitionOk

`func (o *SubWorkflowParams) GetWorkflowDefinitionOk() (*map[string]interface{}, bool)`

GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinition

`func (o *SubWorkflowParams) SetWorkflowDefinition(v map[string]interface{})`

SetWorkflowDefinition sets WorkflowDefinition field to given value.

### HasWorkflowDefinition

`func (o *SubWorkflowParams) HasWorkflowDefinition() bool`

HasWorkflowDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


