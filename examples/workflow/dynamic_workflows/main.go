package main

import (
	"examples/workflow/dynamic_workflows/service"
	"examples/workflow/dynamic_workflows/worker"
	"examples/workflow/dynamic_workflows/workflow"
	"fmt"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	sdkworkflow "github.com/conductor-sdk/conductor-go/sdk/workflow"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	if err := runDynamicWorkflowDemo(logger); err != nil {
		logger.Fatal("Dynamic workflow demo failed", zap.Error(err))
	}
}

func runDynamicWorkflowDemo(logger *zap.Logger) error {
	// Initialize Conductor service
	conductorService, err := service.NewConductorService(logger)
	if err != nil {
		return fmt.Errorf("failed to create conductor service: %w", err)
	}

	// Start workers
	if err := conductorService.StartWorkers(getWorkerConfigs()); err != nil {
		return fmt.Errorf("failed to start workers: %w", err)
	}

	logger.Info("=== Dynamic Workflow Demo ===")
	logger.Info("Adding tasks directly in main function - TRUE dynamic workflow building")

	// Scenario 1: Simple workflow
	logger.Info("=== Scenario 1: Simple workflow ===")
	simpleWf := workflow.CreateEmptyDynamicWorkflow(conductorService.WorkflowExecutor)

	// Add tasks directly in main function
	getUserEmail := sdkworkflow.NewSimpleTask("get_user_email", "get_user_email_ref").
		Input("userid", "${workflow.input.userid}")
	simpleWf.Add(getUserEmail)

	sendEmail := sdkworkflow.NewSimpleTask("send_email", "send_email_ref").
		Input("email", "${get_user_email_ref.output.result}").
		Input("subject", "Hello from Dynamic Workflow").
		Input("body", "This is dynamically created")
	simpleWf.Add(sendEmail)

	simpleWf.OutputParameters(map[string]interface{}{
		"email": "${get_user_email_ref.output.result}",
	})

	if err := executeWorkflow(conductorService, simpleWf, "user_a", "Simple Dynamic"); err != nil {
		return err
	}

	// Scenario 2: Complex workflow with HTTP task
	logger.Info("=== Scenario 2: Complex with HTTP task ===")
	complexWf := workflow.CreateEmptyDynamicWorkflow(conductorService.WorkflowExecutor)

	// Add tasks one by one in main function
	getUserEmail2 := sdkworkflow.NewSimpleTask("get_user_email", "get_user_email_ref").
		Input("userid", "${workflow.input.userid}")
	complexWf.Add(getUserEmail2)

	// Add HTTP task directly
	httpTask := sdkworkflow.NewHttpTask("fetch_data_ref", &sdkworkflow.HttpInput{
		Uri:    "https://jsonplaceholder.typicode.com/posts/1",
		Method: "GET",
	})
	complexWf.Add(httpTask)

	// Add final email task
	finalEmail := sdkworkflow.NewSimpleTask("send_email", "send_email_ref").
		Input("email", "${get_user_email_ref.output.result}").
		Input("subject", "Data Fetched").
		Input("body", "API Response: ${fetch_data_ref.output.response.body.title}")
	complexWf.Add(finalEmail)

	complexWf.OutputParameters(map[string]interface{}{
		"email":    "${get_user_email_ref.output.result}",
		"api_data": "${fetch_data_ref.output.response.body}",
	})

	if err := executeWorkflow(conductorService, complexWf, "api_user", "Complex HTTP"); err != nil {
		return err
	}

	// Scenario 3: Conditional task addition based on runtime decision
	logger.Info("=== Scenario 3: Conditional task addition ===")
	conditionalWf := workflow.CreateEmptyDynamicWorkflow(conductorService.WorkflowExecutor)

	// Always add user email
	getUserEmail3 := sdkworkflow.NewSimpleTask("get_user_email", "get_user_email_ref").
		Input("userid", "${workflow.input.userid}")
	conditionalWf.Add(getUserEmail3)

	// Runtime decision - add different tasks based on condition
	userType := "premium" // This could come from any runtime logic
	if userType == "premium" {
		// Add extra email for premium users
		premiumEmail := sdkworkflow.NewSimpleTask("send_email", "premium_email_ref").
			Input("email", "${get_user_email_ref.output.result}").
			Input("subject", "Premium User Benefits").
			Input("body", "Thank you for being a premium user!")
		conditionalWf.Add(premiumEmail)
	}

	// Always add email
	emailTask3 := sdkworkflow.NewSimpleTask("send_email", "email_ref").
		Input("email", "${get_user_email_ref.output.result}").
		Input("subject", "Conditional Workflow").
		Input("body", fmt.Sprintf("Workflow for %s user", userType))
	conditionalWf.Add(emailTask3)

	conditionalWf.OutputParameters(map[string]interface{}{
		"email":     "${get_user_email_ref.output.result}",
		"user_type": userType,
	})

	if err := executeWorkflow(conductorService, conditionalWf, "premium_user", "Conditional"); err != nil {
		return err
	}

	logger.Info("=== All dynamic workflow scenarios completed successfully ===")
	return nil
}

