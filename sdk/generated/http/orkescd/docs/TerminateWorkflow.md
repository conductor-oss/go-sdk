# TerminateWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TerminationReason** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 

## Methods

### NewTerminateWorkflow

`func NewTerminateWorkflow() *TerminateWorkflow`

NewTerminateWorkflow instantiates a new TerminateWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminateWorkflowWithDefaults

`func NewTerminateWorkflowWithDefaults() *TerminateWorkflow`

NewTerminateWorkflowWithDefaults instantiates a new TerminateWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerminationReason

`func (o *TerminateWorkflow) GetTerminationReason() string`

GetTerminationReason returns the TerminationReason field if non-nil, zero value otherwise.

### GetTerminationReasonOk

`func (o *TerminateWorkflow) GetTerminationReasonOk() (*string, bool)`

GetTerminationReasonOk returns a tuple with the TerminationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationReason

`func (o *TerminateWorkflow) SetTerminationReason(v string)`

SetTerminationReason sets TerminationReason field to given value.

### HasTerminationReason

`func (o *TerminateWorkflow) HasTerminationReason() bool`

HasTerminationReason returns a boolean if a field has been set.

### GetWorkflowId

`func (o *TerminateWorkflow) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *TerminateWorkflow) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *TerminateWorkflow) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *TerminateWorkflow) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


