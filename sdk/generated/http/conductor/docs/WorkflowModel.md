# WorkflowModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**ParentWorkflowId** | Pointer to **string** |  | [optional] 
**ParentWorkflowTaskId** | Pointer to **string** |  | [optional] 
**Tasks** | Pointer to [**[]TaskModel**](TaskModel.md) |  | [optional] 
**CorrelationId** | Pointer to **string** |  | [optional] 
**ReRunFromWorkflowId** | Pointer to **string** |  | [optional] 
**ReasonForIncompletion** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**TaskToDomain** | Pointer to **map[string]string** |  | [optional] 
**FailedReferenceTaskNames** | Pointer to **[]string** |  | [optional] 
**FailedTaskNames** | Pointer to **[]string** |  | [optional] 
**WorkflowDefinition** | Pointer to [**WorkflowDef**](WorkflowDef.md) |  | [optional] 
**ExternalInputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**ExternalOutputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Variables** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**LastRetriedTime** | Pointer to **int64** |  | [optional] 
**OwnerApp** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**UpdatedTime** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**FailedTaskId** | Pointer to **string** |  | [optional] 
**PreviousStatus** | Pointer to **string** |  | [optional] 
**WorkflowName** | Pointer to **string** |  | [optional] 
**WorkflowVersion** | Pointer to **int32** |  | [optional] 
**Input** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Output** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewWorkflowModel

`func NewWorkflowModel() *WorkflowModel`

NewWorkflowModel instantiates a new WorkflowModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowModelWithDefaults

`func NewWorkflowModelWithDefaults() *WorkflowModel`

NewWorkflowModelWithDefaults instantiates a new WorkflowModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *WorkflowModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowModel) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowModel) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowModel) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowModel) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowModel) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowModel) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowModel) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowModel) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetParentWorkflowId

`func (o *WorkflowModel) GetParentWorkflowId() string`

GetParentWorkflowId returns the ParentWorkflowId field if non-nil, zero value otherwise.

### GetParentWorkflowIdOk

`func (o *WorkflowModel) GetParentWorkflowIdOk() (*string, bool)`

GetParentWorkflowIdOk returns a tuple with the ParentWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkflowId

`func (o *WorkflowModel) SetParentWorkflowId(v string)`

SetParentWorkflowId sets ParentWorkflowId field to given value.

### HasParentWorkflowId

`func (o *WorkflowModel) HasParentWorkflowId() bool`

HasParentWorkflowId returns a boolean if a field has been set.

### GetParentWorkflowTaskId

`func (o *WorkflowModel) GetParentWorkflowTaskId() string`

GetParentWorkflowTaskId returns the ParentWorkflowTaskId field if non-nil, zero value otherwise.

### GetParentWorkflowTaskIdOk

`func (o *WorkflowModel) GetParentWorkflowTaskIdOk() (*string, bool)`

GetParentWorkflowTaskIdOk returns a tuple with the ParentWorkflowTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkflowTaskId

`func (o *WorkflowModel) SetParentWorkflowTaskId(v string)`

SetParentWorkflowTaskId sets ParentWorkflowTaskId field to given value.

### HasParentWorkflowTaskId

`func (o *WorkflowModel) HasParentWorkflowTaskId() bool`

HasParentWorkflowTaskId returns a boolean if a field has been set.

### GetTasks

`func (o *WorkflowModel) GetTasks() []TaskModel`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *WorkflowModel) GetTasksOk() (*[]TaskModel, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *WorkflowModel) SetTasks(v []TaskModel)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *WorkflowModel) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetCorrelationId

`func (o *WorkflowModel) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *WorkflowModel) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *WorkflowModel) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *WorkflowModel) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetReRunFromWorkflowId

`func (o *WorkflowModel) GetReRunFromWorkflowId() string`

GetReRunFromWorkflowId returns the ReRunFromWorkflowId field if non-nil, zero value otherwise.

### GetReRunFromWorkflowIdOk

`func (o *WorkflowModel) GetReRunFromWorkflowIdOk() (*string, bool)`

GetReRunFromWorkflowIdOk returns a tuple with the ReRunFromWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReRunFromWorkflowId

`func (o *WorkflowModel) SetReRunFromWorkflowId(v string)`

SetReRunFromWorkflowId sets ReRunFromWorkflowId field to given value.

### HasReRunFromWorkflowId

`func (o *WorkflowModel) HasReRunFromWorkflowId() bool`

HasReRunFromWorkflowId returns a boolean if a field has been set.

### GetReasonForIncompletion

`func (o *WorkflowModel) GetReasonForIncompletion() string`

GetReasonForIncompletion returns the ReasonForIncompletion field if non-nil, zero value otherwise.

### GetReasonForIncompletionOk

`func (o *WorkflowModel) GetReasonForIncompletionOk() (*string, bool)`

GetReasonForIncompletionOk returns a tuple with the ReasonForIncompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForIncompletion

`func (o *WorkflowModel) SetReasonForIncompletion(v string)`

SetReasonForIncompletion sets ReasonForIncompletion field to given value.

### HasReasonForIncompletion

`func (o *WorkflowModel) HasReasonForIncompletion() bool`

HasReasonForIncompletion returns a boolean if a field has been set.

### GetEvent

`func (o *WorkflowModel) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WorkflowModel) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WorkflowModel) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WorkflowModel) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetTaskToDomain

