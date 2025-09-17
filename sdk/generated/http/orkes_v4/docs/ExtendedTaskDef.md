# ExtendedTaskDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackoffScaleFactor** | Pointer to **int32** |  | [optional] 
**BaseType** | Pointer to **string** |  | [optional] 
**ConcurrentExecLimit** | Pointer to **int32** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnforceSchema** | Pointer to **bool** |  | [optional] 
**ExecutionNameSpace** | Pointer to **string** |  | [optional] 
**InputKeys** | Pointer to **[]string** |  | [optional] 
**InputSchema** | Pointer to [**SchemaDef**](SchemaDef.md) |  | [optional] 
**InputTemplate** | Pointer to  |  | [optional] 
**IsolationGroupId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OutputKeys** | Pointer to **[]string** |  | [optional] 
**OutputSchema** | Pointer to [**SchemaDef**](SchemaDef.md) |  | [optional] 
**OverwriteTags** | Pointer to **bool** |  | [optional] 
**OwnerApp** | Pointer to **string** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**PollTimeoutSeconds** | Pointer to **int32** |  | [optional] 
**RateLimitFrequencyInSeconds** | Pointer to **int32** |  | [optional] 
**RateLimitPerFrequency** | Pointer to **int32** |  | [optional] 
**ResponseTimeoutSeconds** | Pointer to **int64** |  | [optional] 
**RetryCount** | Pointer to **int32** |  | [optional] 
**RetryDelaySeconds** | Pointer to **int32** |  | [optional] 
**RetryLogic** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**TimeoutPolicy** | Pointer to **string** |  | [optional] 
**TimeoutSeconds** | **int64** |  | 
**TotalTimeoutSeconds** | **int64** |  | 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewExtendedTaskDef

`func NewExtendedTaskDef(timeoutSeconds int64, totalTimeoutSeconds int64, ) *ExtendedTaskDef`

NewExtendedTaskDef instantiates a new ExtendedTaskDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedTaskDefWithDefaults

`func NewExtendedTaskDefWithDefaults() *ExtendedTaskDef`

NewExtendedTaskDefWithDefaults instantiates a new ExtendedTaskDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackoffScaleFactor

`func (o *ExtendedTaskDef) GetBackoffScaleFactor() int32`

GetBackoffScaleFactor returns the BackoffScaleFactor field if non-nil, zero value otherwise.

### GetBackoffScaleFactorOk

`func (o *ExtendedTaskDef) GetBackoffScaleFactorOk() (*int32, bool)`

GetBackoffScaleFactorOk returns a tuple with the BackoffScaleFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoffScaleFactor

`func (o *ExtendedTaskDef) SetBackoffScaleFactor(v int32)`

SetBackoffScaleFactor sets BackoffScaleFactor field to given value.

### HasBackoffScaleFactor

`func (o *ExtendedTaskDef) HasBackoffScaleFactor() bool`

HasBackoffScaleFactor returns a boolean if a field has been set.

### GetBaseType

`func (o *ExtendedTaskDef) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *ExtendedTaskDef) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *ExtendedTaskDef) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *ExtendedTaskDef) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetConcurrentExecLimit

`func (o *ExtendedTaskDef) GetConcurrentExecLimit() int32`

GetConcurrentExecLimit returns the ConcurrentExecLimit field if non-nil, zero value otherwise.

### GetConcurrentExecLimitOk

`func (o *ExtendedTaskDef) GetConcurrentExecLimitOk() (*int32, bool)`

GetConcurrentExecLimitOk returns a tuple with the ConcurrentExecLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentExecLimit

`func (o *ExtendedTaskDef) SetConcurrentExecLimit(v int32)`

SetConcurrentExecLimit sets ConcurrentExecLimit field to given value.

### HasConcurrentExecLimit

`func (o *ExtendedTaskDef) HasConcurrentExecLimit() bool`

HasConcurrentExecLimit returns a boolean if a field has been set.

### GetCreateTime

`func (o *ExtendedTaskDef) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ExtendedTaskDef) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ExtendedTaskDef) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ExtendedTaskDef) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ExtendedTaskDef) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ExtendedTaskDef) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ExtendedTaskDef) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ExtendedTaskDef) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *ExtendedTaskDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtendedTaskDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtendedTaskDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtendedTaskDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnforceSchema

`func (o *ExtendedTaskDef) GetEnforceSchema() bool`

