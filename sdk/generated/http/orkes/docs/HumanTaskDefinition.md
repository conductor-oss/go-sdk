# HumanTaskDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentCompletionStrategy** | Pointer to **string** |  | [optional] 
**Assignments** | Pointer to [**[]HumanTaskAssignment**](HumanTaskAssignment.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**FullTemplate** | Pointer to [**HumanTaskTemplate**](HumanTaskTemplate.md) |  | [optional] 
**TaskTriggers** | Pointer to [**[]HumanTaskTrigger**](HumanTaskTrigger.md) |  | [optional] 
**UserFormTemplate** | Pointer to [**UserFormTemplate**](UserFormTemplate.md) |  | [optional] 

## Methods

### NewHumanTaskDefinition

`func NewHumanTaskDefinition() *HumanTaskDefinition`

NewHumanTaskDefinition instantiates a new HumanTaskDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHumanTaskDefinitionWithDefaults

`func NewHumanTaskDefinitionWithDefaults() *HumanTaskDefinition`

NewHumanTaskDefinitionWithDefaults instantiates a new HumanTaskDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentCompletionStrategy

`func (o *HumanTaskDefinition) GetAssignmentCompletionStrategy() string`

GetAssignmentCompletionStrategy returns the AssignmentCompletionStrategy field if non-nil, zero value otherwise.

### GetAssignmentCompletionStrategyOk

`func (o *HumanTaskDefinition) GetAssignmentCompletionStrategyOk() (*string, bool)`

GetAssignmentCompletionStrategyOk returns a tuple with the AssignmentCompletionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentCompletionStrategy

`func (o *HumanTaskDefinition) SetAssignmentCompletionStrategy(v string)`

SetAssignmentCompletionStrategy sets AssignmentCompletionStrategy field to given value.

### HasAssignmentCompletionStrategy

`func (o *HumanTaskDefinition) HasAssignmentCompletionStrategy() bool`

HasAssignmentCompletionStrategy returns a boolean if a field has been set.

### GetAssignments

`func (o *HumanTaskDefinition) GetAssignments() []HumanTaskAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *HumanTaskDefinition) GetAssignmentsOk() (*[]HumanTaskAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *HumanTaskDefinition) SetAssignments(v []HumanTaskAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *HumanTaskDefinition) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetDisplayName

`func (o *HumanTaskDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HumanTaskDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HumanTaskDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *HumanTaskDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFullTemplate

`func (o *HumanTaskDefinition) GetFullTemplate() HumanTaskTemplate`

GetFullTemplate returns the FullTemplate field if non-nil, zero value otherwise.

### GetFullTemplateOk

`func (o *HumanTaskDefinition) GetFullTemplateOk() (*HumanTaskTemplate, bool)`

GetFullTemplateOk returns a tuple with the FullTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTemplate

`func (o *HumanTaskDefinition) SetFullTemplate(v HumanTaskTemplate)`

SetFullTemplate sets FullTemplate field to given value.

### HasFullTemplate

`func (o *HumanTaskDefinition) HasFullTemplate() bool`

HasFullTemplate returns a boolean if a field has been set.

### GetTaskTriggers

`func (o *HumanTaskDefinition) GetTaskTriggers() []HumanTaskTrigger`

GetTaskTriggers returns the TaskTriggers field if non-nil, zero value otherwise.

### GetTaskTriggersOk

`func (o *HumanTaskDefinition) GetTaskTriggersOk() (*[]HumanTaskTrigger, bool)`

GetTaskTriggersOk returns a tuple with the TaskTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTriggers

`func (o *HumanTaskDefinition) SetTaskTriggers(v []HumanTaskTrigger)`

SetTaskTriggers sets TaskTriggers field to given value.

### HasTaskTriggers

`func (o *HumanTaskDefinition) HasTaskTriggers() bool`

HasTaskTriggers returns a boolean if a field has been set.

### GetUserFormTemplate

`func (o *HumanTaskDefinition) GetUserFormTemplate() UserFormTemplate`

GetUserFormTemplate returns the UserFormTemplate field if non-nil, zero value otherwise.

### GetUserFormTemplateOk

`func (o *HumanTaskDefinition) GetUserFormTemplateOk() (*UserFormTemplate, bool)`

GetUserFormTemplateOk returns a tuple with the UserFormTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFormTemplate

`func (o *HumanTaskDefinition) SetUserFormTemplate(v UserFormTemplate)`

SetUserFormTemplate sets UserFormTemplate field to given value.

### HasUserFormTemplate

`func (o *HumanTaskDefinition) HasUserFormTemplate() bool`

HasUserFormTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


