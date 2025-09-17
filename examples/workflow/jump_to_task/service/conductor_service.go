package service

import (
	"fmt"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/log"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	sdkworker "github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	"go.uber.org/zap"
)

// ConductorService provides direct access to all Conductor SDK dependencies
type ConductorService struct {
	// Direct access to core dependencies - no wrappers needed
	APIClient        *client.APIClient
	TaskRunner       *sdkworker.TaskRunner
	WorkflowExecutor *executor.WorkflowExecutor

	// WorkflowClient is actually a WorkflowResourceApiService under the hood
	// so we can cast it to access additional methods like JumpToTask
	WorkflowClient client.WorkflowClient

	// Logger for service-level operations
	Logger *zap.Logger
}

// NewConductorService creates a new ConductorService with all dependencies initialized
func NewConductorService(logger *zap.Logger) (*ConductorService, error) {
	// Initialize logger
	if logger == nil {
		logger = zap.Must(zap.NewProduction())
	}

	// Set SDK logger
	log.SetLogger(log.NewZap(logger))

	// Initialize API client
	// Initialize API client (simplified to use environment config)
	apiClient := client.NewAPIClientFromEnv()

	// Initialize core components
	workflowExecutor := executor.NewWorkflowExecutor(apiClient)
	workflowClient := client.NewWorkflowClient(apiClient) // This returns WorkflowResourceApiService
	taskRunner := sdkworker.NewTaskRunnerWithApiClient(apiClient)

	return &ConductorService{
		APIClient:        apiClient,
		TaskRunner:       taskRunner,
		WorkflowExecutor: workflowExecutor,
		WorkflowClient:   workflowClient,
		Logger:           logger,
	}, nil
}

// WorkerConfig holds configuration for a worker
type WorkerConfig struct {
	TaskType     string
	Handler      func(*model.Task) (interface{}, error)
	Concurrency  int
	PollInterval time.Duration
	Description  string
}

// StartWorkers is a utility function to start multiple workers at once
func (s *ConductorService) StartWorkers(workers []WorkerConfig) error {
	for _, worker := range workers {
		// Call TaskRunner.StartWorker directly
		if err := s.TaskRunner.StartWorker(worker.TaskType, worker.Handler, worker.Concurrency, worker.PollInterval); err != nil {
			s.Logger.Error("Failed to start worker", zap.Error(err))
			return fmt.Errorf("failed to start worker: %w", err)
		}

		s.Logger.Debug("Started worker",
			zap.String("task_type", worker.TaskType),
			zap.Int("concurrency", worker.Concurrency),
			zap.Duration("poll_interval", worker.PollInterval))
	}

	s.Logger.Info("Started all workers", zap.Int("worker_count", len(workers)))
	return nil
}

func (s *ConductorService) CheckedJumpResults(workflowId, scenario string) error {
	// Call WorkflowExecutor.GetWorkflow directly
	wf, err := s.WorkflowExecutor.GetWorkflow(workflowId, true)
	if err != nil {
		s.Logger.Error("Failed to get final workflow state", zap.Error(err))
		return fmt.Errorf("failed to get final workflow state: %w", err)
	}

	s.Logger.Info("Jump results analysis",
		zap.String("scenario", scenario),
		zap.String("workflow_id", workflowId),
		zap.String("final_status", string(wf.Status)),
		zap.Int("total_tasks", len(wf.Tasks)))

	var skippedTasks, executedTasks []string

	for _, task := range wf.Tasks {
		s.Logger.Debug("Task analysis",
			zap.String("task_ref", task.ReferenceTaskName),
			zap.String("status", string(task.Status)))

		switch string(task.Status) {
		case "SKIPPED":
			skippedTasks = append(skippedTasks, task.ReferenceTaskName)
		case "COMPLETED", "IN_PROGRESS":
			executedTasks = append(executedTasks, task.ReferenceTaskName)
		}
	}

	s.Logger.Info("Task execution summary",
		zap.String("scenario", scenario),
		zap.Strings("skipped_tasks", skippedTasks),
		zap.Strings("executed_tasks", executedTasks))

	return nil
}
