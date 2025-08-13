//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package worker

import (
	"context"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// TaskContext extends context.Context with Conductor task/workflow metadata.
type TaskContext interface {
	context.Context
	GetWorkflowInstanceID() string
	GetWorkflowType() string
	GetTaskID() string
	GetTaskType() string
	GetRetryCount() int
	GetRetriedTaskID() string
	GetPollCount() int
}

type workflowContext struct {
	context.Context
	workflowInstanceID string
	workflowType       string
	taskID             string
	taskType           string
	retryCount         int
	retriedTaskID      string
	pollCount          int
}

func (w *workflowContext) GetWorkflowInstanceID() string { return w.workflowInstanceID }
func (w *workflowContext) GetWorkflowType() string       { return w.workflowType }
func (w *workflowContext) GetTaskID() string             { return w.taskID }
func (w *workflowContext) GetTaskType() string           { return w.taskType }
func (w *workflowContext) GetRetryCount() int            { return w.retryCount }
func (w *workflowContext) GetRetriedTaskID() string      { return w.retriedTaskID }
func (w *workflowContext) GetPollCount() int             { return w.pollCount }

// getWorkflowContext builds a TaskContext with enriched metadata from a model.Task.
// The context.Context parameter should be the first parameter as per Go conventions.
func getWorkflowContext(parent context.Context, t *model.Task) TaskContext {
	if parent == nil {
		parent = context.Background()
	}
	if t == nil {
		return &workflowContext{Context: parent}
	}
	return &workflowContext{
		Context:            parent,
		workflowInstanceID: t.WorkflowInstanceId,
		workflowType:       t.WorkflowType,
		taskID:             t.TaskId,
		taskType:           t.TaskDefName,
		retryCount:         int(t.RetryCount),
		retriedTaskID:      t.RetriedTaskId,
		pollCount:          int(t.PollCount),
	}
}
