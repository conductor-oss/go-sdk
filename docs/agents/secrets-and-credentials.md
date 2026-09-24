# Secrets & Credentials

## The contract that must port exactly

There is one source of truth for credentials: the Conductor server. The SDK never reads an env
var, a `.env` file, or a keyring for them. python-sdk's `docs/security.md` states the invariant:

> Agent tools declare required credentials; a capable server delivers resolved values only in
> task runtime metadata. Missing credentials fail before tool execution.

The five steps, and where each lives in Go:

1. **Declare** — `tool.WithCredentials("GH_TOKEN")` on a tool. Names, never values. Go cannot
   inspect a function body to see which secrets it uses, so this declaration is what tells the
   server what to resolve. `ToolDef.Credentials` is a typed field that serializes under the tool's
   `config`, which is where the server's compiler looks.
2. **Register** — the names travel with the tool in `agentConfig`.
3. **Resolve** — server-side, against its store. Opaque to the SDK.
4. **Deliver** — on `Task.RuntimeMetadata` of the specific task handed to the poller. Wire-only:
   never persisted to task input, never in execution history, never a separate fetch.
5. **Consume** — `ai.Secret(ctx, "GH_TOKEN")`. A missing name is `ErrCredentialNotFound`,
   fail-closed, with the delivered names in the message. No fallback to the process environment.

`Secret` is a lookup, not a fetch: the runtime moves `RuntimeMetadata` into the task context before
calling the handler (`dispatch.go`), so `Secret` is only usable inside a tool handler and does no I/O.

```go
func openPR(ctx context.Context, in PRIn) (PRResult, error) {
    token, err := ai.Secret(ctx, "GH_TOKEN")
    if err != nil {
        return PRResult{Error: err.Error()}, nil   // tell the model; do not fail the task
    }
    ...
}
```

## What did not port: `os.environ` mutation

Python offers `inject_via_env`, which writes secrets into the process environment under a
process-wide lock so libraries that read env vars find them. Go does not do this. There is no
ambient-env-reading library ecosystem to accommodate, and mutating a shared process environment
from concurrent goroutines is exactly the hazard the Python lock papers over.

The one legitimate need for env vars — a tool that shells out — is served by `ai.SecretsEnv(ctx,
names...)`, which returns `KEY=VALUE` strings for `exec.Cmd.Env`, scoped to the child process and
never touching the parent's environment. `TestSecretsEnvForSubprocess` covers it.

## Redaction

`Agent.MaskedFields` names fields the server masks in stored execution data. Credentials never
enter task input or output in the first place, so there is nothing to redact there; the risk is a
tool echoing a secret in its result, which is the tool author's responsibility.

## Verified end to end

| test | what it proves |
|---|---|
| `TestTeamWithSecret` | a declared credential is delivered to a sub-agent's tool and used |
| `TestSecretRequiresDeclaration` | an undeclared name is `ErrCredentialNotFound`, not an empty string |
| `TestSecretsEnvForSubprocess` | `SecretsEnv` reaches a child process; the parent env is untouched |

## Summary: Python mechanism → Go

| Python | Go | Why |
|---|---|---|
| `os.environ` mutation + `RLock` | none | no ambient consumers; concurrency hazard |
| `contextvars` accessor (`get_secret`) | `ai.Secret(ctx, name)` — `context.Context` is the propagation | same ergonomics, Go's native mechanism |
| per-subprocess env | `ai.SecretsEnv` → `exec.Cmd.Env` | scoped to the child |
| `Task.runtime_metadata` delivery | `Task.RuntimeMetadata` (unchanged) | the contract worth preserving; server-side and language-agnostic |
| credential names under tool `config` | same | golden fixtures `02_tools_worker` and `17_tools_mcp` pin it; `15_execution_and_creds` pins agent-level `credentials` |
