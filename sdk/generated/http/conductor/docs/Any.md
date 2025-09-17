# Any

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**Value** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**DefaultInstanceForType** | Pointer to [**Any**](Any.md) |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** | Simplified schema for ParserAny (original had circular references) | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**TypeUrl** | Pointer to **string** |  | [optional] 
**TypeUrlBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**MemoizedSerializedSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewAny

`func NewAny() *Any`

NewAny instantiates a new Any object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnyWithDefaults

`func NewAnyWithDefaults() *Any`

NewAnyWithDefaults instantiates a new Any object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnknownFields

`func (o *Any) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *Any) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *Any) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *Any) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetValue

`func (o *Any) GetValue() ByteString`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Any) GetValueOk() (*ByteString, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Any) SetValue(v ByteString)`

SetValue sets Value field to given value.

### HasValue

`func (o *Any) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *Any) GetDefaultInstanceForType() Any`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *Any) GetDefaultInstanceForTypeOk() (*Any, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *Any) SetDefaultInstanceForType(v Any)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *Any) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetParserForType

`func (o *Any) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *Any) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *Any) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *Any) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetSerializedSize

`func (o *Any) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *Any) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *Any) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *Any) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetTypeUrl

`func (o *Any) GetTypeUrl() string`

GetTypeUrl returns the TypeUrl field if non-nil, zero value otherwise.

### GetTypeUrlOk

`func (o *Any) GetTypeUrlOk() (*string, bool)`

GetTypeUrlOk returns a tuple with the TypeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeUrl

`func (o *Any) SetTypeUrl(v string)`

SetTypeUrl sets TypeUrl field to given value.

### HasTypeUrl

`func (o *Any) HasTypeUrl() bool`

HasTypeUrl returns a boolean if a field has been set.

### GetTypeUrlBytes

`func (o *Any) GetTypeUrlBytes() ByteString`

GetTypeUrlBytes returns the TypeUrlBytes field if non-nil, zero value otherwise.

### GetTypeUrlBytesOk

`func (o *Any) GetTypeUrlBytesOk() (*ByteString, bool)`

GetTypeUrlBytesOk returns a tuple with the TypeUrlBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeUrlBytes

`func (o *Any) SetTypeUrlBytes(v ByteString)`

SetTypeUrlBytes sets TypeUrlBytes field to given value.

### HasTypeUrlBytes

`func (o *Any) HasTypeUrlBytes() bool`

HasTypeUrlBytes returns a boolean if a field has been set.

### GetInitialized

`func (o *Any) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *Any) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *Any) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *Any) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *Any) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *Any) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *Any) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *Any) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *Any) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *Any) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *Any) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *Any) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetAllFields

`func (o *Any) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *Any) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *Any) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *Any) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetMemoizedSerializedSize

`func (o *Any) GetMemoizedSerializedSize() int32`

GetMemoizedSerializedSize returns the MemoizedSerializedSize field if non-nil, zero value otherwise.

### GetMemoizedSerializedSizeOk

`func (o *Any) GetMemoizedSerializedSizeOk() (*int32, bool)`

GetMemoizedSerializedSizeOk returns a tuple with the MemoizedSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizedSerializedSize

`func (o *Any) SetMemoizedSerializedSize(v int32)`

SetMemoizedSerializedSize sets MemoizedSerializedSize field to given value.

### HasMemoizedSerializedSize

`func (o *Any) HasMemoizedSerializedSize() bool`

HasMemoizedSerializedSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


