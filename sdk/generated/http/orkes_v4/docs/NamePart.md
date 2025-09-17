# NamePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to  |  | [optional] 
**DefaultInstanceForType** | Pointer to [**NamePart**](NamePart.md) |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**IsExtension** | Pointer to **bool** |  | [optional] 
**MemoizedSerializedSize** | Pointer to **int32** |  | [optional] 
**NamePart** | Pointer to **string** |  | [optional] 
**NamePartBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** |  | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewNamePart

`func NewNamePart() *NamePart`

NewNamePart instantiates a new NamePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamePartWithDefaults

`func NewNamePartWithDefaults() *NamePart`

NewNamePartWithDefaults instantiates a new NamePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *NamePart) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *NamePart) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *NamePart) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *NamePart) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### SetAllFieldsNil

`func (o *NamePart) SetAllFieldsNil(b bool)`

 SetAllFieldsNil sets the value for AllFields to be an explicit nil

### UnsetAllFields
`func (o *NamePart) UnsetAllFields()`

UnsetAllFields ensures that no value is present for AllFields, not even an explicit nil
### GetDefaultInstanceForType

`func (o *NamePart) GetDefaultInstanceForType() NamePart`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *NamePart) GetDefaultInstanceForTypeOk() (*NamePart, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *NamePart) SetDefaultInstanceForType(v NamePart)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *NamePart) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *NamePart) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *NamePart) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *NamePart) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *NamePart) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *NamePart) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *NamePart) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *NamePart) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *NamePart) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *NamePart) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *NamePart) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *NamePart) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *NamePart) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetIsExtension

`func (o *NamePart) GetIsExtension() bool`

GetIsExtension returns the IsExtension field if non-nil, zero value otherwise.

### GetIsExtensionOk

`func (o *NamePart) GetIsExtensionOk() (*bool, bool)`

GetIsExtensionOk returns a tuple with the IsExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExtension

`func (o *NamePart) SetIsExtension(v bool)`

SetIsExtension sets IsExtension field to given value.

### HasIsExtension

`func (o *NamePart) HasIsExtension() bool`

HasIsExtension returns a boolean if a field has been set.

### GetMemoizedSerializedSize

`func (o *NamePart) GetMemoizedSerializedSize() int32`

GetMemoizedSerializedSize returns the MemoizedSerializedSize field if non-nil, zero value otherwise.

### GetMemoizedSerializedSizeOk

`func (o *NamePart) GetMemoizedSerializedSizeOk() (*int32, bool)`

GetMemoizedSerializedSizeOk returns a tuple with the MemoizedSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizedSerializedSize

`func (o *NamePart) SetMemoizedSerializedSize(v int32)`

SetMemoizedSerializedSize sets MemoizedSerializedSize field to given value.

### HasMemoizedSerializedSize

`func (o *NamePart) HasMemoizedSerializedSize() bool`

HasMemoizedSerializedSize returns a boolean if a field has been set.

### GetNamePart

`func (o *NamePart) GetNamePart() string`

GetNamePart returns the NamePart field if non-nil, zero value otherwise.

### GetNamePartOk

`func (o *NamePart) GetNamePartOk() (*string, bool)`

GetNamePartOk returns a tuple with the NamePart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePart

`func (o *NamePart) SetNamePart(v string)`

SetNamePart sets NamePart field to given value.

### HasNamePart

`func (o *NamePart) HasNamePart() bool`

HasNamePart returns a boolean if a field has been set.

### GetNamePartBytes

`func (o *NamePart) GetNamePartBytes() ByteString`

GetNamePartBytes returns the NamePartBytes field if non-nil, zero value otherwise.

### GetNamePartBytesOk

`func (o *NamePart) GetNamePartBytesOk() (*ByteString, bool)`

GetNamePartBytesOk returns a tuple with the NamePartBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePartBytes

`func (o *NamePart) SetNamePartBytes(v ByteString)`

SetNamePartBytes sets NamePartBytes field to given value.

### HasNamePartBytes

`func (o *NamePart) HasNamePartBytes() bool`

HasNamePartBytes returns a boolean if a field has been set.

### GetParserForType

`func (o *NamePart) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *NamePart) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *NamePart) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *NamePart) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetSerializedSize

`func (o *NamePart) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *NamePart) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *NamePart) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *NamePart) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetUnknownFields

`func (o *NamePart) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *NamePart) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *NamePart) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *NamePart) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


