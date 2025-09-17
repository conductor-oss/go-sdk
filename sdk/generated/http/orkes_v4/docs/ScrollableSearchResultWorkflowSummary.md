# ScrollableSearchResultWorkflowSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]WorkflowSummary**](WorkflowSummary.md) |  | [optional] 
**TotalHits** | Pointer to **int64** |  | [optional] 

## Methods

### NewScrollableSearchResultWorkflowSummary

`func NewScrollableSearchResultWorkflowSummary() *ScrollableSearchResultWorkflowSummary`

NewScrollableSearchResultWorkflowSummary instantiates a new ScrollableSearchResultWorkflowSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScrollableSearchResultWorkflowSummaryWithDefaults

`func NewScrollableSearchResultWorkflowSummaryWithDefaults() *ScrollableSearchResultWorkflowSummary`

NewScrollableSearchResultWorkflowSummaryWithDefaults instantiates a new ScrollableSearchResultWorkflowSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *ScrollableSearchResultWorkflowSummary) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *ScrollableSearchResultWorkflowSummary) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *ScrollableSearchResultWorkflowSummary) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.

### HasQueryId

`func (o *ScrollableSearchResultWorkflowSummary) HasQueryId() bool`

HasQueryId returns a boolean if a field has been set.

### GetResults

`func (o *ScrollableSearchResultWorkflowSummary) GetResults() []WorkflowSummary`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ScrollableSearchResultWorkflowSummary) GetResultsOk() (*[]WorkflowSummary, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ScrollableSearchResultWorkflowSummary) SetResults(v []WorkflowSummary)`

SetResults sets Results field to given value.

### HasResults

`func (o *ScrollableSearchResultWorkflowSummary) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalHits

`func (o *ScrollableSearchResultWorkflowSummary) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *ScrollableSearchResultWorkflowSummary) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *ScrollableSearchResultWorkflowSummary) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *ScrollableSearchResultWorkflowSummary) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


