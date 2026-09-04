// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
package integration_tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createTestWorkflowWithYield creates a workflow with YIELD task for testing return strategies
func createTestWorkflowWithYield(t *testing.T, workflowName string, uniqueSuffix string) *workflow.ConductorWorkflow {
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name(workflowName).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var1").Input("var_value", "first")).
		Add(workflow.NewYieldTask("yield_task").Input("yield_input_"+uniqueSuffix, "yield_input_"+uniqueSuffix)).
		Add(workflow.NewSetVariableTask("set_var2").Input("var_value", "second"))

	err := testdata.ValidateWorkflowRegistration(wf)
	require.NoError(t, err, "Failed to register workflow")

	return wf
}

// createStartWorkflowRequest creates a StartWorkflowRequest for testing
func createStartWorkflowRequest(workflowName string, testInput string) model.StartWorkflowRequest {
	return model.StartWorkflowRequest{
		Name:    workflowName,
		Version: 1,
		Input: map[string]interface{}{
			"test_input": testInput,
		},
	}
}

// createTestWorkflowsWithSubWorkflow creates main and sub-workflows for testing BLOCKING_WORKFLOW strategy
func createTestWorkflowsWithSubWorkflow(t *testing.T, uniqueSuffix string) (mainWfName, subWfName string) {
	mainWfName = "test_main_wf_" + uniqueSuffix
	subWfName = "test_sub_wf_" + uniqueSuffix

	// First, create and register the sub-workflow with a YIELD task
	// This will be the workflow that blocks execution
	subWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name(subWfName).
		Version(1).
		Description("Sub-workflow with YIELD task").
		Add(workflow.NewSetVariableTask("sub_set_var").Input("var_value", "sub_workflow_value")).
		Add(workflow.NewYieldTask("yield_task")). // This task will block execution
		Add(workflow.NewSetVariableTask("sub_final_var").Input("final_value", "completed"))

	err := testdata.ValidateWorkflowRegistration(subWorkflow)
	require.NoError(t, err, "Failed to register sub-workflow")

	// Now create and register the main workflow that calls the sub-workflow
	mainWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name(mainWfName).
		Version(1).
		Description("Main workflow that calls sub-workflow").
		Add(workflow.NewSetVariableTask("main_set_var").Input("var_value", "main_workflow_value")).
		Add(workflow.NewSubWorkflowTask("sub_workflow_ref", subWfName, 1)). // Call the sub-workflow
		Add(workflow.NewSetVariableTask("main_final_var").Input("final_value", "all_done"))

	err = testdata.ValidateWorkflowRegistration(mainWorkflow)
	require.NoError(t, err, "Failed to register main workflow")

	// Setup cleanup for both workflows
	t.Cleanup(func() {
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			mainWfName,
			1,
		)
		assert.NoError(t, err, "Failed to unregister main workflow")

		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			subWfName,
			1,
		)
		assert.NoError(t, err, "Failed to unregister sub-workflow")
	})

	return mainWfName, subWfName
}

