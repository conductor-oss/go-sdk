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
	"fmt"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
)

// uniqueSuffix helps avoid name collisions in CI between parallel runs.
func uniqueSuffix() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func legacyHandler(tk *model.Task) (interface{}, error) {
	return map[string]interface{}{
		"message":   "processed by legacy worker",
		"taskId":    tk.TaskId,
		"timestamp": time.Now().Unix(),
	}, nil
}

func simpleHandler(ctx context.Context, in interface{}) (interface{}, error) {
	return map[string]interface{}{
		"message":   "processed by simple task",
		"timestamp": time.Now().Unix(),
	}, nil
}

func typedHandler(ctx worker.TaskContext, data testdata.TestDataIn) (testdata.TestDataOut, error) {
	return testdata.TestDataOut{
		Message: "processed by typed worker",
		Sum:     data.B + len(data.A),
		TaskId:  ctx.TaskID(),
	}, nil
}

// TestLegacyWorkerIntegration validates legacy StartWorker flow with raw map output.
func TestLegacyWorkerIntegration(t *testing.T) {
	t.Parallel()
	suffix := uniqueSuffix()
	taskName := "TEST_GO_LEGACY_TASK_" + suffix
	refName := taskName

	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WF_LEGACY_" + suffix).
		Version(1).
		Add(workflow.NewSimpleTask(taskName, refName).Input("test", "test_value")).
		OutputParameters(map[string]interface{}{
			"message": fmt.Sprintf("${%s.output.message}", refName),
		})

	if err := wf.Register(true); err != nil {
		t.Fatal(err)
	}

	if err := testdata.TaskRunner.StartWorker(taskName, legacyHandler, 2, testdata.WorkerPollInterval); err != nil {
		t.Fatal(err)
	}

	expected := map[string]interface{}{"message": "processed by legacy worker"}
	if err := testdata.ValidateWorkflowWithOutput(wf, testdata.WorkflowValidationTimeout, model.CompletedWorkflow, expected); err != nil {
		t.Fatal(err)
	}
}

func TestRegularWorkerIntegration(t *testing.T) {
	t.Parallel()
	suffix := uniqueSuffix()
	taskName := "TEST_GO_REGULAR_WORKER_" + suffix

	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WF_REGULAR_" + suffix).
		Version(1).
		Add(workflow.NewSimpleTask(taskName, taskName).Input("k", "v")).
		OutputParameters(map[string]interface{}{
			"message": fmt.Sprintf("${%s.output.message}", taskName),
		})
	if err := wf.Register(true); err != nil {
		t.Fatal(err)
	}

	w := worker.NewWorker(
		taskName,
		legacyHandler,
		worker.WithBatchSize(1),
		worker.WithPollInterval(testdata.WorkerPollInterval),
	)
	if err := testdata.TaskRunner.RegisterWorker(w); err != nil {
		t.Fatal(err)
	}

	expected := map[string]interface{}{"message": "processed by legacy worker"}
	if err := testdata.ValidateWorkflowWithOutput(wf, testdata.WorkflowValidationTimeout, model.CompletedWorkflow, expected); err != nil {
		t.Fatal(err)
	}
}

// TestTypedWorkerIntegration validates typed worker with struct input/output.
func TestTypedWorkerIntegration(t *testing.T) {
	t.Parallel()
	suffix := uniqueSuffix()
	taskName := "TEST_GO_TYPED_TASK_" + suffix
	refName := taskName

	tw := worker.NewTypedWorker(
		taskName,
		typedHandler,
		worker.WithBatchSize(1),
		worker.WithPollInterval(testdata.WorkerPollInterval),
	)
	if err := testdata.TaskRunner.RegisterWorker(tw); err != nil {
		t.Fatal(err)
	}

	task := workflow.NewSimpleTask(taskName, refName).InputMap(map[string]interface{}{"a": "xyz", "b": 10})
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WF_TYPED_" + suffix).
		Version(1).
		Add(task).
		OutputParameters(map[string]interface{}{
			"message": fmt.Sprintf("${%s.output.message}", refName),
			"sum":     fmt.Sprintf("${%s.output.sum}", refName),
		})
	if err := wf.Register(true); err != nil {
		t.Fatal(err)
	}

	expected := map[string]interface{}{
		"message": "processed by typed worker",
		"sum":     float64(10 + len("xyz")),
	}
	if err := testdata.ValidateWorkflowWithOutput(wf, testdata.WorkflowValidationTimeout, model.CompletedWorkflow, expected); err != nil {
		t.Fatal(err)
	}
}

// TestSimpleTypedWorkerIntegration validates typed worker with generic map in/out.
func TestSimpleTypedWorkerIntegration(t *testing.T) {
	t.Parallel()
	suffix := uniqueSuffix()
	taskName := "TEST_GO_TYPED_MAP_TASK_" + suffix
	refName := taskName

	tw := worker.NewSimpleTypedWorker(
		taskName,
		simpleHandler,
		worker.WithBatchSize(2),
		worker.WithPollInterval(testdata.WorkerPollInterval),
	)
	if err := testdata.TaskRunner.RegisterWorker(tw); err != nil {
		t.Fatal(err)
	}

	task := workflow.NewSimpleTask(taskName, refName).InputMap(map[string]interface{}{"foo": "bar", "k": 5})
	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WF_TYPED_MAP_" + suffix).
		Version(1).
		Add(task).
		OutputParameters(map[string]interface{}{
			"message": fmt.Sprintf("${%s.output.message}", refName),
		})
	if err := wf.Register(true); err != nil {
		t.Fatal(err)
	}

	expected := map[string]interface{}{"message": "processed by simple task"}
	if err := testdata.ValidateWorkflowWithOutput(wf, testdata.WorkflowValidationTimeout, model.CompletedWorkflow, expected); err != nil {
		t.Fatal(err)
	}
}