`func (o *WorkflowModel) GetTaskToDomain() map[string]string`

GetTaskToDomain returns the TaskToDomain field if non-nil, zero value otherwise.

### GetTaskToDomainOk

`func (o *WorkflowModel) GetTaskToDomainOk() (*map[string]string, bool)`

GetTaskToDomainOk returns a tuple with the TaskToDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskToDomain

`func (o *WorkflowModel) SetTaskToDomain(v map[string]string)`

SetTaskToDomain sets TaskToDomain field to given value.

### HasTaskToDomain

`func (o *WorkflowModel) HasTaskToDomain() bool`

HasTaskToDomain returns a boolean if a field has been set.

### GetFailedReferenceTaskNames

`func (o *WorkflowModel) GetFailedReferenceTaskNames() []string`

GetFailedReferenceTaskNames returns the FailedReferenceTaskNames field if non-nil, zero value otherwise.

### GetFailedReferenceTaskNamesOk

`func (o *WorkflowModel) GetFailedReferenceTaskNamesOk() (*[]string, bool)`

GetFailedReferenceTaskNamesOk returns a tuple with the FailedReferenceTaskNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReferenceTaskNames

`func (o *WorkflowModel) SetFailedReferenceTaskNames(v []string)`

SetFailedReferenceTaskNames sets FailedReferenceTaskNames field to given value.

### HasFailedReferenceTaskNames

`func (o *WorkflowModel) HasFailedReferenceTaskNames() bool`

HasFailedReferenceTaskNames returns a boolean if a field has been set.

### GetFailedTaskNames

`func (o *WorkflowModel) GetFailedTaskNames() []string`

GetFailedTaskNames returns the FailedTaskNames field if non-nil, zero value otherwise.

### GetFailedTaskNamesOk

`func (o *WorkflowModel) GetFailedTaskNamesOk() (*[]string, bool)`

GetFailedTaskNamesOk returns a tuple with the FailedTaskNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedTaskNames

`func (o *WorkflowModel) SetFailedTaskNames(v []string)`

SetFailedTaskNames sets FailedTaskNames field to given value.

### HasFailedTaskNames

`func (o *WorkflowModel) HasFailedTaskNames() bool`

HasFailedTaskNames returns a boolean if a field has been set.

### GetWorkflowDefinition

`func (o *WorkflowModel) GetWorkflowDefinition() WorkflowDef`

GetWorkflowDefinition returns the WorkflowDefinition field if non-nil, zero value otherwise.

### GetWorkflowDefinitionOk

`func (o *WorkflowModel) GetWorkflowDefinitionOk() (*WorkflowDef, bool)`

GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinition

`func (o *WorkflowModel) SetWorkflowDefinition(v WorkflowDef)`

SetWorkflowDefinition sets WorkflowDefinition field to given value.

### HasWorkflowDefinition

`func (o *WorkflowModel) HasWorkflowDefinition() bool`

HasWorkflowDefinition returns a boolean if a field has been set.

### GetExternalInputPayloadStoragePath

`func (o *WorkflowModel) GetExternalInputPayloadStoragePath() string`

GetExternalInputPayloadStoragePath returns the ExternalInputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalInputPayloadStoragePathOk

`func (o *WorkflowModel) GetExternalInputPayloadStoragePathOk() (*string, bool)`

GetExternalInputPayloadStoragePathOk returns a tuple with the ExternalInputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInputPayloadStoragePath

`func (o *WorkflowModel) SetExternalInputPayloadStoragePath(v string)`

SetExternalInputPayloadStoragePath sets ExternalInputPayloadStoragePath field to given value.

### HasExternalInputPayloadStoragePath

`func (o *WorkflowModel) HasExternalInputPayloadStoragePath() bool`

HasExternalInputPayloadStoragePath returns a boolean if a field has been set.

### GetExternalOutputPayloadStoragePath

`func (o *WorkflowModel) GetExternalOutputPayloadStoragePath() string`

GetExternalOutputPayloadStoragePath returns the ExternalOutputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalOutputPayloadStoragePathOk

`func (o *WorkflowModel) GetExternalOutputPayloadStoragePathOk() (*string, bool)`

GetExternalOutputPayloadStoragePathOk returns a tuple with the ExternalOutputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOutputPayloadStoragePath

`func (o *WorkflowModel) SetExternalOutputPayloadStoragePath(v string)`

