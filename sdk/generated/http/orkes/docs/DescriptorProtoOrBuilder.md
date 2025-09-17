# DescriptorProtoOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**DefaultInstanceForType** | Pointer to **map[string]interface{}** | Simplified schema for Message (original had circular references) | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**EnumTypeCount** | Pointer to **int32** |  | [optional] 
**EnumTypeList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnumTypeOrBuilderList** | Pointer to [**[]EnumDescriptorProtoOrBuilder**](EnumDescriptorProtoOrBuilder.md) |  | [optional] 
**ExtensionCount** | Pointer to **int32** |  | [optional] 
**ExtensionList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ExtensionOrBuilderList** | Pointer to [**[]FieldDescriptorProtoOrBuilder**](FieldDescriptorProtoOrBuilder.md) |  | [optional] 
**ExtensionRangeCount** | Pointer to **int32** |  | [optional] 
**ExtensionRangeList** | Pointer to [**[]ExtensionRange**](ExtensionRange.md) |  | [optional] 
**ExtensionRangeOrBuilderList** | Pointer to [**[]ExtensionRangeOrBuilder**](ExtensionRangeOrBuilder.md) |  | [optional] 
**FieldCount** | Pointer to **int32** |  | [optional] 
**FieldList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**FieldOrBuilderList** | Pointer to [**[]FieldDescriptorProtoOrBuilder**](FieldDescriptorProtoOrBuilder.md) |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**NestedTypeCount** | Pointer to **int32** |  | [optional] 
**OneofDeclCount** | Pointer to **int32** |  | [optional] 
**OneofDeclList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OneofDeclOrBuilderList** | Pointer to [**[]OneofDescriptorProtoOrBuilder**](OneofDescriptorProtoOrBuilder.md) |  | [optional] 
**Options** | Pointer to [**MessageOptions**](MessageOptions.md) |  | [optional] 
**OptionsOrBuilder** | Pointer to [**MessageOptionsOrBuilder**](MessageOptionsOrBuilder.md) |  | [optional] 
**ReservedNameCount** | Pointer to **int32** |  | [optional] 
**ReservedNameList** | Pointer to **[]string** |  | [optional] 
**ReservedRangeCount** | Pointer to **int32** |  | [optional] 
**ReservedRangeList** | Pointer to [**[]ReservedRange**](ReservedRange.md) |  | [optional] 
**ReservedRangeOrBuilderList** | Pointer to [**[]ReservedRangeOrBuilder**](ReservedRangeOrBuilder.md) |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewDescriptorProtoOrBuilder

`func NewDescriptorProtoOrBuilder() *DescriptorProtoOrBuilder`

NewDescriptorProtoOrBuilder instantiates a new DescriptorProtoOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescriptorProtoOrBuilderWithDefaults

`func NewDescriptorProtoOrBuilderWithDefaults() *DescriptorProtoOrBuilder`

NewDescriptorProtoOrBuilderWithDefaults instantiates a new DescriptorProtoOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *DescriptorProtoOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *DescriptorProtoOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *DescriptorProtoOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *DescriptorProtoOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *DescriptorProtoOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *DescriptorProtoOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *DescriptorProtoOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *DescriptorProtoOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *DescriptorProtoOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *DescriptorProtoOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *DescriptorProtoOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *DescriptorProtoOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetEnumTypeCount

`func (o *DescriptorProtoOrBuilder) GetEnumTypeCount() int32`

GetEnumTypeCount returns the EnumTypeCount field if non-nil, zero value otherwise.

### GetEnumTypeCountOk

`func (o *DescriptorProtoOrBuilder) GetEnumTypeCountOk() (*int32, bool)`

GetEnumTypeCountOk returns a tuple with the EnumTypeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumTypeCount

`func (o *DescriptorProtoOrBuilder) SetEnumTypeCount(v int32)`

SetEnumTypeCount sets EnumTypeCount field to given value.

