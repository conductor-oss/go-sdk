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
	"encoding/json"
	"fmt"
	"strings"
)

// Workers for the agent's own callables and conditions. The server compiles
// each into a task named "<agent>_<suffix>" and waits for a worker to answer
// it; a task nothing polls stays SCHEDULED and the run never finishes, so each
// must be registered whenever the agent sets the field that produces it.

const terminationSuffix = "termination"

// loopStateIn is what the server sends these tasks. Result is untyped because a
// turn may produce content blocks, rendered to text as Python's str() does.
type loopStateIn struct {
	Result    any              `json:"result"`
	Messages  []map[string]any `json:"messages,omitempty"`
	Iteration int              `json:"iteration"`
}

func (in loopStateIn) state() StopWhenState {
	return StopWhenState{Result: resultText(in.Result), Messages: in.Messages, Iteration: in.Iteration}
}

func resultText(v any) string {
	switch value := v.(type) {
	case nil:
		return ""
	case string:
		return value
	}
	raw, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprint(v)
	}
	return string(raw)
}

type shouldContinueOut struct {
	ShouldContinue bool   `json:"should_continue"`
	Reason         string `json:"reason,omitempty"`
}

type routerIn struct {
	Prompt string `json:"prompt"`
}

type routerOut struct {
	SelectedAgent string `json:"selected_agent"`
}

// terminationHandler tells the server whether the condition lets the loop go on.
func terminationHandler(cond TerminationCondition) func(context.Context, loopStateIn) (shouldContinueOut, error) {
	return func(_ context.Context, in loopStateIn) (shouldContinueOut, error) {
		stop, reason := evaluateTermination(cond, in.state())
		return shouldContinueOut{ShouldContinue: !stop, Reason: reason}, nil
	}
}

// stopWhenHandler inverts the predicate's "stop" into the server's "continue".
// An error continues the loop, as the Python worker does.
func (f StopWhenFunc) stopWhenHandler() func(context.Context, loopStateIn) (shouldContinueOut, error) {
	return func(ctx context.Context, in loopStateIn) (shouldContinueOut, error) {
		stop, err := f(ctx, in.state())
		if err != nil {
			return shouldContinueOut{ShouldContinue: true}, nil
		}
		return shouldContinueOut{ShouldContinue: !stop}, nil
	}
}

// routerHandler falls back to the first sub-agent when the function errors.
func (f RouterFunc) routerHandler(fallback string) func(context.Context, routerIn) (routerOut, error) {
	return func(ctx context.Context, in routerIn) (routerOut, error) {
		name, err := f(ctx, in.Prompt)
		if err != nil || name == "" {
			return routerOut{SelectedAgent: fallback}, nil
		}
		return routerOut{SelectedAgent: name}, nil
	}
}

// evaluateTermination produces the Python conditions' verdicts and reasons.
// TerminationCondition is sealed, so an unknown one only continues the loop.
func evaluateTermination(cond TerminationCondition, state StopWhenState) (stop bool, reason string) {
	switch c := cond.(type) {
	case TextMentionTermination:
		return evaluateTextMention(c, state)
	case StopMessageTermination:
		return evaluateStopMessage(c, state)
	case MaxMessageTermination:
		return evaluateMaxMessage(c, state)
	case TokenUsageTermination:
		// Token counts are absent from this task's input, so it never fires here.
		return false, ""
	case andTermination:
		var reasons []string
		for _, inner := range c.conditions {
			stop, why := evaluateTermination(inner, state)
			if !stop {
				return false, ""
			}
			if why != "" {
				reasons = append(reasons, why)
			}
		}
		return true, strings.Join(reasons, " AND ")
	case orTermination:
		for _, inner := range c.conditions {
			if stop, why := evaluateTermination(inner, state); stop {
				return true, why
			}
		}
	}
	return false, ""
}

func evaluateTextMention(c TextMentionTermination, state StopWhenState) (bool, string) {
	result, text := state.Result, c.Text
	if !c.CaseSensitive {
		result, text = strings.ToLower(result), strings.ToLower(text)
	}
	if strings.Contains(result, text) {
		return true, fmt.Sprintf("Text '%s' found in output", c.Text)
	}
	return false, ""
}

func evaluateStopMessage(c StopMessageTermination, state StopWhenState) (bool, string) {
	message := c.StopMessage
	if message == "" {
		message = defaultStopMessage
	}
	if strings.TrimSpace(state.Result) == message {
		return true, fmt.Sprintf("Stop message '%s' received", message)
	}
	return false, ""
}

// Counts the history the server sent, or the iteration when it sent none.
func evaluateMaxMessage(c MaxMessageTermination, state StopWhenState) (bool, string) {
	count := len(state.Messages)
	if count == 0 {
		count = state.Iteration
	}
	if count >= c.MaxMessages {
		return true, fmt.Sprintf("Message count (%d) >= limit (%d)", count, c.MaxMessages)
	}
	return false, ""
}

// systemWorkers returns one worker per field the server compiles into a task.
func (a *Agent) systemWorkers() []ToolDef {
	var out []ToolDef
	if a.Termination != nil {
		out = append(out, ToolDef{
			Name:    a.workerTaskName(terminationSuffix),
			Handler: terminationHandler(a.Termination),
		})
	}
	if a.StopWhen != nil {
		out = append(out, ToolDef{
			Name:    a.workerTaskName(stopWhenSuffix),
			Handler: a.StopWhen.stopWhenHandler(),
		})
	}
	if a.RouterFunc != nil {
		var fallback string
		if len(a.Agents) > 0 {
			fallback = a.Agents[0].Name
		}
		out = append(out, ToolDef{
			Name:    a.workerTaskName(routerSuffix),
			Handler: a.RouterFunc.routerHandler(fallback),
		})
	}
	return out
}
