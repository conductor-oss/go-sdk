//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package client

//nolint:gocognit,gocyclo // This file contains generated mapper functions with high complexity
import (
	"encoding/json"
	"math"

	"github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
	"github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/human"
	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
)

// ============================================================================
// CONSOLIDATED DTO MAPPERS WITH CONSISTENT NAMING
// ============================================================================
// Naming Convention:
// - toDomain* : Generated models -> Domain models
// - toGenerated* : Domain models -> Generated models
// ============================================================================

// toGeneratedAuthorizationRequest converts rbac.AuthorizationRequest to orkes.AuthorizationRequest
func toGeneratedAuthorizationRequest(domainReq rbac.AuthorizationRequest) orkes.AuthorizationRequest {
	return orkes.AuthorizationRequest{
		Access: domainReq.Access,
		Subject: orkes.SubjectRef{
			Type: &domainReq.Subject.Type_,
			Id:   domainReq.Subject.Id,
		},
		Target: orkes.TargetRef{
			Type: domainReq.Target.Type_,
			Id:   domainReq.Target.Id,
		},
	}
}

// toDomainConductorUser converts map[string]interface{} to rbac.ConductorUser
func toDomainConductorUser(genUser map[string]interface{}) *rbac.ConductorUser {
	if genUser == nil {
		return nil
	}

	domain := &rbac.ConductorUser{}

	// Map boolean fields
	if val, ok := genUser["applicationUser"].(bool); ok {
		domain.ApplicationUser = val
	}
	if val, ok := genUser["encryptedId"].(bool); ok {
		domain.EncryptedId = val
	}

	// Map string fields
	if val, ok := genUser["encryptedIdDisplayValue"].(string); ok {
		domain.EncryptedIdDisplayValue = val
	}
	if val, ok := genUser["id"].(string); ok {
		domain.Id = val
	}
	if val, ok := genUser["name"].(string); ok {
		domain.Name = val
	}
	if val, ok := genUser["uuid"].(string); ok {
		domain.Uuid = val
	}

	// Map Groups: []interface{} -> []rbac.Group
	if groupsData, ok := genUser["groups"].([]interface{}); ok {
		groups := make([]rbac.Group, len(groupsData))
		for i, groupData := range groupsData {
			if groupMap, ok := groupData.(map[string]interface{}); ok {
				groups[i] = toDomainGroupFromMap(groupMap)
			}
		}
		domain.Groups = groups
	}

	// Map Roles: []interface{} -> []rbac.Role
	if rolesData, ok := genUser["roles"].([]interface{}); ok {
		roles := make([]rbac.Role, len(rolesData))
		for i, roleData := range rolesData {
			if roleMap, ok := roleData.(map[string]interface{}); ok {
				roles[i] = toDomainRoleFromMap(roleMap)
			}
		}
		domain.Roles = roles
	}

	return domain
}

// toDomainGroupFromMap converts map[string]interface{} to rbac.Group
func toDomainGroupFromMap(groupMap map[string]interface{}) rbac.Group {
	domain := rbac.Group{}

	// Map string fields
	if val, ok := groupMap["id"].(string); ok {
		domain.Id = val
	}
	if val, ok := groupMap["description"].(string); ok {
		domain.Description = val
	}

	// Map DefaultAccess: map[string]interface{} -> map[string][]string
	if defaultAccessData, ok := groupMap["defaultAccess"].(map[string]interface{}); ok {
		defaultAccess := make(map[string][]string)
		for key, value := range defaultAccessData {
			if strSlice, ok := value.([]string); ok {
				defaultAccess[key] = strSlice
			} else if interfaceSlice, ok := value.([]interface{}); ok {
				strSlice := make([]string, len(interfaceSlice))
				for i, v := range interfaceSlice {
					if str, ok := v.(string); ok {
						strSlice[i] = str
					}
				}
				defaultAccess[key] = strSlice
			}
		}
		domain.DefaultAccess = defaultAccess
	}

	// Map Roles: []interface{} -> []rbac.Role
	if rolesData, ok := groupMap["roles"].([]interface{}); ok {
		roles := make([]rbac.Role, len(rolesData))
		for i, roleData := range rolesData {
			if roleMap, ok := roleData.(map[string]interface{}); ok {
				roles[i] = toDomainRoleFromMap(roleMap)
			}
		}
		domain.Roles = roles
	}

	return domain
}

// toDomainRoleFromMap converts map[string]interface{} to rbac.Role
func toDomainRoleFromMap(roleMap map[string]interface{}) rbac.Role {
	domain := rbac.Role{}

	// Map string fields
	if val, ok := roleMap["name"].(string); ok {
		domain.Name = val
	}

	// Map Permissions: []interface{} -> []rbac.Permission
	if permissionsData, ok := roleMap["permissions"].([]interface{}); ok {
		permissions := make([]rbac.Permission, len(permissionsData))
		for i, permissionData := range permissionsData {
			if permissionMap, ok := permissionData.(map[string]interface{}); ok {
				permissions[i] = toDomainPermissionFromMap(permissionMap)
			}
		}
		domain.Permissions = permissions
	}

	return domain
}

// toDomainPermissionFromMap converts map[string]interface{} to rbac.Permission
func toDomainPermissionFromMap(permissionMap map[string]interface{}) rbac.Permission {
	domain := rbac.Permission{}

	// Map string fields
	if val, ok := permissionMap["name"].(string); ok {
		domain.Name = val
	}

	return domain
}

// toDomainConductorUsers converts []orkes.ConductorUser to []rbac.ConductorUser
func toDomainConductorUsers(genUsers []orkes.ConductorUser) []rbac.ConductorUser {
	if genUsers == nil {
		return nil
	}

	result := make([]rbac.ConductorUser, len(genUsers))
	for i, genUser := range genUsers {
		result[i] = toDomainConductorUserFromOrkes(&genUser)
	}
	return result
}

// toDomainConductorUserFromOrkes converts orkes.ConductorUser to rbac.ConductorUser
func toDomainConductorUserFromOrkes(genUser *orkes.ConductorUser) rbac.ConductorUser {
	if genUser == nil {
		return rbac.ConductorUser{}
	}

	// Comprehensive field-by-field mapping from orkes.ConductorUser to rbac.ConductorUser
	domain := rbac.ConductorUser{}

	// Map boolean fields (pointers to values)
	if genUser.ApplicationUser != nil {
		domain.ApplicationUser = *genUser.ApplicationUser
	}
	if genUser.EncryptedId != nil {
		domain.EncryptedId = *genUser.EncryptedId
	}

	// Map string fields (pointers to values)
	if genUser.EncryptedIdDisplayValue != nil {
		domain.EncryptedIdDisplayValue = *genUser.EncryptedIdDisplayValue
	}
	if genUser.Id != nil {
		domain.Id = *genUser.Id
	}
	if genUser.Name != nil {
		domain.Name = *genUser.Name
	}
	if genUser.Uuid != nil {
		domain.Uuid = *genUser.Uuid
	}

	// Map Groups: []orkes.Group -> []rbac.Group
	if genUser.Groups != nil {
		domain.Groups = toDomainGroupsFromGenerated(genUser.Groups)
	}

	// Map Roles: []orkes.Role -> []rbac.Role
	if genUser.Roles != nil {
		domain.Roles = toDomainRolesFromGenerated(genUser.Roles)
	}

	return domain
}

// toDomainRolesFromGenerated converts []orkes.Role to []rbac.Role
func toDomainRolesFromGenerated(genRoles []orkes.Role) []rbac.Role {
	if genRoles == nil {
		return nil
	}

	result := make([]rbac.Role, len(genRoles))
	for i, genRole := range genRoles {
		result[i] = toDomainRoleFromOrkes(&genRole)
	}
	return result
}

// toDomainRoleFromOrkes converts orkes.Role to rbac.Role
func toDomainRoleFromOrkes(genRole *orkes.Role) rbac.Role {
	if genRole == nil {
		return rbac.Role{}
	}

	// Comprehensive field-by-field mapping from orkes.Role to rbac.Role
	domain := rbac.Role{}

	// Map string fields (pointers to values)
	if genRole.Name != nil {
		domain.Name = *genRole.Name
	}

	// Map Permissions: []orkes.Permission -> []rbac.Permission
	if genRole.Permissions != nil {
		permissions := make([]rbac.Permission, len(genRole.Permissions))
		for i, genPermission := range genRole.Permissions {
			permissions[i] = toDomainPermissionFromOrkes(&genPermission)
		}
		domain.Permissions = permissions
	}

	return domain
}

// toDomainPermissionFromOrkes converts orkes.Permission to rbac.Permission
func toDomainPermissionFromOrkes(genPermission *orkes.Permission) rbac.Permission {
	if genPermission == nil {
		return rbac.Permission{}
	}

	// Comprehensive field-by-field mapping from orkes.Permission to rbac.Permission
	domain := rbac.Permission{}

	// Map string fields (pointers to values)
	if genPermission.Name != nil {
		domain.Name = *genPermission.Name
	}

	return domain
}

// toDomainConductorApplication converts interface{} to rbac.ConductorApplication
func toDomainConductorApplication(result interface{}) *rbac.ConductorApplication {
	if result == nil {
		return nil
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		return nil
	}

	app := &rbac.ConductorApplication{}
	if id, ok := resultMap["id"].(string); ok {
		app.Id = id
	}
	if name, ok := resultMap["name"].(string); ok {
		app.Name = name
	}
	if createdBy, ok := resultMap["createdBy"].(string); ok {
		app.CreatedBy = createdBy
	}
	if updatedBy, ok := resultMap["updatedBy"].(string); ok {
		app.UpdatedBy = updatedBy
	}
	if createTime, ok := resultMap["createTime"].(int64); ok {
		app.CreateTime = createTime
	} else if createTime, ok := resultMap["createTime"].(float64); ok {
		app.CreateTime = int64(createTime)
	}
	if updateTime, ok := resultMap["updateTime"].(int64); ok {
		app.UpdateTime = updateTime
	} else if updateTime, ok := resultMap["updateTime"].(float64); ok {
		app.UpdateTime = int64(updateTime)
	}
	return app
}

// toDomainGroupsFromGenerated converts []orkes.Group to []rbac.Group using field-by-field mapping
func toDomainGroupsFromGenerated(genGroups []orkes.Group) []rbac.Group {
	if genGroups == nil {
		return nil
	}

	result := make([]rbac.Group, len(genGroups))
	for i := range genGroups {
		genGroup := genGroups[i]
		domain := rbac.Group{}

		// DefaultAccess: *map[string][]string -> map[string][]string
		if genGroup.DefaultAccess != nil {
			// Make a shallow copy to avoid aliasing
			copied := make(map[string][]string, len(*genGroup.DefaultAccess))
			for k, v := range *genGroup.DefaultAccess {
				if v == nil {
					continue
				}
				vv := make([]string, len(v))
				copy(vv, v)
				copied[k] = vv
			}
			domain.DefaultAccess = copied
		}

		// Description: *string -> string
		if genGroup.Description != nil {
			domain.Description = *genGroup.Description
		}

		// Id: *string -> string
		if genGroup.Id != nil {
			domain.Id = *genGroup.Id
		}

		// Roles: []orkes.Role -> []rbac.Role
		if genGroup.Roles != nil {
			domain.Roles = toDomainRolesFromGenerated(genGroup.Roles)
		}

		result[i] = domain
	}
	return result
}

// toGeneratedUpsertGroupRequest converts rbac.UpsertGroupRequest to orkes.UpsertGroupRequest
func toGeneratedUpsertGroupRequest(domainReq rbac.UpsertGroupRequest) orkes.UpsertGroupRequest {
	result := orkes.UpsertGroupRequest{}

	// DefaultAccess: map[string][]string -> *map[string][]string
	if domainReq.DefaultAccess != nil {
		copied := make(map[string][]string, len(domainReq.DefaultAccess))
		for k, v := range domainReq.DefaultAccess {
			if v == nil {
				continue
			}
			vv := make([]string, len(v))
			copy(vv, v)
			copied[k] = vv
		}
		result.DefaultAccess = &copied
	}

	// Description: string -> string (required)
	result.Description = domainReq.Description

	// Roles: []string -> []string
	if domainReq.Roles != nil {
		rolesCopy := make([]string, len(domainReq.Roles))
		copy(rolesCopy, domainReq.Roles)
		result.Roles = rolesCopy
	}

	return result
}

// ============================================================================
// Secret Model Mappers
// ============================================================================

// toGeneratedTag converts model.Tag to orkes.Tag
func toGeneratedTag(domainTag model.Tag) orkes.Tag {
	return orkes.Tag{
		Key:   &domainTag.Key,
		Type:  &domainTag.Type,
		Value: &domainTag.Value,
	}
}

// toGeneratedTagFromTagObject converts model.TagObject to orkes.Tag
func toGeneratedTagFromTagObject(domainTag model.TagObject) orkes.Tag {
	return orkes.Tag{
		Key:   &domainTag.Key,
		Type:  &domainTag.Type,
		Value: &domainTag.Value,
	}
}

// toGeneratedWorkflowTask converts model.WorkflowTask to orkes.WorkflowTask
func toGeneratedWorkflowTask(domain model.WorkflowTask) orkes.WorkflowTask {
	gen := orkes.WorkflowTask{}

	// Map required fields
	gen.Name = domain.Name
	gen.TaskReferenceName = domain.TaskReferenceName

	gen.Description = ToPointer(domain.Description)
	gen.Type = ToPointer(domain.Type_)
	gen.DynamicTaskNameParam = ToPointer(domain.DynamicTaskNameParam)
	gen.CaseValueParam = ToPointer(domain.CaseValueParam)
	gen.CaseExpression = ToPointer(domain.CaseExpression)
	gen.ScriptExpression = ToPointer(domain.ScriptExpression)
	gen.EvaluatorType = ToPointer(domain.EvaluatorType)
	gen.Expression = ToPointer(domain.Expression)
	gen.StartDelay = ToPointer(domain.StartDelay)
	gen.DynamicForkJoinTasksParam = ToPointer(domain.DynamicForkJoinTasksParam)
	gen.DynamicForkTasksParam = ToPointer(domain.DynamicForkTasksParam)
	gen.DynamicForkTasksInputParamName = ToPointer(domain.DynamicForkTasksInputParamName)
	gen.Optional = ToPointer(domain.Optional)
	gen.Sink = ToPointer(domain.Sink)

	// Map complex fields that need conversion
	if domain.DecisionCases != nil {
		genDecisionCases := make(map[string][]orkes.WorkflowTask)
		for key, tasks := range domain.DecisionCases {
			genTasks := make([]orkes.WorkflowTask, len(tasks))
			for i, task := range tasks {
				genTasks[i] = toGeneratedWorkflowTask(task)
			}
			genDecisionCases[key] = genTasks
		}
		gen.DecisionCases = &genDecisionCases
	}

	if domain.DefaultCase != nil {
		genDefaultCase := make([]orkes.WorkflowTask, len(domain.DefaultCase))
		for i, task := range domain.DefaultCase {
			genDefaultCase[i] = toGeneratedWorkflowTask(task)
		}
		gen.DefaultCase = genDefaultCase
	}

	if domain.ForkTasks != nil {
		genForkTasks := make([][]orkes.WorkflowTask, len(domain.ForkTasks))
		for i, taskGroup := range domain.ForkTasks {
			genTaskGroup := make([]orkes.WorkflowTask, len(taskGroup))
			for j, task := range taskGroup {
				genTaskGroup[j] = toGeneratedWorkflowTask(task)
			}
			genForkTasks[i] = genTaskGroup
		}
		gen.ForkTasks = genForkTasks
	}

	if domain.SubWorkflowParam != nil {
		genSubWorkflowParam := toGeneratedSubWorkflowParams(domain.SubWorkflowParam)
		gen.SubWorkflowParam = &genSubWorkflowParam
	}

	if domain.TaskDefinition != nil {
		genTaskDef := toGeneratedTaskDef(domain.TaskDefinition)
		gen.TaskDefinition = &genTaskDef
	}

	if domain.InputParameters != nil {
		gen.InputParameters = domain.InputParameters
	}

	gen.DefaultExclusiveJoinTask = domain.DefaultExclusiveJoinTask

	gen.LoopCondition = ToPointer(domain.LoopCondition)
	if domain.LoopOver != nil {
		genLoopOver := make([]orkes.WorkflowTask, len(domain.LoopOver))
		for i, loopTask := range domain.LoopOver {
			genLoopOver[i] = toGeneratedWorkflowTask(loopTask)
		}
		gen.LoopOver = genLoopOver
	}

	return gen
}

// toGeneratedSubWorkflowParams converts model.SubWorkflowParams to orkes.SubWorkflowParams
func toGeneratedSubWorkflowParams(domain *model.SubWorkflowParams) orkes.SubWorkflowParams {
	if domain == nil {
		return orkes.SubWorkflowParams{}
	}

	gen := orkes.SubWorkflowParams{
		Name:         ToPointer(domain.Name),
		Version:      ToPointer(int32(domain.Version)),
		TaskToDomain: &domain.TaskToDomain,
	}

	if domain.WorkflowDefinition != nil {
		gen.WorkflowDefinition = toMapFromWorkflowDef(domain.WorkflowDefinition)
	}

	return gen
}

// toGeneratedTaskDef converts model.TaskDef to orkes.TaskDef
func toGeneratedTaskDef(domain *model.TaskDef) orkes.TaskDef {
	if domain == nil {
		return orkes.TaskDef{}
	}

	var retryLogicPtr *string
	if domain.RetryLogic != "" {
		retryLogicPtr = &domain.RetryLogic
	}

	var timeoutPolicyPtr *string
	if domain.TimeoutPolicy != "" {
		timeoutPolicyPtr = &domain.TimeoutPolicy
	}

	responseTimeoutSeconds := domain.ResponseTimeoutSeconds
	if responseTimeoutSeconds == 0 {
		responseTimeoutSeconds = 1 // Minimum required by server
	}

	backoffScaleFactor := domain.BackoffScaleFactor
	if backoffScaleFactor == 0 {
		backoffScaleFactor = 1.0 // Default value
	}

	return orkes.TaskDef{
		Name:                        domain.Name,
		Description:                 ToPointer(domain.Description),
		RetryCount:                  ToPointer(domain.RetryCount),
		TimeoutSeconds:              domain.TimeoutSeconds,
		TotalTimeoutSeconds:         domain.TimeoutSeconds, // Use same value for both
		InputKeys:                   domain.InputKeys,
		OutputKeys:                  domain.OutputKeys,
		TimeoutPolicy:               timeoutPolicyPtr,
		RetryLogic:                  retryLogicPtr,
		RetryDelaySeconds:           ToPointer(domain.RetryDelaySeconds),
		ResponseTimeoutSeconds:      ToPointer(responseTimeoutSeconds),
		ConcurrentExecLimit:         ToPointer(domain.ConcurrentExecLimit),
		InputTemplate:               domain.InputTemplate,
		RateLimitPerFrequency:       ToPointer(domain.RateLimitPerFrequency),
		RateLimitFrequencyInSeconds: ToPointer(domain.RateLimitFrequencyInSeconds),
		IsolationGroupId:            ToPointer(domain.IsolationGroupId),
		ExecutionNameSpace:          ToPointer(domain.ExecutionNameSpace),
		OwnerEmail:                  ToPointer(domain.OwnerEmail),
		PollTimeoutSeconds:          ToPointer(domain.PollTimeoutSeconds),
		BackoffScaleFactor:          ToPointer(backoffScaleFactor),
	}
}

// toGeneratedTags converts []model.Tag to []orkes.Tag
func toGeneratedTags(domainTags []model.Tag) []orkes.Tag {
	result := make([]orkes.Tag, len(domainTags))
	for i, domainTag := range domainTags {
		result[i] = toGeneratedTag(domainTag)
	}
	return result
}

// toDomainTagsFromGenerated converts []orkes.Tag to []model.Tag
func toDomainTagsFromGenerated(genTags []orkes.Tag) []model.Tag {
	result := make([]model.Tag, len(genTags))
	for i, genTag := range genTags {
		result[i] = toDomainTagFromGenerated(&genTag)
	}
	return result
}

// toDomainSecretsFromGenerated converts []orkes.ExtendedSecret to []model.Secret
func toDomainSecretsFromGenerated(genSecrets []orkes.ExtendedSecret) []model.Secret {
	if genSecrets == nil {
		return nil
	}
	result := make([]model.Secret, len(genSecrets))
	for i, genSecret := range genSecrets {
		result[i] = toDomainSecretFromGenerated(&genSecret)
	}
	return result
}

// toDomainSecretFromGenerated converts orkes.ExtendedSecret to model.Secret
func toDomainSecretFromGenerated(gen *orkes.ExtendedSecret) model.Secret {
	if gen == nil {
		return model.Secret{}
	}

	secret := model.Secret{}
	secret.Name = GetPointerValue(gen.Name, "")

	// Convert tags
	if gen.Tags != nil {
		secret.Tags = make([]model.Tag, len(gen.Tags))
		for i, genTag := range gen.Tags {
			secret.Tags[i] = toDomainTagFromGenerated(&genTag)
		}
	}

	return secret
}

// toDomainTagFromGenerated converts orkes.Tag to model.Tag
func toDomainTagFromGenerated(gen *orkes.Tag) model.Tag {
	if gen == nil {
		return model.Tag{}
	}

	return model.Tag{
		Key:   GetPointerValue(gen.Key, ""),
		Type:  GetPointerValue(gen.Type, ""),
		Value: GetPointerValue(gen.Value, ""),
	}
}

// ============================================================================
// Webhook Model Mappers
// ============================================================================

// toGeneratedWebhookConfig converts model.WebhookConfig to orkes.WebhookConfig
func toGeneratedWebhookConfig(domainConfig model.WebhookConfig) orkes.WebhookConfig {
	gen := orkes.WebhookConfig{}

	// Convert string fields
	if domainConfig.CreatedBy != "" {
		gen.CreatedBy = &domainConfig.CreatedBy
	}
	if domainConfig.HeaderKey != "" {
		gen.HeaderKey = &domainConfig.HeaderKey
	}
	if domainConfig.Id != "" {
		gen.Id = &domainConfig.Id
	}
	if domainConfig.Name != "" {
		gen.Name = &domainConfig.Name
	}
	if domainConfig.SecretKey != "" {
		gen.SecretKey = &domainConfig.SecretKey
	}
	if domainConfig.SecretValue != "" {
		gen.SecretValue = &domainConfig.SecretValue
	}
	if domainConfig.SourcePlatform != "" {
		gen.SourcePlatform = &domainConfig.SourcePlatform
	}
	if domainConfig.Verifier != "" {
		gen.Verifier = &domainConfig.Verifier
	}

	// Convert boolean field
	gen.UrlVerified = &domainConfig.UrlVerified

	// Convert Headers map
	if domainConfig.Headers != nil {
		gen.Headers = &domainConfig.Headers
	}

	if domainConfig.ReceiverWorkflowNamesToVersions != nil {
		gen.ReceiverWorkflowNamesToVersions = &domainConfig.ReceiverWorkflowNamesToVersions
	}

	if domainConfig.WorkflowsToStart != nil {
		gen.WorkflowsToStart = make(map[string]interface{})
		for k, v := range domainConfig.WorkflowsToStart {
			gen.WorkflowsToStart[k] = v
		}
	}

	// Convert WebhookExecutionHistory array
	if domainConfig.WebhookExecutionHistory != nil {
		gen.WebhookExecutionHistory = make([]orkes.WebhookExecutionHistory, len(domainConfig.WebhookExecutionHistory))
		for i, history := range domainConfig.WebhookExecutionHistory {
			gen.WebhookExecutionHistory[i] = toGeneratedWebhookExecutionHistory(history)
		}
	}

	return gen
}

// toGeneratedWebhookExecutionHistory converts model.WebhookExecutionHistory to orkes.WebhookExecutionHistory
func toGeneratedWebhookExecutionHistory(domain model.WebhookExecutionHistory) orkes.WebhookExecutionHistory {
	gen := orkes.WebhookExecutionHistory{}

	// Convert string fields
	if domain.EventId != "" {
		gen.EventId = &domain.EventId
	}
	if domain.Payload != "" {
		gen.Payload = &domain.Payload
	}

	// Convert boolean field
	gen.Matched = &domain.Matched

	// Convert int64 field
	gen.TimeStamp = &domain.TimeStamp

	// Convert WorkflowIds array
	if domain.WorkflowIds != nil {
		gen.WorkflowIds = make([]string, len(domain.WorkflowIds))
		copy(gen.WorkflowIds, domain.WorkflowIds)
	}

	return gen
}

// toDomainWebhookConfigFromGenerated converts orkes.WebhookConfig to model.WebhookConfig
func toDomainWebhookConfigFromGenerated(genConfig orkes.WebhookConfig) model.WebhookConfig {
	domain := model.WebhookConfig{}

	// Convert string fields
	domain.CreatedBy = GetPointerValue(genConfig.CreatedBy, "")
	domain.HeaderKey = GetPointerValue(genConfig.HeaderKey, "")
	domain.Id = GetPointerValue(genConfig.Id, "")
	domain.Name = GetPointerValue(genConfig.Name, "")
	domain.SecretKey = GetPointerValue(genConfig.SecretKey, "")
	domain.SecretValue = GetPointerValue(genConfig.SecretValue, "")
	domain.SourcePlatform = GetPointerValue(genConfig.SourcePlatform, "")
	domain.Verifier = GetPointerValue(genConfig.Verifier, "")

	// Convert boolean field
	domain.UrlVerified = GetPointerValue(genConfig.UrlVerified, false)

	// Convert Headers map
	if genConfig.Headers != nil {
		domain.Headers = *genConfig.Headers
	}

	// Convert ReceiverWorkflowNamesToVersions map
	if genConfig.ReceiverWorkflowNamesToVersions != nil {
		domain.ReceiverWorkflowNamesToVersions = *genConfig.ReceiverWorkflowNamesToVersions
	}

	if genConfig.WorkflowsToStart != nil {
		domain.WorkflowsToStart = make(map[string]int32)
		for k, v := range genConfig.WorkflowsToStart {
			switch val := v.(type) {
			case int32:
				domain.WorkflowsToStart[k] = val
			case float64:
				domain.WorkflowsToStart[k] = int32(val)
			case int:
				// Check for overflow before conversion
				if val >= math.MinInt32 && val <= math.MaxInt32 {
					domain.WorkflowsToStart[k] = int32(val)
				}
			default:
				// Skip invalid types
				continue
			}
		}
	}

	// Convert WebhookExecutionHistory array
	if genConfig.WebhookExecutionHistory != nil {
		domain.WebhookExecutionHistory = make([]model.WebhookExecutionHistory, len(genConfig.WebhookExecutionHistory))
		for i, genHistory := range genConfig.WebhookExecutionHistory {
			domain.WebhookExecutionHistory[i] = toDomainWebhookExecutionHistoryFromGenerated(&genHistory)
		}
	}

	return domain
}

// toDomainWebhookExecutionHistoryFromGenerated converts orkes.WebhookExecutionHistory to model.WebhookExecutionHistory
func toDomainWebhookExecutionHistoryFromGenerated(gen *orkes.WebhookExecutionHistory) model.WebhookExecutionHistory {
	if gen == nil {
		return model.WebhookExecutionHistory{}
	}

	domain := model.WebhookExecutionHistory{}

	// Convert string fields
	domain.EventId = GetPointerValue(gen.EventId, "")
	domain.Payload = GetPointerValue(gen.Payload, "")

	// Convert boolean field
	domain.Matched = GetPointerValue(gen.Matched, false)

	// Convert int64 field
	domain.TimeStamp = GetPointerValue(gen.TimeStamp, int64(0))

	// Convert WorkflowIds array
	if gen.WorkflowIds != nil {
		domain.WorkflowIds = make([]string, len(gen.WorkflowIds))
		copy(domain.WorkflowIds, gen.WorkflowIds)
	}

	return domain
}

