# BulkResponseWorkflowModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkErrorResults** | Pointer to **map[string]string** |  | [optional] 
**BulkSuccessfulResults** | Pointer to [**[]WorkflowModel**](WorkflowModel.md) |  | [optional] 

## Methods

### NewBulkResponseWorkflowModel

`func NewBulkResponseWorkflowModel() *BulkResponseWorkflowModel`

NewBulkResponseWorkflowModel instantiates a new BulkResponseWorkflowModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResponseWorkflowModelWithDefaults

`func NewBulkResponseWorkflowModelWithDefaults() *BulkResponseWorkflowModel`

NewBulkResponseWorkflowModelWithDefaults instantiates a new BulkResponseWorkflowModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkErrorResults

`func (o *BulkResponseWorkflowModel) GetBulkErrorResults() map[string]string`

GetBulkErrorResults returns the BulkErrorResults field if non-nil, zero value otherwise.

### GetBulkErrorResultsOk

`func (o *BulkResponseWorkflowModel) GetBulkErrorResultsOk() (*map[string]string, bool)`

GetBulkErrorResultsOk returns a tuple with the BulkErrorResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkErrorResults

`func (o *BulkResponseWorkflowModel) SetBulkErrorResults(v map[string]string)`

SetBulkErrorResults sets BulkErrorResults field to given value.

### HasBulkErrorResults

`func (o *BulkResponseWorkflowModel) HasBulkErrorResults() bool`

HasBulkErrorResults returns a boolean if a field has been set.

### GetBulkSuccessfulResults

`func (o *BulkResponseWorkflowModel) GetBulkSuccessfulResults() []WorkflowModel`

GetBulkSuccessfulResults returns the BulkSuccessfulResults field if non-nil, zero value otherwise.

### GetBulkSuccessfulResultsOk

`func (o *BulkResponseWorkflowModel) GetBulkSuccessfulResultsOk() (*[]WorkflowModel, bool)`

GetBulkSuccessfulResultsOk returns a tuple with the BulkSuccessfulResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSuccessfulResults

`func (o *BulkResponseWorkflowModel) SetBulkSuccessfulResults(v []WorkflowModel)`

SetBulkSuccessfulResults sets BulkSuccessfulResults field to given value.

### HasBulkSuccessfulResults

`func (o *BulkResponseWorkflowModel) HasBulkSuccessfulResults() bool`

HasBulkSuccessfulResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


