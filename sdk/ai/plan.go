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
	"fmt"
	"reflect"
)

// Plan is a ready-made plan for a StrategyPlanExecute agent.
//
// Passing one to Run or Start via WithPlan skips the planner LLM: the server
// compiles these steps verbatim and carries them out with the agent's tools.
// That makes a run deterministic — the same steps, in the same order, with the
// same arguments — which is what the planner cannot promise.
//
// The serialized form is the wire format the server's plan compiler consumes
// and is shared with the Java (org.conductoross.conductor.ai.plans) and Python
// (conductor.ai.agents.plans) SDKs; a plan written in one reads in the others.
type Plan struct {
	// Steps is the DAG of operations. At least one is required: the server
	// treats an empty plan as "no plan" and falls back to the planner, which
	// would silently undo the point of passing one.
	Steps []Step
	// Validation lists checks to run after the steps complete.
	Validation []Validation
	// OnSuccess runs when validation passes; OnFailure when it does not.
	OnSuccess []Action
	OnFailure []Action
}

// Step is a node in the plan. Steps run in declaration order unless DependsOn
// says otherwise; a step starts once every step it depends on has completed.
type Step struct {
	// ID names the step. Other steps refer to its output with Ref{ID}.
	ID string
	// Operations are the tool calls this step makes, in order — or at once
	// when Parallel is set.
	Operations []Op
	// DependsOn lists the step IDs this step waits for.
	DependsOn []string
	// Parallel runs Operations concurrently instead of one after another.
	Parallel bool
}

// Op is one tool call inside a step.
//
// Exactly one of Args or Generate must be set. Args calls the tool with
// literal values, resolved deterministically at compile time; Generate asks an
// LLM to write the arguments at run time from Instructions.
type Op struct {
	// Tool is the name of one of the agent's tools.
	Tool string
	// Args are the tool's arguments. A Ref value anywhere in the tree — at
	// the top level or nested in a map or slice — is replaced with the
	// referenced step's output when the plan runs.
	Args map[string]any
	// Generate defers argument construction to an LLM call.
	Generate *Generate
}

// Ref stands in for the whole output of an earlier step.
//
// It carries the complete result object, not a field of it, so the receiving
// tool declares a parameter of the producing tool's output type. There is no
// field selection ("weather.temp_f"); to pass one field, have the producing
// tool return it on its own or let Generate pick it out.
type Ref struct {
	StepID string
}

// Generate has an LLM write a tool's arguments at run time.
type Generate struct {
	// Instructions say what the arguments should contain.
	Instructions string
	// OutputSchema is an example of the JSON object the LLM must produce,
	// with the tool's argument names as keys. Give real placeholder values
	// — {"temp_f": 0}, not {"temp_f": <integer>} — because the server parses
	// it as JSON to learn which keys to wire into the tool.
	OutputSchema string
	// MaxTokens caps the LLM's reply. Zero leaves the server default.
	MaxTokens int
	// Context is extra material for the LLM: a string, or a Ref so it sees
	// an earlier step's actual output rather than the reference itself.
	Context any
}

// Validation is a post-execution check the plan runs against its own results.
type Validation struct {
	Tool string
	Args map[string]any
	// SuccessCondition is a server-evaluated expression over the tool's
	// output; empty means the tool completing is the check.
	SuccessCondition string
}

// Action is a tool call run after validation, from OnSuccess or OnFailure.
type Action struct {
	Tool string
	Args map[string]any
}

// Validate reports the first structural problem with the plan.
//
// Everything checked here would otherwise surface as a compile failure on the
// server, or worse as a step that runs with nothing bound — a Ref to a step
// that does not exist resolves to nothing, and the tool still runs.
func (p *Plan) Validate() error {
	if p == nil {
		return fmt.Errorf("plan is nil")
	}
	if len(p.Steps) == 0 {
		return fmt.Errorf("plan has no steps")
	}
	ids, err := p.stepIDs()
	if err != nil {
		return err
	}
	for _, s := range p.Steps {
		if err := s.validate(ids); err != nil {
			return err
		}
	}
	return p.validateHooks(ids)
}