// toDomainWebhookConfigsFromGenerated converts []orkes.WebhookConfig to []model.WebhookConfig
func toDomainWebhookConfigsFromGenerated(genConfigs []orkes.WebhookConfig) []model.WebhookConfig {
	if genConfigs == nil {
		return nil
	}
	result := make([]model.WebhookConfig, len(genConfigs))
	for i, genConfig := range genConfigs {
		result[i] = toDomainWebhookConfigFromGenerated(genConfig)
	}
	return result
}

// ============================================================================
// Workflow State Model Mappers
// ============================================================================

// toDomainWorkflowStateFromGenerated converts orkes.WorkflowStatus to model.WorkflowState
func toDomainWorkflowStateFromGenerated(genStatus orkes.WorkflowStatus) model.WorkflowState {
	domain := model.WorkflowState{}

	// Convert string fields
	domain.WorkflowId = GetPointerValue(genStatus.WorkflowId, "")
	domain.CorrelationId = GetPointerValue(genStatus.CorrelationId, "")
	domain.Status = GetPointerValue(genStatus.Status, "")

	if genStatus.Output != nil {
		domain.Output = genStatus.Output
	}

	if genStatus.Variables != nil {
		domain.Variables = genStatus.Variables
	}

	return domain
}

// ============================================================================
// Workflow Bulk Model Mappers
// ============================================================================

// toDomainBulkResponseFromGenerated converts orkes.BulkResponse to model.BulkResponse
func toDomainBulkResponseFromGenerated(genResponse orkes.BulkResponse) model.BulkResponse {
	result := model.BulkResponse{}

	result.BulkSuccessfulResults = genResponse.BulkSuccessfulResults
	if genResponse.BulkErrorResults != nil {
		result.BulkErrorResults = *genResponse.BulkErrorResults
	}

	return result
}

// ============================================================================
// Integration Model Mappers
// ============================================================================

// toDomainIntegration converts orkes.Integration to integration.Integration
func toDomainIntegration(genIntegration orkes.Integration) integration.Integration {

	result := integration.Integration{
		Category:    GetPointerValue(genIntegration.Category, ""),
		CreatedBy:   GetPointerValue(genIntegration.CreatedBy, ""),
		Description: GetPointerValue(genIntegration.Description, ""),
		Enabled:     GetPointerValue(genIntegration.Enabled, false),
		Name:        GetPointerValue(genIntegration.Name, ""),
		Type_:       GetPointerValue(genIntegration.Type, ""),
		UpdatedBy:   GetPointerValue(genIntegration.UpdatedBy, ""),
	}

	if genIntegration.CreateTime != nil {
		result.CreatedOn = *genIntegration.CreateTime
	}
	if genIntegration.UpdateTime != nil {
		result.UpdatedOn = *genIntegration.UpdateTime
	}
	if genIntegration.ModelsCount != nil {
		result.ModelsCount = *genIntegration.ModelsCount
	}

	// Handle configuration - flatten the nested map structure
	if genIntegration.Configuration != nil {
		result.Configuration = genIntegration.Configuration
	}

	// Handle tags conversion
	if len(genIntegration.Tags) > 0 {
		result.Tags = make([]model.TagObject, len(genIntegration.Tags))
		for i, tag := range genIntegration.Tags {
			result.Tags[i] = model.TagObject{
				Key:   GetPointerValue(tag.Key, ""),
				Value: GetPointerValue(tag.Value, ""),
				Type:  GetPointerValue(tag.Type, ""),
			}
		}
	}

	return result
}

// toDomainIntegrationsFromGenerated converts []orkes.Integration to []model.Integration
//
//nolint:gocognit,gocyclo // Generated mapper function with high complexity
func toDomainIntegrationsFromGenerated(genIntegrations []orkes.Integration) []model.Integration {
	if genIntegrations == nil {
		return nil
	}
	result := make([]model.Integration, len(genIntegrations))
	for i := range genIntegrations {
		gen := genIntegrations[i]
		m := model.Integration{}

		if gen.Category != nil {
			m.Category = *gen.Category
		}
		if gen.CreateTime != nil {
			m.CreateTime = *gen.CreateTime
		}
		if gen.CreatedBy != nil {
			m.CreatedBy = *gen.CreatedBy
		}
		if gen.Description != nil {
			m.Description = *gen.Description
		}
		if gen.Enabled != nil {
			m.Enabled = *gen.Enabled
		}
		if gen.ModelsCount != nil {
			m.ModelsCount = *gen.ModelsCount
		}
		if gen.Name != nil {
			m.Name = *gen.Name
		}
		if gen.OwnerApp != nil {
			m.OwnerApp = *gen.OwnerApp
		}
		if gen.Type != nil {
			m.Type_ = *gen.Type
		}
		if gen.UpdateTime != nil {
			m.UpdateTime = *gen.UpdateTime
		}
		if gen.UpdatedBy != nil {
			m.UpdatedBy = *gen.UpdatedBy
		}

		// Maps and slices
		if gen.Configuration != nil {
			conf := make(map[string]interface{}, len(gen.Configuration))
			for k, v := range gen.Configuration {
				conf[k] = v
			}
			m.Configuration = conf
		}
		if gen.Tags != nil {
			m.Tags = toDomainTagsFromGenerated(gen.Tags)
		}
		if gen.Apis != nil {
			m.Apis = toDomainModelIntegrationApisFromGenerated(gen.Apis)
		}

		result[i] = m
	}
	return result
}

// toDomainIntegrationApi converts orkes.IntegrationApi to integration.IntegrationApi
func toDomainIntegrationApi(genApi orkes.IntegrationApi) integration.IntegrationApi {
	res := integration.IntegrationApi{}
	if genApi.Api != nil {
		res.Api = *genApi.Api
	}
	if genApi.Configuration != nil {
		res.Configuration = genApi.Configuration
	}
	if genApi.CreatedBy != nil {
		res.CreatedBy = *genApi.CreatedBy
	}
	if genApi.CreateTime != nil {
		res.CreatedOn = *genApi.CreateTime
	}
	if genApi.Description != nil {
		res.Description = *genApi.Description
	}
	if genApi.Enabled != nil {
		res.Enabled = *genApi.Enabled
	}
	if genApi.IntegrationName != nil {
		res.IntegrationName = *genApi.IntegrationName
	}
	// Tags for this type are []model.TagObject
	if genApi.Tags != nil {
		res.Tags = make([]model.TagObject, len(genApi.Tags))
		for i := range genApi.Tags {
			res.Tags[i] = model.TagObject{
				Key:   GetPointerValue(genApi.Tags[i].Key, ""),
				Type:  GetPointerValue(genApi.Tags[i].Type, ""),
				Value: GetPointerValue(genApi.Tags[i].Value, ""),
			}
		}
	}
	if genApi.UpdatedBy != nil {
		res.UpdatedBy = *genApi.UpdatedBy
	}
	if genApi.UpdateTime != nil {
		res.UpdatedOn = *genApi.UpdateTime
	}
	return res
}

// toDomainModelIntegrationApisFromGenerated converts []orkes.IntegrationApi to []model.IntegrationApi
func toDomainModelIntegrationApisFromGenerated(genApis []orkes.IntegrationApi) []model.IntegrationApi {
	if genApis == nil {
		return nil
	}
	res := make([]model.IntegrationApi, len(genApis))
	for i := range genApis {
		res[i] = toDomainModelIntegrationApiFromGenerated(&genApis[i])
	}
	return res
}

// toDomainModelIntegrationApiFromGenerated converts orkes.IntegrationApi to model.IntegrationApi
func toDomainModelIntegrationApiFromGenerated(gen *orkes.IntegrationApi) model.IntegrationApi {
	if gen == nil {
		return model.IntegrationApi{}
	}
	m := model.IntegrationApi{}
	if gen.Api != nil {
		m.Api = *gen.Api
	}
	if gen.Configuration != nil {
		m.Configuration = gen.Configuration
	}
	if gen.CreateTime != nil {
		m.CreateTime = *gen.CreateTime
	}
	if gen.CreatedBy != nil {
		m.CreatedBy = *gen.CreatedBy
	}
	if gen.Description != nil {
		m.Description = *gen.Description
	}
	if gen.Enabled != nil {
		m.Enabled = *gen.Enabled
	}
	if gen.IntegrationName != nil {
		m.IntegrationName = *gen.IntegrationName
	}
	if gen.OwnerApp != nil {
		m.OwnerApp = *gen.OwnerApp
	}
	if gen.Tags != nil {
		m.Tags = toDomainTagsFromGenerated(gen.Tags)
	}
	if gen.UpdateTime != nil {
		m.UpdateTime = *gen.UpdateTime
	}
	if gen.UpdatedBy != nil {
		m.UpdatedBy = *gen.UpdatedBy
	}
	return m
}

// toDomainIntegrationApis converts []orkes.IntegrationApi to []integration.IntegrationApi
func toDomainIntegrationApis(genApis []orkes.IntegrationApi) []integration.IntegrationApi {
	result := make([]integration.IntegrationApi, len(genApis))
	for i, genApi := range genApis {
		result[i] = toDomainIntegrationApi(genApi)
	}
	return result
}

// toDomainIntegrationDef converts generated IntegrationDef to domain IntegrationDef
func toDomainIntegrationDef(genDef *orkes.IntegrationDef) model.IntegrationDef {
	if genDef == nil {
		return model.IntegrationDef{}
	}

	// Map simple scalar fields
	result := model.IntegrationDef{
		Name:          GetPointerValue(genDef.Name, ""),
		Category:      GetPointerValue(genDef.Category, ""),
		CategoryLabel: GetPointerValue(genDef.CategoryLabel, ""),
		Description:   GetPointerValue(genDef.Description, ""),
		Enabled:       GetPointerValue(genDef.Enabled, false),
		IconName:      GetPointerValue(genDef.IconName, ""),
		Type_:         GetPointerValue(genDef.Type, ""),
	}

	// Map tags
	if genDef.Tags != nil {
		result.Tags = make([]string, len(genDef.Tags))
		copy(result.Tags, genDef.Tags)
	}

	// Map configuration (nested form fields)
	if genDef.Configuration != nil {
		result.Configuration = make([]model.IntegrationDefFormField, len(genDef.Configuration))
		for i, ff := range genDef.Configuration {
			result.Configuration[i] = toDomainIntegrationDefFormField(ff)
		}
	}

	return result
}

// toDomainIntegrationDefFormField converts generated IntegrationDefFormField to domain IntegrationDefFormField
func toDomainIntegrationDefFormField(gen orkes.IntegrationDefFormField) model.IntegrationDefFormField {
	field := model.IntegrationDefFormField{
		DefaultValue: GetPointerValue(gen.DefaultValue, ""),
		Description:  GetPointerValue(gen.Description, ""),
		FieldName:    GetPointerValue(gen.FieldName, ""),
		FieldType:    GetPointerValue(gen.FieldType, ""),
		Label:        GetPointerValue(gen.Label, ""),
		Optional:     GetPointerValue(gen.Optional, false),
		Value:        GetPointerValue(gen.Value, ""),
	}
	if gen.ValueOptions != nil {
		field.ValueOptions = make([]model.Option, len(gen.ValueOptions))
		for i, opt := range gen.ValueOptions {
			field.ValueOptions[i] = toDomainOption(opt)
		}
	}
	return field
}

// toDomainOption converts generated Option to domain Option
func toDomainOption(gen orkes.Option) model.Option {
	return model.Option{
		Label: GetPointerValue(gen.Label, ""),
		Value: GetPointerValue(gen.Value, ""),
	}
}

// ============================================================================
// Prompt/MessageTemplate Model Mappers
// ============================================================================

// toDomainPromptTemplate converts orkes.MessageTemplate to integration.PromptTemplate
func toDomainPromptTemplate(genTemplate *orkes.MessageTemplate) integration.PromptTemplate {
	result := integration.PromptTemplate{}

	if genTemplate == nil {
		return result
	}

	// Map basic string fields
	if genTemplate.CreatedBy != nil {
		result.CreatedBy = *genTemplate.CreatedBy
	}
	if genTemplate.Description != nil {
		result.Description = *genTemplate.Description
	}
	if genTemplate.Name != nil {
		result.Name = *genTemplate.Name
	}
	if genTemplate.Template != nil {
		result.Template = *genTemplate.Template
	}
	if genTemplate.UpdatedBy != nil {
		result.UpdatedBy = *genTemplate.UpdatedBy
	}

	// Map timestamp fields
	if genTemplate.CreateTime != nil {
		result.CreatedOn = *genTemplate.CreateTime
	}
	if genTemplate.UpdateTime != nil {
		result.UpdatedOn = *genTemplate.UpdateTime
	}

	// Map string slices
	if len(genTemplate.Integrations) > 0 {
		result.Integrations = genTemplate.Integrations
	}
	if len(genTemplate.Variables) > 0 {
		result.Variables = genTemplate.Variables
	}

	// Map tags
	if len(genTemplate.Tags) > 0 {
		result.Tags = make([]model.TagObject, len(genTemplate.Tags))
		for i, tag := range genTemplate.Tags {
			result.Tags[i] = model.TagObject{
				Key:   GetPointerValue(tag.Key, ""),
				Type:  GetPointerValue(tag.Type, ""),
				Value: GetPointerValue(tag.Value, ""),
			}
		}
	}

	return result
}

// toDomainPromptTemplates converts []orkes.MessageTemplate to []integration.PromptTemplate
func toDomainPromptTemplates(genTemplates []orkes.MessageTemplate) []integration.PromptTemplate {
	result := make([]integration.PromptTemplate, len(genTemplates))
	for i, genTemplate := range genTemplates {
		result[i] = toDomainPromptTemplate(&genTemplate)
	}
	return result
}

// toGeneratedPromptTemplate converts integration.PromptTemplate to orkes.MessageTemplate
func toGeneratedPromptTemplate(domainTemplate integration.PromptTemplate) orkes.MessageTemplate {
	result := orkes.MessageTemplate{}

	// Map basic string fields
	if domainTemplate.CreatedBy != "" {
		result.CreatedBy = &domainTemplate.CreatedBy
	}
	if domainTemplate.Description != "" {
		result.Description = &domainTemplate.Description
	}
	if domainTemplate.Name != "" {
		result.Name = &domainTemplate.Name
	}
	if domainTemplate.Template != "" {
		result.Template = &domainTemplate.Template
	}
	if domainTemplate.UpdatedBy != "" {
		result.UpdatedBy = &domainTemplate.UpdatedBy
	}

	// Map timestamp fields
	if domainTemplate.CreatedOn > 0 {
		result.CreateTime = &domainTemplate.CreatedOn
	}
	if domainTemplate.UpdatedOn > 0 {
		result.UpdateTime = &domainTemplate.UpdatedOn
	}

	// Map string slices
	if len(domainTemplate.Integrations) > 0 {
		result.Integrations = domainTemplate.Integrations
	}
	if len(domainTemplate.Variables) > 0 {
		result.Variables = domainTemplate.Variables
	}

	// Map tags
	if len(domainTemplate.Tags) > 0 {
		result.Tags = make([]orkes.Tag, len(domainTemplate.Tags))
		for i, tag := range domainTemplate.Tags {
			result.Tags[i] = orkes.Tag{
				Key:   &tag.Key,
				Type:  &tag.Type,
				Value: &tag.Value,
			}
		}
	}

	return result
}

// toGeneratedPromptTemplates converts []integration.PromptTemplate to []orkes.MessageTemplate
func toGeneratedPromptTemplates(domainTemplates []integration.PromptTemplate) []orkes.MessageTemplate {
	result := make([]orkes.MessageTemplate, len(domainTemplates))
	for i, domainTemplate := range domainTemplates {
		result[i] = toGeneratedPromptTemplate(domainTemplate)
	}
	return result
}

// toGeneratedPromptTemplateTestRequest converts domain PromptTemplateTestRequest to generated PromptTemplateTestRequest
func toGeneratedPromptTemplateTestRequest(req *model.PromptTemplateTestRequest) *orkes.PromptTemplateTestRequest {
	if req == nil {
		return nil
	}
	return &orkes.PromptTemplateTestRequest{
		Prompt:          ToPointer(req.Prompt),
		LlmProvider:     ToPointer(req.LlmProvider),
		Model:           ToPointer(req.Model),
		Temperature:     ToPointer(req.Temperature),
		TopP:            ToPointer(req.TopP),
		PromptVariables: req.PromptVariables,
	}
}

// ============================================================================
// Task/Metadata Model Mappers
// ============================================================================

// toDomainTaskDef converts map[string]interface{} to model.TaskDef
//
//nolint:gocognit,gocyclo // Generated mapper function with high complexity
func toDomainTaskDef(resultMap map[string]interface{}) model.TaskDef {
	def := model.TaskDef{}

	// Basic string fields
	if name, ok := resultMap["name"].(string); ok {
		def.Name = name
	}
	if desc, ok := resultMap["description"].(string); ok {
		def.Description = desc
	}
	if ownerApp, ok := resultMap["ownerApp"].(string); ok {
		def.OwnerApp = ownerApp
	}
	if createdBy, ok := resultMap["createdBy"].(string); ok {
		def.CreatedBy = createdBy
	}
	if updatedBy, ok := resultMap["updatedBy"].(string); ok {
		def.UpdatedBy = updatedBy
	}
	if timeoutPolicy, ok := resultMap["timeoutPolicy"].(string); ok {
		def.TimeoutPolicy = timeoutPolicy
	}
	if retryLogic, ok := resultMap["retryLogic"].(string); ok {
		def.RetryLogic = retryLogic
	}
	if isolationGroupId, ok := resultMap["isolationGroupId"].(string); ok {
		def.IsolationGroupId = isolationGroupId
	}
	if executionNameSpace, ok := resultMap["executionNameSpace"].(string); ok {
		def.ExecutionNameSpace = executionNameSpace
	}
	if ownerEmail, ok := resultMap["ownerEmail"].(string); ok {
		def.OwnerEmail = ownerEmail
	}

	// Numeric fields (int64)
	if createTime, ok := resultMap["createTime"].(float64); ok {
		def.CreateTime = int64(createTime)
	}
	if updateTime, ok := resultMap["updateTime"].(float64); ok {
		def.UpdateTime = int64(updateTime)
	}
	if timeoutSeconds, ok := resultMap["timeoutSeconds"].(float64); ok {
		def.TimeoutSeconds = int64(timeoutSeconds)
	}
	if responseTimeoutSeconds, ok := resultMap["responseTimeoutSeconds"].(float64); ok {
		def.ResponseTimeoutSeconds = int64(responseTimeoutSeconds)
	}

	// Numeric fields (int32)
	if retryCount, ok := resultMap["retryCount"].(float64); ok {
		def.RetryCount = int32(retryCount)
	}
	if pollTimeoutSeconds, ok := resultMap["pollTimeoutSeconds"].(float64); ok {
		def.PollTimeoutSeconds = int32(pollTimeoutSeconds)
	}
	if retryDelaySeconds, ok := resultMap["retryDelaySeconds"].(float64); ok {
		def.RetryDelaySeconds = int32(retryDelaySeconds)
	}
	if concurrentExecLimit, ok := resultMap["concurrentExecLimit"].(float64); ok {
		def.ConcurrentExecLimit = int32(concurrentExecLimit)
	}
	if rateLimitPerFrequency, ok := resultMap["rateLimitPerFrequency"].(float64); ok {
		def.RateLimitPerFrequency = int32(rateLimitPerFrequency)
	}
	if rateLimitFrequencyInSeconds, ok := resultMap["rateLimitFrequencyInSeconds"].(float64); ok {
		def.RateLimitFrequencyInSeconds = int32(rateLimitFrequencyInSeconds)
	}
	if backoffScaleFactor, ok := resultMap["backoffScaleFactor"].(float64); ok {
		def.BackoffScaleFactor = int32(backoffScaleFactor)
	}

	// Boolean fields
	if overwriteTags, ok := resultMap["overwriteTags"].(bool); ok {
		def.OverwriteTags = overwriteTags
	}

	// Handle input/output parameters as arrays
	if inputKeys, ok := resultMap["inputKeys"].([]interface{}); ok {
		def.InputKeys = make([]string, len(inputKeys))
		for i, key := range inputKeys {
			if keyStr, ok := key.(string); ok {
				def.InputKeys[i] = keyStr
			}
		}
	}
	if outputKeys, ok := resultMap["outputKeys"].([]interface{}); ok {
		def.OutputKeys = make([]string, len(outputKeys))
		for i, key := range outputKeys {
			if keyStr, ok := key.(string); ok {
				def.OutputKeys[i] = keyStr
			}
		}
	}

	// Handle input template as map
	if inputTemplate, ok := resultMap["inputTemplate"].(map[string]interface{}); ok {
		def.InputTemplate = inputTemplate
	}

	// Handle tags array - convert to TagObject if needed
	if tags, ok := resultMap["tags"].([]interface{}); ok {
		def.Tags = make([]model.TagObject, len(tags))
		for i, tag := range tags {
			if tagMap, ok := tag.(map[string]interface{}); ok {
				jsonData, err := json.Marshal(tagMap)
				if err != nil {
					continue
				}
				if err := json.Unmarshal(jsonData, &def.Tags[i]); err != nil {
					continue
				}
			}
		}
	}

	return def
}

// toDomainTaskListSearchResultSummaryFromConductorGenerated converts *orkes.TaskListSearchResultSummary to model.TaskListSearchResultSummary
func toDomainTaskListSearchResultSummaryFromConductorGenerated(gen *orkes.TaskListSearchResultSummary) model.TaskListSearchResultSummary {
	if gen == nil {
		return model.TaskListSearchResultSummary{}
	}

	domain := model.TaskListSearchResultSummary{}

	// Map Results: []orkes.Task -> []model.Task
	if gen.Results != nil {
		domain.Results = toDomainTasksFromGenerated(gen.Results)
	}

	// Map Summary: *map[string]int64 -> map[string]int64
	if gen.Summary != nil {
		domain.Summary = *gen.Summary
	}

	// Map TotalHits: *int64 -> int64
	if gen.TotalHits != nil {
		domain.TotalHits = *gen.TotalHits
	}

	return domain
}

// toDomainTask converts *orkes.Task to model.Task
func toDomainTask(gen *orkes.Task) model.Task {
	if gen == nil {
		return model.Task{}
	}

	return model.Task{
		TaskType:                         gen.GetTaskType(),
		Status:                           model.TaskResultStatus(gen.GetStatus()),
		InputData:                        gen.GetInputData(),
		ReferenceTaskName:                gen.GetReferenceTaskName(),
		RetryCount:                       gen.GetRetryCount(),
		Seq:                              gen.GetSeq(),
		CorrelationId:                    gen.GetCorrelationId(),
		PollCount:                        gen.GetPollCount(),
		TaskDefName:                      gen.GetTaskDefName(),
		ScheduledTime:                    gen.GetScheduledTime(),
		StartTime:                        gen.GetStartTime(),
		EndTime:                          gen.GetEndTime(),
		UpdateTime:                       gen.GetUpdateTime(),
		StartDelayInSeconds:              gen.GetStartDelayInSeconds(),
		RetriedTaskId:                    gen.GetRetriedTaskId(),
		Retried:                          gen.GetRetried(),
		Executed:                         gen.GetExecuted(),
		CallbackFromWorker:               gen.GetCallbackFromWorker(),
		ResponseTimeoutSeconds:           gen.GetResponseTimeoutSeconds(),
		WorkflowInstanceId:               gen.GetWorkflowInstanceId(),
		WorkflowType:                     gen.GetWorkflowType(),
		TaskId:                           gen.GetTaskId(),
		ReasonForIncompletion:            gen.GetReasonForIncompletion(),
		CallbackAfterSeconds:             gen.GetCallbackAfterSeconds(),
		WorkerId:                         gen.GetWorkerId(),
		OutputData:                       gen.GetOutputData(),
		WorkflowTask:                     toDomainWorkflowTask(gen.GetWorkflowTask()),
		Domain:                           gen.GetDomain(),
		RateLimitPerFrequency:            gen.GetRateLimitPerFrequency(),
		RateLimitFrequencyInSeconds:      gen.GetRateLimitFrequencyInSeconds(),
		ExternalInputPayloadStoragePath:  gen.GetExternalInputPayloadStoragePath(),
		ExternalOutputPayloadStoragePath: gen.GetExternalOutputPayloadStoragePath(),
		WorkflowPriority:                 gen.GetWorkflowPriority(),
		ExecutionNameSpace:               gen.GetExecutionNameSpace(),
		IsolationGroupId:                 gen.GetIsolationGroupId(),
		Iteration:                        gen.GetIteration(),
		QueueWaitTime:                    gen.GetQueueWaitTime(),
		TaskDefinition:                   func() *model.TaskDef { t := toDomainTaskDefPtr(gen.TaskDefinition); return &t }(),
	}
}

// toDomainWorkflowTask converts orkes.WorkflowTask to *model.WorkflowTask
func toDomainWorkflowTask(task orkes.WorkflowTask) *model.WorkflowTask {
	modelTask := model.WorkflowTask{
		Name:                           task.Name,
		TaskReferenceName:              task.TaskReferenceName,
		Description:                    GetPointerValue(task.Description, ""),
		InputParameters:                task.InputParameters,
		Type_:                          GetPointerValue(task.Type, ""),
		Optional:                       GetPointerValue(task.Optional, false),
		StartDelay:                     GetPointerValue(task.StartDelay, int32(0)),
		AsyncComplete:                  GetPointerValue(task.AsyncComplete, false),
		RetryCount:                     GetPointerValue(task.RetryCount, int32(0)),
		RateLimited:                    GetPointerValue(task.RateLimited, false),
		LoopCondition:                  GetPointerValue(task.LoopCondition, ""),
		ScriptExpression:               GetPointerValue(task.ScriptExpression, ""),
		JoinOn:                         task.JoinOn,
		Sink:                           GetPointerValue(task.Sink, ""),
		Expression:                     GetPointerValue(task.Expression, ""),
		DynamicTaskNameParam:           GetPointerValue(task.DynamicTaskNameParam, ""),
		DynamicForkTasksParam:          GetPointerValue(task.DynamicForkTasksParam, ""),
		DynamicForkTasksInputParamName: GetPointerValue(task.DynamicForkTasksInputParamName, ""),
	}
	return &modelTask
}

// toDomainTaskDefPtr converts *orkes.TaskDef to *model.TaskDef
func toDomainTaskDefPtr(genDef *orkes.TaskDef) model.TaskDef {
	if genDef == nil {
		return model.TaskDef{}
	}
	return model.TaskDef{
		Name:                        genDef.Name,
		Description:                 GetPointerValue(genDef.Description, ""),
		RetryCount:                  GetPointerValue(genDef.RetryCount, 0),
		TimeoutSeconds:              genDef.TimeoutSeconds,
		InputKeys:                   genDef.InputKeys,
		OutputKeys:                  genDef.OutputKeys,
		TimeoutPolicy:               GetPointerValue(genDef.TimeoutPolicy, ""),
		RetryLogic:                  GetPointerValue(genDef.RetryLogic, ""),
		RetryDelaySeconds:           GetPointerValue(genDef.RetryDelaySeconds, 0),
		ResponseTimeoutSeconds:      GetPointerValue(genDef.ResponseTimeoutSeconds, 0),
		ConcurrentExecLimit:         GetPointerValue(genDef.ConcurrentExecLimit, 0),
		InputTemplate:               genDef.InputTemplate,
		RateLimitPerFrequency:       GetPointerValue(genDef.RateLimitPerFrequency, 0),
		RateLimitFrequencyInSeconds: GetPointerValue(genDef.RateLimitFrequencyInSeconds, 0),
		IsolationGroupId:            GetPointerValue(genDef.IsolationGroupId, ""),
		ExecutionNameSpace:          GetPointerValue(genDef.ExecutionNameSpace, ""),
		OwnerEmail:                  GetPointerValue(genDef.OwnerEmail, ""),
		PollTimeoutSeconds:          GetPointerValue(genDef.PollTimeoutSeconds, 0),
		BackoffScaleFactor:          GetPointerValue(genDef.BackoffScaleFactor, 0),
	}
}

