package integration_tests

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/antihax/optional"
	"github.com/google/uuid"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const retryLimit = 5

func TestWorkflowCreation(t *testing.T) {
	executor := testdata.WorkflowExecutor

	wf := workflow.NewConductorWorkflow(executor).
		Name("PopulationMinMax").
		Version(1).
		Description("Simple Population Min Max workflow").
		Add(testdata.NewSetStateVariableTask(workflow.NewSimpleTask("set_state", "set_state")))
	err := wf.Register(true)
	if err != nil {
		t.Fatalf("Failed to register workflow: %s, reason: %s", wf.GetName(), err.Error())
	}

	workflow := testdata.NewKitchenSinkWorkflow(testdata.WorkflowExecutor)
	err = workflow.Register(true)
	require.NoError(t, err)

	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	startWorkers()
	run, err := executeWorkflowWithRetries(workflow, map[string]interface{}{
		"key1": "input1",
		"key2": 101,
	})
	require.NoError(t, err)

	assert.NotEmpty(t, run, "Workflow is null", run)
	workflowId := run.WorkflowId
	require.NotEmpty(t, workflowId)

	workflowUpdated, err := testdata.WaitForWorkflowCompletion(run.WorkflowId, testdata.ExtendedValidationTimeout)
	require.NoError(t, err)
	assert.Equal(t, model.CompletedWorkflow, workflowUpdated.Status)
	assert.Equal(t, "input1", run.Input["key1"])
}

func TestRemoveWorkflow(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	executor := testdata.WorkflowExecutor
	wf := workflow.NewConductorWorkflow(executor)
	wf.Name("temp_wf_" + strconv.Itoa(time.Now().Nanosecond())).Version(1)
	wf = wf.Add(workflow.NewSetVariableTask("set_var").Input("var_value", 42))
	err := wf.Register(true)

	assert.NoError(t, err, "Failed to register workflow")

	id, err := executor.StartWorkflow(&model.StartWorkflowRequest{Name: wf.GetName()})
	assert.NoError(t, err, "Failed to start workflow")

	execution, err := executor.GetWorkflow(id, true)
	assert.NoError(t, err, "Failed to get workflow execution")
	assert.Equal(t, model.CompletedWorkflow, execution.Status, "Workflow is not in the completed state")

	err = executor.RemoveWorkflow(id)
	assert.NoError(t, err, "Failed to remove workflow execution")

	_, err = executor.GetWorkflow(id, true)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no such workflow by Id")

	_, err = testdata.MetadataClient.UnregisterWorkflowDef(
		context.Background(),
		wf.GetName(),
		wf.GetVersion(),
	)
	assert.NoError(t, err, "Failed to delete workflow definition ", err)
}

func TestExecuteWorkflow(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	executor := testdata.WorkflowExecutor
	wf := workflow.NewConductorWorkflow(executor).
		Name("temp_wf_2_" + strconv.Itoa(time.Now().Nanosecond())).
		Version(1).
		OwnerEmail("test@orkes.io")
	wf = wf.Add(workflow.NewSetVariableTask("set_var").Input("var_value", 42))
	wf.OutputParameters(map[string]interface{}{
		"param1": "Test",
		"param2": 123,
	})
	err := wf.Register(true)

	assert.NoError(t, err, "Failed to register workflow")
	version := wf.GetVersion()
	run, err := executeWorkflowWithRetriesWithStartWorkflowRequest(
		&model.StartWorkflowRequest{
			Name:    wf.GetName(),
			Version: version,
		},
	)
	assert.NoError(t, err, "Failed to start workflow")
	assert.Equal(t, model.CompletedWorkflow, run.Status)

	execution, err := testdata.WaitForWorkflowStatus(run.WorkflowId, []model.WorkflowStatus{model.CompletedWorkflow}, 10*time.Second)
	assert.NoError(t, err, "Failed to get workflow execution")

	got := execution.WorkflowDefinition
	want := wf.ToWorkflowDef()
	require.NotNil(t, got)

	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Version, got.Version)
	if assert.Len(t, got.Tasks, len(want.Tasks)) {
		assert.Equal(t, want.Tasks[0].Name, got.Tasks[0].Name)
		assert.Equal(t, want.Tasks[0].TaskReferenceName, got.Tasks[0].TaskReferenceName)
		assert.Equal(t, want.Tasks[0].Type_, got.Tasks[0].Type_)
	}

	_, err = testdata.MetadataClient.UnregisterWorkflowDef(
		context.Background(),
		wf.GetName(),
		wf.GetVersion(),
	)
	assert.NoError(t, err, "Failed to delete workflow definition ", err)
}

func TestExecuteWorkflowWithCorrelationIds(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	executor := testdata.WorkflowExecutor
	correlationId1 := "correlationId1-" + uuid.New().String()
	correlationId2 := "correlationId2-" + uuid.New().String()

	httpTaskWorkflow1 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_HTTP" + correlationId1).
		OwnerEmail("test@orkes.io").
		Version(1).
		Add(testdata.TestHttpTask)
	httpTaskWorkflow2 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_HTTP" + correlationId2).
		OwnerEmail("test@orkes.io").
		Version(1).
		Add(testdata.TestHttpTask)
	_, err := httpTaskWorkflow1.StartWorkflow(&model.StartWorkflowRequest{CorrelationId: correlationId1})
	require.NoError(t, err)

	_, err = httpTaskWorkflow2.StartWorkflow(&model.StartWorkflowRequest{CorrelationId: correlationId2})
	require.NoError(t, err)

	time.Sleep(3 * time.Second)
	workflows, err := executor.GetByCorrelationIdsAndNames(true, true,
		[]string{correlationId1, correlationId2}, []string{httpTaskWorkflow1.GetName(), httpTaskWorkflow2.GetName()})
	require.NoError(t, err)
	require.Contains(t, workflows, correlationId1)
	assert.Contains(t, workflows, correlationId2)
	assert.NotEmpty(t, workflows[correlationId1])
	assert.NotEmpty(t, workflows[correlationId2])
	assert.Equal(t, workflows[correlationId1][0].CorrelationId, correlationId1)
	assert.Equal(t, workflows[correlationId2][0].CorrelationId, correlationId2)
}

func TestTerminateWorkflowWithFailure(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	executor := testdata.WorkflowExecutor
	wf := workflow.NewConductorWorkflow(executor).
		Name("TEST_GO_SET_VAR_USED_AS_FAILURE").
		Version(1).
		Add(workflow.NewSetVariableTask("set_var").Input("var_value", 42))
	err := testdata.ValidateWorkflowRegistration(wf)
	if err != nil {
		t.Fatal(err)
	}

	workflowWait := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_WAIT_CONDUCTOR").
		Version(1).
		Add(workflow.NewWaitTask("termination_wait")).
		FailureWorkflow(wf.GetName())
	err = testdata.ValidateWorkflowRegistration(workflowWait)
	if err != nil {
		t.Fatal(err)
	}

	id, err := workflowWait.StartWorkflow(&model.StartWorkflowRequest{})
	if err != nil {
		t.Fatal(err)
	}
	err = executor.TerminateWithFailure(id, "Terminated to trigger failure workflow", true)
	if err != nil {
		t.Fatal(err)
	}
	terminatedWfStatus, err := executor.GetWorkflow(id, false)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, terminatedWfStatus.Output["conductor.failure_workflow"])
}

