# Suggestions Tracker

> Single file tracking all Lovable suggestions. Update in-place — do not create per-suggestion files.
> When a suggestion is completed, move it to the **Completed** section with date and notes.

---

## Convention

- **suggestionId**: `S-<NNN>` (sequential)
- **status**: `open` | `in-progress` | `done` | `rejected`
- When done: move to Completed section, add completion date and notes
- When rejected: move to Rejected section, add reason

---

## Open Suggestions

### S-002: Promote `errcore.VarTwoNoType` from ❓ to ✅ in Cycle 6

- **createdAt:** 2026-05-05
- **source:** Lovable (Cycle 13, §15 audit)
- **affectedProject:** enum-v4 spec
- **description:** `VarTwoNoType` was scored ❓ but IS cross-referenced from multiple spec files.
- **rationale:** Under spec-internal-consistency dimension it qualifies as ✅.
- **proposed change:** Backport promotion when Task AC runs.
- **acceptance criteria:** Cycle 6 audit report row 16 updated.
- **status:** open (deferred — folded into Task AC consistency-dimension re-audit; will land with the §15 sweep after the freeze lifts)

### S-111: Surface the GoConvey-only sub-pattern in spec/06

- **createdAt:** 2026-05-06
- **source:** Lovable (Cycle 37 — D-CVS-64 carry-forward)
- **affectedProject:** enum-v4 spec
- **description:** `spec/06-testing-guidelines/02-test-case-types.md` and `05-assertion-patterns.md` present the `coretests`-framework path (`CaseV1`/`args.Map`/`BaseTestCase`/custom `ShouldBeEqualMap`) as the only option. `enum-v4`'s own `tests/creationtests/` package is a worked example of a **simpler GoConvey-only sub-pattern** (plain `So(actual, ShouldEqual, expected)` + AAA comments + plain `[]*Wrapper` / `map[K]V` registries, no `args.*` bundling, no `BaseTestCase` extension) that the spec doesn't acknowledge.
- **rationale:** Documenting the alternative makes the spec honest about consumer freedom and lets readers pick the lighter path when they don't need `args.Map` argument bundling.
- **proposed change:** Add a top-of-file note to both files cross-linking `enum-v4/tests/creationtests/` (e.g. `AllEnums_ContractsTesting_test.go` for the diff-pattern). Map the diff-based assertion claim onto `enumimpl.DynamicMap.LogShouldDiffMessage(...) + So(diff, ShouldBeEmpty)` as the local-only equivalent.
- **acceptance criteria:** Both files include the GoConvey-only callout. Cross-ref to `tests/creationtests/AllEnums_ContractsTesting_test.go` resolves.
- **status:** open
- **risk:** Low — additive cosmetic note; falls under Task AC consistency-dimension re-audit.

---

## Completed Suggestions

### S-109: Cycle-15 deep-probe of `tests/creationtests/` patterns to clear 21 ❓

- **completed:** 2026-05-06 (Cycle 37)
- **source:** Lovable (Cycle 27 — carry-forward)
- **resolution:** Read all 14 files under `enum-v4/tests/creationtests/` and ran the symbol-set probe `rg -n 'coretests\.|args\.|results\.|CaseV1|CaseNilSafe|GenericGherkins|GetAssert|ShouldBeEqualMap|ShouldBeSafe|InvokeWithPanicRecovery|BaseTestCase' tests/creationtests/` → **zero hits**. Confirmed `enum-v4` deliberately does NOT consume the upstream `coretests`/`args`/`results` framework; instead uses GoConvey + 2 local wrapper structs + module-level slice/map registries + AAA comments. Settled the 10 Cycle-15 ❓ items: **1 promoted ✅** (claim 20, diff-based assertion pattern via `enumimpl.DynamicMap.LogShouldDiffMessage` + `So(diff, ShouldBeEmpty)` in `AllEnums_ContractsTesting_test.go`), **9 annotated ⓘ "upstream-only"** (no `enum-v4` evidence available; remain blocked by Task AB for upstream-clone promotion). Cycle-15 verifiable subset 22/22 → 23/23 (still 100%); spec/06 unknown ❓ pool **10 → 0**. New finding D-CVS-64 (LOW) carried forward as **S-111**. Audit file: `spec/07-code-vs-spec-audits/29-cycle37-S109-creationtests-deep-probe.md`. Spec changelog → spec-v0.44.0. `package.json` 0.6.0 → 0.7.0.
- **acceptance criteria:** ✅ New audit cycle entry on scoreboard (`Cycle 37`). ✅ Cycle-15 ❓ pool reduced from 10 → 0 unknown (1 promoted, 9 annotated upstream-only). ✅ Direct source evidence cited for every promotion/annotation.