// executeWorkflow registers and executes a dynamically built workflow
func executeWorkflow(conductorService *service.ConductorService, wf *sdkworkflow.ConductorWorkflow, userid string, scenarioName string) error {

	// Register the dynamically built workflow
	if err := wf.Register(true); err != nil {
		conductorService.Logger.Error("Failed to register dynamic workflow",
			zap.String("scenario", scenarioName),
			zap.Error(err))
		return fmt.Errorf("failed to register workflow: %w", err)
	}
	conductorService.Logger.Info("Successfully registered dynamic workflow",
		zap.String("scenario", scenarioName))

	// Execute the workflow
	workflowId, err := conductorService.WorkflowExecutor.StartWorkflow(&model.StartWorkflowRequest{
		Name:    workflow.DynamicWorkflowName,
		Version: 1,
		Input: map[string]interface{}{
			"userid": userid,
		},
	})
	if err != nil {
		conductorService.Logger.Error("Failed to start workflow", zap.Error(err))
		return fmt.Errorf("failed to start workflow: %w", err)
	}

	conductorService.Logger.Info("Started dynamic workflow",
		zap.String("workflow_id", workflowId),
		zap.String("scenario", scenarioName))

	// Monitor execution
	monitorChannel, err := conductorService.WorkflowExecutor.MonitorExecution(workflowId)
	if err != nil {
		conductorService.Logger.Error("Failed to monitor execution", zap.Error(err))
		return fmt.Errorf("failed to monitor execution: %w", err)
	}

	// Wait for completion
	select {
	case run := <-monitorChannel:
		conductorService.Logger.Info("Dynamic workflow completed",
			zap.String("workflow_id", run.WorkflowId),
			zap.String("status", string(run.Status)),
			zap.String("scenario", scenarioName))

		if len(run.Tasks) > 0 {
			executedTasks := make([]string, len(run.Tasks))
			for i, task := range run.Tasks {
				executedTasks[i] = task.ReferenceTaskName
			}
			conductorService.Logger.Info("Tasks executed dynamically",
				zap.Strings("executed_tasks", executedTasks),
				zap.Int("total_tasks", len(run.Tasks)),
				zap.String("scenario", scenarioName))
		}

	case <-time.After(30 * time.Second):
		conductorService.Logger.Warn("Workflow monitoring timed out", zap.String("scenario", scenarioName))
	}

	return nil
}

// getWorkerConfigs returns worker configurations
func getWorkerConfigs() []service.WorkerConfig {
	return []service.WorkerConfig{
		{
			TaskType:     "get_user_email",
			Handler:      worker.GetUserEmailWorker,
			Concurrency:  1,
			PollInterval: 100 * time.Millisecond,
			Description:  "Get user email address",
		},
		{
			TaskType:     "send_email",
			Handler:      worker.SendEmailWorker,
			Concurrency:  1,
			PollInterval: 100 * time.Millisecond,
			Description:  "Send email notification",
		},
	}
}
