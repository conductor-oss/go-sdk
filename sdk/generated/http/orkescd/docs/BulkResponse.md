# BulkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkErrorResults** | Pointer to **map[string]string** |  | [optional] 
**BulkSuccessfulResults** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewBulkResponse

`func NewBulkResponse() *BulkResponse`

NewBulkResponse instantiates a new BulkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResponseWithDefaults

`func NewBulkResponseWithDefaults() *BulkResponse`

NewBulkResponseWithDefaults instantiates a new BulkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkErrorResults

`func (o *BulkResponse) GetBulkErrorResults() map[string]string`

GetBulkErrorResults returns the BulkErrorResults field if non-nil, zero value otherwise.

### GetBulkErrorResultsOk

`func (o *BulkResponse) GetBulkErrorResultsOk() (*map[string]string, bool)`

GetBulkErrorResultsOk returns a tuple with the BulkErrorResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkErrorResults

`func (o *BulkResponse) SetBulkErrorResults(v map[string]string)`

SetBulkErrorResults sets BulkErrorResults field to given value.

### HasBulkErrorResults

`func (o *BulkResponse) HasBulkErrorResults() bool`

HasBulkErrorResults returns a boolean if a field has been set.

### GetBulkSuccessfulResults

`func (o *BulkResponse) GetBulkSuccessfulResults() []map[string]interface{}`

GetBulkSuccessfulResults returns the BulkSuccessfulResults field if non-nil, zero value otherwise.

### GetBulkSuccessfulResultsOk

`func (o *BulkResponse) GetBulkSuccessfulResultsOk() (*[]map[string]interface{}, bool)`

GetBulkSuccessfulResultsOk returns a tuple with the BulkSuccessfulResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSuccessfulResults

`func (o *BulkResponse) SetBulkSuccessfulResults(v []map[string]interface{})`

SetBulkSuccessfulResults sets BulkSuccessfulResults field to given value.

### HasBulkSuccessfulResults

`func (o *BulkResponse) HasBulkSuccessfulResults() bool`

HasBulkSuccessfulResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


