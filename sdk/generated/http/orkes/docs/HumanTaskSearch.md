# HumanTaskSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | Pointer to [**[]HumanTaskUser**](HumanTaskUser.md) |  | [optional] 
**Claimants** | Pointer to [**[]HumanTaskUser**](HumanTaskUser.md) |  | [optional] 
**DefinitionNames** | Pointer to **[]string** |  | [optional] 
**DisplayNames** | Pointer to **[]string** |  | [optional] 
**FullTextQuery** | Pointer to **string** |  | [optional] 
**SearchType** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**States** | Pointer to **[]string** |  | [optional] 
**TaskInputQuery** | Pointer to **string** |  | [optional] 
**TaskOutputQuery** | Pointer to **string** |  | [optional] 
**TaskRefNames** | Pointer to **[]string** |  | [optional] 
**UpdateEndTime** | Pointer to **int64** |  | [optional] 
**UpdateStartTime** | Pointer to **int64** |  | [optional] 
**WorkflowIds** | Pointer to **[]string** |  | [optional] 
**WorkflowNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHumanTaskSearch

`func NewHumanTaskSearch() *HumanTaskSearch`

NewHumanTaskSearch instantiates a new HumanTaskSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHumanTaskSearchWithDefaults

`func NewHumanTaskSearchWithDefaults() *HumanTaskSearch`

NewHumanTaskSearchWithDefaults instantiates a new HumanTaskSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *HumanTaskSearch) GetAssignees() []HumanTaskUser`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *HumanTaskSearch) GetAssigneesOk() (*[]HumanTaskUser, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *HumanTaskSearch) SetAssignees(v []HumanTaskUser)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *HumanTaskSearch) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetClaimants

`func (o *HumanTaskSearch) GetClaimants() []HumanTaskUser`

GetClaimants returns the Claimants field if non-nil, zero value otherwise.

### GetClaimantsOk

`func (o *HumanTaskSearch) GetClaimantsOk() (*[]HumanTaskUser, bool)`

GetClaimantsOk returns a tuple with the Claimants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimants

`func (o *HumanTaskSearch) SetClaimants(v []HumanTaskUser)`

SetClaimants sets Claimants field to given value.

### HasClaimants

`func (o *HumanTaskSearch) HasClaimants() bool`

HasClaimants returns a boolean if a field has been set.

### GetDefinitionNames

`func (o *HumanTaskSearch) GetDefinitionNames() []string`

GetDefinitionNames returns the DefinitionNames field if non-nil, zero value otherwise.

### GetDefinitionNamesOk

`func (o *HumanTaskSearch) GetDefinitionNamesOk() (*[]string, bool)`

GetDefinitionNamesOk returns a tuple with the DefinitionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionNames

`func (o *HumanTaskSearch) SetDefinitionNames(v []string)`

SetDefinitionNames sets DefinitionNames field to given value.

### HasDefinitionNames

`func (o *HumanTaskSearch) HasDefinitionNames() bool`

HasDefinitionNames returns a boolean if a field has been set.

### GetDisplayNames

`func (o *HumanTaskSearch) GetDisplayNames() []string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *HumanTaskSearch) GetDisplayNamesOk() (*[]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *HumanTaskSearch) SetDisplayNames(v []string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *HumanTaskSearch) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### GetFullTextQuery

`func (o *HumanTaskSearch) GetFullTextQuery() string`

GetFullTextQuery returns the FullTextQuery field if non-nil, zero value otherwise.

### GetFullTextQueryOk

`func (o *HumanTaskSearch) GetFullTextQueryOk() (*string, bool)`

GetFullTextQueryOk returns a tuple with the FullTextQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTextQuery

`func (o *HumanTaskSearch) SetFullTextQuery(v string)`

SetFullTextQuery sets FullTextQuery field to given value.

### HasFullTextQuery

`func (o *HumanTaskSearch) HasFullTextQuery() bool`

HasFullTextQuery returns a boolean if a field has been set.

### GetSearchType

`func (o *HumanTaskSearch) GetSearchType() string`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *HumanTaskSearch) GetSearchTypeOk() (*string, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *HumanTaskSearch) SetSearchType(v string)`

SetSearchType sets SearchType field to given value.

### HasSearchType

`func (o *HumanTaskSearch) HasSearchType() bool`

HasSearchType returns a boolean if a field has been set.

### GetSize

`func (o *HumanTaskSearch) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HumanTaskSearch) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HumanTaskSearch) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *HumanTaskSearch) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStart

`func (o *HumanTaskSearch) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *HumanTaskSearch) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *HumanTaskSearch) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *HumanTaskSearch) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStates

`func (o *HumanTaskSearch) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *HumanTaskSearch) GetStatesOk() (*[]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *HumanTaskSearch) SetStates(v []string)`

SetStates sets States field to given value.

### HasStates

`func (o *HumanTaskSearch) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetTaskInputQuery

`func (o *HumanTaskSearch) GetTaskInputQuery() string`