### HasEnumTypeCount

`func (o *DescriptorProtoOrBuilder) HasEnumTypeCount() bool`

HasEnumTypeCount returns a boolean if a field has been set.

### GetEnumTypeList

`func (o *DescriptorProtoOrBuilder) GetEnumTypeList() []map[string]interface{}`

GetEnumTypeList returns the EnumTypeList field if non-nil, zero value otherwise.

### GetEnumTypeListOk

`func (o *DescriptorProtoOrBuilder) GetEnumTypeListOk() (*[]map[string]interface{}, bool)`

GetEnumTypeListOk returns a tuple with the EnumTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumTypeList

`func (o *DescriptorProtoOrBuilder) SetEnumTypeList(v []map[string]interface{})`

SetEnumTypeList sets EnumTypeList field to given value.

### HasEnumTypeList

`func (o *DescriptorProtoOrBuilder) HasEnumTypeList() bool`

HasEnumTypeList returns a boolean if a field has been set.

### GetEnumTypeOrBuilderList

`func (o *DescriptorProtoOrBuilder) GetEnumTypeOrBuilderList() []EnumDescriptorProtoOrBuilder`

GetEnumTypeOrBuilderList returns the EnumTypeOrBuilderList field if non-nil, zero value otherwise.

### GetEnumTypeOrBuilderListOk

`func (o *DescriptorProtoOrBuilder) GetEnumTypeOrBuilderListOk() (*[]EnumDescriptorProtoOrBuilder, bool)`

GetEnumTypeOrBuilderListOk returns a tuple with the EnumTypeOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumTypeOrBuilderList

`func (o *DescriptorProtoOrBuilder) SetEnumTypeOrBuilderList(v []EnumDescriptorProtoOrBuilder)`

SetEnumTypeOrBuilderList sets EnumTypeOrBuilderList field to given value.

### HasEnumTypeOrBuilderList

`func (o *DescriptorProtoOrBuilder) HasEnumTypeOrBuilderList() bool`

HasEnumTypeOrBuilderList returns a boolean if a field has been set.

### GetExtensionCount

`func (o *DescriptorProtoOrBuilder) GetExtensionCount() int32`

GetExtensionCount returns the ExtensionCount field if non-nil, zero value otherwise.

### GetExtensionCountOk

`func (o *DescriptorProtoOrBuilder) GetExtensionCountOk() (*int32, bool)`

GetExtensionCountOk returns a tuple with the ExtensionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionCount

`func (o *DescriptorProtoOrBuilder) SetExtensionCount(v int32)`

SetExtensionCount sets ExtensionCount field to given value.

### HasExtensionCount

`func (o *DescriptorProtoOrBuilder) HasExtensionCount() bool`

HasExtensionCount returns a boolean if a field has been set.

### GetExtensionList

`func (o *DescriptorProtoOrBuilder) GetExtensionList() []map[string]interface{}`

GetExtensionList returns the ExtensionList field if non-nil, zero value otherwise.

### GetExtensionListOk

`func (o *DescriptorProtoOrBuilder) GetExtensionListOk() (*[]map[string]interface{}, bool)`

GetExtensionListOk returns a tuple with the ExtensionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionList

`func (o *DescriptorProtoOrBuilder) SetExtensionList(v []map[string]interface{})`

SetExtensionList sets ExtensionList field to given value.

### HasExtensionList

`func (o *DescriptorProtoOrBuilder) HasExtensionList() bool`

HasExtensionList returns a boolean if a field has been set.

### GetExtensionOrBuilderList

`func (o *DescriptorProtoOrBuilder) GetExtensionOrBuilderList() []FieldDescriptorProtoOrBuilder`

GetExtensionOrBuilderList returns the ExtensionOrBuilderList field if non-nil, zero value otherwise.

### GetExtensionOrBuilderListOk

`func (o *DescriptorProtoOrBuilder) GetExtensionOrBuilderListOk() (*[]FieldDescriptorProtoOrBuilder, bool)`

