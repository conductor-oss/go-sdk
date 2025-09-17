# RerunWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** |  | [optional] 
**ReRunFromTaskId** | Pointer to **string** |  | [optional] 
**ReRunFromWorkflowId** | Pointer to **string** |  | [optional] 
**TaskInput** | Pointer to  |  | [optional] 
**WorkflowInput** | Pointer to  |  | [optional] 

## Methods

### NewRerunWorkflowRequest

`func NewRerunWorkflowRequest() *RerunWorkflowRequest`

NewRerunWorkflowRequest instantiates a new RerunWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRerunWorkflowRequestWithDefaults

`func NewRerunWorkflowRequestWithDefaults() *RerunWorkflowRequest`

NewRerunWorkflowRequestWithDefaults instantiates a new RerunWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *RerunWorkflowRequest) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *RerunWorkflowRequest) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *RerunWorkflowRequest) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *RerunWorkflowRequest) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetReRunFromTaskId

`func (o *RerunWorkflowRequest) GetReRunFromTaskId() string`

GetReRunFromTaskId returns the ReRunFromTaskId field if non-nil, zero value otherwise.

### GetReRunFromTaskIdOk

`func (o *RerunWorkflowRequest) GetReRunFromTaskIdOk() (*string, bool)`

GetReRunFromTaskIdOk returns a tuple with the ReRunFromTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReRunFromTaskId

`func (o *RerunWorkflowRequest) SetReRunFromTaskId(v string)`

SetReRunFromTaskId sets ReRunFromTaskId field to given value.

### HasReRunFromTaskId

`func (o *RerunWorkflowRequest) HasReRunFromTaskId() bool`

HasReRunFromTaskId returns a boolean if a field has been set.

### GetReRunFromWorkflowId

`func (o *RerunWorkflowRequest) GetReRunFromWorkflowId() string`

GetReRunFromWorkflowId returns the ReRunFromWorkflowId field if non-nil, zero value otherwise.

### GetReRunFromWorkflowIdOk

`func (o *RerunWorkflowRequest) GetReRunFromWorkflowIdOk() (*string, bool)`

GetReRunFromWorkflowIdOk returns a tuple with the ReRunFromWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReRunFromWorkflowId

`func (o *RerunWorkflowRequest) SetReRunFromWorkflowId(v string)`

SetReRunFromWorkflowId sets ReRunFromWorkflowId field to given value.

### HasReRunFromWorkflowId

`func (o *RerunWorkflowRequest) HasReRunFromWorkflowId() bool`

HasReRunFromWorkflowId returns a boolean if a field has been set.

### GetTaskInput

`func (o *RerunWorkflowRequest) GetTaskInput() map[string]map[string]interface{}`

GetTaskInput returns the TaskInput field if non-nil, zero value otherwise.

### GetTaskInputOk

`func (o *RerunWorkflowRequest) GetTaskInputOk() (*map[string]map[string]interface{}, bool)`

GetTaskInputOk returns a tuple with the TaskInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInput

`func (o *RerunWorkflowRequest) SetTaskInput(v map[string]map[string]interface{})`

SetTaskInput sets TaskInput field to given value.

### HasTaskInput

`func (o *RerunWorkflowRequest) HasTaskInput() bool`

HasTaskInput returns a boolean if a field has been set.

### SetTaskInputNil

`func (o *RerunWorkflowRequest) SetTaskInputNil(b bool)`

 SetTaskInputNil sets the value for TaskInput to be an explicit nil

### UnsetTaskInput
`func (o *RerunWorkflowRequest) UnsetTaskInput()`

UnsetTaskInput ensures that no value is present for TaskInput, not even an explicit nil
### GetWorkflowInput

`func (o *RerunWorkflowRequest) GetWorkflowInput() map[string]map[string]interface{}`

GetWorkflowInput returns the WorkflowInput field if non-nil, zero value otherwise.

### GetWorkflowInputOk

`func (o *RerunWorkflowRequest) GetWorkflowInputOk() (*map[string]map[string]interface{}, bool)`

GetWorkflowInputOk returns a tuple with the WorkflowInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInput

`func (o *RerunWorkflowRequest) SetWorkflowInput(v map[string]map[string]interface{})`

SetWorkflowInput sets WorkflowInput field to given value.

### HasWorkflowInput

`func (o *RerunWorkflowRequest) HasWorkflowInput() bool`

HasWorkflowInput returns a boolean if a field has been set.

### SetWorkflowInputNil

`func (o *RerunWorkflowRequest) SetWorkflowInputNil(b bool)`

 SetWorkflowInputNil sets the value for WorkflowInput to be an explicit nil

### UnsetWorkflowInput
`func (o *RerunWorkflowRequest) UnsetWorkflowInput()`

UnsetWorkflowInput ensures that no value is present for WorkflowInput, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