GetTaskInputQuery returns the TaskInputQuery field if non-nil, zero value otherwise.

### GetTaskInputQueryOk

`func (o *HumanTaskSearch) GetTaskInputQueryOk() (*string, bool)`

GetTaskInputQueryOk returns a tuple with the TaskInputQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInputQuery

`func (o *HumanTaskSearch) SetTaskInputQuery(v string)`

SetTaskInputQuery sets TaskInputQuery field to given value.

### HasTaskInputQuery

`func (o *HumanTaskSearch) HasTaskInputQuery() bool`

HasTaskInputQuery returns a boolean if a field has been set.

### GetTaskOutputQuery

`func (o *HumanTaskSearch) GetTaskOutputQuery() string`

GetTaskOutputQuery returns the TaskOutputQuery field if non-nil, zero value otherwise.

### GetTaskOutputQueryOk

`func (o *HumanTaskSearch) GetTaskOutputQueryOk() (*string, bool)`

GetTaskOutputQueryOk returns a tuple with the TaskOutputQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOutputQuery

`func (o *HumanTaskSearch) SetTaskOutputQuery(v string)`

SetTaskOutputQuery sets TaskOutputQuery field to given value.

### HasTaskOutputQuery

`func (o *HumanTaskSearch) HasTaskOutputQuery() bool`

HasTaskOutputQuery returns a boolean if a field has been set.

### GetTaskRefNames

`func (o *HumanTaskSearch) GetTaskRefNames() []string`

GetTaskRefNames returns the TaskRefNames field if non-nil, zero value otherwise.

### GetTaskRefNamesOk

`func (o *HumanTaskSearch) GetTaskRefNamesOk() (*[]string, bool)`

GetTaskRefNamesOk returns a tuple with the TaskRefNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRefNames

`func (o *HumanTaskSearch) SetTaskRefNames(v []string)`

SetTaskRefNames sets TaskRefNames field to given value.

### HasTaskRefNames

`func (o *HumanTaskSearch) HasTaskRefNames() bool`

HasTaskRefNames returns a boolean if a field has been set.

### GetUpdateEndTime

`func (o *HumanTaskSearch) GetUpdateEndTime() int64`

GetUpdateEndTime returns the UpdateEndTime field if non-nil, zero value otherwise.

### GetUpdateEndTimeOk

`func (o *HumanTaskSearch) GetUpdateEndTimeOk() (*int64, bool)`

GetUpdateEndTimeOk returns a tuple with the UpdateEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateEndTime

`func (o *HumanTaskSearch) SetUpdateEndTime(v int64)`

SetUpdateEndTime sets UpdateEndTime field to given value.

### HasUpdateEndTime

`func (o *HumanTaskSearch) HasUpdateEndTime() bool`

HasUpdateEndTime returns a boolean if a field has been set.

### GetUpdateStartTime

`func (o *HumanTaskSearch) GetUpdateStartTime() int64`

GetUpdateStartTime returns the UpdateStartTime field if non-nil, zero value otherwise.

### GetUpdateStartTimeOk

`func (o *HumanTaskSearch) GetUpdateStartTimeOk() (*int64, bool)`

GetUpdateStartTimeOk returns a tuple with the UpdateStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStartTime

`func (o *HumanTaskSearch) SetUpdateStartTime(v int64)`

SetUpdateStartTime sets UpdateStartTime field to given value.

### HasUpdateStartTime

`func (o *HumanTaskSearch) HasUpdateStartTime() bool`

HasUpdateStartTime returns a boolean if a field has been set.

### GetWorkflowIds

`func (o *HumanTaskSearch) GetWorkflowIds() []string`

GetWorkflowIds returns the WorkflowIds field if non-nil, zero value otherwise.

### GetWorkflowIdsOk

`func (o *HumanTaskSearch) GetWorkflowIdsOk() (*[]string, bool)`

GetWorkflowIdsOk returns a tuple with the WorkflowIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowIds

`func (o *HumanTaskSearch) SetWorkflowIds(v []string)`

SetWorkflowIds sets WorkflowIds field to given value.

### HasWorkflowIds

`func (o *HumanTaskSearch) HasWorkflowIds() bool`

HasWorkflowIds returns a boolean if a field has been set.

### GetWorkflowNames

`func (o *HumanTaskSearch) GetWorkflowNames() []string`

GetWorkflowNames returns the WorkflowNames field if non-nil, zero value otherwise.

### GetWorkflowNamesOk

`func (o *HumanTaskSearch) GetWorkflowNamesOk() (*[]string, bool)`

GetWorkflowNamesOk returns a tuple with the WorkflowNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowNames

`func (o *HumanTaskSearch) SetWorkflowNames(v []string)`

SetWorkflowNames sets WorkflowNames field to given value.

### HasWorkflowNames

`func (o *HumanTaskSearch) HasWorkflowNames() bool`

HasWorkflowNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


