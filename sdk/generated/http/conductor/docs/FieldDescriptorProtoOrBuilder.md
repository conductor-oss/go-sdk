# FieldDescriptorProtoOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**TypeName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**JsonName** | Pointer to **string** |  | [optional] 
**Proto3Optional** | Pointer to **bool** |  | [optional] 
**OneofIndex** | Pointer to **int32** |  | [optional] 
**Extendee** | Pointer to **string** |  | [optional] 
**NameBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**Options** | Pointer to [**FieldOptions**](FieldOptions.md) |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**OptionsOrBuilder** | Pointer to [**FieldOptionsOrBuilder**](FieldOptionsOrBuilder.md) |  | [optional] 
**DefaultValueBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**TypeNameBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**ExtendeeBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**JsonNameBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**DefaultInstanceForType** | Pointer to **map[string]interface{}** | Simplified schema for Message (original had circular references) | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 

## Methods

### NewFieldDescriptorProtoOrBuilder

`func NewFieldDescriptorProtoOrBuilder() *FieldDescriptorProtoOrBuilder`

NewFieldDescriptorProtoOrBuilder instantiates a new FieldDescriptorProtoOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldDescriptorProtoOrBuilderWithDefaults

`func NewFieldDescriptorProtoOrBuilderWithDefaults() *FieldDescriptorProtoOrBuilder`

NewFieldDescriptorProtoOrBuilderWithDefaults instantiates a new FieldDescriptorProtoOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FieldDescriptorProtoOrBuilder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldDescriptorProtoOrBuilder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldDescriptorProtoOrBuilder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FieldDescriptorProtoOrBuilder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypeName

`func (o *FieldDescriptorProtoOrBuilder) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *FieldDescriptorProtoOrBuilder) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *FieldDescriptorProtoOrBuilder) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *FieldDescriptorProtoOrBuilder) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetType

`func (o *FieldDescriptorProtoOrBuilder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldDescriptorProtoOrBuilder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldDescriptorProtoOrBuilder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FieldDescriptorProtoOrBuilder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultValue

`func (o *FieldDescriptorProtoOrBuilder) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *FieldDescriptorProtoOrBuilder) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *FieldDescriptorProtoOrBuilder) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *FieldDescriptorProtoOrBuilder) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetNumber

`func (o *FieldDescriptorProtoOrBuilder) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *FieldDescriptorProtoOrBuilder) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *FieldDescriptorProtoOrBuilder) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *FieldDescriptorProtoOrBuilder) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetJsonName

`func (o *FieldDescriptorProtoOrBuilder) GetJsonName() string`

GetJsonName returns the JsonName field if non-nil, zero value otherwise.

### GetJsonNameOk

`func (o *FieldDescriptorProtoOrBuilder) GetJsonNameOk() (*string, bool)`

GetJsonNameOk returns a tuple with the JsonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonName

`func (o *FieldDescriptorProtoOrBuilder) SetJsonName(v string)`

SetJsonName sets JsonName field to given value.

### HasJsonName

`func (o *FieldDescriptorProtoOrBuilder) HasJsonName() bool`

HasJsonName returns a boolean if a field has been set.

### GetProto3Optional

`func (o *FieldDescriptorProtoOrBuilder) GetProto3Optional() bool`

GetProto3Optional returns the Proto3Optional field if non-nil, zero value otherwise.

### GetProto3OptionalOk

`func (o *FieldDescriptorProtoOrBuilder) GetProto3OptionalOk() (*bool, bool)`

GetProto3OptionalOk returns a tuple with the Proto3Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto3Optional

`func (o *FieldDescriptorProtoOrBuilder) SetProto3Optional(v bool)`

SetProto3Optional sets Proto3Optional field to given value.

### HasProto3Optional

`func (o *FieldDescriptorProtoOrBuilder) HasProto3Optional() bool`

HasProto3Optional returns a boolean if a field has been set.

