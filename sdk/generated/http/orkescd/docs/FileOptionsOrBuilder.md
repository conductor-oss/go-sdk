# FileOptionsOrBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CcEnableArenas** | Pointer to **bool** |  | [optional] 
**CcGenericServices** | Pointer to **bool** |  | [optional] 
**CsharpNamespace** | Pointer to **string** |  | [optional] 
**CsharpNamespaceBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**DefaultInstanceForType** | Pointer to **map[string]interface{}** | Simplified schema for Message (original had circular references) | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**DescriptorForType** | Pointer to **map[string]interface{}** | Simplified schema for Descriptor (original had circular references) | [optional] 
**Features** | Pointer to [**FeatureSet**](FeatureSet.md) |  | [optional] 
**FeaturesOrBuilder** | Pointer to [**FeatureSetOrBuilder**](FeatureSetOrBuilder.md) |  | [optional] 
**GoPackage** | Pointer to **string** |  | [optional] 
**GoPackageBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**InitializationErrorString** | Pointer to **string** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**JavaGenerateEqualsAndHash** | Pointer to **bool** |  | [optional] 
**JavaGenericServices** | Pointer to **bool** |  | [optional] 
**JavaMultipleFiles** | Pointer to **bool** |  | [optional] 
**JavaOuterClassname** | Pointer to **string** |  | [optional] 
**JavaOuterClassnameBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**JavaPackage** | Pointer to **string** |  | [optional] 
**JavaPackageBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**JavaStringCheckUtf8** | Pointer to **bool** |  | [optional] 
**ObjcClassPrefix** | Pointer to **string** |  | [optional] 
**ObjcClassPrefixBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**OptimizeFor** | Pointer to **string** |  | [optional] 
**PhpClassPrefix** | Pointer to **string** |  | [optional] 
**PhpClassPrefixBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**PhpGenericServices** | Pointer to **bool** |  | [optional] 
**PhpMetadataNamespace** | Pointer to **string** |  | [optional] 
**PhpMetadataNamespaceBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**PhpNamespace** | Pointer to **string** |  | [optional] 
**PhpNamespaceBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**PyGenericServices** | Pointer to **bool** |  | [optional] 
**RubyPackage** | Pointer to **string** |  | [optional] 
**RubyPackageBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**SwiftPrefix** | Pointer to **string** |  | [optional] 
**SwiftPrefixBytes** | Pointer to [**ByteString**](ByteString.md) |  | [optional] 
**UninterpretedOptionCount** | Pointer to **int32** |  | [optional] 
**UninterpretedOptionList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UninterpretedOptionOrBuilderList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UnknownFields** | Pointer to [**UnknownFieldSet**](UnknownFieldSet.md) |  | [optional] 

## Methods

### NewFileOptionsOrBuilder

`func NewFileOptionsOrBuilder() *FileOptionsOrBuilder`

NewFileOptionsOrBuilder instantiates a new FileOptionsOrBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileOptionsOrBuilderWithDefaults

`func NewFileOptionsOrBuilderWithDefaults() *FileOptionsOrBuilder`

NewFileOptionsOrBuilderWithDefaults instantiates a new FileOptionsOrBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *FileOptionsOrBuilder) GetAllFields() map[string]map[string]interface{}`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *FileOptionsOrBuilder) GetAllFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *FileOptionsOrBuilder) SetAllFields(v map[string]map[string]interface{})`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *FileOptionsOrBuilder) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetCcEnableArenas

`func (o *FileOptionsOrBuilder) GetCcEnableArenas() bool`

GetCcEnableArenas returns the CcEnableArenas field if non-nil, zero value otherwise.

### GetCcEnableArenasOk

`func (o *FileOptionsOrBuilder) GetCcEnableArenasOk() (*bool, bool)`

GetCcEnableArenasOk returns a tuple with the CcEnableArenas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEnableArenas

`func (o *FileOptionsOrBuilder) SetCcEnableArenas(v bool)`

SetCcEnableArenas sets CcEnableArenas field to given value.

### HasCcEnableArenas

`func (o *FileOptionsOrBuilder) HasCcEnableArenas() bool`

HasCcEnableArenas returns a boolean if a field has been set.

### GetCcGenericServices

`func (o *FileOptionsOrBuilder) GetCcGenericServices() bool`

