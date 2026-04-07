package worker

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/stretchr/testify/assert"
)

func newTestTaskRunner() *TaskRunner {
	apiClient := client.NewAPIClient(
		nil,
		settings.NewHttpSettings("http://localhost:0/invalid"),
	)
	return NewTaskRunnerWithApiClient(apiClient)
}

func TestUpdateTaskWithRetry_ContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	tr := newTestTaskRunner().WithBaseContext(ctx)

	taskResult := &model.TaskResult{
		TaskId:             "test-task-id",
		WorkflowInstanceId: "test-workflow-id",
	}

	err := tr.updateTaskWithRetry("test-task", taskResult)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context cancelled during task update retry")
}
