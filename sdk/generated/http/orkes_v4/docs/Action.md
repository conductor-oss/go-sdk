# Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**CompleteTask** | Pointer to [**TaskDetails**](TaskDetails.md) |  | [optional] 
**ExpandInlineJSON** | Pointer to **bool** |  | [optional] 
**FailTask** | Pointer to [**TaskDetails**](TaskDetails.md) |  | [optional] 
**StartWorkflow** | Pointer to [**StartWorkflowRequest**](StartWorkflowRequest.md) |  | [optional] 
**TerminateWorkflow** | Pointer to [**TerminateWorkflow**](TerminateWorkflow.md) |  | [optional] 
**UpdateWorkflowVariables** | Pointer to [**UpdateWorkflowVariables**](UpdateWorkflowVariables.md) |  | [optional] 

## Methods

### NewAction

`func NewAction() *Action`

NewAction instantiates a new Action object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionWithDefaults

`func NewActionWithDefaults() *Action`

NewActionWithDefaults instantiates a new Action object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Action) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Action) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Action) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Action) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCompleteTask

`func (o *Action) GetCompleteTask() TaskDetails`

GetCompleteTask returns the CompleteTask field if non-nil, zero value otherwise.

### GetCompleteTaskOk

`func (o *Action) GetCompleteTaskOk() (*TaskDetails, bool)`

GetCompleteTaskOk returns a tuple with the CompleteTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteTask

`func (o *Action) SetCompleteTask(v TaskDetails)`

SetCompleteTask sets CompleteTask field to given value.

### HasCompleteTask

`func (o *Action) HasCompleteTask() bool`

HasCompleteTask returns a boolean if a field has been set.

### GetExpandInlineJSON

`func (o *Action) GetExpandInlineJSON() bool`

GetExpandInlineJSON returns the ExpandInlineJSON field if non-nil, zero value otherwise.

### GetExpandInlineJSONOk

`func (o *Action) GetExpandInlineJSONOk() (*bool, bool)`

GetExpandInlineJSONOk returns a tuple with the ExpandInlineJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandInlineJSON

`func (o *Action) SetExpandInlineJSON(v bool)`

SetExpandInlineJSON sets ExpandInlineJSON field to given value.

### HasExpandInlineJSON

`func (o *Action) HasExpandInlineJSON() bool`

HasExpandInlineJSON returns a boolean if a field has been set.

### GetFailTask

`func (o *Action) GetFailTask() TaskDetails`

GetFailTask returns the FailTask field if non-nil, zero value otherwise.

### GetFailTaskOk

`func (o *Action) GetFailTaskOk() (*TaskDetails, bool)`

GetFailTaskOk returns a tuple with the FailTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailTask

`func (o *Action) SetFailTask(v TaskDetails)`

SetFailTask sets FailTask field to given value.

### HasFailTask

`func (o *Action) HasFailTask() bool`

HasFailTask returns a boolean if a field has been set.

### GetStartWorkflow

`func (o *Action) GetStartWorkflow() StartWorkflowRequest`

GetStartWorkflow returns the StartWorkflow field if non-nil, zero value otherwise.

### GetStartWorkflowOk

`func (o *Action) GetStartWorkflowOk() (*StartWorkflowRequest, bool)`

GetStartWorkflowOk returns a tuple with the StartWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWorkflow

`func (o *Action) SetStartWorkflow(v StartWorkflowRequest)`

SetStartWorkflow sets StartWorkflow field to given value.

### HasStartWorkflow

`func (o *Action) HasStartWorkflow() bool`

HasStartWorkflow returns a boolean if a field has been set.

### GetTerminateWorkflow

`func (o *Action) GetTerminateWorkflow() TerminateWorkflow`

GetTerminateWorkflow returns the TerminateWorkflow field if non-nil, zero value otherwise.

### GetTerminateWorkflowOk

`func (o *Action) GetTerminateWorkflowOk() (*TerminateWorkflow, bool)`

GetTerminateWorkflowOk returns a tuple with the TerminateWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminateWorkflow

`func (o *Action) SetTerminateWorkflow(v TerminateWorkflow)`

SetTerminateWorkflow sets TerminateWorkflow field to given value.

### HasTerminateWorkflow

`func (o *Action) HasTerminateWorkflow() bool`

HasTerminateWorkflow returns a boolean if a field has been set.

### GetUpdateWorkflowVariables

`func (o *Action) GetUpdateWorkflowVariables() UpdateWorkflowVariables`

GetUpdateWorkflowVariables returns the UpdateWorkflowVariables field if non-nil, zero value otherwise.

### GetUpdateWorkflowVariablesOk

`func (o *Action) GetUpdateWorkflowVariablesOk() (*UpdateWorkflowVariables, bool)`

GetUpdateWorkflowVariablesOk returns a tuple with the UpdateWorkflowVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateWorkflowVariables

`func (o *Action) SetUpdateWorkflowVariables(v UpdateWorkflowVariables)`

SetUpdateWorkflowVariables sets UpdateWorkflowVariables field to given value.

### HasUpdateWorkflowVariables

`func (o *Action) HasUpdateWorkflowVariables() bool`

HasUpdateWorkflowVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