GetExtensionOrBuilderListOk returns a tuple with the ExtensionOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionOrBuilderList

`func (o *DescriptorProtoOrBuilder) SetExtensionOrBuilderList(v []FieldDescriptorProtoOrBuilder)`

SetExtensionOrBuilderList sets ExtensionOrBuilderList field to given value.

### HasExtensionOrBuilderList

`func (o *DescriptorProtoOrBuilder) HasExtensionOrBuilderList() bool`

HasExtensionOrBuilderList returns a boolean if a field has been set.

### GetExtensionRangeCount

`func (o *DescriptorProtoOrBuilder) GetExtensionRangeCount() int32`

GetExtensionRangeCount returns the ExtensionRangeCount field if non-nil, zero value otherwise.

### GetExtensionRangeCountOk

`func (o *DescriptorProtoOrBuilder) GetExtensionRangeCountOk() (*int32, bool)`

GetExtensionRangeCountOk returns a tuple with the ExtensionRangeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionRangeCount

`func (o *DescriptorProtoOrBuilder) SetExtensionRangeCount(v int32)`

SetExtensionRangeCount sets ExtensionRangeCount field to given value.

### HasExtensionRangeCount

`func (o *DescriptorProtoOrBuilder) HasExtensionRangeCount() bool`

HasExtensionRangeCount returns a boolean if a field has been set.

### GetExtensionRangeList

`func (o *DescriptorProtoOrBuilder) GetExtensionRangeList() []ExtensionRange`

GetExtensionRangeList returns the ExtensionRangeList field if non-nil, zero value otherwise.

### GetExtensionRangeListOk

`func (o *DescriptorProtoOrBuilder) GetExtensionRangeListOk() (*[]ExtensionRange, bool)`

GetExtensionRangeListOk returns a tuple with the ExtensionRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionRangeList

`func (o *DescriptorProtoOrBuilder) SetExtensionRangeList(v []ExtensionRange)`

SetExtensionRangeList sets ExtensionRangeList field to given value.

### HasExtensionRangeList

`func (o *DescriptorProtoOrBuilder) HasExtensionRangeList() bool`

HasExtensionRangeList returns a boolean if a field has been set.

### GetExtensionRangeOrBuilderList

`func (o *DescriptorProtoOrBuilder) GetExtensionRangeOrBuilderList() []ExtensionRangeOrBuilder`

GetExtensionRangeOrBuilderList returns the ExtensionRangeOrBuilderList field if non-nil, zero value otherwise.

### GetExtensionRangeOrBuilderListOk

`func (o *DescriptorProtoOrBuilder) GetExtensionRangeOrBuilderListOk() (*[]ExtensionRangeOrBuilder, bool)`

GetExtensionRangeOrBuilderListOk returns a tuple with the ExtensionRangeOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionRangeOrBuilderList

`func (o *DescriptorProtoOrBuilder) SetExtensionRangeOrBuilderList(v []ExtensionRangeOrBuilder)`

SetExtensionRangeOrBuilderList sets ExtensionRangeOrBuilderList field to given value.

### HasExtensionRangeOrBuilderList

`func (o *DescriptorProtoOrBuilder) HasExtensionRangeOrBuilderList() bool`

HasExtensionRangeOrBuilderList returns a boolean if a field has been set.

### GetFieldCount

`func (o *DescriptorProtoOrBuilder) GetFieldCount() int32`

GetFieldCount returns the FieldCount field if non-nil, zero value otherwise.

### GetFieldCountOk

`func (o *DescriptorProtoOrBuilder) GetFieldCountOk() (*int32, bool)`

GetFieldCountOk returns a tuple with the FieldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCount

`func (o *DescriptorProtoOrBuilder) SetFieldCount(v int32)`

SetFieldCount sets FieldCount field to given value.

### HasFieldCount

