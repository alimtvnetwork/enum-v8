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

### S-108: Restore missing `scripts/coverage/Generate-CoveragePrompts.ps1`

- **createdAt:** 2026-05-06
- **source:** Lovable (Cycle 27 — AB scripts deep-probe, surfaced D-CVS-62)
- **affectedProject:** enum-v4
- **description:** `scripts/CoverageRunner.psm1:313-316` and `scripts/PackageCoverage.psm1:150` both call `& $promptScript ... -BatchSize 500` against `scripts/coverage/Generate-CoveragePrompts.ps1`, but that file is **missing** from the repo. The call-sites guard with `if (Test-Path $promptScript)` so the missing-file case silently no-ops — meaning the documented "per-batch coverage prompt files capped at 500 functions" feature (`spec/03-powershell-test-run/06-coverage-prompt-generator.md`) does not run today.
- **rationale:** Either the spec promises a feature that doesn't exist (drift) OR the script was lost in a previous refactor. Since the call-sites are still wired, restoration is the lower-friction fix.
- **proposed change:** Restore the script (read `spec/03-powershell-test-run/06-coverage-prompt-generator.md` for the contract: input = cover.out + go-tool-cover func output; output = batched prompt files with `-BatchSize 500`; per-package grouping).
- **acceptance criteria:** `Test-Path scripts/coverage/Generate-CoveragePrompts.ps1` → True; `./run.ps1 -tc` produces files under `data/prompts/`.
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
