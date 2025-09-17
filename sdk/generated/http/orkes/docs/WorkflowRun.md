# WorkflowRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Output** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**ResponseType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TargetWorkflowId** | Pointer to **string** |  | [optional] 
**TargetWorkflowStatus** | Pointer to **string** |  | [optional] 
**Tasks** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**Variables** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowRun

`func NewWorkflowRun() *WorkflowRun`

NewWorkflowRun instantiates a new WorkflowRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunWithDefaults

`func NewWorkflowRunWithDefaults() *WorkflowRun`

NewWorkflowRunWithDefaults instantiates a new WorkflowRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *WorkflowRun) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *WorkflowRun) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *WorkflowRun) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *WorkflowRun) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetCreateTime

`func (o *WorkflowRun) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *WorkflowRun) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *WorkflowRun) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *WorkflowRun) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WorkflowRun) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkflowRun) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkflowRun) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkflowRun) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowRun) GetInput() map[string]map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowRun) GetInputOk() (*map[string]map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowRun) SetInput(v map[string]map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowRun) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowRun) GetOutput() map[string]map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowRun) GetOutputOk() (*map[string]map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowRun) SetOutput(v map[string]map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowRun) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetPriority

`func (o *WorkflowRun) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkflowRun) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkflowRun) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *WorkflowRun) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRequestId

`func (o *WorkflowRun) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *WorkflowRun) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *WorkflowRun) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *WorkflowRun) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetResponseType

`func (o *WorkflowRun) GetResponseType() string`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *WorkflowRun) GetResponseTypeOk() (*string, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *WorkflowRun) SetResponseType(v string)`

SetResponseType sets ResponseType field to given value.

### HasResponseType

`func (o *WorkflowRun) HasResponseType() bool`

HasResponseType returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetWorkflowId

`func (o *WorkflowRun) GetTargetWorkflowId() string`

GetTargetWorkflowId returns the TargetWorkflowId field if non-nil, zero value otherwise.

### GetTargetWorkflowIdOk

`func (o *WorkflowRun) GetTargetWorkflowIdOk() (*string, bool)`

GetTargetWorkflowIdOk returns a tuple with the TargetWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetWorkflowId

`func (o *WorkflowRun) SetTargetWorkflowId(v string)`

SetTargetWorkflowId sets TargetWorkflowId field to given value.

### HasTargetWorkflowId

`func (o *WorkflowRun) HasTargetWorkflowId() bool`

HasTargetWorkflowId returns a boolean if a field has been set.

### GetTargetWorkflowStatus

`func (o *WorkflowRun) GetTargetWorkflowStatus() string`

GetTargetWorkflowStatus returns the TargetWorkflowStatus field if non-nil, zero value otherwise.

### GetTargetWorkflowStatusOk

`func (o *WorkflowRun) GetTargetWorkflowStatusOk() (*string, bool)`

GetTargetWorkflowStatusOk returns a tuple with the TargetWorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetWorkflowStatus

`func (o *WorkflowRun) SetTargetWorkflowStatus(v string)`

SetTargetWorkflowStatus sets TargetWorkflowStatus field to given value.

### HasTargetWorkflowStatus

`func (o *WorkflowRun) HasTargetWorkflowStatus() bool`

HasTargetWorkflowStatus returns a boolean if a field has been set.

### GetTasks

`func (o *WorkflowRun) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *WorkflowRun) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *WorkflowRun) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *WorkflowRun) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetUpdateTime

`func (o *WorkflowRun) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *WorkflowRun) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *WorkflowRun) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *WorkflowRun) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetVariables

`func (o *WorkflowRun) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *WorkflowRun) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *WorkflowRun) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *WorkflowRun) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowRun) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowRun) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowRun) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowRun) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


