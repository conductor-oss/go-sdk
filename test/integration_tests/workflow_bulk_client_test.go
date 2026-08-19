//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package integration_tests

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestWorkflowBulkDelete tests the Delete endpoint for bulk workflow operations
func TestWorkflowBulkDelete(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	testdata.SkipIfOSS(t, "POST /workflow/bulk/delete is not implemented by plain OSS Conductor, confirmed empirically (404 'No static resource api/workflow/bulk/delete')")

	// Create a simple workflow for testing
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_BULK_DELETE_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var").Input("var_value", "bulk_delete_test"))

	err := testdata.ValidateWorkflowRegistration(wf)
	require.NoError(t, err, "Failed to register workflow")

	defer func() {
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	}()

	// Start multiple workflows
	workflowIds := make([]string, 3)
	for i := 0; i < 3; i++ {
		startRequest := &model.StartWorkflowRequest{
			Name:    wf.GetName(),
			Version: 1,
			Input:   map[string]interface{}{"instance": i, "test": "bulk_delete"},
		}

		workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
		require.NoError(t, err, "Failed to start workflow %d", i)
		workflowIds[i] = workflowId
	}

	// Wait for workflows to complete
	err = testdata.WaitForMultipleWorkflowsCompletion(workflowIds, testdata.WorkflowValidationTimeout)
	require.NoError(t, err, "Failed to wait for workflows to complete")

	// Test bulk delete
	ctx := context.Background()

	response, _, err := testdata.WorkflowBulkClient.Delete(ctx, workflowIds)
	require.NoError(t, err, "Failed to delete workflows in bulk")

	// Verify bulk response
	assert.Equal(t, len(response.BulkErrorResults), 0, "Bulk error results should be empty")
	assert.Equal(t, len(response.BulkSuccessfulResults), len(workflowIds),
		"Bulk successful results should be equal to workflow count")

	// Verify workflows are deleted. Bulk delete is accepted before it has taken effect,
	// so poll rather than reading once — a single immediate read is a race that shows up
	// intermittently in CI.
	for _, id := range workflowIds {
		var lastErr error
		deadline := time.Now().Add(testdata.WorkflowValidationTimeout)
		for time.Now().Before(deadline) {
			_, lastErr = testdata.WorkflowExecutor.GetWorkflow(id, false)
			if lastErr != nil {
				break
			}
			time.Sleep(250 * time.Millisecond)
		}
		require.Error(t, lastErr, "Workflow %s should be deleted", id)
		assert.Contains(t, lastErr.Error(), "no such workflow by Id", "Error should indicate workflow not found")
	}
}

// TestWorkflowBulkPause tests the PauseWorkflow1 endpoint for bulk workflow operations
func TestWorkflowBulkPause(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a long-running workflow for testing
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_BULK_PAUSE_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewWaitForDurationTask("wait_for_pause", 60*time.Second))

	err := testdata.ValidateWorkflowRegistration(wf)
	require.NoError(t, err, "Failed to register workflow")

	// Start multiple workflows
	workflowIds := make([]string, 3)
	for i := 0; i < 3; i++ {
		startRequest := &model.StartWorkflowRequest{
			Name:    wf.GetName(),
			Version: 1,
			Input:   map[string]interface{}{"instance": i, "test": "bulk_pause"},
		}

		workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
		require.NoError(t, err, "Failed to start workflow %d", i)
		workflowIds[i] = workflowId
	}

	defer func() {
		for _, id := range workflowIds {
			err = testdata.WorkflowExecutor.RemoveWorkflow(id)
			assert.NoError(t, err, "Failed to remove workflow %s", id)
		}

		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	}()

	// Wait for workflows to be running
	for i, id := range workflowIds {
		workflow, err := testdata.WaitForWorkflowRunning(id, testdata.WorkflowValidationTimeout)
		assert.NoError(t, err, "Failed to get workflow %d", i)
		assert.Equal(t, model.RunningWorkflow, workflow.Status, "Workflow %d should be running", i)
	}

	// Test bulk pause
	ctx := context.Background()

	response, _, err := testdata.WorkflowBulkClient.Pause(ctx, workflowIds)
	require.NoError(t, err, "Failed to pause workflows in bulk")

	// Verify bulk response
	assert.Equal(t, len(response.BulkErrorResults), 0, "Bulk error results should be empty")
	assert.Equal(t, len(response.BulkSuccessfulResults), len(workflowIds),
		"Bulk successful results should be equal to workflow count")

	// Verify workflows are paused
	for _, id := range workflowIds {
		workflow, err := testdata.WaitForWorkflowStatus(id, []model.WorkflowStatus{model.PausedWorkflow}, testdata.WorkflowValidationTimeout)
		require.NoError(t, err, "Failed to get workflow %s", id)
		assert.Equal(t, model.PausedWorkflow, workflow.Status, "Workflow %s should be paused", id)
	}
}

