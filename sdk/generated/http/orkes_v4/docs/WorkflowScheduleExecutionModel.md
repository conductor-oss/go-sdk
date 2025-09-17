# WorkflowScheduleExecutionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | Pointer to **string** |  | [optional] 
**ExecutionTime** | Pointer to **int64** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**QueueMsgId** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ScheduleName** | Pointer to **string** |  | [optional] 
**ScheduledTime** | Pointer to **int64** |  | [optional] 
**StackTrace** | Pointer to **string** |  | [optional] 
**StartWorkflowRequest** | Pointer to [**StartWorkflowRequest**](StartWorkflowRequest.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**WorkflowName** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowScheduleExecutionModel

`func NewWorkflowScheduleExecutionModel() *WorkflowScheduleExecutionModel`

NewWorkflowScheduleExecutionModel instantiates a new WorkflowScheduleExecutionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowScheduleExecutionModelWithDefaults

`func NewWorkflowScheduleExecutionModelWithDefaults() *WorkflowScheduleExecutionModel`

NewWorkflowScheduleExecutionModelWithDefaults instantiates a new WorkflowScheduleExecutionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *WorkflowScheduleExecutionModel) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *WorkflowScheduleExecutionModel) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *WorkflowScheduleExecutionModel) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *WorkflowScheduleExecutionModel) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetExecutionTime

`func (o *WorkflowScheduleExecutionModel) GetExecutionTime() int64`

GetExecutionTime returns the ExecutionTime field if non-nil, zero value otherwise.

### GetExecutionTimeOk

`func (o *WorkflowScheduleExecutionModel) GetExecutionTimeOk() (*int64, bool)`

GetExecutionTimeOk returns a tuple with the ExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTime

`func (o *WorkflowScheduleExecutionModel) SetExecutionTime(v int64)`

SetExecutionTime sets ExecutionTime field to given value.

### HasExecutionTime

`func (o *WorkflowScheduleExecutionModel) HasExecutionTime() bool`

HasExecutionTime returns a boolean if a field has been set.

### GetOrgId

`func (o *WorkflowScheduleExecutionModel) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WorkflowScheduleExecutionModel) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WorkflowScheduleExecutionModel) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *WorkflowScheduleExecutionModel) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetQueueMsgId

`func (o *WorkflowScheduleExecutionModel) GetQueueMsgId() string`

GetQueueMsgId returns the QueueMsgId field if non-nil, zero value otherwise.

### GetQueueMsgIdOk

`func (o *WorkflowScheduleExecutionModel) GetQueueMsgIdOk() (*string, bool)`

GetQueueMsgIdOk returns a tuple with the QueueMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMsgId

`func (o *WorkflowScheduleExecutionModel) SetQueueMsgId(v string)`

SetQueueMsgId sets QueueMsgId field to given value.

### HasQueueMsgId

`func (o *WorkflowScheduleExecutionModel) HasQueueMsgId() bool`

HasQueueMsgId returns a boolean if a field has been set.

### GetReason

`func (o *WorkflowScheduleExecutionModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WorkflowScheduleExecutionModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WorkflowScheduleExecutionModel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WorkflowScheduleExecutionModel) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetScheduleName

`func (o *WorkflowScheduleExecutionModel) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *WorkflowScheduleExecutionModel) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *WorkflowScheduleExecutionModel) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *WorkflowScheduleExecutionModel) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### GetScheduledTime

`func (o *WorkflowScheduleExecutionModel) GetScheduledTime() int64`

GetScheduledTime returns the ScheduledTime field if non-nil, zero value otherwise.

### GetScheduledTimeOk

`func (o *WorkflowScheduleExecutionModel) GetScheduledTimeOk() (*int64, bool)`

GetScheduledTimeOk returns a tuple with the ScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTime

`func (o *WorkflowScheduleExecutionModel) SetScheduledTime(v int64)`

SetScheduledTime sets ScheduledTime field to given value.

### HasScheduledTime

`func (o *WorkflowScheduleExecutionModel) HasScheduledTime() bool`

HasScheduledTime returns a boolean if a field has been set.

### GetStackTrace

`func (o *WorkflowScheduleExecutionModel) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *WorkflowScheduleExecutionModel) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *WorkflowScheduleExecutionModel) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *WorkflowScheduleExecutionModel) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.

### GetStartWorkflowRequest

`func (o *WorkflowScheduleExecutionModel) GetStartWorkflowRequest() StartWorkflowRequest`

GetStartWorkflowRequest returns the StartWorkflowRequest field if non-nil, zero value otherwise.

### GetStartWorkflowRequestOk

`func (o *WorkflowScheduleExecutionModel) GetStartWorkflowRequestOk() (*StartWorkflowRequest, bool)`

GetStartWorkflowRequestOk returns a tuple with the StartWorkflowRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWorkflowRequest

`func (o *WorkflowScheduleExecutionModel) SetStartWorkflowRequest(v StartWorkflowRequest)`

SetStartWorkflowRequest sets StartWorkflowRequest field to given value.

### HasStartWorkflowRequest

`func (o *WorkflowScheduleExecutionModel) HasStartWorkflowRequest() bool`

HasStartWorkflowRequest returns a boolean if a field has been set.

### GetState

`func (o *WorkflowScheduleExecutionModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkflowScheduleExecutionModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkflowScheduleExecutionModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *WorkflowScheduleExecutionModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowScheduleExecutionModel) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowScheduleExecutionModel) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowScheduleExecutionModel) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowScheduleExecutionModel) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowName

`func (o *WorkflowScheduleExecutionModel) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *WorkflowScheduleExecutionModel) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *WorkflowScheduleExecutionModel) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.

### HasWorkflowName

`func (o *WorkflowScheduleExecutionModel) HasWorkflowName() bool`

HasWorkflowName returns a boolean if a field has been set.

### GetZoneId

`func (o *WorkflowScheduleExecutionModel) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *WorkflowScheduleExecutionModel) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *WorkflowScheduleExecutionModel) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *WorkflowScheduleExecutionModel) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


