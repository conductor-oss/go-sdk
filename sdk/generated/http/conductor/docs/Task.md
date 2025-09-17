# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**InputData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
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
**OutputData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**WorkflowTask** | Pointer to [**WorkflowTask**](WorkflowTask.md) |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**RateLimitPerFrequency** | Pointer to **int32** |  | [optional] 
**RateLimitFrequencyInSeconds** | Pointer to **int32** |  | [optional] 
**ExternalInputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**ExternalOutputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**WorkflowPriority** | Pointer to **int32** |  | [optional] 
**ExecutionNameSpace** | Pointer to **string** |  | [optional] 
**IsolationGroupId** | Pointer to **string** |  | [optional] 
**Iteration** | Pointer to **int32** |  | [optional] 
**SubWorkflowId** | Pointer to **string** |  | [optional] 
**SubworkflowChanged** | Pointer to **bool** |  | [optional] 
**FirstStartTime** | Pointer to **int64** |  | [optional] 
**ParentTaskId** | Pointer to **string** |  | [optional] 
**QueueWaitTime** | Pointer to **int64** |  | [optional] 
**LoopOverTask** | Pointer to **bool** |  | [optional] 
**TaskDefinition** | Pointer to [**TaskDef**](TaskDef.md) |  | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskType

`func (o *Task) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *Task) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *Task) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *Task) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetStatus

`func (o *Task) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Task) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Task) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Task) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInputData

`func (o *Task) GetInputData() map[string]map[string]interface{}`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *Task) GetInputDataOk() (*map[string]map[string]interface{}, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *Task) SetInputData(v map[string]map[string]interface{})`

SetInputData sets InputData field to given value.

### HasInputData

`func (o *Task) HasInputData() bool`

HasInputData returns a boolean if a field has been set.

### GetReferenceTaskName

`func (o *Task) GetReferenceTaskName() string`

GetReferenceTaskName returns the ReferenceTaskName field if non-nil, zero value otherwise.

### GetReferenceTaskNameOk

`func (o *Task) GetReferenceTaskNameOk() (*string, bool)`

GetReferenceTaskNameOk returns a tuple with the ReferenceTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTaskName

`func (o *Task) SetReferenceTaskName(v string)`

SetReferenceTaskName sets ReferenceTaskName field to given value.

### HasReferenceTaskName

`func (o *Task) HasReferenceTaskName() bool`

HasReferenceTaskName returns a boolean if a field has been set.

### GetRetryCount

`func (o *Task) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *Task) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *Task) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *Task) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetSeq

`func (o *Task) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *Task) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *Task) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *Task) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetCorrelationId

`func (o *Task) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *Task) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *Task) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *Task) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetPollCount

`func (o *Task) GetPollCount() int32`

GetPollCount returns the PollCount field if non-nil, zero value otherwise.

### GetPollCountOk

`func (o *Task) GetPollCountOk() (*int32, bool)`

GetPollCountOk returns a tuple with the PollCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollCount

`func (o *Task) SetPollCount(v int32)`

SetPollCount sets PollCount field to given value.

### HasPollCount

`func (o *Task) HasPollCount() bool`

HasPollCount returns a boolean if a field has been set.

### GetTaskDefName

`func (o *Task) GetTaskDefName() string`

GetTaskDefName returns the TaskDefName field if non-nil, zero value otherwise.

### GetTaskDefNameOk

`func (o *Task) GetTaskDefNameOk() (*string, bool)`

GetTaskDefNameOk returns a tuple with the TaskDefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefName

`func (o *Task) SetTaskDefName(v string)`

SetTaskDefName sets TaskDefName field to given value.

### HasTaskDefName

`func (o *Task) HasTaskDefName() bool`

HasTaskDefName returns a boolean if a field has been set.

### GetScheduledTime

`func (o *Task) GetScheduledTime() int64`

GetScheduledTime returns the ScheduledTime field if non-nil, zero value otherwise.

### GetScheduledTimeOk

`func (o *Task) GetScheduledTimeOk() (*int64, bool)`

GetScheduledTimeOk returns a tuple with the ScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTime

`func (o *Task) SetScheduledTime(v int64)`

SetScheduledTime sets ScheduledTime field to given value.

### HasScheduledTime

`func (o *Task) HasScheduledTime() bool`

HasScheduledTime returns a boolean if a field has been set.

