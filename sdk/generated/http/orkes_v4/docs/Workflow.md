# Workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**ExternalInputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**ExternalOutputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**FailedReferenceTaskNames** | Pointer to **[]string** |  | [optional] 
**FailedTaskNames** | Pointer to **[]string** |  | [optional] 
**History** | Pointer to [**[]Workflow**](Workflow.md) |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **map[string]interface{}** |  | [optional] 
**LastRetriedTime** | Pointer to **int64** |  | [optional] 
**Output** | Pointer to **map[string]interface{}** |  | [optional] 
**OwnerApp** | Pointer to **string** |  | [optional] 
**ParentWorkflowId** | Pointer to **string** |  | [optional] 
**ParentWorkflowTaskId** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**RateLimitKey** | Pointer to **string** |  | [optional] 
**RateLimited** | Pointer to **bool** |  | [optional] 
**ReRunFromWorkflowId** | Pointer to **string** |  | [optional] 
**ReasonForIncompletion** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TaskToDomain** | Pointer to  |  | [optional] 
**Tasks** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **map[string]interface{}** |  | [optional] 
**WorkflowDefinition** | Pointer to [**WorkflowDef**](WorkflowDef.md) |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**WorkflowName** | Pointer to **string** |  | [optional] 
**WorkflowVersion** | Pointer to **int32** |  | [optional] 

## Methods

### NewWorkflow

`func NewWorkflow() *Workflow`

NewWorkflow instantiates a new Workflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWithDefaults

`func NewWorkflowWithDefaults() *Workflow`

NewWorkflowWithDefaults instantiates a new Workflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *Workflow) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *Workflow) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *Workflow) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *Workflow) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetCreateTime

`func (o *Workflow) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Workflow) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Workflow) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Workflow) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Workflow) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Workflow) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Workflow) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Workflow) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEndTime

`func (o *Workflow) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Workflow) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Workflow) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Workflow) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEvent

`func (o *Workflow) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Workflow) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Workflow) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *Workflow) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetExternalInputPayloadStoragePath

`func (o *Workflow) GetExternalInputPayloadStoragePath() string`

GetExternalInputPayloadStoragePath returns the ExternalInputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalInputPayloadStoragePathOk

`func (o *Workflow) GetExternalInputPayloadStoragePathOk() (*string, bool)`

GetExternalInputPayloadStoragePathOk returns a tuple with the ExternalInputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInputPayloadStoragePath

`func (o *Workflow) SetExternalInputPayloadStoragePath(v string)`

SetExternalInputPayloadStoragePath sets ExternalInputPayloadStoragePath field to given value.

### HasExternalInputPayloadStoragePath

`func (o *Workflow) HasExternalInputPayloadStoragePath() bool`

HasExternalInputPayloadStoragePath returns a boolean if a field has been set.

### GetExternalOutputPayloadStoragePath

`func (o *Workflow) GetExternalOutputPayloadStoragePath() string`

GetExternalOutputPayloadStoragePath returns the ExternalOutputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalOutputPayloadStoragePathOk

`func (o *Workflow) GetExternalOutputPayloadStoragePathOk() (*string, bool)`

GetExternalOutputPayloadStoragePathOk returns a tuple with the ExternalOutputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOutputPayloadStoragePath

`func (o *Workflow) SetExternalOutputPayloadStoragePath(v string)`

SetExternalOutputPayloadStoragePath sets ExternalOutputPayloadStoragePath field to given value.

### HasExternalOutputPayloadStoragePath

`func (o *Workflow) HasExternalOutputPayloadStoragePath() bool`

HasExternalOutputPayloadStoragePath returns a boolean if a field has been set.

### GetFailedReferenceTaskNames

`func (o *Workflow) GetFailedReferenceTaskNames() []string`

GetFailedReferenceTaskNames returns the FailedReferenceTaskNames field if non-nil, zero value otherwise.

### GetFailedReferenceTaskNamesOk

`func (o *Workflow) GetFailedReferenceTaskNamesOk() (*[]string, bool)`

GetFailedReferenceTaskNamesOk returns a tuple with the FailedReferenceTaskNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReferenceTaskNames

`func (o *Workflow) SetFailedReferenceTaskNames(v []string)`

SetFailedReferenceTaskNames sets FailedReferenceTaskNames field to given value.

### HasFailedReferenceTaskNames

`func (o *Workflow) HasFailedReferenceTaskNames() bool`

HasFailedReferenceTaskNames returns a boolean if a field has been set.

### GetFailedTaskNames

`func (o *Workflow) GetFailedTaskNames() []string`

GetFailedTaskNames returns the FailedTaskNames field if non-nil, zero value otherwise.

### GetFailedTaskNamesOk

`func (o *Workflow) GetFailedTaskNamesOk() (*[]string, bool)`

GetFailedTaskNamesOk returns a tuple with the FailedTaskNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedTaskNames

`func (o *Workflow) SetFailedTaskNames(v []string)`

SetFailedTaskNames sets FailedTaskNames field to given value.

### HasFailedTaskNames

`func (o *Workflow) HasFailedTaskNames() bool`

HasFailedTaskNames returns a boolean if a field has been set.

### GetHistory

`func (o *Workflow) GetHistory() []Workflow`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *Workflow) GetHistoryOk() (*[]Workflow, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *Workflow) SetHistory(v []Workflow)`

