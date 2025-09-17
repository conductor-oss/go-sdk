# FileOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 
**DefaultInstanceForType** | Pointer to [**FileOptions**](FileOptions.md) |  | [optional] 
**ParserForType** | Pointer to **map[string]interface{}** |  | [optional] 
**SerializedSize** | Pointer to **int32** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**JavaStringCheckUtf8** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to [**FeatureSet**](FeatureSet.md) |  | [optional] 
**UninterpretedOptionOrBuilderList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**JavaPackage** | Pointer to **string** |  | [optional] 
**OptimizeFor** | Pointer to **string** |  | [optional] 
**JavaPackageBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**JavaOuterClassname** | Pointer to **string** |  | [optional] 
**JavaOuterClassnameBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**JavaMultipleFiles** | Pointer to **bool** |  | [optional] 
**JavaGenerateEqualsAndHash** | Pointer to **bool** |  | [optional] 
**CcGenericServices** | Pointer to **bool** |  | [optional] 
**GoPackage** | Pointer to **string** |  | [optional] 
**GoPackageBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**CcEnableArenas** | Pointer to **bool** |  | [optional] 
**ObjcClassPrefix** | Pointer to **string** |  | [optional] 
**CsharpNamespace** | Pointer to **string** |  | [optional] 
**SwiftPrefix** | Pointer to **string** |  | [optional] 
**JavaGenericServices** | Pointer to **bool** |  | [optional] 
**PyGenericServices** | Pointer to **bool** |  | [optional] 
**PhpGenericServices** | Pointer to **bool** |  | [optional] 
**ObjcClassPrefixBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**CsharpNamespaceBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**SwiftPrefixBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**PhpClassPrefixBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**PhpNamespaceBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**PhpMetadataNamespace** | Pointer to **string** |  | [optional] 
**PhpMetadataNamespaceBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**RubyPackageBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**PhpClassPrefix** | Pointer to **string** |  | [optional] 
**PhpNamespace** | Pointer to **string** |  | [optional] 
**RubyPackage** | Pointer to **string** |  | [optional] 
**FeaturesOrBuilder** | Pointer to [**FeatureSetOrBuilder**](FeatureSetOrBuilder.md) |  | [optional] 
**UninterpretedOptionList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UninterpretedOptionCount** | Pointer to **int32** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**AllFieldsRaw** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**MemoizedSerializedSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewFileOptions

`func NewFileOptions() *FileOptions`

NewFileOptions instantiates a new FileOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileOptionsWithDefaults

`func NewFileOptionsWithDefaults() *FileOptions`

NewFileOptionsWithDefaults instantiates a new FileOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnknownFields

`func (o *FileOptions) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *FileOptions) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *FileOptions) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *FileOptions) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *FileOptions) GetDefaultInstanceForType() FileOptions`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *FileOptions) GetDefaultInstanceForTypeOk() (*FileOptions, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *FileOptions) SetDefaultInstanceForType(v FileOptions)`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *FileOptions) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetParserForType

`func (o *FileOptions) GetParserForType() map[string]interface{}`

GetParserForType returns the ParserForType field if non-nil, zero value otherwise.

### GetParserForTypeOk

`func (o *FileOptions) GetParserForTypeOk() (*map[string]interface{}, bool)`

GetParserForTypeOk returns a tuple with the ParserForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParserForType

`func (o *FileOptions) SetParserForType(v map[string]interface{})`

SetParserForType sets ParserForType field to given value.

### HasParserForType

`func (o *FileOptions) HasParserForType() bool`

HasParserForType returns a boolean if a field has been set.

### GetSerializedSize

`func (o *FileOptions) GetSerializedSize() int32`

GetSerializedSize returns the SerializedSize field if non-nil, zero value otherwise.

### GetSerializedSizeOk

`func (o *FileOptions) GetSerializedSizeOk() (*int32, bool)`

GetSerializedSizeOk returns a tuple with the SerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedSize

`func (o *FileOptions) SetSerializedSize(v int32)`

SetSerializedSize sets SerializedSize field to given value.

### HasSerializedSize

`func (o *FileOptions) HasSerializedSize() bool`

HasSerializedSize returns a boolean if a field has been set.

### GetDeprecated

`func (o *FileOptions) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *FileOptions) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *FileOptions) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *FileOptions) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetJavaStringCheckUtf8

