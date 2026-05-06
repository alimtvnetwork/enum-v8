# Pending Issues — Consolidated Tracker

> Single file for all pending issues. Update in-place.
> When resolved, move to Resolved section with date.

---

## Open Issues

*(PI-001 moved to Resolved — see below)*

### PI-002: Cross-spec stale `integratedtests/` paths (Task AH)

- **severity:** HIGH
- **description:** Multiple spec files outside `spec/01-app/` still reference `tests/integratedtests/` which doesn't exist. Known targets: `spec/06-testing-guidelines/01-folder-structure.md`, `spec/03-powershell-test-run/`, `spec/04-tooling/04-bootstrap-into-new-repo.md`, `spec/02-app-issues/02-internal-package-coverage-policy.md`, `spec/00-llm-integration-guide.md` line 36.
- **owner:** AI (audit cycles)
- **plan:** ✅ COMPLETE — resolved by re-framing (consumer-coverage callouts) across Cycles 12/15/17/18 + Cycle 34 (S-003/S-004 close-out). All `integratedtests/` references now correctly document the upstream `core-v9` consumer layout with explicit upstream-vs-enum-v5 scope disclaimers. Move to Resolved on the next cycle that touches this file.

### PI-003: 148 ❓ claims unresolved (Task AB)

- **severity:** MEDIUM
- **description:** 148 claims across `spec/01-app/` scored ❓ because they reference upstream `core-v9` APIs that have zero `enum-v5` consumers and no mirrored source. Need upstream source fetch to verify.
- **owner:** AI + User (fetch access)
- **blocks:** Full spec verification, re-audit of §07/§09 (Task AC)

### PI-004: `spec/06-testing-guidelines/` never audited

- **severity:** HIGH
- **description:** 9 files, most-referenced spec directory for any implementation AI. Contains known stale `integratedtests/` reference. Cycle 15 target.
- **owner:** AI
- **plan:** Next audit cycle (AA)

---

## Resolved Issues

### D-CVS-62: Missing `scripts/coverage/` utilities

- **resolved:** 2026-05-06 (Cycles 31 + 32, S-108 + S-110)
- **fix:** Cycle 31 (S-108) restored the auto-invoked `Generate-CoveragePrompts.ps1`. Cycle 32 (S-110) restored the three standalone utilities (`Get-UncoveredLines.ps1`, `Get-FunctionCoverage.ps1`, `Get-PackageCoverageReport.ps1`). All four scripts smoke-tested via nix-pwsh. `scripts/coverage/` and `spec/03-powershell-test-run/06-coverage-prompt-generator.md` are now in lockstep.

### PI-001: Upstream `core-v9` `go.mod` module path mismatch (Task W + AG)

- **resolved:** 2026-05-05
- **fix:** User renamed upstream `go.mod` → `module github.com/alimtvnetwork/core-v9`, tagged `v1.5.8`. AI dropped `replace` bridge in `enum-v5/go.mod`, pinned `core-v9 v1.5.8`.

_(Consolidated from `.lovable/pending-issues/01-core-v9-go-mod-rename.md` and `02-cross-spec-stale-paths.md` — those files are now superseded by this tracker)_
