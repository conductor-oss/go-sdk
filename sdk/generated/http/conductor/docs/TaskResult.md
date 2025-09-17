# TaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowInstanceId** | **string** |  | 
**TaskId** | **string** |  | 
**ReasonForIncompletion** | Pointer to **string** |  | [optional] 
**CallbackAfterSeconds** | Pointer to **int64** |  | [optional] 
**WorkerId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**OutputData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Logs** | Pointer to [**[]TaskExecLog**](TaskExecLog.md) |  | [optional] 
**ExternalOutputPayloadStoragePath** | Pointer to **string** |  | [optional] 
**SubWorkflowId** | Pointer to **string** |  | [optional] 
**ExtendLease** | Pointer to **bool** |  | [optional] 

## Methods

### NewTaskResult

`func NewTaskResult(workflowInstanceId string, taskId string, ) *TaskResult`

NewTaskResult instantiates a new TaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResultWithDefaults

`func NewTaskResultWithDefaults() *TaskResult`

NewTaskResultWithDefaults instantiates a new TaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowInstanceId

`func (o *TaskResult) GetWorkflowInstanceId() string`

GetWorkflowInstanceId returns the WorkflowInstanceId field if non-nil, zero value otherwise.

### GetWorkflowInstanceIdOk

`func (o *TaskResult) GetWorkflowInstanceIdOk() (*string, bool)`

GetWorkflowInstanceIdOk returns a tuple with the WorkflowInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInstanceId

`func (o *TaskResult) SetWorkflowInstanceId(v string)`

SetWorkflowInstanceId sets WorkflowInstanceId field to given value.


### GetTaskId

`func (o *TaskResult) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskResult) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskResult) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetReasonForIncompletion

`func (o *TaskResult) GetReasonForIncompletion() string`

GetReasonForIncompletion returns the ReasonForIncompletion field if non-nil, zero value otherwise.

### GetReasonForIncompletionOk

`func (o *TaskResult) GetReasonForIncompletionOk() (*string, bool)`

GetReasonForIncompletionOk returns a tuple with the ReasonForIncompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForIncompletion

`func (o *TaskResult) SetReasonForIncompletion(v string)`

SetReasonForIncompletion sets ReasonForIncompletion field to given value.

### HasReasonForIncompletion

`func (o *TaskResult) HasReasonForIncompletion() bool`

HasReasonForIncompletion returns a boolean if a field has been set.

### GetCallbackAfterSeconds

`func (o *TaskResult) GetCallbackAfterSeconds() int64`

GetCallbackAfterSeconds returns the CallbackAfterSeconds field if non-nil, zero value otherwise.

### GetCallbackAfterSecondsOk

`func (o *TaskResult) GetCallbackAfterSecondsOk() (*int64, bool)`

GetCallbackAfterSecondsOk returns a tuple with the CallbackAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackAfterSeconds

`func (o *TaskResult) SetCallbackAfterSeconds(v int64)`

SetCallbackAfterSeconds sets CallbackAfterSeconds field to given value.

### HasCallbackAfterSeconds

`func (o *TaskResult) HasCallbackAfterSeconds() bool`

HasCallbackAfterSeconds returns a boolean if a field has been set.

### GetWorkerId

`func (o *TaskResult) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *TaskResult) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *TaskResult) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *TaskResult) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetStatus

`func (o *TaskResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOutputData

`func (o *TaskResult) GetOutputData() map[string]map[string]interface{}`

GetOutputData returns the OutputData field if non-nil, zero value otherwise.

### GetOutputDataOk

`func (o *TaskResult) GetOutputDataOk() (*map[string]map[string]interface{}, bool)`

GetOutputDataOk returns a tuple with the OutputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputData

`func (o *TaskResult) SetOutputData(v map[string]map[string]interface{})`

SetOutputData sets OutputData field to given value.

### HasOutputData

`func (o *TaskResult) HasOutputData() bool`

HasOutputData returns a boolean if a field has been set.

### GetLogs

`func (o *TaskResult) GetLogs() []TaskExecLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *TaskResult) GetLogsOk() (*[]TaskExecLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *TaskResult) SetLogs(v []TaskExecLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *TaskResult) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetExternalOutputPayloadStoragePath

`func (o *TaskResult) GetExternalOutputPayloadStoragePath() string`

GetExternalOutputPayloadStoragePath returns the ExternalOutputPayloadStoragePath field if non-nil, zero value otherwise.

### GetExternalOutputPayloadStoragePathOk

`func (o *TaskResult) GetExternalOutputPayloadStoragePathOk() (*string, bool)`

GetExternalOutputPayloadStoragePathOk returns a tuple with the ExternalOutputPayloadStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOutputPayloadStoragePath

`func (o *TaskResult) SetExternalOutputPayloadStoragePath(v string)`

SetExternalOutputPayloadStoragePath sets ExternalOutputPayloadStoragePath field to given value.

### HasExternalOutputPayloadStoragePath

`func (o *TaskResult) HasExternalOutputPayloadStoragePath() bool`

HasExternalOutputPayloadStoragePath returns a boolean if a field has been set.

### GetSubWorkflowId

`func (o *TaskResult) GetSubWorkflowId() string`

GetSubWorkflowId returns the SubWorkflowId field if non-nil, zero value otherwise.

### GetSubWorkflowIdOk

`func (o *TaskResult) GetSubWorkflowIdOk() (*string, bool)`

GetSubWorkflowIdOk returns a tuple with the SubWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWorkflowId

`func (o *TaskResult) SetSubWorkflowId(v string)`

SetSubWorkflowId sets SubWorkflowId field to given value.

### HasSubWorkflowId

`func (o *TaskResult) HasSubWorkflowId() bool`

HasSubWorkflowId returns a boolean if a field has been set.

### GetExtendLease

`func (o *TaskResult) GetExtendLease() bool`

GetExtendLease returns the ExtendLease field if non-nil, zero value otherwise.

### GetExtendLeaseOk

`func (o *TaskResult) GetExtendLeaseOk() (*bool, bool)`

GetExtendLeaseOk returns a tuple with the ExtendLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendLease

`func (o *TaskResult) SetExtendLease(v bool)`

SetExtendLease sets ExtendLease field to given value.

### HasExtendLease

`func (o *TaskResult) HasExtendLease() bool`

HasExtendLease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