### GetOneofIndex

`func (o *FieldDescriptorProtoOrBuilder) GetOneofIndex() int32`

GetOneofIndex returns the OneofIndex field if non-nil, zero value otherwise.

### GetOneofIndexOk

`func (o *FieldDescriptorProtoOrBuilder) GetOneofIndexOk() (*int32, bool)`

GetOneofIndexOk returns a tuple with the OneofIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneofIndex

`func (o *FieldDescriptorProtoOrBuilder) SetOneofIndex(v int32)`

SetOneofIndex sets OneofIndex field to given value.

### HasOneofIndex

`func (o *FieldDescriptorProtoOrBuilder) HasOneofIndex() bool`

HasOneofIndex returns a boolean if a field has been set.

### GetExtendee

`func (o *FieldDescriptorProtoOrBuilder) GetExtendee() string`

GetExtendee returns the Extendee field if non-nil, zero value otherwise.

### GetExtendeeOk

`func (o *FieldDescriptorProtoOrBuilder) GetExtendeeOk() (*string, bool)`

GetExtendeeOk returns a tuple with the Extendee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendee

`func (o *FieldDescriptorProtoOrBuilder) SetExtendee(v string)`

SetExtendee sets Extendee field to given value.

### HasExtendee

`func (o *FieldDescriptorProtoOrBuilder) HasExtendee() bool`

HasExtendee returns a boolean if a field has been set.

### GetNameBytes

`func (o *FieldDescriptorProtoOrBuilder) GetNameBytes() ByteString`

GetNameBytes returns the NameBytes field if non-nil, zero value otherwise.

### GetNameBytesOk

`func (o *FieldDescriptorProtoOrBuilder) GetNameBytesOk() (*ByteString, bool)`

GetNameBytesOk returns a tuple with the NameBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameBytes

`func (o *FieldDescriptorProtoOrBuilder) SetNameBytes(v ByteString)`

SetNameBytes sets NameBytes field to given value.

### HasNameBytes

`func (o *FieldDescriptorProtoOrBuilder) HasNameBytes() bool`

HasNameBytes returns a boolean if a field has been set.

### GetOptions

`func (o *FieldDescriptorProtoOrBuilder) GetOptions() FieldOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FieldDescriptorProtoOrBuilder) GetOptionsOk() (*FieldOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FieldDescriptorProtoOrBuilder) SetOptions(v FieldOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FieldDescriptorProtoOrBuilder) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetLabel

`func (o *FieldDescriptorProtoOrBuilder) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FieldDescriptorProtoOrBuilder) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FieldDescriptorProtoOrBuilder) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FieldDescriptorProtoOrBuilder) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOptionsOrBuilder

`func (o *FieldDescriptorProtoOrBuilder) GetOptionsOrBuilder() FieldOptionsOrBuilder`

GetOptionsOrBuilder returns the OptionsOrBuilder field if non-nil, zero value otherwise.

### GetOptionsOrBuilderOk

`func (o *FieldDescriptorProtoOrBuilder) GetOptionsOrBuilderOk() (*FieldOptionsOrBuilder, bool)`

GetOptionsOrBuilderOk returns a tuple with the OptionsOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsOrBuilder

`func (o *FieldDescriptorProtoOrBuilder) SetOptionsOrBuilder(v FieldOptionsOrBuilder)`

SetOptionsOrBuilder sets OptionsOrBuilder field to given value.

### HasOptionsOrBuilder

`func (o *FieldDescriptorProtoOrBuilder) HasOptionsOrBuilder() bool`

HasOptionsOrBuilder returns a boolean if a field has been set.

### GetDefaultValueBytes

`func (o *FieldDescriptorProtoOrBuilder) GetDefaultValueBytes() ByteString`

GetDefaultValueBytes returns the DefaultValueBytes field if non-nil, zero value otherwise.

### GetDefaultValueBytesOk

