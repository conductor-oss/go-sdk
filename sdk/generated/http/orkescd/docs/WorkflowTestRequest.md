# WorkflowTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ExternalInputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 
**IdempotencyStrategy** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**Priority** | Pointer to **int32** |  | [optional] 
**SubWorkflowTestRequest** | Pointer to [**map[string]WorkflowTestRequest**](WorkflowTestRequest.md) |  | [optional] 
**TaskRefToMockOutput** | Pointer to [**map[string][]TaskMock**](array.md) |  | [optional] 
**TaskToDomain** | Pointer to **map[string]string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**WorkflowDef** | Pointer to [**WorkflowDef**](WorkflowDef.md) |  | [optional] 

## Methods

### NewWorkflowTestRequest

`func NewWorkflowTestRequest(name string, ) *WorkflowTestRequest`

NewWorkflowTestRequest instantiates a new WorkflowTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTestRequestWithDefaults

`func NewWorkflowTestRequestWithDefaults() *WorkflowTestRequest`

NewWorkflowTestRequestWithDefaults instantiates a new WorkflowTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *WorkflowTestRequest) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *WorkflowTestRequest) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *WorkflowTestRequest) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *WorkflowTestRequest) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WorkflowTestRequest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkflowTestRequest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkflowTestRequest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkflowTestRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetExternalInputPayloadStoragePath

`func (o *WorkflowTestRequest) GetExternalInputPayloadStoragePath() string`

GetExternalInputPayloadStoragePath returns the ExternalInputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalInputPayloadStoragePathOk

`func (o *WorkflowTestRequest) GetExternalInputPayloadStoragePathOk() (*string, bool)`

GetExternalInputPayloadStoragePathOk returns a tuple with the ExternalInputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInputPayloadStoragePath

`func (o *WorkflowTestRequest) SetExternalInputPayloadStoragePath(v string)`

SetExternalInputPayloadStoragePath sets ExternalInputPayloadStoragePath field to given value.

### HasExternalInputPayloadStoragePath

`func (o *WorkflowTestRequest) HasExternalInputPayloadStoragePath() bool`

HasExternalInputPayloadStoragePath returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *WorkflowTestRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *WorkflowTestRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *WorkflowTestRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *WorkflowTestRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetIdempotencyStrategy

`func (o *WorkflowTestRequest) GetIdempotencyStrategy() string`

GetIdempotencyStrategy returns the IdempotencyStrategy field if non-nil, zero value otherwise.

### GetIdempotencyStrategyOk

`func (o *WorkflowTestRequest) GetIdempotencyStrategyOk() (*string, bool)`

GetIdempotencyStrategyOk returns a tuple with the IdempotencyStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyStrategy

`func (o *WorkflowTestRequest) SetIdempotencyStrategy(v string)`

SetIdempotencyStrategy sets IdempotencyStrategy field to given value.

### HasIdempotencyStrategy

`func (o *WorkflowTestRequest) HasIdempotencyStrategy() bool`

HasIdempotencyStrategy returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowTestRequest) GetInput() map[string]map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowTestRequest) GetInputOk() (*map[string]map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowTestRequest) SetInput(v map[string]map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowTestRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTestRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTestRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTestRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *WorkflowTestRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkflowTestRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkflowTestRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *WorkflowTestRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSubWorkflowTestRequest

`func (o *WorkflowTestRequest) GetSubWorkflowTestRequest() map[string]WorkflowTestRequest`

GetSubWorkflowTestRequest returns the SubWorkflowTestRequest field if non-nil, zero value otherwise.

### GetSubWorkflowTestRequestOk

`func (o *WorkflowTestRequest) GetSubWorkflowTestRequestOk() (*map[string]WorkflowTestRequest, bool)`

GetSubWorkflowTestRequestOk returns a tuple with the SubWorkflowTestRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWorkflowTestRequest

`func (o *WorkflowTestRequest) SetSubWorkflowTestRequest(v map[string]WorkflowTestRequest)`

SetSubWorkflowTestRequest sets SubWorkflowTestRequest field to given value.

### HasSubWorkflowTestRequest

`func (o *WorkflowTestRequest) HasSubWorkflowTestRequest() bool`

HasSubWorkflowTestRequest returns a boolean if a field has been set.

### GetTaskRefToMockOutput

`func (o *WorkflowTestRequest) GetTaskRefToMockOutput() map[string][]TaskMock`

GetTaskRefToMockOutput returns the TaskRefToMockOutput field if non-nil, zero value otherwise.

### GetTaskRefToMockOutputOk

`func (o *WorkflowTestRequest) GetTaskRefToMockOutputOk() (*map[string][]TaskMock, bool)`

GetTaskRefToMockOutputOk returns a tuple with the TaskRefToMockOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRefToMockOutput

`func (o *WorkflowTestRequest) SetTaskRefToMockOutput(v map[string][]TaskMock)`

SetTaskRefToMockOutput sets TaskRefToMockOutput field to given value.

### HasTaskRefToMockOutput

`func (o *WorkflowTestRequest) HasTaskRefToMockOutput() bool`

HasTaskRefToMockOutput returns a boolean if a field has been set.

### GetTaskToDomain

`func (o *WorkflowTestRequest) GetTaskToDomain() map[string]string`

GetTaskToDomain returns the TaskToDomain field if non-nil, zero value otherwise.

### GetTaskToDomainOk

`func (o *WorkflowTestRequest) GetTaskToDomainOk() (*map[string]string, bool)`

GetTaskToDomainOk returns a tuple with the TaskToDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskToDomain

`func (o *WorkflowTestRequest) SetTaskToDomain(v map[string]string)`

SetTaskToDomain sets TaskToDomain field to given value.

### HasTaskToDomain

`func (o *WorkflowTestRequest) HasTaskToDomain() bool`

HasTaskToDomain returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowTestRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowTestRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowTestRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowTestRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkflowDef

`func (o *WorkflowTestRequest) GetWorkflowDef() WorkflowDef`

GetWorkflowDef returns the WorkflowDef field if non-nil, zero value otherwise.

### GetWorkflowDefOk

`func (o *WorkflowTestRequest) GetWorkflowDefOk() (*WorkflowDef, bool)`

GetWorkflowDefOk returns a tuple with the WorkflowDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDef

`func (o *WorkflowTestRequest) SetWorkflowDef(v WorkflowDef)`

SetWorkflowDef sets WorkflowDef field to given value.

### HasWorkflowDef

`func (o *WorkflowTestRequest) HasWorkflowDef() bool`

HasWorkflowDef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


