//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

import "context"

// Go functions are never sent to the server. Each one is registered as a
// Conductor worker and referenced in agentConfig by a derived task name, so
// the name is the contract between the serializer and the worker registry:
// the runtime must register a worker under exactly the name emitted here.
//
// The suffixes match the Python SDK, which builds the same names in its
// serializer and its worker registration.
const (
	stopWhenSuffix = "stop_when"
	routerSuffix   = "router_fn"
	gateSuffix     = "gate"
)

// workerTaskName derives the task name for one of the agent's callables.
func (a *Agent) workerTaskName(suffix string) string {
	return a.Name + "_" + suffix
}

// workerRef is the wire shape for a callable: {"taskName": "..."}.
func workerRef(taskName string) map[string]any {
	return map[string]any{"taskName": taskName}
}

// StopWhenState is the state handed to a StopWhenFunc after each turn.
//
// It mirrors the context the Python worker builds — result, messages and
// iteration — so a predicate ported between SDKs sees the same inputs.
type StopWhenState struct {
	// Result is the agent's output so far.
	Result string
	// Messages is the conversation history at this point.
	Messages []map[string]any
	// Iteration counts loop passes, starting at 0.
	Iteration int
}

// RouterFunc picks which sub-agent handles a prompt, by name.
//
// The returned name must be one of the agent's sub-agents. The runtime
// registers it as a worker under "<agent>_router_fn"; if it returns an unknown
// name or fails, the server falls back to the first sub-agent, matching the
// Python worker's behaviour.
//
// Use Router instead when the choice is better made by an LLM: that path is a
// nested agent the server runs, with no worker round trip.
type RouterFunc func(ctx context.Context, prompt string) (string, error)

// StopWhenFunc decides whether to stop the agent loop after a turn.
//
// Returning true stops the loop. The runtime registers it as a worker under
// "<agent>_stop_when" and inverts the result into the should_continue flag
// the server expects, matching the Python worker's contract.
//
// Use this for logic the server cannot express; use TerminationCondition for
// the cases it can, since those are evaluated server-side without a worker
// round trip.
type StopWhenFunc func(ctx context.Context, state StopWhenState) (bool, error)