GetCcGenericServices returns the CcGenericServices field if non-nil, zero value otherwise.

### GetCcGenericServicesOk

`func (o *FileOptionsOrBuilder) GetCcGenericServicesOk() (*bool, bool)`

GetCcGenericServicesOk returns a tuple with the CcGenericServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcGenericServices

`func (o *FileOptionsOrBuilder) SetCcGenericServices(v bool)`

SetCcGenericServices sets CcGenericServices field to given value.

### HasCcGenericServices

`func (o *FileOptionsOrBuilder) HasCcGenericServices() bool`

HasCcGenericServices returns a boolean if a field has been set.

### GetCsharpNamespace

`func (o *FileOptionsOrBuilder) GetCsharpNamespace() string`

GetCsharpNamespace returns the CsharpNamespace field if non-nil, zero value otherwise.

### GetCsharpNamespaceOk

`func (o *FileOptionsOrBuilder) GetCsharpNamespaceOk() (*string, bool)`

GetCsharpNamespaceOk returns a tuple with the CsharpNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsharpNamespace

`func (o *FileOptionsOrBuilder) SetCsharpNamespace(v string)`

SetCsharpNamespace sets CsharpNamespace field to given value.

### HasCsharpNamespace

`func (o *FileOptionsOrBuilder) HasCsharpNamespace() bool`

HasCsharpNamespace returns a boolean if a field has been set.

### GetCsharpNamespaceBytes

`func (o *FileOptionsOrBuilder) GetCsharpNamespaceBytes() ByteString`

GetCsharpNamespaceBytes returns the CsharpNamespaceBytes field if non-nil, zero value otherwise.

### GetCsharpNamespaceBytesOk

`func (o *FileOptionsOrBuilder) GetCsharpNamespaceBytesOk() (*ByteString, bool)`

GetCsharpNamespaceBytesOk returns a tuple with the CsharpNamespaceBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsharpNamespaceBytes

`func (o *FileOptionsOrBuilder) SetCsharpNamespaceBytes(v ByteString)`

SetCsharpNamespaceBytes sets CsharpNamespaceBytes field to given value.

### HasCsharpNamespaceBytes

`func (o *FileOptionsOrBuilder) HasCsharpNamespaceBytes() bool`

HasCsharpNamespaceBytes returns a boolean if a field has been set.

### GetDefaultInstanceForType

`func (o *FileOptionsOrBuilder) GetDefaultInstanceForType() map[string]interface{}`

GetDefaultInstanceForType returns the DefaultInstanceForType field if non-nil, zero value otherwise.

### GetDefaultInstanceForTypeOk

`func (o *FileOptionsOrBuilder) GetDefaultInstanceForTypeOk() (*map[string]interface{}, bool)`

GetDefaultInstanceForTypeOk returns a tuple with the DefaultInstanceForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstanceForType

`func (o *FileOptionsOrBuilder) SetDefaultInstanceForType(v map[string]interface{})`

SetDefaultInstanceForType sets DefaultInstanceForType field to given value.

### HasDefaultInstanceForType

`func (o *FileOptionsOrBuilder) HasDefaultInstanceForType() bool`

HasDefaultInstanceForType returns a boolean if a field has been set.

### GetDeprecated

