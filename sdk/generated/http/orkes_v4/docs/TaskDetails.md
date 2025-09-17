# TaskDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Output** | Pointer to  |  | [optional] 
**OutputMessage** | Pointer to [**Any**](Any.md) |  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] 
**TaskRefName** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskDetails

`func NewTaskDetails() *TaskDetails`

NewTaskDetails instantiates a new TaskDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskDetailsWithDefaults

`func NewTaskDetailsWithDefaults() *TaskDetails`

NewTaskDetailsWithDefaults instantiates a new TaskDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutput

`func (o *TaskDetails) GetOutput() map[string]map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *TaskDetails) GetOutputOk() (*map[string]map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *TaskDetails) SetOutput(v map[string]map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *TaskDetails) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *TaskDetails) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *TaskDetails) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetOutputMessage

`func (o *TaskDetails) GetOutputMessage() Any`

GetOutputMessage returns the OutputMessage field if non-nil, zero value otherwise.

### GetOutputMessageOk

`func (o *TaskDetails) GetOutputMessageOk() (*Any, bool)`

GetOutputMessageOk returns a tuple with the OutputMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputMessage

`func (o *TaskDetails) SetOutputMessage(v Any)`

SetOutputMessage sets OutputMessage field to given value.

### HasOutputMessage

`func (o *TaskDetails) HasOutputMessage() bool`

HasOutputMessage returns a boolean if a field has been set.

### GetTaskId

`func (o *TaskDetails) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskDetails) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskDetails) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TaskDetails) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTaskRefName

`func (o *TaskDetails) GetTaskRefName() string`

GetTaskRefName returns the TaskRefName field if non-nil, zero value otherwise.

### GetTaskRefNameOk

`func (o *TaskDetails) GetTaskRefNameOk() (*string, bool)`

GetTaskRefNameOk returns a tuple with the TaskRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRefName

`func (o *TaskDetails) SetTaskRefName(v string)`

SetTaskRefName sets TaskRefName field to given value.

### HasTaskRefName

`func (o *TaskDetails) HasTaskRefName() bool`

HasTaskRefName returns a boolean if a field has been set.

### GetWorkflowId

`func (o *TaskDetails) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *TaskDetails) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *TaskDetails) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *TaskDetails) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


