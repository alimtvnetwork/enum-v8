# Changelog

All notable changes to **enum-v5** are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

The release pipeline extracts the matching `## [vX.Y.Z]` section as the
GitHub Release body — keep entries small, sectioned, and human-readable.

---

## [Unreleased]

### Changed
- **Cycle 48 (2026-05-06)** — **S-115 SHIPPED (audit-probe correctness — sentinel-aware upstream-clone detection).** Empirically falsified the original S-115 framing: `Get-UpstreamPackages` already walks `coredata/*` recursively (`Get-ChildItem -Recurse`) and a live indexing run against `/tmp/core-v9-upstream` correctly resolved `coreonce.NewAnyErrorOnce`, `coregeneric.Hashmap`, `corestr.New`, `coredynamic` (177 packages indexed). The real defect was operator-side — when the upstream clone is **missing**, audit probes that read source via `rg`/`grep` directly silently return 0 hits and produce false fabrication conclusions (R-CVS-01/02/03 were all this same drift class). Built two guardrails in `scripts/spec-api-check.psm1` v1.2.0: (1) **`Test-UpstreamClone` exported helper** returning `{ Ok; Path; Reason; PackageCount }` with `-AutoClone` for one-shot remediation, sentinel = `coredata/coregeneric`, reasons `missing`/`sentinel-missing`/`clone-failed`/`ok`. (2) **Sentinel-missing warning inside `Get-UpstreamPackages`** so every spec-api-check run surfaces the drift even if the helper isn't called explicitly. Wired into `scripts/CoveragePreChecks.psm1` "Spec-API Lint" phase: replaces the plain `Test-Path` skip with `Test-UpstreamClone` so a directory that exists but lacks the sentinel (wrong branch, partial clone) skips with a precise reason instead of producing false negatives. Smoke test `tests/scripts/Test-UpstreamClone.ps1` covers 4 cases — all 7 assertions pass via `nix run nixpkgs#powershell`. `package.json` 0.17.0 → 0.18.0.
- **Cycle 47 (2026-05-06)** — AB-residual deep-probe of `spec/01-app/10-reflection-and-dynamic.md` (Cycle-22 carry-over). 5 of 6 ❓ resolved: item 6 (`isany.DeepEqual`) → ✅ verbatim; item 5 (`coreonce` lazy-binding) → ❌ NEW **C-CVS-63 HIGH** + R-CVS-03 retraction (package exists at `coredata/coreonce/` but is typed memoization not reflection-binding); item 2 (`corejson` unsafe fast-path) → ❌ NEW **C-CVS-64 HIGH** (zero `unsafe` imports); item 3 (type-cache) → retained ❓ plausible-no-emitter; items 1+4 → out-of-band. Spawned AJ-17b + AJ-19b (BLOCKED, folded). NEW S-115 (harden `Get-UpstreamPackages` to walk `coredata/`). 🎉 **AB-residual deep-probe sweep across `spec/01-app/` COMPLETE** — all 7 AB cycles' ❓ pools resolved; zero active probe targets remain. AB-residual `spec/01-app/` ❓ pool 11 → 6 (all OOB). §10 verifiable 38.5% → 37.5%. Cumulative AB ❌ 51 → 53 (CRITICAL 23, HIGH +2). Audit file: `spec/07-code-vs-spec-audits/36-cycle47-AB-residual-spec01-reflection.md`. Spec changelog → spec-v0.52.0.
- **Cycle 46 (2026-05-06)** — AB-residual deep-probe of `spec/01-app/08-validators.md` (Cycle-21 carry-over). All 6 ❓ items resolved: 2 ❓→✅ (row 43 `errcore.VarTwoNoType`+`ValidationFailedType`, row 45 `regexnew.New.Lazy`), 3 ❓→ⓘ "upstream-only" per S-109 (rows 42, 44), 1 row 46 → out-of-band advisory. **No new findings.** §08 ❓ pool fully cleared (6 → 0); §08 verifiable score 33.3% → 42.9%. AB-residual `spec/01-app/` ❓ pool 17 → 11 (only Cycle 22 reflection still has active probe targets). Cumulative AB ❌ unchanged at 51 (CRITICAL 23). Audit file: `spec/07-code-vs-spec-audits/35-cycle46-AB-residual-spec01-validators.md`. Spec changelog → spec-v0.51.0.
- **Cycle 45 (2026-05-06)** — AB-residual deep-probe of `spec/01-app/11-versioning.md` (Cycle-23 carry-over). Single ❓ resolved via `/tmp/core-v9-upstream` evidence: row 6 (`coreversion` ↔ `coregeneric.Collection` interop) → ❌ DEMOTION — zero `coregeneric` imports in `coreversion/`; ships own hand-rolled `VersionsCollection`. **NEW C-CVS-62 HIGH**. §11 ❓ pool fully cleared (1 → 0); §11 verifiable score unchanged at 18.2%. AB-residual `spec/01-app/` ❓ pool 18 → 17. Cumulative AB ❌ 50 → 51 (CRITICAL still 23). Spawned AJ-21b (BLOCKED, folded into AJ-21). No code or spec rewrites. Audit file: `spec/07-code-vs-spec-audits/34-cycle45-AB-residual-spec01-versioning.md`. Spec changelog → spec-v0.50.0.
- **Cycle 44 (2026-05-06)** — AB-residual deep-probe of `spec/01-app/07-conditional-and-utilities.md` (Cycle-20 carry-over). 2 of 3 ❓ items resolved via `/tmp/core-v9-upstream` evidence: row 51 (`LazyLock` lazy+cache behaviour ✅; NEW D-CVS-66 LOW mechanism-name drift: `sync.Mutex`+`isCompiled` guard, NOT `sync.Once`), row 52 (`corecmp` constants `CompareEqual/Less/Greater` = 0/-1/1 ✅ verbatim). Row 50 (advisory pitfall) deferred to Task AC. AB-residual `spec/01-app/` ❓ pool 20 → 18. §07 verifiable score 70.6% → 73.7%. Spawned AJ-04b (BLOCKED). No code or spec rewrites; audit promotion only. Audit file: `spec/07-code-vs-spec-audits/33-cycle44-AB-residual-spec01-conditional.md`. Spec changelog → spec-v0.49.0.
- **Cycle 43 (2026-05-06)** — AB-residual deep-probe of `spec/01-app/09-converters.md` (Cycle-19 carry-over). 4 of 8 ❓ items resolved via `/tmp/core-v9-upstream` evidence: 3 ❓→✅ (rows 57, 58, 60), 1 ❓→❌ (row 62: NEW C-CVS-61 CRITICAL — `errcore.OverflowType` fabricated). NEW D-CVS-65 (LOW): `BytesTo.PrettyJsonString` should be `converters.PrettyJson.Bytes.Safe`. AB-residual `spec/01-app/` ❓ pool 24 → 20. §09 verifiable score 66.7% → 68.4%. Spawned AJ-03b + AJ-44 (BLOCKED by spec/01-app freeze). No code or spec rewrites; audit promotion only. Audit file: `spec/07-code-vs-spec-audits/32-cycle43-AB-residual-spec01-converters.md`. Spec changelog → spec-v0.48.0.
- **Cycle 40 — S-114 SHIPPED (`Resolve-TestSuiteRoot` helper + 4 callsite fixes)** — Discovered while scoping AL: `scripts/PackageCoverage.psm1` (lines 45 + 62) hard-coded `./tests/integratedtests/$pkg/...` even though this repo's tests live under `tests/creationtests/` (Core-memory rule). The TCP command would silently fail on every package. Same bug pattern in `scripts/TestRunner.psm1` (TP command, lines 35/42/44), `scripts/Help.psm1` (`Invoke-IntegratedTests`, lines 27/31), and `scripts/PreCommitCheck.psm1:55`. Root-cause fix: added a shared `Resolve-TestSuiteRoot` helper to `scripts/Utilities.psm1` (~30 LOC) that probes both candidate roots and prefers `creationtests`, with optional `-Package` parameter for per-package resolution and graceful fallback to `creationtests` when neither exists (so the downstream `go test` error is the user-facing diagnostic — "no Go files in ..."). Refactored all four callsites to use the helper instead of hard-coding the path. `CoverageRunner.psm1` (lines 103-107) was already correct (loops over both names) — left untouched. Smoke-tested via 5 cases at `tests/scripts/Test-ResolveTestSuiteRoot.ps1`: (1) creationtests-only layout → `creationtests`; (2) `-Package` known → `creationtests`; (3) `-Package` missing → fallback default; (4) legacy integratedtests-only layout → `integratedtests`; (5) live enum-v5 repo → `creationtests`. **All 5 pass.** Module-import smoke also confirms `Resolve-TestSuiteRoot` is exported and all four refactored modules parse cleanly. `package.json` 0.9.0 → 0.10.0.
- **Cycle 39 — S-111 SHIPPED (GoConvey-only sub-pattern callouts)** — Closes the carry-forward from Cycle 37 (D-CVS-64). Added two new spec sections that surface the GoConvey-only sub-pattern `enum-v5` itself uses: (1) `spec/06-testing-guidelines/02-test-case-types.md` gains a "Sub-Pattern: GoConvey-Only (Local Wrapper)" block between the Style D footnote and the `CaseV1` heading — describes when to use it, gives a worked example from `tests/creationtests/AllEnums_ContractsTesting_test.go` (`Convey` + `EnumTestWrapper` registry + `LogShouldDiffMessage` + `So(diff, ShouldBeEmpty)`), and provides an equivalence table mapping upstream primitives (`CaseV1`, `coretests.GetAssert.ShouldBeEqualMap`, `args.Map`, `t.Run`, `tc.ShouldBeEqualFirst`) to their GoConvey-only counterparts; (2) `spec/06-testing-guidelines/05-assertion-patterns.md` gains a "Sub-Pattern: GoConvey-Only Diff Assertion" block appended after the named-map-types pitfall — explains why the diff-based pattern returns the empty string on success and a human-readable diff on failure (matching `ShouldBeEqualMap` ergonomics with no `coretests` dependency), lists five companion assertions, cross-links the worked example. Cosmetic spec-only change; no code or test edits. Treated as editorial under the `spec/01-app/` freeze (the freeze covers `spec/01-app/`; this edit lives in `spec/06-testing-guidelines/`). Spec changelog → **spec-v0.45.0**. `package.json` 0.8.0 → 0.9.0.
- **Cycle 38 — S-112 + S-113 SHIPPED (truthful Git-Pull phase + remote-probe skip)** — Two related defects in the `tc` runner's first phase, fixed in one pass against `scripts/TestRunnerCore.psm1` + `scripts/CoverageRunner.psm1`. **S-113 (root cause):** `Invoke-GitPull` ran `git pull` unconditionally and emitted a confusing `remote: Repository not found / fatal: repository '…/enum-v5.git/' not found` whenever the local clone had a misconfigured/private/missing `origin`. Added a two-step early probe BEFORE any `git pull`: (1) `git remote get-url origin` — if `origin` isn't configured, return `Status='skip'` with message `no origin remote` and a friendly `No 'origin' remote configured — skipping pull` line; (2) `git ls-remote --exit-code origin HEAD` — if the remote is unreachable, return `Status='skip'` with message `remote unreachable` and surface the URL so the user can fix it. Only when both probes pass do we now actually invoke `git pull`. **S-112 (truthfulness):** `Invoke-GitPull` was void-return so callers couldn't tell pass from soft-fail; `CoverageRunner.psm1:27-30` hard-coded `Register-Phase "Git Pull" "pass" "pulled from remote"` regardless of outcome — making the Phase Summary lie (`✓ Git Pull pulled from remote` rendered after `✗ git pull failed`). Refactored `Invoke-GitPull` to return `[pscustomobject]@{ Status='pass'|'warn'|'skip'; Message=... }`; `Invoke-FetchLatest` now returns `@{ GitPull; Tidy }` with the same shape for `go mod tidy`; `CoverageRunner.psm1` uses both results to register the phases truthfully. The dashboard already supports `skip → ⊘` and `warn → ⚠` glyphs (`scripts/DashboardPhases.psm1:45-51`), so no UI changes needed. Smoke-tested via `pwsh -File tests/scripts/Test-InvokeGitPull.ps1` against a fresh `git init` repo: Test 1 (no origin) → `Status=skip / Message="no origin remote"`; Test 2 (bogus unreachable origin) → `Status=skip / Message="remote unreachable"`. Both pass. The user's reported run (`./run.ps1 -tc` against a clone with stale `origin = github.com/alimtvnetwork/enum-v5.git`) will now show `⊘ Git Pull remote unreachable` instead of `✓ Git Pull pulled from remote` — the phase summary stops lying. Added regression test at `tests/scripts/Test-InvokeGitPull.ps1`. `package.json` 0.7.0 → 0.8.0.
- **Cycle 37 — S-109 SHIPPED (`tests/creationtests/` deep-probe of Cycle-15 ❓ pool)** — Settled all 10 ❓ items left over from Cycle 15 in `spec/06-testing-guidelines/` by direct inspection of every file under `enum-v5/tests/creationtests/` (14 files). Probe commands `rg -n 'coretests\.|coretestcases\.|args\.Map|args\.One|args\.Six|args\.Holder|args\.LeftRight|CaseV1|CaseNilSafe|GenericGherkins|GetAssert|ShouldBeEqualMap|ShouldBeSafe|InvokeWithPanicRecovery|results\.Result|results\.ResultAny|results\.ExpectAnyError|BaseTestCase' tests/creationtests/` returned **zero hits** — confirms `enum-v5` does not consume the upstream `coretests`/`args`/`results` framework documented in spec/06 at all; instead it uses ubiquitous GoConvey (`Convey`/`So`/`ShouldEqual`/`ShouldResemble`/`ShouldBeNil`/`ShouldBeTrue`/`ShouldBeEmpty`) over two **local** wrapper structs (`EnumTestWrapper`, `PathPatternTypeCreationTestWrapper`) with module-level registries (`var allEnumGeneralTestCases = []*EnumTestWrapper{...}`, `var pathPatternTypeCreationTestCases = [...]PathPatternTypeCreationTestWrapper{...}`, `var allScriptCreationTestCases = map[string]ScriptType{...}`) and AAA comments. Outcome: **1 ❓ → ✅** (claim 20 — diff-based assertion pattern is behaviourally evidenced via `actualEnumDynamicMap.LogShouldDiffMessage(true, header, expected); So(diff, ShouldBeEmpty)` in `AllEnums_ContractsTesting_test.go:42-47`); **9 ❓ → ⓘ "upstream-only" annotated** (CaseV1/CaseNilSafe/GenericGherkins, args.*, results.*/InvokeWithPanicRecovery, ShouldBeEqual*/ShouldBeSafe upstream-custom assertions, the 5 sub-claims of `07-diagnostics-output-standards.md`, `08-good-vs-bad.md` examples, `09-creating-custom-cases.md` `BaseTestCase` extension pattern). The 9 ⓘ items remain blocked by Task **AB** for upstream-clone promotion but are no longer "unknown". Cycle-15 verifiable subset grows 22/22 → 23/23 (still 100%); spec/06 unknown ❓ pool drops **10 → 0**. New LOW finding **D-CVS-64** raised: `02-test-case-types.md` + `05-assertion-patterns.md` don't surface the **GoConvey-only sub-pattern** that `enum-v5` itself is a worked example of (plain `So(...)` + AAA + plain registries, no `args.Map`/`BaseTestCase`); tracked as carry-forward suggestion **S-111** (cosmetic, non-blocking, deferrable to Task AC). Audit file: `spec/07-code-vs-spec-audits/29-cycle37-S109-creationtests-deep-probe.md`. Spec changelog → **spec-v0.44.0**. `package.json` 0.6.0 → 0.7.0.
- **Cycle 36 — S-103 SHIPPED (portable runner spec reorg)** — Moved the two explicitly-portable runner files (`spec/03-powershell-test-run/08-generic-go-test-coverage-runner.md` and `09-ai-agent-complete-reference.md`) into a new `spec/03-powershell-test-run/portable/` sub-directory and renumbered them to `01-` and `02-` inside it. Added `spec/03-powershell-test-run/portable/README.md` explaining the scope split (portable vs `enum-v5`-specific), listing the two files, and codifying three editor rules to keep the portability promise intact (no enum-v5-specific paths/flags here, `tests/integratedtests/` references describe upstream `core-v9` consumer layout, keep portability promise explicit per file). Updated the two live cross-refs to the new paths: `spec/00-llm-integration-guide.md` line 2380 (AI-agent test command reference) and `spec/04-tooling/03-powershell-implementation.md` line 456 (file-table row); also fixed the table row inside the moved `02-ai-agent-complete-reference.md` that pointed to its sibling. Historical references in `spec/CHANGELOG.md` Cycle-16 entry, `spec/07-code-vs-spec-audits/17-cycle16-powershell-test-run.md`, and `spec/99-audits/01-original-11-step-plan.md` are intentionally left as-is — they document audit history at the time those cycles ran. Acceptance check `rg -n 'spec/03-powershell-test-run/(08|09)-' spec/ --glob '!spec/CHANGELOG.md' --glob '!spec/07-code-vs-spec-audits/**' --glob '!spec/99-audits/**'` returns zero hits. The structural split replaces reliance on the Cycle-16 top-of-file consumer-coverage callouts (D-CVS-47/48) with a directory-level signal that's harder to miss. Spec changelog → **spec-v0.43.0**. `package.json` bumped 0.5.0 → 0.6.0.
- **Cycle 35 — S-104 + S-105 SHIPPED (historical-naming callout + index-drift CI guard)** — Two carry-forward suggestions retired in one pass. **S-104:** added a prominent top-of-file callout to `cross-repo/core-v8/README.md` explaining the four invariants future editors must respect — (1) the `core-v8` directory name is historical and intentional (mirrors a separate upstream repo), (2) the actual import path used by `enum-v5` source is `github.com/alimtvnetwork/core-v9` (renamed 2026-05-05, tagged `v1.5.8`), (3) spec/script text that references *this directory* must always write `cross-repo/core-v8/` even when the surrounding sentence is about `core-v9` content, (4) the historical `enum-v1` / `core-v8` body references must NOT be rewritten (Core-memory rule). This closes the Cycle-17 root cause of D-CVS-49/52/53/55 (5 broken `cross-repo/core-v9/` paths) at the point of truth instead of per-cite-site clarification. Body content untouched per Core memory. **S-105:** added `scripts/ci/check-issues-index-drift.py` plus 5 unittest cases in `scripts/ci/test_check_issues_index_drift.py` (all pass) and a new `issues-index-drift` job in `.github/workflows/ci-guards.yml` that depends on `python-tests` and mirrors the `collision-audit` job pattern. The script extracts every `| NN |`-prefixed row from `spec/02-app-issues/00-issues-index.md` (canonical) and `spec/02-app-issues/README.md` (human-readable), then compares both **row count AND id-set** so it catches the original Cycle-18 failure mode (stale-by-4-rows for ~14 days) AND the subtler same-count-different-id case. On drift it exits 1 with `Missing from README: [...]` / `Missing from index: [...]` diffs; on missing files it exits 2. Live repo reports `OK: spec/02-app-issues index in sync (9 rows).` so the guard is adopted at a clean baseline. `package.json` bumped 0.4.0 → 0.5.0.
- **Cycle 34 — backlog hygiene sweep (S-001 / S-003 / S-004 closed)** — Three open suggestions retired in one pass. **S-001** (pin Go toolchain to 1.22 as a Task-W stopgap) closed as **obsolete**: Tasks W + AG already removed the dual-path `replace` bridge by renaming upstream to `module github.com/alimtvnetwork/core-v9` + tagging `v1.5.8` and pinning `enum-v5/go.mod` to `core-v9 v1.5.8` directly (per Core memory). Pinning to Go 1.22 today would mask a working modern setup and re-introduce the lock-in risk the original suggestion itself flagged. **S-003** (stale `integratedtests` path in `spec/06-testing-guidelines/01-folder-structure.md`) closed as **already-resolved**: line 3 of that file already carries the upstream-scope disclaimer added in an earlier audit cycle (*"⚠️ Scope: the layout below describes upstream `core-v9`. enum-v5 uses a single shared `tests/creationtests/` package…"*) — the `integratedtests` references in the body intentionally document the upstream consumer layout, not stale paths. **S-004** (stale `integratedtests` references in `spec/00-llm-integration-guide.md`) resolved via **re-framing** to match the Cycle-12/15/17/18 pattern: added an inline upstream-vs-enum-v5 scope callout above the Test-Folder-Structure code fence (line 826) cross-linking both `spec/01-app/13-testing-patterns.md` §6.1 and `spec/06-testing-guidelines/01-folder-structure.md`. The Decision-Matrix Style-D row (line 36) is now disambiguated by the same callout. PI-002 (Cross-spec stale `integratedtests/` paths) plan marked complete in `.lovable/memory/pending-issues/01-all-pending-issues.md`. Open-suggestions list shrinks: S-001 → done, S-003 → done, S-004 → done, S-002 → deferred to Task AC. `package.json` bumped 0.3.0 → 0.4.0.
- **Cycle 33 — S-107 SHIPPED (goconvey-failure summarizer for `failing-tests.txt`)** — Added `Get-GoconveyFailureSummary` to `scripts/TestLogWriter.psm1` (also exported from the module). For each failed test's captured block, pairs every `Expected:` line with the nearest following `Actual:` / `(Line N)` / `Message:` in an 8-line look-ahead window, breaking on the next `Expected:` so a multi-failure block produces multiple triplets. `Write-TestLogs` now prepends a compact `── Failure summary (N) ──` section to each `FAIL:` entry, listing `#i Expected: ... | Actual: ... (Line N) [Message]` BEFORE the noisy raw block (which is preserved underneath for full context). This addresses the long-standing ergonomics issue diagnosed in Cycle 24 where `failing-tests.txt`'s Pass-2 collector buries goconvey `Failures:` blocks under the trailing `N total assertions` / `* assertions: N ✓` noise — the real signal (Expected vs Actual + line number) was always present but invisible. Smoke-tested via `nix run nixpkgs#powershell` on a synthetic 2-failure goconvey block: both triplets extracted correctly (`Expected=<expected-value-1>|Actual=<actual-value-1>|Line=42` and `Expected=5|Actual=7|Line=88|Message=counts mismatch`), empty blocks return an empty array, non-goconvey blocks (e.g. plain `panic:` traces) return an empty array with no false positives. `package.json` bumped 0.2.0 → 0.3.0 per the mandatory-minor-bump rule.
- **Cycle 32 — S-110 SHIPPED (restored 3 standalone coverage utilities)** — Built the sibling helpers documented alongside S-108 in `spec/03-powershell-test-run/06-coverage-prompt-generator.md`. (1) `scripts/coverage/Get-UncoveredLines.ps1` returns a single object `{SourceFile, UncoveredCount, Ranges}` for one source file using the same `coverage.out` block parser and range-collapse formatter as the S-108 main script. (2) `scripts/coverage/Get-FunctionCoverage.ps1` parses `go tool cover -func` output (string array, multi-line string, or file path), filters strictly below `-Threshold` (default 100.0), and returns objects sorted ascending by coverage. (3) `scripts/coverage/Get-PackageCoverageReport.ps1` combines both for a single `-Package`, with `-Format text|markdown|json` and optional `-OutputFile`. All three avoid the `$Input` automatic-variable shadowing pitfall (use `$Source`). End-to-end smoke-tested via `nix run nixpkgs#powershell` against the same 6-block synthetic `coverage.out` + 3-line func output used for S-108: text format renders `SplitLeftRight  40.0%   L8-L12, L18` with the file path on the next line; markdown produces a valid 4-column table; JSON parses cleanly with the expected fields. **D-CVS-62 is now fully closed** — `scripts/coverage/` and the spec are in lockstep. `package.json` bumped 0.1.0 → 0.2.0 per the mandatory-minor-bump rule.
- **Cycle 31 — S-108 SHIPPED (restored missing `scripts/coverage/Generate-CoveragePrompts.ps1`)** — Created the script that has been silently no-op'd by the `if (Test-Path $promptScript)` guards in `scripts/CoverageRunner.psm1:313-316` and `scripts/PackageCoverage.psm1:145-150` since at least Cycle 27 (D-CVS-62). Implementation (~210 LOC, PS 5.1+) follows `spec/03-powershell-test-run/06-coverage-prompt-generator.md` exactly: parses `coverage.out` for zero-count statement blocks (`<file>:<sl>.<sc>,<el>.<ec> <stmts> <count>`), parses `go tool cover -func` lines (`<file.go>:<line>:<TAB><Func><TAB><pct>%`), filters to sub-100% functions, sorts ascending by coverage, batches at `-BatchSize` (default 500) into `coverage-prompt-N.txt`, emits `prompts-summary.json`. Range collapse renders contiguous uncovered lines as `L15-L17, L22`. Avoids PowerShell's `$Input` automatic-variable shadowing pitfall (uses `$Source` for the helper param). **Smoke-tested end-to-end** via `nix run nixpkgs#powershell` against synthetic 6-block `coverage.out` + 3-line func output: produced exactly the spec sample format (correct sort order — SplitLeftRight 40.0% before NewError 66.7% — and correct ranges). Sibling standalone utilities (`Get-UncoveredLines.ps1`, `Get-FunctionCoverage.ps1`, `Get-PackageCoverageReport.ps1`) listed in the same spec section are still missing — tracked as new suggestion **S-110**. `package.json` bumped 0.0.0 → 0.1.0 per Core memory's mandatory-minor-bump rule for code changes outside `.release/`.
- **Cycle 27 — AB residual deep-probe of `scripts/*.psm1` + `.github/workflows/*.yml`** — first dedicated "scripts deep-probe" audit cycle. Promoted **11 ❓** from runner-internal/workflow-internal claims using direct file evidence: §03 Cycle-16 claim 8 (parallel default vs `--sync` opt-in confirmed at `scripts/CoverageRunner.psm1:124,196,212`), claim 12 (JSON super-set: actual schema in `scripts/PreCommitCheck.psm1:169` writes 7 fields — `timestamp`, `passed`, `checkedCount`, `passedCount`, `failedCount`, `source`, `failures[]` — vs spec example showing 6, missing `source`), claim 14 (threading model = `min(packages, 2×CPU)` PS7 runspace pool at line 213), claim 16 → ⚠️ (D-CVS-62: `scripts/coverage/Generate-CoveragePrompts.ps1` MISSING but referenced with `if (Test-Path)` guard so silently no-ops). §04 Cycle-17 claims 7/9/10/11/15 confirmed via `.github/workflows/ci.yml`, `.github/workflows/release.yml`, dashboard modules, and `run.ps1:75-89` module-loading. Surfaced **D-CVS-62** (LOW: missing prompt-generator script → suggestion S-108) and **D-CVS-63** (LOW: spec JSON schema missing `source` field — fixed in same cycle by editing `spec/03-powershell-test-run/04-pre-commit-api-checker.md`). AB-residual ❓ count: **53 → 42**. Cumulative AB ❌ unchanged at **49** (24 CRITICAL). New audit file `spec/07-code-vs-spec-audits/28-cycle27-AB-scripts-deep-probe.md`. Spec changelog → `spec-v0.42.0`.
- **Cycle 30 — S-106 v2 SHIPPED (Go AST signature lint)** — added `scripts/specapisig/` (Go program, ~280 LOC) that AST-parses upstream `core-v9` + local enum-v5 and emits a JSON signature index of every exported top-level func/method (parameters with names+types, results, variadic flag, file:line). Paired PowerShell driver `scripts/spec-api-sig-check.psm1` v1.0.0 walks every `pkg.Symbol(...)` call-site in `spec/01-app/`, splits args by balanced parens (string-aware), and verifies arity against the upstream signature. Variadic candidates accept `expected-1..N` args. **End-to-end design verified via Python port:** scanned 163 spec call-sites and correctly flagged all 4 `errcore.VarTwo(...)` sites as 4-arg or 3-arg calls vs. the real 5-arg `(isIncludeType bool, firstName string, firstValue any, secondName string, secondValue any) string` signature — exactly the C-CVS-44 class of defects v1 cannot catch (60 OK, 99 unresolved/handled-by-v1, 4 arity mismatches). Wired into both `scripts/CoveragePreChecks.psm1` (`Spec-API Sig` dashboard phase, runs after the v1 lint, regenerates the index every run via `go run ./scripts/specapisig`) and `.github/workflows/ci-guards.yml` (new "Build Go signature index" + "Run S-106 v2" steps appended to the `spec-api-lint` job, sharing the same strict-mode toggle as v1). With v2 in place, **arity drift can no longer escape into spec/**, complementing v1's name-fabrication coverage.
- **Cycle 29 — S-106 v1.1 (false-positive cleanup)** — bumped `scripts/spec-api-check.psm1` to v1.1.0. Three fixes: (1) **indented-fence detection** — fences nested inside numbered-list items (e.g. `   ```go`) were treated as prose, so all bindings/refs inside leaked; the regex is now `^\s*```(\w*)`. (2) **local enum-v5 indexing** — `Get-UpstreamPackages` now also walks the project root (with skip-list `node_modules|cross-repo|tests|scripts|spec|src|public|data|cmd|assets|configs|internal`) so spec references to `compressformats`/`logtype`/`inttype`/etc. resolve. (3) **expanded allow-list** — added Go stdlib (`unsafe`, `runtime`, `path`, `filepath`, `url`, `net`, `rand`, `crypto`, `sha256`, `base64`), template/pseudo-package names (`emailvalidator`, `corev8`, `expected`, `validator`, `downstream`, `registry`), and a CommonLocalVarNames bucket (`tc`, `col`, `svc`, `cart`, `safe`, `payload`, `pattern`, `result`, `input`, `status`, `err`, `opts`, `cfg`, `req`, `resp`, `ctx`, `val`, `item`, `items`, `row`, `rows`, `msg`, `data`, `out`, `buf`) that frequently appear with elided bindings. Added receiver-name detection (`func (it Variant) ...` now binds `it` as a local) and `vN`-versioned local skip (`v1`, `v2`, …). Verified via Python port: **package-fabrication false positives drop 22 → 0** (34 refs → 0 refs); 43 unique sym-fabrications remain — all mapped to existing AB findings (C-CVS-11..59) that AJ-01..43 will resolve. The lint is now signal-clean enough to enable `-StrictExitCode` without spurious noise the moment the AJ rewrites land. Indexed-package count grew 182 → ~259 (upstream + local).
- **Cycle 28 — S-106 wired into GitHub Actions** — added `spec-api-lint` job to `.github/workflows/ci-guards.yml`. Clones upstream `core-v9 @ v1.5.8` to `/tmp/core-v9-upstream`, then runs `Invoke-SpecApiCheck` via PowerShell. Smart strict-mode toggle: PRs touching `spec/` or `scripts/spec-api-check.psm1` run with `-StrictExitCode` (any fabrication fails the job); all other runs are warn-only so the 49 known AB-flagged fabrications don't block unrelated PRs while AJ-01..43 are still gated by the `spec/01-app/` freeze. Closes the "CI workflow gate" remaining-tasks item from Cycle 27.
- **Cycle 27 — S-106 wired into `run.ps1 -tc` pre-checks** — added a "Spec-API Lint" phase to `scripts/CoveragePreChecks.psm1` that runs after the safeTest boundary lint and before the Go auto-fixer. Soft gate by default (warn-only) so the 49 known fabrications (AJ-01..43 backlog) don't block test runs while the freeze is in place. New flags wired through `run.ps1`: `--no-spec-api` (skip) and `--strict-spec-api` (fail TC on any fabrication — for CI). The phase auto-skips when `/tmp/core-v9-upstream` (or `scripts/spec-api-check.psm1`, or `spec/01-app/`) is absent, with a helpful clone hint. This satisfies the "wire S-106 into pre-checks" item from Cycle 26's remaining-tasks list and gives the dashboard a permanent line of defense against new fabrications landing in PRs.
- **Cycle 26 — S-106 lint BUILT + self-audit retractions** — created `scripts/spec-api-check.psm1` v1.0.0 (presence-only fabrication lint over 182 upstream packages / 10,216 symbols). First run produced **R-CVS-01** (retracts C-CVS-29: `coredynamic` exists at `coredata/coredynamic/` — 20+ files; the 8 specific methods like `AllFields`/`SetField`/`InvokeMethod` remain confirmed sym-fabrications) and **R-CVS-02** (retracts C-CVS-51: `corestr` exists at `coredata/corestr/` — 30+ files; `StringBuilder`/`IsValidUTF8` remain sym-fabrications). Both original audits ran `find . -type d -name {pkg}` from upstream root and missed the nested `coredata/` parent. Cumulative AB ❌ drops 50 → **49** across 7 sections; CRITICAL count drops 24 → **22**. Added **C-CVS-60** for residual low-impact sym-fabs (`coronce.New`, etc.). AJ-15 split into AJ-15a (path-qualify §2 → `coredata/coredynamic`) + AJ-15b (purge fabricated symbols); AJ-36/37/38 re-scoped to keep `corestr` package and purge fabricated symbols only. S-106 v1.0 limits: presence-only — does NOT catch arity/return-type drift (S-106 v2 needs Go AST pass). With the lint in place, AJ-01..43 may now safely proceed once the freeze is lifted. See `spec/07-code-vs-spec-audits/27-cycle26-S106-self-audit-retractions.md`. Spec changelog → `spec-v0.41.0`.
- **Cycle 25 spec audit (Task AB pass 7) — COMPLETES AB sweep of `spec/01-app/`** — `spec/01-app/16-security.md`: 13 ❓ → **3 ✅ / 1 ⚠️ / 9 ❌ / 0 ❓** (verifiable score **66.7 %**). §16 inherits fabrication: trust-boundary rules cite fabricated `coredynamic.*` (Cycle 22 inheritance), `corevalidator.New.Line/Slice` fluent (Cycle 21 inheritance), and a never-existed `corestr` package. Surfaced **9 NEW contradictions** (C-CVS-51..59) + **1 NEW drift** (D-CVS-61): six CRITICAL (`corestr` package doesn't exist anywhere in upstream — `StringBuilder`/`IsValidUTF8` fabricated, use stdlib `strings.Builder`/`unicode/utf8.ValidString`; `errcore.InvalidInput.MergeError(...)` won't compile — `InvalidInput` not exposed as `errcore` category, only `ShouldBe`/`Expected`/`StackEnhance` are; `coredynamic.AllFields`/`SetField`/`InvokeMethod` rules unactionable; trust-boundary §6 example uses fabricated `corevalidator.New.Line` fluent + `result.IsFailed()` shape), two HIGH (`corevalidator.New.Slice.MaxLength(N)` cited in 4 places — all fabricated; `errcore.VarTwo` example reproduces C-CVS-44 missing-`isIncludeType` defect — folded into AJ-28), one LOW drift (`coregeneric` import path should be `coredata/coregeneric`). Spawned AJ-35..43 (BLOCKED). Cumulative AB ❌ count: **50** across 7 sections (~54 % fabrication rate, 24 CRITICAL ~48 %). **🎉 AB sweep of `spec/01-app/` is COMPLETE** — all 7 sections that held ≥10 ❓ have been promoted. **S-106 lint remains MANDATORY before any AJ rewrite.** See `spec/07-code-vs-spec-audits/26-cycle25-AB-security.md`. Spec changelog → `spec-v0.40.0`.
- **Cycle 24 spec audit (Task AB pass 6)** — `spec/01-app/15-observability.md`: 13 ❓ → **6 ✅ / 7 ❌ / 0 ❓** (verifiable score **74.1 %** — drops from clean baseline). Surfaced **7 NEW contradictions** (C-CVS-44..50): four CRITICAL (`errcore.VarTwo` example missing mandatory leading `isIncludeType bool` parameter — spec example won't compile; `VarTwo`/`VarTwoNoType`/`MessageVarMap` return `string` not `error` — entire §2 misframes helpers as error builders; test-failure output format `Test #N — {scenario}: should be equal\n  expected:\n  actual:` fabricated with zero matches in `coretests/results/`; `errcore.HandleErr` does NOT attach stack-enhanced wrapping — implementation is just `panic(err.Error())`, so §3 rule 2 cites the wrong rationale), two HIGH (`VarTwoNoType` is `VarTwo(false, ...)` not a distinct helper; `coretests/results/ResultAny.go` does not exist — real files are `Result.go`/`ResultAssert.go`/`Results.go`), one MEDIUM (`StackEnhance` documented as 2-method but exposes 8 including the `*Skip` family used by 24 internal call-sites). Spawned AJ-28..34 (BLOCKED). Cumulative AB ❌ count: **41** across 6 sections (~52 % fabrication rate, 18 CRITICAL). **S-106 lint remains MANDATORY before any AJ rewrite.** See `spec/07-code-vs-spec-audits/25-cycle24-AB-observability.md`. Spec changelog → `spec-v0.39.0`.
- **Fixed `Test_AllEnums_ContractsTesting`** — regenerated `tests/creationtests/allEnumGeneralTestCases.go` fixtures to match upstream `core-v9 v1.5.8`. Two drift causes: (1) `RangeNamesCsv()` formatter changed `Name[N]` → `Name(N)`, and (2) several upstream/local enums gained or renamed members (e.g. `stringcompareas` added `Glob`/`NonGlob`, renamed `Contains`→`IsContains`). Workflow: temporarily renamed `GenerateTestCases` → `Test_GenerateTestCases`, ran it to regenerate the fixture body via `generateAllEnumGeneralTestCases(false)`, spliced output into the file (preserving original imports header), restored the original name, and escaped a literal `"` map key in `quotes.Quote`. Test now passes (`ok tests/creationtests`).
- **Diagnosed `failing-tests.txt` log gap** — `scripts/TestLogWriter.psm1`'s Pass-2 collector buries goconvey "Failures:" blocks under thousands of "N total assertions" lines, making real failure causes invisible. Suggestion **S-107** opened to add a goconvey-failure summarizer that surfaces `Expected:`/`Actual:`/`Line N:` triplets at the top of each failed-test block.
- **Cycle 23 spec audit (Task AB pass 5)** — `spec/01-app/11-versioning.md`: 11 ❓ → **2 ✅ / 8 ❌ / 1 ❓** (verifiable score **18.2 % — new WORST score in project**). Surfaced **8 NEW contradictions** (C-CVS-37..43): four CRITICAL (`coreversion.Parse` constructor fabricated — real is `New.Default(s) Version` no error; method-style `LessThan/Equal/GreaterThanOrEqual` fabricated — real is package-level `Compare(left, right *Version) corecomparator.Compare`; `versionindexes` constants `V1/V2/V8` "version eras" fabricated — real consts `Major=0/Minor=1/Patch=2/Build=3/Invalid=4` index version-component positions; package path `versionindexes/` wrong — real path `enums/versionindexes/`'s parent fabrication for the conceptual framing), three HIGH (typed accessors `Major()/Minor()/Patch()` fabricated — real is public-field struct; `errcore.FailedToConvertType` wrapping rationale fabricated — zero `errcore` references in package; package-path drift), one LOW (`String()` not guaranteed to return `"v1.2.3"`). Spawned AJ-21..27 (BLOCKED). Cumulative AB ❌ count: **34** across 5 sections (~55 % fabrication rate, ~44 % CRITICAL severity). **S-106 lint is now MANDATORY before any AJ rewrite.** §11 is uniquely bad: C-CVS-43 fabricates a different *purpose* for the package (eras vs. component positions), not just a wrong API surface. See `spec/07-code-vs-spec-audits/24-cycle23-AB-versioning.md`. Spec changelog → `spec-v0.38.0`.
- **Cycle 22 spec audit (Task AB pass 4)** — `spec/01-app/10-reflection-and-dynamic.md`: 19 claims → **5 ✅ / 8 ❌ / 6 ❓** (verifiable score **38.5 % — second-worst in project after §08**). Surfaced **8 NEW contradictions** (C-CVS-29..36): five CRITICAL (entire `coredynamic` package fabricated — `coredynamic/` directory does not exist in upstream `core-v9 v1.5.8`; `reflectcore.IsPointer/IsStruct/...` bare-function predicates; `reflectcore.WalkFields`; `reflectcore.GetTag`; `reflectcore.DerefAll` — all absent) and three HIGH (`internal/reflectinternal` "off-limits" framing misleading because `reflectcore/vars.go` publicly re-exports 15 internal symbols; no `errcore` panic-wrapping facade; `InvokeMethod`/`HasMethod` mistake-row guidance fabricated). Real `reflectcore` is a thin re-export shim over `internal/reflectinternal`; field walking lives in `reflectcore/reflectmodel.FieldProcessor`. Spawned AJ-15..20 (BLOCKED). Cumulative AB ❌ count: **26** across 4 sections (~45 % fabrication rate). **S-106 lint is now MANDATORY before any AJ rewrite.** See `spec/07-code-vs-spec-audits/23-cycle22-AB-reflection-and-dynamic.md`. Spec changelog → `spec-v0.37.0`.
- **Cycle 21 spec audit (Task AB pass 3)** — `spec/01-app/08-validators.md`: 18 ❓ → **4 ✅ / 8 ❌ / 6 ❓** (verifiable score **33.3 % — lowest in project**). Almost the entire chapter is fabricated. Surfaced **8 NEW contradictions** (C-CVS-21..28): four CRITICAL (5-method validator contract + fluent-builder API + `Validate(input) Result` + "PowerShell parses validator output" attribution pipeline all fabricated) and four HIGH (`RangeValidator`, `StringCompareAs` reclassified as parameter enum, `Result` type, custom-validator template). Real `corevalidator/` exposes `LineValidator{LineNumber, TextValidator}` with `IsMatch(lineNumber, content, isCaseSensitive) bool`. All blocked pending freeze waiver. Spawned AJ-08..14. Cumulative AB ❌ count: **18** across 3 sections (~41 % fabrication rate). **Strongly recommend S-106 lint lands before any AJ rewrite.** See `spec/07-code-vs-spec-audits/22-cycle21-AB-validators.md`. Spec changelog → `spec-v0.36.0`.
- **Cycle 20 spec audit (Task AB pass 2)** — `spec/01-app/07-conditional-and-utilities.md`: 17 ❓ → **12 ✅ / 5 ❌ / 3 ❓** (verifiable score 70.6 %). Surfaced **5 NEW contradictions** (C-CVS-16..20): three HIGH (`TypedErrorFunctionsExecuteResults` wrong shape; `coremath` claims 7 type families but has 3; `Collection.ToMap()` fabricated) and two CRITICAL (`namevalue.NewInstance` constructor doesn't exist + entire `keymk.New.Compile(...)` snippet fabricated). All blocked pending freeze waiver. Spawned AJ-04..07. Cumulative AB ❌ count: **10** across 2 sections (~25 % fabrication rate). See `spec/07-code-vs-spec-audits/21-cycle20-AB-conditional-and-utilities.md`. Spec changelog → `spec-v0.35.0`.
- **Cycle 19 spec audit (Task AB pass 1)** — upstream `core-v9 v1.5.8`
  cloned to `/tmp/core-v9-upstream`; first ❓→ground-truth promotion pass on
  `spec/01-app/09-converters.md`. Result: **23 ❓ → 10 ✅ / 5 ❌ / 8 ❓**
  (verifiable score 66.7 %). Surfaced **5 NEW contradictions**
  (C-CVS-11..15): four HIGH (`StringTo.Integer64`, `StringTo.Float32`,
  `StringTo.Bool`, `PrettyJson.String`/`.FromAny`) and one CRITICAL
  (entire `typesconv` §2 + §4.3 example fabricated). All 5 are blocked
  pending a freeze waiver for `spec/01-app/`. Spawned action items
  AJ-01..03. Also corrected 2 stale Core-memory items in `mem://index.md`
  (M-CVS-01: `enum-v3`→`enum-v5` module name; M-CVS-02: upstream `go.mod`
  rename declared complete + `replace` bridge removal). See
  `spec/07-code-vs-spec-audits/20-cycle19-AB-converters-promotion.md`.
  Spec changelog → `spec-v0.34.0`.
- **Cycle 18 spec audit (Task AA + closes Task AH)** — closed
  `spec/02-app-issues/` (11 files, 402 lines) at **100 % verifiable** (21 ✅ /
  5 ❓ audit-history). Raised and resolved **5 LOW drifts (D-CVS-56 →
  D-CVS-60)**: 1 stale README index (5 open vs reality 9 resolved) + 4
  upstream-vs-`enum-v5` scope footnotes on the historical resolution files
  (`02-internal-package-coverage-policy.md`, `03-getassert-undocumented-api.md`,
  `04-testwrappers-public-surface.md`, `05-missing-params-go-files.md`).
  **🎉 Marks Task AH (cross-`spec/` cleanup sweep) Done** — every directory
  under `spec/` outside the immutable history folders is now baselined. See
  `spec/07-code-vs-spec-audits/19-cycle18-app-issues.md`. Spec changelog
  bumped to **spec-v0.33.0**.
- **Cycle 17 spec audit (Task AA + AH partial)** — closed `spec/04-tooling/`
  (10 files, 2 553 lines) at **100 % verifiable** (22 ✅ / 8 ❓ workflow-
  internal). Raised and resolved **7 LOW drifts (D-CVS-49 → D-CVS-55)** in the
  same cycle: 2 broken `cross-repo/core-v9/` paths in `00-overview.md`, 1
  missing-precedent in `04-bootstrap-into-new-repo.md` §7 (the AH-tracked
  occurrence), and 4 stale `enum-v2`/`cross-repo/core-v9` tokens in
  `06-cross-repo-sync.md` (line 80 template comment carried both stale tokens).
  Each fix includes a Core-memory clarification that `cross-repo/core-v8/`
  intentionally keeps its historical name even though the import path is
  `core-v9`. See `spec/07-code-vs-spec-audits/18-cycle17-tooling.md`. Spec
  changelog bumped to **spec-v0.32.0**.
- **Cycle 16 spec audit (Task AA + AH partial)** — closed
  `spec/03-powershell-test-run/` (9 files, 2 519 lines) at **100 % verifiable**
  (22 ✅ / 6 ❓; the 6 ❓ are runner-internal behaviours requiring a direct
  `scripts/*.psm1` probe). Raised and resolved **5 LOW drifts (D-CVS-44 →
  D-CVS-48)** in the same cycle via top-of-file consumer-coverage callouts
  (`01-overview.md`, `04-pre-commit-api-checker.md`, `08-generic-go-test-coverage-runner.md`,
  `09-ai-agent-complete-reference.md`) plus one inline rewrite
  (`06-coverage-prompt-generator.md` line 71). Folds in Task AH debt for this
  directory. See `spec/07-code-vs-spec-audits/17-cycle16-powershell-test-run.md`.
  Spec changelog bumped to **spec-v0.31.0**.
- **`spec/01-app/` DRIFT-FROZEN (Task AI)** — declared the directory closed for
  code-vs-spec drift work in `spec/CHANGELOG.md` as **spec-v0.30.0**. Allowed
  future edits: AB-driven ❓→✅ promotions, AC re-audit of §07/§09,
  upstream-API-change additions (paired with a new audit cycle row), typo/
  formatting fixes. Drift work moves to `spec/03-powershell-test-run/`,
  `spec/04-tooling/`, `spec/02-app-issues/` (Cycles 16+). Scoreboard top-line
  updated with 🧊 freeze marker.
- **Cycle 15 spec audit (Task AA)** — closed `spec/06-testing-guidelines/`
  directory at **100 % of its verifiable subset** (32 claims sampled across 10
  files; 22 ✅ / 10 ❓ pending task AB). Resolved one LOW drift (D-CVS-43) by
  adding an `enum-v5` consumer-coverage callout to `spec/06-testing-guidelines/README.md`
  and a `⚠️ Scope` warning to `01-folder-structure.md`. See
  `spec/07-code-vs-spec-audits/16-cycle15-testing-guidelines.md`. Spec changelog
  bumped to **spec-v0.29.0**.
- **core-v9 API migration (Task AM)** — Applied all confirmed `core-v9 v1.5.8`
  namespace rewrites across `enum-v5` Go source: `coredynamic.TypeName(...)` →
  `coredynamic.SafeTypeName(...)`, `converters.AnyToValueString(x)` →
  `converters.AnyTo.ValueString(x)`, `converters.Any.ToFullNameValueString` →
  `converters.AnyTo.ToFullNameValueString`, and the remaining
  `converters.StringTo*` calls → `converters.StringTo.*` methods.

### Fixed
- **Task AM / `tests/creationtests` compile blocker** — removed the obsolete
  package-level converter calls that caused `undefined: converters.StringToInteger`,
  `undefined: converters.StringToIntegerWithDefault`,
  `undefined: converters.StringToIntegerDefault`, and `undefined: converters.StringToByte`
  after upgrading to `core-v9 v1.5.8`.

- **`scripts/CoverageRunner.psm1`, `scripts/CoverageCompileCheck.psm1`** — second
  pass at the `Blocked: (root) — : no such file or directory` failure that
  surfaced during parallel `./run.ps1 -tc` runs. Even with the prior multi-root
  probe, a failing `go list ./...` (e.g. when the upstream `core-v9`
  module-loader error fires) could leak stderr fragments past the regex filter
  as a single-package "phantom (root)" entry, which then aborted the whole run
  before any coverage % was reported. Hardening:
  1. Discovery now captures `$LASTEXITCODE` from `go list`, anchors the
     keep-filter to the actual `module` declared in `go.mod`, rejects lines
     containing whitespace / `...` / known stderr prefixes (`go:`, `warning:`,
     `matched no packages`, `package`, `can't load`, `cannot find`, `err:`,
     `# `), and aborts loudly with the raw `go list` output if nothing valid
     survives — instead of silently producing one bogus package path.
  2. `Get-PackageShortName` replaces the bare `'.*(integratedtests|creationtests)/?'`
     regex. It always returns a non-empty label (trailing test segment → last
     path segment → full path) so blocked-package summaries never collapse to
     the unhelpful `(root)`. Blocked log lines now include the full import path
     next to the short name.
  3. All three early-abort code paths in `Invoke-TestCoverage` now call
     `Write-PhaseSummaryBox` before returning, so the dashboard summary (and
     any `Coverage Run: fail` phase status) renders even when the pipeline
     aborts before the merge step. This is what was hiding the coverage
     percentage — the run never reached `Write-CoverageConsoleSummary` because
     the aborted summary box was suppressed.

### Previously
- **`scripts/CoverageRunner.psm1`** — pre-coverage compile check no longer
  aborts with `Blocked: : no such file or directory`. Two root causes:
  1. Test discovery was hard-coded to `./tests/integratedtests/...`, which
     does not exist in this repo (tests live under `./tests/creationtests/`).
     `go list` errors were being coerced into a single empty package path
     and handed to `go test`. The probe now tries both directory names,
     skips paths not present on disk, and filters `go list` warning/error
     lines so only valid import paths reach the compile check.
  2. The in-package-test scan hard-coded the `core-v9` module prefix when
     stripping import paths to filesystem-relative paths. It now reads the
     `module` line from `go.mod` so the same script works in `enum-v5`,
     `core-v9`, and any other module.
- **`scripts/CoverageRunner.psm1`, `scripts/CoverageCompileCheck.psm1`** —
  `shortName` regex updated from `'.*integratedtests/?'` to
  `'.*(integratedtests|creationtests)/?'` so blocked-package labels render
  the trailing path segment in either layout.

### Added
- **CI/CD pipeline** (`.github/workflows/ci.yml`) — SHA dedup, `golangci-lint`,
  `govulncheck`, 4-suite parallel test matrix, aggregated test summary,
  60% coverage gate, and `go build ./...` smoke check.
- **Weekly vulnerability scan** (`.github/workflows/vulncheck.yml`) — scheduled
  `govulncheck` run with two-tier classification (fail on third-party,
  warn on stdlib).
- **Release workflow** (`.github/workflows/release.yml`) — triggers on
  `release/**` branches and `v*` tags; produces source archives, SHA-256
  checksums, and a GitHub Release whose body is extracted from this file.
- **Reusable CI guards** (`.github/workflows/ci-guards.yml`):
  - `scripts/ci/check-collisions.py` — per-package identifier collision
    audit (cross-file, case-insensitive, intra-file). GOOS/GOARCH build-tag
    siblings and Exported/unexported accessor pairs are recognised and
    excluded from false positives.
  - `scripts/ci/lint-baseline-diff.py` — lint gate that fails only on
    **new** `golangci-lint` findings; baseline cached per `main` push and
    seeded from `.ci-baselines/golangci-lint.json` on cold cache.
- **Spec docs**: `spec/04-tooling/00-overview.md` (tooling index/landing
  page), `02-release-pipeline.md`, `03-vulnerability-scanning.md`,
  `04-bootstrap-into-new-repo.md`, `04-ci-guards.md`,
  `05-branch-protection.md`, `06-cross-repo-sync.md`.
- **`CONTRIBUTING.md`** — local dev (`./run.sh`), commit conventions,
  release procedure.
- `.golangci.yml`, `CODEOWNERS`, `.github/PULL_REQUEST_TEMPLATE.md`
  (now with structured local / CI-guard / cross-repo checklists),
  `.gitattributes`.
- **Dependabot** (`.github/dependabot.yml`) — weekly `gomod` and
  `github-actions` updates, grouped minor/patch PRs, scheduled Mondays
  09:00 Asia/Kuala_Lumpur.
- **Cross-repo staging** under `cross-repo/core-v9/` — mirrored
  workflows, `.golangci.yml`, `dependabot.yml`, baselines, and a
  README install guide so the `core-v9` upstream can adopt the same
  CI surface (governed by `spec/04-tooling/06-cross-repo-sync.md`).
- **Python regression tests** for the CI guards:
  `scripts/ci/test_check_collisions.py` (22 cases covering build-tag
  collapsing, accessor pairs, decl parsing, string/comment skipping,
  per-package scoping) and `scripts/ci/test_lint_baseline_diff.py`
  (15 cases covering load/identity rules, seeding mode, gate mode,
  exit codes, summary counts).

### Changed
- Module path migrated from `gitlab.com/auk-go/enum` to
  `github.com/alimtvnetwork/enum-v5`.
- **Core dependency renamed** `github.com/alimtvnetwork/core-v9` →
  `github.com/alimtvnetwork/core-v9` across all 307 source files
  (`go.mod`, all package imports, spec docs, CI configs, coverage
  scripts, PR template). The `cross-repo/core-v9/` staging directory
  name is intentionally retained — it tracks the upstream repo name,
  not the module path. Pseudo-version pin
  `v1.5.6-0.20260423064907-72bcd64c06b5` carries over unchanged.
- Dependency `gitlab.com/auk-go/core` replaced with
  `github.com/alimtvnetwork/core-v9`, pinned to pseudo-version
  `v0.0.0-20260423064907-72bcd64c06b5` (commit `72bcd64` on
  `feature/1.5.6`) so CI can resolve the module deterministically.
- **`go.mod` pseudo-version downgraded** from
  `v1.5.6-0.<date>-<sha>` to `v0.0.0-<date>-<sha>`. The `v1.5.6-0.`
  form requires a preceding `v1.5.5` tag on the upstream `core-v9`
  repo, which doesn't exist (the v8 repo had v1.5.5; v9 was just
  renamed and has no tags yet). The `v0.0.0-` form has no predecessor
  requirement. Re-pin to a real `vX.Y.Z` tag once `core-v9` upstream
  ships its first tagged release.
- **`go.mod` rename-bridge `replace` directive** — the upstream
  `core-v9` repo's own `go.mod` still declares
  `module github.com/alimtvnetwork/core-v9`; Go enforces import-path
  / module-path equality so the v9 path can't load directly. Until
  upstream commits its `module github.com/alimtvnetwork/core-v9`
  line, `replace github.com/alimtvnetwork/core-v9 =>
  github.com/alimtvnetwork/core-v9 v0.0.0-<date>-<sha>` resolves the
  v9 import path to the v8 artifact at the same pinned commit. All
  source-code imports stay on `core-v9`; only the resolution target
  is bridged. Remove the `replace` once upstream's `go.mod` is fixed.

### CI
- `ci-guards.yml` gained a `python-tests` job that runs all
  `scripts/ci/test_*.py` via `unittest discover`. The existing
  `collision-audit` and `lint-baseline-diff` jobs now `needs:
  python-tests` so a broken gate script fails fast before the slower
  jobs spend CI minutes producing meaningless results.
- `scripts/CoveragePreChecks.psm1` — auto-fixer and bracecheck steps
  now skip gracefully (with `Register-Phase ... "skip"`) when
  `scripts/autofix/` or `scripts/bracecheck/` are absent from the
  repo, instead of hard-failing the entire `./run.ps1 -tc` run.
- **`scripts/bracecheck/`** (NEW Go tool, ~210 lines + README) —
  fast syntax pre-check. Lexical brace/bracket/paren balance
  validation (skips strings, runes, comments) plus a full
  `parser.AllErrors` pass over every `.go` file. Reports issues as
  `<relpath>:<line>:<col>: <message>`. Verified clean on 637 files.
- **`scripts/autofix/`** (NEW Go tool, ~165 lines + README) —
  conservative auto-fixer. Trims trailing whitespace, collapses 3+
  blank lines to 2, ensures one trailing newline, runs
  `format.Source`. Idempotent. Files that don't parse are skipped
  with a warning so bracecheck pinpoints the syntax issue. Supports
  `--dry-run`. With both tools restored, `./run.ps1 -tc` no longer
  prints the "scripts/autofix/ not present" skip notice.
- **`.github/workflows/python-tests.yml`** (NEW) — standalone runner
  for the CI-guard Python tests, triggered on `v*` tags, manual
  dispatch, and `scripts/ci/**` changes. Matrix tests across Python
  3.10/3.11/3.12. Complements the in-line `python-tests` job in
  `ci-guards.yml` by also catching releases and long-lived branches.

### Docs
- `CONTRIBUTING.md` — pre-push checklist rewritten as checkboxes
  mirroring `.github/PULL_REQUEST_TEMPLATE.md`; Spec References now
  links to `spec/04-tooling/00-overview.md` plus all six tooling
  spec files (00–06).
- `.ci-baselines/README.md` — fully documents the seed-then-gate
  workflow: seeding mode (warnings, never blocks), gating mode
  (NEW/FIXED/UNCHANGED diff with exit codes), mode-transition
  commands, and reviewer guidance for baseline drift.

