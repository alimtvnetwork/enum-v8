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

### S-001: Pin Go toolchain to 1.22 as stopgap for Task W

- **createdAt:** 2026-05-05
- **source:** Lovable (Cycle 13)
- **affectedProject:** enum-v4
- **description:** Go 1.25 rejects the dual-path `replace` bridge. Pinning to Go 1.22 would unblock builds.
- **rationale:** Allows development to continue while waiting for upstream `core-v9` `go.mod` rename.
- **proposed change:** Add `toolchain go1.22.0` to `go.mod`.
- **acceptance criteria:** `./run.ps1 -tc` passes with Go 1.22 toolchain.
- **status:** open
- **risk:** Masks the underlying issue; locks to older toolchain.

### S-002: Promote `errcore.VarTwoNoType` from ❓ to ✅ in Cycle 6

- **createdAt:** 2026-05-05
- **source:** Lovable (Cycle 13, §15 audit)
- **affectedProject:** enum-v4 spec
- **description:** `VarTwoNoType` was scored ❓ but IS cross-referenced from multiple spec files.
- **rationale:** Under spec-internal-consistency dimension it qualifies as ✅.
- **proposed change:** Backport promotion when Task AC runs.
- **acceptance criteria:** Cycle 6 audit report row 16 updated.
- **status:** open

### S-003: Fix `spec/06-testing-guidelines/01-folder-structure.md` stale path

- **createdAt:** 2026-05-05
- **source:** Lovable (reliability report)
- **affectedProject:** enum-v4 spec
- **description:** Line 13 references `tests/integratedtests/` which doesn't exist. Should be `tests/creationtests/`.
- **rationale:** This is the #1 failure risk for any AI following the spec to write tests.
- **proposed change:** Replace `integratedtests` with `creationtests` throughout the file.
- **acceptance criteria:** `rg integratedtests spec/06-testing-guidelines/01-folder-structure.md` returns 0 hits.
- **status:** open

### S-004: Fix `spec/00-llm-integration-guide.md` stale test path reference

- **createdAt:** 2026-05-05
- **source:** Lovable (reliability report)
- **affectedProject:** enum-v4 spec
- **description:** Line 36 references `tests/integratedtests/` in the decision matrix table.
- **rationale:** First file any AI reads; stale path causes immediate misdirection.
- **proposed change:** Replace stale reference with `tests/creationtests/`.
- **acceptance criteria:** `rg integratedtests spec/00-llm-integration-guide.md` returns only the anti-pattern callout lines.
- **status:** open

### S-109: Cycle-15 deep-probe of `tests/creationtests/` patterns to clear 21 ❓

- **createdAt:** 2026-05-06
- **source:** Lovable (Cycle 27 — carry-forward)
- **affectedProject:** enum-v4 spec
- **description:** 21 ❓ remain in `spec/06-testing-guidelines/` (Cycle 15 baseline). They're behavioural/observational claims about the Goconvey + `EnumTestWrapper` registry pattern that need a fresh probe of `tests/creationtests/` source — distinct probe technique from Cycle 27's grep-the-script approach.
- **rationale:** Last large pool of unresolved ❓ outside the AJ/AC backlog.
- **proposed change:** Run a dedicated cycle: read `tests/creationtests/EnumTestWrapper.go`, `tests/creationtests/all*.go`, and the `Test_AllEnums_*` registry; promote each Cycle-15 ❓ against direct source evidence.
- **acceptance criteria:** Cycle-15 audit file shows ≤ 5 ❓ remaining; new cycle entry on the scoreboard.
- **status:** open

---

## Completed Suggestions

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

- **createdAt:** 2026-05-06
- **source:** Lovable (Cycle 16 carry-forward)
- **affectedProject:** enum-v4 (spec only)
- **description:** `spec/03-powershell-test-run/08-generic-go-test-coverage-runner.md` and `09-ai-agent-complete-reference.md` are explicitly portable ("any Go module / repository", "self-contained reference for any AI agent"). They currently sit alongside `enum-v4`-specific runner files and required top-of-file consumer-coverage callouts in Cycle 16 (D-CVS-47, D-CVS-48) to disambiguate scope.
- **rationale:** A structural split (sub-directory) is more discoverable than a notational one (callouts) and would let future portable-runner edits ship without touching the `enum-v4`-specific files.
- **proposed change:** Move `08-` and `09-` into a new `spec/03-powershell-test-run/portable/` sub-directory; renumber to `01-`/`02-` inside it; update cross-refs in `spec/04-tooling/` and `spec/06-testing-guidelines/` if any.
- **acceptance criteria:** `rg -nc 'spec/03-powershell-test-run/(08|09)-' spec/` returns zero hits after the move; new portable sub-directory has its own README explaining the scope split.
- **status:** open
- **risk:** Low — structural reorganisation, no normative content change.

### S-104: Add `cross-repo/core-v8/README.md` historical-naming top-of-file note

- **createdAt:** 2026-05-06
- **source:** Lovable (Cycle 17 carry-forward)
- **affectedProject:** enum-v4
- **description:** Cycle 17 surfaced 5 broken `cross-repo/core-v9/` paths in spec text (D-CVS-49, -52, -53, -55) — all stemming from the same root cause: readers who know the import path is `core-v9` instinctively type the wrong directory name (the actual directory keeps its historical `core-v8` name per Core memory). A prominent note inside `cross-repo/core-v8/README.md` would prevent future authors from making the same mistake.
- **rationale:** Reduces future drift class entirely (point-of-truth fix vs. per-cite-site clarification).
- **proposed change:** Add a top-of-file callout to `cross-repo/core-v8/README.md` explaining: (1) the directory name is historical, (2) the actual import path is `github.com/alimtvnetwork/core-v9`, (3) anyone editing spec text should write `cross-repo/core-v8/` even when discussing `core-v9` content.
- **acceptance criteria:** `cross-repo/core-v8/README.md` head section includes the explanation; future audit cycles can cite it instead of repeating the Core-memory note inline.
- **status:** open
- **risk:** None — informational README edit only.

### S-105: CI guard — `spec/02-app-issues/README.md` index drift detector

- **createdAt:** 2026-05-06
- **source:** Lovable (Cycle 18 carry-forward)
- **affectedProject:** enum-v4
- **description:** Cycle 18 found `spec/02-app-issues/README.md` was stale by 4 issues for ~14 days. A trivial CI guard comparing the row count of `00-issues-index.md` (canonical) vs `README.md` would prevent recurrence.
- **proposed change:** Add a check to `.github/workflows/ci-guards.yml` (or `scripts/ci/`) that fails if the row counts differ.
- **acceptance criteria:** Removing a row from README CI fails; both in sync CI passes.
- **status:** open
- **risk:** None.

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
