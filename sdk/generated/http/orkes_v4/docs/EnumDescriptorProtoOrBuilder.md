# EnumDescriptorProtoOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to  |  | [optional] 
**DefaultInstanceForType** | Pointer to  | Simplified schema for Message (original had circular references) | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**Options** | Pointer to [**EnumOptions**](EnumOptions.md) |  | [optional] 
**OptionsOrBuilder** | Pointer to [**EnumOptionsOrBuilder**](EnumOptionsOrBuilder.md) |  | [optional] 
**ReservedNameCount** | Pointer to **int32** |  | [optional] 
**ReservedNameList** | Pointer to **[]string** |  | [optional] 
**ReservedRangeCount** | Pointer to **int32** |  | [optional] 
**ReservedRangeList** | Pointer to [**[]EnumReservedRange**](EnumReservedRange.md) |  | [optional] 
**ReservedRangeOrBuilderList** | Pointer to [**[]EnumReservedRangeOrBuilder**](EnumReservedRangeOrBuilder.md) |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**ValueCount** | Pointer to **int32** |  | [optional] 
**ValueList** | Pointer to [**[]EnumValueDescriptorProto**](EnumValueDescriptorProto.md) |  | [optional] 
**ValueOrBuilderList** | Pointer to [**[]EnumValueDescriptorProtoOrBuilder**](EnumValueDescriptorProtoOrBuilder.md) |  | [optional] 

## Methods

### NewEnumDescriptorProtoOrBuilder

`func NewEnumDescriptorProtoOrBuilder() *EnumDescriptorProtoOrBuilder`

NewEnumDescriptorProtoOrBuilder instantiates a new EnumDescriptorProtoOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumDescriptorProtoOrBuilderWithDefaults

`func NewEnumDescriptorProtoOrBuilderWithDefaults() *EnumDescriptorProtoOrBuilder`

NewEnumDescriptorProtoOrBuilderWithDefaults instantiates a new EnumDescriptorProtoOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *EnumDescriptorProtoOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *EnumDescriptorProtoOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *EnumDescriptorProtoOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *EnumDescriptorProtoOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### SetAllFieldsNil

`func (o *EnumDescriptorProtoOrBuilder) SetAllFieldsNil(b bool)`

 SetAllFieldsNil sets the value for AllFields to be an explicit nil

### UnsetAllFields
`func (o *EnumDescriptorProtoOrBuilder) UnsetAllFields()`

UnsetAllFields ensures that no value is present for AllFields, not even an explicit nil
### GetDefaultInstanceForType

`func (o *EnumDescriptorProtoOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *EnumDescriptorProtoOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *EnumDescriptorProtoOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *EnumDescriptorProtoOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### SetDefaultInstanceForTypeNil

`func (o *EnumDescriptorProtoOrBuilder) SetDefaultInstanceForTypeNil(b bool)`

 SetDefaultInstanceForTypeNil sets the value for DefaultInstanceForType to be an explicit nil

### UnsetDefaultInstanceForType
`func (o *EnumDescriptorProtoOrBuilder) UnsetDefaultInstanceForType()`

UnsetDefaultInstanceForType ensures that no value is present for DefaultInstanceForType, not even an explicit nil
### GetDescriptorForType

`func (o *EnumDescriptorProtoOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *EnumDescriptorProtoOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *EnumDescriptorProtoOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *EnumDescriptorProtoOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *EnumDescriptorProtoOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *EnumDescriptorProtoOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *EnumDescriptorProtoOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *EnumDescriptorProtoOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *EnumDescriptorProtoOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *EnumDescriptorProtoOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *EnumDescriptorProtoOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *EnumDescriptorProtoOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetName

`func (o *EnumDescriptorProtoOrBuilder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnumDescriptorProtoOrBuilder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnumDescriptorProtoOrBuilder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnumDescriptorProtoOrBuilder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameBytes

`func (o *EnumDescriptorProtoOrBuilder) GetNameBytes() ByteString`

GetNameBytes returns the NameBytes field if non-nil, zero value otherwise.

### GetNameBytesOk

`func (o *EnumDescriptorProtoOrBuilder) GetNameBytesOk() (*ByteString, bool)`

GetNameBytesOk returns a tuple with the NameBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameBytes

`func (o *EnumDescriptorProtoOrBuilder) SetNameBytes(v ByteString)`

SetNameBytes sets NameBytes field to given value.

### HasNameBytes

`func (o *EnumDescriptorProtoOrBuilder) HasNameBytes() bool`

HasNameBytes returns a boolean if a field has been set.

### GetOptions

`func (o *EnumDescriptorProtoOrBuilder) GetOptions() EnumOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *EnumDescriptorProtoOrBuilder) GetOptionsOk() (*EnumOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *EnumDescriptorProtoOrBuilder) SetOptions(v EnumOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *EnumDescriptorProtoOrBuilder) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOptionsOrBuilder

