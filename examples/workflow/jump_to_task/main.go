package main

import (
	"context"
	"examples/workflow/jump_to_task/service"
	"examples/workflow/jump_to_task/worker"
	"examples/workflow/jump_to_task/workflow"
	jumpworkflow "examples/workflow/jump_to_task/workflow"
	"fmt"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	if err := runJumpToTaskDemo(logger); err != nil {
		logger.Fatal("JumpToTask demo failed", zap.Error(err))
	}
}

func runJumpToTaskDemo(logger *zap.Logger) error {
	// Initialize Conductor service with all dependencies
	conductorService, err := service.NewConductorService(logger)
	if err != nil {
		return fmt.Errorf("failed to create conductor service: %w", err)
	}

	// Register workflows using the service
	workflows := []struct {
		createFunc func(*executor.WorkflowExecutor) (*workflow.WorkflowDefinition, error)
		name       string
	}{
		{
			createFunc: func(executor *executor.WorkflowExecutor) (*workflow.WorkflowDefinition, error) {
				return &workflow.WorkflowDefinition{
					Workflow: jumpworkflow.CreateSimpleJumpWorkflow(executor),
					Name:     "Simple Jump Workflow",
				}, nil
			},
			name: "Simple Jump Workflow",
		},
		{
			createFunc: func(executor *executor.WorkflowExecutor) (*workflow.WorkflowDefinition, error) {
				return &workflow.WorkflowDefinition{
					Workflow: jumpworkflow.CreateWaitBasedJumpWorkflow(executor),
					Name:     "Wait-based Jump Workflow",
				}, nil
			},
			name: "Wait-based Jump Workflow",
		},
	}

	for _, wf := range workflows {
		workflowDef, err := wf.createFunc(conductorService.WorkflowExecutor)
		if err != nil {
			return fmt.Errorf("failed to create workflow %s: %w", wf.name, err)
		}

		// Call workflow.Register directly - no wrapper needed
		if err := workflowDef.Workflow.Register(true); err != nil {
			conductorService.Logger.Error("Failed to register workflow",
				zap.String("name", wf.name),
				zap.Error(err))
			return fmt.Errorf("failed to register workflow %s: %w", wf.name, err)
		}
		conductorService.Logger.Info("Successfully registered workflow", zap.String("name", wf.name))
	}

	if err := conductorService.StartWorkers(gettWorkerConfigs()); err != nil {
		return fmt.Errorf("failed to start workers: %w", err)
	}

	// Run the main JumpToTask demonstration
	ctx := context.Background()

	logger.Info("=== JumpToTask API Demonstration ===")
	logger.Info("This example demonstrates jumping to specific tasks in running workflows")

	if err := simpleJump(ctx, conductorService); err != nil {
		logger.Error("Simple jump demonstration failed", zap.Error(err))
		return err
	}

	if err := waitTaskJump(ctx, conductorService); err != nil {
		logger.Error("Wait task jump demonstration failed", zap.Error(err))
		return err
	}

	logger.Info("=== All JumpToTask demonstrations completed successfully ===")
	return nil
}

// startWorkers starts all task workers needed for the demonstration using the service
func gettWorkerConfigs() []service.WorkerConfig {
	// Workers with clear, descriptive naming following best practices
	return []service.WorkerConfig{
		// Simple workflow workers
		{
			TaskType:     "process_initial_data",
			Handler:      worker.ProcessInitialDataWorker,
			Concurrency:  1,
			PollInterval: 200 * time.Millisecond,
			Description:  "Process initial workflow data",
		},
		{
			TaskType:     "validate_business_data",
			Handler:      worker.ValidateBusinessDataWorker,
			Concurrency:  1,
			PollInterval: 200 * time.Millisecond,
			Description:  "Validate business data integrity",
		},
		{
			TaskType:     "apply_business_logic",
			Handler:      worker.ApplyBusinessLogicWorker,
			Concurrency:  1,
			PollInterval: 200 * time.Millisecond,
			Description:  "Apply business logic rules",
		},
		{
			TaskType:     "complete_final_process",
			Handler:      worker.CompleteFinalProcessWorker,
			Concurrency:  1,
			PollInterval: 200 * time.Millisecond,
			Description:  "Complete final workflow processing",
		},
		// Wait workflow workers
		{
			TaskType:     "initialize_workflow_state",
			Handler:      worker.InitializeWorkflowStateWorker,
			Concurrency:  1,
			PollInterval: 200 * time.Millisecond,
			Description:  "Initialize workflow state",
		},
		{
			TaskType:     "process_middle_step",
			Handler:      worker.ProcessMiddleStepWorker,
			Concurrency:  1,
			PollInterval: 200 * time.Millisecond,
			Description:  "Process middle workflow step",
		},
	}
}

