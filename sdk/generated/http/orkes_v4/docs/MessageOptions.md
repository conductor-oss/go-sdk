# MessageOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to  |  | [optional] 
**AllFieldsRaw** | Pointer to  |  | [optional] 
**DefaultInstanceForType** | Pointer to [**MessageOptions**](MessageOptions.md) |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**DeprecatedLegacyJsonFieldConflicts** | Pointer to **bool** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**Features** | Pointer to [**FeatureSet**](FeatureSet.md) |  | [optional] 
**FeaturesOrBuilder** | Pointer to [**FeatureSetOrBuilder**](FeatureSetOrBuilder.md) |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**MapEntry** | Pointer to **bool** |  | [optional] 
**MemoizedSerializedSize** | Pointer to **int32** |  | [optional] 
**MessageSetWireFormat** | Pointer to **bool** |  | [optional] 
**NoStandardDescriptorAccessor** | Pointer to **bool** |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** |  | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**UninterpretedOptionCount** | Pointer to **int32** |  | [optional] 
**UninterpretedOptionList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UninterpretedOptionOrBuilderList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewMessageOptions

`func NewMessageOptions() *MessageOptions`

NewMessageOptions instantiates a new MessageOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageOptionsWithDefaults

`func NewMessageOptionsWithDefaults() *MessageOptions`

NewMessageOptionsWithDefaults instantiates a new MessageOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *MessageOptions) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *MessageOptions) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *MessageOptions) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *MessageOptions) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### SetAllFieldsNil

`func (o *MessageOptions) SetAllFieldsNil(b bool)`

 SetAllFieldsNil sets the value for AllFields to be an explicit nil

### UnsetAllFields
`func (o *MessageOptions) UnsetAllFields()`

UnsetAllFields ensures that no value is present for AllFields, not even an explicit nil
### GetAllFieldsRaw

`func (o *MessageOptions) GetAllFieldsRaw() map[string]map[string]interface{}`

GetAllFieldsRaw returns the AllFieldsRaw field if non-nil, zero value otherwise.

### GetAllFieldsRawOk

`func (o *MessageOptions) GetAllFieldsRawOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsRawOk returns a tuple with the AllFieldsRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFieldsRaw

`func (o *MessageOptions) SetAllFieldsRaw(v map[string]map[string]interface{})`

SetAllFieldsRaw sets AllFieldsRaw field to given value.

### HasAllFieldsRaw

`func (o *MessageOptions) HasAllFieldsRaw() bool`

HasAllFieldsRaw returns a boolean if a field has been set.

### SetAllFieldsRawNil

`func (o *MessageOptions) SetAllFieldsRawNil(b bool)`

 SetAllFieldsRawNil sets the value for AllFieldsRaw to be an explicit nil

### UnsetAllFieldsRaw
`func (o *MessageOptions) UnsetAllFieldsRaw()`

UnsetAllFieldsRaw ensures that no value is present for AllFieldsRaw, not even an explicit nil
### GetDefaultInstanceForType

`func (o *MessageOptions) GetDefaultInstanceForType() MessageOptions`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *MessageOptions) GetDefaultInstanceForTypeOk() (*MessageOptions, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *MessageOptions) SetDefaultInstanceForType(v MessageOptions)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *MessageOptions) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDeprecated

`func (o *MessageOptions) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *MessageOptions) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *MessageOptions) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *MessageOptions) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDeprecatedLegacyJsonFieldConflicts

`func (o *MessageOptions) GetDeprecatedLegacyJsonFieldConflicts() bool`

GetDeprecatedLegacyJsonFieldConflicts returns the DeprecatedLegacyJsonFieldConflicts field if non-nil, zero value otherwise.

### GetDeprecatedLegacyJsonFieldConflictsOk

`func (o *MessageOptions) GetDeprecatedLegacyJsonFieldConflictsOk() (*bool, bool)`

