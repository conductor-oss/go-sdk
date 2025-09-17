# TaskModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ReferenceTaskName** | Pointer to **string** |  | [optional] 
**RetryCount** | Pointer to **int32** |  | [optional] 
**Seq** | Pointer to **int32** |  | [optional] 
**CorrelationId** | Pointer to **string** |  | [optional] 
**PollCount** | Pointer to **int32** |  | [optional] 
**TaskDefName** | Pointer to **string** |  | [optional] 
**ScheduledTime** | Pointer to **int64** |  | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**StartDelayInSeconds** | Pointer to **int32** |  | [optional] 
**RetriedTaskId** | Pointer to **string** |  | [optional] 
**Retried** | Pointer to **bool** |  | [optional] 
**Executed** | Pointer to **bool** |  | [optional] 
**CallbackFromWorker** | Pointer to **bool** |  | [optional] 
**ResponseTimeoutSeconds** | Pointer to **int64** |  | [optional] 
**WorkflowInstanceId** | Pointer to **string** |  | [optional] 
**WorkflowType** | Pointer to **string** |  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] 
**ReasonForIncompletion** | Pointer to **string** |  | [optional] 
**CallbackAfterSeconds** | Pointer to **int64** |  | [optional] 
**WorkerId** | Pointer to **string** |  | [optional] 
**WorkflowTask** | Pointer to [**WorkflowTask**](WorkflowTask.md) |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**InputMessage** | Pointer to [**Any**](Any.md) |  | [optional] 
**OutputMessage** | Pointer to [**Any**](Any.md) |  | [optional] 
**RateLimitPerFrequency** | Pointer to **int32** |  | [optional] 
**RateLimitFrequencyInSeconds** | Pointer to **int32** |  | [optional] 
**ExternalInputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**ExternalOutputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**WorkflowPriority** | Pointer to **int32** |  | [optional] 
**ExecutionNameSpace** | Pointer to **string** |  | [optional] 
**IsolationGroupId** | Pointer to **string** |  | [optional] 
**Iteration** | Pointer to **int32** |  | [optional] 
**SubWorkflowId** | Pointer to **string** |  | [optional] 
**WaitTimeout** | Pointer to **int64** |  | [optional] 
**SubworkflowChanged** | Pointer to **bool** |  | [optional] 
**QueueWaitTime** | Pointer to **int64** |  | [optional] 
**LoopOverTask** | Pointer to **bool** |  | [optional] 
**InputData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**OutputData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewTaskModel

`func NewTaskModel() *TaskModel`

NewTaskModel instantiates a new TaskModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskModelWithDefaults

`func NewTaskModelWithDefaults() *TaskModel`

NewTaskModelWithDefaults instantiates a new TaskModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskType

`func (o *TaskModel) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *TaskModel) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *TaskModel) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *TaskModel) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetStatus

`func (o *TaskModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReferenceTaskName

`func (o *TaskModel) GetReferenceTaskName() string`

GetReferenceTaskName returns the ReferenceTaskName field if non-nil, zero value otherwise.

### GetReferenceTaskNameOk

`func (o *TaskModel) GetReferenceTaskNameOk() (*string, bool)`

GetReferenceTaskNameOk returns a tuple with the ReferenceTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTaskName

`func (o *TaskModel) SetReferenceTaskName(v string)`

SetReferenceTaskName sets ReferenceTaskName field to given value.

### HasReferenceTaskName

`func (o *TaskModel) HasReferenceTaskName() bool`

HasReferenceTaskName returns a boolean if a field has been set.

### GetRetryCount

`func (o *TaskModel) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *TaskModel) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *TaskModel) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *TaskModel) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetSeq

`func (o *TaskModel) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *TaskModel) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *TaskModel) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *TaskModel) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetCorrelationId

`func (o *TaskModel) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *TaskModel) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *TaskModel) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *TaskModel) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetPollCount

`func (o *TaskModel) GetPollCount() int32`

GetPollCount returns the PollCount field if non-nil, zero value otherwise.

### GetPollCountOk

`func (o *TaskModel) GetPollCountOk() (*int32, bool)`

GetPollCountOk returns a tuple with the PollCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollCount

`func (o *TaskModel) SetPollCount(v int32)`

SetPollCount sets PollCount field to given value.

### HasPollCount

`func (o *TaskModel) HasPollCount() bool`

HasPollCount returns a boolean if a field has been set.

### GetTaskDefName

`func (o *TaskModel) GetTaskDefName() string`

GetTaskDefName returns the TaskDefName field if non-nil, zero value otherwise.

### GetTaskDefNameOk

`func (o *TaskModel) GetTaskDefNameOk() (*string, bool)`

GetTaskDefNameOk returns a tuple with the TaskDefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefName

`func (o *TaskModel) SetTaskDefName(v string)`

SetTaskDefName sets TaskDefName field to given value.

### HasTaskDefName

`func (o *TaskModel) HasTaskDefName() bool`

HasTaskDefName returns a boolean if a field has been set.

### GetScheduledTime

`func (o *TaskModel) GetScheduledTime() int64`

GetScheduledTime returns the ScheduledTime field if non-nil, zero value otherwise.

### GetScheduledTimeOk

`func (o *TaskModel) GetScheduledTimeOk() (*int64, bool)`

GetScheduledTimeOk returns a tuple with the ScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTime

`func (o *TaskModel) SetScheduledTime(v int64)`

SetScheduledTime sets ScheduledTime field to given value.

### HasScheduledTime

`func (o *TaskModel) HasScheduledTime() bool`

HasScheduledTime returns a boolean if a field has been set.

### GetStartTime

`func (o *TaskModel) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TaskModel) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TaskModel) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TaskModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *TaskModel) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TaskModel) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TaskModel) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TaskModel) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *TaskModel) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *TaskModel) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *TaskModel) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *TaskModel) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetStartDelayInSeconds

`func (o *TaskModel) GetStartDelayInSeconds() int32`

GetStartDelayInSeconds returns the StartDelayInSeconds field if non-nil, zero value otherwise.

### GetStartDelayInSecondsOk

`func (o *TaskModel) GetStartDelayInSecondsOk() (*int32, bool)`

GetStartDelayInSecondsOk returns a tuple with the StartDelayInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDelayInSeconds

`func (o *TaskModel) SetStartDelayInSeconds(v int32)`

SetStartDelayInSeconds sets StartDelayInSeconds field to given value.

### HasStartDelayInSeconds

`func (o *TaskModel) HasStartDelayInSeconds() bool`

HasStartDelayInSeconds returns a boolean if a field has been set.

### GetRetriedTaskId

`func (o *TaskModel) GetRetriedTaskId() string`

GetRetriedTaskId returns the RetriedTaskId field if non-nil, zero value otherwise.

### GetRetriedTaskIdOk

`func (o *TaskModel) GetRetriedTaskIdOk() (*string, bool)`

GetRetriedTaskIdOk returns a tuple with the RetriedTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetriedTaskId

`func (o *TaskModel) SetRetriedTaskId(v string)`

SetRetriedTaskId sets RetriedTaskId field to given value.

### HasRetriedTaskId

`func (o *TaskModel) HasRetriedTaskId() bool`

HasRetriedTaskId returns a boolean if a field has been set.

### GetRetried

`func (o *TaskModel) GetRetried() bool`

GetRetried returns the Retried field if non-nil, zero value otherwise.

### GetRetriedOk

`func (o *TaskModel) GetRetriedOk() (*bool, bool)`

GetRetriedOk returns a tuple with the Retried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetried

`func (o *TaskModel) SetRetried(v bool)`

SetRetried sets Retried field to given value.

### HasRetried

`func (o *TaskModel) HasRetried() bool`

HasRetried returns a boolean if a field has been set.

### GetExecuted

`func (o *TaskModel) GetExecuted() bool`

GetExecuted returns the Executed field if non-nil, zero value otherwise.

### GetExecutedOk

`func (o *TaskModel) GetExecutedOk() (*bool, bool)`

GetExecutedOk returns a tuple with the Executed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuted

`func (o *TaskModel) SetExecuted(v bool)`

