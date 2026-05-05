# 01 — Code-vs-Spec Drift Scoreboard (Living Document)

> **Single source of truth** for code-vs-spec drift. Updated after every cycle.

## Current MEASURED drift score: **§03 100.0 / §04 100.0 / §05 100.0 / §06 100.0 / §08 100.0 / §10 100.0 / §11 100.0 / §12 100.0 (verifiable)** *(8 sections audited + 2 baseline-only — §07, §09 — all closed)*

> §12 (cmd-entrypoints) Cycle 10 closed by resolving **1 HIGH contradiction + 1 drift** (+1 latent drift logged for later): (a) **C-CVS-10** — spec §1 asserted "no `cmd/` directory" / "no `main` package" but `enum-v2/cmd/main/main.go` exists with `package main` and `func main()`; rewrote §1 as a "library-first, smoke-test allowed" policy that distinguishes upstream `core-v9` (zero `cmd/`) from this module (one permitted smoke-test harness `cmd/main/`); rule narrowed to "no additional `cmd/<name>/` entrypoints"; cross-linked `cmd/README.md`. (b) **D-CVS-32** — §3 line 71 stale `tests/integratedtests/` rewritten to `tests/creationtests/` (5th occurrence; final `rg integratedtests` sweep deferred to task **AH**). (c) **D-CVS-35** (NEW, deferred) — §5 cites `04-bootstrap-into-new-repo.md` but `spec/04-tooling/` only contains `00..03`. See [`11-cycle10-cmd-entrypoints.md`](./11-cycle10-cmd-entrypoints.md).
>
> §11 (versioning) Cycle 9 closed by resolving **2 contradictions + 3 drifts**: **C-CVS-09a/b** (mojibake `core-v9 → core-v9` at §3:95 and §4:112 rewritten as the historical `core-v8` → `core-v9` migration), **D-CVS-27** (`tests/integratedtests/` → `tests/creationtests/` at §4:108), **D-CVS-30** (`versionindexes.V8` comment corrected: V8 = legacy era, V9 = current), **D-CVS-31** (4 stale `.lovable/user-preferences` citations rewritten to `mem://index.md` Core). See [`10-cycle9-versioning.md`](./10-cycle9-versioning.md).

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
| 2026-05-04 | 4 (closed) | `01-app/06-data-structures.md` | 20 | 14 | 0 | 0 | 6 | **100.0%** *(verifiable)* |
| 2026-05-04 | 5 (baseline) | `01-app/07-conditional-and-utilities.md` | 17 | 0 | 0 | 0 | 17 | **N/A** *(no verifiable subset)* |
| 2026-05-05 | 6 (baseline) | `01-app/08-validators.md` | 19 | 0 | 1 | 0 | 18 | **0.0%** *(verifiable)* |
| 2026-05-05 | 6 (closed)   | `01-app/08-validators.md` | 19 | 1 | 0 | 0 | 18 | **100.0%** *(verifiable)* |
| 2026-05-05 | 7 (baseline) | `01-app/09-converters.md` | 23 | 0 | 0 | 0 | 23 | **N/A** *(no verifiable subset)* |
| 2026-05-05 | 8 (baseline / closed) | `01-app/10-reflection-and-dynamic.md` | 19 | 4 | 0 | 0 | 15 | **100.0%** *(verifiable)* |
| 2026-05-05 | 9 (baseline) | `01-app/11-versioning.md` | 20 | 4 | 3 | 2 | 11 | **44.4%** *(verifiable)* |
| 2026-05-05 | 9 (closed)   | `01-app/11-versioning.md` | 20 | 9 | 0 | 0 | 11 | **100.0%** *(verifiable)* |
| 2026-05-05 | 10 (baseline) | `01-app/12-cmd-entrypoints.md` | 22 | 9 | 3 | 4 | 6 | **56.3%** *(verifiable)* |
| 2026-05-05 | 10 (closed)   | `01-app/12-cmd-entrypoints.md` | 22 | 16 | 0 | 0 | 6 | **100.0%** *(verifiable)* |

