# ConnectivityTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Successful** | Pointer to **bool** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectivityTestResult

`func NewConnectivityTestResult() *ConnectivityTestResult`

NewConnectivityTestResult instantiates a new ConnectivityTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectivityTestResultWithDefaults

`func NewConnectivityTestResultWithDefaults() *ConnectivityTestResult`

NewConnectivityTestResultWithDefaults instantiates a new ConnectivityTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ConnectivityTestResult) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ConnectivityTestResult) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ConnectivityTestResult) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ConnectivityTestResult) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSuccessful

`func (o *ConnectivityTestResult) GetSuccessful() bool`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *ConnectivityTestResult) GetSuccessfulOk() (*bool, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *ConnectivityTestResult) SetSuccessful(v bool)`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *ConnectivityTestResult) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### GetWorkflowId

`func (o *ConnectivityTestResult) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ConnectivityTestResult) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ConnectivityTestResult) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *ConnectivityTestResult) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


