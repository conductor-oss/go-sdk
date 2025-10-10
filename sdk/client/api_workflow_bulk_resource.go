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
	"net/http"
	"net/url"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// Linger please
var (
	_ context.Context
)

type WorkflowBulkResourceApiService struct {
	*APIClient
}

// Deprecated: Use Pause instead.
func (a *WorkflowBulkResourceApiService) PauseWorkflow1(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	return a.Pause(ctx, body)
}

// Pause pauses the list of workflows.
func (a *WorkflowBulkResourceApiService) Pause(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	var result model.BulkResponse

	localVarPath := "/workflow/bulk/pause"

	resp, err := a.Put(ctx, localVarPath, body, &result)
	if err != nil {
		return model.BulkResponse{}, resp, err
	}
	return result, resp, nil
}

// WorkflowBulkResourceApiRestartOpts Optional parameters for Restart.
type WorkflowBulkResourceApiRestartOpts struct {
	// UseLatestDefinitions if set to true, the restarted workflow will use the latest definition.
	UseLatestDefinitions optional.Bool
}

// Deprecated: Use WorkflowBulkResourceApiRestartOpts instead.
type WorkflowBulkResourceApiRestart1Opts = WorkflowBulkResourceApiRestartOpts

// Restart restarts the list of workflows.
func (a *WorkflowBulkResourceApiService) Restart(ctx context.Context, body []string, localVarOptionals *WorkflowBulkResourceApiRestartOpts) (model.BulkResponse, *http.Response, error) {
	var result model.BulkResponse

	path := "/workflow/bulk/restart"

	queryParams := url.Values{}
	if localVarOptionals != nil && localVarOptionals.UseLatestDefinitions.IsSet() {
		queryParams.Add("useLatestDefinitions", parameterToString(localVarOptionals.UseLatestDefinitions.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, &result)
	if err != nil {
		return model.BulkResponse{}, resp, err
	}
	return result, resp, nil
}

// Deprecated: Use Resume instead.
func (a *WorkflowBulkResourceApiService) ResumeWorkflow(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	return a.Resume(ctx, body)
}

// Resume resumes the list of workflows.
func (a *WorkflowBulkResourceApiService) Resume(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	var result model.BulkResponse

	path := "/workflow/bulk/resume"

	resp, err := a.Put(ctx, path, body, &result)
	if err != nil {
		return model.BulkResponse{}, resp, err
	}
	return result, resp, nil
}

// Retry retries the last failed task for each workflow from the list.
func (a *WorkflowBulkResourceApiService) Retry(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	var result model.BulkResponse

	path := "/workflow/bulk/retry"

	resp, err := a.Post(ctx, path, body, &result)
	if err != nil {
		return model.BulkResponse{}, resp, err
	}
	return result, resp, nil
}

// Deprecated: Use Retry instead.
func (a *WorkflowBulkResourceApiService) Retry1(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	return a.Retry(ctx, body)
}

// WorkflowBulkResourceApiTerminateOpts Optional parameters for Terminate.
type WorkflowBulkResourceApiTerminateOpts struct {
	// Reason a reason for termination.
	Reason optional.String
	// TriggerFailureWorkflow if set to true, the associated compensation flow  will be triggered.
	TriggerFailureWorkflow optional.Bool
}

// Terminate terminates workflows execution.
func (a *WorkflowBulkResourceApiService) Terminate(ctx context.Context, body []string, opts *WorkflowBulkResourceApiTerminateOpts) (model.BulkResponse, *http.Response, error) {
	var result model.BulkResponse

	path := "/workflow/bulk/terminate"

	queryParams := url.Values{}
	if opts != nil && opts.Reason.IsSet() {
		queryParams.Add("reason", parameterToString(opts.Reason.Value(), ""))
	}
	if opts != nil && opts.TriggerFailureWorkflow.IsSet() {
		queryParams.Add("triggerFailureWorkflow", parameterToString(opts.TriggerFailureWorkflow.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, &result)
	if err != nil {
		return model.BulkResponse{}, resp, err
	}
	return result, resp, nil
}

// Delete permanently removes workflows from the system.
func (a *WorkflowBulkResourceApiService) Delete(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	var result model.BulkResponse

	path := "/workflow/bulk/delete"

	resp, err := a.Post(ctx, path, body, &result)
	if err != nil {
		return model.BulkResponse{}, resp, err
	}
	return result, resp, nil
}