GetEnforceSchema returns the EnforceSchema field if non-nil, zero value otherwise.

### GetEnforceSchemaOk

`func (o *ExtendedTaskDef) GetEnforceSchemaOk() (*bool, bool)`

GetEnforceSchemaOk returns a tuple with the EnforceSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSchema

`func (o *ExtendedTaskDef) SetEnforceSchema(v bool)`

SetEnforceSchema sets EnforceSchema field to given value.

### HasEnforceSchema

`func (o *ExtendedTaskDef) HasEnforceSchema() bool`

HasEnforceSchema returns a boolean if a field has been set.

### GetExecutionNameSpace

`func (o *ExtendedTaskDef) GetExecutionNameSpace() string`

GetExecutionNameSpace returns the ExecutionNameSpace field if non-nil, zero value otherwise.

### GetExecutionNameSpaceOk

`func (o *ExtendedTaskDef) GetExecutionNameSpaceOk() (*string, bool)`

GetExecutionNameSpaceOk returns a tuple with the ExecutionNameSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionNameSpace

`func (o *ExtendedTaskDef) SetExecutionNameSpace(v string)`

SetExecutionNameSpace sets ExecutionNameSpace field to given value.

### HasExecutionNameSpace

`func (o *ExtendedTaskDef) HasExecutionNameSpace() bool`

HasExecutionNameSpace returns a boolean if a field has been set.

### GetInputKeys

`func (o *ExtendedTaskDef) GetInputKeys() []string`

GetInputKeys returns the InputKeys field if non-nil, zero value otherwise.

### GetInputKeysOk

`func (o *ExtendedTaskDef) GetInputKeysOk() (*[]string, bool)`

GetInputKeysOk returns a tuple with the InputKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputKeys

`func (o *ExtendedTaskDef) SetInputKeys(v []string)`

SetInputKeys sets InputKeys field to given value.

### HasInputKeys

`func (o *ExtendedTaskDef) HasInputKeys() bool`

HasInputKeys returns a boolean if a field has been set.

### GetInputSchema

`func (o *ExtendedTaskDef) GetInputSchema() SchemaDef`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *ExtendedTaskDef) GetInputSchemaOk() (*SchemaDef, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *ExtendedTaskDef) SetInputSchema(v SchemaDef)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *ExtendedTaskDef) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetInputTemplate

`func (o *ExtendedTaskDef) GetInputTemplate() map[string]map[string]interface{}`

GetInputTemplate returns the InputTemplate field if non-nil, zero value otherwise.

### GetInputTemplateOk

`func (o *ExtendedTaskDef) GetInputTemplateOk() (*map[string]map[string]interface{}, bool)`

GetInputTemplateOk returns a tuple with the InputTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTemplate

`func (o *ExtendedTaskDef) SetInputTemplate(v map[string]map[string]interface{})`

SetInputTemplate sets InputTemplate field to given value.

### HasInputTemplate

`func (o *ExtendedTaskDef) HasInputTemplate() bool`

HasInputTemplate returns a boolean if a field has been set.

### SetInputTemplateNil

`func (o *ExtendedTaskDef) SetInputTemplateNil(b bool)`

 SetInputTemplateNil sets the value for InputTemplate to be an explicit nil

### UnsetInputTemplate
`func (o *ExtendedTaskDef) UnsetInputTemplate()`

UnsetInputTemplate ensures that no value is present for InputTemplate, not even an explicit nil
### GetIsolationGroupId

`func (o *ExtendedTaskDef) GetIsolationGroupId() string`

GetIsolationGroupId returns the IsolationGroupId field if non-nil, zero value otherwise.

### GetIsolationGroupIdOk

`func (o *ExtendedTaskDef) GetIsolationGroupIdOk() (*string, bool)`

GetIsolationGroupIdOk returns a tuple with the IsolationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationGroupId

`func (o *ExtendedTaskDef) SetIsolationGroupId(v string)`

SetIsolationGroupId sets IsolationGroupId field to given value.

### HasIsolationGroupId

`func (o *ExtendedTaskDef) HasIsolationGroupId() bool`

HasIsolationGroupId returns a boolean if a field has been set.

### GetName

