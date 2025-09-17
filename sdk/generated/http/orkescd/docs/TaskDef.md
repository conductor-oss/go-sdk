# TaskDef

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
**InputTemplate** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**IsolationGroupId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**OutputKeys** | Pointer to **[]string** |  | [optional] 
**OutputSchema** | Pointer to [**SchemaDef**](SchemaDef.md) |  | [optional] 
**OwnerApp** | Pointer to **string** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**PollTimeoutSeconds** | Pointer to **int32** |  | [optional] 
**RateLimitFrequencyInSeconds** | Pointer to **int32** |  | [optional] 
**RateLimitPerFrequency** | Pointer to **int32** |  | [optional] 
**ResponseTimeoutSeconds** | Pointer to **int64** |  | [optional] 
**RetryCount** | Pointer to **int32** |  | [optional] 
**RetryDelaySeconds** | Pointer to **int32** |  | [optional] 
**RetryLogic** | Pointer to **string** |  | [optional] 
**TimeoutPolicy** | Pointer to **string** |  | [optional] 
**TimeoutSeconds** | **int64** |  | 
**TotalTimeoutSeconds** | **int64** |  | 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskDef

`func NewTaskDef(name string, timeoutSeconds int64, totalTimeoutSeconds int64, ) *TaskDef`

NewTaskDef instantiates a new TaskDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskDefWithDefaults

`func NewTaskDefWithDefaults() *TaskDef`

NewTaskDefWithDefaults instantiates a new TaskDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackoffScaleFactor

`func (o *TaskDef) GetBackoffScaleFactor() int32`

GetBackoffScaleFactor returns the BackoffScaleFactor field if non-nil, zero value otherwise.

### GetBackoffScaleFactorOk

`func (o *TaskDef) GetBackoffScaleFactorOk() (*int32, bool)`

GetBackoffScaleFactorOk returns a tuple with the BackoffScaleFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoffScaleFactor

`func (o *TaskDef) SetBackoffScaleFactor(v int32)`

SetBackoffScaleFactor sets BackoffScaleFactor field to given value.

### HasBackoffScaleFactor

`func (o *TaskDef) HasBackoffScaleFactor() bool`

HasBackoffScaleFactor returns a boolean if a field has been set.

### GetBaseType

`func (o *TaskDef) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *TaskDef) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *TaskDef) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *TaskDef) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetConcurrentExecLimit

`func (o *TaskDef) GetConcurrentExecLimit() int32`

GetConcurrentExecLimit returns the ConcurrentExecLimit field if non-nil, zero value otherwise.

### GetConcurrentExecLimitOk

`func (o *TaskDef) GetConcurrentExecLimitOk() (*int32, bool)`

GetConcurrentExecLimitOk returns a tuple with the ConcurrentExecLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentExecLimit

`func (o *TaskDef) SetConcurrentExecLimit(v int32)`

SetConcurrentExecLimit sets ConcurrentExecLimit field to given value.

### HasConcurrentExecLimit

`func (o *TaskDef) HasConcurrentExecLimit() bool`

HasConcurrentExecLimit returns a boolean if a field has been set.

### GetCreateTime

`func (o *TaskDef) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *TaskDef) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *TaskDef) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *TaskDef) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *TaskDef) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TaskDef) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TaskDef) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TaskDef) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *TaskDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnforceSchema

`func (o *TaskDef) GetEnforceSchema() bool`

GetEnforceSchema returns the EnforceSchema field if non-nil, zero value otherwise.

### GetEnforceSchemaOk

`func (o *TaskDef) GetEnforceSchemaOk() (*bool, bool)`

GetEnforceSchemaOk returns a tuple with the EnforceSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSchema

`func (o *TaskDef) SetEnforceSchema(v bool)`

SetEnforceSchema sets EnforceSchema field to given value.

### HasEnforceSchema

`func (o *TaskDef) HasEnforceSchema() bool`

HasEnforceSchema returns a boolean if a field has been set.

### GetExecutionNameSpace

`func (o *TaskDef) GetExecutionNameSpace() string`

GetExecutionNameSpace returns the ExecutionNameSpace field if non-nil, zero value otherwise.

### GetExecutionNameSpaceOk

`func (o *TaskDef) GetExecutionNameSpaceOk() (*string, bool)`

GetExecutionNameSpaceOk returns a tuple with the ExecutionNameSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionNameSpace

`func (o *TaskDef) SetExecutionNameSpace(v string)`

SetExecutionNameSpace sets ExecutionNameSpace field to given value.

### HasExecutionNameSpace

`func (o *TaskDef) HasExecutionNameSpace() bool`

HasExecutionNameSpace returns a boolean if a field has been set.

### GetInputKeys

`func (o *TaskDef) GetInputKeys() []string`