GetDeprecatedLegacyJsonFieldConflictsOk returns a tuple with the DeprecatedLegacyJsonFieldConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedLegacyJsonFieldConflicts

`func (o *MessageOptions) SetDeprecatedLegacyJsonFieldConflicts(v bool)`

SetDeprecatedLegacyJsonFieldConflicts sets DeprecatedLegacyJsonFieldConflicts field to given value.

### HasDeprecatedLegacyJsonFieldConflicts

`func (o *MessageOptions) HasDeprecatedLegacyJsonFieldConflicts() bool`

HasDeprecatedLegacyJsonFieldConflicts returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *MessageOptions) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *MessageOptions) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *MessageOptions) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *MessageOptions) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetFeatures

`func (o *MessageOptions) GetFeatures() FeatureSet`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *MessageOptions) GetFeaturesOk() (*FeatureSet, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *MessageOptions) SetFeatures(v FeatureSet)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *MessageOptions) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesOrBuilder

`func (o *MessageOptions) GetFeaturesOrBuilder() FeatureSetOrBuilder`

GetFeaturesOrBuilder returns the FeaturesOrBuilder field if non-nil, zero value otherwise.

### GetFeaturesOrBuilderOk

`func (o *MessageOptions) GetFeaturesOrBuilderOk() (*FeatureSetOrBuilder, bool)`

GetFeaturesOrBuilderOk returns a tuple with the FeaturesOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesOrBuilder

`func (o *MessageOptions) SetFeaturesOrBuilder(v FeatureSetOrBuilder)`

SetFeaturesOrBuilder sets FeaturesOrBuilder field to given value.

### HasFeaturesOrBuilder

`func (o *MessageOptions) HasFeaturesOrBuilder() bool`

HasFeaturesOrBuilder returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *MessageOptions) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *MessageOptions) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *MessageOptions) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *MessageOptions) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *MessageOptions) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *MessageOptions) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *MessageOptions) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *MessageOptions) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetMapEntry

`func (o *MessageOptions) GetMapEntry() bool`

GetMapEntry returns the MapEntry field if non-nil, zero value otherwise.

### GetMapEntryOk

`func (o *MessageOptions) GetMapEntryOk() (*bool, bool)`

GetMapEntryOk returns a tuple with the MapEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapEntry

`func (o *MessageOptions) SetMapEntry(v bool)`

SetMapEntry sets MapEntry field to given value.

### HasMapEntry

`func (o *MessageOptions) HasMapEntry() bool`

HasMapEntry returns a boolean if a field has been set.

### GetMemoizedSerializedSize

`func (o *MessageOptions) GetMemoizedSerializedSize() int32`

GetMemoizedSerializedSize returns the MemoizedSerializedSize field if non-nil, zero value otherwise.

### GetMemoizedSerializedSizeOk

`func (o *MessageOptions) GetMemoizedSerializedSizeOk() (*int32, bool)`

GetMemoizedSerializedSizeOk returns a tuple with the MemoizedSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizedSerializedSize

`func (o *MessageOptions) SetMemoizedSerializedSize(v int32)`

SetMemoizedSerializedSize sets MemoizedSerializedSize field to given value.

### HasMemoizedSerializedSize

`func (o *MessageOptions) HasMemoizedSerializedSize() bool`

HasMemoizedSerializedSize returns a boolean if a field has been set.

### GetMessageSetWireFormat

`func (o *MessageOptions) GetMessageSetWireFormat() bool`

GetMessageSetWireFormat returns the MessageSetWireFormat field if non-nil, zero value otherwise.

### GetMessageSetWireFormatOk

`func (o *MessageOptions) GetMessageSetWireFormatOk() (*bool, bool)`

GetMessageSetWireFormatOk returns a tuple with the MessageSetWireFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSetWireFormat

`func (o *MessageOptions) SetMessageSetWireFormat(v bool)`

SetMessageSetWireFormat sets MessageSetWireFormat field to given value.

### HasMessageSetWireFormat