func TestExecuteWorkflowWithReturnStrategy(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV52)

	uniqueSuffix := uuid.New().String()
	workflowName := "TEST_GO_WORKFLOW_RETURN_STRATEGY_" + uniqueSuffix

	// Create workflow using reusable function
	createdWorkflow := createTestWorkflowWithYield(t, workflowName, uniqueSuffix)

	// Create start request using reusable function
	startRequest := createStartWorkflowRequest(workflowName, "return_strategy_test")

	// Test with TARGET_WORKFLOW strategy (default)
	t.Run("TARGET_WORKFLOW", func(t *testing.T) {
		opts := client.DefaultExecuteWorkflowOpts()
		opts.ReturnStrategy = model.ReturnTargetWorkflow
		opts.RequestID = fmt.Sprintf("req-%s-target", uniqueSuffix)

		response, err := testdata.WorkflowClient.ExecuteWorkflowWithReturnStrategy(
			context.Background(),
			startRequest,
			opts,
		)
		require.NoError(t, err, "Failed to execute workflow with TARGET_WORKFLOW strategy")
		require.NotNil(t, response, "Response should not be nil")

		// Verify response
		assert.Equal(t, model.ReturnTargetWorkflow, response.ResponseType, "Response type should be TARGET_WORKFLOW")
		assert.NotEmpty(t, response.WorkflowId, "Workflow ID should not be empty")

		// Verify essential fields
		assert.Equal(t, response.WorkflowId, response.TargetWorkflowId, "WorkflowId should match TargetWorkflowId")
		assert.NotNil(t, response.Input, "Input should not be nil")
		assert.Contains(t, response.Input, "test_input", "Input should contain test_input")
		assert.Equal(t, "return_strategy_test", response.Input["test_input"], "Input value should match")
		assert.Equal(t, model.RunningWorkflow.String(), response.Status, "Task status should be IN_PROGRESS")
		assert.Equal(t, model.RunningWorkflow, response.TargetWorkflowStatus, "Target workflow status should be RUNNING")

		// Cleanup
		err = testdata.WorkflowExecutor.RemoveWorkflow(response.WorkflowId)
		assert.NoError(t, err, "Failed to remove workflow")
	})

	// Test with BLOCKING_WORKFLOW strategy
	t.Run("BLOCKING_WORKFLOW", func(t *testing.T) {
		testdata.SkipIfOSS(t, ossGapBlockingWorkflow)

		// Create workflows using reusable function
		mainWfName, _ := createTestWorkflowsWithSubWorkflow(t, uniqueSuffix)

		// Create a StartWorkflowRequest for the main workflow
		mainStartRequest := model.StartWorkflowRequest{
			Name:    mainWfName,
			Version: 1,
			Input: map[string]interface{}{
				"test_input": "blocking_workflow_test",
			},
		}

		opts := client.DefaultExecuteWorkflowOpts()
		opts.ReturnStrategy = model.ReturnBlockingWorkflow
		opts.RequestID = fmt.Sprintf("req-%s-blocking-wf", uniqueSuffix)

		response, err := testdata.WorkflowClient.ExecuteWorkflowWithReturnStrategy(
			context.Background(),
			mainStartRequest,
			opts,
		)
		require.NoError(t, err, "Failed to execute workflow with BLOCKING_WORKFLOW strategy")
		require.NotNil(t, response, "Response should not be nil")

		// Verify response
		assert.Equal(t, model.ReturnBlockingWorkflow, response.ResponseType, "Response type should be BLOCKING_WORKFLOW")
		assert.NotEmpty(t, response.WorkflowId, "Workflow ID should not be empty")
		assert.NotEmpty(t, response.TargetWorkflowId, "Target workflow ID should not be empty")
		assert.Equal(t, model.RunningWorkflow.String(), response.Status, "Task status should be IN_PROGRESS")
		assert.Equal(t, model.RunningWorkflow, response.TargetWorkflowStatus, "Target workflow status should be RUNNING")

		// The main workflow ID and the blocking workflow ID should be different
		assert.NotEqual(t, response.TargetWorkflowId, response.WorkflowId,
			"Target workflow ID should be different from workflow ID for a sub-workflow")

		// Verify that we received the blocking sub-workflow, not the main workflow
		// Find the YIELD task
		var yieldTask *model.Task
		for i := range response.Tasks {
			if response.Tasks[i].TaskType == "YIELD" {
				yieldTask = &response.Tasks[i]
				break
			}
		}

		// Verify we found the YIELD task in the blocking workflow
		assert.NotNil(t, yieldTask, "YIELD task should be present in the blocking workflow")
		if yieldTask != nil {
			assert.Equal(t, "yield_task", yieldTask.ReferenceTaskName, "YIELD task reference name should match")
			assert.Equal(t, model.InProgressTask, yieldTask.Status, "YIELD task should be in IN_PROGRESS status")
			assert.Equal(t, "YIELD", yieldTask.TaskType, "Task type should be YIELD")
		}

		// Verify that the blocking workflow is indeed the sub-workflow
		// by checking for the sub-workflow specific task
		var subSetVarTask *model.Task
		for i := range response.Tasks {
			if response.Tasks[i].ReferenceTaskName == "sub_set_var" {
				subSetVarTask = &response.Tasks[i]
				break
			}
		}
		assert.NotNil(t, subSetVarTask, "Sub-workflow task 'sub_set_var' should be present")

		// Additional verification that we're looking at the sub-workflow
		assert.NotContains(t, response.Tasks, "main_set_var",
			"Main workflow task 'main_set_var' should not be present in the blocking workflow")

		// Cleanup - we need to terminate all workflows (main and sub-workflows)
		// First terminate and remove the main workflow
		err = testdata.WorkflowExecutor.TerminateWithFailure(response.TargetWorkflowId, "Test cleanup", false)
		assert.NoError(t, err, "Failed to terminate target workflow")

		err = testdata.WorkflowExecutor.RemoveWorkflow(response.TargetWorkflowId)
		assert.NoError(t, err, "Failed to remove target workflow")

		// Then remove the blocking workflow
		err = testdata.WorkflowExecutor.RemoveWorkflow(response.WorkflowId)
		assert.NoError(t, err, "Failed to remove blocking workflow")
	})

	// Test with BLOCKING_TASK strategy
	t.Run("BLOCKING_TASK", func(t *testing.T) {
		testdata.SkipIfOSS(t, ossGapBlockingTask)

		opts := client.DefaultExecuteWorkflowOpts()
		opts.ReturnStrategy = model.ReturnBlockingTask
		opts.RequestID = fmt.Sprintf("req-%s-blocking-task", uniqueSuffix)

		response, err := testdata.WorkflowClient.ExecuteWorkflowWithReturnStrategy(
			context.Background(),
			startRequest,
			opts,
		)
		require.NoError(t, err, "Failed to execute workflow with BLOCKING_TASK strategy")
		require.NotNil(t, response, "Response should not be nil")

		// Verify response
		assert.Equal(t, model.ReturnBlockingTask, response.ResponseType, "Response type should be BLOCKING_TASK")
		assert.NotEmpty(t, response.TaskId, "Task ID should not be empty")
		assert.NotEmpty(t, response.WorkflowId, "Workflow ID should not be empty")
		assert.Equal(t, "YIELD", response.TaskType, "Task type should be YIELD")
		assert.Equal(t, "yield_task", response.ReferenceTaskName, "Reference task name should match")
		assert.Equal(t, "yield_task", response.TaskDefName, "Task definition name should match")
		assert.Equal(t, workflowName, response.WorkflowType, "Workflow type should match the workflow name")

		assert.Equal(t, model.InProgressTask.String(), response.Status, "Task status should be IN_PROGRESS")
		assert.Equal(t, model.RunningWorkflow, response.TargetWorkflowStatus, "Target workflow status should be RUNNING")

		// Verify input data is present
		assert.NotNil(t, response.Input, "Input data should not be nil")
		assert.Contains(t, response.Input, "yield_input_"+uniqueSuffix, "Input data should contain yield input")
		assert.Equal(t, "yield_input_"+uniqueSuffix, response.Input["yield_input_"+uniqueSuffix], "Input data should match expected value")

		// Cleanup
		err = testdata.WorkflowExecutor.RemoveWorkflow(response.WorkflowId)
		assert.NoError(t, err, "Failed to remove workflow")
	})

	// Test with BLOCKING_TASK_INPUT strategy
	t.Run("BLOCKING_TASK_INPUT", func(t *testing.T) {
		testdata.SkipIfOSS(t, ossGapBlockingTask)

		opts := client.DefaultExecuteWorkflowOpts()
		opts.ReturnStrategy = model.ReturnBlockingTaskInput
		opts.RequestID = fmt.Sprintf("req-%s-blocking-task-input", uniqueSuffix)

		response, err := testdata.WorkflowClient.ExecuteWorkflowWithReturnStrategy(
			context.Background(),
			startRequest,
			opts,
		)
		require.NoError(t, err, "Failed to execute workflow with BLOCKING_TASK_INPUT strategy")
		require.NotNil(t, response, "Response should not be nil")

		// Verify response
		assert.Equal(t, model.ReturnBlockingTaskInput, response.ResponseType, "Response type should be BLOCKING_TASK_INPUT")
		assert.NotEmpty(t, response.TaskId, "Task ID should not be empty")
		assert.NotEmpty(t, response.WorkflowId, "Workflow ID should not be empty")
		assert.Equal(t, "YIELD", response.TaskType, "Task type should be YIELD")
		assert.Equal(t, "yield_task", response.ReferenceTaskName, "Reference task name should match")
		assert.Equal(t, "yield_task", response.TaskDefName, "Task definition name should match")
		assert.Equal(t, workflowName, response.WorkflowType, "Workflow type should match the workflow name")
		assert.Equal(t, model.InProgressTask.String(), response.Status, "Task status should be IN_PROGRESS")
		assert.Equal(t, model.RunningWorkflow, response.TargetWorkflowStatus, "Target workflow status should be RUNNING")

		// Verify input data is present
		assert.NotNil(t, response.Input, "Input data should not be nil")
		assert.Contains(t, response.Input, "yield_input_"+uniqueSuffix, "Input data should contain yield input")
		assert.Equal(t, "yield_input_"+uniqueSuffix, response.Input["yield_input_"+uniqueSuffix], "Input data should match expected value")
		// Cleanup
		err = testdata.WorkflowExecutor.RemoveWorkflow(response.WorkflowId)
		assert.NoError(t, err, "Failed to remove workflow")
	})

	t.Cleanup(func() {
		_, err := testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			createdWorkflow.GetName(),
			createdWorkflow.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	})
}

