# HealthCheckStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthResults** | Pointer to [**[]Health**](Health.md) |  | [optional] 
**SuppressedHealthResults** | Pointer to [**[]Health**](Health.md) |  | [optional] 
**Healthy** | Pointer to **bool** |  | [optional] 

## Methods

### NewHealthCheckStatus

`func NewHealthCheckStatus() *HealthCheckStatus`

NewHealthCheckStatus instantiates a new HealthCheckStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckStatusWithDefaults

`func NewHealthCheckStatusWithDefaults() *HealthCheckStatus`

NewHealthCheckStatusWithDefaults instantiates a new HealthCheckStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthResults

`func (o *HealthCheckStatus) GetHealthResults() []Health`

GetHealthResults returns the HealthResults field if non-nil, zero value otherwise.

### GetHealthResultsOk

`func (o *HealthCheckStatus) GetHealthResultsOk() (*[]Health, bool)`

GetHealthResultsOk returns a tuple with the HealthResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthResults

`func (o *HealthCheckStatus) SetHealthResults(v []Health)`

SetHealthResults sets HealthResults field to given value.

### HasHealthResults

`func (o *HealthCheckStatus) HasHealthResults() bool`

HasHealthResults returns a boolean if a field has been set.

### GetSuppressedHealthResults

`func (o *HealthCheckStatus) GetSuppressedHealthResults() []Health`

GetSuppressedHealthResults returns the SuppressedHealthResults field if non-nil, zero value otherwise.

### GetSuppressedHealthResultsOk

`func (o *HealthCheckStatus) GetSuppressedHealthResultsOk() (*[]Health, bool)`

GetSuppressedHealthResultsOk returns a tuple with the SuppressedHealthResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedHealthResults

`func (o *HealthCheckStatus) SetSuppressedHealthResults(v []Health)`

SetSuppressedHealthResults sets SuppressedHealthResults field to given value.

### HasSuppressedHealthResults

`func (o *HealthCheckStatus) HasSuppressedHealthResults() bool`

HasSuppressedHealthResults returns a boolean if a field has been set.

### GetHealthy

`func (o *HealthCheckStatus) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *HealthCheckStatus) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *HealthCheckStatus) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *HealthCheckStatus) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


