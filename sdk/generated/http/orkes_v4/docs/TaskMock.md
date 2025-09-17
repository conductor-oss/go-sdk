# TaskMock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionTime** | Pointer to **int64** |  | [optional] 
**Output** | Pointer to  |  | [optional] 
**QueueWaitTime** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskMock

`func NewTaskMock() *TaskMock`

NewTaskMock instantiates a new TaskMock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskMockWithDefaults

`func NewTaskMockWithDefaults() *TaskMock`

NewTaskMockWithDefaults instantiates a new TaskMock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionTime

`func (o *TaskMock) GetExecutionTime() int64`

GetExecutionTime returns the ExecutionTime field if non-nil, zero value otherwise.

### GetExecutionTimeOk

`func (o *TaskMock) GetExecutionTimeOk() (*int64, bool)`

GetExecutionTimeOk returns a tuple with the ExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTime

`func (o *TaskMock) SetExecutionTime(v int64)`

SetExecutionTime sets ExecutionTime field to given value.

### HasExecutionTime

`func (o *TaskMock) HasExecutionTime() bool`

HasExecutionTime returns a boolean if a field has been set.

### GetOutput

`func (o *TaskMock) GetOutput() map[string]map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *TaskMock) GetOutputOk() (*map[string]map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *TaskMock) SetOutput(v map[string]map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *TaskMock) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *TaskMock) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *TaskMock) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetQueueWaitTime

`func (o *TaskMock) GetQueueWaitTime() int64`

GetQueueWaitTime returns the QueueWaitTime field if non-nil, zero value otherwise.

### GetQueueWaitTimeOk

`func (o *TaskMock) GetQueueWaitTimeOk() (*int64, bool)`

GetQueueWaitTimeOk returns a tuple with the QueueWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueWaitTime

`func (o *TaskMock) SetQueueWaitTime(v int64)`

SetQueueWaitTime sets QueueWaitTime field to given value.

### HasQueueWaitTime

`func (o *TaskMock) HasQueueWaitTime() bool`

HasQueueWaitTime returns a boolean if a field has been set.

### GetStatus

`func (o *TaskMock) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskMock) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskMock) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskMock) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


