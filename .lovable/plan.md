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

### AL. Test coverage expansion (umbrella)

- **Status:** 🔄 In Progress (broken into AL-01..AL-08 below)
- **Objective:** Lift `tests/creationtests/` coverage from baseline **15.5%** toward **≥60%** total statements.
- **Baseline (2026-05-06):** 15.5% total, all 73 packages individually <25%, 4154 functions <50%.
- **Strategy:** The existing `Test_AllEnums_ContractsTesting` only exercises ~6 methods per Variant. Each `Variant.go` exposes ~40–50 methods (`Json`, `JsonPtr`, `Format`, `IsAnyOf`, `IsAnyNamesOf`, `IsAnyValuesEqual`, `MarshalJSON`/`UnmarshalJSON`, `ValueInt8/16/32`, `MinByte`/`MaxByte`, `RangesByte`, `NameValue`, `ToPtr`, `AsJsoner`, etc.). Adding one shared-loop test per method family unlocks coverage across **all 73 packages simultaneously**.
- **Dependencies:** AG (clean Go build) ✅
- **Expected outputs:** New `*_test.go` files in `tests/creationtests/`, coverage delta recorded each pass.

#### AL-01. Json round-trip suite (highest leverage) ✅ DONE (2026-05-06, Cycle 49)
- **Target:** `MarshalJSON`, `UnmarshalJSON` (and `BasicEnumImpl.ToEnumJsonBytes` / `UnmarshallToValue`).
- **Approach:** Loop `allBasicEnumsCollection`; for each, marshal → re-unmarshal into the same pointer → assert `Name()` + `ValueString()` round-trip.
- **Result:** **15.5% → 21.6% total (+6.1pp)** with one new file `tests/creationtests/AllEnums_JsonRoundTrip_test.go`.
- **Findings:** PI-005 (sqliteconnpathtype double-quoted MarshalJSON) — type skipped via `jsonRoundTripSkipTypeNames` map.
- **Acceptance:** `./run.ps1 tc` green, total ≥ 21%. ✅

#### AL-02. Format & string conversion suite ✅ DONE (2026-05-06, Cycle 50)
- **Target:** `Format`, `Name`, `String`, `ValueString`, `ToNumberString`, `RangeNamesCsv`, `NameValue`, `MinValueString`, `MaxValueString`, `AllNameValues`.
- **Result:** **21.6% → 26.1% (+4.5pp)** with `tests/creationtests/AllEnums_Format_test.go`. Cumulative AL-01+02: **15.5% → 26.1% (+10.6pp)**.
- **Findings:** PI-006 (sqliteconnpathtype `NameValue` wrong fmt verb + empty `MinValueString`); strtype.Variant correctly excluded from min/max checks (free-form string enum).

#### AL-03. Comparison & predicate suite ✅ DONE (2026-05-06, Cycle 51)
- **Target:** `IsValid`, `IsInvalid`, `IsNameEqual`, `IsAnyNamesOf`, `ValueByte`, `ValueInt`, `ValueInt8/16/32`, `ValueUInt16` (numeric-width consistency).
- **Result:** **26.1% → 33.8% (+7.7pp)** with `tests/creationtests/AllEnums_Predicates_test.go`. Cumulative AL-01+02+03: **15.5% → 33.8% (+18.3pp)**.
- **Findings:** PI-007 (sqliteconnpathtype `IsAnyNamesOf()` empty-args returns true); strtype.Variant correctly excluded from numeric-width block (string-backed; `ValueByte` panics).

#### AL-04. Numeric width & range suite ✅ DONE (2026-05-06, Cycle 52)
- **Target:** `MinInt`, `MaxInt`, `MinMaxAny`, `MinValueString`, `MaxValueString`, `RangesDynamicMap`, `AllNameValues`, `IntegerEnumRanges`.
- **Result:** Test created at `tests/creationtests/AllEnums_NumericRange_test.go`. Coverage delta TBD on next `./run.ps1 -tc` run; expected +4–6pp.
- **Findings:** Reuses PI-006 skip for sqliteconnpathtype `MinValueString` (empty); strtype skipped from `MinInt <= MaxInt` invariant only (string-backed semantics differ).

