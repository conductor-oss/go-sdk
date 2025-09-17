# LocationOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**DefaultInstanceForType** | Pointer to **map[string]interface{}** | Simplified schema for Message (original had circular references) | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**LeadingComments** | Pointer to **string** |  | [optional] 
**LeadingCommentsBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**LeadingDetachedCommentsCount** | Pointer to **int32** |  | [optional] 
**LeadingDetachedCommentsList** | Pointer to **[]string** |  | [optional] 
**PathCount** | Pointer to **int32** |  | [optional] 
**PathList** | Pointer to **[]int32** |  | [optional] 
**SpanCount** | Pointer to **int32** |  | [optional] 
**SpanList** | Pointer to **[]int32** |  | [optional] 
**TrailingComments** | Pointer to **string** |  | [optional] 
**TrailingCommentsBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewLocationOrBuilder

`func NewLocationOrBuilder() *LocationOrBuilder`

NewLocationOrBuilder instantiates a new LocationOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationOrBuilderWithDefaults

`func NewLocationOrBuilderWithDefaults() *LocationOrBuilder`

NewLocationOrBuilderWithDefaults instantiates a new LocationOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *LocationOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *LocationOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *LocationOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *LocationOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *LocationOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *LocationOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *LocationOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *LocationOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *LocationOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *LocationOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *LocationOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *LocationOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *LocationOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *LocationOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *LocationOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *LocationOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *LocationOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *LocationOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *LocationOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *LocationOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetLeadingComments

`func (o *LocationOrBuilder) GetLeadingComments() string`

GetLeadingComments returns the LeadingComments field if non-nil, zero value otherwise.

### GetLeadingCommentsOk

`func (o *LocationOrBuilder) GetLeadingCommentsOk() (*string, bool)`

GetLeadingCommentsOk returns a tuple with the LeadingComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadingComments

`func (o *LocationOrBuilder) SetLeadingComments(v string)`

SetLeadingComments sets LeadingComments field to given value.

### HasLeadingComments

`func (o *LocationOrBuilder) HasLeadingComments() bool`

HasLeadingComments returns a boolean if a field has been set.

### GetLeadingCommentsBytes

`func (o *LocationOrBuilder) GetLeadingCommentsBytes() ByteString`

GetLeadingCommentsBytes returns the LeadingCommentsBytes field if non-nil, zero value otherwise.

### GetLeadingCommentsBytesOk

`func (o *LocationOrBuilder) GetLeadingCommentsBytesOk() (*ByteString, bool)`

GetLeadingCommentsBytesOk returns a tuple with the LeadingCommentsBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadingCommentsBytes

`func (o *LocationOrBuilder) SetLeadingCommentsBytes(v ByteString)`

SetLeadingCommentsBytes sets LeadingCommentsBytes field to given value.

### HasLeadingCommentsBytes

`func (o *LocationOrBuilder) HasLeadingCommentsBytes() bool`

HasLeadingCommentsBytes returns a boolean if a field has been set.

### GetLeadingDetachedCommentsCount

`func (o *LocationOrBuilder) GetLeadingDetachedCommentsCount() int32`

GetLeadingDetachedCommentsCount returns the LeadingDetachedCommentsCount field if non-nil, zero value otherwise.

### GetLeadingDetachedCommentsCountOk

`func (o *LocationOrBuilder) GetLeadingDetachedCommentsCountOk() (*int32, bool)`

GetLeadingDetachedCommentsCountOk returns a tuple with the LeadingDetachedCommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadingDetachedCommentsCount

`func (o *LocationOrBuilder) SetLeadingDetachedCommentsCount(v int32)`

SetLeadingDetachedCommentsCount sets LeadingDetachedCommentsCount field to given value.

### HasLeadingDetachedCommentsCount

`func (o *LocationOrBuilder) HasLeadingDetachedCommentsCount() bool`

HasLeadingDetachedCommentsCount returns a boolean if a field has been set.

### GetLeadingDetachedCommentsList

`func (o *LocationOrBuilder) GetLeadingDetachedCommentsList() []string`

GetLeadingDetachedCommentsList returns the LeadingDetachedCommentsList field if non-nil, zero value otherwise.

### GetLeadingDetachedCommentsListOk

`func (o *LocationOrBuilder) GetLeadingDetachedCommentsListOk() (*[]string, bool)`

GetLeadingDetachedCommentsListOk returns a tuple with the LeadingDetachedCommentsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadingDetachedCommentsList

