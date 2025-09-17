# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CircuitBreakerConfig** | Pointer to [**OrkesCircuitBreakerConfig**](OrkesCircuitBreakerConfig.md) |  | [optional] 

## Methods

### NewConfig

`func NewConfig() *Config`

NewConfig instantiates a new Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWithDefaults

`func NewConfigWithDefaults() *Config`

NewConfigWithDefaults instantiates a new Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuitBreakerConfig

`func (o *Config) GetCircuitBreakerConfig() OrkesCircuitBreakerConfig`

GetCircuitBreakerConfig returns the CircuitBreakerConfig field if non-nil, zero value otherwise.

### GetCircuitBreakerConfigOk

`func (o *Config) GetCircuitBreakerConfigOk() (*OrkesCircuitBreakerConfig, bool)`

GetCircuitBreakerConfigOk returns a tuple with the CircuitBreakerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreakerConfig

`func (o *Config) SetCircuitBreakerConfig(v OrkesCircuitBreakerConfig)`

SetCircuitBreakerConfig sets CircuitBreakerConfig field to given value.

### HasCircuitBreakerConfig

`func (o *Config) HasCircuitBreakerConfig() bool`

HasCircuitBreakerConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