GetInputKeys returns the InputKeys field if non-nil, zero value otherwise.

### GetInputKeysOk

`func (o *TaskDef) GetInputKeysOk() (*[]string, bool)`

GetInputKeysOk returns a tuple with the InputKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputKeys

`func (o *TaskDef) SetInputKeys(v []string)`

SetInputKeys sets InputKeys field to given value.

### HasInputKeys

`func (o *TaskDef) HasInputKeys() bool`

HasInputKeys returns a boolean if a field has been set.

### GetInputSchema

`func (o *TaskDef) GetInputSchema() SchemaDef`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *TaskDef) GetInputSchemaOk() (*SchemaDef, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *TaskDef) SetInputSchema(v SchemaDef)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *TaskDef) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetInputTemplate

`func (o *TaskDef) GetInputTemplate() map[string]map[string]interface{}`

GetInputTemplate returns the InputTemplate field if non-nil, zero value otherwise.

### GetInputTemplateOk

`func (o *TaskDef) GetInputTemplateOk() (*map[string]map[string]interface{}, bool)`

GetInputTemplateOk returns a tuple with the InputTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTemplate

`func (o *TaskDef) SetInputTemplate(v map[string]map[string]interface{})`

SetInputTemplate sets InputTemplate field to given value.

### HasInputTemplate

`func (o *TaskDef) HasInputTemplate() bool`

HasInputTemplate returns a boolean if a field has been set.

### GetIsolationGroupId

`func (o *TaskDef) GetIsolationGroupId() string`

GetIsolationGroupId returns the IsolationGroupId field if non-nil, zero value otherwise.

### GetIsolationGroupIdOk

`func (o *TaskDef) GetIsolationGroupIdOk() (*string, bool)`

GetIsolationGroupIdOk returns a tuple with the IsolationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationGroupId

`func (o *TaskDef) SetIsolationGroupId(v string)`

SetIsolationGroupId sets IsolationGroupId field to given value.

### HasIsolationGroupId

`func (o *TaskDef) HasIsolationGroupId() bool`

HasIsolationGroupId returns a boolean if a field has been set.

### GetName

`func (o *TaskDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskDef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskDef) SetName(v string)`

SetName sets Name field to given value.


### GetOutputKeys

`func (o *TaskDef) GetOutputKeys() []string`

GetOutputKeys returns the OutputKeys field if non-nil, zero value otherwise.

### GetOutputKeysOk

`func (o *TaskDef) GetOutputKeysOk() (*[]string, bool)`

GetOutputKeysOk returns a tuple with the OutputKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputKeys

`func (o *TaskDef) SetOutputKeys(v []string)`

SetOutputKeys sets OutputKeys field to given value.

### HasOutputKeys

`func (o *TaskDef) HasOutputKeys() bool`

HasOutputKeys returns a boolean if a field has been set.

### GetOutputSchema

