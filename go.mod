module github.com/alimtvnetwork/enum-v1

go 1.17.8

require (
	github.com/smartystreets/goconvey v1.8.1
	github.com/alimtvnetwork/core-v8 v1.5.5
	golang.org/x/sys v0.13.0
)

require (
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/smarty/assertions v1.15.1 // indirect
)

// core-v8 v1.5.5 lives on the feature/1.5.6 branch (no tagged release on default branch yet).
// Pseudo-version pinned to commit 72bcd64c06b5 (2026-04-23T06:49:07Z) on feature/1.5.6.
// Re-resolve with `go mod tidy` after upstream pushes new commits.
replace github.com/alimtvnetwork/core-v8 => github.com/alimtvnetwork/core-v8 v1.5.6-0.20260423064907-72bcd64c06b5
