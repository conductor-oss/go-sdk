# Agent config golden files

These JSON documents are the cross-SDK wire contract for agent definitions.
Each one is the `agentConfig` that the **Python** SDK's `AgentConfigSerializer`
produces for a given agent, captured verbatim. The Go serializer must produce
the same document for the equivalent Go agent.

The server compiles `agentConfig` into a Conductor workflow, so a difference
here is a behavioural difference, not a formatting one.

## Running the test

    go test ./sdk/ai/ -run TestGoldenAgentConfig -v

`TestGoldenCoverage` prints how many fixtures are implemented. Fixtures whose
wire surface does not exist in Go yet are mapped to `nil` in `goldenFixtures`
and report as skipped rather than failing.

## Regenerating

Requires a checkout of [conductor-oss/python-sdk](https://github.com/conductor-oss/python-sdk).
No server or network access is needed; the serializer is pure.

    cd /path/to/python-sdk
    PYTHONPATH=src python3 /path/to/go-sdk/sdk/ai/testdata/agent_config/generate_fixtures.py \
        --out /path/to/go-sdk/sdk/ai/testdata/agent_config

Regenerate only when the wire format itself changes, and review the diff: a
change here means every SDK has to follow. If a golden file changes because of
a Python SDK upgrade, that is a signal to check the Java and TypeScript SDKs too.

## Notes

Comparison is **semantic**, not byte-for-byte. Both sides are parsed into
`map[string]any` before comparison, because key ordering legitimately differs
between languages. The files are written with sorted keys and two-space indent
purely to keep diffs readable.

The fixtures deliberately cover the parts of the format that are easy to get
subtly wrong:

- fields omitted when unset versus sent as a zero value
- `strategy` emitted only when the agent actually declares sub-agents
- derived worker task names such as `terminating_stop_when` and `router_fn_router_fn`
- nested `and` / `or` termination trees
- `credentials` landing under a tool's `config`, not at the tool's top level