`func (o *TaskDef) GetOutputSchema() SchemaDef`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *TaskDef) GetOutputSchemaOk() (*SchemaDef, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *TaskDef) SetOutputSchema(v SchemaDef)`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *TaskDef) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetOwnerApp

`func (o *TaskDef) GetOwnerApp() string`

GetOwnerApp returns the OwnerApp field if non-nil, zero value otherwise.

### GetOwnerAppOk

`func (o *TaskDef) GetOwnerAppOk() (*string, bool)`

GetOwnerAppOk returns a tuple with the OwnerApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerApp

`func (o *TaskDef) SetOwnerApp(v string)`

SetOwnerApp sets OwnerApp field to given value.

### HasOwnerApp

`func (o *TaskDef) HasOwnerApp() bool`

HasOwnerApp returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *TaskDef) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *TaskDef) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *TaskDef) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *TaskDef) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetPollTimeoutSeconds

`func (o *TaskDef) GetPollTimeoutSeconds() int32`

GetPollTimeoutSeconds returns the PollTimeoutSeconds field if non-nil, zero value otherwise.

### GetPollTimeoutSecondsOk

`func (o *TaskDef) GetPollTimeoutSecondsOk() (*int32, bool)`

GetPollTimeoutSecondsOk returns a tuple with the PollTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollTimeoutSeconds

`func (o *TaskDef) SetPollTimeoutSeconds(v int32)`

SetPollTimeoutSeconds sets PollTimeoutSeconds field to given value.

### HasPollTimeoutSeconds

`func (o *TaskDef) HasPollTimeoutSeconds() bool`

HasPollTimeoutSeconds returns a boolean if a field has been set.

### GetRateLimitFrequencyInSeconds

`func (o *TaskDef) GetRateLimitFrequencyInSeconds() int32`

GetRateLimitFrequencyInSeconds returns the RateLimitFrequencyInSeconds field if non-nil, zero value otherwise.

### GetRateLimitFrequencyInSecondsOk

`func (o *TaskDef) GetRateLimitFrequencyInSecondsOk() (*int32, bool)`

GetRateLimitFrequencyInSecondsOk returns a tuple with the RateLimitFrequencyInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitFrequencyInSeconds

`func (o *TaskDef) SetRateLimitFrequencyInSeconds(v int32)`

SetRateLimitFrequencyInSeconds sets RateLimitFrequencyInSeconds field to given value.

### HasRateLimitFrequencyInSeconds

`func (o *TaskDef) HasRateLimitFrequencyInSeconds() bool`

HasRateLimitFrequencyInSeconds returns a boolean if a field has been set.

### GetRateLimitPerFrequency

`func (o *TaskDef) GetRateLimitPerFrequency() int32`

GetRateLimitPerFrequency returns the RateLimitPerFrequency field if non-nil, zero value otherwise.

### GetRateLimitPerFrequencyOk

`func (o *TaskDef) GetRateLimitPerFrequencyOk() (*int32, bool)`

GetRateLimitPerFrequencyOk returns a tuple with the RateLimitPerFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerFrequency

`func (o *TaskDef) SetRateLimitPerFrequency(v int32)`

SetRateLimitPerFrequency sets RateLimitPerFrequency field to given value.

### HasRateLimitPerFrequency

`func (o *TaskDef) HasRateLimitPerFrequency() bool`

HasRateLimitPerFrequency returns a boolean if a field has been set.

### GetResponseTimeoutSeconds

`func (o *TaskDef) GetResponseTimeoutSeconds() int64`

GetResponseTimeoutSeconds returns the ResponseTimeoutSeconds field if non-nil, zero value otherwise.

### GetResponseTimeoutSecondsOk

`func (o *TaskDef) GetResponseTimeoutSecondsOk() (*int64, bool)`

GetResponseTimeoutSecondsOk returns a tuple with the ResponseTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeoutSeconds

`func (o *TaskDef) SetResponseTimeoutSeconds(v int64)`

SetResponseTimeoutSeconds sets ResponseTimeoutSeconds field to given value.

### HasResponseTimeoutSeconds

`func (o *TaskDef) HasResponseTimeoutSeconds() bool`

HasResponseTimeoutSeconds returns a boolean if a field has been set.

### GetRetryCount

`func (o *TaskDef) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *TaskDef) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *TaskDef) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *TaskDef) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *TaskDef) GetRetryDelaySeconds() int32`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *TaskDef) GetRetryDelaySecondsOk() (*int32, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *TaskDef) SetRetryDelaySeconds(v int32)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *TaskDef) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetRetryLogic

`func (o *TaskDef) GetRetryLogic() string`

GetRetryLogic returns the RetryLogic field if non-nil, zero value otherwise.

### GetRetryLogicOk

`func (o *TaskDef) GetRetryLogicOk() (*string, bool)`

GetRetryLogicOk returns a tuple with the RetryLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryLogic

`func (o *TaskDef) SetRetryLogic(v string)`

SetRetryLogic sets RetryLogic field to given value.

### HasRetryLogic

`func (o *TaskDef) HasRetryLogic() bool`

HasRetryLogic returns a boolean if a field has been set.

### GetTimeoutPolicy

`func (o *TaskDef) GetTimeoutPolicy() string`

GetTimeoutPolicy returns the TimeoutPolicy field if non-nil, zero value otherwise.

### GetTimeoutPolicyOk

`func (o *TaskDef) GetTimeoutPolicyOk() (*string, bool)`

GetTimeoutPolicyOk returns a tuple with the TimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutPolicy

`func (o *TaskDef) SetTimeoutPolicy(v string)`

SetTimeoutPolicy sets TimeoutPolicy field to given value.

### HasTimeoutPolicy

`func (o *TaskDef) HasTimeoutPolicy() bool`

HasTimeoutPolicy returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *TaskDef) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *TaskDef) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *TaskDef) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### GetTotalTimeoutSeconds

`func (o *TaskDef) GetTotalTimeoutSeconds() int64`

GetTotalTimeoutSeconds returns the TotalTimeoutSeconds field if non-nil, zero value otherwise.

### GetTotalTimeoutSecondsOk

`func (o *TaskDef) GetTotalTimeoutSecondsOk() (*int64, bool)`

GetTotalTimeoutSecondsOk returns a tuple with the TotalTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeoutSeconds

`func (o *TaskDef) SetTotalTimeoutSeconds(v int64)`

SetTotalTimeoutSeconds sets TotalTimeoutSeconds field to given value.


### GetUpdateTime

`func (o *TaskDef) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *TaskDef) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *TaskDef) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *TaskDef) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *TaskDef) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *TaskDef) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *TaskDef) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *TaskDef) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


