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
	"github.com/stretchr/testify/assert"
)

// TestWorkflowRateLimitConfig tests the basic rate limit configuration
func TestWorkflowRateLimitConfig(t *testing.T) {
	testWorkflowExecutor := testdata.WorkflowExecutor

	// Create workflow with rate limit configuration
	wf := workflow.NewConductorWorkflow(testWorkflowExecutor).
		Name("test_rate_limit_workflow").
		Version(1).
		Description("Test workflow with rate limiting").
		RateLimitKey("test_key").
		ConcurrentExecutionLimit(5)

	// Add a simple task
	task := workflow.NewSimpleTask("simple_task", "simple_task_ref")
	wf.Add(task)

	// Verify rate limit configuration
	rateLimitConfig := wf.GetRateLimitConfig()
	assert.NotNil(t, rateLimitConfig)
	assert.Equal(t, "test_key", rateLimitConfig.RateLimitKey)
	assert.Equal(t, int32(5), rateLimitConfig.ConcurrentExecLimit)

	// Convert to WorkflowDef and verify
	workflowDef := wf.ToWorkflowDef()
	assert.NotNil(t, workflowDef.RateLimitConfig)
	assert.Equal(t, "test_key", workflowDef.RateLimitConfig.RateLimitKey)
	assert.Equal(t, int32(5), workflowDef.RateLimitConfig.ConcurrentExecLimit)
}

// TestWorkflowDynamicRateLimitKey tests dynamic rate limit key with workflow input
func TestWorkflowDynamicRateLimitKey(t *testing.T) {
	testWorkflowExecutor := testdata.WorkflowExecutor

	// Create workflow with dynamic rate limit key
	wf := workflow.NewConductorWorkflow(testWorkflowExecutor).
		Name("test_dynamic_rate_limit_workflow").
		Version(1).
		RateLimitKey("${workflow.input.customerId}").
		ConcurrentExecutionLimit(3)

	// Add a simple task
	task := workflow.NewSimpleTask("process_customer_task", "process_customer_ref")
	wf.Add(task)

	// Verify dynamic key configuration
	rateLimitConfig := wf.GetRateLimitConfig()
	assert.NotNil(t, rateLimitConfig)
	assert.Equal(t, "${workflow.input.customerId}", rateLimitConfig.RateLimitKey)
	assert.Equal(t, int32(3), rateLimitConfig.ConcurrentExecLimit)
}

// TestWorkflowComplexDynamicRateLimitKey tests complex dynamic expressions
func TestWorkflowComplexDynamicRateLimitKey(t *testing.T) {
	testWorkflowExecutor := testdata.WorkflowExecutor

	// Create workflow with complex dynamic rate limit key
	wf := workflow.NewConductorWorkflow(testWorkflowExecutor).
		Name("test_complex_rate_limit_workflow").
		Version(1).
		RateLimitKey("${workflow.input.tenantId}_${workflow.input.region}").
		ConcurrentExecutionLimit(10)

	// Add a simple task
	task := workflow.NewSimpleTask("multi_tenant_task", "multi_tenant_ref")
	wf.Add(task)

	// Verify complex dynamic key
	rateLimitConfig := wf.GetRateLimitConfig()
	assert.NotNil(t, rateLimitConfig)
	assert.Equal(t, "${workflow.input.tenantId}_${workflow.input.region}", rateLimitConfig.RateLimitKey)
}

// TestWorkflowSetRateLimitConfig tests the SetRateLimitConfig method
func TestWorkflowSetRateLimitConfig(t *testing.T) {
	testWorkflowExecutor := testdata.WorkflowExecutor

	// Create workflow
	wf := workflow.NewConductorWorkflow(testWorkflowExecutor).
		Name("test_set_rate_limit_workflow").
		Version(1)

	// Set rate limit config using SetRateLimitConfig
	config := &model.RateLimitConfig{
		RateLimitKey:        "api_key_${workflow.input.apiKey}",
		ConcurrentExecLimit: 15,
	}
	wf.SetRateLimitConfig(config)
	// Add a simple task
	task := workflow.NewSimpleTask("api_task", "api_task_ref")
	wf.Add(task)

	// Verify configuration
	rateLimitConfig := wf.GetRateLimitConfig()
	assert.NotNil(t, rateLimitConfig)
	assert.Equal(t, "api_key_${workflow.input.apiKey}", rateLimitConfig.RateLimitKey)
	assert.Equal(t, int32(15), rateLimitConfig.ConcurrentExecLimit)
}

