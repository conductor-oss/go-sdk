//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/antihax/optional"

	"github.com/conductor-sdk/conductor-go/sdk/client"
)

// Per-execution worker domains, the mechanism behind Agent.Stateful.
//
// A stateful run sends the server a run id. The server routes that run's tasks
// to a queue of the same name and records the mapping on the workflow, and
// this process starts workers that poll that queue. Calls in one run therefore
// reach the workers of the process that started it, and two runs never take
// each other's tasks.
//
// Without a run id nothing changes: the server schedules into the shared queue
// and workers poll it undomained, which is what every non-stateful run does.

// newRunID returns the run id a stateful agent's start request carries, or an
// empty string when nothing in the agent tree is stateful. The format matches
// the Python SDK's uuid4().hex: 32 lowercase hex characters, no dashes.
func newRunID(agent *Agent) string {
	if agent == nil || !hasStatefulTools(agent) {
		return ""
	}
	var buf [16]byte
	if _, err := rand.Read(buf[:]); err != nil {
		// Only a broken system randomness source reaches here. Losing the
		// domain would silently drop the isolation the caller asked for, so
		// fall back to a value that is still unique to this run.
		return fmt.Sprintf("%032x", buf)
	}
	return hex.EncodeToString(buf[:])
}

// hasStatefulTools reports whether the agent, any of its tools, or anything
// nested under it asks for a per-execution domain.
func hasStatefulTools(a *Agent) bool {
	if a == nil {
		return false
	}
	if a.Stateful {
		return true
	}
	for _, t := range a.Tools {
		if t.Stateful {
			return true
		}
	}
	for _, sub := range a.nestedAgents() {
		if hasStatefulTools(sub) {
			return true
		}
	}
	return false
}

// startStatefulWorkers registers this run's workers once the run exists and
// its routing is known. It does nothing for a run without a run id, whose
// workers were already started before the run began.
func (r *Runtime) startStatefulWorkers(ctx context.Context, agent *Agent, executionID, runID string) error {
	if runID == "" {
		return nil
	}
	return r.registerWorkers(agent, r.workerDomains(ctx, executionID))
}

// workerDomains is the run's task-to-domain map: for each task name, the queue
// the server routes it to. A name the server left out is scheduled on the
// shared queue, so its worker polls undomained — that is how a nested skill's
// own tasks behave, since the server routes only the tools it compiled for the
// agent itself.
func (r *Runtime) workerDomains(ctx context.Context, executionID string) map[string]string {
	wf, _, err := r.workflow.GetExecutionStatus(ctx, executionID,
		&client.WorkflowResourceApiGetExecutionStatusOpts{IncludeTasks: optional.NewBool(false)})
	if err != nil {
		return nil
	}
	return wf.TaskToDomain
}
