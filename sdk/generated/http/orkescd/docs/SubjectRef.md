# SubjectRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | Pointer to **string** | User, role or group | [optional] 

## Methods

### NewSubjectRef

`func NewSubjectRef(id string, ) *SubjectRef`

NewSubjectRef instantiates a new SubjectRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectRefWithDefaults

`func NewSubjectRefWithDefaults() *SubjectRef`

NewSubjectRefWithDefaults instantiates a new SubjectRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubjectRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubjectRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubjectRef) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SubjectRef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubjectRef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubjectRef) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubjectRef) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


