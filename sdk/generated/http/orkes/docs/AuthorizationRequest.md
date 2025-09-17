# AuthorizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **[]string** | The set of access which is granted or removed | 
**Subject** | [**SubjectRef**](SubjectRef.md) |  | 
**Target** | [**TargetRef**](TargetRef.md) |  | 

## Methods

### NewAuthorizationRequest

`func NewAuthorizationRequest(access []string, subject SubjectRef, target TargetRef, ) *AuthorizationRequest`

NewAuthorizationRequest instantiates a new AuthorizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationRequestWithDefaults

`func NewAuthorizationRequestWithDefaults() *AuthorizationRequest`

NewAuthorizationRequestWithDefaults instantiates a new AuthorizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *AuthorizationRequest) GetAccess() []string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AuthorizationRequest) GetAccessOk() (*[]string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AuthorizationRequest) SetAccess(v []string)`

SetAccess sets Access field to given value.


### GetSubject

`func (o *AuthorizationRequest) GetSubject() SubjectRef`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AuthorizationRequest) GetSubjectOk() (*SubjectRef, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AuthorizationRequest) SetSubject(v SubjectRef)`

SetSubject sets Subject field to given value.


### GetTarget

`func (o *AuthorizationRequest) GetTarget() TargetRef`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AuthorizationRequest) GetTargetOk() (*TargetRef, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AuthorizationRequest) SetTarget(v TargetRef)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


