# UpsertUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | User&#39;s full name | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpsertUserRequest

`func NewUpsertUserRequest() *UpsertUserRequest`

NewUpsertUserRequest instantiates a new UpsertUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertUserRequestWithDefaults

`func NewUpsertUserRequestWithDefaults() *UpsertUserRequest`

NewUpsertUserRequestWithDefaults instantiates a new UpsertUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *UpsertUserRequest) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UpsertUserRequest) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UpsertUserRequest) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UpsertUserRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetName

`func (o *UpsertUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpsertUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpsertUserRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpsertUserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *UpsertUserRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpsertUserRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpsertUserRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpsertUserRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


