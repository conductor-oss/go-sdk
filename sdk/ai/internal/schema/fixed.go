//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package schema

// Schemas the server defines rather than a Go type: these tool types have no
// handler, since the server dispatches them, so their schemas are fixed and
// the other SDKs send exactly these documents. They live here because the tool
// constructors and the wire-conformance tests both need them and cannot import
// each other. Each call returns a fresh map, since callers may add to a schema.

// HumanInput is the default schema for a human tool: one question to present.
func HumanInput() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"question": map[string]any{
				"type":        "string",
				"description": "The question or prompt to present to the human.",
			},
		},
		"required": []string{"question"},
	}
}

// AgentRequest is the schema for an agent-as-tool: the request to delegate.
func AgentRequest() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"request": map[string]any{
				"type":        "string",
				"description": "The request or question to send to this agent.",
			},
		},
		"required": []string{"request"},
	}
}

// EmptyObject is the schema for a tool that takes no model-supplied arguments.
func EmptyObject() map[string]any {
	return map[string]any{"type": "object", "properties": map[string]any{}}
}