SetExecuted sets Executed field to given value.

### HasExecuted

`func (o *TaskModel) HasExecuted() bool`

HasExecuted returns a boolean if a field has been set.

### GetCallbackFromWorker

`func (o *TaskModel) GetCallbackFromWorker() bool`

GetCallbackFromWorker returns the CallbackFromWorker field if non-nil, zero value otherwise.

### GetCallbackFromWorkerOk

`func (o *TaskModel) GetCallbackFromWorkerOk() (*bool, bool)`

GetCallbackFromWorkerOk returns a tuple with the CallbackFromWorker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackFromWorker

`func (o *TaskModel) SetCallbackFromWorker(v bool)`

SetCallbackFromWorker sets CallbackFromWorker field to given value.

### HasCallbackFromWorker

`func (o *TaskModel) HasCallbackFromWorker() bool`

HasCallbackFromWorker returns a boolean if a field has been set.

### GetResponseTimeoutSeconds

`func (o *TaskModel) GetResponseTimeoutSeconds() int64`

GetResponseTimeoutSeconds returns the ResponseTimeoutSeconds field if non-nil, zero value otherwise.

### GetResponseTimeoutSecondsOk

`func (o *TaskModel) GetResponseTimeoutSecondsOk() (*int64, bool)`

GetResponseTimeoutSecondsOk returns a tuple with the ResponseTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeoutSeconds

`func (o *TaskModel) SetResponseTimeoutSeconds(v int64)`

SetResponseTimeoutSeconds sets ResponseTimeoutSeconds field to given value.

### HasResponseTimeoutSeconds

`func (o *TaskModel) HasResponseTimeoutSeconds() bool`

HasResponseTimeoutSeconds returns a boolean if a field has been set.

### GetWorkflowInstanceId

`func (o *TaskModel) GetWorkflowInstanceId() string`

GetWorkflowInstanceId returns the WorkflowInstanceId field if non-nil, zero value otherwise.

### GetWorkflowInstanceIdOk

`func (o *TaskModel) GetWorkflowInstanceIdOk() (*string, bool)`

GetWorkflowInstanceIdOk returns a tuple with the WorkflowInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInstanceId

`func (o *TaskModel) SetWorkflowInstanceId(v string)`

SetWorkflowInstanceId sets WorkflowInstanceId field to given value.

### HasWorkflowInstanceId

`func (o *TaskModel) HasWorkflowInstanceId() bool`

HasWorkflowInstanceId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *TaskModel) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *TaskModel) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *TaskModel) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *TaskModel) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### GetTaskId

`func (o *TaskModel) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskModel) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskModel) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TaskModel) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetReasonForIncompletion

`func (o *TaskModel) GetReasonForIncompletion() string`

GetReasonForIncompletion returns the ReasonForIncompletion field if non-nil, zero value otherwise.

### GetReasonForIncompletionOk

`func (o *TaskModel) GetReasonForIncompletionOk() (*string, bool)`

GetReasonForIncompletionOk returns a tuple with the ReasonForIncompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForIncompletion

`func (o *TaskModel) SetReasonForIncompletion(v string)`

SetReasonForIncompletion sets ReasonForIncompletion field to given value.

### HasReasonForIncompletion

`func (o *TaskModel) HasReasonForIncompletion() bool`

HasReasonForIncompletion returns a boolean if a field has been set.

### GetCallbackAfterSeconds

`func (o *TaskModel) GetCallbackAfterSeconds() int64`

GetCallbackAfterSeconds returns the CallbackAfterSeconds field if non-nil, zero value otherwise.

### GetCallbackAfterSecondsOk

`func (o *TaskModel) GetCallbackAfterSecondsOk() (*int64, bool)`

GetCallbackAfterSecondsOk returns a tuple with the CallbackAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackAfterSeconds

`func (o *TaskModel) SetCallbackAfterSeconds(v int64)`

SetCallbackAfterSeconds sets CallbackAfterSeconds field to given value.

### HasCallbackAfterSeconds

`func (o *TaskModel) HasCallbackAfterSeconds() bool`

HasCallbackAfterSeconds returns a boolean if a field has been set.

### GetWorkerId

`func (o *TaskModel) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *TaskModel) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *TaskModel) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *TaskModel) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetWorkflowTask

`func (o *TaskModel) GetWorkflowTask() WorkflowTask`

GetWorkflowTask returns the WorkflowTask field if non-nil, zero value otherwise.

### GetWorkflowTaskOk

`func (o *TaskModel) GetWorkflowTaskOk() (*WorkflowTask, bool)`

GetWorkflowTaskOk returns a tuple with the WorkflowTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTask

`func (o *TaskModel) SetWorkflowTask(v WorkflowTask)`

SetWorkflowTask sets WorkflowTask field to given value.

### HasWorkflowTask

`func (o *TaskModel) HasWorkflowTask() bool`

HasWorkflowTask returns a boolean if a field has been set.

### GetDomain

`func (o *TaskModel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TaskModel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TaskModel) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *TaskModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetInputMessage

`func (o *TaskModel) GetInputMessage() Any`

GetInputMessage returns the InputMessage field if non-nil, zero value otherwise.

### GetInputMessageOk

`func (o *TaskModel) GetInputMessageOk() (*Any, bool)`

GetInputMessageOk returns a tuple with the InputMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessage

`func (o *TaskModel) SetInputMessage(v Any)`

SetInputMessage sets InputMessage field to given value.

### HasInputMessage

`func (o *TaskModel) HasInputMessage() bool`

HasInputMessage returns a boolean if a field has been set.

### GetOutputMessage

`func (o *TaskModel) GetOutputMessage() Any`

GetOutputMessage returns the OutputMessage field if non-nil, zero value otherwise.

### GetOutputMessageOk

`func (o *TaskModel) GetOutputMessageOk() (*Any, bool)`

GetOutputMessageOk returns a tuple with the OutputMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputMessage

`func (o *TaskModel) SetOutputMessage(v Any)`

SetOutputMessage sets OutputMessage field to given value.

### HasOutputMessage

`func (o *TaskModel) HasOutputMessage() bool`

HasOutputMessage returns a boolean if a field has been set.

### GetRateLimitPerFrequency

`func (o *TaskModel) GetRateLimitPerFrequency() int32`

GetRateLimitPerFrequency returns the RateLimitPerFrequency field if non-nil, zero value otherwise.

### GetRateLimitPerFrequencyOk

`func (o *TaskModel) GetRateLimitPerFrequencyOk() (*int32, bool)`

GetRateLimitPerFrequencyOk returns a tuple with the RateLimitPerFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerFrequency

`func (o *TaskModel) SetRateLimitPerFrequency(v int32)`

SetRateLimitPerFrequency sets RateLimitPerFrequency field to given value.

### HasRateLimitPerFrequency

`func (o *TaskModel) HasRateLimitPerFrequency() bool`

HasRateLimitPerFrequency returns a boolean if a field has been set.

### GetRateLimitFrequencyInSeconds

`func (o *TaskModel) GetRateLimitFrequencyInSeconds() int32`

GetRateLimitFrequencyInSeconds returns the RateLimitFrequencyInSeconds field if non-nil, zero value otherwise.

### GetRateLimitFrequencyInSecondsOk

`func (o *TaskModel) GetRateLimitFrequencyInSecondsOk() (*int32, bool)`

GetRateLimitFrequencyInSecondsOk returns a tuple with the RateLimitFrequencyInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitFrequencyInSeconds

`func (o *TaskModel) SetRateLimitFrequencyInSeconds(v int32)`

SetRateLimitFrequencyInSeconds sets RateLimitFrequencyInSeconds field to given value.

### HasRateLimitFrequencyInSeconds

`func (o *TaskModel) HasRateLimitFrequencyInSeconds() bool`

HasRateLimitFrequencyInSeconds returns a boolean if a field has been set.

### GetExternalInputPayloadStoragePath

`func (o *TaskModel) GetExternalInputPayloadStoragePath() string`

GetExternalInputPayloadStoragePath returns the ExternalInputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalInputPayloadStoragePathOk

`func (o *TaskModel) GetExternalInputPayloadStoragePathOk() (*string, bool)`

GetExternalInputPayloadStoragePathOk returns a tuple with the ExternalInputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInputPayloadStoragePath

`func (o *TaskModel) SetExternalInputPayloadStoragePath(v string)`

SetExternalInputPayloadStoragePath sets ExternalInputPayloadStoragePath field to given value.

### HasExternalInputPayloadStoragePath

`func (o *TaskModel) HasExternalInputPayloadStoragePath() bool`

HasExternalInputPayloadStoragePath returns a boolean if a field has been set.

### GetExternalOutputPayloadStoragePath

`func (o *TaskModel) GetExternalOutputPayloadStoragePath() string`

GetExternalOutputPayloadStoragePath returns the ExternalOutputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalOutputPayloadStoragePathOk

`func (o *TaskModel) GetExternalOutputPayloadStoragePathOk() (*string, bool)`

GetExternalOutputPayloadStoragePathOk returns a tuple with the ExternalOutputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOutputPayloadStoragePath

`func (o *TaskModel) SetExternalOutputPayloadStoragePath(v string)`

SetExternalOutputPayloadStoragePath sets ExternalOutputPayloadStoragePath field to given value.

### HasExternalOutputPayloadStoragePath

`func (o *TaskModel) HasExternalOutputPayloadStoragePath() bool`

HasExternalOutputPayloadStoragePath returns a boolean if a field has been set.

### GetWorkflowPriority

`func (o *TaskModel) GetWorkflowPriority() int32`

GetWorkflowPriority returns the WorkflowPriority field if non-nil, zero value otherwise.

### GetWorkflowPriorityOk

`func (o *TaskModel) GetWorkflowPriorityOk() (*int32, bool)`

GetWorkflowPriorityOk returns a tuple with the WorkflowPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowPriority

`func (o *TaskModel) SetWorkflowPriority(v int32)`

SetWorkflowPriority sets WorkflowPriority field to given value.

### HasWorkflowPriority

`func (o *TaskModel) HasWorkflowPriority() bool`

HasWorkflowPriority returns a boolean if a field has been set.

### GetExecutionNameSpace

`func (o *TaskModel) GetExecutionNameSpace() string`

GetExecutionNameSpace returns the ExecutionNameSpace field if non-nil, zero value otherwise.

### GetExecutionNameSpaceOk

`func (o *TaskModel) GetExecutionNameSpaceOk() (*string, bool)`

GetExecutionNameSpaceOk returns a tuple with the ExecutionNameSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionNameSpace

`func (o *TaskModel) SetExecutionNameSpace(v string)`

SetExecutionNameSpace sets ExecutionNameSpace field to given value.

### HasExecutionNameSpace

`func (o *TaskModel) HasExecutionNameSpace() bool`

HasExecutionNameSpace returns a boolean if a field has been set.

### GetIsolationGroupId

`func (o *TaskModel) GetIsolationGroupId() string`

GetIsolationGroupId returns the IsolationGroupId field if non-nil, zero value otherwise.

### GetIsolationGroupIdOk

`func (o *TaskModel) GetIsolationGroupIdOk() (*string, bool)`

GetIsolationGroupIdOk returns a tuple with the IsolationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationGroupId

`func (o *TaskModel) SetIsolationGroupId(v string)`

SetIsolationGroupId sets IsolationGroupId field to given value.

### HasIsolationGroupId

`func (o *TaskModel) HasIsolationGroupId() bool`

HasIsolationGroupId returns a boolean if a field has been set.

### GetIteration

`func (o *TaskModel) GetIteration() int32`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *TaskModel) GetIterationOk() (*int32, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *TaskModel) SetIteration(v int32)`

