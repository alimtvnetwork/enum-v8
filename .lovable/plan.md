# Active Plan — enum-v5

> Single source of truth for the project roadmap. Letter IDs are stable across sessions.
> Last updated: 2026-05-06 (Cycle 19 — AB pass 1 done; AJ-01..03 spawned, blocked by `spec/01-app/` freeze).

---

## Phase 1: Unblock Build (CRITICAL)

### W. Upstream `core-v9` `go.mod` rename + tag `v1.5.8`

- **Status:** ✅ Done (2026-05-05)

### AG. Drop the `replace` bridge and pin clean `core-v9 v1.5.8`

- **Status:** ✅ Done (2026-05-05)
- **Result:** `replace` directive removed, `require core-v9 v1.5.8` pinned.

### AM. Fix broken `core-v9` API calls (converter + coredynamic migration)

- **Status:** ✅ Done for reported blocker — `tests/creationtests` compile check passes in sandbox
- **Objective:** Update all `enum-v5` source files that use old `core-v8`-era function signatures (`converters.AnyToValueString`, `coredynamic.TypeName`, etc.) to the new struct-namespace API (`converters.AnyTo.ValueString`, `coredynamic.SafeTypeName`, etc.).
- **Dependencies:** None for the reported `creationtests` compile blocker.
- **Expected outputs:** All affected `.go` files updated, `go build ./...` passes.
- **Acceptance criteria:** `go build ./...` succeeds with `core-v9 v1.5.8`.
- **Reference:** `.lovable/memory/06-core-v9-api-migration.md`

---

## Phase 2: Spec Accuracy (HIGH)

### AA. Continue spec-audit cycles

- **Status:** 🔄 In Progress
- **Objective:** Audit remaining spec directories for code-vs-spec drift.
- **Dependencies:** None
- **Expected outputs:** Audit report per cycle in `spec/07-code-vs-spec-audits/`, scoreboard updates, spec fixes.
- **Acceptance criteria:** Each audited directory reaches ≥95% verifiable.
- **Cycle plan:**

| Cycle | Target | Priority | Reason |
|-------|--------|----------|--------|
| 15 | `spec/06-testing-guidelines/` (9 files) | **P0** | Most-referenced, known stale paths, never audited |
| 16 | `spec/03-powershell-test-run/` (9 files) | P1 | Toolchain docs, likely stale |
| 17 | `spec/04-tooling/` (10 files) | P1 | CI pipeline, bootstrap |
| 18 | `spec/02-app-issues/` (10 files) | P2 | All resolved, quick verification |
| 19 | `spec/05-failing-tests/` (25 files) | P3 | Post-mortems, reference-only |
| 20 | `spec/00-llm-integration-guide.md` | P1 | 2386-line monolith, needs reconciliation with audited specs |

### AH. Cross-`spec/` cleanup sweep

- **Status:** 🔄 In Progress — folded into upcoming directory audits
- **Objective:** Replace all stale `tests/integratedtests/` and `enum-v1` references outside `cross-repo/`.
- **Dependencies:** None
- **Expected outputs:** `rg integratedtests spec/` returns only anti-pattern callout lines.
- **Remaining targets:**
  - `spec/06-testing-guidelines/01-folder-structure.md` (line 13)
  - `spec/03-powershell-test-run/` (4 files, TBD)
  - `spec/04-tooling/04-bootstrap-into-new-repo.md`
  - `spec/02-app-issues/02-internal-package-coverage-policy.md`
  - `spec/00-llm-integration-guide.md` (line 36)

### AB. Fetch upstream `core-v9` source

- **Status:** ⏳ Pending
- **Objective:** Get upstream `core-v9` source into a workspace path so auditor can verify 148 ❓ claims.
- **Dependencies:** Fetch access to upstream repo
- **Expected outputs:** ❓ claims promoted to ✅ or ❌.
- **Acceptance criteria:** ❓ count drops to <10.

### AC. Re-audit §07 and §09

- **Status:** ⏳ Pending — waits on AB
- **Objective:** Apply spec-internal-consistency dimension to the two baseline-only sections.
- **Dependencies:** AB
- **Expected outputs:** Updated audit reports, possible spec fixes.

### AI. Mark `spec/01-app/` as frozen

- **Status:** ⏳ Pending
- **Objective:** Add freeze entry to `spec/CHANGELOG.md`.
- **Dependencies:** None (can do anytime)
- **Expected outputs:** Changelog entry stating frozen status.
- **Acceptance criteria:** `spec/CHANGELOG.md` has freeze entry with date and cycle reference.

---

## Phase 3: Implementation

### AJ. Implement spec fixes from Cycle 15 findings

- **Status:** 📋 Planned
- **Objective:** Apply all drift/contradiction fixes found in Cycle 15 (`spec/06-testing-guidelines/`).
- **Dependencies:** AA (Cycle 15)
- **Expected outputs:** Updated spec files, scoreboard, changelog bump.

### AK. New enum package creation (template validation)

- **Status:** 📋 Planned
- **Objective:** Create 1 new enum package end-to-end following spec §05 recipe, validating the spec is complete.
- **Dependencies:** AG (clean Go build)
- **Expected outputs:** New enum package with full method set, registered in `tests/creationtests/`.
- **Acceptance criteria:** `go build ./newpkg/...` and `go test ./tests/creationtests/...` pass.

### AL. Test coverage expansion

- **Status:** 📋 Planned
- **Objective:** Add test cases for untested enum packages using `tests/creationtests/` pattern.
- **Dependencies:** AG (clean Go build), AA Cycle 15 (accurate test guidelines)
- **Expected outputs:** New test files, coverage increase.

---

## Phase 4: Manual / Parked

### A. Manual `cross-repo/core-v8/` push

- **Status:** ⏭️ Manual user action (credential-bound)
- **Trigger:** When main-repo CI files change, mirror then user pushes.

---

## Next Task Selection

**Recommended next task:** Pick from this list:

1. **AA / Cycle 15** — Audit `spec/06-testing-guidelines/` (highest leverage spec task)
2. **AI** — Mark `spec/01-app/` as frozen (quick win, 5 minutes)
3. **AB** — Fetch upstream `core-v9` source for ❓ claim verification

---

## Completed

### Cycles 1–14 — `spec/01-app/` directory audit

- **Completed:** 2026-05-04 → 2026-05-05
- **Result:** All 14 numbered files audited. 12 at 100% verifiable. 2 baseline-only.
- **Findings:** C-CVS-01..10 (10 contradictions), D-CVS-01..43 (43 drifts), all resolved.

### W + AG — Unblock Build

- **Completed:** 2026-05-05
- **Result:** Upstream `core-v9` `go.mod` fixed, `replace` bridge removed, `require core-v9 v1.5.8` pinned.
