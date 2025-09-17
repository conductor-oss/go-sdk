# UpgradeWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**TaskOutput** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**WorkflowInput** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewUpgradeWorkflowRequest

`func NewUpgradeWorkflowRequest(name string, ) *UpgradeWorkflowRequest`

NewUpgradeWorkflowRequest instantiates a new UpgradeWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeWorkflowRequestWithDefaults

`func NewUpgradeWorkflowRequestWithDefaults() *UpgradeWorkflowRequest`

NewUpgradeWorkflowRequestWithDefaults instantiates a new UpgradeWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpgradeWorkflowRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpgradeWorkflowRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpgradeWorkflowRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTaskOutput

`func (o *UpgradeWorkflowRequest) GetTaskOutput() map[string]map[string]interface{}`

GetTaskOutput returns the TaskOutput field if non-nil, zero value otherwise.

### GetTaskOutputOk

`func (o *UpgradeWorkflowRequest) GetTaskOutputOk() (*map[string]map[string]interface{}, bool)`

GetTaskOutputOk returns a tuple with the TaskOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOutput

`func (o *UpgradeWorkflowRequest) SetTaskOutput(v map[string]map[string]interface{})`

SetTaskOutput sets TaskOutput field to given value.

### HasTaskOutput

`func (o *UpgradeWorkflowRequest) HasTaskOutput() bool`

HasTaskOutput returns a boolean if a field has been set.

### GetVersion

`func (o *UpgradeWorkflowRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpgradeWorkflowRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpgradeWorkflowRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpgradeWorkflowRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkflowInput

`func (o *UpgradeWorkflowRequest) GetWorkflowInput() map[string]map[string]interface{}`

GetWorkflowInput returns the WorkflowInput field if non-nil, zero value otherwise.

### GetWorkflowInputOk

`func (o *UpgradeWorkflowRequest) GetWorkflowInputOk() (*map[string]map[string]interface{}, bool)`

GetWorkflowInputOk returns a tuple with the WorkflowInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInput

`func (o *UpgradeWorkflowRequest) SetWorkflowInput(v map[string]map[string]interface{})`

SetWorkflowInput sets WorkflowInput field to given value.

### HasWorkflowInput

`func (o *UpgradeWorkflowRequest) HasWorkflowInput() bool`

HasWorkflowInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