func TestExecuteWorkflowSync(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	executor := testdata.WorkflowExecutor
	wf := workflow.NewConductorWorkflow(executor).
		Name("temp_wf_3_" + strconv.Itoa(time.Now().Nanosecond())).
		Version(1).
		OwnerEmail("test@orkes.io")
	wf = wf.Add(workflow.NewSetVariableTask("set_var").Input("var_value", 42))
	wf.OutputParameters(map[string]interface{}{
		"param1": "Test",
		"param2": 123,
	})
	err := wf.Register(true)

	assert.NoError(t, err, "Failed to register workflow")
	run, err := executeWorkflowWithRetries(wf, map[string]interface{}{
		"key1": "input1",
		"key2": 101,
	})
	if err != nil {
		t.Fatalf("Failed to complete the workflow, reason: %s", err)
	}
	assert.NotEmpty(t, run, "Workflow is null", run)
	assert.Equal(t, model.CompletedWorkflow, run.Status)

	execution, err := executor.GetWorkflow(run.WorkflowId, true)
	assert.NoError(t, err, "Failed to get workflow execution")
	assert.Equal(t, model.CompletedWorkflow, execution.Status, "Workflow is not in the completed state")

	_, err = testdata.MetadataClient.UnregisterWorkflowDef(
		context.Background(),
		wf.GetName(),
		wf.GetVersion(),
	)
	assert.NoError(t, err, "Failed to delete workflow definition ", err)
}

// TestTerminateWorkflow tests the Terminate API endpoint
func TestTerminateWorkflow(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a long-running workflow with a wait task
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_TERMINATE_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewWaitForDurationTask("wait_for_termination", 60*time.Second))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Start the workflow
	workflowId, err := wf.StartWorkflowWithInput(map[string]interface{}{"test": "terminate"})
	assert.NoError(t, err, "Failed to start workflow")

	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	})

	// Wait for workflow to be running
	runningWorkflow, err := testdata.WaitForWorkflowRunning(workflowId, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to be running")
	assert.Equal(t, model.RunningWorkflow, runningWorkflow.Status, "Workflow should be running")

	// Terminate the workflow
	opts := &client.WorkflowResourceApiTerminateOpts{
		Reason: optional.NewString("Test termination"),
	}
	_, err = testdata.WorkflowClient.Terminate(context.Background(), workflowId, opts)
	assert.NoError(t, err, "Failed to terminate workflow")

	// Verify workflow is terminated
	terminatedWorkflow, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.TerminatedWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow termination")
	assert.Equal(t, model.TerminatedWorkflow, terminatedWorkflow.Status, "Workflow should be terminated")
}

// TestRetryWorkflow tests the Retry API endpoint after workflow termination
func TestRetryWorkflow(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a workflow with a wait task that will run long enough for us to terminate it
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_RETRY_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewWaitForDurationTask("wait_for_retry", 5*time.Second))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Start the workflow
	workflowId, err := wf.StartWorkflowWithInput(map[string]interface{}{"test": "retry"})
	assert.NoError(t, err, "Failed to start workflow")

	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	})

	// Wait for workflow to be running
	runningWorkflow, err := testdata.WaitForWorkflowRunning(workflowId, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to be running")
	assert.Equal(t, model.RunningWorkflow, runningWorkflow.Status, "Workflow should be running")

	// Terminate the workflow
	terminateOpts := &client.WorkflowResourceApiTerminateOpts{
		Reason: optional.NewString("Terminating for retry test"),
	}
	_, err = testdata.WorkflowClient.Terminate(context.Background(), workflowId, terminateOpts)
	assert.NoError(t, err, "Failed to terminate workflow")

	// Verify workflow is terminated
	terminatedWorkflow, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.TerminatedWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow termination")
	assert.Equal(t, model.TerminatedWorkflow, terminatedWorkflow.Status, "Workflow should be terminated")

	// Retry the workflow with all available options
	retryOpts := &client.WorkflowResourceApiRetryOpts{
		ResumeSubworkflowTasks: optional.NewBool(false),
		RetryIfRetriedByParent: optional.NewBool(true),
	}
	_, err = testdata.WorkflowClient.Retry(context.Background(), workflowId, retryOpts)
	assert.NoError(t, err, "Failed to retry workflow")

	// Verify workflow is running again after retry
	retriedWorkflow, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.RunningWorkflow, model.CompletedWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to get workflow after retry")
	assert.Contains(t, []model.WorkflowStatus{model.RunningWorkflow, model.CompletedWorkflow}, retriedWorkflow.Status, "Workflow should be running after retry")
}

