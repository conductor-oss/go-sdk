# SourceCodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**DefaultInstanceForType** | Pointer to [**SourceCodeInfo**](SourceCodeInfo.md) |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**LocationCount** | Pointer to **int32** |  | [optional] 
**LocationList** | Pointer to [**[]Location**](Location.md) |  | [optional] 
**LocationOrBuilderList** | Pointer to [**[]LocationOrBuilder**](LocationOrBuilder.md) |  | [optional] 
**MemoizedSerializedSize** | Pointer to **int32** |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** |  | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewSourceCodeInfo

`func NewSourceCodeInfo() *SourceCodeInfo`

NewSourceCodeInfo instantiates a new SourceCodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceCodeInfoWithDefaults

`func NewSourceCodeInfoWithDefaults() *SourceCodeInfo`

NewSourceCodeInfoWithDefaults instantiates a new SourceCodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *SourceCodeInfo) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *SourceCodeInfo) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *SourceCodeInfo) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *SourceCodeInfo) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *SourceCodeInfo) GetDefaultInstanceForType() SourceCodeInfo`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *SourceCodeInfo) GetDefaultInstanceForTypeOk() (*SourceCodeInfo, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *SourceCodeInfo) SetDefaultInstanceForType(v SourceCodeInfo)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *SourceCodeInfo) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *SourceCodeInfo) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *SourceCodeInfo) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *SourceCodeInfo) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *SourceCodeInfo) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *SourceCodeInfo) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *SourceCodeInfo) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *SourceCodeInfo) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *SourceCodeInfo) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *SourceCodeInfo) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *SourceCodeInfo) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *SourceCodeInfo) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *SourceCodeInfo) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetLocationCount

`func (o *SourceCodeInfo) GetLocationCount() int32`

GetLocationCount returns the LocationCount field if non-nil, zero value otherwise.

### GetLocationCountOk

`func (o *SourceCodeInfo) GetLocationCountOk() (*int32, bool)`

GetLocationCountOk returns a tuple with the LocationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCount

`func (o *SourceCodeInfo) SetLocationCount(v int32)`

SetLocationCount sets LocationCount field to given value.

### HasLocationCount

`func (o *SourceCodeInfo) HasLocationCount() bool`

HasLocationCount returns a boolean if a field has been set.

### GetLocationList

`func (o *SourceCodeInfo) GetLocationList() []Location`

GetLocationList returns the LocationList field if non-nil, zero value otherwise.

### GetLocationListOk

`func (o *SourceCodeInfo) GetLocationListOk() (*[]Location, bool)`

GetLocationListOk returns a tuple with the LocationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationList

`func (o *SourceCodeInfo) SetLocationList(v []Location)`

SetLocationList sets LocationList field to given value.

### HasLocationList

`func (o *SourceCodeInfo) HasLocationList() bool`

HasLocationList returns a boolean if a field has been set.

### GetLocationOrBuilderList

`func (o *SourceCodeInfo) GetLocationOrBuilderList() []LocationOrBuilder`

GetLocationOrBuilderList returns the LocationOrBuilderList field if non-nil, zero value otherwise.

### GetLocationOrBuilderListOk

`func (o *SourceCodeInfo) GetLocationOrBuilderListOk() (*[]LocationOrBuilder, bool)`

GetLocationOrBuilderListOk returns a tuple with the LocationOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationOrBuilderList

`func (o *SourceCodeInfo) SetLocationOrBuilderList(v []LocationOrBuilder)`

SetLocationOrBuilderList sets LocationOrBuilderList field to given value.

### HasLocationOrBuilderList

`func (o *SourceCodeInfo) HasLocationOrBuilderList() bool`

HasLocationOrBuilderList returns a boolean if a field has been set.

### GetMemoizedSerializedSize

`func (o *SourceCodeInfo) GetMemoizedSerializedSize() int32`

GetMemoizedSerializedSize returns the MemoizedSerializedSize field if non-nil, zero value otherwise.

### GetMemoizedSerializedSizeOk

`func (o *SourceCodeInfo) GetMemoizedSerializedSizeOk() (*int32, bool)`

GetMemoizedSerializedSizeOk returns a tuple with the MemoizedSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizedSerializedSize

`func (o *SourceCodeInfo) SetMemoizedSerializedSize(v int32)`

SetMemoizedSerializedSize sets MemoizedSerializedSize field to given value.

### HasMemoizedSerializedSize

`func (o *SourceCodeInfo) HasMemoizedSerializedSize() bool`

HasMemoizedSerializedSize returns a boolean if a field has been set.

### GetParserForType

`func (o *SourceCodeInfo) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *SourceCodeInfo) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *SourceCodeInfo) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *SourceCodeInfo) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetSerializedSize

`func (o *SourceCodeInfo) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *SourceCodeInfo) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *SourceCodeInfo) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *SourceCodeInfo) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetUnknownFields

`func (o *SourceCodeInfo) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *SourceCodeInfo) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *SourceCodeInfo) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *SourceCodeInfo) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


