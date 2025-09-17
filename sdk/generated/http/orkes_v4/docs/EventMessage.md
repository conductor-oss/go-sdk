# EventMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** |  | [optional] 
**EventExecutions** | Pointer to [**[]ExtendedEventExecution**](ExtendedEventExecution.md) |  | [optional] 
**EventTarget** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**FullPayload** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewEventMessage

`func NewEventMessage() *EventMessage`

NewEventMessage instantiates a new EventMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventMessageWithDefaults

`func NewEventMessageWithDefaults() *EventMessage`

NewEventMessageWithDefaults instantiates a new EventMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EventMessage) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventMessage) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventMessage) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventMessage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEventExecutions

`func (o *EventMessage) GetEventExecutions() []ExtendedEventExecution`

GetEventExecutions returns the EventExecutions field if non-nil, zero value otherwise.

### GetEventExecutionsOk

`func (o *EventMessage) GetEventExecutionsOk() (*[]ExtendedEventExecution, bool)`

GetEventExecutionsOk returns a tuple with the EventExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventExecutions

`func (o *EventMessage) SetEventExecutions(v []ExtendedEventExecution)`

SetEventExecutions sets EventExecutions field to given value.

### HasEventExecutions

`func (o *EventMessage) HasEventExecutions() bool`

HasEventExecutions returns a boolean if a field has been set.

### GetEventTarget

`func (o *EventMessage) GetEventTarget() string`

GetEventTarget returns the EventTarget field if non-nil, zero value otherwise.

### GetEventTargetOk

`func (o *EventMessage) GetEventTargetOk() (*string, bool)`

GetEventTargetOk returns a tuple with the EventTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTarget

`func (o *EventMessage) SetEventTarget(v string)`

SetEventTarget sets EventTarget field to given value.

### HasEventTarget

`func (o *EventMessage) HasEventTarget() bool`

HasEventTarget returns a boolean if a field has been set.

### GetEventType

`func (o *EventMessage) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventMessage) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventMessage) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventMessage) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetFullPayload

`func (o *EventMessage) GetFullPayload() map[string]interface{}`

GetFullPayload returns the FullPayload field if non-nil, zero value otherwise.

### GetFullPayloadOk

`func (o *EventMessage) GetFullPayloadOk() (*map[string]interface{}, bool)`

GetFullPayloadOk returns a tuple with the FullPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPayload

`func (o *EventMessage) SetFullPayload(v map[string]interface{})`

SetFullPayload sets FullPayload field to given value.

### HasFullPayload

`func (o *EventMessage) HasFullPayload() bool`

HasFullPayload returns a boolean if a field has been set.

### GetId

`func (o *EventMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *EventMessage) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *EventMessage) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *EventMessage) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *EventMessage) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPayload

`func (o *EventMessage) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *EventMessage) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *EventMessage) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *EventMessage) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetStatus

`func (o *EventMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventMessage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *EventMessage) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *EventMessage) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *EventMessage) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *EventMessage) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