### S-004: Fix `spec/00-llm-integration-guide.md` stale test path reference

- **completed:** 2026-05-06 (Cycle 34)
- **source:** Lovable (reliability report, 2026-05-05)
- **resolution:** Re-framing applied (not path replacement). The `tests/integratedtests/` references on lines 36 (Decision Matrix Style D row) and 826 (Test Folder Structure section) correctly document the **upstream `core-v9` consumer layout**; `enum-v4` itself uses the single shared `tests/creationtests/` package. Added an inline scope callout above the line-826 code fence that explains the upstream-vs-enum-v4 split and cross-links both `spec/01-app/13-testing-patterns.md` §6.1 and `spec/06-testing-guidelines/01-folder-structure.md` (which already carries the same disclaimer on its own line 3). Pattern matches the resolution model used in Cycles 12/15/17/18 for sibling `integratedtests` references.
- **acceptance criteria:** ✅ `tests/integratedtests/` references in `spec/00-llm-integration-guide.md` are now scope-disambiguated. ✅ No misdirection risk for AI readers — the callout lands before the code block.

### S-003: Fix `spec/06-testing-guidelines/01-folder-structure.md` stale path

- **completed:** 2026-05-06 (Cycle 34) — already-resolved discovery
- **source:** Lovable (reliability report, 2026-05-05)
- **resolution:** Verified obsolete. `spec/06-testing-guidelines/01-folder-structure.md` line 3 already carries the upstream-scope disclaimer added in earlier audit cycles: *"⚠️ Scope: the layout below describes upstream `core-v9`. enum-v4 uses a single shared `tests/creationtests/` package — see [spec/01-app/13-testing-patterns.md §6.1]…"*. The `integratedtests` references in the file body intentionally document the upstream consumer layout, not enum-v4. Closing as **resolved-by-prior-cycle** with no edit needed.
- **acceptance criteria:** ✅ Scope disclaimer present (line 3). ✅ No misdirection risk.

### S-001: Pin Go toolchain to 1.22 as stopgap for Task W

- **completed:** 2026-05-06 (Cycle 34) — closed as **obsolete**
- **source:** Lovable (Cycle 13, 2026-05-05)
- **resolution:** Stopgap no longer relevant. The underlying problem (Go 1.25 rejecting the dual-path `replace` bridge to upstream `core-v9`) was fixed at the source by Tasks W + AG (2026-05-05): upstream `core-v9` was renamed to `module github.com/alimtvnetwork/core-v9` and tagged `v1.5.8`; `enum-v4/go.mod` now pins `core-v9 v1.5.8` directly with no `replace` bridge (per Core memory). Pinning to Go 1.22 would now mask the working modern setup. **Do not** add `toolchain go1.22.0` to `go.mod`.
- **acceptance criteria:** ✅ `replace` bridge removed (Task AG). ✅ `enum-v4` builds against current toolchain without the stopgap. ✅ Risk noted in original suggestion (toolchain lock-in) avoided.

### S-107: Goconvey-failure summarizer for `failing-tests.txt`

- **completed:** 2026-05-06 (Cycle 33)
- **source:** Lovable (carry-forward from Cycle 24 — `TestLogWriter.psm1` Pass-2 collector buries goconvey failures under `N total assertions` lines)
- **resolution:** Added `Get-GoconveyFailureSummary` to `scripts/TestLogWriter.psm1` (also exported). For each failed test's captured block, walks the lines and pairs every `Expected:` with its nearest following `Actual:` / `(Line N)` / `Message:` (8-line look-ahead window, stops when the next `Expected:` starts). `Write-TestLogs` now prepends a `── Failure summary (N) ──` mini-section to each `FAIL:` block in `failing-tests.txt`, listing each triplet as `#i Expected: ... | Actual: ... (Line N) [Message]` BEFORE the noisy raw block. The original raw block is preserved underneath for full context. Smoke-tested via `nix run nixpkgs#powershell` on a synthetic 2-failure goconvey block: extracted both triplets correctly (incl. optional Message field), empty blocks return 0, non-goconvey blocks (e.g. plain `panic:`) return 0 with no false positives. `package.json` 0.2.0 → 0.3.0.
- **acceptance criteria:** ✅ Failed-test blocks now lead with the goconvey Expected/Actual triplets. ✅ Existing raw-block content is preserved (no information loss). ✅ Helper handles empty + non-goconvey blocks safely.