// toDomainSearchResultTask converts *conductor.SearchResultTask to model.SearchResultTask
func toDomainSearchResultTask(gen *conductor.SearchResultTask) model.SearchResultTask {
	if gen == nil {
		return model.SearchResultTask{}
	}

	// Convert conductor tasks to domain tasks
	var tasks []model.Task
	if gen.Results != nil {
		tasks = make([]model.Task, len(gen.Results))
		for i, conductorTask := range gen.Results {
			tasks[i] = toDomainTaskFromConductor(&conductorTask)

		}
	}

	return model.SearchResultTask{
		TotalHits: GetPointerValue(gen.TotalHits, 0),
		Results:   tasks,
	}
}

// toGeneratedUpgradeWorkflowRequest converts domain UpgradeWorkflowRequest to generated model
func toGeneratedUpgradeWorkflowRequest(request *model.UpgradeWorkflowRequest) orkes.UpgradeWorkflowRequest {
	if request == nil {
		return orkes.UpgradeWorkflowRequest{}
	}

	genRequest := orkes.UpgradeWorkflowRequest{}

	genRequest.Name = request.Name

	if request.Version != 0 {
		genRequest.Version = &request.Version
	}

	if request.TaskOutput != nil {
		genRequest.TaskOutput = request.TaskOutput
	}

	if request.WorkflowInput != nil {
		genRequest.WorkflowInput = request.WorkflowInput
	}

	return genRequest
}

// toGeneratedWorkflowStateUpdate converts domain WorkflowStateUpdate to generated model
func toGeneratedWorkflowStateUpdate(update *model.WorkflowStateUpdate) orkes.WorkflowStateUpdate {
	if update == nil {
		return orkes.WorkflowStateUpdate{}
	}

	genUpdate := orkes.WorkflowStateUpdate{}

	if update.TaskReferenceName != "" {
		genUpdate.TaskReferenceName = &update.TaskReferenceName
	}

	if update.TaskResult != nil {
		genUpdate.TaskResult = toGeneratedTaskResult(update.TaskResult)
	}

	if update.Variables != nil {
		genUpdate.Variables = update.Variables
	}

	return genUpdate
}

// toGeneratedTaskResult converts domain TaskResult to generated model
func toGeneratedTaskResult(result *model.TaskResult) *orkes.TaskResult {
	if result == nil {
		return nil
	}

	return &orkes.TaskResult{
		WorkflowInstanceId:               result.WorkflowInstanceId,
		TaskId:                           result.TaskId,
		ReasonForIncompletion:            ToPointer(result.ReasonForIncompletion),
		CallbackAfterSeconds:             ToPointer(result.CallbackAfterSeconds),
		WorkerId:                         ToPointer(result.WorkerId),
		Status:                           ToPointer(string(result.Status)),
		OutputData:                       result.OutputData,
		Logs:                             toGeneratedTaskExecLogs(result.Logs),
		ExternalOutputPayloadStoragePath: ToPointer(result.ExternalOutputPayloadStoragePath),
		SubWorkflowId:                    ToPointer(result.SubWorkflowId),
	}
}

// toGeneratedTaskExecLogs converts domain TaskExecLog slice to generated slice
func toGeneratedTaskExecLogs(logs []model.TaskExecLog) []orkes.TaskExecLog {
	if logs == nil {
		return nil
	}

	genLogs := make([]orkes.TaskExecLog, len(logs))
	for i, log := range logs {
		genLogs[i] = orkes.TaskExecLog{
			Log:         ToPointer(log.Log),
			TaskId:      ToPointer(log.TaskId),
			CreatedTime: ToPointer(log.CreatedTime),
		}
	}

	return genLogs
}

// ============================================================================
// Environment Model Mappers
// ============================================================================

// toDomainEnvironmentVariables converts []orkes.EnvironmentVariable to []model.EnvironmentVariable
func toDomainEnvironmentVariables(genEnvVars []orkes.EnvironmentVariable) []model.EnvironmentVariable {
	result := make([]model.EnvironmentVariable, len(genEnvVars))
	for i, genEnvVar := range genEnvVars {
		result[i] = model.EnvironmentVariable{
			Name:  GetPointerValue(genEnvVar.Name, ""),
			Value: GetPointerValue(genEnvVar.Value, ""),
			Tags:  toDomainTags(genEnvVar.Tags),
		}
	}
	return result
}

// toDomainExternalStorageLocation converts conductor.ExternalStorageLocation to model.ExternalStorageLocation
func toDomainExternalStorageLocation(conductorLoc *conductor.ExternalStorageLocation) model.ExternalStorageLocation {
	domainLoc := model.ExternalStorageLocation{}
	if conductorLoc != nil {
		if conductorLoc.Uri != nil {
			domainLoc.Uri = *conductorLoc.Uri
		}
		if conductorLoc.Path != nil {
			domainLoc.Path = *conductorLoc.Path
		}
	}
	return domainLoc
}

// toDomainTags converts generated tags to domain tags
func toDomainTags(genTags []orkes.Tag) []model.Tag {
	if genTags == nil {
		return nil
	}
	result := make([]model.Tag, len(genTags))
	for i, genTag := range genTags {
		result[i] = model.Tag{
			Key:   GetPointerValue(genTag.Key, ""),
			Type:  GetPointerValue(genTag.Type, ""),
			Value: GetPointerValue(genTag.Value, ""),
		}
	}
	return result
}

// ============================================================================
// Event Handler Model Mappers
// ============================================================================

// toDomainEventHandlerFromOrkes converts orkes.EventHandler to model.EventHandler
//
//nolint:gocognit,gocyclo // Generated mapper function with high complexity
func toDomainEventHandlerFromOrkes(orkesHandler orkes.EventHandler) model.EventHandler {
	domainHandler := model.EventHandler{}

	// Map basic fields
	if orkesHandler.Name != nil {
		domainHandler.Name = *orkesHandler.Name
	}
	if orkesHandler.Event != nil {
		domainHandler.Event = *orkesHandler.Event
	}
	if orkesHandler.Condition != nil {
		domainHandler.Condition = *orkesHandler.Condition
	}
	if orkesHandler.Description != nil {
		domainHandler.Description = *orkesHandler.Description
	}
	if orkesHandler.EvaluatorType != nil {
		domainHandler.EvaluatorType = *orkesHandler.EvaluatorType
	}
	if orkesHandler.Active != nil {
		domainHandler.Active = *orkesHandler.Active
	}

	if len(orkesHandler.Actions) > 0 {
		domainHandler.Actions = make([]model.Action, len(orkesHandler.Actions))
		for i := range orkesHandler.Actions {
			oa := orkesHandler.Actions[i]
			act := model.Action{}
			if oa.Action != nil {
				act.Action = *oa.Action
			}
			if oa.StartWorkflow != nil {
				sw := model.StartWorkflow{}
				sw.Name = oa.StartWorkflow.Name
				if oa.StartWorkflow.Version != nil {
					sw.Version = *oa.StartWorkflow.Version
				}
				if oa.StartWorkflow.Input != nil {
					sw.Input = oa.StartWorkflow.Input
				}
				act.StartWorkflow = &sw
			}
			if oa.CompleteTask != nil {
				td := model.TaskDetails{TaskId: GetPointerValue(oa.CompleteTask.TaskId, "")}
				if oa.CompleteTask.Output != nil {
					td.Output = oa.CompleteTask.Output
				}
				td.WorkflowId = GetPointerValue(oa.CompleteTask.WorkflowId, "")
				td.TaskRefName = GetPointerValue(oa.CompleteTask.TaskRefName, "")
				act.CompleteTask = &td
			}
			if oa.FailTask != nil {
				td := model.TaskDetails{TaskId: GetPointerValue(oa.FailTask.TaskId, "")}
				if oa.FailTask.Output != nil {
					td.Output = oa.FailTask.Output
				}
				td.WorkflowId = GetPointerValue(oa.FailTask.WorkflowId, "")
				td.TaskRefName = GetPointerValue(oa.FailTask.TaskRefName, "")
				act.FailTask = &td
			}
			if oa.ExpandInlineJSON != nil {
				act.ExpandInlineJSON = *oa.ExpandInlineJSON
			}
			domainHandler.Actions[i] = act
		}
	}

	return domainHandler
}

// toDomainEventHandlersFromOrkes converts []orkes.EventHandler to []model.EventHandler
func toDomainEventHandlersFromOrkes(orkesHandlers []orkes.EventHandler) []model.EventHandler {
	result := make([]model.EventHandler, len(orkesHandlers))
	for i, orkesHandler := range orkesHandlers {
		result[i] = toDomainEventHandlerFromOrkes(orkesHandler)
	}
	return result
}

// toGeneratedEventHandlerForOrkes converts model.EventHandler to orkes.EventHandler
func toGeneratedEventHandlerForOrkes(domainHandler model.EventHandler) orkes.EventHandler {
	orkesHandler := orkes.EventHandler{
		Name:  &domainHandler.Name,
		Event: &domainHandler.Event,
	}

	// Map optional fields
	if domainHandler.Condition != "" {
		orkesHandler.Condition = &domainHandler.Condition
	}
	if domainHandler.Description != "" {
		orkesHandler.Description = &domainHandler.Description
	}
	if domainHandler.EvaluatorType != "" {
		orkesHandler.EvaluatorType = &domainHandler.EvaluatorType
	}
	// Active field - always set the pointer
	orkesHandler.Active = &domainHandler.Active

	// Map actions field-by-field
	if len(domainHandler.Actions) > 0 {
		orkesHandler.Actions = make([]orkes.Action, len(domainHandler.Actions))
		for i := range domainHandler.Actions {
			orkesHandler.Actions[i] = toGeneratedActionForOrkes(&domainHandler.Actions[i])
		}
	}

	return orkesHandler
}

// ============================================================================
// Workflow Model Mappers
// ============================================================================

// toDomainWorkflow converts *orkes.Workflow to model.Workflow
//
//nolint:gocognit,gocyclo // Generated mapper function with high complexity
func toDomainWorkflow(orkesWorkflow *orkes.Workflow) model.Workflow {
	if orkesWorkflow == nil {
		return model.Workflow{}
	}

	domainWorkflow := model.Workflow{}

	if orkesWorkflow.WorkflowId != nil {
		domainWorkflow.WorkflowId = *orkesWorkflow.WorkflowId
	}
	if orkesWorkflow.CorrelationId != nil {
		domainWorkflow.CorrelationId = *orkesWorkflow.CorrelationId
	}
	if orkesWorkflow.Status != nil {
		domainWorkflow.Status = model.WorkflowStatus(*orkesWorkflow.Status)
	}
	if orkesWorkflow.StartTime != nil {
		domainWorkflow.StartTime = *orkesWorkflow.StartTime
	}
	if orkesWorkflow.EndTime != nil {
		domainWorkflow.EndTime = *orkesWorkflow.EndTime
	}
	if orkesWorkflow.CreateTime != nil {
		domainWorkflow.CreateTime = *orkesWorkflow.CreateTime
	}
	if orkesWorkflow.UpdateTime != nil {
		domainWorkflow.UpdateTime = *orkesWorkflow.UpdateTime
	}
	if orkesWorkflow.CreatedBy != nil {
		domainWorkflow.CreatedBy = *orkesWorkflow.CreatedBy
	}
	if orkesWorkflow.Priority != nil {
		domainWorkflow.Priority = *orkesWorkflow.Priority
	}

	if orkesWorkflow.Input != nil {
		domainWorkflow.Input = orkesWorkflow.Input
	}

	if orkesWorkflow.Output != nil {
		domainWorkflow.Output = orkesWorkflow.Output
	}

	// Map variables
	if orkesWorkflow.Variables != nil {
		domainWorkflow.Variables = orkesWorkflow.Variables
	}

	// Map tasks
	if len(orkesWorkflow.Tasks) > 0 {
		domainWorkflow.Tasks = toDomainTasksFromGenerated(orkesWorkflow.Tasks)
	}

	// Map workflow definition when provided by API
	if orkesWorkflow.WorkflowDefinition != nil {
		wfDef := toDomainWorkflowDefFromGenerated(orkesWorkflow.WorkflowDefinition)
		domainWorkflow.WorkflowDefinition = &wfDef
	}

	return domainWorkflow
}

// toDomainWorkflowRun converts *orkes.WorkflowRun to model.WorkflowRun
func toDomainWorkflowRun(orkesRun *orkes.WorkflowRun) model.WorkflowRun {
	domainRun := model.WorkflowRun{}

	if orkesRun != nil {
		if orkesRun.WorkflowId != nil {
			domainRun.WorkflowId = *orkesRun.WorkflowId
		}
		if orkesRun.CorrelationId != nil {
			domainRun.CorrelationId = *orkesRun.CorrelationId
		}
		if orkesRun.Status != nil {
			domainRun.Status = *orkesRun.Status
		}
		if orkesRun.Priority != nil {
			domainRun.Priority = *orkesRun.Priority
		}
		if orkesRun.CreateTime != nil {
			domainRun.CreateTime = *orkesRun.CreateTime
		}
		if orkesRun.UpdateTime != nil {
			domainRun.UpdateTime = *orkesRun.UpdateTime
		}
		if orkesRun.CreatedBy != nil {
			domainRun.CreatedBy = *orkesRun.CreatedBy
		}

		if orkesRun.Input != nil {
			domainRun.Input = orkesRun.Input
		}

		if orkesRun.Output != nil {
			domainRun.Output = orkesRun.Output
		}

		// Map tasks
		if len(orkesRun.Tasks) > 0 {
			domainRun.Tasks = toDomainTasksFromGenerated(orkesRun.Tasks)
		}
	}

	return domainRun
}

// toDomainSearchResultWorkflowSummary converts *orkes.ScrollableSearchResultWorkflowSummary to model.SearchResultWorkflowSummary
func toDomainSearchResultWorkflowSummary(orkesResult *orkes.ScrollableSearchResultWorkflowSummary) model.SearchResultWorkflowSummary {
	domainResult := model.SearchResultWorkflowSummary{}

	if orkesResult != nil {
		if orkesResult.TotalHits != nil {
			domainResult.TotalHits = *orkesResult.TotalHits
		}

		if len(orkesResult.Results) > 0 {
			domainResult.Results = make([]model.WorkflowSummary, len(orkesResult.Results))
			for i, summary := range orkesResult.Results {
				domainResult.Results[i] = toDomainWorkflowSummaryFromGenerated(&summary)
			}
		}
	}

	return domainResult
}

// toDomainStartWorkflowRequestFromGenerated converts orkes.StartWorkflowRequest to model.StartWorkflowRequest
func toDomainStartWorkflowRequestFromGenerated(genRequest orkes.StartWorkflowRequest) model.StartWorkflowRequest {
	result := model.StartWorkflowRequest{}

	// Map basic fields
	if genRequest.Name != "" {
		result.Name = genRequest.Name
	}
	if genRequest.Version != nil {
		result.Version = int32(*genRequest.Version)
	}
	if genRequest.CorrelationId != nil {
		result.CorrelationId = *genRequest.CorrelationId
	}
	if genRequest.Priority != nil {
		result.Priority = *genRequest.Priority
	}

	// Map input
	if genRequest.Input != nil {
		result.Input = genRequest.Input
	}

	// Map task to domain mapping
	if genRequest.TaskToDomain != nil {
		result.TaskToDomain = *genRequest.TaskToDomain
	}

	return result
}

// toGeneratedStartWorkflowRequestForExecute converts model.StartWorkflowRequest to orkes.StartWorkflowRequest
func toGeneratedStartWorkflowRequestForExecute(domainRequest *model.StartWorkflowRequest) orkes.StartWorkflowRequest {
	orkesRequest := orkes.StartWorkflowRequest{}

	// Map basic fields
	if domainRequest.Name != "" {
		orkesRequest.Name = domainRequest.Name
	}
	if domainRequest.Version > 0 {
		version := int32(domainRequest.Version)
		orkesRequest.Version = &version
	}
	if domainRequest.CorrelationId != "" {
		correlationId := domainRequest.CorrelationId
		orkesRequest.CorrelationId = &correlationId
	}
	if domainRequest.Priority > 0 {
		orkesRequest.Priority = &domainRequest.Priority
	}

	// Map input
	if domainRequest.Input != nil {
		orkesRequest.Input = GetInputAsMap(domainRequest.Input)

	}

	// Map task to domain mapping
	if len(domainRequest.TaskToDomain) > 0 {
		orkesRequest.TaskToDomain = &domainRequest.TaskToDomain
	}

	// Map inline workflow definition if provided
	if domainRequest.WorkflowDef != nil {
		orkesRequest.WorkflowDef = toGeneratedWorkflowDef(domainRequest.WorkflowDef)
	}

	// Map additional fields
	if domainRequest.ExternalInputPayloadStoragePath != "" {
		orkesRequest.ExternalInputPayloadStoragePath = &domainRequest.ExternalInputPayloadStoragePath
	}
	if domainRequest.IdempotencyKey != "" {
		orkesRequest.IdempotencyKey = &domainRequest.IdempotencyKey
	}
	if domainRequest.IdempotencyStrategy != "" {
		strategy := string(domainRequest.IdempotencyStrategy)
		orkesRequest.IdempotencyStrategy = &strategy
	}
	if domainRequest.CreatedBy != "" {
		orkesRequest.CreatedBy = &domainRequest.CreatedBy
	}

	return orkesRequest
}

// toGeneratedRerunWorkflowRequest converts model.RerunWorkflowRequest to orkes.RerunWorkflowRequest
func toGeneratedRerunWorkflowRequest(domainRequest *model.RerunWorkflowRequest) orkes.RerunWorkflowRequest {
	if domainRequest == nil {
		return orkes.RerunWorkflowRequest{}
	}

	orkesRequest := orkes.RerunWorkflowRequest{}

	if domainRequest.ReRunFromWorkflowId != "" {
		orkesRequest.ReRunFromWorkflowId = &domainRequest.ReRunFromWorkflowId
	}
	if domainRequest.ReRunFromTaskId != "" {
		orkesRequest.ReRunFromTaskId = &domainRequest.ReRunFromTaskId
	}
	if len(domainRequest.TaskInput) > 0 {
		orkesRequest.TaskInput = domainRequest.TaskInput
	}

	return orkesRequest
}

// toGeneratedSkipTaskRequest converts model.SkipTaskRequest to orkes.SkipTaskRequest
func toGeneratedSkipTaskRequest(domainRequest *model.SkipTaskRequest) orkes.SkipTaskRequest {
	if domainRequest == nil {
		return orkes.SkipTaskRequest{}
	}

	orkesRequest := orkes.SkipTaskRequest{}

	if domainRequest.TaskInput != nil {
		orkesRequest.TaskInput = domainRequest.TaskInput
	}

	if domainRequest.TaskOutput != nil {
		orkesRequest.TaskOutput = domainRequest.TaskOutput
	}

	return orkesRequest
}

// toDomainWorkflowRunFromSignalResponse converts orkes.SignalResponse to model.WorkflowRun
func toDomainWorkflowRunFromSignalResponse(signalResponse *orkes.SignalResponse) model.WorkflowRun {
	if signalResponse == nil {
		return model.WorkflowRun{}
	}

	workflowRun := model.WorkflowRun{}

	// Convert string pointer fields to strings
	workflowRun.CorrelationId = GetPointerValue(signalResponse.CorrelationId, "")
	workflowRun.RequestId = GetPointerValue(signalResponse.RequestId, "")
	workflowRun.WorkflowId = GetPointerValue(signalResponse.WorkflowId, "")
	workflowRun.TargetWorkflowId = GetPointerValue(signalResponse.TargetWorkflowId, "")

	// Convert ResponseType from string pointer to string
	workflowRun.ResponseType = GetPointerValue(signalResponse.ResponseType, "")

	// Convert TargetWorkflowStatus from string pointer to string
	workflowRun.TargetWorkflowStatus = GetPointerValue(signalResponse.TargetWorkflowStatus, "")

	// Map overall status if present in extended SignalResponse
	if signalResponse.Status != "" {
		workflowRun.Status = signalResponse.Status
	}

	// Fallback to target workflow status if overall status missing
	if workflowRun.Status == "" && workflowRun.TargetWorkflowStatus != "" {
		workflowRun.Status = workflowRun.TargetWorkflowStatus
	}

	// Convert Input from nested map to flat map
	if signalResponse.Input != nil {
		workflowRun.Input = signalResponse.Input
	}

	// Convert Output from nested map to flat map
	if signalResponse.Output != nil {
		workflowRun.Output = signalResponse.Output
	}

	return workflowRun
}

// ============================================================================
// Conductor Model Mappers
// ============================================================================

// Additional helper functions for backward compatibility
func toConductorApplicationFromExtendedConductorApplication(gen *orkes.ExtendedConductorApplication) rbac.ConductorApplication {
	if gen == nil {
		return rbac.ConductorApplication{}
	}

	tags := make([]model.Tag, len(gen.Tags))
	for i, tag := range gen.Tags {
		tags[i] = toDomainTagFromGenerated(&tag)
	}

	return rbac.ConductorApplication{
		Id:         gen.GetId(),
		Name:       gen.GetName(),
		CreatedBy:  gen.GetCreatedBy(),
		UpdatedBy:  gen.GetUpdatedBy(),
		CreateTime: gen.GetCreateTime(),
		UpdateTime: gen.GetUpdateTime(),
		Tags:       tags,
	}
}

// toDomainWorkflowTasksFromGenerated converts []orkes.WorkflowTask to []model.WorkflowTask
func toDomainWorkflowTasksFromGenerated(genTasks []orkes.WorkflowTask) []model.WorkflowTask {
	if genTasks == nil {
		return nil
	}

	result := make([]model.WorkflowTask, len(genTasks))
	for i, genTask := range genTasks {
		result[i] = toDomainWorkflowTaskFromGenerated(&genTask)
	}
	return result
}

// toDomainWorkflowTaskFromGenerated converts orkes.WorkflowTask to model.WorkflowTask
//
//nolint:gocognit,gocyclo // Generated mapper function with high complexity
func toDomainWorkflowTaskFromGenerated(gen *orkes.WorkflowTask) model.WorkflowTask {
	if gen == nil {
		return model.WorkflowTask{}
	}

	domain := model.WorkflowTask{}

	// Map required fields
	domain.Name = gen.Name
	domain.TaskReferenceName = gen.TaskReferenceName

	// Map string fields (pointers to values)
	if gen.Description != nil {
		domain.Description = *gen.Description
	}
	if gen.Type != nil {
		domain.Type_ = *gen.Type
	}
	if gen.DynamicTaskNameParam != nil {
		domain.DynamicTaskNameParam = *gen.DynamicTaskNameParam
	}
	if gen.CaseValueParam != nil {
		domain.CaseValueParam = *gen.CaseValueParam
	}
	if gen.CaseExpression != nil {
		domain.CaseExpression = *gen.CaseExpression
	}
	if gen.ScriptExpression != nil {
		domain.ScriptExpression = *gen.ScriptExpression
	}
	if gen.DynamicForkJoinTasksParam != nil {
		domain.DynamicForkJoinTasksParam = *gen.DynamicForkJoinTasksParam
	}
	if gen.DynamicForkTasksParam != nil {
		domain.DynamicForkTasksParam = *gen.DynamicForkTasksParam
	}
	if gen.DynamicForkTasksInputParamName != nil {
		domain.DynamicForkTasksInputParamName = *gen.DynamicForkTasksInputParamName
	}
	if gen.Sink != nil {
		domain.Sink = *gen.Sink
	}
	if gen.LoopCondition != nil {
		domain.LoopCondition = *gen.LoopCondition
	}
	if gen.EvaluatorType != nil {
		domain.EvaluatorType = *gen.EvaluatorType
	}
	if gen.Expression != nil {
		domain.Expression = *gen.Expression
	}
	if gen.WorkflowTaskType != nil {
		domain.WorkflowTaskType = *gen.WorkflowTaskType
	}

	// Map numeric fields (pointers to values)
	if gen.StartDelay != nil {
		domain.StartDelay = *gen.StartDelay
	}
	if gen.RetryCount != nil {
		domain.RetryCount = *gen.RetryCount
	}

	// Map boolean fields (pointers to values)
	if gen.Optional != nil {
		domain.Optional = *gen.Optional
	}
	if gen.RateLimited != nil {
		domain.RateLimited = *gen.RateLimited
	}
	if gen.AsyncComplete != nil {
		domain.AsyncComplete = *gen.AsyncComplete
	}

	// Map slice fields
	if gen.JoinOn != nil {
		domain.JoinOn = make([]string, len(gen.JoinOn))
		copy(domain.JoinOn, gen.JoinOn)
	}
	if gen.DefaultExclusiveJoinTask != nil {
		domain.DefaultExclusiveJoinTask = make([]string, len(gen.DefaultExclusiveJoinTask))
		copy(domain.DefaultExclusiveJoinTask, gen.DefaultExclusiveJoinTask)
	}

	// Map map fields with nested conversion
	if gen.InputParameters != nil {
		domain.InputParameters = gen.InputParameters
	}

	// Map complex nested structures
	if gen.DecisionCases != nil {
		decisionCases := make(map[string][]model.WorkflowTask)
		for key, tasks := range *gen.DecisionCases {
			decisionCases[key] = toDomainWorkflowTasksFromGenerated(tasks)
		}
		domain.DecisionCases = decisionCases
	}

	if gen.DefaultCase != nil {
		domain.DefaultCase = toDomainWorkflowTasksFromGenerated(gen.DefaultCase)
	}

	if gen.ForkTasks != nil {
		forkTasks := make([][]model.WorkflowTask, len(gen.ForkTasks))
		for i, taskGroup := range gen.ForkTasks {
			forkTasks[i] = toDomainWorkflowTasksFromGenerated(taskGroup)
		}
		domain.ForkTasks = forkTasks
	}

	if gen.LoopOver != nil {
		domain.LoopOver = toDomainWorkflowTasksFromGenerated(gen.LoopOver)
	}

	// Map SubWorkflowParam
	if gen.SubWorkflowParam != nil {
		domain.SubWorkflowParam = toDomainSubWorkflowParamsFromGenerated(gen.SubWorkflowParam)
	}

	// Map TaskDefinition
	if gen.TaskDefinition != nil {
		domain.TaskDefinition = toDomainTaskDefFromGenerated(gen.TaskDefinition)
	}

	// Map CacheConfig
	if gen.CacheConfig != nil {
		domain.CacheConfig = toDomainCacheConfigFromGenerated(gen.CacheConfig)
	}

	return domain
}