`func (o *FileOptionsOrBuilder) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *FileOptionsOrBuilder) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *FileOptionsOrBuilder) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *FileOptionsOrBuilder) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescriptorForType

`func (o *FileOptionsOrBuilder) GetDescriptorForType() map[string]interface{}`

GetDescriptorForType returns the DescriptorForType field if non-nil, zero value otherwise.

### GetDescriptorForTypeOk

`func (o *FileOptionsOrBuilder) GetDescriptorForTypeOk() (*map[string]interface{}, bool)`

GetDescriptorForTypeOk returns a tuple with the DescriptorForType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptorForType

`func (o *FileOptionsOrBuilder) SetDescriptorForType(v map[string]interface{})`

SetDescriptorForType sets DescriptorForType field to given value.

### HasDescriptorForType

`func (o *FileOptionsOrBuilder) HasDescriptorForType() bool`

HasDescriptorForType returns a boolean if a field has been set.

### GetFeatures

`func (o *FileOptionsOrBuilder) GetFeatures() FeatureSet`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FileOptionsOrBuilder) GetFeaturesOk() (*FeatureSet, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FileOptionsOrBuilder) SetFeatures(v FeatureSet)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FileOptionsOrBuilder) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesOrBuilder

`func (o *FileOptionsOrBuilder) GetFeaturesOrBuilder() FeatureSetOrBuilder`

GetFeaturesOrBuilder returns the FeaturesOrBuilder field if non-nil, zero value otherwise.

### GetFeaturesOrBuilderOk

`func (o *FileOptionsOrBuilder) GetFeaturesOrBuilderOk() (*FeatureSetOrBuilder, bool)`

GetFeaturesOrBuilderOk returns a tuple with the FeaturesOrBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesOrBuilder

`func (o *FileOptionsOrBuilder) SetFeaturesOrBuilder(v FeatureSetOrBuilder)`

SetFeaturesOrBuilder sets FeaturesOrBuilder field to given value.

### HasFeaturesOrBuilder

`func (o *FileOptionsOrBuilder) HasFeaturesOrBuilder() bool`

HasFeaturesOrBuilder returns a boolean if a field has been set.

### GetGoPackage

`func (o *FileOptionsOrBuilder) GetGoPackage() string`

GetGoPackage returns the GoPackage field if non-nil, zero value otherwise.

### GetGoPackageOk

`func (o *FileOptionsOrBuilder) GetGoPackageOk() (*string, bool)`

GetGoPackageOk returns a tuple with the GoPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoPackage

`func (o *FileOptionsOrBuilder) SetGoPackage(v string)`

SetGoPackage sets GoPackage field to given value.

### HasGoPackage

`func (o *FileOptionsOrBuilder) HasGoPackage() bool`

HasGoPackage returns a boolean if a field has been set.

### GetGoPackageBytes

`func (o *FileOptionsOrBuilder) GetGoPackageBytes() ByteString`

GetGoPackageBytes returns the GoPackageBytes field if non-nil, zero value otherwise.

### GetGoPackageBytesOk

`func (o *FileOptionsOrBuilder) GetGoPackageBytesOk() (*ByteString, bool)`

GetGoPackageBytesOk returns a tuple with the GoPackageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoPackageBytes

`func (o *FileOptionsOrBuilder) SetGoPackageBytes(v ByteString)`

SetGoPackageBytes sets GoPackageBytes field to given value.

### HasGoPackageBytes

`func (o *FileOptionsOrBuilder) HasGoPackageBytes() bool`

HasGoPackageBytes returns a boolean if a field has been set.

### GetInitializationErrorString

`func (o *FileOptionsOrBuilder) GetInitializationErrorString() string`

GetInitializationErrorString returns the InitializationErrorString field if non-nil, zero value otherwise.

### GetInitializationErrorStringOk

`func (o *FileOptionsOrBuilder) GetInitializationErrorStringOk() (*string, bool)`

GetInitializationErrorStringOk returns a tuple with the InitializationErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationErrorString

`func (o *FileOptionsOrBuilder) SetInitializationErrorString(v string)`

SetInitializationErrorString sets InitializationErrorString field to given value.

### HasInitializationErrorString

`func (o *FileOptionsOrBuilder) HasInitializationErrorString() bool`

HasInitializationErrorString returns a boolean if a field has been set.

### GetInitialized

`func (o *FileOptionsOrBuilder) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *FileOptionsOrBuilder) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *FileOptionsOrBuilder) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *FileOptionsOrBuilder) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetJavaGenerateEqualsAndHash

`func (o *FileOptionsOrBuilder) GetJavaGenerateEqualsAndHash() bool`

GetJavaGenerateEqualsAndHash returns the JavaGenerateEqualsAndHash field if non-nil, zero value otherwise.

### GetJavaGenerateEqualsAndHashOk

`func (o *FileOptionsOrBuilder) GetJavaGenerateEqualsAndHashOk() (*bool, bool)`

GetJavaGenerateEqualsAndHashOk returns a tuple with the JavaGenerateEqualsAndHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaGenerateEqualsAndHash

`func (o *FileOptionsOrBuilder) SetJavaGenerateEqualsAndHash(v bool)`

SetJavaGenerateEqualsAndHash sets JavaGenerateEqualsAndHash field to given value.

### HasJavaGenerateEqualsAndHash

