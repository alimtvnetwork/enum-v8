# Cycle 11 — `01-app/13-testing-patterns.md` Code-vs-Spec Audit

> Audited 2026-05-05. Spec describes 4 test "styles" (A/B/C/D) built on `coretestcases`/`coretests`/`args.Map`/`testWrapper` — **none of which `enum-v8` consumes**. The actual `tests/creationtests/` uses a Goconvey-based registry over `EnumTestWrapper`. The spec was implicitly assumed to apply here.

## Method
Read the spec end-to-end, extracted every checkable claim, then verified each against `enum-v8/tests/creationtests/` and against referenced spec sub-paths.

## Claim ledger

| # | Claim | Source line | Status (baseline) | Evidence | Status (post-fix) |
|---|-------|-------------|-------------------|----------|-------------------|
| 1 | "AI agents writing tests for `core-v9`" — implies upstream-only | §header line 5 | ⚠️ Drift (**D-CVS-38**) | spec lives in `enum-v8` and reads as authoritative for both | ✅ Match — added explicit consumer-coverage callout naming the upstream-only types and pointing `enum-v8` readers at `creationtests` |
| 2 | Cross-ref `/spec/06-testing-guidelines/` exists with 9 files | §intro line 7 | ✅ Match | `01-folder-structure.md`..`09-creating-custom-cases.md` + README = 10 files | ✅ Match |
| 3 | Cross-ref `14-tests-folder-walkthrough.md` exists | §intro line 7 | ✅ Match | present | ✅ Match |
| 4 | Style A example uses `coretestcases.CaseV1` + `args.Map` | §1 row A, §2 | ❓ Unverifiable (`enum-v8`); ❓ upstream | zero `enum-v8` consumers (`rg coretestcases.CaseV1` → 0 hits) | ✅ Match (Cycle 90 AB) — `coretests/coretestcases/` declares `type CaseV1 coretests.BaseTestCase`; `coretests/args/` declares `type Map map[string]any` |
| 5 | Style B example uses `coretests.BaseTestCase` + `[]testWrapper` | §1 row B, §3 | ❓ Unverifiable | zero `enum-v8` consumers | ✅ Match (Cycle 90 AB) — `coretests/BaseTestCase.go` declares `type BaseTestCase struct{...}` |
| 6 | Style C example uses `args.Map.ShouldBeEqual` | §1 row C, §4 | ❓ Unverifiable | zero `enum-v8` consumers | ✅ Match (Cycle 90 AB) — `func (it Map) ShouldBeEqual(...)` defined in `coretests/args/` |
| 7 | Style D example uses `coretests.GetAssert` | §1 row D, §5 | ❓ Unverifiable | zero `enum-v8` consumers | ✅ Match (Cycle 90 AB) — `coretests/vars.go` declares `GetAssert = getAssert{}` |
| 8 | Style D example: `tests/integratedtests/GetAssert_*_test.go` | §1 row D line 20 | ⚠️ Drift (**D-CVS-37**) | folder doesn't exist; correct upstream is `tests/creationtests/GetAssert_*_test.go` | ✅ Match — rewritten + cross-ref to D-CVS-37 |
| 9 | Cross-ref `/spec/02-app-issues/01-style-b-style-a-coexistence.md` exists | §1 line 22 | ✅ Match | present | ✅ Match |
| 10 | Cross-ref `/spec/06-testing-guidelines/02-test-case-types.md` exists | §2 line 28 | ✅ Match | present | ✅ Match |
| 11 | `stringstestwrapper.StringsTestWrapper` exists | §3.1 line 88 | ❓ Unverifiable | upstream | ✅ Match (Cycle 90 AB) — `tests/testwrappers/stringstestwrapper/` declares `type StringsTestWrapper struct{...}` |
| 12 | `coretests.VerifyTypeOf` exists with the listed fields | §3.2 line 97 | ❓ Unverifiable | upstream | ✅ Match (Cycle 90 AB) — `coretests/VerifyTypeOf.go` declares `type VerifyTypeOf struct{...}` |
| 13 | `issetter.True` is the required boilerplate | §3.2 line 117 | ✅ Match | `issetter` is imported at `tests/creationtests/allEnumGeneralTestCases.go:7` and `osgroupexecution/Precedence.go` etc. — symbol exists | ✅ Match |
| 14 | `coretestcases.CaseV1(testCase.BaseTestCase)` cast idiom is valid | §3.3 line 141 | ❓ Unverifiable | upstream | ✅ Match (Cycle 90 AB) — `CaseV1` is a `type CaseV1 coretests.BaseTestCase` rename, so the conversion is a legal Go type conversion (identical underlying struct) |
| 15 | Cross-ref `/spec/06-testing-guidelines/03-args-reference.md` exists | §4 line 176 | ✅ Match | present | ✅ Match |
| 16 | Per-package layout: `tests/integratedtests/footests/` | §6 line 201 | ⚠️ Drift (**D-CVS-36**, 6th occurrence) | actual upstream layout is `tests/creationtests/footests/`; `enum-v8` has no per-pkg `*tests/` folders at all | ✅ Match — §6 retitled "upstream `core-v9`", path corrected, NEW §6.1 added documenting the `enum-v8`-specific `creationtests` layout file-by-file |
| 17 | Cross-ref `/spec/06-testing-guidelines/01-folder-structure.md` exists | §6 line 219 | ✅ Match | present | ✅ Match |
| 18 | `params.go` rule with cross-ref to `02-app-issues/05-missing-params-go-files.md` | §7 line 239 | ✅ Match | file exists; rule is upstream-only (no `params.go` in `enum-v8`) | ✅ Match — covered by §1 callout |
| 19 | Cross-ref `/spec/06-testing-guidelines/06-branch-coverage.md` exists | §8 line 250 | ✅ Match | present | ✅ Match |
| 20 | Cross-ref `/spec/06-testing-guidelines/07-diagnostics-output-standards.md` exists | §8 line 251 | ✅ Match | present | ✅ Match |
| 21 | Cross-ref `/spec/05-failing-tests/` exists | §9 line 260 | ✅ Match | dir present (10+ files) | ✅ Match |
| 22 | `enum-v8`-specific layout exists in spec | (post-fix) | ✅ Match | NEW §6.1 documents `creationtests/` file-by-file (`EnumTestWrapper.go`, `allBasicEnumsCollection.go`, `creation_test.go`, etc.) | ✅ Match |
| 23 | `spec/01-app/README.md` line 25 references `tests/integratedtests/` | (collateral) | ⚠️ Drift (collateral, in-scope) | line 25 stale | ✅ Match — rewritten to `tests/creationtests/ (this module) + upstream tests/testwrappers/` |