// TestExecuteAndGetTarget tests the ExecuteAndGetTarget method
func TestExecuteAndGetTarget(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV52)

	uniqueSuffix := uuid.New().String()
	workflowName := "TEST_GO_WORKFLOW_GET_TARGET_" + uniqueSuffix

	// Create workflow using reusable function
	createdWorkflow := createTestWorkflowWithYield(t, workflowName, uniqueSuffix)

	// Create start request using reusable function
	startRequest := createStartWorkflowRequest(workflowName, "get_target_test")

	// Execute workflow with ExecuteAndGetTarget
	workflowRun, _, err := testdata.WorkflowClient.ExecuteAndGetTarget(
		context.Background(),
		startRequest,
		fmt.Sprintf("req-%s-get-target", uniqueSuffix),
		workflowName,
		1,
		[]string{},
		10,
		string(model.DurableConsistency),
	)
	require.NoError(t, err, "Failed to execute workflow with ExecuteAndGetTarget")

	// Verify workflow run with comprehensive checks based on API response example
	assert.NotEmpty(t, workflowRun.WorkflowId, "Workflow ID should not be empty")
	assert.Equal(t, model.ReturnTargetWorkflow, workflowRun.ResponseType, "Response type should be TARGET_WORKFLOW")

	// Verify essential fields
	assert.Equal(t, workflowRun.WorkflowId, workflowRun.TargetWorkflowId, "WorkflowId should match TargetWorkflowId")
	assert.NotNil(t, workflowRun.Input, "Input should not be nil")
	assert.Contains(t, workflowRun.Input, "test_input", "Input should contain test_input")
	assert.Equal(t, "get_target_test", workflowRun.Input["test_input"], "Input value should match")
	assert.Equal(t, model.RunningWorkflow, workflowRun.TargetWorkflowStatus, "Target workflow status should be RUNNING")
	assert.Equal(t, model.RunningWorkflow, workflowRun.Status, "Workflow status should be RUNNING")

	t.Cleanup(func() {
		err := testdata.WorkflowExecutor.RemoveWorkflow(workflowRun.WorkflowId)
		assert.NoError(t, err, "Failed to remove workflow")

		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			createdWorkflow.GetName(),
			createdWorkflow.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	})
}

