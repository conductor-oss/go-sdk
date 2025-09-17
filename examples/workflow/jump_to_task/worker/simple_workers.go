package worker

import (
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// Simple workflow workers with clear, descriptive naming

// ProcessInitialDataWorker handles the initial data processing task
// Task Name: "process_initial_data"
func ProcessInitialDataWorker(task *model.Task) (interface{}, error) {
	time.Sleep(1 * time.Second)
	return map[string]interface{}{
		"processedData": "Data processed in initial step",
		"timestamp":     time.Now().Unix(),
		"status":        "completed",
	}, nil
}

// ValidateBusinessDataWorker handles data validation tasks
// Task Name: "validate_business_data"
func ValidateBusinessDataWorker(task *model.Task) (interface{}, error) {
	time.Sleep(2 * time.Second)
	return map[string]interface{}{
		"validatedData": "Data validated successfully",
		"isValid":       true,
		"timestamp":     time.Now().Unix(),
	}, nil
}

// ApplyBusinessLogicWorker applies business logic processing
// Task Name: "apply_business_logic"
func ApplyBusinessLogicWorker(task *model.Task) (interface{}, error) {
	time.Sleep(2 * time.Second)
	return map[string]interface{}{
		"businessResult": "Business logic applied",
		"score":          95,
		"timestamp":      time.Now().Unix(),
	}, nil
}

// CompleteFinalProcessWorker handles final processing tasks
// Task Name: "complete_final_process"
func CompleteFinalProcessWorker(task *model.Task) (interface{}, error) {
	time.Sleep(1 * time.Second)
	return map[string]interface{}{
		"finalResult": "Final processing completed",
		"summary":     "Workflow completed successfully",
		"timestamp":   time.Now().Unix(),
	}, nil
}
