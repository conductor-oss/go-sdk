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
	"strconv"
	"strings"
)

// Position is where a guardrail runs relative to the thing it guards.
type Position string

const (
	// PositionInput checks what goes in, before the model or tool sees it.
	// Kept for wire parity with the other SDKs; current servers do not act on
	// it. Agent-level guardrails are compiled at PositionOutput only, so an
	// input guardrail on an agent is accepted and ignored. Tool guardrails
	// always inspect the tool's arguments before the call, whichever position
	// they declare — there is no check on a tool's result.
	PositionInput Position = "input"
	// PositionOutput checks what comes out. This is the default.
	PositionOutput Position = "output"
)

// OnFail is what happens when a guardrail rejects.
type OnFail string

const (
	// OnFailRaise fails the run.
	OnFailRaise OnFail = "raise"
	// OnFailRetry sends the guardrail's message back to the model and lets it
	// try again, up to MaxRetries.
	OnFailRetry OnFail = "retry"
	// OnFailFix substitutes the corrected output the guardrail returns.
	OnFailFix OnFail = "fix"
	// OnFailHuman escalates to a person. Output position only.
	OnFailHuman OnFail = "human"
)

var validPositions = map[Position]struct{}{
	PositionInput: {}, PositionOutput: {},
}

var validOnFail = map[OnFail]struct{}{
	OnFailRaise: {}, OnFailRetry: {}, OnFailFix: {}, OnFailHuman: {},
}

// Defaults applied when a field is left at its zero value, matching Python.
const (
	defaultMaxRetries        = 3
	defaultRegexName         = "regex_guardrail"
	defaultLLMName           = "llm_guardrail"
	defaultRegexMode         = "block"
	defaultGuardrailPosition = PositionOutput
	defaultGuardrailOnFail   = OnFailRaise
)

// guardrailBase holds what every guardrail sends, whatever its type.
type guardrailBase struct {
	// Name identifies the guardrail. Each concrete type supplies a default.
	Name string
	// Position defaults to PositionOutput.
	Position Position
	// OnFail defaults to OnFailRaise.
	OnFail OnFail
	// MaxRetries defaults to 3 and applies to OnFailRetry.
	MaxRetries int
}

func (g guardrailBase) baseConfig(defaultName string) map[string]any {
	name := g.Name
	if name == "" {
		name = defaultName
	}
	pos := g.Position
	if pos == "" {
		pos = defaultGuardrailPosition
	}
	return map[string]any{
		"name":       name,
		"position":   pos,
		"onFail":     g.onFailOrDefault(),
		"maxRetries": g.maxRetriesOrDefault(),
	}
}

func (g guardrailBase) onFailOrDefault() OnFail {
	if g.OnFail == "" {
		return defaultGuardrailOnFail
	}
	return g.OnFail
}

func (g guardrailBase) maxRetriesOrDefault() int {
	if g.MaxRetries <= 0 {
		return defaultMaxRetries
	}
	return g.MaxRetries
}

// validate checks the closed enums. Positions and failure modes are a fixed
// set in the shared schema, so a typo is a bug worth catching here rather than
// a server-side compile error.
func (g guardrailBase) validate(kind string) error {
	if g.Position != "" {
		if _, ok := validPositions[g.Position]; !ok {
			return fmt.Errorf("invalid position %q on %s guardrail %q", g.Position, kind, g.Name)
		}
	}
	if g.OnFail != "" {
		if _, ok := validOnFail[g.OnFail]; !ok {
			return fmt.Errorf("invalid onFail %q on %s guardrail %q", g.OnFail, kind, g.Name)
		}
	}
	if g.OnFail == OnFailHuman && g.Position == PositionInput {
		return fmt.Errorf("guardrail %q: onFail=human is only valid at position=output", g.Name)
	}
	return nil
}

// RegexGuardrail matches content against patterns, server side. No worker is
// involved, so it costs nothing per call.
type RegexGuardrail struct {
	guardrailBase
	// Patterns are regular expressions. Required.
	Patterns []string
	// Mode is "block" (default) to reject a match, or "allow" to reject
	// anything that does not match.
	Mode string
	// Message is the feedback the model sees on a retry.
	Message string
}

func (g *RegexGuardrail) guardrailConfig() map[string]any {
	cfg := g.baseConfig(defaultRegexName)
	cfg["guardrailType"] = "regex"
	cfg["patterns"] = g.Patterns
	mode := g.Mode
	if mode == "" {
		mode = defaultRegexMode
	}
	cfg["mode"] = mode
	if g.Message != "" {
		cfg["message"] = g.Message
	}
	return cfg
}

