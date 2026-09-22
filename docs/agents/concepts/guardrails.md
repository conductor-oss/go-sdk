# Guardrails

Guardrails reject, retry or repair what an agent or a tool produces. This page covers the three
guardrail types, where they attach, and what current servers actually enforce.

`Agent.Guardrails` and `ToolDef.Guardrails` are both `[]ai.Guardrail`. `Guardrail` is a sealed
interface; the implementations are `*RegexGuardrail`, `*LLMGuardrail` and `*CustomGuardrail`, all
used as pointers.

| Type | Evaluated by | Needs a worker |
|---|---|---|
| `RegexGuardrail` | the server, from `Patterns` | no |
| `LLMGuardrail` | the server, one judging call per check | no |
| `CustomGuardrail` | your `Check` function | yes, in your process |

## Shared fields

Every guardrail embeds the same four fields. They are promoted from an unexported struct, so they
cannot be set in a composite literal — assign them after construction:

```go
g := &ai.RegexGuardrail{Patterns: []string{`\b\d{3}-\d{2}-\d{4}\b`}}
g.Name = "no_ssn"
g.Position = ai.PositionOutput
g.OnFail = ai.OnFailRetry
g.MaxRetries = 2
```

| Field | Default | Meaning |
|---|---|---|
| `Name` | `"regex_guardrail"` / `"llm_guardrail"`; required on `CustomGuardrail` | Identifies the guardrail, and is the task name for a custom one. |
| `Position` | `PositionOutput` | `PositionInput` or `PositionOutput`. See [What the server enforces](#what-the-server-enforces). |
| `OnFail` | `OnFailRaise` | What happens on rejection. |
| `MaxRetries` | `3` | Retry budget for `OnFailRetry`. |

`OnFail` values:

| Value | Effect |
|---|---|
| `ai.OnFailRaise` | Fail the run. |
| `ai.OnFailRetry` | Send the message back to the model, up to `MaxRetries`; past that, raise. |
| `ai.OnFailFix` | Substitute `GuardrailResult.FixedOutput`; raise when it is empty. |
| `ai.OnFailHuman` | Escalate to a person. Rejected by `Agent.Validate` at `PositionInput`. |

## RegexGuardrail

```go
noEmail := &ai.RegexGuardrail{
    Patterns: []string{`[\w.+-]+@[\w-]+\.[\w.-]+`},
    Mode:     "block",
    Message:  "Response must not contain email addresses. Redact them.",
}
noEmail.Name = "no_email"
noEmail.OnFail = ai.OnFailRetry
```

`Patterns` is required. `Mode` is `"block"` (default, reject a match) or `"allow"` (reject a
non-match). `Message` is the feedback the model sees on a retry.

## LLMGuardrail

```go
tone := &ai.LLMGuardrail{
    Model:     "openai/gpt-4o",
    Policy:    "The response must be professional and free of slang.",
    MaxTokens: 256,
}
tone.Name = "professional_tone"
```

`Model` and `Policy` are required. `MaxTokens` bounds the judging call and is omitted when zero.

## CustomGuardrail

`NewCustomGuardrail(name, check)` returns a `*CustomGuardrail` with defaults applied. The check is
a `GuardrailFunc`:

```go
type GuardrailFunc func(ctx context.Context, in ai.GuardrailInput) (ai.GuardrailResult, error)
```

```go
func noPII(ctx context.Context, in ai.GuardrailInput) (ai.GuardrailResult, error) {
    if ccPattern.MatchString(in.Content) {
        return ai.GuardrailResult{
            Passed:  false,
            Message: "Redact all card numbers before responding.",
        }, nil
    }
    return ai.GuardrailResult{Passed: true}, nil
}

guard := ai.NewCustomGuardrail("no_pii", noPII)
guard.OnFail = ai.OnFailRetry

agent := &ai.Agent{
    Name:       "support_agent",
    Model:      "openai/gpt-4o",
    Tools:      ai.Tools(tool.Func("get_customer", getCustomer, "Look up a customer.")),
    Guardrails: []ai.Guardrail{guard},
}
```

`GuardrailInput`:

| Field | Type | Meaning |
|---|---|---|
| `Content` | `string` | The text under inspection; non-text content arrives as JSON. |
| `Iteration` | `int` | 0 on the first check, then one per retry this guardrail caused. |
| `ToolCalls` | `[]any` | The turn's tool calls when the server sends them, else nil. |

`GuardrailResult`:

| Field | Type | Meaning |
|---|---|---|
| `Passed` | `bool` | Whether the content is acceptable. |
| `Message` | `string` | Retry feedback to the model, and the reason a run is stopped. |
| `FixedOutput` | `string` | Replacement content for `OnFailFix`. |

A returned error is a rejection, not a pass: a broken guardrail never lets content through.

### It runs as a worker in your process

The runtime registers one Conductor worker per custom guardrail — the agent's own and its tools' —
under the guardrail's `Name`, which is also the `taskName` sent to the server. Names must therefore
be unique across the agent. Set `External: true` for a guardrail whose worker runs elsewhere:
nothing is registered locally, and `Check` may be nil.

## Tool guardrails

```go
lookup := tool.Func("get_customer", getCustomer, "Look up a customer.")
lookup.Guardrails = []ai.Guardrail{guard}
```

For a worker tool, the SDK also evaluates the tool's guardrails around the handler in this process:
`PositionInput` ones against the call arguments as JSON before the handler runs, `PositionOutput`
ones against its return value. `OnFailRaise` fails the task; `OnFailFix` substitutes `FixedOutput`;
anything else returns `{"error": ..., "blocked": true}` to the model. `LLMGuardrail` is skipped
there — this SDK makes no model calls — so only the server enforces it.

## What the server enforces

Two facts the field names do not show:

- Current servers compile **agent-level** guardrails at output position only. An agent guardrail
  declaring `PositionInput` is accepted, serialized and then ignored.
- A **tool** guardrail always inspects the tool's arguments before the call, whatever `Position` it
  declares. The server never gates a tool result.

`Position` is kept for wire parity with the Python and Java SDKs.

## Validation

`Agent.Validate` rejects an unknown `Position` or `OnFail`, `OnFailHuman` at `PositionInput`, a
`RegexGuardrail` with no `Patterns` or a `Mode` other than `block`/`allow`, an `LLMGuardrail`
missing `Model` or `Policy`, and a `CustomGuardrail` with no `Name` or with neither `Check` nor
`External`.

## Next steps

Continue with [termination](./termination.md), [callbacks](./callbacks.md), or [streaming and human-in-the-loop](./streaming-hitl.md) for `OnFailHuman`.
