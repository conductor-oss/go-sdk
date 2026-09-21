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
	// PositionInput checks what goes in; kept for wire parity, as current servers
	// ignore it: agent guardrails compile at output only, and a tool guardrail
	// always inspects arguments before the call, never a result.
	PositionInput Position = "input"
	// PositionOutput checks what comes out. This is the default.
	PositionOutput Position = "output"
)

// OnFail is what happens when a guardrail rejects.
type OnFail string

const (
	// OnFailRaise fails the run.
	OnFailRaise OnFail = "raise"
	// OnFailRetry sends the message back to the model to retry, up to MaxRetries.
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

// validate checks the closed enums the shared schema fixes, so a typo fails here.
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

// RegexGuardrail matches content against patterns server side, with no worker.
type RegexGuardrail struct {
	guardrailBase
	// Patterns are regular expressions. Required.
	Patterns []string
	// Mode is "block" (default) to reject a match, or "allow" to reject a non-match.
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

// LLMGuardrail asks a model whether the content satisfies a policy, one call each.
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

// CustomGuardrail runs a Go function as a Conductor worker, dispatched by Name.
type CustomGuardrail struct {
	guardrailBase
	// Check reports whether the content passes. Required unless External is set.
	Check GuardrailFunc
	// External marks a guardrail whose worker runs elsewhere, registering nothing.
	External bool
}

// NewCustomGuardrail builds a custom guardrail; the other fields take defaults.
func NewCustomGuardrail(name string, check GuardrailFunc) *CustomGuardrail {
	return &CustomGuardrail{guardrailBase: guardrailBase{Name: name}, Check: check}
}

// GuardrailFunc is a custom check. passed=false rejects the content, and an error
// counts as a failure too, so a broken guardrail never lets content through. It
// takes a struct rather than the bare string Python and Java pass because the
// server also sends the retry iteration and the turn's tool calls.
type GuardrailFunc func(ctx context.Context, in GuardrailInput) (GuardrailResult, error)

// GuardrailInput is what the server hands a custom guardrail for one check. It
// has no Position: the server does not send one, and a guardrail knows its own.
type GuardrailInput struct {
	// Content is the text under inspection; non-text content arrives as JSON.
	Content string
	// Iteration is 0 on the first check, then one per retry this guardrail caused.
	Iteration int
	// ToolCalls are the turn's tool calls when the server includes them, else nil.
	ToolCalls []any
}

// GuardrailResult is a custom guardrail's verdict.
type GuardrailResult struct {
	// Passed reports whether the content is acceptable.
	Passed bool
	// Message is the retry feedback to the model, and the reason a run is stopped.
	Message string
	// FixedOutput is the corrected content for OnFailFix; empty fails the run.
	FixedOutput string
}

func (g *CustomGuardrail) guardrailConfig() map[string]any {
	cfg := g.baseConfig(g.Name)
	if g.External {
		cfg["guardrailType"] = "external"
	} else {
		cfg["guardrailType"] = "custom"
	}
	// The server dispatches to this name, and the runtime registers a worker under it.
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

// validateGuardrails checks every guardrail on an agent or tool. Guardrail stays
// a minimal interface so a caller can implement one, so the hook is optional.
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

// guardrailIn is the task input as the server sends it: the content, sent under
// several aliases of which "content" is canonical, and the live loop iteration.
type guardrailIn struct {
	Content   any   `json:"content"`
	Iteration any   `json:"iteration"`
	ToolCalls []any `json:"toolCalls"`
}

// guardrailOut is the task output, in the shape Python's GuardrailEntry returns.
// The server reads passed, message, on_fail and fixed_output; on_fail drives the
// retry/raise/fix switch and is always sent, as without it a failure raises.
type guardrailOut struct {
	Passed         bool    `json:"passed"`
	Message        string  `json:"message"`
	OnFail         string  `json:"on_fail"`
	FixedOutput    *string `json:"fixed_output"`
	GuardrailName  string  `json:"guardrail_name"`
	ShouldContinue bool    `json:"should_continue"`
}

// contentOf matches Python's _stringify_content: text as is, nil empty, else JSON.
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

// iterationOf reads the loop counter: float64 from JSON, string from old servers.
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

// guardrailHandler adapts Check to the worker contract as Python's GuardrailEntry
// does: a retry past MaxRetries, or a fix with nothing to substitute, raises.
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

// customGuardrails returns the guardrails this process must serve: the agent's
// own and its tools', minus server-evaluated regex and LLM ones and external ones.
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
