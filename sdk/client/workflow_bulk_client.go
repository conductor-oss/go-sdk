package client

import (
	"context"
	"net/http"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// WorkflowBulkClient is the client for the workflow bulk API
type WorkflowBulkClient interface {
	// Deprecated: Use Pause instead.
	PauseWorkflow1(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
	// Pause pauses the list of workflows.
	Pause(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
	// Restart restarts the list of completed workflow.
	Restart(ctx context.Context, workflowIds []string, opts *WorkflowBulkResourceApiRestartOpts) (model.BulkResponse, *http.Response, error)
	// Resume resumes the list of workflows.
	Resume(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
	// Deprecated: Use Resume instead.
	ResumeWorkflow(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
	// Retry retries the last failed task for each workflow from the list.
	Retry(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
	// Deprecated: Use Retry instead.
	Retry1(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
	// Terminate terminates the list of workflows.
	Terminate(ctx context.Context, workflowIds []string, opts *WorkflowBulkResourceApiTerminateOpts) (model.BulkResponse, *http.Response, error)
	// Delete permanently removes workflows from the system.
	Delete(ctx context.Context, workflowIDs []string) (model.BulkResponse, *http.Response, error)
}

// NewWorkflowBulkClient creates a new WorkflowBulkClient
func NewWorkflowBulkClient(client *APIClient) WorkflowBulkClient {
	return &WorkflowBulkResourceApiService{client}
}
