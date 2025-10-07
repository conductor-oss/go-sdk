package workflow

// GetWorkflowTask Task to get a workflow
type GetWorkflowTask struct {
	Task
}

// NewGetWorkflowTask Create a new Get Workflow Task
func NewGetWorkflowTask(taskRefName string, workflowID string, includeTasks bool) *GetWorkflowTask {
	return &GetWorkflowTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			taskType:          GET_WORKFLOW,
			inputParameters: map[string]interface{}{
				"id":           workflowID,
				"includeTasks": includeTasks,
			},
		},
	}
}
