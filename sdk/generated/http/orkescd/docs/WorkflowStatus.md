# WorkflowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** |  | [optional] 
**Output** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowStatus

`func NewWorkflowStatus() *WorkflowStatus`

NewWorkflowStatus instantiates a new WorkflowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStatusWithDefaults

`func NewWorkflowStatusWithDefaults() *WorkflowStatus`

NewWorkflowStatusWithDefaults instantiates a new WorkflowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *WorkflowStatus) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *WorkflowStatus) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *WorkflowStatus) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *WorkflowStatus) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowStatus) GetOutput() map[string]map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowStatus) GetOutputOk() (*map[string]map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowStatus) SetOutput(v map[string]map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowStatus) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVariables

`func (o *WorkflowStatus) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *WorkflowStatus) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *WorkflowStatus) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *WorkflowStatus) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowStatus) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowStatus) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowStatus) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowStatus) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