// stepIDs collects the step IDs, rejecting a missing or repeated one.
func (p *Plan) stepIDs() (map[string]struct{}, error) {
	ids := make(map[string]struct{}, len(p.Steps))
	for i, s := range p.Steps {
		if s.ID == "" {
			return nil, fmt.Errorf("plan step %d has no id", i)
		}
		if _, dup := ids[s.ID]; dup {
			return nil, fmt.Errorf("plan step id %q is used twice", s.ID)
		}
		ids[s.ID] = struct{}{}
	}
	return ids, nil
}

func (s Step) validate(ids map[string]struct{}) error {
	for _, dep := range s.DependsOn {
		if _, ok := ids[dep]; !ok {
			return fmt.Errorf("plan step %q depends on unknown step %q", s.ID, dep)
		}
	}
	for i, op := range s.Operations {
		if err := op.validate(ids); err != nil {
			return fmt.Errorf("plan step %q op %d: %w", s.ID, i, err)
		}
	}
	return nil
}

// validateHooks covers the tool calls that run after the steps: validation,
// on_success and on_failure.
func (p *Plan) validateHooks(ids map[string]struct{}) error {
	for i, v := range p.Validation {
		if v.Tool == "" {
			return fmt.Errorf("plan validation %d has no tool", i)
		}
		if err := checkRefs(v.Args, ids); err != nil {
			return fmt.Errorf("plan validation %q: %w", v.Tool, err)
		}
	}
	if err := validateActions("on_success", p.OnSuccess, ids); err != nil {
		return err
	}
	return validateActions("on_failure", p.OnFailure, ids)
}

func validateActions(group string, actions []Action, ids map[string]struct{}) error {
	for i, a := range actions {
		if a.Tool == "" {
			return fmt.Errorf("plan %s action %d has no tool", group, i)
		}
		if err := checkRefs(a.Args, ids); err != nil {
			return fmt.Errorf("plan %s action %q: %w", group, a.Tool, err)
		}
	}
	return nil
}

func (o Op) validate(ids map[string]struct{}) error {
	if o.Tool == "" {
		return fmt.Errorf("op has no tool")
	}
	if (o.Args == nil) == (o.Generate == nil) {
		return fmt.Errorf("op %q: exactly one of Args or Generate must be set", o.Tool)
	}
	if o.Generate != nil {
		if o.Generate.Instructions == "" || o.Generate.OutputSchema == "" {
			return fmt.Errorf("op %q: Generate needs Instructions and OutputSchema", o.Tool)
		}
		return checkRefs(o.Generate.Context, ids)
	}
	return checkRefs(o.Args, ids)
}

// checkRefs walks a value tree and rejects any Ref to a step not in ids.
func checkRefs(v any, ids map[string]struct{}) error {
	switch x := v.(type) {
	case Ref:
		return x.check(ids)
	case *Ref:
		if x == nil {
			return nil
		}
		return x.check(ids)
	case map[string]any:
		for _, sub := range x {
			if err := checkRefs(sub, ids); err != nil {
				return err
			}
		}
		return nil
	default:
		for _, sub := range anySlice(v) {
			if err := checkRefs(sub, ids); err != nil {
				return err
			}
		}
		return nil
	}
}

func (r Ref) check(ids map[string]struct{}) error {
	if r.StepID == "" {
		return fmt.Errorf("Ref has no step id")
	}
	if _, ok := ids[r.StepID]; !ok {
		return fmt.Errorf("Ref to unknown step %q", r.StepID)
	}
	return nil
}

