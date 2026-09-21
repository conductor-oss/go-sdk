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

// Go functions are never sent to the server: each is registered as a Conductor
// worker and referenced in agentConfig by the derived task name
// "<agent>_<suffix>", so the runtime must register a worker under exactly the
// name emitted here. The suffixes match the Python SDK.
const (
	stopWhenSuffix = "stop_when"
	routerSuffix   = "router_fn"
	gateSuffix     = "gate"
)

func (a *Agent) workerTaskName(suffix string) string {
	return a.Name + "_" + suffix
}

// workerRef is the wire shape for a callable: {"taskName": "..."}.
func workerRef(taskName string) map[string]any {
	return map[string]any{"taskName": taskName}
}

// StopWhenState is the state handed to a StopWhenFunc after each turn. It
// mirrors the Python worker's context, so a ported predicate sees the same inputs.
type StopWhenState struct {
	// Result is the agent's output so far.
	Result string
	// Messages is the conversation history at this point.
	Messages []map[string]any
	// Iteration counts loop passes, starting at 0.
	Iteration int
}

// RouterFunc picks which sub-agent handles a prompt, by name. The name must be
// one of the agent's sub-agents; an unknown name or an error falls back to the
// first sub-agent, matching the Python worker. Use Router when an LLM should
// make the choice: the server runs that nested agent with no worker round trip.
type RouterFunc func(ctx context.Context, prompt string) (string, error)

// StopWhenFunc decides whether to stop the agent loop after a turn; true stops
// it. The runtime inverts the result into the should_continue flag the server
// expects. Use TerminationCondition for checks the server can evaluate itself,
// since those need no worker round trip.
type StopWhenFunc func(ctx context.Context, state StopWhenState) (bool, error)