// TestExecuteAndGetBlockingWorkflow tests the ExecuteAndGetBlockingWorkflow method
// which returns the state of the workflow that is currently blocking execution (may be a sub-workflow)
func TestExecuteAndGetBlockingWorkflow(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV52)
	testdata.SkipIfOSS(t, ossGapBlockingWorkflow)

	// Create unique names for our workflows
	uniqueSuffix := uuid.New().String()

	// Create workflows using reusable function
	mainWfName, subWfName := createTestWorkflowsWithSubWorkflow(t, uniqueSuffix)

	// Create a unique request ID
	requestId := fmt.Sprintf("req-blocking-wf-%s", uniqueSuffix)

	// Create a StartWorkflowRequest for the main workflow
	startRequest := model.StartWorkflowRequest{
		Name:    mainWfName,
		Version: 1,
		Input: map[string]interface{}{
			"test_input": "blocking_workflow_test",
		},
	}

	// Execute workflow with ExecuteAndGetBlockingWorkflow
	workflowRun, _, err := testdata.WorkflowClient.ExecuteAndGetBlockingWorkflow(
		context.Background(),
		startRequest,
		requestId,
		mainWfName,
		1,
		[]string{},
		10,
		string(model.DurableConsistency),
	)
	require.NoError(t, err, "Failed to execute workflow with ExecuteAndGetBlockingWorkflow")

	// Get both the main workflow and the sub-workflow for detailed verification
	mainWorkflowExecution, err := testdata.WorkflowExecutor.GetWorkflow(workflowRun.TargetWorkflowId, true)
	require.NoError(t, err, "Failed to get main workflow")
	require.NotNil(t, mainWorkflowExecution, "Main workflow should not be nil")

	subWorkflowExecution, err := testdata.WorkflowExecutor.GetWorkflow(workflowRun.WorkflowId, true)
	require.NoError(t, err, "Failed to get sub-workflow")
	require.NotNil(t, subWorkflowExecution, "Sub-workflow should not be nil")

	// Verify response type is BLOCKING_WORKFLOW
	assert.Equal(t, model.ReturnBlockingWorkflow, workflowRun.ResponseType, "Response type should be BLOCKING_WORKFLOW")
	assert.NotEqual(t, workflowRun.TargetWorkflowId, workflowRun.WorkflowId,
		"Target workflow ID should be different from workflow ID for a sub-workflow")
	assert.Equal(t, model.RunningWorkflow, workflowRun.TargetWorkflowStatus, "Target workflow status should be RUNNING")
	assert.Equal(t, model.RunningWorkflow, workflowRun.Status, "Workflow status should be RUNNING")

	// Verify that the TargetWorkflowId is indeed the main workflow
	assert.Equal(t, mainWfName, mainWorkflowExecution.WorkflowName,
		"Target workflow should be the main workflow")

	// Verify that the WorkflowId is indeed the sub-workflow
	assert.Equal(t, subWfName, subWorkflowExecution.WorkflowName,
		"Blocking workflow should be the sub-workflow")

	// Verify the parent-child relationship between workflows
	var subWorkflowTask *model.Task
	for _, task := range mainWorkflowExecution.Tasks {
		if task.ReferenceTaskName == "sub_workflow_ref" {
			subWorkflowTask = &task
			break
		}
	}
	require.NotNil(t, subWorkflowTask, "Main workflow should have a sub-workflow task")

	// Verify that the sub-workflow task in the main workflow references the correct sub-workflow ID
	assert.Equal(t, workflowRun.WorkflowId, subWorkflowTask.SubWorkflowId,
		"SubWorkflowId in main workflow should match the blocking workflow ID")

	// Verify that we received the blocking sub-workflow, not the main workflow
	// Find the YIELD task
	var yieldTask *model.Task
	for i := range workflowRun.Tasks {
		if workflowRun.Tasks[i].TaskType == "YIELD" {
			yieldTask = &workflowRun.Tasks[i]
			break
		}
	}

	// Verify we found the YIELD task in the blocking workflow
	assert.NotNil(t, yieldTask, "YIELD task should be present in the blocking workflow")
	if yieldTask != nil {
		assert.Equal(t, "yield_task", yieldTask.ReferenceTaskName, "YIELD task reference name should match")
		assert.Equal(t, model.InProgressTask, yieldTask.Status, "YIELD task should be in IN_PROGRESS status")
		assert.Equal(t, "YIELD", yieldTask.TaskType, "Task type should be YIELD")
	}

	// Cleanup - we need to terminate all workflows (main and sub-workflows)
	t.Cleanup(func() {
		// First terminate and remove the main workflow
		err = testdata.WorkflowExecutor.TerminateWithFailure(workflowRun.TargetWorkflowId, "Test cleanup", false)
		assert.NoError(t, err, "Failed to terminate target workflow")

		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowRun.TargetWorkflowId)
		assert.NoError(t, err, "Failed to remove target workflow")

		// Then remove the blocking workflow
		err = testdata.WorkflowExecutor.RemoveWorkflow(workflowRun.WorkflowId)
		assert.NoError(t, err, "Failed to remove blocking workflow")
	})
}

