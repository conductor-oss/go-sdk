package client

import (
	"context"
	"net/http"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

type WorkflowBulkClient interface {
	Delete(ctx context.Context, workflowIDs []string) (model.BulkResponse, *http.Response, error)
	Pause(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
	Restart(ctx context.Context, workflowIds []string, opts *WorkflowBulkResourceApiRestartOpts) (model.BulkResponse, *http.Response, error)
	Resume(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
	Retry(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
	Terminate(ctx context.Context, workflowIds []string, opts *WorkflowBulkResourceApiTerminateOpts) (model.BulkResponse, *http.Response, error)
}

func NewWorkflowBulkClient(client *APIClient) WorkflowBulkClient {
	return &WorkflowBulkResourceApiService{client}
}
