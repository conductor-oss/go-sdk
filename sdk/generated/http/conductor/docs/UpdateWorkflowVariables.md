# UpdateWorkflowVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**AppendArray** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateWorkflowVariables

`func NewUpdateWorkflowVariables() *UpdateWorkflowVariables`

NewUpdateWorkflowVariables instantiates a new UpdateWorkflowVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkflowVariablesWithDefaults

`func NewUpdateWorkflowVariablesWithDefaults() *UpdateWorkflowVariables`

NewUpdateWorkflowVariablesWithDefaults instantiates a new UpdateWorkflowVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *UpdateWorkflowVariables) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *UpdateWorkflowVariables) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *UpdateWorkflowVariables) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *UpdateWorkflowVariables) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetVariables

`func (o *UpdateWorkflowVariables) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *UpdateWorkflowVariables) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *UpdateWorkflowVariables) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *UpdateWorkflowVariables) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetAppendArray

`func (o *UpdateWorkflowVariables) GetAppendArray() bool`

GetAppendArray returns the AppendArray field if non-nil, zero value otherwise.

### GetAppendArrayOk

`func (o *UpdateWorkflowVariables) GetAppendArrayOk() (*bool, bool)`

GetAppendArrayOk returns a tuple with the AppendArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendArray

`func (o *UpdateWorkflowVariables) SetAppendArray(v bool)`

SetAppendArray sets AppendArray field to given value.

### HasAppendArray

`func (o *UpdateWorkflowVariables) HasAppendArray() bool`

HasAppendArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


