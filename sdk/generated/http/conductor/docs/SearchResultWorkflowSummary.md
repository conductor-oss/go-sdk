# SearchResultWorkflowSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalHits** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]WorkflowSummary**](WorkflowSummary.md) |  | [optional] 

## Methods

### NewSearchResultWorkflowSummary

`func NewSearchResultWorkflowSummary() *SearchResultWorkflowSummary`

NewSearchResultWorkflowSummary instantiates a new SearchResultWorkflowSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultWorkflowSummaryWithDefaults

`func NewSearchResultWorkflowSummaryWithDefaults() *SearchResultWorkflowSummary`

NewSearchResultWorkflowSummaryWithDefaults instantiates a new SearchResultWorkflowSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalHits

`func (o *SearchResultWorkflowSummary) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *SearchResultWorkflowSummary) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *SearchResultWorkflowSummary) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *SearchResultWorkflowSummary) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetResults

`func (o *SearchResultWorkflowSummary) GetResults() []WorkflowSummary`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchResultWorkflowSummary) GetResultsOk() (*[]WorkflowSummary, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchResultWorkflowSummary) SetResults(v []WorkflowSummary)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchResultWorkflowSummary) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