`func (o *LocationOrBuilder) SetLeadingDetachedCommentsList(v []string)`

SetLeadingDetachedCommentsList sets LeadingDetachedCommentsList field to given value.

### HasLeadingDetachedCommentsList

`func (o *LocationOrBuilder) HasLeadingDetachedCommentsList() bool`

HasLeadingDetachedCommentsList returns a boolean if a field has been set.

### GetPathCount

`func (o *LocationOrBuilder) GetPathCount() int32`

GetPathCount returns the PathCount field if non-nil, zero value otherwise.

### GetPathCountOk

`func (o *LocationOrBuilder) GetPathCountOk() (*int32, bool)`

GetPathCountOk returns a tuple with the PathCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathCount

`func (o *LocationOrBuilder) SetPathCount(v int32)`

SetPathCount sets PathCount field to given value.

### HasPathCount

`func (o *LocationOrBuilder) HasPathCount() bool`

HasPathCount returns a boolean if a field has been set.

### GetPathList

`func (o *LocationOrBuilder) GetPathList() []int32`

GetPathList returns the PathList field if non-nil, zero value otherwise.

### GetPathListOk

`func (o *LocationOrBuilder) GetPathListOk() (*[]int32, bool)`

GetPathListOk returns a tuple with the PathList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathList

`func (o *LocationOrBuilder) SetPathList(v []int32)`

SetPathList sets PathList field to given value.

### HasPathList

`func (o *LocationOrBuilder) HasPathList() bool`

HasPathList returns a boolean if a field has been set.

### GetSpanCount

`func (o *LocationOrBuilder) GetSpanCount() int32`

GetSpanCount returns the SpanCount field if non-nil, zero value otherwise.

### GetSpanCountOk

`func (o *LocationOrBuilder) GetSpanCountOk() (*int32, bool)`

GetSpanCountOk returns a tuple with the SpanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanCount

`func (o *LocationOrBuilder) SetSpanCount(v int32)`

SetSpanCount sets SpanCount field to given value.

### HasSpanCount

`func (o *LocationOrBuilder) HasSpanCount() bool`

HasSpanCount returns a boolean if a field has been set.

### GetSpanList

`func (o *LocationOrBuilder) GetSpanList() []int32`

GetSpanList returns the SpanList field if non-nil, zero value otherwise.

### GetSpanListOk

`func (o *LocationOrBuilder) GetSpanListOk() (*[]int32, bool)`

GetSpanListOk returns a tuple with the SpanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanList

`func (o *LocationOrBuilder) SetSpanList(v []int32)`

SetSpanList sets SpanList field to given value.

### HasSpanList

`func (o *LocationOrBuilder) HasSpanList() bool`

HasSpanList returns a boolean if a field has been set.

### GetTrailingComments

`func (o *LocationOrBuilder) GetTrailingComments() string`

GetTrailingComments returns the TrailingComments field if non-nil, zero value otherwise.

### GetTrailingCommentsOk

`func (o *LocationOrBuilder) GetTrailingCommentsOk() (*string, bool)`

GetTrailingCommentsOk returns a tuple with the TrailingComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingComments

`func (o *LocationOrBuilder) SetTrailingComments(v string)`

SetTrailingComments sets TrailingComments field to given value.

### HasTrailingComments

`func (o *LocationOrBuilder) HasTrailingComments() bool`

HasTrailingComments returns a boolean if a field has been set.

### GetTrailingCommentsBytes

`func (o *LocationOrBuilder) GetTrailingCommentsBytes() ByteString`

GetTrailingCommentsBytes returns the TrailingCommentsBytes field if non-nil, zero value otherwise.

### GetTrailingCommentsBytesOk

`func (o *LocationOrBuilder) GetTrailingCommentsBytesOk() (*ByteString, bool)`

GetTrailingCommentsBytesOk returns a tuple with the TrailingCommentsBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingCommentsBytes

`func (o *LocationOrBuilder) SetTrailingCommentsBytes(v ByteString)`

SetTrailingCommentsBytes sets TrailingCommentsBytes field to given value.

### HasTrailingCommentsBytes

`func (o *LocationOrBuilder) HasTrailingCommentsBytes() bool`

HasTrailingCommentsBytes returns a boolean if a field has been set.

### GetUnknownFields

`func (o *LocationOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *LocationOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *LocationOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *LocationOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


