# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to  |  | [optional] 
**DefaultInstanceForType** | Pointer to [**Location**](Location.md) |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**LeadingComments** | Pointer to **string** |  | [optional] 
**LeadingCommentsBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**LeadingDetachedCommentsCount** | Pointer to **int32** |  | [optional] 
**LeadingDetachedCommentsList** | Pointer to **[]string** |  | [optional] 
**MemoizedSerializedSize** | Pointer to **int32** |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** |  | [optional] 
**PathCount** | Pointer to **int32** |  | [optional] 
**PathList** | Pointer to **[]int32** |  | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**SpanCount** | Pointer to **int32** |  | [optional] 
**SpanList** | Pointer to **[]int32** |  | [optional] 
**TrailingComments** | Pointer to **string** |  | [optional] 
**TrailingCommentsBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewLocation

`func NewLocation() *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *Location) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *Location) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *Location) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *Location) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### SetAllFieldsNil

`func (o *Location) SetAllFieldsNil(b bool)`

 SetAllFieldsNil sets the value for AllFields to be an explicit nil

### UnsetAllFields
`func (o *Location) UnsetAllFields()`

UnsetAllFields ensures that no value is present for AllFields, not even an explicit nil
### GetDefaultInstanceForType

`func (o *Location) GetDefaultInstanceForType() Location`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *Location) GetDefaultInstanceForTypeOk() (*Location, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *Location) SetDefaultInstanceForType(v Location)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *Location) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *Location) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *Location) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *Location) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *Location) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *Location) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *Location) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *Location) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *Location) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *Location) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *Location) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *Location) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *Location) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetLeadingComments

`func (o *Location) GetLeadingComments() string`

GetLeadingComments returns the LeadingComments field if non-nil, zero value otherwise.

### GetLeadingCommentsOk

`func (o *Location) GetLeadingCommentsOk() (*string, bool)`

GetLeadingCommentsOk returns a tuple with the LeadingComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadingComments

`func (o *Location) SetLeadingComments(v string)`

SetLeadingComments sets LeadingComments field to given value.

### HasLeadingComments

`func (o *Location) HasLeadingComments() bool`

HasLeadingComments returns a boolean if a field has been set.

### GetLeadingCommentsBytes

`func (o *Location) GetLeadingCommentsBytes() ByteString`

GetLeadingCommentsBytes returns the LeadingCommentsBytes field if non-nil, zero value otherwise.

### GetLeadingCommentsBytesOk

`func (o *Location) GetLeadingCommentsBytesOk() (*ByteString, bool)`

GetLeadingCommentsBytesOk returns a tuple with the LeadingCommentsBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadingCommentsBytes

`func (o *Location) SetLeadingCommentsBytes(v ByteString)`

SetLeadingCommentsBytes sets LeadingCommentsBytes field to given value.

### HasLeadingCommentsBytes

`func (o *Location) HasLeadingCommentsBytes() bool`

HasLeadingCommentsBytes returns a boolean if a field has been set.

### GetLeadingDetachedCommentsCount

`func (o *Location) GetLeadingDetachedCommentsCount() int32`

GetLeadingDetachedCommentsCount returns the LeadingDetachedCommentsCount field if non-nil, zero value otherwise.

### GetLeadingDetachedCommentsCountOk

`func (o *Location) GetLeadingDetachedCommentsCountOk() (*int32, bool)`

GetLeadingDetachedCommentsCountOk returns a tuple with the LeadingDetachedCommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadingDetachedCommentsCount

`func (o *Location) SetLeadingDetachedCommentsCount(v int32)`

SetLeadingDetachedCommentsCount sets LeadingDetachedCommentsCount field to given value.

### HasLeadingDetachedCommentsCount

`func (o *Location) HasLeadingDetachedCommentsCount() bool`

HasLeadingDetachedCommentsCount returns a boolean if a field has been set.

### GetLeadingDetachedCommentsList

`func (o *Location) GetLeadingDetachedCommentsList() []string`

GetLeadingDetachedCommentsList returns the LeadingDetachedCommentsList field if non-nil, zero value otherwise.

### GetLeadingDetachedCommentsListOk

`func (o *Location) GetLeadingDetachedCommentsListOk() (*[]string, bool)`

GetLeadingDetachedCommentsListOk returns a tuple with the LeadingDetachedCommentsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadingDetachedCommentsList

`func (o *Location) SetLeadingDetachedCommentsList(v []string)`

SetLeadingDetachedCommentsList sets LeadingDetachedCommentsList field to given value.

### HasLeadingDetachedCommentsList

`func (o *Location) HasLeadingDetachedCommentsList() bool`

HasLeadingDetachedCommentsList returns a boolean if a field has been set.

### GetMemoizedSerializedSize

`func (o *Location) GetMemoizedSerializedSize() int32`

GetMemoizedSerializedSize returns the MemoizedSerializedSize field if non-nil, zero value otherwise.

### GetMemoizedSerializedSizeOk

`func (o *Location) GetMemoizedSerializedSizeOk() (*int32, bool)`

GetMemoizedSerializedSizeOk returns a tuple with the MemoizedSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizedSerializedSize

`func (o *Location) SetMemoizedSerializedSize(v int32)`

SetMemoizedSerializedSize sets MemoizedSerializedSize field to given value.

### HasMemoizedSerializedSize

`func (o *Location) HasMemoizedSerializedSize() bool`

HasMemoizedSerializedSize returns a boolean if a field has been set.

### GetParserForType

`func (o *Location) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *Location) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *Location) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *Location) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetPathCount

`func (o *Location) GetPathCount() int32`

GetPathCount returns the PathCount field if non-nil, zero value otherwise.

### GetPathCountOk

`func (o *Location) GetPathCountOk() (*int32, bool)`

GetPathCountOk returns a tuple with the PathCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathCount

`func (o *Location) SetPathCount(v int32)`

SetPathCount sets PathCount field to given value.

### HasPathCount

`func (o *Location) HasPathCount() bool`

HasPathCount returns a boolean if a field has been set.

### GetPathList

`func (o *Location) GetPathList() []int32`

GetPathList returns the PathList field if non-nil, zero value otherwise.

### GetPathListOk

`func (o *Location) GetPathListOk() (*[]int32, bool)`

GetPathListOk returns a tuple with the PathList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathList

`func (o *Location) SetPathList(v []int32)`

SetPathList sets PathList field to given value.

### HasPathList

`func (o *Location) HasPathList() bool`

HasPathList returns a boolean if a field has been set.

### GetSerializedSize

`func (o *Location) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *Location) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *Location) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *Location) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetSpanCount

`func (o *Location) GetSpanCount() int32`

GetSpanCount returns the SpanCount field if non-nil, zero value otherwise.

### GetSpanCountOk

`func (o *Location) GetSpanCountOk() (*int32, bool)`

GetSpanCountOk returns a tuple with the SpanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanCount

`func (o *Location) SetSpanCount(v int32)`

SetSpanCount sets SpanCount field to given value.

### HasSpanCount

`func (o *Location) HasSpanCount() bool`

HasSpanCount returns a boolean if a field has been set.

### GetSpanList

`func (o *Location) GetSpanList() []int32`

GetSpanList returns the SpanList field if non-nil, zero value otherwise.

### GetSpanListOk

`func (o *Location) GetSpanListOk() (*[]int32, bool)`

GetSpanListOk returns a tuple with the SpanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanList

`func (o *Location) SetSpanList(v []int32)`

SetSpanList sets SpanList field to given value.

### HasSpanList

`func (o *Location) HasSpanList() bool`

HasSpanList returns a boolean if a field has been set.

### GetTrailingComments

`func (o *Location) GetTrailingComments() string`

GetTrailingComments returns the TrailingComments field if non-nil, zero value otherwise.

### GetTrailingCommentsOk

`func (o *Location) GetTrailingCommentsOk() (*string, bool)`

GetTrailingCommentsOk returns a tuple with the TrailingComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingComments

`func (o *Location) SetTrailingComments(v string)`

SetTrailingComments sets TrailingComments field to given value.

### HasTrailingComments

`func (o *Location) HasTrailingComments() bool`

HasTrailingComments returns a boolean if a field has been set.

### GetTrailingCommentsBytes

`func (o *Location) GetTrailingCommentsBytes() ByteString`

GetTrailingCommentsBytes returns the TrailingCommentsBytes field if non-nil, zero value otherwise.

### GetTrailingCommentsBytesOk

`func (o *Location) GetTrailingCommentsBytesOk() (*ByteString, bool)`

GetTrailingCommentsBytesOk returns a tuple with the TrailingCommentsBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingCommentsBytes

`func (o *Location) SetTrailingCommentsBytes(v ByteString)`

SetTrailingCommentsBytes sets TrailingCommentsBytes field to given value.

### HasTrailingCommentsBytes

`func (o *Location) HasTrailingCommentsBytes() bool`

HasTrailingCommentsBytes returns a boolean if a field has been set.

### GetUnknownFields

`func (o *Location) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *Location) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *Location) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *Location) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


