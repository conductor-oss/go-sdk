# ExtendedSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewExtendedSecret

`func NewExtendedSecret() *ExtendedSecret`

NewExtendedSecret instantiates a new ExtendedSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedSecretWithDefaults

`func NewExtendedSecretWithDefaults() *ExtendedSecret`

NewExtendedSecretWithDefaults instantiates a new ExtendedSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtendedSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtendedSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtendedSecret) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExtendedSecret) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *ExtendedSecret) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExtendedSecret) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExtendedSecret) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExtendedSecret) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


