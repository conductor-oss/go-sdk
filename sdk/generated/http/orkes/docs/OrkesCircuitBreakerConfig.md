# OrkesCircuitBreakerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticTransitionFromOpenToHalfOpenEnabled** | Pointer to **bool** |  | [optional] 
**FailureRateThreshold** | Pointer to **float32** |  | [optional] 
**MaxWaitDurationInHalfOpenState** | Pointer to **int64** |  | [optional] 
**MinimumNumberOfCalls** | Pointer to **int32** |  | [optional] 
**PermittedNumberOfCallsInHalfOpenState** | Pointer to **int32** |  | [optional] 
**SlidingWindowSize** | Pointer to **int32** |  | [optional] 
**SlowCallDurationThreshold** | Pointer to **int64** |  | [optional] 
**SlowCallRateThreshold** | Pointer to **float32** |  | [optional] 
**WaitDurationInOpenState** | Pointer to **int64** |  | [optional] 

## Methods

### NewOrkesCircuitBreakerConfig

`func NewOrkesCircuitBreakerConfig() *OrkesCircuitBreakerConfig`

NewOrkesCircuitBreakerConfig instantiates a new OrkesCircuitBreakerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrkesCircuitBreakerConfigWithDefaults

`func NewOrkesCircuitBreakerConfigWithDefaults() *OrkesCircuitBreakerConfig`

NewOrkesCircuitBreakerConfigWithDefaults instantiates a new OrkesCircuitBreakerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticTransitionFromOpenToHalfOpenEnabled

`func (o *OrkesCircuitBreakerConfig) GetAutomaticTransitionFromOpenToHalfOpenEnabled() bool`

GetAutomaticTransitionFromOpenToHalfOpenEnabled returns the AutomaticTransitionFromOpenToHalfOpenEnabled field if non-nil, zero value otherwise.

### GetAutomaticTransitionFromOpenToHalfOpenEnabledOk

`func (o *OrkesCircuitBreakerConfig) GetAutomaticTransitionFromOpenToHalfOpenEnabledOk() (*bool, bool)`

GetAutomaticTransitionFromOpenToHalfOpenEnabledOk returns a tuple with the AutomaticTransitionFromOpenToHalfOpenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticTransitionFromOpenToHalfOpenEnabled

`func (o *OrkesCircuitBreakerConfig) SetAutomaticTransitionFromOpenToHalfOpenEnabled(v bool)`

SetAutomaticTransitionFromOpenToHalfOpenEnabled sets AutomaticTransitionFromOpenToHalfOpenEnabled field to given value.

### HasAutomaticTransitionFromOpenToHalfOpenEnabled

`func (o *OrkesCircuitBreakerConfig) HasAutomaticTransitionFromOpenToHalfOpenEnabled() bool`

HasAutomaticTransitionFromOpenToHalfOpenEnabled returns a boolean if a field has been set.

### GetFailureRateThreshold

`func (o *OrkesCircuitBreakerConfig) GetFailureRateThreshold() float32`

GetFailureRateThreshold returns the FailureRateThreshold field if non-nil, zero value otherwise.

### GetFailureRateThresholdOk

`func (o *OrkesCircuitBreakerConfig) GetFailureRateThresholdOk() (*float32, bool)`

GetFailureRateThresholdOk returns a tuple with the FailureRateThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureRateThreshold

`func (o *OrkesCircuitBreakerConfig) SetFailureRateThreshold(v float32)`

SetFailureRateThreshold sets FailureRateThreshold field to given value.

### HasFailureRateThreshold

`func (o *OrkesCircuitBreakerConfig) HasFailureRateThreshold() bool`

HasFailureRateThreshold returns a boolean if a field has been set.

### GetMaxWaitDurationInHalfOpenState

`func (o *OrkesCircuitBreakerConfig) GetMaxWaitDurationInHalfOpenState() int64`

GetMaxWaitDurationInHalfOpenState returns the MaxWaitDurationInHalfOpenState field if non-nil, zero value otherwise.

### GetMaxWaitDurationInHalfOpenStateOk

`func (o *OrkesCircuitBreakerConfig) GetMaxWaitDurationInHalfOpenStateOk() (*int64, bool)`

GetMaxWaitDurationInHalfOpenStateOk returns a tuple with the MaxWaitDurationInHalfOpenState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWaitDurationInHalfOpenState

`func (o *OrkesCircuitBreakerConfig) SetMaxWaitDurationInHalfOpenState(v int64)`

SetMaxWaitDurationInHalfOpenState sets MaxWaitDurationInHalfOpenState field to given value.

### HasMaxWaitDurationInHalfOpenState

`func (o *OrkesCircuitBreakerConfig) HasMaxWaitDurationInHalfOpenState() bool`

HasMaxWaitDurationInHalfOpenState returns a boolean if a field has been set.

### GetMinimumNumberOfCalls

`func (o *OrkesCircuitBreakerConfig) GetMinimumNumberOfCalls() int32`

GetMinimumNumberOfCalls returns the MinimumNumberOfCalls field if non-nil, zero value otherwise.

### GetMinimumNumberOfCallsOk

`func (o *OrkesCircuitBreakerConfig) GetMinimumNumberOfCallsOk() (*int32, bool)`