`func (o *EnumDescriptorProtoOrBuilder) GetOptionsOrBuilder() EnumOptionsOrBuilder`

GetOptionsOrBuilder returns the OptionsOrBuilder field if non-nil, zero value otherwise.

### GetOptionsOrBuilderOk

`func (o *EnumDescriptorProtoOrBuilder) GetOptionsOrBuilderOk() (*EnumOptionsOrBuilder, bool)`

GetOptionsOrBuilderOk returns a tuple with the OptionsOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsOrBuilder

`func (o *EnumDescriptorProtoOrBuilder) SetOptionsOrBuilder(v EnumOptionsOrBuilder)`

SetOptionsOrBuilder sets OptionsOrBuilder field to given value.

### HasOptionsOrBuilder

`func (o *EnumDescriptorProtoOrBuilder) HasOptionsOrBuilder() bool`

HasOptionsOrBuilder returns a boolean if a field has been set.

### GetReservedNameCount

`func (o *EnumDescriptorProtoOrBuilder) GetReservedNameCount() int32`

GetReservedNameCount returns the ReservedNameCount field if non-nil, zero value otherwise.

### GetReservedNameCountOk

`func (o *EnumDescriptorProtoOrBuilder) GetReservedNameCountOk() (*int32, bool)`

GetReservedNameCountOk returns a tuple with the ReservedNameCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedNameCount

`func (o *EnumDescriptorProtoOrBuilder) SetReservedNameCount(v int32)`

SetReservedNameCount sets ReservedNameCount field to given value.

### HasReservedNameCount

`func (o *EnumDescriptorProtoOrBuilder) HasReservedNameCount() bool`

HasReservedNameCount returns a boolean if a field has been set.

### GetReservedNameList

`func (o *EnumDescriptorProtoOrBuilder) GetReservedNameList() []string`

GetReservedNameList returns the ReservedNameList field if non-nil, zero value otherwise.

### GetReservedNameListOk

`func (o *EnumDescriptorProtoOrBuilder) GetReservedNameListOk() (*[]string, bool)`

GetReservedNameListOk returns a tuple with the ReservedNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedNameList

`func (o *EnumDescriptorProtoOrBuilder) SetReservedNameList(v []string)`

SetReservedNameList sets ReservedNameList field to given value.

### HasReservedNameList

`func (o *EnumDescriptorProtoOrBuilder) HasReservedNameList() bool`

HasReservedNameList returns a boolean if a field has been set.

### GetReservedRangeCount

`func (o *EnumDescriptorProtoOrBuilder) GetReservedRangeCount() int32`

GetReservedRangeCount returns the ReservedRangeCount field if non-nil, zero value otherwise.

### GetReservedRangeCountOk

`func (o *EnumDescriptorProtoOrBuilder) GetReservedRangeCountOk() (*int32, bool)`

GetReservedRangeCountOk returns a tuple with the ReservedRangeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedRangeCount

`func (o *EnumDescriptorProtoOrBuilder) SetReservedRangeCount(v int32)`

SetReservedRangeCount sets ReservedRangeCount field to given value.

### HasReservedRangeCount

`func (o *EnumDescriptorProtoOrBuilder) HasReservedRangeCount() bool`

HasReservedRangeCount returns a boolean if a field has been set.

### GetReservedRangeList

`func (o *EnumDescriptorProtoOrBuilder) GetReservedRangeList() []EnumReservedRange`

GetReservedRangeList returns the ReservedRangeList field if non-nil, zero value otherwise.

### GetReservedRangeListOk

