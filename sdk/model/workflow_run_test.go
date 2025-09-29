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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflowRunStatusMethods(t *testing.T) {
	tests := []struct {
		status       WorkflowStatus
		isRunning    bool
		isPaused     bool
		isComplete   bool
		isFailed     bool
		isTimedOut   bool
		isTerminated bool
	}{
		{RunningWorkflow, true, false, false, false, false, false},
		{PausedWorkflow, false, true, false, false, false, false},
		{CompletedWorkflow, false, false, true, false, false, false},
		{FailedWorkflow, false, false, false, true, false, false},
		{TimedOutWorkflow, false, false, false, false, true, false},
		{TerminatedWorkflow, false, false, false, false, false, true},
	}

	for _, tt := range tests {
		t.Run(string(tt.status), func(t *testing.T) {
			w := &WorkflowRun{Status: tt.status}

			assert.Equal(t, tt.isRunning, w.IsRunning())
			assert.Equal(t, tt.isPaused, w.IsPaused())
			assert.Equal(t, tt.isComplete, w.IsCompleted())
			assert.Equal(t, tt.isFailed, w.IsFailed())
			assert.Equal(t, tt.isTimedOut, w.IsTimedOut())
			assert.Equal(t, tt.isTerminated, w.IsTerminated())
		})
	}
}

func TestWorkflowRunGetTaskByReferenceName(t *testing.T) {
	w := &WorkflowRun{
		Tasks: []Task{
			{TaskId: "id1", ReferenceTaskName: "task1", Status: CompletedTask},
			{TaskId: "id2", ReferenceTaskName: "task2", Status: InProgressTask},
		},
	}

	// Found
	task := w.GetTaskByReferenceName("task1")
	assert.NotNil(t, task)
	assert.Equal(t, "id1", task.TaskId)

	// Not found
	task = w.GetTaskByReferenceName("nonexistent")
	assert.Nil(t, task)
}

func TestWorkflowRunGetTasksByStatus(t *testing.T) {
	w := &WorkflowRun{
		Tasks: []Task{
			{ReferenceTaskName: "task1", Status: CompletedTask},
			{ReferenceTaskName: "task2", Status: InProgressTask},
			{ReferenceTaskName: "task3", Status: FailedTask},
			{ReferenceTaskName: "task4", Status: FailedWithTerminalErrorTask},
		},
	}

	// Single status
	completed := w.GetTasksByStatus(CompletedTask)
	assert.Len(t, completed, 1)
	assert.Equal(t, "task1", completed[0].ReferenceTaskName)

	// Multiple statuses
	failed := w.GetTasksByStatus(FailedTask, FailedWithTerminalErrorTask)
	assert.Len(t, failed, 2)
	assert.Equal(t, "task3", failed[0].ReferenceTaskName)
	assert.Equal(t, "task4", failed[1].ReferenceTaskName)
}

func TestWorkflowRunConvenienceMethods(t *testing.T) {
	w := &WorkflowRun{
		Tasks: []Task{
			{ReferenceTaskName: "task1", Status: CompletedTask},
			{ReferenceTaskName: "task2", Status: InProgressTask},
			{ReferenceTaskName: "task3", Status: FailedTask},
			{ReferenceTaskName: "task4", Status: ScheduledTask},
			{ReferenceTaskName: "task5", Status: FailedWithTerminalErrorTask},
		},
	}

	assert.Len(t, w.GetCompletedTasks(), 1)
	assert.Len(t, w.GetFailedTasks(), 2) // Both FailedTask and FailedWithTerminalErrorTask
}

func TestWorkflowRunEmptyWorkflow(t *testing.T) {
	w := &WorkflowRun{Tasks: []Task{}}
	assert.Len(t, w.GetCompletedTasks(), 0)
	assert.Len(t, w.GetFailedTasks(), 0)
	assert.Nil(t, w.GetTaskByReferenceName("any"))
}

func TestWorkflowRunGetFailedTasks(t *testing.T) {
	w := &WorkflowRun{
		Tasks: []Task{
			{ReferenceTaskName: "task1", Status: FailedTask},
			{ReferenceTaskName: "task2", Status: FailedWithTerminalErrorTask},
			{ReferenceTaskName: "task3", Status: CompletedTask},
			{ReferenceTaskName: "task4", Status: InProgressTask},
		},
	}

	failedTasks := w.GetFailedTasks()
	assert.Len(t, failedTasks, 2)

	// Check that both FailedTask and FailedWithTerminalErrorTask are included
	statuses := make(map[TaskResultStatus]bool)
	for _, task := range failedTasks {
		statuses[task.Status] = true
	}
	assert.True(t, statuses[FailedTask])
	assert.True(t, statuses[FailedWithTerminalErrorTask])
}

func TestWorkflowRunGetCompletedTasks(t *testing.T) {
	w := &WorkflowRun{
		Tasks: []Task{
			{ReferenceTaskName: "task1", Status: CompletedTask},
			{ReferenceTaskName: "task2", Status: CompletedTask},
			{ReferenceTaskName: "task3", Status: FailedTask},
			{ReferenceTaskName: "task4", Status: InProgressTask},
		},
	}

	completedTasks := w.GetCompletedTasks()
	assert.Len(t, completedTasks, 2)

	// Verify all returned tasks have CompletedTask status
	for _, task := range completedTasks {
		assert.Equal(t, CompletedTask, task.Status)
	}
}
