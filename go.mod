module github.com/alimtvnetwork/enum-v2

go 1.17.8

require (
	github.com/smartystreets/goconvey v1.8.1
	github.com/alimtvnetwork/core-v9 v0.0.0-20260504132614-f978d4d81b5e
	golang.org/x/sys v0.13.0
)

require (
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/smarty/assertions v1.15.1 // indirect
)

// ── Temporary rename bridge ──────────────────────────────────────────
// The upstream `core-v9` GitHub repo was renamed from `core-v8`, but its
// own `go.mod` on the `release/v1.5.7` branch STILL declares
// `module github.com/alimtvnetwork/core-v8` (verified 2026-05-05).
// Go enforces that the import path match the declared module path, so it
// rejects loading the v9 repo under the v9 path. Until upstream commits
// the `module github.com/alimtvnetwork/core-v9` line AND tags a release,
// we resolve the v9 import path to the v8 module artifact at the pinned
// commit (HEAD of release/v1.5.7).
//
// NOTE: Bumping this pointer pulls in the latest upstream code but does
// NOT fix the `internal/` blocker — Go enforces the internal/ rule
// against the cached module's declared path (`core-v8`), so consumers
// under `enum-v2/...` are still rejected for any package that
// transitively imports `core-v9/internal/...`.
//
// This block is removed (replaced by `v1.5.7` tag pin) once upstream:
//   1. Edits go.mod: `module github.com/alimtvnetwork/core-v9`
//   2. Rewrites all `core-v8` imports → `core-v9` in source
//   3. Tags `v1.5.7` on the renamed commit
replace github.com/alimtvnetwork/core-v9 => github.com/alimtvnetwork/core-v8 v0.0.0-20260504132614-f978d4d81b5e
