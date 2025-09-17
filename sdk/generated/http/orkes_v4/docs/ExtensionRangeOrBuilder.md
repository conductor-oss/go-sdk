# ExtensionRangeOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to  |  | [optional] 
**DefaultInstanceForType** | Pointer to  | Simplified schema for Message (original had circular references) | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to [**ExtensionRangeOptions**](ExtensionRangeOptions.md) |  | [optional] 
**OptionsOrBuilder** | Pointer to [**ExtensionRangeOptionsOrBuilder**](ExtensionRangeOptionsOrBuilder.md) |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewExtensionRangeOrBuilder

`func NewExtensionRangeOrBuilder() *ExtensionRangeOrBuilder`

NewExtensionRangeOrBuilder instantiates a new ExtensionRangeOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionRangeOrBuilderWithDefaults

`func NewExtensionRangeOrBuilderWithDefaults() *ExtensionRangeOrBuilder`

NewExtensionRangeOrBuilderWithDefaults instantiates a new ExtensionRangeOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *ExtensionRangeOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *ExtensionRangeOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *ExtensionRangeOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *ExtensionRangeOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### SetAllFieldsNil

`func (o *ExtensionRangeOrBuilder) SetAllFieldsNil(b bool)`

 SetAllFieldsNil sets the value for AllFields to be an explicit nil

### UnsetAllFields
`func (o *ExtensionRangeOrBuilder) UnsetAllFields()`

UnsetAllFields ensures that no value is present for AllFields, not even an explicit nil
### GetDefaultInstanceForType

`func (o *ExtensionRangeOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *ExtensionRangeOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *ExtensionRangeOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *ExtensionRangeOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### SetDefaultInstanceForTypeNil

`func (o *ExtensionRangeOrBuilder) SetDefaultInstanceForTypeNil(b bool)`

 SetDefaultInstanceForTypeNil sets the value for DefaultInstanceForType to be an explicit nil

### UnsetDefaultInstanceForType
`func (o *ExtensionRangeOrBuilder) UnsetDefaultInstanceForType()`

UnsetDefaultInstanceForType ensures that no value is present for DefaultInstanceForType, not even an explicit nil
### GetDescriptorForType

`func (o *ExtensionRangeOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *ExtensionRangeOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *ExtensionRangeOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *ExtensionRangeOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetEnd

`func (o *ExtensionRangeOrBuilder) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ExtensionRangeOrBuilder) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ExtensionRangeOrBuilder) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ExtensionRangeOrBuilder) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *ExtensionRangeOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *ExtensionRangeOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *ExtensionRangeOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *ExtensionRangeOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *ExtensionRangeOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *ExtensionRangeOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *ExtensionRangeOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *ExtensionRangeOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetOptions

`func (o *ExtensionRangeOrBuilder) GetOptions() ExtensionRangeOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ExtensionRangeOrBuilder) GetOptionsOk() (*ExtensionRangeOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ExtensionRangeOrBuilder) SetOptions(v ExtensionRangeOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ExtensionRangeOrBuilder) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOptionsOrBuilder

`func (o *ExtensionRangeOrBuilder) GetOptionsOrBuilder() ExtensionRangeOptionsOrBuilder`

GetOptionsOrBuilder returns the OptionsOrBuilder field if non-nil, zero value otherwise.

### GetOptionsOrBuilderOk

`func (o *ExtensionRangeOrBuilder) GetOptionsOrBuilderOk() (*ExtensionRangeOptionsOrBuilder, bool)`

GetOptionsOrBuilderOk returns a tuple with the OptionsOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsOrBuilder

`func (o *ExtensionRangeOrBuilder) SetOptionsOrBuilder(v ExtensionRangeOptionsOrBuilder)`

SetOptionsOrBuilder sets OptionsOrBuilder field to given value.

### HasOptionsOrBuilder

`func (o *ExtensionRangeOrBuilder) HasOptionsOrBuilder() bool`

HasOptionsOrBuilder returns a boolean if a field has been set.

### GetStart

`func (o *ExtensionRangeOrBuilder) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ExtensionRangeOrBuilder) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ExtensionRangeOrBuilder) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ExtensionRangeOrBuilder) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUnknownFields

`func (o *ExtensionRangeOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *ExtensionRangeOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *ExtensionRangeOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *ExtensionRangeOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


