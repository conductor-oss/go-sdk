# FieldOptionsOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to  |  | [optional] 
**Ctype** | Pointer to **string** |  | [optional] 
**DebugRedact** | Pointer to **bool** |  | [optional] 
**DefaultInstanceForType** | Pointer to  | Simplified schema for Message (original had circular references) | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**EditionDefaultsCount** | Pointer to **int32** |  | [optional] 
**EditionDefaultsList** | Pointer to [**[]EditionDefault**](EditionDefault.md) |  | [optional] 
**EditionDefaultsOrBuilderList** | Pointer to [**[]EditionDefaultOrBuilder**](EditionDefaultOrBuilder.md) |  | [optional] 
**Features** | Pointer to [**FeatureSet**](FeatureSet.md) |  | [optional] 
**FeaturesOrBuilder** | Pointer to [**FeatureSetOrBuilder**](FeatureSetOrBuilder.md) |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**Jstype** | Pointer to **string** |  | [optional] 
**Lazy** | Pointer to **bool** |  | [optional] 
**Packed** | Pointer to **bool** |  | [optional] 
**Retention** | Pointer to **string** |  | [optional] 
**TargetsCount** | Pointer to **int32** |  | [optional] 
**TargetsList** | Pointer to **[]string** |  | [optional] 
**UninterpretedOptionCount** | Pointer to **int32** |  | [optional] 
**UninterpretedOptionList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UninterpretedOptionOrBuilderList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**UnverifiedLazy** | Pointer to **bool** |  | [optional] 
**Weak** | Pointer to **bool** |  | [optional] 

## Methods

### NewFieldOptionsOrBuilder

`func NewFieldOptionsOrBuilder() *FieldOptionsOrBuilder`

NewFieldOptionsOrBuilder instantiates a new FieldOptionsOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldOptionsOrBuilderWithDefaults

`func NewFieldOptionsOrBuilderWithDefaults() *FieldOptionsOrBuilder`

NewFieldOptionsOrBuilderWithDefaults instantiates a new FieldOptionsOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *FieldOptionsOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *FieldOptionsOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *FieldOptionsOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *FieldOptionsOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### SetAllFieldsNil

`func (o *FieldOptionsOrBuilder) SetAllFieldsNil(b bool)`

 SetAllFieldsNil sets the value for AllFields to be an explicit nil

### UnsetAllFields
`func (o *FieldOptionsOrBuilder) UnsetAllFields()`

UnsetAllFields ensures that no value is present for AllFields, not even an explicit nil
### GetCtype

`func (o *FieldOptionsOrBuilder) GetCtype() string`

GetCtype returns the Ctype field if non-nil, zero value otherwise.

### GetCtypeOk

`func (o *FieldOptionsOrBuilder) GetCtypeOk() (*string, bool)`

GetCtypeOk returns a tuple with the Ctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtype

`func (o *FieldOptionsOrBuilder) SetCtype(v string)`

SetCtype sets Ctype field to given value.

### HasCtype

`func (o *FieldOptionsOrBuilder) HasCtype() bool`

HasCtype returns a boolean if a field has been set.

### GetDebugRedact

`func (o *FieldOptionsOrBuilder) GetDebugRedact() bool`

GetDebugRedact returns the DebugRedact field if non-nil, zero value otherwise.

### GetDebugRedactOk

`func (o *FieldOptionsOrBuilder) GetDebugRedactOk() (*bool, bool)`

GetDebugRedactOk returns a tuple with the DebugRedact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugRedact

`func (o *FieldOptionsOrBuilder) SetDebugRedact(v bool)`

SetDebugRedact sets DebugRedact field to given value.

### HasDebugRedact

`func (o *FieldOptionsOrBuilder) HasDebugRedact() bool`

HasDebugRedact returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *FieldOptionsOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *FieldOptionsOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *FieldOptionsOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *FieldOptionsOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### SetDefaultInstanceForTypeNil

`func (o *FieldOptionsOrBuilder) SetDefaultInstanceForTypeNil(b bool)`

 SetDefaultInstanceForTypeNil sets the value for DefaultInstanceForType to be an explicit nil

