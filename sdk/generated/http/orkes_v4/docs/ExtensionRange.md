# ExtensionRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to  |  | [optional] 
**DefaultInstanceForType** | Pointer to [**ExtensionRange**](ExtensionRange.md) |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**MemoizedSerializedSize** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**ExtensionRangeOptions**](ExtensionRangeOptions.md) |  | [optional] 
**OptionsOrBuilder** | Pointer to [**ExtensionRangeOptionsOrBuilder**](ExtensionRangeOptionsOrBuilder.md) |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** |  | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewExtensionRange

`func NewExtensionRange() *ExtensionRange`

NewExtensionRange instantiates a new ExtensionRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionRangeWithDefaults

`func NewExtensionRangeWithDefaults() *ExtensionRange`

NewExtensionRangeWithDefaults instantiates a new ExtensionRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *ExtensionRange) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *ExtensionRange) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *ExtensionRange) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *ExtensionRange) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### SetAllFieldsNil

`func (o *ExtensionRange) SetAllFieldsNil(b bool)`

 SetAllFieldsNil sets the value for AllFields to be an explicit nil

### UnsetAllFields
`func (o *ExtensionRange) UnsetAllFields()`

UnsetAllFields ensures that no value is present for AllFields, not even an explicit nil
### GetDefaultInstanceForType

`func (o *ExtensionRange) GetDefaultInstanceForType() ExtensionRange`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *ExtensionRange) GetDefaultInstanceForTypeOk() (*ExtensionRange, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *ExtensionRange) SetDefaultInstanceForType(v ExtensionRange)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *ExtensionRange) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *ExtensionRange) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *ExtensionRange) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *ExtensionRange) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *ExtensionRange) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetEnd

`func (o *ExtensionRange) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ExtensionRange) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ExtensionRange) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ExtensionRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *ExtensionRange) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *ExtensionRange) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *ExtensionRange) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *ExtensionRange) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *ExtensionRange) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *ExtensionRange) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *ExtensionRange) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *ExtensionRange) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetMemoizedSerializedSize

`func (o *ExtensionRange) GetMemoizedSerializedSize() int32`

GetMemoizedSerializedSize returns the MemoizedSerializedSize field if non-nil, zero value otherwise.

### GetMemoizedSerializedSizeOk

`func (o *ExtensionRange) GetMemoizedSerializedSizeOk() (*int32, bool)`

GetMemoizedSerializedSizeOk returns a tuple with the MemoizedSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizedSerializedSize

`func (o *ExtensionRange) SetMemoizedSerializedSize(v int32)`

SetMemoizedSerializedSize sets MemoizedSerializedSize field to given value.

### HasMemoizedSerializedSize

`func (o *ExtensionRange) HasMemoizedSerializedSize() bool`

HasMemoizedSerializedSize returns a boolean if a field has been set.

### GetOptions

`func (o *ExtensionRange) GetOptions() ExtensionRangeOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ExtensionRange) GetOptionsOk() (*ExtensionRangeOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ExtensionRange) SetOptions(v ExtensionRangeOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ExtensionRange) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOptionsOrBuilder

`func (o *ExtensionRange) GetOptionsOrBuilder() ExtensionRangeOptionsOrBuilder`

GetOptionsOrBuilder returns the OptionsOrBuilder field if non-nil, zero value otherwise.

### GetOptionsOrBuilderOk

`func (o *ExtensionRange) GetOptionsOrBuilderOk() (*ExtensionRangeOptionsOrBuilder, bool)`

GetOptionsOrBuilderOk returns a tuple with the OptionsOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsOrBuilder

`func (o *ExtensionRange) SetOptionsOrBuilder(v ExtensionRangeOptionsOrBuilder)`

SetOptionsOrBuilder sets OptionsOrBuilder field to given value.

### HasOptionsOrBuilder

`func (o *ExtensionRange) HasOptionsOrBuilder() bool`

HasOptionsOrBuilder returns a boolean if a field has been set.

### GetParserForType

`func (o *ExtensionRange) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *ExtensionRange) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *ExtensionRange) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *ExtensionRange) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetSerializedSize

`func (o *ExtensionRange) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *ExtensionRange) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *ExtensionRange) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *ExtensionRange) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetStart

`func (o *ExtensionRange) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ExtensionRange) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ExtensionRange) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ExtensionRange) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUnknownFields

`func (o *ExtensionRange) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *ExtensionRange) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *ExtensionRange) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *ExtensionRange) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