// TestWorkflowBulkResume tests the ResumeWorkflow1 endpoint for bulk workflow operations
func TestWorkflowBulkResume(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a long-running workflow for testing
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_BULK_RESUME_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewWaitForDurationTask("wait_for_resume", 60*time.Second))

	err := testdata.ValidateWorkflowRegistration(wf)
	require.NoError(t, err, "Failed to register workflow")

	// Start multiple workflows
	workflowIds := make([]string, 3)
	for i := 0; i < 3; i++ {
		startRequest := &model.StartWorkflowRequest{
			Name:    wf.GetName(),
			Version: 1,
			Input:   map[string]interface{}{"instance": i, "test": "bulk_resume"},
		}

		workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
		require.NoError(t, err, "Failed to start workflow %d", i)
		workflowIds[i] = workflowId
	}

	defer func() {
		for _, id := range workflowIds {
			err = testdata.WorkflowExecutor.RemoveWorkflow(id)
			assert.NoError(t, err, "Failed to remove workflow %s", id)
		}

		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	}()

	// Wait for workflows to be running
	for _, id := range workflowIds {
		workflow, err := testdata.WaitForWorkflowStatus(id, []model.WorkflowStatus{model.RunningWorkflow}, testdata.WorkflowValidationTimeout)
		require.NoError(t, err, "Failed to get workflow %s", id)
		assert.Equal(t, model.RunningWorkflow, workflow.Status, "Workflow %s should be running", id)
	}

	// Pause workflows first
	ctx := context.Background()

	_, _, err = testdata.WorkflowBulkClient.Pause(ctx, workflowIds)
	require.NoError(t, err, "Failed to pause workflows in bulk")

	// Verify workflows are paused
	for _, id := range workflowIds {
		workflow, err := testdata.WaitForWorkflowStatus(id, []model.WorkflowStatus{model.PausedWorkflow}, testdata.WorkflowValidationTimeout)
		require.NoError(t, err, "Failed to get workflow %s", id)
		assert.Equal(t, model.PausedWorkflow, workflow.Status, "Workflow %s should be paused", id)
	}

	// Test bulk resume
	response, _, err := testdata.WorkflowBulkClient.Resume(ctx, workflowIds)
	require.NoError(t, err, "Failed to resume workflows in bulk")

	// Verify bulk response
	assert.Equal(t, len(response.BulkErrorResults), 0, "Bulk error results should be empty")
	assert.Equal(t, len(response.BulkSuccessfulResults), len(workflowIds),
		"Bulk successful results should be equal to workflow count")

	// Verify workflows are running again
	for _, id := range workflowIds {
		workflow, err := testdata.WaitForWorkflowStatus(id, []model.WorkflowStatus{model.RunningWorkflow}, testdata.WorkflowValidationTimeout)
		require.NoError(t, err, "Failed to get workflow %s", id)
		assert.Equal(t, model.RunningWorkflow, workflow.Status, "Workflow %s should be running after resume", id)
	}
}