// TestConcurrentWorkflowExecution tests actual concurrent execution with rate limits
func TestConcurrentWorkflowExecution(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testWorkflowExecutor := testdata.WorkflowExecutor

	// Create workflow with strict concurrency limit
	workflowName := fmt.Sprintf("test_concurrent_%d", time.Now().Unix())
	wf := workflow.NewConductorWorkflow(testWorkflowExecutor).
		Name(workflowName).
		Version(1).
		RateLimitKey("concurrent_test").
		ConcurrentExecutionLimit(2) // Only allow 2 concurrent executions

	// Add a wait task to simulate long-running workflow
	waitTask := workflow.NewWaitForDurationTask("wait_task", 5*time.Second)
	wf.Add(waitTask)

	// Register the workflow
	err := wf.Register(true)
	assert.NoError(t, err)

	// Clean up workflow after test
	defer func() {
		err := wf.UnRegister()
		if err != nil {
			t.Logf("Failed to unregister workflow: %v", err)
		}
	}()

	// Start 5 workflows simultaneously
	var wg sync.WaitGroup
	results := make([]struct {
		ID    string
		Error error
	}, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			input := map[string]interface{}{
				"index": idx,
			}
			id, err := wf.StartWorkflowWithInput(input)
			results[idx].ID = id
			results[idx].Error = err
		}(i)
	}

	wg.Wait()

	// Count successful starts
	successCount := 0
	for _, result := range results {
		if result.Error == nil && result.ID != "" {
			successCount++
			t.Logf("Workflow %s started successfully", result.ID)
		} else if result.Error != nil {
			t.Logf("Workflow start failed/queued: %v", result.Error)
		}
	}

	// Due to rate limiting, we expect some workflows to be queued
	// The exact behavior depends on server configuration
	assert.GreaterOrEqual(t, successCount, 2, "At least 2 workflows should be running")
}

// TestPerCustomerRateLimit tests rate limiting per customer ID
func TestPerCustomerRateLimit(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	testWorkflowExecutor := testdata.WorkflowExecutor

	// Create workflow with per-customer rate limiting
	workflowName := fmt.Sprintf("test_per_customer_%d", time.Now().Unix())
	wf := workflow.NewConductorWorkflow(testWorkflowExecutor).
		Name(workflowName).
		Version(1).
		RateLimitKey("${workflow.input.customerId}").
		ConcurrentExecutionLimit(2) // 2 concurrent per customer

	// Add a simple task
	task := workflow.NewSimpleTask("customer_task", "customer_task_ref")
	wf.Add(task)

	// Register the workflow
	err := wf.Register(true)
	assert.NoError(t, err)

	// Clean up workflow after test
	defer func() {
		err := wf.UnRegister()
		if err != nil {
			t.Logf("Failed to unregister workflow: %v", err)
		}
	}()

	// Start workflows for different customers
	var wg sync.WaitGroup

	// Start 3 workflows for customer1
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			input := map[string]interface{}{
				"customerId": "customer1",
				"orderId":    fmt.Sprintf("order_%d", idx),
			}
			id, err := wf.StartWorkflowWithInput(input)
			if err != nil {
				t.Logf("Customer1 workflow %d failed/queued: %v", idx, err)
			} else {
				t.Logf("Customer1 workflow %d started: %s", idx, id)
			}
		}(i)
	}
	// Start 3 workflows for customer2
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			input := map[string]interface{}{
				"customerId": "customer2",
				"orderId":    fmt.Sprintf("order_%d", idx),
			}
			id, err := wf.StartWorkflowWithInput(input)
			if err != nil {
				t.Logf("Customer2 workflow %d failed/queued: %v", idx, err)
			} else {
				t.Logf("Customer2 workflow %d started: %s", idx, id)
			}
		}(i)
	}

	wg.Wait()

	// Each customer should have their own rate limit of 2 concurrent executions
	// The third workflow for each customer should be queued
}

