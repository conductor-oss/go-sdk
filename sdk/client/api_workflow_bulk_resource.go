//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package client

import (
	"context"
	"errors"
	"net/http"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

type WorkflowBulkResourceApiService struct {
	*APIClient
}

// NewWorkflowBulkResourceApiService creates a new WorkflowBulkResourceApiService instance
func NewWorkflowBulkResourceApiService(client *APIClient) *WorkflowBulkResourceApiService {
	return &WorkflowBulkResourceApiService{APIClient: client}
}

// Delete Permanently remove workflows from the system
func (a *WorkflowBulkResourceApiService) Delete(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	req := a.http_orkes.WorkflowBulkResourceAPI.Delete(ctx).RequestBody(body)

	genResp, resp, err := req.Execute()
	if err != nil {
		return model.BulkResponse{}, resp, wrapGeneratedError(err, resp)
	}

	// Convert using mapper
	result := toDomainBulkResponseFromGenerated(*genResp)
	return result, resp, nil
}

// Pause Pause the list of workflows
func (a *WorkflowBulkResourceApiService) Pause(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	req := a.http_orkes.WorkflowBulkResourceAPI.PauseWorkflow1(ctx).RequestBody(body)

	genResp, resp, err := req.Execute()
	if err != nil {
		return model.BulkResponse{}, resp, wrapGeneratedError(err, resp)
	}

	if genResp == nil {
		return model.BulkResponse{}, resp, wrapGeneratedError(errors.New("bulk response not found"), resp)
	}

	// Convert using mapper
	result := toDomainBulkResponseFromGenerated(*genResp)
	return result, resp, nil
}

// WorkflowBulkResourceApiRestartOpts Optional parameters for Restart
type WorkflowBulkResourceApiRestartOpts struct {
	UseLatestDefinitions optional.Bool
}

// Restart Restart the list of workflows
func (a *WorkflowBulkResourceApiService) Restart(ctx context.Context, body []string, opts *WorkflowBulkResourceApiRestartOpts) (model.BulkResponse, *http.Response, error) {
	req := a.http_orkes.WorkflowBulkResourceAPI.Restart1(ctx).RequestBody(body)

	if opts != nil && opts.UseLatestDefinitions.IsSet() {
		req = req.UseLatestDefinitions(opts.UseLatestDefinitions.Value())
	}
	genResp, resp, err := req.Execute()
	if err != nil {
		return model.BulkResponse{}, resp, wrapGeneratedError(err, resp)
	}

	if genResp == nil {
		return model.BulkResponse{}, resp, wrapGeneratedError(errors.New("bulk response not found"), resp)
	}

	// Convert using mapper
	result := toDomainBulkResponseFromGenerated(*genResp)
	return result, resp, nil
}

// Resume Resume the list of workflows
func (a *WorkflowBulkResourceApiService) Resume(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	req := a.http_orkes.WorkflowBulkResourceAPI.ResumeWorkflow1(ctx).RequestBody(body)

	genResp, resp, err := req.Execute()
	if err != nil {
		return model.BulkResponse{}, resp, wrapGeneratedError(err, resp)
	}

	if genResp == nil {
		return model.BulkResponse{}, resp, wrapGeneratedError(errors.New("bulk response not found"), resp)
	}

	// Convert using mapper
	result := toDomainBulkResponseFromGenerated(*genResp)
	return result, resp, nil
}

// Retry Retry the last failed task for each workflow from the list
func (a *WorkflowBulkResourceApiService) Retry(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	req := a.http_orkes.WorkflowBulkResourceAPI.Retry1(ctx).RequestBody(body)

	genResp, resp, err := req.Execute()
	if err != nil {
		return model.BulkResponse{}, resp, wrapGeneratedError(err, resp)
	}

	if genResp == nil {
		return model.BulkResponse{}, resp, wrapGeneratedError(errors.New("bulk response not found"), resp)
	}

	// Convert using mapper
	result := toDomainBulkResponseFromGenerated(*genResp)
	return result, resp, nil
}

// WorkflowBulkResourceApiTerminateOpts Optional parameters for Terminate
type WorkflowBulkResourceApiTerminateOpts struct {
	Reason                 optional.String
	TriggerFailureWorkflow optional.Bool
}

// Terminate Terminate workflows execution
func (a *WorkflowBulkResourceApiService) Terminate(ctx context.Context, body []string, opts *WorkflowBulkResourceApiTerminateOpts) (model.BulkResponse, *http.Response, error) {
	req := a.http_orkes.WorkflowBulkResourceAPI.Terminate(ctx).RequestBody(body)

	if opts != nil && opts.Reason.IsSet() {
		req = req.Reason(opts.Reason.Value())
	}
	if opts != nil && opts.TriggerFailureWorkflow.IsSet() {
		req = req.TriggerFailureWorkflow(opts.TriggerFailureWorkflow.Value())
	}

	genResp, resp, err := req.Execute()
	if err != nil {
		return model.BulkResponse{}, resp, wrapGeneratedError(err, resp)
	}

	if genResp == nil {
		return model.BulkResponse{}, resp, wrapGeneratedError(errors.New("bulk response not found"), resp)
	}

	// Convert using mapper
	result := toDomainBulkResponseFromGenerated(*genResp)
	return result, resp, nil
}
