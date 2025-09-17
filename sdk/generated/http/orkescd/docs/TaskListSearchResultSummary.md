# TaskListSearchResultSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**Summary** | Pointer to **map[string]int64** |  | [optional] 
**TotalHits** | Pointer to **int64** |  | [optional] 

## Methods

### NewTaskListSearchResultSummary

`func NewTaskListSearchResultSummary() *TaskListSearchResultSummary`

NewTaskListSearchResultSummary instantiates a new TaskListSearchResultSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskListSearchResultSummaryWithDefaults

`func NewTaskListSearchResultSummaryWithDefaults() *TaskListSearchResultSummary`

NewTaskListSearchResultSummaryWithDefaults instantiates a new TaskListSearchResultSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *TaskListSearchResultSummary) GetResults() []Task`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TaskListSearchResultSummary) GetResultsOk() (*[]Task, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TaskListSearchResultSummary) SetResults(v []Task)`

SetResults sets Results field to given value.

### HasResults

`func (o *TaskListSearchResultSummary) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSummary

`func (o *TaskListSearchResultSummary) GetSummary() map[string]int64`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TaskListSearchResultSummary) GetSummaryOk() (*map[string]int64, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TaskListSearchResultSummary) SetSummary(v map[string]int64)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *TaskListSearchResultSummary) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTotalHits

`func (o *TaskListSearchResultSummary) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *TaskListSearchResultSummary) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *TaskListSearchResultSummary) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *TaskListSearchResultSummary) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