### UnsetDefaultInstanceForType
`func (o *FieldOptionsOrBuilder) UnsetDefaultInstanceForType()`

UnsetDefaultInstanceForType ensures that no value is present for DefaultInstanceForType, not even an explicit nil
### GetDeprecated

`func (o *FieldOptionsOrBuilder) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *FieldOptionsOrBuilder) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *FieldOptionsOrBuilder) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *FieldOptionsOrBuilder) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *FieldOptionsOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *FieldOptionsOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *FieldOptionsOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *FieldOptionsOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetEditionDefaultsCount

`func (o *FieldOptionsOrBuilder) GetEditionDefaultsCount() int32`

GetEditionDefaultsCount returns the EditionDefaultsCount field if non-nil, zero value otherwise.

### GetEditionDefaultsCountOk

`func (o *FieldOptionsOrBuilder) GetEditionDefaultsCountOk() (*int32, bool)`

GetEditionDefaultsCountOk returns a tuple with the EditionDefaultsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionDefaultsCount

`func (o *FieldOptionsOrBuilder) SetEditionDefaultsCount(v int32)`

SetEditionDefaultsCount sets EditionDefaultsCount field to given value.

### HasEditionDefaultsCount

`func (o *FieldOptionsOrBuilder) HasEditionDefaultsCount() bool`

HasEditionDefaultsCount returns a boolean if a field has been set.

### GetEditionDefaultsList

`func (o *FieldOptionsOrBuilder) GetEditionDefaultsList() []EditionDefault`

GetEditionDefaultsList returns the EditionDefaultsList field if non-nil, zero value otherwise.

### GetEditionDefaultsListOk

`func (o *FieldOptionsOrBuilder) GetEditionDefaultsListOk() (*[]EditionDefault, bool)`

GetEditionDefaultsListOk returns a tuple with the EditionDefaultsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionDefaultsList

`func (o *FieldOptionsOrBuilder) SetEditionDefaultsList(v []EditionDefault)`

SetEditionDefaultsList sets EditionDefaultsList field to given value.

### HasEditionDefaultsList

`func (o *FieldOptionsOrBuilder) HasEditionDefaultsList() bool`

HasEditionDefaultsList returns a boolean if a field has been set.

### GetEditionDefaultsOrBuilderList

`func (o *FieldOptionsOrBuilder) GetEditionDefaultsOrBuilderList() []EditionDefaultOrBuilder`

GetEditionDefaultsOrBuilderList returns the EditionDefaultsOrBuilderList field if non-nil, zero value otherwise.

### GetEditionDefaultsOrBuilderListOk

`func (o *FieldOptionsOrBuilder) GetEditionDefaultsOrBuilderListOk() (*[]EditionDefaultOrBuilder, bool)`

GetEditionDefaultsOrBuilderListOk returns a tuple with the EditionDefaultsOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionDefaultsOrBuilderList

`func (o *FieldOptionsOrBuilder) SetEditionDefaultsOrBuilderList(v []EditionDefaultOrBuilder)`

SetEditionDefaultsOrBuilderList sets EditionDefaultsOrBuilderList field to given value.

### HasEditionDefaultsOrBuilderList

`func (o *FieldOptionsOrBuilder) HasEditionDefaultsOrBuilderList() bool`

HasEditionDefaultsOrBuilderList returns a boolean if a field has been set.

### GetFeatures

`func (o *FieldOptionsOrBuilder) GetFeatures() FeatureSet`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FieldOptionsOrBuilder) GetFeaturesOk() (*FeatureSet, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FieldOptionsOrBuilder) SetFeatures(v FeatureSet)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FieldOptionsOrBuilder) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesOrBuilder

`func (o *FieldOptionsOrBuilder) GetFeaturesOrBuilder() FeatureSetOrBuilder`

GetFeaturesOrBuilder returns the FeaturesOrBuilder field if non-nil, zero value otherwise.

### GetFeaturesOrBuilderOk

`func (o *FieldOptionsOrBuilder) GetFeaturesOrBuilderOk() (*FeatureSetOrBuilder, bool)`

GetFeaturesOrBuilderOk returns a tuple with the FeaturesOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesOrBuilder

`func (o *FieldOptionsOrBuilder) SetFeaturesOrBuilder(v FeatureSetOrBuilder)`

SetFeaturesOrBuilder sets FeaturesOrBuilder field to given value.

### HasFeaturesOrBuilder

`func (o *FieldOptionsOrBuilder) HasFeaturesOrBuilder() bool`

HasFeaturesOrBuilder returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *FieldOptionsOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *FieldOptionsOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *FieldOptionsOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *FieldOptionsOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *FieldOptionsOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *FieldOptionsOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *FieldOptionsOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *FieldOptionsOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetJstype

`func (o *FieldOptionsOrBuilder) GetJstype() string`

GetJstype returns the Jstype field if non-nil, zero value otherwise.

### GetJstypeOk

`func (o *FieldOptionsOrBuilder) GetJstypeOk() (*string, bool)`

GetJstypeOk returns a tuple with the Jstype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJstype

`func (o *FieldOptionsOrBuilder) SetJstype(v string)`

SetJstype sets Jstype field to given value.

### HasJstype

`func (o *FieldOptionsOrBuilder) HasJstype() bool`

HasJstype returns a boolean if a field has been set.

### GetLazy

`func (o *FieldOptionsOrBuilder) GetLazy() bool`

GetLazy returns the Lazy field if non-nil, zero value otherwise.

### GetLazyOk

`func (o *FieldOptionsOrBuilder) GetLazyOk() (*bool, bool)`

GetLazyOk returns a tuple with the Lazy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLazy

`func (o *FieldOptionsOrBuilder) SetLazy(v bool)`

SetLazy sets Lazy field to given value.

### HasLazy

`func (o *FieldOptionsOrBuilder) HasLazy() bool`

HasLazy returns a boolean if a field has been set.

### GetPacked

`func (o *FieldOptionsOrBuilder) GetPacked() bool`

GetPacked returns the Packed field if non-nil, zero value otherwise.

### GetPackedOk

`func (o *FieldOptionsOrBuilder) GetPackedOk() (*bool, bool)`

GetPackedOk returns a tuple with the Packed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacked

`func (o *FieldOptionsOrBuilder) SetPacked(v bool)`

SetPacked sets Packed field to given value.

### HasPacked

`func (o *FieldOptionsOrBuilder) HasPacked() bool`

HasPacked returns a boolean if a field has been set.

### GetRetention

`func (o *FieldOptionsOrBuilder) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *FieldOptionsOrBuilder) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *FieldOptionsOrBuilder) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *FieldOptionsOrBuilder) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetTargetsCount

`func (o *FieldOptionsOrBuilder) GetTargetsCount() int32`

GetTargetsCount returns the TargetsCount field if non-nil, zero value otherwise.

### GetTargetsCountOk

`func (o *FieldOptionsOrBuilder) GetTargetsCountOk() (*int32, bool)`

GetTargetsCountOk returns a tuple with the TargetsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsCount

`func (o *FieldOptionsOrBuilder) SetTargetsCount(v int32)`

SetTargetsCount sets TargetsCount field to given value.

### HasTargetsCount

`func (o *FieldOptionsOrBuilder) HasTargetsCount() bool`

HasTargetsCount returns a boolean if a field has been set.

### GetTargetsList

`func (o *FieldOptionsOrBuilder) GetTargetsList() []string`

GetTargetsList returns the TargetsList field if non-nil, zero value otherwise.

### GetTargetsListOk

`func (o *FieldOptionsOrBuilder) GetTargetsListOk() (*[]string, bool)`

GetTargetsListOk returns a tuple with the TargetsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsList

`func (o *FieldOptionsOrBuilder) SetTargetsList(v []string)`

SetTargetsList sets TargetsList field to given value.

### HasTargetsList

`func (o *FieldOptionsOrBuilder) HasTargetsList() bool`

HasTargetsList returns a boolean if a field has been set.

### GetUninterpretedOptionCount

`func (o *FieldOptionsOrBuilder) GetUninterpretedOptionCount() int32`

