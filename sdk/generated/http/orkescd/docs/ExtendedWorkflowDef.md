# ExtendedWorkflowDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheConfig** | Pointer to [**CacheConfig**](CacheConfig.md) |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnforceSchema** | Pointer to **bool** |  | [optional] 
**FailureWorkflow** | Pointer to **string** |  | [optional] 
**InputParameters** | Pointer to **[]string** |  | [optional] 
**InputSchema** | Pointer to [**SchemaDef**](SchemaDef.md) |  | [optional] 
**InputTemplate** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**MaskedFields** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**OutputParameters** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**OutputSchema** | Pointer to [**SchemaDef**](SchemaDef.md) |  | [optional] 
**OverwriteTags** | Pointer to **bool** |  | [optional] 
**OwnerApp** | Pointer to **string** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**RateLimitConfig** | Pointer to [**RateLimitConfig**](RateLimitConfig.md) |  | [optional] 
**Restartable** | Pointer to **bool** |  | [optional] 
**SchemaVersion** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Tasks** | [**[]WorkflowTask**](WorkflowTask.md) |  | 
**TimeoutPolicy** | Pointer to **string** |  | [optional] 
**TimeoutSeconds** | **int64** |  | 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**WorkflowStatusListenerEnabled** | Pointer to **bool** |  | [optional] 
**WorkflowStatusListenerSink** | Pointer to **string** |  | [optional] 

## Methods

### NewExtendedWorkflowDef

`func NewExtendedWorkflowDef(name string, tasks []WorkflowTask, timeoutSeconds int64, ) *ExtendedWorkflowDef`

NewExtendedWorkflowDef instantiates a new ExtendedWorkflowDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedWorkflowDefWithDefaults

`func NewExtendedWorkflowDefWithDefaults() *ExtendedWorkflowDef`

NewExtendedWorkflowDefWithDefaults instantiates a new ExtendedWorkflowDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheConfig

`func (o *ExtendedWorkflowDef) GetCacheConfig() CacheConfig`

GetCacheConfig returns the CacheConfig field if non-nil, zero value otherwise.

### GetCacheConfigOk

`func (o *ExtendedWorkflowDef) GetCacheConfigOk() (*CacheConfig, bool)`

GetCacheConfigOk returns a tuple with the CacheConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheConfig

`func (o *ExtendedWorkflowDef) SetCacheConfig(v CacheConfig)`

SetCacheConfig sets CacheConfig field to given value.

### HasCacheConfig

`func (o *ExtendedWorkflowDef) HasCacheConfig() bool`

HasCacheConfig returns a boolean if a field has been set.

### GetCreateTime

`func (o *ExtendedWorkflowDef) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ExtendedWorkflowDef) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ExtendedWorkflowDef) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ExtendedWorkflowDef) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ExtendedWorkflowDef) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ExtendedWorkflowDef) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ExtendedWorkflowDef) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ExtendedWorkflowDef) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *ExtendedWorkflowDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtendedWorkflowDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtendedWorkflowDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtendedWorkflowDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnforceSchema

`func (o *ExtendedWorkflowDef) GetEnforceSchema() bool`

GetEnforceSchema returns the EnforceSchema field if non-nil, zero value otherwise.

### GetEnforceSchemaOk

`func (o *ExtendedWorkflowDef) GetEnforceSchemaOk() (*bool, bool)`

GetEnforceSchemaOk returns a tuple with the EnforceSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSchema

`func (o *ExtendedWorkflowDef) SetEnforceSchema(v bool)`

SetEnforceSchema sets EnforceSchema field to given value.

### HasEnforceSchema

`func (o *ExtendedWorkflowDef) HasEnforceSchema() bool`

HasEnforceSchema returns a boolean if a field has been set.

### GetFailureWorkflow

`func (o *ExtendedWorkflowDef) GetFailureWorkflow() string`

GetFailureWorkflow returns the FailureWorkflow field if non-nil, zero value otherwise.

### GetFailureWorkflowOk

`func (o *ExtendedWorkflowDef) GetFailureWorkflowOk() (*string, bool)`

GetFailureWorkflowOk returns a tuple with the FailureWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureWorkflow

`func (o *ExtendedWorkflowDef) SetFailureWorkflow(v string)`

SetFailureWorkflow sets FailureWorkflow field to given value.

### HasFailureWorkflow

`func (o *ExtendedWorkflowDef) HasFailureWorkflow() bool`

HasFailureWorkflow returns a boolean if a field has been set.

### GetInputParameters

`func (o *ExtendedWorkflowDef) GetInputParameters() []string`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *ExtendedWorkflowDef) GetInputParametersOk() (*[]string, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *ExtendedWorkflowDef) SetInputParameters(v []string)`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *ExtendedWorkflowDef) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### GetInputSchema

