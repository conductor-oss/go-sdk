# StartWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ExternalInputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 
**IdempotencyStrategy** | Pointer to **string** |  | [optional] 
**Input** | Pointer to  |  | [optional] 
**Name** | **string** |  | 
**Priority** | Pointer to **int32** |  | [optional] 
**TaskToDomain** | Pointer to  |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**WorkflowDef** | Pointer to [**WorkflowDef**](WorkflowDef.md) |  | [optional] 

## Methods

### NewStartWorkflowRequest

`func NewStartWorkflowRequest(name string, ) *StartWorkflowRequest`

NewStartWorkflowRequest instantiates a new StartWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartWorkflowRequestWithDefaults

`func NewStartWorkflowRequestWithDefaults() *StartWorkflowRequest`

NewStartWorkflowRequestWithDefaults instantiates a new StartWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *StartWorkflowRequest) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *StartWorkflowRequest) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *StartWorkflowRequest) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *StartWorkflowRequest) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *StartWorkflowRequest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *StartWorkflowRequest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *StartWorkflowRequest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *StartWorkflowRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetExternalInputPayloadStoragePath

`func (o *StartWorkflowRequest) GetExternalInputPayloadStoragePath() string`

GetExternalInputPayloadStoragePath returns the ExternalInputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalInputPayloadStoragePathOk

`func (o *StartWorkflowRequest) GetExternalInputPayloadStoragePathOk() (*string, bool)`

GetExternalInputPayloadStoragePathOk returns a tuple with the ExternalInputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInputPayloadStoragePath

`func (o *StartWorkflowRequest) SetExternalInputPayloadStoragePath(v string)`

SetExternalInputPayloadStoragePath sets ExternalInputPayloadStoragePath field to given value.

### HasExternalInputPayloadStoragePath

`func (o *StartWorkflowRequest) HasExternalInputPayloadStoragePath() bool`

HasExternalInputPayloadStoragePath returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *StartWorkflowRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *StartWorkflowRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *StartWorkflowRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *StartWorkflowRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetIdempotencyStrategy

`func (o *StartWorkflowRequest) GetIdempotencyStrategy() string`

GetIdempotencyStrategy returns the IdempotencyStrategy field if non-nil, zero value otherwise.

### GetIdempotencyStrategyOk

`func (o *StartWorkflowRequest) GetIdempotencyStrategyOk() (*string, bool)`

GetIdempotencyStrategyOk returns a tuple with the IdempotencyStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyStrategy

`func (o *StartWorkflowRequest) SetIdempotencyStrategy(v string)`

SetIdempotencyStrategy sets IdempotencyStrategy field to given value.

### HasIdempotencyStrategy

`func (o *StartWorkflowRequest) HasIdempotencyStrategy() bool`

HasIdempotencyStrategy returns a boolean if a field has been set.

### GetInput

`func (o *StartWorkflowRequest) GetInput() map[string]map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *StartWorkflowRequest) GetInputOk() (*map[string]map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *StartWorkflowRequest) SetInput(v map[string]map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *StartWorkflowRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *StartWorkflowRequest) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *StartWorkflowRequest) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetName

`func (o *StartWorkflowRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StartWorkflowRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StartWorkflowRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *StartWorkflowRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *StartWorkflowRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *StartWorkflowRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *StartWorkflowRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTaskToDomain

`func (o *StartWorkflowRequest) GetTaskToDomain() map[string]string`

GetTaskToDomain returns the TaskToDomain field if non-nil, zero value otherwise.

### GetTaskToDomainOk

`func (o *StartWorkflowRequest) GetTaskToDomainOk() (*map[string]string, bool)`

GetTaskToDomainOk returns a tuple with the TaskToDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskToDomain

`func (o *StartWorkflowRequest) SetTaskToDomain(v map[string]string)`

SetTaskToDomain sets TaskToDomain field to given value.

### HasTaskToDomain

`func (o *StartWorkflowRequest) HasTaskToDomain() bool`

HasTaskToDomain returns a boolean if a field has been set.

### SetTaskToDomainNil

`func (o *StartWorkflowRequest) SetTaskToDomainNil(b bool)`

 SetTaskToDomainNil sets the value for TaskToDomain to be an explicit nil

### UnsetTaskToDomain
`func (o *StartWorkflowRequest) UnsetTaskToDomain()`

UnsetTaskToDomain ensures that no value is present for TaskToDomain, not even an explicit nil
### GetVersion

`func (o *StartWorkflowRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StartWorkflowRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StartWorkflowRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StartWorkflowRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkflowDef

`func (o *StartWorkflowRequest) GetWorkflowDef() WorkflowDef`

GetWorkflowDef returns the WorkflowDef field if non-nil, zero value otherwise.

### GetWorkflowDefOk

`func (o *StartWorkflowRequest) GetWorkflowDefOk() (*WorkflowDef, bool)`

GetWorkflowDefOk returns a tuple with the WorkflowDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDef

`func (o *StartWorkflowRequest) SetWorkflowDef(v WorkflowDef)`

SetWorkflowDef sets WorkflowDef field to given value.

### HasWorkflowDef

`func (o *StartWorkflowRequest) HasWorkflowDef() bool`

HasWorkflowDef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