func (g *RegexGuardrail) validateGuardrail() error {
	if len(g.Patterns) == 0 {
		return fmt.Errorf("regex guardrail %q has no patterns", g.Name)
	}
	if g.Mode != "" && g.Mode != "block" && g.Mode != "allow" {
		return fmt.Errorf("invalid mode %q: must be block or allow", g.Mode)
	}
	return g.validate("regex")
}

// LLMGuardrail asks a model whether the content satisfies a policy. It costs a
// model call per check, so prefer a small fast model.
type LLMGuardrail struct {
	guardrailBase
	// Model in "provider/model" form. Required.
	Model string
	// Policy describes what to check for. Required.
	Policy string
	// MaxTokens bounds the judging call. Omitted when zero.
	MaxTokens int
}

func (g *LLMGuardrail) guardrailConfig() map[string]any {
	cfg := g.baseConfig(defaultLLMName)
	cfg["guardrailType"] = "llm"
	cfg["model"] = g.Model
	cfg["policy"] = g.Policy
	if g.MaxTokens > 0 {
		cfg["maxTokens"] = g.MaxTokens
	}
	return cfg
}

func (g *LLMGuardrail) validateGuardrail() error {
	if g.Model == "" {
		return fmt.Errorf("llm guardrail %q has no model", g.Name)
	}
	if g.Policy == "" {
		return fmt.Errorf("llm guardrail %q has no policy", g.Name)
	}
	return g.validate("llm")
}

// CustomGuardrail runs a Go function as a Conductor worker. Name is the task
// name the server dispatches to, so it is required rather than defaulted: Go
// cannot read a function's name at runtime the way Python can.
type CustomGuardrail struct {
	guardrailBase
	// Check receives the content and reports whether it passes. Required
	// unless External is set.
	Check GuardrailFunc
	// External marks a guardrail whose worker runs elsewhere, so this process
	// registers nothing for it.
	External bool
}

// NewCustomGuardrail builds a custom guardrail with the two fields every one
// needs. Position, OnFail and MaxRetries take their defaults and can be set on
// the result.
func NewCustomGuardrail(name string, check GuardrailFunc) *CustomGuardrail {
	return &CustomGuardrail{guardrailBase: guardrailBase{Name: name}, Check: check}
}

// GuardrailFunc is a custom check. Returning passed=false rejects the content,
// and an error counts as a failure too, so a broken guardrail never lets
// content through.
//
// It takes a struct rather than the bare string Python and Java pass because
// the server sends more than the content — the retry iteration and the turn's
// tool calls — and a struct lets a check use them, and lets fields be added
// later without changing every caller.
type GuardrailFunc func(ctx context.Context, in GuardrailInput) (GuardrailResult, error)

// GuardrailInput is what the server hands a custom guardrail for one check.
// Only fields the server actually sends are here; Position is not one of
// them, and a guardrail knows its own Position from its declaration anyway.
type GuardrailInput struct {
	// Content is the text under inspection. Non-text content arrives as JSON,
	// the way Python's _stringify_content renders it.
	Content string
	// Iteration is which pass this is: 0 on the first check, then one more
	// for each retry the guardrail itself caused.
	Iteration int
	// ToolCalls are the turn's tool calls, when the server includes them.
	// Nil otherwise.
	ToolCalls []any
}

// GuardrailResult is a custom guardrail's verdict.
type GuardrailResult struct {
	// Passed reports whether the content is acceptable.
	Passed bool
	// Message is the feedback sent back to the model on a retry, and the
	// failure reason when the run is stopped.
	Message string
	// FixedOutput is the corrected content, used when OnFail is OnFailFix.
	// Empty means there is no fix, and a fix guardrail then fails the run.
	FixedOutput string
}

func (g *CustomGuardrail) guardrailConfig() map[string]any {
	cfg := g.baseConfig(g.Name)
	if g.External {
		cfg["guardrailType"] = "external"
	} else {
		cfg["guardrailType"] = "custom"
	}
	// The task name is the guardrail's name: the server dispatches to it, and
	// the runtime registers a worker under the same name.
	cfg["taskName"] = cfg["name"]
	return cfg
}

func (g *CustomGuardrail) validateGuardrail() error {
	if g.Name == "" {
		return fmt.Errorf("a custom guardrail needs a Name: it is the task name the " +
			"server dispatches to, and Go cannot derive it from the function")
	}
	if g.Check == nil && !g.External {
		return fmt.Errorf("custom guardrail %q has no Check function; set External "+
			"if its worker runs elsewhere", g.Name)
	}
	return g.validate("custom")
}

