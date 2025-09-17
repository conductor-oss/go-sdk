# GenerateTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | Pointer to **string** |  | [optional] 
**KeySecret** | Pointer to **string** |  | [optional] 

## Methods

### NewGenerateTokenRequest

`func NewGenerateTokenRequest() *GenerateTokenRequest`

NewGenerateTokenRequest instantiates a new GenerateTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateTokenRequestWithDefaults

`func NewGenerateTokenRequestWithDefaults() *GenerateTokenRequest`

NewGenerateTokenRequestWithDefaults instantiates a new GenerateTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *GenerateTokenRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *GenerateTokenRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *GenerateTokenRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *GenerateTokenRequest) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetKeySecret

`func (o *GenerateTokenRequest) GetKeySecret() string`

GetKeySecret returns the KeySecret field if non-nil, zero value otherwise.

### GetKeySecretOk

`func (o *GenerateTokenRequest) GetKeySecretOk() (*string, bool)`

GetKeySecretOk returns a tuple with the KeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySecret

`func (o *GenerateTokenRequest) SetKeySecret(v string)`

SetKeySecret sets KeySecret field to given value.

### HasKeySecret

`func (o *GenerateTokenRequest) HasKeySecret() bool`

HasKeySecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


