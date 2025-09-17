# ConductorUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationUser** | Pointer to **bool** |  | [optional] 
**EncryptedId** | Pointer to **bool** |  | [optional] 
**EncryptedIdDisplayValue** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]Group**](Group.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrkesWorkersApp** | Pointer to **bool** |  | [optional] 
**Roles** | Pointer to [**[]Role**](Role.md) |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewConductorUser

`func NewConductorUser() *ConductorUser`

NewConductorUser instantiates a new ConductorUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConductorUserWithDefaults

`func NewConductorUserWithDefaults() *ConductorUser`

NewConductorUserWithDefaults instantiates a new ConductorUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationUser

`func (o *ConductorUser) GetApplicationUser() bool`

GetApplicationUser returns the ApplicationUser field if non-nil, zero value otherwise.

### GetApplicationUserOk

`func (o *ConductorUser) GetApplicationUserOk() (*bool, bool)`

GetApplicationUserOk returns a tuple with the ApplicationUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUser

`func (o *ConductorUser) SetApplicationUser(v bool)`

SetApplicationUser sets ApplicationUser field to given value.

### HasApplicationUser

`func (o *ConductorUser) HasApplicationUser() bool`

HasApplicationUser returns a boolean if a field has been set.

### GetEncryptedId

`func (o *ConductorUser) GetEncryptedId() bool`

GetEncryptedId returns the EncryptedId field if non-nil, zero value otherwise.

### GetEncryptedIdOk

`func (o *ConductorUser) GetEncryptedIdOk() (*bool, bool)`

GetEncryptedIdOk returns a tuple with the EncryptedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedId

`func (o *ConductorUser) SetEncryptedId(v bool)`

SetEncryptedId sets EncryptedId field to given value.

### HasEncryptedId

`func (o *ConductorUser) HasEncryptedId() bool`

HasEncryptedId returns a boolean if a field has been set.

### GetEncryptedIdDisplayValue

`func (o *ConductorUser) GetEncryptedIdDisplayValue() string`

GetEncryptedIdDisplayValue returns the EncryptedIdDisplayValue field if non-nil, zero value otherwise.

### GetEncryptedIdDisplayValueOk

`func (o *ConductorUser) GetEncryptedIdDisplayValueOk() (*string, bool)`

GetEncryptedIdDisplayValueOk returns a tuple with the EncryptedIdDisplayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedIdDisplayValue

`func (o *ConductorUser) SetEncryptedIdDisplayValue(v string)`

SetEncryptedIdDisplayValue sets EncryptedIdDisplayValue field to given value.

### HasEncryptedIdDisplayValue

`func (o *ConductorUser) HasEncryptedIdDisplayValue() bool`

HasEncryptedIdDisplayValue returns a boolean if a field has been set.

### GetGroups

`func (o *ConductorUser) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ConductorUser) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ConductorUser) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ConductorUser) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetId

`func (o *ConductorUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConductorUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConductorUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConductorUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConductorUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConductorUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConductorUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConductorUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrkesWorkersApp

`func (o *ConductorUser) GetOrkesWorkersApp() bool`

GetOrkesWorkersApp returns the OrkesWorkersApp field if non-nil, zero value otherwise.

### GetOrkesWorkersAppOk

`func (o *ConductorUser) GetOrkesWorkersAppOk() (*bool, bool)`

GetOrkesWorkersAppOk returns a tuple with the OrkesWorkersApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrkesWorkersApp

`func (o *ConductorUser) SetOrkesWorkersApp(v bool)`

SetOrkesWorkersApp sets OrkesWorkersApp field to given value.

### HasOrkesWorkersApp

`func (o *ConductorUser) HasOrkesWorkersApp() bool`

HasOrkesWorkersApp returns a boolean if a field has been set.

### GetRoles

`func (o *ConductorUser) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ConductorUser) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ConductorUser) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ConductorUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUuid

`func (o *ConductorUser) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ConductorUser) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ConductorUser) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ConductorUser) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


