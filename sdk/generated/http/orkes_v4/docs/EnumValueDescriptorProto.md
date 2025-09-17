# EnumValueDescriptorProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to  |  | [optional] 
**DefaultInstanceForType** | Pointer to [**EnumValueDescriptorProto**](EnumValueDescriptorProto.md) |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**MemoizedSerializedSize** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**EnumValueOptions**](EnumValueOptions.md) |  | [optional] 
**OptionsOrBuilder** | Pointer to [**EnumValueOptionsOrBuilder**](EnumValueOptionsOrBuilder.md) |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** |  | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewEnumValueDescriptorProto

`func NewEnumValueDescriptorProto() *EnumValueDescriptorProto`

NewEnumValueDescriptorProto instantiates a new EnumValueDescriptorProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumValueDescriptorProtoWithDefaults

`func NewEnumValueDescriptorProtoWithDefaults() *EnumValueDescriptorProto`

NewEnumValueDescriptorProtoWithDefaults instantiates a new EnumValueDescriptorProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *EnumValueDescriptorProto) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *EnumValueDescriptorProto) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *EnumValueDescriptorProto) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *EnumValueDescriptorProto) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### SetAllFieldsNil

`func (o *EnumValueDescriptorProto) SetAllFieldsNil(b bool)`

 SetAllFieldsNil sets the value for AllFields to be an explicit nil

### UnsetAllFields
`func (o *EnumValueDescriptorProto) UnsetAllFields()`

UnsetAllFields ensures that no value is present for AllFields, not even an explicit nil
### GetDefaultInstanceForType

`func (o *EnumValueDescriptorProto) GetDefaultInstanceForType() EnumValueDescriptorProto`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *EnumValueDescriptorProto) GetDefaultInstanceForTypeOk() (*EnumValueDescriptorProto, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *EnumValueDescriptorProto) SetDefaultInstanceForType(v EnumValueDescriptorProto)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *EnumValueDescriptorProto) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *EnumValueDescriptorProto) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *EnumValueDescriptorProto) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *EnumValueDescriptorProto) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *EnumValueDescriptorProto) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *EnumValueDescriptorProto) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *EnumValueDescriptorProto) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *EnumValueDescriptorProto) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *EnumValueDescriptorProto) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *EnumValueDescriptorProto) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *EnumValueDescriptorProto) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *EnumValueDescriptorProto) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *EnumValueDescriptorProto) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetMemoizedSerializedSize

`func (o *EnumValueDescriptorProto) GetMemoizedSerializedSize() int32`

GetMemoizedSerializedSize returns the MemoizedSerializedSize field if non-nil, zero value otherwise.

### GetMemoizedSerializedSizeOk

`func (o *EnumValueDescriptorProto) GetMemoizedSerializedSizeOk() (*int32, bool)`

GetMemoizedSerializedSizeOk returns a tuple with the MemoizedSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizedSerializedSize

`func (o *EnumValueDescriptorProto) SetMemoizedSerializedSize(v int32)`

SetMemoizedSerializedSize sets MemoizedSerializedSize field to given value.

### HasMemoizedSerializedSize

`func (o *EnumValueDescriptorProto) HasMemoizedSerializedSize() bool`

HasMemoizedSerializedSize returns a boolean if a field has been set.

### GetName

`func (o *EnumValueDescriptorProto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnumValueDescriptorProto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnumValueDescriptorProto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnumValueDescriptorProto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameBytes

`func (o *EnumValueDescriptorProto) GetNameBytes() ByteString`

GetNameBytes returns the NameBytes field if non-nil, zero value otherwise.

### GetNameBytesOk

`func (o *EnumValueDescriptorProto) GetNameBytesOk() (*ByteString, bool)`

GetNameBytesOk returns a tuple with the NameBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameBytes

`func (o *EnumValueDescriptorProto) SetNameBytes(v ByteString)`

SetNameBytes sets NameBytes field to given value.

### HasNameBytes

`func (o *EnumValueDescriptorProto) HasNameBytes() bool`

HasNameBytes returns a boolean if a field has been set.

### GetNumber

`func (o *EnumValueDescriptorProto) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *EnumValueDescriptorProto) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *EnumValueDescriptorProto) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *EnumValueDescriptorProto) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetOptions

`func (o *EnumValueDescriptorProto) GetOptions() EnumValueOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *EnumValueDescriptorProto) GetOptionsOk() (*EnumValueOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *EnumValueDescriptorProto) SetOptions(v EnumValueOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *EnumValueDescriptorProto) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOptionsOrBuilder

`func (o *EnumValueDescriptorProto) GetOptionsOrBuilder() EnumValueOptionsOrBuilder`

GetOptionsOrBuilder returns the OptionsOrBuilder field if non-nil, zero value otherwise.

### GetOptionsOrBuilderOk

`func (o *EnumValueDescriptorProto) GetOptionsOrBuilderOk() (*EnumValueOptionsOrBuilder, bool)`

GetOptionsOrBuilderOk returns a tuple with the OptionsOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsOrBuilder

`func (o *EnumValueDescriptorProto) SetOptionsOrBuilder(v EnumValueOptionsOrBuilder)`

SetOptionsOrBuilder sets OptionsOrBuilder field to given value.

### HasOptionsOrBuilder

`func (o *EnumValueDescriptorProto) HasOptionsOrBuilder() bool`

HasOptionsOrBuilder returns a boolean if a field has been set.

### GetParserForType

`func (o *EnumValueDescriptorProto) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *EnumValueDescriptorProto) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *EnumValueDescriptorProto) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *EnumValueDescriptorProto) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetSerializedSize

`func (o *EnumValueDescriptorProto) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *EnumValueDescriptorProto) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *EnumValueDescriptorProto) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *EnumValueDescriptorProto) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetUnknownFields

`func (o *EnumValueDescriptorProto) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *EnumValueDescriptorProto) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *EnumValueDescriptorProto) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *EnumValueDescriptorProto) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


