# SearchResultHandledEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]HandledEventResponse**](HandledEventResponse.md) |  | [optional] 
**TotalHits** | Pointer to **int64** |  | [optional] 

## Methods

### NewSearchResultHandledEventResponse

`func NewSearchResultHandledEventResponse() *SearchResultHandledEventResponse`

NewSearchResultHandledEventResponse instantiates a new SearchResultHandledEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultHandledEventResponseWithDefaults

`func NewSearchResultHandledEventResponseWithDefaults() *SearchResultHandledEventResponse`

NewSearchResultHandledEventResponseWithDefaults instantiates a new SearchResultHandledEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SearchResultHandledEventResponse) GetResults() []HandledEventResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchResultHandledEventResponse) GetResultsOk() (*[]HandledEventResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchResultHandledEventResponse) SetResults(v []HandledEventResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchResultHandledEventResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalHits

`func (o *SearchResultHandledEventResponse) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *SearchResultHandledEventResponse) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *SearchResultHandledEventResponse) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *SearchResultHandledEventResponse) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