### S-110: Restore standalone coverage utilities documented alongside S-108

- **completed:** 2026-05-06 (Cycle 32)
- **source:** Lovable (Cycle 31 follow-up to S-108 completion)
- **resolution:** Created the three standalone utilities listed in `spec/03-powershell-test-run/06-coverage-prompt-generator.md`:
  - `scripts/coverage/Get-UncoveredLines.ps1` — emits a single-row object `{SourceFile, UncoveredCount, Ranges}` with collapsed line ranges (e.g. `L15-L17, L22`).
  - `scripts/coverage/Get-FunctionCoverage.ps1` — parses `go tool cover -func` output, filters below `-Threshold` (default 100.0), returns objects sorted ascending by coverage.
  - `scripts/coverage/Get-PackageCoverageReport.ps1` — combines both for a single `-Package`; supports `-Format text|markdown|json` and optional `-OutputFile`.
  All three reuse the same regexes and range-collapse helper as the S-108 main script and avoid the `$Input` automatic-variable shadowing pitfall (use `$Source`). Smoke-tested via `nix run nixpkgs#powershell` against synthetic 6-block `coverage.out` + 3-line func output: each utility produces exactly the spec-shaped output (correct sort order, correct ranges, all 3 formats). `package.json` 0.1.0 → 0.2.0.
- **acceptance criteria:** ✅ All four scripts now exist in `scripts/coverage/`. ✅ Spec drift D-CVS-62 is **fully** closed (was partially closed by S-108).

### S-108: Restore missing `scripts/coverage/Generate-CoveragePrompts.ps1`

- **completed:** 2026-05-06 (Cycle 31)
- **source:** Lovable (Cycle 27 — AB scripts deep-probe, surfaced D-CVS-62)
- **resolution:** Created `scripts/coverage/Generate-CoveragePrompts.ps1` per the contract in `spec/03-powershell-test-run/06-coverage-prompt-generator.md`. Parses `coverage.out` (statement blocks with `count==0`) + `go tool cover -func` output, sorts functions ascending by coverage, batches at `-BatchSize` (default 500) into `data/prompts/coverage-prompt-N.txt`, emits `prompts-summary.json`. Smoke-tested end-to-end via `nix run nixpkgs#powershell` against synthetic `coverage.out` + func output: produced exactly the spec sample format including range-collapsed uncovered-line ranges (`L15-L17, L22`) and ascending coverage sort. Both call-sites in `scripts/CoverageRunner.psm1:313-316` and `scripts/PackageCoverage.psm1:150` now resolve. Standalone utilities (`Get-UncoveredLines.ps1` etc.) tracked separately as **S-110**.
- **acceptance criteria:** ✅ `Test-Path scripts/coverage/Generate-CoveragePrompts.ps1` → True. ✅ Smoke run produces `coverage-prompt-1.txt` matching spec format and `prompts-summary.json`.

### S-100: Add `cmd/main/` smoke-test policy carve-out to spec §12

- **completedAt:** 2026-05-05 (Cycle 10, fix C-CVS-10)
- **notes:** Spec §12 rewritten as "library-first, smoke-test allowed" policy.

### S-101: Rewrite §06 around `SimpleSlice`/`Hashset`/`SimpleStringOnce`

- **completedAt:** 2026-05-04 (Cycle 4, fixes D-CVS-22/23/24)
- **notes:** Replaced fictional `coreonce.New.String` with actual constructors.

### S-102: Add consumer-coverage callouts for upstream-only API

- **completedAt:** 2026-05-04 → 2026-05-05 (D-CVS-25, D-CVS-38, D-CVS-42)
- **notes:** Three sections now carry explicit upstream-only callouts.

### S-103: Extract portable runner specs into `spec/03-portable/` sub-directory

