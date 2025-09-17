# EnumOptionsOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deprecated** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to [**FeatureSet**](FeatureSet.md) |  | [optional] 
**DeprecatedLegacyJsonFieldConflicts** | Pointer to **bool** |  | [optional] 
**UninterpretedOptionOrBuilderList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**AllowAlias** | Pointer to **bool** |  | [optional] 
**FeaturesOrBuilder** | Pointer to [**FeatureSetOrBuilder**](FeatureSetOrBuilder.md) |  | [optional] 
**UninterpretedOptionList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UninterpretedOptionCount** | Pointer to **int32** |  | [optional] 
**DefaultInstanceForType** | Pointer to **map[string]interface{}** | Simplified schema for Message (original had circular references) | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 

## Methods

### NewEnumOptionsOrBuilder

`func NewEnumOptionsOrBuilder() *EnumOptionsOrBuilder`

NewEnumOptionsOrBuilder instantiates a new EnumOptionsOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumOptionsOrBuilderWithDefaults

`func NewEnumOptionsOrBuilderWithDefaults() *EnumOptionsOrBuilder`

NewEnumOptionsOrBuilderWithDefaults instantiates a new EnumOptionsOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeprecated

`func (o *EnumOptionsOrBuilder) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *EnumOptionsOrBuilder) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *EnumOptionsOrBuilder) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *EnumOptionsOrBuilder) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetFeatures

`func (o *EnumOptionsOrBuilder) GetFeatures() FeatureSet`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *EnumOptionsOrBuilder) GetFeaturesOk() (*FeatureSet, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *EnumOptionsOrBuilder) SetFeatures(v FeatureSet)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *EnumOptionsOrBuilder) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetDeprecatedLegacyJsonFieldConflicts

`func (o *EnumOptionsOrBuilder) GetDeprecatedLegacyJsonFieldConflicts() bool`

GetDeprecatedLegacyJsonFieldConflicts returns the DeprecatedLegacyJsonFieldConflicts field if non-nil, zero value otherwise.

### GetDeprecatedLegacyJsonFieldConflictsOk

`func (o *EnumOptionsOrBuilder) GetDeprecatedLegacyJsonFieldConflictsOk() (*bool, bool)`

GetDeprecatedLegacyJsonFieldConflictsOk returns a tuple with the DeprecatedLegacyJsonFieldConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedLegacyJsonFieldConflicts

`func (o *EnumOptionsOrBuilder) SetDeprecatedLegacyJsonFieldConflicts(v bool)`

SetDeprecatedLegacyJsonFieldConflicts sets DeprecatedLegacyJsonFieldConflicts field to given value.

### HasDeprecatedLegacyJsonFieldConflicts

`func (o *EnumOptionsOrBuilder) HasDeprecatedLegacyJsonFieldConflicts() bool`

HasDeprecatedLegacyJsonFieldConflicts returns a boolean if a field has been set.

### GetUninterpretedOptionOrBuilderList

`func (o *EnumOptionsOrBuilder) GetUninterpretedOptionOrBuilderList() []map[string]interface{}`

GetUninterpretedOptionOrBuilderList returns the UninterpretedOptionOrBuilderList field if non-nil, zero value otherwise.

### GetUninterpretedOptionOrBuilderListOk

`func (o *EnumOptionsOrBuilder) GetUninterpretedOptionOrBuilderListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionOrBuilderListOk returns a tuple with the UninterpretedOptionOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionOrBuilderList

`func (o *EnumOptionsOrBuilder) SetUninterpretedOptionOrBuilderList(v []map[string]interface{})`

SetUninterpretedOptionOrBuilderList sets UninterpretedOptionOrBuilderList field to given value.

### HasUninterpretedOptionOrBuilderList

`func (o *EnumOptionsOrBuilder) HasUninterpretedOptionOrBuilderList() bool`

HasUninterpretedOptionOrBuilderList returns a boolean if a field has been set.

### GetAllowAlias

