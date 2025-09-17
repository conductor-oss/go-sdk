# FieldOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**AllFieldsRaw** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Ctype** | Pointer to **string** |  | [optional] 
**DebugRedact** | Pointer to **bool** |  | [optional] 
**DefaultInstanceForType** | Pointer to [**FieldOptions**](FieldOptions.md) |  | [optional] 
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
**MemoizedSerializedSize** | Pointer to **int32** |  | [optional] 
**Packed** | Pointer to **bool** |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** |  | [optional] 
**Retention** | Pointer to **string** |  | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**TargetsCount** | Pointer to **int32** |  | [optional] 
**TargetsList** | Pointer to **[]string** |  | [optional] 
**UninterpretedOptionCount** | Pointer to **int32** |  | [optional] 
**UninterpretedOptionList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UninterpretedOptionOrBuilderList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**UnverifiedLazy** | Pointer to **bool** |  | [optional] 
**Weak** | Pointer to **bool** |  | [optional] 

## Methods

### NewFieldOptions

`func NewFieldOptions() *FieldOptions`

NewFieldOptions instantiates a new FieldOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldOptionsWithDefaults

`func NewFieldOptionsWithDefaults() *FieldOptions`

NewFieldOptionsWithDefaults instantiates a new FieldOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *FieldOptions) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *FieldOptions) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *FieldOptions) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *FieldOptions) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetAllFieldsRaw

`func (o *FieldOptions) GetAllFieldsRaw() map[string]map[string]interface{}`

GetAllFieldsRaw returns the AllFieldsRaw field if non-nil, zero value otherwise.

### GetAllFieldsRawOk

`func (o *FieldOptions) GetAllFieldsRawOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsRawOk returns a tuple with the AllFieldsRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFieldsRaw

`func (o *FieldOptions) SetAllFieldsRaw(v map[string]map[string]interface{})`

SetAllFieldsRaw sets AllFieldsRaw field to given value.

### HasAllFieldsRaw

`func (o *FieldOptions) HasAllFieldsRaw() bool`

HasAllFieldsRaw returns a boolean if a field has been set.

### GetCtype

`func (o *FieldOptions) GetCtype() string`

GetCtype returns the Ctype field if non-nil, zero value otherwise.

### GetCtypeOk

`func (o *FieldOptions) GetCtypeOk() (*string, bool)`

GetCtypeOk returns a tuple with the Ctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtype

`func (o *FieldOptions) SetCtype(v string)`

SetCtype sets Ctype field to given value.

### HasCtype

`func (o *FieldOptions) HasCtype() bool`

HasCtype returns a boolean if a field has been set.

### GetDebugRedact

`func (o *FieldOptions) GetDebugRedact() bool`

GetDebugRedact returns the DebugRedact field if non-nil, zero value otherwise.

### GetDebugRedactOk

`func (o *FieldOptions) GetDebugRedactOk() (*bool, bool)`

GetDebugRedactOk returns a tuple with the DebugRedact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugRedact

`func (o *FieldOptions) SetDebugRedact(v bool)`

SetDebugRedact sets DebugRedact field to given value.

### HasDebugRedact

`func (o *FieldOptions) HasDebugRedact() bool`

HasDebugRedact returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *FieldOptions) GetDefaultInstanceForType() FieldOptions`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *FieldOptions) GetDefaultInstanceForTypeOk() (*FieldOptions, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *FieldOptions) SetDefaultInstanceForType(v FieldOptions)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *FieldOptions) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDeprecated

`func (o *FieldOptions) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *FieldOptions) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *FieldOptions) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *FieldOptions) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *FieldOptions) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *FieldOptions) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *FieldOptions) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *FieldOptions) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetEditionDefaultsCount

`func (o *FieldOptions) GetEditionDefaultsCount() int32`

GetEditionDefaultsCount returns the EditionDefaultsCount field if non-nil, zero value otherwise.

### GetEditionDefaultsCountOk

`func (o *FieldOptions) GetEditionDefaultsCountOk() (*int32, bool)`

GetEditionDefaultsCountOk returns a tuple with the EditionDefaultsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionDefaultsCount

`func (o *FieldOptions) SetEditionDefaultsCount(v int32)`

SetEditionDefaultsCount sets EditionDefaultsCount field to given value.

### HasEditionDefaultsCount

`func (o *FieldOptions) HasEditionDefaultsCount() bool`

HasEditionDefaultsCount returns a boolean if a field has been set.

### GetEditionDefaultsList

`func (o *FieldOptions) GetEditionDefaultsList() []EditionDefault`

GetEditionDefaultsList returns the EditionDefaultsList field if non-nil, zero value otherwise.

### GetEditionDefaultsListOk

`func (o *FieldOptions) GetEditionDefaultsListOk() (*[]EditionDefault, bool)`

GetEditionDefaultsListOk returns a tuple with the EditionDefaultsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionDefaultsList

`func (o *FieldOptions) SetEditionDefaultsList(v []EditionDefault)`

SetEditionDefaultsList sets EditionDefaultsList field to given value.

### HasEditionDefaultsList

`func (o *FieldOptions) HasEditionDefaultsList() bool`

HasEditionDefaultsList returns a boolean if a field has been set.

### GetEditionDefaultsOrBuilderList

`func (o *FieldOptions) GetEditionDefaultsOrBuilderList() []EditionDefaultOrBuilder`

GetEditionDefaultsOrBuilderList returns the EditionDefaultsOrBuilderList field if non-nil, zero value otherwise.

### GetEditionDefaultsOrBuilderListOk

`func (o *FieldOptions) GetEditionDefaultsOrBuilderListOk() (*[]EditionDefaultOrBuilder, bool)`

GetEditionDefaultsOrBuilderListOk returns a tuple with the EditionDefaultsOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionDefaultsOrBuilderList

`func (o *FieldOptions) SetEditionDefaultsOrBuilderList(v []EditionDefaultOrBuilder)`

SetEditionDefaultsOrBuilderList sets EditionDefaultsOrBuilderList field to given value.

### HasEditionDefaultsOrBuilderList

`func (o *FieldOptions) HasEditionDefaultsOrBuilderList() bool`

HasEditionDefaultsOrBuilderList returns a boolean if a field has been set.

### GetFeatures

`func (o *FieldOptions) GetFeatures() FeatureSet`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FieldOptions) GetFeaturesOk() (*FeatureSet, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FieldOptions) SetFeatures(v FeatureSet)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FieldOptions) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesOrBuilder

`func (o *FieldOptions) GetFeaturesOrBuilder() FeatureSetOrBuilder`

GetFeaturesOrBuilder returns the FeaturesOrBuilder field if non-nil, zero value otherwise.

### GetFeaturesOrBuilderOk

`func (o *FieldOptions) GetFeaturesOrBuilderOk() (*FeatureSetOrBuilder, bool)`

GetFeaturesOrBuilderOk returns a tuple with the FeaturesOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesOrBuilder

`func (o *FieldOptions) SetFeaturesOrBuilder(v FeatureSetOrBuilder)`

SetFeaturesOrBuilder sets FeaturesOrBuilder field to given value.

### HasFeaturesOrBuilder

`func (o *FieldOptions) HasFeaturesOrBuilder() bool`

HasFeaturesOrBuilder returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *FieldOptions) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *FieldOptions) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *FieldOptions) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *FieldOptions) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *FieldOptions) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *FieldOptions) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *FieldOptions) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *FieldOptions) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetJstype

`func (o *FieldOptions) GetJstype() string`

GetJstype returns the Jstype field if non-nil, zero value otherwise.

### GetJstypeOk

`func (o *FieldOptions) GetJstypeOk() (*string, bool)`

GetJstypeOk returns a tuple with the Jstype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJstype

`func (o *FieldOptions) SetJstype(v string)`

SetJstype sets Jstype field to given value.

### HasJstype

`func (o *FieldOptions) HasJstype() bool`

HasJstype returns a boolean if a field has been set.

### GetLazy

`func (o *FieldOptions) GetLazy() bool`

GetLazy returns the Lazy field if non-nil, zero value otherwise.

### GetLazyOk

`func (o *FieldOptions) GetLazyOk() (*bool, bool)`

GetLazyOk returns a tuple with the Lazy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLazy

`func (o *FieldOptions) SetLazy(v bool)`

SetLazy sets Lazy field to given value.

### HasLazy

`func (o *FieldOptions) HasLazy() bool`

HasLazy returns a boolean if a field has been set.

### GetMemoizedSerializedSize

`func (o *FieldOptions) GetMemoizedSerializedSize() int32`

GetMemoizedSerializedSize returns the MemoizedSerializedSize field if non-nil, zero value otherwise.

### GetMemoizedSerializedSizeOk

`func (o *FieldOptions) GetMemoizedSerializedSizeOk() (*int32, bool)`

GetMemoizedSerializedSizeOk returns a tuple with the MemoizedSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizedSerializedSize

`func (o *FieldOptions) SetMemoizedSerializedSize(v int32)`

SetMemoizedSerializedSize sets MemoizedSerializedSize field to given value.

### HasMemoizedSerializedSize

`func (o *FieldOptions) HasMemoizedSerializedSize() bool`

HasMemoizedSerializedSize returns a boolean if a field has been set.

### GetPacked

`func (o *FieldOptions) GetPacked() bool`

GetPacked returns the Packed field if non-nil, zero value otherwise.

### GetPackedOk

`func (o *FieldOptions) GetPackedOk() (*bool, bool)`

GetPackedOk returns a tuple with the Packed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacked

`func (o *FieldOptions) SetPacked(v bool)`

SetPacked sets Packed field to given value.

### HasPacked

`func (o *FieldOptions) HasPacked() bool`

HasPacked returns a boolean if a field has been set.

### GetParserForType

`func (o *FieldOptions) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *FieldOptions) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *FieldOptions) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *FieldOptions) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetRetention

`func (o *FieldOptions) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *FieldOptions) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *FieldOptions) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *FieldOptions) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetSerializedSize

`func (o *FieldOptions) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *FieldOptions) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *FieldOptions) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *FieldOptions) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetTargetsCount

`func (o *FieldOptions) GetTargetsCount() int32`

GetTargetsCount returns the TargetsCount field if non-nil, zero value otherwise.

### GetTargetsCountOk

`func (o *FieldOptions) GetTargetsCountOk() (*int32, bool)`

GetTargetsCountOk returns a tuple with the TargetsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsCount

`func (o *FieldOptions) SetTargetsCount(v int32)`

SetTargetsCount sets TargetsCount field to given value.

### HasTargetsCount

`func (o *FieldOptions) HasTargetsCount() bool`

HasTargetsCount returns a boolean if a field has been set.

### GetTargetsList

`func (o *FieldOptions) GetTargetsList() []string`

GetTargetsList returns the TargetsList field if non-nil, zero value otherwise.

### GetTargetsListOk

`func (o *FieldOptions) GetTargetsListOk() (*[]string, bool)`

GetTargetsListOk returns a tuple with the TargetsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsList

`func (o *FieldOptions) SetTargetsList(v []string)`

SetTargetsList sets TargetsList field to given value.

### HasTargetsList

`func (o *FieldOptions) HasTargetsList() bool`

HasTargetsList returns a boolean if a field has been set.

### GetUninterpretedOptionCount

`func (o *FieldOptions) GetUninterpretedOptionCount() int32`

GetUninterpretedOptionCount returns the UninterpretedOptionCount field if non-nil, zero value otherwise.

### GetUninterpretedOptionCountOk

`func (o *FieldOptions) GetUninterpretedOptionCountOk() (*int32, bool)`

GetUninterpretedOptionCountOk returns a tuple with the UninterpretedOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionCount

`func (o *FieldOptions) SetUninterpretedOptionCount(v int32)`

SetUninterpretedOptionCount sets UninterpretedOptionCount field to given value.

### HasUninterpretedOptionCount

`func (o *FieldOptions) HasUninterpretedOptionCount() bool`

HasUninterpretedOptionCount returns a boolean if a field has been set.

### GetUninterpretedOptionList

`func (o *FieldOptions) GetUninterpretedOptionList() []map[string]interface{}`

GetUninterpretedOptionList returns the UninterpretedOptionList field if non-nil, zero value otherwise.

### GetUninterpretedOptionListOk

`func (o *FieldOptions) GetUninterpretedOptionListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionListOk returns a tuple with the UninterpretedOptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionList

`func (o *FieldOptions) SetUninterpretedOptionList(v []map[string]interface{})`

SetUninterpretedOptionList sets UninterpretedOptionList field to given value.

### HasUninterpretedOptionList

`func (o *FieldOptions) HasUninterpretedOptionList() bool`

HasUninterpretedOptionList returns a boolean if a field has been set.

### GetUninterpretedOptionOrBuilderList

`func (o *FieldOptions) GetUninterpretedOptionOrBuilderList() []map[string]interface{}`

GetUninterpretedOptionOrBuilderList returns the UninterpretedOptionOrBuilderList field if non-nil, zero value otherwise.

### GetUninterpretedOptionOrBuilderListOk

`func (o *FieldOptions) GetUninterpretedOptionOrBuilderListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionOrBuilderListOk returns a tuple with the UninterpretedOptionOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionOrBuilderList

`func (o *FieldOptions) SetUninterpretedOptionOrBuilderList(v []map[string]interface{})`

SetUninterpretedOptionOrBuilderList sets UninterpretedOptionOrBuilderList field to given value.

### HasUninterpretedOptionOrBuilderList

`func (o *FieldOptions) HasUninterpretedOptionOrBuilderList() bool`

HasUninterpretedOptionOrBuilderList returns a boolean if a field has been set.

### GetUnknownFields

`func (o *FieldOptions) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *FieldOptions) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *FieldOptions) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *FieldOptions) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetUnverifiedLazy

`func (o *FieldOptions) GetUnverifiedLazy() bool`

GetUnverifiedLazy returns the UnverifiedLazy field if non-nil, zero value otherwise.

### GetUnverifiedLazyOk

`func (o *FieldOptions) GetUnverifiedLazyOk() (*bool, bool)`

GetUnverifiedLazyOk returns a tuple with the UnverifiedLazy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnverifiedLazy

`func (o *FieldOptions) SetUnverifiedLazy(v bool)`

SetUnverifiedLazy sets UnverifiedLazy field to given value.

### HasUnverifiedLazy

`func (o *FieldOptions) HasUnverifiedLazy() bool`

HasUnverifiedLazy returns a boolean if a field has been set.

### GetWeak

`func (o *FieldOptions) GetWeak() bool`

GetWeak returns the Weak field if non-nil, zero value otherwise.

### GetWeakOk

`func (o *FieldOptions) GetWeakOk() (*bool, bool)`

GetWeakOk returns a tuple with the Weak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeak

`func (o *FieldOptions) SetWeak(v bool)`

SetWeak sets Weak field to given value.

### HasWeak

`func (o *FieldOptions) HasWeak() bool`

HasWeak returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


