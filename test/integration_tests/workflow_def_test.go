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
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHttpTask(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	httpTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_HTTP").
		OwnerEmail("test@orkes.io").
		Version(1).
		WorkflowStatusListenerEnabled(true).
		Add(testdata.TestHttpTask)
	err := testdata.ValidateWorkflow(httpTaskWorkflow, testdata.WorkflowValidationTimeout, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflowBulk(httpTaskWorkflow, testdata.WorkflowValidationTimeout, testdata.WorkflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflowDeletion(httpTaskWorkflow)
	if err != nil {
		t.Fatal(
			"Failed to delete workflow. Reason: ", err.Error(),
		)
	}
}

func SimpleTask(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	err := testdata.ValidateTaskRegistration(*testdata.TestSimpleTask.ToTaskDef())
	if err != nil {
		t.Fatal(err)
	}
	simpleTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SIMPLE").
		Version(1).
		Add(testdata.TestSimpleTask)
	err = testdata.TaskRunner.StartWorker(
		testdata.TestSimpleTask.ReferenceName(),
		testdata.SimpleWorker,
		testdata.WorkerQty,
		testdata.WorkerPollInterval,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflow(simpleTaskWorkflow, testdata.WorkflowValidationTimeout, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflowBulk(simpleTaskWorkflow, testdata.WorkflowValidationTimeout, testdata.WorkflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.TaskRunner.DecreaseBatchSize(
		testdata.TestSimpleTask.ReferenceName(),
		testdata.WorkerQty,
	)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflowDeletion(simpleTaskWorkflow)
	if err != nil {
		t.Fatal(
			"Failed to delete workflow. Reason: ", err.Error(),
		)
	}
}

func SimpleTaskWithoutRetryCount(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	taskToRegister := testdata.TestSimpleTask.ToTaskDef()
	taskToRegister.RetryCount = 0
	err := testdata.ValidateTaskRegistration(*taskToRegister)
	if err != nil {
		t.Fatal(err)
	}
	simpleTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SIMPLE").
		Version(1).
		Add(testdata.TestSimpleTask)
	err = testdata.TaskRunner.StartWorker(
		testdata.TestSimpleTask.ReferenceName(),
		testdata.SimpleWorker,
		testdata.WorkerQty,
		testdata.WorkerPollInterval,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflow(simpleTaskWorkflow, testdata.WorkflowValidationTimeout, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflowBulk(simpleTaskWorkflow, testdata.WorkflowValidationTimeout, testdata.WorkflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.TaskRunner.DecreaseBatchSize(
		testdata.TestSimpleTask.ReferenceName(),
		testdata.WorkerQty,
	)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflowDeletion(simpleTaskWorkflow)
	if err != nil {
		t.Fatal(
			"Failed to delete workflow. Reason: ", err.Error(),
		)
	}
}

func TestInlineTask(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	inlineTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_INLINE_TASK").
		Version(1).
		Add(testdata.TestInlineTask)
	err := testdata.ValidateWorkflow(inlineTaskWorkflow, testdata.WorkflowValidationTimeout, model.CompletedWorkflow)
	if err != nil {
		t.Fatal(err)
	}
	err = testdata.ValidateWorkflowBulk(inlineTaskWorkflow, testdata.WorkflowValidationTimeout, testdata.WorkflowBulkQty)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflowDeletion(inlineTaskWorkflow)
	if err != nil {
		t.Fatal(
			"Failed to delete workflow. Reason: ", err.Error(),
		)
	}
}

func TestSqsEventTask(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	workflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_EVENT_SQS").
		Version(1).
		Add(testdata.TestSqsEventTask)
	err := testdata.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflowDeletion(workflow)
	if err != nil {
		t.Fatal(
			"Failed to delete workflow. Reason: ", err.Error(),
		)
	}
}

func TestConductorEventTask(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	workflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_EVENT_CONDUCTOR").
		Version(1).
		Add(testdata.TestConductorEventTask)
	err := testdata.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflowDeletion(workflow)
	if err != nil {
		t.Fatal(
			"Failed to delete workflow. Reason: ", err.Error(),
		)
	}
}

func TestKafkaPublishTask(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	workflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_KAFKA_PUBLISH").
		Version(1).
		Add(testdata.TestKafkaPublishTask)
	err := testdata.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflowDeletion(workflow)
	if err != nil {
		t.Fatal(
			"Failed to delete workflow. Reason: ", err.Error(),
		)
	}
}

func TestDoWhileTask(t *testing.T) {

}

func TestTerminateTask(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	workflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_TERMINATE").
		Version(1).
		Add(testdata.TestTerminateTask)
	err := testdata.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflowDeletion(workflow)
	if err != nil {
		t.Fatal(
			"Failed to delete workflow. Reason: ", err.Error(),
		)
	}
}

func TestSwitchTask(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	workflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SWITCH").
		Version(1).
		Add(testdata.TestSwitchTask)
	err := testdata.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}

	err = testdata.ValidateWorkflowDeletion(workflow)
	if err != nil {
		t.Fatal(
			"Failed to delete workflow. Reason: ", err.Error(),
		)
	}
}

func TestDynamicForkWorkflow(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("dynamic_workflow_array_sub_workflow").
		Version(1).
		Add(createDynamicForkTask())
	err := wf.Register(true)
	if err != nil {
		t.Fatal()
	}

	err = testdata.ValidateWorkflowDeletion(wf)
	if err != nil {
		t.Fatal(
			"Failed to delete workflow. Reason: ", err.Error(),
		)
	}
}

func createDynamicForkTask() *workflow.DynamicForkTask {
	return workflow.NewDynamicForkTaskWithoutPrepareTask(
		"dynamic_workflow_array_sub_workflow",
	).Input(
		"forkTaskWorkflow", "extract_user",
	).Input(
		"forkTaskInputs", []map[string]interface{}{
			{
				"input": "value1",
			},
			{
				"sub_workflow_2_inputs": map[string]interface{}{
					"key":  "value",
					"key2": 23,
				},
			},
		},
	)
}

func TestComplexSwitchWorkflow(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	wf := testdata.GetWorkflowWithComplexSwitchTask()
	err := testdata.ValidateWorkflowRegistration(wf)
	if err != nil {
		t.Fatal(err)
	}
	receivedWf, _, err := testdata.MetadataClient.Get(context.Background(), wf.GetName(), nil)
	if err != nil {
		t.Fatal(err)
	}
	counter := countMultipleSwitchInnerTasks(receivedWf.Tasks...)
	assert.Equal(t, 7, counter)

	err = testdata.ValidateWorkflowDeletion(wf)
	if err != nil {
		t.Fatal(
			"Failed to delete workflow. Reason: ", err.Error(),
		)
	}
}

func TestRegisterWorkflow_SwitchEmptyDefaultCase(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_SWITCH_EMPTY_DEFAULT_" + time.Now().Format("150405")).
		Version(1).
		Description("A test workflow").OwnerEmail("owner@example.com")

	switchTest := workflow.NewSwitchTask("test_switch", "function() { return 'true'; }").
		UseJavascript(true).
		Input("myInput", "foo").
		DefaultCase().
		SwitchCase("true", workflow.NewJQTask("jq_true", ".").Optional(true))

	wf.Add(switchTest)

	err := wf.Register(true)
	assert.NoError(t, err)
}

func TestGetWorkflowTask(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	uuid := uuid.New().String()
	// First, create and run a simple workflow that we'll reference
	sourceWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_SOURCE_WORKFLOW_" + uuid).
		OwnerEmail("owner@example.com").
		Version(1).
		Add(testdata.TestHttpTask)

	err := testdata.ValidateWorkflowRegistration(sourceWorkflow)
	assert.NoError(t, err)

	// Start the source workflow to get its ID
	sourceWorkflowId, err := sourceWorkflow.StartWorkflowWithInput(make(map[string]interface{}))
	require.NoError(t, err)

	// Wait for source workflow to complete
	sourceWorkflowChannel, err := testdata.WorkflowExecutor.MonitorExecution(sourceWorkflowId)
	require.NoError(t, err)
	_, err = executor.WaitForWorkflowCompletionUntilTimeout(sourceWorkflowChannel, testdata.WorkflowValidationTimeout)
	require.NoError(t, err)

	// // Now create a workflow with GET_WORKFLOW task that references the completed workflow
	getWorkflowTask := workflow.NewGetWorkflowTask("get_workflow_ref", sourceWorkflowId, true)

	getWorkflowWF := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_GET_WORKFLOW_" + uuid).
		Version(1).
		OwnerEmail("owner@example.com").
		Add(getWorkflowTask)

	err = testdata.ValidateWorkflowRegistration(getWorkflowWF)
	assert.NoError(t, err)
	// Start the workflow and wait for it to complete
	getWorkflowId, err := getWorkflowWF.StartWorkflowWithInput(make(map[string]interface{}))
	require.NoError(t, err)

	getWorkflowChannel, err := testdata.WorkflowExecutor.MonitorExecution(getWorkflowId)
	require.NoError(t, err)
	completedWf, err := executor.WaitForWorkflowCompletionUntilTimeout(getWorkflowChannel, testdata.WorkflowValidationTimeout)
	require.NoError(t, err)

	// Verify the workflow completed successfully
	assert.Equal(t, model.CompletedWorkflow, completedWf.Status)

	// Extract the workflow data from the output
	outputMap, ok := completedWf.Output["workflow"].(map[string]interface{})
	require.True(t, ok, "Expected output to contain workflow details")

	// Verify key fields from the retrieved workflow
	t.Log("Verifying GET_WORKFLOW task retrieved the correct workflow data")

	// 1. Verify basic workflow metadata
	assert.Equal(t, sourceWorkflowId, outputMap["workflowId"], "Retrieved workflow ID should match source workflow ID")
	assert.Equal(t, "COMPLETED", outputMap["status"], "Retrieved workflow status should be COMPLETED")
	assert.Equal(t, "TEST_GO_SOURCE_WORKFLOW_"+uuid, outputMap["workflowName"], "Retrieved workflow name should match source workflow name")
	assert.Equal(t, float64(1), outputMap["workflowVersion"], "Retrieved workflow version should match source workflow version")

	// 2. Verify task details (if includeTasks=true)
	tasks, ok := outputMap["tasks"].([]interface{})
	require.True(t, ok, "Expected workflow output to contain tasks")
	require.GreaterOrEqual(t, len(tasks), 1, "Expected at least one task in the workflow")

	// Get the first task
	firstTask, ok := tasks[0].(map[string]interface{})
	require.True(t, ok, "Expected task to be a map")

	// 3. Verify the task details
	assert.Equal(t, "HTTP", firstTask["taskType"], "Task type should be HTTP")
	assert.Equal(t, "COMPLETED", firstTask["status"], "Task status should be COMPLETED")
	assert.Equal(t, "TEST_GO_TASK_HTTP", firstTask["referenceTaskName"], "Task reference name should match")

	// Clean up
	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(getWorkflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		err = testdata.WorkflowExecutor.RemoveWorkflow(sourceWorkflowId)
		assert.NoError(t, err, "Failed to remove workflow")

		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			sourceWorkflow.GetName(),
			sourceWorkflow.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			getWorkflowWF.GetName(),
			getWorkflowWF.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	})
}

func TestGetWorkflowTaskWithDynamicInput(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	uuid := uuid.New().String()
	// First, create and run a simple workflow that we'll reference
	sourceWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_SOURCE_WORKFLOW_DYNAMIC_" + uuid).
		Version(1).
		Add(testdata.TestHttpTask)

	err := testdata.ValidateWorkflowRegistration(sourceWorkflow)
	assert.NoError(t, err)

	// Start the source workflow to get its ID
	sourceWorkflowId, err := sourceWorkflow.StartWorkflowWithInput(make(map[string]interface{}))
	require.NoError(t, err)

	// Wait for source workflow to complete
	sourceWorkflowChannel, err := testdata.WorkflowExecutor.MonitorExecution(sourceWorkflowId)
	require.NoError(t, err)
	_, err = executor.WaitForWorkflowCompletionUntilTimeout(sourceWorkflowChannel, testdata.WorkflowValidationTimeout)
	require.NoError(t, err)

	// Now create a workflow with GET_WORKFLOW task that uses dynamic parameter reference
	// The workflowId will be passed as input to the workflow
	getWorkflowTask := workflow.NewGetWorkflowTask("get_workflow_ref", "${workflow.input.targetWorkflowId}", true)
	getWorkflowWF := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_GET_WORKFLOW_DYNAMIC_" + uuid).
		Version(1).
		Add(getWorkflowTask)

	err = testdata.ValidateWorkflowRegistration(getWorkflowWF)
	assert.NoError(t, err)

	// Start the workflow with the source workflow ID as input
	workflowInput := map[string]interface{}{
		"targetWorkflowId": sourceWorkflowId,
	}

	dynamicWorkflowId, err := getWorkflowWF.StartWorkflowWithInput(workflowInput)
	require.NoError(t, err)

	// Wait for workflow to complete
	dynamicWorkflowChannel, err := testdata.WorkflowExecutor.MonitorExecution(dynamicWorkflowId)
	require.NoError(t, err)
	completedWf, err := executor.WaitForWorkflowCompletionUntilTimeout(dynamicWorkflowChannel, testdata.WorkflowValidationTimeout)
	require.NoError(t, err)

	// Verify the workflow completed successfully
	assert.Equal(t, model.CompletedWorkflow, completedWf.Status)

	// Extract the workflow data from the output
	outputMap, ok := completedWf.Output["workflow"].(map[string]interface{})
	require.True(t, ok, "Expected output to contain workflow details")

	// Verify the workflow data matches the source workflow
	t.Log("Verifying GET_WORKFLOW task retrieved the correct workflow data (dynamic input)")

	// 1. Verify basic workflow metadata
	assert.Equal(t, sourceWorkflowId, outputMap["workflowId"], "Retrieved workflow ID should match source workflow ID")
	assert.Equal(t, "COMPLETED", outputMap["status"], "Retrieved workflow status should be COMPLETED")
	assert.Equal(t, "TEST_GO_SOURCE_WORKFLOW_DYNAMIC_"+uuid, outputMap["workflowName"], "Retrieved workflow name should match source workflow name")
	assert.Equal(t, float64(1), outputMap["workflowVersion"], "Retrieved workflow version should match source workflow version")

	// 2. Verify task details (if includeTasks=true)
	tasks, ok := outputMap["tasks"].([]interface{})
	require.True(t, ok, "Expected workflow output to contain tasks")
	require.GreaterOrEqual(t, len(tasks), 1, "Expected at least one task in the workflow")

	// Get the first task
	firstTask, ok := tasks[0].(map[string]interface{})
	require.True(t, ok, "Expected task to be a map")

	// 3. Verify the task details
	assert.Equal(t, "HTTP", firstTask["taskType"], "Task type should be HTTP")
	assert.Equal(t, "COMPLETED", firstTask["status"], "Task status should be COMPLETED")
	assert.Equal(t, "TEST_GO_TASK_HTTP", firstTask["referenceTaskName"], "Task reference name should match")

	// Clean up
	t.Cleanup(func() {
		err = testdata.WorkflowExecutor.RemoveWorkflow(dynamicWorkflowId)
		assert.NoError(t, err, "Failed to remove workflow")
		err = testdata.WorkflowExecutor.RemoveWorkflow(sourceWorkflowId)
		assert.NoError(t, err, "Failed to remove workflow")

		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			sourceWorkflow.GetName(),
			sourceWorkflow.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
		_, err = testdata.MetadataClient.UnregisterWorkflowDef(
			context.Background(),
			getWorkflowWF.GetName(),
			getWorkflowWF.GetVersion(),
		)
		assert.NoError(t, err, "Failed to remove workflow definition")
	})
}

func TestYieldTask(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV52)

	uid := uuid.New().String()
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_YIELD_" + uid).
		OwnerEmail("owner@example.com").
		Version(1).
		Add(testdata.TestHttpTask).
		Add(workflow.NewYieldTask("simple_ref_1" + uid))

	// Register workflow
	err := testdata.ValidateWorkflowRegistration(wf)
	require.NoError(t, err)

	// Start workflow
	workflowId, err := wf.StartWorkflowWithInput(map[string]interface{}{"key": "value"})
	require.NoError(t, err)

	// Monitor execution
	ch, err := testdata.WorkflowExecutor.MonitorExecution(workflowId)
	require.NoError(t, err)

	output := map[string]interface{}{
		"taskOutPut": "Output passed using the API",
	}
	err = testdata.WorkflowExecutor.SignalWorkflowTaskAsync(workflowId, model.CompletedTask, output)
	require.NoError(t, err)

	// Wait for completion
	_, err = executor.WaitForWorkflowCompletionUntilTimeout(ch, testdata.WorkflowValidationTimeout)
	require.NoError(t, err)

	wfAfter, wfErr := testdata.WorkflowExecutor.GetWorkflow(workflowId, true)
	require.NoError(t, wfErr)
	require.NotNil(t, wfAfter)

	// Validate the YIELD task output was set via the signal
	var yieldTask *model.Task
	for i := range wfAfter.Tasks {
		if wfAfter.Tasks[i].ReferenceTaskName == "simple_ref_1"+uid {
			yieldTask = &wfAfter.Tasks[i]
			break
		}
	}
	require.NotNil(t, yieldTask, "expected to find YIELD task by reference name")
	assert.Equal(t, "YIELD", yieldTask.TaskType)
	assert.Equal(t, model.CompletedTask, yieldTask.Status)

	var out string
	if v, ok := yieldTask.OutputData["taskOutPut"].(string); ok {
		out = v
	}
	assert.Equal(t, "Output passed using the API", out)

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
}

func countMultipleSwitchInnerTasks(tasks ...model.WorkflowTask) int {
	counter := 0
	for _, task := range tasks {
		counter += countSwitchInnerTasks(task)
	}
	return counter
}

func countSwitchInnerTasks(task model.WorkflowTask) int {
	counter := 1
	if task.Type_ != "SWITCH" {
		return counter
	}
	for _, value := range task.DecisionCases {
		counter += countMultipleSwitchInnerTasks(value...)
	}
	return counter
}
