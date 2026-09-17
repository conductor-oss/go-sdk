//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

// Strategy determines how an agent orchestrates its sub-agents.
//
// A strategy is only sent to the server when the agent actually declares
// sub-agents (via Agents, Planner or Fallback); a leaf agent has no strategy
// on the wire regardless of what this field holds.
type Strategy string

const (
	// StrategyHandoff lets each sub-agent transfer control to another. Default.
	StrategyHandoff Strategy = "handoff"
	// StrategySequential runs sub-agents in declaration order.
	StrategySequential Strategy = "sequential"
	// StrategyParallel fans out to every sub-agent at once.
	StrategyParallel Strategy = "parallel"
	// StrategyRouter dispatches to one sub-agent chosen by Router.
	StrategyRouter Strategy = "router"
	// StrategyRoundRobin cycles through sub-agents in order.
	StrategyRoundRobin Strategy = "round_robin"
	// StrategyRandom picks a sub-agent at random.
	StrategyRandom Strategy = "random"
	// StrategySwarm lets sub-agents hand off freely, guided by Handoffs.
	StrategySwarm Strategy = "swarm"
	// StrategyManual defers sub-agent selection to a worker.
	StrategyManual Strategy = "manual"
	// StrategyPlanExecute plans with Planner, then executes the plan's steps.
	StrategyPlanExecute Strategy = "plan_execute"
)

// validStrategies mirrors the Python SDK's Strategy enum. Kept as a set so
// Validate can reject unknown values with the same message shape.
var validStrategies = map[Strategy]struct{}{
	StrategyHandoff:     {},
	StrategySequential:  {},
	StrategyParallel:    {},
	StrategyRouter:      {},
	StrategyRoundRobin:  {},
	StrategyRandom:      {},
	StrategySwarm:       {},
	StrategyManual:      {},
	StrategyPlanExecute: {},
}
