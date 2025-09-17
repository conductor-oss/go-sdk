# HumanTaskEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to [**HumanTaskUser**](HumanTaskUser.md) |  | [optional] 
**Claimant** | Pointer to [**HumanTaskUser**](HumanTaskUser.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **int64** |  | [optional] 
**DefinitionName** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**HumanTaskDef** | Pointer to [**HumanTaskDefinition**](HumanTaskDefinition.md) |  | [optional] 
**Input** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Output** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**OwnerApp** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] 
**TaskRefName** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**UpdatedOn** | Pointer to **int64** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**WorkflowName** | Pointer to **string** |  | [optional] 

## Methods

### NewHumanTaskEntry

`func NewHumanTaskEntry() *HumanTaskEntry`

NewHumanTaskEntry instantiates a new HumanTaskEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHumanTaskEntryWithDefaults

`func NewHumanTaskEntryWithDefaults() *HumanTaskEntry`

NewHumanTaskEntryWithDefaults instantiates a new HumanTaskEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *HumanTaskEntry) GetAssignee() HumanTaskUser`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *HumanTaskEntry) GetAssigneeOk() (*HumanTaskUser, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *HumanTaskEntry) SetAssignee(v HumanTaskUser)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *HumanTaskEntry) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetClaimant

`func (o *HumanTaskEntry) GetClaimant() HumanTaskUser`

GetClaimant returns the Claimant field if non-nil, zero value otherwise.

### GetClaimantOk

`func (o *HumanTaskEntry) GetClaimantOk() (*HumanTaskUser, bool)`

GetClaimantOk returns a tuple with the Claimant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimant

`func (o *HumanTaskEntry) SetClaimant(v HumanTaskUser)`

SetClaimant sets Claimant field to given value.

### HasClaimant

`func (o *HumanTaskEntry) HasClaimant() bool`

HasClaimant returns a boolean if a field has been set.

### GetCreatedBy

`func (o *HumanTaskEntry) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *HumanTaskEntry) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *HumanTaskEntry) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *HumanTaskEntry) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *HumanTaskEntry) GetCreatedOn() int64`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *HumanTaskEntry) GetCreatedOnOk() (*int64, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *HumanTaskEntry) SetCreatedOn(v int64)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *HumanTaskEntry) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDefinitionName

`func (o *HumanTaskEntry) GetDefinitionName() string`

GetDefinitionName returns the DefinitionName field if non-nil, zero value otherwise.

### GetDefinitionNameOk

`func (o *HumanTaskEntry) GetDefinitionNameOk() (*string, bool)`

GetDefinitionNameOk returns a tuple with the DefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionName

`func (o *HumanTaskEntry) SetDefinitionName(v string)`

SetDefinitionName sets DefinitionName field to given value.

### HasDefinitionName

`func (o *HumanTaskEntry) HasDefinitionName() bool`

HasDefinitionName returns a boolean if a field has been set.

### GetDisplayName

`func (o *HumanTaskEntry) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HumanTaskEntry) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HumanTaskEntry) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *HumanTaskEntry) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHumanTaskDef

`func (o *HumanTaskEntry) GetHumanTaskDef() HumanTaskDefinition`

GetHumanTaskDef returns the HumanTaskDef field if non-nil, zero value otherwise.

### GetHumanTaskDefOk

`func (o *HumanTaskEntry) GetHumanTaskDefOk() (*HumanTaskDefinition, bool)`

GetHumanTaskDefOk returns a tuple with the HumanTaskDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanTaskDef

`func (o *HumanTaskEntry) SetHumanTaskDef(v HumanTaskDefinition)`

SetHumanTaskDef sets HumanTaskDef field to given value.

### HasHumanTaskDef

`func (o *HumanTaskEntry) HasHumanTaskDef() bool`

HasHumanTaskDef returns a boolean if a field has been set.

### GetInput

`func (o *HumanTaskEntry) GetInput() map[string]map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *HumanTaskEntry) GetInputOk() (*map[string]map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *HumanTaskEntry) SetInput(v map[string]map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *HumanTaskEntry) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *HumanTaskEntry) GetOutput() map[string]map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *HumanTaskEntry) GetOutputOk() (*map[string]map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *HumanTaskEntry) SetOutput(v map[string]map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *HumanTaskEntry) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetOwnerApp

`func (o *HumanTaskEntry) GetOwnerApp() string`

GetOwnerApp returns the OwnerApp field if non-nil, zero value otherwise.

### GetOwnerAppOk

`func (o *HumanTaskEntry) GetOwnerAppOk() (*string, bool)`

GetOwnerAppOk returns a tuple with the OwnerApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerApp

`func (o *HumanTaskEntry) SetOwnerApp(v string)`

SetOwnerApp sets OwnerApp field to given value.

### HasOwnerApp

`func (o *HumanTaskEntry) HasOwnerApp() bool`

HasOwnerApp returns a boolean if a field has been set.

### GetState

`func (o *HumanTaskEntry) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HumanTaskEntry) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HumanTaskEntry) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HumanTaskEntry) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTaskId

`func (o *HumanTaskEntry) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *HumanTaskEntry) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *HumanTaskEntry) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *HumanTaskEntry) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTaskRefName

`func (o *HumanTaskEntry) GetTaskRefName() string`

GetTaskRefName returns the TaskRefName field if non-nil, zero value otherwise.

### GetTaskRefNameOk

`func (o *HumanTaskEntry) GetTaskRefNameOk() (*string, bool)`

GetTaskRefNameOk returns a tuple with the TaskRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRefName

`func (o *HumanTaskEntry) SetTaskRefName(v string)`

SetTaskRefName sets TaskRefName field to given value.

### HasTaskRefName

`func (o *HumanTaskEntry) HasTaskRefName() bool`

HasTaskRefName returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *HumanTaskEntry) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *HumanTaskEntry) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *HumanTaskEntry) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *HumanTaskEntry) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *HumanTaskEntry) GetUpdatedOn() int64`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *HumanTaskEntry) GetUpdatedOnOk() (*int64, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *HumanTaskEntry) SetUpdatedOn(v int64)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *HumanTaskEntry) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetWorkflowId

`func (o *HumanTaskEntry) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *HumanTaskEntry) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *HumanTaskEntry) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *HumanTaskEntry) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowName

`func (o *HumanTaskEntry) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *HumanTaskEntry) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *HumanTaskEntry) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.

### HasWorkflowName

`func (o *HumanTaskEntry) HasWorkflowName() bool`

HasWorkflowName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


