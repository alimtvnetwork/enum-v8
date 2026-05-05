# Upstream `core-v9` `go.mod` Module Path Mismatch

## Description

`enum-v3` source imports `github.com/alimtvnetwork/core-v9/...` but the upstream `core-v9` repository's own `go.mod` declares `module github.com/alimtvnetwork/core-v8`. A `replace` bridge in `enum-v3/go.mod` redirects resolution to the cached `core-v8` artifact, but Go 1.25 rejects this with `used for two different module paths`, and Go's `internal/` rule rejects any `enum-v3` consumer reaching into `core-v9/.../internal/...` because it sees the cached path as `core-v8`.

## Root Cause

Upstream `core-v9` repo has not yet had its `go.mod` rewritten to `module github.com/alimtvnetwork/core-v9`, nor has a release been tagged under the new path.

## Steps to Reproduce

1. `cd enum-v3`
2. `./run.ps1 -tc`
3. Observe: `STATUS ✗ BLOCKED` with downloads of `core-v8` artifacts and downstream `internal/` rejections.

Expected: clean build with `core-v9` resolved as `core-v9`.
Actual: dual-path rejection from Go 1.25.

## Attempted Solutions

- [x] Pin `replace` to a SHA-based pseudo-version `v0.0.0-20260423064907-72bcd64c06b5` — works for download, fails on `internal/` rule.
- [x] Pin `replace` to `v1.5.6` — Go 1.25 rejects with `used for two different module paths`.
- [x] Try pseudo-version `v1.5.6-0.<date>-<sha>` — module proxy rejects (no `v1.5.5` predecessor tag).
- [ ] Pin Go toolchain to 1.22 in `enum-v3/go.mod` — offered as stopgap, user has not accepted.
- [ ] **Real fix:** Edit upstream `core-v9` `go.mod` + tag `v1.5.8` (Task W) — manual user action.

## Priority

High — blocks all `core-v9` package consumption that touches `internal/`.

## Blocked By

Task W (upstream manual action, user-owned).

## Related

- `.lovable/cicd-issues/01-go-mod-rename-bridge.md`
- `.lovable/cicd-issues/02-internal-package-rule.md`
- `.lovable/cicd-issues/03-pseudo-version-rejected.md`
- `.lovable/memory/02-go-mod-bridge.md`
