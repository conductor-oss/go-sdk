# CacheConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**TtlInSecond** | Pointer to **int32** |  | [optional] 

## Methods

### NewCacheConfig

`func NewCacheConfig() *CacheConfig`

NewCacheConfig instantiates a new CacheConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheConfigWithDefaults

`func NewCacheConfigWithDefaults() *CacheConfig`

NewCacheConfigWithDefaults instantiates a new CacheConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CacheConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CacheConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CacheConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CacheConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetTtlInSecond

`func (o *CacheConfig) GetTtlInSecond() int32`

GetTtlInSecond returns the TtlInSecond field if non-nil, zero value otherwise.

### GetTtlInSecondOk

`func (o *CacheConfig) GetTtlInSecondOk() (*int32, bool)`

GetTtlInSecondOk returns a tuple with the TtlInSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlInSecond

`func (o *CacheConfig) SetTtlInSecond(v int32)`

SetTtlInSecond sets TtlInSecond field to given value.

### HasTtlInSecond

`func (o *CacheConfig) HasTtlInSecond() bool`

HasTtlInSecond returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


