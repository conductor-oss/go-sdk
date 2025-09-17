# ConnectivityTestInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to  |  | [optional] 
**Sink** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectivityTestInput

`func NewConnectivityTestInput() *ConnectivityTestInput`

NewConnectivityTestInput instantiates a new ConnectivityTestInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectivityTestInputWithDefaults

`func NewConnectivityTestInputWithDefaults() *ConnectivityTestInput`

NewConnectivityTestInputWithDefaults instantiates a new ConnectivityTestInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *ConnectivityTestInput) GetInput() map[string]map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ConnectivityTestInput) GetInputOk() (*map[string]map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ConnectivityTestInput) SetInput(v map[string]map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ConnectivityTestInput) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *ConnectivityTestInput) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ConnectivityTestInput) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetSink

`func (o *ConnectivityTestInput) GetSink() string`

GetSink returns the Sink field if non-nil, zero value otherwise.

### GetSinkOk

`func (o *ConnectivityTestInput) GetSinkOk() (*string, bool)`

GetSinkOk returns a tuple with the Sink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSink

`func (o *ConnectivityTestInput) SetSink(v string)`

SetSink sets Sink field to given value.

### HasSink

`func (o *ConnectivityTestInput) HasSink() bool`

HasSink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


