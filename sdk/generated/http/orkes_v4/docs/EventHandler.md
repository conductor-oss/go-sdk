# EventHandler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]Action**](Action.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Condition** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EvaluatorType** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewEventHandler

`func NewEventHandler() *EventHandler`

NewEventHandler instantiates a new EventHandler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHandlerWithDefaults

`func NewEventHandlerWithDefaults() *EventHandler`

NewEventHandlerWithDefaults instantiates a new EventHandler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *EventHandler) GetActions() []Action`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *EventHandler) GetActionsOk() (*[]Action, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *EventHandler) SetActions(v []Action)`

SetActions sets Actions field to given value.

### HasActions

`func (o *EventHandler) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetActive

`func (o *EventHandler) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EventHandler) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EventHandler) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EventHandler) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCondition

`func (o *EventHandler) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *EventHandler) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *EventHandler) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *EventHandler) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetCreatedBy

`func (o *EventHandler) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EventHandler) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EventHandler) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EventHandler) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *EventHandler) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventHandler) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventHandler) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventHandler) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvaluatorType

`func (o *EventHandler) GetEvaluatorType() string`

GetEvaluatorType returns the EvaluatorType field if non-nil, zero value otherwise.

### GetEvaluatorTypeOk

`func (o *EventHandler) GetEvaluatorTypeOk() (*string, bool)`

GetEvaluatorTypeOk returns a tuple with the EvaluatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorType

`func (o *EventHandler) SetEvaluatorType(v string)`

SetEvaluatorType sets EvaluatorType field to given value.

### HasEvaluatorType

`func (o *EventHandler) HasEvaluatorType() bool`

HasEvaluatorType returns a boolean if a field has been set.

### GetEvent

`func (o *EventHandler) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventHandler) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventHandler) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *EventHandler) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetName

`func (o *EventHandler) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventHandler) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventHandler) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventHandler) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *EventHandler) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *EventHandler) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *EventHandler) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *EventHandler) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetTags

`func (o *EventHandler) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EventHandler) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EventHandler) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EventHandler) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


