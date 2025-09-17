# WebhookConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**EvaluatorType** | Pointer to **string** |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 
**HeaderKey** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReceiverWorkflowNamesToVersions** | Pointer to **map[string]int32** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**SecretValue** | Pointer to **string** |  | [optional] 
**SourcePlatform** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**UrlVerified** | Pointer to **bool** |  | [optional] 
**Verifier** | Pointer to **string** |  | [optional] 
**WebhookExecutionHistory** | Pointer to [**[]WebhookExecutionHistory**](WebhookExecutionHistory.md) |  | [optional] 
**WorkflowsToStart** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewWebhookConfig

`func NewWebhookConfig() *WebhookConfig`

NewWebhookConfig instantiates a new WebhookConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookConfigWithDefaults

`func NewWebhookConfigWithDefaults() *WebhookConfig`

NewWebhookConfigWithDefaults instantiates a new WebhookConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *WebhookConfig) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WebhookConfig) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WebhookConfig) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WebhookConfig) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEvaluatorType

`func (o *WebhookConfig) GetEvaluatorType() string`

GetEvaluatorType returns the EvaluatorType field if non-nil, zero value otherwise.

### GetEvaluatorTypeOk

`func (o *WebhookConfig) GetEvaluatorTypeOk() (*string, bool)`

GetEvaluatorTypeOk returns a tuple with the EvaluatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorType

`func (o *WebhookConfig) SetEvaluatorType(v string)`

SetEvaluatorType sets EvaluatorType field to given value.

### HasEvaluatorType

`func (o *WebhookConfig) HasEvaluatorType() bool`

HasEvaluatorType returns a boolean if a field has been set.

### GetExpression

`func (o *WebhookConfig) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *WebhookConfig) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *WebhookConfig) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *WebhookConfig) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetHeaderKey

`func (o *WebhookConfig) GetHeaderKey() string`

GetHeaderKey returns the HeaderKey field if non-nil, zero value otherwise.

### GetHeaderKeyOk

`func (o *WebhookConfig) GetHeaderKeyOk() (*string, bool)`

GetHeaderKeyOk returns a tuple with the HeaderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderKey

`func (o *WebhookConfig) SetHeaderKey(v string)`

SetHeaderKey sets HeaderKey field to given value.

### HasHeaderKey

`func (o *WebhookConfig) HasHeaderKey() bool`

HasHeaderKey returns a boolean if a field has been set.

### GetHeaders

`func (o *WebhookConfig) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookConfig) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookConfig) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WebhookConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetId

`func (o *WebhookConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WebhookConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReceiverWorkflowNamesToVersions

`func (o *WebhookConfig) GetReceiverWorkflowNamesToVersions() map[string]int32`

GetReceiverWorkflowNamesToVersions returns the ReceiverWorkflowNamesToVersions field if non-nil, zero value otherwise.

### GetReceiverWorkflowNamesToVersionsOk

`func (o *WebhookConfig) GetReceiverWorkflowNamesToVersionsOk() (*map[string]int32, bool)`

GetReceiverWorkflowNamesToVersionsOk returns a tuple with the ReceiverWorkflowNamesToVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverWorkflowNamesToVersions

`func (o *WebhookConfig) SetReceiverWorkflowNamesToVersions(v map[string]int32)`

SetReceiverWorkflowNamesToVersions sets ReceiverWorkflowNamesToVersions field to given value.

### HasReceiverWorkflowNamesToVersions

`func (o *WebhookConfig) HasReceiverWorkflowNamesToVersions() bool`

HasReceiverWorkflowNamesToVersions returns a boolean if a field has been set.

### GetSecretKey

`func (o *WebhookConfig) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *WebhookConfig) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *WebhookConfig) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *WebhookConfig) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSecretValue

`func (o *WebhookConfig) GetSecretValue() string`

GetSecretValue returns the SecretValue field if non-nil, zero value otherwise.

### GetSecretValueOk

`func (o *WebhookConfig) GetSecretValueOk() (*string, bool)`