// anySlice returns v's elements when v is a slice of interface values — []any
// or a named type with that shape — and nil otherwise. Slices of concrete
// types (e.g. []string) cannot hold a Ref, so there is nothing to walk.
func anySlice(v any) []any {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Slice || rv.Type().Elem().Kind() != reflect.Interface {
		return nil
	}
	out := make([]any, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		out = append(out, rv.Index(i).Interface())
	}
	return out
}

// toPayload renders the plan in the shape the server's plan compiler reads.
//
// Optional lists are omitted rather than sent empty, and booleans only when
// true, matching Plan.toJson() in Java and Plan.to_dict() in Python so the
// three SDKs produce the same document for the same plan.
func (p *Plan) toPayload() map[string]any {
	steps := make([]any, 0, len(p.Steps))
	for _, s := range p.Steps {
		steps = append(steps, s.toPayload())
	}
	out := map[string]any{"steps": steps}
	if len(p.Validation) > 0 {
		vs := make([]any, 0, len(p.Validation))
		for _, v := range p.Validation {
			vs = append(vs, v.toPayload())
		}
		out["validation"] = vs
	}
	if len(p.OnSuccess) > 0 {
		out["on_success"] = actionsPayload(p.OnSuccess)
	}
	if len(p.OnFailure) > 0 {
		out["on_failure"] = actionsPayload(p.OnFailure)
	}
	return out
}

func (s Step) toPayload() map[string]any {
	ops := make([]any, 0, len(s.Operations))
	for _, op := range s.Operations {
		ops = append(ops, op.toPayload())
	}
	out := map[string]any{"id": s.ID, "operations": ops}
	if len(s.DependsOn) > 0 {
		deps := make([]any, 0, len(s.DependsOn))
		for _, d := range s.DependsOn {
			deps = append(deps, d)
		}
		out["depends_on"] = deps
	}
	if s.Parallel {
		out["parallel"] = true
	}
	return out
}

func (o Op) toPayload() map[string]any {
	out := map[string]any{"tool": o.Tool}
	if o.Args != nil {
		out["args"] = serializePlanValue(o.Args)
	}
	if o.Generate != nil {
		out["generate"] = o.Generate.toPayload()
	}
	return out
}

func (g *Generate) toPayload() map[string]any {
	out := map[string]any{
		"instructions":  g.Instructions,
		"output_schema": g.OutputSchema,
	}
	if g.MaxTokens > 0 {
		out["max_tokens"] = g.MaxTokens
	}
	if g.Context != nil {
		out["context"] = serializePlanValue(g.Context)
	}
	return out
}

func (v Validation) toPayload() map[string]any {
	out := map[string]any{"tool": v.Tool}
	if v.Args != nil {
		out["args"] = serializePlanValue(v.Args)
	}
	if v.SuccessCondition != "" {
		out["success_condition"] = v.SuccessCondition
	}
	return out
}

func actionsPayload(actions []Action) []any {
	out := make([]any, 0, len(actions))
	for _, a := range actions {
		m := map[string]any{"tool": a.Tool}
		if a.Args != nil {
			m["args"] = serializePlanValue(a.Args)
		}
		out = append(out, m)
	}
	return out
}

// toPayload is the wire form of a Ref: {"$ref": "<step id>"}.
func (r Ref) toPayload() map[string]any {
	return map[string]any{"$ref": r.StepID}
}

// serializePlanValue copies a value tree, replacing every Ref with its wire
// form. Maps and slices are walked; anything else passes through unchanged.
func serializePlanValue(v any) any {
	switch x := v.(type) {
	case Ref:
		return x.toPayload()
	case *Ref:
		if x == nil {
			return nil
		}
		return x.toPayload()
	case map[string]any:
		out := make(map[string]any, len(x))
		for k, sub := range x {
			out[k] = serializePlanValue(sub)
		}
		return out
	default:
		items := anySlice(v)
		if items == nil {
			return v
		}
		out := make([]any, 0, len(items))
		for _, sub := range items {
			out = append(out, serializePlanValue(sub))
		}
		return out
	}
}
