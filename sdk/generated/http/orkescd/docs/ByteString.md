# ByteString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Empty** | Pointer to **bool** |  | [optional] 
**ValidUtf8** | Pointer to **bool** |  | [optional] 

## Methods

### NewByteString

`func NewByteString() *ByteString`

NewByteString instantiates a new ByteString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByteStringWithDefaults

`func NewByteStringWithDefaults() *ByteString`

NewByteStringWithDefaults instantiates a new ByteString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmpty

`func (o *ByteString) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *ByteString) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *ByteString) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *ByteString) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetValidUtf8

`func (o *ByteString) GetValidUtf8() bool`

GetValidUtf8 returns the ValidUtf8 field if non-nil, zero value otherwise.

### GetValidUtf8Ok

`func (o *ByteString) GetValidUtf8Ok() (*bool, bool)`

GetValidUtf8Ok returns a tuple with the ValidUtf8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUtf8

`func (o *ByteString) SetValidUtf8(v bool)`

SetValidUtf8 sets ValidUtf8 field to given value.

### HasValidUtf8

`func (o *ByteString) HasValidUtf8() bool`

HasValidUtf8 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