### GetStartTime

`func (o *Task) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Task) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Task) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Task) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *Task) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Task) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Task) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Task) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Task) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Task) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Task) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *Task) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetStartDelayInSeconds

`func (o *Task) GetStartDelayInSeconds() int32`

GetStartDelayInSeconds returns the StartDelayInSeconds field if non-nil, zero value otherwise.

### GetStartDelayInSecondsOk

`func (o *Task) GetStartDelayInSecondsOk() (*int32, bool)`

GetStartDelayInSecondsOk returns a tuple with the StartDelayInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDelayInSeconds

`func (o *Task) SetStartDelayInSeconds(v int32)`

SetStartDelayInSeconds sets StartDelayInSeconds field to given value.

### HasStartDelayInSeconds

`func (o *Task) HasStartDelayInSeconds() bool`

HasStartDelayInSeconds returns a boolean if a field has been set.

### GetRetriedTaskId

`func (o *Task) GetRetriedTaskId() string`

GetRetriedTaskId returns the RetriedTaskId field if non-nil, zero value otherwise.

### GetRetriedTaskIdOk

`func (o *Task) GetRetriedTaskIdOk() (*string, bool)`

GetRetriedTaskIdOk returns a tuple with the RetriedTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetriedTaskId

`func (o *Task) SetRetriedTaskId(v string)`

SetRetriedTaskId sets RetriedTaskId field to given value.

### HasRetriedTaskId

`func (o *Task) HasRetriedTaskId() bool`

HasRetriedTaskId returns a boolean if a field has been set.

### GetRetried

`func (o *Task) GetRetried() bool`

GetRetried returns the Retried field if non-nil, zero value otherwise.

### GetRetriedOk

`func (o *Task) GetRetriedOk() (*bool, bool)`

GetRetriedOk returns a tuple with the Retried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetried

`func (o *Task) SetRetried(v bool)`

SetRetried sets Retried field to given value.

### HasRetried

`func (o *Task) HasRetried() bool`

HasRetried returns a boolean if a field has been set.

### GetExecuted

`func (o *Task) GetExecuted() bool`

GetExecuted returns the Executed field if non-nil, zero value otherwise.

### GetExecutedOk

`func (o *Task) GetExecutedOk() (*bool, bool)`

GetExecutedOk returns a tuple with the Executed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuted

`func (o *Task) SetExecuted(v bool)`

SetExecuted sets Executed field to given value.

### HasExecuted

`func (o *Task) HasExecuted() bool`

HasExecuted returns a boolean if a field has been set.

### GetCallbackFromWorker

`func (o *Task) GetCallbackFromWorker() bool`

GetCallbackFromWorker returns the CallbackFromWorker field if non-nil, zero value otherwise.

### GetCallbackFromWorkerOk

`func (o *Task) GetCallbackFromWorkerOk() (*bool, bool)`

GetCallbackFromWorkerOk returns a tuple with the CallbackFromWorker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackFromWorker

`func (o *Task) SetCallbackFromWorker(v bool)`

SetCallbackFromWorker sets CallbackFromWorker field to given value.

### HasCallbackFromWorker

`func (o *Task) HasCallbackFromWorker() bool`

HasCallbackFromWorker returns a boolean if a field has been set.

### GetResponseTimeoutSeconds

`func (o *Task) GetResponseTimeoutSeconds() int64`

GetResponseTimeoutSeconds returns the ResponseTimeoutSeconds field if non-nil, zero value otherwise.

### GetResponseTimeoutSecondsOk

`func (o *Task) GetResponseTimeoutSecondsOk() (*int64, bool)`

GetResponseTimeoutSecondsOk returns a tuple with the ResponseTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeoutSeconds

`func (o *Task) SetResponseTimeoutSeconds(v int64)`

SetResponseTimeoutSeconds sets ResponseTimeoutSeconds field to given value.

### HasResponseTimeoutSeconds

`func (o *Task) HasResponseTimeoutSeconds() bool`

HasResponseTimeoutSeconds returns a boolean if a field has been set.

### GetWorkflowInstanceId

`func (o *Task) GetWorkflowInstanceId() string`

GetWorkflowInstanceId returns the WorkflowInstanceId field if non-nil, zero value otherwise.

### GetWorkflowInstanceIdOk

`func (o *Task) GetWorkflowInstanceIdOk() (*string, bool)`

GetWorkflowInstanceIdOk returns a tuple with the WorkflowInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInstanceId

`func (o *Task) SetWorkflowInstanceId(v string)`

SetWorkflowInstanceId sets WorkflowInstanceId field to given value.

### HasWorkflowInstanceId

`func (o *Task) HasWorkflowInstanceId() bool`

HasWorkflowInstanceId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *Task) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *Task) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *Task) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *Task) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### GetTaskId

`func (o *Task) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *Task) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *Task) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *Task) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetReasonForIncompletion

`func (o *Task) GetReasonForIncompletion() string`

GetReasonForIncompletion returns the ReasonForIncompletion field if non-nil, zero value otherwise.

### GetReasonForIncompletionOk

`func (o *Task) GetReasonForIncompletionOk() (*string, bool)`

GetReasonForIncompletionOk returns a tuple with the ReasonForIncompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForIncompletion

`func (o *Task) SetReasonForIncompletion(v string)`

SetReasonForIncompletion sets ReasonForIncompletion field to given value.

### HasReasonForIncompletion

`func (o *Task) HasReasonForIncompletion() bool`

HasReasonForIncompletion returns a boolean if a field has been set.

### GetCallbackAfterSeconds

`func (o *Task) GetCallbackAfterSeconds() int64`

GetCallbackAfterSeconds returns the CallbackAfterSeconds field if non-nil, zero value otherwise.

### GetCallbackAfterSecondsOk

`func (o *Task) GetCallbackAfterSecondsOk() (*int64, bool)`

GetCallbackAfterSecondsOk returns a tuple with the CallbackAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackAfterSeconds

`func (o *Task) SetCallbackAfterSeconds(v int64)`

SetCallbackAfterSeconds sets CallbackAfterSeconds field to given value.

### HasCallbackAfterSeconds

`func (o *Task) HasCallbackAfterSeconds() bool`

HasCallbackAfterSeconds returns a boolean if a field has been set.

### GetWorkerId

`func (o *Task) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *Task) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *Task) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *Task) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetOutputData

`func (o *Task) GetOutputData() map[string]map[string]interface{}`

GetOutputData returns the OutputData field if non-nil, zero value otherwise.

### GetOutputDataOk

`func (o *Task) GetOutputDataOk() (*map[string]map[string]interface{}, bool)`

GetOutputDataOk returns a tuple with the OutputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputData

`func (o *Task) SetOutputData(v map[string]map[string]interface{})`

SetOutputData sets OutputData field to given value.

### HasOutputData

`func (o *Task) HasOutputData() bool`

HasOutputData returns a boolean if a field has been set.

### GetWorkflowTask

`func (o *Task) GetWorkflowTask() WorkflowTask`

GetWorkflowTask returns the WorkflowTask field if non-nil, zero value otherwise.

### GetWorkflowTaskOk

`func (o *Task) GetWorkflowTaskOk() (*WorkflowTask, bool)`

GetWorkflowTaskOk returns a tuple with the WorkflowTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTask

`func (o *Task) SetWorkflowTask(v WorkflowTask)`

SetWorkflowTask sets WorkflowTask field to given value.

### HasWorkflowTask

`func (o *Task) HasWorkflowTask() bool`

HasWorkflowTask returns a boolean if a field has been set.

### GetDomain

`func (o *Task) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Task) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Task) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Task) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetRateLimitPerFrequency

`func (o *Task) GetRateLimitPerFrequency() int32`

GetRateLimitPerFrequency returns the RateLimitPerFrequency field if non-nil, zero value otherwise.

### GetRateLimitPerFrequencyOk

`func (o *Task) GetRateLimitPerFrequencyOk() (*int32, bool)`

GetRateLimitPerFrequencyOk returns a tuple with the RateLimitPerFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerFrequency

`func (o *Task) SetRateLimitPerFrequency(v int32)`

SetRateLimitPerFrequency sets RateLimitPerFrequency field to given value.

### HasRateLimitPerFrequency

`func (o *Task) HasRateLimitPerFrequency() bool`

HasRateLimitPerFrequency returns a boolean if a field has been set.

### GetRateLimitFrequencyInSeconds

`func (o *Task) GetRateLimitFrequencyInSeconds() int32`

GetRateLimitFrequencyInSeconds returns the RateLimitFrequencyInSeconds field if non-nil, zero value otherwise.

