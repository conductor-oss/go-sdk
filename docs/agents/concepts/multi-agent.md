# Multi-agent

How one `ai.Agent` orchestrates others: strategies, routing, handoffs, gates, and plan-execute.

There is one agent type. An agent becomes a multi-agent system when it declares sub-agents
in `Agents`, or fills the plan-execute slots `Planner` or `Fallback`. `Strategy` is sent to
the server only then.

```go
parent := &ai.Agent{
    Name:     "support_team",
    Model:    "openai/gpt-4o",
    Agents:   []*ai.Agent{billing, technical},
    Strategy: ai.StrategySequential,
}
```

## Strategies

| Constant | Wire value | Behaviour |
|---|---|---|
| `ai.StrategyHandoff` | `handoff` | Default. Each sub-agent may transfer control to another. |
| `ai.StrategySequential` | `sequential` | Sub-agents run in declaration order. |
| `ai.StrategyParallel` | `parallel` | All sub-agents run at once. |
| `ai.StrategyRouter` | `router` | One sub-agent, chosen by `Router` or `RouterFunc`. |
| `ai.StrategyRoundRobin` | `round_robin` | Sub-agents cycle in order. |
| `ai.StrategyRandom` | `random` | A sub-agent is picked at random. |
| `ai.StrategySwarm` | `swarm` | Sub-agents hand off freely, guided by `Handoffs`. |
| `ai.StrategyManual` | `manual` | Sub-agent selection is left to a worker. |
| `ai.StrategyPlanExecute` | `plan_execute` | `Planner` emits a plan; its steps execute. |

An unknown strategy is rejected by `Agent.Validate`, which runtime methods call before
serializing.

Under `StrategyParallel`, `Synthesize` controls the final LLM step that merges the
sub-agents' results. Nil keeps the server's default of true; `ai.Ptr(false)` skips it.

## Nesting

`Agents` holds `*ai.Agent`, so a sub-agent may declare sub-agents of its own, each with its
own strategy. A sub-agent with an empty `Model` inherits the parent's at compile time.

```go
engineering := &ai.Agent{
    Name:     "engineering_lead",
    Agents:   []*ai.Agent{backendDev, frontendDev},
    Strategy: ai.StrategyHandoff,
}
ceo := &ai.Agent{
    Name:     "ceo",
    Model:    "openai/gpt-4o",
    Agents:   []*ai.Agent{engineering, marketing},
    Strategy: ai.StrategySwarm,
}
```

`Validate` recurses into `Agents`, `Router`, `Planner` and `Fallback`. The runtime starts
workers for the whole tree: sub-agents, agents used as tools, router, planner and fallback.

An agent can also be a plain tool instead of a participant in the strategy, with
`tool.Agent(name, agent, description)`. The parent gets its result; control never
transfers. See [tools](./tools.md).

## Router vs RouterFunc

`StrategyRouter` requires exactly one of:

| Field | Type | Selection runs |
|---|---|---|
| `Router` | `*ai.Agent` | On the server, as a nested LLM agent whose output names a sub-agent |
| `RouterFunc` | `ai.RouterFunc` | In your process, as a worker named `<agent>_router_fn` |

```go
// LLM router
parent := &ai.Agent{
    Name:     "dispatcher",
    Model:    "openai/gpt-4o",
    Agents:   []*ai.Agent{billing, technical},
    Strategy: ai.StrategyRouter,
    Router: &ai.Agent{
        Name:         "classifier",
        Instructions: "Reply with exactly one word: billing or technical.",
    },
}

// Go router
parent.Router = nil
parent.RouterFunc = func(ctx context.Context, prompt string) (string, error) {
    if strings.Contains(prompt, "invoice") {
        return "billing", nil
    }
    return "technical", nil
}
```

Setting both is a validation error. Either one is serialized whatever the strategy, but only
`StrategyRouter` acts on it.

## Handoffs

`Handoffs` is a list of `ai.HandoffCondition`, usually with `StrategySwarm`. The interface is
sealed: the server knows a fixed set of types, so only these three implement it.

| Type | Fields | Fires when |
|---|---|---|
| `*ai.OnTextMention` | `Target`, `Text` | The active agent's text contains `Text` |
| `*ai.OnToolResult` | `Target`, `ToolName`, `ResultContains` | `ToolName` returns; `ResultContains` narrows it, empty means any result |
| `*ai.OnCondition` | `Target`, `Condition` | An `ai.HandoffFunc` in your process returns true |

```go
team := &ai.Agent{
    Name:     "support_team",
    Model:    "openai/gpt-4o",
    Agents:   []*ai.Agent{support, billing, escalation},
    Strategy: ai.StrategySwarm,
    Handoffs: []ai.HandoffCondition{
        &ai.OnTextMention{Target: "billing", Text: "refund"},
        &ai.OnToolResult{Target: "escalation", ToolName: "check_severity", ResultContains: "P1"},
        &ai.OnCondition{Target: "billing", Condition: func(ctx context.Context, s ai.HandoffState) (bool, error) {
            return strings.Contains(s.Result, "invoice"), nil
        }},
    },
    AllowedTransitions: map[string][]string{
        "support": {"billing", "escalation"},
    },
}
```

`HandoffFunc` is `func(ctx context.Context, state ai.HandoffState) (bool, error)`.
`HandoffState` carries `Result`, `Conversation`, `Context`, `ActiveAgent` and `ToolResults`.
An error is treated as "do not hand off".

Each `OnCondition` runs as a worker named `<agent>_handoff_<target>`, so a parent may hold
only one `OnCondition` per target; a second is rejected by `Validate` rather than colliding
on the task name. Every `Target`, and every name in `AllowedTransitions`, must be one of the
parent's sub-agents.