// TestRerunWorkflow tests the Rerun API endpoint
func TestRerunWorkflow(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a simple workflow with multiple tasks
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_RERUN_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var1").Input("var_value", "first")).
		Add(workflow.NewSetVariableTask("set_var2").Input("var_value", "second"))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Start the workflow manually to get workflow ID for rerun
	workflowId, err := wf.StartWorkflowWithInput(map[string]interface{}{"test": "rerun"})
	assert.NoError(t, err, "Failed to start workflow for rerun test")

	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	})

	// Wait for workflow completion
	completedWorkflow, err := testdata.WaitForWorkflowCompletion(workflowId, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow completion")
	assert.Equal(t, model.CompletedWorkflow, completedWorkflow.Status, "Workflow should be completed")
	assert.Len(t, completedWorkflow.Tasks, 2, "Workflow should have 2 tasks")

	// Rerun from the second task with comprehensive parameters
	rerunRequest := model.RerunWorkflowRequest{
		ReRunFromTaskId: completedWorkflow.Tasks[1].TaskId,
		TaskInput: map[string]interface{}{
			"var_value": "rerun_value",
		},
		WorkflowInput: map[string]interface{}{
			"rerun_test": "true",
			"timestamp":  time.Now().Unix(),
		},
		CorrelationId: "rerun-correlation-" + uniqueSuffix,
	}

	rerunId, _, err := testdata.WorkflowClient.Rerun(context.Background(), rerunRequest, workflowId)
	assert.NoError(t, err, "Failed to rerun workflow")
	assert.Equal(t, workflowId, rerunId, "Rerun should return the same workflow ID")

	// Wait for rerun completion
	rerunWorkflow, err := testdata.WaitForWorkflowCompletion(rerunId, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for rerun workflow completion")
	assert.Equal(t, model.CompletedWorkflow, rerunWorkflow.Status, "Rerun workflow should be completed")

	var set1Count, set2Count int32
	var latestSetVar2 *model.Task
	var maxSeq int32
	for _, tk := range rerunWorkflow.Tasks {
		switch tk.ReferenceTaskName {
		case "set_var1":
			set1Count++
		case "set_var2":
			set2Count++
			if tk.Seq > maxSeq {
				maxSeq = tk.Seq
				c := tk // copy loop var
				latestSetVar2 = &c
			}
		}
	}

	// Assert upstream not re-executed; downstream re-executed
	assert.Equal(t, int32(1), set1Count, "set_var1 (upstream) must NOT be re-executed on rerun-from set_var2")
	assert.Equal(t, int32(2), set2Count, "set_var2 (downstream) must be executed twice (initial + rerun)")
	require.NotNil(t, latestSetVar2, "expected to capture latest set_var2 instance")
	assert.Equal(t, model.CompletedTask, latestSetVar2.Status, "latest set_var2 must complete")

	val, ok := latestSetVar2.InputData["var_value"]
	assert.True(t, ok, "latest set_var2 must have var_value in InputData")
	assert.Equal(t, "rerun_value", val)

}

func TestRetryWorkflowWithDecide(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	workflowName := "TEST_GO_FAILING_WORKFLOW_RETRY_DECIDE_" + uniqueSuffix

	// Create a workflow that will fail quickly
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name(workflowName).
		Version(1).
		Add(workflow.NewInlineTask("failing_inline_task",
			"function() { throw new Error('test error'); }();"))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register failing workflow")

	// Start the workflow and wait for it to fail
	workflowId, err := wf.StartWorkflowWithInput(map[string]interface{}{"test": "retry_decide"})
	assert.NoError(t, err, "Failed to start workflow")

	defer func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	}()

	// Wait for workflow to fail
	failedWorkflow, err := testdata.WaitForWorkflowStatus(workflowId,
		[]model.WorkflowStatus{model.FailedWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to fail")
	assert.Equal(t, model.FailedWorkflow, failedWorkflow.Status, "Workflow should be failed")

	_, err = testdata.WorkflowClient.Retry(context.Background(), workflowId, nil)
	assert.NoError(t, err, "Failed to retry workflow")

	// Verify workflow is running after retry
	workflowAfterRetry, err := testdata.WaitForWorkflowStatus(workflowId,
		[]model.WorkflowStatus{model.RunningWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to get workflow after retry")
	assert.Equal(t, model.RunningWorkflow, workflowAfterRetry.Status, "Workflow should be running after retry")
	assert.Equal(t, workflowId, workflowAfterRetry.WorkflowId, "Workflow ID should be the same")

	_, err = testdata.WorkflowClient.Decide(context.Background(), workflowId)
	assert.NoError(t, err, "Failed to trigger decide")

	workflowAfterDecide, err := testdata.WaitForWorkflowStatus(workflowId,
		[]model.WorkflowStatus{model.FailedWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to fail")

	// After decide, the workflow should be properly failed
	assert.Equal(t, model.FailedWorkflow, workflowAfterDecide.Status,
		"Workflow should be failed after decide unsticks it")
}

// TestRestartWorkflow tests the Restart API endpoint
func TestRestartWorkflow(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a simple workflow
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	// v1: single task
	wfV1 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_RESTART_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var").Input("var_value", "initial"))

	// v2: add an extra task so we can observe a change when useLatestDefinitions=true
	wfV2 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_RESTART_" + uniqueSuffix).
		Version(2).
		Add(workflow.NewSetVariableTask("set_var").Input("var_value", "v2")).
		Add(workflow.NewSetVariableTask("set_var_extra").Input("var_value", "extra"))

	require.NoError(t, testdata.ValidateWorkflowRegistration(wfV1), "register v1")
	require.NoError(t, testdata.ValidateWorkflowRegistration(wfV2), "register v2")

	// Start and complete v1
	workflowId, err := wfV1.StartWorkflowWithInput(map[string]interface{}{"test": "restart"})
	require.NoError(t, err, "start workflow")

	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
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
	})

	// Wait for workflow completion
	completedWorkflow, err := testdata.WaitForWorkflowCompletion(workflowId, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow completion")
	assert.Equal(t, model.CompletedWorkflow, completedWorkflow.Status, "Workflow should be completed")

	var initialSetVarTaskID string
	var initialSetVarSeq int32
	for _, tk := range completedWorkflow.Tasks {
		if tk.ReferenceTaskName == "set_var" {
			initialSetVarTaskID = tk.TaskId
			initialSetVarSeq = tk.Seq
			break
		}
	}
	require.NotEmpty(t, initialSetVarTaskID, "expected set_var in initial run")

	// Restart the workflow with options
	opts := &client.WorkflowResourceApiRestartOpts{
		UseLatestDefinitions: optional.NewBool(true),
	}
	_, err = testdata.WorkflowClient.Restart(context.Background(), workflowId, opts)
	assert.NoError(t, err, "Failed to restart workflow")

	// Wait for restarted workflow completion
	restartedWorkflow, err := testdata.WaitForWorkflowCompletion(workflowId, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for restarted workflow completion")
	assert.Equal(t, model.CompletedWorkflow, restartedWorkflow.Status, "Restarted workflow should be completed")

	var sawNewSetVar bool
	var sawSetVarExtra bool
	for _, tk := range restartedWorkflow.Tasks {
		if tk.ReferenceTaskName == "set_var" {
			if tk.TaskId != initialSetVarTaskID || tk.Seq > initialSetVarSeq {
				sawNewSetVar = true
			}
		}
		if tk.ReferenceTaskName == "set_var_extra" {
			sawSetVarExtra = true
		}
	}
	require.True(t, sawNewSetVar, "set_var must re-execute after restart")
	require.True(t, sawSetVarExtra, "v2 task set_var_extra should appear when useLatestDefinitions=true")
}

func TestUpgradeRunningWorkflowToVersion(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create version 1 of the workflow - use a WAIT task that waits indefinitely until signaled
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	workflowName := "TEST_GO_WORKFLOW_UPGRADE_" + uniqueSuffix

	wfV1 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name(workflowName).
		Version(1).
		Add(workflow.NewWaitForDurationTask("wait_for_upgrade", 3*time.Second)) // Task that will stay scheduled

	preV2 := workflow.NewSetVariableTask("pre_v2").Input("marker", "introduced_in_v2")
	postV2 := workflow.NewSetVariableTask("post_v2").Input("done", true)

	wfV2 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name(workflowName).
		Version(2).
		Add(preV2).
		Add(workflow.NewWaitForDurationTask("wait_for_upgrade", 3*time.Second)).
		Add(postV2)

	require.NoError(t, testdata.ValidateWorkflowRegistration(wfV1), "register v1")
	require.NoError(t, testdata.ValidateWorkflowRegistration(wfV2), "register v2")

	// Start the workflow explicitly with version 1
	startRequest := &model.StartWorkflowRequest{
		Name:    workflowName,
		Version: 1,
		Input:   map[string]interface{}{"upgrade_test": "value"}, // Explicitly specify version 1
	}

	workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
	assert.NoError(t, err, "Failed to start workflow")

	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wfV1.GetName(),
			wfV1.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition v1")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wfV2.GetName(),
			wfV2.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition v2")
	})

	// Wait for workflow to be running with version 1
	workflowAfterRunning, err := testdata.WaitForWorkflowStatus(workflowId,
		[]model.WorkflowStatus{model.RunningWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to run")
	assert.Equal(t, model.RunningWorkflow, workflowAfterRunning.Status, "Workflow should be running")
	assert.Equal(t, "value", workflowAfterRunning.Input["upgrade_test"], "Workflow input should be value")

	// Now upgrade the running workflow
	upgradeRequest := model.UpgradeWorkflowRequest{
		Name:    workflowName,
		Version: 2,
		WorkflowInput: map[string]interface{}{
			"upgrade_test_upgrade": "value_upgrade",
		},
	}
	_, err = testdata.WorkflowClient.UpgradeRunningWorkflowToVersion(
		context.Background(),
		upgradeRequest,
		workflowId,
	)
	assert.NoError(t, err, "Failed to upgrade workflow")

	// Verify the workflow has been upgraded and completes with version 2
	upgradedWorkflow, err := testdata.WaitForWorkflowStatus(workflowId,
		[]model.WorkflowStatus{model.CompletedWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to get upgraded workflow")

	assert.Equal(t, "value_upgrade", upgradedWorkflow.Input["upgrade_test_upgrade"], "Workflow input should be value_upgrade")

	var preFound, postFound bool
	for _, tsk := range upgradedWorkflow.Tasks {
		switch tsk.ReferenceTaskName {
		case "pre_v2":
			preFound = true
			assert.Equal(t, "SKIPPED", string(tsk.Status), "pre_v2 must be SKIPPED after upgrade")
		case "post_v2":
			postFound = true
			if upgradedWorkflow.Status == model.CompletedWorkflow {
				assert.Equal(t, "COMPLETED", string(tsk.Status), "post_v2 should complete on v2 path")
			}
		}
	}
	require.True(t, preFound, "v2 graph must contain pre_v2")
	require.True(t, postFound, "v2 graph must contain post_v2")
}

// TestStartWorkflowWithRequest tests the StartWorkflowWithRequest API endpoint
func TestStartWorkflowWithRequest(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a simple workflow
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_START_REQUEST_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var").Input("var_value", "start_request_test"))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Create a comprehensive StartWorkflowRequest
	startRequest := model.StartWorkflowRequest{
		Name:    wf.GetName(),
		Version: 1,
		Input: map[string]interface{}{
			"test_key": "test_value",
			"number":   42,
			"boolean":  true,
		},
		CorrelationId: "test-correlation-" + uniqueSuffix,
		Priority:      5,
		TaskToDomain: map[string]string{
			"SET_VARIABLE": "test-domain",
		},
		ExternalInputPayloadStoragePath: "",
	}

	// Test StartWorkflowWithRequest
	workflowId, _, err := testdata.WorkflowClient.StartWorkflowWithRequest(
		context.Background(),
		startRequest,
	)

	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	})

	assert.NoError(t, err, "Failed to start workflow with request")
	assert.NotEmpty(t, workflowId, "Workflow ID should not be empty")

	completedWorkflow, err := testdata.WaitForWorkflowCompletion(workflowId, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow completion")
	assert.Equal(t, model.CompletedWorkflow, completedWorkflow.Status, "Workflow should be completed")

	// Verify workflow properties
	assert.Equal(t, startRequest.CorrelationId, completedWorkflow.CorrelationId, "Correlation ID should match")
	assert.Equal(t, int32(startRequest.Priority), completedWorkflow.Priority, "Priority should match")
}

func startWorkers() {
	testdata.TaskRunner.StartWorker("simple_task", testdata.SimpleWorker, 10, 100*time.Millisecond)
	testdata.TaskRunner.StartWorker("dynamic_fork_prep", testdata.DynamicForkWorker, 3, 100*time.Millisecond)
}

func executeWorkflowWithRetries(wf *workflow.ConductorWorkflow, workflowInput interface{}) (*model.WorkflowRun, error) {
	for attempt := 0; attempt < retryLimit; attempt += 1 {
		workflowRun, err := wf.ExecuteWorkflowWithInput(workflowInput, "")
		if err != nil {
			time.Sleep(time.Duration(attempt+2) * time.Second)
			fmt.Println("Failed to execute workflow, reason: " + err.Error())
			continue
		}
		return workflowRun, nil
	}
	return nil, fmt.Errorf("exhausted retries for workflow execution")
}

func executeWorkflowWithRetriesWithStartWorkflowRequest(startWorkflowRequest *model.StartWorkflowRequest) (*model.WorkflowRun, error) {
	for attempt := 1; attempt <= retryLimit; attempt += 1 {
		workflowRun, err := testdata.WorkflowExecutor.ExecuteWorkflow(startWorkflowRequest, "")
		if err != nil {
			time.Sleep(time.Duration(attempt+2) * time.Second)
			fmt.Printf("Failed to execute workflow, reason: %s", err.Error())
			continue
		}
		return workflowRun, nil
	}
	return nil, fmt.Errorf("exhausted retries for workflow execution")
}

// TestGetWorkflowState tests the GetWorkflowState API endpoint
func TestGetWorkflowState(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a workflow with variables
	uniqueSuffix := fmt.Sprintf("%d-%s", time.Now().UnixNano(), uuid.New().String())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_STATE_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var1").Input("var_value", "state_test_1")).
		Add(workflow.NewSetVariableTask("set_var2").Input("var_value", "state_test_2"))

	wf.OutputParameters(map[string]interface{}{
		"output_param1": "value1",
		"output_param2": "value2",
	})

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Start the workflow
	startRequest := &model.StartWorkflowRequest{
		Name:    wf.GetName(),
		Version: 1,
		Input:   map[string]interface{}{"state_test_input": "input_value"},
	}

	workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
	assert.NoError(t, err, "Failed to start workflow")

	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	})

	completedWorkflow, err := testdata.WaitForWorkflowCompletion(workflowId, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow completion")
	assert.Equal(t, model.CompletedWorkflow, completedWorkflow.Status, "Workflow should be completed")

	// Test GetWorkflowState with output and variables
	workflowState, _, err := testdata.WorkflowClient.GetWorkflowState(
		context.Background(),
		workflowId,
		true, // includeOutput
		true, // includeVariables
	)
	assert.NoError(t, err, "Failed to get workflow state")
	assert.NotNil(t, workflowState, "Workflow state should not be nil")

	// Verify workflow state contains expected data
	assert.Equal(t, workflowId, workflowState.WorkflowId, "Workflow ID should match")
	assert.Equal(t, string(model.CompletedWorkflow), workflowState.Status, "Workflow status should match")
	assert.NotNil(t, workflowState.Output, "Output should not be nil")
	assert.NotNil(t, workflowState.Variables, "Variables should not be nil")

	// Check output parameters
	assert.Equal(t, "value1", workflowState.Output["output_param1"], "Output should match")
	assert.Equal(t, "value2", workflowState.Output["output_param2"], "Output should match")

	assert.Equal(t, "state_test_2", workflowState.Variables["var_value"], "Variable value should match the last set value")
}

// TestGetRunningWorkflow tests the GetRunningWorkflow API endpoint
func TestGetRunningWorkflow(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	// Create a long-running workflow
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_RUNNING_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewWaitForDurationTask("wait_for_running", 30*time.Second))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Start multiple instances of the workflow
	numWorkflows := 3
	workflowIds := make([]string, numWorkflows)

	for i := 0; i < numWorkflows; i++ {
		startRequest := &model.StartWorkflowRequest{
			Name:    wf.GetName(),
			Version: 1,
			Input:   map[string]interface{}{"instance": i, "running_test": "multiple"},
		}

		workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
		require.NoError(t, err, "Failed to start workflow instance %d", i)
		workflowIds[i] = workflowId
	}

	// Setup cleanup
	t.Cleanup(func() {
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
	})

	err = testdata.WaitForMultipleWorkflowsStatus(workflowIds, []model.WorkflowStatus{model.RunningWorkflow}, testdata.ExtendedValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflows to be running")

	// Test GetRunningWorkflow
	opts := &client.WorkflowResourceApiGetRunningWorkflowOpts{
		Version:   optional.NewInt32(1),
		StartTime: optional.NewInt64(time.Now().Add(-1*time.Hour).Unix() * 1000), // 1 hour ago
		EndTime:   optional.NewInt64(time.Now().Add(1*time.Hour).Unix() * 1000),  // 1 hour from now
	}

	var (
		runningWorkflows []string
	)
	err = testdata.RetryCondition(3, 1*time.Second,
		func() error {
			runningWorkflows, _, err = testdata.WorkflowClient.GetRunningWorkflow(context.Background(), wf.GetName(), opts)
			return err
		},
		func() bool {
			return len(runningWorkflows) == numWorkflows
		})

	assert.NoError(t, err, "Failed to get running workflows")
	assert.NotNil(t, runningWorkflows, "Running workflows result should not be nil")

	// Verify all our workflows are in the running workflows list
	foundCount := 0
	for _, id := range workflowIds {
		for _, runningWf := range runningWorkflows {
			if runningWf == id {
				foundCount++
				break
			}
		}
	}

	assert.Equal(t, numWorkflows, foundCount, "All %d workflows should be in the running workflows list", numWorkflows)

	// Verify the individual workflows are running
	for i, id := range runningWorkflows {
		workflow, err := testdata.WorkflowExecutor.GetWorkflow(id, false)
		assert.NoError(t, err, "Failed to get workflow %d", i)
		assert.Equal(t, model.RunningWorkflow, workflow.Status, "Workflow %d should be running", i)
	}
}

// TestGetWorkflowsBatch tests the GetWorkflowsBatch API endpoint
func TestGetWorkflowsBatch(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create multiple correlation IDs and workflows
	correlationIds := make([]string, 3)
	workflowIds := make([]string, 3)
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())

	// Create a simple workflow
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_BATCH_GET_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var").Input("var_value", "batch_get_test"))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	for i := 0; i < 3; i++ {
		correlationIds[i] = "test-batch-get-correlation-" + strconv.Itoa(i) + "-" + uuid.New().String()

		startRequest := &model.StartWorkflowRequest{
			Name:          wf.GetName(),
			Version:       1,
			CorrelationId: correlationIds[i],
			Input: map[string]interface{}{
				"index": i,
				"test":  "batch_get",
			},
		}

		workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
		assert.NoError(t, err, "Failed to start workflow %d", i)
		workflowIds[i] = workflowId
	}

	// Setup cleanup
	t.Cleanup(func() {
		for _, id := range workflowIds {
			err = testdata.WorkflowExecutor.RemoveWorkflow(id)
			if err != nil {
				t.Logf("Warning: Failed to remove workflow %s: %v", id, err)
			}
		}

		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		if err != nil {
			t.Logf("Warning: Failed to remove workflow definition: %v", err)
		}
	})
	err = testdata.WaitForMultipleWorkflowsCompletion(workflowIds, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow completion")

	// Test GetWorkflowsBatch - create a map with specific keys expected by the API
	// The API expects a map with "workflowNames" and "correlationIds" as keys
	batchRequest := make(map[string][]string)
	batchRequest["workflowNames"] = []string{wf.GetName()}
	batchRequest["correlationIds"] = correlationIds

	opts := &client.WorkflowResourceApiGetWorkflowsOpts{
		IncludeClosed: optional.NewBool(true),
		IncludeTasks:  optional.NewBool(false),
	}
	var (
		workflowsMap map[string][]model.Workflow
	)
	err = testdata.RetryCondition(3, 1*time.Second,
		func() error {
			workflowsMap, _, err = testdata.WorkflowClient.GetWorkflowsBatch(
				context.Background(),
				batchRequest,
				opts,
			)
			return err
		},
		func() bool {
			return len(workflowsMap) == len(correlationIds)

		})
	assert.NoError(t, err, "Failed to get workflows batch")
	assert.NotNil(t, workflowsMap, "Workflows map should not be nil")

	// The result should contain correlation IDs as keys
	// Create a map to track which correlation IDs we've found
	foundCorrelationIds := make(map[string]bool)
	allWorkflows := []model.Workflow{}

	for corrId, workflows := range workflowsMap {
		foundCorrelationIds[corrId] = true

		// Verify each workflow has the correct status and correlation ID
		for _, workflow := range workflows {
			assert.Equal(t, corrId, workflow.CorrelationId, "Workflow should have correct correlation ID")
			assert.Equal(t, model.CompletedWorkflow, workflow.Status, "Workflow should be completed")
			allWorkflows = append(allWorkflows, workflow)
		}
	}

	// Verify all correlation IDs were found
	for _, corrId := range correlationIds {
		assert.True(t, foundCorrelationIds[corrId], "Should find workflow with correlation ID %s", corrId)
	}

	// Verify we found the right number of workflows
	assert.Len(t, allWorkflows, len(correlationIds), "Should find %d workflows", len(correlationIds))

	// Test with a subset of correlation IDs
	// Create a request with only the first correlation ID
	subsetRequest := make(map[string][]string)
	subsetRequest["workflowNames"] = []string{wf.GetName()}
	subsetRequest["correlationIds"] = []string{correlationIds[0]}

	subsetResult, _, err := testdata.WorkflowClient.GetWorkflowsBatch(
		context.Background(),
		subsetRequest,
		opts,
	)
	assert.NoError(t, err, "Failed to get workflows batch with subset request")
	assert.NotNil(t, subsetResult, "Subset result should not be nil")

	// Verify the subset result contains only one workflow with the first correlation ID
	subsetWorkflows, exists := subsetResult[correlationIds[0]]
	assert.True(t, exists, "Should find workflows for correlation ID %s", correlationIds[0])
	assert.Len(t, subsetWorkflows, 1, "Should find 1 workflow in subset")
}

// TestGetWorkflowsByCorrelationId tests the GetWorkflowsByCorrelationId API endpoint
func TestGetWorkflowsByCorrelationId(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a unique correlation ID
	correlationId := "test-correlation-" + uuid.New().String()
	uniqueSuffix := fmt.Sprintf("%d-%s", time.Now().UnixNano(), uuid.New().String())

	// Create a simple workflow
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_CORRELATION_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var").Input("var_value", "correlation_test"))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Start multiple workflows with the same correlation ID
	workflowIds := make([]string, 3)
	for i := 0; i < 3; i++ {
		startRequest := &model.StartWorkflowRequest{
			Name:          wf.GetName(),
			Version:       1,
			CorrelationId: correlationId,
			Input: map[string]interface{}{
				"instance": i,
				"test":     "correlation",
			},
		}

		workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
		assert.NoError(t, err, "Failed to start workflow %d", i)

		workflowIds[i] = workflowId
	}

	// Wait for workflows to complete
	err = testdata.WaitForMultipleWorkflowsCompletion(workflowIds, 30*time.Second)
	assert.NoError(t, err, "Failed to wait for workflows completion")

	// Test GetWorkflowsByCorrelationId
	opts := &client.WorkflowResourceApiGetWorkflowsOpts{
		IncludeClosed: optional.NewBool(true),
		IncludeTasks:  optional.NewBool(false),
	}

	var (
		workflows []model.Workflow
	)
	err = testdata.RetryCondition(3, 1*time.Second,
		func() error {
			workflows, _, err = testdata.WorkflowClient.GetWorkflowsByCorrelationId(
				context.Background(),
				wf.GetName(),
				correlationId,
				opts,
			)
			return err
		},
		func() bool {
			return len(workflows) == len(workflowIds)

		})
	assert.NoError(t, err, "Failed to get workflows by correlation ID")
	assert.NotNil(t, workflows, "Workflows should not be nil")
	assert.Len(t, workflows, 3, "Should find 3 workflows with the correlation ID")

	// Verify all workflows have the correct correlation ID
	for _, wf := range workflows {
		assert.Equal(t, correlationId, wf.CorrelationId, "Workflow should have the correct correlation ID")
		assert.Equal(t, model.CompletedWorkflow, wf.Status, "Workflow should be completed")
	}

	// Test with includeTasks = true
	optsWithTasks := &client.WorkflowResourceApiGetWorkflowsOpts{
		IncludeClosed: optional.NewBool(true),
		IncludeTasks:  optional.NewBool(true),
	}

	workflowsWithTasks, _, err := testdata.WorkflowClient.GetWorkflowsByCorrelationId(
		context.Background(),
		wf.GetName(),
		correlationId,
		optsWithTasks,
	)
	assert.NoError(t, err, "Failed to get workflows with tasks by correlation ID")
	assert.NotNil(t, workflowsWithTasks, "Workflows with tasks should not be nil")
	assert.Len(t, workflowsWithTasks, 3, "Should find 3 workflows with tasks")

	// Verify tasks are included
	for _, wf := range workflowsWithTasks {
		assert.True(t, len(wf.Tasks) > 0, "Workflow should have tasks")
	}

	t.Cleanup(func() {
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
	})
}

// TestPauseResumeWorkflow tests the Pause and Resume Workflow APIs
func TestPauseResumeWorkflow(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a long-running workflow with a wait task
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_PAUSE_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewWaitForDurationTask("wait_for_pause", 60*time.Second))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Start the workflow
	workflowId, err := wf.StartWorkflowWithInput(map[string]interface{}{"test": "pause"})
	assert.NoError(t, err, "Failed to start workflow")

	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	})

	// Wait for workflow to be running
	runningWorkflow, err := testdata.WaitForWorkflowRunning(workflowId, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to be running")
	assert.Equal(t, model.RunningWorkflow, runningWorkflow.Status, "Workflow should be running")

	// Pause the workflow
	_, err = testdata.WorkflowClient.PauseWorkflow(context.Background(), workflowId)
	assert.NoError(t, err, "Failed to pause workflow")

	// Verify workflow is paused
	pausedWorkflow, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.PausedWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to be paused")
	assert.Equal(t, model.PausedWorkflow, pausedWorkflow.Status, "Workflow should be paused")

	// Resume the workflow to test the full cycle
	_, err = testdata.WorkflowClient.ResumeWorkflow(context.Background(), workflowId)
	assert.NoError(t, err, "Failed to resume workflow")

	// Verify workflow is running again after resume
	resumedWorkflow, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.RunningWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to resume")
	assert.Equal(t, model.RunningWorkflow, resumedWorkflow.Status, "Workflow should be running after resume")
}

// TestResetWorkflowCallbackTime tests the ResetWorkflow API endpoint with tasks that have non-zero CallbackAfterSeconds
// This test manually updates a task to have CallbackAfterSeconds > 0, then verifies reset functionality
func TestResetWorkflowCallbackTime(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	taskName := "RESET_CALLBACK_TASK_" + uniqueSuffix

	// Create a simple workflow
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_RESET_CALLBACK_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSimpleTask("reset_callback_ref", taskName))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Register task definition
	taskDef := model.TaskDef{
		Name:                   taskName,
		TimeoutSeconds:         300,
		RetryCount:             0,
		TimeoutPolicy:          "TIME_OUT_WF",
		RetryLogic:             "FIXED",
		RetryDelaySeconds:      1,
		ResponseTimeoutSeconds: 300,
	}
	err = testdata.ValidateTaskRegistration(taskDef)
	assert.NoError(t, err, "Failed to register task")

	// We will manually update the task using TaskClient APIs

	// Start the workflow
	startRequest := &model.StartWorkflowRequest{
		Name:    wf.GetName(),
		Version: 1,
		Input:   map[string]interface{}{"callback_test": "value"},
	}

	workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
	assert.NoError(t, err, "Failed to start workflow")

	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			wf.GetName(),
			wf.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	})

	// Wait for the workflow to start and task to be scheduled
	workflowWithTasks, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.RunningWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to start")

	// Find the scheduled task
	var scheduledTask *model.Task
	for _, task := range workflowWithTasks.Tasks {
		if task.ReferenceTaskName == taskName {
			scheduledTask = &task
			break
		}
	}
	assert.NotNil(t, scheduledTask, "Should find the scheduled task")
	assert.Equal(t, scheduledTask.CallbackAfterSeconds, int64(0), "Task should have non-zero CallbackAfterSeconds")

	// Manually update the task to IN_PROGRESS status with CallbackAfterSeconds = 45
	taskResult := &model.TaskResult{
		TaskId:               scheduledTask.TaskId,
		WorkflowInstanceId:   scheduledTask.WorkflowInstanceId,
		Status:               model.InProgressTask,
		CallbackAfterSeconds: 45, // Set callback after 45 seconds
		OutputData: map[string]interface{}{
			"message": "Manually set to IN_PROGRESS with callback",
		},
	}

	_, _, err = testdata.TaskClient.UpdateTask(context.Background(), taskResult)
	assert.NoError(t, err, "Failed to update task to IN_PROGRESS")

	workflowWithTasks, err = testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.RunningWorkflow, model.CompletedWorkflow}, testdata.ExtendedValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to be running or completed")

	// Find the updated task
	var taskBeforeReset *model.Task
	for _, task := range workflowWithTasks.Tasks {
		if task.ReferenceTaskName == taskName {
			taskBeforeReset = &task
			break
		}
	}
	assert.NotNil(t, taskBeforeReset, "Should find the updated task")

	// Capture the original CallbackAfterSeconds
	originalCallbackAfterSeconds := taskBeforeReset.CallbackAfterSeconds

	// Verify we got non-zero CallbackAfterSeconds (should be 45)
	assert.Equal(t, originalCallbackAfterSeconds, int64(45),
		"Task should have non-zero CallbackAfterSeconds after manual update, got: %d", originalCallbackAfterSeconds)

	// Call ResetWorkflow - this should reset callback times for non-terminal tasks
	_, err = testdata.WorkflowClient.ResetWorkflow(context.Background(), workflowId)
	assert.NoError(t, err, "Failed to reset workflow")

	// Get workflow after reset
	resetWorkflow, err := testdata.WorkflowExecutor.GetWorkflow(workflowId, true)
	assert.NoError(t, err, "Failed to get workflow after reset")

	// Find the task after reset
	var taskAfterReset *model.Task
	for _, task := range resetWorkflow.Tasks {
		if task.ReferenceTaskName == taskName {
			taskAfterReset = &task
			break
		}
	}
	assert.NotNil(t, taskAfterReset, "Task should still exist after reset")

	// MAIN VERIFICATION: Check that CallbackAfterSeconds was reset to 0
	assert.Equal(t, int64(0), taskAfterReset.CallbackAfterSeconds,
		"CallbackAfterSeconds should be reset to 0, but got: %d", taskAfterReset.CallbackAfterSeconds)

	// Verify that the callback time was reset from 45 to 0
	assert.NotEqual(t, originalCallbackAfterSeconds, taskAfterReset.CallbackAfterSeconds,
		"CallbackAfterSeconds should have changed from %d to 0", originalCallbackAfterSeconds)

	// The task should remain in its current state (status doesn't change with reset)
	assert.True(t,
		taskAfterReset.Status == model.InProgressTask || taskAfterReset.Status == model.ScheduledTask,
		"Task should be in IN_PROGRESS or SCHEDULED state after reset, got: %s", taskAfterReset.Status)
}