`func (o *MessageOptions) HasMessageSetWireFormat() bool`

HasMessageSetWireFormat returns a boolean if a field has been set.

### GetNoStandardDescriptorAccessor

`func (o *MessageOptions) GetNoStandardDescriptorAccessor() bool`

GetNoStandardDescriptorAccessor returns the NoStandardDescriptorAccessor field if non-nil, zero value otherwise.

### GetNoStandardDescriptorAccessorOk

`func (o *MessageOptions) GetNoStandardDescriptorAccessorOk() (*bool, bool)`

GetNoStandardDescriptorAccessorOk returns a tuple with the NoStandardDescriptorAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoStandardDescriptorAccessor

`func (o *MessageOptions) SetNoStandardDescriptorAccessor(v bool)`

SetNoStandardDescriptorAccessor sets NoStandardDescriptorAccessor field to given value.

### HasNoStandardDescriptorAccessor

`func (o *MessageOptions) HasNoStandardDescriptorAccessor() bool`

HasNoStandardDescriptorAccessor returns a boolean if a field has been set.

### GetParserForType

`func (o *MessageOptions) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *MessageOptions) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *MessageOptions) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *MessageOptions) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetSerializedSize

`func (o *MessageOptions) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *MessageOptions) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *MessageOptions) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *MessageOptions) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetUninterpretedOptionCount

`func (o *MessageOptions) GetUninterpretedOptionCount() int32`

GetUninterpretedOptionCount returns the UninterpretedOptionCount field if non-nil, zero value otherwise.

### GetUninterpretedOptionCountOk

`func (o *MessageOptions) GetUninterpretedOptionCountOk() (*int32, bool)`

GetUninterpretedOptionCountOk returns a tuple with the UninterpretedOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionCount

`func (o *MessageOptions) SetUninterpretedOptionCount(v int32)`

SetUninterpretedOptionCount sets UninterpretedOptionCount field to given value.

### HasUninterpretedOptionCount

`func (o *MessageOptions) HasUninterpretedOptionCount() bool`

HasUninterpretedOptionCount returns a boolean if a field has been set.

### GetUninterpretedOptionList

`func (o *MessageOptions) GetUninterpretedOptionList() []map[string]interface{}`

GetUninterpretedOptionList returns the UninterpretedOptionList field if non-nil, zero value otherwise.

### GetUninterpretedOptionListOk

`func (o *MessageOptions) GetUninterpretedOptionListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionListOk returns a tuple with the UninterpretedOptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionList

`func (o *MessageOptions) SetUninterpretedOptionList(v []map[string]interface{})`

SetUninterpretedOptionList sets UninterpretedOptionList field to given value.

### HasUninterpretedOptionList

`func (o *MessageOptions) HasUninterpretedOptionList() bool`

HasUninterpretedOptionList returns a boolean if a field has been set.

### GetUninterpretedOptionOrBuilderList

`func (o *MessageOptions) GetUninterpretedOptionOrBuilderList() []map[string]interface{}`

GetUninterpretedOptionOrBuilderList returns the UninterpretedOptionOrBuilderList field if non-nil, zero value otherwise.

### GetUninterpretedOptionOrBuilderListOk

`func (o *MessageOptions) GetUninterpretedOptionOrBuilderListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionOrBuilderListOk returns a tuple with the UninterpretedOptionOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionOrBuilderList

`func (o *MessageOptions) SetUninterpretedOptionOrBuilderList(v []map[string]interface{})`

SetUninterpretedOptionOrBuilderList sets UninterpretedOptionOrBuilderList field to given value.

### HasUninterpretedOptionOrBuilderList

`func (o *MessageOptions) HasUninterpretedOptionOrBuilderList() bool`

HasUninterpretedOptionOrBuilderList returns a boolean if a field has been set.

### GetUnknownFields

`func (o *MessageOptions) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *MessageOptions) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *MessageOptions) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *MessageOptions) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


