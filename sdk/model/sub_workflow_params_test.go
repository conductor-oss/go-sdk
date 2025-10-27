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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkflowDefinitionParam_UnmarshalJSON_Expression(t *testing.T) {
	jsonData := []byte(`"${workflow.input.wf_def}"`)

	var param WorkflowDefinitionParam
	err := json.Unmarshal(jsonData, &param)

	require.NoError(t, err)
	assert.True(t, param.IsExpression())
	assert.False(t, param.IsDefinition())

	expr, ok := param.GetExpression()
	assert.True(t, ok)
	assert.Equal(t, "${workflow.input.wf_def}", expr)

	def, ok := param.GetDefinition()
	assert.False(t, ok)
	assert.Nil(t, def)
}

func TestWorkflowDefinitionParam_UnmarshalJSON_WorkflowDef(t *testing.T) {
	jsonData := []byte(`{
		"name": "test_workflow",
		"version": 1,
		"description": "Test workflow"
	}`)

	var param WorkflowDefinitionParam
	err := json.Unmarshal(jsonData, &param)

	require.NoError(t, err)
	assert.False(t, param.IsExpression())
	assert.True(t, param.IsDefinition())

	expr, ok := param.GetExpression()
	assert.False(t, ok)
	assert.Equal(t, "", expr)

	def, ok := param.GetDefinition()
	assert.True(t, ok)
	assert.NotNil(t, def)
	assert.Equal(t, "test_workflow", def.Name)
	assert.Equal(t, int32(1), def.Version)
	assert.Equal(t, "Test workflow", def.Description)
}

func TestWorkflowDefinitionParam_MarshalJSON_Expression(t *testing.T) {
	param := NewWorkflowDefinitionExpression("${workflow.input.wf_def}")

	data, err := json.Marshal(param)

	require.NoError(t, err)
	assert.JSONEq(t, `"${workflow.input.wf_def}"`, string(data))
}

func TestWorkflowDefinitionParam_MarshalJSON_WorkflowDef(t *testing.T) {
	workflowDef := &WorkflowDef{
		Name:        "test_workflow",
		Version:     1,
		Description: "Test workflow",
	}
	param := NewWorkflowDefinitionObject(workflowDef)

	data, err := json.Marshal(param)

	require.NoError(t, err)
	assert.Contains(t, string(data), `"name":"test_workflow"`)
	assert.Contains(t, string(data), `"version":1`)
}

func TestWorkflowDefinitionParam_MarshalJSON_Nil(t *testing.T) {
	var param *WorkflowDefinitionParam

	data, err := json.Marshal(param)

	require.NoError(t, err)
	assert.Equal(t, "null", string(data))
}

func TestWorkflowDefinitionParam_Constructors(t *testing.T) {
	t.Run("NewWorkflowDefinitionExpression", func(t *testing.T) {
		param := NewWorkflowDefinitionExpression("${workflow.input.test}")

		assert.True(t, param.IsExpression())
		assert.False(t, param.IsDefinition())

		expr, ok := param.GetExpression()
		assert.True(t, ok)
		assert.Equal(t, "${workflow.input.test}", expr)
	})

	t.Run("NewWorkflowDefinitionObject", func(t *testing.T) {
		workflowDef := &WorkflowDef{Name: "test"}
		param := NewWorkflowDefinitionObject(workflowDef)

		assert.False(t, param.IsExpression())
		assert.True(t, param.IsDefinition())

		def, ok := param.GetDefinition()
		assert.True(t, ok)
		assert.Equal(t, "test", def.Name)
	})
}

func TestWorkflowDefinitionParam_NilHandling(t *testing.T) {
	var param *WorkflowDefinitionParam

	assert.False(t, param.IsExpression())
	assert.False(t, param.IsDefinition())

	expr, ok := param.GetExpression()
	assert.False(t, ok)
	assert.Equal(t, "", expr)

	def, ok := param.GetDefinition()
	assert.False(t, ok)
	assert.Nil(t, def)
}

