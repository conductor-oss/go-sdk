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
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestConcurrentWorkflowExecution tests actual concurrent execution with rate limits
func TestConcurrentWorkflowExecutionNew(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	testWorkflowExecutor := testdata.WorkflowExecutor

	concurrentLimit := int32(2)
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
	wf = wf.Add(workflow.NewSetVariableTask("set_var").Input("var_value", 42)).Add(workflow.NewSetVariableTask("set_var_2").Input("var_value_2", 43))

	// Register the workflow
	err := wf.Register(true)
	assert.NoError(t, err)

	ids := make([]string, 5)

	for i := 0; i < len(ids); i++ {
		input := map[string]interface{}{
			"index": i,
		}
		var id string
		err := testdata.RetryTimeout(3, 500*time.Millisecond, func() error {
			var startErr error
			id, startErr = wf.StartWorkflowWithInput(input)
			return startErr
		})
		require.NoError(t, err)
		ids[i] = id
	}

	require.NoError(t, testdata.WaitForMultipleWorkflowsCompletion(ids, testdata.ExtendedValidationTimeout))

	for _, id := range ids {
		execution, err := testWorkflowExecutor.GetWorkflow(id, false)
		// Cross check:
		// 1. The workflow definition contains a rate limit configuration that was set when the flow was created.
		require.NoError(t, err)
		require.NotNil(t, execution)
		assert.Equal(t, execution.WorkflowDefinition.RateLimitConfig.RateLimitKey, rateLimitKey)
		assert.Equal(t, execution.WorkflowDefinition.RateLimitConfig.ConcurrentExecLimit, concurrentLimit)
	}

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
func TestPerCustomerRateLimitNew(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	testWorkflowExecutor := testdata.WorkflowExecutor
	concurrentLimit := int32(2)
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
		Add(workflow.NewSetVariableTask("TEST_GO_SET_VAR_2").Input("var_value_2", 43))

	// Register the workflow
	err := wf.Register(true)
	assert.NoError(t, err)

	type CustomerWorkflow struct {
		CustomerID string
		WorkflowID string
	}

	customers := []string{"customer_A_" + uuid, "customer_B_" + uuid}
	workflowsPerCustomer := 4

	allWorkflows := make([]CustomerWorkflow, 0)
	allWorkflowIds := make([]string, 0)

	for _, customerId := range customers {
		for i := 0; i < workflowsPerCustomer; i++ {
			input := map[string]interface{}{
				"customerId": customerId,
				"index":      i,
			}

			var id string
			err := testdata.RetryTimeout(3, 500*time.Millisecond, func() error {
				var startErr error
				id, startErr = wf.StartWorkflowWithInput(input)
				return startErr
			})
			require.NoError(t, err)

			allWorkflows = append(allWorkflows, CustomerWorkflow{
				CustomerID: customerId,
				WorkflowID: id,
			})
			allWorkflowIds = append(allWorkflowIds, id)
		}
	}

	// Wait for workflows to complete
	require.NoError(t, testdata.WaitForMultipleWorkflowsCompletion(allWorkflowIds, testdata.ExtendedValidationTimeout))

	// Count results
	for _, cw := range allWorkflows {
		// Cross check:
		// 1. The workflow definition contains a rate limit configuration that was set when the flow was created.
		execution, err := testWorkflowExecutor.GetWorkflow(cw.WorkflowID, false)
		require.NoError(t, err)
		require.NotNil(t, execution)
		assert.Equal(t, execution.WorkflowDefinition.RateLimitConfig.RateLimitKey, "${workflow.input.customerId}")
		assert.Equal(t, concurrentLimit, execution.WorkflowDefinition.RateLimitConfig.ConcurrentExecLimit)
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