// TestWorkflowBulkRestart tests the Restart endpoint for bulk workflow operations
func TestWorkflowBulkRestart(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a simple workflow for testing
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wfV1 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_BULK_RESTART_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var").Input("var_value", "bulk_restart_test"))

	wfV2 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_BULK_RESTART_" + uniqueSuffix).
		Version(2).
		Add(workflow.NewSetVariableTask("set_var").Input("var_value", "v2")).
		Add(workflow.NewSetVariableTask("set_var_extra").Input("var_value", "extra"))

	require.NoError(t, testdata.ValidateWorkflowRegistration(wfV1), "register v1")
	require.NoError(t, testdata.ValidateWorkflowRegistration(wfV2), "register v2")

	// Start multiple workflows
	workflowIds := make([]string, 3)
	for i := 0; i < 3; i++ {
		startRequest := &model.StartWorkflowRequest{
			Name:    wfV1.GetName(),
			Version: 1,
			Input:   map[string]interface{}{"instance": i, "test": "bulk_restart"},
		}

		workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
		require.NoError(t, err, "Failed to start workflow %d", i)
		workflowIds[i] = workflowId
	}
	defer func() {
		for _, id := range workflowIds {
			err := testdata.WorkflowExecutor.RemoveWorkflow(id)
			assert.NoError(t, err, "Failed to remove workflow %s", id)
		}

		_, err := testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wfV1.GetName(),
			wfV1.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wfV2.GetName(),
			wfV2.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	}()

	// Wait for workflows to complete
	err := testdata.WaitForMultipleWorkflowsCompletion(workflowIds, testdata.WorkflowValidationTimeout)
	require.NoError(t, err, "Failed to wait for workflows to complete")

	// Test bulk restart with useLatestDefinitions=false (should restart with v1)
	ctx := context.Background()

	optsV1 := &client.WorkflowBulkResourceApiRestartOpts{
		UseLatestDefinitions: optional.NewBool(false),
	}

	responseV1, _, err := testdata.WorkflowBulkClient.Restart(ctx, workflowIds, optsV1)
	require.NoError(t, err, "Failed to restart workflows in bulk with v1")
	assert.Equal(t, len(responseV1.BulkErrorResults), 0, "Bulk error results should be empty")
	assert.Equal(t, len(responseV1.BulkSuccessfulResults), len(workflowIds),
		"Bulk successful results should be equal to workflow count")

	// Wait for restarted workflows to complete with v1
	for _, id := range workflowIds {
		workflow, err := testdata.WaitForWorkflowStatus(id, []model.WorkflowStatus{model.CompletedWorkflow}, testdata.WorkflowValidationTimeout)
		require.NoError(t, err, "Failed to get workflow %s after v1 restart", id)
		assert.Equal(t, model.CompletedWorkflow, workflow.Status, "Workflow %s should be completed after v1 restart", id)

		// Verify v1 workflow structure (should only have set_var task, no set_var_extra)
		var hasSetVarExtra bool
		for _, task := range workflow.Tasks {
			if task.ReferenceTaskName == "set_var_extra" {
				hasSetVarExtra = true
				break
			}
		}
		assert.False(t, hasSetVarExtra, "Workflow %s should not have set_var_extra task after v1 restart", id)
	}

	// Now test bulk restart with useLatestDefinitions=true (should restart with v2)
	optsV2 := &client.WorkflowBulkResourceApiRestartOpts{
		UseLatestDefinitions: optional.NewBool(true),
	}

	responseV2, _, err := testdata.WorkflowBulkClient.Restart(ctx, workflowIds, optsV2)
	require.NoError(t, err, "Failed to restart workflows in bulk with v2")
	assert.Equal(t, len(responseV2.BulkErrorResults), 0, "Bulk error results should be empty")
	assert.Equal(t, len(responseV2.BulkSuccessfulResults), len(workflowIds),
		"Bulk successful results should be equal to workflow count")

	// Wait for restarted workflows to complete with v2
	for _, id := range workflowIds {
		workflow, err := testdata.WaitForWorkflowStatus(id, []model.WorkflowStatus{model.CompletedWorkflow}, testdata.WorkflowValidationTimeout)
		require.NoError(t, err, "Failed to get workflow %s after v2 restart", id)
		assert.Equal(t, model.CompletedWorkflow, workflow.Status, "Workflow %s should be completed after v2 restart", id)

		// Verify v2 workflow structure (should have both set_var and set_var_extra tasks)
		var hasSetVar bool
		var hasSetVarExtra bool
		for _, task := range workflow.Tasks {
			if task.ReferenceTaskName == "set_var" {
				hasSetVar = true
			}
			if task.ReferenceTaskName == "set_var_extra" {
				hasSetVarExtra = true
			}
		}
		assert.True(t, hasSetVar, "Workflow %s should have set_var task after v2 restart", id)
		assert.True(t, hasSetVarExtra, "Workflow %s should have set_var_extra task after v2 restart", id)
	}
}

