# WorkflowDef

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
**InputTemplate** | Pointer to  |  | [optional] 
**MaskedFields** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to  |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OutputParameters** | Pointer to  |  | [optional] 
**OutputSchema** | Pointer to [**SchemaDef**](SchemaDef.md) |  | [optional] 
**OwnerApp** | Pointer to **string** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**RateLimitConfig** | Pointer to [**RateLimitConfig**](RateLimitConfig.md) |  | [optional] 
**Restartable** | Pointer to **bool** |  | [optional] 
**SchemaVersion** | Pointer to **int32** |  | [optional] 
**Tasks** | [**[]WorkflowTask**](WorkflowTask.md) |  | 
**TimeoutPolicy** | Pointer to **string** |  | [optional] 
**TimeoutSeconds** | **int64** |  | 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to  |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**WorkflowStatusListenerEnabled** | Pointer to **bool** |  | [optional] 
**WorkflowStatusListenerSink** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowDef

`func NewWorkflowDef(tasks []WorkflowTask, timeoutSeconds int64, ) *WorkflowDef`

NewWorkflowDef instantiates a new WorkflowDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDefWithDefaults

`func NewWorkflowDefWithDefaults() *WorkflowDef`

NewWorkflowDefWithDefaults instantiates a new WorkflowDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheConfig

`func (o *WorkflowDef) GetCacheConfig() CacheConfig`

GetCacheConfig returns the CacheConfig field if non-nil, zero value otherwise.

### GetCacheConfigOk

`func (o *WorkflowDef) GetCacheConfigOk() (*CacheConfig, bool)`

GetCacheConfigOk returns a tuple with the CacheConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheConfig

`func (o *WorkflowDef) SetCacheConfig(v CacheConfig)`

SetCacheConfig sets CacheConfig field to given value.

### HasCacheConfig

`func (o *WorkflowDef) HasCacheConfig() bool`

HasCacheConfig returns a boolean if a field has been set.

### GetCreateTime

`func (o *WorkflowDef) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *WorkflowDef) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *WorkflowDef) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *WorkflowDef) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WorkflowDef) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkflowDef) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkflowDef) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkflowDef) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnforceSchema

`func (o *WorkflowDef) GetEnforceSchema() bool`

GetEnforceSchema returns the EnforceSchema field if non-nil, zero value otherwise.

### GetEnforceSchemaOk

`func (o *WorkflowDef) GetEnforceSchemaOk() (*bool, bool)`

GetEnforceSchemaOk returns a tuple with the EnforceSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSchema

`func (o *WorkflowDef) SetEnforceSchema(v bool)`

SetEnforceSchema sets EnforceSchema field to given value.

### HasEnforceSchema

`func (o *WorkflowDef) HasEnforceSchema() bool`

HasEnforceSchema returns a boolean if a field has been set.

### GetFailureWorkflow

`func (o *WorkflowDef) GetFailureWorkflow() string`

GetFailureWorkflow returns the FailureWorkflow field if non-nil, zero value otherwise.

### GetFailureWorkflowOk

`func (o *WorkflowDef) GetFailureWorkflowOk() (*string, bool)`

GetFailureWorkflowOk returns a tuple with the FailureWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureWorkflow

`func (o *WorkflowDef) SetFailureWorkflow(v string)`

SetFailureWorkflow sets FailureWorkflow field to given value.

### HasFailureWorkflow

`func (o *WorkflowDef) HasFailureWorkflow() bool`

HasFailureWorkflow returns a boolean if a field has been set.

### GetInputParameters

`func (o *WorkflowDef) GetInputParameters() []string`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowDef) GetInputParametersOk() (*[]string, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowDef) SetInputParameters(v []string)`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowDef) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### GetInputSchema

`func (o *WorkflowDef) GetInputSchema() SchemaDef`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *WorkflowDef) GetInputSchemaOk() (*SchemaDef, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *WorkflowDef) SetInputSchema(v SchemaDef)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *WorkflowDef) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetInputTemplate

`func (o *WorkflowDef) GetInputTemplate() map[string]map[string]interface{}`

