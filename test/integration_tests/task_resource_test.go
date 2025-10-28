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
	"net/http"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUpdateTaskRefByName(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	uuid := uuid.New().String()
	simpleTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_UPDATE_TASK_" + uuid).
		Version(1).
		Add(testdata.TestSimpleTask)

	err := testdata.ValidateWorkflowRegistration(simpleTaskWorkflow)

	require.NoError(t, err, "Failed to register workflow")

	workflowId, response, err := testdata.WorkflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		simpleTaskWorkflow.GetName(),
		nil,
	)
	require.NoError(t, err, "Failed to start workflow")
	require.Equal(t, http.StatusOK, response.StatusCode, "Expected status code 200, got %d", response.StatusCode)
	require.NotEmpty(t, workflowId, "Workflow ID is empty")
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
	require.NoError(t, err, "Failed to updated task by ref name")
	require.Equal(t, http.StatusOK, response.StatusCode, "Expected status code 200, got %d", response.StatusCode)
	require.NotEmpty(t, returnValue, "Return value is empty")
	errorChannel := make(chan error)
	go testdata.ValidateWorkflowDaemon(
		5*time.Second,
		errorChannel,
		workflowId,
		outputData,
		model.CompletedWorkflow,
	)
	err = <-errorChannel
	require.NoError(t, err, "Failed to validate workflow")

	err = testdata.ValidateWorkflowDeletion(simpleTaskWorkflow)
	require.NoError(t, err, "Failed to delete workflow")
}
