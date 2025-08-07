package main

import (
	hello_world "examples/hello_world/src"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/log"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"

	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

var (
	apiClient        = client.NewAPIClientFromEnv()
	taskRunner       = worker.NewTaskRunnerWithApiClient(apiClient)
	workflowExecutor = executor.NewWorkflowExecutor(apiClient)
)

func main() {
	// Start the Greet Worker. This worker will process "greet" tasks.
	taskRunner.StartWorker("greet", hello_world.Greet, 1, time.Millisecond*100)

	// This is used to register the Workflow, it's a one-time process. You can comment from here
	wf := hello_world.CreateWorkflow(workflowExecutor)
	err := wf.Register(true)
	if err != nil {
		log.Error("Failed to register workflow", "error", err)
		return
	}
	// Till Here after registering the workflow

	// Start the greetings workflow
	id, err := workflowExecutor.StartWorkflow(
		&model.StartWorkflowRequest{
			Name:    "greetings",
			Version: 1,
			Input: map[string]string{
				"name": "Gopher",
			},
		},
	)

	if err != nil {
		log.Error("Failed to start workflow", "error", err)
		return
	}
	log.Info("Started workflow", "id", id)

	// Get a channel to monitor the workflow execution -
	// Note: This is useful in case of short duration workflows that completes in few seconds.
	channel, _ := workflowExecutor.MonitorExecution(id)
	run := <-channel
	log.Info("Output of the workflow", "output", run.Output)
}
