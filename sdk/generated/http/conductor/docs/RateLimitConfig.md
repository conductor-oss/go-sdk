# RateLimitConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RateLimitKey** | Pointer to **string** |  | [optional] 
**ConcurrentExecLimit** | Pointer to **int32** |  | [optional] 

## Methods

### NewRateLimitConfig

`func NewRateLimitConfig() *RateLimitConfig`

NewRateLimitConfig instantiates a new RateLimitConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitConfigWithDefaults

`func NewRateLimitConfigWithDefaults() *RateLimitConfig`

NewRateLimitConfigWithDefaults instantiates a new RateLimitConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRateLimitKey

`func (o *RateLimitConfig) GetRateLimitKey() string`

GetRateLimitKey returns the RateLimitKey field if non-nil, zero value otherwise.

### GetRateLimitKeyOk

`func (o *RateLimitConfig) GetRateLimitKeyOk() (*string, bool)`

GetRateLimitKeyOk returns a tuple with the RateLimitKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitKey

`func (o *RateLimitConfig) SetRateLimitKey(v string)`

SetRateLimitKey sets RateLimitKey field to given value.

### HasRateLimitKey

`func (o *RateLimitConfig) HasRateLimitKey() bool`

HasRateLimitKey returns a boolean if a field has been set.

### GetConcurrentExecLimit

`func (o *RateLimitConfig) GetConcurrentExecLimit() int32`

GetConcurrentExecLimit returns the ConcurrentExecLimit field if non-nil, zero value otherwise.

### GetConcurrentExecLimitOk

`func (o *RateLimitConfig) GetConcurrentExecLimitOk() (*int32, bool)`

GetConcurrentExecLimitOk returns a tuple with the ConcurrentExecLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentExecLimit

`func (o *RateLimitConfig) SetConcurrentExecLimit(v int32)`

SetConcurrentExecLimit sets ConcurrentExecLimit field to given value.

### HasConcurrentExecLimit

`func (o *RateLimitConfig) HasConcurrentExecLimit() bool`

HasConcurrentExecLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