## Open drift findings

**D-CVS-35** (low, deferred) — `spec/01-app/12-cmd-entrypoints.md` §5 cites `/spec/04-tooling/04-bootstrap-into-new-repo.md` but `spec/04-tooling/` only contains `00..03`. Logged in cycle 10 but not in scope for fix; address in a future cleanup pass.

All 8 audited-and-closed sections (§03, §04, §05, §06, §08, §10, §11, §12) are at 100 % of their verifiable subsets. §07 and §09 have no verifiable subset (zero `enum-v2` consumers). Remaining ❓s — 17 §07 + 18 §08 + 23 §09 + 15 §10 + 11 §11 + 6 §12 + 7 §04 + 1 §05 + 6 §06 = **104 ❓** total — require upstream `core-v9` source (task **AB**).

## Resolved drift findings

| ID | Title | Resolved at | Fix location | Path taken |
|----|-------|-------------|--------------|------------|
| D-CVS-01 | Spec §03 line 4 says "consumes `core-v9`" — stale | 2026-05-04 | `spec/01-app/03-import-conventions.md:4` | s/core-v9/core-v9/ |
| D-CVS-02 | Spec §03 line 88 says path "ends in `core-v9`" — stale | 2026-05-04 | `spec/01-app/03-import-conventions.md:88` | s/core-v9/core-v9/ + s/corev8/corev9/ |
| D-CVS-03 | Spec §03 line 94 prose/example mismatch (v8 vs v9) | 2026-05-04 | `spec/01-app/03-import-conventions.md:94` | s/core-v9/core-v9/ |
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
| C-CVS-06 | §4 "never `encoding/json` directly" rule violated by `inttype` | 2026-05-04 | `spec/01-app/06-data-structures.md` §4 "Rule (with documented exceptions)" | Documented the two legitimate exceptions: `MarshalJSON` → `json.Marshal` for primitive emission; `*json.Number` parameter type |
| C-CVS-07 | `corejson.Serialize.ToString` / `Serialize.Raw` example didn't compile | 2026-05-04 | `spec/01-app/06-data-structures.md` §4 code block | Replaced with the actually-used `Serialize.ToBytesErr` / `Deserialize.BytesTo` + `*Result` wrapper |
| C-CVS-08 | `corepayload.New.PayloadWrapper.UsingInstruction(...)` example unverifiable | 2026-05-04 | `spec/01-app/06-data-structures.md` §6 | Added explicit "upstream-only" callout deferring field-set verification to task **AB** |
| D-CVS-20 | `corejson.Serialize.ToString` / `Raw` listed but never called | 2026-05-04 | `spec/01-app/06-data-structures.md` §4 | Replaced with `Serialize.ToBytesErr(...) → *Result` |
| D-CVS-21 | `corejson.Deserialize.UsingBytes` / `FromTo` listed but never called | 2026-05-04 | `spec/01-app/06-data-structures.md` §4 | Replaced with `Deserialize.BytesTo(bytes, &target)` |
| D-CVS-22 | `coreonce.New.String(producer)` namespace doesn't match real top-level constructors | 2026-05-04 | `spec/01-app/06-data-structures.md` §5 | Rewrote §5 around `coreonce.NewAnyOnce` / `NewByteOnce` |
| D-CVS-23 | `corestr` shown as "thread-safe list of strings"; real surface is `Hashset`/`SimpleSlice`/`SimpleStringOnce` | 2026-05-04 | `spec/01-app/06-data-structures.md` §3 | Rewrote §3 around `New.Hashset` / `New.SimpleSlice` / `SimpleStringOnce` |
| D-CVS-24 | `coreonce` "covers all common types" overstated | 2026-05-04 | `spec/01-app/06-data-structures.md` §5 | Softened to "common typed wrappers"; cross-referenced `corestr.SimpleStringOnce` |
| D-CVS-25 | `coregeneric` and `corepayload` presented as first-class but have no `enum-v2` consumers | 2026-05-04 | `spec/01-app/06-data-structures.md` §1, §2, §6 + §7 decision matrix | Added "Consumer-coverage note" in §1 + ⚠️ "upstream-only" callouts in §2 and §6; §7 matrix now marks each row with `enum-v2` verification status |
| D-CVS-26 | §08 §6 references nonexistent `tests/integratedtests/<pkg>tests/` for validator tests | 2026-05-05 | `spec/01-app/08-validators.md` §6 line 347 | Rewrote to `tests/creationtests/<pkg>tests/` + cross-ref to C-CVS-01 / D-CVS-17 (mirrors the §03 / §05 fixes) |
| C-CVS-09a | §11 §3 line 95 says "`core-v9` → `core-v9`" (mojibake from bulk v8→v9 rename) | 2026-05-05 | `spec/01-app/11-versioning.md:95` | Rewrote to "the historical `core-v8` → `core-v9` migration" — legitimate historical reference |
| C-CVS-09b | §11 §4 line 112 says "module path changes (`core-v9` → `core-v9`)" — same mojibake | 2026-05-05 | `spec/01-app/11-versioning.md:112` | Rewrote to "the historical `core-v8` → `core-v9` migration is the canonical example" |
| D-CVS-27 | §11 §4 line 108 references `tests/integratedtests/` (4th occurrence of this pattern) | 2026-05-05 | `spec/01-app/11-versioning.md:108` | Rewrote to `tests/creationtests/` + cross-ref to C-CVS-01 / D-CVS-17 / D-CVS-26 |
| D-CVS-30 | §11 §2 line 59 comment claims `versionindexes.V8 // 8 (current era — core-v9)` — contradictory | 2026-05-05 | `spec/01-app/11-versioning.md:59` | Rewrote to `// 8 (legacy era; the current core-v9 era is V9)` |
| D-CVS-31 | §11 cites `.lovable/user-preferences line 8` (file does not exist in `enum-v2`) in 4 places | 2026-05-05 | `spec/01-app/11-versioning.md` lines 5, 78, 133, 156 | Rewrote all 4 citations to point only to `mem://index.md` Core (which exists and carries the rule) |

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
| ✅ Cycle 4 baseline on §06 | **35.7** verifiable on §06 | 2026-05-04 |
| ✅ Resolve §06 contradictions C-CVS-06..08 + apply D-CVS-20..25 | **100.0** verifiable on §06 | 2026-05-04 |
| ✅ Cycle 6 baseline on §08 + apply D-CVS-26 | **100.0** verifiable on §08 | 2026-05-05 |
| ✅ Cycle 7 baseline on §09 (no drifts; all upstream-only) | **N/A** *(no verifiable subset)* | 2026-05-05 |
| ✅ Cycle 8 baseline+closed on §10 (4 MUST/MUST-NOT rules verified, 0 violations) | **100.0** verifiable on §10 | 2026-05-05 |
| ✅ Cycle 9 baseline on §11 | **44.4** verifiable on §11 | 2026-05-05 |
| ✅ Resolve §11 contradictions C-CVS-09a/b + apply D-CVS-27, D-CVS-30, D-CVS-31 | **100.0** verifiable on §11 | 2026-05-05 |
| 🚧 Fetch `core-v9` source (task **AB**) → resolve **98 ❓** total: 17 §07 + 18 §08 + 23 §09 + 15 §10 + 11 §11 + 7 §04 + 1 §05 + 6 §06 | — | pending |
| 🚧 Audit all 16 sections of `01-app/` | 16/16 | **9/16 baseline (7 closed, 2 baseline-only)** |
| 🎯 Reach ≥95 % aggregate match rate | ≥ 95 | ✅ (verifiable subset) |
| 🎯 Zero ❌ contradictions | 0 | ✅ |
