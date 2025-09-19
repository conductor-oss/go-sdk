//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

// RerunWorkflowRequest is the request body for the Rerun Workflow.
type RerunWorkflowRequest struct {
	// ReRunFromWorkflowId the unique identifier of the workflow to be rerun.
	ReRunFromWorkflowId string `json:"reRunFromWorkflowId,omitempty"`
	// WorkflowInput a map of inputs for the rerun workflow.
	WorkflowInput map[string]interface{} `json:"workflowInput,omitempty"`
	// ReRunFromTaskId the unique identifier of the task to rerun the workflow from.
	ReRunFromTaskId string `json:"reRunFromTaskId,omitempty"`
	// TaskInput a map of inputs for the rerun task.
	TaskInput map[string]interface{} `json:"taskInput,omitempty"`
	// CorrelationId the unique identifier used to correlate the current workflow execution with other executions of the same workflow.
	CorrelationId string `json:"correlationId,omitempty"`
}
