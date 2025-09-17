# EnumReservedRangeOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to  |  | [optional] 
**DefaultInstanceForType** | Pointer to  | Simplified schema for Message (original had circular references) | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewEnumReservedRangeOrBuilder

`func NewEnumReservedRangeOrBuilder() *EnumReservedRangeOrBuilder`

NewEnumReservedRangeOrBuilder instantiates a new EnumReservedRangeOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumReservedRangeOrBuilderWithDefaults

`func NewEnumReservedRangeOrBuilderWithDefaults() *EnumReservedRangeOrBuilder`

NewEnumReservedRangeOrBuilderWithDefaults instantiates a new EnumReservedRangeOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *EnumReservedRangeOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *EnumReservedRangeOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *EnumReservedRangeOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *EnumReservedRangeOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### SetAllFieldsNil

`func (o *EnumReservedRangeOrBuilder) SetAllFieldsNil(b bool)`

 SetAllFieldsNil sets the value for AllFields to be an explicit nil

### UnsetAllFields
`func (o *EnumReservedRangeOrBuilder) UnsetAllFields()`

UnsetAllFields ensures that no value is present for AllFields, not even an explicit nil
### GetDefaultInstanceForType

`func (o *EnumReservedRangeOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *EnumReservedRangeOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *EnumReservedRangeOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *EnumReservedRangeOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### SetDefaultInstanceForTypeNil

`func (o *EnumReservedRangeOrBuilder) SetDefaultInstanceForTypeNil(b bool)`

 SetDefaultInstanceForTypeNil sets the value for DefaultInstanceForType to be an explicit nil

### UnsetDefaultInstanceForType
`func (o *EnumReservedRangeOrBuilder) UnsetDefaultInstanceForType()`

UnsetDefaultInstanceForType ensures that no value is present for DefaultInstanceForType, not even an explicit nil
### GetDescriptorForType

`func (o *EnumReservedRangeOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *EnumReservedRangeOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *EnumReservedRangeOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *EnumReservedRangeOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetEnd

`func (o *EnumReservedRangeOrBuilder) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *EnumReservedRangeOrBuilder) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *EnumReservedRangeOrBuilder) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *EnumReservedRangeOrBuilder) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *EnumReservedRangeOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *EnumReservedRangeOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *EnumReservedRangeOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *EnumReservedRangeOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *EnumReservedRangeOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *EnumReservedRangeOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *EnumReservedRangeOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *EnumReservedRangeOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetStart

`func (o *EnumReservedRangeOrBuilder) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EnumReservedRangeOrBuilder) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EnumReservedRangeOrBuilder) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *EnumReservedRangeOrBuilder) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUnknownFields

`func (o *EnumReservedRangeOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *EnumReservedRangeOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *EnumReservedRangeOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *EnumReservedRangeOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


