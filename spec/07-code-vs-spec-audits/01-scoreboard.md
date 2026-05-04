# 01 — Code-vs-Spec Drift Scoreboard (Living Document)

> **Single source of truth** for code-vs-spec drift. Updated after every cycle.

## Current MEASURED drift score: **§03 100.0 / §04 100.0 / §05 47.1 (verifiable)** *(3 sections audited)*

> §03 + §04 closed. §05 (enum-system) Cycle 3 baseline: 18 claims, 8 ✅ / 6 ⚠️ / **3 ❌** / 1 ❓. The 3 contradictions are real — spec mandates "first constant must be `Invalid`" but 10 enum packages put a non-Invalid sentinel first (`Default`, `Unspecified`, `Uninitialized`, `InvalidIndex = -1`). The 6 drifts include a fictional `consts.go` file split, a `reflectinternal.TypeName(...)` example that can't compile from `enum-v1`, and a stale `tests/integratedtests/` reference (same C-CVS-01 issue already fixed in §03). See [`04-cycle3-enum-system.md`](./04-cycle3-enum-system.md).

## Cycle history

| Date | Cycle | Spec audited | Claims | ✅ Match | ⚠️ Drift | ❌ Contradiction | ❓ Unverifiable | Score |
|------|-------|--------------|--------|---------|---------|------------------|----------------|-------|
| 2026-05-04 | 1 (baseline) | `01-app/03-import-conventions.md` | 12 | 5 | 5 | 2 | 0 | **41.7%** |
| 2026-05-04 | 1 (post-LOW) | `01-app/03-import-conventions.md` | 12 | 10 | 0 | 2 | 0 | **83.3%** |
| 2026-05-04 | 1 (closed) | `01-app/03-import-conventions.md` | 12 | 12 | 0 | 0 | 0 | **100.0%** |
| 2026-05-04 | 2 (baseline) | `01-app/04-error-system.md` | 18 | 3 | 8 | 0 | 7 | **27.3%** *(verifiable)* |
| 2026-05-04 | 2 (closed) | `01-app/04-error-system.md` | 18 | 11 | 0 | 0 | 7 | **100.0%** *(verifiable)* |
| 2026-05-04 | 3 (baseline) | `01-app/05-enum-system.md` | 18 | 8 | 6 | 3 | 1 | **47.1%** *(verifiable)* |

## Open drift findings

| ID | Title | Severity | Spec ref | Code ref | Resolution path |
|----|-------|----------|----------|----------|-----------------|
| C-CVS-03 | Spec mandates first const = `Invalid`; 10 enums use other sentinels | HIGH | `01-app/05-enum-system.md` §4 Step 1 | `compressformats`, `compresslevels`, `envtype`, `inttype`, `logtype`, `revokereason`, `scripttype`, `sqljointype`, `strtype`, `taskpriority` | Reframe as "sentinel first" with allowed names (`Invalid`, `Default`, `Unspecified`, `Uninitialized`, `InvalidIndex`) |
| C-CVS-04 | Recipe imports `core-v9/internal/reflectinternal` (cross-module `internal/` is forbidden by Go) | HIGH | `01-app/05-enum-system.md` §4 Step 2 | zero packages do this | Replace with string-literal type name OR `DefaultAllCases(firstItem, ranges[:])` |
| C-CVS-05 | "Zero-value sentinel" rule contradicted by `inttype.InvalidIndex Variant = -1` | HIGH | `01-app/05-enum-system.md` §4 Step 1 | `inttype/Variant.go` | Document the `-1` form for signed-int enums |
| D-CVS-14 | Recipe says `<Type>.go` but actual filename is `Variant.go` in 64/71 packages | LOW | §4 Step 3 | every enum package | Document type-name + `Variant.go` convention |
| D-CVS-15 | Recipe shows separate `consts.go`; no enum has one — type + iota + methods all in `<TypeName>.go` | MED | §4 | every enum package | Collapse Step 1 + Step 3 into single-file recipe |
| D-CVS-16 | §6 factory table missing `*AllCases` family (10+1 call sites); `CreateUsingMap` listed but never used | MED | §6 | `enumimpl.New.BasicByte.{DefaultAllCases,DefaultWithAliasMapAllCases,UsingFirstItemSliceAllCases,UsingFirstItemSliceAliasMap,CreateUsingSlicePlusAliasMapOptions,CreateUsingStringersSpread}` | Add `*AllCases` rows; remove unused `CreateUsingMap` |
| D-CVS-17 | §8 says tests live in `tests/integratedtests/<pkg>tests/` — same as C-CVS-01 | MED | §8 | `tests/creationtests/` shared registry | Mirror C-CVS-01 fix from §03 |
| D-CVS-18 | `reflectinternal.TypeName(Invalid)` example unrunnable from `enum-v1` | MED | §4 Step 2 | zero usage | Replace with real pattern |
| D-CVS-19 | "Predicate file-split rule (>6 OR >20 lines)" never enforced (`pathpatterntype` has 113 in one file) | LOW | §4 "Predicate file-split rule" | `pathpatterntype/Variant.go` | Soften to guideline matching practice |

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
| 🚧 Apply MED + LOW spec fixes for §04 (D-CVS-06..13) | ✅ 100.0 verifiable on §04 | 2026-05-04 |
| ✅ Cycle 3 baseline on §05 | **47.1** verifiable on §05 | 2026-05-04 |
| 🚧 Resolve §05 contradictions C-CVS-03..05 (HIGH) | target ≥ 70% on §05 | pending (next task **AD**) |
| 🚧 Apply LOW + MED spec fixes for §05 (D-CVS-14..19) | target 100% on §05 | pending (next task **AD**) |
| 🚧 Fetch `core-v9` source (task **AB**) → resolve 7 ❓ on §04 + 1 ❓ on §05 | — | pending |
| 🚧 Audit all 16 sections of `01-app/` | 16/16 | **3/16 done** |
| 🎯 Reach ≥95% aggregate match rate | ≥ 95 | Pending |
| 🎯 Zero ❌ contradictions | 0 (currently **3** on §05) | ❌ |
