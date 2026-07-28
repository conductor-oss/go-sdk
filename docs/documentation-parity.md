# Java/Go documentation parity

The Go SDK follows the same documentation information architecture as the
[Java SDK](https://github.com/conductor-oss/java-sdk/tree/main/docs) while mapping
instructions to the Go public API. This page makes the intentional differences
explicit so a missing page is not mistaken for missing support — or for support
that does not exist.

| Java documentation capability | Go documentation counterpart | Status |
|---|---|---|
| Connection, quickstart, workflows, workers | [Core guides](README.md#build) | Supported with Go APIs |
| Security, observability, upgrading | [Operations guides](README.md#operate) | Supported with Go APIs |
| Server setup, workflow lifecycle, workflow testing | Not yet written | APIs exist; tracked in [#266](https://github.com/conductor-oss/go-sdk/issues/266) |
| Reliability, deployment/scaling, debugging | Not yet written | APIs exist; tracked in [#267](https://github.com/conductor-oss/go-sdk/issues/267) |
| Schema client, schedules and events | Not yet written | `SchemaClient`, `SchedulerClient`, `EventHandlerClient` exist; tracked in [#268](https://github.com/conductor-oss/go-sdk/issues/268) |
| API map, curated examples | Not yet written | Tracked in [#269](https://github.com/conductor-oss/go-sdk/issues/269) |
| Agent concepts, runtime, API/client, definition contract | **No counterpart** | The Go SDK does not provide a Conductor agent runtime. This is a capability gap, not a documentation gap. |
| Google ADK and framework bridges | **No counterpart** | Requires the agent runtime |
| Workflow-scoped `FileClient` | **No counterpart** | [Compatibility](compatibility.md#capability-differences-from-the-java-and-python-sdks) |
| Spring and Spring Boot integration | **No counterpart** | Java-specific; use Go application hosting patterns instead |

## Stated coverage gap

Of the 20 public client interfaces in `sdk/client`, **17 have no prose
documentation** at the time of writing. Ten of those have no canonical Java page
that would ever cover them, so they are not a parity gap and are not tracked as
one. They stay discoverable through the API map
([#269](https://github.com/conductor-oss/go-sdk/issues/269)), which is intended as
an exhaustive coverage ledger rather than an illustrative list.

The security surface is called out separately: `security.md` currently documents
TLS only, while Java's equivalent also covers secrets and access control. Tracked
in [#272](https://github.com/conductor-oss/go-sdk/issues/272).

## Maintenance rule

When adding a Java-style guide or a Go capability, update this map and the
[documentation hub](README.md) in the same change. Do not claim a Java-only client
or framework feature is available in Go — and do not imply Go lacks AI support
when what it lacks is the agent runtime.
