# UnknownFieldSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultInstanceForType** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** | Simplified schema for Parser (original had circular references) | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**SerializedSizeAsMessageSet** | Pointer to **int32** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 

## Methods

### NewUnknownFieldSet

`func NewUnknownFieldSet() *UnknownFieldSet`

NewUnknownFieldSet instantiates a new UnknownFieldSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnknownFieldSetWithDefaults

`func NewUnknownFieldSetWithDefaults() *UnknownFieldSet`

NewUnknownFieldSetWithDefaults instantiates a new UnknownFieldSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultInstanceForType

`func (o *UnknownFieldSet) GetDefaultInstanceForType() UnknownFieldSet`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *UnknownFieldSet) GetDefaultInstanceForTypeOk() (*UnknownFieldSet, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *UnknownFieldSet) SetDefaultInstanceForType(v UnknownFieldSet)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *UnknownFieldSet) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetParserForType

`func (o *UnknownFieldSet) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *UnknownFieldSet) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *UnknownFieldSet) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *UnknownFieldSet) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetSerializedSize

`func (o *UnknownFieldSet) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *UnknownFieldSet) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *UnknownFieldSet) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *UnknownFieldSet) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetSerializedSizeAsMessageSet

`func (o *UnknownFieldSet) GetSerializedSizeAsMessageSet() int32`

GetSerializedSizeAsMessageSet returns the SerializedSizeAsMessageSet field if non-nil, zero value otherwise.

### GetSerializedSizeAsMessageSetOk

`func (o *UnknownFieldSet) GetSerializedSizeAsMessageSetOk() (*int32, bool)`

GetSerializedSizeAsMessageSetOk returns a tuple with the SerializedSizeAsMessageSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSizeAsMessageSet

`func (o *UnknownFieldSet) SetSerializedSizeAsMessageSet(v int32)`

SetSerializedSizeAsMessageSet sets SerializedSizeAsMessageSet field to given value.

### HasSerializedSizeAsMessageSet

`func (o *UnknownFieldSet) HasSerializedSizeAsMessageSet() bool`

HasSerializedSizeAsMessageSet returns a boolean if a field has been set.

### GetInitialized

`func (o *UnknownFieldSet) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *UnknownFieldSet) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *UnknownFieldSet) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *UnknownFieldSet) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