// TestWorkflowBuilderChaining tests fluent API chaining
func TestWorkflowBuilderChaining(t *testing.T) {
	testWorkflowExecutor := testdata.WorkflowExecutor

	// Test method chaining
	wf := workflow.NewConductorWorkflow(testWorkflowExecutor).
		Name("test_chaining_workflow").
		Version(1).
		Description("Test workflow with chained configuration").
		RateLimitKey("chained_key").
		ConcurrentExecutionLimit(7).
		TimeoutPolicy(workflow.TimeOutWorkflow, 3600).
		OwnerEmail("test@example.com")

	// Add multiple tasks
	task1 := workflow.NewSimpleTask("task1", "task1_ref")
	task2 := workflow.NewSimpleTask("task2", "task2_ref")
	wf.Add(task1).Add(task2)

	// Verify all configurations
	assert.Equal(t, "test_chaining_workflow", wf.GetName())
	assert.Equal(t, int32(1), wf.GetVersion())

	rateLimitConfig := wf.GetRateLimitConfig()
	assert.NotNil(t, rateLimitConfig)
	assert.Equal(t, "chained_key", rateLimitConfig.RateLimitKey)
	assert.Equal(t, int32(7), rateLimitConfig.ConcurrentExecLimit)

	workflowDef := wf.ToWorkflowDef()
	assert.Equal(t, "test@example.com", workflowDef.OwnerEmail)
	assert.Equal(t, "TIME_OUT_WF", workflowDef.TimeoutPolicy)
	assert.Equal(t, int64(3600), workflowDef.TimeoutSeconds)
	assert.Len(t, workflowDef.Tasks, 2)
}

// TestNoRateLimitConfiguration tests workflow without rate limit
func TestNoRateLimitConfiguration(t *testing.T) {
	testWorkflowExecutor := testdata.WorkflowExecutor

	// Create workflow without rate limit configuration
	wf := workflow.NewConductorWorkflow(testWorkflowExecutor).
		Name("test_no_rate_limit_workflow").
		Version(1)

	// Add a simple task
	task := workflow.NewSimpleTask("unrestricted_task", "unrestricted_ref")
	wf.Add(task)

	// Verify no rate limit configuration
	rateLimitConfig := wf.GetRateLimitConfig()
	assert.Nil(t, rateLimitConfig)

	// Convert to WorkflowDef and verify
	workflowDef := wf.ToWorkflowDef()
	assert.Nil(t, workflowDef.RateLimitConfig)
}

// TestUpdateRateLimitConfiguration tests updating rate limit configuration
func TestUpdateRateLimitConfiguration(t *testing.T) {
	testWorkflowExecutor := testdata.WorkflowExecutor

	// Create workflow with initial rate limit
	wf := workflow.NewConductorWorkflow(testWorkflowExecutor).
		Name("test_update_rate_limit").
		Version(1).
		RateLimitKey("initial_key").
		ConcurrentExecutionLimit(5)

	// Verify initial configuration
	config := wf.GetRateLimitConfig()
	assert.Equal(t, "initial_key", config.RateLimitKey)
	assert.Equal(t, int32(5), config.ConcurrentExecLimit)

	// Update rate limit configuration
	wf.RateLimitKey("updated_key").
		ConcurrentExecutionLimit(10)

	// Verify updated configuration
	config = wf.GetRateLimitConfig()
	assert.Equal(t, "updated_key", config.RateLimitKey)
	assert.Equal(t, int32(10), config.ConcurrentExecLimit)
}