`func (o *FileOptionsOrBuilder) HasJavaGenerateEqualsAndHash() bool`

HasJavaGenerateEqualsAndHash returns a boolean if a field has been set.

### GetJavaGenericServices

`func (o *FileOptionsOrBuilder) GetJavaGenericServices() bool`

GetJavaGenericServices returns the JavaGenericServices field if non-nil, zero value otherwise.

### GetJavaGenericServicesOk

`func (o *FileOptionsOrBuilder) GetJavaGenericServicesOk() (*bool, bool)`

GetJavaGenericServicesOk returns a tuple with the JavaGenericServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaGenericServices

`func (o *FileOptionsOrBuilder) SetJavaGenericServices(v bool)`

SetJavaGenericServices sets JavaGenericServices field to given value.

### HasJavaGenericServices

`func (o *FileOptionsOrBuilder) HasJavaGenericServices() bool`

HasJavaGenericServices returns a boolean if a field has been set.

### GetJavaMultipleFiles

`func (o *FileOptionsOrBuilder) GetJavaMultipleFiles() bool`

GetJavaMultipleFiles returns the JavaMultipleFiles field if non-nil, zero value otherwise.

### GetJavaMultipleFilesOk

`func (o *FileOptionsOrBuilder) GetJavaMultipleFilesOk() (*bool, bool)`

GetJavaMultipleFilesOk returns a tuple with the JavaMultipleFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaMultipleFiles

`func (o *FileOptionsOrBuilder) SetJavaMultipleFiles(v bool)`

SetJavaMultipleFiles sets JavaMultipleFiles field to given value.

### HasJavaMultipleFiles

`func (o *FileOptionsOrBuilder) HasJavaMultipleFiles() bool`

HasJavaMultipleFiles returns a boolean if a field has been set.

### GetJavaOuterClassname

`func (o *FileOptionsOrBuilder) GetJavaOuterClassname() string`

GetJavaOuterClassname returns the JavaOuterClassname field if non-nil, zero value otherwise.

### GetJavaOuterClassnameOk

`func (o *FileOptionsOrBuilder) GetJavaOuterClassnameOk() (*string, bool)`

GetJavaOuterClassnameOk returns a tuple with the JavaOuterClassname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaOuterClassname

`func (o *FileOptionsOrBuilder) SetJavaOuterClassname(v string)`

SetJavaOuterClassname sets JavaOuterClassname field to given value.

### HasJavaOuterClassname

`func (o *FileOptionsOrBuilder) HasJavaOuterClassname() bool`

HasJavaOuterClassname returns a boolean if a field has been set.

### GetJavaOuterClassnameBytes

`func (o *FileOptionsOrBuilder) GetJavaOuterClassnameBytes() ByteString`

GetJavaOuterClassnameBytes returns the JavaOuterClassnameBytes field if non-nil, zero value otherwise.

### GetJavaOuterClassnameBytesOk

`func (o *FileOptionsOrBuilder) GetJavaOuterClassnameBytesOk() (*ByteString, bool)`

GetJavaOuterClassnameBytesOk returns a tuple with the JavaOuterClassnameBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaOuterClassnameBytes

`func (o *FileOptionsOrBuilder) SetJavaOuterClassnameBytes(v ByteString)`

SetJavaOuterClassnameBytes sets JavaOuterClassnameBytes field to given value.

### HasJavaOuterClassnameBytes

`func (o *FileOptionsOrBuilder) HasJavaOuterClassnameBytes() bool`

HasJavaOuterClassnameBytes returns a boolean if a field has been set.

### GetJavaPackage

`func (o *FileOptionsOrBuilder) GetJavaPackage() string`

GetJavaPackage returns the JavaPackage field if non-nil, zero value otherwise.

### GetJavaPackageOk

`func (o *FileOptionsOrBuilder) GetJavaPackageOk() (*string, bool)`

GetJavaPackageOk returns a tuple with the JavaPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaPackage

`func (o *FileOptionsOrBuilder) SetJavaPackage(v string)`

SetJavaPackage sets JavaPackage field to given value.

### HasJavaPackage

`func (o *FileOptionsOrBuilder) HasJavaPackage() bool`

HasJavaPackage returns a boolean if a field has been set.

### GetJavaPackageBytes

`func (o *FileOptionsOrBuilder) GetJavaPackageBytes() ByteString`

