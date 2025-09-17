# WorkflowScheduleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CronExpression** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Paused** | Pointer to **bool** |  | [optional] 
**PausedReason** | Pointer to **string** |  | [optional] 
**QueueMsgId** | Pointer to **string** |  | [optional] 
**RunCatchupScheduleInstances** | Pointer to **bool** |  | [optional] 
**ScheduleEndTime** | Pointer to **int64** |  | [optional] 
**ScheduleStartTime** | Pointer to **int64** |  | [optional] 
**StartWorkflowRequest** | Pointer to [**StartWorkflowRequest**](StartWorkflowRequest.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **int64** |  | [optional] 
**ZoneId** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowScheduleModel

`func NewWorkflowScheduleModel() *WorkflowScheduleModel`

NewWorkflowScheduleModel instantiates a new WorkflowScheduleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowScheduleModelWithDefaults

`func NewWorkflowScheduleModelWithDefaults() *WorkflowScheduleModel`

NewWorkflowScheduleModelWithDefaults instantiates a new WorkflowScheduleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *WorkflowScheduleModel) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *WorkflowScheduleModel) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *WorkflowScheduleModel) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *WorkflowScheduleModel) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WorkflowScheduleModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkflowScheduleModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkflowScheduleModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkflowScheduleModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCronExpression

`func (o *WorkflowScheduleModel) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *WorkflowScheduleModel) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *WorkflowScheduleModel) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *WorkflowScheduleModel) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowScheduleModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowScheduleModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowScheduleModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowScheduleModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkflowScheduleModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowScheduleModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowScheduleModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowScheduleModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *WorkflowScheduleModel) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WorkflowScheduleModel) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WorkflowScheduleModel) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *WorkflowScheduleModel) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPaused

`func (o *WorkflowScheduleModel) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *WorkflowScheduleModel) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *WorkflowScheduleModel) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *WorkflowScheduleModel) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetPausedReason

`func (o *WorkflowScheduleModel) GetPausedReason() string`

GetPausedReason returns the PausedReason field if non-nil, zero value otherwise.

### GetPausedReasonOk

`func (o *WorkflowScheduleModel) GetPausedReasonOk() (*string, bool)`

GetPausedReasonOk returns a tuple with the PausedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedReason

`func (o *WorkflowScheduleModel) SetPausedReason(v string)`

SetPausedReason sets PausedReason field to given value.

### HasPausedReason

`func (o *WorkflowScheduleModel) HasPausedReason() bool`

HasPausedReason returns a boolean if a field has been set.

### GetQueueMsgId

`func (o *WorkflowScheduleModel) GetQueueMsgId() string`

GetQueueMsgId returns the QueueMsgId field if non-nil, zero value otherwise.

### GetQueueMsgIdOk

`func (o *WorkflowScheduleModel) GetQueueMsgIdOk() (*string, bool)`

GetQueueMsgIdOk returns a tuple with the QueueMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMsgId

`func (o *WorkflowScheduleModel) SetQueueMsgId(v string)`

SetQueueMsgId sets QueueMsgId field to given value.

### HasQueueMsgId

`func (o *WorkflowScheduleModel) HasQueueMsgId() bool`

HasQueueMsgId returns a boolean if a field has been set.

### GetRunCatchupScheduleInstances

`func (o *WorkflowScheduleModel) GetRunCatchupScheduleInstances() bool`

GetRunCatchupScheduleInstances returns the RunCatchupScheduleInstances field if non-nil, zero value otherwise.

### GetRunCatchupScheduleInstancesOk

`func (o *WorkflowScheduleModel) GetRunCatchupScheduleInstancesOk() (*bool, bool)`

GetRunCatchupScheduleInstancesOk returns a tuple with the RunCatchupScheduleInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCatchupScheduleInstances

`func (o *WorkflowScheduleModel) SetRunCatchupScheduleInstances(v bool)`

SetRunCatchupScheduleInstances sets RunCatchupScheduleInstances field to given value.

### HasRunCatchupScheduleInstances

`func (o *WorkflowScheduleModel) HasRunCatchupScheduleInstances() bool`

HasRunCatchupScheduleInstances returns a boolean if a field has been set.

### GetScheduleEndTime

`func (o *WorkflowScheduleModel) GetScheduleEndTime() int64`

GetScheduleEndTime returns the ScheduleEndTime field if non-nil, zero value otherwise.

### GetScheduleEndTimeOk

`func (o *WorkflowScheduleModel) GetScheduleEndTimeOk() (*int64, bool)`

GetScheduleEndTimeOk returns a tuple with the ScheduleEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEndTime

`func (o *WorkflowScheduleModel) SetScheduleEndTime(v int64)`

SetScheduleEndTime sets ScheduleEndTime field to given value.

### HasScheduleEndTime

`func (o *WorkflowScheduleModel) HasScheduleEndTime() bool`

HasScheduleEndTime returns a boolean if a field has been set.

### GetScheduleStartTime

`func (o *WorkflowScheduleModel) GetScheduleStartTime() int64`

GetScheduleStartTime returns the ScheduleStartTime field if non-nil, zero value otherwise.

### GetScheduleStartTimeOk

`func (o *WorkflowScheduleModel) GetScheduleStartTimeOk() (*int64, bool)`

GetScheduleStartTimeOk returns a tuple with the ScheduleStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStartTime

`func (o *WorkflowScheduleModel) SetScheduleStartTime(v int64)`

SetScheduleStartTime sets ScheduleStartTime field to given value.

### HasScheduleStartTime

`func (o *WorkflowScheduleModel) HasScheduleStartTime() bool`

HasScheduleStartTime returns a boolean if a field has been set.

### GetStartWorkflowRequest

`func (o *WorkflowScheduleModel) GetStartWorkflowRequest() StartWorkflowRequest`

GetStartWorkflowRequest returns the StartWorkflowRequest field if non-nil, zero value otherwise.

### GetStartWorkflowRequestOk

`func (o *WorkflowScheduleModel) GetStartWorkflowRequestOk() (*StartWorkflowRequest, bool)`

GetStartWorkflowRequestOk returns a tuple with the StartWorkflowRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWorkflowRequest

`func (o *WorkflowScheduleModel) SetStartWorkflowRequest(v StartWorkflowRequest)`

SetStartWorkflowRequest sets StartWorkflowRequest field to given value.

### HasStartWorkflowRequest

`func (o *WorkflowScheduleModel) HasStartWorkflowRequest() bool`

HasStartWorkflowRequest returns a boolean if a field has been set.

### GetTags

`func (o *WorkflowScheduleModel) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkflowScheduleModel) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkflowScheduleModel) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkflowScheduleModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkflowScheduleModel) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkflowScheduleModel) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkflowScheduleModel) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkflowScheduleModel) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *WorkflowScheduleModel) GetUpdatedTime() int64`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *WorkflowScheduleModel) GetUpdatedTimeOk() (*int64, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *WorkflowScheduleModel) SetUpdatedTime(v int64)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *WorkflowScheduleModel) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetZoneId

`func (o *WorkflowScheduleModel) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *WorkflowScheduleModel) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *WorkflowScheduleModel) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *WorkflowScheduleModel) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


