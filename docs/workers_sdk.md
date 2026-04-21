# Writing Workers with the Go SDK

A worker is responsible for executing a task. 
Operator and System tasks are handled by the Conductor server, while user defined tasks needs to have a worker created that awaits the work to be scheduled by the server for it to be executed.

Worker framework provides features such as polling threads, metrics and server communication.

### Design Principles for Workers
Each worker embodies design pattern and follows certain basic principles:

1. Workers are stateless and do not implement a workflow specific logic. 
2. Each worker executes a very specific task and produces well-defined output given specific inputs. 
3. Workers are meant to be idempotent (or should handle cases where the task that partially executed gets rescheduled due to timeouts etc.)
4. Workers do not implement the logic to handle retries etc, that is taken care by the Conductor server.

### Creating Task Workers
Task worker is implemented using a function that confirms to the following function
```go
type ExecuteTaskFunction func(t *Task) (interface{}, error)
```

Worker returns a struct as the output of the task execution.  The struct MUST be serializable to a JSON map.
If an `error` is returned, the task is marked as `FAILED`

#### Task worker that returns a struct

```go

//TaskOutput struct that represents the output of the task execution
type TaskOutput struct {
    Keys    []string
    Message string
    Value   float64
}

//SimpleWorker function accepts Task as input and returns TaskOutput as result
//If there is a failure, error can be returned and the task will be marked as FAILED
func SimpleWorker(t *model.Task) (interface{}, error) {
    taskResult := &TaskOutput{
        Keys:    []string{"Key1", "Key2"},
        Message: "Hello World",
        Value:   rand.ExpFloat64(),
    }
    return taskResult, nil
}
```

#### Controlling execution for long-running tasks
For the long-running tasks you might want to spawn another process/routine and update the status of the task at a later point and complete the
execution function without actually marking the task as `COMPLETED`.  Use `TaskResult` struct that allows you to specify more fined grained control.

Here is an example of a task execution function that returns with `IN_PROGRESS` status asking server to push the task again in 60 seconds.
```go
func LongRunningTaskWorker(t *model.Task) (interface{}, error) {
	taskResult := model.NewTaskResult(t)
	taskResult.OutputData = map[string]interface{}{}
    
	//Keep the status as IN_PROGRESS
	taskResult.Status = task_result_status.IN_PROGRESS
	//Time after which the task should be sent back to worker
	taskResult.CallbackAfterSeconds = 60
	return taskResult, nil
}
```

## Type-safe Worker API and contextual execution (Go 1.23+)

The SDK includes a new worker API that adds:

- Type-safe handlers via generics: `TypedWorker[TIn, TOut]`
- Contextual execution: `TaskContext` with workflow/task metadata
- Clear per-task configuration via options: `WithBatchSize`, `WithPollInterval`, `WithPollTimeout`, `WithDomain`, `WithBaseContext`
- Unified registration using `RegisterWorker` / `RegisterWorkers`

### Worker with options

```go
api := client.NewAPIClientFromEnv()
runner := worker.NewTaskRunnerWithApiClient(api)

w := worker.NewWorker(
    "greet",
    func(t *model.Task) (interface{}, error) {
        name := fmt.Sprintf("%v", t.InputData["person_to_be_greated"]) // map input to your task
        return map[string]any{"hello": "Hello, " + name}, nil
    },
    worker.WithBatchSize(2),
    worker.WithPollInterval(250*time.Millisecond),
    worker.WithPollTimeout(5*time.Second), // negative uses server default; zero leaves unchanged
    worker.WithDomain("dev"),
)

if err := runner.RegisterWorker(w); err != nil {
    panic(err)
}

runner.WaitWorkers()
```

### TypedWorker with structured I/O and TaskContext