`func (o *FileOptions) GetJavaStringCheckUtf8() bool`

GetJavaStringCheckUtf8 returns the JavaStringCheckUtf8 field if non-nil, zero value otherwise.

### GetJavaStringCheckUtf8Ok

`func (o *FileOptions) GetJavaStringCheckUtf8Ok() (*bool, bool)`

GetJavaStringCheckUtf8Ok returns a tuple with the JavaStringCheckUtf8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaStringCheckUtf8

`func (o *FileOptions) SetJavaStringCheckUtf8(v bool)`

SetJavaStringCheckUtf8 sets JavaStringCheckUtf8 field to given value.

### HasJavaStringCheckUtf8

`func (o *FileOptions) HasJavaStringCheckUtf8() bool`

HasJavaStringCheckUtf8 returns a boolean if a field has been set.

### GetFeatures

`func (o *FileOptions) GetFeatures() FeatureSet`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FileOptions) GetFeaturesOk() (*FeatureSet, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FileOptions) SetFeatures(v FeatureSet)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FileOptions) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetUninterpretedOptionOrBuilderList

`func (o *FileOptions) GetUninterpretedOptionOrBuilderList() []map[string]interface{}`

GetUninterpretedOptionOrBuilderList returns the UninterpretedOptionOrBuilderList field if non-nil, zero value otherwise.

### GetUninterpretedOptionOrBuilderListOk