SetHistory sets History field to given value.

### HasHistory

`func (o *Workflow) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *Workflow) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *Workflow) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *Workflow) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *Workflow) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetInput

`func (o *Workflow) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Workflow) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Workflow) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *Workflow) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetLastRetriedTime

`func (o *Workflow) GetLastRetriedTime() int64`

GetLastRetriedTime returns the LastRetriedTime field if non-nil, zero value otherwise.

### GetLastRetriedTimeOk

`func (o *Workflow) GetLastRetriedTimeOk() (*int64, bool)`

GetLastRetriedTimeOk returns a tuple with the LastRetriedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRetriedTime

`func (o *Workflow) SetLastRetriedTime(v int64)`

SetLastRetriedTime sets LastRetriedTime field to given value.

### HasLastRetriedTime

`func (o *Workflow) HasLastRetriedTime() bool`

HasLastRetriedTime returns a boolean if a field has been set.

### GetOutput

`func (o *Workflow) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *Workflow) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *Workflow) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *Workflow) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetOwnerApp

`func (o *Workflow) GetOwnerApp() string`

GetOwnerApp returns the OwnerApp field if non-nil, zero value otherwise.

### GetOwnerAppOk

`func (o *Workflow) GetOwnerAppOk() (*string, bool)`

GetOwnerAppOk returns a tuple with the OwnerApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerApp

`func (o *Workflow) SetOwnerApp(v string)`

SetOwnerApp sets OwnerApp field to given value.

### HasOwnerApp

`func (o *Workflow) HasOwnerApp() bool`

HasOwnerApp returns a boolean if a field has been set.

### GetParentWorkflowId

`func (o *Workflow) GetParentWorkflowId() string`

GetParentWorkflowId returns the ParentWorkflowId field if non-nil, zero value otherwise.

### GetParentWorkflowIdOk

`func (o *Workflow) GetParentWorkflowIdOk() (*string, bool)`

GetParentWorkflowIdOk returns a tuple with the ParentWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkflowId

`func (o *Workflow) SetParentWorkflowId(v string)`

SetParentWorkflowId sets ParentWorkflowId field to given value.

### HasParentWorkflowId

`func (o *Workflow) HasParentWorkflowId() bool`

HasParentWorkflowId returns a boolean if a field has been set.

### GetParentWorkflowTaskId

`func (o *Workflow) GetParentWorkflowTaskId() string`

GetParentWorkflowTaskId returns the ParentWorkflowTaskId field if non-nil, zero value otherwise.

### GetParentWorkflowTaskIdOk

`func (o *Workflow) GetParentWorkflowTaskIdOk() (*string, bool)`

GetParentWorkflowTaskIdOk returns a tuple with the ParentWorkflowTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkflowTaskId

`func (o *Workflow) SetParentWorkflowTaskId(v string)`

SetParentWorkflowTaskId sets ParentWorkflowTaskId field to given value.

### HasParentWorkflowTaskId

`func (o *Workflow) HasParentWorkflowTaskId() bool`

HasParentWorkflowTaskId returns a boolean if a field has been set.

### GetPriority

`func (o *Workflow) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Workflow) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Workflow) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Workflow) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRateLimitKey

`func (o *Workflow) GetRateLimitKey() string`

GetRateLimitKey returns the RateLimitKey field if non-nil, zero value otherwise.

### GetRateLimitKeyOk

`func (o *Workflow) GetRateLimitKeyOk() (*string, bool)`

GetRateLimitKeyOk returns a tuple with the RateLimitKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitKey

`func (o *Workflow) SetRateLimitKey(v string)`

SetRateLimitKey sets RateLimitKey field to given value.

### HasRateLimitKey

`func (o *Workflow) HasRateLimitKey() bool`

HasRateLimitKey returns a boolean if a field has been set.

### GetRateLimited

`func (o *Workflow) GetRateLimited() bool`

GetRateLimited returns the RateLimited field if non-nil, zero value otherwise.

### GetRateLimitedOk

`func (o *Workflow) GetRateLimitedOk() (*bool, bool)`

GetRateLimitedOk returns a tuple with the RateLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimited

`func (o *Workflow) SetRateLimited(v bool)`

