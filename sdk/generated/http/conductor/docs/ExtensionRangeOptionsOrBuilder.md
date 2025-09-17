# ExtensionRangeOptionsOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**FeatureSet**](FeatureSet.md) |  | [optional] 
**UninterpretedOptionOrBuilderList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**DeclarationCount** | Pointer to **int32** |  | [optional] 
**DeclarationOrBuilderList** | Pointer to [**[]DeclarationOrBuilder**](DeclarationOrBuilder.md) |  | [optional] 
**DeclarationList** | Pointer to [**[]Declaration**](Declaration.md) |  | [optional] 
**Verification** | Pointer to **string** |  | [optional] 
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

### NewExtensionRangeOptionsOrBuilder

`func NewExtensionRangeOptionsOrBuilder() *ExtensionRangeOptionsOrBuilder`

NewExtensionRangeOptionsOrBuilder instantiates a new ExtensionRangeOptionsOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionRangeOptionsOrBuilderWithDefaults

`func NewExtensionRangeOptionsOrBuilderWithDefaults() *ExtensionRangeOptionsOrBuilder`

NewExtensionRangeOptionsOrBuilderWithDefaults instantiates a new ExtensionRangeOptionsOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *ExtensionRangeOptionsOrBuilder) GetFeatures() FeatureSet`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ExtensionRangeOptionsOrBuilder) GetFeaturesOk() (*FeatureSet, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ExtensionRangeOptionsOrBuilder) SetFeatures(v FeatureSet)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ExtensionRangeOptionsOrBuilder) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetUninterpretedOptionOrBuilderList

`func (o *ExtensionRangeOptionsOrBuilder) GetUninterpretedOptionOrBuilderList() []map[string]interface{}`

GetUninterpretedOptionOrBuilderList returns the UninterpretedOptionOrBuilderList field if non-nil, zero value otherwise.

### GetUninterpretedOptionOrBuilderListOk

`func (o *ExtensionRangeOptionsOrBuilder) GetUninterpretedOptionOrBuilderListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionOrBuilderListOk returns a tuple with the UninterpretedOptionOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionOrBuilderList

`func (o *ExtensionRangeOptionsOrBuilder) SetUninterpretedOptionOrBuilderList(v []map[string]interface{})`

SetUninterpretedOptionOrBuilderList sets UninterpretedOptionOrBuilderList field to given value.

### HasUninterpretedOptionOrBuilderList

`func (o *ExtensionRangeOptionsOrBuilder) HasUninterpretedOptionOrBuilderList() bool`

HasUninterpretedOptionOrBuilderList returns a boolean if a field has been set.

### GetDeclarationCount

`func (o *ExtensionRangeOptionsOrBuilder) GetDeclarationCount() int32`

GetDeclarationCount returns the DeclarationCount field if non-nil, zero value otherwise.

### GetDeclarationCountOk

`func (o *ExtensionRangeOptionsOrBuilder) GetDeclarationCountOk() (*int32, bool)`

GetDeclarationCountOk returns a tuple with the DeclarationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarationCount

`func (o *ExtensionRangeOptionsOrBuilder) SetDeclarationCount(v int32)`

SetDeclarationCount sets DeclarationCount field to given value.

### HasDeclarationCount

`func (o *ExtensionRangeOptionsOrBuilder) HasDeclarationCount() bool`

HasDeclarationCount returns a boolean if a field has been set.

### GetDeclarationOrBuilderList

`func (o *ExtensionRangeOptionsOrBuilder) GetDeclarationOrBuilderList() []DeclarationOrBuilder`

GetDeclarationOrBuilderList returns the DeclarationOrBuilderList field if non-nil, zero value otherwise.

### GetDeclarationOrBuilderListOk

`func (o *ExtensionRangeOptionsOrBuilder) GetDeclarationOrBuilderListOk() (*[]DeclarationOrBuilder, bool)`

GetDeclarationOrBuilderListOk returns a tuple with the DeclarationOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarationOrBuilderList

`func (o *ExtensionRangeOptionsOrBuilder) SetDeclarationOrBuilderList(v []DeclarationOrBuilder)`

SetDeclarationOrBuilderList sets DeclarationOrBuilderList field to given value.

### HasDeclarationOrBuilderList

`func (o *ExtensionRangeOptionsOrBuilder) HasDeclarationOrBuilderList() bool`

HasDeclarationOrBuilderList returns a boolean if a field has been set.

### GetDeclarationList

`func (o *ExtensionRangeOptionsOrBuilder) GetDeclarationList() []Declaration`

GetDeclarationList returns the DeclarationList field if non-nil, zero value otherwise.