// toDomainSubWorkflowParamsFromGenerated converts orkes.SubWorkflowParams to model.SubWorkflowParams
func toDomainSubWorkflowParamsFromGenerated(gen *orkes.SubWorkflowParams) *model.SubWorkflowParams {
	if gen == nil {
		return nil
	}

	domain := &model.SubWorkflowParams{}

	if gen.Name != nil {
		domain.Name = *gen.Name
	}

	if gen.Version != nil {
		domain.Version = *gen.Version
	}

	if gen.TaskToDomain != nil {
		domain.TaskToDomain = *gen.TaskToDomain
	}

	if gen.WorkflowDefinition != nil {
		domain.WorkflowDefinition = fromMapToWorkflowDef(gen.WorkflowDefinition)
	}

	return domain
}

// toDomainTaskDefFromGenerated converts orkes.TaskDef to model.TaskDef
//
//nolint:gocyclo // Generated mapper function with high complexity
func toDomainTaskDefFromGenerated(gen *orkes.TaskDef) *model.TaskDef {
	if gen == nil {
		return nil
	}

	domain := &model.TaskDef{}

	domain.Name = gen.Name
	domain.TimeoutSeconds = gen.TimeoutSeconds
	if gen.RetryCount != nil {
		domain.RetryCount = *gen.RetryCount
	}

	// Map string fields (pointers to values)
	if gen.OwnerApp != nil {
		domain.OwnerApp = *gen.OwnerApp
	}
	if gen.CreatedBy != nil {
		domain.CreatedBy = *gen.CreatedBy
	}
	if gen.UpdatedBy != nil {
		domain.UpdatedBy = *gen.UpdatedBy
	}
	if gen.Description != nil {
		domain.Description = *gen.Description
	}
	if gen.TimeoutPolicy != nil {
		domain.TimeoutPolicy = *gen.TimeoutPolicy
	}
	if gen.RetryLogic != nil {
		domain.RetryLogic = *gen.RetryLogic
	}
	if gen.IsolationGroupId != nil {
		domain.IsolationGroupId = *gen.IsolationGroupId
	}
	if gen.ExecutionNameSpace != nil {
		domain.ExecutionNameSpace = *gen.ExecutionNameSpace
	}
	if gen.OwnerEmail != nil {
		domain.OwnerEmail = *gen.OwnerEmail
	}

	// Map numeric fields (pointers to values)
	if gen.CreateTime != nil {
		domain.CreateTime = *gen.CreateTime
	}
	if gen.UpdateTime != nil {
		domain.UpdateTime = *gen.UpdateTime
	}
	if gen.RetryDelaySeconds != nil {
		domain.RetryDelaySeconds = *gen.RetryDelaySeconds
	}
	if gen.ResponseTimeoutSeconds != nil {
		domain.ResponseTimeoutSeconds = *gen.ResponseTimeoutSeconds
	}
	if gen.ConcurrentExecLimit != nil {
		domain.ConcurrentExecLimit = *gen.ConcurrentExecLimit
	}
	if gen.RateLimitPerFrequency != nil {
		domain.RateLimitPerFrequency = *gen.RateLimitPerFrequency
	}
	if gen.RateLimitFrequencyInSeconds != nil {
		domain.RateLimitFrequencyInSeconds = *gen.RateLimitFrequencyInSeconds
	}
	if gen.PollTimeoutSeconds != nil {
		domain.PollTimeoutSeconds = *gen.PollTimeoutSeconds
	}
	if gen.BackoffScaleFactor != nil {
		domain.BackoffScaleFactor = *gen.BackoffScaleFactor
	}

	// Map slice fields
	if gen.InputKeys != nil {
		domain.InputKeys = make([]string, len(gen.InputKeys))
		copy(domain.InputKeys, gen.InputKeys)
	}
	if gen.OutputKeys != nil {
		domain.OutputKeys = make([]string, len(gen.OutputKeys))
		copy(domain.OutputKeys, gen.OutputKeys)
	}

	// Map map fields with nested conversion
	if gen.InputTemplate != nil {
		domain.InputTemplate = gen.InputTemplate
	}

	return domain
}

// toDomainCacheConfigFromGenerated converts orkes.CacheConfig to model.CacheConfig
func toDomainCacheConfigFromGenerated(gen *orkes.CacheConfig) *model.CacheConfig {
	if gen == nil {
		return nil
	}

	domain := &model.CacheConfig{}

	if gen.Key != nil {
		domain.Key = *gen.Key
	}

	if gen.TtlInSecond != nil {
		domain.TtlInSeconds = int(*gen.TtlInSecond)
	}

	return domain
}

// toDomainWorkflowDefFromGenerated converts generated WorkflowDef to domain model
//
//nolint:gocyclo // Generated mapper function with high complexity
func toDomainWorkflowDefFromGenerated(gen *orkes.WorkflowDef) model.WorkflowDef {
	if gen == nil {
		return model.WorkflowDef{}
	}

	domain := model.WorkflowDef{}

	domain.Name = gen.Name
	domain.TimeoutSeconds = gen.TimeoutSeconds

	// Map string fields (pointers to values)
	if gen.OwnerApp != nil {
		domain.OwnerApp = *gen.OwnerApp
	}
	if gen.CreatedBy != nil {
		domain.CreatedBy = *gen.CreatedBy
	}
	if gen.UpdatedBy != nil {
		domain.UpdatedBy = *gen.UpdatedBy
	}
	if gen.Description != nil {
		domain.Description = *gen.Description
	}
	if gen.FailureWorkflow != nil {
		domain.FailureWorkflow = *gen.FailureWorkflow
	}
	if gen.OwnerEmail != nil {
		domain.OwnerEmail = *gen.OwnerEmail
	}
	if gen.TimeoutPolicy != nil {
		domain.TimeoutPolicy = *gen.TimeoutPolicy
	}

	// Map numeric fields (pointers to values)
	if gen.CreateTime != nil {
		domain.CreateTime = *gen.CreateTime
	}
	if gen.UpdateTime != nil {
		domain.UpdateTime = *gen.UpdateTime
	}
	if gen.Version != nil {
		domain.Version = *gen.Version
	}
	if gen.SchemaVersion != nil {
		domain.SchemaVersion = *gen.SchemaVersion
	}

	// Map boolean fields (pointers to values)
	if gen.Restartable != nil {
		domain.Restartable = *gen.Restartable
	}
	if gen.WorkflowStatusListenerEnabled != nil {
		domain.WorkflowStatusListenerEnabled = *gen.WorkflowStatusListenerEnabled
	}

	// Map slice fields
	if gen.InputParameters != nil {
		domain.InputParameters = make([]string, len(gen.InputParameters))
		copy(domain.InputParameters, gen.InputParameters)
	}

	// Map Tasks: []orkes.WorkflowTask -> []model.WorkflowTask
	if gen.Tasks != nil {
		domain.Tasks = toDomainWorkflowTasksFromGenerated(gen.Tasks)
	}

	// Map map fields with nested conversion
	if gen.OutputParameters != nil {
		domain.OutputParameters = gen.OutputParameters
	}
	if gen.Variables != nil {
		domain.Variables = gen.Variables
	}
	if gen.InputTemplate != nil {
		domain.InputTemplate = gen.InputTemplate
	}

	return domain
}

// toGeneratedTagFromTagString converts model.TagString to orkes.Tag
func toGeneratedTagFromTagString(domain model.TagString) orkes.Tag {
	gen := orkes.Tag{}

	if domain.Key != "" {
		gen.Key = &domain.Key
	}
	if domain.Type_ != "" {
		gen.Type = &domain.Type_
	}
	if domain.Value != "" {
		gen.Value = &domain.Value
	}

	return gen
}

// toDomainTagObjectsFromGenerated converts []orkes.Tag to []model.TagObject
func toDomainTagObjectsFromGenerated(genTags []orkes.Tag) []model.TagObject {
	if genTags == nil {
		return nil
	}

	result := make([]model.TagObject, len(genTags))
	for i, genTag := range genTags {
		result[i] = toDomainTagObjectFromGenerated(&genTag)
	}
	return result
}

// toDomainTagObjectFromGenerated converts orkes.Tag to model.TagObject
func toDomainTagObjectFromGenerated(gen *orkes.Tag) model.TagObject {
	if gen == nil {
		return model.TagObject{}
	}

	domain := model.TagObject{}

	if gen.Key != nil {
		domain.Key = *gen.Key
	}
	if gen.Type != nil {
		domain.Type = *gen.Type
	}
	if gen.Value != nil {
		domain.Value = *gen.Value
	}

	return domain
}

// ============================================================================
// SERVICE REGISTRY DTO MAPPERS WITH CONSISTENT NAMING
// ============================================================================

// toGeneratedServiceRegistry converts model.ServiceRegistry to orkes.ServiceRegistry
func toGeneratedServiceRegistry(domainRegistry *model.ServiceRegistry) *orkes.ServiceRegistry {
	if domainRegistry == nil {
		return nil
	}

	serviceRegistry := &orkes.ServiceRegistry{
		Name:       ToPointer(domainRegistry.Name),
		ServiceURI: ToPointer(domainRegistry.ServiceURI),
		Type:       ToPointer(domainRegistry.Type_),
	}

	if domainRegistry.Config != nil {
		serviceRegistry.Config = &orkes.Config{
			CircuitBreakerConfig: toGeneratedServiceRegistryCircuitBreakerConfig(domainRegistry.Config.CircuitBreakerConfig),
		}
	}
	if domainRegistry.Methods != nil {
		serviceRegistry.Methods = toGeneratedServiceRegistryMethods(domainRegistry.Methods)
	}
	if domainRegistry.RequestParams != nil {
		serviceRegistry.RequestParams = toGeneratedServiceRegistryRequestParams(domainRegistry.RequestParams)
	}

	return serviceRegistry
}

func toGeneratedServiceRegistryCircuitBreakerConfig(domainServiceRegistryCircuitBreakerConfig *model.OrkesCircuitBreakerConfig) *orkes.OrkesCircuitBreakerConfig {
	if domainServiceRegistryCircuitBreakerConfig == nil {
		return nil
	}

	return &orkes.OrkesCircuitBreakerConfig{
		AutomaticTransitionFromOpenToHalfOpenEnabled: ToPointer(domainServiceRegistryCircuitBreakerConfig.AutomaticTransitionFromOpenToHalfOpenEnabled),
		FailureRateThreshold:                         ToPointer(domainServiceRegistryCircuitBreakerConfig.FailureRateThreshold),
		MaxWaitDurationInHalfOpenState:               ToPointer(domainServiceRegistryCircuitBreakerConfig.MaxWaitDurationInHalfOpenState),
		MinimumNumberOfCalls:                         ToPointer(domainServiceRegistryCircuitBreakerConfig.MinimumNumberOfCalls),
		PermittedNumberOfCallsInHalfOpenState:        ToPointer(domainServiceRegistryCircuitBreakerConfig.PermittedNumberOfCallsInHalfOpenState),
		SlidingWindowSize:                            ToPointer(domainServiceRegistryCircuitBreakerConfig.SlidingWindowSize),
		SlowCallDurationThreshold:                    ToPointer(domainServiceRegistryCircuitBreakerConfig.SlowCallDurationThreshold),
		SlowCallRateThreshold:                        ToPointer(domainServiceRegistryCircuitBreakerConfig.SlowCallRateThreshold),
		WaitDurationInOpenState:                      ToPointer(domainServiceRegistryCircuitBreakerConfig.WaitDurationInOpenState),
	}
}

func toGeneratedServiceRegistryMethods(domainServiceRegistryMethods []model.ServiceMethod) []orkes.ServiceMethod {
	if domainServiceRegistryMethods == nil {
		return nil
	}

	methods := make([]orkes.ServiceMethod, len(domainServiceRegistryMethods))
	for i, method := range domainServiceRegistryMethods {
		methods[i] = toGeneratedServiceRegistryMethod(method)
	}
	return methods
}

func toGeneratedServiceRegistryMethod(domainServiceRegistryMethod model.ServiceMethod) orkes.ServiceMethod {
	return orkes.ServiceMethod{
		Id:            ToPointer(domainServiceRegistryMethod.Id),
		InputType:     ToPointer(domainServiceRegistryMethod.InputType),
		MethodName:    ToPointer(domainServiceRegistryMethod.MethodName),
		MethodType:    ToPointer(domainServiceRegistryMethod.MethodType),
		OperationName: ToPointer(domainServiceRegistryMethod.OperationName),
		OutputType:    ToPointer(domainServiceRegistryMethod.OutputType),
		ExampleInput:  domainServiceRegistryMethod.ExampleInput,
		RequestParams: toGeneratedServiceRegistryRequestParams(domainServiceRegistryMethod.RequestParams),
	}
}

func toGeneratedServiceRegistryRequestParam(domainServiceRegistryRequestParam model.RequestParam) orkes.RequestParam {
	return orkes.RequestParam{
		Name:     ToPointer(domainServiceRegistryRequestParam.Name),
		Type:     ToPointer(domainServiceRegistryRequestParam.Type_),
		Required: ToPointer(domainServiceRegistryRequestParam.Required),
		Schema:   toGeneratedSchema(domainServiceRegistryRequestParam.Schema),
	}
}

func toGeneratedServiceRegistryRequestParams(domainServiceRegistryRequestParams []model.RequestParam) []orkes.RequestParam {
	if domainServiceRegistryRequestParams == nil {
		return nil
	}

	requestParams := make([]orkes.RequestParam, len(domainServiceRegistryRequestParams))
	for i, requestParam := range domainServiceRegistryRequestParams {
		requestParams[i] = toGeneratedServiceRegistryRequestParam(requestParam)
	}
	return requestParams
}

func toGeneratedSchema(domainSchema *model.Schema) *orkes.Schema {
	if domainSchema == nil {
		return nil
	}
	return &orkes.Schema{
		DefaultValue: convertToMapInterface(domainSchema.DefaultValue),
		Format:       ToPointer(domainSchema.Format),
		Type:         ToPointer(domainSchema.Type_),
	}
}

// toDomainServiceRegistryFromGenerated converts orkes.ServiceRegistry to model.ServiceRegistry
func toDomainServiceRegistryFromGenerated(genRegistry *orkes.ServiceRegistry) model.ServiceRegistry {
	if genRegistry == nil {
		return model.ServiceRegistry{}
	}

	return model.ServiceRegistry{
		Name:          GetPointerValue(genRegistry.Name, ""),
		ServiceURI:    GetPointerValue(genRegistry.ServiceURI, ""),
		Type_:         GetPointerValue(genRegistry.Type, ""),
		Config:        toDomainServiceRegistryConfig(genRegistry.Config),
		Methods:       toDomainServiceRegistryMethods(genRegistry.Methods),
		RequestParams: toDomainServiceRegistryRequestParams(genRegistry.RequestParams),
	}
}

// toDomainServiceRegistriesFromGenerated converts []orkes.ServiceRegistry to []model.ServiceRegistry
func toDomainServiceRegistriesFromGenerated(genRegistries []orkes.ServiceRegistry) []model.ServiceRegistry {
	if genRegistries == nil {
		return nil
	}

	registries := make([]model.ServiceRegistry, len(genRegistries))
	for i, genRegistry := range genRegistries {
		registries[i] = toDomainServiceRegistryFromGenerated(&genRegistry)
	}
	return registries
}

// toDomainServiceRegistryConfig converts *orkes.Config to *model.Config
func toDomainServiceRegistryConfig(gen *orkes.Config) *model.Config {
	if gen == nil {
		return nil
	}
	cfg := &model.Config{}
	if !orkes.IsNil(gen.CircuitBreakerConfig) {
		cfg.CircuitBreakerConfig = toDomainOrkesCircuitBreakerConfig(gen.CircuitBreakerConfig)
	}
	return cfg
}

// toDomainOrkesCircuitBreakerConfig converts *orkes.OrkesCircuitBreakerConfig to *model.OrkesCircuitBreakerConfig
func toDomainOrkesCircuitBreakerConfig(gen *orkes.OrkesCircuitBreakerConfig) *model.OrkesCircuitBreakerConfig {
	if gen == nil {
		return nil
	}
	return &model.OrkesCircuitBreakerConfig{
		AutomaticTransitionFromOpenToHalfOpenEnabled: GetPointerValue(gen.AutomaticTransitionFromOpenToHalfOpenEnabled, false),
		FailureRateThreshold:                         GetPointerValue(gen.FailureRateThreshold, float32(0)),
		MaxWaitDurationInHalfOpenState:               GetPointerValue(gen.MaxWaitDurationInHalfOpenState, int64(0)),
		MinimumNumberOfCalls:                         GetPointerValue(gen.MinimumNumberOfCalls, int32(0)),
		PermittedNumberOfCallsInHalfOpenState:        GetPointerValue(gen.PermittedNumberOfCallsInHalfOpenState, int32(0)),
		SlidingWindowSize:                            GetPointerValue(gen.SlidingWindowSize, int32(0)),
		SlowCallDurationThreshold:                    GetPointerValue(gen.SlowCallDurationThreshold, int64(0)),
		SlowCallRateThreshold:                        GetPointerValue(gen.SlowCallRateThreshold, float32(0)),
		WaitDurationInOpenState:                      GetPointerValue(gen.WaitDurationInOpenState, int64(0)),
	}
}

// toDomainServiceRegistryMethods converts []orkes.ServiceMethod to []model.ServiceMethod
func toDomainServiceRegistryMethods(genMethods []orkes.ServiceMethod) []model.ServiceMethod {
	if genMethods == nil {
		return nil
	}
	methods := make([]model.ServiceMethod, len(genMethods))
	for i := range genMethods {
		gm := genMethods[i]
		methods[i] = model.ServiceMethod{
			ExampleInput:  gm.ExampleInput,
			Id:            GetPointerValue(gm.Id, int64(0)),
			InputType:     GetPointerValue(gm.InputType, ""),
			MethodName:    GetPointerValue(gm.MethodName, ""),
			MethodType:    GetPointerValue(gm.MethodType, ""),
			OperationName: GetPointerValue(gm.OperationName, ""),
			OutputType:    GetPointerValue(gm.OutputType, ""),
			RequestParams: toDomainRequestParamsFromGenerated(gm.RequestParams),
		}
	}
	return methods
}

// toDomainServiceRegistryRequestParams converts []orkes.RequestParam to []model.RequestParam
func toDomainServiceRegistryRequestParams(genParams []orkes.RequestParam) []model.RequestParam {
	return toDomainRequestParamsFromGenerated(genParams)
}

// GetMapValue gets a value from map with default fallback
func GetMapValue[T any](m map[string]interface{}, key string, defaultValue T) T {
	if m == nil {
		return defaultValue
	}
	if value, ok := m[key]; ok {
		if typedValue, ok := value.(T); ok {
			return typedValue
		}
	}
	return defaultValue
}

// convertToMapInterface converts *interface{} to map[string]interface{}
func convertToMapInterface(input *interface{}) map[string]interface{} {
	if input == nil {
		return nil
	}
	if mapVal, ok := (*input).(map[string]interface{}); ok {
		return mapVal
	}
	return map[string]interface{}{"value": *input}
}

// toDomainAccessKeysResponse converts map[string]interface{} to rbac.AccessKeyResponse
func toDomainAccessKeysResponse(genKey map[string]interface{}) rbac.AccessKeyResponse {
	key := rbac.AccessKeyResponse{}

	if id, ok := genKey["id"].(string); ok {
		key.Id = id
	}
	if status, ok := genKey["status"].(string); ok {
		key.Status = status
	}
	if createdAt, ok := genKey["createdAt"].(float64); ok {
		key.CreatedAt = int64(createdAt)
	} else if createdAt, ok := genKey["createdAt"].(int64); ok {
		key.CreatedAt = createdAt
	}

	return key
}

// toDomainAccessKeysResponseFromGenerated converts a generated response (map or []interface{}) to []AccessKeyResponse
//
//nolint:gocognit,gocyclo // Generated mapper function with high complexity
func toDomainAccessKeysResponseFromGenerated(gen interface{}) []rbac.AccessKeyResponse {
	if gen == nil {
		return nil
	}

	// New (fixed) case: server returns a plain array
	if arr, ok := gen.([]interface{}); ok {
		out := make([]rbac.AccessKeyResponse, 0, len(arr))
		for _, it := range arr {
			if m, ok := it.(map[string]interface{}); ok {
				out = append(out, toDomainAccessKeysResponse(m))
			}
		}
		return out
	}

	// Legacy case: generator expected map[string]interface{} with nested arrays
	if mgen, ok := gen.(map[string]interface{}); ok {
		trySlices := []string{"items", "result", "results", "accessKeys", "data"}
		for _, key := range trySlices {
			if raw, ok := mgen[key]; ok {
				if arr, ok := raw.([]interface{}); ok {
					out := make([]rbac.AccessKeyResponse, 0, len(arr))
					for _, it := range arr {
						if m, ok := it.(map[string]interface{}); ok {
							out = append(out, toDomainAccessKeysResponse(m))
						}
					}
					return out
				}
			}
		}

		// Some generators wrongly wrap arrays under numeric string keys or map["0"], etc.
		for _, v := range mgen {
			if arr, ok := v.([]interface{}); ok {
				out := make([]rbac.AccessKeyResponse, 0, len(arr))
				for _, it := range arr {
					if m, ok := it.(map[string]interface{}); ok {
						out = append(out, toDomainAccessKeysResponse(m))
					}
				}
				return out
			}
		}

		// If the response itself was a single key object
		if id, ok := mgen["id"].(string); ok && id != "" {
			return []rbac.AccessKeyResponse{toDomainAccessKeysResponse(mgen)}
		}
	}

	return nil
}

// toGeneratedHumanTaskSearch converts domain HumanTaskSearch to generated HumanTaskSearch
func toGeneratedHumanTaskSearch(domain human.HumanTaskSearch) orkes.HumanTaskSearch {
	// Convert HumanTaskUser slices
	assignees := make([]orkes.HumanTaskUser, 0, len(domain.Assignees))
	for _, assignee := range domain.Assignees {
		assignees = append(assignees, orkes.HumanTaskUser{
			User:     ToPointer(assignee.User),
			UserType: ToPointer(assignee.UserType),
		})
	}

	claimants := make([]orkes.HumanTaskUser, 0, len(domain.Claimants))
	for _, claimant := range domain.Claimants {
		claimants = append(claimants, orkes.HumanTaskUser{
			User:     ToPointer(claimant.User),
			UserType: ToPointer(claimant.UserType),
		})
	}

	return orkes.HumanTaskSearch{
		Assignees:       assignees,
		Claimants:       claimants,
		DefinitionNames: domain.DefinitionNames,
		DisplayNames:    domain.DisplayNames,
		FullTextQuery:   ToPointer(domain.FullTextQuery),
		SearchType:      ToPointer(domain.SearchType),
		Size:            ToPointer(domain.Size),
		Start:           ToPointer(domain.Start),
		States:          domain.States,
		TaskInputQuery:  ToPointer(domain.TaskInputQuery),
		TaskOutputQuery: ToPointer(domain.TaskOutputQuery),
		TaskRefNames:    domain.TaskRefNames,
		UpdateEndTime:   ToPointer(domain.UpdateEndTime),
		UpdateStartTime: ToPointer(domain.UpdateStartTime),
		WorkflowNames:   domain.WorkflowNames,
	}
}

// toDomainHumanTaskSearchResult converts generated HumanTaskSearchResult to domain HumanTaskSearchResult
func toDomainHumanTaskSearchResult(gen *orkes.HumanTaskSearchResult) human.HumanTaskSearchResult {
	// Convert HumanTaskEntry slices
	results := make([]human.HumanTaskEntry, 0, len(gen.Results))
	for _, entry := range gen.Results {
		// Convert HumanTaskUser pointers
		var assignee *human.HumanTaskUser
		if entry.Assignee != nil {
			assignee = &human.HumanTaskUser{
				User:     GetPointerValue(entry.Assignee.User, ""),
				UserType: GetPointerValue(entry.Assignee.UserType, ""),
			}
		}

		var claimant *human.HumanTaskUser
		if entry.Claimant != nil {
			claimant = &human.HumanTaskUser{
				User:     GetPointerValue(entry.Claimant.User, ""),
				UserType: GetPointerValue(entry.Claimant.UserType, ""),
			}
		}

		results = append(results, human.HumanTaskEntry{
			TaskId:         GetPointerValue(entry.TaskId, ""),
			TaskRefName:    GetPointerValue(entry.TaskRefName, ""),
			WorkflowId:     GetPointerValue(entry.WorkflowId, ""),
			DisplayName:    GetPointerValue(entry.DisplayName, ""),
			State:          GetPointerValue(entry.State, ""),
			Assignee:       assignee,
			Claimant:       claimant,
			Input:          entry.Input,
			Output:         entry.Output,
			DefinitionName: GetPointerValue(entry.DefinitionName, ""),
		})
	}

	return human.HumanTaskSearchResult{
		Hits:          GetPointerValue(gen.Hits, int32(0)),
		PageSizeLimit: GetPointerValue(gen.PageSizeLimit, int32(0)),
		Results:       results,
		Start:         GetPointerValue(gen.Start, int32(0)),
		TotalHits:     GetPointerValue(gen.TotalHits, int64(0)),
	}
}

// convertNestedMapToInterface converts map[string]map[string]interface{} to map[string]interface{}

// toDomainHumanTaskEntry converts generated HumanTaskEntry to domain HumanTaskEntry
func toDomainHumanTaskEntry(gen *orkes.HumanTaskEntry) human.HumanTaskEntry {
	// Convert HumanTaskUser pointers
	var assignee *human.HumanTaskUser
	if gen.Assignee != nil {
		assignee = &human.HumanTaskUser{
			User:     GetPointerValue(gen.Assignee.User, ""),
			UserType: GetPointerValue(gen.Assignee.UserType, ""),
		}
	}

	var claimant *human.HumanTaskUser
	if gen.Claimant != nil {
		claimant = &human.HumanTaskUser{
			User:     GetPointerValue(gen.Claimant.User, ""),
			UserType: GetPointerValue(gen.Claimant.UserType, ""),
		}
	}

	return human.HumanTaskEntry{
		TaskId:         GetPointerValue(gen.TaskId, ""),
		TaskRefName:    GetPointerValue(gen.TaskRefName, ""),
		WorkflowId:     GetPointerValue(gen.WorkflowId, ""),
		DisplayName:    GetPointerValue(gen.DisplayName, ""),
		State:          GetPointerValue(gen.State, ""),
		Assignee:       assignee,
		Claimant:       claimant,
		Input:          gen.Input,
		Output:         gen.Output,
		DefinitionName: GetPointerValue(gen.DefinitionName, ""),
	}
}