func TestWorkflowTest(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())

	httpTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_TEST_" + uniqueSuffix).
		Version(1).
		Add(testdata.TestHttpTask)

	err := testdata.ValidateWorkflowRegistration(httpTaskWorkflow)
	assert.NoError(t, err, "Failed to register workflow")

	t.Cleanup(func() {
		_, err := testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			httpTaskWorkflow.GetName(),
			httpTaskWorkflow.GetVersion(),
		)
		assert.NoError(t, err, "Failed to unregister workflow definition")
	})

	// Create a task mock for the HTTP task to simulate its output
	taskMocks := make(map[string][]model.TaskMock)
	taskMocks[testdata.TestHttpTask.ReferenceName()] = []model.TaskMock{
		{
			Status: "COMPLETED",
			Output: map[string]interface{}{
				"response": map[string]interface{}{
					"body": map[string]interface{}{
						"testKey": "testValue",
					},
					"statusCode": 200,
				},
			},
			QueueWaitTime: 5, // 5 milliseconds
		},
	}

	// Prepare the workflow test request
	testRequest := model.WorkflowTestRequest{
		Name:                httpTaskWorkflow.GetName(),
		Version:             httpTaskWorkflow.GetVersion(),
		Input:               map[string]interface{}{"inputParam1": "testValue1"},
		TaskRefToMockOutput: taskMocks,
		WorkflowDef:         httpTaskWorkflow.ToWorkflowDef(),
	}

	// Call the test workflow API
	workflowResult, _, err := testdata.WorkflowClient.TestWorkflow(context.Background(), testRequest)
	// Validate response
	assert.NoError(t, err, "Failed to test workflow")
	// Validate workflow result
	assert.Equal(t, model.CompletedWorkflow, workflowResult.Status, "Expected workflow status COMPLETED, got %s", workflowResult.Status)
	// Validate tasks in workflow result
	assert.Equal(t, 1, len(workflowResult.Tasks), "Expected 1 task in workflow result, got %d", len(workflowResult.Tasks))
	// Validate http task output
	httpTask := workflowResult.Tasks[0]
	assert.Equal(t, model.CompletedTask, httpTask.Status, "Expected HTTP task status COMPLETED, got %s", httpTask.Status)

}