### GetDeclarationListOk

`func (o *ExtensionRangeOptionsOrBuilder) GetDeclarationListOk() (*[]Declaration, bool)`

GetDeclarationListOk returns a tuple with the DeclarationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarationList

`func (o *ExtensionRangeOptionsOrBuilder) SetDeclarationList(v []Declaration)`

SetDeclarationList sets DeclarationList field to given value.

### HasDeclarationList

`func (o *ExtensionRangeOptionsOrBuilder) HasDeclarationList() bool`

HasDeclarationList returns a boolean if a field has been set.

### GetVerification

`func (o *ExtensionRangeOptionsOrBuilder) GetVerification() string`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *ExtensionRangeOptionsOrBuilder) GetVerificationOk() (*string, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *ExtensionRangeOptionsOrBuilder) SetVerification(v string)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *ExtensionRangeOptionsOrBuilder) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### GetFeaturesOrBuilder

`func (o *ExtensionRangeOptionsOrBuilder) GetFeaturesOrBuilder() FeatureSetOrBuilder`

GetFeaturesOrBuilder returns the FeaturesOrBuilder field if non-nil, zero value otherwise.

### GetFeaturesOrBuilderOk

`func (o *ExtensionRangeOptionsOrBuilder) GetFeaturesOrBuilderOk() (*FeatureSetOrBuilder, bool)`

GetFeaturesOrBuilderOk returns a tuple with the FeaturesOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesOrBuilder

`func (o *ExtensionRangeOptionsOrBuilder) SetFeaturesOrBuilder(v FeatureSetOrBuilder)`

SetFeaturesOrBuilder sets FeaturesOrBuilder field to given value.

### HasFeaturesOrBuilder

`func (o *ExtensionRangeOptionsOrBuilder) HasFeaturesOrBuilder() bool`

HasFeaturesOrBuilder returns a boolean if a field has been set.

### GetUninterpretedOptionList

`func (o *ExtensionRangeOptionsOrBuilder) GetUninterpretedOptionList() []map[string]interface{}`

GetUninterpretedOptionList returns the UninterpretedOptionList field if non-nil, zero value otherwise.

### GetUninterpretedOptionListOk

`func (o *ExtensionRangeOptionsOrBuilder) GetUninterpretedOptionListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionListOk returns a tuple with the UninterpretedOptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionList

`func (o *ExtensionRangeOptionsOrBuilder) SetUninterpretedOptionList(v []map[string]interface{})`

SetUninterpretedOptionList sets UninterpretedOptionList field to given value.

### HasUninterpretedOptionList

`func (o *ExtensionRangeOptionsOrBuilder) HasUninterpretedOptionList() bool`

HasUninterpretedOptionList returns a boolean if a field has been set.

### GetUninterpretedOptionCount

`func (o *ExtensionRangeOptionsOrBuilder) GetUninterpretedOptionCount() int32`

GetUninterpretedOptionCount returns the UninterpretedOptionCount field if non-nil, zero value otherwise.

### GetUninterpretedOptionCountOk

`func (o *ExtensionRangeOptionsOrBuilder) GetUninterpretedOptionCountOk() (*int32, bool)`

GetUninterpretedOptionCountOk returns a tuple with the UninterpretedOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionCount

`func (o *ExtensionRangeOptionsOrBuilder) SetUninterpretedOptionCount(v int32)`

SetUninterpretedOptionCount sets UninterpretedOptionCount field to given value.

### HasUninterpretedOptionCount

`func (o *ExtensionRangeOptionsOrBuilder) HasUninterpretedOptionCount() bool`

HasUninterpretedOptionCount returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *ExtensionRangeOptionsOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *ExtensionRangeOptionsOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *ExtensionRangeOptionsOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *ExtensionRangeOptionsOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *ExtensionRangeOptionsOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *ExtensionRangeOptionsOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *ExtensionRangeOptionsOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *ExtensionRangeOptionsOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *ExtensionRangeOptionsOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *ExtensionRangeOptionsOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *ExtensionRangeOptionsOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *ExtensionRangeOptionsOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetAllFields

`func (o *ExtensionRangeOptionsOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *ExtensionRangeOptionsOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *ExtensionRangeOptionsOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *ExtensionRangeOptionsOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetUnknownFields

`func (o *ExtensionRangeOptionsOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *ExtensionRangeOptionsOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *ExtensionRangeOptionsOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *ExtensionRangeOptionsOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetInitialized

`func (o *ExtensionRangeOptionsOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *ExtensionRangeOptionsOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *ExtensionRangeOptionsOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *ExtensionRangeOptionsOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


