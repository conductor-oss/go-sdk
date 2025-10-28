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
	"github.com/stretchr/testify/require"
)

func TestWorkerBatchSize(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	simpleTaskWorkflow := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SIMPLE").
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
	err = testdata.ValidateWorkflowBulk(simpleTaskWorkflow, testdata.ExtendedValidationTimeout, testdata.WorkflowBulkQty)
	require.NoError(t, err)
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
	err = testdata.ValidateWorkflowBulk(simpleTaskWorkflow, testdata.ExtendedValidationTimeout, testdata.WorkflowBulkQty)
	require.NoError(t, err)
}

func TestFaultyWorker(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	taskName := "TEST_GO_FAULTY_TASK"
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_FAULTY_WORKFLOW").
		Version(1).
		Add(workflow.NewSimpleTask(taskName, taskName))
	err := wf.Register(true)
	require.NoError(t, err)
	err = testdata.TaskRunner.StartWorker(
		taskName,
		testdata.FaultyWorker,
		5,
		testdata.WorkerPollInterval,
	)
	require.NoError(t, err)
	err = testdata.ValidateWorkflow(wf, 5*time.Second, model.FailedWorkflow)
	require.NoError(t, err)
}

func TestWorkerWithNonRetryableError(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	taskName := "TEST_GO_NON_RETRYABLE_ERROR_TASK"
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_NON_RETRYABLE_ERROR_WF").
		Version(1).
		Add(workflow.NewSimpleTask(taskName, taskName))
	err := wf.Register(true)
	require.NoError(t, err)
	err = testdata.TaskRunner.StartWorker(
		taskName,
		testdata.FaultyWorker,
		5,
		testdata.WorkerPollInterval,
	)
	require.NoError(t, err)
	err = testdata.ValidateWorkflow(wf, 5*time.Second, model.FailedWorkflow)
	require.NoError(t, err)
}
