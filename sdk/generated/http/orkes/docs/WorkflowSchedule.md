# WorkflowSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CronExpression** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Paused** | Pointer to **bool** |  | [optional] 
**PausedReason** | Pointer to **string** |  | [optional] 
**RunCatchupScheduleInstances** | Pointer to **bool** |  | [optional] 
**ScheduleEndTime** | Pointer to **int64** |  | [optional] 
**ScheduleStartTime** | Pointer to **int64** |  | [optional] 
**StartWorkflowRequest** | Pointer to [**StartWorkflowRequest**](StartWorkflowRequest.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **int64** |  | [optional] 
**ZoneId** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowSchedule

`func NewWorkflowSchedule() *WorkflowSchedule`

NewWorkflowSchedule instantiates a new WorkflowSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowScheduleWithDefaults

`func NewWorkflowScheduleWithDefaults() *WorkflowSchedule`

NewWorkflowScheduleWithDefaults instantiates a new WorkflowSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *WorkflowSchedule) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *WorkflowSchedule) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *WorkflowSchedule) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *WorkflowSchedule) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WorkflowSchedule) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkflowSchedule) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkflowSchedule) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkflowSchedule) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCronExpression

`func (o *WorkflowSchedule) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *WorkflowSchedule) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *WorkflowSchedule) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *WorkflowSchedule) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowSchedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowSchedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowSchedule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowSchedule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkflowSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaused

`func (o *WorkflowSchedule) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *WorkflowSchedule) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *WorkflowSchedule) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *WorkflowSchedule) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetPausedReason

`func (o *WorkflowSchedule) GetPausedReason() string`

GetPausedReason returns the PausedReason field if non-nil, zero value otherwise.

### GetPausedReasonOk

`func (o *WorkflowSchedule) GetPausedReasonOk() (*string, bool)`

GetPausedReasonOk returns a tuple with the PausedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedReason

`func (o *WorkflowSchedule) SetPausedReason(v string)`

SetPausedReason sets PausedReason field to given value.

### HasPausedReason

`func (o *WorkflowSchedule) HasPausedReason() bool`

HasPausedReason returns a boolean if a field has been set.

### GetRunCatchupScheduleInstances

`func (o *WorkflowSchedule) GetRunCatchupScheduleInstances() bool`

GetRunCatchupScheduleInstances returns the RunCatchupScheduleInstances field if non-nil, zero value otherwise.

### GetRunCatchupScheduleInstancesOk

`func (o *WorkflowSchedule) GetRunCatchupScheduleInstancesOk() (*bool, bool)`

GetRunCatchupScheduleInstancesOk returns a tuple with the RunCatchupScheduleInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCatchupScheduleInstances

`func (o *WorkflowSchedule) SetRunCatchupScheduleInstances(v bool)`

SetRunCatchupScheduleInstances sets RunCatchupScheduleInstances field to given value.

### HasRunCatchupScheduleInstances

`func (o *WorkflowSchedule) HasRunCatchupScheduleInstances() bool`

HasRunCatchupScheduleInstances returns a boolean if a field has been set.

### GetScheduleEndTime

`func (o *WorkflowSchedule) GetScheduleEndTime() int64`

GetScheduleEndTime returns the ScheduleEndTime field if non-nil, zero value otherwise.

### GetScheduleEndTimeOk

`func (o *WorkflowSchedule) GetScheduleEndTimeOk() (*int64, bool)`

GetScheduleEndTimeOk returns a tuple with the ScheduleEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEndTime

`func (o *WorkflowSchedule) SetScheduleEndTime(v int64)`

SetScheduleEndTime sets ScheduleEndTime field to given value.

### HasScheduleEndTime

`func (o *WorkflowSchedule) HasScheduleEndTime() bool`

HasScheduleEndTime returns a boolean if a field has been set.

### GetScheduleStartTime

`func (o *WorkflowSchedule) GetScheduleStartTime() int64`

GetScheduleStartTime returns the ScheduleStartTime field if non-nil, zero value otherwise.

### GetScheduleStartTimeOk

`func (o *WorkflowSchedule) GetScheduleStartTimeOk() (*int64, bool)`

GetScheduleStartTimeOk returns a tuple with the ScheduleStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStartTime

`func (o *WorkflowSchedule) SetScheduleStartTime(v int64)`

SetScheduleStartTime sets ScheduleStartTime field to given value.

### HasScheduleStartTime

`func (o *WorkflowSchedule) HasScheduleStartTime() bool`

HasScheduleStartTime returns a boolean if a field has been set.

### GetStartWorkflowRequest

`func (o *WorkflowSchedule) GetStartWorkflowRequest() StartWorkflowRequest`

GetStartWorkflowRequest returns the StartWorkflowRequest field if non-nil, zero value otherwise.

### GetStartWorkflowRequestOk

`func (o *WorkflowSchedule) GetStartWorkflowRequestOk() (*StartWorkflowRequest, bool)`

GetStartWorkflowRequestOk returns a tuple with the StartWorkflowRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWorkflowRequest

`func (o *WorkflowSchedule) SetStartWorkflowRequest(v StartWorkflowRequest)`

SetStartWorkflowRequest sets StartWorkflowRequest field to given value.

### HasStartWorkflowRequest

`func (o *WorkflowSchedule) HasStartWorkflowRequest() bool`

HasStartWorkflowRequest returns a boolean if a field has been set.

### GetTags

`func (o *WorkflowSchedule) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkflowSchedule) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkflowSchedule) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkflowSchedule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkflowSchedule) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkflowSchedule) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkflowSchedule) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkflowSchedule) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *WorkflowSchedule) GetUpdatedTime() int64`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *WorkflowSchedule) GetUpdatedTimeOk() (*int64, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *WorkflowSchedule) SetUpdatedTime(v int64)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *WorkflowSchedule) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetZoneId

`func (o *WorkflowSchedule) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *WorkflowSchedule) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *WorkflowSchedule) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *WorkflowSchedule) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