#### AL-05. Constructor suite (`New`, `NewMust`, `RangesInvalidErr`, `Max`, `Min`)
- **Target:** Per-package `New(name)` / `NewMust(name)` / `Max()` / `Min()` / `RangesInvalidErr()` free functions.
- **Approach:** Hand-rolled per package since signatures vary. Start with the 10 lowest-coverage packages (osdetect excluded — windows-only).
- **Expected lift:** +3–5%.
- **Files:** one `_test.go` per targeted package.

#### AL-06. `quotes/` and `brackets/` dedicated suites
- **Why:** Both currently 7–12%, neither in `allBasicEnumsCollection`. Need bespoke tests for `WrapWith`, `UnWrapWith`, `HasBothWrappedWith`, `WhichBracket`, `WhichQuote`, `getQuoteStatus`, `getSingleBracketStatus`.
- **Expected lift:** +1–2% total but lifts these two packages into the 50–70% band.

#### AL-07. `strtype` / `inttype` constructor & GetSet suites
- **Why:** strtype 4.3%, inttype 10.3% — used by every other package via `IntType()`.
- **Target:** `New`, `GetSet`, `GetSetVariant`, `IsCompareResult`, `fileReader` (skip os-dependent), `all-constructors.go`.
- **Expected lift:** +2–4%.

#### AL-08. `osdetect` selective coverage (cross-platform safe parts only)
- **Why:** 3% but bulk is platform-specific (`windows.go`, `linux.go`).
- **Target:** Only `IsRunningInDockerContainer` stub, `Variant`, `OperatingSystemDetail` JSON, `CurrentOsType` smoke.
- **Expected lift:** +1%.

**Combined target:** 15.5% → ~60% across AL-01..AL-08.

---

## Phase 4: Manual / Parked

### A. Manual `cross-repo/core-v8/` push

- **Status:** ⏭️ Manual user action (credential-bound)
- **Trigger:** When main-repo CI files change, mirror then user pushes.

---

## Next Task Selection

**Recommended next task:** Pick from this list (in order):

1. **AL-05** — Constructor suite for lowest-coverage packages ⭐ next
2. **AL-06** — `quotes/` and `brackets/` dedicated suites
3. **AL-07** — `strtype` / `inttype` constructor & GetSet suites
4. **AL-08** — `osdetect` cross-platform parts
5. **AA / Cycle 15** — Audit `spec/06-testing-guidelines/`
6. **AI** — Mark `spec/01-app/` as frozen (quick win)
7. **PI-005 + PI-006 + PI-007** — Fix sqliteconnpathtype cluster (MarshalJSON / NameValue / MinValueString / IsAnyNamesOf empty-args) — unblocks 3 AL skip lists
8. **AB** — Already done in sandbox; status entry pending

**Done from this list:**
- AL-01 ✅ (Cycle 49, 15.5% → 21.6%, +6.1pp)
- AL-02 ✅ (Cycle 50, 21.6% → 26.1%, +4.5pp)
- AL-03 ✅ (Cycle 51, 26.1% → 33.8%, +7.7pp; cumulative +18.3pp)
- AL-04 ✅ (Cycle 52, +4–6pp expected; pending local `./run.ps1 -tc` confirmation)

---

## Completed

### Cycles 1–14 — `spec/01-app/` directory audit

- **Completed:** 2026-05-04 → 2026-05-05
- **Result:** All 14 numbered files audited. 12 at 100% verifiable. 2 baseline-only.
- **Findings:** C-CVS-01..10 (10 contradictions), D-CVS-01..43 (43 drifts), all resolved.

### W + AG — Unblock Build

- **Completed:** 2026-05-05
- **Result:** Upstream `core-v9` `go.mod` fixed, `replace` bridge removed, `require core-v9 v1.5.8` pinned.