// TestExecuteAndGetBlockingTask tests the ExecuteAndGetBlockingTask method
func TestExecuteAndGetBlockingTask(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV52)
	testdata.SkipIfOSS(t, ossGapBlockingTask)

	uniqueSuffix := uuid.New().String()
	workflowName := "TEST_GO_WORKFLOW_GET_BLOCKING_TASK_" + uniqueSuffix

	// Create workflow using reusable function
	createdWorkflow := createTestWorkflowWithYield(t, workflowName, uniqueSuffix)

	// Create start request using reusable function
	startRequest := createStartWorkflowRequest(workflowName, "get_blocking_task_test")

	// Execute workflow with ExecuteAndGetBlockingTask
	taskRun, _, err := testdata.WorkflowClient.ExecuteAndGetBlockingTask(
		context.Background(),
		startRequest,
		fmt.Sprintf("req-%s-get-blocking-task", uniqueSuffix),
		workflowName,
		1,
		[]string{},
		10,
		string(model.DurableConsistency),
	)
	require.NoError(t, err, "Failed to execute workflow with ExecuteAndGetBlockingTask")

	// Verify task run with comprehensive field checks based on API documentation
	assert.NotEmpty(t, taskRun.TaskId, "Task ID should not be empty")
	assert.NotEmpty(t, taskRun.WorkflowId, "Workflow ID should not be empty")
	assert.Equal(t, "YIELD", taskRun.TaskType, "Task type should be YIELD")
	assert.Equal(t, "yield_task", taskRun.ReferenceTaskName, "Reference task name should match")
	assert.Equal(t, "yield_task", taskRun.TaskDefName, "Task definition name should match")
	assert.Equal(t, workflowName, taskRun.WorkflowType, "Workflow type should match the workflow name")

	assert.Equal(t, model.InProgressTask, taskRun.Status, "Task status should be IN_PROGRESS")
	assert.Equal(t, model.RunningWorkflow, taskRun.TargetWorkflowStatus, "Target workflow status should be RUNNING")

	// Verify input data is present
	assert.NotNil(t, taskRun.InputData, "Input data should not be nil")
	assert.Contains(t, taskRun.InputData, "yield_input_"+uniqueSuffix, "Input data should contain yield input")
	assert.Equal(t, "yield_input_"+uniqueSuffix, taskRun.InputData["yield_input_"+uniqueSuffix], "Input data should match expected value")

	t.Cleanup(func() {
		err := testdata.WorkflowExecutor.RemoveWorkflow(taskRun.WorkflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			createdWorkflow.GetName(),
			createdWorkflow.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	})
}

