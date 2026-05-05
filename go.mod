module github.com/alimtvnetwork/enum-v2

go 1.17.8

require (
	github.com/smartystreets/goconvey v1.8.1
	github.com/alimtvnetwork/core-v9 v0.0.0-20260423064907-72bcd64c06b5
	golang.org/x/sys v0.13.0
)

require (
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/smarty/assertions v1.15.1 // indirect
)

// ── Temporary rename bridge ──────────────────────────────────────────
// The upstream `core-v9` GitHub repo was renamed from `core-v8`, but its
// own `go.mod` still declares `module github.com/alimtvnetwork/core-v8`.
// Go enforces that the import path match the declared module path, so it
// rejects loading the v9 repo under the v9 path. Until the upstream
// `go.mod` is updated to declare `core-v9`, we resolve the v9 import
// path to the v8 module artifact at the pinned commit.
//
// This block is removed (replaced by a real-tag pin) once upstream
// commits the `module github.com/alimtvnetwork/core-v9` line and tags a
// release.
replace github.com/alimtvnetwork/core-v9 => github.com/alimtvnetwork/core-v8 v0.0.0-20260423064907-72bcd64c06b5