`func (o *FileOptions) GetUninterpretedOptionOrBuilderListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionOrBuilderListOk returns a tuple with the UninterpretedOptionOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionOrBuilderList

`func (o *FileOptions) SetUninterpretedOptionOrBuilderList(v []map[string]interface{})`

SetUninterpretedOptionOrBuilderList sets UninterpretedOptionOrBuilderList field to given value.

### HasUninterpretedOptionOrBuilderList

`func (o *FileOptions) HasUninterpretedOptionOrBuilderList() bool`

HasUninterpretedOptionOrBuilderList returns a boolean if a field has been set.

### GetJavaPackage

`func (o *FileOptions) GetJavaPackage() string`

GetJavaPackage returns the JavaPackage field if non-nil, zero value otherwise.

### GetJavaPackageOk

`func (o *FileOptions) GetJavaPackageOk() (*string, bool)`

GetJavaPackageOk returns a tuple with the JavaPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaPackage

`func (o *FileOptions) SetJavaPackage(v string)`

SetJavaPackage sets JavaPackage field to given value.

### HasJavaPackage

`func (o *FileOptions) HasJavaPackage() bool`

HasJavaPackage returns a boolean if a field has been set.

### GetOptimizeFor

`func (o *FileOptions) GetOptimizeFor() string`

GetOptimizeFor returns the OptimizeFor field if non-nil, zero value otherwise.

### GetOptimizeForOk

`func (o *FileOptions) GetOptimizeForOk() (*string, bool)`

GetOptimizeForOk returns a tuple with the OptimizeFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeFor

`func (o *FileOptions) SetOptimizeFor(v string)`

SetOptimizeFor sets OptimizeFor field to given value.

### HasOptimizeFor

`func (o *FileOptions) HasOptimizeFor() bool`

HasOptimizeFor returns a boolean if a field has been set.

### GetJavaPackageBytes

`func (o *FileOptions) GetJavaPackageBytes() ByteString`

GetJavaPackageBytes returns the JavaPackageBytes field if non-nil, zero value otherwise.

### GetJavaPackageBytesOk

`func (o *FileOptions) GetJavaPackageBytesOk() (*ByteString, bool)`

GetJavaPackageBytesOk returns a tuple with the JavaPackageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaPackageBytes

`func (o *FileOptions) SetJavaPackageBytes(v ByteString)`

SetJavaPackageBytes sets JavaPackageBytes field to given value.

### HasJavaPackageBytes

`func (o *FileOptions) HasJavaPackageBytes() bool`

HasJavaPackageBytes returns a boolean if a field has been set.

### GetJavaOuterClassname

`func (o *FileOptions) GetJavaOuterClassname() string`

GetJavaOuterClassname returns the JavaOuterClassname field if non-nil, zero value otherwise.

### GetJavaOuterClassnameOk

`func (o *FileOptions) GetJavaOuterClassnameOk() (*string, bool)`

GetJavaOuterClassnameOk returns a tuple with the JavaOuterClassname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaOuterClassname

`func (o *FileOptions) SetJavaOuterClassname(v string)`

SetJavaOuterClassname sets JavaOuterClassname field to given value.

### HasJavaOuterClassname

`func (o *FileOptions) HasJavaOuterClassname() bool`

HasJavaOuterClassname returns a boolean if a field has been set.

### GetJavaOuterClassnameBytes

`func (o *FileOptions) GetJavaOuterClassnameBytes() ByteString`

GetJavaOuterClassnameBytes returns the JavaOuterClassnameBytes field if non-nil, zero value otherwise.

### GetJavaOuterClassnameBytesOk

`func (o *FileOptions) GetJavaOuterClassnameBytesOk() (*ByteString, bool)`

GetJavaOuterClassnameBytesOk returns a tuple with the JavaOuterClassnameBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaOuterClassnameBytes

`func (o *FileOptions) SetJavaOuterClassnameBytes(v ByteString)`

SetJavaOuterClassnameBytes sets JavaOuterClassnameBytes field to given value.

### HasJavaOuterClassnameBytes

`func (o *FileOptions) HasJavaOuterClassnameBytes() bool`

HasJavaOuterClassnameBytes returns a boolean if a field has been set.

### GetJavaMultipleFiles

`func (o *FileOptions) GetJavaMultipleFiles() bool`

GetJavaMultipleFiles returns the JavaMultipleFiles field if non-nil, zero value otherwise.

### GetJavaMultipleFilesOk

`func (o *FileOptions) GetJavaMultipleFilesOk() (*bool, bool)`

GetJavaMultipleFilesOk returns a tuple with the JavaMultipleFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaMultipleFiles

`func (o *FileOptions) SetJavaMultipleFiles(v bool)`

SetJavaMultipleFiles sets JavaMultipleFiles field to given value.

### HasJavaMultipleFiles

`func (o *FileOptions) HasJavaMultipleFiles() bool`

HasJavaMultipleFiles returns a boolean if a field has been set.

### GetJavaGenerateEqualsAndHash

`func (o *FileOptions) GetJavaGenerateEqualsAndHash() bool`

GetJavaGenerateEqualsAndHash returns the JavaGenerateEqualsAndHash field if non-nil, zero value otherwise.

### GetJavaGenerateEqualsAndHashOk

`func (o *FileOptions) GetJavaGenerateEqualsAndHashOk() (*bool, bool)`

GetJavaGenerateEqualsAndHashOk returns a tuple with the JavaGenerateEqualsAndHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaGenerateEqualsAndHash

`func (o *FileOptions) SetJavaGenerateEqualsAndHash(v bool)`

SetJavaGenerateEqualsAndHash sets JavaGenerateEqualsAndHash field to given value.

### HasJavaGenerateEqualsAndHash

`func (o *FileOptions) HasJavaGenerateEqualsAndHash() bool`

HasJavaGenerateEqualsAndHash returns a boolean if a field has been set.

### GetCcGenericServices

`func (o *FileOptions) GetCcGenericServices() bool`

GetCcGenericServices returns the CcGenericServices field if non-nil, zero value otherwise.

### GetCcGenericServicesOk

`func (o *FileOptions) GetCcGenericServicesOk() (*bool, bool)`

GetCcGenericServicesOk returns a tuple with the CcGenericServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcGenericServices

`func (o *FileOptions) SetCcGenericServices(v bool)`

SetCcGenericServices sets CcGenericServices field to given value.

### HasCcGenericServices

`func (o *FileOptions) HasCcGenericServices() bool`

HasCcGenericServices returns a boolean if a field has been set.

### GetGoPackage

`func (o *FileOptions) GetGoPackage() string`

GetGoPackage returns the GoPackage field if non-nil, zero value otherwise.

### GetGoPackageOk

`func (o *FileOptions) GetGoPackageOk() (*string, bool)`

GetGoPackageOk returns a tuple with the GoPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoPackage

`func (o *FileOptions) SetGoPackage(v string)`

SetGoPackage sets GoPackage field to given value.

### HasGoPackage

`func (o *FileOptions) HasGoPackage() bool`

HasGoPackage returns a boolean if a field has been set.

### GetGoPackageBytes

`func (o *FileOptions) GetGoPackageBytes() ByteString`

GetGoPackageBytes returns the GoPackageBytes field if non-nil, zero value otherwise.

### GetGoPackageBytesOk

`func (o *FileOptions) GetGoPackageBytesOk() (*ByteString, bool)`

GetGoPackageBytesOk returns a tuple with the GoPackageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoPackageBytes

`func (o *FileOptions) SetGoPackageBytes(v ByteString)`

SetGoPackageBytes sets GoPackageBytes field to given value.

### HasGoPackageBytes

`func (o *FileOptions) HasGoPackageBytes() bool`

HasGoPackageBytes returns a boolean if a field has been set.

### GetCcEnableArenas

`func (o *FileOptions) GetCcEnableArenas() bool`

GetCcEnableArenas returns the CcEnableArenas field if non-nil, zero value otherwise.

### GetCcEnableArenasOk

`func (o *FileOptions) GetCcEnableArenasOk() (*bool, bool)`

GetCcEnableArenasOk returns a tuple with the CcEnableArenas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEnableArenas

`func (o *FileOptions) SetCcEnableArenas(v bool)`

SetCcEnableArenas sets CcEnableArenas field to given value.

### HasCcEnableArenas

`func (o *FileOptions) HasCcEnableArenas() bool`

HasCcEnableArenas returns a boolean if a field has been set.

### GetObjcClassPrefix

`func (o *FileOptions) GetObjcClassPrefix() string`

GetObjcClassPrefix returns the ObjcClassPrefix field if non-nil, zero value otherwise.

### GetObjcClassPrefixOk

`func (o *FileOptions) GetObjcClassPrefixOk() (*string, bool)`

GetObjcClassPrefixOk returns a tuple with the ObjcClassPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjcClassPrefix

`func (o *FileOptions) SetObjcClassPrefix(v string)`

SetObjcClassPrefix sets ObjcClassPrefix field to given value.

### HasObjcClassPrefix

`func (o *FileOptions) HasObjcClassPrefix() bool`

HasObjcClassPrefix returns a boolean if a field has been set.

### GetCsharpNamespace

`func (o *FileOptions) GetCsharpNamespace() string`

GetCsharpNamespace returns the CsharpNamespace field if non-nil, zero value otherwise.

### GetCsharpNamespaceOk

`func (o *FileOptions) GetCsharpNamespaceOk() (*string, bool)`

GetCsharpNamespaceOk returns a tuple with the CsharpNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsharpNamespace

`func (o *FileOptions) SetCsharpNamespace(v string)`

SetCsharpNamespace sets CsharpNamespace field to given value.

### HasCsharpNamespace

`func (o *FileOptions) HasCsharpNamespace() bool`

HasCsharpNamespace returns a boolean if a field has been set.

### GetSwiftPrefix

`func (o *FileOptions) GetSwiftPrefix() string`

GetSwiftPrefix returns the SwiftPrefix field if non-nil, zero value otherwise.

### GetSwiftPrefixOk

`func (o *FileOptions) GetSwiftPrefixOk() (*string, bool)`

GetSwiftPrefixOk returns a tuple with the SwiftPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftPrefix

`func (o *FileOptions) SetSwiftPrefix(v string)`

SetSwiftPrefix sets SwiftPrefix field to given value.

### HasSwiftPrefix

`func (o *FileOptions) HasSwiftPrefix() bool`

HasSwiftPrefix returns a boolean if a field has been set.

### GetJavaGenericServices

`func (o *FileOptions) GetJavaGenericServices() bool`

GetJavaGenericServices returns the JavaGenericServices field if non-nil, zero value otherwise.

### GetJavaGenericServicesOk

`func (o *FileOptions) GetJavaGenericServicesOk() (*bool, bool)`

GetJavaGenericServicesOk returns a tuple with the JavaGenericServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaGenericServices

`func (o *FileOptions) SetJavaGenericServices(v bool)`

SetJavaGenericServices sets JavaGenericServices field to given value.

### HasJavaGenericServices

`func (o *FileOptions) HasJavaGenericServices() bool`

HasJavaGenericServices returns a boolean if a field has been set.

### GetPyGenericServices

`func (o *FileOptions) GetPyGenericServices() bool`

GetPyGenericServices returns the PyGenericServices field if non-nil, zero value otherwise.

### GetPyGenericServicesOk

`func (o *FileOptions) GetPyGenericServicesOk() (*bool, bool)`

GetPyGenericServicesOk returns a tuple with the PyGenericServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPyGenericServices

`func (o *FileOptions) SetPyGenericServices(v bool)`

SetPyGenericServices sets PyGenericServices field to given value.

### HasPyGenericServices

`func (o *FileOptions) HasPyGenericServices() bool`

HasPyGenericServices returns a boolean if a field has been set.

### GetPhpGenericServices

`func (o *FileOptions) GetPhpGenericServices() bool`

GetPhpGenericServices returns the PhpGenericServices field if non-nil, zero value otherwise.

### GetPhpGenericServicesOk

`func (o *FileOptions) GetPhpGenericServicesOk() (*bool, bool)`

GetPhpGenericServicesOk returns a tuple with the PhpGenericServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpGenericServices

`func (o *FileOptions) SetPhpGenericServices(v bool)`

SetPhpGenericServices sets PhpGenericServices field to given value.

### HasPhpGenericServices

`func (o *FileOptions) HasPhpGenericServices() bool`

HasPhpGenericServices returns a boolean if a field has been set.

### GetObjcClassPrefixBytes

`func (o *FileOptions) GetObjcClassPrefixBytes() ByteString`

GetObjcClassPrefixBytes returns the ObjcClassPrefixBytes field if non-nil, zero value otherwise.

### GetObjcClassPrefixBytesOk

`func (o *FileOptions) GetObjcClassPrefixBytesOk() (*ByteString, bool)`

GetObjcClassPrefixBytesOk returns a tuple with the ObjcClassPrefixBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjcClassPrefixBytes

`func (o *FileOptions) SetObjcClassPrefixBytes(v ByteString)`

SetObjcClassPrefixBytes sets ObjcClassPrefixBytes field to given value.

### HasObjcClassPrefixBytes

`func (o *FileOptions) HasObjcClassPrefixBytes() bool`

HasObjcClassPrefixBytes returns a boolean if a field has been set.

### GetCsharpNamespaceBytes

`func (o *FileOptions) GetCsharpNamespaceBytes() ByteString`

GetCsharpNamespaceBytes returns the CsharpNamespaceBytes field if non-nil, zero value otherwise.

### GetCsharpNamespaceBytesOk

`func (o *FileOptions) GetCsharpNamespaceBytesOk() (*ByteString, bool)`

GetCsharpNamespaceBytesOk returns a tuple with the CsharpNamespaceBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsharpNamespaceBytes

`func (o *FileOptions) SetCsharpNamespaceBytes(v ByteString)`

SetCsharpNamespaceBytes sets CsharpNamespaceBytes field to given value.

### HasCsharpNamespaceBytes

`func (o *FileOptions) HasCsharpNamespaceBytes() bool`

HasCsharpNamespaceBytes returns a boolean if a field has been set.

### GetSwiftPrefixBytes

`func (o *FileOptions) GetSwiftPrefixBytes() ByteString`

GetSwiftPrefixBytes returns the SwiftPrefixBytes field if non-nil, zero value otherwise.

### GetSwiftPrefixBytesOk

`func (o *FileOptions) GetSwiftPrefixBytesOk() (*ByteString, bool)`

GetSwiftPrefixBytesOk returns a tuple with the SwiftPrefixBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftPrefixBytes

`func (o *FileOptions) SetSwiftPrefixBytes(v ByteString)`

SetSwiftPrefixBytes sets SwiftPrefixBytes field to given value.

### HasSwiftPrefixBytes

`func (o *FileOptions) HasSwiftPrefixBytes() bool`

HasSwiftPrefixBytes returns a boolean if a field has been set.

### GetPhpClassPrefixBytes

`func (o *FileOptions) GetPhpClassPrefixBytes() ByteString`

GetPhpClassPrefixBytes returns the PhpClassPrefixBytes field if non-nil, zero value otherwise.

### GetPhpClassPrefixBytesOk

`func (o *FileOptions) GetPhpClassPrefixBytesOk() (*ByteString, bool)`

GetPhpClassPrefixBytesOk returns a tuple with the PhpClassPrefixBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpClassPrefixBytes

`func (o *FileOptions) SetPhpClassPrefixBytes(v ByteString)`

SetPhpClassPrefixBytes sets PhpClassPrefixBytes field to given value.

### HasPhpClassPrefixBytes

`func (o *FileOptions) HasPhpClassPrefixBytes() bool`

HasPhpClassPrefixBytes returns a boolean if a field has been set.

### GetPhpNamespaceBytes

`func (o *FileOptions) GetPhpNamespaceBytes() ByteString`

GetPhpNamespaceBytes returns the PhpNamespaceBytes field if non-nil, zero value otherwise.

### GetPhpNamespaceBytesOk

`func (o *FileOptions) GetPhpNamespaceBytesOk() (*ByteString, bool)`

GetPhpNamespaceBytesOk returns a tuple with the PhpNamespaceBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpNamespaceBytes

`func (o *FileOptions) SetPhpNamespaceBytes(v ByteString)`

SetPhpNamespaceBytes sets PhpNamespaceBytes field to given value.

### HasPhpNamespaceBytes

`func (o *FileOptions) HasPhpNamespaceBytes() bool`

HasPhpNamespaceBytes returns a boolean if a field has been set.

### GetPhpMetadataNamespace

`func (o *FileOptions) GetPhpMetadataNamespace() string`

GetPhpMetadataNamespace returns the PhpMetadataNamespace field if non-nil, zero value otherwise.

### GetPhpMetadataNamespaceOk

`func (o *FileOptions) GetPhpMetadataNamespaceOk() (*string, bool)`

GetPhpMetadataNamespaceOk returns a tuple with the PhpMetadataNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpMetadataNamespace

`func (o *FileOptions) SetPhpMetadataNamespace(v string)`

SetPhpMetadataNamespace sets PhpMetadataNamespace field to given value.

### HasPhpMetadataNamespace

`func (o *FileOptions) HasPhpMetadataNamespace() bool`

HasPhpMetadataNamespace returns a boolean if a field has been set.

### GetPhpMetadataNamespaceBytes

`func (o *FileOptions) GetPhpMetadataNamespaceBytes() ByteString`

GetPhpMetadataNamespaceBytes returns the PhpMetadataNamespaceBytes field if non-nil, zero value otherwise.

### GetPhpMetadataNamespaceBytesOk

`func (o *FileOptions) GetPhpMetadataNamespaceBytesOk() (*ByteString, bool)`

GetPhpMetadataNamespaceBytesOk returns a tuple with the PhpMetadataNamespaceBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpMetadataNamespaceBytes

`func (o *FileOptions) SetPhpMetadataNamespaceBytes(v ByteString)`

SetPhpMetadataNamespaceBytes sets PhpMetadataNamespaceBytes field to given value.

### HasPhpMetadataNamespaceBytes

`func (o *FileOptions) HasPhpMetadataNamespaceBytes() bool`

HasPhpMetadataNamespaceBytes returns a boolean if a field has been set.

### GetRubyPackageBytes

`func (o *FileOptions) GetRubyPackageBytes() ByteString`

GetRubyPackageBytes returns the RubyPackageBytes field if non-nil, zero value otherwise.

### GetRubyPackageBytesOk

`func (o *FileOptions) GetRubyPackageBytesOk() (*ByteString, bool)`

GetRubyPackageBytesOk returns a tuple with the RubyPackageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRubyPackageBytes

`func (o *FileOptions) SetRubyPackageBytes(v ByteString)`

SetRubyPackageBytes sets RubyPackageBytes field to given value.

### HasRubyPackageBytes

`func (o *FileOptions) HasRubyPackageBytes() bool`

HasRubyPackageBytes returns a boolean if a field has been set.

### GetPhpClassPrefix

`func (o *FileOptions) GetPhpClassPrefix() string`

GetPhpClassPrefix returns the PhpClassPrefix field if non-nil, zero value otherwise.

### GetPhpClassPrefixOk

`func (o *FileOptions) GetPhpClassPrefixOk() (*string, bool)`

GetPhpClassPrefixOk returns a tuple with the PhpClassPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpClassPrefix

`func (o *FileOptions) SetPhpClassPrefix(v string)`

SetPhpClassPrefix sets PhpClassPrefix field to given value.

### HasPhpClassPrefix

`func (o *FileOptions) HasPhpClassPrefix() bool`

HasPhpClassPrefix returns a boolean if a field has been set.

### GetPhpNamespace

`func (o *FileOptions) GetPhpNamespace() string`

GetPhpNamespace returns the PhpNamespace field if non-nil, zero value otherwise.

### GetPhpNamespaceOk

`func (o *FileOptions) GetPhpNamespaceOk() (*string, bool)`

GetPhpNamespaceOk returns a tuple with the PhpNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpNamespace

`func (o *FileOptions) SetPhpNamespace(v string)`

SetPhpNamespace sets PhpNamespace field to given value.

### HasPhpNamespace

`func (o *FileOptions) HasPhpNamespace() bool`

HasPhpNamespace returns a boolean if a field has been set.

### GetRubyPackage

`func (o *FileOptions) GetRubyPackage() string`

GetRubyPackage returns the RubyPackage field if non-nil, zero value otherwise.

### GetRubyPackageOk

`func (o *FileOptions) GetRubyPackageOk() (*string, bool)`

GetRubyPackageOk returns a tuple with the RubyPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRubyPackage

`func (o *FileOptions) SetRubyPackage(v string)`

SetRubyPackage sets RubyPackage field to given value.

### HasRubyPackage

`func (o *FileOptions) HasRubyPackage() bool`

HasRubyPackage returns a boolean if a field has been set.

### GetFeaturesOrBuilder

`func (o *FileOptions) GetFeaturesOrBuilder() FeatureSetOrBuilder`

GetFeaturesOrBuilder returns the FeaturesOrBuilder field if non-nil, zero value otherwise.

### GetFeaturesOrBuilderOk

`func (o *FileOptions) GetFeaturesOrBuilderOk() (*FeatureSetOrBuilder, bool)`

GetFeaturesOrBuilderOk returns a tuple with the FeaturesOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesOrBuilder

`func (o *FileOptions) SetFeaturesOrBuilder(v FeatureSetOrBuilder)`

SetFeaturesOrBuilder sets FeaturesOrBuilder field to given value.

### HasFeaturesOrBuilder

`func (o *FileOptions) HasFeaturesOrBuilder() bool`

HasFeaturesOrBuilder returns a boolean if a field has been set.

### GetUninterpretedOptionList

`func (o *FileOptions) GetUninterpretedOptionList() []map[string]interface{}`

GetUninterpretedOptionList returns the UninterpretedOptionList field if non-nil, zero value otherwise.

### GetUninterpretedOptionListOk

`func (o *FileOptions) GetUninterpretedOptionListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionListOk returns a tuple with the UninterpretedOptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionList

`func (o *FileOptions) SetUninterpretedOptionList(v []map[string]interface{})`

SetUninterpretedOptionList sets UninterpretedOptionList field to given value.

### HasUninterpretedOptionList

`func (o *FileOptions) HasUninterpretedOptionList() bool`

HasUninterpretedOptionList returns a boolean if a field has been set.

### GetUninterpretedOptionCount

`func (o *FileOptions) GetUninterpretedOptionCount() int32`

GetUninterpretedOptionCount returns the UninterpretedOptionCount field if non-nil, zero value otherwise.

### GetUninterpretedOptionCountOk

`func (o *FileOptions) GetUninterpretedOptionCountOk() (*int32, bool)`

GetUninterpretedOptionCountOk returns a tuple with the UninterpretedOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionCount

`func (o *FileOptions) SetUninterpretedOptionCount(v int32)`

SetUninterpretedOptionCount sets UninterpretedOptionCount field to given value.

### HasUninterpretedOptionCount

`func (o *FileOptions) HasUninterpretedOptionCount() bool`

HasUninterpretedOptionCount returns a boolean if a field has been set.

### GetInitialized

`func (o *FileOptions) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *FileOptions) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *FileOptions) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *FileOptions) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *FileOptions) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *FileOptions) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *FileOptions) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *FileOptions) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *FileOptions) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *FileOptions) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *FileOptions) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *FileOptions) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetAllFields

`func (o *FileOptions) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *FileOptions) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *FileOptions) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *FileOptions) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetAllFieldsRaw

`func (o *FileOptions) GetAllFieldsRaw() map[string]map[string]interface{}`

GetAllFieldsRaw returns the AllFieldsRaw field if non-nil, zero value otherwise.

### GetAllFieldsRawOk

`func (o *FileOptions) GetAllFieldsRawOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsRawOk returns a tuple with the AllFieldsRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFieldsRaw

`func (o *FileOptions) SetAllFieldsRaw(v map[string]map[string]interface{})`

SetAllFieldsRaw sets AllFieldsRaw field to given value.

### HasAllFieldsRaw

`func (o *FileOptions) HasAllFieldsRaw() bool`

HasAllFieldsRaw returns a boolean if a field has been set.

### GetMemoizedSerializedSize

`func (o *FileOptions) GetMemoizedSerializedSize() int32`

GetMemoizedSerializedSize returns the MemoizedSerializedSize field if non-nil, zero value otherwise.

### GetMemoizedSerializedSizeOk

`func (o *FileOptions) GetMemoizedSerializedSizeOk() (*int32, bool)`

GetMemoizedSerializedSizeOk returns a tuple with the MemoizedSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizedSerializedSize

`func (o *FileOptions) SetMemoizedSerializedSize(v int32)`

SetMemoizedSerializedSize sets MemoizedSerializedSize field to given value.

### HasMemoizedSerializedSize

`func (o *FileOptions) HasMemoizedSerializedSize() bool`

HasMemoizedSerializedSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