GetSecretValueOk returns a tuple with the SecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretValue

`func (o *WebhookConfig) SetSecretValue(v string)`

SetSecretValue sets SecretValue field to given value.

### HasSecretValue

`func (o *WebhookConfig) HasSecretValue() bool`

HasSecretValue returns a boolean if a field has been set.

### GetSourcePlatform

`func (o *WebhookConfig) GetSourcePlatform() string`

GetSourcePlatform returns the SourcePlatform field if non-nil, zero value otherwise.

### GetSourcePlatformOk

`func (o *WebhookConfig) GetSourcePlatformOk() (*string, bool)`

GetSourcePlatformOk returns a tuple with the SourcePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePlatform

`func (o *WebhookConfig) SetSourcePlatform(v string)`

SetSourcePlatform sets SourcePlatform field to given value.

### HasSourcePlatform

`func (o *WebhookConfig) HasSourcePlatform() bool`

HasSourcePlatform returns a boolean if a field has been set.

### GetTags

`func (o *WebhookConfig) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WebhookConfig) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WebhookConfig) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WebhookConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUrlVerified

`func (o *WebhookConfig) GetUrlVerified() bool`

GetUrlVerified returns the UrlVerified field if non-nil, zero value otherwise.

### GetUrlVerifiedOk

`func (o *WebhookConfig) GetUrlVerifiedOk() (*bool, bool)`

GetUrlVerifiedOk returns a tuple with the UrlVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlVerified

`func (o *WebhookConfig) SetUrlVerified(v bool)`

SetUrlVerified sets UrlVerified field to given value.

### HasUrlVerified

`func (o *WebhookConfig) HasUrlVerified() bool`

HasUrlVerified returns a boolean if a field has been set.

### GetVerifier

`func (o *WebhookConfig) GetVerifier() string`

GetVerifier returns the Verifier field if non-nil, zero value otherwise.

### GetVerifierOk

`func (o *WebhookConfig) GetVerifierOk() (*string, bool)`

GetVerifierOk returns a tuple with the Verifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifier

`func (o *WebhookConfig) SetVerifier(v string)`

SetVerifier sets Verifier field to given value.

### HasVerifier

`func (o *WebhookConfig) HasVerifier() bool`

HasVerifier returns a boolean if a field has been set.

### GetWebhookExecutionHistory

`func (o *WebhookConfig) GetWebhookExecutionHistory() []WebhookExecutionHistory`

GetWebhookExecutionHistory returns the WebhookExecutionHistory field if non-nil, zero value otherwise.

### GetWebhookExecutionHistoryOk

`func (o *WebhookConfig) GetWebhookExecutionHistoryOk() (*[]WebhookExecutionHistory, bool)`

GetWebhookExecutionHistoryOk returns a tuple with the WebhookExecutionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookExecutionHistory

`func (o *WebhookConfig) SetWebhookExecutionHistory(v []WebhookExecutionHistory)`

SetWebhookExecutionHistory sets WebhookExecutionHistory field to given value.

### HasWebhookExecutionHistory

`func (o *WebhookConfig) HasWebhookExecutionHistory() bool`

HasWebhookExecutionHistory returns a boolean if a field has been set.

### GetWorkflowsToStart

`func (o *WebhookConfig) GetWorkflowsToStart() map[string]map[string]interface{}`

GetWorkflowsToStart returns the WorkflowsToStart field if non-nil, zero value otherwise.

### GetWorkflowsToStartOk

`func (o *WebhookConfig) GetWorkflowsToStartOk() (*map[string]map[string]interface{}, bool)`

GetWorkflowsToStartOk returns a tuple with the WorkflowsToStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowsToStart

`func (o *WebhookConfig) SetWorkflowsToStart(v map[string]map[string]interface{})`

SetWorkflowsToStart sets WorkflowsToStart field to given value.

### HasWorkflowsToStart

`func (o *WebhookConfig) HasWorkflowsToStart() bool`

HasWorkflowsToStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


