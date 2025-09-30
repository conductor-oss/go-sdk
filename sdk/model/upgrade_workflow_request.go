// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package model

// UpgradeWorkflowRequest is the request body for the UpgradeRunningWorkflowToVersion.
type UpgradeWorkflowRequest struct {
	// Name the name of the workflow definition.
	Name string `json:"name"`
	// TaskOutput a map of task outputs for any skipped tasks,
	// with the key as the task reference name, and the value as the task output object.
	TaskOutput map[string]interface{} `json:"taskOutput,omitempty"`
	// Version the version to which the workflow is to be updated.
	Version int32 `json:"version,omitempty"`
	// WorkflowInput a map of inputs for the upgraded workflow execution,
	// with the parameter name as the key and its input value as the value.
	WorkflowInput map[string]interface{} `json:"workflowInput,omitempty"`
}
