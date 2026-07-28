# Documentation standard

Every primary guide must include:

- Audience and prerequisites.
- An explicit OSS/Orkes capability label when behavior differs.
- A security note when credentials, user data, or external side effects are involved.
- Runnable commands or a clear **Fragment** label linked to a complete example under `examples/`.
- Expected result, common failure modes, cleanup where resources are started, and next-step links.

Keep examples repository-native and compiling against the module in this repo.
Use `go get github.com/conductor-sdk/conductor-go` rather than stale version
literals, and treat [pkg.go.dev](https://pkg.go.dev/github.com/conductor-sdk/conductor-go)
as the signature source of truth. CI validates internal links and curated paths.

## Pages that predate this standard

This standard was adopted **forward-looking**. It binds new and edited pages. The
guides that existed when it was introduced — `workers.md`, `workflows.md`,
`upgrading.md`, `security.md`, `observability.md`, and
`connection-authentication.md` — do not yet conform, deliberately: retrofitting
them in the same change as a large mechanical rename would have made the diff
unreviewable.

That gap is tracked in
[#270](https://github.com/conductor-oss/go-sdk/issues/270), not left implicit.
Conformance is not enforced mechanically in CI; link and path validity is.
