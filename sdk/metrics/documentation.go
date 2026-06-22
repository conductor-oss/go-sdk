//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package metrics

type MetricDocumentation string

// Legacy metric documentation.
const (
	EXTERNAL_PAYLOAD_USED_DOC     MetricDocumentation = "Incremented each time external payload storage is used"
	TASK_ACK_ERROR_DOC            MetricDocumentation = "Task ack has encountered an exception"
	TASK_ACK_FAILED_DOC           MetricDocumentation = "Task ack failed"
	TASK_EXECUTE_ERROR_DOC        MetricDocumentation = "Execution error"
	TASK_EXECUTE_TIME_DOC         MetricDocumentation = "Time to execute a task (deprecated gauge, milliseconds; see task_execute_time_seconds histogram)"
	TASK_EXECUTION_QUEUE_FULL_DOC MetricDocumentation = "Counter to record execution queue has saturated"
	TASK_EXECUTION_STARTED_DOC    MetricDocumentation = "Incremented each time a polled task is dispatched to the worker function"
	TASK_PAUSED_DOC               MetricDocumentation = "Counter for number of times the task has been polled, when the worker has been paused"
	TASK_POLL_DOC                 MetricDocumentation = "Incremented each time polling is done"
	TASK_POLL_ERROR_DOC           MetricDocumentation = "Client error when polling for a task queue"
	TASK_POLL_TIME_DOC            MetricDocumentation = "Time to poll for a batch of tasks (deprecated gauge, seconds; see task_poll_time_seconds histogram)"
	TASK_RESULT_SIZE_DOC          MetricDocumentation = "Records output payload size of a task (deprecated name; see task_result_size_bytes)"
	TASK_UPDATE_ERROR_DOC         MetricDocumentation = "Task status cannot be updated back to server"
	TASK_UPDATE_TIME_DOC          MetricDocumentation = "Time to update for a task (deprecated gauge, milliseconds; see task_update_time_seconds histogram)"
	THREAD_UNCAUGHT_EXCEPTION_DOC MetricDocumentation = "Uncaught exceptions raised inside worker goroutines"
	WORKFLOW_START_ERROR_DOC      MetricDocumentation = "Counter for workflow start errors"
	WORKFLOW_INPUT_SIZE_DOC       MetricDocumentation = "Records input payload size of a workflow (deprecated name; see workflow_input_size_bytes)"
)

// Canonical metric documentation.
const (
	TASK_POLL_TIME_SECONDS_DOC          MetricDocumentation = "Task poll latency in seconds"
	TASK_EXECUTE_TIME_SECONDS_DOC       MetricDocumentation = "Task execution latency in seconds"
	TASK_UPDATE_TIME_SECONDS_DOC        MetricDocumentation = "Task update (result-report) latency in seconds"
	HTTP_API_CLIENT_REQUEST_SECONDS_DOC MetricDocumentation = "HTTP API client request latency in seconds"
	TASK_RESULT_SIZE_BYTES_DOC          MetricDocumentation = "Records output payload size of a task in bytes"
	WORKFLOW_INPUT_SIZE_BYTES_DOC       MetricDocumentation = "Records input payload size of a workflow in bytes"
	ACTIVE_WORKERS_DOC                  MetricDocumentation = "Current number of worker goroutines actively executing a task"
)
