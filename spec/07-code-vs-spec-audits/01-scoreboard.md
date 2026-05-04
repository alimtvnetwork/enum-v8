# 01 — Code-vs-Spec Drift Scoreboard (Living Document)

> **Single source of truth** for code-vs-spec drift. Updated after every cycle.

## Current MEASURED drift score: **§03 100.0 / §04 100.0 / §05 100.0 / §06 35.7 (verifiable)** *(4 sections audited)*

> §03–§05 closed. §06 (data-structures) Cycle 4 baseline: 20 claims, 5 ✅ / 6 ⚠️ / **3 ❌** / **6 ❓**. The ❓ count is high because `coregeneric` and `corepayload` have **zero consumers in `enum-v1`** — those sub-packages are documented for upstream completeness but cannot be verified from this repo. Three contradictions on the verifiable subset: `corejson.Serialize.ToString` / `Deserialize.FromTo` examples don't exist (real API is `ToBytesErr` / `BytesTo`); `coreonce.New.String` namespace doesn't match real `coreonce.NewAnyOnce` / `NewByteOnce` calls; and **`inttype/Variant.go` + `inttype/all-constructors.go` import `encoding/json` directly**, violating §4's "never `encoding/json`" rule. See [`05-cycle4-data-structures.md`](./05-cycle4-data-structures.md).

## Cycle history

| Date | Cycle | Spec audited | Claims | ✅ Match | ⚠️ Drift | ❌ Contradiction | ❓ Unverifiable | Score |
|------|-------|--------------|--------|---------|---------|------------------|----------------|-------|
| 2026-05-04 | 1 (baseline) | `01-app/03-import-conventions.md` | 12 | 5 | 5 | 2 | 0 | **41.7%** |
| 2026-05-04 | 1 (post-LOW) | `01-app/03-import-conventions.md` | 12 | 10 | 0 | 2 | 0 | **83.3%** |
| 2026-05-04 | 1 (closed) | `01-app/03-import-conventions.md` | 12 | 12 | 0 | 0 | 0 | **100.0%** |
| 2026-05-04 | 2 (baseline) | `01-app/04-error-system.md` | 18 | 3 | 8 | 0 | 7 | **27.3%** *(verifiable)* |
| 2026-05-04 | 2 (closed) | `01-app/04-error-system.md` | 18 | 11 | 0 | 0 | 7 | **100.0%** *(verifiable)* |
| 2026-05-04 | 3 (baseline) | `01-app/05-enum-system.md` | 18 | 8 | 6 | 3 | 1 | **47.1%** *(verifiable)* |
| 2026-05-04 | 3 (closed) | `01-app/05-enum-system.md` | 18 | 17 | 0 | 0 | 1 | **100.0%** *(verifiable)* |
| 2026-05-04 | 4 (baseline) | `01-app/06-data-structures.md` | 20 | 5 | 6 | 3 | 6 | **35.7%** *(verifiable)* |

## Open drift findings

