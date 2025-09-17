# MessageOptionsOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deprecated** | Pointer to **bool** |  | [optional] 
**MessageSetWireFormat** | Pointer to **bool** |  | [optional] 
**MapEntry** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to [**FeatureSet**](FeatureSet.md) |  | [optional] 
**DeprecatedLegacyJsonFieldConflicts** | Pointer to **bool** |  | [optional] 
**UninterpretedOptionOrBuilderList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**NoStandardDescriptorAccessor** | Pointer to **bool** |  | [optional] 
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

### NewMessageOptionsOrBuilder

`func NewMessageOptionsOrBuilder() *MessageOptionsOrBuilder`

NewMessageOptionsOrBuilder instantiates a new MessageOptionsOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageOptionsOrBuilderWithDefaults

`func NewMessageOptionsOrBuilderWithDefaults() *MessageOptionsOrBuilder`

NewMessageOptionsOrBuilderWithDefaults instantiates a new MessageOptionsOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeprecated

`func (o *MessageOptionsOrBuilder) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *MessageOptionsOrBuilder) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *MessageOptionsOrBuilder) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *MessageOptionsOrBuilder) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetMessageSetWireFormat

`func (o *MessageOptionsOrBuilder) GetMessageSetWireFormat() bool`

GetMessageSetWireFormat returns the MessageSetWireFormat field if non-nil, zero value otherwise.

### GetMessageSetWireFormatOk

`func (o *MessageOptionsOrBuilder) GetMessageSetWireFormatOk() (*bool, bool)`

GetMessageSetWireFormatOk returns a tuple with the MessageSetWireFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSetWireFormat

`func (o *MessageOptionsOrBuilder) SetMessageSetWireFormat(v bool)`

SetMessageSetWireFormat sets MessageSetWireFormat field to given value.

### HasMessageSetWireFormat

`func (o *MessageOptionsOrBuilder) HasMessageSetWireFormat() bool`

HasMessageSetWireFormat returns a boolean if a field has been set.

### GetMapEntry

`func (o *MessageOptionsOrBuilder) GetMapEntry() bool`

GetMapEntry returns the MapEntry field if non-nil, zero value otherwise.

### GetMapEntryOk

`func (o *MessageOptionsOrBuilder) GetMapEntryOk() (*bool, bool)`

GetMapEntryOk returns a tuple with the MapEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapEntry

`func (o *MessageOptionsOrBuilder) SetMapEntry(v bool)`

SetMapEntry sets MapEntry field to given value.

### HasMapEntry

`func (o *MessageOptionsOrBuilder) HasMapEntry() bool`

HasMapEntry returns a boolean if a field has been set.

### GetFeatures

`func (o *MessageOptionsOrBuilder) GetFeatures() FeatureSet`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *MessageOptionsOrBuilder) GetFeaturesOk() (*FeatureSet, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *MessageOptionsOrBuilder) SetFeatures(v FeatureSet)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *MessageOptionsOrBuilder) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetDeprecatedLegacyJsonFieldConflicts

`func (o *MessageOptionsOrBuilder) GetDeprecatedLegacyJsonFieldConflicts() bool`

GetDeprecatedLegacyJsonFieldConflicts returns the DeprecatedLegacyJsonFieldConflicts field if non-nil, zero value otherwise.

### GetDeprecatedLegacyJsonFieldConflictsOk

`func (o *MessageOptionsOrBuilder) GetDeprecatedLegacyJsonFieldConflictsOk() (*bool, bool)`

GetDeprecatedLegacyJsonFieldConflictsOk returns a tuple with the DeprecatedLegacyJsonFieldConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedLegacyJsonFieldConflicts

`func (o *MessageOptionsOrBuilder) SetDeprecatedLegacyJsonFieldConflicts(v bool)`

SetDeprecatedLegacyJsonFieldConflicts sets DeprecatedLegacyJsonFieldConflicts field to given value.

### HasDeprecatedLegacyJsonFieldConflicts

`func (o *MessageOptionsOrBuilder) HasDeprecatedLegacyJsonFieldConflicts() bool`

HasDeprecatedLegacyJsonFieldConflicts returns a boolean if a field has been set.

### GetUninterpretedOptionOrBuilderList

`func (o *MessageOptionsOrBuilder) GetUninterpretedOptionOrBuilderList() []map[string]interface{}`

GetUninterpretedOptionOrBuilderList returns the UninterpretedOptionOrBuilderList field if non-nil, zero value otherwise.

### GetUninterpretedOptionOrBuilderListOk

`func (o *MessageOptionsOrBuilder) GetUninterpretedOptionOrBuilderListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionOrBuilderListOk returns a tuple with the UninterpretedOptionOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionOrBuilderList

