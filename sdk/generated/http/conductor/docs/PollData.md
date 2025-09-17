# PollData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueName** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**WorkerId** | Pointer to **string** |  | [optional] 
**LastPollTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewPollData

`func NewPollData() *PollData`

NewPollData instantiates a new PollData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollDataWithDefaults

`func NewPollDataWithDefaults() *PollData`

NewPollDataWithDefaults instantiates a new PollData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueName

`func (o *PollData) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *PollData) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *PollData) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *PollData) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### GetDomain

`func (o *PollData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PollData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PollData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PollData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetWorkerId

`func (o *PollData) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *PollData) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *PollData) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *PollData) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetLastPollTime

`func (o *PollData) GetLastPollTime() int64`

GetLastPollTime returns the LastPollTime field if non-nil, zero value otherwise.

### GetLastPollTimeOk

`func (o *PollData) GetLastPollTimeOk() (*int64, bool)`

GetLastPollTimeOk returns a tuple with the LastPollTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPollTime

`func (o *PollData) SetLastPollTime(v int64)`

SetLastPollTime sets LastPollTime field to given value.

### HasLastPollTime

`func (o *PollData) HasLastPollTime() bool`

HasLastPollTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