// toGeneratedHumanTaskTemplateFromSearch converts domain HumanTaskSearch to generated HumanTaskTemplate
// This is a workaround since the interface expects HumanTaskSearch but API expects HumanTaskTemplate
//
//nolint:gocyclo // Generated mapper function with high complexity
func toGeneratedHumanTaskTemplateFromSearch(domain human.HumanTaskSearch) orkes.HumanTaskTemplate {
	// Create a template with search criteria embedded in JsonSchema
	jsonSchema := make(map[string]map[string]interface{})

	// Embed search criteria in the schema
	if domain.FullTextQuery != "" {
		jsonSchema["fullTextQuery"] = map[string]interface{}{"value": domain.FullTextQuery}
	}
	if domain.SearchType != "" {
		jsonSchema["searchType"] = map[string]interface{}{"value": domain.SearchType}
	}
	if len(domain.States) > 0 {
		jsonSchema["states"] = map[string]interface{}{"value": domain.States}
	}
	if len(domain.DefinitionNames) > 0 {
		jsonSchema["definitionNames"] = map[string]interface{}{"value": domain.DefinitionNames}
	}
	if len(domain.WorkflowNames) > 0 {
		jsonSchema["workflowNames"] = map[string]interface{}{"value": domain.WorkflowNames}
	}
	if len(domain.TaskRefNames) > 0 {
		jsonSchema["taskRefNames"] = map[string]interface{}{"value": domain.TaskRefNames}
	}
	if len(domain.DisplayNames) > 0 {
		jsonSchema["displayNames"] = map[string]interface{}{"value": domain.DisplayNames}
	}
	if domain.TaskInputQuery != "" {
		jsonSchema["taskInputQuery"] = map[string]interface{}{"value": domain.TaskInputQuery}
	}
	if domain.TaskOutputQuery != "" {
		jsonSchema["taskOutputQuery"] = map[string]interface{}{"value": domain.TaskOutputQuery}
	}
	if domain.UpdateStartTime > 0 {
		jsonSchema["updateStartTime"] = map[string]interface{}{"value": domain.UpdateStartTime}
	}
	if domain.UpdateEndTime > 0 {
		jsonSchema["updateEndTime"] = map[string]interface{}{"value": domain.UpdateEndTime}
	}
	if domain.Size > 0 {
		jsonSchema["size"] = map[string]interface{}{"value": domain.Size}
	}
	if domain.Start > 0 {
		jsonSchema["start"] = map[string]interface{}{"value": domain.Start}
	}

	// Convert assignees and claimants
	if len(domain.Assignees) > 0 {
		assignees := make([]map[string]interface{}, len(domain.Assignees))
		for i, assignee := range domain.Assignees {
			assignees[i] = map[string]interface{}{
				"user":     assignee.User,
				"userType": assignee.UserType,
			}
		}
		jsonSchema["assignees"] = map[string]interface{}{"value": assignees}
	}

	if len(domain.Claimants) > 0 {
		claimants := make([]map[string]interface{}, len(domain.Claimants))
		for i, claimant := range domain.Claimants {
			claimants[i] = map[string]interface{}{
				"user":     claimant.User,
				"userType": claimant.UserType,
			}
		}
		jsonSchema["claimants"] = map[string]interface{}{"value": claimants}
	}

	// Create template UI structure
	templateUI := make(map[string]interface{})
	templateUI["searchCriteria"] = map[string]interface{}{
		"type": "search",
		"data": jsonSchema,
	}

	// Cast jsonSchema to map[string]interface{} for generated type
	flatSchema := make(map[string]interface{}, len(jsonSchema))
	for k, v := range jsonSchema {
		flatSchema[k] = v
	}

	return orkes.HumanTaskTemplate{
		Name:       "search_template", // Default name for search-based templates
		JsonSchema: flatSchema,
		TemplateUI: templateUI,
		Version:    1,
	}
}

// toDomainHumanTaskSearchFromTemplate converts generated HumanTaskTemplate to domain HumanTaskSearch
// This extracts search criteria from the template's JsonSchema
//
//nolint:gocognit,gocyclo // Generated mapper function with high complexity
func toDomainHumanTaskSearchFromTemplate(gen *orkes.HumanTaskTemplate) human.HumanTaskSearch {
	if gen == nil {
		return human.HumanTaskSearch{}
	}

	result := human.HumanTaskSearch{}

	// Extract search criteria from JsonSchema
	if gen.JsonSchema != nil {
		jsonSchema := gen.JsonSchema
		// Extract basic search fields
		if node, ok := jsonSchema["fullTextQuery"]; ok {
			if m, ok := node.(map[string]interface{}); ok {
				if v, ok := m["value"].(string); ok {
					result.FullTextQuery = v
				}
			}
		}
		if node, ok := jsonSchema["searchType"]; ok {
			if m, ok := node.(map[string]interface{}); ok {
				if v, ok := m["value"].(string); ok {
					result.SearchType = v
				}
			}
		}
		if node, ok := jsonSchema["taskInputQuery"]; ok {
			if m, ok := node.(map[string]interface{}); ok {
				if v, ok := m["value"].(string); ok {
					result.TaskInputQuery = v
				}
			}
		}
		if node, ok := jsonSchema["taskOutputQuery"]; ok {
			if m, ok := node.(map[string]interface{}); ok {
				if v, ok := m["value"].(string); ok {
					result.TaskOutputQuery = v
				}
			}
		}
		if node, ok := jsonSchema["updateStartTime"]; ok {
			if m, ok := node.(map[string]interface{}); ok {
				switch v := m["value"].(type) {
				case int64:
					result.UpdateStartTime = v
				case float64:
					result.UpdateStartTime = int64(v)
				}
			}
		}
		if node, ok := jsonSchema["updateEndTime"]; ok {
			if m, ok := node.(map[string]interface{}); ok {
				switch v := m["value"].(type) {
				case int64:
					result.UpdateEndTime = v
				case float64:
					result.UpdateEndTime = int64(v)
				}
			}
		}
		if node, ok := jsonSchema["size"]; ok {
			if m, ok := node.(map[string]interface{}); ok {
				switch v := m["value"].(type) {
				case int32:
					result.Size = v
				case float64:
					result.Size = int32(v)
				}
			}
		}
		if node, ok := jsonSchema["start"]; ok {
			if m, ok := node.(map[string]interface{}); ok {
				switch v := m["value"].(type) {
				case int32:
					result.Start = v
				case float64:
					result.Start = int32(v)
				}
			}
		}

		// Extract string arrays
		if node, ok := jsonSchema["states"]; ok {
			if arrNode, ok := node.(map[string]interface{}); ok {
				if arr, ok := arrNode["value"].([]interface{}); ok {
					result.States = make([]string, 0, len(arr))
					for _, v := range arr {
						if s, ok := v.(string); ok {
							result.States = append(result.States, s)
						}
					}
				}
			}
		}
		if node, ok := jsonSchema["definitionNames"]; ok {
			if arrNode, ok := node.(map[string]interface{}); ok {
				if arr, ok := arrNode["value"].([]interface{}); ok {
					result.DefinitionNames = make([]string, 0, len(arr))
					for _, v := range arr {
						if s, ok := v.(string); ok {
							result.DefinitionNames = append(result.DefinitionNames, s)
						}
					}
				}
			}
		}
		if node, ok := jsonSchema["workflowNames"]; ok {
			if arrNode, ok := node.(map[string]interface{}); ok {
				if arr, ok := arrNode["value"].([]interface{}); ok {
					result.WorkflowNames = make([]string, 0, len(arr))
					for _, v := range arr {
						if s, ok := v.(string); ok {
							result.WorkflowNames = append(result.WorkflowNames, s)
						}
					}
				}
			}
		}
		if node, ok := jsonSchema["taskRefNames"]; ok {
			if arrNode, ok := node.(map[string]interface{}); ok {
				if arr, ok := arrNode["value"].([]interface{}); ok {
					result.TaskRefNames = make([]string, 0, len(arr))
					for _, v := range arr {
						if s, ok := v.(string); ok {
							result.TaskRefNames = append(result.TaskRefNames, s)
						}
					}
				}
			}
		}
		if node, ok := jsonSchema["displayNames"]; ok {
			if arrNode, ok := node.(map[string]interface{}); ok {
				if arr, ok := arrNode["value"].([]interface{}); ok {
					result.DisplayNames = make([]string, 0, len(arr))
					for _, v := range arr {
						if s, ok := v.(string); ok {
							result.DisplayNames = append(result.DisplayNames, s)
						}
					}
				}
			}
		}
	}

	return result
}

// toDomainHumanTaskTemplates converts []orkes.HumanTaskTemplate to []human.HumanTaskSearch
func toDomainHumanTaskTemplates(genTemplates []orkes.HumanTaskTemplate) []human.HumanTaskSearch {
	if genTemplates == nil {
		return nil
	}

	result := make([]human.HumanTaskSearch, len(genTemplates))
	for i, template := range genTemplates {
		result[i] = toDomainHumanTaskSearchFromTemplate(&template)
	}
	return result
}

// toGeneratedIntegrationUpdate converts domain IntegrationUpdate to generated IntegrationUpdate
func toGeneratedIntegrationUpdate(update integration.IntegrationUpdate) orkes.IntegrationUpdate {
	genUpdate := orkes.IntegrationUpdate{
		Category:    ToPointer(update.Category),
		Description: ToPointer(update.Description),
		Type:        ToPointer(update.Type_),
		Enabled:     ToPointer(update.Enabled),
	}

	if update.Configuration != nil {
		// Convert ConfigKey map to simple map structure expected by generated client
		config := make(map[string]interface{})
		for k, v := range update.Configuration {
			config[string(k)] = v
		}
		genUpdate.Configuration = config
	}

	return genUpdate
}

// toGeneratedIntegrationApiUpdate converts domain IntegrationApiUpdate to generated IntegrationApiUpdate
func toGeneratedIntegrationApiUpdate(update integration.IntegrationApiUpdate) orkes.IntegrationApiUpdate {
	genUpdate := orkes.IntegrationApiUpdate{
		Enabled:     ToPointer(update.Enabled),
		Description: ToPointer(update.Description),
	}

	if update.Configuration != nil {
		genUpdate.Configuration = update.Configuration
	}

	return genUpdate
}

// toGeneratedEventLogs converts []model.EventLog to []orkes.EventLog
func toGeneratedEventLogs(domainLogs []model.EventLog) []orkes.EventLog {
	if domainLogs == nil {
		return nil
	}

	result := make([]orkes.EventLog, len(domainLogs))
	for i, domainLog := range domainLogs {
		result[i] = orkes.EventLog{
			Id:          ToPointer(domainLog.Id),
			EventType:   ToPointer(domainLog.EventType),
			Event:       ToPointer(domainLog.Event),
			CreatedAt:   ToPointer(domainLog.CreatedAt),
			HandlerName: ToPointer(domainLog.HandlerName),
			TaskId:      ToPointer(domainLog.TaskId),
			WorkerId:    ToPointer(domainLog.WorkerId),
		}
	}
	return result
}

// toExtendedWorkflowDef converts model.WorkflowDef to orkes.ExtendedWorkflowDef
func toExtendedWorkflowDef(domainDef *model.WorkflowDef) orkes.ExtendedWorkflowDef {
	if domainDef == nil {
		return orkes.ExtendedWorkflowDef{}
	}

	// Convert tasks
	var genTasks []orkes.WorkflowTask
	if domainDef.Tasks != nil {
		genTasks = make([]orkes.WorkflowTask, len(domainDef.Tasks))
		for i, task := range domainDef.Tasks {
			genTasks[i] = toGeneratedWorkflowTask(task)
		}
	}

	// Convert tags
	var genTags []orkes.Tag
	if domainDef.Tags != nil {
		genTags = make([]orkes.Tag, len(domainDef.Tags))
		for i, tag := range domainDef.Tags {
			genTags[i] = toGeneratedTagFromTagObject(tag)
		}
	}

	// Set timeout seconds with default if missing
	timeoutSeconds := domainDef.TimeoutSeconds
	if timeoutSeconds == 0 {
		// Default to 60s if not provided to satisfy server requirement
		timeoutSeconds = 60
	}

	var timeoutPolicyPtr *string
	if domainDef.TimeoutPolicy != "" {
		timeoutPolicyPtr = &domainDef.TimeoutPolicy
	}

	// Default schemaVersion to 2 if not provided
	schemaVersion := domainDef.SchemaVersion
	if schemaVersion == 0 {
		schemaVersion = 2
	}

	return orkes.ExtendedWorkflowDef{
		Name:                          domainDef.Name,
		Description:                   ToPointer(domainDef.Description),
		Version:                       ToPointer(domainDef.Version),
		Tasks:                         genTasks,
		InputParameters:               domainDef.InputParameters,
		OutputParameters:              domainDef.OutputParameters,
		FailureWorkflow:               ToPointer(domainDef.FailureWorkflow),
		SchemaVersion:                 ToPointer(int32(schemaVersion)),
		OwnerEmail:                    ToPointer(domainDef.OwnerEmail),
		TimeoutPolicy:                 timeoutPolicyPtr,
		TimeoutSeconds:                timeoutSeconds,
		Variables:                     domainDef.Variables,
		InputTemplate:                 domainDef.InputTemplate,
		Restartable:                   ToPointer(domainDef.Restartable),
		WorkflowStatusListenerEnabled: ToPointer(domainDef.WorkflowStatusListenerEnabled),
		Tags:                          genTags,
		OverwriteTags:                 ToPointer(domainDef.OverwriteTags),
	}
}

// toExtendedTaskDef converts model.TaskDef to orkes.ExtendedTaskDef
func toExtendedTaskDef(domainDef *model.TaskDef) orkes.ExtendedTaskDef {
	if domainDef == nil {
		return orkes.ExtendedTaskDef{}
	}

	// Convert tags
	var genTags []orkes.Tag
	if domainDef.Tags != nil {
		genTags = make([]orkes.Tag, len(domainDef.Tags))
		for i, tag := range domainDef.Tags {
			genTags[i] = toGeneratedTagFromTagObject(tag)
		}
	}

	// Only set RetryLogic and TimeoutPolicy if they're not empty to avoid JSON parse errors
	var retryLogicPtr *string
	if domainDef.RetryLogic != "" {
		retryLogicPtr = &domainDef.RetryLogic
	}

	var timeoutPolicyPtr *string
	if domainDef.TimeoutPolicy != "" {
		timeoutPolicyPtr = &domainDef.TimeoutPolicy
	}

	// Set default values for required fields
	responseTimeoutSeconds := domainDef.ResponseTimeoutSeconds
	if responseTimeoutSeconds == 0 {
		responseTimeoutSeconds = 1 // Minimum required by server
	}

	backoffScaleFactor := domainDef.BackoffScaleFactor
	if backoffScaleFactor == 0 {
		backoffScaleFactor = 1.0 // Default value
	}

	return orkes.ExtendedTaskDef{
		Name:                        domainDef.Name,
		Description:                 ToPointer(domainDef.Description),
		RetryCount:                  ToPointer(domainDef.RetryCount),
		TimeoutSeconds:              domainDef.TimeoutSeconds,
		TotalTimeoutSeconds:         domainDef.TimeoutSeconds, // Use same value for both
		InputKeys:                   domainDef.InputKeys,
		OutputKeys:                  domainDef.OutputKeys,
		TimeoutPolicy:               timeoutPolicyPtr,
		RetryLogic:                  retryLogicPtr,
		RetryDelaySeconds:           ToPointer(domainDef.RetryDelaySeconds),
		ResponseTimeoutSeconds:      ToPointer(responseTimeoutSeconds),
		ConcurrentExecLimit:         ToPointer(domainDef.ConcurrentExecLimit),
		InputTemplate:               domainDef.InputTemplate,
		RateLimitPerFrequency:       ToPointer(domainDef.RateLimitPerFrequency),
		RateLimitFrequencyInSeconds: ToPointer(domainDef.RateLimitFrequencyInSeconds),
		IsolationGroupId:            ToPointer(domainDef.IsolationGroupId),
		ExecutionNameSpace:          ToPointer(domainDef.ExecutionNameSpace),
		OwnerEmail:                  ToPointer(domainDef.OwnerEmail),
		PollTimeoutSeconds:          ToPointer(domainDef.PollTimeoutSeconds),
		BackoffScaleFactor:          ToPointer(backoffScaleFactor),
		Tags:                        genTags,
		OverwriteTags:               ToPointer(domainDef.OverwriteTags),
	}
}

// ============================================================================
// Metadata Tags Mappers
// ============================================================================

// toGeneratedTagsFromMetadataTags converts []model.MetadataTag to []orkes.Tag
func toGeneratedTagsFromMetadataTags(metadataTags []model.MetadataTag) []orkes.Tag {
	if metadataTags == nil {
		return nil
	}

	result := make([]orkes.Tag, len(metadataTags))
	for i, tag := range metadataTags {
		result[i] = orkes.Tag{
			Key:   ToPointer(tag.Key),
			Value: ToPointer(tag.Value),
		}
	}
	return result
}

// toDomainMetadataTagsFromGenerated converts []orkes.Tag to []model.MetadataTag
func toDomainMetadataTagsFromGenerated(genTags []orkes.Tag) []model.MetadataTag {
	if genTags == nil {
		return nil
	}

	result := make([]model.MetadataTag, len(genTags))
	for i, tag := range genTags {
		result[i] = model.MetadataTag{
			Key:   GetPointerValue(tag.Key, ""),
			Value: GetPointerValue(tag.Value, ""),
		}
	}
	return result
}

// toExtendedWorkflowDefWithTags converts model.WorkflowDef with tags to orkes.ExtendedWorkflowDef
func toExtendedWorkflowDefWithTags(domainDef *model.WorkflowDef, tags []model.MetadataTag, overwriteTags bool) orkes.ExtendedWorkflowDef {
	if domainDef == nil {
		return orkes.ExtendedWorkflowDef{}
	}

	// Convert to ExtendedWorkflowDef
	extDef := toExtendedWorkflowDef(domainDef)

	// Add tags if provided
	if len(tags) > 0 {
		extDef.Tags = toGeneratedTagsFromMetadataTags(tags)
		extDef.OverwriteTags = ToPointer(overwriteTags)
	}

	return extDef
}

// toExtendedTaskDefWithTags converts model.TaskDef with tags to orkes.ExtendedTaskDef
func toExtendedTaskDefWithTags(domainDef *model.TaskDef, tags []model.MetadataTag, overwriteTags bool) orkes.ExtendedTaskDef {
	if domainDef == nil {
		return orkes.ExtendedTaskDef{}
	}

	// Convert to ExtendedTaskDef
	extDef := toExtendedTaskDef(domainDef)

	// Add tags if provided
	if len(tags) > 0 {
		extDef.Tags = toGeneratedTagsFromMetadataTags(tags)
		extDef.OverwriteTags = ToPointer(overwriteTags)
	}

	return extDef
}

// ============================================================================
// Scheduler Model Mappers
// ============================================================================

// toDomainSaveScheduleRequestFromModel converts orkes.WorkflowScheduleModel to model.SaveScheduleRequest
func toDomainSaveScheduleRequestFromModel(genModel *orkes.WorkflowScheduleModel) model.SaveScheduleRequest {
	result := model.SaveScheduleRequest{}

	if genModel == nil {
		return result
	}

	// Map basic string fields
	if genModel.CreatedBy != nil {
		result.CreatedBy = *genModel.CreatedBy
	}
	if genModel.CronExpression != nil {
		result.CronExpression = *genModel.CronExpression
	}
	if genModel.Name != nil {
		result.Name = *genModel.Name
	}
	if genModel.UpdatedBy != nil {
		result.UpdatedBy = *genModel.UpdatedBy
	}

	// Map boolean fields
	if genModel.Paused != nil {
		result.Paused = *genModel.Paused
	}
	if genModel.RunCatchupScheduleInstances != nil {
		result.RunCatchupScheduleInstances = *genModel.RunCatchupScheduleInstances
	}

	// Map timestamp fields
	if genModel.ScheduleEndTime != nil {
		result.ScheduleEndTime = *genModel.ScheduleEndTime
	}
	if genModel.ScheduleStartTime != nil {
		result.ScheduleStartTime = *genModel.ScheduleStartTime
	}

	// Map StartWorkflowRequest
	if genModel.StartWorkflowRequest != nil {
		// Convert orkes.StartWorkflowRequest to model.StartWorkflowRequest
		domainStartReq := toDomainStartWorkflowRequestFromGenerated(*genModel.StartWorkflowRequest)
		result.StartWorkflowRequest = &domainStartReq
	}

	return result
}

// toDomainSaveScheduleRequestFromWorkflowSchedule converts orkes.WorkflowSchedule to model.SaveScheduleRequest
func toDomainSaveScheduleRequestFromWorkflowSchedule(genSchedule *orkes.WorkflowSchedule) model.SaveScheduleRequest {
	result := model.SaveScheduleRequest{}

	if genSchedule == nil {
		return result
	}

	// Map basic string fields
	if genSchedule.CreatedBy != nil {
		result.CreatedBy = *genSchedule.CreatedBy
	}
	if genSchedule.CronExpression != nil {
		result.CronExpression = *genSchedule.CronExpression
	}
	if genSchedule.Name != nil {
		result.Name = *genSchedule.Name
	}
	if genSchedule.UpdatedBy != nil {
		result.UpdatedBy = *genSchedule.UpdatedBy
	}

	// Map boolean fields
	if genSchedule.Paused != nil {
		result.Paused = *genSchedule.Paused
	}
	if genSchedule.RunCatchupScheduleInstances != nil {
		result.RunCatchupScheduleInstances = *genSchedule.RunCatchupScheduleInstances
	}

	// Map timestamp fields
	if genSchedule.ScheduleEndTime != nil {
		result.ScheduleEndTime = *genSchedule.ScheduleEndTime
	}
	if genSchedule.ScheduleStartTime != nil {
		result.ScheduleStartTime = *genSchedule.ScheduleStartTime
	}

	// Map StartWorkflowRequest
	if genSchedule.StartWorkflowRequest != nil {
		// Convert orkes.StartWorkflowRequest to model.StartWorkflowRequest
		domainStartReq := toDomainStartWorkflowRequestFromGenerated(*genSchedule.StartWorkflowRequest)
		result.StartWorkflowRequest = &domainStartReq
	}

	return result
}

// toDomainWorkflowScheduleModelsFromGenerated converts []orkes.WorkflowScheduleModel to []model.WorkflowScheduleModel
func toDomainWorkflowScheduleModelsFromGenerated(genModels []orkes.WorkflowScheduleModel) []model.WorkflowScheduleModel {
	result := make([]model.WorkflowScheduleModel, len(genModels))
	for i, genModel := range genModels {
		result[i] = toDomainWorkflowScheduleModel(&genModel)
	}
	return result
}

// toDomainWorkflowScheduleModel converts orkes.WorkflowScheduleModel to model.WorkflowScheduleModel
//
//nolint:gocyclo // Generated mapper function with high complexity
func toDomainWorkflowScheduleModel(genModel *orkes.WorkflowScheduleModel) model.WorkflowScheduleModel {
	result := model.WorkflowScheduleModel{}

	if genModel == nil {
		return result
	}

	// Map basic string fields
	if genModel.CreatedBy != nil {
		result.CreatedBy = *genModel.CreatedBy
	}
	if genModel.CronExpression != nil {
		result.CronExpression = *genModel.CronExpression
	}
	if genModel.Description != nil {
		result.Description = *genModel.Description
	}
	if genModel.Name != nil {
		result.Name = *genModel.Name
	}
	if genModel.PausedReason != nil {
		result.PausedReason = *genModel.PausedReason
	}
	if genModel.UpdatedBy != nil {
		result.UpdatedBy = *genModel.UpdatedBy
	}
	if genModel.ZoneId != nil {
		result.ZoneId = *genModel.ZoneId
	}

	// Map boolean fields
	if genModel.Paused != nil {
		result.Paused = *genModel.Paused
	}
	if genModel.RunCatchupScheduleInstances != nil {
		result.RunCatchupScheduleInstances = *genModel.RunCatchupScheduleInstances
	}

	// Map timestamp fields
	if genModel.CreateTime != nil {
		result.CreateTime = *genModel.CreateTime
	}
	if genModel.ScheduleEndTime != nil {
		result.ScheduleEndTime = *genModel.ScheduleEndTime
	}
	if genModel.ScheduleStartTime != nil {
		result.ScheduleStartTime = *genModel.ScheduleStartTime
	}
	if genModel.UpdatedTime != nil {
		result.UpdatedTime = *genModel.UpdatedTime
	}

	// Map StartWorkflowRequest
	if genModel.StartWorkflowRequest != nil {
		// Convert orkes.StartWorkflowRequest to model.StartWorkflowRequest
		domainStartReq := toDomainStartWorkflowRequestFromGenerated(*genModel.StartWorkflowRequest)
		result.StartWorkflowRequest = &domainStartReq
	}

	// Map tags
	if len(genModel.Tags) > 0 {
		result.Tags = toDomainTags(genModel.Tags)
	}

	return result
}

// toGeneratedSaveScheduleRequest converts model.SaveScheduleRequest to orkes.SaveScheduleRequest
func toGeneratedSaveScheduleRequest(domainRequest *model.SaveScheduleRequest) orkes.SaveScheduleRequest {
	result := orkes.SaveScheduleRequest{}

	if domainRequest == nil {
		return result
	}

	// Map required fields
	result.CronExpression = domainRequest.CronExpression
	result.Name = domainRequest.Name

	// Map optional string fields
	if domainRequest.CreatedBy != "" {
		result.CreatedBy = &domainRequest.CreatedBy
	}
	if domainRequest.UpdatedBy != "" {
		result.UpdatedBy = &domainRequest.UpdatedBy
	}

	// Map optional boolean fields
	if domainRequest.Paused {
		result.Paused = &domainRequest.Paused
	}
	if domainRequest.RunCatchupScheduleInstances {
		result.RunCatchupScheduleInstances = &domainRequest.RunCatchupScheduleInstances
	}

	// Map optional timestamp fields
	if domainRequest.ScheduleEndTime > 0 {
		result.ScheduleEndTime = &domainRequest.ScheduleEndTime
	}
	if domainRequest.ScheduleStartTime > 0 {
		result.ScheduleStartTime = &domainRequest.ScheduleStartTime
	}

	// Map StartWorkflowRequest
	if domainRequest.StartWorkflowRequest != nil {
		result.StartWorkflowRequest = toGeneratedStartWorkflowRequestForExecute(domainRequest.StartWorkflowRequest)
	}

	return result
}

// toDomainSearchResultWorkflowScheduleFromGenerated converts orkes.SearchResultWorkflowScheduleExecutionModel to model.SearchResultWorkflowSchedule
func toDomainSearchResultWorkflowScheduleFromGenerated(genResult *orkes.SearchResultWorkflowScheduleExecutionModel) model.SearchResultWorkflowSchedule {
	result := model.SearchResultWorkflowSchedule{}

	if genResult == nil {
		return result
	}

	// Map total hits
	if genResult.TotalHits != nil {
		result.TotalHits = *genResult.TotalHits
	}

	// Map results - convert WorkflowScheduleExecutionModel to WorkflowScheduleExecutionModel
	if len(genResult.Results) > 0 {
		result.Results = make([]model.WorkflowScheduleExecutionModel, len(genResult.Results))
		for i, execModel := range genResult.Results {
			// Convert orkes.WorkflowScheduleExecutionModel to model.WorkflowScheduleExecutionModel
			result.Results[i] = toDomainWorkflowScheduleExecutionModel(&execModel)
		}
	}

	return result
}

// toDomainWorkflowScheduleExecutionModel converts orkes.WorkflowScheduleExecutionModel to model.WorkflowScheduleExecutionModel
func toDomainWorkflowScheduleExecutionModel(execModel *orkes.WorkflowScheduleExecutionModel) model.WorkflowScheduleExecutionModel {
	result := model.WorkflowScheduleExecutionModel{}

	if execModel == nil {
		return result
	}

	// Map basic string fields
	if execModel.ExecutionId != nil {
		result.ExecutionId = *execModel.ExecutionId
	}
	if execModel.Reason != nil {
		result.Reason = *execModel.Reason
	}
	if execModel.ScheduleName != nil {
		result.ScheduleName = *execModel.ScheduleName
	}
	if execModel.StackTrace != nil {
		result.StackTrace = *execModel.StackTrace
	}
	if execModel.State != nil {
		result.State = *execModel.State
	}
	if execModel.WorkflowId != nil {
		result.WorkflowId = *execModel.WorkflowId
	}
	if execModel.WorkflowName != nil {
		result.WorkflowName = *execModel.WorkflowName
	}
	if execModel.ZoneId != nil {
		result.ZoneId = *execModel.ZoneId
	}

	// Map timestamp fields
	if execModel.ExecutionTime != nil {
		result.ExecutionTime = *execModel.ExecutionTime
	}
	if execModel.ScheduledTime != nil {
		result.ScheduledTime = *execModel.ScheduledTime
	}

	// Map StartWorkflowRequest
	if execModel.StartWorkflowRequest != nil {
		// Convert orkes.StartWorkflowRequest to model.StartWorkflowRequest
		domainStartReq := toDomainStartWorkflowRequestFromGenerated(*execModel.StartWorkflowRequest)
		result.StartWorkflowRequest = &domainStartReq
	}

	return result
}

