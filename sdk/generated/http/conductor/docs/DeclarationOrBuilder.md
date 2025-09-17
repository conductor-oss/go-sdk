# DeclarationOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**Repeated** | Pointer to **bool** |  | [optional] 
**FullNameBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**TypeBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**Reserved** | Pointer to **bool** |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**DefaultInstanceForType** | Pointer to **map[string]interface{}** | Simplified schema for Message (original had circular references) | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 

## Methods

### NewDeclarationOrBuilder

`func NewDeclarationOrBuilder() *DeclarationOrBuilder`

NewDeclarationOrBuilder instantiates a new DeclarationOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclarationOrBuilderWithDefaults

`func NewDeclarationOrBuilderWithDefaults() *DeclarationOrBuilder`

NewDeclarationOrBuilderWithDefaults instantiates a new DeclarationOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeclarationOrBuilder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeclarationOrBuilder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeclarationOrBuilder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeclarationOrBuilder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNumber

`func (o *DeclarationOrBuilder) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *DeclarationOrBuilder) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *DeclarationOrBuilder) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *DeclarationOrBuilder) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetFullName

`func (o *DeclarationOrBuilder) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *DeclarationOrBuilder) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *DeclarationOrBuilder) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *DeclarationOrBuilder) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetRepeated

`func (o *DeclarationOrBuilder) GetRepeated() bool`

GetRepeated returns the Repeated field if non-nil, zero value otherwise.

### GetRepeatedOk

`func (o *DeclarationOrBuilder) GetRepeatedOk() (*bool, bool)`

GetRepeatedOk returns a tuple with the Repeated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeated

`func (o *DeclarationOrBuilder) SetRepeated(v bool)`

SetRepeated sets Repeated field to given value.

### HasRepeated

`func (o *DeclarationOrBuilder) HasRepeated() bool`

HasRepeated returns a boolean if a field has been set.

### GetFullNameBytes

`func (o *DeclarationOrBuilder) GetFullNameBytes() ByteString`

GetFullNameBytes returns the FullNameBytes field if non-nil, zero value otherwise.

### GetFullNameBytesOk

`func (o *DeclarationOrBuilder) GetFullNameBytesOk() (*ByteString, bool)`

GetFullNameBytesOk returns a tuple with the FullNameBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullNameBytes

`func (o *DeclarationOrBuilder) SetFullNameBytes(v ByteString)`

SetFullNameBytes sets FullNameBytes field to given value.

### HasFullNameBytes

`func (o *DeclarationOrBuilder) HasFullNameBytes() bool`

HasFullNameBytes returns a boolean if a field has been set.

### GetTypeBytes

`func (o *DeclarationOrBuilder) GetTypeBytes() ByteString`

GetTypeBytes returns the TypeBytes field if non-nil, zero value otherwise.

### GetTypeBytesOk

`func (o *DeclarationOrBuilder) GetTypeBytesOk() (*ByteString, bool)`

GetTypeBytesOk returns a tuple with the TypeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeBytes

`func (o *DeclarationOrBuilder) SetTypeBytes(v ByteString)`

SetTypeBytes sets TypeBytes field to given value.

### HasTypeBytes

`func (o *DeclarationOrBuilder) HasTypeBytes() bool`

HasTypeBytes returns a boolean if a field has been set.

### GetReserved

`func (o *DeclarationOrBuilder) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *DeclarationOrBuilder) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *DeclarationOrBuilder) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *DeclarationOrBuilder) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *DeclarationOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *DeclarationOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *DeclarationOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *DeclarationOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *DeclarationOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *DeclarationOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *DeclarationOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *DeclarationOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetAllFields

`func (o *DeclarationOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *DeclarationOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *DeclarationOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *DeclarationOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *DeclarationOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *DeclarationOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *DeclarationOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *DeclarationOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetUnknownFields

`func (o *DeclarationOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *DeclarationOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *DeclarationOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *DeclarationOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetInitialized

`func (o *DeclarationOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *DeclarationOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *DeclarationOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *DeclarationOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