GetInputTemplate returns the InputTemplate field if non-nil, zero value otherwise.

### GetInputTemplateOk

`func (o *WorkflowDef) GetInputTemplateOk() (*map[string]map[string]interface{}, bool)`

GetInputTemplateOk returns a tuple with the InputTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTemplate

`func (o *WorkflowDef) SetInputTemplate(v map[string]map[string]interface{})`

SetInputTemplate sets InputTemplate field to given value.

### HasInputTemplate

`func (o *WorkflowDef) HasInputTemplate() bool`

HasInputTemplate returns a boolean if a field has been set.

### SetInputTemplateNil

`func (o *WorkflowDef) SetInputTemplateNil(b bool)`

 SetInputTemplateNil sets the value for InputTemplate to be an explicit nil

### UnsetInputTemplate
`func (o *WorkflowDef) UnsetInputTemplate()`

UnsetInputTemplate ensures that no value is present for InputTemplate, not even an explicit nil
### GetMaskedFields

`func (o *WorkflowDef) GetMaskedFields() []string`

GetMaskedFields returns the MaskedFields field if non-nil, zero value otherwise.

### GetMaskedFieldsOk

`func (o *WorkflowDef) GetMaskedFieldsOk() (*[]string, bool)`

GetMaskedFieldsOk returns a tuple with the MaskedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedFields

`func (o *WorkflowDef) SetMaskedFields(v []string)`

SetMaskedFields sets MaskedFields field to given value.

### HasMaskedFields

`func (o *WorkflowDef) HasMaskedFields() bool`

HasMaskedFields returns a boolean if a field has been set.

### GetMetadata

`func (o *WorkflowDef) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WorkflowDef) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WorkflowDef) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WorkflowDef) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *WorkflowDef) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WorkflowDef) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *WorkflowDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowDef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowDef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowDef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputParameters

`func (o *WorkflowDef) GetOutputParameters() map[string]map[string]interface{}`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *WorkflowDef) GetOutputParametersOk() (*map[string]map[string]interface{}, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *WorkflowDef) SetOutputParameters(v map[string]map[string]interface{})`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *WorkflowDef) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### SetOutputParametersNil

`func (o *WorkflowDef) SetOutputParametersNil(b bool)`

 SetOutputParametersNil sets the value for OutputParameters to be an explicit nil

### UnsetOutputParameters
`func (o *WorkflowDef) UnsetOutputParameters()`

UnsetOutputParameters ensures that no value is present for OutputParameters, not even an explicit nil
### GetOutputSchema

