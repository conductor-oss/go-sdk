// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package model

// WorkflowStateUpdate is the request body for the UpdateWorkflowAndTaskState endpoint.
type WorkflowStateUpdate struct {
	// TaskReferenceName the reference name of the task to update.
	TaskReferenceName string `json:"taskReferenceName,omitempty"`
	// TaskResult the result of the task to update.
	TaskResult *TaskResult `json:"taskResult,omitempty"`
	// Variables the variables to update.
	Variables map[string]interface{} `json:"variables,omitempty"`
}