SetIteration sets Iteration field to given value.

### HasIteration

`func (o *TaskModel) HasIteration() bool`

HasIteration returns a boolean if a field has been set.

### GetSubWorkflowId

`func (o *TaskModel) GetSubWorkflowId() string`

GetSubWorkflowId returns the SubWorkflowId field if non-nil, zero value otherwise.

### GetSubWorkflowIdOk

`func (o *TaskModel) GetSubWorkflowIdOk() (*string, bool)`

GetSubWorkflowIdOk returns a tuple with the SubWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWorkflowId

`func (o *TaskModel) SetSubWorkflowId(v string)`

SetSubWorkflowId sets SubWorkflowId field to given value.

### HasSubWorkflowId

`func (o *TaskModel) HasSubWorkflowId() bool`

HasSubWorkflowId returns a boolean if a field has been set.

### GetWaitTimeout

`func (o *TaskModel) GetWaitTimeout() int64`

GetWaitTimeout returns the WaitTimeout field if non-nil, zero value otherwise.

### GetWaitTimeoutOk

`func (o *TaskModel) GetWaitTimeoutOk() (*int64, bool)`

GetWaitTimeoutOk returns a tuple with the WaitTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimeout

`func (o *TaskModel) SetWaitTimeout(v int64)`

SetWaitTimeout sets WaitTimeout field to given value.

