# Active Plan — enum-v2

> Single source of truth for the project roadmap. Letter IDs are stable across sessions.
> Last updated: 2026-05-05 (after Cycle 14 / `spec/01-app/16-security.md` audit).

## Active tasks

### W. Upstream `core-v9` `go.mod` rename + tag `v1.5.8`

- **Status:** 🚫 Blocked — manual upstream action required (user must edit a different repo).
- **Why it matters:** Go 1.25 rejects the current dual-path `replace` bridge with `core-v8@v1.5.6 used for two different module paths`. Until upstream's own `go.mod` declares `module github.com/alimtvnetwork/core-v9` and a release is tagged, `enum-v2` cannot consume any `core-v9` package that transitively imports an `internal/` package.
- **Acceptance:** Upstream `core-v9` repo `go.mod` line 1 = `module github.com/alimtvnetwork/core-v9`; release `v1.5.8` tagged.

### AG. Drop the `replace` bridge and pin clean `core-v9 v1.5.8`

- **Status:** ⏳ Pending — waits on **W**.
- **Acceptance:** `enum-v2/go.mod` carries `require github.com/alimtvnetwork/core-v9 v1.5.8` with no `replace` directive; `./run.ps1 -tc` passes.

### AA. Continue spec-audit cycles (next directory)

- **Status:** 🔄 In Progress.
- **Done in this loop:** Cycles 1–14 closed `spec/01-app/§03–§16` (12 sections at 100% verifiable, 2 baseline-only).
- **Next target (Cycle 15):** `spec/06-testing-guidelines/` (highest leverage — heavily referenced by §13/§14/§15, never audited; combine with AH stale-path sweep).
- **Other candidates:** `spec/02-app-issues/`, `spec/03-powershell-test-run/`, `spec/04-tooling/`.

### AB. Fetch upstream `core-v9` source to resolve ❓ claims

- **Status:** ⏳ Pending — needs a checkout/fetch of the upstream repo into a workspace path that the auditor can `rg`.
- **Why:** 148 ❓ claims across `spec/01-app/` (17 §07 + 18 §08 + 23 §09 + 15 §10 + 11 §11 + 6 §12 + 8 §13 + 10 §14 + 13 §15 + 13 §16 + 7 §04 + 1 §05 + 6 §06).

### AC. Re-audit §07 and §09 against spec-internal-consistency dimension

- **Status:** ⏳ Pending — new dimension introduced in Cycle 13. Both sections currently sit at "N/A — no verifiable subset" but likely have promotable claims (cross-ref resolution, no-contradiction checks).

### AH. Cross-`spec/` cleanup sweep

- **Status:** 🔄 In Progress — to be folded into upcoming directory audits.
- **Remaining targets:**
  - `spec/03-powershell-test-run/` (4 files) — sweep for `tests/integratedtests/`.
  - `spec/04-tooling/04-bootstrap-into-new-repo.md` — same.
  - `spec/02-app-issues/02-internal-package-coverage-policy.md` — same.
  - Stray `enum-v1` strings outside `cross-repo/core-v8/`.

### AI. Mark `spec/01-app/` as frozen for code-vs-spec drift

- **Status:** ⏳ Pending — opened in Cycle 14.
- **Acceptance:** `spec/CHANGELOG.md` carries an entry: "`spec/01-app/` frozen at 2026-05-05 — Cycle 14; 12/14 sections at 100% verifiable, 2 baseline-only (§07, §09) deferred to AB."

### A. Manual `cross-repo/core-v8/` push

- **Status:** ⏭️ Manual user action (credential-bound). Parked under the user's responsibility.
- **Trigger:** Whenever main-repo CI (`.github/workflows/`, `.golangci.yml`, `.github/dependabot.yml`, `.ci-baselines/`, `scripts/ci/test_*.py`, `CHANGELOG.md`) changes — mirror first, then user pushes.

## Completed

### Cycles 1–14 — `spec/01-app/` directory audit

- **Completed:** 2026-05-04 → 2026-05-05.
- **Result:** All 14 numbered files audited. 12 closed at 100% verifiable. 2 baseline-only (§07, §09) — no verifiable subset until AB lands.
- **Drift findings resolved:** C-CVS-01 through C-CVS-10, D-CVS-01 through D-CVS-43.
- **Evidence:** `spec/07-code-vs-spec-audits/02-cycle1-…` through `15-cycle14-security.md` + scoreboard in `01-scoreboard.md`.
- **Per-file table:**

| File | Cycle | Status |
|---|---|---|
| `03-import-conventions.md` | 1  | ✅ 100.0% |
| `04-error-system.md` | 2  | ✅ 100.0% |
| `05-enum-system.md` | 3  | ✅ 100.0% |
| `06-data-structures.md` | 4  | ✅ 100.0% |
| `07-conditional-and-utilities.md` | 5  | ⚪ baseline-only |
| `08-validators.md` | 6  | ✅ 100.0% |
| `09-converters.md` | 7  | ⚪ baseline-only |
| `10-reflection-and-dynamic.md` | 8  | ✅ 100.0% |
| `11-versioning.md` | 9  | ✅ 100.0% |
| `12-cmd-entrypoints.md` | 10 | ✅ 100.0% |
| `13-testing-patterns.md` | 11 | ✅ 100.0% |
| `14-tests-folder-walkthrough.md` | 12 | ✅ 100.0% |
| `15-observability.md` | 13 | ✅ 100.0% (zero edits) |
| `16-security.md` | 14 | ✅ 100.0% (zero edits) |