### GetRateLimitFrequencyInSecondsOk

`func (o *Task) GetRateLimitFrequencyInSecondsOk() (*int32, bool)`

GetRateLimitFrequencyInSecondsOk returns a tuple with the RateLimitFrequencyInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitFrequencyInSeconds

`func (o *Task) SetRateLimitFrequencyInSeconds(v int32)`

SetRateLimitFrequencyInSeconds sets RateLimitFrequencyInSeconds field to given value.

### HasRateLimitFrequencyInSeconds

`func (o *Task) HasRateLimitFrequencyInSeconds() bool`

HasRateLimitFrequencyInSeconds returns a boolean if a field has been set.

### GetExternalInputPayloadStoragePath

`func (o *Task) GetExternalInputPayloadStoragePath() string`

GetExternalInputPayloadStoragePath returns the ExternalInputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalInputPayloadStoragePathOk

`func (o *Task) GetExternalInputPayloadStoragePathOk() (*string, bool)`

GetExternalInputPayloadStoragePathOk returns a tuple with the ExternalInputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInputPayloadStoragePath

`func (o *Task) SetExternalInputPayloadStoragePath(v string)`

SetExternalInputPayloadStoragePath sets ExternalInputPayloadStoragePath field to given value.

### HasExternalInputPayloadStoragePath

`func (o *Task) HasExternalInputPayloadStoragePath() bool`

HasExternalInputPayloadStoragePath returns a boolean if a field has been set.

### GetExternalOutputPayloadStoragePath

`func (o *Task) GetExternalOutputPayloadStoragePath() string`

GetExternalOutputPayloadStoragePath returns the ExternalOutputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalOutputPayloadStoragePathOk

`func (o *Task) GetExternalOutputPayloadStoragePathOk() (*string, bool)`

GetExternalOutputPayloadStoragePathOk returns a tuple with the ExternalOutputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOutputPayloadStoragePath

`func (o *Task) SetExternalOutputPayloadStoragePath(v string)`

SetExternalOutputPayloadStoragePath sets ExternalOutputPayloadStoragePath field to given value.

### HasExternalOutputPayloadStoragePath

`func (o *Task) HasExternalOutputPayloadStoragePath() bool`

HasExternalOutputPayloadStoragePath returns a boolean if a field has been set.

### GetWorkflowPriority

`func (o *Task) GetWorkflowPriority() int32`

GetWorkflowPriority returns the WorkflowPriority field if non-nil, zero value otherwise.

### GetWorkflowPriorityOk

`func (o *Task) GetWorkflowPriorityOk() (*int32, bool)`

GetWorkflowPriorityOk returns a tuple with the WorkflowPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowPriority

`func (o *Task) SetWorkflowPriority(v int32)`

SetWorkflowPriority sets WorkflowPriority field to given value.

### HasWorkflowPriority

`func (o *Task) HasWorkflowPriority() bool`

HasWorkflowPriority returns a boolean if a field has been set.

### GetExecutionNameSpace

`func (o *Task) GetExecutionNameSpace() string`

GetExecutionNameSpace returns the ExecutionNameSpace field if non-nil, zero value otherwise.

### GetExecutionNameSpaceOk

`func (o *Task) GetExecutionNameSpaceOk() (*string, bool)`

GetExecutionNameSpaceOk returns a tuple with the ExecutionNameSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionNameSpace

`func (o *Task) SetExecutionNameSpace(v string)`

SetExecutionNameSpace sets ExecutionNameSpace field to given value.

### HasExecutionNameSpace

`func (o *Task) HasExecutionNameSpace() bool`

HasExecutionNameSpace returns a boolean if a field has been set.

### GetIsolationGroupId

`func (o *Task) GetIsolationGroupId() string`

GetIsolationGroupId returns the IsolationGroupId field if non-nil, zero value otherwise.

### GetIsolationGroupIdOk

`func (o *Task) GetIsolationGroupIdOk() (*string, bool)`

GetIsolationGroupIdOk returns a tuple with the IsolationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationGroupId

`func (o *Task) SetIsolationGroupId(v string)`

SetIsolationGroupId sets IsolationGroupId field to given value.

### HasIsolationGroupId

`func (o *Task) HasIsolationGroupId() bool`

HasIsolationGroupId returns a boolean if a field has been set.

### GetIteration

