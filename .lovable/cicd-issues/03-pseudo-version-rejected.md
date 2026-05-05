# 03 — Pseudo-Version `v1.5.6-0.<date>-<sha>` Rejected by Module Proxy

## Symptom

Attempting to pin the `replace` bridge with a pseudo-version like:

```
replace github.com/alimtvnetwork/core-v9 =>
    github.com/alimtvnetwork/core-v8 v1.5.6-0.20260504132614-f978d4d81b5e
```

…fails with `unknown revision` from the Go module proxy.

## Root Cause

Go's pseudo-version format `vX.Y.Z-0.<date>-<sha>` requires a **predecessor tag** of `vX.Y.(Z-1)` to exist on the source repo. There is no `v1.5.5` tag on `core-v8`, so the proxy refuses to construct the pseudo-version.

## Fix / Workaround

**Don't use this pattern.** Stick with the SHA-pinned pseudo-version `v0.0.0-<date>-<sha>` (which has no predecessor-tag requirement) until Task W lands and `v1.5.8` is published cleanly.

## Status

✅ Documented. Do not re-attempt this pattern.

## Related

- `.lovable/cicd-issues/01-go-mod-rename-bridge.md`
- `.lovable/strictly-avoid.md` § "Module paths & renames"
