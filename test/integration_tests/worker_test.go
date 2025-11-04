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
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkerBatchSize(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	uuid := uuid.New().String()
	simpleTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SIMPLE_" + uuid).
		Version(1).
		Add(testdata.TestSimpleTask)
	err := testdata.TaskRunner.StartWorker(
		testdata.TestSimpleTask.ReferenceName(),
		testdata.SimpleWorker,
		5,
		testdata.WorkerPollInterval,
	)
	require.NoError(t, err)
	time.Sleep(1 * time.Second)
	require.Equal(t, 5, testdata.TaskRunner.GetBatchSizeForTask(testdata.TestSimpleTask.ReferenceName()), "Unexpected batch size")
	runningWorkflows, err := testdata.ValidateWorkflowBulk(simpleTaskWorkflow, testdata.ExtendedValidationTimeout, testdata.WorkflowBulkQty)
	require.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, testdata.ValidateWorkflowDeletion(simpleTaskWorkflow), "Failed to delete workflow")
		assert.NoError(t, testdata.ValidateWorkflowExecutionsRunningDeletion(runningWorkflows), "Failed to delete workflow executions")
	})

	err = testdata.TaskRunner.SetBatchSize(
		testdata.TestSimpleTask.ReferenceName(),
		0,
	)
	require.NoError(t, err)
	time.Sleep(1 * time.Second)
	require.Equal(t, 0, testdata.TaskRunner.GetBatchSizeForTask(testdata.TestSimpleTask.ReferenceName()), "Unexpected batch size")
	err = testdata.TaskRunner.SetBatchSize(
		testdata.TestSimpleTask.ReferenceName(),
		8,
	)
	require.NoError(t, err)
	time.Sleep(1 * time.Second)
	require.Equal(t, 8, testdata.TaskRunner.GetBatchSizeForTask(testdata.TestSimpleTask.ReferenceName()), "Unexpected batch size")
	updatedRunningWorkflows, err := testdata.ValidateWorkflowBulk(simpleTaskWorkflow, testdata.ExtendedValidationTimeout, testdata.WorkflowBulkQty)
	require.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, testdata.ValidateWorkflowExecutionsRunningDeletion(updatedRunningWorkflows), "Failed to delete workflow executions")
	})
}

func TestFaultyWorker(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	uuid := uuid.New().String()
	taskName := "TEST_GO_FAULTY_TASK_" + uuid
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_FAULTY_WORKFLOW_" + uuid).
		Version(1).
		Add(workflow.NewSimpleTask(taskName, taskName))
	require.NoError(t, wf.Register(true))
	t.Cleanup(func() {
		assert.NoError(t, testdata.ValidateWorkflowDeletion(wf), "Failed to delete workflow")
	})

	err := testdata.TaskRunner.StartWorker(
		taskName,
		testdata.FaultyWorker,
		5,
		testdata.WorkerPollInterval,
	)
	require.NoError(t, err)
	completedWorkflow, err := testdata.ValidateWorkflow(wf, 5*time.Second, model.FailedWorkflow)
	require.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, testdata.ValidateWorkflowExecutionDeletion(completedWorkflow), "Failed to delete workflow execution")
	})

}

func TestWorkerWithNonRetryableError(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	uuid := uuid.New().String()
	taskName := "TEST_GO_NON_RETRYABLE_ERROR_TASK_" + uuid
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_NON_RETRYABLE_ERROR_WF_" + uuid).
		Version(1).
		Add(workflow.NewSimpleTask(taskName, taskName))

	require.NoError(t, wf.Register(true))
	t.Cleanup(func() {
		assert.NoError(t, testdata.ValidateWorkflowDeletion(wf), "Failed to delete workflow")
	})

	err := testdata.TaskRunner.StartWorker(
		taskName,
		testdata.FaultyWorker,
		5,
		testdata.WorkerPollInterval,
	)
	require.NoError(t, err)
	completedWorkflow, err := testdata.ValidateWorkflow(wf, 5*time.Second, model.FailedWorkflow)
	require.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, testdata.ValidateWorkflowExecutionDeletion(completedWorkflow), "Failed to delete workflow execution")
	})
}
