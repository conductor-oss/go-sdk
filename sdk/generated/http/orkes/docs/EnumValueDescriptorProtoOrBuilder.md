# EnumValueDescriptorProtoOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**DefaultInstanceForType** | Pointer to **map[string]interface{}** | Simplified schema for Message (original had circular references) | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**EnumValueOptions**](EnumValueOptions.md) |  | [optional] 
**OptionsOrBuilder** | Pointer to [**EnumValueOptionsOrBuilder**](EnumValueOptionsOrBuilder.md) |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewEnumValueDescriptorProtoOrBuilder

`func NewEnumValueDescriptorProtoOrBuilder() *EnumValueDescriptorProtoOrBuilder`

NewEnumValueDescriptorProtoOrBuilder instantiates a new EnumValueDescriptorProtoOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumValueDescriptorProtoOrBuilderWithDefaults

`func NewEnumValueDescriptorProtoOrBuilderWithDefaults() *EnumValueDescriptorProtoOrBuilder`

NewEnumValueDescriptorProtoOrBuilderWithDefaults instantiates a new EnumValueDescriptorProtoOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *EnumValueDescriptorProtoOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *EnumValueDescriptorProtoOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *EnumValueDescriptorProtoOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *EnumValueDescriptorProtoOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *EnumValueDescriptorProtoOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *EnumValueDescriptorProtoOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *EnumValueDescriptorProtoOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *EnumValueDescriptorProtoOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *EnumValueDescriptorProtoOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *EnumValueDescriptorProtoOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *EnumValueDescriptorProtoOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *EnumValueDescriptorProtoOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *EnumValueDescriptorProtoOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *EnumValueDescriptorProtoOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *EnumValueDescriptorProtoOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *EnumValueDescriptorProtoOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *EnumValueDescriptorProtoOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *EnumValueDescriptorProtoOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *EnumValueDescriptorProtoOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *EnumValueDescriptorProtoOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetName

`func (o *EnumValueDescriptorProtoOrBuilder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnumValueDescriptorProtoOrBuilder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnumValueDescriptorProtoOrBuilder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnumValueDescriptorProtoOrBuilder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameBytes

`func (o *EnumValueDescriptorProtoOrBuilder) GetNameBytes() ByteString`

GetNameBytes returns the NameBytes field if non-nil, zero value otherwise.

### GetNameBytesOk

`func (o *EnumValueDescriptorProtoOrBuilder) GetNameBytesOk() (*ByteString, bool)`

GetNameBytesOk returns a tuple with the NameBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameBytes

`func (o *EnumValueDescriptorProtoOrBuilder) SetNameBytes(v ByteString)`

SetNameBytes sets NameBytes field to given value.

### HasNameBytes

`func (o *EnumValueDescriptorProtoOrBuilder) HasNameBytes() bool`

HasNameBytes returns a boolean if a field has been set.

### GetNumber

`func (o *EnumValueDescriptorProtoOrBuilder) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *EnumValueDescriptorProtoOrBuilder) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *EnumValueDescriptorProtoOrBuilder) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *EnumValueDescriptorProtoOrBuilder) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetOptions

`func (o *EnumValueDescriptorProtoOrBuilder) GetOptions() EnumValueOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *EnumValueDescriptorProtoOrBuilder) GetOptionsOk() (*EnumValueOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *EnumValueDescriptorProtoOrBuilder) SetOptions(v EnumValueOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *EnumValueDescriptorProtoOrBuilder) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOptionsOrBuilder

`func (o *EnumValueDescriptorProtoOrBuilder) GetOptionsOrBuilder() EnumValueOptionsOrBuilder`

GetOptionsOrBuilder returns the OptionsOrBuilder field if non-nil, zero value otherwise.

### GetOptionsOrBuilderOk

`func (o *EnumValueDescriptorProtoOrBuilder) GetOptionsOrBuilderOk() (*EnumValueOptionsOrBuilder, bool)`

GetOptionsOrBuilderOk returns a tuple with the OptionsOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsOrBuilder

`func (o *EnumValueDescriptorProtoOrBuilder) SetOptionsOrBuilder(v EnumValueOptionsOrBuilder)`

SetOptionsOrBuilder sets OptionsOrBuilder field to given value.

### HasOptionsOrBuilder

`func (o *EnumValueDescriptorProtoOrBuilder) HasOptionsOrBuilder() bool`

HasOptionsOrBuilder returns a boolean if a field has been set.

### GetUnknownFields

`func (o *EnumValueDescriptorProtoOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *EnumValueDescriptorProtoOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *EnumValueDescriptorProtoOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *EnumValueDescriptorProtoOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


