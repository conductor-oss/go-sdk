# StartWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**CorrelationId** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**TaskToDomain** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewStartWorkflow

`func NewStartWorkflow() *StartWorkflow`

NewStartWorkflow instantiates a new StartWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartWorkflowWithDefaults

`func NewStartWorkflowWithDefaults() *StartWorkflow`

NewStartWorkflowWithDefaults instantiates a new StartWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StartWorkflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StartWorkflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StartWorkflow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StartWorkflow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *StartWorkflow) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StartWorkflow) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StartWorkflow) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StartWorkflow) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCorrelationId

`func (o *StartWorkflow) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *StartWorkflow) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *StartWorkflow) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *StartWorkflow) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetInput

`func (o *StartWorkflow) GetInput() map[string]map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *StartWorkflow) GetInputOk() (*map[string]map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *StartWorkflow) SetInput(v map[string]map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *StartWorkflow) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetTaskToDomain

`func (o *StartWorkflow) GetTaskToDomain() map[string]string`

GetTaskToDomain returns the TaskToDomain field if non-nil, zero value otherwise.

### GetTaskToDomainOk

`func (o *StartWorkflow) GetTaskToDomainOk() (*map[string]string, bool)`

GetTaskToDomainOk returns a tuple with the TaskToDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskToDomain

`func (o *StartWorkflow) SetTaskToDomain(v map[string]string)`

SetTaskToDomain sets TaskToDomain field to given value.

### HasTaskToDomain

`func (o *StartWorkflow) HasTaskToDomain() bool`

HasTaskToDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


