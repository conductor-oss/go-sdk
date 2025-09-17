# HumanTaskTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**JsonSchema** | **map[string]map[string]interface{}** |  | 
**Name** | **string** |  | 
**OwnerApp** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**TemplateUI** | **map[string]map[string]interface{}** |  | 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Version** | **int32** |  | 

## Methods

### NewHumanTaskTemplate

`func NewHumanTaskTemplate(jsonSchema map[string]map[string]interface{}, name string, templateUI map[string]map[string]interface{}, version int32, ) *HumanTaskTemplate`

NewHumanTaskTemplate instantiates a new HumanTaskTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHumanTaskTemplateWithDefaults

`func NewHumanTaskTemplateWithDefaults() *HumanTaskTemplate`

NewHumanTaskTemplateWithDefaults instantiates a new HumanTaskTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *HumanTaskTemplate) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HumanTaskTemplate) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HumanTaskTemplate) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HumanTaskTemplate) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *HumanTaskTemplate) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *HumanTaskTemplate) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *HumanTaskTemplate) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *HumanTaskTemplate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetJsonSchema

`func (o *HumanTaskTemplate) GetJsonSchema() map[string]map[string]interface{}`

GetJsonSchema returns the JsonSchema field if non-nil, zero value otherwise.

### GetJsonSchemaOk

`func (o *HumanTaskTemplate) GetJsonSchemaOk() (*map[string]map[string]interface{}, bool)`

GetJsonSchemaOk returns a tuple with the JsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchema

`func (o *HumanTaskTemplate) SetJsonSchema(v map[string]map[string]interface{})`

SetJsonSchema sets JsonSchema field to given value.


### GetName

`func (o *HumanTaskTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HumanTaskTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HumanTaskTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerApp

`func (o *HumanTaskTemplate) GetOwnerApp() string`

GetOwnerApp returns the OwnerApp field if non-nil, zero value otherwise.

### GetOwnerAppOk

`func (o *HumanTaskTemplate) GetOwnerAppOk() (*string, bool)`

GetOwnerAppOk returns a tuple with the OwnerApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerApp

`func (o *HumanTaskTemplate) SetOwnerApp(v string)`

SetOwnerApp sets OwnerApp field to given value.

### HasOwnerApp

`func (o *HumanTaskTemplate) HasOwnerApp() bool`

HasOwnerApp returns a boolean if a field has been set.

### GetTags

`func (o *HumanTaskTemplate) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HumanTaskTemplate) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HumanTaskTemplate) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HumanTaskTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplateUI

`func (o *HumanTaskTemplate) GetTemplateUI() map[string]map[string]interface{}`

GetTemplateUI returns the TemplateUI field if non-nil, zero value otherwise.

### GetTemplateUIOk

`func (o *HumanTaskTemplate) GetTemplateUIOk() (*map[string]map[string]interface{}, bool)`

GetTemplateUIOk returns a tuple with the TemplateUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUI

`func (o *HumanTaskTemplate) SetTemplateUI(v map[string]map[string]interface{})`

SetTemplateUI sets TemplateUI field to given value.


### GetUpdateTime

`func (o *HumanTaskTemplate) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *HumanTaskTemplate) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *HumanTaskTemplate) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *HumanTaskTemplate) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *HumanTaskTemplate) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *HumanTaskTemplate) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *HumanTaskTemplate) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *HumanTaskTemplate) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVersion

`func (o *HumanTaskTemplate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HumanTaskTemplate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HumanTaskTemplate) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