GetJavaPackageBytes returns the JavaPackageBytes field if non-nil, zero value otherwise.

### GetJavaPackageBytesOk

`func (o *FileOptionsOrBuilder) GetJavaPackageBytesOk() (*ByteString, bool)`

GetJavaPackageBytesOk returns a tuple with the JavaPackageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaPackageBytes

`func (o *FileOptionsOrBuilder) SetJavaPackageBytes(v ByteString)`

SetJavaPackageBytes sets JavaPackageBytes field to given value.

### HasJavaPackageBytes

`func (o *FileOptionsOrBuilder) HasJavaPackageBytes() bool`

HasJavaPackageBytes returns a boolean if a field has been set.

### GetJavaStringCheckUtf8

`func (o *FileOptionsOrBuilder) GetJavaStringCheckUtf8() bool`

GetJavaStringCheckUtf8 returns the JavaStringCheckUtf8 field if non-nil, zero value otherwise.

### GetJavaStringCheckUtf8Ok

`func (o *FileOptionsOrBuilder) GetJavaStringCheckUtf8Ok() (*bool, bool)`

GetJavaStringCheckUtf8Ok returns a tuple with the JavaStringCheckUtf8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaStringCheckUtf8

`func (o *FileOptionsOrBuilder) SetJavaStringCheckUtf8(v bool)`

SetJavaStringCheckUtf8 sets JavaStringCheckUtf8 field to given value.

### HasJavaStringCheckUtf8

`func (o *FileOptionsOrBuilder) HasJavaStringCheckUtf8() bool`

HasJavaStringCheckUtf8 returns a boolean if a field has been set.

### GetObjcClassPrefix

`func (o *FileOptionsOrBuilder) GetObjcClassPrefix() string`

GetObjcClassPrefix returns the ObjcClassPrefix field if non-nil, zero value otherwise.

### GetObjcClassPrefixOk

`func (o *FileOptionsOrBuilder) GetObjcClassPrefixOk() (*string, bool)`

GetObjcClassPrefixOk returns a tuple with the ObjcClassPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjcClassPrefix

`func (o *FileOptionsOrBuilder) SetObjcClassPrefix(v string)`

SetObjcClassPrefix sets ObjcClassPrefix field to given value.

### HasObjcClassPrefix

`func (o *FileOptionsOrBuilder) HasObjcClassPrefix() bool`

HasObjcClassPrefix returns a boolean if a field has been set.

### GetObjcClassPrefixBytes

`func (o *FileOptionsOrBuilder) GetObjcClassPrefixBytes() ByteString`

GetObjcClassPrefixBytes returns the ObjcClassPrefixBytes field if non-nil, zero value otherwise.

### GetObjcClassPrefixBytesOk

`func (o *FileOptionsOrBuilder) GetObjcClassPrefixBytesOk() (*ByteString, bool)`

GetObjcClassPrefixBytesOk returns a tuple with the ObjcClassPrefixBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjcClassPrefixBytes

`func (o *FileOptionsOrBuilder) SetObjcClassPrefixBytes(v ByteString)`

SetObjcClassPrefixBytes sets ObjcClassPrefixBytes field to given value.

### HasObjcClassPrefixBytes

`func (o *FileOptionsOrBuilder) HasObjcClassPrefixBytes() bool`

HasObjcClassPrefixBytes returns a boolean if a field has been set.

### GetOptimizeFor

`func (o *FileOptionsOrBuilder) GetOptimizeFor() string`

GetOptimizeFor returns the OptimizeFor field if non-nil, zero value otherwise.

### GetOptimizeForOk

`func (o *FileOptionsOrBuilder) GetOptimizeForOk() (*string, bool)`

GetOptimizeForOk returns a tuple with the OptimizeFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeFor

`func (o *FileOptionsOrBuilder) SetOptimizeFor(v string)`

SetOptimizeFor sets OptimizeFor field to given value.

### HasOptimizeFor

`func (o *FileOptionsOrBuilder) HasOptimizeFor() bool`

HasOptimizeFor returns a boolean if a field has been set.

### GetPhpClassPrefix

`func (o *FileOptionsOrBuilder) GetPhpClassPrefix() string`

GetPhpClassPrefix returns the PhpClassPrefix field if non-nil, zero value otherwise.

### GetPhpClassPrefixOk