`func (o *ExtendedWorkflowDef) GetInputSchema() SchemaDef`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *ExtendedWorkflowDef) GetInputSchemaOk() (*SchemaDef, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *ExtendedWorkflowDef) SetInputSchema(v SchemaDef)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *ExtendedWorkflowDef) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetInputTemplate

`func (o *ExtendedWorkflowDef) GetInputTemplate() map[string]map[string]interface{}`

GetInputTemplate returns the InputTemplate field if non-nil, zero value otherwise.

### GetInputTemplateOk

`func (o *ExtendedWorkflowDef) GetInputTemplateOk() (*map[string]map[string]interface{}, bool)`

GetInputTemplateOk returns a tuple with the InputTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTemplate

`func (o *ExtendedWorkflowDef) SetInputTemplate(v map[string]map[string]interface{})`

SetInputTemplate sets InputTemplate field to given value.

### HasInputTemplate

`func (o *ExtendedWorkflowDef) HasInputTemplate() bool`

HasInputTemplate returns a boolean if a field has been set.

### GetMaskedFields

`func (o *ExtendedWorkflowDef) GetMaskedFields() []string`

GetMaskedFields returns the MaskedFields field if non-nil, zero value otherwise.

### GetMaskedFieldsOk

`func (o *ExtendedWorkflowDef) GetMaskedFieldsOk() (*[]string, bool)`

GetMaskedFieldsOk returns a tuple with the MaskedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedFields

`func (o *ExtendedWorkflowDef) SetMaskedFields(v []string)`

SetMaskedFields sets MaskedFields field to given value.

### HasMaskedFields

`func (o *ExtendedWorkflowDef) HasMaskedFields() bool`

HasMaskedFields returns a boolean if a field has been set.

### GetMetadata

`func (o *ExtendedWorkflowDef) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExtendedWorkflowDef) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExtendedWorkflowDef) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ExtendedWorkflowDef) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ExtendedWorkflowDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtendedWorkflowDef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtendedWorkflowDef) SetName(v string)`

SetName sets Name field to given value.


### GetOutputParameters

`func (o *ExtendedWorkflowDef) GetOutputParameters() map[string]map[string]interface{}`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *ExtendedWorkflowDef) GetOutputParametersOk() (*map[string]map[string]interface{}, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *ExtendedWorkflowDef) SetOutputParameters(v map[string]map[string]interface{})`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *ExtendedWorkflowDef) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### GetOutputSchema

`func (o *ExtendedWorkflowDef) GetOutputSchema() SchemaDef`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *ExtendedWorkflowDef) GetOutputSchemaOk() (*SchemaDef, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *ExtendedWorkflowDef) SetOutputSchema(v SchemaDef)`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *ExtendedWorkflowDef) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetOverwriteTags

`func (o *ExtendedWorkflowDef) GetOverwriteTags() bool`

GetOverwriteTags returns the OverwriteTags field if non-nil, zero value otherwise.

### GetOverwriteTagsOk

`func (o *ExtendedWorkflowDef) GetOverwriteTagsOk() (*bool, bool)`

GetOverwriteTagsOk returns a tuple with the OverwriteTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteTags

`func (o *ExtendedWorkflowDef) SetOverwriteTags(v bool)`

SetOverwriteTags sets OverwriteTags field to given value.

### HasOverwriteTags

`func (o *ExtendedWorkflowDef) HasOverwriteTags() bool`

HasOverwriteTags returns a boolean if a field has been set.

### GetOwnerApp

`func (o *ExtendedWorkflowDef) GetOwnerApp() string`

GetOwnerApp returns the OwnerApp field if non-nil, zero value otherwise.

### GetOwnerAppOk

`func (o *ExtendedWorkflowDef) GetOwnerAppOk() (*string, bool)`

GetOwnerAppOk returns a tuple with the OwnerApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerApp

`func (o *ExtendedWorkflowDef) SetOwnerApp(v string)`

SetOwnerApp sets OwnerApp field to given value.

### HasOwnerApp

`func (o *ExtendedWorkflowDef) HasOwnerApp() bool`

HasOwnerApp returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *ExtendedWorkflowDef) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *ExtendedWorkflowDef) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *ExtendedWorkflowDef) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *ExtendedWorkflowDef) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetRateLimitConfig

`func (o *ExtendedWorkflowDef) GetRateLimitConfig() RateLimitConfig`

GetRateLimitConfig returns the RateLimitConfig field if non-nil, zero value otherwise.

### GetRateLimitConfigOk

`func (o *ExtendedWorkflowDef) GetRateLimitConfigOk() (*RateLimitConfig, bool)`

GetRateLimitConfigOk returns a tuple with the RateLimitConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitConfig

`func (o *ExtendedWorkflowDef) SetRateLimitConfig(v RateLimitConfig)`

