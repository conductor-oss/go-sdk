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

// Per-execution worker domains, the mechanism behind Agent.Stateful: a run id
// doubles as a queue name, the server records the mapping on the workflow and
// routes that run's tasks there, and this process polls that queue, so two
// runs never take each other's tasks. Without a run id the server schedules
// into the shared queue and workers poll it undomained.

// newRunID returns the run id a stateful start request carries, empty when
// nothing in the tree is stateful. Format: the Python SDK's uuid4().hex — 32
// lowercase hex characters, no dashes.
func newRunID(agent *Agent) string {
	if agent == nil || !hasStatefulTools(agent) {
		return ""
	}
	var buf [16]byte
	if _, err := rand.Read(buf[:]); err != nil {
		// Dropping the domain would silently lose the caller's isolation.
		return fmt.Sprintf("%032x", buf)
	}
	return hex.EncodeToString(buf[:])
}

// hasStatefulTools reports whether anything in the tree asks for a domain.
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
// its routing is known; a run without a run id started its workers already.
func (r *Runtime) startStatefulWorkers(ctx context.Context, agent *Agent, executionID, runID string) error {
	if runID == "" {
		return nil
	}
	return r.registerWorkers(agent, r.workerDomains(ctx, executionID))
}

// workerDomains maps each task name to the queue the server routes it to. A
// name the server left out is scheduled on the shared queue, which is how a
// nested skill's tasks behave: the server routes only the agent's own tools.
func (r *Runtime) workerDomains(ctx context.Context, executionID string) map[string]string {
	wf, _, err := r.workflow.GetExecutionStatus(ctx, executionID,
		&client.WorkflowResourceApiGetExecutionStatusOpts{IncludeTasks: optional.NewBool(false)})
	if err != nil {
		return nil
	}
	return wf.TaskToDomain
}
