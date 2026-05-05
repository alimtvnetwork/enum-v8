# Cycle 12 — `01-app/14-tests-folder-walkthrough.md` Code-vs-Spec Audit

> Audited 2026-05-05. Spec catalogues `tests/integratedtests/` layout, `tests/testwrappers/` packages, and `coretests.GetAssert` helper inventory — **none of which `enum-v2` consumes**. Mirrors the §13 finding from cycle 11. Also exposed that cycle 11's "01-app/ is clean" claim was premature — 7 more `integratedtests` hits remained across `01-package-map.md`, `02-design-philosophy.md`, and `14-…` itself.

## Method
Read end-to-end, extracted every checkable claim, then verified each against `enum-v2/tests/creationtests/` and against referenced spec sub-paths.

## Claim ledger

| # | Claim | Source line | Status (baseline) | Evidence | Status (post-fix) |
|---|-------|-------------|-------------------|----------|-------------------|
| 1 | Audience: helpers in `tests/testwrappers/` | header | ⚠️ Drift (**D-CVS-42**) | spec lives in `enum-v2`; `rg tests/testwrappers` over `enum-v2/` source = 0 hits | ✅ Match — added consumer-coverage callout |
| 2 | Cross-ref `13-testing-patterns.md` exists | line 7 | ✅ Match | present | ✅ Match |
| 3 | "`tests/integratedtests/` — One Folder per Source Package" | §1 line 11 | ⚠️ Drift (**D-CVS-39**, 7th occurrence) | folder doesn't exist; correct upstream path is `tests/creationtests/` | ✅ Match — §1 retitled "**upstream `core-v9`**", path corrected, scope warning added |
| 4 | 50+ subfolders listed (`argstests/`, `anycmptests/`, …) | §1 lines 14-31 | ❓ Unverifiable | upstream | ❓ Unverifiable (→ AB) |
| 5 | `GetAssert_*_test.go` at top of `integratedtests/` | §1 line 31 | ⚠️ Drift (subsumed by D-CVS-39) | path → `tests/creationtests/` | ✅ Match |
| 6 | Naming rule: folder = source-pkg + `tests`; pkg name = same as folder; `_test.go` runners only; `_testcases.go` no `import "testing"`; one `NilReceiver_test.go` per package | §1 lines 36-40 | ❓ Unverifiable (`enum-v2` has no per-pkg dirs); applies to upstream | ❓ Unverifiable (→ AB) |
| 7 | Cross-ref `/spec/06-testing-guidelines/01-folder-structure.md` exists | §1 line 42 | ✅ Match | present | ✅ Match |
| 8 | Cross-ref `/spec/02-app-issues/04-testwrappers-public-surface.md` exists | §2 line 50 | ✅ Match | present | ✅ Match |
| 9 | `stringstestwrapper.StringsTestWrapper` API: `Arrange()/Expected() []string` | §2.1 lines 60-67 | ❓ Unverifiable | upstream | ❓ Unverifiable (→ AB) |
| 10 | `chmodhelpertestwrappers` files (RwxCompile, RwxInstruction, VerifyRwx*) | §2.2 | ❓ Unverifiable | upstream | ❓ Unverifiable (→ AB) |
| 11 | `coredynamictestwrappers.ReflectSetFromTo*` files | §2.3 | ❓ Unverifiable | upstream | ❓ Unverifiable (→ AB) |
| 12 | `corevalidatortestwrappers.{Segment,Slice,Text}ValidatorWrapper.go` | §2.4 | ❓ Unverifiable | upstream | ❓ Unverifiable (→ AB) |
| 13 | Cross-ref `/spec/02-app-issues/03-getassert-undocumented-api.md` exists | §3 line 112 | ✅ Match | present | ✅ Match |
| 14 | "observed from `tests/integratedtests/GetAssert_*_test.go`" | §3 line 112 | ⚠️ Drift (**D-CVS-41**) | wrong upstream path | ✅ Match — corrected to `tests/creationtests/GetAssert_*_test.go` |
| 15 | `GetAssert.Quick / SortedArray / SortedArrayNoPrint / SortedMessage / ToString / ToStrings / ToStringsWithSpace / AnyToDoubleQuoteLines / AnyToStringDoubleQuoteLine / ConvertLinesToDoubleQuoteThenString / ErrorToLinesWithSpaces / ErrorToLinesWithSpacesDefault / SimpleTestCaseWrapper` (13 methods) | §3.1 lines 117-130 | ❓ Unverifiable | upstream | ❓ Unverifiable (→ AB) |
| 16 | Canonical pattern: `formatter := asserter.X` then `actualSlice.Adds(...)` | §3.2 | ❓ Unverifiable | upstream | ❓ Unverifiable (→ AB) |
| 17 | "Add to `GetAssert` when 3+ pkgs use it AND deterministic AND pure formatting" | §3.3 | ✅ Match (process rule, in-spec) | self-consistent governance rule | ✅ Match |
| 18 | `coretestcases.CaseV1(testCase.BaseTestCase)` cast safe (same memory layout) | §4 lines 159-166 | ❓ Unverifiable | upstream | ❓ Unverifiable (→ AB) |
| 19 | Walkthrough example uses `tests/integratedtests/widgettests/` | §5 line 175 | ⚠️ Drift (**D-CVS-40**, 8th occurrence) | wrong upstream path | ✅ Match — corrected to `tests/creationtests/widgettests/` + `enum-v2` redirect note |
| 20 | Cross-ref `13-testing-patterns.md §3` (Style B) exists | §5 line 191 | ✅ Match | present | ✅ Match |
| 21 | Cross-ref `/spec/05-failing-tests/` exists with 25 post-mortems | §6 line 197 | ✅ Match | dir present, 26 files (`01..26-…`) | ✅ Match |
| 22 | 5 referenced failing-tests files exist (`02-…`, `12-…`, `13-…`, `18-…`, `22-…`) | §6 lines 201-205 | ✅ Match | all 5 present | ✅ Match |
| 23 | Collateral: `01-app/01-package-map.md` §8 + 4 bullets reference `tests/integratedtests/` (5 hits) | (collateral) | ⚠️ Drift (in-scope) | stale | ✅ Match — §8 retitled "upstream"; scope warning added; 4 bullets corrected to `tests/creationtests/` |
| 24 | Collateral: `01-app/02-design-philosophy.md` line 183 references `tests/integratedtests/footests/` | (collateral) | ⚠️ Drift (in-scope) | stale | ✅ Match — corrected + `enum-v2` redirect note |

