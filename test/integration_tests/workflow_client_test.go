package integration_tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/require"
)

func TestWorkflowTest(t *testing.T) {
	httpTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_TEST").
		Version(1).
		Add(testdata.TestHttpTask)

	err := testdata.ValidateWorkflowRegistration(httpTaskWorkflow)

	if err != nil {
		t.Fatal(
			"Failed to register workflow. Reason: ", err.Error(),
		)
	}

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
		// Optionally include the workflow definition directly
		WorkflowDef: httpTaskWorkflow.ToWorkflowDef(),
	}

	// Call the test workflow API
	workflowResult, resp, err := testdata.WorkflowClient.TestWorkflow(context.Background(), testRequest)

	// Validate response
	if err != nil {
		t.Fatalf("Failed to test workflow. API Error: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}

	// Validate workflow result
	if workflowResult.Status != "COMPLETED" {
		t.Fatalf("Expected workflow status COMPLETED, got %s", workflowResult.Status)
	}

	// Validate tasks in workflow result
	if len(workflowResult.Tasks) != 1 {
		t.Fatalf("Expected 1 task in workflow result, got %d", len(workflowResult.Tasks))
	}

	// Validate http task output
	httpTask := workflowResult.Tasks[0]
	if httpTask.Status != "COMPLETED" {
		t.Fatalf("Expected HTTP task status COMPLETED, got %s", httpTask.Status)
	}
}

func TestJumpToTask(t *testing.T) {
	input := workflow.HttpInput{
		Method: "GET",
		Uri:    "http://httpbin:8081/api/hello?name=Test123",
	}
	workflowDef := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_JUMP").
		Version(1).
		Add(testdata.TestSimpleTask).
		Add(workflow.NewHttpTask("http_ref_1", &input))

	err := testdata.ValidateWorkflowRegistration(workflowDef)
	require.NoError(t, err)
	// Start a workflow instance
	workflowInput := map[string]interface{}{
		"testKey": "testValue",
	}

	startRequest := &model.StartWorkflowRequest{
		Name:    workflowDef.GetName(),
		Version: 1,
		Input:   workflowInput,
	}

	workflowId, err := testdata.WorkflowExecutor.StartWorkflow(startRequest)
	require.NoError(t, err)

	t.Logf("Started workflow with ID: %s", workflowId)

	// Verify workflow is running
	_, err = testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.RunningWorkflow}, testdata.WorkflowValidationTimeout)
	require.NoError(t, err)

	// Define custom input for the task we're jumping to
	jumpTaskInput := map[string]interface{}{
		"uri":    "http://httpbin:8081/api/hello?name=Test1",
		"method": "GET",
	}

	// Jump to the third task, skipping the second
	opts := &client.WorkflowResourceApiJumpToTaskOpts{
		TaskReferenceName: optional.NewString("http_ref_1"),
	}

	resp, err := testdata.WorkflowClient.JumpToTask(
		context.Background(),
		jumpTaskInput,
		workflowId,
		opts,
	)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, resp.StatusCode)

	// Get the updated workflow
	updatedWorkflow, err := testdata.WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.RunningWorkflow, model.CompletedWorkflow}, testdata.WorkflowValidationTimeout)
	require.NoError(t, err)

	// Verify that the second task was skipped and the third task is being executed or completed
	var firstTaskSkipped bool
	var secondTaskActive bool

	for _, task := range updatedWorkflow.Tasks {
		if task.ReferenceTaskName == "TEST_GO_TASK_SIMPLE" {
			if task.Status == "SKIPPED" {
				firstTaskSkipped = true
			}
		}

		if task.ReferenceTaskName == "http_ref_1" {
			if task.Status == "IN_PROGRESS" || task.Status == "COMPLETED" {
				secondTaskActive = true
			}
		}
	}

	require.True(t, firstTaskSkipped)
	require.True(t, secondTaskActive)

	t.Log("Successfully tested jump to task functionality")
}