`func (o *MessageOptionsOrBuilder) SetUninterpretedOptionOrBuilderList(v []map[string]interface{})`

SetUninterpretedOptionOrBuilderList sets UninterpretedOptionOrBuilderList field to given value.

### HasUninterpretedOptionOrBuilderList

`func (o *MessageOptionsOrBuilder) HasUninterpretedOptionOrBuilderList() bool`

HasUninterpretedOptionOrBuilderList returns a boolean if a field has been set.

### GetNoStandardDescriptorAccessor

`func (o *MessageOptionsOrBuilder) GetNoStandardDescriptorAccessor() bool`

GetNoStandardDescriptorAccessor returns the NoStandardDescriptorAccessor field if non-nil, zero value otherwise.

### GetNoStandardDescriptorAccessorOk

`func (o *MessageOptionsOrBuilder) GetNoStandardDescriptorAccessorOk() (*bool, bool)`

GetNoStandardDescriptorAccessorOk returns a tuple with the NoStandardDescriptorAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoStandardDescriptorAccessor

`func (o *MessageOptionsOrBuilder) SetNoStandardDescriptorAccessor(v bool)`

SetNoStandardDescriptorAccessor sets NoStandardDescriptorAccessor field to given value.

### HasNoStandardDescriptorAccessor

`func (o *MessageOptionsOrBuilder) HasNoStandardDescriptorAccessor() bool`

HasNoStandardDescriptorAccessor returns a boolean if a field has been set.

### GetFeaturesOrBuilder

`func (o *MessageOptionsOrBuilder) GetFeaturesOrBuilder() FeatureSetOrBuilder`

GetFeaturesOrBuilder returns the FeaturesOrBuilder field if non-nil, zero value otherwise.

### GetFeaturesOrBuilderOk

`func (o *MessageOptionsOrBuilder) GetFeaturesOrBuilderOk() (*FeatureSetOrBuilder, bool)`

GetFeaturesOrBuilderOk returns a tuple with the FeaturesOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesOrBuilder

`func (o *MessageOptionsOrBuilder) SetFeaturesOrBuilder(v FeatureSetOrBuilder)`

SetFeaturesOrBuilder sets FeaturesOrBuilder field to given value.

### HasFeaturesOrBuilder

`func (o *MessageOptionsOrBuilder) HasFeaturesOrBuilder() bool`

HasFeaturesOrBuilder returns a boolean if a field has been set.

### GetUninterpretedOptionList

`func (o *MessageOptionsOrBuilder) GetUninterpretedOptionList() []map[string]interface{}`

GetUninterpretedOptionList returns the UninterpretedOptionList field if non-nil, zero value otherwise.

### GetUninterpretedOptionListOk

`func (o *MessageOptionsOrBuilder) GetUninterpretedOptionListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionListOk returns a tuple with the UninterpretedOptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionList

`func (o *MessageOptionsOrBuilder) SetUninterpretedOptionList(v []map[string]interface{})`

SetUninterpretedOptionList sets UninterpretedOptionList field to given value.

### HasUninterpretedOptionList

`func (o *MessageOptionsOrBuilder) HasUninterpretedOptionList() bool`

HasUninterpretedOptionList returns a boolean if a field has been set.

### GetUninterpretedOptionCount

`func (o *MessageOptionsOrBuilder) GetUninterpretedOptionCount() int32`

GetUninterpretedOptionCount returns the UninterpretedOptionCount field if non-nil, zero value otherwise.

### GetUninterpretedOptionCountOk

`func (o *MessageOptionsOrBuilder) GetUninterpretedOptionCountOk() (*int32, bool)`

GetUninterpretedOptionCountOk returns a tuple with the UninterpretedOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionCount

`func (o *MessageOptionsOrBuilder) SetUninterpretedOptionCount(v int32)`

SetUninterpretedOptionCount sets UninterpretedOptionCount field to given value.

### HasUninterpretedOptionCount

`func (o *MessageOptionsOrBuilder) HasUninterpretedOptionCount() bool`

HasUninterpretedOptionCount returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *MessageOptionsOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *MessageOptionsOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *MessageOptionsOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *MessageOptionsOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *MessageOptionsOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *MessageOptionsOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *MessageOptionsOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *MessageOptionsOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *MessageOptionsOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *MessageOptionsOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *MessageOptionsOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *MessageOptionsOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetAllFields

`func (o *MessageOptionsOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *MessageOptionsOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *MessageOptionsOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *MessageOptionsOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetUnknownFields

`func (o *MessageOptionsOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *MessageOptionsOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *MessageOptionsOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *MessageOptionsOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetInitialized

`func (o *MessageOptionsOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *MessageOptionsOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *MessageOptionsOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *MessageOptionsOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


