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

// A tool's own guardrails run in the worker, around the handler, as the
// Python worker's run_tool_task does. The server gates a tool's input before
// it dispatches the call, but a guardrail on the tool's output has no other
// place to run: the result goes from the worker straight back to the model.
//
// Input guardrails see the call's arguments as JSON; output guardrails see the
// handler's return, as is for a string and as JSON otherwise. A failed input
// guardrail with onFail raise fails the task, and any other failure returns a
// blocked marker in place of running the tool. A failed output guardrail with
// onFail fix substitutes the fixed output, raise fails the task, and any other
// failure returns a blocked marker in place of the result.

// evaluateLocally runs one guardrail in this process. Regex guardrails match
// here with the block/allow semantics the server uses; custom guardrails call
// their Check function. An LLM guardrail is not evaluated and evaluated is
// false: the Python worker calls the model itself for one, and this SDK makes
// no model calls of its own, so on a tool an LLM guardrail is enforced only
// where the server enforces it.
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
		// The same verdicts and default messages as the Python RegexGuardrail._check.
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

// guardrailName, guardrailPosition and guardrailOnFail read the effective
// values a guardrail serializes, defaults applied.
func guardrailName(g Guardrail) string { return fmt.Sprint(g.guardrailConfig()["name"]) }

func guardrailPosition(g Guardrail) Position {
	return Position(fmt.Sprint(g.guardrailConfig()["position"]))
}

func guardrailOnFail(g Guardrail) OnFail { return OnFail(fmt.Sprint(g.guardrailConfig()["onFail"])) }

// blockedResult is what the model sees in place of a tool result a guardrail
// refused, the shape the Python worker returns.
func blockedResult(message string) map[string]any {
	return map[string]any{"error": message, "blocked": true}
}

// checkToolInput runs the tool's input guardrails over the call's arguments.
// It returns a replacement result when a guardrail blocks the call, or an
// error when the guardrail's onFail is raise.
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

// checkToolOutput runs the tool's output guardrails over the handler's return
// and gives back what the task should carry: the value itself, the fixed
// output, or a blocked marker. onFail raise is an error.
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

// guardrailContent is the text a guardrail sees for a value: a string as is,
// anything else as JSON, as the Python worker does.
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
