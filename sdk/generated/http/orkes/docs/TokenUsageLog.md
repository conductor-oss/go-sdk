# TokenUsageLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Api** | Pointer to **string** |  | [optional] 
**CompletionTokens** | Pointer to **int32** |  | [optional] 
**IntegrationName** | Pointer to **string** |  | [optional] 
**PeriodStart** | Pointer to **int64** |  | [optional] 
**PromptTokens** | Pointer to **int32** |  | [optional] 
**TotalTokens** | Pointer to **int32** |  | [optional] 

## Methods

### NewTokenUsageLog

`func NewTokenUsageLog() *TokenUsageLog`

NewTokenUsageLog instantiates a new TokenUsageLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenUsageLogWithDefaults

`func NewTokenUsageLogWithDefaults() *TokenUsageLog`

NewTokenUsageLogWithDefaults instantiates a new TokenUsageLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApi

`func (o *TokenUsageLog) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *TokenUsageLog) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *TokenUsageLog) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *TokenUsageLog) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *TokenUsageLog) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *TokenUsageLog) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *TokenUsageLog) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *TokenUsageLog) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetIntegrationName

`func (o *TokenUsageLog) GetIntegrationName() string`

GetIntegrationName returns the IntegrationName field if non-nil, zero value otherwise.

### GetIntegrationNameOk

`func (o *TokenUsageLog) GetIntegrationNameOk() (*string, bool)`

GetIntegrationNameOk returns a tuple with the IntegrationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationName

`func (o *TokenUsageLog) SetIntegrationName(v string)`

SetIntegrationName sets IntegrationName field to given value.

### HasIntegrationName

`func (o *TokenUsageLog) HasIntegrationName() bool`

HasIntegrationName returns a boolean if a field has been set.

### GetPeriodStart

`func (o *TokenUsageLog) GetPeriodStart() int64`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *TokenUsageLog) GetPeriodStartOk() (*int64, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *TokenUsageLog) SetPeriodStart(v int64)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *TokenUsageLog) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPromptTokens

`func (o *TokenUsageLog) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *TokenUsageLog) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *TokenUsageLog) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *TokenUsageLog) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetTotalTokens

`func (o *TokenUsageLog) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *TokenUsageLog) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *TokenUsageLog) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *TokenUsageLog) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


