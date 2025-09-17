# MethodDescriptorProtoOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**NameBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**InputType** | Pointer to **string** |  | [optional] 
**OutputType** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**MethodOptions**](MethodOptions.md) |  | [optional] 
**InputTypeBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**OutputTypeBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**ClientStreaming** | Pointer to **bool** |  | [optional] 
**ServerStreaming** | Pointer to **bool** |  | [optional] 
**OptionsOrBuilder** | Pointer to [**MethodOptionsOrBuilder**](MethodOptionsOrBuilder.md) |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**DefaultInstanceForType** | Pointer to **map[string]interface{}** | Simplified schema for Message (original had circular references) | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 

## Methods

### NewMethodDescriptorProtoOrBuilder

`func NewMethodDescriptorProtoOrBuilder() *MethodDescriptorProtoOrBuilder`

NewMethodDescriptorProtoOrBuilder instantiates a new MethodDescriptorProtoOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMethodDescriptorProtoOrBuilderWithDefaults

`func NewMethodDescriptorProtoOrBuilderWithDefaults() *MethodDescriptorProtoOrBuilder`

NewMethodDescriptorProtoOrBuilderWithDefaults instantiates a new MethodDescriptorProtoOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MethodDescriptorProtoOrBuilder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MethodDescriptorProtoOrBuilder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MethodDescriptorProtoOrBuilder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MethodDescriptorProtoOrBuilder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameBytes

`func (o *MethodDescriptorProtoOrBuilder) GetNameBytes() ByteString`

GetNameBytes returns the NameBytes field if non-nil, zero value otherwise.

### GetNameBytesOk

`func (o *MethodDescriptorProtoOrBuilder) GetNameBytesOk() (*ByteString, bool)`

GetNameBytesOk returns a tuple with the NameBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameBytes

`func (o *MethodDescriptorProtoOrBuilder) SetNameBytes(v ByteString)`

SetNameBytes sets NameBytes field to given value.

### HasNameBytes

`func (o *MethodDescriptorProtoOrBuilder) HasNameBytes() bool`

HasNameBytes returns a boolean if a field has been set.

### GetInputType

`func (o *MethodDescriptorProtoOrBuilder) GetInputType() string`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *MethodDescriptorProtoOrBuilder) GetInputTypeOk() (*string, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *MethodDescriptorProtoOrBuilder) SetInputType(v string)`

SetInputType sets InputType field to given value.

### HasInputType

`func (o *MethodDescriptorProtoOrBuilder) HasInputType() bool`

HasInputType returns a boolean if a field has been set.

### GetOutputType

`func (o *MethodDescriptorProtoOrBuilder) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *MethodDescriptorProtoOrBuilder) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *MethodDescriptorProtoOrBuilder) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.

### HasOutputType

`func (o *MethodDescriptorProtoOrBuilder) HasOutputType() bool`

HasOutputType returns a boolean if a field has been set.

### GetOptions