`AllowedTransitions` maps a sub-agent name to the targets it may reach. Empty means no
restriction.

## Gates

A gate decides, after an agent inside a `StrategySequential` pipeline finishes, whether the
pipeline continues. Set `Gate` on the agent producing the output.

| Type | Evaluated | Stops when |
|---|---|---|
| `ai.TextGate` | Server side | The output contains `Text`; case-sensitive unless `IgnoreCase` |
| `ai.GateFunc` | A worker named `<agent>_gate` | Your function returns false |

```go
reviewer := &ai.Agent{
    Name: "reviewer",
    Gate: ai.TextGate{Text: "ESCALATE", IgnoreCase: true},
}

checker := &ai.Agent{
    Name: "checker",
    Gate: ai.GateFunc(func(ctx context.Context, s ai.GateState) (bool, error) {
        return !strings.Contains(s.Result, "STOP"), nil
    }),
}
```

`GateFunc` returns true to continue. `GateState` carries only `Result`. An error continues
the pipeline.

## Plan-execute

Under `StrategyPlanExecute` the `Planner` agent writes a plan over the parent's `Tools`, and
`Fallback` runs if the plan cannot be carried out. `FallbackMaxTurns` bounds the fallback;
zero leaves the limit to the server. `PlannerContext` is extra material for the planner and
is valid only under this strategy.

```go
harness := &ai.Agent{
    Name:     "research_pac",
    Model:    "openai/gpt-4o",
    Strategy: ai.StrategyPlanExecute,
    Tools:    ai.Tools(tool.Func("get_weather", getWeather, "Return today's weather for a city.")),
    Planner: &ai.Agent{
        Name:         "research_pac_planner",
        Instructions: "Plan the steps using the available tools.",
    },
    Fallback: &ai.Agent{
        Name:         "research_pac_fallback",
        Instructions: "The plan failed. Recover with the available tools.",
    },
    FallbackMaxTurns: 4,
}
```

### EnablePlanning is not plan-execute

The names invite confusion, so they are worth separating.

`EnablePlanning` asks one ordinary agent to think before it acts:

```go
agent := &ai.Agent{
    Name:           "researcher",
    Model:          "openai/gpt-4o",
    Tools:          ai.Tools(searchTool, summarizeTool),
    EnablePlanning: true,
}
```

The server adds a sentence to that agent's instructions: plan the steps first, work through them
with your tools, and check progress as you go. Nothing else changes. Same agent, same tools, same
tasks in the UI, and no `Planner` or sub-agents involved.

| | `EnablePlanning` | `StrategyPlanExecute` |
|---|---|---|
| What it is | extra wording in the prompt | an orchestration strategy |
| The plan | stays in the model's head | a real `ai.Plan` the server compiles and runs |
| Needs sub-agents | no | yes, a `Planner`, and usually a `Fallback` |
| Costs | nothing | a planning turn, plus the steps |

Reach for `EnablePlanning` when one agent keeps skipping steps on a multi-step job. Reach for
plan-execute when you need the plan to be inspectable, resumable and reviewable before it runs.

### Supplying the plan yourself

`ai.WithPlan(plan)` passes a ready-made `ai.Plan` to `Run` or `Start`, and the planner never
runs. The server still requires the `Planner` slot, because the strategy compiles around it.
`WithPlan` on any other strategy is an error.

```go
plan := &ai.Plan{Steps: []ai.Step{
    {ID: "weather", Operations: []ai.Op{
        {Tool: "get_weather", Args: map[string]any{"city": "Oslo"}},
    }},
    {ID: "packing", DependsOn: []string{"weather"}, Operations: []ai.Op{
        {Tool: "packing_advice", Args: map[string]any{"weather": ai.Ref{StepID: "weather"}}},
    }},
}}

res, err := rt.Run(ctx, harness, "What should I pack for Oslo?", ai.WithPlan(plan))
```

| Type | Fields | Notes |
|---|---|---|
| `ai.Plan` | `Steps`, `Validation`, `OnSuccess`, `OnFailure` | At least one step; an empty plan reads as "no plan" and the planner runs |
| `ai.Step` | `ID`, `Operations`, `DependsOn`, `Parallel` | Starts once every `DependsOn` step has completed; `Parallel` runs its `Operations` at once |
| `ai.Op` | `Tool`, `Args`, `Generate` | Exactly one of `Args` or `Generate` |
| `ai.Ref` | `StepID` | Stands for the whole output of an earlier step, anywhere in `Args` |
| `ai.Generate` | `Instructions`, `OutputSchema`, `MaxTokens`, `Context` | An LLM writes the arguments at run time |

`Ref` has no field selection: the receiving tool takes a parameter of the producing tool's
output type. To use part of an output, return just that field from the producing tool, or
use `Generate`. `Generate.OutputSchema` is an example JSON object keyed by the tool's
argument names, with real placeholder values, because the server parses it to learn the
keys. `Generate.Context` may itself be a `Ref`, so the LLM sees the earlier step's actual
output.

`Plan.Validate` rejects a missing or duplicate step ID, a `DependsOn` on an unknown step, a
`Ref` to an unknown step, and an `Op` that sets both or neither of `Args` and `Generate`.
`Run` and `Start` call it before starting.

`Validation` entries run after the steps, each a tool call with an optional
`SuccessCondition` the server evaluates over its output. `OnSuccess` and `OnFailure` are
`ai.Action` tool calls run afterwards.

## Bounding the loop

Every open-ended composition needs a limit. `MaxTurns` defaults to 25. `Termination` is a
server-evaluated condition tree and `StopWhen` is Go logic run as a worker; see
[termination](./termination.md).

## Next steps

Continue with [termination](./termination.md), [deploy, serve, run](./deploy-serve-run.md), or the [agent definition reference](../reference/agent-definition.md).