// toDomainWorkflowScheduleModelFromExecutionModel converts orkes.WorkflowScheduleExecutionModel to model.WorkflowScheduleModel

// toDomainWorkflowScheduleModelsFromSaveRequests converts []model.SaveScheduleRequest to []model.WorkflowScheduleModel
func toDomainWorkflowScheduleModelsFromSaveRequests(saveRequests []model.SaveScheduleRequest) []model.WorkflowScheduleModel {
	result := make([]model.WorkflowScheduleModel, len(saveRequests))
	for i, saveRequest := range saveRequests {
		result[i] = toDomainWorkflowScheduleModelFromSaveRequest(&saveRequest)
	}
	return result
}

// toDomainWorkflowScheduleModelFromSaveRequest converts model.SaveScheduleRequest to model.WorkflowScheduleModel
func toDomainWorkflowScheduleModelFromSaveRequest(saveRequest *model.SaveScheduleRequest) model.WorkflowScheduleModel {
	result := model.WorkflowScheduleModel{}

	if saveRequest == nil {
		return result
	}

	// Map fields from SaveScheduleRequest
	result.CreatedBy = saveRequest.CreatedBy
	result.CronExpression = saveRequest.CronExpression
	result.Name = saveRequest.Name
	result.Paused = saveRequest.Paused
	result.RunCatchupScheduleInstances = saveRequest.RunCatchupScheduleInstances
	result.ScheduleEndTime = saveRequest.ScheduleEndTime
	result.ScheduleStartTime = saveRequest.ScheduleStartTime
	result.UpdatedBy = saveRequest.UpdatedBy

	// Map StartWorkflowRequest
	if saveRequest.StartWorkflowRequest != nil {
		result.StartWorkflowRequest = saveRequest.StartWorkflowRequest
	}

	return result
}

// toDomainWorkflowScheduleFromSaveRequest converts model.SaveScheduleRequest to model.WorkflowSchedule
func toDomainWorkflowScheduleFromSaveRequest(saveRequest *model.SaveScheduleRequest) model.WorkflowSchedule {
	result := model.WorkflowSchedule{}

	if saveRequest == nil {
		return result
	}

	// Map fields from SaveScheduleRequest
	result.CreatedBy = saveRequest.CreatedBy
	result.CronExpression = saveRequest.CronExpression
	result.Name = saveRequest.Name
	result.Paused = saveRequest.Paused
	result.RunCatchupScheduleInstances = saveRequest.RunCatchupScheduleInstances
	result.ScheduleEndTime = saveRequest.ScheduleEndTime
	result.ScheduleStartTime = saveRequest.ScheduleStartTime
	result.UpdatedBy = saveRequest.UpdatedBy

	// Map StartWorkflowRequest
	if saveRequest.StartWorkflowRequest != nil {
		result.StartWorkflowRequest = saveRequest.StartWorkflowRequest
	}

	return result
}

// ============================================================================
// Conductor Model Mappers
// ============================================================================

// toGeneratedWorkflowDefForConductor converts model.WorkflowDef to conductor.WorkflowDef
//
//nolint:gocognit,gocyclo // Generated mapper function with high complexity
func toGeneratedWorkflowDefForConductor(domainWorkflow model.WorkflowDef) conductor.WorkflowDef {
	result := conductor.WorkflowDef{}

	// Required
	result.Name = domainWorkflow.Name
	result.TimeoutSeconds = domainWorkflow.TimeoutSeconds

	// Simple pointers
	if domainWorkflow.OwnerApp != "" {
		result.OwnerApp = &domainWorkflow.OwnerApp
	}
	if domainWorkflow.CreateTime != 0 {
		result.CreateTime = &domainWorkflow.CreateTime
	}
	if domainWorkflow.UpdateTime != 0 {
		result.UpdateTime = &domainWorkflow.UpdateTime
	}
	if domainWorkflow.CreatedBy != "" {
		result.CreatedBy = &domainWorkflow.CreatedBy
	}
	if domainWorkflow.UpdatedBy != "" {
		result.UpdatedBy = &domainWorkflow.UpdatedBy
	}
	if domainWorkflow.Description != "" {
		result.Description = &domainWorkflow.Description
	}
	if domainWorkflow.Version != 0 {
		v := domainWorkflow.Version
		result.Version = &v
	}
	if domainWorkflow.FailureWorkflow != "" {
		result.FailureWorkflow = &domainWorkflow.FailureWorkflow
	}
	if domainWorkflow.SchemaVersion != 0 {
		sv := domainWorkflow.SchemaVersion
		result.SchemaVersion = &sv
	}
	result.Restartable = &domainWorkflow.Restartable
	result.WorkflowStatusListenerEnabled = &domainWorkflow.WorkflowStatusListenerEnabled
	if domainWorkflow.OwnerEmail != "" {
		result.OwnerEmail = &domainWorkflow.OwnerEmail
	}
	if domainWorkflow.TimeoutPolicy != "" {
		result.TimeoutPolicy = &domainWorkflow.TimeoutPolicy
	}

	// Slices
	if domainWorkflow.InputParameters != nil {
		result.InputParameters = domainWorkflow.InputParameters
	}

	// Maps: convert map[string]interface{} -> map[string]map[string]interface{}
	if domainWorkflow.OutputParameters != nil {
		result.OutputParameters = domainWorkflow.OutputParameters
	}
	if domainWorkflow.Variables != nil {
		result.Variables = domainWorkflow.Variables
	}
	if domainWorkflow.InputTemplate != nil {
		result.InputTemplate = domainWorkflow.InputTemplate
	}

	// Tasks
	if domainWorkflow.Tasks != nil {
		result.Tasks = toGeneratedWorkflowTasksForConductor(domainWorkflow.Tasks)
	}

	return result
}

// toGeneratedWorkflowTasksForConductor converts []model.WorkflowTask to []conductor.WorkflowTask
//
//nolint:gocognit,gocyclo // Generated mapper function with high complexity
func toGeneratedWorkflowTasksForConductor(tasks []model.WorkflowTask) []conductor.WorkflowTask {
	if tasks == nil {
		return nil
	}
	res := make([]conductor.WorkflowTask, len(tasks))
	for i := range tasks {
		t := tasks[i]
		wt := conductor.WorkflowTask{
			Name:              t.Name,
			TaskReferenceName: t.TaskReferenceName,
		}
		if t.Description != "" {
			wt.Description = &t.Description
		}
		if t.InputParameters != nil {
			wt.InputParameters = t.InputParameters
		}
		if t.Type_ != "" {
			wt.Type = &t.Type_
		}
		if t.DynamicTaskNameParam != "" {
			wt.DynamicTaskNameParam = &t.DynamicTaskNameParam
		}
		if t.CaseValueParam != "" {
			wt.CaseValueParam = &t.CaseValueParam
		}
		if t.CaseExpression != "" {
			wt.CaseExpression = &t.CaseExpression
		}
		if t.ScriptExpression != "" {
			wt.ScriptExpression = &t.ScriptExpression
		}
		if t.DecisionCases != nil {
			dc := make(map[string][]conductor.WorkflowTask, len(t.DecisionCases))
			for k, arr := range t.DecisionCases {
				dcarr := make([]conductor.WorkflowTask, len(arr))
				for j := range arr {
					dcarr[j] = toGeneratedWorkflowTasksForConductor([]model.WorkflowTask{arr[j]})[0]
				}
				dc[k] = dcarr
			}
			wt.DecisionCases = &dc
		}
		if t.DynamicForkJoinTasksParam != "" {
			wt.DynamicForkJoinTasksParam = &t.DynamicForkJoinTasksParam
		}
		if t.DynamicForkTasksParam != "" {
			wt.DynamicForkTasksParam = &t.DynamicForkTasksParam
		}
		if t.DynamicForkTasksInputParamName != "" {
			wt.DynamicForkTasksInputParamName = &t.DynamicForkTasksInputParamName
		}
		if t.DefaultCase != nil {
			dc := make([]conductor.WorkflowTask, len(t.DefaultCase))
			for j := range t.DefaultCase {
				dc[j] = toGeneratedWorkflowTasksForConductor([]model.WorkflowTask{t.DefaultCase[j]})[0]
			}
			wt.DefaultCase = dc
		}
		if t.ForkTasks != nil {
			forks := make([][]conductor.WorkflowTask, len(t.ForkTasks))
			for j := range t.ForkTasks {
				forks[j] = toGeneratedWorkflowTasksForConductor(t.ForkTasks[j])
			}
			wt.ForkTasks = forks
		}
		if t.StartDelay != 0 {
			wt.StartDelay = &t.StartDelay
		}
		if t.SubWorkflowParam != nil {
			sp := conductor.SubWorkflowParams{}
			if t.SubWorkflowParam.Name != "" {
				sp.Name = &t.SubWorkflowParam.Name
			}
			if t.SubWorkflowParam.Version != 0 {
				v := int32(t.SubWorkflowParam.Version)
				sp.Version = &v
			}
			if t.SubWorkflowParam.TaskToDomain != nil {
				sp.TaskToDomain = &t.SubWorkflowParam.TaskToDomain
			}
			wt.SubWorkflowParam = &sp
		}
		if t.JoinOn != nil {
			wt.JoinOn = t.JoinOn
		}
		if t.Sink != "" {
			wt.Sink = &t.Sink
		}
		if t.Optional {
			wt.Optional = &t.Optional
		}
		if t.TaskDefinition != nil {
			td := conductor.TaskDef{Name: t.TaskDefinition.Name}
			if t.TaskDefinition.Description != "" {
				td.Description = &t.TaskDefinition.Description
			}
			if t.TaskDefinition.RetryCount != 0 {
				td.RetryCount = &t.TaskDefinition.RetryCount
			}
			td.TimeoutSeconds = t.TaskDefinition.TimeoutSeconds
			td.InputKeys = t.TaskDefinition.InputKeys
			td.OutputKeys = t.TaskDefinition.OutputKeys
			if t.TaskDefinition.TimeoutPolicy != "" {
				td.TimeoutPolicy = &t.TaskDefinition.TimeoutPolicy
			}
			if t.TaskDefinition.RetryLogic != "" {
				td.RetryLogic = &t.TaskDefinition.RetryLogic
			}
			if t.TaskDefinition.RetryDelaySeconds != 0 {
				td.RetryDelaySeconds = &t.TaskDefinition.RetryDelaySeconds
			}
			if t.TaskDefinition.ResponseTimeoutSeconds != 0 {
				td.ResponseTimeoutSeconds = &t.TaskDefinition.ResponseTimeoutSeconds
			}
			if t.TaskDefinition.ConcurrentExecLimit != 0 {
				td.ConcurrentExecLimit = &t.TaskDefinition.ConcurrentExecLimit
			}
			if t.TaskDefinition.InputTemplate != nil {
				td.InputTemplate = t.TaskDefinition.InputTemplate
			}
			if t.TaskDefinition.RateLimitPerFrequency != 0 {
				td.RateLimitPerFrequency = &t.TaskDefinition.RateLimitPerFrequency
			}
			if t.TaskDefinition.RateLimitFrequencyInSeconds != 0 {
				td.RateLimitFrequencyInSeconds = &t.TaskDefinition.RateLimitFrequencyInSeconds
			}
			if t.TaskDefinition.IsolationGroupId != "" {
				td.IsolationGroupId = &t.TaskDefinition.IsolationGroupId
			}
			if t.TaskDefinition.ExecutionNameSpace != "" {
				td.ExecutionNameSpace = &t.TaskDefinition.ExecutionNameSpace
			}
			if t.TaskDefinition.OwnerEmail != "" {
				td.OwnerEmail = &t.TaskDefinition.OwnerEmail
			}
			if t.TaskDefinition.PollTimeoutSeconds != 0 {
				td.PollTimeoutSeconds = &t.TaskDefinition.PollTimeoutSeconds
			}
			if t.TaskDefinition.BackoffScaleFactor != 0 {
				td.BackoffScaleFactor = &t.TaskDefinition.BackoffScaleFactor
			}
			wt.TaskDefinition = &td
		}
		if t.RateLimited {
			wt.RateLimited = &t.RateLimited
		}
		if t.DefaultExclusiveJoinTask != nil {
			wt.DefaultExclusiveJoinTask = t.DefaultExclusiveJoinTask
		}
		if t.AsyncComplete {
			wt.AsyncComplete = &t.AsyncComplete
		}
		if t.LoopCondition != "" {
			wt.LoopCondition = &t.LoopCondition
		}
		if t.LoopOver != nil {
			wt.LoopOver = toGeneratedWorkflowTasksForConductor(t.LoopOver)
		}
		if t.RetryCount != 0 {
			wt.RetryCount = &t.RetryCount
		}
		if t.EvaluatorType != "" {
			wt.EvaluatorType = &t.EvaluatorType
		}
		if t.Expression != "" {
			wt.Expression = &t.Expression
		}
		if t.WorkflowTaskType != "" {
			wt.WorkflowTaskType = &t.WorkflowTaskType
		}
		res[i] = wt
	}
	return res
}

// toGeneratedWorkflowDef converts domain WorkflowDef to generated model
//
//nolint:gocyclo // Generated mapper function with high complexity
func toGeneratedWorkflowDef(domain *model.WorkflowDef) *orkes.WorkflowDef {
	if domain == nil {
		return nil
	}

	// Comprehensive field-by-field mapping from model.WorkflowDef to orkes.WorkflowDef
	gen := &orkes.WorkflowDef{}

	// Required fields
	if domain.Name != "" {
		gen.Name = domain.Name
	}
	if domain.TimeoutSeconds != 0 {
		gen.TimeoutSeconds = domain.TimeoutSeconds
	}

	// Convert Tasks: []model.WorkflowTask -> []orkes.WorkflowTask
	if domain.Tasks != nil {
		gen.Tasks = toGeneratedWorkflowTasks(domain.Tasks)
	}

	// String fields to pointers
	if domain.OwnerApp != "" {
		gen.OwnerApp = &domain.OwnerApp
	}
	if domain.CreatedBy != "" {
		gen.CreatedBy = &domain.CreatedBy
	}
	if domain.UpdatedBy != "" {
		gen.UpdatedBy = &domain.UpdatedBy
	}
	if domain.Description != "" {
		gen.Description = &domain.Description
	}
	if domain.FailureWorkflow != "" {
		gen.FailureWorkflow = &domain.FailureWorkflow
	}
	if domain.OwnerEmail != "" {
		gen.OwnerEmail = &domain.OwnerEmail
	}
	if domain.TimeoutPolicy != "" {
		gen.TimeoutPolicy = &domain.TimeoutPolicy
	}

	// Numeric fields to pointers
	if domain.CreateTime != 0 {
		gen.CreateTime = &domain.CreateTime
	}
	if domain.UpdateTime != 0 {
		gen.UpdateTime = &domain.UpdateTime
	}
	if domain.Version != 0 {
		gen.Version = &domain.Version
	}
	if domain.SchemaVersion != 0 {
		gen.SchemaVersion = &domain.SchemaVersion
	}

	// Boolean fields to pointers
	gen.Restartable = &domain.Restartable
	gen.WorkflowStatusListenerEnabled = &domain.WorkflowStatusListenerEnabled

	// Slice fields (direct assignment)
	if domain.InputParameters != nil {
		gen.InputParameters = domain.InputParameters
	}

	// Map fields with nested conversion
	if domain.OutputParameters != nil {
		gen.OutputParameters = domain.OutputParameters
	}
	if domain.Variables != nil {
		gen.Variables = domain.Variables
	}
	if domain.InputTemplate != nil {
		gen.InputTemplate = domain.InputTemplate
	}

	return gen
}

// toGeneratedWorkflowTasks converts []model.WorkflowTask to []orkes.WorkflowTask
func toGeneratedWorkflowTasks(tasks []model.WorkflowTask) []orkes.WorkflowTask {
	if tasks == nil {
		return nil
	}

	result := make([]orkes.WorkflowTask, len(tasks))
	for i, task := range tasks {
		result[i] = toGeneratedWorkflowTask(task)
	}
	return result
}

// toDomainCircuitBreakerTransitionResponseFromGenerated converts orkes.CircuitBreakerTransitionResponse to model.CircuitBreakerTransitionResponse
func toDomainCircuitBreakerTransitionResponseFromGenerated(gen *orkes.CircuitBreakerTransitionResponse) model.CircuitBreakerTransitionResponse {
	if gen == nil {
		return model.CircuitBreakerTransitionResponse{}
	}

	return model.CircuitBreakerTransitionResponse{
		CurrentState:        GetPointerValue(gen.CurrentState, ""),
		Message:             GetPointerValue(gen.Message, ""),
		PreviousState:       GetPointerValue(gen.PreviousState, ""),
		Service:             GetPointerValue(gen.Service, ""),
		TransitionTimestamp: GetPointerValue(gen.TransitionTimestamp, int64(0)),
	}
}

// toDomainServiceMethodsFromGenerated converts []orkes.ServiceMethod to []model.ServiceMethod
func toDomainServiceMethodsFromGenerated(genMethods []orkes.ServiceMethod) []model.ServiceMethod {
	if genMethods == nil {
		return nil
	}

	methods := make([]model.ServiceMethod, len(genMethods))
	for i := range genMethods {
		gm := genMethods[i]
		methods[i] = model.ServiceMethod{
			ExampleInput:  gm.ExampleInput,
			Id:            GetPointerValue(gm.Id, 0),
			InputType:     GetPointerValue(gm.InputType, ""),
			MethodName:    GetPointerValue(gm.MethodName, ""),
			MethodType:    GetPointerValue(gm.MethodType, ""),
			OperationName: GetPointerValue(gm.OperationName, ""),
			OutputType:    GetPointerValue(gm.OutputType, ""),
			RequestParams: toDomainRequestParamsFromGenerated(gm.RequestParams),
		}
	}
	return methods
}

// toDomainProtoRegistryEntriesFromGenerated converts []orkes.ProtoRegistryEntry to []model.ProtoRegistryEntry
func toDomainProtoRegistryEntriesFromGenerated(genEntries []orkes.ProtoRegistryEntry) []model.ProtoRegistryEntry {
	if genEntries == nil {
		return nil
	}

	entries := make([]model.ProtoRegistryEntry, len(genEntries))
	for i := range genEntries {
		ge := genEntries[i]
		entries[i] = model.ProtoRegistryEntry{
			Data:        GetPointerValue(ge.Data, ""),
			Filename:    GetPointerValue(ge.Filename, ""),
			ServiceName: GetPointerValue(ge.ServiceName, ""),
		}
	}
	return entries
}

// toDomainRequestParamsFromGenerated converts []orkes.RequestParam to []model.RequestParam
func toDomainRequestParamsFromGenerated(genParams []orkes.RequestParam) []model.RequestParam {
	if genParams == nil {
		return nil
	}
	params := make([]model.RequestParam, len(genParams))
	for i := range genParams {
		gp := genParams[i]
		params[i] = model.RequestParam{
			Name:     GetPointerValue(gp.Name, ""),
			Required: GetPointerValue(gp.Required, false),
			Type_:    GetPointerValue(gp.Type, ""),
			Schema:   toDomainSchemaFromGenerated(gp.Schema),
		}
	}
	return params
}

// toDomainSchemaFromGenerated converts *orkes.Schema to *model.Schema
func toDomainSchemaFromGenerated(gen *orkes.Schema) *model.Schema {
	if gen == nil {
		return nil
	}
	return &model.Schema{
		DefaultValue: convertFromMapInterface(gen.DefaultValue),
		Format:       GetPointerValue(gen.Format, ""),
		Type_:        GetPointerValue(gen.Type, ""),
	}
}

// toDomainSearchResultWorkflowSummaryFromScrollableGenerated converts orkes.ScrollableSearchResultWorkflowSummary to model.SearchResultWorkflowSummary

// toDomainSearchResultWorkflowSummaryFromGenerated converts orkes.ScrollableSearchResultWorkflowSummary to model.SearchResultWorkflowSummary
func toDomainSearchResultWorkflowSummaryFromConductorGenerated(gen *conductor.SearchResultWorkflowSummary) model.SearchResultWorkflowSummary {
	if gen == nil {
		return model.SearchResultWorkflowSummary{}
	}
	result := model.SearchResultWorkflowSummary{
		TotalHits: GetPointerValue(gen.TotalHits, int64(0)),
		Results:   nil,
	}
	if !orkes.IsNil(gen.Results) {
		result.Results = make([]model.WorkflowSummary, len(gen.Results))
		for i := range gen.Results {
			ws := gen.Results[i]
			result.Results[i] = toDomainWorkflowSummaryFromConductorGenerated(&ws)
		}
	}
	return result
}

// toDomainSearchResultWorkflowSummaryFromGenerated converts orkes.ScrollableSearchResultWorkflowSummary to model.SearchResultWorkflowSummary
func toDomainSearchResultWorkflowFromConductorGenerated(gen *conductor.SearchResultWorkflow) model.SearchResultWorkflow {
	if gen == nil {
		return model.SearchResultWorkflow{}
	}
	result := model.SearchResultWorkflow{
		TotalHits: GetPointerValue(gen.TotalHits, int64(0)),
		Results:   nil,
	}
	if !orkes.IsNil(gen.Results) {
		result.Results = make([]model.Workflow, len(gen.Results))
		for i := range gen.Results {
			ws := gen.Results[i]
			result.Results[i] = toDomainWorkflowFromConductorGenerated(&ws)
		}
	}
	return result
}

func toDomainWorkflowFromConductorGenerated(gen *conductor.Workflow) model.Workflow {
	if gen == nil {
		return model.Workflow{}
	}
	return model.Workflow{
		WorkflowId:                       GetPointerValue(gen.WorkflowId, ""),
		CorrelationId:                    GetPointerValue(gen.CorrelationId, ""),
		Status:                           model.WorkflowStatus(GetPointerValue(gen.Status, "")),
		StartTime:                        GetPointerValue(gen.StartTime, int64(0)),
		EndTime:                          GetPointerValue(gen.EndTime, int64(0)),
		Priority:                         GetPointerValue(gen.Priority, int32(0)),
		Input:                            gen.Input,
		Output:                           gen.Output,
		ReasonForIncompletion:            GetPointerValue(gen.ReasonForIncompletion, ""),
		Event:                            GetPointerValue(gen.Event, ""),
		FailedReferenceTaskNames:         gen.FailedReferenceTaskNames,
		ExternalInputPayloadStoragePath:  GetPointerValue(gen.ExternalInputPayloadStoragePath, ""),
		ExternalOutputPayloadStoragePath: GetPointerValue(gen.ExternalOutputPayloadStoragePath, ""),
		// Additional fields from conductor.Workflow
		CreateTime:           GetPointerValue(gen.CreateTime, int64(0)),
		UpdateTime:           GetPointerValue(gen.UpdateTime, int64(0)),
		CreatedBy:            GetPointerValue(gen.CreatedBy, ""),
		UpdatedBy:            GetPointerValue(gen.UpdatedBy, ""),
		OwnerApp:             GetPointerValue(gen.OwnerApp, ""),
		ParentWorkflowId:     GetPointerValue(gen.ParentWorkflowId, ""),
		ParentWorkflowTaskId: GetPointerValue(gen.ParentWorkflowTaskId, ""),
		ReRunFromWorkflowId:  GetPointerValue(gen.ReRunFromWorkflowId, ""),
		IdempotencyKey:       GetPointerValue(gen.IdempotencyKey, ""),
		RateLimitKey:         GetPointerValue(gen.RateLimitKey, ""),
		RateLimited:          GetPointerValue(gen.RateLimited, false),
		WorkflowName:         GetPointerValue(gen.WorkflowName, ""),
		WorkflowVersion:      GetPointerValue(gen.WorkflowVersion, int32(0)),
		LastRetriedTime:      GetPointerValue(gen.LastRetriedTime, int64(0)),
		FailedTaskNames:      gen.FailedTaskNames,
		Variables:            gen.Variables,
		TaskToDomain:         GetPointerValue(gen.TaskToDomain, map[string]string{}),
		// Convert Tasks using existing mapper
		Tasks: toDomainTasksFromConductor(gen.Tasks),
		// Convert WorkflowDefinition using existing mapper
		WorkflowDefinition: toDomainWorkflowDefFromConductorGenerated(gen.WorkflowDefinition),
		// Convert History recursively
		History: toDomainWorkflowsFromConductorGenerated(gen.History),
	}
}

// toDomainWorkflowDefFromConductorGenerated converts conductor.WorkflowDef to model.WorkflowDef
//
//nolint:gocyclo // Generated mapper function with high complexity
func toDomainWorkflowDefFromConductorGenerated(gen *conductor.WorkflowDef) *model.WorkflowDef {
	if gen == nil {
		return nil
	}

	result := &model.WorkflowDef{
		// Required fields
		Name:           gen.Name,
		TimeoutSeconds: gen.TimeoutSeconds,
	}

	// Optional string fields
	if gen.OwnerApp != nil {
		result.OwnerApp = *gen.OwnerApp
	}
	if gen.CreatedBy != nil {
		result.CreatedBy = *gen.CreatedBy
	}
	if gen.UpdatedBy != nil {
		result.UpdatedBy = *gen.UpdatedBy
	}
	if gen.Description != nil {
		result.Description = *gen.Description
	}
	if gen.FailureWorkflow != nil {
		result.FailureWorkflow = *gen.FailureWorkflow
	}
	if gen.OwnerEmail != nil {
		result.OwnerEmail = *gen.OwnerEmail
	}
	if gen.TimeoutPolicy != nil {
		result.TimeoutPolicy = *gen.TimeoutPolicy
	}

	// Optional numeric fields
	if gen.CreateTime != nil {
		result.CreateTime = *gen.CreateTime
	}
	if gen.UpdateTime != nil {
		result.UpdateTime = *gen.UpdateTime
	}
	if gen.Version != nil {
		result.Version = *gen.Version
	}
	if gen.SchemaVersion != nil {
		result.SchemaVersion = *gen.SchemaVersion
	}

	// Optional boolean fields
	if gen.Restartable != nil {
		result.Restartable = *gen.Restartable
	}
	if gen.WorkflowStatusListenerEnabled != nil {
		result.WorkflowStatusListenerEnabled = *gen.WorkflowStatusListenerEnabled
	}

	// Slice fields
	if gen.InputParameters != nil {
		result.InputParameters = make([]string, len(gen.InputParameters))
		copy(result.InputParameters, gen.InputParameters)
	}

	// Map fields that need conversion from nested to flat
	if gen.OutputParameters != nil {
		result.OutputParameters = gen.OutputParameters
	}
	if gen.Variables != nil {
		result.Variables = gen.Variables
	}
	if gen.InputTemplate != nil {
		result.InputTemplate = gen.InputTemplate
	}

	if gen.Tasks != nil {
		result.Tasks = make([]model.WorkflowTask, len(gen.Tasks))
		for i, task := range gen.Tasks {
			result.Tasks[i] = toDomainModelWorkflowTaskFromConductor(&task)
		}
	}

	return result
}

