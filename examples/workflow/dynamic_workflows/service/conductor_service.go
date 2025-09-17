package service

import (
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/log"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	"go.uber.org/zap"
)

// ConductorService encapsulates all Conductor SDK clients and workers
type ConductorService struct {
	Logger           *zap.Logger
	APIClient        *client.APIClient
	WorkflowExecutor *executor.WorkflowExecutor
	WorkflowClient   client.WorkflowClient
	TaskClient       client.TaskClient
	TaskRunner       *worker.TaskRunner
}

// WorkerConfig defines worker configuration
type WorkerConfig struct {
	TaskType     string
	Handler      func(*model.Task) (interface{}, error)
	Concurrency  int
	PollInterval time.Duration
	Description  string
}

// NewConductorService creates a new Conductor service instance
func NewConductorService(logger *zap.Logger) (*ConductorService, error) {
	// Set SDK logger
	log.SetLogger(log.NewZap(logger))

	// Initialize API client from environment variables
	apiClient := client.NewAPIClientFromEnv()

	// Create clients
	workflowExecutor := executor.NewWorkflowExecutor(apiClient)
	workflowClient := client.NewWorkflowClient(apiClient)
	taskClient := client.NewTaskClient(apiClient)
	taskRunner := worker.NewTaskRunnerWithApiClient(apiClient)

	return &ConductorService{
		Logger:           logger,
		APIClient:        apiClient,
		WorkflowExecutor: workflowExecutor,
		WorkflowClient:   workflowClient,
		TaskClient:       taskClient,
		TaskRunner:       taskRunner,
	}, nil
}

// StartWorkers starts all configured workers
func (cs *ConductorService) StartWorkers(configs []WorkerConfig) error {
	for _, config := range configs {
		cs.Logger.Info("Registering worker",
			zap.String("task_type", config.TaskType),
			zap.String("description", config.Description),
			zap.Int("concurrency", config.Concurrency))

		cs.TaskRunner.StartWorker(
			config.TaskType,
			config.Handler,
			config.Concurrency,
			config.PollInterval,
		)
	}

	cs.Logger.Info("All workers started successfully",
		zap.Int("worker_count", len(configs)))
	return nil
}