`func (o *DescriptorProtoOrBuilder) HasFieldCount() bool`

HasFieldCount returns a boolean if a field has been set.

### GetFieldList

`func (o *DescriptorProtoOrBuilder) GetFieldList() []map[string]interface{}`

GetFieldList returns the FieldList field if non-nil, zero value otherwise.

### GetFieldListOk

`func (o *DescriptorProtoOrBuilder) GetFieldListOk() (*[]map[string]interface{}, bool)`

GetFieldListOk returns a tuple with the FieldList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldList

`func (o *DescriptorProtoOrBuilder) SetFieldList(v []map[string]interface{})`

SetFieldList sets FieldList field to given value.

### HasFieldList

`func (o *DescriptorProtoOrBuilder) HasFieldList() bool`

HasFieldList returns a boolean if a field has been set.

### GetFieldOrBuilderList

`func (o *DescriptorProtoOrBuilder) GetFieldOrBuilderList() []FieldDescriptorProtoOrBuilder`

GetFieldOrBuilderList returns the FieldOrBuilderList field if non-nil, zero value otherwise.

### GetFieldOrBuilderListOk

`func (o *DescriptorProtoOrBuilder) GetFieldOrBuilderListOk() (*[]FieldDescriptorProtoOrBuilder, bool)`

GetFieldOrBuilderListOk returns a tuple with the FieldOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldOrBuilderList

`func (o *DescriptorProtoOrBuilder) SetFieldOrBuilderList(v []FieldDescriptorProtoOrBuilder)`

SetFieldOrBuilderList sets FieldOrBuilderList field to given value.

### HasFieldOrBuilderList

`func (o *DescriptorProtoOrBuilder) HasFieldOrBuilderList() bool`

HasFieldOrBuilderList returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *DescriptorProtoOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *DescriptorProtoOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *DescriptorProtoOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *DescriptorProtoOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *DescriptorProtoOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *DescriptorProtoOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *DescriptorProtoOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *DescriptorProtoOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetName

