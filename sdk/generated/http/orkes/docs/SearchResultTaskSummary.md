# SearchResultTaskSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]TaskSummary**](TaskSummary.md) |  | [optional] 
**TotalHits** | Pointer to **int64** |  | [optional] 

## Methods

### NewSearchResultTaskSummary

`func NewSearchResultTaskSummary() *SearchResultTaskSummary`

NewSearchResultTaskSummary instantiates a new SearchResultTaskSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultTaskSummaryWithDefaults

`func NewSearchResultTaskSummaryWithDefaults() *SearchResultTaskSummary`

NewSearchResultTaskSummaryWithDefaults instantiates a new SearchResultTaskSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SearchResultTaskSummary) GetResults() []TaskSummary`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchResultTaskSummary) GetResultsOk() (*[]TaskSummary, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchResultTaskSummary) SetResults(v []TaskSummary)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchResultTaskSummary) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalHits

`func (o *SearchResultTaskSummary) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *SearchResultTaskSummary) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *SearchResultTaskSummary) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *SearchResultTaskSummary) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