func TestSubWorkflowParams_UnmarshalJSON_WithExpression(t *testing.T) {
	// This is the actual JSON from the issue description
	jsonData := []byte(`{
		"name": "sub_workflow",
		"taskReferenceName": "sub_workflow_ref",
		"inputParameters": {},
		"type": "SUB_WORKFLOW",
		"subWorkflowParam": {
			"name": "some_workflow",
			"version": 1,
			"workflowDefinition": "${workflow.input.wf_def}"
		}
	}`)

	var task struct {
		SubWorkflowParam *SubWorkflowParams `json:"subWorkflowParam"`
	}

	err := json.Unmarshal(jsonData, &task)

	require.NoError(t, err)
	require.NotNil(t, task.SubWorkflowParam)
	assert.Equal(t, "some_workflow", task.SubWorkflowParam.Name)
	assert.Equal(t, int32(1), task.SubWorkflowParam.Version)

	require.NotNil(t, task.SubWorkflowParam.WorkflowDefinition)
	assert.True(t, task.SubWorkflowParam.WorkflowDefinition.IsExpression())

	expr, ok := task.SubWorkflowParam.WorkflowDefinition.GetExpression()
	assert.True(t, ok)
	assert.Equal(t, "${workflow.input.wf_def}", expr)
}

func TestSubWorkflowParams_UnmarshalJSON_WithWorkflowDef(t *testing.T) {
	jsonData := []byte(`{
		"name": "some_workflow",
		"version": 1,
		"workflowDefinition": {
			"name": "inline_workflow",
			"version": 2,
			"description": "Inline workflow definition"
		}
	}`)

	var params SubWorkflowParams
	err := json.Unmarshal(jsonData, &params)

	require.NoError(t, err)
	assert.Equal(t, "some_workflow", params.Name)
	assert.Equal(t, int32(1), params.Version)

	require.NotNil(t, params.WorkflowDefinition)
	assert.True(t, params.WorkflowDefinition.IsDefinition())

	def, ok := params.WorkflowDefinition.GetDefinition()
	assert.True(t, ok)
	assert.Equal(t, "inline_workflow", def.Name)
	assert.Equal(t, int32(2), def.Version)
}

func TestSubWorkflowParams_MarshalJSON_RoundTrip(t *testing.T) {
	t.Run("With Expression", func(t *testing.T) {
		params := SubWorkflowParams{
			Name:               "test_workflow",
			Version:            1,
			WorkflowDefinition: NewWorkflowDefinitionExpression("${workflow.input.wf_def}"),
		}

		data, err := json.Marshal(params)
		require.NoError(t, err)

		var unmarshaled SubWorkflowParams
		err = json.Unmarshal(data, &unmarshaled)
		require.NoError(t, err)

		assert.Equal(t, params.Name, unmarshaled.Name)
		assert.Equal(t, params.Version, unmarshaled.Version)
		assert.True(t, unmarshaled.WorkflowDefinition.IsExpression())

		expr, ok := unmarshaled.WorkflowDefinition.GetExpression()
		assert.True(t, ok)
		assert.Equal(t, "${workflow.input.wf_def}", expr)
	})

	t.Run("With WorkflowDef", func(t *testing.T) {
		params := SubWorkflowParams{
			Name:    "test_workflow",
			Version: 1,
			WorkflowDefinition: NewWorkflowDefinitionObject(&WorkflowDef{
				Name:        "inline",
				Version:     2,
				Description: "Test",
			}),
		}

		data, err := json.Marshal(params)
		require.NoError(t, err)

		var unmarshaled SubWorkflowParams
		err = json.Unmarshal(data, &unmarshaled)
		require.NoError(t, err)

		assert.Equal(t, params.Name, unmarshaled.Name)
		assert.Equal(t, params.Version, unmarshaled.Version)
		assert.True(t, unmarshaled.WorkflowDefinition.IsDefinition())

		def, ok := unmarshaled.WorkflowDefinition.GetDefinition()
		assert.True(t, ok)
		assert.Equal(t, "inline", def.Name)
	})
}