// TestExecuteAndGetBlockingTaskInput tests the ExecuteAndGetBlockingTaskInput method
func TestExecuteAndGetBlockingTaskInput(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV52)
	testdata.SkipIfOSS(t, ossGapBlockingTask)

	uniqueSuffix := uuid.New().String()
	workflowName := "TEST_GO_WORKFLOW_GET_BLOCKING_TASK_INPUT_" + uniqueSuffix

	// Create workflow using reusable function
	createdWorkflow := createTestWorkflowWithYield(t, workflowName, uniqueSuffix)

	// Create start request using reusable function
	startRequest := createStartWorkflowRequest(workflowName, "get_blocking_task_input_test")

	// Execute workflow with ExecuteAndGetBlockingTaskInput
	taskRun, _, err := testdata.WorkflowClient.ExecuteAndGetBlockingTaskInput(
		context.Background(),
		startRequest,
		fmt.Sprintf("req-%s-get-blocking-task-input", uniqueSuffix),
		workflowName,
		1,
		[]string{},
		10,
		string(model.DurableConsistency),
	)
	require.NoError(t, err, "Failed to execute workflow with ExecuteAndGetBlockingTaskInput")

	// Verify task run with comprehensive field checks based on API documentation
	assert.NotEmpty(t, taskRun.TaskId, "Task ID should not be empty")
	assert.NotEmpty(t, taskRun.WorkflowId, "Workflow ID should not be empty")
	assert.Equal(t, "YIELD", taskRun.TaskType, "Task type should be YIELD")
	assert.Equal(t, "yield_task", taskRun.ReferenceTaskName, "Reference task name should match")
	assert.Equal(t, "yield_task", taskRun.TaskDefName, "Task definition name should match")
	assert.Equal(t, workflowName, taskRun.WorkflowType, "Workflow type should match the workflow name")
	assert.Equal(t, model.InProgressTask, taskRun.Status, "Task status should be IN_PROGRESS")
	assert.Equal(t, model.RunningWorkflow, taskRun.TargetWorkflowStatus, "Target workflow status should be RUNNING")

	// Verify input data is present
	assert.NotNil(t, taskRun.InputData, "Input data should not be nil")
	assert.Contains(t, taskRun.InputData, "yield_input_"+uniqueSuffix, "Input data should contain yield input")
	assert.Equal(t, "yield_input_"+uniqueSuffix, taskRun.InputData["yield_input_"+uniqueSuffix], "Input data should match expected value")

	t.Cleanup(func() {
		err := testdata.WorkflowExecutor.RemoveWorkflow(taskRun.WorkflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			createdWorkflow.GetName(),
			createdWorkflow.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	})
}
