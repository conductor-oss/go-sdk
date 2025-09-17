# HumanTaskAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to [**HumanTaskUser**](HumanTaskUser.md) |  | [optional] 
**SlaMinutes** | Pointer to **int64** |  | [optional] 

## Methods

### NewHumanTaskAssignment

`func NewHumanTaskAssignment() *HumanTaskAssignment`

NewHumanTaskAssignment instantiates a new HumanTaskAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHumanTaskAssignmentWithDefaults

`func NewHumanTaskAssignmentWithDefaults() *HumanTaskAssignment`

NewHumanTaskAssignmentWithDefaults instantiates a new HumanTaskAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *HumanTaskAssignment) GetAssignee() HumanTaskUser`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *HumanTaskAssignment) GetAssigneeOk() (*HumanTaskUser, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *HumanTaskAssignment) SetAssignee(v HumanTaskUser)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *HumanTaskAssignment) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetSlaMinutes

`func (o *HumanTaskAssignment) GetSlaMinutes() int64`

GetSlaMinutes returns the SlaMinutes field if non-nil, zero value otherwise.

### GetSlaMinutesOk

`func (o *HumanTaskAssignment) GetSlaMinutesOk() (*int64, bool)`

GetSlaMinutesOk returns a tuple with the SlaMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaMinutes

`func (o *HumanTaskAssignment) SetSlaMinutes(v int64)`

SetSlaMinutes sets SlaMinutes field to given value.

### HasSlaMinutes

`func (o *HumanTaskAssignment) HasSlaMinutes() bool`

HasSlaMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


