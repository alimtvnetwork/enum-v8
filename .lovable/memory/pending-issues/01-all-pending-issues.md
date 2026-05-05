# Pending Issues — Consolidated Tracker

> Single file for all pending issues. Update in-place.
> When resolved, move to Resolved section with date.

---

## Open Issues

### PI-001: Upstream `core-v9` `go.mod` module path mismatch (Task W)

- **severity:** CRITICAL / BLOCKING
- **description:** Upstream `core-v9` repo's `go.mod` still declares `module github.com/alimtvnetwork/core-v8`. Go 1.25 rejects the `replace` bridge. Any `core-v9` package importing `internal/` is rejected for `enum-v3` consumers.
- **owner:** User (manual upstream action)
- **blocks:** Task AG (drop replace bridge), full Go build, any `internal/`-dependent imports
- **attempted fixes:** pseudo-version pin, v1.5.6 pin, Go 1.22 toolchain pin (not accepted)
- **real fix:** Edit upstream `go.mod` line 1 → `module github.com/alimtvnetwork/core-v9`, tag `v1.5.8`

### PI-002: Cross-spec stale `integratedtests/` paths (Task AH)

- **severity:** HIGH
- **description:** Multiple spec files outside `spec/01-app/` still reference `tests/integratedtests/` which doesn't exist. Known targets: `spec/06-testing-guidelines/01-folder-structure.md`, `spec/03-powershell-test-run/`, `spec/04-tooling/04-bootstrap-into-new-repo.md`, `spec/02-app-issues/02-internal-package-coverage-policy.md`, `spec/00-llm-integration-guide.md` line 36.
- **owner:** AI (audit cycles)
- **plan:** Fold into Cycle 15+ directory audits

### PI-003: 148 ❓ claims unresolved (Task AB)

- **severity:** MEDIUM
- **description:** 148 claims across `spec/01-app/` scored ❓ because they reference upstream `core-v9` APIs that have zero `enum-v3` consumers and no mirrored source. Need upstream source fetch to verify.
- **owner:** AI + User (fetch access)
- **blocks:** Full spec verification, re-audit of §07/§09 (Task AC)

### PI-004: `spec/06-testing-guidelines/` never audited

- **severity:** HIGH
- **description:** 9 files, most-referenced spec directory for any implementation AI. Contains known stale `integratedtests/` reference. Cycle 15 target.
- **owner:** AI
- **plan:** Next audit cycle (AA)

---

## Resolved Issues

_(Consolidated from `.lovable/pending-issues/01-core-v9-go-mod-rename.md` and `02-cross-spec-stale-paths.md` — those files are now superseded by this tracker)_