- **completed:** 2026-05-06 (Cycle 36)
- **source:** Lovable (Cycle 16 carry-forward)
- **resolution:** Moved `spec/03-powershell-test-run/08-generic-go-test-coverage-runner.md` → `spec/03-powershell-test-run/portable/01-generic-go-test-coverage-runner.md` and `09-ai-agent-complete-reference.md` → `portable/02-ai-agent-complete-reference.md`. Added `spec/03-powershell-test-run/portable/README.md` with the scope-split explanation, the 2-file table, and 3 editor rules. Updated 2 live cross-refs (`spec/00-llm-integration-guide.md:2380`, `spec/04-tooling/03-powershell-implementation.md:456`) and 1 internal cross-ref inside the moved `02-ai-agent-complete-reference.md`. Historical references in `spec/CHANGELOG.md`, `spec/07-code-vs-spec-audits/17-cycle16-powershell-test-run.md`, and `spec/99-audits/01-original-11-step-plan.md` left as-is (audit history). Spec changelog → spec-v0.43.0. `package.json` 0.5.0 → 0.6.0.
- **acceptance criteria:** ✅ `rg -n 'spec/03-powershell-test-run/(08|09)-' spec/ --glob '!spec/CHANGELOG.md' --glob '!spec/07-code-vs-spec-audits/**' --glob '!spec/99-audits/**'` returns zero hits. ✅ `spec/03-powershell-test-run/portable/README.md` exists and documents the split.

### S-104: Add `cross-repo/core-v8/README.md` historical-naming top-of-file note

- **completed:** 2026-05-06 (Cycle 35)
- **source:** Lovable (Cycle 17 carry-forward)
- **resolution:** Added a prominent top-of-file callout to `cross-repo/core-v8/README.md` covering: (1) the directory name is historical and intentional (mirrors a separate upstream repo), (2) the actual import path used by `enum-v4` source is `github.com/alimtvnetwork/core-v9` (renamed 2026-05-05, tagged `v1.5.8`), (3) editors must always write `cross-repo/core-v8/` even when the surrounding sentence is about `core-v9` content, (4) the historical `enum-v1` / `core-v8` body references must NOT be rewritten (Core-memory rule). Body untouched. Closes the Cycle-17 root cause of D-CVS-49/52/53/55 at the point of truth.
- **acceptance criteria:** ✅ Head section includes the 4-point explanation. ✅ Body content preserved.

### S-105: CI guard — `spec/02-app-issues/` index-drift detector

- **completed:** 2026-05-06 (Cycle 35)
- **source:** Lovable (Cycle 18 carry-forward)
- **resolution:** Added `scripts/ci/check-issues-index-drift.py` + `scripts/ci/test_check_issues_index_drift.py` (5 unittest cases, all pass) and a new `issues-index-drift` job in `.github/workflows/ci-guards.yml` (depends on `python-tests`, mirrors the `collision-audit` job). Script extracts every `| NN |`-prefixed row from `00-issues-index.md` (canonical) and `README.md` (human), compares **row count AND id-set**, exits 1 with `Missing from README` / `Missing from index` diffs on drift, exits 2 on missing files. Live repo reports `OK: spec/02-app-issues index in sync (9 rows).` — adopted at clean baseline.
- **acceptance criteria:** ✅ Removing a row from README → CI fails (`test_row_count_mismatch_fails`). ✅ Both in sync → CI passes (`test_in_sync_passes` + live-repo smoke). ✅ Same count, mismatched ids → CI fails (`test_id_set_mismatch_same_count_fails`).

### S-106: `spec-api-check.psm1` — automate code-fence vs API verification

- **createdAt:** 2026-05-06
- **source:** Lovable (Cycle 19 carry-forward)
- **affectedProject:** enum-v4
- **description:** Cycle 19 (AB pass 1) found 5 ❌ contradictions in `spec/01-app/09-converters.md` whose root cause is identical: **the spec was authored against an internal mental model of the API, not against grep output**. C-CVS-15 in particular fabricated an entire `typesconv` numeric-widening surface that does not exist.
- **rationale:** The same author pattern is almost certainly present in the other §07/§08/§10/§11/§15/§16 ❓ pools (124 ❓ remaining). A one-shot lint can prevent the next 50+ ❌s from ever shipping.
- **proposed change:** Add `scripts/spec-api-check.psm1` that, for every Go code-fenced block in `spec/01-app/`, extracts `<pkg>.<symbol>` references and runs `go vet` (or `go list -f '{{.Imports}}'`) against a synthesised stub `.go` file in a tmp module rooted at `core-v9 v1.5.8`. Failures should print the spec file + line.
- **acceptance criteria:** Removing `Integer` from `converters.stringTo` makes the lint fail; current spec text after AJ-01..03 land makes it pass; CI-friendly.
- **status:** open
- **risk:** Tooling addition only — no change to spec or production code.

---

## Rejected Suggestions

_(none yet)_