// simpleJump shows jumping in a complex workflow with HTTP tasks
func simpleJump(ctx context.Context, conductorService *service.ConductorService) error {
	conductorService.Logger.Info("=== 1. Simple Jump Demonstration ===")

	// Start workflow by calling WorkflowExecutor directly
	workflowId, err := conductorService.WorkflowExecutor.StartWorkflow(&model.StartWorkflowRequest{
		Name:    jumpworkflow.SimpleJumpWorkflowName,
		Version: 1,
		Input: map[string]interface{}{
			"demo_type": "simple_jump",
			"timestamp": time.Now().Unix(),
		},
	})
	if err != nil {
		conductorService.Logger.Error("Failed to start simple workflow", zap.Error(err))
		return fmt.Errorf("failed to start simple workflow: %w", err)
	}
	conductorService.Logger.Info("Started simple workflow", zap.String("workflow_id", workflowId))

	// Wait for first task to complete
	time.Sleep(1 * time.Second)

	// Jump to HTTP task using WorkflowResourceService directly
	conductorService.Logger.Info("Jumping to call_external_api, skipping apply_business_logic task")
	jumpInput := map[string]interface{}{
		"uri":    "https://jsonplaceholder.typicode.com/posts/2",
		"method": "GET",
	}

	opts := &client.WorkflowResourceApiJumpToTaskOpts{
		TaskReferenceName: optional.NewString("call_external_api"),
	}

	// Cast WorkflowClient to WorkflowResourceApiService to access JumpToTask
	workflowService := conductorService.WorkflowClient.(*client.WorkflowResourceApiService)
	resp, err := workflowService.JumpToTask(ctx, jumpInput, workflowId, opts)
	if err != nil {
		conductorService.Logger.Error("Failed to jump to task", zap.String("workflow_id", workflowId), zap.Error(err))
		return fmt.Errorf("failed to perform simple jump: %w", err)
	}
	conductorService.Logger.Info("Successfully jumped to task",
		zap.String("workflow_id", workflowId),
		zap.String("target_task", "call_external_api"),
		zap.Int("status_code", resp.StatusCode))

	// Wait for HTTP task to complete
	time.Sleep(3 * time.Second)

	return conductorService.CheckedJumpResults(workflowId, "Simple Jump")
}

// waitTaskJump shows jumping over wait tasks
func waitTaskJump(ctx context.Context, conductorService *service.ConductorService) error {
	conductorService.Logger.Info("=== 2. Wait Task Jump Demonstration ===")

	// Start wait-based workflow by calling WorkflowExecutor directly
	workflowId, err := conductorService.WorkflowExecutor.StartWorkflow(&model.StartWorkflowRequest{
		Name:    jumpworkflow.WaitBasedJumpWorkflowName,
		Version: 1,
		Input: map[string]interface{}{
			"demo_type": "wait_jump",
			"timestamp": time.Now().Unix(),
		},
	})
	if err != nil {
		conductorService.Logger.Error("Failed to start wait workflow", zap.Error(err))
		return fmt.Errorf("failed to start wait workflow: %w", err)
	}
	conductorService.Logger.Info("Started wait-based workflow", zap.String("workflow_id", workflowId))

	// Wait for first task, then jump before wait task starts
	time.Sleep(2 * time.Second)

	// Jump to HTTP task using WorkflowResourceService directly
	conductorService.Logger.Info("Jumping to execute_api_call, skipping wait_for_delay and process_middle_step tasks")
	jumpInput := map[string]interface{}{
		"uri":    "https://httpbin.org/json",
		"method": "GET",
	}

	opts := &client.WorkflowResourceApiJumpToTaskOpts{
		TaskReferenceName: optional.NewString("execute_api_call"),
	}

	// Cast WorkflowClient to WorkflowResourceApiService to access JumpToTask
	workflowService := conductorService.WorkflowClient.(*client.WorkflowResourceApiService)
	resp, err := workflowService.JumpToTask(ctx, jumpInput, workflowId, opts)
	if err != nil {
		conductorService.Logger.Error("Failed to jump to task", zap.String("workflow_id", workflowId), zap.Error(err))
		return fmt.Errorf("failed to perform wait jump: %w", err)
	}
	conductorService.Logger.Info("Successfully jumped to task",
		zap.String("workflow_id", workflowId),
		zap.String("target_task", "execute_api_call"),
		zap.Int("status_code", resp.StatusCode))

	// Wait for completion
	time.Sleep(3 * time.Second)

	return conductorService.CheckedJumpResults(workflowId, "Wait Task Jump")
}