`func (o *FileOptionsOrBuilder) GetPhpClassPrefixOk() (*string, bool)`

GetPhpClassPrefixOk returns a tuple with the PhpClassPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpClassPrefix

`func (o *FileOptionsOrBuilder) SetPhpClassPrefix(v string)`

SetPhpClassPrefix sets PhpClassPrefix field to given value.

### HasPhpClassPrefix

`func (o *FileOptionsOrBuilder) HasPhpClassPrefix() bool`

HasPhpClassPrefix returns a boolean if a field has been set.

### GetPhpClassPrefixBytes

`func (o *FileOptionsOrBuilder) GetPhpClassPrefixBytes() ByteString`

GetPhpClassPrefixBytes returns the PhpClassPrefixBytes field if non-nil, zero value otherwise.

### GetPhpClassPrefixBytesOk

`func (o *FileOptionsOrBuilder) GetPhpClassPrefixBytesOk() (*ByteString, bool)`

GetPhpClassPrefixBytesOk returns a tuple with the PhpClassPrefixBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpClassPrefixBytes

`func (o *FileOptionsOrBuilder) SetPhpClassPrefixBytes(v ByteString)`

SetPhpClassPrefixBytes sets PhpClassPrefixBytes field to given value.

### HasPhpClassPrefixBytes

`func (o *FileOptionsOrBuilder) HasPhpClassPrefixBytes() bool`

HasPhpClassPrefixBytes returns a boolean if a field has been set.

### GetPhpGenericServices

`func (o *FileOptionsOrBuilder) GetPhpGenericServices() bool`

GetPhpGenericServices returns the PhpGenericServices field if non-nil, zero value otherwise.

### GetPhpGenericServicesOk

`func (o *FileOptionsOrBuilder) GetPhpGenericServicesOk() (*bool, bool)`

GetPhpGenericServicesOk returns a tuple with the PhpGenericServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpGenericServices

`func (o *FileOptionsOrBuilder) SetPhpGenericServices(v bool)`

SetPhpGenericServices sets PhpGenericServices field to given value.

### HasPhpGenericServices

`func (o *FileOptionsOrBuilder) HasPhpGenericServices() bool`

HasPhpGenericServices returns a boolean if a field has been set.

### GetPhpMetadataNamespace

`func (o *FileOptionsOrBuilder) GetPhpMetadataNamespace() string`

GetPhpMetadataNamespace returns the PhpMetadataNamespace field if non-nil, zero value otherwise.

### GetPhpMetadataNamespaceOk

`func (o *FileOptionsOrBuilder) GetPhpMetadataNamespaceOk() (*string, bool)`

GetPhpMetadataNamespaceOk returns a tuple with the PhpMetadataNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpMetadataNamespace

`func (o *FileOptionsOrBuilder) SetPhpMetadataNamespace(v string)`

SetPhpMetadataNamespace sets PhpMetadataNamespace field to given value.

### HasPhpMetadataNamespace

`func (o *FileOptionsOrBuilder) HasPhpMetadataNamespace() bool`

HasPhpMetadataNamespace returns a boolean if a field has been set.

### GetPhpMetadataNamespaceBytes

`func (o *FileOptionsOrBuilder) GetPhpMetadataNamespaceBytes() ByteString`

GetPhpMetadataNamespaceBytes returns the PhpMetadataNamespaceBytes field if non-nil, zero value otherwise.

### GetPhpMetadataNamespaceBytesOk

`func (o *FileOptionsOrBuilder) GetPhpMetadataNamespaceBytesOk() (*ByteString, bool)`

GetPhpMetadataNamespaceBytesOk returns a tuple with the PhpMetadataNamespaceBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpMetadataNamespaceBytes

`func (o *FileOptionsOrBuilder) SetPhpMetadataNamespaceBytes(v ByteString)`

SetPhpMetadataNamespaceBytes sets PhpMetadataNamespaceBytes field to given value.

### HasPhpMetadataNamespaceBytes

`func (o *FileOptionsOrBuilder) HasPhpMetadataNamespaceBytes() bool`

HasPhpMetadataNamespaceBytes returns a boolean if a field has been set.

### GetPhpNamespace

`func (o *FileOptionsOrBuilder) GetPhpNamespace() string`

