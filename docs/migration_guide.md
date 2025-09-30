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

## Additional Notes

### ArchiveWorkflow Parameter (Deprecated)

The `ArchiveWorkflow` parameter in `WorkflowResourceApiDeleteOpts` is deprecated and has no effect when configured:

```go
type WorkflowResourceApiDeleteOpts struct {
    // Deprecated: There is no effect when configured.
    ArchiveWorkflow optional.Bool
}
```

## Timeline

- **Current Version**: Deprecated methods are still available but marked for removal
- **Next Major Version**: Deprecated methods will be removed
- **Recommended Action**: Migrate to the new method names as soon as possible