// toDomainWorkflowsFromConductorGenerated converts []conductor.Workflow to []model.Workflow
func toDomainWorkflowsFromConductorGenerated(genWorkflows []conductor.Workflow) []model.Workflow {
	if genWorkflows == nil {
		return nil
	}
	workflows := make([]model.Workflow, len(genWorkflows))
	for i, gen := range genWorkflows {
		workflows[i] = toDomainWorkflowFromConductorGenerated(&gen)
	}
	return workflows
}

// toDomainWorkflowSummaryFromGenerated converts orkes.WorkflowSummary to model.WorkflowSummary
func toDomainWorkflowSummaryFromGenerated(gen *orkes.WorkflowSummary) model.WorkflowSummary {
	if gen == nil {
		return model.WorkflowSummary{}
	}
	return model.WorkflowSummary{
		WorkflowType:                     GetPointerValue(gen.WorkflowType, ""),
		Version:                          GetPointerValue(gen.Version, int32(0)),
		WorkflowId:                       GetPointerValue(gen.WorkflowId, ""),
		CorrelationId:                    GetPointerValue(gen.CorrelationId, ""),
		StartTime:                        GetPointerValue(gen.StartTime, ""),
		UpdateTime:                       GetPointerValue(gen.UpdateTime, ""),
		EndTime:                          GetPointerValue(gen.EndTime, ""),
		Status:                           GetPointerValue(gen.Status, ""),
		Input:                            GetPointerValue(gen.Input, ""),
		Output:                           GetPointerValue(gen.Output, ""),
		ReasonForIncompletion:            GetPointerValue(gen.ReasonForIncompletion, ""),
		ExecutionTime:                    GetPointerValue(gen.ExecutionTime, int64(0)),
		Event:                            GetPointerValue(gen.Event, ""),
		FailedReferenceTaskNames:         GetPointerValue(gen.FailedReferenceTaskNames, ""),
		ExternalInputPayloadStoragePath:  GetPointerValue(gen.ExternalInputPayloadStoragePath, ""),
		ExternalOutputPayloadStoragePath: GetPointerValue(gen.ExternalOutputPayloadStoragePath, ""),
		Priority:                         GetPointerValue(gen.Priority, int32(0)),
		OutputSize:                       GetPointerValue(gen.OutputSize, int64(0)),
		InputSize:                        GetPointerValue(gen.InputSize, int64(0)),
	}
}

func toDomainWorkflowSummaryFromConductorGenerated(gen *conductor.WorkflowSummary) model.WorkflowSummary {
	if gen == nil {
		return model.WorkflowSummary{}
	}
	return model.WorkflowSummary{
		WorkflowType:                     GetPointerValue(gen.WorkflowType, ""),
		Version:                          GetPointerValue(gen.Version, int32(0)),
		WorkflowId:                       GetPointerValue(gen.WorkflowId, ""),
		CorrelationId:                    GetPointerValue(gen.CorrelationId, ""),
		StartTime:                        GetPointerValue(gen.StartTime, ""),
		UpdateTime:                       GetPointerValue(gen.UpdateTime, ""),
		EndTime:                          GetPointerValue(gen.EndTime, ""),
		Status:                           GetPointerValue(gen.Status, ""),
		Input:                            GetPointerValue(gen.Input, ""),
		Output:                           GetPointerValue(gen.Output, ""),
		ReasonForIncompletion:            GetPointerValue(gen.ReasonForIncompletion, ""),
		ExecutionTime:                    GetPointerValue(gen.ExecutionTime, int64(0)),
		Event:                            GetPointerValue(gen.Event, ""),
		FailedReferenceTaskNames:         GetPointerValue(gen.FailedReferenceTaskNames, ""),
		ExternalInputPayloadStoragePath:  GetPointerValue(gen.ExternalInputPayloadStoragePath, ""),
		ExternalOutputPayloadStoragePath: GetPointerValue(gen.ExternalOutputPayloadStoragePath, ""),
		Priority:                         GetPointerValue(gen.Priority, int32(0)),
		OutputSize:                       GetPointerValue(gen.OutputSize, int64(0)),
		InputSize:                        GetPointerValue(gen.InputSize, int64(0)),
	}
}

// convertFromNestedMapInterface converts map[string]map[string]interface{} to map[string]interface{}
func convertFromNestedMapInterface(nested map[string]map[string]interface{}) map[string]interface{} {
	if nested == nil {
		return nil
	}
	flat := make(map[string]interface{}, len(nested))
	for k, v := range nested {
		flat[k] = v
	}
	return flat
}

// convertFromMapInterface converts any interface{} (including primitives, maps, arrays) into *interface{} for domain
func convertFromMapInterface(v interface{}) *interface{} {
	if v == nil {
		return nil
	}
	any := v
	return &any
}

// toDomainTasksFromGenerated converts []orkes.Task to []model.Task
func toDomainTasksFromGenerated(genTasks []orkes.Task) []model.Task {
	if genTasks == nil {
		return nil
	}

	tasks := make([]model.Task, len(genTasks))
	for i, gen := range genTasks {
		tasks[i] = toDomainTask(&gen)
	}
	return tasks
}

// toGeneratedUpsertUserRequest converts rbac.UpsertUserRequest to orkes.UpsertUserRequest
func toGeneratedUpsertUserRequest(domain rbac.UpsertUserRequest) orkes.UpsertUserRequest {
	return orkes.UpsertUserRequest{
		Groups: domain.Groups,
		Name:   domain.Name,
		Roles:  domain.Roles,
	}
}

// toDomainSignalResponseFromGenerated converts orkes.SignalResponse to model.SignalResponse
//
//nolint:gocyclo // Generated mapper function with high complexity
func toDomainSignalResponseFromGenerated(gen *orkes.SignalResponse) model.SignalResponse {
	if gen == nil {
		return model.SignalResponse{}
	}

	resp := model.SignalResponse{}
	resp.CorrelationID = GetPointerValue(gen.CorrelationId, "")
	resp.RequestID = GetPointerValue(gen.RequestId, "")
	resp.WorkflowId = GetPointerValue(gen.WorkflowId, "")
	resp.TargetWorkflowId = GetPointerValue(gen.TargetWorkflowId, "")

	// ResponseType
	rType := GetPointerValue(gen.ResponseType, "")
	switch rType {
	case string(model.ReturnTargetWorkflow):
		resp.ResponseType = model.ReturnTargetWorkflow
	case string(model.ReturnBlockingWorkflow):
		resp.ResponseType = model.ReturnBlockingWorkflow
	case string(model.ReturnBlockingTask):
		resp.ResponseType = model.ReturnBlockingTask
	case string(model.ReturnBlockingTaskInput):
		resp.ResponseType = model.ReturnBlockingTaskInput
	}

	// TargetWorkflowStatus
	wStatus := GetPointerValue(gen.TargetWorkflowStatus, "")
	switch wStatus {
	case string(model.RunningWorkflow):
		resp.TargetWorkflowStatus = model.RunningWorkflow
	case string(model.CompletedWorkflow):
		resp.TargetWorkflowStatus = model.CompletedWorkflow
	case string(model.FailedWorkflow):
		resp.TargetWorkflowStatus = model.FailedWorkflow
	case string(model.TerminatedWorkflow):
		resp.TargetWorkflowStatus = model.TerminatedWorkflow
	case string(model.TimedOutWorkflow):
		resp.TargetWorkflowStatus = model.TimedOutWorkflow
	}

	// Input/Output flatten
	if gen.Input != nil {
		resp.Input = gen.Input
	}
	if gen.Output != nil {
		resp.Output = gen.Output
	}

	// Map additional fields that were added to the generated SignalResponse
	resp.CreatedBy = gen.CreatedBy
	resp.CreateTime = gen.CreateTime

	// Map Tasks field
	if gen.Tasks != nil {
		resp.Tasks = make([]model.Task, len(gen.Tasks))
		for i, taskInterface := range gen.Tasks {
			if taskData, err := json.Marshal(taskInterface); err == nil {
				var task model.Task
				if err := json.Unmarshal(taskData, &task); err == nil {
					resp.Tasks[i] = task
				}
			}
		}
	}

	// Convert string status to WorkflowStatus
	if gen.Status != "" {
		switch gen.Status {
		case string(model.RunningWorkflow):
			resp.Status = model.RunningWorkflow
		case string(model.CompletedWorkflow):
			resp.Status = model.CompletedWorkflow
		case string(model.FailedWorkflow):
			resp.Status = model.FailedWorkflow
		case string(model.TerminatedWorkflow):
			resp.Status = model.TerminatedWorkflow
		case string(model.TimedOutWorkflow):
			resp.Status = model.TimedOutWorkflow
		}
	}

	resp.UpdateTime = gen.UpdateTime

	// Task-specific fields for BLOCKING_TASK strategies
	resp.TaskType = gen.TaskType
	resp.TaskId = gen.TaskId
	resp.ReferenceTaskName = gen.ReferenceTaskName
	resp.RetryCount = gen.RetryCount
	resp.TaskDefName = gen.TaskDefName
	resp.WorkflowType = gen.WorkflowType

	return resp
}

// toDomainExternalStorageLocationFromGenerated converts conductor.ExternalStorageLocation to model.ExternalStorageLocation

// toDomainPollDataFromMap converts interface{} to model.PollData
func toDomainPollDataFromMap(data interface{}) model.PollData {
	pd := model.PollData{}
	if data == nil {
		return pd
	}

	// Fast-path: already a map
	if m, ok := data.(map[string]any); ok {
		if v, ok := m["queueName"].(string); ok {
			pd.QueueName = v
		}
		if v, ok := m["domain"].(string); ok {
			pd.Domain = v
		}
		if v, ok := m["workerId"].(string); ok {
			pd.WorkerId = v
		}
		switch v := m["lastPollTime"].(type) {
		case float64:
			pd.LastPollTime = int64(v)
		case int64:
			pd.LastPollTime = v
		case int:
			pd.LastPollTime = int64(v)
		case json.Number:
			if iv, err := v.Int64(); err == nil {
				pd.LastPollTime = iv
			}
		}
		return pd
	}

	// Try JSON round-trip into struct directly
	if raw, err := json.Marshal(data); err == nil {
		var s model.PollData
		if err := json.Unmarshal(raw, &s); err == nil {
			return s
		}
		// Fallback: unmarshal into generic map then extract
		var gm map[string]any
		if err := json.Unmarshal(raw, &gm); err == nil {
			return toDomainPollDataFromMap(gm)
		}
	}

	return pd
}

// toDomainPollDataListFromGenerated converts []orkes.PollData to []model.PollData
func toDomainPollDataListFromGenerated(genPollData []orkes.PollData) []model.PollData {
	if genPollData == nil {
		return nil
	}
	pollData := make([]model.PollData, len(genPollData))
	for i, gen := range genPollData {
		pollData[i] = toDomainPollDataFromGenerated(&gen)
	}
	return pollData
}

// toDomainPollDataFromGenerated converts orkes.PollData to model.PollData
func toDomainPollDataFromGenerated(gen *orkes.PollData) model.PollData {
	if gen == nil {
		return model.PollData{}
	}

	return model.PollData{
		QueueName:    GetPointerValue(gen.QueueName, ""),
		Domain:       GetPointerValue(gen.Domain, ""),
		WorkerId:     GetPointerValue(gen.WorkerId, ""),
		LastPollTime: GetPointerValue(gen.LastPollTime, int64(0)),
	}
}

// toDomainTaskExecLogsFromGenerated converts []orkes.TaskExecLog to []model.TaskExecLog
func toDomainTaskExecLogsFromGenerated(genLogs []orkes.TaskExecLog) []model.TaskExecLog {
	if genLogs == nil {
		return nil
	}
	logs := make([]model.TaskExecLog, len(genLogs))
	for i, gen := range genLogs {
		logs[i] = toDomainTaskExecLogFromGenerated(&gen)
	}
	return logs
}

// toDomainTaskExecLogFromGenerated converts orkes.TaskExecLog to model.TaskExecLog
func toDomainTaskExecLogFromGenerated(gen *orkes.TaskExecLog) model.TaskExecLog {
	if gen == nil {
		return model.TaskExecLog{}
	}

	return model.TaskExecLog{
		Log:         GetPointerValue(gen.Log, ""),
		TaskId:      GetPointerValue(gen.TaskId, ""),
		CreatedTime: GetPointerValue(gen.CreatedTime, int64(0)),
	}
}

// toDomainSearchResultTaskSummaryFromGenerated converts orkes.SearchResultTaskSummary to model.SearchResultTaskSummary
func toDomainSearchResultTaskSummaryFromGenerated(gen *orkes.SearchResultTaskSummary) model.SearchResultTaskSummary {
	if gen == nil {
		return model.SearchResultTaskSummary{}
	}

	result := model.SearchResultTaskSummary{
		TotalHits: GetPointerValue(gen.TotalHits, int64(0)),
	}

	// Convert Results
	if gen.Results != nil {
		result.Results = make([]model.TaskSummary, len(gen.Results))
		for i, genTaskSummary := range gen.Results {
			result.Results[i] = toDomainTaskSummaryFromGenerated(&genTaskSummary)
		}
	}

	return result
}

// toDomainTaskSummaryFromGenerated converts orkes.TaskSummary to model.TaskSummary
func toDomainTaskSummaryFromGenerated(gen *orkes.TaskSummary) model.TaskSummary {
	if gen == nil {
		return model.TaskSummary{}
	}

	return model.TaskSummary{
		WorkflowId:                       GetPointerValue(gen.WorkflowId, ""),
		WorkflowType:                     GetPointerValue(gen.WorkflowType, ""),
		CorrelationId:                    GetPointerValue(gen.CorrelationId, ""),
		ScheduledTime:                    GetPointerValue(gen.ScheduledTime, ""),
		StartTime:                        GetPointerValue(gen.StartTime, ""),
		UpdateTime:                       GetPointerValue(gen.UpdateTime, ""),
		EndTime:                          GetPointerValue(gen.EndTime, ""),
		Status:                           GetPointerValue(gen.Status, ""),
		ReasonForIncompletion:            GetPointerValue(gen.ReasonForIncompletion, ""),
		ExecutionTime:                    GetPointerValue(gen.ExecutionTime, int64(0)),
		QueueWaitTime:                    GetPointerValue(gen.QueueWaitTime, int64(0)),
		TaskDefName:                      GetPointerValue(gen.TaskDefName, ""),
		TaskType:                         GetPointerValue(gen.TaskType, ""),
		Input:                            GetPointerValue(gen.Input, ""),
		Output:                           GetPointerValue(gen.Output, ""),
		TaskId:                           GetPointerValue(gen.TaskId, ""),
		ExternalInputPayloadStoragePath:  GetPointerValue(gen.ExternalInputPayloadStoragePath, ""),
		ExternalOutputPayloadStoragePath: GetPointerValue(gen.ExternalOutputPayloadStoragePath, ""),
		WorkflowPriority:                 GetPointerValue(gen.WorkflowPriority, int32(0)),
	}
}

// toDomainTasksFromConductor converts []conductor.Task to []model.Task
func toDomainTasksFromConductor(genTasks []conductor.Task) []model.Task {
	if genTasks == nil {
		return nil
	}
	tasks := make([]model.Task, len(genTasks))
	for i, gen := range genTasks {
		tasks[i] = toDomainTaskFromConductor(&gen)
	}
	return tasks
}

// toDomainTaskFromConductor converts *conductor.Task to model.Task using field-by-field mapping
func toDomainTaskFromConductor(gen *conductor.Task) model.Task {
	if gen == nil {
		return model.Task{}
	}

	domain := model.Task{
		TaskType:                         gen.GetTaskType(),
		Status:                           model.TaskResultStatus(gen.GetStatus()),
		ReferenceTaskName:                gen.GetReferenceTaskName(),
		RetryCount:                       gen.GetRetryCount(),
		Seq:                              gen.GetSeq(),
		CorrelationId:                    gen.GetCorrelationId(),
		PollCount:                        gen.GetPollCount(),
		TaskDefName:                      gen.GetTaskDefName(),
		ScheduledTime:                    gen.GetScheduledTime(),
		StartTime:                        gen.GetStartTime(),
		EndTime:                          gen.GetEndTime(),
		UpdateTime:                       gen.GetUpdateTime(),
		StartDelayInSeconds:              gen.GetStartDelayInSeconds(),
		RetriedTaskId:                    gen.GetRetriedTaskId(),
		Retried:                          gen.GetRetried(),
		Executed:                         gen.GetExecuted(),
		CallbackFromWorker:               gen.GetCallbackFromWorker(),
		ResponseTimeoutSeconds:           gen.GetResponseTimeoutSeconds(),
		WorkflowInstanceId:               gen.GetWorkflowInstanceId(),
		WorkflowType:                     gen.GetWorkflowType(),
		TaskId:                           gen.GetTaskId(),
		ReasonForIncompletion:            gen.GetReasonForIncompletion(),
		CallbackAfterSeconds:             gen.GetCallbackAfterSeconds(),
		WorkerId:                         gen.GetWorkerId(),
		Domain:                           gen.GetDomain(),
		RateLimitPerFrequency:            gen.GetRateLimitPerFrequency(),
		RateLimitFrequencyInSeconds:      gen.GetRateLimitFrequencyInSeconds(),
		ExternalInputPayloadStoragePath:  gen.GetExternalInputPayloadStoragePath(),
		ExternalOutputPayloadStoragePath: gen.GetExternalOutputPayloadStoragePath(),
		WorkflowPriority:                 gen.GetWorkflowPriority(),
		ExecutionNameSpace:               gen.GetExecutionNameSpace(),
		IsolationGroupId:                 gen.GetIsolationGroupId(),
		Iteration:                        gen.GetIteration(),
		QueueWaitTime:                    gen.GetQueueWaitTime(),
	}

	// InputData: map[string]map[string]interface{} -> map[string]interface{}
	if input := gen.GetInputData(); input != nil {
		domain.InputData = make(map[string]interface{}, len(input))
		for k, v := range input {
			domain.InputData[k] = v
		}
	}

	// OutputData: map[string]map[string]interface{} -> map[string]interface{}
	if output := gen.GetOutputData(); output != nil {
		domain.OutputData = make(map[string]interface{}, len(output))
		for k, v := range output {
			domain.OutputData[k] = v
		}
	}

	// WorkflowTask (use Ok accessor since getter returns by value)
	if wtPtr, ok := gen.GetWorkflowTaskOk(); ok && wtPtr != nil {
		converted := toDomainModelWorkflowTaskFromConductor(wtPtr)
		domain.WorkflowTask = &converted
	}

	// TaskDefinition
	if td := gen.TaskDefinition; td != nil {
		t := toDomainTaskDefPtrFromConductor(td)
		domain.TaskDefinition = &t
	}

	return domain
}

// toDomainModelWorkflowTaskFromConductor converts conductor.WorkflowTask to *model.WorkflowTask
func toDomainModelWorkflowTaskFromConductor(task *conductor.WorkflowTask) model.WorkflowTask {
	if task == nil {
		return model.WorkflowTask{}
	}
	// Convert InputParameters from map[string]map[string]interface{} to map[string]interface{}
	var inputParams map[string]interface{}
	if task.InputParameters != nil {
		inputParams = make(map[string]interface{}, len(task.InputParameters))
		for k, v := range task.InputParameters {
			inputParams[k] = v
		}
	}

	mt := model.WorkflowTask{
		Name:                           task.Name,
		TaskReferenceName:              task.TaskReferenceName,
		Description:                    GetPointerValue(task.Description, ""),
		InputParameters:                inputParams,
		Type_:                          GetPointerValue(task.Type, ""),
		Optional:                       GetPointerValue(task.Optional, false),
		StartDelay:                     GetPointerValue(task.StartDelay, int32(0)),
		AsyncComplete:                  GetPointerValue(task.AsyncComplete, false),
		RetryCount:                     GetPointerValue(task.RetryCount, int32(0)),
		RateLimited:                    GetPointerValue(task.RateLimited, false),
		LoopCondition:                  GetPointerValue(task.LoopCondition, ""),
		ScriptExpression:               GetPointerValue(task.ScriptExpression, ""),
		JoinOn:                         task.JoinOn,
		Sink:                           GetPointerValue(task.Sink, ""),
		Expression:                     GetPointerValue(task.Expression, ""),
		DynamicTaskNameParam:           GetPointerValue(task.DynamicTaskNameParam, ""),
		DynamicForkTasksParam:          GetPointerValue(task.DynamicForkTasksParam, ""),
		DynamicForkTasksInputParamName: GetPointerValue(task.DynamicForkTasksInputParamName, ""),
	}
	return mt
}

// toDomainTaskDefPtrFromConductor converts *conductor.TaskDef to model.TaskDef
func toDomainTaskDefPtrFromConductor(genDef *conductor.TaskDef) model.TaskDef {
	if genDef == nil {
		return model.TaskDef{}
	}
	// Convert InputTemplate map nesting
	var inputTemplate map[string]interface{}
	if genDef.InputTemplate != nil {
		inputTemplate = make(map[string]interface{}, len(genDef.InputTemplate))
		for k, v := range genDef.InputTemplate {
			inputTemplate[k] = v
		}
	}

	return model.TaskDef{
		Name:                        genDef.Name,
		Description:                 GetPointerValue(genDef.Description, ""),
		RetryCount:                  GetPointerValue(genDef.RetryCount, 0),
		TimeoutSeconds:              genDef.TimeoutSeconds,
		InputKeys:                   genDef.InputKeys,
		OutputKeys:                  genDef.OutputKeys,
		TimeoutPolicy:               GetPointerValue(genDef.TimeoutPolicy, ""),
		RetryLogic:                  GetPointerValue(genDef.RetryLogic, ""),
		RetryDelaySeconds:           GetPointerValue(genDef.RetryDelaySeconds, 0),
		ResponseTimeoutSeconds:      GetPointerValue(genDef.ResponseTimeoutSeconds, 0),
		ConcurrentExecLimit:         GetPointerValue(genDef.ConcurrentExecLimit, 0),
		InputTemplate:               inputTemplate,
		RateLimitPerFrequency:       GetPointerValue(genDef.RateLimitPerFrequency, 0),
		RateLimitFrequencyInSeconds: GetPointerValue(genDef.RateLimitFrequencyInSeconds, 0),
		IsolationGroupId:            GetPointerValue(genDef.IsolationGroupId, ""),
		ExecutionNameSpace:          GetPointerValue(genDef.ExecutionNameSpace, ""),
		OwnerEmail:                  GetPointerValue(genDef.OwnerEmail, ""),
		PollTimeoutSeconds:          GetPointerValue(genDef.PollTimeoutSeconds, 0),
		BackoffScaleFactor:          GetPointerValue(genDef.BackoffScaleFactor, 0),
	}
}

// toGeneratedServiceMethod converts model.ServiceMethod to orkes.ServiceMethod
func toGeneratedServiceMethod(domain *model.ServiceMethod) orkes.ServiceMethod {
	if domain == nil {
		return orkes.ServiceMethod{}
	}

	gen := orkes.ServiceMethod{}

	// Convert ExampleInput from map[string]interface{} to map[string]map[string]interface{}
	if domain.ExampleInput != nil {
		gen.ExampleInput = domain.ExampleInput
	}

	// Convert Id
	if domain.Id != 0 {
		gen.Id = &domain.Id
	}

	// Convert string fields
	if domain.InputType != "" {
		gen.InputType = &domain.InputType
	}
	if domain.MethodName != "" {
		gen.MethodName = &domain.MethodName
	}
	if domain.MethodType != "" {
		gen.MethodType = &domain.MethodType
	}
	if domain.OperationName != "" {
		gen.OperationName = &domain.OperationName
	}
	if domain.OutputType != "" {
		gen.OutputType = &domain.OutputType
	}

	// Convert RequestParams
	if domain.RequestParams != nil {
		gen.RequestParams = make([]orkes.RequestParam, len(domain.RequestParams))
		for i, param := range domain.RequestParams {
			gen.RequestParams[i] = toGeneratedServiceRegistryRequestParam(param)
		}
	}

	return gen
}

// toDomainGrantedAccessResponseFromMap converts map[string]interface{} to rbac.GrantedAccessResponse
func toDomainGrantedAccessResponseFromMap(data map[string]interface{}) rbac.GrantedAccessResponse {
	if data == nil {
		return rbac.GrantedAccessResponse{}
	}

	result := rbac.GrantedAccessResponse{}

	// Extract grantedAccess array from map
	if grantedAccessData, ok := data["grantedAccess"].([]interface{}); ok {
		result.GrantedAccess = make([]rbac.GrantedAccess, len(grantedAccessData))
		for i, accessData := range grantedAccessData {
			if accessMap, ok := accessData.(map[string]interface{}); ok {
				result.GrantedAccess[i] = toDomainGrantedAccessFromMap(accessMap)
			}
		}
	}

	return result
}

// toDomainGrantedAccessFromMap converts map[string]interface{} to rbac.GrantedAccess
func toDomainGrantedAccessFromMap(data map[string]interface{}) rbac.GrantedAccess {
	if data == nil {
		return rbac.GrantedAccess{}
	}

	access := rbac.GrantedAccess{}

	// Extract access array
	if accessData, ok := data["access"].([]interface{}); ok {
		access.Access = make([]string, len(accessData))
		for i, v := range accessData {
			if str, ok := v.(string); ok {
				access.Access[i] = str
			}
		}
	}

	// Extract tag
	if tag, ok := data["tag"].(string); ok {
		access.Tag = tag
	}

	// Extract target
	if targetData, ok := data["target"].(map[string]interface{}); ok {
		access.Target = toDomainTargetRefFromMap(targetData)
	}

	return access
}

// toDomainTargetRefFromMap converts map[string]interface{} to rbac.TargetRef
func toDomainTargetRefFromMap(data map[string]interface{}) *rbac.TargetRef {
	if data == nil {
		return nil
	}

	target := &rbac.TargetRef{}

	// Extract id
	if id, ok := data["id"].(string); ok {
		target.Id = id
	}

	// Extract type
	if type_, ok := data["type"].(string); ok {
		target.Type_ = type_
	}

	return target
}

// toDomainGrantedAccessResponseFromGenerated converts orkes.GrantedAccessResponse to rbac.GrantedAccessResponse
func toDomainGrantedAccessResponseFromGenerated(gen *orkes.GrantedAccessResponse) rbac.GrantedAccessResponse {
	if gen == nil {
		return rbac.GrantedAccessResponse{}
	}

	result := rbac.GrantedAccessResponse{}

	// Convert GrantedAccess array
	if gen.GrantedAccess != nil {
		result.GrantedAccess = make([]rbac.GrantedAccess, len(gen.GrantedAccess))
		for i, genAccess := range gen.GrantedAccess {
			result.GrantedAccess[i] = toDomainGrantedAccessFromGenerated(&genAccess)
		}
	}

	return result
}

// toDomainGrantedAccessFromGenerated converts orkes.GrantedAccess to rbac.GrantedAccess
func toDomainGrantedAccessFromGenerated(gen *orkes.GrantedAccess) rbac.GrantedAccess {
	if gen == nil {
		return rbac.GrantedAccess{}
	}

	access := rbac.GrantedAccess{}

	// Convert Access array
	if gen.Access != nil {
		access.Access = make([]string, len(gen.Access))
		copy(access.Access, gen.Access)
	}

	// Convert Tag
	access.Tag = GetPointerValue(gen.Tag, "")

	// Convert Target
	if gen.Target != nil {
		access.Target = toDomainTargetRefFromGenerated(gen.Target)
	}

	return access
}

