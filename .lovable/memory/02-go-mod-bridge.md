# `go.mod` rename bridge

## The problem

`enum-v2` source code imports `github.com/alimtvnetwork/core-v9/...`. The upstream `core-v9` repository's own `go.mod` STILL declares:

```
module github.com/alimtvnetwork/core-v8
```

This mismatch means a clean `require github.com/alimtvnetwork/core-v9 vX.Y.Z` would fail: Go can't resolve a module whose declared path differs from the import path.

## The bridge (current state)

`enum-v2/go.mod` carries:

```
replace github.com/alimtvnetwork/core-v9 => github.com/alimtvnetwork/core-v8 v0.0.0-20260423064907-72bcd64c06b5
```

This redirects `core-v9` import-path resolution to the cached `core-v8` module artifact.

## Why the bridge is INSUFFICIENT

Go's `internal/` rule is enforced against the **cached module's declared path** — `core-v8`. Any consumer under `enum-v2/...` that transitively imports a `core-v9/.../internal/...` package is rejected, because to the toolchain it looks like a foreign module reaching into another module's `internal/`.

Symptom (Go 1.25):

```
github.com/alimtvnetwork/core-v8@v1.5.6 used for two different module paths
```

## Failed attempts

- **Pin `core-v8 v1.5.6`** — Go 1.25 explicitly rejects dual-path use.
- **Pseudo-version `v1.5.6-0.<date>-<sha>`** — Needs a `v1.5.5` predecessor tag that doesn't exist; module proxy rejects.
- **Pin SHA-only pseudo-version** — Same: requires a chained predecessor tag.

## The only real fix (Task W)

1. Edit upstream `core-v9` repo's `go.mod`:
   ```
   module github.com/alimtvnetwork/core-v9
   ```
2. Sweep its source imports `core-v8` → `core-v9`.
3. Tag a release `v1.5.8`.
4. In `enum-v2`: drop the `replace` directive and pin `require github.com/alimtvnetwork/core-v9 v1.5.8` (Task AG).

## Stopgap option (suggestion, not yet accepted)

Pin the toolchain to Go 1.22 in `enum-v2/go.mod` — older Go is more permissive about the dual-path situation. Trade-off: silently masks the issue and locks the project to an older toolchain.

## Don'ts

- Do NOT reintroduce `core-v8` in source imports outside `cross-repo/core-v8/`.
- Do NOT propose pseudo-versions that require non-existent predecessor tags.
- Do NOT delete the `replace` directive without first landing W.
