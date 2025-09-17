# FeatureSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to  |  | [optional] 
**AllFieldsRaw** | Pointer to  |  | [optional] 
**DefaultInstanceForType** | Pointer to [**FeatureSet**](FeatureSet.md) |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**EnumType** | Pointer to **string** |  | [optional] 
**FieldPresence** | Pointer to **string** |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**JsonFormat** | Pointer to **string** |  | [optional] 
**MemoizedSerializedSize** | Pointer to **int32** |  | [optional] 
**MessageEncoding** | Pointer to **string** |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** |  | [optional] 
**RepeatedFieldEncoding** | Pointer to **string** |  | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**Utf8Validation** | Pointer to **string** |  | [optional] 

## Methods

### NewFeatureSet

`func NewFeatureSet() *FeatureSet`

NewFeatureSet instantiates a new FeatureSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureSetWithDefaults

`func NewFeatureSetWithDefaults() *FeatureSet`

NewFeatureSetWithDefaults instantiates a new FeatureSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *FeatureSet) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *FeatureSet) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *FeatureSet) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *FeatureSet) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### SetAllFieldsNil

`func (o *FeatureSet) SetAllFieldsNil(b bool)`

 SetAllFieldsNil sets the value for AllFields to be an explicit nil

### UnsetAllFields
`func (o *FeatureSet) UnsetAllFields()`

UnsetAllFields ensures that no value is present for AllFields, not even an explicit nil
### GetAllFieldsRaw

`func (o *FeatureSet) GetAllFieldsRaw() map[string]map[string]interface{}`

GetAllFieldsRaw returns the AllFieldsRaw field if non-nil, zero value otherwise.

### GetAllFieldsRawOk

