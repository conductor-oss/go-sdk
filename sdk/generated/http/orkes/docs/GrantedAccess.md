# GrantedAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **[]string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Target** | Pointer to [**TargetRef**](TargetRef.md) |  | [optional] 

## Methods

### NewGrantedAccess

`func NewGrantedAccess() *GrantedAccess`

NewGrantedAccess instantiates a new GrantedAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantedAccessWithDefaults

`func NewGrantedAccessWithDefaults() *GrantedAccess`

NewGrantedAccessWithDefaults instantiates a new GrantedAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *GrantedAccess) GetAccess() []string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *GrantedAccess) GetAccessOk() (*[]string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *GrantedAccess) SetAccess(v []string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *GrantedAccess) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetTag

`func (o *GrantedAccess) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GrantedAccess) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GrantedAccess) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GrantedAccess) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTarget

`func (o *GrantedAccess) GetTarget() TargetRef`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GrantedAccess) GetTargetOk() (*TargetRef, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GrantedAccess) SetTarget(v TargetRef)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *GrantedAccess) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


