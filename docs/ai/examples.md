# Worked Examples

Three examples of the shipped API. Each maps to an e2e test in `test/ai_e2e/` that runs it
against a live server and asserts the worker-side effect.

## 1. A tool and a guardrail — the hello world

```go
type weatherIn  struct{ City string `json:"city"` }
type weatherOut struct {
    City      string `json:"city"`
    TempF     int    `json:"temp_f"`
    Condition string `json:"condition"`
}

func getWeather(ctx context.Context, in weatherIn) (weatherOut, error) {
    return weatherOut{City: in.City, TempF: 72, Condition: "sunny"}, nil
}

noKeys := ai.NewCustomGuardrail("no_api_keys", func(_ context.Context, in ai.GuardrailInput) (ai.GuardrailResult, error) {
    if strings.Contains(in.Content, "sk-live-") {
        return ai.GuardrailResult{Passed: false, Message: "Remove the API key before answering."}, nil
    }
    return ai.GuardrailResult{Passed: true}, nil
})
noKeys.OnFail = ai.OnFailRetry

agent := &ai.Agent{
    Name:         "weather",
    Model:        "openai/gpt-4o-mini",
    Instructions: "Answer in one short sentence.",
    Tools:        []ai.ToolDef{tool.Func("get_weather", "Current temperature for a city", getWeather)},
    Guardrails:   []ai.Guardrail{noKeys},
}

rt := ai.NewRuntime(ai.Config{})
defer rt.Shutdown()
res, err := rt.Run(ctx, agent, "What's the weather in Oslo?")
```

The schemas for `get_weather` come from `weatherIn`/`weatherOut` by reflection — including the
**output** schema, which a plan step downstream can name fields from. The guardrail runs as a
worker named `no_api_keys`; the server dispatches to it after each reply.

Tests: `TestToolCall` (the tool's `72` reaches the answer), `TestCustomGuardrailPasses` /
`TestCustomGuardrailRaises`.

## 2. A swarm with a predicate handoff and a credentialed tool

```go
func openPR(ctx context.Context, in prIn) (prOut, error) {
    token, err := ai.Secret(ctx, "GH_TOKEN")     // delivered with the task, never from env
    if err != nil {
        return prOut{Error: err.Error()}, nil       // tell the model rather than fail the task
    }
    ...
}

support := &ai.Agent{
    Name:     "support",
    Model:    model,
    Strategy: ai.StrategySwarm,
    Agents: []*ai.Agent{
        {Name: "triage",  Model: model, Instructions: "Triage the request in one sentence."},
        {Name: "billing", Model: model, Instructions: "Handle billing.",
            Tools: []ai.ToolDef{tool.Func("open_pr", "Open a refund PR", openPR, tool.WithCredentials("GH_TOKEN"))}},
    },
    Handoffs: []ai.HandoffCondition{
        &ai.OnCondition{Target: "billing", Condition: func(_ context.Context, s ai.HandoffState) (bool, error) {
            return s.ActiveAgent == "triage" && strings.Contains(strings.ToLower(s.Result), "refund"), nil
        }},
    },
    AllowedTransitions: map[string][]string{"triage": {"billing"}},
}
```

The predicate runs as the worker `support_handoff_billing` — the name the server dispatches to.
`s.ActiveAgent` is already resolved from the server's index to `"triage"`. `GH_TOKEN` is resolved
server-side and arrives on the task; the tool never sees the environment.

Tests: `TestOnConditionHandoff`, `TestTeamWithSecret`, `TestSecretRequiresDeclaration`.

## 3. Streaming with human approval

```go
ops := &ai.Agent{
    Name: "ops", Model: model, Temperature: ai.Ptr(0.0),
    Tools: []ai.ToolDef{
        tool.Func("check_service",       "Health of a service",  checkService),
        tool.Func("restart_service",     "Restart a service",    restartService),
        tool.Func("delete_service_data", "Delete data. Destructive.", deleteServiceData, tool.RequiresApproval()),
    },
}

h, _ := rt.Start(ctx, ops, "Payments is down: check it, restart it, then clear its stale cache.")
events, _ := h.Events(ctx)
go func() { for ev := range events { log.Println(ev.Name) } }()

// Control flow polls; the stream is evidence, not the trigger.
for !terminal {
    if waiting, _ := h.Waiting(ctx); waiting {
        _ = h.Approve(ctx)
    }
    ...
}
res, _ := h.Result(ctx)
```

`RequiresApproval()` makes the server pause before dispatching that one tool and emit a `waiting`
event. `Approve` waits for the server to confirm the waiting state first, because the event lands
slightly before the server will accept a response.

Test: `TestStreamingWithApproval` — asserts `approve` precedes `tool:delete_service_data` in the
recorded sequence and that the guarded tool ran exactly once.

## Also verified

`TestPlanExecute` (a planner chains an `int` from one tool into another — the case that exposed
the server's schema handling), `TestCodeExecution`, `TestCLICommand`,
`TestMCPToolResultReachesTheAnswer`, `TestSecretsEnvForSubprocess`.