`func (o *FeatureSet) GetAllFieldsRawOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsRawOk returns a tuple with the AllFieldsRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFieldsRaw

`func (o *FeatureSet) SetAllFieldsRaw(v map[string]map[string]interface{})`

SetAllFieldsRaw sets AllFieldsRaw field to given value.

### HasAllFieldsRaw

`func (o *FeatureSet) HasAllFieldsRaw() bool`

HasAllFieldsRaw returns a boolean if a field has been set.

### SetAllFieldsRawNil

`func (o *FeatureSet) SetAllFieldsRawNil(b bool)`

 SetAllFieldsRawNil sets the value for AllFieldsRaw to be an explicit nil

### UnsetAllFieldsRaw
`func (o *FeatureSet) UnsetAllFieldsRaw()`

UnsetAllFieldsRaw ensures that no value is present for AllFieldsRaw, not even an explicit nil
### GetDefaultInstanceForType

`func (o *FeatureSet) GetDefaultInstanceForType() FeatureSet`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *FeatureSet) GetDefaultInstanceForTypeOk() (*FeatureSet, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *FeatureSet) SetDefaultInstanceForType(v FeatureSet)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *FeatureSet) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *FeatureSet) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *FeatureSet) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *FeatureSet) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *FeatureSet) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetEnumType

`func (o *FeatureSet) GetEnumType() string`

GetEnumType returns the EnumType field if non-nil, zero value otherwise.

### GetEnumTypeOk

`func (o *FeatureSet) GetEnumTypeOk() (*string, bool)`

GetEnumTypeOk returns a tuple with the EnumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumType

`func (o *FeatureSet) SetEnumType(v string)`

SetEnumType sets EnumType field to given value.

### HasEnumType

`func (o *FeatureSet) HasEnumType() bool`

HasEnumType returns a boolean if a field has been set.

### GetFieldPresence

`func (o *FeatureSet) GetFieldPresence() string`

GetFieldPresence returns the FieldPresence field if non-nil, zero value otherwise.

### GetFieldPresenceOk

`func (o *FeatureSet) GetFieldPresenceOk() (*string, bool)`

GetFieldPresenceOk returns a tuple with the FieldPresence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldPresence

`func (o *FeatureSet) SetFieldPresence(v string)`

SetFieldPresence sets FieldPresence field to given value.

### HasFieldPresence

`func (o *FeatureSet) HasFieldPresence() bool`

HasFieldPresence returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *FeatureSet) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *FeatureSet) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *FeatureSet) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *FeatureSet) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *FeatureSet) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *FeatureSet) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *FeatureSet) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *FeatureSet) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetJsonFormat

`func (o *FeatureSet) GetJsonFormat() string`

GetJsonFormat returns the JsonFormat field if non-nil, zero value otherwise.

### GetJsonFormatOk

`func (o *FeatureSet) GetJsonFormatOk() (*string, bool)`

GetJsonFormatOk returns a tuple with the JsonFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonFormat

`func (o *FeatureSet) SetJsonFormat(v string)`

SetJsonFormat sets JsonFormat field to given value.

### HasJsonFormat

`func (o *FeatureSet) HasJsonFormat() bool`

HasJsonFormat returns a boolean if a field has been set.

### GetMemoizedSerializedSize

`func (o *FeatureSet) GetMemoizedSerializedSize() int32`

GetMemoizedSerializedSize returns the MemoizedSerializedSize field if non-nil, zero value otherwise.

### GetMemoizedSerializedSizeOk

`func (o *FeatureSet) GetMemoizedSerializedSizeOk() (*int32, bool)`

GetMemoizedSerializedSizeOk returns a tuple with the MemoizedSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizedSerializedSize

`func (o *FeatureSet) SetMemoizedSerializedSize(v int32)`

SetMemoizedSerializedSize sets MemoizedSerializedSize field to given value.

### HasMemoizedSerializedSize

`func (o *FeatureSet) HasMemoizedSerializedSize() bool`

HasMemoizedSerializedSize returns a boolean if a field has been set.

### GetMessageEncoding

`func (o *FeatureSet) GetMessageEncoding() string`

GetMessageEncoding returns the MessageEncoding field if non-nil, zero value otherwise.

### GetMessageEncodingOk

`func (o *FeatureSet) GetMessageEncodingOk() (*string, bool)`

GetMessageEncodingOk returns a tuple with the MessageEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEncoding

`func (o *FeatureSet) SetMessageEncoding(v string)`

SetMessageEncoding sets MessageEncoding field to given value.

### HasMessageEncoding

`func (o *FeatureSet) HasMessageEncoding() bool`

HasMessageEncoding returns a boolean if a field has been set.

### GetParserForType

`func (o *FeatureSet) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *FeatureSet) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *FeatureSet) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *FeatureSet) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetRepeatedFieldEncoding

`func (o *FeatureSet) GetRepeatedFieldEncoding() string`

GetRepeatedFieldEncoding returns the RepeatedFieldEncoding field if non-nil, zero value otherwise.

### GetRepeatedFieldEncodingOk

`func (o *FeatureSet) GetRepeatedFieldEncodingOk() (*string, bool)`

GetRepeatedFieldEncodingOk returns a tuple with the RepeatedFieldEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatedFieldEncoding

`func (o *FeatureSet) SetRepeatedFieldEncoding(v string)`

SetRepeatedFieldEncoding sets RepeatedFieldEncoding field to given value.

### HasRepeatedFieldEncoding

`func (o *FeatureSet) HasRepeatedFieldEncoding() bool`

HasRepeatedFieldEncoding returns a boolean if a field has been set.

### GetSerializedSize

`func (o *FeatureSet) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *FeatureSet) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *FeatureSet) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *FeatureSet) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetUnknownFields

`func (o *FeatureSet) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *FeatureSet) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *FeatureSet) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *FeatureSet) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetUtf8Validation

`func (o *FeatureSet) GetUtf8Validation() string`

GetUtf8Validation returns the Utf8Validation field if non-nil, zero value otherwise.

### GetUtf8ValidationOk

`func (o *FeatureSet) GetUtf8ValidationOk() (*string, bool)`

GetUtf8ValidationOk returns a tuple with the Utf8Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtf8Validation

`func (o *FeatureSet) SetUtf8Validation(v string)`

SetUtf8Validation sets Utf8Validation field to given value.

### HasUtf8Validation

`func (o *FeatureSet) HasUtf8Validation() bool`

HasUtf8Validation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


