# Structured output

Constraining an agent's final answer to the shape of a Go struct with `Agent.OutputType`, and
reading it back off `AgentResult`.

## Declaring the type

Set `OutputType` to a zero value of the struct. The SDK derives a JSON Schema from it and the
server constrains the model to that schema.

```go
type Ticket struct {
    Summary  string
    Priority int
    Tags     []string
}

agent := &ai.Agent{
    Name:         "ticket_writer",
    Model:        "openai/gpt-4o",
    Instructions: "Turn the report into a ticket.",
    OutputType:   Ticket{},
}
```

On the wire this is `outputType: {"schema": ..., "className": "Ticket"}`; `className` is the Go
type's name. A pointer type is dereferenced first, so `OutputType: &Ticket{}` sends the same
document.

## How the schema is derived

Reflection over the struct, by the same rules that give a `tool.Func` handler its input schema:

- A field is advertised under its `json` tag, or under its name in **snake_case** when it has
  none — `Summary` is `summary`, `AccountID` is `account_id`, `TempF` is `temp_f`,
  `HTTPStatus` is `http_status`.
- Declaration order is preserved. `required` is a positional array, and the server copies the
  property order into text the model reads.
- A field is **required** unless it is a pointer or its tag carries `omitempty` — that pair is
  Go's stand-in for a Python default.
- `json:"-"` and unexported fields are skipped. An embedded struct's fields are lifted into the
  parent, as `encoding/json` does.

Go types map as: `string` → `string`; `bool` → `boolean`; every integer kind → `integer`;
`float32`/`float64` → `number`; slice or array → `array` with `items` (`[]byte` → `string`);
`map` → `object` with `additionalProperties`; struct → nested `object`; `any` → an open schema.

The document carries types and required-ness only — no descriptions, titles, formats or
constraints. Go and Java send types only; Python additionally carries Pydantic's per-field
titles, and the wire schema leaves the inner document open so all three are accepted.

So the type above yields:

```json
{
  "type": "object",
  "properties": {
    "summary":  {"type": "string"},
    "priority": {"type": "integer"},
    "tags":     {"type": "array", "items": {"type": "string"}}
  },
  "required": ["summary", "priority", "tags"]
}
```

Making a field optional:

```go
type Ticket struct {
    Summary  string
    Priority int
    Assignee *string           // optional: pointer
    Notes    string  `json:"notes,omitempty"` // optional: omitempty
    Internal string  `json:"-"`               // not sent
}
```

## Reading the result

```go
result, err := runtime.Run(ctx, agent, "The checkout page 500s on submit.")
if err != nil {
    return err
}

var ticket Ticket
if err := json.Unmarshal([]byte(result.Output), &ticket); err != nil {
    return err
}
fmt.Println(ticket.Summary, ticket.Priority)
```

`AgentResult`:

| Field | Meaning |
|---|---|
| `ExecutionID` | the execution this result came from |
| `Status` | `StatusRunning`, `StatusCompleted`, `StatusFailed`, `StatusTerminated`, `StatusTimedOut`, `StatusPaused` or `StatusWaiting`; `Status.Terminal()` reports whether it will progress no further |
| `Output` | the final answer, as text |
| `FinishReason` | why the loop stopped, e.g. `"STOP"` or `"MAX_TURNS"` |
| `Error` | the failure reason when `Status` is `StatusFailed` |
| `TokenUsage` | `PromptTokens`, `CompletionTokens`, `TotalTokens`; zero when the server does not report them |
| `Raw` | the untouched status document |

`Output` is a `string`: a structured answer arrives as the JSON document in that string, which is
what the example decodes. Fields are read from the status document leniently, so if a server
build returns the value as a JSON object rather than as text, `Output` is empty and the value is
still there under `output.result` in `Raw`.

`result.PrintResult()` writes the result to stdout in a boxed layout — the output, then status,
execution id and token usage, or the error when the run failed. It is the counterpart of the
Python SDK's `print_result` and Java's `printResult`.

Note that `Runtime.Run` returns the `*AgentResult` *and* a non-nil error when the run failed, so
a failed run is inspectable rather than opaque.

## Next steps

Continue with [tools](./tools.md), [guardrails](./guardrails.md), or the [runtime reference](../reference/runtime.md).