GetPhpNamespace returns the PhpNamespace field if non-nil, zero value otherwise.

### GetPhpNamespaceOk

`func (o *FileOptionsOrBuilder) GetPhpNamespaceOk() (*string, bool)`

GetPhpNamespaceOk returns a tuple with the PhpNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpNamespace

`func (o *FileOptionsOrBuilder) SetPhpNamespace(v string)`

SetPhpNamespace sets PhpNamespace field to given value.

### HasPhpNamespace

`func (o *FileOptionsOrBuilder) HasPhpNamespace() bool`

HasPhpNamespace returns a boolean if a field has been set.

### GetPhpNamespaceBytes

`func (o *FileOptionsOrBuilder) GetPhpNamespaceBytes() ByteString`

GetPhpNamespaceBytes returns the PhpNamespaceBytes field if non-nil, zero value otherwise.

### GetPhpNamespaceBytesOk

`func (o *FileOptionsOrBuilder) GetPhpNamespaceBytesOk() (*ByteString, bool)`

GetPhpNamespaceBytesOk returns a tuple with the PhpNamespaceBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhpNamespaceBytes

`func (o *FileOptionsOrBuilder) SetPhpNamespaceBytes(v ByteString)`

SetPhpNamespaceBytes sets PhpNamespaceBytes field to given value.

### HasPhpNamespaceBytes

`func (o *FileOptionsOrBuilder) HasPhpNamespaceBytes() bool`

HasPhpNamespaceBytes returns a boolean if a field has been set.

### GetPyGenericServices

`func (o *FileOptionsOrBuilder) GetPyGenericServices() bool`

GetPyGenericServices returns the PyGenericServices field if non-nil, zero value otherwise.

### GetPyGenericServicesOk

`func (o *FileOptionsOrBuilder) GetPyGenericServicesOk() (*bool, bool)`

GetPyGenericServicesOk returns a tuple with the PyGenericServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPyGenericServices

`func (o *FileOptionsOrBuilder) SetPyGenericServices(v bool)`

SetPyGenericServices sets PyGenericServices field to given value.

### HasPyGenericServices

`func (o *FileOptionsOrBuilder) HasPyGenericServices() bool`

HasPyGenericServices returns a boolean if a field has been set.

### GetRubyPackage

`func (o *FileOptionsOrBuilder) GetRubyPackage() string`

GetRubyPackage returns the RubyPackage field if non-nil, zero value otherwise.

### GetRubyPackageOk

`func (o *FileOptionsOrBuilder) GetRubyPackageOk() (*string, bool)`

GetRubyPackageOk returns a tuple with the RubyPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRubyPackage

`func (o *FileOptionsOrBuilder) SetRubyPackage(v string)`

SetRubyPackage sets RubyPackage field to given value.

### HasRubyPackage

`func (o *FileOptionsOrBuilder) HasRubyPackage() bool`

HasRubyPackage returns a boolean if a field has been set.

### GetRubyPackageBytes

`func (o *FileOptionsOrBuilder) GetRubyPackageBytes() ByteString`

GetRubyPackageBytes returns the RubyPackageBytes field if non-nil, zero value otherwise.

### GetRubyPackageBytesOk

`func (o *FileOptionsOrBuilder) GetRubyPackageBytesOk() (*ByteString, bool)`

GetRubyPackageBytesOk returns a tuple with the RubyPackageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRubyPackageBytes

`func (o *FileOptionsOrBuilder) SetRubyPackageBytes(v ByteString)`

SetRubyPackageBytes sets RubyPackageBytes field to given value.

### HasRubyPackageBytes

`func (o *FileOptionsOrBuilder) HasRubyPackageBytes() bool`

HasRubyPackageBytes returns a boolean if a field has been set.

### GetSwiftPrefix

`func (o *FileOptionsOrBuilder) GetSwiftPrefix() string`

GetSwiftPrefix returns the SwiftPrefix field if non-nil, zero value otherwise.

### GetSwiftPrefixOk

`func (o *FileOptionsOrBuilder) GetSwiftPrefixOk() (*string, bool)`

GetSwiftPrefixOk returns a tuple with the SwiftPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftPrefix

`func (o *FileOptionsOrBuilder) SetSwiftPrefix(v string)`

SetSwiftPrefix sets SwiftPrefix field to given value.

### HasSwiftPrefix