GetMinimumNumberOfCallsOk returns a tuple with the MinimumNumberOfCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNumberOfCalls

`func (o *OrkesCircuitBreakerConfig) SetMinimumNumberOfCalls(v int32)`

SetMinimumNumberOfCalls sets MinimumNumberOfCalls field to given value.

### HasMinimumNumberOfCalls

`func (o *OrkesCircuitBreakerConfig) HasMinimumNumberOfCalls() bool`

HasMinimumNumberOfCalls returns a boolean if a field has been set.

### GetPermittedNumberOfCallsInHalfOpenState

`func (o *OrkesCircuitBreakerConfig) GetPermittedNumberOfCallsInHalfOpenState() int32`

GetPermittedNumberOfCallsInHalfOpenState returns the PermittedNumberOfCallsInHalfOpenState field if non-nil, zero value otherwise.

### GetPermittedNumberOfCallsInHalfOpenStateOk

`func (o *OrkesCircuitBreakerConfig) GetPermittedNumberOfCallsInHalfOpenStateOk() (*int32, bool)`

GetPermittedNumberOfCallsInHalfOpenStateOk returns a tuple with the PermittedNumberOfCallsInHalfOpenState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermittedNumberOfCallsInHalfOpenState

`func (o *OrkesCircuitBreakerConfig) SetPermittedNumberOfCallsInHalfOpenState(v int32)`

SetPermittedNumberOfCallsInHalfOpenState sets PermittedNumberOfCallsInHalfOpenState field to given value.

### HasPermittedNumberOfCallsInHalfOpenState

`func (o *OrkesCircuitBreakerConfig) HasPermittedNumberOfCallsInHalfOpenState() bool`

HasPermittedNumberOfCallsInHalfOpenState returns a boolean if a field has been set.

### GetSlidingWindowSize

`func (o *OrkesCircuitBreakerConfig) GetSlidingWindowSize() int32`

GetSlidingWindowSize returns the SlidingWindowSize field if non-nil, zero value otherwise.

### GetSlidingWindowSizeOk

`func (o *OrkesCircuitBreakerConfig) GetSlidingWindowSizeOk() (*int32, bool)`

GetSlidingWindowSizeOk returns a tuple with the SlidingWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlidingWindowSize

`func (o *OrkesCircuitBreakerConfig) SetSlidingWindowSize(v int32)`

SetSlidingWindowSize sets SlidingWindowSize field to given value.

### HasSlidingWindowSize

`func (o *OrkesCircuitBreakerConfig) HasSlidingWindowSize() bool`

HasSlidingWindowSize returns a boolean if a field has been set.

### GetSlowCallDurationThreshold

`func (o *OrkesCircuitBreakerConfig) GetSlowCallDurationThreshold() int64`

GetSlowCallDurationThreshold returns the SlowCallDurationThreshold field if non-nil, zero value otherwise.

### GetSlowCallDurationThresholdOk

`func (o *OrkesCircuitBreakerConfig) GetSlowCallDurationThresholdOk() (*int64, bool)`

GetSlowCallDurationThresholdOk returns a tuple with the SlowCallDurationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowCallDurationThreshold

`func (o *OrkesCircuitBreakerConfig) SetSlowCallDurationThreshold(v int64)`

SetSlowCallDurationThreshold sets SlowCallDurationThreshold field to given value.

### HasSlowCallDurationThreshold

`func (o *OrkesCircuitBreakerConfig) HasSlowCallDurationThreshold() bool`

HasSlowCallDurationThreshold returns a boolean if a field has been set.

### GetSlowCallRateThreshold

`func (o *OrkesCircuitBreakerConfig) GetSlowCallRateThreshold() float32`

GetSlowCallRateThreshold returns the SlowCallRateThreshold field if non-nil, zero value otherwise.

### GetSlowCallRateThresholdOk

`func (o *OrkesCircuitBreakerConfig) GetSlowCallRateThresholdOk() (*float32, bool)`

GetSlowCallRateThresholdOk returns a tuple with the SlowCallRateThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowCallRateThreshold

`func (o *OrkesCircuitBreakerConfig) SetSlowCallRateThreshold(v float32)`

SetSlowCallRateThreshold sets SlowCallRateThreshold field to given value.

### HasSlowCallRateThreshold

`func (o *OrkesCircuitBreakerConfig) HasSlowCallRateThreshold() bool`

HasSlowCallRateThreshold returns a boolean if a field has been set.

### GetWaitDurationInOpenState

`func (o *OrkesCircuitBreakerConfig) GetWaitDurationInOpenState() int64`

GetWaitDurationInOpenState returns the WaitDurationInOpenState field if non-nil, zero value otherwise.

### GetWaitDurationInOpenStateOk

`func (o *OrkesCircuitBreakerConfig) GetWaitDurationInOpenStateOk() (*int64, bool)`

GetWaitDurationInOpenStateOk returns a tuple with the WaitDurationInOpenState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDurationInOpenState

`func (o *OrkesCircuitBreakerConfig) SetWaitDurationInOpenState(v int64)`

SetWaitDurationInOpenState sets WaitDurationInOpenState field to given value.

### HasWaitDurationInOpenState

`func (o *OrkesCircuitBreakerConfig) HasWaitDurationInOpenState() bool`

HasWaitDurationInOpenState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