func TestJumpToTask(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV52)

	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())

	input := workflow.HttpInput{
		Method: "GET",
		Uri:    "http://httpbin:8081/api/hello?name=Test123",
	}
	workflowDef := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_JUMP_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewWaitForDurationTask("wait_ref_1", 5*time.Second)).
		Add(workflow.NewHttpTask("http_ref_1", &input))

	err := testdata.ValidateWorkflowRegistration(workflowDef)
	assert.NoError(t, err, "Failed to register workflow")

	// Start a workflow instance
	startRequest := &model.StartWorkflowRequest{
		Name:    workflowDef.GetName(),
		Version: 1,
		Input:   map[string]interface{}{"testKey": "testValue"},
	}

	workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
	assert.NoError(t, err, "Failed to start workflow")

	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), workflowDef.GetName(), workflowDef.GetVersion())
		assert.NoError(t, err, "Failed to remove workflow definition")
	})

	runningWorkflow, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.RunningWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to start")
	assert.Equal(t, model.RunningWorkflow, runningWorkflow.Status, "Workflow should be running")

	// Jump to the HTTP task, skipping the wait task
	jumpInput := map[string]interface{}{
		"uri":    "http://httpbin:8081/api/hello?name=JumpTest",
		"method": "GET",
	}

	opts := &client.WorkflowResourceApiJumpToTaskOpts{
		TaskReferenceName: optional.NewString("http_ref_1"),
	}

	_, err = testdata.WorkflowClient.JumpToTask(
		context.Background(),
		jumpInput,
		workflowId,
		opts,
	)

	assert.NoError(t, err, "Failed to jump to task")

	// Wait for the jump to be processed and workflow to continue
	updatedWorkflow, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.CompletedWorkflow}, 30*time.Second)
	assert.NoError(t, err, "Failed to wait for workflow after jump")

	// Verify that the wait task was skipped and the HTTP task is active
	var waitTaskSkipped bool
	var httpTaskActive bool

	for _, task := range updatedWorkflow.Tasks {
		if task.Status == "SKIPPED" {
			waitTaskSkipped = true
		}

		if task.ReferenceTaskName == "http_ref_1" {
			if task.Status == "IN_PROGRESS" || task.Status == "COMPLETED" {
				httpTaskActive = true
			}
		}

	}
	assert.True(t, waitTaskSkipped, "Expected the wait task to be skipped")
	assert.True(t, httpTaskActive, "Expected the HTTP task to be active after jumping")
}

