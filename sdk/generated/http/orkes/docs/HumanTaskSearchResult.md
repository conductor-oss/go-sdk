# HumanTaskSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to **int32** |  | [optional] 
**PageSizeLimit** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**[]HumanTaskEntry**](HumanTaskEntry.md) |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**TotalHits** | Pointer to **int64** |  | [optional] 

## Methods

### NewHumanTaskSearchResult

`func NewHumanTaskSearchResult() *HumanTaskSearchResult`

NewHumanTaskSearchResult instantiates a new HumanTaskSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHumanTaskSearchResultWithDefaults

`func NewHumanTaskSearchResultWithDefaults() *HumanTaskSearchResult`

NewHumanTaskSearchResultWithDefaults instantiates a new HumanTaskSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *HumanTaskSearchResult) GetHits() int32`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *HumanTaskSearchResult) GetHitsOk() (*int32, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *HumanTaskSearchResult) SetHits(v int32)`

SetHits sets Hits field to given value.

### HasHits

`func (o *HumanTaskSearchResult) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetPageSizeLimit

`func (o *HumanTaskSearchResult) GetPageSizeLimit() int32`

GetPageSizeLimit returns the PageSizeLimit field if non-nil, zero value otherwise.

### GetPageSizeLimitOk

`func (o *HumanTaskSearchResult) GetPageSizeLimitOk() (*int32, bool)`

GetPageSizeLimitOk returns a tuple with the PageSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSizeLimit

`func (o *HumanTaskSearchResult) SetPageSizeLimit(v int32)`

SetPageSizeLimit sets PageSizeLimit field to given value.

### HasPageSizeLimit

`func (o *HumanTaskSearchResult) HasPageSizeLimit() bool`

HasPageSizeLimit returns a boolean if a field has been set.

### GetResults

`func (o *HumanTaskSearchResult) GetResults() []HumanTaskEntry`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *HumanTaskSearchResult) GetResultsOk() (*[]HumanTaskEntry, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *HumanTaskSearchResult) SetResults(v []HumanTaskEntry)`

SetResults sets Results field to given value.

### HasResults

`func (o *HumanTaskSearchResult) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStart

`func (o *HumanTaskSearchResult) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *HumanTaskSearchResult) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *HumanTaskSearchResult) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *HumanTaskSearchResult) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTotalHits

`func (o *HumanTaskSearchResult) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *HumanTaskSearchResult) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *HumanTaskSearchResult) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *HumanTaskSearchResult) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


