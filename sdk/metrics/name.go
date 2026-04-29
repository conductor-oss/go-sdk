//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package metrics

type MetricName string

// Legacy metric names. Emitted by the legacy MetricsCollector implementation.
const (
	EXTERNAL_PAYLOAD_USED     MetricName = "external_payload_used"
	TASK_EXECUTE_ERROR        MetricName = "task_execute_error"
	TASK_EXECUTE_TIME         MetricName = "task_execute_time"
	TASK_EXECUTION_QUEUE_FULL MetricName = "task_execution_queue_full"
	TASK_PAUSED               MetricName = "task_paused"
	TASK_POLL                 MetricName = "task_poll"
	TASK_POLL_ERROR           MetricName = "task_poll_error"
	TASK_POLL_TIME            MetricName = "task_poll_time"
	TASK_RESULT_SIZE          MetricName = "task_result_size"
	TASK_UPDATE_ERROR         MetricName = "task_update_error"
	TASK_UPDATE_TIME          MetricName = "task_update_time"
	THREAD_UNCAUGHT_EXCEPTION MetricName = "thread_uncaught_exceptions"
	WORKFLOW_INPUT_SIZE       MetricName = "workflow_input_size"
	WORKFLOW_START_ERROR      MetricName = "workflow_start_error"
)

// Canonical metric names. Emitted by the canonical MetricsCollector implementation.
const (
	TASK_POLL_TOTAL                  MetricName = "task_poll_total"
	TASK_POLL_ERROR_TOTAL            MetricName = "task_poll_error_total"
	TASK_EXECUTION_STARTED_TOTAL     MetricName = "task_execution_started_total"
	TASK_EXECUTE_ERROR_TOTAL         MetricName = "task_execute_error_total"
	TASK_UPDATE_ERROR_TOTAL          MetricName = "task_update_error_total"
	TASK_ACK_ERROR_TOTAL             MetricName = "task_ack_error_total"
	TASK_ACK_FAILED_TOTAL            MetricName = "task_ack_failed_total"
	TASK_EXECUTION_QUEUE_FULL_TOTAL  MetricName = "task_execution_queue_full_total"
	TASK_PAUSED_TOTAL                MetricName = "task_paused_total"
	THREAD_UNCAUGHT_EXCEPTIONS_TOTAL MetricName = "thread_uncaught_exceptions_total"
	EXTERNAL_PAYLOAD_USED_TOTAL      MetricName = "external_payload_used_total"
	WORKFLOW_START_ERROR_TOTAL       MetricName = "workflow_start_error_total"

	TASK_POLL_TIME_SECONDS          MetricName = "task_poll_time_seconds"
	TASK_EXECUTE_TIME_SECONDS       MetricName = "task_execute_time_seconds"
	TASK_UPDATE_TIME_SECONDS        MetricName = "task_update_time_seconds"
	HTTP_API_CLIENT_REQUEST_SECONDS MetricName = "http_api_client_request_seconds"

	TASK_RESULT_SIZE_BYTES    MetricName = "task_result_size_bytes"
	WORKFLOW_INPUT_SIZE_BYTES MetricName = "workflow_input_size_bytes"

	ACTIVE_WORKERS MetricName = "active_workers"
)