// TestSkipTaskFromWorkflow tests the SkipTaskFromWorkflow API endpoint
func TestSkipTaskFromWorkflow(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a workflow with multiple tasks
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SKIP_TASK_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var1").Input("var_value", "first")).
		Add(workflow.NewWaitForDurationTask("wait_task", 3*time.Second)).
		Add(workflow.NewWaitForDurationTask("wait_task_skip", 3*time.Second)).
		Add(workflow.NewSetVariableTask("set_var2").Input("var_value", "third"))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Start the workflow
	startRequest := &model.StartWorkflowRequest{
		Name:    wf.GetName(),
		Version: 1,
		Input:   map[string]interface{}{"skip_test": "value"},
	}

	workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
	assert.NoError(t, err, "Failed to start workflow")

	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), wf.GetName(), wf.GetVersion())
		assert.NoError(t, err, "Failed to remove workflow definition")
	})

	// Wait for workflow to be running
	runningWorkflow, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.RunningWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to be running")
	assert.Equal(t, model.RunningWorkflow, runningWorkflow.Status, "Workflow should be running")

	// Get workflow with tasks to find the wait task
	var waitTaskId string
	// Look for the wait task that's in progress
	for _, task := range runningWorkflow.Tasks {
		if task.ReferenceTaskName == "wait_task" && task.Status == "IN_PROGRESS" {
			waitTaskId = task.TaskId
			break
		}
	}

	assert.NotEmpty(t, waitTaskId, "Should have found the wait task ID")

	// Skip the wait task
	skipRequest := model.SkipTaskRequest{
		TaskInput: map[string]interface{}{
			"skipped": true,
			"reason":  "Test skip functionality",
		},
		TaskOutput: map[string]interface{}{
			"result": "skipped_by_test",
		},
	}

	_, err = testdata.WorkflowClient.SkipTaskFromWorkflow(
		context.Background(),
		workflowId,
		"wait_task_skip",
		skipRequest,
	)
	assert.NoError(t, err, "Failed to skip task")

	updatedWorkflow, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.RunningWorkflow, model.CompletedWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow after skip")

	var waitTaskSkipped bool
	for _, task := range updatedWorkflow.Tasks {
		if task.ReferenceTaskName == "wait_task_skip" && task.Status == model.SkippedTask {
			waitTaskSkipped = true
			// Verify the skip task output
			if task.OutputData != nil {
				assert.Equal(t, "skipped_by_test", task.OutputData["result"], "Skip task output should match")
			}
			break
		}
	}
	assert.True(t, waitTaskSkipped, "Wait task should be skipped")
}