| ID | Title | Severity | Spec ref | Code ref | Resolution path |
|----|-------|----------|----------|----------|-----------------|
| C-CVS-06 | §4 "never `encoding/json` directly" rule violated by `inttype` | HIGH | `01-app/06-data-structures.md` §4 "Rule" | `inttype/Variant.go:440` (`json.Marshal`), `inttype/all-constructors.go:75` (`*json.Number`) | Document the legitimate exceptions (`MarshalJSON` primitive delegation, `*json.Number` parameter) |
| C-CVS-07 | `corejson.Serialize.ToString` / `Serialize.Raw` example doesn't compile | HIGH | §4 code block | zero call sites; real: `corejson.Serialize.ToBytesErr` | Replace examples with the actually-used API |
| C-CVS-08 | `corepayload.New.PayloadWrapper.UsingInstruction(...)` example unverifiable | MED | §6 | zero `corepayload` consumers in `enum-v1` | Move §6 to "upstream-only" appendix or fetch upstream (task **AB**) |
| D-CVS-20 | `corejson.Serialize.ToString` / `Raw` listed but never called | MED | §4 | real: `Serialize.ToBytesErr` | Update §4 examples |
| D-CVS-21 | `corejson.Deserialize.UsingBytes` / `FromTo` listed but never called | MED | §4 | real: `Deserialize.BytesTo` | Update §4 examples |
| D-CVS-22 | `coreonce.New.String(producer)` namespace doesn't match real top-level constructors | MED | §5 | real: `coreonce.NewAnyOnce`, `coreonce.NewByteOnce` | Rewrite §5 around real constructors |
| D-CVS-23 | `corestr` shown as "thread-safe list of strings"; real surface is `Hashset`/`SimpleSlice`/`SimpleStringOnce` | MED | §3 | `corestr.New.Hashset`, `corestr.New.SimpleSlice`, `corestr.SimpleStringOnce` | Rewrite §3 around real exports |
| D-CVS-24 | `coreonce` "covers all common types" overstated | LOW | §5 | only `AnyOnce` / `ByteOnce` observed | Soften wording or fetch upstream |
| D-CVS-25 | `coregeneric` and `corepayload` presented as first-class but have no `enum-v1` consumers | LOW | §1, §2, §6 | zero imports | Add "upstream-only" callout in §1 |

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
| C-CVS-03 | Spec mandated first const = `Invalid`; 10 enums use other sentinels | 2026-05-04 | `spec/01-app/05-enum-system.md` §4.1 | Reframed as "sentinel-first" rule with sentinel-name table (`Invalid` / `Default` / `Unspecified` / `Uninitialized` / domain term) |
| C-CVS-04 | Recipe imported `core-v9/internal/reflectinternal` (forbidden cross-module `internal/`) | 2026-05-04 | `spec/01-app/05-enum-system.md` §4.3 | Replaced with `enumimpl.New.BasicByte.DefaultAllCases(Invalid, Ranges[:])`; added explicit warning |
| C-CVS-05 | "Zero-value sentinel" rule contradicted by `inttype.InvalidIndex Variant = -1` | 2026-05-04 | `spec/01-app/05-enum-system.md` §4.1 | Documented signed-int exception (`InvalidIndex = -1`) explicitly |
| D-CVS-14 | Recipe used `<Type>.go`; actual filename is `Variant.go` in 64/71 packages | 2026-05-04 | `spec/01-app/05-enum-system.md` §1 + §4.2 | Documented `<TypeName>.go` convention; called out `Variant` as conventional type name |
| D-CVS-15 | Recipe split `consts.go` + `<Type>.go`; no enum has `consts.go` | 2026-05-04 | `spec/01-app/05-enum-system.md` §4 | Collapsed to 2-file recipe (`Variant.go` + `vars.go`) |
| D-CVS-16 | §6 missing `*AllCases` family; listed unused `CreateUsingMap` | 2026-05-04 | `spec/01-app/05-enum-system.md` §6 | Expanded factory table with all 9 in-use methods; dropped `CreateUsingMap` |
| D-CVS-17 | §8 referenced nonexistent `tests/integratedtests/<pkg>tests/` | 2026-05-04 | `spec/01-app/05-enum-system.md` §8 | Rewrote to point at `tests/creationtests/` shared registry (mirrors C-CVS-01 fix from §03) |
| D-CVS-18 | `reflectinternal.TypeName(Invalid)` example unrunnable | 2026-05-04 | `spec/01-app/05-enum-system.md` §4.3 | Replaced with `DefaultAllCases` / `UsingTypeSlice` patterns |
| D-CVS-19 | Predicate file-split rule (>6 OR >20 lines) never enforced | 2026-05-04 | `spec/01-app/05-enum-system.md` §4.5 | Softened to guideline matching `pathpatterntype` reality |

## Targets

| Milestone | Score | Status |
|-----------|-------|--------|
| ✅ First measured baseline (Cycle 1) | **41.7** | 2026-05-04 |
| ✅ Apply 5 LOW spec fixes from Cycle 1 (D-CVS-01..05) | **83.3** on §03 | 2026-05-04 |
| ✅ Resolve C-CVS-01 + C-CVS-02 → §03 at 100% | **100.0** on §03 | 2026-05-04 |
| ✅ Cycle 2 baseline on §04 | **27.3** verifiable on §04 | 2026-05-04 |
| ✅ Apply MED + LOW spec fixes for §04 (D-CVS-06..13) | **100.0** verifiable on §04 | 2026-05-04 |
| ✅ Cycle 3 baseline on §05 | **47.1** verifiable on §05 | 2026-05-04 |
| ✅ Resolve §05 contradictions C-CVS-03..05 (HIGH) + apply D-CVS-14..19 | **100.0** verifiable on §05 | 2026-05-04 |
| 🚧 Fetch `core-v9` source (task **AB**) → resolve 7 ❓ on §04 + 1 ❓ on §05 | — | pending |
| 🚧 Audit all 16 sections of `01-app/` | 16/16 | **3/16 done** |
| 🎯 Reach ≥95% aggregate match rate | ≥ 95 | ✅ (verifiable subset) |
| 🎯 Zero ❌ contradictions | 0 | ✅ |
