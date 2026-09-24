//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

import "fmt"

// defaultContextMaxBytes is the server's cap on a fetched context document (Python's Context.max_bytes).
const defaultContextMaxBytes = 16384

// PlanContext is material handed to the Planner of a StrategyPlanExecute agent
// with the prompt: exactly one of literal Text or a URL the server fetches
// (Python's plans.Context).
type PlanContext struct {
	Text string
	URL  string
	// Headers are sent when fetching URL.
	Headers map[string]string
	// Optional lets planning proceed when URL cannot be fetched (Python's required=False).
	Optional bool
	// MaxBytes caps a fetched document. Zero means the server default.
	MaxBytes int
}

func (c PlanContext) validate() error {
	if (c.Text == "") == (c.URL == "") {
		return fmt.Errorf("PlanContext: exactly one of Text or URL must be set")
	}
	return nil
}

// config mirrors Context.to_dict: text alone, or url with only non-default settings.
func (c PlanContext) config() map[string]any {
	out := map[string]any{}
	if c.Text != "" {
		out["text"] = c.Text
	}
	if c.URL != "" {
		out["url"] = c.URL
		if len(c.Headers) > 0 {
			out["headers"] = c.Headers
		}
		if c.Optional {
			out["required"] = false
		}
		if c.MaxBytes != 0 && c.MaxBytes != defaultContextMaxBytes {
			out["maxBytes"] = c.MaxBytes
		}
	}
	return out
}

// PrefillToolCall runs a tool with fixed arguments before the first LLM turn and
// puts its result into the conversation. The tool need not be in Agent.Tools: the
// runtime registers its worker either way, so it never lacks a poller.
type PrefillToolCall struct {
	Tool      ToolDef
	Arguments map[string]any
}

// Prefill builds a PrefillToolCall, the counterpart of Python's my_tool.call(arg=value).
func Prefill(t ToolDef, arguments map[string]any) PrefillToolCall {
	return PrefillToolCall{Tool: t, Arguments: arguments}
}

func (p PrefillToolCall) config() map[string]any {
	args := p.Arguments
	if args == nil {
		args = map[string]any{}
	}
	return map[string]any{"toolName": p.Tool.Name, "arguments": args}
}