// validateGuardrails checks every guardrail on an agent or tool. Guardrail
// itself stays a minimal interface so a caller can implement one; the
// validation hook is optional.
func validateGuardrails(owner string, gs []Guardrail) error {
	for _, g := range gs {
		v, ok := g.(interface{ validateGuardrail() error })
		if !ok {
			continue
		}
		if err := v.validateGuardrail(); err != nil {
			return fmt.Errorf("%s: %w", owner, err)
		}
	}
	return nil
}

// guardrailIn is the task input as the server sends it: the content under
// several aliases (content is the canonical one) and the live loop iteration.
// Both are bound loosely because the content is whatever the model produced
// and the iteration may arrive as a number, a numeric string, or nothing.
type guardrailIn struct {
	Content   any   `json:"content"`
	Iteration any   `json:"iteration"`
	ToolCalls []any `json:"toolCalls"`
}

// guardrailOut is the task output, in the shape Python's GuardrailEntry
// returns. The server's normalize step reads passed, message, on_fail and
// fixed_output; on_fail drives the retry/raise/fix switch, so it is always
// sent — without it a failure always raises.
type guardrailOut struct {
	Passed         bool    `json:"passed"`
	Message        string  `json:"message"`
	OnFail         string  `json:"on_fail"`
	FixedOutput    *string `json:"fixed_output"`
	GuardrailName  string  `json:"guardrail_name"`
	ShouldContinue bool    `json:"should_continue"`
}

// contentOf renders the content the way Python's _stringify_content does:
// text as is, nothing as empty, anything else as JSON.
func contentOf(v any) string {
	switch x := v.(type) {
	case nil:
		return ""
	case string:
		return x
	default:
		b, err := json.Marshal(x)
		if err != nil {
			return fmt.Sprint(x)
		}
		return string(b)
	}
}

// iterationOf reads the loop counter, which JSON delivers as a float64 and
// older servers as a string. Anything else counts as iteration zero.
func iterationOf(v any) int {
	switch x := v.(type) {
	case float64:
		return int(x)
	case int:
		return x
	case string:
		if i, err := strconv.Atoi(strings.TrimSpace(x)); err == nil {
			return i
		}
	}
	return 0
}

// guardrailHandler adapts Check to the worker contract, mirroring Python's
// GuardrailEntry: a retry past MaxRetries and a fix with nothing to substitute
// both escalate to raise, and a Check that errors is a failure rather than a
// pass, so a broken guardrail never lets content through.
func (g *CustomGuardrail) guardrailHandler() func(context.Context, guardrailIn) (guardrailOut, error) {
	onFail := g.onFailOrDefault()
	maxRetries := g.maxRetriesOrDefault()
	return func(ctx context.Context, in guardrailIn) (guardrailOut, error) {
		res, err := g.Check(ctx, GuardrailInput{
			Content:   contentOf(in.Content),
			Iteration: iterationOf(in.Iteration),
			ToolCalls: in.ToolCalls,
		})
		if err != nil {
			res = GuardrailResult{Message: "Guardrail error: " + err.Error()}
		}
		if res.Passed {
			return guardrailOut{Passed: true, OnFail: "pass"}, nil
		}
		var fixed *string
		if res.FixedOutput != "" {
			fixed = &res.FixedOutput
		}
		fail := onFail
		if fail == OnFailRetry && iterationOf(in.Iteration) >= maxRetries {
			fail = OnFailRaise
		}
		if fail == OnFailFix && fixed == nil {
			fail = OnFailRaise
		}
		return guardrailOut{
			Message:        res.Message,
			OnFail:         string(fail),
			FixedOutput:    fixed,
			GuardrailName:  g.Name,
			ShouldContinue: fail == OnFailRetry,
		}, nil
	}
}

// customGuardrails returns the guardrails this process must serve: the
// agent's own and its tools', skipping regex and LLM guardrails, which the
// server evaluates itself, and external ones served elsewhere.
func (a *Agent) customGuardrails() []*CustomGuardrail {
	var out []*CustomGuardrail
	collect := func(gs []Guardrail) {
		for _, g := range gs {
			if c, ok := g.(*CustomGuardrail); ok && c.Check != nil && !c.External {
				out = append(out, c)
			}
		}
	}
	collect(a.Guardrails)
	for _, t := range a.Tools {
		collect(t.Guardrails)
	}
	return out
}
