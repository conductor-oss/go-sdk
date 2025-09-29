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

func TestWorkflowStatusMethods(t *testing.T) {
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
			w := &Workflow{Status: tt.status}

			assert.Equal(t, tt.isRunning, w.IsRunning())
			assert.Equal(t, tt.isPaused, w.IsPaused())
			assert.Equal(t, tt.isComplete, w.IsCompleted())
			assert.Equal(t, tt.isFailed, w.IsFailed())
			assert.Equal(t, tt.isTimedOut, w.IsTimedOut())
			assert.Equal(t, tt.isTerminated, w.IsTerminated())
		})
	}
}

func TestGetInProgressTasks(t *testing.T) {
	w := &Workflow{
		Tasks: []Task{
			{ReferenceTaskName: "task1", Status: InProgressTask},
			{ReferenceTaskName: "task2", Status: CompletedTask},
			{ReferenceTaskName: "task3", Status: InProgressTask},
		},
	}

	tasks := w.GetInProgressTasks()
	assert.Len(t, tasks, 2)
	assert.Equal(t, "task1", tasks[0].ReferenceTaskName)
	assert.Equal(t, "task3", tasks[1].ReferenceTaskName)
}

func TestGetTaskByReferenceName(t *testing.T) {
	w := &Workflow{
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

func TestGetTasksByStatus(t *testing.T) {
	w := &Workflow{
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

func TestConvenienceMethods(t *testing.T) {
	w := &Workflow{
		Tasks: []Task{
			{ReferenceTaskName: "task1", Status: CompletedTask},
			{ReferenceTaskName: "task2", Status: InProgressTask},
			{ReferenceTaskName: "task3", Status: FailedTask},
			{ReferenceTaskName: "task4", Status: ScheduledTask},
			{ReferenceTaskName: "task5", Status: FailedWithTerminalErrorTask},
		},
	}

	assert.Len(t, w.GetCompletedTasks(), 1)
	assert.Len(t, w.GetInProgressTasks(), 1)
	assert.Len(t, w.GetFailedTasks(), 2) // Both FailedTask and FailedWithTerminalErrorTask
	assert.Len(t, w.GetScheduledTasks(), 1)
}

func TestEmptyWorkflow(t *testing.T) {
	w := &Workflow{Tasks: []Task{}}
	assert.Len(t, w.GetInProgressTasks(), 0)
	assert.Len(t, w.GetCompletedTasks(), 0)
	assert.Len(t, w.GetFailedTasks(), 0)
	assert.Len(t, w.GetScheduledTasks(), 0)
	assert.Nil(t, w.GetTaskByReferenceName("any"))
}