`func (o *ExtendedTaskDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtendedTaskDef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtendedTaskDef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExtendedTaskDef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputKeys

`func (o *ExtendedTaskDef) GetOutputKeys() []string`

GetOutputKeys returns the OutputKeys field if non-nil, zero value otherwise.

### GetOutputKeysOk

`func (o *ExtendedTaskDef) GetOutputKeysOk() (*[]string, bool)`

GetOutputKeysOk returns a tuple with the OutputKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputKeys

`func (o *ExtendedTaskDef) SetOutputKeys(v []string)`

SetOutputKeys sets OutputKeys field to given value.

### HasOutputKeys

`func (o *ExtendedTaskDef) HasOutputKeys() bool`

HasOutputKeys returns a boolean if a field has been set.

### GetOutputSchema

`func (o *ExtendedTaskDef) GetOutputSchema() SchemaDef`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *ExtendedTaskDef) GetOutputSchemaOk() (*SchemaDef, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *ExtendedTaskDef) SetOutputSchema(v SchemaDef)`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *ExtendedTaskDef) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetOverwriteTags

`func (o *ExtendedTaskDef) GetOverwriteTags() bool`

GetOverwriteTags returns the OverwriteTags field if non-nil, zero value otherwise.

### GetOverwriteTagsOk

`func (o *ExtendedTaskDef) GetOverwriteTagsOk() (*bool, bool)`

GetOverwriteTagsOk returns a tuple with the OverwriteTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteTags

`func (o *ExtendedTaskDef) SetOverwriteTags(v bool)`

SetOverwriteTags sets OverwriteTags field to given value.

### HasOverwriteTags

`func (o *ExtendedTaskDef) HasOverwriteTags() bool`

HasOverwriteTags returns a boolean if a field has been set.

### GetOwnerApp

`func (o *ExtendedTaskDef) GetOwnerApp() string`

GetOwnerApp returns the OwnerApp field if non-nil, zero value otherwise.

### GetOwnerAppOk

`func (o *ExtendedTaskDef) GetOwnerAppOk() (*string, bool)`

GetOwnerAppOk returns a tuple with the OwnerApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerApp

`func (o *ExtendedTaskDef) SetOwnerApp(v string)`

SetOwnerApp sets OwnerApp field to given value.

### HasOwnerApp

`func (o *ExtendedTaskDef) HasOwnerApp() bool`

HasOwnerApp returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *ExtendedTaskDef) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *ExtendedTaskDef) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *ExtendedTaskDef) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *ExtendedTaskDef) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetPollTimeoutSeconds

`func (o *ExtendedTaskDef) GetPollTimeoutSeconds() int32`

GetPollTimeoutSeconds returns the PollTimeoutSeconds field if non-nil, zero value otherwise.

### GetPollTimeoutSecondsOk

`func (o *ExtendedTaskDef) GetPollTimeoutSecondsOk() (*int32, bool)`

GetPollTimeoutSecondsOk returns a tuple with the PollTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollTimeoutSeconds

`func (o *ExtendedTaskDef) SetPollTimeoutSeconds(v int32)`

SetPollTimeoutSeconds sets PollTimeoutSeconds field to given value.

### HasPollTimeoutSeconds

`func (o *ExtendedTaskDef) HasPollTimeoutSeconds() bool`

HasPollTimeoutSeconds returns a boolean if a field has been set.

### GetRateLimitFrequencyInSeconds

`func (o *ExtendedTaskDef) GetRateLimitFrequencyInSeconds() int32`

GetRateLimitFrequencyInSeconds returns the RateLimitFrequencyInSeconds field if non-nil, zero value otherwise.

### GetRateLimitFrequencyInSecondsOk

`func (o *ExtendedTaskDef) GetRateLimitFrequencyInSecondsOk() (*int32, bool)`

GetRateLimitFrequencyInSecondsOk returns a tuple with the RateLimitFrequencyInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitFrequencyInSeconds

`func (o *ExtendedTaskDef) SetRateLimitFrequencyInSeconds(v int32)`

SetRateLimitFrequencyInSeconds sets RateLimitFrequencyInSeconds field to given value.

### HasRateLimitFrequencyInSeconds

`func (o *ExtendedTaskDef) HasRateLimitFrequencyInSeconds() bool`

HasRateLimitFrequencyInSeconds returns a boolean if a field has been set.

### GetRateLimitPerFrequency

`func (o *ExtendedTaskDef) GetRateLimitPerFrequency() int32`

GetRateLimitPerFrequency returns the RateLimitPerFrequency field if non-nil, zero value otherwise.

### GetRateLimitPerFrequencyOk