`func (o *DescriptorProtoOrBuilder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescriptorProtoOrBuilder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescriptorProtoOrBuilder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DescriptorProtoOrBuilder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameBytes

`func (o *DescriptorProtoOrBuilder) GetNameBytes() ByteString`

GetNameBytes returns the NameBytes field if non-nil, zero value otherwise.

### GetNameBytesOk

`func (o *DescriptorProtoOrBuilder) GetNameBytesOk() (*ByteString, bool)`

GetNameBytesOk returns a tuple with the NameBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameBytes

`func (o *DescriptorProtoOrBuilder) SetNameBytes(v ByteString)`

SetNameBytes sets NameBytes field to given value.

### HasNameBytes

`func (o *DescriptorProtoOrBuilder) HasNameBytes() bool`

HasNameBytes returns a boolean if a field has been set.

### GetNestedTypeCount

`func (o *DescriptorProtoOrBuilder) GetNestedTypeCount() int32`

GetNestedTypeCount returns the NestedTypeCount field if non-nil, zero value otherwise.

### GetNestedTypeCountOk

`func (o *DescriptorProtoOrBuilder) GetNestedTypeCountOk() (*int32, bool)`

GetNestedTypeCountOk returns a tuple with the NestedTypeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedTypeCount

`func (o *DescriptorProtoOrBuilder) SetNestedTypeCount(v int32)`

SetNestedTypeCount sets NestedTypeCount field to given value.

### HasNestedTypeCount

`func (o *DescriptorProtoOrBuilder) HasNestedTypeCount() bool`

HasNestedTypeCount returns a boolean if a field has been set.

### GetOneofDeclCount

`func (o *DescriptorProtoOrBuilder) GetOneofDeclCount() int32`

GetOneofDeclCount returns the OneofDeclCount field if non-nil, zero value otherwise.

### GetOneofDeclCountOk

`func (o *DescriptorProtoOrBuilder) GetOneofDeclCountOk() (*int32, bool)`

GetOneofDeclCountOk returns a tuple with the OneofDeclCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneofDeclCount

`func (o *DescriptorProtoOrBuilder) SetOneofDeclCount(v int32)`

SetOneofDeclCount sets OneofDeclCount field to given value.

### HasOneofDeclCount

`func (o *DescriptorProtoOrBuilder) HasOneofDeclCount() bool`

HasOneofDeclCount returns a boolean if a field has been set.

### GetOneofDeclList

`func (o *DescriptorProtoOrBuilder) GetOneofDeclList() []map[string]interface{}`

GetOneofDeclList returns the OneofDeclList field if non-nil, zero value otherwise.

### GetOneofDeclListOk

`func (o *DescriptorProtoOrBuilder) GetOneofDeclListOk() (*[]map[string]interface{}, bool)`

GetOneofDeclListOk returns a tuple with the OneofDeclList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneofDeclList

`func (o *DescriptorProtoOrBuilder) SetOneofDeclList(v []map[string]interface{})`

SetOneofDeclList sets OneofDeclList field to given value.

### HasOneofDeclList

`func (o *DescriptorProtoOrBuilder) HasOneofDeclList() bool`

HasOneofDeclList returns a boolean if a field has been set.

### GetOneofDeclOrBuilderList

`func (o *DescriptorProtoOrBuilder) GetOneofDeclOrBuilderList() []OneofDescriptorProtoOrBuilder`

GetOneofDeclOrBuilderList returns the OneofDeclOrBuilderList field if non-nil, zero value otherwise.

### GetOneofDeclOrBuilderListOk

`func (o *DescriptorProtoOrBuilder) GetOneofDeclOrBuilderListOk() (*[]OneofDescriptorProtoOrBuilder, bool)`

GetOneofDeclOrBuilderListOk returns a tuple with the OneofDeclOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneofDeclOrBuilderList

`func (o *DescriptorProtoOrBuilder) SetOneofDeclOrBuilderList(v []OneofDescriptorProtoOrBuilder)`

SetOneofDeclOrBuilderList sets OneofDeclOrBuilderList field to given value.

### HasOneofDeclOrBuilderList

`func (o *DescriptorProtoOrBuilder) HasOneofDeclOrBuilderList() bool`

HasOneofDeclOrBuilderList returns a boolean if a field has been set.

### GetOptions

`func (o *DescriptorProtoOrBuilder) GetOptions() MessageOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DescriptorProtoOrBuilder) GetOptionsOk() (*MessageOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DescriptorProtoOrBuilder) SetOptions(v MessageOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DescriptorProtoOrBuilder) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOptionsOrBuilder

`func (o *DescriptorProtoOrBuilder) GetOptionsOrBuilder() MessageOptionsOrBuilder`

GetOptionsOrBuilder returns the OptionsOrBuilder field if non-nil, zero value otherwise.

### GetOptionsOrBuilderOk

`func (o *DescriptorProtoOrBuilder) GetOptionsOrBuilderOk() (*MessageOptionsOrBuilder, bool)`

GetOptionsOrBuilderOk returns a tuple with the OptionsOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsOrBuilder

`func (o *DescriptorProtoOrBuilder) SetOptionsOrBuilder(v MessageOptionsOrBuilder)`

SetOptionsOrBuilder sets OptionsOrBuilder field to given value.

### HasOptionsOrBuilder

`func (o *DescriptorProtoOrBuilder) HasOptionsOrBuilder() bool`

HasOptionsOrBuilder returns a boolean if a field has been set.

### GetReservedNameCount

`func (o *DescriptorProtoOrBuilder) GetReservedNameCount() int32`

GetReservedNameCount returns the ReservedNameCount field if non-nil, zero value otherwise.

### GetReservedNameCountOk

`func (o *DescriptorProtoOrBuilder) GetReservedNameCountOk() (*int32, bool)`