`func (o *EnumDescriptorProtoOrBuilder) GetReservedRangeListOk() (*[]EnumReservedRange, bool)`

GetReservedRangeListOk returns a tuple with the ReservedRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedRangeList

`func (o *EnumDescriptorProtoOrBuilder) SetReservedRangeList(v []EnumReservedRange)`

SetReservedRangeList sets ReservedRangeList field to given value.

### HasReservedRangeList

`func (o *EnumDescriptorProtoOrBuilder) HasReservedRangeList() bool`

HasReservedRangeList returns a boolean if a field has been set.

### GetReservedRangeOrBuilderList

`func (o *EnumDescriptorProtoOrBuilder) GetReservedRangeOrBuilderList() []EnumReservedRangeOrBuilder`

GetReservedRangeOrBuilderList returns the ReservedRangeOrBuilderList field if non-nil, zero value otherwise.

### GetReservedRangeOrBuilderListOk

`func (o *EnumDescriptorProtoOrBuilder) GetReservedRangeOrBuilderListOk() (*[]EnumReservedRangeOrBuilder, bool)`

GetReservedRangeOrBuilderListOk returns a tuple with the ReservedRangeOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedRangeOrBuilderList

`func (o *EnumDescriptorProtoOrBuilder) SetReservedRangeOrBuilderList(v []EnumReservedRangeOrBuilder)`

SetReservedRangeOrBuilderList sets ReservedRangeOrBuilderList field to given value.

### HasReservedRangeOrBuilderList

`func (o *EnumDescriptorProtoOrBuilder) HasReservedRangeOrBuilderList() bool`

HasReservedRangeOrBuilderList returns a boolean if a field has been set.

### GetUnknownFields

`func (o *EnumDescriptorProtoOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *EnumDescriptorProtoOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *EnumDescriptorProtoOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *EnumDescriptorProtoOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetValueCount

`func (o *EnumDescriptorProtoOrBuilder) GetValueCount() int32`

GetValueCount returns the ValueCount field if non-nil, zero value otherwise.

### GetValueCountOk

`func (o *EnumDescriptorProtoOrBuilder) GetValueCountOk() (*int32, bool)`

GetValueCountOk returns a tuple with the ValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCount

`func (o *EnumDescriptorProtoOrBuilder) SetValueCount(v int32)`

SetValueCount sets ValueCount field to given value.

### HasValueCount

`func (o *EnumDescriptorProtoOrBuilder) HasValueCount() bool`

HasValueCount returns a boolean if a field has been set.

### GetValueList

`func (o *EnumDescriptorProtoOrBuilder) GetValueList() []EnumValueDescriptorProto`

GetValueList returns the ValueList field if non-nil, zero value otherwise.

### GetValueListOk

`func (o *EnumDescriptorProtoOrBuilder) GetValueListOk() (*[]EnumValueDescriptorProto, bool)`

GetValueListOk returns a tuple with the ValueList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueList

`func (o *EnumDescriptorProtoOrBuilder) SetValueList(v []EnumValueDescriptorProto)`

SetValueList sets ValueList field to given value.

### HasValueList

`func (o *EnumDescriptorProtoOrBuilder) HasValueList() bool`

HasValueList returns a boolean if a field has been set.

### GetValueOrBuilderList

`func (o *EnumDescriptorProtoOrBuilder) GetValueOrBuilderList() []EnumValueDescriptorProtoOrBuilder`

GetValueOrBuilderList returns the ValueOrBuilderList field if non-nil, zero value otherwise.

### GetValueOrBuilderListOk

`func (o *EnumDescriptorProtoOrBuilder) GetValueOrBuilderListOk() (*[]EnumValueDescriptorProtoOrBuilder, bool)`

GetValueOrBuilderListOk returns a tuple with the ValueOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueOrBuilderList

`func (o *EnumDescriptorProtoOrBuilder) SetValueOrBuilderList(v []EnumValueDescriptorProtoOrBuilder)`

SetValueOrBuilderList sets ValueOrBuilderList field to given value.

### HasValueOrBuilderList

`func (o *EnumDescriptorProtoOrBuilder) HasValueOrBuilderList() bool`

HasValueOrBuilderList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