`func (o *ExtendedTaskDef) GetRateLimitPerFrequencyOk() (*int32, bool)`

GetRateLimitPerFrequencyOk returns a tuple with the RateLimitPerFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerFrequency

`func (o *ExtendedTaskDef) SetRateLimitPerFrequency(v int32)`

SetRateLimitPerFrequency sets RateLimitPerFrequency field to given value.

### HasRateLimitPerFrequency

`func (o *ExtendedTaskDef) HasRateLimitPerFrequency() bool`

HasRateLimitPerFrequency returns a boolean if a field has been set.

### GetResponseTimeoutSeconds

`func (o *ExtendedTaskDef) GetResponseTimeoutSeconds() int64`

GetResponseTimeoutSeconds returns the ResponseTimeoutSeconds field if non-nil, zero value otherwise.

### GetResponseTimeoutSecondsOk

`func (o *ExtendedTaskDef) GetResponseTimeoutSecondsOk() (*int64, bool)`

GetResponseTimeoutSecondsOk returns a tuple with the ResponseTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeoutSeconds

`func (o *ExtendedTaskDef) SetResponseTimeoutSeconds(v int64)`

SetResponseTimeoutSeconds sets ResponseTimeoutSeconds field to given value.

### HasResponseTimeoutSeconds

`func (o *ExtendedTaskDef) HasResponseTimeoutSeconds() bool`

HasResponseTimeoutSeconds returns a boolean if a field has been set.

### GetRetryCount

`func (o *ExtendedTaskDef) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *ExtendedTaskDef) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *ExtendedTaskDef) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *ExtendedTaskDef) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *ExtendedTaskDef) GetRetryDelaySeconds() int32`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *ExtendedTaskDef) GetRetryDelaySecondsOk() (*int32, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *ExtendedTaskDef) SetRetryDelaySeconds(v int32)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *ExtendedTaskDef) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetRetryLogic

`func (o *ExtendedTaskDef) GetRetryLogic() string`

GetRetryLogic returns the RetryLogic field if non-nil, zero value otherwise.

### GetRetryLogicOk

`func (o *ExtendedTaskDef) GetRetryLogicOk() (*string, bool)`

GetRetryLogicOk returns a tuple with the RetryLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryLogic

`func (o *ExtendedTaskDef) SetRetryLogic(v string)`

SetRetryLogic sets RetryLogic field to given value.

### HasRetryLogic

`func (o *ExtendedTaskDef) HasRetryLogic() bool`

HasRetryLogic returns a boolean if a field has been set.

### GetTags

`func (o *ExtendedTaskDef) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExtendedTaskDef) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExtendedTaskDef) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExtendedTaskDef) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeoutPolicy

`func (o *ExtendedTaskDef) GetTimeoutPolicy() string`

GetTimeoutPolicy returns the TimeoutPolicy field if non-nil, zero value otherwise.

### GetTimeoutPolicyOk

`func (o *ExtendedTaskDef) GetTimeoutPolicyOk() (*string, bool)`

GetTimeoutPolicyOk returns a tuple with the TimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutPolicy

`func (o *ExtendedTaskDef) SetTimeoutPolicy(v string)`

SetTimeoutPolicy sets TimeoutPolicy field to given value.

### HasTimeoutPolicy

`func (o *ExtendedTaskDef) HasTimeoutPolicy() bool`

HasTimeoutPolicy returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *ExtendedTaskDef) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ExtendedTaskDef) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ExtendedTaskDef) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### GetTotalTimeoutSeconds

`func (o *ExtendedTaskDef) GetTotalTimeoutSeconds() int64`

GetTotalTimeoutSeconds returns the TotalTimeoutSeconds field if non-nil, zero value otherwise.

### GetTotalTimeoutSecondsOk

`func (o *ExtendedTaskDef) GetTotalTimeoutSecondsOk() (*int64, bool)`

GetTotalTimeoutSecondsOk returns a tuple with the TotalTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeoutSeconds

`func (o *ExtendedTaskDef) SetTotalTimeoutSeconds(v int64)`

SetTotalTimeoutSeconds sets TotalTimeoutSeconds field to given value.


### GetUpdateTime

`func (o *ExtendedTaskDef) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *ExtendedTaskDef) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *ExtendedTaskDef) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *ExtendedTaskDef) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ExtendedTaskDef) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ExtendedTaskDef) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ExtendedTaskDef) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ExtendedTaskDef) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


