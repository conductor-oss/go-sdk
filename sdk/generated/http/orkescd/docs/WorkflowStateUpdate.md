# WorkflowStateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskReferenceName** | Pointer to **string** |  | [optional] 
**TaskResult** | Pointer to [**TaskResult**](TaskResult.md) |  | [optional] 
**Variables** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewWorkflowStateUpdate

`func NewWorkflowStateUpdate() *WorkflowStateUpdate`

NewWorkflowStateUpdate instantiates a new WorkflowStateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStateUpdateWithDefaults

`func NewWorkflowStateUpdateWithDefaults() *WorkflowStateUpdate`

NewWorkflowStateUpdateWithDefaults instantiates a new WorkflowStateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskReferenceName

`func (o *WorkflowStateUpdate) GetTaskReferenceName() string`

GetTaskReferenceName returns the TaskReferenceName field if non-nil, zero value otherwise.

### GetTaskReferenceNameOk

`func (o *WorkflowStateUpdate) GetTaskReferenceNameOk() (*string, bool)`

GetTaskReferenceNameOk returns a tuple with the TaskReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskReferenceName

`func (o *WorkflowStateUpdate) SetTaskReferenceName(v string)`

SetTaskReferenceName sets TaskReferenceName field to given value.

### HasTaskReferenceName

`func (o *WorkflowStateUpdate) HasTaskReferenceName() bool`

HasTaskReferenceName returns a boolean if a field has been set.

### GetTaskResult

`func (o *WorkflowStateUpdate) GetTaskResult() TaskResult`

GetTaskResult returns the TaskResult field if non-nil, zero value otherwise.

### GetTaskResultOk

`func (o *WorkflowStateUpdate) GetTaskResultOk() (*TaskResult, bool)`

GetTaskResultOk returns a tuple with the TaskResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskResult

`func (o *WorkflowStateUpdate) SetTaskResult(v TaskResult)`

SetTaskResult sets TaskResult field to given value.

### HasTaskResult

`func (o *WorkflowStateUpdate) HasTaskResult() bool`

HasTaskResult returns a boolean if a field has been set.

### GetVariables

`func (o *WorkflowStateUpdate) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *WorkflowStateUpdate) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *WorkflowStateUpdate) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *WorkflowStateUpdate) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