// TestWorkflowBulkRetry tests the Retry1 endpoint for bulk workflow operations
func TestWorkflowBulkRetry(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a failing workflow for testing
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_BULK_RETRY_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewWaitForDurationTask("wait_for_retry", 10*time.Second))

	err := testdata.ValidateWorkflowRegistration(wf)
	require.NoError(t, err, "Failed to register workflow")

	// Start multiple workflows
	workflowIds := make([]string, 3)
	for i := 0; i < 3; i++ {
		startRequest := &model.StartWorkflowRequest{
			Name:    wf.GetName(),
			Version: 1,
			Input:   map[string]interface{}{"instance": i, "test": "bulk_retry"},
		}

		workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
		require.NoError(t, err, "Failed to start workflow %d", i)
		workflowIds[i] = workflowId
	}

	defer func() {
		for _, id := range workflowIds {
			err := testdata.WorkflowExecutor.RemoveWorkflow(id)
			assert.NoError(t, err, "Failed to remove workflow %s", id)
		}

		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	}()

	// Wait for workflows to be running
	for _, id := range workflowIds {
		runningWorkflow, err := testdata.WaitForWorkflowRunning(id, testdata.WorkflowValidationTimeout)
		assert.NoError(t, err, "Failed to wait for workflow to be running")
		assert.Equal(t, model.RunningWorkflow, runningWorkflow.Status, "Workflow should be running")
	}

	// Test bulk retry
	ctx := context.Background()

	_, _, err = testdata.WorkflowBulkClient.Terminate(ctx, workflowIds, nil)
	require.NoError(t, err, "Failed to terminate workflows in bulk")

	// Verify workflow is terminated
	for _, id := range workflowIds {
		terminatedWorkflow, err := testdata.WaitForWorkflowStatus(id, []model.WorkflowStatus{model.TerminatedWorkflow}, testdata.WorkflowValidationTimeout)
		assert.NoError(t, err, "Failed to wait for workflow termination")
		assert.Equal(t, model.TerminatedWorkflow, terminatedWorkflow.Status, "Workflow should be terminated")
	}

	response, _, err := testdata.WorkflowBulkClient.Retry(ctx, workflowIds)
	require.NoError(t, err, "Failed to retry workflows in bulk")

	// Verify bulk response
	assert.Equal(t, len(response.BulkErrorResults), 0, "Bulk error results should be empty")
	assert.Equal(t, len(response.BulkSuccessfulResults), len(workflowIds),
		"Bulk successful results should be equal to workflow count")

	// Wait for retried workflows to be running or completed
	for _, id := range workflowIds {
		workflow, err := testdata.WaitForWorkflowStatus(id, []model.WorkflowStatus{model.RunningWorkflow, model.CompletedWorkflow}, testdata.WorkflowValidationTimeout)
		require.NoError(t, err, "Failed to get workflow %s", id)
		assert.Contains(t, []model.WorkflowStatus{model.RunningWorkflow, model.CompletedWorkflow},
			workflow.Status, "Workflow %s should be running, completed, or failed after retry", id)
	}
}

// TestWorkflowBulkTerminate tests the Terminate endpoint for bulk workflow operations
func TestWorkflowBulkTerminate(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a long-running workflow for testing
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_BULK_TERMINATE_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewWaitForDurationTask("wait_for_terminate", 60*time.Second))

	err := testdata.ValidateWorkflowRegistration(wf)
	require.NoError(t, err, "Failed to register workflow")

	// Start multiple workflows
	workflowIds := make([]string, 3)
	for i := 0; i < 3; i++ {
		startRequest := &model.StartWorkflowRequest{
			Name:    wf.GetName(),
			Version: 1,
			Input:   map[string]interface{}{"instance": i, "test": "bulk_terminate"},
		}

		workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
		require.NoError(t, err, "Failed to start workflow %d", i)
		workflowIds[i] = workflowId
	}
	defer func() {
		for _, id := range workflowIds {
			err := testdata.WorkflowExecutor.RemoveWorkflow(id)
			assert.NoError(t, err, "Failed to remove workflow %s", id)
		}

		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	}()

	// Wait for workflows to be running
	for _, id := range workflowIds {
		workflow, err := testdata.WaitForWorkflowStatus(id, []model.WorkflowStatus{model.RunningWorkflow}, testdata.WorkflowValidationTimeout)
		require.NoError(t, err, "Failed to get workflow %s", id)
		assert.Equal(t, model.RunningWorkflow, workflow.Status, "Workflow %s should be running", id)
	}

	// Test bulk terminate with options
	ctx := context.Background()

	opts := &client.WorkflowBulkResourceApiTerminateOpts{
		Reason:                 optional.NewString("Bulk termination test"),
		TriggerFailureWorkflow: optional.NewBool(false),
	}

	response, _, err := testdata.WorkflowBulkClient.Terminate(ctx, workflowIds, opts)
	require.NoError(t, err, "Failed to terminate workflows in bulk")

	// Verify bulk response
	assert.Equal(t, len(response.BulkErrorResults), 0, "Bulk error results should be empty")
	assert.Equal(t, len(response.BulkSuccessfulResults), len(workflowIds),
		"Bulk successful results should be equal to workflow count")

	// Verify workflows are terminated
	for _, id := range workflowIds {
		workflow, err := testdata.WaitForWorkflowStatus(id, []model.WorkflowStatus{model.TerminatedWorkflow}, testdata.WorkflowValidationTimeout)
		require.NoError(t, err, "Failed to get workflow %s", id)
		assert.Equal(t, model.TerminatedWorkflow, workflow.Status, "Workflow %s should be terminated", id)
	}
}
