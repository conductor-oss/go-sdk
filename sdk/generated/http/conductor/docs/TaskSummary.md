# TaskSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | Pointer to **string** |  | [optional] 
**WorkflowType** | Pointer to **string** |  | [optional] 
**CorrelationId** | Pointer to **string** |  | [optional] 
**ScheduledTime** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ReasonForIncompletion** | Pointer to **string** |  | [optional] 
**ExecutionTime** | Pointer to **int64** |  | [optional] 
**QueueWaitTime** | Pointer to **int64** |  | [optional] 
**TaskDefName** | Pointer to **string** |  | [optional] 
**TaskType** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **string** |  | [optional] 
**Output** | Pointer to **string** |  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] 
**ExternalInputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**ExternalOutputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**WorkflowPriority** | Pointer to **int32** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskSummary

`func NewTaskSummary() *TaskSummary`

NewTaskSummary instantiates a new TaskSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskSummaryWithDefaults

`func NewTaskSummaryWithDefaults() *TaskSummary`

NewTaskSummaryWithDefaults instantiates a new TaskSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *TaskSummary) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *TaskSummary) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *TaskSummary) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *TaskSummary) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *TaskSummary) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *TaskSummary) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *TaskSummary) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *TaskSummary) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### GetCorrelationId

`func (o *TaskSummary) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *TaskSummary) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *TaskSummary) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *TaskSummary) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetScheduledTime

`func (o *TaskSummary) GetScheduledTime() string`

GetScheduledTime returns the ScheduledTime field if non-nil, zero value otherwise.

### GetScheduledTimeOk

`func (o *TaskSummary) GetScheduledTimeOk() (*string, bool)`

GetScheduledTimeOk returns a tuple with the ScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTime

`func (o *TaskSummary) SetScheduledTime(v string)`

SetScheduledTime sets ScheduledTime field to given value.

### HasScheduledTime

`func (o *TaskSummary) HasScheduledTime() bool`

HasScheduledTime returns a boolean if a field has been set.

### GetStartTime

`func (o *TaskSummary) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TaskSummary) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TaskSummary) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TaskSummary) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *TaskSummary) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *TaskSummary) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *TaskSummary) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *TaskSummary) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetEndTime

`func (o *TaskSummary) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TaskSummary) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TaskSummary) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TaskSummary) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStatus

`func (o *TaskSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReasonForIncompletion

`func (o *TaskSummary) GetReasonForIncompletion() string`

GetReasonForIncompletion returns the ReasonForIncompletion field if non-nil, zero value otherwise.

### GetReasonForIncompletionOk

`func (o *TaskSummary) GetReasonForIncompletionOk() (*string, bool)`

GetReasonForIncompletionOk returns a tuple with the ReasonForIncompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForIncompletion

`func (o *TaskSummary) SetReasonForIncompletion(v string)`

SetReasonForIncompletion sets ReasonForIncompletion field to given value.

### HasReasonForIncompletion

`func (o *TaskSummary) HasReasonForIncompletion() bool`

HasReasonForIncompletion returns a boolean if a field has been set.

### GetExecutionTime

`func (o *TaskSummary) GetExecutionTime() int64`

GetExecutionTime returns the ExecutionTime field if non-nil, zero value otherwise.

### GetExecutionTimeOk

`func (o *TaskSummary) GetExecutionTimeOk() (*int64, bool)`

GetExecutionTimeOk returns a tuple with the ExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTime

`func (o *TaskSummary) SetExecutionTime(v int64)`

SetExecutionTime sets ExecutionTime field to given value.

### HasExecutionTime

`func (o *TaskSummary) HasExecutionTime() bool`

HasExecutionTime returns a boolean if a field has been set.

### GetQueueWaitTime

`func (o *TaskSummary) GetQueueWaitTime() int64`

GetQueueWaitTime returns the QueueWaitTime field if non-nil, zero value otherwise.

### GetQueueWaitTimeOk

`func (o *TaskSummary) GetQueueWaitTimeOk() (*int64, bool)`

GetQueueWaitTimeOk returns a tuple with the QueueWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueWaitTime

`func (o *TaskSummary) SetQueueWaitTime(v int64)`

SetQueueWaitTime sets QueueWaitTime field to given value.

### HasQueueWaitTime

`func (o *TaskSummary) HasQueueWaitTime() bool`

HasQueueWaitTime returns a boolean if a field has been set.

### GetTaskDefName

`func (o *TaskSummary) GetTaskDefName() string`

GetTaskDefName returns the TaskDefName field if non-nil, zero value otherwise.

### GetTaskDefNameOk

`func (o *TaskSummary) GetTaskDefNameOk() (*string, bool)`

GetTaskDefNameOk returns a tuple with the TaskDefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefName

`func (o *TaskSummary) SetTaskDefName(v string)`

SetTaskDefName sets TaskDefName field to given value.

### HasTaskDefName

`func (o *TaskSummary) HasTaskDefName() bool`

HasTaskDefName returns a boolean if a field has been set.

### GetTaskType

`func (o *TaskSummary) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *TaskSummary) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *TaskSummary) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *TaskSummary) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetInput

`func (o *TaskSummary) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *TaskSummary) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *TaskSummary) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *TaskSummary) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *TaskSummary) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *TaskSummary) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *TaskSummary) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *TaskSummary) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetTaskId

`func (o *TaskSummary) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskSummary) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskSummary) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TaskSummary) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetExternalInputPayloadStoragePath

`func (o *TaskSummary) GetExternalInputPayloadStoragePath() string`

GetExternalInputPayloadStoragePath returns the ExternalInputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalInputPayloadStoragePathOk

`func (o *TaskSummary) GetExternalInputPayloadStoragePathOk() (*string, bool)`

GetExternalInputPayloadStoragePathOk returns a tuple with the ExternalInputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInputPayloadStoragePath

`func (o *TaskSummary) SetExternalInputPayloadStoragePath(v string)`

SetExternalInputPayloadStoragePath sets ExternalInputPayloadStoragePath field to given value.

### HasExternalInputPayloadStoragePath

`func (o *TaskSummary) HasExternalInputPayloadStoragePath() bool`

HasExternalInputPayloadStoragePath returns a boolean if a field has been set.

### GetExternalOutputPayloadStoragePath

`func (o *TaskSummary) GetExternalOutputPayloadStoragePath() string`

GetExternalOutputPayloadStoragePath returns the ExternalOutputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalOutputPayloadStoragePathOk

`func (o *TaskSummary) GetExternalOutputPayloadStoragePathOk() (*string, bool)`

GetExternalOutputPayloadStoragePathOk returns a tuple with the ExternalOutputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOutputPayloadStoragePath

`func (o *TaskSummary) SetExternalOutputPayloadStoragePath(v string)`

SetExternalOutputPayloadStoragePath sets ExternalOutputPayloadStoragePath field to given value.

### HasExternalOutputPayloadStoragePath

`func (o *TaskSummary) HasExternalOutputPayloadStoragePath() bool`

HasExternalOutputPayloadStoragePath returns a boolean if a field has been set.

### GetWorkflowPriority

`func (o *TaskSummary) GetWorkflowPriority() int32`

GetWorkflowPriority returns the WorkflowPriority field if non-nil, zero value otherwise.

### GetWorkflowPriorityOk

`func (o *TaskSummary) GetWorkflowPriorityOk() (*int32, bool)`

GetWorkflowPriorityOk returns a tuple with the WorkflowPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowPriority

`func (o *TaskSummary) SetWorkflowPriority(v int32)`

SetWorkflowPriority sets WorkflowPriority field to given value.

### HasWorkflowPriority

`func (o *TaskSummary) HasWorkflowPriority() bool`

HasWorkflowPriority returns a boolean if a field has been set.

### GetDomain

`func (o *TaskSummary) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TaskSummary) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TaskSummary) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *TaskSummary) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


