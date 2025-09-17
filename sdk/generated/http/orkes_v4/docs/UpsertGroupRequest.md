# UpsertGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultAccess** | Pointer to  | a default Map&lt;TargetType, Set&lt;Access&gt; to share permissions, allowed target types: WORKFLOW_DEF, TASK_DEF, WORKFLOW_SCHEDULE | [optional] 
**Description** | Pointer to **string** | A general description of the group | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpsertGroupRequest

`func NewUpsertGroupRequest() *UpsertGroupRequest`

NewUpsertGroupRequest instantiates a new UpsertGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertGroupRequestWithDefaults

`func NewUpsertGroupRequestWithDefaults() *UpsertGroupRequest`

NewUpsertGroupRequestWithDefaults instantiates a new UpsertGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultAccess

`func (o *UpsertGroupRequest) GetDefaultAccess() map[string]interface{}`

GetDefaultAccess returns the DefaultAccess field if non-nil, zero value otherwise.

### GetDefaultAccessOk

`func (o *UpsertGroupRequest) GetDefaultAccessOk() (*map[string]interface{}, bool)`

GetDefaultAccessOk returns a tuple with the DefaultAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccess

`func (o *UpsertGroupRequest) SetDefaultAccess(v map[string]interface{})`

SetDefaultAccess sets DefaultAccess field to given value.

### HasDefaultAccess

`func (o *UpsertGroupRequest) HasDefaultAccess() bool`

HasDefaultAccess returns a boolean if a field has been set.

### SetDefaultAccessNil

`func (o *UpsertGroupRequest) SetDefaultAccessNil(b bool)`

 SetDefaultAccessNil sets the value for DefaultAccess to be an explicit nil

### UnsetDefaultAccess
`func (o *UpsertGroupRequest) UnsetDefaultAccess()`

UnsetDefaultAccess ensures that no value is present for DefaultAccess, not even an explicit nil
### GetDescription

`func (o *UpsertGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpsertGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpsertGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpsertGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRoles

`func (o *UpsertGroupRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpsertGroupRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpsertGroupRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpsertGroupRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


