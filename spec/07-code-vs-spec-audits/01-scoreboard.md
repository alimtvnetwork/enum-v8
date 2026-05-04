# 01 — Code-vs-Spec Drift Scoreboard (Living Document)

> **Single source of truth** for code-vs-spec drift. Updated after every cycle.

## Current MEASURED drift score: **§03 100.0 / §04 27.3 (verifiable subset)** *(2 sections audited)*

> §03 closed at 12/12. §04 (error-system) audited in Cycle 2 with 18 claims: 3 ✅, 8 ⚠️ (spec is incomplete vs consumer usage), 7 ❓ (unverifiable without upstream `core-v9` source), 0 ❌. Score on the verifiable 11 = 27.3%. Sandbox lacks Go + upstream source; aspirational APIs intentionally marked ❓ rather than ❌ to avoid false contradictions. See [`03-cycle2-error-system.md`](./03-cycle2-error-system.md).

## Cycle history

| Date | Cycle | Spec audited | Claims | ✅ Match | ⚠️ Drift | ❌ Contradiction | ❓ Unverifiable | Score |
|------|-------|--------------|--------|---------|---------|------------------|----------------|-------|
| 2026-05-04 | 1 (baseline) | `01-app/03-import-conventions.md` | 12 | 5 | 5 | 2 | 0 | **41.7%** |
| 2026-05-04 | 1 (post-LOW) | `01-app/03-import-conventions.md` | 12 | 10 | 0 | 2 | 0 | **83.3%** |
| 2026-05-04 | 1 (closed) | `01-app/03-import-conventions.md` | 12 | 12 | 0 | 0 | 0 | **100.0%** |
| 2026-05-04 | 2 (baseline) | `01-app/04-error-system.md` | 18 | 3 | 8 | 0 | 7 | **27.3%** *(verifiable)* |

## Open drift findings

| ID | Title | Severity | Spec ref | Code ref | Resolution path |
|----|-------|----------|----------|----------|-----------------|
| D-CVS-06 | `errcore.MustBeEmpty` undocumented | MED | `01-app/04-error-system.md` §1 | 8+ call sites incl. `dbdrivertype/connectionStringCompiler.go:144` | Add row to §1 table; clarify vs `HandleErr` |
| D-CVS-07 | `errcore.RawErrCollection` undocumented | MED | `01-app/04-error-system.md` §1 | `osdetect/windowsSystemDetailGenerator_windows.go:16` | Add §1.5 "Error Accumulation" |
| D-CVS-08 | `<RawErrorType>.ErrorRefOnly` undocumented | MED | `01-app/04-error-system.md` §1.2 | `errcore.OutOfRangeType.ErrorRefOnly` etc. | Add row to §1.2 constructor table |
| D-CVS-09 | `<RawErrorType>.CombineWithAnother` undocumented | LOW | §1.2 | `errcore.FailedToParseType.CombineWithAnother` | Add to §1.2 + cross-link `MergeError` |
| D-CVS-10 | `errcore.MessageWithRef` undocumented | LOW | §1.4 | source usage | Add row to §1.4 |
| D-CVS-11 | `errcore.RangeNotMeet` undocumented | LOW | §1 | source usage | Add §1.6 "Enum-Specific Builders" |
| D-CVS-12 | `errcore.ToError` / `ToString` undocumented | LOW | §1 | `osdetect/vars.go:111` | Add §1.7 "Conversion Helpers" |
| D-CVS-13 | `RawErrorType` examples missing `FailedToExecuteType`/`NotSupportedType`/`PathInvalidErrorType`/`ComparatorShouldBeWithinRangeType` | LOW | §1.1 | direct usage | Expand §1.1 examples or footnote upstream enumeration |

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