GetReservedNameCountOk returns a tuple with the ReservedNameCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedNameCount

`func (o *DescriptorProtoOrBuilder) SetReservedNameCount(v int32)`

SetReservedNameCount sets ReservedNameCount field to given value.

### HasReservedNameCount

`func (o *DescriptorProtoOrBuilder) HasReservedNameCount() bool`

HasReservedNameCount returns a boolean if a field has been set.

### GetReservedNameList

`func (o *DescriptorProtoOrBuilder) GetReservedNameList() []string`

GetReservedNameList returns the ReservedNameList field if non-nil, zero value otherwise.

### GetReservedNameListOk

`func (o *DescriptorProtoOrBuilder) GetReservedNameListOk() (*[]string, bool)`

GetReservedNameListOk returns a tuple with the ReservedNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedNameList

`func (o *DescriptorProtoOrBuilder) SetReservedNameList(v []string)`

SetReservedNameList sets ReservedNameList field to given value.

### HasReservedNameList

`func (o *DescriptorProtoOrBuilder) HasReservedNameList() bool`

HasReservedNameList returns a boolean if a field has been set.

### GetReservedRangeCount

`func (o *DescriptorProtoOrBuilder) GetReservedRangeCount() int32`

GetReservedRangeCount returns the ReservedRangeCount field if non-nil, zero value otherwise.

### GetReservedRangeCountOk

`func (o *DescriptorProtoOrBuilder) GetReservedRangeCountOk() (*int32, bool)`

GetReservedRangeCountOk returns a tuple with the ReservedRangeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedRangeCount

`func (o *DescriptorProtoOrBuilder) SetReservedRangeCount(v int32)`

SetReservedRangeCount sets ReservedRangeCount field to given value.

### HasReservedRangeCount

`func (o *DescriptorProtoOrBuilder) HasReservedRangeCount() bool`

HasReservedRangeCount returns a boolean if a field has been set.

### GetReservedRangeList

`func (o *DescriptorProtoOrBuilder) GetReservedRangeList() []ReservedRange`

GetReservedRangeList returns the ReservedRangeList field if non-nil, zero value otherwise.

### GetReservedRangeListOk

`func (o *DescriptorProtoOrBuilder) GetReservedRangeListOk() (*[]ReservedRange, bool)`

GetReservedRangeListOk returns a tuple with the ReservedRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedRangeList

`func (o *DescriptorProtoOrBuilder) SetReservedRangeList(v []ReservedRange)`

SetReservedRangeList sets ReservedRangeList field to given value.

### HasReservedRangeList

`func (o *DescriptorProtoOrBuilder) HasReservedRangeList() bool`

HasReservedRangeList returns a boolean if a field has been set.

### GetReservedRangeOrBuilderList

`func (o *DescriptorProtoOrBuilder) GetReservedRangeOrBuilderList() []ReservedRangeOrBuilder`

GetReservedRangeOrBuilderList returns the ReservedRangeOrBuilderList field if non-nil, zero value otherwise.

### GetReservedRangeOrBuilderListOk

`func (o *DescriptorProtoOrBuilder) GetReservedRangeOrBuilderListOk() (*[]ReservedRangeOrBuilder, bool)`

GetReservedRangeOrBuilderListOk returns a tuple with the ReservedRangeOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedRangeOrBuilderList

`func (o *DescriptorProtoOrBuilder) SetReservedRangeOrBuilderList(v []ReservedRangeOrBuilder)`

SetReservedRangeOrBuilderList sets ReservedRangeOrBuilderList field to given value.

### HasReservedRangeOrBuilderList

`func (o *DescriptorProtoOrBuilder) HasReservedRangeOrBuilderList() bool`

HasReservedRangeOrBuilderList returns a boolean if a field has been set.

### GetUnknownFields

`func (o *DescriptorProtoOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *DescriptorProtoOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *DescriptorProtoOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *DescriptorProtoOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


