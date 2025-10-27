//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

import (
	"encoding/json"
	"fmt"
)

type SubWorkflowParams struct {
	Name               string                   `json:"name"`
	Version            int32                    `json:"version,omitempty"`
	TaskToDomain       map[string]string        `json:"taskToDomain,omitempty"`
	WorkflowDefinition *WorkflowDefinitionParam `json:"workflowDefinition,omitempty"`
}

// WorkflowDefinitionParam represents a workflow definition that can be either
// a string expression (e.g., "${workflow.input.wf_def}") or an actual WorkflowDef object.
type WorkflowDefinitionParam struct {
	expression *string
	definition *WorkflowDef
}

// NewWorkflowDefinitionExpression creates a WorkflowDefinitionParam from a string expression.
func NewWorkflowDefinitionExpression(expr string) *WorkflowDefinitionParam {
	return &WorkflowDefinitionParam{
		expression: &expr,
	}
}

// NewWorkflowDefinitionParam creates a WorkflowDefinitionParam from a WorkflowDef object.
func NewWorkflowDefinitionParam(def *WorkflowDef) *WorkflowDefinitionParam {
	return &WorkflowDefinitionParam{
		definition: def,
	}
}

// IsExpression returns true if this parameter contains a string expression.
func (w *WorkflowDefinitionParam) IsExpression() bool {
	return w != nil && w.expression != nil
}

// IsDefinition returns true if this parameter contains a WorkflowDef object.
func (w *WorkflowDefinitionParam) IsDefinition() bool {
	return w != nil && w.definition != nil
}

// GetExpression returns the string expression and true if this is an expression,
// otherwise returns empty string and false.
func (w *WorkflowDefinitionParam) GetExpression() (string, bool) {
	if w != nil && w.expression != nil {
		return *w.expression, true
	}
	return "", false
}

// GetDefinition returns the WorkflowDef object and true if this is a definition,
// otherwise returns nil and false.
func (w *WorkflowDefinitionParam) GetDefinition() (*WorkflowDef, bool) {
	if w != nil && w.definition != nil {
		return w.definition, true
	}
	return nil, false
}

// UnmarshalJSON implements json.Unmarshaler to handle both string and object types.
func (w *WorkflowDefinitionParam) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as string first (for expressions like "${workflow.input.wf_def}")
	var expr string
	if err := json.Unmarshal(data, &expr); err == nil {
		w.expression = &expr
		w.definition = nil
		return nil
	}

	// Try to unmarshal as WorkflowDef object
	var def WorkflowDef
	if err := json.Unmarshal(data, &def); err == nil {
		w.definition = &def
		w.expression = nil
		return nil
	}

	return fmt.Errorf("workflowDefinition must be either a string expression or a WorkflowDef object")
}

// MarshalJSON implements json.Marshaler to serialize based on the stored type.
func (w *WorkflowDefinitionParam) MarshalJSON() ([]byte, error) {
	if w == nil {
		return json.Marshal(nil)
	}

	if w.expression != nil {
		return json.Marshal(*w.expression)
	}

	if w.definition != nil {
		return json.Marshal(w.definition)
	}

	return json.Marshal(nil)
}