`func (o *Task) GetIteration() int32`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *Task) GetIterationOk() (*int32, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *Task) SetIteration(v int32)`

SetIteration sets Iteration field to given value.

### HasIteration

`func (o *Task) HasIteration() bool`

HasIteration returns a boolean if a field has been set.

### GetSubWorkflowId

`func (o *Task) GetSubWorkflowId() string`

GetSubWorkflowId returns the SubWorkflowId field if non-nil, zero value otherwise.

### GetSubWorkflowIdOk

`func (o *Task) GetSubWorkflowIdOk() (*string, bool)`

GetSubWorkflowIdOk returns a tuple with the SubWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWorkflowId

`func (o *Task) SetSubWorkflowId(v string)`

SetSubWorkflowId sets SubWorkflowId field to given value.

### HasSubWorkflowId

`func (o *Task) HasSubWorkflowId() bool`

HasSubWorkflowId returns a boolean if a field has been set.

### GetSubworkflowChanged

`func (o *Task) GetSubworkflowChanged() bool`

GetSubworkflowChanged returns the SubworkflowChanged field if non-nil, zero value otherwise.

### GetSubworkflowChangedOk

`func (o *Task) GetSubworkflowChangedOk() (*bool, bool)`

GetSubworkflowChangedOk returns a tuple with the SubworkflowChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubworkflowChanged

`func (o *Task) SetSubworkflowChanged(v bool)`

SetSubworkflowChanged sets SubworkflowChanged field to given value.

### HasSubworkflowChanged

`func (o *Task) HasSubworkflowChanged() bool`

HasSubworkflowChanged returns a boolean if a field has been set.

### GetFirstStartTime

`func (o *Task) GetFirstStartTime() int64`

GetFirstStartTime returns the FirstStartTime field if non-nil, zero value otherwise.

### GetFirstStartTimeOk

`func (o *Task) GetFirstStartTimeOk() (*int64, bool)`

GetFirstStartTimeOk returns a tuple with the FirstStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstStartTime

`func (o *Task) SetFirstStartTime(v int64)`

SetFirstStartTime sets FirstStartTime field to given value.

### HasFirstStartTime

`func (o *Task) HasFirstStartTime() bool`

HasFirstStartTime returns a boolean if a field has been set.

### GetParentTaskId

`func (o *Task) GetParentTaskId() string`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *Task) GetParentTaskIdOk() (*string, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *Task) SetParentTaskId(v string)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *Task) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetQueueWaitTime

`func (o *Task) GetQueueWaitTime() int64`

GetQueueWaitTime returns the QueueWaitTime field if non-nil, zero value otherwise.

### GetQueueWaitTimeOk

`func (o *Task) GetQueueWaitTimeOk() (*int64, bool)`

GetQueueWaitTimeOk returns a tuple with the QueueWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueWaitTime

`func (o *Task) SetQueueWaitTime(v int64)`

SetQueueWaitTime sets QueueWaitTime field to given value.

### HasQueueWaitTime

`func (o *Task) HasQueueWaitTime() bool`

HasQueueWaitTime returns a boolean if a field has been set.

### GetLoopOverTask

`func (o *Task) GetLoopOverTask() bool`

GetLoopOverTask returns the LoopOverTask field if non-nil, zero value otherwise.

### GetLoopOverTaskOk

`func (o *Task) GetLoopOverTaskOk() (*bool, bool)`

GetLoopOverTaskOk returns a tuple with the LoopOverTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopOverTask

`func (o *Task) SetLoopOverTask(v bool)`

SetLoopOverTask sets LoopOverTask field to given value.

### HasLoopOverTask

`func (o *Task) HasLoopOverTask() bool`

HasLoopOverTask returns a boolean if a field has been set.

### GetTaskDefinition

`func (o *Task) GetTaskDefinition() TaskDef`

GetTaskDefinition returns the TaskDefinition field if non-nil, zero value otherwise.

### GetTaskDefinitionOk

`func (o *Task) GetTaskDefinitionOk() (*TaskDef, bool)`

GetTaskDefinitionOk returns a tuple with the TaskDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinition

`func (o *Task) SetTaskDefinition(v TaskDef)`

SetTaskDefinition sets TaskDefinition field to given value.

### HasTaskDefinition

`func (o *Task) HasTaskDefinition() bool`

HasTaskDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


