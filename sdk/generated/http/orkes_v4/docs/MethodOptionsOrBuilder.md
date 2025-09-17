# MethodOptionsOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to  |  | [optional] 
**DefaultInstanceForType** | Pointer to  | Simplified schema for Message (original had circular references) | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**Features** | Pointer to [**FeatureSet**](FeatureSet.md) |  | [optional] 
**FeaturesOrBuilder** | Pointer to [**FeatureSetOrBuilder**](FeatureSetOrBuilder.md) |  | [optional] 
**IdempotencyLevel** | Pointer to **string** |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**UninterpretedOptionCount** | Pointer to **int32** |  | [optional] 
**UninterpretedOptionList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UninterpretedOptionOrBuilderList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewMethodOptionsOrBuilder

`func NewMethodOptionsOrBuilder() *MethodOptionsOrBuilder`

NewMethodOptionsOrBuilder instantiates a new MethodOptionsOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMethodOptionsOrBuilderWithDefaults

`func NewMethodOptionsOrBuilderWithDefaults() *MethodOptionsOrBuilder`

NewMethodOptionsOrBuilderWithDefaults instantiates a new MethodOptionsOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *MethodOptionsOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *MethodOptionsOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *MethodOptionsOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *MethodOptionsOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### SetAllFieldsNil

`func (o *MethodOptionsOrBuilder) SetAllFieldsNil(b bool)`

 SetAllFieldsNil sets the value for AllFields to be an explicit nil

### UnsetAllFields
`func (o *MethodOptionsOrBuilder) UnsetAllFields()`

UnsetAllFields ensures that no value is present for AllFields, not even an explicit nil
### GetDefaultInstanceForType

`func (o *MethodOptionsOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *MethodOptionsOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *MethodOptionsOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *MethodOptionsOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### SetDefaultInstanceForTypeNil

`func (o *MethodOptionsOrBuilder) SetDefaultInstanceForTypeNil(b bool)`

 SetDefaultInstanceForTypeNil sets the value for DefaultInstanceForType to be an explicit nil

### UnsetDefaultInstanceForType
`func (o *MethodOptionsOrBuilder) UnsetDefaultInstanceForType()`

UnsetDefaultInstanceForType ensures that no value is present for DefaultInstanceForType, not even an explicit nil
### GetDeprecated

`func (o *MethodOptionsOrBuilder) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *MethodOptionsOrBuilder) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *MethodOptionsOrBuilder) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *MethodOptionsOrBuilder) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *MethodOptionsOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *MethodOptionsOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *MethodOptionsOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *MethodOptionsOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetFeatures

`func (o *MethodOptionsOrBuilder) GetFeatures() FeatureSet`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *MethodOptionsOrBuilder) GetFeaturesOk() (*FeatureSet, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *MethodOptionsOrBuilder) SetFeatures(v FeatureSet)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *MethodOptionsOrBuilder) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesOrBuilder

`func (o *MethodOptionsOrBuilder) GetFeaturesOrBuilder() FeatureSetOrBuilder`

GetFeaturesOrBuilder returns the FeaturesOrBuilder field if non-nil, zero value otherwise.

### GetFeaturesOrBuilderOk

`func (o *MethodOptionsOrBuilder) GetFeaturesOrBuilderOk() (*FeatureSetOrBuilder, bool)`

GetFeaturesOrBuilderOk returns a tuple with the FeaturesOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesOrBuilder

`func (o *MethodOptionsOrBuilder) SetFeaturesOrBuilder(v FeatureSetOrBuilder)`

SetFeaturesOrBuilder sets FeaturesOrBuilder field to given value.

### HasFeaturesOrBuilder

`func (o *MethodOptionsOrBuilder) HasFeaturesOrBuilder() bool`

HasFeaturesOrBuilder returns a boolean if a field has been set.

### GetIdempotencyLevel

`func (o *MethodOptionsOrBuilder) GetIdempotencyLevel() string`

GetIdempotencyLevel returns the IdempotencyLevel field if non-nil, zero value otherwise.

### GetIdempotencyLevelOk

`func (o *MethodOptionsOrBuilder) GetIdempotencyLevelOk() (*string, bool)`

GetIdempotencyLevelOk returns a tuple with the IdempotencyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyLevel

`func (o *MethodOptionsOrBuilder) SetIdempotencyLevel(v string)`

SetIdempotencyLevel sets IdempotencyLevel field to given value.

### HasIdempotencyLevel

`func (o *MethodOptionsOrBuilder) HasIdempotencyLevel() bool`

HasIdempotencyLevel returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *MethodOptionsOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *MethodOptionsOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *MethodOptionsOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *MethodOptionsOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *MethodOptionsOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *MethodOptionsOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *MethodOptionsOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *MethodOptionsOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetUninterpretedOptionCount

`func (o *MethodOptionsOrBuilder) GetUninterpretedOptionCount() int32`

GetUninterpretedOptionCount returns the UninterpretedOptionCount field if non-nil, zero value otherwise.

### GetUninterpretedOptionCountOk

`func (o *MethodOptionsOrBuilder) GetUninterpretedOptionCountOk() (*int32, bool)`

GetUninterpretedOptionCountOk returns a tuple with the UninterpretedOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionCount

`func (o *MethodOptionsOrBuilder) SetUninterpretedOptionCount(v int32)`

SetUninterpretedOptionCount sets UninterpretedOptionCount field to given value.

### HasUninterpretedOptionCount

`func (o *MethodOptionsOrBuilder) HasUninterpretedOptionCount() bool`

HasUninterpretedOptionCount returns a boolean if a field has been set.

### GetUninterpretedOptionList

`func (o *MethodOptionsOrBuilder) GetUninterpretedOptionList() []map[string]interface{}`

GetUninterpretedOptionList returns the UninterpretedOptionList field if non-nil, zero value otherwise.

### GetUninterpretedOptionListOk

`func (o *MethodOptionsOrBuilder) GetUninterpretedOptionListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionListOk returns a tuple with the UninterpretedOptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionList

`func (o *MethodOptionsOrBuilder) SetUninterpretedOptionList(v []map[string]interface{})`

SetUninterpretedOptionList sets UninterpretedOptionList field to given value.

### HasUninterpretedOptionList

`func (o *MethodOptionsOrBuilder) HasUninterpretedOptionList() bool`

HasUninterpretedOptionList returns a boolean if a field has been set.

### GetUninterpretedOptionOrBuilderList

`func (o *MethodOptionsOrBuilder) GetUninterpretedOptionOrBuilderList() []map[string]interface{}`

GetUninterpretedOptionOrBuilderList returns the UninterpretedOptionOrBuilderList field if non-nil, zero value otherwise.

### GetUninterpretedOptionOrBuilderListOk

`func (o *MethodOptionsOrBuilder) GetUninterpretedOptionOrBuilderListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionOrBuilderListOk returns a tuple with the UninterpretedOptionOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionOrBuilderList

`func (o *MethodOptionsOrBuilder) SetUninterpretedOptionOrBuilderList(v []map[string]interface{})`

SetUninterpretedOptionOrBuilderList sets UninterpretedOptionOrBuilderList field to given value.

### HasUninterpretedOptionOrBuilderList

`func (o *MethodOptionsOrBuilder) HasUninterpretedOptionOrBuilderList() bool`

HasUninterpretedOptionOrBuilderList returns a boolean if a field has been set.

### GetUnknownFields

`func (o *MethodOptionsOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *MethodOptionsOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *MethodOptionsOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *MethodOptionsOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


