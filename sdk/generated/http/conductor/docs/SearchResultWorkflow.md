# SearchResultWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalHits** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]Workflow**](Workflow.md) |  | [optional] 

## Methods

### NewSearchResultWorkflow

`func NewSearchResultWorkflow() *SearchResultWorkflow`

NewSearchResultWorkflow instantiates a new SearchResultWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultWorkflowWithDefaults

`func NewSearchResultWorkflowWithDefaults() *SearchResultWorkflow`

NewSearchResultWorkflowWithDefaults instantiates a new SearchResultWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalHits

`func (o *SearchResultWorkflow) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *SearchResultWorkflow) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *SearchResultWorkflow) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *SearchResultWorkflow) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetResults

`func (o *SearchResultWorkflow) GetResults() []Workflow`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchResultWorkflow) GetResultsOk() (*[]Workflow, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchResultWorkflow) SetResults(v []Workflow)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchResultWorkflow) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


