# ServiceDescriptorProtoOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**DefaultInstanceForType** | Pointer to **map[string]interface{}** | Simplified schema for Message (original had circular references) | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**MethodCount** | Pointer to **int32** |  | [optional] 
**MethodList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**MethodOrBuilderList** | Pointer to [**[]MethodDescriptorProtoOrBuilder**](MethodDescriptorProtoOrBuilder.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**Options** | Pointer to [**ServiceOptions**](ServiceOptions.md) |  | [optional] 
**OptionsOrBuilder** | Pointer to [**ServiceOptionsOrBuilder**](ServiceOptionsOrBuilder.md) |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewServiceDescriptorProtoOrBuilder

`func NewServiceDescriptorProtoOrBuilder() *ServiceDescriptorProtoOrBuilder`

NewServiceDescriptorProtoOrBuilder instantiates a new ServiceDescriptorProtoOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDescriptorProtoOrBuilderWithDefaults

`func NewServiceDescriptorProtoOrBuilderWithDefaults() *ServiceDescriptorProtoOrBuilder`

NewServiceDescriptorProtoOrBuilderWithDefaults instantiates a new ServiceDescriptorProtoOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *ServiceDescriptorProtoOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *ServiceDescriptorProtoOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *ServiceDescriptorProtoOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *ServiceDescriptorProtoOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *ServiceDescriptorProtoOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *ServiceDescriptorProtoOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *ServiceDescriptorProtoOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *ServiceDescriptorProtoOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *ServiceDescriptorProtoOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *ServiceDescriptorProtoOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *ServiceDescriptorProtoOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *ServiceDescriptorProtoOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *ServiceDescriptorProtoOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *ServiceDescriptorProtoOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *ServiceDescriptorProtoOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *ServiceDescriptorProtoOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *ServiceDescriptorProtoOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *ServiceDescriptorProtoOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *ServiceDescriptorProtoOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *ServiceDescriptorProtoOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetMethodCount

`func (o *ServiceDescriptorProtoOrBuilder) GetMethodCount() int32`

GetMethodCount returns the MethodCount field if non-nil, zero value otherwise.

### GetMethodCountOk

`func (o *ServiceDescriptorProtoOrBuilder) GetMethodCountOk() (*int32, bool)`

GetMethodCountOk returns a tuple with the MethodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodCount

`func (o *ServiceDescriptorProtoOrBuilder) SetMethodCount(v int32)`

SetMethodCount sets MethodCount field to given value.

### HasMethodCount

`func (o *ServiceDescriptorProtoOrBuilder) HasMethodCount() bool`

HasMethodCount returns a boolean if a field has been set.

### GetMethodList

`func (o *ServiceDescriptorProtoOrBuilder) GetMethodList() []map[string]interface{}`

GetMethodList returns the MethodList field if non-nil, zero value otherwise.

### GetMethodListOk

`func (o *ServiceDescriptorProtoOrBuilder) GetMethodListOk() (*[]map[string]interface{}, bool)`

GetMethodListOk returns a tuple with the MethodList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodList

`func (o *ServiceDescriptorProtoOrBuilder) SetMethodList(v []map[string]interface{})`

SetMethodList sets MethodList field to given value.

### HasMethodList

`func (o *ServiceDescriptorProtoOrBuilder) HasMethodList() bool`

HasMethodList returns a boolean if a field has been set.

### GetMethodOrBuilderList

`func (o *ServiceDescriptorProtoOrBuilder) GetMethodOrBuilderList() []MethodDescriptorProtoOrBuilder`

GetMethodOrBuilderList returns the MethodOrBuilderList field if non-nil, zero value otherwise.

### GetMethodOrBuilderListOk

`func (o *ServiceDescriptorProtoOrBuilder) GetMethodOrBuilderListOk() (*[]MethodDescriptorProtoOrBuilder, bool)`

GetMethodOrBuilderListOk returns a tuple with the MethodOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodOrBuilderList

`func (o *ServiceDescriptorProtoOrBuilder) SetMethodOrBuilderList(v []MethodDescriptorProtoOrBuilder)`

SetMethodOrBuilderList sets MethodOrBuilderList field to given value.

### HasMethodOrBuilderList

`func (o *ServiceDescriptorProtoOrBuilder) HasMethodOrBuilderList() bool`

HasMethodOrBuilderList returns a boolean if a field has been set.

### GetName

`func (o *ServiceDescriptorProtoOrBuilder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceDescriptorProtoOrBuilder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceDescriptorProtoOrBuilder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceDescriptorProtoOrBuilder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameBytes

`func (o *ServiceDescriptorProtoOrBuilder) GetNameBytes() ByteString`

GetNameBytes returns the NameBytes field if non-nil, zero value otherwise.

### GetNameBytesOk

`func (o *ServiceDescriptorProtoOrBuilder) GetNameBytesOk() (*ByteString, bool)`

GetNameBytesOk returns a tuple with the NameBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameBytes

`func (o *ServiceDescriptorProtoOrBuilder) SetNameBytes(v ByteString)`

SetNameBytes sets NameBytes field to given value.

### HasNameBytes

`func (o *ServiceDescriptorProtoOrBuilder) HasNameBytes() bool`

HasNameBytes returns a boolean if a field has been set.

### GetOptions

`func (o *ServiceDescriptorProtoOrBuilder) GetOptions() ServiceOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ServiceDescriptorProtoOrBuilder) GetOptionsOk() (*ServiceOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ServiceDescriptorProtoOrBuilder) SetOptions(v ServiceOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ServiceDescriptorProtoOrBuilder) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOptionsOrBuilder

`func (o *ServiceDescriptorProtoOrBuilder) GetOptionsOrBuilder() ServiceOptionsOrBuilder`

GetOptionsOrBuilder returns the OptionsOrBuilder field if non-nil, zero value otherwise.

### GetOptionsOrBuilderOk

`func (o *ServiceDescriptorProtoOrBuilder) GetOptionsOrBuilderOk() (*ServiceOptionsOrBuilder, bool)`

GetOptionsOrBuilderOk returns a tuple with the OptionsOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsOrBuilder

`func (o *ServiceDescriptorProtoOrBuilder) SetOptionsOrBuilder(v ServiceOptionsOrBuilder)`

SetOptionsOrBuilder sets OptionsOrBuilder field to given value.

### HasOptionsOrBuilder

`func (o *ServiceDescriptorProtoOrBuilder) HasOptionsOrBuilder() bool`

HasOptionsOrBuilder returns a boolean if a field has been set.

### GetUnknownFields

`func (o *ServiceDescriptorProtoOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *ServiceDescriptorProtoOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *ServiceDescriptorProtoOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *ServiceDescriptorProtoOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