`func (o *MethodDescriptorProtoOrBuilder) GetOptions() MethodOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MethodDescriptorProtoOrBuilder) GetOptionsOk() (*MethodOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MethodDescriptorProtoOrBuilder) SetOptions(v MethodOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MethodDescriptorProtoOrBuilder) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetInputTypeBytes

`func (o *MethodDescriptorProtoOrBuilder) GetInputTypeBytes() ByteString`

GetInputTypeBytes returns the InputTypeBytes field if non-nil, zero value otherwise.

### GetInputTypeBytesOk

`func (o *MethodDescriptorProtoOrBuilder) GetInputTypeBytesOk() (*ByteString, bool)`

GetInputTypeBytesOk returns a tuple with the InputTypeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTypeBytes

`func (o *MethodDescriptorProtoOrBuilder) SetInputTypeBytes(v ByteString)`

SetInputTypeBytes sets InputTypeBytes field to given value.

### HasInputTypeBytes

`func (o *MethodDescriptorProtoOrBuilder) HasInputTypeBytes() bool`

HasInputTypeBytes returns a boolean if a field has been set.

### GetOutputTypeBytes

`func (o *MethodDescriptorProtoOrBuilder) GetOutputTypeBytes() ByteString`

GetOutputTypeBytes returns the OutputTypeBytes field if non-nil, zero value otherwise.

### GetOutputTypeBytesOk

`func (o *MethodDescriptorProtoOrBuilder) GetOutputTypeBytesOk() (*ByteString, bool)`

GetOutputTypeBytesOk returns a tuple with the OutputTypeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTypeBytes

`func (o *MethodDescriptorProtoOrBuilder) SetOutputTypeBytes(v ByteString)`

SetOutputTypeBytes sets OutputTypeBytes field to given value.

### HasOutputTypeBytes

`func (o *MethodDescriptorProtoOrBuilder) HasOutputTypeBytes() bool`

HasOutputTypeBytes returns a boolean if a field has been set.

### GetClientStreaming

`func (o *MethodDescriptorProtoOrBuilder) GetClientStreaming() bool`

GetClientStreaming returns the ClientStreaming field if non-nil, zero value otherwise.

### GetClientStreamingOk

`func (o *MethodDescriptorProtoOrBuilder) GetClientStreamingOk() (*bool, bool)`

GetClientStreamingOk returns a tuple with the ClientStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientStreaming

`func (o *MethodDescriptorProtoOrBuilder) SetClientStreaming(v bool)`

SetClientStreaming sets ClientStreaming field to given value.

### HasClientStreaming

`func (o *MethodDescriptorProtoOrBuilder) HasClientStreaming() bool`

HasClientStreaming returns a boolean if a field has been set.

### GetServerStreaming

`func (o *MethodDescriptorProtoOrBuilder) GetServerStreaming() bool`

GetServerStreaming returns the ServerStreaming field if non-nil, zero value otherwise.

### GetServerStreamingOk

`func (o *MethodDescriptorProtoOrBuilder) GetServerStreamingOk() (*bool, bool)`

GetServerStreamingOk returns a tuple with the ServerStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStreaming

`func (o *MethodDescriptorProtoOrBuilder) SetServerStreaming(v bool)`

SetServerStreaming sets ServerStreaming field to given value.

### HasServerStreaming

`func (o *MethodDescriptorProtoOrBuilder) HasServerStreaming() bool`

HasServerStreaming returns a boolean if a field has been set.

### GetOptionsOrBuilder

`func (o *MethodDescriptorProtoOrBuilder) GetOptionsOrBuilder() MethodOptionsOrBuilder`

GetOptionsOrBuilder returns the OptionsOrBuilder field if non-nil, zero value otherwise.

### GetOptionsOrBuilderOk

`func (o *MethodDescriptorProtoOrBuilder) GetOptionsOrBuilderOk() (*MethodOptionsOrBuilder, bool)`

GetOptionsOrBuilderOk returns a tuple with the OptionsOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsOrBuilder

`func (o *MethodDescriptorProtoOrBuilder) SetOptionsOrBuilder(v MethodOptionsOrBuilder)`

SetOptionsOrBuilder sets OptionsOrBuilder field to given value.

### HasOptionsOrBuilder

`func (o *MethodDescriptorProtoOrBuilder) HasOptionsOrBuilder() bool`

HasOptionsOrBuilder returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *MethodDescriptorProtoOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *MethodDescriptorProtoOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *MethodDescriptorProtoOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *MethodDescriptorProtoOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *MethodDescriptorProtoOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *MethodDescriptorProtoOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *MethodDescriptorProtoOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *MethodDescriptorProtoOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetAllFields

`func (o *MethodDescriptorProtoOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *MethodDescriptorProtoOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *MethodDescriptorProtoOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *MethodDescriptorProtoOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *MethodDescriptorProtoOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *MethodDescriptorProtoOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *MethodDescriptorProtoOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *MethodDescriptorProtoOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetUnknownFields

`func (o *MethodDescriptorProtoOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *MethodDescriptorProtoOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *MethodDescriptorProtoOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *MethodDescriptorProtoOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetInitialized

`func (o *MethodDescriptorProtoOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *MethodDescriptorProtoOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *MethodDescriptorProtoOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *MethodDescriptorProtoOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