`func (o *WorkflowDef) GetOutputSchema() SchemaDef`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *WorkflowDef) GetOutputSchemaOk() (*SchemaDef, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *WorkflowDef) SetOutputSchema(v SchemaDef)`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *WorkflowDef) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetOwnerApp

`func (o *WorkflowDef) GetOwnerApp() string`

GetOwnerApp returns the OwnerApp field if non-nil, zero value otherwise.

### GetOwnerAppOk

`func (o *WorkflowDef) GetOwnerAppOk() (*string, bool)`

GetOwnerAppOk returns a tuple with the OwnerApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerApp

`func (o *WorkflowDef) SetOwnerApp(v string)`

SetOwnerApp sets OwnerApp field to given value.

### HasOwnerApp

`func (o *WorkflowDef) HasOwnerApp() bool`

HasOwnerApp returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *WorkflowDef) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *WorkflowDef) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *WorkflowDef) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *WorkflowDef) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetRateLimitConfig

`func (o *WorkflowDef) GetRateLimitConfig() RateLimitConfig`

GetRateLimitConfig returns the RateLimitConfig field if non-nil, zero value otherwise.

### GetRateLimitConfigOk

`func (o *WorkflowDef) GetRateLimitConfigOk() (*RateLimitConfig, bool)`

GetRateLimitConfigOk returns a tuple with the RateLimitConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitConfig

`func (o *WorkflowDef) SetRateLimitConfig(v RateLimitConfig)`

SetRateLimitConfig sets RateLimitConfig field to given value.

### HasRateLimitConfig

`func (o *WorkflowDef) HasRateLimitConfig() bool`

HasRateLimitConfig returns a boolean if a field has been set.

### GetRestartable

`func (o *WorkflowDef) GetRestartable() bool`

GetRestartable returns the Restartable field if non-nil, zero value otherwise.

### GetRestartableOk

`func (o *WorkflowDef) GetRestartableOk() (*bool, bool)`

GetRestartableOk returns a tuple with the Restartable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartable

`func (o *WorkflowDef) SetRestartable(v bool)`

SetRestartable sets Restartable field to given value.

### HasRestartable

`func (o *WorkflowDef) HasRestartable() bool`

HasRestartable returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *WorkflowDef) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *WorkflowDef) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *WorkflowDef) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *WorkflowDef) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetTasks

`func (o *WorkflowDef) GetTasks() []WorkflowTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *WorkflowDef) GetTasksOk() (*[]WorkflowTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *WorkflowDef) SetTasks(v []WorkflowTask)`

SetTasks sets Tasks field to given value.


### GetTimeoutPolicy

`func (o *WorkflowDef) GetTimeoutPolicy() string`

GetTimeoutPolicy returns the TimeoutPolicy field if non-nil, zero value otherwise.

### GetTimeoutPolicyOk

`func (o *WorkflowDef) GetTimeoutPolicyOk() (*string, bool)`

GetTimeoutPolicyOk returns a tuple with the TimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutPolicy

`func (o *WorkflowDef) SetTimeoutPolicy(v string)`

SetTimeoutPolicy sets TimeoutPolicy field to given value.

### HasTimeoutPolicy

`func (o *WorkflowDef) HasTimeoutPolicy() bool`

HasTimeoutPolicy returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *WorkflowDef) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *WorkflowDef) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *WorkflowDef) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### GetUpdateTime

`func (o *WorkflowDef) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *WorkflowDef) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *WorkflowDef) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *WorkflowDef) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkflowDef) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkflowDef) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkflowDef) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkflowDef) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVariables

`func (o *WorkflowDef) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *WorkflowDef) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *WorkflowDef) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *WorkflowDef) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *WorkflowDef) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *WorkflowDef) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetVersion

`func (o *WorkflowDef) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowDef) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowDef) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowDef) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkflowStatusListenerEnabled

`func (o *WorkflowDef) GetWorkflowStatusListenerEnabled() bool`

GetWorkflowStatusListenerEnabled returns the WorkflowStatusListenerEnabled field if non-nil, zero value otherwise.

### GetWorkflowStatusListenerEnabledOk

`func (o *WorkflowDef) GetWorkflowStatusListenerEnabledOk() (*bool, bool)`

GetWorkflowStatusListenerEnabledOk returns a tuple with the WorkflowStatusListenerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatusListenerEnabled

`func (o *WorkflowDef) SetWorkflowStatusListenerEnabled(v bool)`

SetWorkflowStatusListenerEnabled sets WorkflowStatusListenerEnabled field to given value.

### HasWorkflowStatusListenerEnabled

`func (o *WorkflowDef) HasWorkflowStatusListenerEnabled() bool`

HasWorkflowStatusListenerEnabled returns a boolean if a field has been set.

### GetWorkflowStatusListenerSink

`func (o *WorkflowDef) GetWorkflowStatusListenerSink() string`

GetWorkflowStatusListenerSink returns the WorkflowStatusListenerSink field if non-nil, zero value otherwise.

### GetWorkflowStatusListenerSinkOk

`func (o *WorkflowDef) GetWorkflowStatusListenerSinkOk() (*string, bool)`

GetWorkflowStatusListenerSinkOk returns a tuple with the WorkflowStatusListenerSink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatusListenerSink

`func (o *WorkflowDef) SetWorkflowStatusListenerSink(v string)`

SetWorkflowStatusListenerSink sets WorkflowStatusListenerSink field to given value.

### HasWorkflowStatusListenerSink

`func (o *WorkflowDef) HasWorkflowStatusListenerSink() bool`

HasWorkflowStatusListenerSink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