SetRateLimited sets RateLimited field to given value.

### HasRateLimited

`func (o *Workflow) HasRateLimited() bool`

HasRateLimited returns a boolean if a field has been set.

### GetReRunFromWorkflowId

`func (o *Workflow) GetReRunFromWorkflowId() string`

GetReRunFromWorkflowId returns the ReRunFromWorkflowId field if non-nil, zero value otherwise.

### GetReRunFromWorkflowIdOk

`func (o *Workflow) GetReRunFromWorkflowIdOk() (*string, bool)`

GetReRunFromWorkflowIdOk returns a tuple with the ReRunFromWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReRunFromWorkflowId

`func (o *Workflow) SetReRunFromWorkflowId(v string)`

SetReRunFromWorkflowId sets ReRunFromWorkflowId field to given value.

### HasReRunFromWorkflowId

`func (o *Workflow) HasReRunFromWorkflowId() bool`

HasReRunFromWorkflowId returns a boolean if a field has been set.

### GetReasonForIncompletion

`func (o *Workflow) GetReasonForIncompletion() string`

GetReasonForIncompletion returns the ReasonForIncompletion field if non-nil, zero value otherwise.

### GetReasonForIncompletionOk

`func (o *Workflow) GetReasonForIncompletionOk() (*string, bool)`

GetReasonForIncompletionOk returns a tuple with the ReasonForIncompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForIncompletion

`func (o *Workflow) SetReasonForIncompletion(v string)`

SetReasonForIncompletion sets ReasonForIncompletion field to given value.

### HasReasonForIncompletion

`func (o *Workflow) HasReasonForIncompletion() bool`

HasReasonForIncompletion returns a boolean if a field has been set.

### GetStartTime

`func (o *Workflow) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Workflow) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Workflow) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Workflow) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *Workflow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Workflow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Workflow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Workflow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskToDomain

`func (o *Workflow) GetTaskToDomain() map[string]string`

GetTaskToDomain returns the TaskToDomain field if non-nil, zero value otherwise.

### GetTaskToDomainOk

`func (o *Workflow) GetTaskToDomainOk() (*map[string]string, bool)`

GetTaskToDomainOk returns a tuple with the TaskToDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskToDomain

`func (o *Workflow) SetTaskToDomain(v map[string]string)`

SetTaskToDomain sets TaskToDomain field to given value.

### HasTaskToDomain

`func (o *Workflow) HasTaskToDomain() bool`

HasTaskToDomain returns a boolean if a field has been set.

### SetTaskToDomainNil

`func (o *Workflow) SetTaskToDomainNil(b bool)`

 SetTaskToDomainNil sets the value for TaskToDomain to be an explicit nil

### UnsetTaskToDomain
`func (o *Workflow) UnsetTaskToDomain()`

UnsetTaskToDomain ensures that no value is present for TaskToDomain, not even an explicit nil
### GetTasks

`func (o *Workflow) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Workflow) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Workflow) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *Workflow) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Workflow) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Workflow) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Workflow) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *Workflow) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Workflow) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Workflow) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Workflow) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Workflow) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVariables

`func (o *Workflow) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Workflow) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Workflow) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Workflow) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetWorkflowDefinition

`func (o *Workflow) GetWorkflowDefinition() WorkflowDef`

GetWorkflowDefinition returns the WorkflowDefinition field if non-nil, zero value otherwise.

### GetWorkflowDefinitionOk

`func (o *Workflow) GetWorkflowDefinitionOk() (*WorkflowDef, bool)`

GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinition

`func (o *Workflow) SetWorkflowDefinition(v WorkflowDef)`

SetWorkflowDefinition sets WorkflowDefinition field to given value.

### HasWorkflowDefinition

`func (o *Workflow) HasWorkflowDefinition() bool`

HasWorkflowDefinition returns a boolean if a field has been set.

### GetWorkflowId

`func (o *Workflow) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *Workflow) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *Workflow) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *Workflow) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowName

`func (o *Workflow) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *Workflow) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *Workflow) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.

### HasWorkflowName

`func (o *Workflow) HasWorkflowName() bool`

HasWorkflowName returns a boolean if a field has been set.

### GetWorkflowVersion

`func (o *Workflow) GetWorkflowVersion() int32`

GetWorkflowVersion returns the WorkflowVersion field if non-nil, zero value otherwise.

### GetWorkflowVersionOk

`func (o *Workflow) GetWorkflowVersionOk() (*int32, bool)`

GetWorkflowVersionOk returns a tuple with the WorkflowVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowVersion

`func (o *Workflow) SetWorkflowVersion(v int32)`

SetWorkflowVersion sets WorkflowVersion field to given value.

### HasWorkflowVersion

`func (o *Workflow) HasWorkflowVersion() bool`

HasWorkflowVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


