# SaveScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**CronExpression** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Paused** | Pointer to **bool** |  | [optional] 
**RunCatchupScheduleInstances** | Pointer to **bool** |  | [optional] 
**ScheduleEndTime** | Pointer to **int64** |  | [optional] 
**ScheduleStartTime** | Pointer to **int64** |  | [optional] 
**StartWorkflowRequest** | [**StartWorkflowRequest**](StartWorkflowRequest.md) |  | 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **string** |  | [optional] 

## Methods

### NewSaveScheduleRequest

`func NewSaveScheduleRequest(startWorkflowRequest StartWorkflowRequest, ) *SaveScheduleRequest`

NewSaveScheduleRequest instantiates a new SaveScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveScheduleRequestWithDefaults

`func NewSaveScheduleRequestWithDefaults() *SaveScheduleRequest`

NewSaveScheduleRequestWithDefaults instantiates a new SaveScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *SaveScheduleRequest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SaveScheduleRequest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SaveScheduleRequest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SaveScheduleRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCronExpression

`func (o *SaveScheduleRequest) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *SaveScheduleRequest) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *SaveScheduleRequest) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *SaveScheduleRequest) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetDescription

`func (o *SaveScheduleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaveScheduleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaveScheduleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SaveScheduleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SaveScheduleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaveScheduleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaveScheduleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaveScheduleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaused

`func (o *SaveScheduleRequest) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *SaveScheduleRequest) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *SaveScheduleRequest) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *SaveScheduleRequest) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetRunCatchupScheduleInstances

`func (o *SaveScheduleRequest) GetRunCatchupScheduleInstances() bool`

GetRunCatchupScheduleInstances returns the RunCatchupScheduleInstances field if non-nil, zero value otherwise.

### GetRunCatchupScheduleInstancesOk

`func (o *SaveScheduleRequest) GetRunCatchupScheduleInstancesOk() (*bool, bool)`

GetRunCatchupScheduleInstancesOk returns a tuple with the RunCatchupScheduleInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCatchupScheduleInstances

`func (o *SaveScheduleRequest) SetRunCatchupScheduleInstances(v bool)`

SetRunCatchupScheduleInstances sets RunCatchupScheduleInstances field to given value.

### HasRunCatchupScheduleInstances

`func (o *SaveScheduleRequest) HasRunCatchupScheduleInstances() bool`

HasRunCatchupScheduleInstances returns a boolean if a field has been set.

### GetScheduleEndTime

`func (o *SaveScheduleRequest) GetScheduleEndTime() int64`

GetScheduleEndTime returns the ScheduleEndTime field if non-nil, zero value otherwise.

### GetScheduleEndTimeOk

`func (o *SaveScheduleRequest) GetScheduleEndTimeOk() (*int64, bool)`

GetScheduleEndTimeOk returns a tuple with the ScheduleEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEndTime

`func (o *SaveScheduleRequest) SetScheduleEndTime(v int64)`

SetScheduleEndTime sets ScheduleEndTime field to given value.

### HasScheduleEndTime

`func (o *SaveScheduleRequest) HasScheduleEndTime() bool`

HasScheduleEndTime returns a boolean if a field has been set.

### GetScheduleStartTime

`func (o *SaveScheduleRequest) GetScheduleStartTime() int64`

GetScheduleStartTime returns the ScheduleStartTime field if non-nil, zero value otherwise.

### GetScheduleStartTimeOk

`func (o *SaveScheduleRequest) GetScheduleStartTimeOk() (*int64, bool)`

GetScheduleStartTimeOk returns a tuple with the ScheduleStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStartTime

`func (o *SaveScheduleRequest) SetScheduleStartTime(v int64)`

SetScheduleStartTime sets ScheduleStartTime field to given value.

### HasScheduleStartTime

`func (o *SaveScheduleRequest) HasScheduleStartTime() bool`

HasScheduleStartTime returns a boolean if a field has been set.

### GetStartWorkflowRequest

`func (o *SaveScheduleRequest) GetStartWorkflowRequest() StartWorkflowRequest`

GetStartWorkflowRequest returns the StartWorkflowRequest field if non-nil, zero value otherwise.

### GetStartWorkflowRequestOk

`func (o *SaveScheduleRequest) GetStartWorkflowRequestOk() (*StartWorkflowRequest, bool)`

GetStartWorkflowRequestOk returns a tuple with the StartWorkflowRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWorkflowRequest

`func (o *SaveScheduleRequest) SetStartWorkflowRequest(v StartWorkflowRequest)`

SetStartWorkflowRequest sets StartWorkflowRequest field to given value.


### GetUpdatedBy

`func (o *SaveScheduleRequest) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *SaveScheduleRequest) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *SaveScheduleRequest) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *SaveScheduleRequest) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetZoneId

`func (o *SaveScheduleRequest) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *SaveScheduleRequest) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *SaveScheduleRequest) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *SaveScheduleRequest) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


