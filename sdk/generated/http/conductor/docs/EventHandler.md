# EventHandler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Event** | **string** |  | 
**Condition** | Pointer to **string** |  | [optional] 
**Actions** | [**[]Action**](Action.md) |  | 
**Active** | Pointer to **bool** |  | [optional] 
**EvaluatorType** | Pointer to **string** |  | [optional] 

## Methods

### NewEventHandler

`func NewEventHandler(name string, event string, actions []Action, ) *EventHandler`

NewEventHandler instantiates a new EventHandler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHandlerWithDefaults

`func NewEventHandlerWithDefaults() *EventHandler`

NewEventHandlerWithDefaults instantiates a new EventHandler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


