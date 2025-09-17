# FeatureSetOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnumType** | Pointer to **string** |  | [optional] 
**RepeatedFieldEncoding** | Pointer to **string** |  | [optional] 
**Utf8Validation** | Pointer to **string** |  | [optional] 
**FieldPresence** | Pointer to **string** |  | [optional] 
**MessageEncoding** | Pointer to **string** |  | [optional] 
**JsonFormat** | Pointer to **string** |  | [optional] 
**DefaultInstanceForType** | Pointer to **map[string]interface{}** | Simplified schema for Message (original had circular references) | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 

## Methods

### NewFeatureSetOrBuilder

`func NewFeatureSetOrBuilder() *FeatureSetOrBuilder`

NewFeatureSetOrBuilder instantiates a new FeatureSetOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureSetOrBuilderWithDefaults

`func NewFeatureSetOrBuilderWithDefaults() *FeatureSetOrBuilder`

NewFeatureSetOrBuilderWithDefaults instantiates a new FeatureSetOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnumType

`func (o *FeatureSetOrBuilder) GetEnumType() string`

GetEnumType returns the EnumType field if non-nil, zero value otherwise.

### GetEnumTypeOk

`func (o *FeatureSetOrBuilder) GetEnumTypeOk() (*string, bool)`

GetEnumTypeOk returns a tuple with the EnumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumType

`func (o *FeatureSetOrBuilder) SetEnumType(v string)`

SetEnumType sets EnumType field to given value.

### HasEnumType

`func (o *FeatureSetOrBuilder) HasEnumType() bool`

HasEnumType returns a boolean if a field has been set.

### GetRepeatedFieldEncoding

`func (o *FeatureSetOrBuilder) GetRepeatedFieldEncoding() string`

GetRepeatedFieldEncoding returns the RepeatedFieldEncoding field if non-nil, zero value otherwise.

### GetRepeatedFieldEncodingOk

`func (o *FeatureSetOrBuilder) GetRepeatedFieldEncodingOk() (*string, bool)`

GetRepeatedFieldEncodingOk returns a tuple with the RepeatedFieldEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatedFieldEncoding

`func (o *FeatureSetOrBuilder) SetRepeatedFieldEncoding(v string)`

SetRepeatedFieldEncoding sets RepeatedFieldEncoding field to given value.

### HasRepeatedFieldEncoding

`func (o *FeatureSetOrBuilder) HasRepeatedFieldEncoding() bool`

HasRepeatedFieldEncoding returns a boolean if a field has been set.

### GetUtf8Validation

`func (o *FeatureSetOrBuilder) GetUtf8Validation() string`

GetUtf8Validation returns the Utf8Validation field if non-nil, zero value otherwise.

### GetUtf8ValidationOk

`func (o *FeatureSetOrBuilder) GetUtf8ValidationOk() (*string, bool)`

GetUtf8ValidationOk returns a tuple with the Utf8Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtf8Validation

`func (o *FeatureSetOrBuilder) SetUtf8Validation(v string)`

SetUtf8Validation sets Utf8Validation field to given value.

### HasUtf8Validation

`func (o *FeatureSetOrBuilder) HasUtf8Validation() bool`

HasUtf8Validation returns a boolean if a field has been set.

### GetFieldPresence

`func (o *FeatureSetOrBuilder) GetFieldPresence() string`

GetFieldPresence returns the FieldPresence field if non-nil, zero value otherwise.

### GetFieldPresenceOk

`func (o *FeatureSetOrBuilder) GetFieldPresenceOk() (*string, bool)`

GetFieldPresenceOk returns a tuple with the FieldPresence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldPresence

`func (o *FeatureSetOrBuilder) SetFieldPresence(v string)`

SetFieldPresence sets FieldPresence field to given value.

### HasFieldPresence

`func (o *FeatureSetOrBuilder) HasFieldPresence() bool`

HasFieldPresence returns a boolean if a field has been set.

### GetMessageEncoding

`func (o *FeatureSetOrBuilder) GetMessageEncoding() string`

GetMessageEncoding returns the MessageEncoding field if non-nil, zero value otherwise.

### GetMessageEncodingOk

`func (o *FeatureSetOrBuilder) GetMessageEncodingOk() (*string, bool)`

GetMessageEncodingOk returns a tuple with the MessageEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEncoding

`func (o *FeatureSetOrBuilder) SetMessageEncoding(v string)`

SetMessageEncoding sets MessageEncoding field to given value.

### HasMessageEncoding

`func (o *FeatureSetOrBuilder) HasMessageEncoding() bool`

HasMessageEncoding returns a boolean if a field has been set.

### GetJsonFormat

`func (o *FeatureSetOrBuilder) GetJsonFormat() string`

GetJsonFormat returns the JsonFormat field if non-nil, zero value otherwise.

### GetJsonFormatOk

`func (o *FeatureSetOrBuilder) GetJsonFormatOk() (*string, bool)`

GetJsonFormatOk returns a tuple with the JsonFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonFormat

`func (o *FeatureSetOrBuilder) SetJsonFormat(v string)`

SetJsonFormat sets JsonFormat field to given value.

### HasJsonFormat

`func (o *FeatureSetOrBuilder) HasJsonFormat() bool`

HasJsonFormat returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *FeatureSetOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *FeatureSetOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *FeatureSetOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *FeatureSetOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *FeatureSetOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *FeatureSetOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *FeatureSetOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *FeatureSetOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *FeatureSetOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *FeatureSetOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *FeatureSetOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *FeatureSetOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetAllFields

`func (o *FeatureSetOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *FeatureSetOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *FeatureSetOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *FeatureSetOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetUnknownFields

`func (o *FeatureSetOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *FeatureSetOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *FeatureSetOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *FeatureSetOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetInitialized

`func (o *FeatureSetOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *FeatureSetOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *FeatureSetOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *FeatureSetOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


