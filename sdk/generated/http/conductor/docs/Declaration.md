# Declaration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**Repeated** | Pointer to **bool** |  | [optional] 
**DefaultInstanceForType** | Pointer to [**Declaration**](Declaration.md) |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** | Simplified schema for ParserDeclaration (original had circular references) | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**FullNameBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**TypeBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**Reserved** | Pointer to **bool** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**MemoizedSerializedSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewDeclaration

`func NewDeclaration() *Declaration`

NewDeclaration instantiates a new Declaration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclarationWithDefaults

`func NewDeclarationWithDefaults() *Declaration`

NewDeclarationWithDefaults instantiates a new Declaration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnknownFields

`func (o *Declaration) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *Declaration) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *Declaration) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *Declaration) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetType

`func (o *Declaration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Declaration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Declaration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Declaration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNumber

`func (o *Declaration) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Declaration) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Declaration) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Declaration) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetFullName

`func (o *Declaration) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Declaration) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Declaration) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *Declaration) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetRepeated

`func (o *Declaration) GetRepeated() bool`

GetRepeated returns the Repeated field if non-nil, zero value otherwise.

### GetRepeatedOk

`func (o *Declaration) GetRepeatedOk() (*bool, bool)`

GetRepeatedOk returns a tuple with the Repeated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeated

`func (o *Declaration) SetRepeated(v bool)`

SetRepeated sets Repeated field to given value.

### HasRepeated

`func (o *Declaration) HasRepeated() bool`

HasRepeated returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *Declaration) GetDefaultInstanceForType() Declaration`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *Declaration) GetDefaultInstanceForTypeOk() (*Declaration, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *Declaration) SetDefaultInstanceForType(v Declaration)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *Declaration) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetParserForType

`func (o *Declaration) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *Declaration) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *Declaration) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *Declaration) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetSerializedSize

`func (o *Declaration) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *Declaration) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *Declaration) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *Declaration) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetFullNameBytes

`func (o *Declaration) GetFullNameBytes() ByteString`

GetFullNameBytes returns the FullNameBytes field if non-nil, zero value otherwise.

### GetFullNameBytesOk

`func (o *Declaration) GetFullNameBytesOk() (*ByteString, bool)`

GetFullNameBytesOk returns a tuple with the FullNameBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullNameBytes

`func (o *Declaration) SetFullNameBytes(v ByteString)`

SetFullNameBytes sets FullNameBytes field to given value.

### HasFullNameBytes

`func (o *Declaration) HasFullNameBytes() bool`

HasFullNameBytes returns a boolean if a field has been set.

### GetTypeBytes

`func (o *Declaration) GetTypeBytes() ByteString`

GetTypeBytes returns the TypeBytes field if non-nil, zero value otherwise.

### GetTypeBytesOk

`func (o *Declaration) GetTypeBytesOk() (*ByteString, bool)`

GetTypeBytesOk returns a tuple with the TypeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeBytes

`func (o *Declaration) SetTypeBytes(v ByteString)`

SetTypeBytes sets TypeBytes field to given value.

### HasTypeBytes

`func (o *Declaration) HasTypeBytes() bool`

HasTypeBytes returns a boolean if a field has been set.

### GetReserved

`func (o *Declaration) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *Declaration) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *Declaration) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *Declaration) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetInitialized

`func (o *Declaration) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *Declaration) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *Declaration) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *Declaration) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *Declaration) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *Declaration) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *Declaration) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *Declaration) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *Declaration) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *Declaration) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *Declaration) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *Declaration) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetAllFields

`func (o *Declaration) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *Declaration) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *Declaration) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *Declaration) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetMemoizedSerializedSize

`func (o *Declaration) GetMemoizedSerializedSize() int32`

GetMemoizedSerializedSize returns the MemoizedSerializedSize field if non-nil, zero value otherwise.

### GetMemoizedSerializedSizeOk

`func (o *Declaration) GetMemoizedSerializedSizeOk() (*int32, bool)`

GetMemoizedSerializedSizeOk returns a tuple with the MemoizedSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizedSerializedSize

`func (o *Declaration) SetMemoizedSerializedSize(v int32)`

SetMemoizedSerializedSize sets MemoizedSerializedSize field to given value.

### HasMemoizedSerializedSize

`func (o *Declaration) HasMemoizedSerializedSize() bool`

HasMemoizedSerializedSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


