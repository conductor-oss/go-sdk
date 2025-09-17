# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultAccess** | Pointer to  |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to [**[]Role**](Role.md) |  | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultAccess

`func (o *Group) GetDefaultAccess() map[string]interface{}`

GetDefaultAccess returns the DefaultAccess field if non-nil, zero value otherwise.

### GetDefaultAccessOk

`func (o *Group) GetDefaultAccessOk() (*map[string]interface{}, bool)`

GetDefaultAccessOk returns a tuple with the DefaultAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccess

`func (o *Group) SetDefaultAccess(v map[string]interface{})`

SetDefaultAccess sets DefaultAccess field to given value.

### HasDefaultAccess

`func (o *Group) HasDefaultAccess() bool`

HasDefaultAccess returns a boolean if a field has been set.

### SetDefaultAccessNil

`func (o *Group) SetDefaultAccessNil(b bool)`

 SetDefaultAccessNil sets the value for DefaultAccess to be an explicit nil

### UnsetDefaultAccess
`func (o *Group) UnsetDefaultAccess()`

UnsetDefaultAccess ensures that no value is present for DefaultAccess, not even an explicit nil
### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoles

`func (o *Group) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Group) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Group) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Group) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