**Total:** 24 claims · baseline 8 ✅ / 6 ⚠️ / 0 ❌ / 10 ❓ → **57.1 % verifiable** (8/14).
**Post-fix:** 14 ✅ / 0 ⚠️ / 0 ❌ / 10 ❓ → **100 % verifiable**.

## Findings opened & closed in this cycle

### D-CVS-39 — `tests/integratedtests/` per-package layout in §14 (7th occurrence)
- §1 prescribed `tests/integratedtests/` as the canonical layout with 16+ subfolder examples — folder doesn't exist; correct upstream path is `tests/creationtests/`.
- Fix: §1 retitled "**`tests/creationtests/` *(upstream)* — One Folder per Source Package**"; scope warning added pointing `enum-v2` readers at §13's §6.1; tree diagram updated; cross-link to all six prior occurrences (C-CVS-01 / D-CVS-17 / D-CVS-26 / D-CVS-27 / D-CVS-32 / D-CVS-36).

### D-CVS-40 — `tests/integratedtests/widgettests/` walkthrough example (8th occurrence)
- §5 line 175 worked-example folder path corrected to `tests/creationtests/widgettests/` with an inline `enum-v2`-specific redirect ("register the enum in `tests/creationtests/allBasicEnumsCollection.go` instead").

### D-CVS-41 — GetAssert observation source path (9th occurrence)
- §3 stability note rewritten to cite upstream `tests/creationtests/GetAssert_*_test.go`.

### D-CVS-42 — Spec lacks consumer-coverage callout (NEW, mirrors D-CVS-38)
- Same pattern as cycle 11's D-CVS-38: every wrapper, helper, and layout described is upstream-only for `enum-v2`. Added explicit callout naming `tests/testwrappers/`, `coretests.GetAssert`, `coretestcases.CaseV1`, `StringsTestWrapper` and pointing `enum-v2` readers at §13 §6.1.

### Collateral fixes — `01-package-map.md` §8 + `02-design-philosophy.md` line 183
- Cycle 11 declared `spec/01-app/` clean of `integratedtests` references after fixing D-CVS-36, but a re-sweep this cycle found 7 more hits (5 in `01-package-map.md` §8, 1 in `02-design-philosophy.md` line 183, 1 latent in `14-…` outside §1). All fixed in this cycle as in-scope collateral. Sweep now genuinely clean: `rg -n 'tests/integratedtests' spec/01-app/` returns only **intentional anti-pattern callouts** in `05-enum-system.md` line 417 (a "do NOT do this" entry that must keep the wrong path) and the new D-CVS-39/D-CVS-42 references in `13-…` and `14-…` themselves.

### Cycle 11 correction (no new finding)
- Cycle 11's scoreboard claim "Cross-spec sweep status: `rg -n 'tests/integratedtests' spec/01-app/` is now clean" was incomplete — `01-package-map.md` and `02-design-philosophy.md` were missed. Corrected in this cycle. The next AH-style sweep should look at `02-design-philosophy.md`, `01-package-map.md`, and any later `01-app/` files BEFORE auditing them, not after.

## Verifiable subset score

**100.0 %** (14 ✅ / 14 verifiable claims). 10 ❓ (`coretests`, `GetAssert.*` 13 methods, 4 testwrapper packages, `coretestcases.CaseV1` cast — all upstream surface) deferred to **task AB**.

## See also
- [`01-scoreboard.md`](./01-scoreboard.md) — Cycle 12 row + D-CVS-39..42 entries
- Prior `integratedtests` fixes (now 9 spec-corpus occurrences resolved): C-CVS-01 (cycle 1), D-CVS-17 (cycle 3), D-CVS-26 (cycle 6), D-CVS-27 (cycle 9), D-CVS-32 (cycle 10), D-CVS-36 (cycle 11), D-CVS-39/40/41 + 2 collateral (this cycle)
- D-CVS-25 / D-CVS-38 — sibling consumer-coverage callouts