```go
type GreetIn struct {
    Name string `json:"person_to_be_greated"`
}

type GreetOut struct {
    Hello string `json:"hello"`
}

api := client.NewAPIClientFromEnv()
runner := worker.NewTaskRunnerWithApiClient(api)

tw := worker.NewTypedWorker[GreetIn, GreetOut](
    "greet",
    func(ctx worker.TaskContext, in GreetIn) (GreetOut, error) {
        // Access metadata when needed
        _ = ctx.GetWorkflowInstanceID()
        _ = ctx.GetTaskType()
        return GreetOut{Hello: "Hello, " + in.Name}, nil
    },
    worker.WithBatchSize(1),
    worker.WithPollInterval(100*time.Millisecond),
)

if err := runner.RegisterWorker(tw); err != nil {
    panic(err)
}

runner.WaitWorkers()
```

Prefer `NewSimpleTypedWorker` if you want a `func(context.Context, TIn)` signature.

### Register multiple workers

```go
err := runner.RegisterWorkers(
    worker.NewWorker("a", func(t *model.Task) (interface{}, error) { return map[string]any{"ok": true}, nil }),
    worker.NewTypedWorker[In, Out]("b", func(ctx worker.TaskContext, in In) (Out, error) { return Out{}, nil }),
)
if err != nil { panic(err) }
```

### TaskContext reference

`TaskContext` extends `context.Context` and exposes:

- `WorkflowInstanceID() string`
- `WorkflowType() string`
- `TaskID() string`
- `TaskType() string`
- `RetryCount() int`
- `RetriedTaskID() string`
- `PollCount() int`

## Starting Workers
`TaskRunner` interface is used to start the workers, which takes care of polling server for the work, executing worker code and updating the results back to the server.

```go
apiClient := client.NewAPIClient(
    settings.NewAuthenticationSettings(
        KEY,
        SECRET,
    ),
    settings.NewHttpSettings(
    "https://play.orkes.io/api",
))

taskRunner := worker.NewTaskRunnerWithApiClient(apiClient)
//Start polling for a task by name "simple_task", with a batch size of 1 and 1 second interval
//Between polls if there are no tasks available to execute
taskRunner.StartWorker("simple_task", examples.SimpleWorker, 1, time.Second*1)
//Add more StartWorker calls as needed

//Block
taskRunner.WaitWorkers()
```

## Task Management APIs

### Get Task Details
```go
task, err := executor.GetTask(taskId)
```

### Updating the Task result outside the worker implementation
#### Update task by id
```go
output :=  &TaskOutput{
Keys:    []string{"Key1", "Key2"},
Message: "Hello World",
Value:   rand.ExpFloat64(),
}
executor.UpdateTask(taskId, workflowInstanceId, task_result_status.COMPLETED, output)
```

#### Update task by Reference Name
```go
output :=  &TaskOutput{
Keys:    []string{"Key1", "Key2"},
Message: "Hello World",
Value:   rand.ExpFloat64(),
}
executor.UpdateTaskByRefName("task_ref_name", workflowInstanceId, task_result_status.COMPLETED, output)
```