SetExternalOutputPayloadStoragePath sets ExternalOutputPayloadStoragePath field to given value.

### HasExternalOutputPayloadStoragePath

`func (o *WorkflowModel) HasExternalOutputPayloadStoragePath() bool`

HasExternalOutputPayloadStoragePath returns a boolean if a field has been set.

### GetPriority

`func (o *WorkflowModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkflowModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkflowModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *WorkflowModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetVariables

`func (o *WorkflowModel) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *WorkflowModel) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *WorkflowModel) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *WorkflowModel) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetLastRetriedTime

`func (o *WorkflowModel) GetLastRetriedTime() int64`

GetLastRetriedTime returns the LastRetriedTime field if non-nil, zero value otherwise.

### GetLastRetriedTimeOk

`func (o *WorkflowModel) GetLastRetriedTimeOk() (*int64, bool)`

GetLastRetriedTimeOk returns a tuple with the LastRetriedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRetriedTime

`func (o *WorkflowModel) SetLastRetriedTime(v int64)`

SetLastRetriedTime sets LastRetriedTime field to given value.

### HasLastRetriedTime

`func (o *WorkflowModel) HasLastRetriedTime() bool`

HasLastRetriedTime returns a boolean if a field has been set.

### GetOwnerApp

`func (o *WorkflowModel) GetOwnerApp() string`

GetOwnerApp returns the OwnerApp field if non-nil, zero value otherwise.

### GetOwnerAppOk

`func (o *WorkflowModel) GetOwnerAppOk() (*string, bool)`

GetOwnerAppOk returns a tuple with the OwnerApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerApp

`func (o *WorkflowModel) SetOwnerApp(v string)`

SetOwnerApp sets OwnerApp field to given value.

### HasOwnerApp

`func (o *WorkflowModel) HasOwnerApp() bool`

HasOwnerApp returns a boolean if a field has been set.

### GetCreateTime

`func (o *WorkflowModel) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *WorkflowModel) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *WorkflowModel) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *WorkflowModel) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *WorkflowModel) GetUpdatedTime() int64`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *WorkflowModel) GetUpdatedTimeOk() (*int64, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *WorkflowModel) SetUpdatedTime(v int64)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *WorkflowModel) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WorkflowModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkflowModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkflowModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkflowModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkflowModel) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkflowModel) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkflowModel) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkflowModel) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetFailedTaskId

`func (o *WorkflowModel) GetFailedTaskId() string`

GetFailedTaskId returns the FailedTaskId field if non-nil, zero value otherwise.

### GetFailedTaskIdOk

`func (o *WorkflowModel) GetFailedTaskIdOk() (*string, bool)`

GetFailedTaskIdOk returns a tuple with the FailedTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedTaskId

`func (o *WorkflowModel) SetFailedTaskId(v string)`

SetFailedTaskId sets FailedTaskId field to given value.

### HasFailedTaskId

`func (o *WorkflowModel) HasFailedTaskId() bool`

HasFailedTaskId returns a boolean if a field has been set.

### GetPreviousStatus

`func (o *WorkflowModel) GetPreviousStatus() string`

GetPreviousStatus returns the PreviousStatus field if non-nil, zero value otherwise.

### GetPreviousStatusOk

`func (o *WorkflowModel) GetPreviousStatusOk() (*string, bool)`

GetPreviousStatusOk returns a tuple with the PreviousStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStatus

`func (o *WorkflowModel) SetPreviousStatus(v string)`

SetPreviousStatus sets PreviousStatus field to given value.

### HasPreviousStatus

`func (o *WorkflowModel) HasPreviousStatus() bool`

HasPreviousStatus returns a boolean if a field has been set.

### GetWorkflowName

`func (o *WorkflowModel) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *WorkflowModel) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *WorkflowModel) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.

### HasWorkflowName

`func (o *WorkflowModel) HasWorkflowName() bool`

HasWorkflowName returns a boolean if a field has been set.

### GetWorkflowVersion

`func (o *WorkflowModel) GetWorkflowVersion() int32`

GetWorkflowVersion returns the WorkflowVersion field if non-nil, zero value otherwise.

### GetWorkflowVersionOk

`func (o *WorkflowModel) GetWorkflowVersionOk() (*int32, bool)`

GetWorkflowVersionOk returns a tuple with the WorkflowVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowVersion

`func (o *WorkflowModel) SetWorkflowVersion(v int32)`

SetWorkflowVersion sets WorkflowVersion field to given value.

### HasWorkflowVersion

`func (o *WorkflowModel) HasWorkflowVersion() bool`

HasWorkflowVersion returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowModel) GetInput() map[string]map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowModel) GetInputOk() (*map[string]map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowModel) SetInput(v map[string]map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowModel) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowModel) GetOutput() map[string]map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowModel) GetOutputOk() (*map[string]map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowModel) SetOutput(v map[string]map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowModel) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