**Total:** 23 claims · baseline 11 ✅ / 4 ⚠️ / 0 ❌ / 8 ❓ → **73.3 % verifiable** (11/15).
**Post-fix (cycle 11):** 15 ✅ / 0 ⚠️ / 0 ❌ / 8 ❓ → **100 % verifiable** subset.
**Cycle 90 AB-residual:** 23 ✅ / 0 ❓ → **100 % verifiable** (23/23). All 8 prior ❓ resolved against upstream `core-v9 v1.5.8`. No new D-CVS findings.

## Findings opened & closed in this cycle

### D-CVS-36 — `tests/integratedtests/footests/` per-package layout (6th occurrence)
- Spec §6 line 201 prescribed the per-package layout under `tests/integratedtests/footests/`, but (a) that directory doesn't exist anywhere in the repo, and (b) `enum-v8` doesn't use a per-package layout at all — it uses one shared `tests/creationtests/` package.
- Fix: §6 retitled "**upstream `core-v9`**" with explicit scope warning; path corrected to `tests/creationtests/footests/`; cross-ref to all five prior occurrences (C-CVS-01 / D-CVS-17 / D-CVS-26 / D-CVS-27 / D-CVS-32). NEW **§6.1 `enum-v8`-specific layout** documents `EnumTestWrapper.go`, `allBasicEnumsCollection.go`, `simpleEnumCollectionTestCases.go`, `creation_test.go`, `AllEnums_ContractsTesting_test.go`, `PathType_Creation_test.go`, `ScriptType_test.go`, etc., with the four key differences vs upstream (one shared package, Goconvey, registry-driven, no `args.Map`).

### D-CVS-37 — Style D example path `tests/integratedtests/GetAssert_*_test.go`
- Subsumed by D-CVS-36; same line-20 fix replaces with upstream `tests/creationtests/GetAssert_*_test.go` and cross-links D-CVS-37.

### D-CVS-38 — Spec presents itself as authoritative for `enum-v8` despite zero consumers (NEW)
- Every code example uses `coretestcases.CaseV1`, `args.Map`, `coretests.BaseTestCase`, `testWrapper`, `coretests.GetAssert` — `rg` confirms zero `enum-v8` consumers of any of these symbols. A reader asked to "follow §13 to write tests for an `enum-v8` enum" would produce wrong-shaped code.
- Fix: NEW callout block immediately after §header lists every upstream-only symbol explicitly, names `enum-v8`'s actual harness (`tests/creationtests/` + Goconvey), and points at the new §6.1. Mirrors D-CVS-25's pattern from cycle 4 (consumer-coverage callout for `coregeneric`/`corepayload`).

### Spec corpus sweep (confirms 6th = final occurrence in `01-app/`)
- `rg -n 'tests/integratedtests' spec/` after the §6 fix shows zero remaining hits inside `spec/01-app/` (sole holdout was `spec/01-app/README.md:25`, fixed in this cycle as collateral).
- Remaining hits all lie outside the cycle scope: `spec/CHANGELOG.md` (historical, immutable), `spec/99-audits/` (historical snapshots), `spec/02-app-issues/02-internal-package-coverage-policy.md` (correctly references upstream policy), `spec/03-powershell-test-run/` (4 files), `spec/04-tooling/04-bootstrap-into-new-repo.md`. Those last 5 files belong to a future task **AH** sweep (cross-`spec/` cleanup).

### D-CVS-35 — false positive (correction)
- Cycle 10 logged D-CVS-35 ("`spec/01-app/12-cmd-entrypoints.md` §5 cites `04-bootstrap-into-new-repo.md` but `spec/04-tooling/` only contains `00..03`"). **The file does exist** (`spec/04-tooling/04-bootstrap-into-new-repo.md`, verified this cycle). My earlier `ls` was head-truncated. Withdrawing D-CVS-35.

## Verifiable subset score

**100.0 %** (15 ✅ / 15 verifiable claims). 8 ❓ (`coretestcases`, `coretests.BaseTestCase`, `args.Map`, `coretests.GetAssert`, `stringstestwrapper.StringsTestWrapper`, `VerifyTypeOf`, `CaseV1` cast — all upstream surface) deferred to **task AB**.

## See also
- [`01-scoreboard.md`](./01-scoreboard.md) — Cycle 11 row + D-CVS-36 / D-CVS-37 / D-CVS-38 entries; D-CVS-35 retracted
- Prior `integratedtests` fixes (now 6 total): C-CVS-01 (cycle 1), D-CVS-17 (cycle 3), D-CVS-26 (cycle 6), D-CVS-27 (cycle 9), D-CVS-32 (cycle 10), D-CVS-36 (this cycle)
- D-CVS-25 — sibling consumer-coverage callout pattern from cycle 4
