package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/metrics"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

const workflowName = "go_simulated_tasks_workflow"

type simulatedWorkerDef struct {
	taskName     string
	codename     string
	sleepSeconds int
}

var simulatedWorkers = []simulatedWorkerDef{
	{"go_worker_0", "quickpulse", 1},
	{"go_worker_1", "whisperlink", 2},
	{"go_worker_2", "shadowfetch", 3},
	{"go_worker_3", "ironforge", 4},
	{"go_worker_4", "deepcrawl", 5},
}

func main() {
	apiClient := client.NewAPIClientFromEnv()
	workflowExecutor := executor.NewWorkflowExecutor(apiClient)
	taskRunner := worker.NewTaskRunnerWithApiClient(apiClient)

	registerMetadata(apiClient, workflowExecutor)

	metricsPort := envIntOrDefault("HARNESS_METRICS_PORT", 9991)
	go metrics.ProvideMetrics(settings.NewMetricsSettings("/metrics", metricsPort))
	fmt.Printf("Prometheus metrics server started on port %d\n", metricsPort)

	workflowsPerSec := envIntOrDefault("HARNESS_WORKFLOWS_PER_SEC", 2)
	batchSize := envIntOrDefault("HARNESS_BATCH_SIZE", 20)
	pollIntervalMs := envIntOrDefault("HARNESS_POLL_INTERVAL_MS", 100)

	for _, def := range simulatedWorkers {
		w := NewSimulatedTaskWorker(def.taskName, def.codename, def.sleepSeconds, batchSize, pollIntervalMs)
		err := taskRunner.StartWorker(
			w.taskName,
			w.Execute,
			w.batchSize,
			w.pollInterval,
		)
		if err != nil {
			fmt.Printf("Failed to start worker %s: %v\n", def.taskName, err)
			os.Exit(1)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	taskRunner.WithBaseContext(ctx)

	governor := NewWorkflowGovernor(workflowExecutor, workflowName, workflowsPerSec)
	go governor.Run(ctx)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	fmt.Println("Shutting down...")
	cancel()
	for _, def := range simulatedWorkers {
		taskRunner.Shutdown(def.taskName)
	}
	taskRunner.WaitWorkers()
}

func registerMetadata(apiClient *client.APIClient, workflowExecutor *executor.WorkflowExecutor) {
	metadataClient := client.NewMetadataClient(apiClient)

	taskDefs := make([]model.TaskDef, 0, len(simulatedWorkers))
	for _, def := range simulatedWorkers {
		taskDefs = append(taskDefs, model.TaskDef{
			Name:                   def.taskName,
			Description:            fmt.Sprintf("Go SDK harness simulated task (%s, default delay %ds)", def.codename, def.sleepSeconds),
			RetryCount:             1,
			TimeoutSeconds:         300,
			ResponseTimeoutSeconds: 300,
		})
	}
	_, err := metadataClient.RegisterTaskDef(context.Background(), taskDefs)
	if err != nil {
		fmt.Printf("Failed to register task definitions: %v\n", err)
		os.Exit(1)
	}

	wf := workflow.NewConductorWorkflow(workflowExecutor).
		Name(workflowName).
		Version(1).
		Description("Go SDK harness simulated task workflow").
		OwnerEmail("go-sdk-harness@conductor.io")

	for _, def := range simulatedWorkers {
		wf.Add(workflow.NewSimpleTask(def.taskName, def.codename))
	}

	err = wf.Register(true)
	if err != nil {
		fmt.Printf("Failed to register workflow: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Registered workflow %s with %d tasks\n", workflowName, len(simulatedWorkers))
}

func envIntOrDefault(key string, defaultVal int) int {
	s := os.Getenv(key)
	if s == "" {
		return defaultVal
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return defaultVal
	}
	return v
}
