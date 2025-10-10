package main

import (
	"context"
	"examples/workflow/lifecycle/workflow"
	"fmt"
	"time"

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
	defer func() {
		err := logger.Sync()
		if err != nil {
			logger.Warn("Failed to sync logger", zap.Error(err))
		}
	}()

	if err := runLifecycleDemo(); err != nil {
		logger.Fatal("Lifecycle demo failed", zap.Error(err))
	}
}

//nolint:gocyclo,gocognit
func runLifecycleDemo() error {
	// Set SDK logger
	log.SetLogger(log.NewZap(zap.Must(zap.NewProduction())))

	// Initialize Conductor clients
	apiClient := client.NewAPIClientFromEnv()
	workflowExecutor := executor.NewWorkflowExecutor(apiClient)

	// Register workflows
	if err := registerWorkflows(workflowExecutor); err != nil {
		logger.Error("Failed to register workflows", zap.Error(err))
		return err
	}

	defer func() {
		logger.Info("Unregistering workflows")
		if err := workflowExecutor.UnRegisterWorkflow(workflow.FailureWorkflowName, 1); err != nil {
			logger.Error("Failed to unregister failure workflow", zap.Error(err))
		}
		if err := workflowExecutor.UnRegisterWorkflow(workflow.LifecycleWorkflowName, 1); err != nil {
			logger.Error("Failed to unregister lifecycle workflow", zap.Error(err))
		}
	}()

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
	wf, err := workflowExecutor.GetWorkflowWithContext(ctx, workflowId, true)
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
	if err = workflowExecutor.PauseWithContext(ctx, workflowId); err != nil {
		logger.Error("Failed to pause workflow", zap.Error(err))
		return err
	}

	logger.Info("Paused workflow")

	wf, err = workflowExecutor.GetWorkflowWithContext(ctx, workflowId, true)
	if err != nil {
		logger.Error("Failed to get workflow status", zap.Error(err))
		return err
	}
	logger.Info("Workflow status after pause", zap.String("status", string(wf.Status)))

	// 4. Resume workflow
	logger.Info("=== 4. Resuming workflow ===")
	if err = workflowExecutor.ResumeWithContext(ctx, workflowId); err != nil {
		logger.Error("Failed to resume workflow", zap.Error(err))
		return err
	}
	logger.Info("Resumed workflow")

	// 5. Complete wait task manually (if exists)
	logger.Info("=== 5. Checking for wait tasks to complete ===")
	wf, err = workflowExecutor.GetWorkflowWithContext(ctx, workflowId, true)
	if err != nil {
		logger.Error("Failed to get workflow status", zap.Error(err))
		return err
	}

	for _, task := range wf.Tasks {
		if task.TaskType == "WAIT" && task.Status == model.InProgressTask {
			logger.Info("Found WAIT task to complete", zap.String("task_id", task.TaskId))

			outputData := map[string]interface{}{
				"completed_by":    "lifecycle_demo",
				"completion_time": time.Now().Unix(),
			}

			if err = workflowExecutor.UpdateTaskWithContext(ctx, task.TaskId, task.WorkflowInstanceId, model.CompletedTask, outputData); err != nil {
				logger.Error("Failed to complete wait task", zap.Error(err))
				return err
			}

			updatedTask, err := workflowExecutor.GetTaskWithContext(ctx, task.TaskId)
			if err != nil {
				logger.Error("Failed to get task", zap.Error(err))
				return err
			}

			logger.Info("Task status after completion", zap.String("status", string(updatedTask.Status)))

			logger.Info("Completed wait task manually")
			break
		}
	}

	// 6. Terminate workflow
	logger.Info("=== 6. Terminating workflow ===")
	if err = workflowExecutor.TerminateWithContext(ctx, workflowId, "Terminating for retry demo"); err != nil {
		logger.Error("Failed to terminate workflow", zap.Error(err))
		return err
	}
	logger.Info("Terminated workflow")

	// 7. Retry workflow
	logger.Info("=== 7. Retrying workflow ===")
	if err = workflowExecutor.RetryWithContext(ctx, workflowId, false); err != nil {
		logger.Error("Failed to retry workflow", zap.Error(err))
		return err
	}
	logger.Info("Retried workflow")

	// Try to pause quickly to prevent full completion
	logger.Info("Pausing workflow to prevent immediate completion")
	if err = workflowExecutor.PauseWithContext(ctx, workflowId); err != nil {
		logger.Error("Failed to pause after retry", zap.Error(err))
		return err
	}

	wf, err = workflowExecutor.GetWorkflowWithContext(ctx, workflowId, true)
	if err != nil {
		logger.Error("Failed to get workflow status", zap.Error(err))
		return err
	}
	logger.Info("Workflow status after retry", zap.String("status", string(wf.Status)))

	// 8. Terminate again for restart demo
	logger.Info("=== 8. Terminating for restart demo ===")
	// Check current status before trying to terminate
	wf, err = workflowExecutor.GetWorkflowWithContext(ctx, workflowId, true)
	if err != nil {
		logger.Error("Failed to get workflow status", zap.Error(err))
		return err
	}

	if wf.Status == model.CompletedWorkflow || wf.Status == model.FailedWorkflow {
		logger.Info("Workflow already in terminal state", zap.String("status", string(wf.Status)))
		logger.Info("Skipping termination - will restart from completed state")
	} else {
		if err = workflowExecutor.Terminate(workflowId, "Terminating for restart demo"); err != nil {
			logger.Error("Failed to terminate workflow", zap.Error(err))
			return err
		}
		logger.Info("Terminated workflow for restart")
	}

	// 9. Restart workflow
	logger.Info("=== 9. Restarting workflow ===")
	if err = workflowExecutor.RestartWithContext(ctx, workflowId, false); err != nil {
		logger.Error("Failed to restart workflow", zap.Error(err))
		return err
	}
	logger.Info("Restarted workflow")

	// 10. Rerun from specific task
	logger.Info("=== 10. Checking for rerun capability ===")
	wf, err = workflowExecutor.GetWorkflowWithContext(ctx, workflowId, true)
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

		var id string
		id, err = workflowExecutor.ReRunWithContext(ctx, workflowId, rerunRequest)
		if err != nil {
			logger.Error("Failed to rerun from task", zap.Error(err))
			return err
		} else {
			logger.Info("Rerun initiated from task", zap.String("id", id))
		}
	} else {
		logger.Warn("No tasks found in workflow", zap.String("workflow_id", workflowId))
	}

	// 11. Final termination with failure workflow trigger
	logger.Info("=== 11. Terminating with failure workflow trigger ===")
	if err = workflowExecutor.TerminateWithFailure(workflowId, "Final termination with failure workflow", true); err != nil {
		logger.Error("Failed to terminate with failure", zap.Error(err))
		return err
	}
	logger.Info("Terminated with failure workflow trigger")

	// 12. Check final status
	logger.Info("=== 12. Final status check ===")
	wf, err = workflowExecutor.GetWorkflowWithContext(ctx, workflowId, true)
	if err != nil {
		logger.Error("Failed to get final status", zap.Error(err))
		return err
	}

	logger.Info("Final workflow status",
		zap.String("workflow_id", workflowId),
		zap.String("status", string(wf.Status)),
		zap.String("reason_for_incompletion", wf.ReasonForIncompletion))

	// 13. Demonstrate RemoveWorkflow API
	logger.Info("=== 13. Demonstrating RemoveWorkflow API ===")

	logger.Info("Checking status of existing workflow before removal", zap.String("workflow_id", workflowId))
	wf, err = workflowExecutor.GetWorkflowWithContext(ctx, workflowId, true)
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

	// Wait for real-time monitoring to complete
	logger.Info("=== 14. Waiting for real-time monitoring to complete ===")
	select {
	case <-monitoringComplete:
		logger.Info("Real-time monitoring completed successfully")
	case <-time.After(5 * time.Second):
		logger.Warn("Real-time monitoring timed out (workflow may have completed before final operations)")
	}

	logger.Info("Removing existing workflow from execution history")
	err = workflowExecutor.RemoveWorkflowWithContext(ctx, workflowId)
	if err != nil {
		logger.Error("Failed to remove workflow", zap.Error(err))
		return err
	}
	logger.Info("Workflow removed successfully from execution history")

	// Verify removal - this should fail
	logger.Info("Verifying workflow removal by attempting to retrieve it")
	_, err = workflowExecutor.GetWorkflowWithContext(ctx, workflowId, true)
	if err != nil {
		logger.Info("Workflow successfully removed from history - retrieval failed as expected",
			zap.String("error", err.Error()))
	} else {
		logger.Warn("Unexpected: workflow still exists after removal")
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