SetRateLimitConfig sets RateLimitConfig field to given value.

### HasRateLimitConfig

`func (o *ExtendedWorkflowDef) HasRateLimitConfig() bool`

HasRateLimitConfig returns a boolean if a field has been set.

### GetRestartable

`func (o *ExtendedWorkflowDef) GetRestartable() bool`

GetRestartable returns the Restartable field if non-nil, zero value otherwise.

### GetRestartableOk

`func (o *ExtendedWorkflowDef) GetRestartableOk() (*bool, bool)`

GetRestartableOk returns a tuple with the Restartable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartable

`func (o *ExtendedWorkflowDef) SetRestartable(v bool)`

SetRestartable sets Restartable field to given value.

### HasRestartable

`func (o *ExtendedWorkflowDef) HasRestartable() bool`

HasRestartable returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ExtendedWorkflowDef) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ExtendedWorkflowDef) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ExtendedWorkflowDef) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ExtendedWorkflowDef) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetTags

`func (o *ExtendedWorkflowDef) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExtendedWorkflowDef) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExtendedWorkflowDef) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExtendedWorkflowDef) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTasks

`func (o *ExtendedWorkflowDef) GetTasks() []WorkflowTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ExtendedWorkflowDef) GetTasksOk() (*[]WorkflowTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ExtendedWorkflowDef) SetTasks(v []WorkflowTask)`

SetTasks sets Tasks field to given value.


### GetTimeoutPolicy

`func (o *ExtendedWorkflowDef) GetTimeoutPolicy() string`

GetTimeoutPolicy returns the TimeoutPolicy field if non-nil, zero value otherwise.

### GetTimeoutPolicyOk

`func (o *ExtendedWorkflowDef) GetTimeoutPolicyOk() (*string, bool)`

GetTimeoutPolicyOk returns a tuple with the TimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutPolicy

`func (o *ExtendedWorkflowDef) SetTimeoutPolicy(v string)`

SetTimeoutPolicy sets TimeoutPolicy field to given value.

### HasTimeoutPolicy

`func (o *ExtendedWorkflowDef) HasTimeoutPolicy() bool`

HasTimeoutPolicy returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *ExtendedWorkflowDef) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ExtendedWorkflowDef) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ExtendedWorkflowDef) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### GetUpdateTime

`func (o *ExtendedWorkflowDef) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *ExtendedWorkflowDef) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *ExtendedWorkflowDef) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *ExtendedWorkflowDef) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ExtendedWorkflowDef) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ExtendedWorkflowDef) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ExtendedWorkflowDef) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ExtendedWorkflowDef) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVariables

`func (o *ExtendedWorkflowDef) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ExtendedWorkflowDef) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ExtendedWorkflowDef) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ExtendedWorkflowDef) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetVersion

`func (o *ExtendedWorkflowDef) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExtendedWorkflowDef) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExtendedWorkflowDef) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExtendedWorkflowDef) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkflowStatusListenerEnabled

`func (o *ExtendedWorkflowDef) GetWorkflowStatusListenerEnabled() bool`

GetWorkflowStatusListenerEnabled returns the WorkflowStatusListenerEnabled field if non-nil, zero value otherwise.

### GetWorkflowStatusListenerEnabledOk

`func (o *ExtendedWorkflowDef) GetWorkflowStatusListenerEnabledOk() (*bool, bool)`

GetWorkflowStatusListenerEnabledOk returns a tuple with the WorkflowStatusListenerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatusListenerEnabled

`func (o *ExtendedWorkflowDef) SetWorkflowStatusListenerEnabled(v bool)`

SetWorkflowStatusListenerEnabled sets WorkflowStatusListenerEnabled field to given value.

### HasWorkflowStatusListenerEnabled

`func (o *ExtendedWorkflowDef) HasWorkflowStatusListenerEnabled() bool`

HasWorkflowStatusListenerEnabled returns a boolean if a field has been set.

### GetWorkflowStatusListenerSink

`func (o *ExtendedWorkflowDef) GetWorkflowStatusListenerSink() string`

GetWorkflowStatusListenerSink returns the WorkflowStatusListenerSink field if non-nil, zero value otherwise.

### GetWorkflowStatusListenerSinkOk

`func (o *ExtendedWorkflowDef) GetWorkflowStatusListenerSinkOk() (*string, bool)`

GetWorkflowStatusListenerSinkOk returns a tuple with the WorkflowStatusListenerSink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatusListenerSink

`func (o *ExtendedWorkflowDef) SetWorkflowStatusListenerSink(v string)`

SetWorkflowStatusListenerSink sets WorkflowStatusListenerSink field to given value.

### HasWorkflowStatusListenerSink

`func (o *ExtendedWorkflowDef) HasWorkflowStatusListenerSink() bool`

HasWorkflowStatusListenerSink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


