# EnumOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**AllFieldsRaw** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**AllowAlias** | Pointer to **bool** |  | [optional] 
**DefaultInstanceForType** | Pointer to [**EnumOptions**](EnumOptions.md) |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**DeprecatedLegacyJsonFieldConflicts** | Pointer to **bool** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**Features** | Pointer to [**FeatureSet**](FeatureSet.md) |  | [optional] 
**FeaturesOrBuilder** | Pointer to [**FeatureSetOrBuilder**](FeatureSetOrBuilder.md) |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**MemoizedSerializedSize** | Pointer to **int32** |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** |  | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**UninterpretedOptionCount** | Pointer to **int32** |  | [optional] 
**UninterpretedOptionList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UninterpretedOptionOrBuilderList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewEnumOptions

`func NewEnumOptions() *EnumOptions`

NewEnumOptions instantiates a new EnumOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumOptionsWithDefaults

`func NewEnumOptionsWithDefaults() *EnumOptions`

NewEnumOptionsWithDefaults instantiates a new EnumOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *EnumOptions) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *EnumOptions) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *EnumOptions) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *EnumOptions) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetAllFieldsRaw

`func (o *EnumOptions) GetAllFieldsRaw() map[string]map[string]interface{}`

GetAllFieldsRaw returns the AllFieldsRaw field if non-nil, zero value otherwise.

### GetAllFieldsRawOk

`func (o *EnumOptions) GetAllFieldsRawOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsRawOk returns a tuple with the AllFieldsRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFieldsRaw

`func (o *EnumOptions) SetAllFieldsRaw(v map[string]map[string]interface{})`

SetAllFieldsRaw sets AllFieldsRaw field to given value.

### HasAllFieldsRaw

`func (o *EnumOptions) HasAllFieldsRaw() bool`

HasAllFieldsRaw returns a boolean if a field has been set.

### GetAllowAlias

`func (o *EnumOptions) GetAllowAlias() bool`

GetAllowAlias returns the AllowAlias field if non-nil, zero value otherwise.

### GetAllowAliasOk

`func (o *EnumOptions) GetAllowAliasOk() (*bool, bool)`

GetAllowAliasOk returns a tuple with the AllowAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAlias

`func (o *EnumOptions) SetAllowAlias(v bool)`

SetAllowAlias sets AllowAlias field to given value.

### HasAllowAlias

`func (o *EnumOptions) HasAllowAlias() bool`

HasAllowAlias returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *EnumOptions) GetDefaultInstanceForType() EnumOptions`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *EnumOptions) GetDefaultInstanceForTypeOk() (*EnumOptions, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *EnumOptions) SetDefaultInstanceForType(v EnumOptions)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *EnumOptions) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDeprecated

`func (o *EnumOptions) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *EnumOptions) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *EnumOptions) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *EnumOptions) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDeprecatedLegacyJsonFieldConflicts

`func (o *EnumOptions) GetDeprecatedLegacyJsonFieldConflicts() bool`

GetDeprecatedLegacyJsonFieldConflicts returns the DeprecatedLegacyJsonFieldConflicts field if non-nil, zero value otherwise.

### GetDeprecatedLegacyJsonFieldConflictsOk

`func (o *EnumOptions) GetDeprecatedLegacyJsonFieldConflictsOk() (*bool, bool)`

GetDeprecatedLegacyJsonFieldConflictsOk returns a tuple with the DeprecatedLegacyJsonFieldConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedLegacyJsonFieldConflicts

`func (o *EnumOptions) SetDeprecatedLegacyJsonFieldConflicts(v bool)`

SetDeprecatedLegacyJsonFieldConflicts sets DeprecatedLegacyJsonFieldConflicts field to given value.

### HasDeprecatedLegacyJsonFieldConflicts

`func (o *EnumOptions) HasDeprecatedLegacyJsonFieldConflicts() bool`

HasDeprecatedLegacyJsonFieldConflicts returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *EnumOptions) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *EnumOptions) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *EnumOptions) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *EnumOptions) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetFeatures

`func (o *EnumOptions) GetFeatures() FeatureSet`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *EnumOptions) GetFeaturesOk() (*FeatureSet, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *EnumOptions) SetFeatures(v FeatureSet)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *EnumOptions) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesOrBuilder

`func (o *EnumOptions) GetFeaturesOrBuilder() FeatureSetOrBuilder`

GetFeaturesOrBuilder returns the FeaturesOrBuilder field if non-nil, zero value otherwise.

### GetFeaturesOrBuilderOk

`func (o *EnumOptions) GetFeaturesOrBuilderOk() (*FeatureSetOrBuilder, bool)`

GetFeaturesOrBuilderOk returns a tuple with the FeaturesOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesOrBuilder

`func (o *EnumOptions) SetFeaturesOrBuilder(v FeatureSetOrBuilder)`

SetFeaturesOrBuilder sets FeaturesOrBuilder field to given value.

### HasFeaturesOrBuilder

`func (o *EnumOptions) HasFeaturesOrBuilder() bool`

HasFeaturesOrBuilder returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *EnumOptions) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *EnumOptions) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *EnumOptions) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *EnumOptions) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *EnumOptions) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *EnumOptions) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *EnumOptions) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *EnumOptions) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetMemoizedSerializedSize

`func (o *EnumOptions) GetMemoizedSerializedSize() int32`

GetMemoizedSerializedSize returns the MemoizedSerializedSize field if non-nil, zero value otherwise.

### GetMemoizedSerializedSizeOk

`func (o *EnumOptions) GetMemoizedSerializedSizeOk() (*int32, bool)`

GetMemoizedSerializedSizeOk returns a tuple with the MemoizedSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizedSerializedSize

`func (o *EnumOptions) SetMemoizedSerializedSize(v int32)`

SetMemoizedSerializedSize sets MemoizedSerializedSize field to given value.

### HasMemoizedSerializedSize

`func (o *EnumOptions) HasMemoizedSerializedSize() bool`

HasMemoizedSerializedSize returns a boolean if a field has been set.

### GetParserForType

`func (o *EnumOptions) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *EnumOptions) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *EnumOptions) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *EnumOptions) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetSerializedSize

`func (o *EnumOptions) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *EnumOptions) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *EnumOptions) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *EnumOptions) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetUninterpretedOptionCount

`func (o *EnumOptions) GetUninterpretedOptionCount() int32`

GetUninterpretedOptionCount returns the UninterpretedOptionCount field if non-nil, zero value otherwise.

### GetUninterpretedOptionCountOk

`func (o *EnumOptions) GetUninterpretedOptionCountOk() (*int32, bool)`

GetUninterpretedOptionCountOk returns a tuple with the UninterpretedOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionCount

`func (o *EnumOptions) SetUninterpretedOptionCount(v int32)`

SetUninterpretedOptionCount sets UninterpretedOptionCount field to given value.

### HasUninterpretedOptionCount

`func (o *EnumOptions) HasUninterpretedOptionCount() bool`

HasUninterpretedOptionCount returns a boolean if a field has been set.

### GetUninterpretedOptionList

`func (o *EnumOptions) GetUninterpretedOptionList() []map[string]interface{}`

GetUninterpretedOptionList returns the UninterpretedOptionList field if non-nil, zero value otherwise.

### GetUninterpretedOptionListOk

`func (o *EnumOptions) GetUninterpretedOptionListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionListOk returns a tuple with the UninterpretedOptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionList

`func (o *EnumOptions) SetUninterpretedOptionList(v []map[string]interface{})`

SetUninterpretedOptionList sets UninterpretedOptionList field to given value.

### HasUninterpretedOptionList

`func (o *EnumOptions) HasUninterpretedOptionList() bool`

HasUninterpretedOptionList returns a boolean if a field has been set.

### GetUninterpretedOptionOrBuilderList

`func (o *EnumOptions) GetUninterpretedOptionOrBuilderList() []map[string]interface{}`

GetUninterpretedOptionOrBuilderList returns the UninterpretedOptionOrBuilderList field if non-nil, zero value otherwise.

### GetUninterpretedOptionOrBuilderListOk

`func (o *EnumOptions) GetUninterpretedOptionOrBuilderListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionOrBuilderListOk returns a tuple with the UninterpretedOptionOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionOrBuilderList

`func (o *EnumOptions) SetUninterpretedOptionOrBuilderList(v []map[string]interface{})`

SetUninterpretedOptionOrBuilderList sets UninterpretedOptionOrBuilderList field to given value.

### HasUninterpretedOptionOrBuilderList

`func (o *EnumOptions) HasUninterpretedOptionOrBuilderList() bool`

HasUninterpretedOptionOrBuilderList returns a boolean if a field has been set.

### GetUnknownFields

`func (o *EnumOptions) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *EnumOptions) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *EnumOptions) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *EnumOptions) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