// toDomainTargetRefFromGenerated converts orkes.TargetRef to rbac.TargetRef
func toDomainTargetRefFromGenerated(gen *orkes.TargetRef) *rbac.TargetRef {
	if gen == nil {
		return nil
	}

	return &rbac.TargetRef{
		Id:    gen.Id,
		Type_: gen.Type,
	}
}

// toGeneratedWorkflowTestRequest converts model.WorkflowTestRequest to orkes.WorkflowTestRequest
//
//nolint:gocyclo // Generated mapper function with high complexity
func toGeneratedWorkflowTestRequest(domain *model.WorkflowTestRequest) orkes.WorkflowTestRequest {
	gen := orkes.WorkflowTestRequest{}

	// Required field
	gen.Name = domain.Name

	// Simple string pointers
	if domain.CorrelationId != "" {
		gen.CorrelationId = &domain.CorrelationId
	}
	if domain.CreatedBy != "" {
		gen.CreatedBy = &domain.CreatedBy
	}
	if domain.ExternalInputPayloadStoragePath != "" {
		gen.ExternalInputPayloadStoragePath = &domain.ExternalInputPayloadStoragePath
	}
	if domain.IdempotencyKey != "" {
		gen.IdempotencyKey = &domain.IdempotencyKey
	}
	if domain.IdempotencyStrategy != "" {
		gen.IdempotencyStrategy = &domain.IdempotencyStrategy
	}

	if domain.Input != nil {
		gen.Input = domain.Input
	}

	// Numeric pointers
	if domain.Priority != 0 {
		gen.Priority = &domain.Priority
	}
	if domain.Version != 0 {
		gen.Version = &domain.Version
	}

	// TaskToDomain
	if domain.TaskToDomain != nil {
		m := make(map[string]string, len(domain.TaskToDomain))
		for k, v := range domain.TaskToDomain {
			m[k] = v
		}
		gen.TaskToDomain = &m
	}

	// TaskRefToMockOutput
	if domain.TaskRefToMockOutput != nil {
		m := make(map[string][]orkes.TaskMock, len(domain.TaskRefToMockOutput))
		for k, mocks := range domain.TaskRefToMockOutput {
			if mocks == nil {
				continue
			}
			arr := make([]orkes.TaskMock, len(mocks))
			for i := range mocks {
				arr[i] = toGeneratedTaskMock(mocks[i])
			}
			m[k] = arr
		}
		gen.TaskRefToMockOutput = &m
	}

	// SubWorkflowTestRequest (recursive)
	if domain.SubWorkflowTestRequest != nil {
		m := make(map[string]orkes.WorkflowTestRequest, len(domain.SubWorkflowTestRequest))
		for k, v := range domain.SubWorkflowTestRequest {
			m[k] = toGeneratedWorkflowTestRequest(&v)
		}
		gen.SubWorkflowTestRequest = &m
	}

	// WorkflowDef (convert domain -> generated) - comprehensive field-by-field mapping
	if domain.WorkflowDef != nil {
		gen.WorkflowDef = toGeneratedWorkflowDef(domain.WorkflowDef)
	}

	return gen
}

// toGeneratedTaskMock converts model.TaskMock to orkes.TaskMock
func toGeneratedTaskMock(domain model.TaskMock) orkes.TaskMock {
	gen := orkes.TaskMock{}
	if domain.ExecutionTime != 0 {
		gen.ExecutionTime = &domain.ExecutionTime
	}
	if domain.QueueWaitTime != 0 {
		gen.QueueWaitTime = &domain.QueueWaitTime
	}
	if domain.Status != "" {
		gen.Status = &domain.Status
	}
	if domain.Output != nil {
		gen.Output = domain.Output
	}
	return gen
}

// toDomainBulkResponseFromConductor converts conductor.BulkResponseString to model.BulkResponse

// toDomainHealthCheckStatusFromOrkes converts a generic map response (ORKES) to model.HealthCheckStatus
func toDomainHealthCheckStatusFromOrkes(gen map[string]interface{}) model.HealthCheckStatus {
	if gen == nil {
		return model.HealthCheckStatus{}
	}
	status := model.HealthCheckStatus{}
	// healthy
	if v, ok := gen["healthy"]; ok {
		if b, ok := v.(bool); ok {
			status.Healthy = b
		}
	}
	// healthResults
	if v, ok := gen["healthResults"]; ok {
		status.HealthResults = toDomainHealthSliceFromAny(v)
	}
	// suppressedHealthResults
	if v, ok := gen["suppressedHealthResults"]; ok {
		status.SuppressedHealthResults = toDomainHealthSliceFromAny(v)
	}
	return status
}

// toDomainHealthSliceFromAny converts any array into []model.Health
func toDomainHealthSliceFromAny(v interface{}) []model.Health {
	arr, ok := v.([]interface{})
	if !ok || arr == nil {
		return nil
	}
	out := make([]model.Health, 0, len(arr))
	for _, item := range arr {
		out = append(out, toDomainHealthFromAny(item))
	}
	return out
}

// toDomainHealthFromAny converts a generic map (ORKES) to model.Health
func toDomainHealthFromAny(v interface{}) model.Health {
	m, ok := v.(map[string]interface{})
	if !ok || m == nil {
		return model.Health{}
	}
	d := model.Health{}
	// healthy
	if hv, ok := m["healthy"]; ok {
		if b, ok := hv.(bool); ok {
			d.Healthy = b
		}
	}
	// errorMessage
	if ev, ok := m["errorMessage"]; ok {
		if s, ok := ev.(string); ok {
			d.ErrorMessage = s
		}
	}
	if dv, ok := m["details"]; ok {
		switch dm := dv.(type) {
		case map[string]interface{}:
			d.Details = dm
		case map[string]map[string]interface{}:
			converted := make(map[string]interface{}, len(dm))
			for k, inner := range dm {
				innerCopy := make(map[string]interface{}, len(inner))
				for ik, iv := range inner {
					innerCopy[ik] = iv
				}
				converted[k] = innerCopy
			}
			d.Details = converted
		}
	}
	return d
}

func toGeneratedActionForOrkes(domain *model.Action) orkes.Action {
	var gen orkes.Action
	if domain == nil {
		return gen
	}
	if domain.Action != "" {
		gen.Action = &domain.Action
	}
	if domain.StartWorkflow != nil {
		sw := orkes.StartWorkflowRequest{
			Name:    domain.StartWorkflow.Name,
			Version: ToPointer(int32(domain.StartWorkflow.Version)),
		}
		if domain.StartWorkflow.Input != nil {
			sw.Input = domain.StartWorkflow.Input
		}
		gen.StartWorkflow = &sw
	}
	if domain.CompleteTask != nil {
		td := orkes.TaskDetails{
			TaskId: ToPointer(domain.CompleteTask.TaskId),
			Output: domain.CompleteTask.Output,
		}
		if domain.CompleteTask.TaskRefName != "" {
			td.TaskRefName = &domain.CompleteTask.TaskRefName
		}
		if domain.CompleteTask.WorkflowId != "" {
			td.WorkflowId = &domain.CompleteTask.WorkflowId
		}
		gen.CompleteTask = &td
	}
	if domain.FailTask != nil {
		td := orkes.TaskDetails{
			TaskId: ToPointer(domain.FailTask.TaskId),
			Output: domain.FailTask.Output,
		}
		if domain.FailTask.TaskRefName != "" {
			td.TaskRefName = &domain.FailTask.TaskRefName
		}
		if domain.FailTask.WorkflowId != "" {
			td.WorkflowId = &domain.FailTask.WorkflowId
		}
		gen.FailTask = &td
	}
	gen.ExpandInlineJSON = &domain.ExpandInlineJSON
	return gen
}

func toMapFromWorkflowDef(domain *model.WorkflowDef) map[string]interface{} {
	if domain == nil {
		return nil
	}
	return map[string]interface{}{
		"ownerApp":                      domain.OwnerApp,
		"createTime":                    domain.CreateTime,
		"updateTime":                    domain.UpdateTime,
		"createdBy":                     domain.CreatedBy,
		"updatedBy":                     domain.UpdatedBy,
		"name":                          domain.Name,
		"description":                   domain.Description,
		"version":                       domain.Version,
		"ownerEmail":                    domain.OwnerEmail,
		"failureWorkflow":               domain.FailureWorkflow,
		"schemaVersion":                 domain.SchemaVersion,
		"restartable":                   domain.Restartable,
		"workflowStatusListenerEnabled": domain.WorkflowStatusListenerEnabled,
		"timeoutPolicy":                 domain.TimeoutPolicy,
		"timeoutSeconds":                domain.TimeoutSeconds,
		"inputParameters":               domain.InputParameters,
		"outputParameters":              domain.OutputParameters,
		"variables":                     domain.Variables,
		"inputTemplate":                 domain.InputTemplate,
		"tags": func() []map[string]interface{} {
			if domain.Tags == nil {
				return nil
			}
			tags := make([]map[string]interface{}, len(domain.Tags))
			for i, t := range domain.Tags {
				tags[i] = map[string]interface{}{
					"key":   t.Key,
					"type":  t.Type,
					"value": t.Value,
				}
			}
			return tags
		}(),
		"overwriteTags": domain.OverwriteTags,
		"rateLimitConfig": func() map[string]interface{} {
			if domain.RateLimitConfig == nil {
				return nil
			}
			return map[string]interface{}{
				"rateLimitKey":        domain.RateLimitConfig.RateLimitKey,
				"concurrentExecLimit": domain.RateLimitConfig.ConcurrentExecLimit,
			}
		}(),
		"tasks": toGenericWorkflowTasks(domain.Tasks),
	}
}

//nolint:gocyclo,gocognit
func fromMapToWorkflowDef(m map[string]interface{}) *model.WorkflowDef {
	if m == nil {
		return nil
	}
	wf := &model.WorkflowDef{}
	if v, ok := m["ownerApp"].(string); ok {
		wf.OwnerApp = v
	}
	if v, ok := m["createTime"].(float64); ok { // JSON numbers decode to float64
		wf.CreateTime = int64(v)
	}
	if v, ok := m["updateTime"].(float64); ok {
		wf.UpdateTime = int64(v)
	}
	if v, ok := m["createdBy"].(string); ok {
		wf.CreatedBy = v
	}
	if v, ok := m["updatedBy"].(string); ok {
		wf.UpdatedBy = v
	}
	if name, ok := m["name"].(string); ok {
		wf.Name = name
	}
	if desc, ok := m["description"].(string); ok {
		wf.Description = desc
	}
	if ver, ok := m["version"].(float64); ok {
		wf.Version = int32(ver)
	}
	if ownerEmail, ok := m["ownerEmail"].(string); ok {
		wf.OwnerEmail = ownerEmail
	}
	if v, ok := m["failureWorkflow"].(string); ok {
		wf.FailureWorkflow = v
	}
	if v, ok := m["schemaVersion"].(float64); ok {
		wf.SchemaVersion = int32(v)
	}
	if v, ok := m["restartable"].(bool); ok {
		wf.Restartable = v
	}
	if v, ok := m["workflowStatusListenerEnabled"].(bool); ok {
		wf.WorkflowStatusListenerEnabled = v
	}
	if v, ok := m["timeoutPolicy"].(string); ok {
		wf.TimeoutPolicy = v
	}
	if timeout, ok := m["timeoutSeconds"].(float64); ok {
		wf.TimeoutSeconds = int64(timeout)
	}
	if ip, ok := m["inputParameters"].([]interface{}); ok {
		wf.InputParameters = make([]string, len(ip))
		for i := range ip {
			if s, sok := ip[i].(string); sok {
				wf.InputParameters[i] = s
			}
		}
	}
	if op, ok := m["outputParameters"].(map[string]interface{}); ok {
		wf.OutputParameters = op
	}
	if vars, ok := m["variables"].(map[string]interface{}); ok {
		wf.Variables = vars
	}
	if tpl, ok := m["inputTemplate"].(map[string]interface{}); ok {
		wf.InputTemplate = tpl
	}
	if tags, ok := m["tags"].([]interface{}); ok {
		wf.Tags = make([]model.TagObject, len(tags))
		for i := range tags {
			if tagMap, ok := tags[i].(map[string]interface{}); ok {
				if kv, ok := tagMap["key"].(string); ok {
					wf.Tags[i].Key = kv
				}
				if tv, ok := tagMap["type"].(string); ok {
					wf.Tags[i].Type = tv
				}
				if vv, ok := tagMap["value"].(string); ok {
					wf.Tags[i].Value = vv
				}
			}
		}
	}
	if v, ok := m["overwriteTags"].(bool); ok {
		wf.OverwriteTags = v
	}
	if rlc, ok := m["rateLimitConfig"].(map[string]interface{}); ok {
		cfg := &model.RateLimitConfig{}
		if k, ok := rlc["rateLimitKey"].(string); ok {
			cfg.RateLimitKey = k
		}
		if c, ok := rlc["concurrentExecLimit"].(float64); ok {
			cfg.ConcurrentExecLimit = int32(c)
		}
		wf.RateLimitConfig = cfg
	}
	if tasks, ok := m["tasks"].([]interface{}); ok {
		wf.Tasks = fromGenericWorkflowTasks(tasks)
	}
	return wf
}

//nolint:gocyclo,gocognit
func toGenericWorkflowTasks(tasks []model.WorkflowTask) []interface{} {
	if tasks == nil {
		return nil
	}
	res := make([]interface{}, len(tasks))
	for i := range tasks {
		m := map[string]interface{}{
			"name":                           tasks[i].Name,
			"taskReferenceName":              tasks[i].TaskReferenceName,
			"description":                    tasks[i].Description,
			"inputParameters":                tasks[i].InputParameters,
			"type":                           tasks[i].Type_,
			"dynamicTaskNameParam":           tasks[i].DynamicTaskNameParam,
			"caseValueParam":                 tasks[i].CaseValueParam,
			"caseExpression":                 tasks[i].CaseExpression,
			"scriptExpression":               tasks[i].ScriptExpression,
			"startDelay":                     tasks[i].StartDelay,
			"joinOn":                         tasks[i].JoinOn,
			"sink":                           tasks[i].Sink,
			"optional":                       tasks[i].Optional,
			"rateLimited":                    tasks[i].RateLimited,
			"asyncComplete":                  tasks[i].AsyncComplete,
			"loopCondition":                  tasks[i].LoopCondition,
			"retryCount":                     tasks[i].RetryCount,
			"expression":                     tasks[i].Expression,
			"dynamicForkJoinTasksParam":      tasks[i].DynamicForkJoinTasksParam,
			"dynamicForkTasksParam":          tasks[i].DynamicForkTasksParam,
			"dynamicForkTasksInputParamName": tasks[i].DynamicForkTasksInputParamName,
		}

		// Only include enum-like strings when non-empty to avoid server coercion errors
		if tasks[i].EvaluatorType != "" {
			m["evaluatorType"] = tasks[i].EvaluatorType
		}
		if tasks[i].WorkflowTaskType != "" {
			m["workflowTaskType"] = tasks[i].WorkflowTaskType
		}

		if tasks[i].DecisionCases != nil {
			dc := make(map[string]interface{}, len(tasks[i].DecisionCases))
			for k, arr := range tasks[i].DecisionCases {
				dc[k] = toGenericWorkflowTasks(arr)
			}
			m["decisionCases"] = dc
		}
		if len(tasks[i].DefaultCase) > 0 {
			m["defaultCase"] = toGenericWorkflowTasks(tasks[i].DefaultCase)
		}
		if len(tasks[i].ForkTasks) > 0 {
			forks := make([]interface{}, len(tasks[i].ForkTasks))
			for j := range tasks[i].ForkTasks {
				forks[j] = toGenericWorkflowTasks(tasks[i].ForkTasks[j])
			}
			m["forkTasks"] = forks
		}
		if tasks[i].SubWorkflowParam != nil {
			sw := map[string]interface{}{
				"name":    tasks[i].SubWorkflowParam.Name,
				"version": tasks[i].SubWorkflowParam.Version,
			}
			if tasks[i].SubWorkflowParam.TaskToDomain != nil {
				sw["taskToDomain"] = tasks[i].SubWorkflowParam.TaskToDomain
			}
			if tasks[i].SubWorkflowParam.WorkflowDefinition != nil {
				sw["workflowDefinition"] = toMapFromWorkflowDef(tasks[i].SubWorkflowParam.WorkflowDefinition)
			}
			m["subWorkflowParam"] = sw
		}
		if len(tasks[i].LoopOver) > 0 {
			m["loopOver"] = toGenericWorkflowTasks(tasks[i].LoopOver)
		}
		if len(tasks[i].DefaultExclusiveJoinTask) > 0 {
			m["defaultExclusiveJoinTask"] = tasks[i].DefaultExclusiveJoinTask
		}
		if tasks[i].TaskDefinition != nil {
			// Minimal pass-through for task definition fields used in JSON
			td := map[string]interface{}{
				"name":                        tasks[i].TaskDefinition.Name,
				"description":                 tasks[i].TaskDefinition.Description,
				"retryCount":                  tasks[i].TaskDefinition.RetryCount,
				"timeoutSeconds":              tasks[i].TaskDefinition.TimeoutSeconds,
				"inputKeys":                   tasks[i].TaskDefinition.InputKeys,
				"outputKeys":                  tasks[i].TaskDefinition.OutputKeys,
				"timeoutPolicy":               tasks[i].TaskDefinition.TimeoutPolicy,
				"retryLogic":                  tasks[i].TaskDefinition.RetryLogic,
				"retryDelaySeconds":           tasks[i].TaskDefinition.RetryDelaySeconds,
				"responseTimeoutSeconds":      tasks[i].TaskDefinition.ResponseTimeoutSeconds,
				"concurrentExecLimit":         tasks[i].TaskDefinition.ConcurrentExecLimit,
				"inputTemplate":               tasks[i].TaskDefinition.InputTemplate,
				"rateLimitPerFrequency":       tasks[i].TaskDefinition.RateLimitPerFrequency,
				"rateLimitFrequencyInSeconds": tasks[i].TaskDefinition.RateLimitFrequencyInSeconds,
				"isolationGroupId":            tasks[i].TaskDefinition.IsolationGroupId,
				"executionNameSpace":          tasks[i].TaskDefinition.ExecutionNameSpace,
				"ownerEmail":                  tasks[i].TaskDefinition.OwnerEmail,
				"pollTimeoutSeconds":          tasks[i].TaskDefinition.PollTimeoutSeconds,
				"backoffScaleFactor":          tasks[i].TaskDefinition.BackoffScaleFactor,
				"overwriteTags":               tasks[i].TaskDefinition.OverwriteTags,
			}
			if tasks[i].TaskDefinition.Tags != nil {
				tags := make([]map[string]interface{}, len(tasks[i].TaskDefinition.Tags))
				for k, t := range tasks[i].TaskDefinition.Tags {
					tags[k] = map[string]interface{}{"key": t.Key, "type": t.Type, "value": t.Value}
				}
				td["tags"] = tags
			}
			m["taskDefinition"] = td
		}
		if tasks[i].CacheConfig != nil {
			m["cacheConfig"] = map[string]interface{}{
				"key":         tasks[i].CacheConfig.Key,
				"ttlInSecond": tasks[i].CacheConfig.TtlInSeconds,
			}
		}
		res[i] = m
	}
	return res
}

//nolint:gocognit,gocyclo // Generated mapper function with high complexity
func fromGenericWorkflowTasks(items []interface{}) []model.WorkflowTask {
	if items == nil {
		return nil
	}
	res := make([]model.WorkflowTask, len(items))
	for i := range items {
		if m, ok := items[i].(map[string]interface{}); ok {
			wt := model.WorkflowTask{}
			if v, ok := m["name"].(string); ok {
				wt.Name = v
			}
			if v, ok := m["taskReferenceName"].(string); ok {
				wt.TaskReferenceName = v
			}
			if v, ok := m["description"].(string); ok {
				wt.Description = v
			}
			if v, ok := m["inputParameters"].(map[string]interface{}); ok {
				wt.InputParameters = v
			}
			if v, ok := m["type"].(string); ok {
				wt.Type_ = v
			}
			if v, ok := m["dynamicTaskNameParam"].(string); ok {
				wt.DynamicTaskNameParam = v
			}
			if v, ok := m["caseValueParam"].(string); ok {
				wt.CaseValueParam = v
			}
			if v, ok := m["caseExpression"].(string); ok {
				wt.CaseExpression = v
			}
			if v, ok := m["optional"].(bool); ok {
				wt.Optional = v
			}
			if v, ok := m["startDelay"].(float64); ok {
				wt.StartDelay = int32(v)
			}
			if v, ok := m["asyncComplete"].(bool); ok {
				wt.AsyncComplete = v
			}
			if v, ok := m["retryCount"].(float64); ok {
				wt.RetryCount = int32(v)
			}
			if v, ok := m["rateLimited"].(bool); ok {
				wt.RateLimited = v
			}
			if v, ok := m["loopCondition"].(string); ok {
				wt.LoopCondition = v
			}
			if v, ok := m["scriptExpression"].(string); ok {
				wt.ScriptExpression = v
			}
			if v, ok := m["joinOn"].([]interface{}); ok {
				wt.JoinOn = make([]string, len(v))
				for j := range v {
					if s, sok := v[j].(string); sok {
						wt.JoinOn[j] = s
					}
				}
			}
			if v, ok := m["sink"].(string); ok {
				wt.Sink = v
			}
			if v, ok := m["expression"].(string); ok {
				wt.Expression = v
			}
			if v, ok := m["dynamicForkTasksParam"].(string); ok {
				wt.DynamicForkTasksParam = v
			}
			if v, ok := m["dynamicForkTasksInputParamName"].(string); ok {
				wt.DynamicForkTasksInputParamName = v
			}
			if v, ok := m["dynamicForkJoinTasksParam"].(string); ok {
				wt.DynamicForkJoinTasksParam = v
			}
			if v, ok := m["evaluatorType"].(string); ok {
				wt.EvaluatorType = v
			}
			if v, ok := m["workflowTaskType"].(string); ok {
				wt.WorkflowTaskType = v
			}
			if v, ok := m["defaultExclusiveJoinTask"].([]interface{}); ok {
				wt.DefaultExclusiveJoinTask = make([]string, len(v))
				for j := range v {
					if s, sok := v[j].(string); sok {
						wt.DefaultExclusiveJoinTask[j] = s
					}
				}
			}
			if v, ok := m["decisionCases"].(map[string]interface{}); ok {
				wt.DecisionCases = make(map[string][]model.WorkflowTask, len(v))
				for key, arr := range v {
					if list, ok := arr.([]interface{}); ok {
						wt.DecisionCases[key] = fromGenericWorkflowTasks(list)
					}
				}
			}
			if v, ok := m["defaultCase"].([]interface{}); ok {
				wt.DefaultCase = fromGenericWorkflowTasks(v)
			}
			if v, ok := m["forkTasks"].([]interface{}); ok {
				wt.ForkTasks = make([][]model.WorkflowTask, len(v))
				for j := range v {
					if inner, ok := v[j].([]interface{}); ok {
						wt.ForkTasks[j] = fromGenericWorkflowTasks(inner)
					}
				}
			}
			if v, ok := m["loopOver"].([]interface{}); ok {
				wt.LoopOver = fromGenericWorkflowTasks(v)
			}
			if sw, ok := m["subWorkflowParam"].(map[string]interface{}); ok {
				p := &model.SubWorkflowParams{}
				if n, ok := sw["name"].(string); ok {
					p.Name = n
				}
				if ver, ok := sw["version"].(float64); ok {
					p.Version = int32(ver)
				}
				if td, ok := sw["taskToDomain"].(map[string]interface{}); ok {
					p.TaskToDomain = make(map[string]string, len(td))
					for k, val := range td {
						if s, sok := val.(string); sok {
							p.TaskToDomain[k] = s
						}
					}
				}
				if wfd, ok := sw["workflowDefinition"].(map[string]interface{}); ok {
					p.WorkflowDefinition = fromMapToWorkflowDef(wfd)
				}
				wt.SubWorkflowParam = p
			}
			if td, ok := m["taskDefinition"].(map[string]interface{}); ok {
				d := &model.TaskDef{}
				if s, ok := td["name"].(string); ok {
					d.Name = s
				}
				if s, ok := td["description"].(string); ok {
					d.Description = s
				}
				if n, ok := td["retryCount"].(float64); ok {
					d.RetryCount = int32(n)
				}
				if n, ok := td["timeoutSeconds"].(float64); ok {
					d.TimeoutSeconds = int64(n)
				}
				if arr, ok := td["inputKeys"].([]interface{}); ok {
					d.InputKeys = make([]string, len(arr))
					for j := range arr {
						if s, sok := arr[j].(string); sok {
							d.InputKeys[j] = s
						}
					}
				}
				if arr, ok := td["outputKeys"].([]interface{}); ok {
					d.OutputKeys = make([]string, len(arr))
					for j := range arr {
						if s, sok := arr[j].(string); sok {
							d.OutputKeys[j] = s
						}
					}
				}
				if s, ok := td["timeoutPolicy"].(string); ok {
					d.TimeoutPolicy = s
				}
				if s, ok := td["retryLogic"].(string); ok {
					d.RetryLogic = s
				}
				if n, ok := td["retryDelaySeconds"].(float64); ok {
					d.RetryDelaySeconds = int32(n)
				}
				if n, ok := td["responseTimeoutSeconds"].(float64); ok {
					d.ResponseTimeoutSeconds = int64(n)
				}
				if n, ok := td["concurrentExecLimit"].(float64); ok {
					d.ConcurrentExecLimit = int32(n)
				}
				if mp, ok := td["inputTemplate"].(map[string]interface{}); ok {
					d.InputTemplate = mp
				}
				if n, ok := td["rateLimitPerFrequency"].(float64); ok {
					d.RateLimitPerFrequency = int32(n)
				}
				if n, ok := td["rateLimitFrequencyInSeconds"].(float64); ok {
					d.RateLimitFrequencyInSeconds = int32(n)
				}
				if s, ok := td["isolationGroupId"].(string); ok {
					d.IsolationGroupId = s
				}
				if s, ok := td["executionNameSpace"].(string); ok {
					d.ExecutionNameSpace = s
				}
				if s, ok := td["ownerEmail"].(string); ok {
					d.OwnerEmail = s
				}
				if n, ok := td["pollTimeoutSeconds"].(float64); ok {
					d.PollTimeoutSeconds = int32(n)
				}
				if n, ok := td["backoffScaleFactor"].(float64); ok {
					d.BackoffScaleFactor = int32(n)
				}
				if arr, ok := td["tags"].([]interface{}); ok {
					d.Tags = make([]model.TagObject, len(arr))
					for j := range arr {
						if tagMap, ok := arr[j].(map[string]interface{}); ok {
							if s, ok := tagMap["key"].(string); ok {
								d.Tags[j].Key = s
							}
							if s, ok := tagMap["type"].(string); ok {
								d.Tags[j].Type = s
							}
							if s, ok := tagMap["value"].(string); ok {
								d.Tags[j].Value = s
							}
						}
					}
				}
				if b, ok := td["overwriteTags"].(bool); ok {
					d.OverwriteTags = b
				}
				wt.TaskDefinition = d
			}
			if cc, ok := m["cacheConfig"].(map[string]interface{}); ok {
				c := &model.CacheConfig{}
				if s, ok := cc["key"].(string); ok {
					c.Key = s
				}
				if n, ok := cc["ttlInSecond"].(float64); ok {
					c.TtlInSeconds = int(n)
				}
				wt.CacheConfig = c
			}
			res[i] = wt
		}
	}
	return res
}
