module github.com/alimtvnetwork/enum-v1

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

// core-v9 has no v1.5.5 tag yet (rename from core-v8 in flight; the v8 repo
// had v1.5.5, v9 does not). The v1.5.6-0.<date>-<sha> pseudo-version Go
// previously expected requires a v1.5.5 predecessor — switched to a
// v0.0.0-<date>-<sha> pseudo-version which has no predecessor requirement.
// Re-pin to a real tag (e.g. v1.5.6) once upstream tags it on core-v9.
replace github.com/alimtvnetwork/core-v9 => github.com/alimtvnetwork/core-v9 v0.0.0-20260423064907-72bcd64c06b5
