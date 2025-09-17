# PromptTemplateTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LlmProvider** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**PromptVariables** | Pointer to  |  | [optional] 
**StopWords** | Pointer to **[]string** |  | [optional] 
**Temperature** | Pointer to **float64** |  | [optional] 
**TopP** | Pointer to **float64** |  | [optional] 

## Methods

### NewPromptTemplateTestRequest

`func NewPromptTemplateTestRequest() *PromptTemplateTestRequest`

NewPromptTemplateTestRequest instantiates a new PromptTemplateTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptTemplateTestRequestWithDefaults

`func NewPromptTemplateTestRequestWithDefaults() *PromptTemplateTestRequest`

NewPromptTemplateTestRequestWithDefaults instantiates a new PromptTemplateTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLlmProvider

`func (o *PromptTemplateTestRequest) GetLlmProvider() string`

GetLlmProvider returns the LlmProvider field if non-nil, zero value otherwise.

### GetLlmProviderOk

`func (o *PromptTemplateTestRequest) GetLlmProviderOk() (*string, bool)`

GetLlmProviderOk returns a tuple with the LlmProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmProvider

`func (o *PromptTemplateTestRequest) SetLlmProvider(v string)`

SetLlmProvider sets LlmProvider field to given value.

### HasLlmProvider

`func (o *PromptTemplateTestRequest) HasLlmProvider() bool`

HasLlmProvider returns a boolean if a field has been set.

### GetModel

`func (o *PromptTemplateTestRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PromptTemplateTestRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PromptTemplateTestRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PromptTemplateTestRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPrompt

`func (o *PromptTemplateTestRequest) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *PromptTemplateTestRequest) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *PromptTemplateTestRequest) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *PromptTemplateTestRequest) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetPromptVariables

`func (o *PromptTemplateTestRequest) GetPromptVariables() map[string]map[string]interface{}`

GetPromptVariables returns the PromptVariables field if non-nil, zero value otherwise.

### GetPromptVariablesOk

`func (o *PromptTemplateTestRequest) GetPromptVariablesOk() (*map[string]map[string]interface{}, bool)`

GetPromptVariablesOk returns a tuple with the PromptVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptVariables

`func (o *PromptTemplateTestRequest) SetPromptVariables(v map[string]map[string]interface{})`

SetPromptVariables sets PromptVariables field to given value.

### HasPromptVariables

`func (o *PromptTemplateTestRequest) HasPromptVariables() bool`

HasPromptVariables returns a boolean if a field has been set.

### SetPromptVariablesNil

`func (o *PromptTemplateTestRequest) SetPromptVariablesNil(b bool)`

 SetPromptVariablesNil sets the value for PromptVariables to be an explicit nil

### UnsetPromptVariables
`func (o *PromptTemplateTestRequest) UnsetPromptVariables()`

UnsetPromptVariables ensures that no value is present for PromptVariables, not even an explicit nil
### GetStopWords

`func (o *PromptTemplateTestRequest) GetStopWords() []string`

GetStopWords returns the StopWords field if non-nil, zero value otherwise.

### GetStopWordsOk

`func (o *PromptTemplateTestRequest) GetStopWordsOk() (*[]string, bool)`

GetStopWordsOk returns a tuple with the StopWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopWords

`func (o *PromptTemplateTestRequest) SetStopWords(v []string)`

SetStopWords sets StopWords field to given value.

### HasStopWords

`func (o *PromptTemplateTestRequest) HasStopWords() bool`

HasStopWords returns a boolean if a field has been set.

### GetTemperature

`func (o *PromptTemplateTestRequest) GetTemperature() float64`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *PromptTemplateTestRequest) GetTemperatureOk() (*float64, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *PromptTemplateTestRequest) SetTemperature(v float64)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *PromptTemplateTestRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTopP

`func (o *PromptTemplateTestRequest) GetTopP() float64`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *PromptTemplateTestRequest) GetTopPOk() (*float64, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *PromptTemplateTestRequest) SetTopP(v float64)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *PromptTemplateTestRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


