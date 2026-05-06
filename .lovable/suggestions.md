# Suggestions

## Active Suggestions

### Pin Go toolchain to 1.22 as a stopgap for Task W

- **Status:** Obsolete — Task W is now ✅ Done. This stopgap is no longer needed.
- **Priority:** N/A
- **Description:** Go 1.25 rejects the dual-path `replace` bridge between `core-v9` (import path) and `core-v8` (cached module path) with `used for two different module paths`. Pinning the toolchain to Go 1.22 in `go.mod` would unblock builds while waiting for the upstream `core-v9` `go.mod` rename + `v1.5.8` tag (Task W). Trade-off: locks the project to an older toolchain and silently masks the underlying issue.
- **Added:** 2026-05-05 (Cycle 13 turn).
- **Superseded:** 2026-05-05 — Task W completed, bridge removed, suggestion no longer applicable.

### Promote `errcore.VarTwoNoType` from ❓ to ✅ in Cycle 6's §08 audit

- **Status:** Pending (cross-cycle promotion — depends on Cycle 13's new "spec-internal-consistency" dimension).
- **Priority:** Low
- **Description:** Cycle 13 introduced spec-internal-consistency as an explicit verifiability dimension. `errcore.VarTwoNoType` was scored ❓ in Cycle 6 row 16 because no `enum-v3` consumer imports it, but it IS cross-referenced from `04-error-system.md:131`, `08-validators.md:240,307,329`, and `15-observability.md`. Under the new dimension it would promote to ✅. Backport mechanically when Task AC runs.
- **Added:** 2026-05-05 (Cycle 13 §15 audit, table row 4).

## Implemented Suggestions

### Add `cmd/main/` smoke-test policy carve-out to spec §12

- **Implemented:** 2026-05-05 (Cycle 10, fix C-CVS-10).
- **Notes:** Spec §12 used to assert "no `cmd/` directory anywhere"; reality: `enum-v3/cmd/main/main.go` is a single permitted smoke-test harness. §12 §1 rewritten as a "library-first, smoke-test allowed" policy. Cross-link to `cmd/README.md` added.

### Rewrite §06 around `SimpleSlice`/`Hashset`/`SimpleStringOnce` instead of fictional `coreonce.New.String`

- **Implemented:** 2026-05-04 (Cycle 4, fixes D-CVS-22/23/24).
- **Notes:** Original §06 §3 / §5 referenced symbols that don't exist. Rewrote around the actual top-level constructors used by `enum-v3`.

### Add consumer-coverage callouts wherever the spec describes upstream-only API

- **Implemented:** 2026-05-04 → 2026-05-05 (D-CVS-25 in §06, D-CVS-38 in §13, D-CVS-42 in §14).
- **Notes:** Three sections now carry explicit "this surface has no `enum-v3` consumer; verify via Task AB" callouts so future readers don't assume verified ✅ status incorrectly.
