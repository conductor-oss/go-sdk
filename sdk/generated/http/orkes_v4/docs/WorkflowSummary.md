# WorkflowSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**ExecutionTime** | Pointer to **int64** |  | [optional] 
**ExternalInputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**ExternalOutputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**FailedReferenceTaskNames** | Pointer to **string** |  | [optional] 
**FailedTaskNames** | Pointer to **[]string** |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **string** |  | [optional] 
**InputSize** | Pointer to **int64** |  | [optional] 
**Output** | Pointer to **string** |  | [optional] 
**OutputSize** | Pointer to **int64** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**ReasonForIncompletion** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TaskToDomain** | Pointer to  |  | [optional] 
**UpdateTime** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**WorkflowType** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowSummary

`func NewWorkflowSummary() *WorkflowSummary`

NewWorkflowSummary instantiates a new WorkflowSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSummaryWithDefaults

`func NewWorkflowSummaryWithDefaults() *WorkflowSummary`

NewWorkflowSummaryWithDefaults instantiates a new WorkflowSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *WorkflowSummary) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *WorkflowSummary) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *WorkflowSummary) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *WorkflowSummary) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WorkflowSummary) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkflowSummary) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkflowSummary) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkflowSummary) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowSummary) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowSummary) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowSummary) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowSummary) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEvent

`func (o *WorkflowSummary) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WorkflowSummary) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WorkflowSummary) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WorkflowSummary) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetExecutionTime

`func (o *WorkflowSummary) GetExecutionTime() int64`

GetExecutionTime returns the ExecutionTime field if non-nil, zero value otherwise.

### GetExecutionTimeOk

`func (o *WorkflowSummary) GetExecutionTimeOk() (*int64, bool)`

GetExecutionTimeOk returns a tuple with the ExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTime

`func (o *WorkflowSummary) SetExecutionTime(v int64)`

SetExecutionTime sets ExecutionTime field to given value.

### HasExecutionTime

`func (o *WorkflowSummary) HasExecutionTime() bool`

HasExecutionTime returns a boolean if a field has been set.

### GetExternalInputPayloadStoragePath

`func (o *WorkflowSummary) GetExternalInputPayloadStoragePath() string`

GetExternalInputPayloadStoragePath returns the ExternalInputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalInputPayloadStoragePathOk

`func (o *WorkflowSummary) GetExternalInputPayloadStoragePathOk() (*string, bool)`

GetExternalInputPayloadStoragePathOk returns a tuple with the ExternalInputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInputPayloadStoragePath

`func (o *WorkflowSummary) SetExternalInputPayloadStoragePath(v string)`

SetExternalInputPayloadStoragePath sets ExternalInputPayloadStoragePath field to given value.

### HasExternalInputPayloadStoragePath

`func (o *WorkflowSummary) HasExternalInputPayloadStoragePath() bool`

HasExternalInputPayloadStoragePath returns a boolean if a field has been set.

### GetExternalOutputPayloadStoragePath

`func (o *WorkflowSummary) GetExternalOutputPayloadStoragePath() string`

GetExternalOutputPayloadStoragePath returns the ExternalOutputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalOutputPayloadStoragePathOk

`func (o *WorkflowSummary) GetExternalOutputPayloadStoragePathOk() (*string, bool)`

GetExternalOutputPayloadStoragePathOk returns a tuple with the ExternalOutputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOutputPayloadStoragePath

`func (o *WorkflowSummary) SetExternalOutputPayloadStoragePath(v string)`

SetExternalOutputPayloadStoragePath sets ExternalOutputPayloadStoragePath field to given value.

### HasExternalOutputPayloadStoragePath

`func (o *WorkflowSummary) HasExternalOutputPayloadStoragePath() bool`

HasExternalOutputPayloadStoragePath returns a boolean if a field has been set.

### GetFailedReferenceTaskNames

`func (o *WorkflowSummary) GetFailedReferenceTaskNames() string`

GetFailedReferenceTaskNames returns the FailedReferenceTaskNames field if non-nil, zero value otherwise.

### GetFailedReferenceTaskNamesOk

`func (o *WorkflowSummary) GetFailedReferenceTaskNamesOk() (*string, bool)`

GetFailedReferenceTaskNamesOk returns a tuple with the FailedReferenceTaskNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReferenceTaskNames

`func (o *WorkflowSummary) SetFailedReferenceTaskNames(v string)`

SetFailedReferenceTaskNames sets FailedReferenceTaskNames field to given value.

### HasFailedReferenceTaskNames

`func (o *WorkflowSummary) HasFailedReferenceTaskNames() bool`

HasFailedReferenceTaskNames returns a boolean if a field has been set.

### GetFailedTaskNames

`func (o *WorkflowSummary) GetFailedTaskNames() []string`

GetFailedTaskNames returns the FailedTaskNames field if non-nil, zero value otherwise.

### GetFailedTaskNamesOk

`func (o *WorkflowSummary) GetFailedTaskNamesOk() (*[]string, bool)`

GetFailedTaskNamesOk returns a tuple with the FailedTaskNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedTaskNames

`func (o *WorkflowSummary) SetFailedTaskNames(v []string)`

SetFailedTaskNames sets FailedTaskNames field to given value.

### HasFailedTaskNames

`func (o *WorkflowSummary) HasFailedTaskNames() bool`

HasFailedTaskNames returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *WorkflowSummary) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *WorkflowSummary) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *WorkflowSummary) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *WorkflowSummary) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowSummary) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowSummary) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowSummary) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowSummary) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetInputSize

`func (o *WorkflowSummary) GetInputSize() int64`

GetInputSize returns the InputSize field if non-nil, zero value otherwise.

### GetInputSizeOk

`func (o *WorkflowSummary) GetInputSizeOk() (*int64, bool)`

GetInputSizeOk returns a tuple with the InputSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSize

`func (o *WorkflowSummary) SetInputSize(v int64)`

SetInputSize sets InputSize field to given value.

### HasInputSize

`func (o *WorkflowSummary) HasInputSize() bool`

HasInputSize returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowSummary) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowSummary) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowSummary) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowSummary) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetOutputSize

`func (o *WorkflowSummary) GetOutputSize() int64`

GetOutputSize returns the OutputSize field if non-nil, zero value otherwise.

### GetOutputSizeOk

`func (o *WorkflowSummary) GetOutputSizeOk() (*int64, bool)`

GetOutputSizeOk returns a tuple with the OutputSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSize

`func (o *WorkflowSummary) SetOutputSize(v int64)`

SetOutputSize sets OutputSize field to given value.

### HasOutputSize

`func (o *WorkflowSummary) HasOutputSize() bool`

HasOutputSize returns a boolean if a field has been set.

### GetPriority

`func (o *WorkflowSummary) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkflowSummary) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkflowSummary) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *WorkflowSummary) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetReasonForIncompletion

`func (o *WorkflowSummary) GetReasonForIncompletion() string`

GetReasonForIncompletion returns the ReasonForIncompletion field if non-nil, zero value otherwise.

### GetReasonForIncompletionOk

`func (o *WorkflowSummary) GetReasonForIncompletionOk() (*string, bool)`

GetReasonForIncompletionOk returns a tuple with the ReasonForIncompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForIncompletion

`func (o *WorkflowSummary) SetReasonForIncompletion(v string)`

SetReasonForIncompletion sets ReasonForIncompletion field to given value.

### HasReasonForIncompletion

`func (o *WorkflowSummary) HasReasonForIncompletion() bool`

HasReasonForIncompletion returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkflowSummary) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowSummary) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowSummary) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkflowSummary) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskToDomain

`func (o *WorkflowSummary) GetTaskToDomain() map[string]string`

GetTaskToDomain returns the TaskToDomain field if non-nil, zero value otherwise.

### GetTaskToDomainOk

`func (o *WorkflowSummary) GetTaskToDomainOk() (*map[string]string, bool)`

GetTaskToDomainOk returns a tuple with the TaskToDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskToDomain

`func (o *WorkflowSummary) SetTaskToDomain(v map[string]string)`

SetTaskToDomain sets TaskToDomain field to given value.

### HasTaskToDomain

`func (o *WorkflowSummary) HasTaskToDomain() bool`

HasTaskToDomain returns a boolean if a field has been set.

### SetTaskToDomainNil

`func (o *WorkflowSummary) SetTaskToDomainNil(b bool)`

 SetTaskToDomainNil sets the value for TaskToDomain to be an explicit nil

### UnsetTaskToDomain
`func (o *WorkflowSummary) UnsetTaskToDomain()`

UnsetTaskToDomain ensures that no value is present for TaskToDomain, not even an explicit nil
### GetUpdateTime

`func (o *WorkflowSummary) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *WorkflowSummary) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *WorkflowSummary) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *WorkflowSummary) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowSummary) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowSummary) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowSummary) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowSummary) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowSummary) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowSummary) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowSummary) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowSummary) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *WorkflowSummary) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowSummary) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowSummary) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *WorkflowSummary) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