`func (o *EnumOptionsOrBuilder) GetAllowAlias() bool`

GetAllowAlias returns the AllowAlias field if non-nil, zero value otherwise.

### GetAllowAliasOk

`func (o *EnumOptionsOrBuilder) GetAllowAliasOk() (*bool, bool)`

GetAllowAliasOk returns a tuple with the AllowAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAlias

`func (o *EnumOptionsOrBuilder) SetAllowAlias(v bool)`

SetAllowAlias sets AllowAlias field to given value.

### HasAllowAlias

`func (o *EnumOptionsOrBuilder) HasAllowAlias() bool`

HasAllowAlias returns a boolean if a field has been set.

### GetFeaturesOrBuilder

`func (o *EnumOptionsOrBuilder) GetFeaturesOrBuilder() FeatureSetOrBuilder`

GetFeaturesOrBuilder returns the FeaturesOrBuilder field if non-nil, zero value otherwise.

### GetFeaturesOrBuilderOk

`func (o *EnumOptionsOrBuilder) GetFeaturesOrBuilderOk() (*FeatureSetOrBuilder, bool)`

GetFeaturesOrBuilderOk returns a tuple with the FeaturesOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesOrBuilder

`func (o *EnumOptionsOrBuilder) SetFeaturesOrBuilder(v FeatureSetOrBuilder)`

SetFeaturesOrBuilder sets FeaturesOrBuilder field to given value.

### HasFeaturesOrBuilder

`func (o *EnumOptionsOrBuilder) HasFeaturesOrBuilder() bool`

HasFeaturesOrBuilder returns a boolean if a field has been set.

### GetUninterpretedOptionList

`func (o *EnumOptionsOrBuilder) GetUninterpretedOptionList() []map[string]interface{}`

GetUninterpretedOptionList returns the UninterpretedOptionList field if non-nil, zero value otherwise.

### GetUninterpretedOptionListOk

`func (o *EnumOptionsOrBuilder) GetUninterpretedOptionListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionListOk returns a tuple with the UninterpretedOptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionList

`func (o *EnumOptionsOrBuilder) SetUninterpretedOptionList(v []map[string]interface{})`

SetUninterpretedOptionList sets UninterpretedOptionList field to given value.

### HasUninterpretedOptionList

`func (o *EnumOptionsOrBuilder) HasUninterpretedOptionList() bool`

HasUninterpretedOptionList returns a boolean if a field has been set.

### GetUninterpretedOptionCount

`func (o *EnumOptionsOrBuilder) GetUninterpretedOptionCount() int32`

GetUninterpretedOptionCount returns the UninterpretedOptionCount field if non-nil, zero value otherwise.

### GetUninterpretedOptionCountOk

`func (o *EnumOptionsOrBuilder) GetUninterpretedOptionCountOk() (*int32, bool)`

GetUninterpretedOptionCountOk returns a tuple with the UninterpretedOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionCount

`func (o *EnumOptionsOrBuilder) SetUninterpretedOptionCount(v int32)`

SetUninterpretedOptionCount sets UninterpretedOptionCount field to given value.

### HasUninterpretedOptionCount

`func (o *EnumOptionsOrBuilder) HasUninterpretedOptionCount() bool`

HasUninterpretedOptionCount returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *EnumOptionsOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *EnumOptionsOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *EnumOptionsOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *EnumOptionsOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *EnumOptionsOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *EnumOptionsOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *EnumOptionsOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *EnumOptionsOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *EnumOptionsOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *EnumOptionsOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *EnumOptionsOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *EnumOptionsOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetAllFields

`func (o *EnumOptionsOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *EnumOptionsOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *EnumOptionsOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *EnumOptionsOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetUnknownFields

`func (o *EnumOptionsOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *EnumOptionsOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *EnumOptionsOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *EnumOptionsOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetInitialized

`func (o *EnumOptionsOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *EnumOptionsOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *EnumOptionsOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *EnumOptionsOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