### Worker Metrics
We use [Prometheus](https://prometheus.io/) to collect metrics.
When enabled the worker starts an HTTP server which is used to publish metrics, which can be hooked up to a prometheus server to scrap and collect metrics.

#### Starting metrics collection
```go
//Start a go routine.  The default settings  exposes port 2112 on /metrics endpoint
go ProvideMetrics(settings.NewDefaultMetricsSettings())
```

Worker SDK emits both legacy metric names (kept for backward compatibility)
and canonical metric names that match the cross-SDK catalog in
[`longrunning-wfstest/sdk-metrics-harmonization.md`](https://github.com/conductor-oss/longrunning-wfstest/blob/main/sdk-metrics-harmonization.md).
Both sets are emitted by every worker; prefer the canonical names for new
dashboards.

#### Canonical metrics

| Name | Type | Labels | Purpose |
| --- | --- | --- | --- |
| `task_poll_total` | Counter | `taskType` | Incremented for every poll request. |
| `task_poll_error_total` | Counter | `taskType`, `exception` | Poll request failed client-side. `exception` is the Go type name. |
| `task_execution_started_total` | Counter | `taskType` | Polled task dispatched to the user worker function. |
| `task_execute_error_total` | Counter | `taskType`, `exception` | Worker function returned a non-nil error. |
| `task_update_error_total` | Counter | `taskType`, `exception` | Task-result update back to the server failed. |
| `task_ack_error_total` | Counter | `taskType`, `exception` | Exception while acknowledging a polled task. |
| `task_ack_failed_total` | Counter | `taskType` | Server returned a non-success ack. |
| `task_execution_queue_full_total` | Counter | `taskType` | Worker's executor queue saturated. |
| `task_paused_total` | Counter | `taskType` | Poll happened while worker was paused. |
| `thread_uncaught_exceptions_total` | Counter | `exception` | Uncaught panic inside a worker goroutine. |
| `external_payload_used_total` | Counter | `entityName`, `operation`, `payload_type` | External payload storage used. |
| `workflow_start_error_total` | Counter | `workflowType`, `exception` | `StartWorkflow` failed client-side. |
| `task_poll_time_seconds` | Histogram | `taskType`, `status` | Task poll latency (seconds). Standard buckets 1ms..10s. |
| `task_execute_time_seconds` | Histogram | `taskType`, `status` | Task execution latency. |
| `task_update_time_seconds` | Histogram | `taskType`, `status` | Task-update latency. |
| `http_api_client_request_seconds` | Histogram | `method`, `uri`, `status` | Every HTTP request made by the SDK's generated API client. |
| `task_result_size_bytes` | Gauge | `taskType` | Last-seen task-result payload size. |
| `workflow_input_size_bytes` | Gauge | `workflowType`, `version` | Last-seen workflow-input payload size. |

Querying percentiles across replicas:

```promql
histogram_quantile(
  0.95,
  sum by (le, taskType) (rate(task_execute_time_seconds_bucket[5m]))
)
```

#### Deprecated metrics

These are still emitted for backward compatibility and will be removed in a
future major release. Prefer the canonical replacements above.

| Deprecated | Replacement | Notes |
| --- | --- | --- |
| `task_poll` (Counter) | `task_poll_total` | Identical semantics. |
| `task_poll_error` (Counter) | `task_poll_error_total` | Canonical adds the `exception` label. |
| `task_execute_error` (Counter) | `task_execute_error_total` | Canonical adds the `exception` label. |
| `task_update_error` (Counter) | `task_update_error_total` | Canonical adds the `exception` label. |
| `task_execution_queue_full` (Counter) | `task_execution_queue_full_total` | Rename-alias. |
| `task_paused` (Counter) | `task_paused_total` | Rename-alias. |
| `thread_uncaught_exceptions` (Counter, no labels) | `thread_uncaught_exceptions_total{exception}` | Canonical populates the previously dropped `message` argument as the `exception` label. |
| `external_payload_used` (Counter) | `external_payload_used_total` | Rename-alias. |
| `workflow_start_error` (Counter) | `workflow_start_error_total` | Canonical adds the `exception` label. |
| `task_poll_time` (Gauge, seconds) | `task_poll_time_seconds` (Histogram) | Gauge kept for dashboards; Histogram aggregates across replicas. |
| `task_execute_time` (Gauge, **milliseconds**) | `task_execute_time_seconds` (Histogram, seconds) | Legacy Gauge emits the wrong unit for historical reasons; the Histogram uses seconds as advertised. |
| `task_update_time` (Gauge, **milliseconds**) | `task_update_time_seconds` (Histogram, seconds) | Same unit note as above. |
| `task_result_size` (Gauge) | `task_result_size_bytes` (Gauge) | Same value, canonical name. |
| `workflow_input_size` (Gauge) | `workflow_input_size_bytes` (Gauge) | Same value, canonical name. |

#### Go runtime metrics

The Prometheus default gatherer also exposes the standard `go_*` and
`process_*` runtime metrics automatically.

Metrics on client side supplement the ones collected from server in
identifying network as well as client-side issues.

### Next: [Create and Execute Workflows](workflow_sdk.md)
