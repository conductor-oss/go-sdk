# SourceCodeInfoOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationList** | Pointer to [**[]Location**](Location.md) |  | [optional] 
**LocationCount** | Pointer to **int32** |  | [optional] 
**LocationOrBuilderList** | Pointer to [**[]LocationOrBuilder**](LocationOrBuilder.md) |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**DefaultInstanceForType** | Pointer to **map[string]interface{}** | Simplified schema for Message (original had circular references) | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 

## Methods

### NewSourceCodeInfoOrBuilder

`func NewSourceCodeInfoOrBuilder() *SourceCodeInfoOrBuilder`

NewSourceCodeInfoOrBuilder instantiates a new SourceCodeInfoOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceCodeInfoOrBuilderWithDefaults

`func NewSourceCodeInfoOrBuilderWithDefaults() *SourceCodeInfoOrBuilder`

NewSourceCodeInfoOrBuilderWithDefaults instantiates a new SourceCodeInfoOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationList

`func (o *SourceCodeInfoOrBuilder) GetLocationList() []Location`

GetLocationList returns the LocationList field if non-nil, zero value otherwise.

### GetLocationListOk

`func (o *SourceCodeInfoOrBuilder) GetLocationListOk() (*[]Location, bool)`

GetLocationListOk returns a tuple with the LocationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationList

`func (o *SourceCodeInfoOrBuilder) SetLocationList(v []Location)`

SetLocationList sets LocationList field to given value.

### HasLocationList

`func (o *SourceCodeInfoOrBuilder) HasLocationList() bool`

HasLocationList returns a boolean if a field has been set.

### GetLocationCount

`func (o *SourceCodeInfoOrBuilder) GetLocationCount() int32`

GetLocationCount returns the LocationCount field if non-nil, zero value otherwise.

### GetLocationCountOk

`func (o *SourceCodeInfoOrBuilder) GetLocationCountOk() (*int32, bool)`

GetLocationCountOk returns a tuple with the LocationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCount

`func (o *SourceCodeInfoOrBuilder) SetLocationCount(v int32)`

SetLocationCount sets LocationCount field to given value.

### HasLocationCount

`func (o *SourceCodeInfoOrBuilder) HasLocationCount() bool`

HasLocationCount returns a boolean if a field has been set.

### GetLocationOrBuilderList

`func (o *SourceCodeInfoOrBuilder) GetLocationOrBuilderList() []LocationOrBuilder`

GetLocationOrBuilderList returns the LocationOrBuilderList field if non-nil, zero value otherwise.

### GetLocationOrBuilderListOk

`func (o *SourceCodeInfoOrBuilder) GetLocationOrBuilderListOk() (*[]LocationOrBuilder, bool)`

GetLocationOrBuilderListOk returns a tuple with the LocationOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationOrBuilderList

`func (o *SourceCodeInfoOrBuilder) SetLocationOrBuilderList(v []LocationOrBuilder)`

SetLocationOrBuilderList sets LocationOrBuilderList field to given value.

### HasLocationOrBuilderList

`func (o *SourceCodeInfoOrBuilder) HasLocationOrBuilderList() bool`

HasLocationOrBuilderList returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *SourceCodeInfoOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *SourceCodeInfoOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *SourceCodeInfoOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *SourceCodeInfoOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *SourceCodeInfoOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *SourceCodeInfoOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *SourceCodeInfoOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *SourceCodeInfoOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetAllFields

`func (o *SourceCodeInfoOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *SourceCodeInfoOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *SourceCodeInfoOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *SourceCodeInfoOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *SourceCodeInfoOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *SourceCodeInfoOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *SourceCodeInfoOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *SourceCodeInfoOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetUnknownFields

`func (o *SourceCodeInfoOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *SourceCodeInfoOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *SourceCodeInfoOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *SourceCodeInfoOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetInitialized

`func (o *SourceCodeInfoOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *SourceCodeInfoOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *SourceCodeInfoOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *SourceCodeInfoOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


