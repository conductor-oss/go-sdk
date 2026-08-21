# Migration Guide: Deprecated Methods

This guide helps you migrate from deprecated methods to their recommended replacements in the Conductor Go SDK.

## WorkflowClient Deprecated Methods

### 1. PauseWorkflow → Pause

**Deprecated Method:**
```go
PauseWorkflow(ctx context.Context, workflowId string) (*http.Response, error)
```

**Replacement:**
```go
Pause(ctx context.Context, workflowId string) (*http.Response, error)
```

**Migration:**
```go
// Before
resp, err := client.PauseWorkflow(ctx, workflowId)

// After
resp, err := client.Pause(ctx, workflowId)
```

### 2. ResetWorkflow → Reset

**Deprecated Method:**
```go
ResetWorkflow(ctx context.Context, workflowId string) (*http.Response, error)
```

**Replacement:**
```go
Reset(ctx context.Context, workflowId string) (*http.Response, error)
```

**Migration:**
```go
// Before
resp, err := client.ResetWorkflow(ctx, workflowId)

// After
resp, err := client.Reset(ctx, workflowId)
```

### 3. ResumeWorkflow → Resume

**Deprecated Method:**
```go
ResumeWorkflow(ctx context.Context, workflowId string) (*http.Response, error)
```

**Replacement:**
```go
Resume(ctx context.Context, workflowId string) (*http.Response, error)
```

**Migration:**
```go
// Before
resp, err := client.ResumeWorkflow(ctx, workflowId)

// After
resp, err := client.Resume(ctx, workflowId)
```

## WorkflowBulkClient Deprecated Methods

### 1. PauseWorkflow1 → Pause

**Deprecated Method:**
```go
PauseWorkflow1(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
```

**Replacement:**
```go
Pause(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
```

**Migration:**
```go
// Before
bulkResp, resp, err := bulkClient.PauseWorkflow1(ctx, workflowIds)

// After
bulkResp, resp, err := bulkClient.Pause(ctx, workflowIds)
```

### 2. ResumeWorkflow → Resume

**Deprecated Method:**
```go
ResumeWorkflow(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
```

**Replacement:**
```go
Resume(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
```

**Migration:**
```go
// Before
bulkResp, resp, err := bulkClient.ResumeWorkflow(ctx, workflowIds)

// After
bulkResp, resp, err := bulkClient.Resume(ctx, workflowIds)
```

### 3. Retry1 → Retry

**Deprecated Method:**
```go
Retry1(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
```

**Replacement:**
```go
Retry(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
```

**Migration:**
```go
// Before
bulkResp, resp, err := bulkClient.Retry1(ctx, workflowIds)

// After
bulkResp, resp, err := bulkClient.Retry(ctx, workflowIds)
```

### 4. WorkflowBulkResourceApiRestart1Opts → WorkflowBulkResourceApiRestartOpts

**Deprecated Type:**
```go
type WorkflowBulkResourceApiRestart1Opts struct {
    UseLatestDefinitions optional.Bool
}
```

**Replacement:**
```go
type WorkflowBulkResourceApiRestartOpts struct {
    UseLatestDefinitions optional.Bool
}
```

**Migration:**
```go
// Before
opts := &client.WorkflowBulkResourceApiRestart1Opts{
    UseLatestDefinitions: optional.NewBool(true),
}

// After
opts := &client.WorkflowBulkResourceApiRestartOpts{
    UseLatestDefinitions: optional.NewBool(true),
}
```

## API Client Deprecated Methods

### 1. NewAuthenticationSettingsFromEnv → settings.NewClientSettingsFromEnv

**Deprecated Method:**
```go
func NewAuthenticationSettingsFromEnv() *settings.AuthenticationSettings
```

**Replacement:**
```go
func NewClientSettingsFromEnv() *settings.ClientSettings
```

**Migration:**
```go
// Before
authSettings := client.NewAuthenticationSettingsFromEnv()

// After
clientSettings := settings.NewClientSettingsFromEnv()
authSettings := clientSettings.Authentication
```

### 2. NewHttpSettingsFromEnv → settings.NewHttpSettingsFromEnv

**Deprecated Method:**
```go
func NewHttpSettingsFromEnv() *settings.HttpSettings
```

**Replacement:**
```go
func NewHttpSettingsFromEnv() *settings.HttpSettings
```

**Migration:**
```go
// Before
httpSettings := client.NewHttpSettingsFromEnv()

// After
httpSettings := settings.NewHttpSettingsFromEnv()
```

## Additional Notes

### ArchiveWorkflow Parameter (Deprecated)

The `ArchiveWorkflow` parameter in `WorkflowResourceApiDeleteOpts` is deprecated and has no effect when configured:

```go
type WorkflowResourceApiDeleteOpts struct {
    // Deprecated: There is no effect when configured.
    ArchiveWorkflow optional.Bool
}
```

### SignalResponse.Status Field Type Change

The `Status` field in `model.SignalResponse` was changed from `WorkflowStatus` to `string` to support both workflow and task statuses in the unified response. Use helper methods like `GetWorkflow()`, `GetBlockingTask()`, `GetWorkflowRun()`, or `GetTaskRun()` which handle parsing automatically, or use `ParseWorkflowStatus()`/`ParseTaskResultStatus()` for manual parsing.

## ApplicationClient Method Signature Changes

### 1. AddRoleToApplicationUser Return Type Change

**Before:**
```go
AddRoleToApplicationUser(ctx context.Context, applicationId string, role string) (interface{}, *http.Response, error)
```

**After:**
```go
AddRoleToApplicationUser(ctx context.Context, applicationId string, role string) (*http.Response, error)
```

**Migration:**
```go
// Before
result, resp, err := client.AddRoleToApplicationUser(ctx, applicationId, role)

// After
resp, err := client.AddRoleToApplicationUser(ctx, applicationId, role)
```

### 2. CreateAccessKey Return Type Change

**Before:**
```go
CreateAccessKey(ctx context.Context, id string) (*rbac.ConductorApplication, *http.Response, error)
```

**After:**
```go
CreateAccessKey(ctx context.Context, id string) (*rbac.CreateAccessKeyResponse, *http.Response, error)
```

**Migration:**
```go
// Before
app, resp, err := client.CreateAccessKey(ctx, id)

// After
accessKey, resp, err := client.CreateAccessKey(ctx, id)
// accessKey now contains Id and Secret fields
```

### 3. DeleteApplication Return Type Change

**Before:**
```go
DeleteApplication(ctx context.Context, id string) (interface{}, *http.Response, error)
```

**After:**
```go
DeleteApplication(ctx context.Context, id string) (*rbac.DeleteApplicationResponse, *http.Response, error)
```

**Migration:**
```go
// Before
result, resp, err := client.DeleteApplication(ctx, id)

// After
response, resp, err := client.DeleteApplication(ctx, id)
// response now contains a Message field with operation result
```

### 4. GetAppByAccessKeyId Return Type Change

**Before:**
```go
GetAppByAccessKeyId(ctx context.Context, accessKeyId string) (interface{}, *http.Response, error)
```

**After:**
```go
GetAppByAccessKeyId(ctx context.Context, accessKeyId string) (*rbac.ConductorApplication, *http.Response, error)
```

**Migration:**
```go
// Before
result, resp, err := client.GetAppByAccessKeyId(ctx, accessKeyId)

// After
app, resp, err := client.GetAppByAccessKeyId(ctx, accessKeyId)
// app is now properly typed as ConductorApplication
```

### 5. RemoveRoleFromApplicationUser Return Type Change

**Before:**
```go
RemoveRoleFromApplicationUser(ctx context.Context, applicationId string, role string) (interface{}, *http.Response, error)
```

**After:**
```go
RemoveRoleFromApplicationUser(ctx context.Context, applicationId string, role string) (*http.Response, error)
```

**Migration:**
```go
// Before
result, resp, err := client.RemoveRoleFromApplicationUser(ctx, applicationId, role)

// After
resp, err := client.RemoveRoleFromApplicationUser(ctx, applicationId, role)
```

### 6. ToggleAccessKeyStatus Return Type Change

**Before:**
```go
ToggleAccessKeyStatus(ctx context.Context, applicationId string, keyId string) (interface{}, *http.Response, error)
```

**After:**
```go
ToggleAccessKeyStatus(ctx context.Context, applicationId string, keyId string) (*rbac.AccessKeyResponse, *http.Response, error)
```

**Migration:**
```go
// Before
result, resp, err := client.ToggleAccessKeyStatus(ctx, applicationId, keyId)

// After
accessKey, resp, err := client.ToggleAccessKeyStatus(ctx, applicationId, keyId)
// accessKey now contains proper AccessKeyResponse with Status field
```

### 7. NewApplicationClient Return Type Change

**Before:**
```go
func NewApplicationClient(apiClient *APIClient) *ApplicationResourceApiService
```

**After:**
```go
func NewApplicationClient(apiClient *APIClient) ApplicationClient
```

**Migration:**
```go
// Before
client := NewApplicationClient(apiClient) // Returns *ApplicationResourceApiService

// After
client := NewApplicationClient(apiClient) // Returns ApplicationClient interface
```

## Timeline

- **Current Version**: Deprecated methods are still available but marked for removal
- **Next Major Version**: Deprecated methods will be removed
- **Recommended Action**: Migrate to the new method names as soon as possible
