package main

import (
	"context"
	"examples/workflow/lifecycle/workflow"
	"fmt"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/log"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func main() {
	logger = zap.Must(zap.NewProduction())
	defer logger.Sync()

	if err := runLifecycleDemo(); err != nil {
		logger.Fatal("Lifecycle demo failed", zap.Error(err))
	}
}

func runLifecycleDemo() error {
	// Set SDK logger
	log.SetLogger(log.NewZap(zap.Must(zap.NewProduction())))

	// Initialize Conductor clients
	apiClient := client.NewAPIClientFromEnv()
	workflowExecutor := executor.NewWorkflowExecutor(apiClient)
	workflowClient := client.NewWorkflowClient(apiClient)
	taskClient := client.NewTaskClient(apiClient)

	// Register workflows
	if err := registerWorkflows(workflowExecutor); err != nil {
		logger.Error("Failed to register workflows", zap.Error(err))
		return err
	}

	// Run comprehensive lifecycle demo
	ctx := context.Background()
	correlationId := fmt.Sprintf("lifecycle_demo_%d", time.Now().Unix())

	// 1. Start workflow with correlation ID and real-time monitoring
	logger.Info("=== 1. Starting workflow with correlation ID and MonitorExecution ===")
	workflowId, err := workflowExecutor.StartWorkflow(&model.StartWorkflowRequest{
		Name:          workflow.LifecycleWorkflowName,
		Version:       1,
		CorrelationId: correlationId,
		Input: map[string]interface{}{
			"demo_type": "lifecycle_operations",
			"timestamp": time.Now().Unix(),
		},
	})
	if err != nil {
		logger.Error("Failed to start workflow", zap.Error(err))
		return err
	}
	logger.Info("Started workflow", zap.String("workflow_id", workflowId), zap.String("correlation_id", correlationId))

	// Set up real-time monitoring using MonitorExecution
	logger.Info("Setting up real-time monitoring with MonitorExecution...")
	monitorChannel, err := workflowExecutor.MonitorExecution(workflowId)
	if err != nil {
		logger.Error("Failed to set up monitoring", zap.Error(err))
		return err
	}

	// Start a goroutine to monitor workflow completion in real-time
	monitoringComplete := make(chan bool)
	go func() {
		defer func() { monitoringComplete <- true }()

		logger.Info("Real-time monitoring started - waiting for workflow completion...")
		run := <-monitorChannel

		logger.Info("Real-time monitoring detected workflow completion!",
			zap.String("workflow_id", run.WorkflowId),
			zap.String("status", string(run.Status)),
			zap.String("reason", run.ReasonForIncompletion))

		if len(run.Tasks) > 0 {
			logger.Info("Final task analysis",
				zap.Int("total_tasks", len(run.Tasks)),
				zap.String("last_task", run.Tasks[len(run.Tasks)-1].ReferenceTaskName))
		}
	}()

	// 2. Check workflow status
	logger.Info("=== 2. Checking workflow status ===")
	time.Sleep(2 * time.Second)
	wf, _, err := workflowClient.GetExecutionStatus(ctx, workflowId,
		&client.WorkflowResourceApiGetExecutionStatusOpts{IncludeTasks: optional.NewBool(true)})
	if err != nil {
		logger.Error("Failed to get workflow status", zap.Error(err))
		return err
	}
	logger.Info("Workflow status",
		zap.String("status", string(wf.Status)),
		zap.Int("tasks_count", len(wf.Tasks)))

	if len(wf.Tasks) > 0 {
		lastTask := wf.Tasks[len(wf.Tasks)-1]
		logger.Info("Currently running task",
			zap.String("task_ref", lastTask.ReferenceTaskName),
			zap.String("status", string(lastTask.Status)))
	} else {
		logger.Error("No tasks found in workflow", zap.String("workflow_id", workflowId))
		return fmt.Errorf("no tasks found in workflow %s", workflowId)
	}

	// 3. Pause workflow
	logger.Info("=== 3. Pausing workflow ===")
	if _, err := workflowClient.PauseWorkflow(ctx, workflowId); err != nil {
		logger.Error("Failed to pause workflow", zap.Error(err))
		return err
	}

	logger.Info("Paused workflow")
	time.Sleep(2 * time.Second)

	wf, _, err = workflowClient.GetExecutionStatus(ctx, workflowId, nil)
	if err != nil {
		logger.Error("Failed to get workflow status", zap.Error(err))
		return err
	}
	logger.Info("Workflow status after pause", zap.String("status", string(wf.Status)))

	// 4. Resume workflow
	logger.Info("=== 4. Resuming workflow ===")
	if _, err := workflowClient.ResumeWorkflow(ctx, workflowId); err != nil {
		logger.Error("Failed to resume workflow", zap.Error(err))
		return err
	}
	logger.Info("Resumed workflow")

	// 5. Complete wait task manually (if exists)
	logger.Info("=== 5. Checking for wait tasks to complete ===")
	time.Sleep(2 * time.Second)
	wf, _, err = workflowClient.GetExecutionStatus(ctx, workflowId,
		&client.WorkflowResourceApiGetExecutionStatusOpts{IncludeTasks: optional.NewBool(true)})
	if err != nil {
		logger.Error("Failed to get workflow status", zap.Error(err))
		return err
	}

	for _, task := range wf.Tasks {
		if task.TaskType == "WAIT" && task.Status == model.InProgressTask {
			logger.Info("Found WAIT task to complete", zap.String("task_id", task.TaskId))

			taskResult := &model.TaskResult{
				WorkflowInstanceId: workflowId,
				TaskId:             task.TaskId,
				Status:             "COMPLETED",
				OutputData: map[string]interface{}{
					"completed_by":    "lifecycle_demo",
					"completion_time": time.Now().Unix(),
				},
			}

			if _, _, err := taskClient.UpdateTask(ctx, taskResult); err != nil {
				logger.Error("Failed to complete wait task", zap.Error(err))
				return err
			}

			logger.Info("Completed wait task manually")
			break
		}
	}

	// 6. Search workflows by correlation ID
	logger.Info("=== 6. Searching workflows by correlation ID ===")
	workflows, _, err := workflowClient.GetWorkflowsByCorrelationId(ctx, workflow.LifecycleWorkflowName, correlationId, nil)
	if err != nil {
		logger.Error("Failed to search by correlation ID", zap.Error(err))
		return err
	} else {
		logger.Info("Found workflows by correlation ID",
			zap.Int("count", len(workflows)),
			zap.String("correlation_id", correlationId))
		for _, wf := range workflows {
			logger.Info("  - Workflow",
				zap.String("id", wf.WorkflowId),
				zap.String("status", string(wf.Status)))
		}
	}

	// 7. Terminate workflow
	logger.Info("=== 7. Terminating workflow ===")
	if err := workflowExecutor.Terminate(workflowId, "Terminating for retry demo"); err != nil {
		logger.Error("Failed to terminate workflow", zap.Error(err))
		return err
	}
	logger.Info("Terminated workflow")

	// 8. Retry workflow
	logger.Info("=== 8. Retrying workflow ===")
	time.Sleep(2 * time.Second)
	if _, err := workflowClient.Retry(ctx, workflowId, nil); err != nil {
		logger.Error("Failed to retry workflow", zap.Error(err))
		return err
	}
	logger.Info("Retried workflow")

	// Wait a bit and pause to prevent completion
	time.Sleep(500 * time.Millisecond)

	// Try to pause quickly to prevent full completion
	logger.Info("Pausing workflow to prevent immediate completion")
	if _, err := workflowClient.PauseWorkflow(ctx, workflowId); err != nil {
		logger.Error("Failed to pause after retry", zap.Error(err))
		return err
	}

	time.Sleep(500 * time.Millisecond)

	wf, _, err = workflowClient.GetExecutionStatus(ctx, workflowId, nil)
	if err != nil {
		logger.Error("Failed to get workflow status", zap.Error(err))
		return err
	}
	logger.Info("Workflow status after retry", zap.String("status", string(wf.Status)))

	// 9. Terminate again for restart demo
	logger.Info("=== 9. Terminating for restart demo ===")
	// Check current status before trying to terminate
	wf, _, err = workflowClient.GetExecutionStatus(ctx, workflowId, nil)
	if err != nil {
		logger.Error("Failed to get workflow status", zap.Error(err))
		return err
	}

	if wf.Status == model.CompletedWorkflow || wf.Status == model.FailedWorkflow {
		logger.Info("Workflow already in terminal state", zap.String("status", string(wf.Status)))
		logger.Info("Skipping termination - will restart from completed state")
	} else {
		if err := workflowExecutor.Terminate(workflowId, "Terminating for restart demo"); err != nil {
			logger.Error("Failed to terminate workflow", zap.Error(err))
			return err
		}
		logger.Info("Terminated workflow for restart")
	}

	// 10. Restart workflow
	logger.Info("=== 10. Restarting workflow ===")
	time.Sleep(1 * time.Second)
	if _, err := workflowClient.Restart(ctx, workflowId, nil); err != nil {
		logger.Error("Failed to restart workflow", zap.Error(err))
		return err
	}
	logger.Info("Restarted workflow")

	// 11. Rerun from specific task
	logger.Info("=== 11. Checking for rerun capability ===")
	time.Sleep(2 * time.Second)
	wf, _, err = workflowClient.GetExecutionStatus(ctx, workflowId,
		&client.WorkflowResourceApiGetExecutionStatusOpts{IncludeTasks: optional.NewBool(true)})
	if err != nil {
		logger.Error("Failed to get workflow status", zap.Error(err))
		return err
	}

	if len(wf.Tasks) > 1 {
		secondTask := wf.Tasks[1]
		logger.Info("Attempting to rerun from second task",
			zap.String("task_id", secondTask.TaskId),
			zap.String("task_ref", secondTask.ReferenceTaskName))

		rerunRequest := model.RerunWorkflowRequest{
			ReRunFromTaskId: secondTask.TaskId,
		}

		if _, _, err := workflowClient.Rerun(ctx, rerunRequest, workflowId); err != nil {
			logger.Error("Failed to rerun from task", zap.Error(err))
			return err
		}
		logger.Info("Rerun initiated from task")
	} else {
		logger.Warn("No tasks found in workflow", zap.String("workflow_id", workflowId))
	}

	// 12. Final termination with failure workflow trigger
	logger.Info("=== 12. Terminating with failure workflow trigger ===")
	time.Sleep(2 * time.Second)
	if err := workflowExecutor.TerminateWithFailure(workflowId, "Final termination with failure workflow", true); err != nil {
		logger.Error("Failed to terminate with failure", zap.Error(err))
		return err
	}
	logger.Info("Terminated with failure workflow trigger")

	// 13. Check final status
	logger.Info("=== 13. Final status check ===")
	time.Sleep(3 * time.Second)
	wf, _, err = workflowClient.GetExecutionStatus(ctx, workflowId, nil)
	if err != nil {
		logger.Error("Failed to get final status", zap.Error(err))
		return err
	}

	logger.Info("Final workflow status",
		zap.String("workflow_id", workflowId),
		zap.String("status", string(wf.Status)),
		zap.String("reason_for_incompletion", wf.ReasonForIncompletion))

	// 14. Check if failure workflow was triggered
	logger.Info("=== 14. Checking for failure workflow execution ===")
	// Note: Failure workflows with simple tasks complete very quickly
	// We check immediately, but they may have already finished
	time.Sleep(1 * time.Second)

	// First check for running failure workflows
	failureWorkflows, _, err := workflowClient.GetRunningWorkflow(ctx, workflow.FailureWorkflowName, nil)
	if err != nil {
		logger.Warn("Failed to get running workflows", zap.Error(err))
		logger.Info("Note: Failure workflow check may not be supported in all Conductor versions")
	} else {
		if len(failureWorkflows) > 0 {
			logger.Info("Found RUNNING failure workflow(s)", zap.Int("count", len(failureWorkflows)))
			for _, fwId := range failureWorkflows {
				logger.Info("  - Failure workflow", zap.String("id", fwId))
			}
		} else {
			logger.Warn("No running failure workflows found")
		}
	}

	// 15. Demonstrate RemoveWorkflow API
	logger.Info("=== 15. Demonstrating RemoveWorkflow API ===")

	logger.Info("Checking status of existing workflow before removal", zap.String("workflow_id", workflowId))
	wf, _, err = workflowClient.GetExecutionStatus(ctx, workflowId, nil)
	if err != nil {
		logger.Error("Failed to get existing workflow status", zap.Error(err))
		return err
	}

	logger.Info("Existing workflow status",
		zap.String("workflow_id", workflowId),
		zap.String("status", string(wf.Status)),
		zap.String("reason_for_incompletion", wf.ReasonForIncompletion))

	// Verify the workflow is in a terminal state
	if wf.Status != model.CompletedWorkflow && wf.Status != model.FailedWorkflow && wf.Status != model.TerminatedWorkflow {
		return fmt.Errorf("workflow must be in a terminal state for removal, current status: %v", wf.Status)
	}

	logger.Info("Workflow is in terminal state and ready for removal")

	// Now demonstrate RemoveWorkflow API

	logger.Info("Removing existing workflow from execution history")
	err = workflowExecutor.RemoveWorkflow(workflowId)
	if err != nil {
		logger.Error("Failed to remove workflow", zap.Error(err))
		return err
	}
	logger.Info("Workflow removed successfully from execution history")

	// Verify removal - this should fail
	logger.Info("Verifying workflow removal by attempting to retrieve it")
	_, _, err = workflowClient.GetExecutionStatus(ctx, workflowId, nil)
	if err != nil {
		logger.Info("Workflow successfully removed from history - retrieval failed as expected",
			zap.String("error", err.Error()))
	} else {
		logger.Warn("Unexpected: workflow still exists after removal")
	}

	// Wait for real-time monitoring to complete
	logger.Info("=== 16. Waiting for real-time monitoring to complete ===")
	select {
	case <-monitoringComplete:
		logger.Info("Real-time monitoring completed successfully")
	case <-time.After(5 * time.Second):
		logger.Warn("Real-time monitoring timed out (workflow may have completed before final operations)")
	}

	logger.Info("=== Lifecycle demo completed successfully ===")
	return nil
}

func registerWorkflows(workflowExecutor *executor.WorkflowExecutor) error {
	// Register failure workflow
	failureWf := workflow.CreateFailureWorkflow(workflowExecutor)
	if err := failureWf.Register(true); err != nil {
		logger.Error("Failed to register failure workflow", zap.Error(err))
		return err
	}
	logger.Info("Registered failure workflow")

	// Create and register main workflow with failure handler
	mainWf := workflow.CreateLifecycleWorkflow(workflowExecutor)
	mainWf.FailureWorkflow(failureWf.GetName())
	if err := mainWf.Register(true); err != nil {
		logger.Error("Failed to register main workflow", zap.Error(err))
		return err
	}
	logger.Info("Registered main workflow")

	return nil
}