// TestUpdateWorkflowAndTaskState tests the UpdateWorkflowAndTaskState API endpoint
func TestUpdateWorkflowAndTaskState(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a workflow with a simple task
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	taskName := "UPDATE_STATE_TASK_" + uniqueSuffix

	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_UPDATE_TASK_STATE_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSimpleTask("update_task_ref", taskName))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Register the task definition
	taskDef := model.TaskDef{
		Name:                   taskName,
		TimeoutSeconds:         60,
		RetryCount:             1,
		TimeoutPolicy:          "TIME_OUT_WF",
		RetryLogic:             "FIXED",
		RetryDelaySeconds:      1,
		ResponseTimeoutSeconds: 10,
	}
	err = testdata.ValidateTaskRegistration(taskDef)
	assert.NoError(t, err, "Failed to register task")

	// Start the workflow
	startRequest := &model.StartWorkflowRequest{
		Name:    wf.GetName(),
		Version: 1,
		Input:   map[string]interface{}{"update_test": "value"},
	}

	workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
	assert.NoError(t, err, "Failed to start workflow")

	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), wf.GetName(), wf.GetVersion())
		assert.NoError(t, err, "Failed to remove workflow definition")
	})

	workflowWithTasks, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.RunningWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to start")

	// Find the task that's scheduled
	var taskToUpdate *model.Task
	for _, task := range workflowWithTasks.Tasks {
		if task.ReferenceTaskName == taskName {
			taskToUpdate = &task
			break
		}
	}
	assert.NotNil(t, taskToUpdate, "Should find the task to update")

	// Update workflow and task state
	updateRequest := model.WorkflowStateUpdate{
		Variables: map[string]interface{}{
			"workflow_updated": true,
			"update_time":      time.Now().Unix(),
		},
		TaskReferenceName: "update_task_ref",
		TaskResult: &model.TaskResult{
			TaskId: taskToUpdate.TaskId,
			Status: model.CompletedTask,
			OutputData: map[string]interface{}{
				"updated_by_test": true,
				"task_completed":  true,
			},
		},
	}

	requestId := "test-request-" + uniqueSuffix
	opts := &client.WorkflowResourceApiUpdateWorkflowAndTaskStateOpts{
		WaitUntilTaskRef: optional.NewString("update_task_ref"),
		WaitForSeconds:   optional.NewInt32(10),
	}

	_, _, err = testdata.WorkflowClient.UpdateWorkflowAndTaskState(
		context.Background(),
		updateRequest,
		requestId,
		workflowId,
		opts,
	)
	assert.NoError(t, err, "Failed to update workflow and task state")

	completedWorkflow, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.CompletedWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to complete")
	assert.Equal(t, model.CompletedWorkflow, completedWorkflow.Status, "Workflow should be completed")

	// Verify workflow status
	assert.Equal(t, model.CompletedWorkflow, completedWorkflow.Status, "Final workflow status should be completed")
	// Verify workflow variables were updated
	assert.NotNil(t, completedWorkflow.Variables, "Workflow variables should exist")
	assert.Equal(t, true, completedWorkflow.Variables["workflow_updated"], "Workflow variable 'workflow_updated' should be true")
	assert.NotNil(t, completedWorkflow.Variables["update_time"], "Workflow variable 'update_time' should exist")

	// Verify task was completed with correct output
	assert.Len(t, completedWorkflow.Tasks, 1, "Should have exactly 1 task")
	completedTask := completedWorkflow.Tasks[0]
	assert.Equal(t, model.CompletedTask, completedTask.Status, "Task should be completed")
	assert.NotNil(t, completedTask.OutputData, "Task should have output data")
	assert.Equal(t, true, completedTask.OutputData["updated_by_test"], "Task output 'updated_by_test' should be true")
	assert.Equal(t, true, completedTask.OutputData["task_completed"], "Task output 'task_completed' should be true")

	// Verify task result was properly applied
	assert.NotEmpty(t, completedTask.TaskId, "Task should have a valid task ID")
	assert.NotEmpty(t, completedTask.EndTime, "Task should have an end time")
}

