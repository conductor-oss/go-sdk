# EnumValueDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to  | Simplified schema for FileDescriptor (original had circular references) | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**EnumValueOptions**](EnumValueOptions.md) |  | [optional] 
**Proto** | Pointer to [**EnumValueDescriptorProto**](EnumValueDescriptorProto.md) |  | [optional] 
**Type** | Pointer to  | Simplified schema for EnumDescriptor (original had circular references) | [optional] 

## Methods

### NewEnumValueDescriptor

`func NewEnumValueDescriptor() *EnumValueDescriptor`

NewEnumValueDescriptor instantiates a new EnumValueDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnumValueDescriptorWithDefaults

`func NewEnumValueDescriptorWithDefaults() *EnumValueDescriptor`

NewEnumValueDescriptorWithDefaults instantiates a new EnumValueDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *EnumValueDescriptor) GetFile() map[string]interface{}`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *EnumValueDescriptor) GetFileOk() (*map[string]interface{}, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *EnumValueDescriptor) SetFile(v map[string]interface{})`

SetFile sets File field to given value.

### HasFile

`func (o *EnumValueDescriptor) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *EnumValueDescriptor) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *EnumValueDescriptor) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetFullName

`func (o *EnumValueDescriptor) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *EnumValueDescriptor) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *EnumValueDescriptor) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *EnumValueDescriptor) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetIndex

`func (o *EnumValueDescriptor) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *EnumValueDescriptor) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *EnumValueDescriptor) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *EnumValueDescriptor) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *EnumValueDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnumValueDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnumValueDescriptor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnumValueDescriptor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *EnumValueDescriptor) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *EnumValueDescriptor) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *EnumValueDescriptor) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *EnumValueDescriptor) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetOptions

`func (o *EnumValueDescriptor) GetOptions() EnumValueOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *EnumValueDescriptor) GetOptionsOk() (*EnumValueOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *EnumValueDescriptor) SetOptions(v EnumValueOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *EnumValueDescriptor) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetProto

`func (o *EnumValueDescriptor) GetProto() EnumValueDescriptorProto`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *EnumValueDescriptor) GetProtoOk() (*EnumValueDescriptorProto, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *EnumValueDescriptor) SetProto(v EnumValueDescriptorProto)`

SetProto sets Proto field to given value.

### HasProto

`func (o *EnumValueDescriptor) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetType

`func (o *EnumValueDescriptor) GetType() map[string]interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnumValueDescriptor) GetTypeOk() (*map[string]interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnumValueDescriptor) SetType(v map[string]interface{})`

SetType sets Type field to given value.

### HasType

`func (o *EnumValueDescriptor) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *EnumValueDescriptor) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *EnumValueDescriptor) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


