package worker

import (
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// Wait-based workflow workers with descriptive, purpose-driven naming

// InitializeWorkflowStateWorker handles workflow initialization in wait-based flows
// Task Name: "initialize_workflow_state"
func InitializeWorkflowStateWorker(task *model.Task) (interface{}, error) {
	return map[string]interface{}{
		"result":    "Workflow state initialized",
		"timestamp": time.Now().Unix(),
	}, nil
}

// ProcessMiddleStepWorker handles intermediate processing tasks
// Task Name: "process_middle_step"
func ProcessMiddleStepWorker(task *model.Task) (interface{}, error) {
	time.Sleep(2 * time.Second)
	return map[string]interface{}{
		"result":    "Middle step processing completed",
		"timestamp": time.Now().Unix(),
	}, nil
}
