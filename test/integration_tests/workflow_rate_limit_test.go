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
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestConcurrentWorkflowExecution tests actual concurrent execution with rate limits
func TestConcurrentWorkflowExecution(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	testWorkflowExecutor := testdata.WorkflowExecutor

	concurrentLimit := int32(3)
	duration := 4 * time.Second
	uuid := uuid.New().String()
	rateLimitKey := "concurrent_test" + uuid
	// Create workflow with strict concurrency limit
	workflowName := fmt.Sprintf("TEST_GO_WORKFLOW_CONCURRENT_%s", uuid)
	wf := workflow.NewConductorWorkflow(testWorkflowExecutor).
		Name(workflowName).
		Version(1).
		RateLimitKey(rateLimitKey).
		ConcurrentExecutionLimit(concurrentLimit)

	// Add a wait task to simulate long-running workflow
	wf = wf.Add(workflow.NewSetVariableTask("set_var").Input("var_value", 42)).Add(workflow.NewWaitForDurationTask("wait_task", duration))

	// Register the workflow
	err := wf.Register(true)
	assert.NoError(t, err)

	// Start 10 workflows simultaneously
	var wg sync.WaitGroup
	ids := make([]string, 10)

	for i := 0; i < len(ids); i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			input := map[string]interface{}{
				"index": idx,
			}
			var id string
			err := testdata.RetryTimeout(3, 500*time.Millisecond, func() error {
				var startErr error
				id, startErr = wf.StartWorkflowWithInput(input)
				return startErr
			})
			require.NoError(t, err)
			ids[idx] = id
		}(i)
	}

	wg.Wait()

	time.Sleep(duration + 1*time.Second)

	completedCount := 0
	for _, id := range ids {
		execution, err := testWorkflowExecutor.GetWorkflow(id, true)
		require.NoError(t, err)
		require.NotNil(t, execution)

		switch execution.Status {
		case model.CompletedWorkflow:
			completedCount++
			// RateLimitKey should be skipped
			assert.Equal(t, execution.RateLimitKey, "")
		case model.RunningWorkflow:
			// RateLimitKey should be set
			assert.Equal(t, execution.RateLimitKey, rateLimitKey)
		}
	}

	assert.GreaterOrEqual(t, completedCount, int(concurrentLimit), "Completed count should be equal to concurrent limit")
	assert.Less(t, completedCount, len(ids), "Completed count should be less than or equal to the number of workflows")

	t.Cleanup(func() {
		for _, id := range ids {
			err = testdata.WorkflowExecutor.RemoveWorkflow(id)
			assert.NoError(t, err, "Failed to remove workflow %s", id)
		}

		err := wf.UnRegister()
		assert.NoError(t, err, "Failed to unregister workflow")

	})
}

// TestPerCustomerRateLimit tests rate limiting per customer ID
func TestPerCustomerRateLimit(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	testWorkflowExecutor := testdata.WorkflowExecutor
	concurrentLimit := int32(3)
	duration := 4 * time.Second
	uuid := uuid.New().String()

	// Create workflow with per-customer rate limiting
	workflowName := fmt.Sprintf("TEST_GO_WORKFLOW_CONCURRENT_CUSTOMER_%s", uuid)
	wf := workflow.NewConductorWorkflow(testWorkflowExecutor).
		Name(workflowName).
		Version(1).
		RateLimitKey("${workflow.input.customerId}").
		ConcurrentExecutionLimit(concurrentLimit)

	// Add a wait task to simulate long-running workflow
	wf = wf.Add(workflow.NewSetVariableTask("TEST_GO_SET_VAR").Input("var_value", 42)).
		Add(workflow.NewWaitForDurationTask("TEST_GO_WAIT_TASK", duration))

	// Register the workflow
	err := wf.Register(true)
	assert.NoError(t, err)

	type CustomerWorkflow struct {
		CustomerID string
		WorkflowID string
	}

	customers := []string{"customer_A_" + uuid, "customer_B_" + uuid}
	workflowsPerCustomer := 6

	allWorkflows := make([]CustomerWorkflow, 0)
	var mu sync.Mutex

	// Start workflows for each customer simultaneously
	var wg sync.WaitGroup
	for _, customerId := range customers {
		for i := 0; i < workflowsPerCustomer; i++ {
			wg.Add(1)
			go func(cId string, idx int) {
				defer wg.Done()

				input := map[string]interface{}{
					"customerId": cId,
					"index":      idx,
				}

				var id string
				err := testdata.RetryTimeout(3, 500*time.Millisecond, func() error {
					var startErr error
					id, startErr = wf.StartWorkflowWithInput(input)
					return startErr
				})
				require.NoError(t, err)

				mu.Lock()
				allWorkflows = append(allWorkflows, CustomerWorkflow{
					CustomerID: cId,
					WorkflowID: id,
				})
				mu.Unlock()
			}(customerId, i)
		}
	}

	wg.Wait()

	// Wait for workflows to complete
	time.Sleep(duration + 1*time.Second)

	// Analyze results per customer
	customerStats := make(map[string]struct {
		Completed int
		Failed    int
		Running   int
	})

	// Count results
	for _, cw := range allWorkflows {
		stats := customerStats[cw.CustomerID]

		execution, err := testWorkflowExecutor.GetWorkflow(cw.WorkflowID, true)
		require.NoError(t, err)
		require.NotNil(t, execution)

		switch execution.Status {
		case model.CompletedWorkflow:
			stats.Completed++
			assert.Equal(t, execution.RateLimitKey, "")
		case model.FailedWorkflow:
			stats.Failed++
		case model.RunningWorkflow:
			stats.Running++
			assert.Equal(t, execution.RateLimitKey, cw.CustomerID)
		}

		customerStats[cw.CustomerID] = stats
	}

	for customerId, stats := range customerStats {
		assert.GreaterOrEqual(t, stats.Completed, int(concurrentLimit),
			"Customer %s should have at least %d completed workflows, got %d",
			customerId, concurrentLimit, stats.Completed)

		assert.Less(t, stats.Completed, workflowsPerCustomer,
			"Customer %s should have at most %d completed workflows, got %d",
			customerId, workflowsPerCustomer, stats.Completed)

		assert.Equal(t, stats.Failed, 0,
			"Customer %s should have no failed workflows, got %d",
			customerId, stats.Failed)

		assert.Equal(t, stats.Running, workflowsPerCustomer-stats.Completed,
			"Customer %s should have %d running workflows, got %d",
			customerId, workflowsPerCustomer-stats.Completed, stats.Running)
	}

	t.Cleanup(func() {
		for _, id := range allWorkflows {
			err = testdata.WorkflowExecutor.RemoveWorkflow(id.WorkflowID)
			assert.NoError(t, err, "Failed to remove workflow %s", id)
		}

		err := wf.UnRegister()
		assert.NoError(t, err, "Failed to unregister workflow")
	})
}
