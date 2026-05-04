# 01 — Code-vs-Spec Drift Scoreboard (Living Document)

> **Single source of truth** for code-vs-spec drift. Updated after every cycle.

## Current MEASURED drift score: **§03 100.0 / §04 100.0 (verifiable subset)** *(2 sections audited, all verifiable claims matched)*

> §03 closed at 12/12. §04 (error-system) extended in spec-v0.22.0 to document the 8 previously-undocumented APIs (`MustBeEmpty`, `RawErrCollection`, `ErrorRefOnly`, `CombineWithAnother`, `MessageWithRef`, `RangeNotMeet`, `ToError`/`ToString`, plus 4 missing `RawErrorType` examples). All 11 verifiable claims now match. 7 ❓ remain pending upstream `core-v9` source fetch (task **AB**).

## Cycle history

| Date | Cycle | Spec audited | Claims | ✅ Match | ⚠️ Drift | ❌ Contradiction | ❓ Unverifiable | Score |
|------|-------|--------------|--------|---------|---------|------------------|----------------|-------|
| 2026-05-04 | 1 (baseline) | `01-app/03-import-conventions.md` | 12 | 5 | 5 | 2 | 0 | **41.7%** |
| 2026-05-04 | 1 (post-LOW) | `01-app/03-import-conventions.md` | 12 | 10 | 0 | 2 | 0 | **83.3%** |
| 2026-05-04 | 1 (closed) | `01-app/03-import-conventions.md` | 12 | 12 | 0 | 0 | 0 | **100.0%** |
| 2026-05-04 | 2 (baseline) | `01-app/04-error-system.md` | 18 | 3 | 8 | 0 | 7 | **27.3%** *(verifiable)* |
| 2026-05-04 | 2 (closed) | `01-app/04-error-system.md` | 18 | 11 | 0 | 0 | 7 | **100.0%** *(verifiable)* |

## Open drift findings

| ID | Title | Severity | Spec ref | Code ref | Resolution path |
|----|-------|----------|----------|----------|-----------------|
| — | None (D-CVS-06..13 resolved in spec-v0.22.0) | — | — | — | — |

## Resolved drift findings

| ID | Title | Resolved at | Fix location | Path taken |
|----|-------|-------------|--------------|------------|
| D-CVS-01 | Spec §03 line 4 says "consumes `core-v8`" — stale | 2026-05-04 | `spec/01-app/03-import-conventions.md:4` | s/core-v8/core-v9/ |
| D-CVS-02 | Spec §03 line 88 says path "ends in `core-v8`" — stale | 2026-05-04 | `spec/01-app/03-import-conventions.md:88` | s/core-v8/core-v9/ + s/corev8/corev9/ |
| D-CVS-03 | Spec §03 line 94 prose/example mismatch (v8 vs v9) | 2026-05-04 | `spec/01-app/03-import-conventions.md:94` | s/core-v8/core-v9/ |
| D-CVS-04 | Spec §03 line 121 conflates "test module" with "core module" | 2026-05-04 | `spec/01-app/03-import-conventions.md:121` | Reworded to be module-generic |
| D-CVS-05 | `coregeneric` canonical-import listing not annotated as optional | 2026-05-04 | `spec/01-app/03-import-conventions.md:61,73` | Inline `// optional` + consumer-coverage note |
| C-CVS-01 | Spec §03 line 129 references nonexistent `tests/integratedtests/` | 2026-05-04 | `spec/01-app/03-import-conventions.md:127-145` | Spec rewritten to `tests/creationtests/` layout (matches code) + cross-ref to upstream `core-v9` per-suite layout |
| C-CVS-02 | Spec §03 line 118 `internal/reflectinternal` example doesn't apply to this repo | 2026-05-04 | `spec/01-app/03-import-conventions.md:113-123` | Section reframed as "in upstream `core-v9` tests"; consumer-side note added explaining cross-module `internal/` is forbidden |
| D-CVS-06 | `errcore.MustBeEmpty` undocumented | 2026-05-04 | `spec/01-app/04-error-system.md` §1 + §1.2 code block | Added table row + code example; clarified vs `HandleErr` |
| D-CVS-07 | `errcore.RawErrCollection` undocumented | 2026-05-04 | `spec/01-app/04-error-system.md` §1.6 (new) | Added "Accumulating Errors" subsection with `osdetect` reference |
| D-CVS-08 | `<RawErrorType>.ErrorRefOnly` undocumented | 2026-05-04 | `spec/01-app/04-error-system.md` §1.2 | Added constructor table row + example |
| D-CVS-09 | `<RawErrorType>.CombineWithAnother` undocumented | 2026-05-04 | `spec/01-app/04-error-system.md` §1.2 | Added as legacy alias of `MergeError` |
| D-CVS-10 | `errcore.MessageWithRef` undocumented | 2026-05-04 | `spec/01-app/04-error-system.md` §1.5 (new) | Added "Reference Helpers" subsection |
| D-CVS-11 | `errcore.RangeNotMeet` undocumented | 2026-05-04 | `spec/01-app/04-error-system.md` §1.5 (new) | Documented alongside `MessageWithRef` |
| D-CVS-12 | `errcore.ToError` / `ToString` undocumented | 2026-05-04 | `spec/01-app/04-error-system.md` §1.7 (new) | Added "Conversion Helpers" subsection |
| D-CVS-13 | `RawErrorType` §1.1 examples incomplete | 2026-05-04 | `spec/01-app/04-error-system.md` §1.1 | Added `FailedToExecuteType`, `NotSupportedType`, `PathInvalidErrorType`, `ComparatorShouldBeWithinRangeType`, `FailedToConvertType` |

## Targets

| Milestone | Score | Status |
|-----------|-------|--------|
| ✅ First measured baseline (Cycle 1) | **41.7** | 2026-05-04 |
| ✅ Apply 5 LOW spec fixes from Cycle 1 (D-CVS-01..05) | **83.3** on §03 | 2026-05-04 |
| ✅ Resolve C-CVS-01 + C-CVS-02 → §03 at 100% | **100.0** on §03 | 2026-05-04 |
| ✅ Cycle 2 baseline on §04 | **27.3** verifiable on §04 | 2026-05-04 |
| 🚧 Apply MED + LOW spec fixes for §04 (D-CVS-06..13) | target ≥ 90% on §04 | pending |
| 🚧 Fetch `core-v9` source (task **AB**) → resolve 7 ❓ on §04 | — | pending |
| 🚧 Audit all 16 sections of `01-app/` | 16/16 | 2/16 done |
| 🎯 Reach ≥95% aggregate match rate | ≥ 95 | Pending |
| 🎯 Zero ❌ contradictions | 0 (currently 0) | ✅ |
