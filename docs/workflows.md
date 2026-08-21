# Authoring Workflows with the Go SDK

## A simple two-step workflow

```go

//API client instance with server address and authentication details
apiClient := client.NewAPIClient(
    settings.NewAuthenticationSettings(
        KEY,
        SECRET,
    ),
    settings.NewHttpSettings(
        "https://play.orkes.io/api",
    ))

//Create new workflow executor
executor := executor.NewWorkflowExecutor(apiClient)

//Create a new ConductorWorkflow instance
conductorWorkflow := workflow.NewConductorWorkflow(executor).
    Name("my_first_workflow").
    Version(1).
    OwnerEmail("developers@orkes.io")

//now, let's add a couple of simple tasks
conductorWorkflow.
	Add(workflow.NewSimpleTask("simple_task_2", "simple_task_1")).
    Add(workflow.NewSimpleTask("simple_task_1", "simple_task_2"))

//Register the workflow with server
conductorWorkflow.Register(true)        //Overwrite the existing definition with the new one
```
### Execute Workflow

#### Using Workflow Executor to start previously registered workflow
```go
//Input can be either a map or a struct that is serializable to a JSON map
workflowInput := map[string]interface{}{}

workflowId, err := executor.StartWorkflow(&model.StartWorkflowRequest{
    Name:  conductorWorkflow.GetName(),
    Input: workflowInput,
})
```

#### Using Workflow Executor to synchronously execute a workflow and get the output as a result
```go
//Input can be either a map or a struct that is serializable to a JSON map
workflowInput := map[string]interface{}{}

workflowRun, err := executor.ExecuteWorkflow(&model.StartWorkflowRequest{Name: wf.GetName(), Version: &version, Input: workflowInput}, "")
//workfowRun is a struct that contains the output of the workflow execution
type WorkflowRun struct {
	CorrelationId        string                 `json:"correlationId,omitempty"`
	CreateTime           int64                  `json:"createTime,omitempty"`
	CreatedBy            string                 `json:"createdBy,omitempty"`
	Input                map[string]interface{} `json:"input,omitempty"`
	Output               map[string]interface{} `json:"output,omitempty"`
	Priority             int32                  `json:"priority,omitempty"`
	RequestId            string                 `json:"requestId,omitempty"`
	ResponseType         ReturnStrategy         `json:"responseType,omitempty"`
	Status               WorkflowStatus         `json:"status,omitempty"`
	TargetWorkflowId     string                 `json:"targetWorkflowId,omitempty"`
	TargetWorkflowStatus WorkflowStatus         `json:"targetWorkflowStatus,omitempty"`
	Tasks                []Task                 `json:"tasks,omitempty"`
	UpdateTime           int64                  `json:"updateTime,omitempty"`
	Variables            map[string]interface{} `json:"variables,omitempty"`
	WorkflowId           string                 `json:"workflowId,omitempty"`
}
```

#### WorkflowRun Helper Methods

`WorkflowRun` provides helper methods to check workflow status and manage tasks:

**Status Check Methods:**
```go
// Check workflow status
workflowRun.IsRunning()     // Returns true if status is RUNNING
workflowRun.IsPaused()      // Returns true if status is PAUSED  
workflowRun.IsFailed()      // Returns true if status is FAILED
workflowRun.IsCompleted()   // Returns true if status is COMPLETED
workflowRun.IsTimedOut()    // Returns true if status is TIMED_OUT
workflowRun.IsTerminated()  // Returns true if status is TERMINATED
```

**Task Management Methods:**
```go
// Get tasks by status
failedTasks := workflowRun.GetFailedTasks()           // Returns tasks with FAILED or FAILED_WITH_TERMINAL_ERROR status
completedTasks := workflowRun.GetCompletedTasks()     // Returns tasks with COMPLETED status
tasks := workflowRun.GetTasksByStatus(model.FailedTask, model.CompletedTask) // Returns tasks with specified statuses

// Get specific task
task := workflowRun.GetTaskByReferenceName("task_ref_name") // Returns task by reference name, nil if not found
```
**Note:** Synchronous workflow execution is useful for workflows that complete in few seconds at max.  For longer running workflows use `StartWorkflow` and use the Id of the workflow to monitor the output.

#### Using struct instance as workflow input
```go
type WorkflowInput struct {
    Name string
    Address []string
}
//...
workflowId, err := executor.StartWorkflow(&model.StartWorkflowRequest{
  Name:  conductorWorkflow.GetName(),
  Input: &WorkflowInput{
  Name: "John Doe",
  Address: []string{"street", "city", "zip"},
  },
})

// Get workflow execution details
workflow, err := executor.GetWorkflowWithContext(context.Background(), workflowId, true)
if err != nil {
    // Handle error
}
```

**Monitor Workflow Execution:**
```go
// Start monitoring workflow execution
executionChannel, err := executor.MonitorExecution(workflowId)
if err != nil {
    // Handle error
}

// Wait for completion
workflow, err := executor.WaitForWorkflowCompletionUntilTimeout(executionChannel, 60*time.Second)
```

#### Workflow Helper Methods

`Workflow` provides additional helper methods for status and task management:

**Status Check Methods:**
```go
workflow.IsRunning()     // Returns true if status is RUNNING
workflow.IsPaused()      // Returns true if status is PAUSED  
workflow.IsFailed()      // Returns true if status is FAILED
workflow.IsCompleted()   // Returns true if status is COMPLETED
workflow.IsTimedOut()    // Returns true if status is TIMED_OUT
workflow.IsTerminated()  // Returns true if status is TERMINATED
```

**Task Management Methods:**
```go
// Get tasks by status
inProgressTasks := workflow.GetInProgressTasks()         // Returns tasks with IN_PROGRESS status
failedTasks := workflow.GetFailedTasks()                // Returns tasks with FAILED or FAILED_WITH_TERMINAL_ERROR status
completedTasks := workflow.GetCompletedTasks()          // Returns tasks with COMPLETED status
scheduledTasks := workflow.GetScheduledTasks()          // Returns tasks with SCHEDULED status
tasks := workflow.GetTasksByStatus(model.FailedTask, model.CompletedTask) // Returns tasks with specified statuses

// Get specific task
task := workflow.GetTaskByReferenceName("task_ref_name") // Returns task by reference name, nil if not found
```

### Workflow Management APIs
Take a look at the [API Docs](https://pkg.go.dev/github.com/conductor-sdk/conductor-go/sdk/workflow/executor) fore more details on how to start, pause, resume, terminate, search and get workflow execution status.

### More Examples
You can find more examples at the following GitHub repository:

https://github.com/conductor-sdk/conductor-examples/tree/main/go-samples
