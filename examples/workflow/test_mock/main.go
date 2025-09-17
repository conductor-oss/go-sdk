package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/log"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func main() {
	logger = zap.Must(zap.NewProduction())
	defer logger.Sync()

	// Set SDK logger
	log.SetLogger(log.NewZap(logger))

	logger.Info("=== Testing & Mocking Patterns Demonstration ===")
	logger.Info("This example demonstrates workflow testing with mocked task outputs")

	if err := runTestingDemo(); err != nil {
		logger.Fatal("Testing demo failed", zap.Error(err))
	}

	logger.Info("=== Testing demonstration completed successfully ===")
}

func runTestingDemo() error {
	// Initialize Conductor API client
	apiClient := client.NewAPIClientFromEnv()
	workflowExecutor := executor.NewWorkflowExecutor(apiClient)

	// 1. Create a simple test workflow
	testWorkflow := createTestWorkflow(workflowExecutor)

	// 2. Register the workflow
	if err := testWorkflow.Register(true); err != nil {
		return fmt.Errorf("failed to register workflow: %w", err)
	}

	// 3. Test workflow with mocked tasks
	if err := testWorkflowWithMocks(apiClient, testWorkflow); err != nil {
		return fmt.Errorf("workflow testing failed: %w", err)
	}

	return nil
}

// createTestWorkflow creates a simple workflow for testing
func createTestWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	// Simple task that validates input
	validateTask := workflow.NewSimpleTask("validate_input", "validate_input").
		Input("data", "${workflow.input.data}")

	// Simple task that processes the data
	processTask := workflow.NewSimpleTask("process_data", "process_data").
		Input("validationResult", "${validate_input.output.isValid}")

	// Simple task that generates output
	outputTask := workflow.NewSimpleTask("generate_output", "generate_output").
		Input("processedData", "${process_data.output.result}")

	return workflow.NewConductorWorkflow(executor).
		Name("TEST_WORKFLOW_MOCKING").
		Version(1).
		Description("Simple workflow for testing and mocking with SIMPLE tasks only").
		Add(validateTask).
		Add(processTask).
		Add(outputTask)
}

// testWorkflowWithMocks demonstrates testing a workflow with mocked task outputs
func testWorkflowWithMocks(apiClient *client.APIClient, testWorkflow *workflow.ConductorWorkflow) error {
	logger.Info("=== Testing workflow with mocked tasks ===")

	// Create mock outputs for tasks
	taskMocks := make(map[string][]model.TaskMock)

	// Mock the validate_input task
	taskMocks["validate_input"] = []model.TaskMock{
		{
			Status: "COMPLETED",
			Output: map[string]interface{}{
				"isValid":        true,
				"validationTime": 50,
			},
			QueueWaitTime: 10, // 10ms queue wait time
		},
	}

	// Mock the process_data task
	taskMocks["process_data"] = []model.TaskMock{
		{
			Status: "COMPLETED",
			Output: map[string]interface{}{
				"processedRecords": 2,
				"processingTime":   100,
				"result":           "success",
			},
			QueueWaitTime: 15,
		},
	}

	// Mock the generate_output task
	taskMocks["generate_output"] = []model.TaskMock{
		{
			Status: "COMPLETED",
			Output: map[string]interface{}{
				"finalResult":     "Test completed successfully",
				"outputGenerated": true,
				"timestamp":       1672531200,
			},
			QueueWaitTime: 5,
		},
	}

	// Prepare the test request
	testRequest := model.WorkflowTestRequest{
		Name:    testWorkflow.GetName(),
		Version: testWorkflow.GetVersion(),
		Input: map[string]interface{}{
			"data": "test input data",
		},
		TaskRefToMockOutput: taskMocks,
	}

	// Execute the test
	ctx := context.Background()
	workflowService := &client.WorkflowResourceApiService{APIClient: apiClient}

	logger.Info("Executing workflow test...")
	workflowResult, resp, err := workflowService.TestWorkflow(ctx, testRequest)
	if err != nil {
		logger.Error("Workflow test failed", zap.Error(err))
		return fmt.Errorf("workflow test failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		logger.Error("Unexpected status code", zap.Int("status_code", resp.StatusCode))
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Validate results
	logger.Info("Workflow test completed",
		zap.String("workflow_status", string(workflowResult.Status)),
		zap.Int("total_tasks", len(workflowResult.Tasks)))

	// Log detailed task results for debugging
	for _, task := range workflowResult.Tasks {
		logger.Info("Task details",
			zap.String("task_ref", task.ReferenceTaskName),
			zap.String("task_type", task.TaskType),
			zap.String("status", string(task.Status)),
			zap.String("reason_for_incompletion", task.ReasonForIncompletion))
	}

	// Check workflow status and provide helpful information
	switch workflowResult.Status {
	case model.CompletedWorkflow:
		logger.Info("Workflow completed successfully with mocked tasks!")
		logger.Info("All SIMPLE tasks were mocked and executed as expected")
	case model.FailedWorkflow:
		logger.Warn(" Workflow failed")
		logger.Info("Reason for failure", zap.String("reason", workflowResult.ReasonForIncompletion))
		logger.Info("Note: This example now uses only SIMPLE tasks which should work with mocks")
	case model.TerminatedWorkflow:
		logger.Warn("Workflow was terminated")
	default:
		logger.Warn("Workflow ended with unexpected status",
			zap.String("status", string(workflowResult.Status)))
	}

	logger.Info("✅ Workflow testing demonstration completed!")
	logger.Info("Key learnings: TestWorkflow API allows testing workflows with mocked SIMPLE task outputs")
	return nil
}
