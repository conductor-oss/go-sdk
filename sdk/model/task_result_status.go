//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

import "fmt"

type TaskResultStatus string

const (
	InProgressTask              TaskResultStatus = "IN_PROGRESS"
	FailedTask                  TaskResultStatus = "FAILED"
	FailedWithTerminalErrorTask TaskResultStatus = "FAILED_WITH_TERMINAL_ERROR"
	CompletedTask               TaskResultStatus = "COMPLETED"
	ScheduledTask               TaskResultStatus = "SCHEDULED"
	SkippedTask                 TaskResultStatus = "SKIPPED"
)

func (t TaskResultStatus) String() string {
	return string(t)
}

func ParseTaskResultStatus(status string) (TaskResultStatus, error) {
	switch status {
	case string(InProgressTask):
		return InProgressTask, nil
	case string(FailedTask):
		return FailedTask, nil
	case string(FailedWithTerminalErrorTask):
		return FailedWithTerminalErrorTask, nil
	case string(CompletedTask):
		return CompletedTask, nil
	case string(ScheduledTask):
		return ScheduledTask, nil
	case string(SkippedTask):
		return SkippedTask, nil
	default:
		return InProgressTask, fmt.Errorf("invalid task result status: %s", status)
	}
}
