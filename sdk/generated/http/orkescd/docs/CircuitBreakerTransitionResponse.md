# CircuitBreakerTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentState** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**PreviousState** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**TransitionTimestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewCircuitBreakerTransitionResponse

`func NewCircuitBreakerTransitionResponse() *CircuitBreakerTransitionResponse`

NewCircuitBreakerTransitionResponse instantiates a new CircuitBreakerTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitBreakerTransitionResponseWithDefaults

`func NewCircuitBreakerTransitionResponseWithDefaults() *CircuitBreakerTransitionResponse`

NewCircuitBreakerTransitionResponseWithDefaults instantiates a new CircuitBreakerTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentState

`func (o *CircuitBreakerTransitionResponse) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *CircuitBreakerTransitionResponse) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *CircuitBreakerTransitionResponse) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *CircuitBreakerTransitionResponse) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetMessage

`func (o *CircuitBreakerTransitionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CircuitBreakerTransitionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CircuitBreakerTransitionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CircuitBreakerTransitionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPreviousState

`func (o *CircuitBreakerTransitionResponse) GetPreviousState() string`

GetPreviousState returns the PreviousState field if non-nil, zero value otherwise.

### GetPreviousStateOk

`func (o *CircuitBreakerTransitionResponse) GetPreviousStateOk() (*string, bool)`

GetPreviousStateOk returns a tuple with the PreviousState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousState

`func (o *CircuitBreakerTransitionResponse) SetPreviousState(v string)`

SetPreviousState sets PreviousState field to given value.

### HasPreviousState

`func (o *CircuitBreakerTransitionResponse) HasPreviousState() bool`

HasPreviousState returns a boolean if a field has been set.

### GetService

`func (o *CircuitBreakerTransitionResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CircuitBreakerTransitionResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CircuitBreakerTransitionResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CircuitBreakerTransitionResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTransitionTimestamp

`func (o *CircuitBreakerTransitionResponse) GetTransitionTimestamp() int64`

GetTransitionTimestamp returns the TransitionTimestamp field if non-nil, zero value otherwise.

### GetTransitionTimestampOk

`func (o *CircuitBreakerTransitionResponse) GetTransitionTimestampOk() (*int64, bool)`

GetTransitionTimestampOk returns a tuple with the TransitionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionTimestamp

`func (o *CircuitBreakerTransitionResponse) SetTransitionTimestamp(v int64)`

SetTransitionTimestamp sets TransitionTimestamp field to given value.

### HasTransitionTimestamp

`func (o *CircuitBreakerTransitionResponse) HasTransitionTimestamp() bool`

HasTransitionTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