GetUninterpretedOptionCount returns the UninterpretedOptionCount field if non-nil, zero value otherwise.

### GetUninterpretedOptionCountOk

`func (o *FieldOptionsOrBuilder) GetUninterpretedOptionCountOk() (*int32, bool)`

GetUninterpretedOptionCountOk returns a tuple with the UninterpretedOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionCount

`func (o *FieldOptionsOrBuilder) SetUninterpretedOptionCount(v int32)`

SetUninterpretedOptionCount sets UninterpretedOptionCount field to given value.

### HasUninterpretedOptionCount

`func (o *FieldOptionsOrBuilder) HasUninterpretedOptionCount() bool`

HasUninterpretedOptionCount returns a boolean if a field has been set.

### GetUninterpretedOptionList

`func (o *FieldOptionsOrBuilder) GetUninterpretedOptionList() []map[string]interface{}`

GetUninterpretedOptionList returns the UninterpretedOptionList field if non-nil, zero value otherwise.

### GetUninterpretedOptionListOk

`func (o *FieldOptionsOrBuilder) GetUninterpretedOptionListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionListOk returns a tuple with the UninterpretedOptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionList

`func (o *FieldOptionsOrBuilder) SetUninterpretedOptionList(v []map[string]interface{})`

SetUninterpretedOptionList sets UninterpretedOptionList field to given value.

### HasUninterpretedOptionList

`func (o *FieldOptionsOrBuilder) HasUninterpretedOptionList() bool`

HasUninterpretedOptionList returns a boolean if a field has been set.

### GetUninterpretedOptionOrBuilderList

`func (o *FieldOptionsOrBuilder) GetUninterpretedOptionOrBuilderList() []map[string]interface{}`

GetUninterpretedOptionOrBuilderList returns the UninterpretedOptionOrBuilderList field if non-nil, zero value otherwise.

### GetUninterpretedOptionOrBuilderListOk

`func (o *FieldOptionsOrBuilder) GetUninterpretedOptionOrBuilderListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionOrBuilderListOk returns a tuple with the UninterpretedOptionOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionOrBuilderList

`func (o *FieldOptionsOrBuilder) SetUninterpretedOptionOrBuilderList(v []map[string]interface{})`

SetUninterpretedOptionOrBuilderList sets UninterpretedOptionOrBuilderList field to given value.

### HasUninterpretedOptionOrBuilderList

`func (o *FieldOptionsOrBuilder) HasUninterpretedOptionOrBuilderList() bool`

HasUninterpretedOptionOrBuilderList returns a boolean if a field has been set.

### GetUnknownFields

`func (o *FieldOptionsOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *FieldOptionsOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *FieldOptionsOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *FieldOptionsOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetUnverifiedLazy

`func (o *FieldOptionsOrBuilder) GetUnverifiedLazy() bool`

GetUnverifiedLazy returns the UnverifiedLazy field if non-nil, zero value otherwise.

### GetUnverifiedLazyOk

`func (o *FieldOptionsOrBuilder) GetUnverifiedLazyOk() (*bool, bool)`

GetUnverifiedLazyOk returns a tuple with the UnverifiedLazy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnverifiedLazy

`func (o *FieldOptionsOrBuilder) SetUnverifiedLazy(v bool)`

SetUnverifiedLazy sets UnverifiedLazy field to given value.

### HasUnverifiedLazy

`func (o *FieldOptionsOrBuilder) HasUnverifiedLazy() bool`

HasUnverifiedLazy returns a boolean if a field has been set.

### GetWeak

`func (o *FieldOptionsOrBuilder) GetWeak() bool`

GetWeak returns the Weak field if non-nil, zero value otherwise.

### GetWeakOk

`func (o *FieldOptionsOrBuilder) GetWeakOk() (*bool, bool)`

GetWeakOk returns a tuple with the Weak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeak

`func (o *FieldOptionsOrBuilder) SetWeak(v bool)`

SetWeak sets Weak field to given value.

### HasWeak

`func (o *FieldOptionsOrBuilder) HasWeak() bool`

HasWeak returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


