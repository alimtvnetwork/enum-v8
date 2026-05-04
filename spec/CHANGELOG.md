# Spec Changelog

> Tracks **versioned releases of the `spec/` deliverable** (this folder).
> Distinct from the `core-v8` Go module's own version (governed by `spec/01-app/11-versioning.md`).
>
> **Rule** (per `.lovable/user-preferences` line 8 + `spec/01-app/11-versioning.md` §3): every meaningful spec edit cycle bumps **at least minor**.
>
> Format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and [Semantic Versioning](https://semver.org/).

---

## [spec-v0.23.0] — 2026-05-04 (Cycle 3 baseline — §05 enum-system audited)

### Added
- **`spec/07-code-vs-spec-audits/04-cycle3-enum-system.md`** — full Cycle 3 audit of `01-app/05-enum-system.md`. 18 claims: 8 ✅ / 6 ⚠️ / **3 ❌** / 1 ❓. Verifiable score: **47.1 %** (8/17). Three real contradictions surfaced (C-CVS-03..05): the "first const = `Invalid`" rule is violated by 10 packages using alternate sentinel names (`Default`, `Unspecified`, `Uninitialized`, `InvalidIndex = -1`); the recipe imports `core-v9/internal/reflectinternal` which `enum-v1` cannot legally do across module boundaries; and `inttype.InvalidIndex Variant = -1` directly contradicts the "zero-value sentinel" wording.
- **6 new drift findings** (D-CVS-14..19): documented vs actual file layout (`Variant.go` everywhere, no `consts.go`), missing `*AllCases` factory family in §6 (10+ call sites), unused `CreateUsingMap` listed, stale `tests/integratedtests/` reference (mirrors C-CVS-01), unrunnable `reflectinternal.TypeName(...)` example, predicate file-split rule that's never enforced.

### Changed
- **`spec/07-code-vs-spec-audits/01-scoreboard.md`** — added Cycle 3 row, 9 new open findings, 3 new milestones; section progress **3/16**. ❌ contradiction count went from 0 → 3 (all on §05).

## [spec-v0.22.0] — 2026-05-04 (Cycle 2 closed — §04 at 100% verifiable)

### Fixed
- **`spec/01-app/04-error-system.md`** — resolves D-CVS-06..13 from Cycle 2.
  - **§1 table**: added rows for `MustBeEmpty`, `RawErrCollection`, `ToError` / `ToString`, `MessageWithRef`, `RangeNotMeet`.
  - **§1.1**: extended `RawErrorType` examples with `FailedToConvertType`, `NotSupportedType`, `PathInvalidErrorType`, `FailedToExecuteType`, `ComparatorShouldBeWithinRangeType`; added footnote pointing to upstream `errcore/RawErrorType.go` for the exhaustive enumeration.
  - **§1.2**: added `ErrorRefOnly(ref)` and `CombineWithAnother(other)` constructor rows; expanded code block with examples for both, `MustBeEmpty`, and the comparison vs `HandleErr`.
  - **New §1.5 "Reference Helpers"**: documents `MessageWithRef` and `RangeNotMeet` with `onofftype/vars.go`, `promptclitype/vars.go`, and `internal/messages/messages.go` references.
  - **New §1.6 "Accumulating Errors"**: documents `RawErrCollection` accumulator pattern with `osdetect/windowsSystemDetailGenerator_windows.go` reference.
  - **New §1.7 "Conversion Helpers"**: documents `ToString` / `ToError` with `osdetect/vars.go` reference.
  - **§7 "When to Use Which API"**: added 9 new rows covering all newly-documented APIs.

### Changed
- **`spec/07-code-vs-spec-audits/01-scoreboard.md`** — Cycle 2 closed: §04 verifiable score moved 27.3 → **100.0**. All 8 open drift findings (D-CVS-06..13) moved to "Resolved". 7 ❓ remain pending task **AB** (fetch upstream `core-v9` source).

## [spec-v0.21.0] — 2026-05-04 (Cycle 2 baseline — §04 error-system audited)

### Added
- **`spec/07-code-vs-spec-audits/03-cycle2-error-system.md`** — full Cycle 2 audit of `01-app/04-error-system.md`. 18 claims extracted: 3 ✅, 8 ⚠️ (spec is incomplete vs consumer usage), 7 ❓ (unverifiable without upstream `core-v9` source — sandbox lacks Go + the module is not vendored), 0 ❌. Verifiable-subset score: **27.3 %** (3/11). New ❓ bucket introduced so aspirational APIs aren't logged as contradictions before upstream source is available.
- **8 new open drift findings** (D-CVS-06..13) covering APIs `enum-v1` actively uses but the spec does not document: `MustBeEmpty`, `RawErrCollection`, `<RawErrorType>.ErrorRefOnly`, `<RawErrorType>.CombineWithAnother`, `MessageWithRef`, `RangeNotMeet`, `ToError`/`ToString`, plus 4 missing `RawErrorType` examples (`FailedToExecuteType`, `NotSupportedType`, `PathInvalidErrorType`, `ComparatorShouldBeWithinRangeType`).

### Changed
- **`spec/07-code-vs-spec-audits/01-scoreboard.md`** — added Cycle 2 row, new ❓ column, 8 open findings, 2 new milestones (apply D-CVS-06..13 spec fixes; fetch upstream `core-v9` source as task **AB**). Section progress: **2/16** done.

## [spec-v0.20.0] — 2026-05-04 (Cycle 1 fully closed — §03 at 100%)

### Fixed
- **`spec/01-app/03-import-conventions.md` §3 "`internal/` access from tests"** — resolves **C-CVS-02**. Section was titled "Common `internal/` packages used by tests" and used `core-v9/internal/reflectinternal` as a live example, but `enum-v1` (this repo) imports zero `internal/` packages. Reframed as a forward-looking explanation that `internal/` access is a **same-module** rule, with the `reflectinternal` example explicitly attributed to the upstream `core-v9` repo's own tests, and a new **consumer-side note** stating that `enum-v1`-style consumers cannot import `core-v9/internal/...` because Go enforces `internal/` at module boundaries.
- **`spec/01-app/03-import-conventions.md` §4 "Test-Package Imports"** — resolves **C-CVS-01**. Old text claimed tests live at `tests/integratedtests/footests/`, which doesn't exist in this repo. Replaced with the actual `enum-v1` layout (`tests/creationtests/` — flat, single `package creationtests`, mix of `*_test.go` and shared fixture `.go` files) plus a cross-reference to the upstream `core-v9` per-suite layout (`tests/<suite>/footests/`). The four shared rules (separate package, normal imports of source, no cycles, same-module `internal/` only) apply to both layouts.
- **`spec/07-code-vs-spec-audits/01-scoreboard.md`** — moved C-CVS-01 and C-CVS-02 from Open → Resolved; §03 score updated 83.3 → **100.0** (12/12). Open-findings list is now empty. Cycle 1 is closed.

### Verified
- `rg -n "integratedtests|core-v8" spec/01-app/03-import-conventions.md` → 0 hits.

---

## [spec-v0.19.0] — 2026-05-04 (Cycle 1 LOW-drift fixes applied)

### Fixed
- **`spec/01-app/03-import-conventions.md`** — applied 5 LOW-severity Cycle 1 findings:
  - **D-CVS-01** (line 4): `consumes core-v8 packages` → `consumes core-v9 packages`.
  - **D-CVS-02** (line 88): `path ends in core-v8` / `not corev8` → `core-v9` / `corev9`.
  - **D-CVS-03** (line 94): `For core-v8, this means:` → `For core-v9, this means:`.
  - **D-CVS-04** (line 121): reworded "rooted at the same `core-v8` module" to a module-generic statement that applies equally to `core-v9`, `enum-v1`, or any other consumer with its own `internal/` tree.
  - **D-CVS-05** (lines 61, 73): annotated `coredata/coregeneric` as `// optional` in the canonical import block and added a sentence noting "not every consumer uses every package — `enum-v1`, for example, currently uses 8 of the 11 listed canonical imports".
- **`spec/07-code-vs-spec-audits/01-scoreboard.md`** — moved D-CVS-01..05 from Open → Resolved; §03 score updated 41.7 → **83.3** (10/12). The two remaining `tests/integratedtests/` and `internal/reflectinternal` contradictions stay open pending a structural decision.

### Verified
- `rg -n "core-v8" spec/01-app/03-import-conventions.md` → 0 hits (clean).

---

## [spec-v0.18.0] — 2026-05-04 (open code-vs-spec audit phase)

### Added
- **`spec/07-code-vs-spec-audits/`** — NEW folder kicking off the implementation-vs-documentation drift audit phase. The spec-only cycle reached its ceiling (97.5–98.0 / 100, 0 open findings) at v0.17.1; this is the inverse: does the code match what the docs claim?
  - `00-index.md` — scope rules + methodology (extract claims → locate evidence → verify ✅/⚠️/❌ → file findings).
  - `01-scoreboard.md` — living drift scoreboard.
  - `02-cycle1-import-conventions.md` — first cycle, audited `01-app/03-import-conventions.md` (12 claims).

### Findings opened (Cycle 1)
- 5 LOW drifts (D-CVS-01..05): stale `core-v8` prose in §03 that didn't get rewritten during the v8→v9 module rename.
- 1 MED + 1 HIGH contradiction (C-CVS-01..02): §03 references a `tests/integratedtests/` directory that doesn't exist in this repo and an `internal/reflectinternal` import nobody uses.
- Initial measured drift score: **41.7%** (5/12 matches). After applying the 5 LOW spec fixes this jumps to ~91.7%.

---

## [spec-v0.17.1] — 2026-04-25 (close F-V16-01 — trust-boundary worked example)

### Added
- `spec/01-app/15-observability.md` §5.1 — End-to-end "trust-boundary handler" worked example (~75 lines including rationale table). Shows inbound untrusted form data → `corevalidator.New.Line` validation → PII-redacted struct copy → structured `slog` log call → generic error returned to caller. Cross-references `08-validators.md`, `16-security.md`, and the existing §5 logging rules.

### Resolved
- **F-V16-01** — Trust-boundary worked example missing in `15-observability.md` §5. The 5th-party mediocre-AI simulation (audit #10) flagged that a junior agent attempting an end-to-end "validate untrusted input + log with PII redaction" task had to assemble pieces from 3 files. Now unified as one canonical example. Projected score uplift: 0.3–0.5 → ~98.0.

### Status
- **0 open / 27 resolved findings.** Spec is at the 97–98 ceiling with no remaining bug-fix work. The natural next phase is **code-vs-spec** (lifts the spec-only directive).

---

## [spec-v0.17.0] — 2026-04-25 (5th-party mediocre-AI reproducibility audit)

### Added
- `spec/99-audits/10-ai-audit-2026-04-25-mediocre-sim.md` — 5th-party simulation audit (mediocre/junior-AI persona, the user's original benchmark).

### Measured
- **MEASURED score: 97.5 / 100** (5th-party mediocre-AI simulation, post spec-v0.16.0). Confirms the v0.16.0 projection range of 97.5–97.8. Score progression across 5 audits: 85.5 → 88.9 → 96.0 → **97.5** (converged inside the 97–98 ceiling).
- 4-task implementation simulation: 3 ✅ PASS + 1 ⚠️ PARTIAL.
- v0.16.0 expansion (`15-observability.md`, `16-security.md`) verified Adequate. Zero regressions.

### Opened (OPTIONAL — spec is at-ceiling regardless)
- **F-V16-01** — `spec/01-app/15-observability.md` §5: trust-boundary worked example missing. A junior agent attempting an end-to-end "validate untrusted input + log with PII redaction" task finds all the rules across `08-validators.md`, `15-observability.md`, `16-security.md` but no single 25-line example unifies them. Severity: very low. Uplift: 0.3–0.5 → ~98.0 if landed.

### Audit-doc updates
- `spec/99-audits/00-index.md` — current-quality block updated to MEASURED 97.5.
- `spec/99-audits/07-scoreboard.md` — 5th-party row added; targets table updated; decision block confirms cycle complete.

### Status
- **Spec-only cycle COMPLETE.** 5 audit passes across 3 distinct model families have converged. The natural next phase is **code-vs-spec** (lifts the spec-only directive).

---

## [spec-v0.16.0] — 2026-04-25 (close F-V14-05 — observability + security expansion)

### Added (new content — not a fix)
- **`spec/01-app/15-observability.md`** (~165 lines) — Diagnostic primitives (`VarTwo`, `VarTwoNoType`, `MessageVarMap`), stack-aware error wrapping (`StackEnhance`), test-failure output format pointer, logging boundaries, tracing/metrics scope, common mistakes. Unifies guidance previously scattered across `04-error-system.md` §5, `13-testing-patterns.md` §8, and `00-llm-integration-guide.md` §StackEnhance.
- **`spec/01-app/16-security.md`** (~175 lines) — Threat model, PII/secret handling, panic policy, allocation/resource safety, reflection safety, input-validation boundary, common mistakes. Unifies guidance previously scattered across `04-error-system.md`, `00-llm-integration-guide.md`, and `02-design-philosophy.md`.

### Edited
- `spec/01-app/README.md` — reading order extended to 17 entries; per-file status table appends rows 15 + 16; expansion note added.

### Resolved
- **F-V14-05** — observability/security split into dedicated spec files. Closes the last open audit finding.

### Audit-doc updates
- `spec/99-audits/09-ai-audit-2026-04-25-claude-sonnet-4.5.md` — banner now reads "ALL F-V14 findings RESOLVED".
- `spec/99-audits/07-scoreboard.md` — open register cleared (0 open / 26 resolved); resolved register includes F-V14-05.
- `spec/99-audits/00-index.md` — current-quality block: 26 resolved / 0 open.

### Score impact
- Projected score: 96.0 → **~97.5–97.8** (at the 97–98 ceiling). All 26 findings across 4 audit cycles now RESOLVED.
- A 5th-party measurement audit is OPTIONAL — convergence pattern shows diminishing returns past audit #09.

---

## [spec-v0.15.1] — 2026-04-25 (close 4 of 5 F-V14 findings)

### Fixed (documentation only — zero implementation changes)
- **F-V14-01** — `spec/01-app/08-validators.md` §2.1: added inline comment after the fluent-builder example explicitly stating `Build()` returns `*LineValidator` (pointer; satisfies `IsSuccessValidator`, `MessageGetter`, `ErrorGetter`). Notes that pointer storage is required because receiver methods are pointer-bound.
- **F-V14-02** — `spec/01-app/10-reflection-and-dynamic.md` §2.1: rewrote the `InvokeMethod` example to (a) state the full signature `InvokeMethod(target any, name string, args ...any) (any, error)` in a comment, (b) explicitly note "no single-return overload exists", and (c) use `_, _ = coredynamic.InvokeMethod(...)` for the discard case so the two-return shape is visible at first glance.
- **F-V14-03** — `spec/01-app/09-converters.md` §1.1: replaced the under-specified `"true"/"false"/"1"/"0"` note with the full `strconv.ParseBool` accepted-set (case-sensitive: `1`, `t`, `T`, `TRUE`, `true`, `True`, `0`, `f`, `F`, `FALSE`, `false`, `False`); explicit rejection of whitespace and `yes`/`no`/`on`/`off`.
- **F-V14-04** — `spec/01-app/12-cmd-entrypoints.md` §1: appended a paragraph to the "no `cmd/`" rule clarifying it is **PR-review-enforced** (no `go vet` lint, no CI check, no PowerShell guard exists). Pointer to `spec/02-app-issues/` for anyone wanting machine enforcement.

### Deferred
- **F-V14-05** — observability/security split into dedicated spec files. Genuinely new content (not a fix); deferred to **spec-v0.16.0+**.

### Audit-doc updates
- `spec/99-audits/09-ai-audit-2026-04-25-claude-sonnet-4.5.md` — banner marking F-V14-01..04 RESOLVED; F-V14-05 deferred.
- `spec/99-audits/07-scoreboard.md` — open register now lists only F-V14-05; resolved register expanded with the 4 new closures.
- `spec/99-audits/00-index.md` — current-quality block: 25 resolved / 1 open.

### Projected score impact
- 96.0 → **~97.5** (sum of per-finding uplifts: 0.4 + 0.3 + 0.3 + 0.2 = 1.2). At the GPT-5-stated ceiling of 97–98.
- A 5th-party measurement is **optional** — convergence pattern shows diminishing returns past audit #09.

---

## [spec-v0.15.0] — 2026-04-25 (4th-party confirmation audit by Claude Sonnet 4.5)

### Added
- `spec/99-audits/09-ai-audit-2026-04-25-claude-sonnet-4.5.md` — full 4th-party audit report.

### Measured
- **MEASURED score: 96.0 / 100** (Claude Sonnet 4.5, 4th-party, post spec-v0.14.0). Up from GPT-5's 88.9 pre-fix baseline. Just 1.5 points below the 97–98 ceiling.
- 9/9 of v0.14.0's F-V12 fixes verified Adequate. Zero regressions.

### Opened
- **F-V14-01** — `01-app/08-validators.md` §2.1: `LineValidator.Build()` return type unstated (low / S).
- **F-V14-02** — `01-app/10-reflection-and-dynamic.md` §2.1: `coredynamic.InvokeMethod` example discards both returns, contradicting `(any, error)` signature (low / S).
- **F-V14-03** — `01-app/09-converters.md` §1.1: `StringTo.Bool` accepted-input set under-specified (low / S).
- **F-V14-04** — `01-app/12-cmd-entrypoints.md` §1: `cmd/` rule has no documented enforcement mechanism (very low / S).
- **F-V14-05** — DEFERRED to v0.16.0+: no dedicated `15-observability.md` / `16-security.md` (very low / M).

### Audit-doc updates
- `spec/99-audits/00-index.md` — current-quality block updated to MEASURED 96.0.
- `spec/99-audits/07-scoreboard.md` — 4th-party row added; open register lists F-V14-01..05.

### Projected score impact (after F-V14-01..04 closed in v0.15.1)
- 96.0 → ~97.5 (at the GPT-5-stated ceiling).

---

## [spec-v0.14.0] — 2026-04-25 (close all 9 F-V12 findings)

### Fixed (documentation only — zero implementation changes)
- **F-V12-01** — `spec/00-llm-integration-guide.md` §Pattern 7: rewrote suffix-order grammar to the canonical 8-slot order **Base + Filter + Result + Type + Lock + If + Must + Ptr**. Added a worked right-vs-wrong diff example.
- **F-V12-02** — `spec/01-app/01-package-map.md` + `spec/01-app/00-repo-overview.md` §2: appended ⚠️ **LEGACY** label to `coremath` entries with cross-link to integration guide. (Closes the F-NEW-01 ⚠️ Partial gap GPT-5 surfaced.)
- **F-V12-03** — `spec/00-llm-integration-guide.md` §coregeneric: added "Iterator types — import & contract" box with explicit `import "iter"`, full Seq/Seq2 signatures for every container, and a compile-ready usage example.
- **F-V12-04** — `spec/00-llm-integration-guide.md` §Code Style Rules: refined the Constants row to permit direct `""` / `' '` comparisons in fast paths and predicate guards.
- **F-V12-05** — `spec/01-app/04-error-system.md` §1: added `HandleErr(err error)` to the public-API table with signature and behavior; added a runnable example showing the contract (no-op when nil, panics with stack-enhanced wrap when non-nil).
- **F-V12-06** — `spec/01-app/05-enum-system.md` Step 3: added an explicit spelling note distinguishing stdlib `UnmarshalJSON` (one `l`) from the enum engine's `UnmarshallEnumToValue` family (two `l`s) — this is a deliberate historical convention; do NOT "correct" it.
- **F-V12-07** — `spec/00-llm-integration-guide.md` §coregeneric: added "Return-type conventions" boxes to both `Hashset[T]` and `Hashmap[K,V]` clarifying why `*Lock` variants return the receiver while non-`Lock` `*Bool` variants return the boolean signal. Also clarified `Hashset.AddBool` returns `true` if the key ALREADY existed (counterintuitive — explicitly called out).
- **F-V12-08** — **False positive.** `spec/06-testing-guidelines/07-diagnostics-output-standards.md` already existed (78 lines, substantive content covering stack-trace rules, single/multi-line assertion format, alignment). It was simply missing from the audit bundle sent to GPT-5. Verified the link from `01-app/13-testing-patterns.md §8` resolves correctly. No file change needed.
- **F-V12-09** — `spec/01-app/05-enum-system.md` §2: added `BasicEnumValuer` to the Key Interfaces table with cross-reference to Step 3 / F-NEW-07.

### Audit-doc updates
- `spec/99-audits/08-ai-audit-2026-04-25-gpt5-confirm.md` — banner marking all 9 RESOLVED.
- `spec/99-audits/07-scoreboard.md` — open register cleared; resolved register expanded; current state = **0 open / 21 resolved**.
- `spec/99-audits/00-index.md` — current quality block updated.

### Projected score impact
- 88.9 → ~99.4 (sum of per-finding uplifts +10.5), **capped at GPT-5's stated practical ceiling ~97–98**.
- Awaits a 4th-party (Claude Opus 4.5) re-audit to *measure* the new score.

---

## [spec-v0.13.0] — 2026-04-25 (3rd-party confirmation audit)

### Added
- **`spec/99-audits/08-ai-audit-2026-04-25-gpt5-confirm.md`** — independent re-audit by GPT-5 (different model family from prior Gemini passes).

### Findings
- **Measured score: 88.9 / 100** (vs ~94.5 self-projected — overestimate corrected).
- **6 of 7 F-NEW fixes verified ✅ Adequate**; F-NEW-01 marked ⚠️ Partial because `01-app/01-package-map.md` and `01-app/00-repo-overview.md §2` still list `coremath` as active utility (the LEGACY label only landed in the integration guide).
- **9 NEW gaps opened** (F-V12-01..F-V12-09) — full register in `99-audits/07-scoreboard.md`.

### Worst axis
- **Consistency = 82/100** — `coremath` leak across files, suffix-grammar Result-slot omission, Unmarshal/Unmarshall spelling inconsistency.

### Updated
- `spec/99-audits/00-index.md` — table extended; current-quality block updated to 88.9 measured.
- `spec/99-audits/07-scoreboard.md` — score history extended; F-NEW marked RESOLVED with verification verdicts; F-V12 OPEN register added.

### Decision
- Owner directive still active: **spec-only cycle**. All 9 F-V12 fixes will be documentation edits — no Go code, no tests, no PowerShell.

---

## [spec-v0.12.0] — 2026-04-25 (close all 7 F-NEW findings)

### Fixed (documentation only — zero implementation changes)
- **F-NEW-01** — Marked `coremath` as LEGACY in `spec/00-llm-integration-guide.md`; instructs new code to use Go 1.21+ built-in `min` / `max` / `clear`.
- **F-NEW-02** — Added explicit *tokenization rule* + worked example to "Compound `*Or*` Naming in Filter Chains": `AOrB` is a single naming token occupying one slot in the suffix grammar.
- **F-NEW-03** — Added decision matrix at the top of `## coregeneric — Generic Data Structures API Reference` clarifying when to use `Collection[T]` (default, mutex-protected) vs `SimpleSlice[T]` (function-local, perf-sensitive).
- **F-NEW-04** — Elevated the `core` package name vs `core-v8` module path warning to top-level Module Identity blocks in both `spec/00-llm-integration-guide.md` and `spec/01-app/00-repo-overview.md`.
- **F-NEW-05** — Added "Serialization Asymmetry" subsection to `spec/01-app/05-enum-system.md` §7, documenting that `MarshalJSON` MUST emit string names while `UnmarshalJSON` MAY accept names/aliases/numeric strings.
- **F-NEW-06** — Added rationale block to "The CaseV1 Cast Idiom" in `spec/06-testing-guidelines/02-test-case-types.md` explaining why the cast keeps `BaseTestCase` decoupled from assertion libraries.
- **F-NEW-07** — Added comment + rationale to `spec/01-app/05-enum-system.md` §Step 3 explaining all `Value<Type>()` methods are required by `enuminf.BasicEnumValuer` and MUST NOT be deleted.

### Audit-doc updates
- `spec/99-audits/06-ai-audit-2026-04-25-gemini-2.5-pro.md` — banner added marking all 7 findings RESOLVED.
- `spec/99-audits/07-scoreboard.md` — open findings register cleared; resolved register expanded; projected score raised from 85.5 → ~94.5 pending an independent re-audit for confirmation.

### Projected score impact
- 85.5 → ~94.5 (+9.0 weighted points, matches Gemini 2.5 Pro's per-finding uplift sum).
- Subscore expected gains: Unambiguity 75 → ~88; Self-containment 80 → ~90; Consistency 80 → ~90; Completeness 85 → ~90.

---

## [spec-v0.11.0] — 2026-04-25 (audit folder reorganization + new audit)

### Added
- **`spec/99-audits/`** — new folder consolidating every spec-quality audit. All previous `spec/99-*.md` files renamed and moved with serialized prefixes:
  - `99-audit-report.md` → `99-audits/01-original-11-step-plan.md`
  - `99-step11-simulation.md` → `99-audits/02-step11-simulation-cycle1.md`
  - `99-step11-simulation-cycle2.md` → `99-audits/03-step11-simulation-cycle2.md`
  - `99-step11-simulation-cycle3.md` → `99-audits/04-step11-simulation-cycle3.md`
  - `99-ai-audit-2026-04-23.md` → `99-audits/05-ai-audit-2026-04-23-gemini.md`
  - `99-ai-audit-2026-04-25.md` → `99-audits/06-ai-audit-2026-04-25-gemini-2.5-pro.md`
- **`spec/99-audits/00-index.md`** — folder index with reading order, scope rules, and process for adding new audits.
- **`spec/99-audits/07-scoreboard.md`** — living scoreboard.

### Changed
- All cross-references throughout `spec/`, `readme.txt`, and module docs updated to the new `99-audits/NN-…` paths.

### Decision
- **Scope is spec-only.** F-NEW-01..07 will be addressed by editing documentation only.

---

## [spec-v0.5.0] — 2026-04-23 (audit retirement)

### Added
- **`spec/99-audits/`** — new folder consolidating every spec-quality audit. All previous `spec/99-*.md` files renamed and moved with serialized prefixes:
  - `99-audit-report.md` → `99-audits/01-original-11-step-plan.md`
  - `99-step11-simulation.md` → `99-audits/02-step11-simulation-cycle1.md`
  - `99-step11-simulation-cycle2.md` → `99-audits/03-step11-simulation-cycle2.md`
  - `99-step11-simulation-cycle3.md` → `99-audits/04-step11-simulation-cycle3.md`
  - `99-ai-audit-2026-04-23.md` → `99-audits/05-ai-audit-2026-04-23-gemini.md`
  - `99-ai-audit-2026-04-25.md` → `99-audits/06-ai-audit-2026-04-25-gemini-2.5-pro.md`
- **`spec/99-audits/00-index.md`** — folder index with reading order, scope rules, and process for adding new audits.
- **`spec/99-audits/07-scoreboard.md`** — living scoreboard: current score 85.5/100, open findings register (F-NEW-01..07), targets, and explicit "spec-only, no implementation" decision.
- **New independent audit (Gemini 2.5 Pro)** opened **7 new findings** (F-NEW-01..F-NEW-07), all OPEN. Honest score correction: previously self-reported ~96.8 → independently measured **85.5**.

### Changed
- All cross-references throughout `spec/`, `readme.txt`, and module docs updated to the new `99-audits/NN-…` paths.

### Decision
- **Scope is spec-only.** F-NEW-01..07 will be addressed by editing documentation (`00-llm-integration-guide.md`, `05-enum-system.md`, `02-test-case-types.md`). No Go code, tests, or PowerShell will be touched in this cycle.

---

## [spec-v0.5.0] — 2026-04-23 (audit retirement)

### Closed
- **#02** internal-package coverage policy → **wont-fix** (no maintainer; contradiction documented in both surfaces; reopening criteria preserved).
- **#03** `GetAssert` stability declaration → **wont-fix** (discoverability solved; interim "in-tree only" policy works; reopening criteria preserved).
- **#04** `testwrappers` stability declaration → **wont-fix** (inventory complete; reopening criteria preserved).

### State
- **All 9 issues now closed**: 6 resolved with spec edits, 3 wont-fix with explicit reopening criteria.
- **Audit cycle formally retired** — no remaining self-actionable or maintainer-blocked work in this cycle.
- **Spec is production-ready** at v0.5.0 with 97.17% three-cycle reproducibility average.

### Reopening any wont-fix issue
Each wont-fix issue file (#02, #03, #04) contains an explicit "Reopening criteria" section. To reopen: flip the issue file's `Status:` line back to `open`, update `00-issues-index.md`, and bump spec to next minor version.

---

## [spec-v0.4.0] — 2026-04-23 (Cycle 3 audit)

### Added
- **`spec/99-audits/04-step11-simulation-cycle3.md`** — third fresh-AI simulation, enum scenario.
- **`spec/01-app/05-enum-system.md` §4 "Predicate file-split rule"** — explicit threshold (>6 predicates or >20 lines each → separate files).
- **`spec/02-app-issues/09-enum-predicate-file-split.md`** — issue filed and resolved in same cycle.

### Reproducibility
- **97.50%** on Cycle 3 enum scenario.
- **Three-cycle moving average: 97.17%** (C1 96.25% + C2 97.75% + C3 97.50%).
- Confirms **convergence above ≥95% target** with no regression from Followup edits.

### Resolved (issues)
- **#09** Enum predicate file-split rule — resolved by §4 addition.

---

## [spec-v0.3.0] — 2026-04-23 (Cycle 2 audit)

### Added
- **`spec/99-audits/03-step11-simulation-cycle2.md`** — second fresh-AI simulation, converter scenario.
- **`spec/01-app/04-error-system.md` §Boundary Cases** — 6-row table + decision tree distinguishing `FailedToConvertType` from `ValidationFailedType`.
- **`spec/02-app-issues/08-errcore-type-boundary-examples.md`** — issue filed and resolved in same cycle.

### Reproducibility
- **97.75%** on Cycle 2 converter scenario (up +1.5pp from Cycle 1's 96.25%).
- **Combined two-cycle average ≈ 97.0%** — well above the ≥95% target.

### Resolved (issues)
- **#08** `errcore` type boundary examples — resolved by new §"Boundary Cases".

---

## [spec-v0.2.0] — 2026-04-23 (Asia/Kuala_Lumpur, UTC+8)

### Added
- **`spec/01-app/`** — 14 new atomic per-topic architectural files (~3,400 lines total) covering: repo overview, package map, design philosophy, import conventions, error system, enum system, data structures, conditionals/utilities, validators, converters, reflection/dynamic, versioning, cmd entrypoints, testing patterns, tests folder walkthrough.
- **`spec/02-app-issues/`** — issue tracker (7 issues filed; 4 resolved, 3 maintainer-blocked).
- **`spec/04-tooling/04-bootstrap-into-new-repo.md`** — bootstrap guide.
- **`spec/06-testing-guidelines/02-test-case-types.md`** — Style B section (149 lines).
- **`spec/06-testing-guidelines/03-args-reference.md`** — Style C section + empty-map gotcha + `params.go` convention.
- **`spec/99-audits/02-step11-simulation-cycle1.md`** — fresh-AI reproducibility report (96.25% PASS).
- **`spec/01-app/02-design-philosophy.md`** — explicit filename casing rule (PascalCase exported, camelCase unexported, no snake_case).
- **`spec/01-app/08-validators.md`** — canonical validator error message format with 4 worked examples.
- **`spec/01-app/00-llm-integration-guide.md`** — AI reading order + decision matrix (Step 10).

### Changed
- **Renumbered all of `spec/`** to strict `NN-name.md` convention with no duplicate prefixes.
- **Renamed** `00-audit-report.md` → `99-audits/01-original-11-step-plan.md` (parked at slot 99 as meta).
- **Renamed** `02-tooling/` → `04-tooling/`; `testing-guidelines/` → `06-testing-guidelines/`.
- **Promoted** all 14 `spec/01-app/` files from 📝 Draft → ✅ Done after Step 11 verification.

### Resolved (issues)
- **#01** Style B / Style A coexistence — documented in `06-testing-guidelines/02-test-case-types.md`.
- **#05** `params.go` convention — rule added: mandatory for new packages with > 3 cases.
- **#06** Validator error canonical example — added to `08-validators.md` §5.
- **#07** `newCreator.go` filename casing — added to `02-design-philosophy.md` Pillar 1.

### Open (maintainer-blocked)
- **#02** Internal-package coverage policy — needs maintainer governance call.
- **#03** `GetAssert` stability declaration — discoverability done; stability undecided.
- **#04** `testwrappers` stability declaration — inventory done; stability undecided.

### Reproducibility
- **96.25%** average across 4 axes (package layout, test scaffolding, error categorisation, diagnostic format) — exceeds the ≥95% target set in `99-audits/01-original-11-step-plan.md` §9 Step 11.

---

## [spec-v0.1.0] — pre-2026-04-23 (initial scaffold)

### Added
- Initial `spec/` skeleton with `00-audit-report.md` (later renamed to `99-`).
- `00-llm-integration-guide.md` baseline.
- `testing-guidelines/` with 9 numbered files (later renamed to `06-`).
- `02-tooling/` with 3 files (later renamed to `04-`).
- `03-powershell-test-run/` with 9 files.
- `05-failing-tests/` with 25 chronological fix logs.

---

## Version-bump policy (recap)

Per `spec/01-app/11-versioning.md` §3:

| Spec change category | Bump |
|---|---|
| Major restructure (folder renumbering, breaking link changes) | **major** |
| New topic file, new section, new issue resolution | **minor** |
| Typo fix, link repair, formatting | **none** required |

This release (`v0.2.0`) is a **minor** bump — no breaking changes to deep links, but substantial new content.
