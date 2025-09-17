# OneofOptionsOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to  |  | [optional] 
**DefaultInstanceForType** | Pointer to  | Simplified schema for Message (original had circular references) | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**Features** | Pointer to [**FeatureSet**](FeatureSet.md) |  | [optional] 
**FeaturesOrBuilder** | Pointer to [**FeatureSetOrBuilder**](FeatureSetOrBuilder.md) |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**UninterpretedOptionCount** | Pointer to **int32** |  | [optional] 
**UninterpretedOptionList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UninterpretedOptionOrBuilderList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewOneofOptionsOrBuilder

`func NewOneofOptionsOrBuilder() *OneofOptionsOrBuilder`

NewOneofOptionsOrBuilder instantiates a new OneofOptionsOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneofOptionsOrBuilderWithDefaults

`func NewOneofOptionsOrBuilderWithDefaults() *OneofOptionsOrBuilder`

NewOneofOptionsOrBuilderWithDefaults instantiates a new OneofOptionsOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *OneofOptionsOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *OneofOptionsOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *OneofOptionsOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *OneofOptionsOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### SetAllFieldsNil

`func (o *OneofOptionsOrBuilder) SetAllFieldsNil(b bool)`

 SetAllFieldsNil sets the value for AllFields to be an explicit nil

### UnsetAllFields
`func (o *OneofOptionsOrBuilder) UnsetAllFields()`

UnsetAllFields ensures that no value is present for AllFields, not even an explicit nil
### GetDefaultInstanceForType

`func (o *OneofOptionsOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *OneofOptionsOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *OneofOptionsOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *OneofOptionsOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### SetDefaultInstanceForTypeNil

`func (o *OneofOptionsOrBuilder) SetDefaultInstanceForTypeNil(b bool)`

 SetDefaultInstanceForTypeNil sets the value for DefaultInstanceForType to be an explicit nil

### UnsetDefaultInstanceForType
`func (o *OneofOptionsOrBuilder) UnsetDefaultInstanceForType()`

UnsetDefaultInstanceForType ensures that no value is present for DefaultInstanceForType, not even an explicit nil
### GetDescriptorForType

`func (o *OneofOptionsOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *OneofOptionsOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *OneofOptionsOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *OneofOptionsOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetFeatures

`func (o *OneofOptionsOrBuilder) GetFeatures() FeatureSet`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *OneofOptionsOrBuilder) GetFeaturesOk() (*FeatureSet, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *OneofOptionsOrBuilder) SetFeatures(v FeatureSet)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *OneofOptionsOrBuilder) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesOrBuilder

`func (o *OneofOptionsOrBuilder) GetFeaturesOrBuilder() FeatureSetOrBuilder`

GetFeaturesOrBuilder returns the FeaturesOrBuilder field if non-nil, zero value otherwise.

### GetFeaturesOrBuilderOk

`func (o *OneofOptionsOrBuilder) GetFeaturesOrBuilderOk() (*FeatureSetOrBuilder, bool)`

GetFeaturesOrBuilderOk returns a tuple with the FeaturesOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesOrBuilder

`func (o *OneofOptionsOrBuilder) SetFeaturesOrBuilder(v FeatureSetOrBuilder)`

SetFeaturesOrBuilder sets FeaturesOrBuilder field to given value.

### HasFeaturesOrBuilder

`func (o *OneofOptionsOrBuilder) HasFeaturesOrBuilder() bool`

HasFeaturesOrBuilder returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *OneofOptionsOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *OneofOptionsOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *OneofOptionsOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *OneofOptionsOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *OneofOptionsOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *OneofOptionsOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *OneofOptionsOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *OneofOptionsOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetUninterpretedOptionCount

`func (o *OneofOptionsOrBuilder) GetUninterpretedOptionCount() int32`

GetUninterpretedOptionCount returns the UninterpretedOptionCount field if non-nil, zero value otherwise.

### GetUninterpretedOptionCountOk

`func (o *OneofOptionsOrBuilder) GetUninterpretedOptionCountOk() (*int32, bool)`

GetUninterpretedOptionCountOk returns a tuple with the UninterpretedOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionCount

`func (o *OneofOptionsOrBuilder) SetUninterpretedOptionCount(v int32)`

SetUninterpretedOptionCount sets UninterpretedOptionCount field to given value.

### HasUninterpretedOptionCount

`func (o *OneofOptionsOrBuilder) HasUninterpretedOptionCount() bool`

HasUninterpretedOptionCount returns a boolean if a field has been set.

### GetUninterpretedOptionList

`func (o *OneofOptionsOrBuilder) GetUninterpretedOptionList() []map[string]interface{}`

GetUninterpretedOptionList returns the UninterpretedOptionList field if non-nil, zero value otherwise.

### GetUninterpretedOptionListOk

`func (o *OneofOptionsOrBuilder) GetUninterpretedOptionListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionListOk returns a tuple with the UninterpretedOptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionList

`func (o *OneofOptionsOrBuilder) SetUninterpretedOptionList(v []map[string]interface{})`

SetUninterpretedOptionList sets UninterpretedOptionList field to given value.

### HasUninterpretedOptionList

`func (o *OneofOptionsOrBuilder) HasUninterpretedOptionList() bool`

HasUninterpretedOptionList returns a boolean if a field has been set.

### GetUninterpretedOptionOrBuilderList

`func (o *OneofOptionsOrBuilder) GetUninterpretedOptionOrBuilderList() []map[string]interface{}`

GetUninterpretedOptionOrBuilderList returns the UninterpretedOptionOrBuilderList field if non-nil, zero value otherwise.

### GetUninterpretedOptionOrBuilderListOk

`func (o *OneofOptionsOrBuilder) GetUninterpretedOptionOrBuilderListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionOrBuilderListOk returns a tuple with the UninterpretedOptionOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionOrBuilderList

`func (o *OneofOptionsOrBuilder) SetUninterpretedOptionOrBuilderList(v []map[string]interface{})`

SetUninterpretedOptionOrBuilderList sets UninterpretedOptionOrBuilderList field to given value.

### HasUninterpretedOptionOrBuilderList

`func (o *OneofOptionsOrBuilder) HasUninterpretedOptionOrBuilderList() bool`

HasUninterpretedOptionOrBuilderList returns a boolean if a field has been set.

### GetUnknownFields

`func (o *OneofOptionsOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *OneofOptionsOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *OneofOptionsOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *OneofOptionsOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


