package workflow

// YieldTask Task to yield a workflow
type YieldTask struct {
	Task
}

// NewYieldTask creates a new Yield Task.
func NewYieldTask(taskRefName string) *YieldTask {
	return &YieldTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			taskType:          YIELD,
		},
	}
}
