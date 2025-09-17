# ServiceOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**DefaultInstanceForType** | Pointer to [**ServiceOptions**](ServiceOptions.md) |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** |  | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to [**FeatureSet**](FeatureSet.md) |  | [optional] 
**UninterpretedOptionOrBuilderList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**FeaturesOrBuilder** | Pointer to [**FeatureSetOrBuilder**](FeatureSetOrBuilder.md) |  | [optional] 
**UninterpretedOptionList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UninterpretedOptionCount** | Pointer to **int32** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**AllFieldsRaw** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**MemoizedSerializedSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceOptions

`func NewServiceOptions() *ServiceOptions`

NewServiceOptions instantiates a new ServiceOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOptionsWithDefaults

`func NewServiceOptionsWithDefaults() *ServiceOptions`

NewServiceOptionsWithDefaults instantiates a new ServiceOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnknownFields

`func (o *ServiceOptions) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *ServiceOptions) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *ServiceOptions) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *ServiceOptions) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *ServiceOptions) GetDefaultInstanceForType() ServiceOptions`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *ServiceOptions) GetDefaultInstanceForTypeOk() (*ServiceOptions, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *ServiceOptions) SetDefaultInstanceForType(v ServiceOptions)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *ServiceOptions) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetParserForType

`func (o *ServiceOptions) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *ServiceOptions) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *ServiceOptions) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *ServiceOptions) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetSerializedSize

`func (o *ServiceOptions) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *ServiceOptions) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *ServiceOptions) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *ServiceOptions) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetDeprecated

`func (o *ServiceOptions) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *ServiceOptions) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *ServiceOptions) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *ServiceOptions) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetFeatures

`func (o *ServiceOptions) GetFeatures() FeatureSet`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ServiceOptions) GetFeaturesOk() (*FeatureSet, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ServiceOptions) SetFeatures(v FeatureSet)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ServiceOptions) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetUninterpretedOptionOrBuilderList

`func (o *ServiceOptions) GetUninterpretedOptionOrBuilderList() []map[string]interface{}`

GetUninterpretedOptionOrBuilderList returns the UninterpretedOptionOrBuilderList field if non-nil, zero value otherwise.

### GetUninterpretedOptionOrBuilderListOk

`func (o *ServiceOptions) GetUninterpretedOptionOrBuilderListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionOrBuilderListOk returns a tuple with the UninterpretedOptionOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionOrBuilderList

`func (o *ServiceOptions) SetUninterpretedOptionOrBuilderList(v []map[string]interface{})`

SetUninterpretedOptionOrBuilderList sets UninterpretedOptionOrBuilderList field to given value.

### HasUninterpretedOptionOrBuilderList

`func (o *ServiceOptions) HasUninterpretedOptionOrBuilderList() bool`

HasUninterpretedOptionOrBuilderList returns a boolean if a field has been set.

### GetFeaturesOrBuilder

`func (o *ServiceOptions) GetFeaturesOrBuilder() FeatureSetOrBuilder`

GetFeaturesOrBuilder returns the FeaturesOrBuilder field if non-nil, zero value otherwise.

### GetFeaturesOrBuilderOk

`func (o *ServiceOptions) GetFeaturesOrBuilderOk() (*FeatureSetOrBuilder, bool)`

GetFeaturesOrBuilderOk returns a tuple with the FeaturesOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesOrBuilder

`func (o *ServiceOptions) SetFeaturesOrBuilder(v FeatureSetOrBuilder)`

SetFeaturesOrBuilder sets FeaturesOrBuilder field to given value.

### HasFeaturesOrBuilder

`func (o *ServiceOptions) HasFeaturesOrBuilder() bool`

HasFeaturesOrBuilder returns a boolean if a field has been set.

### GetUninterpretedOptionList

`func (o *ServiceOptions) GetUninterpretedOptionList() []map[string]interface{}`

GetUninterpretedOptionList returns the UninterpretedOptionList field if non-nil, zero value otherwise.

### GetUninterpretedOptionListOk

`func (o *ServiceOptions) GetUninterpretedOptionListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionListOk returns a tuple with the UninterpretedOptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionList

`func (o *ServiceOptions) SetUninterpretedOptionList(v []map[string]interface{})`

SetUninterpretedOptionList sets UninterpretedOptionList field to given value.

### HasUninterpretedOptionList

`func (o *ServiceOptions) HasUninterpretedOptionList() bool`

HasUninterpretedOptionList returns a boolean if a field has been set.

### GetUninterpretedOptionCount

`func (o *ServiceOptions) GetUninterpretedOptionCount() int32`

GetUninterpretedOptionCount returns the UninterpretedOptionCount field if non-nil, zero value otherwise.

### GetUninterpretedOptionCountOk

`func (o *ServiceOptions) GetUninterpretedOptionCountOk() (*int32, bool)`

GetUninterpretedOptionCountOk returns a tuple with the UninterpretedOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionCount

`func (o *ServiceOptions) SetUninterpretedOptionCount(v int32)`

SetUninterpretedOptionCount sets UninterpretedOptionCount field to given value.

### HasUninterpretedOptionCount

`func (o *ServiceOptions) HasUninterpretedOptionCount() bool`

HasUninterpretedOptionCount returns a boolean if a field has been set.

### GetInitialized

`func (o *ServiceOptions) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *ServiceOptions) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *ServiceOptions) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *ServiceOptions) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *ServiceOptions) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *ServiceOptions) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *ServiceOptions) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *ServiceOptions) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *ServiceOptions) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *ServiceOptions) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *ServiceOptions) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *ServiceOptions) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetAllFields

`func (o *ServiceOptions) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *ServiceOptions) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *ServiceOptions) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *ServiceOptions) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetAllFieldsRaw

`func (o *ServiceOptions) GetAllFieldsRaw() map[string]map[string]interface{}`

GetAllFieldsRaw returns the AllFieldsRaw field if non-nil, zero value otherwise.

### GetAllFieldsRawOk

`func (o *ServiceOptions) GetAllFieldsRawOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsRawOk returns a tuple with the AllFieldsRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFieldsRaw

`func (o *ServiceOptions) SetAllFieldsRaw(v map[string]map[string]interface{})`

SetAllFieldsRaw sets AllFieldsRaw field to given value.

### HasAllFieldsRaw

`func (o *ServiceOptions) HasAllFieldsRaw() bool`

HasAllFieldsRaw returns a boolean if a field has been set.

### GetMemoizedSerializedSize

`func (o *ServiceOptions) GetMemoizedSerializedSize() int32`

GetMemoizedSerializedSize returns the MemoizedSerializedSize field if non-nil, zero value otherwise.

### GetMemoizedSerializedSizeOk

`func (o *ServiceOptions) GetMemoizedSerializedSizeOk() (*int32, bool)`

GetMemoizedSerializedSizeOk returns a tuple with the MemoizedSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizedSerializedSize

`func (o *ServiceOptions) SetMemoizedSerializedSize(v int32)`

SetMemoizedSerializedSize sets MemoizedSerializedSize field to given value.

### HasMemoizedSerializedSize

`func (o *ServiceOptions) HasMemoizedSerializedSize() bool`

HasMemoizedSerializedSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