### HasWaitTimeout

`func (o *TaskModel) HasWaitTimeout() bool`

HasWaitTimeout returns a boolean if a field has been set.

### GetSubworkflowChanged

`func (o *TaskModel) GetSubworkflowChanged() bool`

GetSubworkflowChanged returns the SubworkflowChanged field if non-nil, zero value otherwise.

### GetSubworkflowChangedOk

`func (o *TaskModel) GetSubworkflowChangedOk() (*bool, bool)`

GetSubworkflowChangedOk returns a tuple with the SubworkflowChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubworkflowChanged

`func (o *TaskModel) SetSubworkflowChanged(v bool)`

SetSubworkflowChanged sets SubworkflowChanged field to given value.

### HasSubworkflowChanged

`func (o *TaskModel) HasSubworkflowChanged() bool`

HasSubworkflowChanged returns a boolean if a field has been set.

### GetQueueWaitTime

`func (o *TaskModel) GetQueueWaitTime() int64`

GetQueueWaitTime returns the QueueWaitTime field if non-nil, zero value otherwise.

### GetQueueWaitTimeOk

`func (o *TaskModel) GetQueueWaitTimeOk() (*int64, bool)`

GetQueueWaitTimeOk returns a tuple with the QueueWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueWaitTime

`func (o *TaskModel) SetQueueWaitTime(v int64)`

SetQueueWaitTime sets QueueWaitTime field to given value.

### HasQueueWaitTime

`func (o *TaskModel) HasQueueWaitTime() bool`

HasQueueWaitTime returns a boolean if a field has been set.

### GetLoopOverTask

`func (o *TaskModel) GetLoopOverTask() bool`

GetLoopOverTask returns the LoopOverTask field if non-nil, zero value otherwise.

### GetLoopOverTaskOk

`func (o *TaskModel) GetLoopOverTaskOk() (*bool, bool)`

GetLoopOverTaskOk returns a tuple with the LoopOverTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopOverTask

`func (o *TaskModel) SetLoopOverTask(v bool)`

SetLoopOverTask sets LoopOverTask field to given value.

### HasLoopOverTask

`func (o *TaskModel) HasLoopOverTask() bool`

HasLoopOverTask returns a boolean if a field has been set.

### GetInputData

`func (o *TaskModel) GetInputData() map[string]map[string]interface{}`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *TaskModel) GetInputDataOk() (*map[string]map[string]interface{}, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *TaskModel) SetInputData(v map[string]map[string]interface{})`

SetInputData sets InputData field to given value.

### HasInputData

`func (o *TaskModel) HasInputData() bool`

HasInputData returns a boolean if a field has been set.

### GetOutputData

`func (o *TaskModel) GetOutputData() map[string]map[string]interface{}`

GetOutputData returns the OutputData field if non-nil, zero value otherwise.

### GetOutputDataOk

`func (o *TaskModel) GetOutputDataOk() (*map[string]map[string]interface{}, bool)`

GetOutputDataOk returns a tuple with the OutputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputData

`func (o *TaskModel) SetOutputData(v map[string]map[string]interface{})`

SetOutputData sets OutputData field to given value.

### HasOutputData

`func (o *TaskModel) HasOutputData() bool`

HasOutputData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