`func (o *FileOptionsOrBuilder) HasSwiftPrefix() bool`

HasSwiftPrefix returns a boolean if a field has been set.

### GetSwiftPrefixBytes

`func (o *FileOptionsOrBuilder) GetSwiftPrefixBytes() ByteString`

GetSwiftPrefixBytes returns the SwiftPrefixBytes field if non-nil, zero value otherwise.

### GetSwiftPrefixBytesOk

`func (o *FileOptionsOrBuilder) GetSwiftPrefixBytesOk() (*ByteString, bool)`

GetSwiftPrefixBytesOk returns a tuple with the SwiftPrefixBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftPrefixBytes

`func (o *FileOptionsOrBuilder) SetSwiftPrefixBytes(v ByteString)`

SetSwiftPrefixBytes sets SwiftPrefixBytes field to given value.

### HasSwiftPrefixBytes

`func (o *FileOptionsOrBuilder) HasSwiftPrefixBytes() bool`

HasSwiftPrefixBytes returns a boolean if a field has been set.

### GetUninterpretedOptionCount

`func (o *FileOptionsOrBuilder) GetUninterpretedOptionCount() int32`

GetUninterpretedOptionCount returns the UninterpretedOptionCount field if non-nil, zero value otherwise.

### GetUninterpretedOptionCountOk

`func (o *FileOptionsOrBuilder) GetUninterpretedOptionCountOk() (*int32, bool)`

GetUninterpretedOptionCountOk returns a tuple with the UninterpretedOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionCount

`func (o *FileOptionsOrBuilder) SetUninterpretedOptionCount(v int32)`

SetUninterpretedOptionCount sets UninterpretedOptionCount field to given value.

### HasUninterpretedOptionCount

`func (o *FileOptionsOrBuilder) HasUninterpretedOptionCount() bool`

HasUninterpretedOptionCount returns a boolean if a field has been set.

### GetUninterpretedOptionList

`func (o *FileOptionsOrBuilder) GetUninterpretedOptionList() []map[string]interface{}`

GetUninterpretedOptionList returns the UninterpretedOptionList field if non-nil, zero value otherwise.

### GetUninterpretedOptionListOk

`func (o *FileOptionsOrBuilder) GetUninterpretedOptionListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionListOk returns a tuple with the UninterpretedOptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionList

`func (o *FileOptionsOrBuilder) SetUninterpretedOptionList(v []map[string]interface{})`

SetUninterpretedOptionList sets UninterpretedOptionList field to given value.

### HasUninterpretedOptionList

`func (o *FileOptionsOrBuilder) HasUninterpretedOptionList() bool`

HasUninterpretedOptionList returns a boolean if a field has been set.

### GetUninterpretedOptionOrBuilderList

`func (o *FileOptionsOrBuilder) GetUninterpretedOptionOrBuilderList() []map[string]interface{}`

GetUninterpretedOptionOrBuilderList returns the UninterpretedOptionOrBuilderList field if non-nil, zero value otherwise.

### GetUninterpretedOptionOrBuilderListOk

`func (o *FileOptionsOrBuilder) GetUninterpretedOptionOrBuilderListOk() (*[]map[string]interface{}, bool)`

GetUninterpretedOptionOrBuilderListOk returns a tuple with the UninterpretedOptionOrBuilderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninterpretedOptionOrBuilderList

`func (o *FileOptionsOrBuilder) SetUninterpretedOptionOrBuilderList(v []map[string]interface{})`

SetUninterpretedOptionOrBuilderList sets UninterpretedOptionOrBuilderList field to given value.

### HasUninterpretedOptionOrBuilderList

`func (o *FileOptionsOrBuilder) HasUninterpretedOptionOrBuilderList() bool`

HasUninterpretedOptionOrBuilderList returns a boolean if a field has been set.

### GetUnknownFields

`func (o *FileOptionsOrBuilder) GetUnknownFields() UnknownFieldSet`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *FileOptionsOrBuilder) GetUnknownFieldsOk() (*UnknownFieldSet, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *FileOptionsOrBuilder) SetUnknownFields(v UnknownFieldSet)`

SetUnknownFields sets UnknownFields field to given value.

### HasUnknownFields

`func (o *FileOptionsOrBuilder) HasUnknownFields() bool`

HasUnknownFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