`func (o *FieldDescriptorProtoOrBuilder) GetDefaultValueBytesOk() (*ByteString, bool)`

GetDefaultValueBytesOk returns a tuple with the DefaultValueBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValueBytes

`func (o *FieldDescriptorProtoOrBuilder) SetDefaultValueBytes(v ByteString)`

SetDefaultValueBytes sets DefaultValueBytes field to given value.

### HasDefaultValueBytes

`func (o *FieldDescriptorProtoOrBuilder) HasDefaultValueBytes() bool`

HasDefaultValueBytes returns a boolean if a field has been set.

### GetTypeNameBytes

`func (o *FieldDescriptorProtoOrBuilder) GetTypeNameBytes() ByteString`

GetTypeNameBytes returns the TypeNameBytes field if non-nil, zero value otherwise.

### GetTypeNameBytesOk

`func (o *FieldDescriptorProtoOrBuilder) GetTypeNameBytesOk() (*ByteString, bool)`

GetTypeNameBytesOk returns a tuple with the TypeNameBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeNameBytes

`func (o *FieldDescriptorProtoOrBuilder) SetTypeNameBytes(v ByteString)`

SetTypeNameBytes sets TypeNameBytes field to given value.

### HasTypeNameBytes

`func (o *FieldDescriptorProtoOrBuilder) HasTypeNameBytes() bool`

HasTypeNameBytes returns a boolean if a field has been set.

### GetExtendeeBytes

`func (o *FieldDescriptorProtoOrBuilder) GetExtendeeBytes() ByteString`

GetExtendeeBytes returns the ExtendeeBytes field if non-nil, zero value otherwise.

### GetExtendeeBytesOk

`func (o *FieldDescriptorProtoOrBuilder) GetExtendeeBytesOk() (*ByteString, bool)`

GetExtendeeBytesOk returns a tuple with the ExtendeeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendeeBytes

`func (o *FieldDescriptorProtoOrBuilder) SetExtendeeBytes(v ByteString)`

SetExtendeeBytes sets ExtendeeBytes field to given value.

### HasExtendeeBytes

`func (o *FieldDescriptorProtoOrBuilder) HasExtendeeBytes() bool`

HasExtendeeBytes returns a boolean if a field has been set.

### GetJsonNameBytes

`func (o *FieldDescriptorProtoOrBuilder) GetJsonNameBytes() ByteString`

GetJsonNameBytes returns the JsonNameBytes field if non-nil, zero value otherwise.

### GetJsonNameBytesOk

`func (o *FieldDescriptorProtoOrBuilder) GetJsonNameBytesOk() (*ByteString, bool)`

GetJsonNameBytesOk returns a tuple with the JsonNameBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonNameBytes

`func (o *FieldDescriptorProtoOrBuilder) SetJsonNameBytes(v ByteString)`

SetJsonNameBytes sets JsonNameBytes field to given value.

### HasJsonNameBytes

`func (o *FieldDescriptorProtoOrBuilder) HasJsonNameBytes() bool`

HasJsonNameBytes returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *FieldDescriptorProtoOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *FieldDescriptorProtoOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *FieldDescriptorProtoOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *FieldDescriptorProtoOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *FieldDescriptorProtoOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *FieldDescriptorProtoOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *FieldDescriptorProtoOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *FieldDescriptorProtoOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetAllFields

`func (o *FieldDescriptorProtoOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *FieldDescriptorProtoOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *FieldDescriptorProtoOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *FieldDescriptorProtoOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *FieldDescriptorProtoOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *FieldDescriptorProtoOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *FieldDescriptorProtoOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *FieldDescriptorProtoOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetUnknownFields

`func (o *FieldDescriptorProtoOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *FieldDescriptorProtoOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *FieldDescriptorProtoOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *FieldDescriptorProtoOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetInitialized

`func (o *FieldDescriptorProtoOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *FieldDescriptorProtoOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *FieldDescriptorProtoOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *FieldDescriptorProtoOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


