module github.com/alimtvnetwork/enum-v2

go 1.17.8

require (
	github.com/smartystreets/goconvey v1.8.1
	github.com/alimtvnetwork/core-v9 v1.5.6
	golang.org/x/sys v0.13.0
)

require (
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/smarty/assertions v1.15.1 // indirect
)

// ── Temporary rename bridge ──────────────────────────────────────────
// The upstream `core-v9` GitHub repo was renamed from `core-v8`, but its
// own `go.mod` (verified 2026-05-05 against tag v1.5.7) STILL declares
// `module github.com/alimtvnetwork/core-v8`. Go enforces that the import
// path match the declared module path, so it rejects loading the v9
// repo under the v9 path. Until upstream edits go.mod to
// `module github.com/alimtvnetwork/core-v9` and re-tags, we resolve the
// v9 import path to the v8 module artifact.
//
// PIN: `v1.5.6` is the latest tag actually published under the
// `core-v8` URL on the Go module proxy. Earlier pseudo-version pins
// (e.g. `v0.0.0-20260504132614-f978d4d81b5e`) referenced commits that
// only exist under the renamed `core-v9` URL, so the proxy 404'd them
// when fetched via `core-v8` — that was the root cause of the
// "invalid version: unknown revision" failures in run.ps1 -tc.
//
// NOTE: Even with this pin working, the bridge is INSUFFICIENT for any
// core-v9 package that transitively imports an `internal/` package —
// Go's internal/ rule is enforced against the cached module's declared
// path (`core-v8`), so consumers under `enum-v2/...` are rejected. The
// only real fix is task **W** (upstream go.mod + source rename + new
// tag), which then unblocks task **AG** (drop this entire replace
// block, pin `core-v9 v1.5.8` cleanly).
replace github.com/alimtvnetwork/core-v9 => github.com/alimtvnetwork/core-v8 v1.5.6
