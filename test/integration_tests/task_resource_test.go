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

	"github.com/antihax/optional"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
)

func TestUpdateTaskRefByName(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	simpleTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_UPDATE_TASK").
		Version(1).
		Add(testdata.TestSimpleTask)

	err := testdata.ValidateWorkflowRegistration(simpleTaskWorkflow)

	if err != nil {
		t.Fatal(
			"Failed to register workflow. Reason: ", err.Error(),
		)
	}

	workflowId, response, err := testdata.WorkflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		simpleTaskWorkflow.GetName(),
		nil,
	)
	if err != nil {
		t.Fatal(
			"Failed to start workflow. Reason: ", err.Error(),
			", workflowId: ", workflowId,
			", response:, ", *response,
		)
	}
	outputData := map[string]interface{}{
		"key": "value",
	}
	returnValue, response, err := testdata.TaskClient.UpdateTaskByRefName(
		context.Background(),
		outputData,
		workflowId,
		testdata.TaskName,
		string(model.CompletedTask),
	)
	if err != nil {
		t.Fatal(
			"Failed to updated task by ref name. Reason: ", err.Error(),
			", workflowId: ", workflowId,
			", return_value: ", returnValue,
			", response:, ", *response,
		)
	}
	errorChannel := make(chan error)
	go testdata.ValidateWorkflowDaemon(
		5*time.Second,
		errorChannel,
		workflowId,
		outputData,
		model.CompletedWorkflow,
	)
	err = <-errorChannel
	if err != nil {
		t.Fatal(
			"Failed to validate workflow. Reason: ", err.Error(),
			", workflowId: ", workflowId,
		)
	} else {
		err := testdata.ValidateWorkflowDeletion(simpleTaskWorkflow)
		if err != nil {
			t.Fatal(
				"Failed to delete workflow. Reason: ", err.Error(),
			)
		}
	}
}

// TestTaskClientAll tests the All method of TaskClient
func TestTaskClientAll(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Test All - Get queue details
	queueDetails, resp, err := testdata.TaskClient.All(context.Background())
	require.NoError(t, err, "TaskClient.All failed")
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, queueDetails)
}

// TestTaskClientSize tests the Size method of TaskClient
func TestTaskClientSize(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Test Size with specific task types
	sizes, resp, err := testdata.TaskClient.Size(context.Background(), &client.TaskResourceApiSizeOpts{
		TaskType: optional.NewInterface([]string{"HTTP", "SIMPLE"}),
	})

	require.NoError(t, err, "TaskClient.Size failed")
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, sizes)
}

// TestTaskClientSearch tests the Search method of TaskClient
func TestTaskClientSearch(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Test Search
	searchResult, resp, err := testdata.TaskClient.Search(context.Background(), &client.TaskResourceApiSearch1Opts{
		Start: optional.NewInt32(0),
		Size:  optional.NewInt32(10),
	})

	require.NoError(t, err, "TaskClient.Search failed")
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, searchResult)
}
