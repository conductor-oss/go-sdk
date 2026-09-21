# CLAUDE.md

Go module `github.com/conductor-sdk/conductor-go`, targeting Go 1.23.

## No explanatory comments on code

Do not write comments that say what the code already says. Exceptions, one or
two sentences each:

- Godoc on exported API, starting with the identifier's name.
- A fact the code cannot show: server behaviour, a wire-format or ordering
  requirement, a rejected alternative.
- A Python or Java SDK parity decision that changes behaviour.

No examples in doc comments except on a package's central type. For calibration,
`sdk/client` is near 9% comment lines and `sdk/workflow` near 16%.

## Commands

```bash
go build ./...
go test ./sdk/... ./test/unit_tests/...   # no server needed
gofmt -l .                                 # must print nothing
golangci-lint run ./...                    # CI pins v2.3, reports only new findings
```

Lint has pre-existing findings in `sdk/client` and `sdk/model`. CI compares
against the pull request's merge base, so leave them alone in an unrelated
change.

## Tests

| Suite | Needs |
|---|---|
| `test/unit_tests`, `sdk/**/*_test.go` | nothing |
| `test/integration_tests` | a server in `CONDUCTOR_SERVER_URL` |
| `test/integration_tests/ai` | `-tags integration` and an LLM provider; see its README |

Golden files under `sdk/ai/testdata/agent_config` pin the agent wire format
against the Python SDK. Changing serialization without updating one is a bug.

## Conventions

- Match the surrounding file; don't reformat or reorder as a side effect.
- Public API is used outside this repo. Adding is cheap, changing is not.
- Commits: one short imperative line, no AI attribution.
