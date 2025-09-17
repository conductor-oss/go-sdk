# WebhookExecutionHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** |  | [optional] 
**Matched** | Pointer to **bool** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to **int64** |  | [optional] 
**WorkflowIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWebhookExecutionHistory

`func NewWebhookExecutionHistory() *WebhookExecutionHistory`

NewWebhookExecutionHistory instantiates a new WebhookExecutionHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookExecutionHistoryWithDefaults

`func NewWebhookExecutionHistoryWithDefaults() *WebhookExecutionHistory`

NewWebhookExecutionHistoryWithDefaults instantiates a new WebhookExecutionHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *WebhookExecutionHistory) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WebhookExecutionHistory) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WebhookExecutionHistory) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *WebhookExecutionHistory) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetMatched

`func (o *WebhookExecutionHistory) GetMatched() bool`

GetMatched returns the Matched field if non-nil, zero value otherwise.

### GetMatchedOk

`func (o *WebhookExecutionHistory) GetMatchedOk() (*bool, bool)`

GetMatchedOk returns a tuple with the Matched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatched

`func (o *WebhookExecutionHistory) SetMatched(v bool)`

SetMatched sets Matched field to given value.

### HasMatched

`func (o *WebhookExecutionHistory) HasMatched() bool`

HasMatched returns a boolean if a field has been set.

### GetPayload

`func (o *WebhookExecutionHistory) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebhookExecutionHistory) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebhookExecutionHistory) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *WebhookExecutionHistory) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetTimeStamp

`func (o *WebhookExecutionHistory) GetTimeStamp() int64`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *WebhookExecutionHistory) GetTimeStampOk() (*int64, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *WebhookExecutionHistory) SetTimeStamp(v int64)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *WebhookExecutionHistory) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetWorkflowIds

`func (o *WebhookExecutionHistory) GetWorkflowIds() []string`

GetWorkflowIds returns the WorkflowIds field if non-nil, zero value otherwise.

### GetWorkflowIdsOk

`func (o *WebhookExecutionHistory) GetWorkflowIdsOk() (*[]string, bool)`

GetWorkflowIdsOk returns a tuple with the WorkflowIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowIds

`func (o *WebhookExecutionHistory) SetWorkflowIds(v []string)`

SetWorkflowIds sets WorkflowIds field to given value.

### HasWorkflowIds

`func (o *WebhookExecutionHistory) HasWorkflowIds() bool`

HasWorkflowIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


