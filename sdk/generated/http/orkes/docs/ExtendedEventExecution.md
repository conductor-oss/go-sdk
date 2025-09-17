# ExtendedEventExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **int64** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**EventHandler** | Pointer to [**EventHandler**](EventHandler.md) |  | [optional] 
**FullMessagePayload** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Output** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Payload** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewExtendedEventExecution

`func NewExtendedEventExecution() *ExtendedEventExecution`

NewExtendedEventExecution instantiates a new ExtendedEventExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedEventExecutionWithDefaults

`func NewExtendedEventExecutionWithDefaults() *ExtendedEventExecution`

NewExtendedEventExecutionWithDefaults instantiates a new ExtendedEventExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ExtendedEventExecution) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ExtendedEventExecution) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ExtendedEventExecution) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ExtendedEventExecution) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreated

`func (o *ExtendedEventExecution) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ExtendedEventExecution) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ExtendedEventExecution) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ExtendedEventExecution) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEvent

`func (o *ExtendedEventExecution) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ExtendedEventExecution) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ExtendedEventExecution) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *ExtendedEventExecution) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetEventHandler

`func (o *ExtendedEventExecution) GetEventHandler() EventHandler`

GetEventHandler returns the EventHandler field if non-nil, zero value otherwise.

### GetEventHandlerOk

`func (o *ExtendedEventExecution) GetEventHandlerOk() (*EventHandler, bool)`

GetEventHandlerOk returns a tuple with the EventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandler

`func (o *ExtendedEventExecution) SetEventHandler(v EventHandler)`

SetEventHandler sets EventHandler field to given value.

### HasEventHandler

`func (o *ExtendedEventExecution) HasEventHandler() bool`

HasEventHandler returns a boolean if a field has been set.

### GetFullMessagePayload

`func (o *ExtendedEventExecution) GetFullMessagePayload() map[string]map[string]interface{}`

GetFullMessagePayload returns the FullMessagePayload field if non-nil, zero value otherwise.

### GetFullMessagePayloadOk

`func (o *ExtendedEventExecution) GetFullMessagePayloadOk() (*map[string]map[string]interface{}, bool)`

GetFullMessagePayloadOk returns a tuple with the FullMessagePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullMessagePayload

`func (o *ExtendedEventExecution) SetFullMessagePayload(v map[string]map[string]interface{})`

SetFullMessagePayload sets FullMessagePayload field to given value.

### HasFullMessagePayload

`func (o *ExtendedEventExecution) HasFullMessagePayload() bool`

HasFullMessagePayload returns a boolean if a field has been set.

### GetId

`func (o *ExtendedEventExecution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtendedEventExecution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtendedEventExecution) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExtendedEventExecution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessageId

`func (o *ExtendedEventExecution) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ExtendedEventExecution) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ExtendedEventExecution) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ExtendedEventExecution) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetName

`func (o *ExtendedEventExecution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtendedEventExecution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtendedEventExecution) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExtendedEventExecution) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *ExtendedEventExecution) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ExtendedEventExecution) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ExtendedEventExecution) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ExtendedEventExecution) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOutput

`func (o *ExtendedEventExecution) GetOutput() map[string]map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ExtendedEventExecution) GetOutputOk() (*map[string]map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ExtendedEventExecution) SetOutput(v map[string]map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ExtendedEventExecution) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetPayload

`func (o *ExtendedEventExecution) GetPayload() map[string]map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ExtendedEventExecution) GetPayloadOk() (*map[string]map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ExtendedEventExecution) SetPayload(v map[string]map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ExtendedEventExecution) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetStatus

`func (o *ExtendedEventExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExtendedEventExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExtendedEventExecution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExtendedEventExecution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *ExtendedEventExecution) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *ExtendedEventExecution) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *ExtendedEventExecution) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *ExtendedEventExecution) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


