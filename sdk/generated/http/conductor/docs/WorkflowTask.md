# WorkflowTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**TaskReferenceName** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**InputParameters** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DynamicTaskNameParam** | Pointer to **string** |  | [optional] 
**CaseValueParam** | Pointer to **string** |  | [optional] 
**CaseExpression** | Pointer to **string** |  | [optional] 
**ScriptExpression** | Pointer to **string** |  | [optional] 
**DecisionCases** | Pointer to [**map[string][]WorkflowTask**](array.md) |  | [optional] 
**DynamicForkJoinTasksParam** | Pointer to **string** |  | [optional] 
**DynamicForkTasksParam** | Pointer to **string** |  | [optional] 
**DynamicForkTasksInputParamName** | Pointer to **string** |  | [optional] 
**DefaultCase** | Pointer to [**[]WorkflowTask**](WorkflowTask.md) |  | [optional] 
**ForkTasks** | Pointer to [**[][]WorkflowTask**]([]WorkflowTask.md) |  | [optional] 
**StartDelay** | Pointer to **int32** |  | [optional] 
**SubWorkflowParam** | Pointer to [**SubWorkflowParams**](SubWorkflowParams.md) |  | [optional] 
**JoinOn** | Pointer to **[]string** |  | [optional] 
**Sink** | Pointer to **string** |  | [optional] 
**Optional** | Pointer to **bool** |  | [optional] 
**TaskDefinition** | Pointer to [**TaskDef**](TaskDef.md) |  | [optional] 
**RateLimited** | Pointer to **bool** |  | [optional] 
**DefaultExclusiveJoinTask** | Pointer to **[]string** |  | [optional] 
**AsyncComplete** | Pointer to **bool** |  | [optional] 
**LoopCondition** | Pointer to **string** |  | [optional] 
**LoopOver** | Pointer to [**[]WorkflowTask**](WorkflowTask.md) |  | [optional] 
**RetryCount** | Pointer to **int32** |  | [optional] 
**EvaluatorType** | Pointer to **string** |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 
**OnStateChange** | Pointer to [**map[string][]StateChangeEvent**](array.md) |  | [optional] 
**JoinStatus** | Pointer to **string** |  | [optional] 
**CacheConfig** | Pointer to [**CacheConfig**](CacheConfig.md) |  | [optional] 
**Permissive** | Pointer to **bool** |  | [optional] 
**WorkflowTaskType** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowTask

`func NewWorkflowTask(name string, taskReferenceName string, ) *WorkflowTask`

NewWorkflowTask instantiates a new WorkflowTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskWithDefaults

`func NewWorkflowTaskWithDefaults() *WorkflowTask`