// TestMultiTaskWorkflowIntegration validates a sequence of mixed legacy and typed workers.
func TestMultiTaskWorkflowIntegration(t *testing.T) {
	t.Parallel()
	suffix := uniqueSuffix()
	t.Logf("Starting TestMultiTaskWorkflowIntegration with suffix: %s", suffix)

	// Tasks
	legacyTaskName := "TEST_GO_LEG_MULTI_" + suffix
	regularTaskName := "TEST_GO_REG_MULTI_" + suffix
	typedTaskName := "TEST_GO_TYP_MULTI_" + suffix
	typedSimpleName := "TEST_GO_TYP_SIMPLE_MULTI_" + suffix

	// Start legacy worker
	if err := testdata.TaskRunner.StartWorker(legacyTaskName, legacyHandler, 1, testdata.WorkerPollInterval); err != nil {
		t.Fatal(err)
	}

	typedWorker := worker.NewTypedWorker(
		typedTaskName,
		typedHandler,
		worker.WithBatchSize(1),
		worker.WithPollInterval(testdata.WorkerPollInterval),
	)

	if err := testdata.TaskRunner.RegisterWorker(typedWorker); err != nil {
		t.Fatal(err)
	}

	typedWithCtxWorker := worker.NewSimpleTypedWorker(
		typedSimpleName,
		simpleHandler,
		worker.WithBatchSize(1),
		worker.WithPollInterval(testdata.WorkerPollInterval),
	)

	if err := testdata.TaskRunner.RegisterWorker(typedWithCtxWorker); err != nil {
		t.Fatal(err)
	}

	regularWorker := worker.NewWorker(
		regularTaskName,
		legacyHandler,
		worker.WithBatchSize(1),
		worker.WithPollInterval(testdata.WorkerPollInterval),
	)

	if err := testdata.TaskRunner.RegisterWorker(regularWorker); err != nil {
		t.Fatal(err)
	}

	wf := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WF_MULTI_" + suffix).
		Version(1).
		Add(workflow.NewSimpleTask(legacyTaskName, legacyTaskName).
			Input("test", "test_value")).
		Add(workflow.NewSimpleTask(typedTaskName, typedTaskName).
			InputMap(map[string]interface{}{"a": "xyz", "b": 10})).
		Add(workflow.NewSimpleTask(typedSimpleName, typedSimpleName).
			InputMap(map[string]interface{}{"foo": "bar", "k": 5})).
		Add(workflow.NewSimpleTask(regularTaskName, regularTaskName).
			InputMap(map[string]interface{}{"test": "test_value"}))

	if err := wf.Register(true); err != nil {
		t.Fatal(err)
	}

	if err := testdata.ValidateWorkflow(wf, testdata.ExtendedValidationTimeout, model.CompletedWorkflow); err != nil {
		t.Fatal(err)
	}
}

// TestParallelExecutionIntegration validates fork-join with two typed workers.
func TestParallelExecutionIntegration(t *testing.T) {
	t.Parallel()
	suffix := uniqueSuffix()
	taskName := "TEST_GO_PARALLEL_" + suffix
	taskName1 := fmt.Sprintf("%s_1", taskName)
	taskName2 := fmt.Sprintf("%s_2", taskName)
	taskName3 := fmt.Sprintf("%s_3", taskName)
	taskName4 := fmt.Sprintf("%s_4", taskName)

	// Workers
	wA := worker.NewTypedWorker(
		taskName,
		typedHandler,
		worker.WithBatchSize(4),
		worker.WithPollInterval(testdata.WorkerPollInterval),
	)

	if err := testdata.TaskRunner.RegisterWorker(wA); err != nil {
		t.Fatal(err)
	}

	w := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WF_PARALLEL_" + suffix).
		Version(1)

	fork := workflow.NewForkTask("parallel_fork_"+suffix,
		[]workflow.TaskInterface{
			workflow.NewSimpleTask(taskName, taskName1).InputMap(map[string]interface{}{"a": "xyz", "b": 10}),
			workflow.NewSimpleTask(taskName, taskName2).InputMap(map[string]interface{}{"a": "xyz", "b": 5}),
		},
		[]workflow.TaskInterface{
			workflow.NewSimpleTask(taskName, taskName3).InputMap(map[string]interface{}{"a": "xyz", "b": 5}),
			workflow.NewSimpleTask(taskName, taskName4).InputMap(map[string]interface{}{"a": "xyz", "b": 10}),
		},
	)
	w.Add(fork).Add(workflow.NewJoinTask("join_ref", taskName1, taskName2, taskName3, taskName4))

	if err := w.Register(true); err != nil {
		t.Fatal(err)
	}

	// No explicit output assertion here; validate completion only
	if err := testdata.ValidateWorkflow(w, testdata.WorkflowValidationTimeout, model.CompletedWorkflow); err != nil {
		t.Fatal(err)
	}
}
