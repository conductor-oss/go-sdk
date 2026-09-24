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
	"regexp"
)

// A tool's own guardrails run in the worker, around the handler, as Python's
// run_tool_task does: the server gates a tool's input before dispatch, but an
// output guardrail has nowhere else to run — the result goes from the worker
// straight back to the model.

// evaluateLocally runs one guardrail in this process: regex matching with the
// server's block/allow semantics, or a custom Check. An LLM guardrail reports
// evaluated false — this SDK makes no model calls — so only the server enforces it.
func evaluateLocally(ctx context.Context, g Guardrail, content string) (result GuardrailResult, evaluated bool, err error) {
	switch g := g.(type) {
	case *RegexGuardrail:
		mode := g.Mode
		if mode == "" {
			mode = defaultRegexMode
		}
		var matched bool
		for _, p := range g.Patterns {
			re, err := regexp.Compile(p)
			if err != nil {
				return GuardrailResult{}, true, fmt.Errorf("guardrail %q: invalid pattern %q: %w", guardrailName(g), p, err)
			}
			if re.MatchString(content) {
				matched = true
				break
			}
		}
		// The verdicts and default messages of Python's RegexGuardrail._check.
		var failure string
		switch {
		case mode == "block" && matched:
			failure = "Content matched a blocked pattern."
		case mode == "allow" && !matched:
			failure = "Content did not match any allowed pattern."
		default:
			return GuardrailResult{Passed: true}, true, nil
		}
		if g.Message != "" {
			failure = g.Message
		}
		return GuardrailResult{Passed: false, Message: failure}, true, nil
	case *CustomGuardrail:
		if g.Check == nil {
			// An external guardrail's worker runs elsewhere.
			return GuardrailResult{}, false, nil
		}
		res, err := g.Check(ctx, GuardrailInput{Content: content})
		return res, true, err
	}
	return GuardrailResult{}, false, nil
}

// guardrailName, guardrailPosition and guardrailOnFail read serialized values,
// defaults applied.
func guardrailName(g Guardrail) string { return fmt.Sprint(g.guardrailConfig()["name"]) }

func guardrailPosition(g Guardrail) Position {
	return Position(fmt.Sprint(g.guardrailConfig()["position"]))
}

func guardrailOnFail(g Guardrail) OnFail { return OnFail(fmt.Sprint(g.guardrailConfig()["onFail"])) }

// blockedResult is the shape the Python worker returns for a refused call.
func blockedResult(message string) map[string]any {
	return map[string]any{"error": message, "blocked": true}
}

// checkToolInput runs the input guardrails over the call's arguments as JSON,
// returning a replacement result when one blocks, or an error for onFail raise.
func checkToolInput(ctx context.Context, t ToolDef, input map[string]any) (map[string]any, error) {
	if len(t.Guardrails) == 0 {
		return nil, nil
	}
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("tool %q: encode input for guardrails: %w", t.Name, err)
	}
	for _, g := range t.Guardrails {
		if guardrailPosition(g) != PositionInput {
			continue
		}
		res, evaluated, err := evaluateLocally(ctx, g, string(raw))
		if err != nil {
			res = GuardrailResult{Message: "Guardrail error: " + err.Error()}
		}
		if !evaluated || res.Passed {
			continue
		}
		if guardrailOnFail(g) == OnFailRaise {
			return nil, fmt.Errorf("tool guardrail %q blocked execution: %s", guardrailName(g), res.Message)
		}
		return blockedResult(fmt.Sprintf("Blocked by guardrail '%s': %s", guardrailName(g), res.Message)), nil
	}
	return nil, nil
}

// checkToolOutput runs the output guardrails over the handler's return and
// gives back the value, the fixed output, or a blocked marker; raise is an error.
func checkToolOutput(ctx context.Context, t ToolDef, value any) (any, error) {
	if len(t.Guardrails) == 0 {
		return value, nil
	}
	for _, g := range t.Guardrails {
		if guardrailPosition(g) != PositionOutput {
			continue
		}
		content, err := guardrailContent(value)
		if err != nil {
			return nil, fmt.Errorf("tool %q: encode output for guardrails: %w", t.Name, err)
		}
		res, evaluated, err := evaluateLocally(ctx, g, content)
		if err != nil {
			res = GuardrailResult{Message: "Guardrail error: " + err.Error()}
		}
		if !evaluated || res.Passed {
			continue
		}
		switch onFail := guardrailOnFail(g); {
		case onFail == OnFailFix && res.FixedOutput != "":
			value = res.FixedOutput
		case onFail == OnFailRaise:
			return nil, fmt.Errorf("tool guardrail %q failed: %s", guardrailName(g), res.Message)
		default:
			value = blockedResult(fmt.Sprintf("Output blocked by guardrail '%s': %s", guardrailName(g), res.Message))
		}
	}
	return value, nil
}

// guardrailContent is the text a guardrail sees: a string as is, else JSON.
func guardrailContent(value any) (string, error) {
	if s, ok := value.(string); ok {
		return s, nil
	}
	raw, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}