NewWorkflowTaskWithDefaults instantiates a new WorkflowTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkflowTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTask) SetName(v string)`

SetName sets Name field to given value.


### GetTaskReferenceName

`func (o *WorkflowTask) GetTaskReferenceName() string`

GetTaskReferenceName returns the TaskReferenceName field if non-nil, zero value otherwise.

### GetTaskReferenceNameOk

`func (o *WorkflowTask) GetTaskReferenceNameOk() (*string, bool)`

GetTaskReferenceNameOk returns a tuple with the TaskReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskReferenceName

`func (o *WorkflowTask) SetTaskReferenceName(v string)`

SetTaskReferenceName sets TaskReferenceName field to given value.


### GetDescription

`func (o *WorkflowTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputParameters

`func (o *WorkflowTask) GetInputParameters() map[string]map[string]interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowTask) GetInputParametersOk() (*map[string]map[string]interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowTask) SetInputParameters(v map[string]map[string]interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowTask) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### GetType

`func (o *WorkflowTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowTask) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDynamicTaskNameParam

`func (o *WorkflowTask) GetDynamicTaskNameParam() string`

GetDynamicTaskNameParam returns the DynamicTaskNameParam field if non-nil, zero value otherwise.

### GetDynamicTaskNameParamOk

`func (o *WorkflowTask) GetDynamicTaskNameParamOk() (*string, bool)`

GetDynamicTaskNameParamOk returns a tuple with the DynamicTaskNameParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicTaskNameParam

`func (o *WorkflowTask) SetDynamicTaskNameParam(v string)`

SetDynamicTaskNameParam sets DynamicTaskNameParam field to given value.

### HasDynamicTaskNameParam

`func (o *WorkflowTask) HasDynamicTaskNameParam() bool`

HasDynamicTaskNameParam returns a boolean if a field has been set.

### GetCaseValueParam

`func (o *WorkflowTask) GetCaseValueParam() string`

GetCaseValueParam returns the CaseValueParam field if non-nil, zero value otherwise.

### GetCaseValueParamOk

`func (o *WorkflowTask) GetCaseValueParamOk() (*string, bool)`

GetCaseValueParamOk returns a tuple with the CaseValueParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseValueParam

`func (o *WorkflowTask) SetCaseValueParam(v string)`

SetCaseValueParam sets CaseValueParam field to given value.

### HasCaseValueParam

`func (o *WorkflowTask) HasCaseValueParam() bool`

HasCaseValueParam returns a boolean if a field has been set.

### GetCaseExpression

`func (o *WorkflowTask) GetCaseExpression() string`

GetCaseExpression returns the CaseExpression field if non-nil, zero value otherwise.

### GetCaseExpressionOk

`func (o *WorkflowTask) GetCaseExpressionOk() (*string, bool)`

GetCaseExpressionOk returns a tuple with the CaseExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExpression

`func (o *WorkflowTask) SetCaseExpression(v string)`

SetCaseExpression sets CaseExpression field to given value.

### HasCaseExpression

`func (o *WorkflowTask) HasCaseExpression() bool`

HasCaseExpression returns a boolean if a field has been set.

### GetScriptExpression

`func (o *WorkflowTask) GetScriptExpression() string`

GetScriptExpression returns the ScriptExpression field if non-nil, zero value otherwise.

### GetScriptExpressionOk

`func (o *WorkflowTask) GetScriptExpressionOk() (*string, bool)`

GetScriptExpressionOk returns a tuple with the ScriptExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptExpression

`func (o *WorkflowTask) SetScriptExpression(v string)`

SetScriptExpression sets ScriptExpression field to given value.

### HasScriptExpression

`func (o *WorkflowTask) HasScriptExpression() bool`

HasScriptExpression returns a boolean if a field has been set.

### GetDecisionCases

`func (o *WorkflowTask) GetDecisionCases() map[string][]WorkflowTask`

GetDecisionCases returns the DecisionCases field if non-nil, zero value otherwise.

### GetDecisionCasesOk

`func (o *WorkflowTask) GetDecisionCasesOk() (*map[string][]WorkflowTask, bool)`

GetDecisionCasesOk returns a tuple with the DecisionCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionCases

`func (o *WorkflowTask) SetDecisionCases(v map[string][]WorkflowTask)`

SetDecisionCases sets DecisionCases field to given value.

### HasDecisionCases

`func (o *WorkflowTask) HasDecisionCases() bool`

HasDecisionCases returns a boolean if a field has been set.

### GetDynamicForkJoinTasksParam

`func (o *WorkflowTask) GetDynamicForkJoinTasksParam() string`

GetDynamicForkJoinTasksParam returns the DynamicForkJoinTasksParam field if non-nil, zero value otherwise.

### GetDynamicForkJoinTasksParamOk

`func (o *WorkflowTask) GetDynamicForkJoinTasksParamOk() (*string, bool)`

GetDynamicForkJoinTasksParamOk returns a tuple with the DynamicForkJoinTasksParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicForkJoinTasksParam

`func (o *WorkflowTask) SetDynamicForkJoinTasksParam(v string)`

SetDynamicForkJoinTasksParam sets DynamicForkJoinTasksParam field to given value.

### HasDynamicForkJoinTasksParam

`func (o *WorkflowTask) HasDynamicForkJoinTasksParam() bool`

HasDynamicForkJoinTasksParam returns a boolean if a field has been set.

### GetDynamicForkTasksParam

`func (o *WorkflowTask) GetDynamicForkTasksParam() string`

GetDynamicForkTasksParam returns the DynamicForkTasksParam field if non-nil, zero value otherwise.

### GetDynamicForkTasksParamOk

`func (o *WorkflowTask) GetDynamicForkTasksParamOk() (*string, bool)`

GetDynamicForkTasksParamOk returns a tuple with the DynamicForkTasksParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicForkTasksParam

`func (o *WorkflowTask) SetDynamicForkTasksParam(v string)`

SetDynamicForkTasksParam sets DynamicForkTasksParam field to given value.

### HasDynamicForkTasksParam

`func (o *WorkflowTask) HasDynamicForkTasksParam() bool`

HasDynamicForkTasksParam returns a boolean if a field has been set.

### GetDynamicForkTasksInputParamName

`func (o *WorkflowTask) GetDynamicForkTasksInputParamName() string`

GetDynamicForkTasksInputParamName returns the DynamicForkTasksInputParamName field if non-nil, zero value otherwise.

### GetDynamicForkTasksInputParamNameOk

`func (o *WorkflowTask) GetDynamicForkTasksInputParamNameOk() (*string, bool)`

GetDynamicForkTasksInputParamNameOk returns a tuple with the DynamicForkTasksInputParamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicForkTasksInputParamName

`func (o *WorkflowTask) SetDynamicForkTasksInputParamName(v string)`

SetDynamicForkTasksInputParamName sets DynamicForkTasksInputParamName field to given value.

### HasDynamicForkTasksInputParamName

`func (o *WorkflowTask) HasDynamicForkTasksInputParamName() bool`

HasDynamicForkTasksInputParamName returns a boolean if a field has been set.

### GetDefaultCase

`func (o *WorkflowTask) GetDefaultCase() []WorkflowTask`

GetDefaultCase returns the DefaultCase field if non-nil, zero value otherwise.

### GetDefaultCaseOk

`func (o *WorkflowTask) GetDefaultCaseOk() (*[]WorkflowTask, bool)`

GetDefaultCaseOk returns a tuple with the DefaultCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCase

`func (o *WorkflowTask) SetDefaultCase(v []WorkflowTask)`

SetDefaultCase sets DefaultCase field to given value.

### HasDefaultCase

`func (o *WorkflowTask) HasDefaultCase() bool`

HasDefaultCase returns a boolean if a field has been set.

### GetForkTasks

`func (o *WorkflowTask) GetForkTasks() [][]WorkflowTask`

GetForkTasks returns the ForkTasks field if non-nil, zero value otherwise.

### GetForkTasksOk

`func (o *WorkflowTask) GetForkTasksOk() (*[][]WorkflowTask, bool)`

GetForkTasksOk returns a tuple with the ForkTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkTasks

`func (o *WorkflowTask) SetForkTasks(v [][]WorkflowTask)`

SetForkTasks sets ForkTasks field to given value.

### HasForkTasks

`func (o *WorkflowTask) HasForkTasks() bool`

HasForkTasks returns a boolean if a field has been set.

### GetStartDelay

`func (o *WorkflowTask) GetStartDelay() int32`

GetStartDelay returns the StartDelay field if non-nil, zero value otherwise.

### GetStartDelayOk

`func (o *WorkflowTask) GetStartDelayOk() (*int32, bool)`

GetStartDelayOk returns a tuple with the StartDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDelay

`func (o *WorkflowTask) SetStartDelay(v int32)`

SetStartDelay sets StartDelay field to given value.

### HasStartDelay

`func (o *WorkflowTask) HasStartDelay() bool`

HasStartDelay returns a boolean if a field has been set.

### GetSubWorkflowParam

`func (o *WorkflowTask) GetSubWorkflowParam() SubWorkflowParams`

GetSubWorkflowParam returns the SubWorkflowParam field if non-nil, zero value otherwise.

### GetSubWorkflowParamOk

`func (o *WorkflowTask) GetSubWorkflowParamOk() (*SubWorkflowParams, bool)`

GetSubWorkflowParamOk returns a tuple with the SubWorkflowParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubWorkflowParam

`func (o *WorkflowTask) SetSubWorkflowParam(v SubWorkflowParams)`

SetSubWorkflowParam sets SubWorkflowParam field to given value.

### HasSubWorkflowParam

`func (o *WorkflowTask) HasSubWorkflowParam() bool`

HasSubWorkflowParam returns a boolean if a field has been set.

### GetJoinOn

`func (o *WorkflowTask) GetJoinOn() []string`

GetJoinOn returns the JoinOn field if non-nil, zero value otherwise.

### GetJoinOnOk

`func (o *WorkflowTask) GetJoinOnOk() (*[]string, bool)`

GetJoinOnOk returns a tuple with the JoinOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinOn

`func (o *WorkflowTask) SetJoinOn(v []string)`

SetJoinOn sets JoinOn field to given value.

### HasJoinOn

`func (o *WorkflowTask) HasJoinOn() bool`

HasJoinOn returns a boolean if a field has been set.

### GetSink

`func (o *WorkflowTask) GetSink() string`

GetSink returns the Sink field if non-nil, zero value otherwise.

### GetSinkOk

`func (o *WorkflowTask) GetSinkOk() (*string, bool)`

GetSinkOk returns a tuple with the Sink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSink

`func (o *WorkflowTask) SetSink(v string)`

SetSink sets Sink field to given value.

### HasSink

`func (o *WorkflowTask) HasSink() bool`

HasSink returns a boolean if a field has been set.

### GetOptional

`func (o *WorkflowTask) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *WorkflowTask) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *WorkflowTask) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *WorkflowTask) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetTaskDefinition

`func (o *WorkflowTask) GetTaskDefinition() TaskDef`

GetTaskDefinition returns the TaskDefinition field if non-nil, zero value otherwise.

### GetTaskDefinitionOk

`func (o *WorkflowTask) GetTaskDefinitionOk() (*TaskDef, bool)`

GetTaskDefinitionOk returns a tuple with the TaskDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinition

`func (o *WorkflowTask) SetTaskDefinition(v TaskDef)`

SetTaskDefinition sets TaskDefinition field to given value.

### HasTaskDefinition

`func (o *WorkflowTask) HasTaskDefinition() bool`

HasTaskDefinition returns a boolean if a field has been set.

### GetRateLimited

`func (o *WorkflowTask) GetRateLimited() bool`

GetRateLimited returns the RateLimited field if non-nil, zero value otherwise.

### GetRateLimitedOk

`func (o *WorkflowTask) GetRateLimitedOk() (*bool, bool)`

GetRateLimitedOk returns a tuple with the RateLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimited

`func (o *WorkflowTask) SetRateLimited(v bool)`

SetRateLimited sets RateLimited field to given value.

### HasRateLimited

`func (o *WorkflowTask) HasRateLimited() bool`

HasRateLimited returns a boolean if a field has been set.

### GetDefaultExclusiveJoinTask

`func (o *WorkflowTask) GetDefaultExclusiveJoinTask() []string`

GetDefaultExclusiveJoinTask returns the DefaultExclusiveJoinTask field if non-nil, zero value otherwise.

### GetDefaultExclusiveJoinTaskOk

`func (o *WorkflowTask) GetDefaultExclusiveJoinTaskOk() (*[]string, bool)`

GetDefaultExclusiveJoinTaskOk returns a tuple with the DefaultExclusiveJoinTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExclusiveJoinTask

`func (o *WorkflowTask) SetDefaultExclusiveJoinTask(v []string)`

SetDefaultExclusiveJoinTask sets DefaultExclusiveJoinTask field to given value.

### HasDefaultExclusiveJoinTask

`func (o *WorkflowTask) HasDefaultExclusiveJoinTask() bool`

HasDefaultExclusiveJoinTask returns a boolean if a field has been set.

### GetAsyncComplete

`func (o *WorkflowTask) GetAsyncComplete() bool`

GetAsyncComplete returns the AsyncComplete field if non-nil, zero value otherwise.

### GetAsyncCompleteOk

`func (o *WorkflowTask) GetAsyncCompleteOk() (*bool, bool)`

GetAsyncCompleteOk returns a tuple with the AsyncComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncComplete

`func (o *WorkflowTask) SetAsyncComplete(v bool)`

SetAsyncComplete sets AsyncComplete field to given value.

### HasAsyncComplete

`func (o *WorkflowTask) HasAsyncComplete() bool`

HasAsyncComplete returns a boolean if a field has been set.

### GetLoopCondition

`func (o *WorkflowTask) GetLoopCondition() string`

GetLoopCondition returns the LoopCondition field if non-nil, zero value otherwise.

### GetLoopConditionOk

`func (o *WorkflowTask) GetLoopConditionOk() (*string, bool)`

GetLoopConditionOk returns a tuple with the LoopCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopCondition

`func (o *WorkflowTask) SetLoopCondition(v string)`

SetLoopCondition sets LoopCondition field to given value.

### HasLoopCondition

`func (o *WorkflowTask) HasLoopCondition() bool`

HasLoopCondition returns a boolean if a field has been set.

### GetLoopOver

`func (o *WorkflowTask) GetLoopOver() []WorkflowTask`

GetLoopOver returns the LoopOver field if non-nil, zero value otherwise.

### GetLoopOverOk

`func (o *WorkflowTask) GetLoopOverOk() (*[]WorkflowTask, bool)`

GetLoopOverOk returns a tuple with the LoopOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopOver

`func (o *WorkflowTask) SetLoopOver(v []WorkflowTask)`

SetLoopOver sets LoopOver field to given value.

### HasLoopOver

`func (o *WorkflowTask) HasLoopOver() bool`

HasLoopOver returns a boolean if a field has been set.

### GetRetryCount

`func (o *WorkflowTask) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *WorkflowTask) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *WorkflowTask) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *WorkflowTask) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetEvaluatorType

`func (o *WorkflowTask) GetEvaluatorType() string`

GetEvaluatorType returns the EvaluatorType field if non-nil, zero value otherwise.

### GetEvaluatorTypeOk

`func (o *WorkflowTask) GetEvaluatorTypeOk() (*string, bool)`

GetEvaluatorTypeOk returns a tuple with the EvaluatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorType

`func (o *WorkflowTask) SetEvaluatorType(v string)`

SetEvaluatorType sets EvaluatorType field to given value.

### HasEvaluatorType

`func (o *WorkflowTask) HasEvaluatorType() bool`

HasEvaluatorType returns a boolean if a field has been set.

### GetExpression

`func (o *WorkflowTask) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *WorkflowTask) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *WorkflowTask) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *WorkflowTask) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetOnStateChange

`func (o *WorkflowTask) GetOnStateChange() map[string][]StateChangeEvent`

GetOnStateChange returns the OnStateChange field if non-nil, zero value otherwise.

### GetOnStateChangeOk

`func (o *WorkflowTask) GetOnStateChangeOk() (*map[string][]StateChangeEvent, bool)`

GetOnStateChangeOk returns a tuple with the OnStateChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnStateChange

`func (o *WorkflowTask) SetOnStateChange(v map[string][]StateChangeEvent)`

SetOnStateChange sets OnStateChange field to given value.

### HasOnStateChange

`func (o *WorkflowTask) HasOnStateChange() bool`

HasOnStateChange returns a boolean if a field has been set.

### GetJoinStatus

`func (o *WorkflowTask) GetJoinStatus() string`

GetJoinStatus returns the JoinStatus field if non-nil, zero value otherwise.

### GetJoinStatusOk

`func (o *WorkflowTask) GetJoinStatusOk() (*string, bool)`

GetJoinStatusOk returns a tuple with the JoinStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinStatus

`func (o *WorkflowTask) SetJoinStatus(v string)`

SetJoinStatus sets JoinStatus field to given value.

### HasJoinStatus

`func (o *WorkflowTask) HasJoinStatus() bool`

HasJoinStatus returns a boolean if a field has been set.

### GetCacheConfig

`func (o *WorkflowTask) GetCacheConfig() CacheConfig`

GetCacheConfig returns the CacheConfig field if non-nil, zero value otherwise.

### GetCacheConfigOk

`func (o *WorkflowTask) GetCacheConfigOk() (*CacheConfig, bool)`

GetCacheConfigOk returns a tuple with the CacheConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheConfig

`func (o *WorkflowTask) SetCacheConfig(v CacheConfig)`

SetCacheConfig sets CacheConfig field to given value.

### HasCacheConfig

`func (o *WorkflowTask) HasCacheConfig() bool`

HasCacheConfig returns a boolean if a field has been set.

### GetPermissive

`func (o *WorkflowTask) GetPermissive() bool`

GetPermissive returns the Permissive field if non-nil, zero value otherwise.

### GetPermissiveOk

`func (o *WorkflowTask) GetPermissiveOk() (*bool, bool)`

GetPermissiveOk returns a tuple with the Permissive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissive

`func (o *WorkflowTask) SetPermissive(v bool)`

SetPermissive sets Permissive field to given value.

### HasPermissive

`func (o *WorkflowTask) HasPermissive() bool`

HasPermissive returns a boolean if a field has been set.

### GetWorkflowTaskType

`func (o *WorkflowTask) GetWorkflowTaskType() string`

GetWorkflowTaskType returns the WorkflowTaskType field if non-nil, zero value otherwise.

### GetWorkflowTaskTypeOk

`func (o *WorkflowTask) GetWorkflowTaskTypeOk() (*string, bool)`

GetWorkflowTaskTypeOk returns a tuple with the WorkflowTaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskType

`func (o *WorkflowTask) SetWorkflowTaskType(v string)`

SetWorkflowTaskType sets WorkflowTaskType field to given value.

### HasWorkflowTaskType

`func (o *WorkflowTask) HasWorkflowTaskType() bool`

HasWorkflowTaskType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