// TestUpdateWorkflowState tests the UpdateWorkflowState API endpoint
func TestUpdateWorkflowState(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a workflow with variables
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_UPDATE_STATE_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewWaitForDurationTask("wait_for_update", 30*time.Second))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Start the workflow
	startRequest := &model.StartWorkflowRequest{
		Name:    wf.GetName(),
		Version: 1,
		Input:   map[string]interface{}{"initial_var": "initial_value"},
	}

	workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
	assert.NoError(t, err, "Failed to start workflow")

	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), wf.GetName(), wf.GetVersion())
		assert.NoError(t, err, "Failed to remove workflow definition")
	})

	runningWorkflow, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.RunningWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to start")
	assert.Equal(t, model.RunningWorkflow, runningWorkflow.Status, "Workflow should be running")

	// Update workflow state with new variables
	stateUpdate := map[string]interface{}{
		"updated_var1": "updated_value1",
		"updated_var2": 42,
		"updated_var3": true,
	}

	_, resp, err := testdata.WorkflowClient.UpdateWorkflowState(
		context.Background(),
		stateUpdate,
		workflowId,
	)
	assert.NoError(t, err, "Failed to update workflow state")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected 200 status code")

	updatedWorkflowState, _, err := testdata.WorkflowClient.GetWorkflowState(
		context.Background(),
		workflowId,
		false, // includeOutput
		true,  // includeVariables
	)
	assert.NoError(t, err, "Failed to get updated workflow state")
	assert.NotNil(t, updatedWorkflowState.Variables, "Variables should not be nil")

	// Check if our variables were set

	assert.Equal(t, "updated_value1", updatedWorkflowState.Variables["updated_var1"], "Variable 1 should match")
	assert.Equal(t, float64(42), updatedWorkflowState.Variables["updated_var2"], "Variable 2 should match")
	assert.Equal(t, true, updatedWorkflowState.Variables["updated_var3"], "Variable 3 should match")
}

// TestGetExecutionStatusTaskList tests the GetExecutionStatusTaskList API endpoint
func TestGetExecutionStatusTaskList(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create a workflow with multiple tasks
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_TASK_LIST_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var1").Input("var_value", "first")).
		Add(workflow.NewSetVariableTask("set_var2").Input("var_value", "second")).
		Add(workflow.NewSetVariableTask("set_var3").Input("var_value", "third"))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Start the workflow
	startRequest := &model.StartWorkflowRequest{
		Name:    wf.GetName(),
		Version: 1,
		Input:   map[string]interface{}{"task_list_test": "value"},
	}

	workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
	assert.NoError(t, err, "Failed to start workflow")

	defer func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), wf.GetName(), wf.GetVersion())
		assert.NoError(t, err, "Failed to remove workflow definition")
	}()

	completedWorkflow, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.CompletedWorkflow}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to complete")
	assert.Equal(t, model.CompletedWorkflow, completedWorkflow.Status, "Workflow should be completed")

	// Test GetExecutionStatusTaskList
	opts := &client.WorkflowResourceAPIGetExecutionStatusTaskListOpts{
		Start:  optional.NewInt32(0),
		Count:  optional.NewInt32(10),
		Status: optional.NewInterface("COMPLETED"),
	}

	taskList, _, err := testdata.WorkflowClient.GetExecutionStatusTaskList(
		context.Background(),
		workflowId,
		opts,
	)
	assert.NoError(t, err, "Failed to get execution status task list")
	assert.NotNil(t, taskList, "Task list should not be nil")
	assert.True(t, len(taskList.Results) >= 3, "Should have at least 3 tasks")

	// Verify task details
	for _, task := range taskList.Results {
		assert.Equal(t, "SET_VARIABLE", task.TaskType, "Task type should be SET_VARIABLE")
		assert.Equal(t, model.CompletedTask, task.Status, "Task should be completed")
		assert.NotEmpty(t, task.TaskId, "Task ID should not be empty")
	}

	// Test without taskType filter
	optsNoFilter := &client.WorkflowResourceAPIGetExecutionStatusTaskListOpts{}

	allTaskList, _, err := testdata.WorkflowClient.GetExecutionStatusTaskList(
		context.Background(),
		workflowId,
		optsNoFilter,
	)
	assert.NoError(t, err, "Failed to get all execution status task list")
	assert.NotNil(t, allTaskList, "All task list should not be nil")
	assert.True(t, len(allTaskList.Results) >= len(taskList.Results), "All task list should have at least as many tasks as filtered list")
}

// TestGetWorkflows tests the GetWorkflows API endpoint (batch correlation IDs)
func TestGetWorkflows(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Create multiple correlation IDs and workflows
	correlationIds := make([]string, 2)
	workflowIds := make([]string, 4) // 2 workflows per correlation ID
	uniqueSuffix := fmt.Sprintf("%d-%s", time.Now().UnixNano(), uuid.New().String())

	// Create a simple workflow
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_BATCH_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var").Input("var_value", "batch_test"))

	err := testdata.ValidateWorkflowRegistration(wf)
	assert.NoError(t, err, "Failed to register workflow")

	// Create correlation IDs and workflows
	for i := 0; i < 2; i++ {
		correlationIds[i] = "test-batch-correlation-" + uuid.New().String()

		// Start 2 workflows for each correlation ID
		for j := 0; j < 2; j++ {
			startRequest := &model.StartWorkflowRequest{
				Name:          wf.GetName(),
				Version:       1,
				CorrelationId: correlationIds[i],
				Input: map[string]interface{}{
					"correlation_index": i,
					"workflow_index":    j,
					"test":              "batch",
				},
			}

			workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
			assert.NoError(t, err, "Failed to start workflow %d-%d", i, j)

			workflowIds[i*2+j] = workflowId
		}
	}

	// Wait for workflows to complete
	err = testdata.WaitForMultipleWorkflowsCompletion(workflowIds, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflow to complete")

	// Test GetWorkflows with batch correlation IDs
	opts := &client.WorkflowResourceApiGetWorkflowsOpts{
		IncludeClosed: optional.NewBool(true),
		IncludeTasks:  optional.NewBool(false),
	}

	var (
		workflowsMap map[string][]model.Workflow
	)
	err = testdata.RetryCondition(3, 1*time.Second,
		func() error {
			workflowsMap, _, err = testdata.WorkflowClient.GetWorkflows(
				context.Background(),
				correlationIds,
				wf.GetName(),
				opts,
			)
			return err
		},
		func() bool {
			return len(workflowsMap) == len(correlationIds)

		})

	assert.NoError(t, err, "Failed to get workflows by batch correlation IDs")
	assert.NotNil(t, workflowsMap, "Workflows map should not be nil")

	// Verify the result contains both correlation IDs
	for _, corrId := range correlationIds {
		workflows, exists := workflowsMap[corrId]
		assert.True(t, exists, "Should find workflows for correlation ID %s", corrId)
		assert.Len(t, workflows, 2, "Should find 2 workflows for correlation ID %s", corrId)

		// Verify each workflow has the correct correlation ID
		for _, wf := range workflows {
			assert.Equal(t, corrId, wf.CorrelationId, "Workflow should have correct correlation ID")
			assert.Equal(t, model.CompletedWorkflow, wf.Status, "Workflow should be completed")
		}
	}

	// Test with includeTasks = true
	optsWithTasks := &client.WorkflowResourceApiGetWorkflowsOpts{
		IncludeClosed: optional.NewBool(true),
		IncludeTasks:  optional.NewBool(true),
	}

	workflowsMapWithTasks, _, err := testdata.WorkflowClient.GetWorkflows(
		context.Background(),
		correlationIds,
		wf.GetName(),
		optsWithTasks,
	)
	assert.NoError(t, err, "Failed to get workflows with tasks by batch correlation IDs")
	assert.NotNil(t, workflowsMapWithTasks, "Workflows map with tasks should not be nil")

	// Verify tasks are included
	for _, corrId := range correlationIds {
		workflows, exists := workflowsMapWithTasks[corrId]
		assert.True(t, exists, "Should find workflows with tasks for correlation ID %s", corrId)
		for _, wf := range workflows {
			assert.True(t, len(wf.Tasks) > 0, "Workflow should have tasks")
		}
	}

	t.Cleanup(func() {
		for _, id := range workflowIds {
			err = testdata.WorkflowExecutor.RemoveWorkflow(id)
			assert.NoError(t, err, "Failed to remove workflow %s", id)
		}
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), wf.GetName(), wf.GetVersion())
		assert.NoError(t, err, "Failed to remove workflow definition")
	})
}
