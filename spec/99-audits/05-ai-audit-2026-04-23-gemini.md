# AI-Driven Spec Audit Report

> **Date**: 2026-04-23 (Asia/Kuala_Lumpur, UTC+8)
> **Auditor**: Gemini 2.5 Pro via Lovable AI Gateway, structured JSON output
> **Spec version audited**: spec-v0.5.0
> **Spec bundle size**: 79 files, ~620 KB, ~155 K tokens
> **Audit scope**: spec ITSELF as a deliverable for autonomous AI development

---

## Headline

| Metric | Value |
|---|---|
| **Overall Score** | **86.3 / 100** |
| **Verdict** | ⚠️ Ship With Caveats |
| **Weighted-axis check** | 86.34 (matches overall) |
| **Findings** | 8 (1 critical, 1 high, 3 medium, 3 low) |
| **Self-reported reproducibility (3 cycles)** | 97.17% |
| **AI-audited reproducibility** | 86.3% |
| **Delta** | **+10.9 pp** (self-report was optimistic) |

### Executive Summary

This is an exceptionally detailed and well-structured specification that has been battle-tested and is actively maintained. Its primary strengths are its self-healing nature, clear testing patterns, and robust error-recovery protocols. However, the spec is undermined by a critical internal contradiction regarding the testing of `internal/` packages, which blocks full automation. Minor ambiguities in rules and the stability of core test helpers lower its otherwise outstanding reproducibility score. With a few targeted fixes to resolve these inconsistencies, this spec will be fully production-ready for autonomous AI development.

---

## Score by Axis

| # | Axis | Score | Weight | Contribution | Rationale |
|---|---|---|---|---|---|
| 1 | **completeness** | 90 | 0.18 | 16.20 | The spec is vast and covers design, implementation, testing, and tooling. Gaps are minor, primarily undocumented limitations of the test framework itself, such as reflection behavior with interfaces. |
| 2 | **unambiguity** | 85 | 0.18 | 15.30 | Strong overall, but ambiguity exists in the 'should' vs 'must' language, the complex enum-splitting rule, and the nuanced decision for when to adopt legacy test Style B. |
| 3 | **internal_consistency** | 70 | 0.15 | 10.50 | A critical contradiction exists between the rule forbidding tests for internal/ packages and the numerous examples of such tests. This is the spec's most significant flaw. |
| 4 | **discoverability** | 98 | 0.10 | 9.80 | Excellent. The AI reading order, detailed tables of contents, and extensive cross-linking make information easy to find. |
| 5 | **executability** | 80 | 0.13 | 10.40 | Most recipes are copy-paste ready. The complex method-naming suffix order is a major exception and is highly prone to AI misinterpretation without more examples of long chains. |
| 6 | **error_recovery** | 98 | 0.08 | 7.84 | Exceptional. The library of past failing tests and the explicit 'Build Error Diagnosis & Fix Protocol' provide a clear path for an AI to recover from common mistakes. |
| 7 | **drift_resistance** | 85 | 0.08 | 6.80 | Good, but the 'wont-fix' issues on API stability and ambiguous 'should' directives create a risk that a mediocre AI will violate unstated conventions. |
| 8 | **verifiability** | 95 | 0.10 | 9.50 | Very strong. The use of structured `args.Map` assertions and diff-based failure reporting makes it easy to verify AI-generated test output. |

**Weakest axes** (lowest 3): `internal_consistency` (70), `executability` (80), `unambiguity` (85)

**Strongest axes** (highest 3): `error_recovery` (98), `discoverability` (98), `verifiability` (95)

---

## Findings — Ranked by Severity & Score Impact

Total score impact: **12.7 points** (explains the gap from 100).

### 🔴 F01 — Contradictory rules for testing `internal/` packages

| | |
|---|---|
| **Severity** | CRITICAL |
| **Score impact** | -5 pts |
| **Effort to fix** | small |
| **Location** | `spec/06-testing-guidelines/06-branch-coverage.md vs spec/02-app-issues/02-internal-package-coverage-policy.md` |

**Why it can fail** (mediocre-AI scenario):  
An AI given a task to add tests to an internal package will face a blocking contradiction: one spec file explicitly forbids it, while other files and the issue tracker acknowledge existing internal tests. The AI may refuse to proceed or guess, violating either the rule or the precedent.

**Evidence**:  
> The rule in `06-branch-coverage.md` states: 'All packages under the internal/ folder are excluded from coverage work.' Issue `02-...-policy.md` confirms the contradiction and closes it as `wont-fix`.

**Fix**:  
Remove the contradiction. EITHER remove the prohibition from `06-branch-coverage.md` and state the policy is 'discouraged but allowed for critical helpers', OR remove all `...internaltests/` packages referenced in the spec.

---

### 🟠 F02 — Method suffix combination order is too complex

| | |
|---|---|
| **Severity** | HIGH |
| **Score impact** | -2.5 pts |
| **Effort to fix** | small |
| **Location** | `spec/00-llm-integration-guide.md` |
| **Status** | ✅ **RESOLVED** at spec-v0.7.0 (2026-04-23) |

**Why it can fail** (mediocre-AI scenario):  
An AI asked to create a method like 'add a non-empty string slice if a flag is true, with a lock' must follow the order: Base+Filter+Type+Lock+If+Must. It could easily generate `AddsSliceNonEmptyLockIf` instead of the correct `AddsNonEmptySliceLockIf`. The rule is stated but under-exemplified.

**Evidence**:  
> The 'Master Suffix Reference Table' in `00-llm-integration-guide.md` specifies a rigid 6-part order for method name suffixes but provides few examples of combinations with more than 2-3 parts. The cognitive load is high.

**Resolution (spec-v0.7.0)**:  
Added a **"Long-Chain Suffix Gallery"** section to `spec/00-llm-integration-guide.md` (after Pattern 7) with:
- 7 worked examples ranging from 3 to 7 tokens (`DeserializePtrMust` → `NonEmptyItemsOrNonWhitespacePtrLockIfMust`)
- 6 explicit anti-patterns showing common AI ordering mistakes
- A 7-step **Slot Validation Checklist** to run before committing any new method name
- Explicit slot-mapping for ambiguous cases (`Trimmed*` as Base-transform, `*Join` as terminal reducer, compound `*Or*` filters)

---

### 🟡 F03 — Ambiguous stability of `GetAssert` and `testwrappers` APIs

| | |
|---|---|
| **Severity** | MEDIUM |
| **Score impact** | -1.5 pts |
| **Effort to fix** | trivial |
| **Location** | `spec/02-app-issues/03-getassert-undocumented-api.md and 04-testwrappers-public-surface.md` |
| **Status** | ✅ **RESOLVED** at spec-v0.8.0 (2026-04-23) |

**Why it can fail** (mediocre-AI scenario):  
A cautious AI agent will see that the stability of core testing helpers is explicitly 'wont-fix' and 'undecided'. It may refuse to use these APIs, leading it to write verbose, non-idiomatic test code from scratch instead of reusing the existing, purpose-built frameworks.

**Evidence**:  
> Issues #03 and #04 are closed as `wont-fix` with the rationale that formal stability is a 'maintainer-only call that does not block spec usage'. This forces the AI to risk using an API the spec flags as potentially unstable.

**Resolution (spec-v0.8.0)**:  
Reclassified both issues `wont-fix` → `resolved` with explicit **"STABLE for in-module use"** declarations. Both files now contain MUST/MAY/MUST-NOT clauses giving AI agents unambiguous guidance: reuse the APIs in-module; do not import from external modules. The narrower question of publishing as external stable packages is preserved as the sole reopening criterion.

---

### 🟡 F04 — Ambiguous 'should' vs 'must' wording encourages rule violations

| | |
|---|---|
| **Severity** | MEDIUM |
| **Score impact** | -1.5 pts |
| **Effort to fix** | medium |
| **Location** | `spec/01-app/10-reflection-and-dynamic.md` |
| **Status** | ✅ **RESOLVED** at spec-v0.9.0 (2026-04-23) |

**Why it can fail** (mediocre-AI scenario):  
A literal AI will interpret 'should' as a suggestion. If the spec says code 'should reach for coredynamic', the AI may decide to import the `reflect` package directly to be more 'efficient', bypassing the spec's safety wrappers and reintroducing panics the framework was designed to prevent.

**Evidence**:  
> `spec/01-app/10-reflection-and-dynamic.md` states: 'New consumer code should reach for coredynamic or reflectcore first.' The intent is mandatory, but the wording is optional.

**Resolution (spec-v0.9.0)**:  
Rewrote the named rule to use **MUST / MUST NOT** explicitly and added an explicit prohibition on importing the `reflect` stdlib package directly. Added a **project-wide "Convention vs Hard Rule"** legend in the same block defining MUST/MUST NOT/MAY (non-negotiable) vs should/rule-of-thumb/prefer (guidance), with a default-to-MUST tiebreaker for safety/correctness rules. A repo-wide audit found this was the only hard-rule `should` in `spec/01-app/`; the remaining `should` usages (e.g. "Reflection should appear only when…") are correctly framed as guidance.

---

### 🟡 F05 — Undocumented testing framework limitations

| | |
|---|---|
| **Severity** | MEDIUM |
| **Score impact** | -1 pts |
| **Effort to fix** | small |
| **Location** | `spec/05-failing-tests/09-failing-tests-round4.md` |
| **Status** | ✅ **RESOLVED** at spec-v0.10.0 (2026-04-23) |

**Why it can fail** (mediocre-AI scenario):  
An AI writing a test for a method that returns an `error` interface will use `VerifyOutArgs`. The test will fail because the underlying reflection logic doesn't match concrete error types (e.g., `*errors.errorString`) to the interface. The AI cannot predict this failure as the limitation is only documented in a failing test log.

**Evidence**:  
> Failing test #Test_I13_VerifyOutArgs_Success notes: '`reflect.TypeOf` returns concrete types, not interface types — `VerifyOutArgs` can never match interface return types.' This is a critical framework limitation absent from the main spec.

**Resolution (spec-v0.10.0)**:  
Added a **"Known Framework Limitations"** section to `spec/06-testing-guidelines/02-test-case-types.md` documenting three quirks (L1: `VerifyOutArgs` interface mismatch; L2: `InvokeFirstAndError` 2-return assumption; L3: `t.Run` failure propagation). Each entry uses a Symptom / Root-cause / MUST-do-workaround / Source format. Added a "Reporting a new limitation" subsection establishing this file — not `05-failing-tests/` — as the canonical discovery point for framework quirks going forward.

---

### 🔵 F06 — Guidance for legacy 'Style B' tests is not sharp enough

| | |
|---|---|
| **Severity** | LOW |
| **Score impact** | -0.5 pts |
| **Effort to fix** | small |
| **Location** | `spec/01-app/13-testing-patterns.md` |

**Why it can fail** (mediocre-AI scenario):  
The spec says to use Style B for 'hetero typed-slice input'. A mediocre AI will struggle to define 'hetero' and will default to the better-documented Style A, producing verbose tests. While not a 'crash', it fails to reproduce the idiomatic patterns present in the codebase.

**Evidence**:  
> The decision matrix in `13-testing-patterns.md` gives a vague condition for Style B, and the resolution of issue #01 just says to 'keep both styles'.

**Fix**:  
Refine the Style B decision matrix with a concrete rule. For example: 'Use Style B only when the input is a slice of structs/interfaces and the output is a `[]string` where each line's format depends on the input row's type.'

---

### 🔵 F07 — Rule for splitting enum predicate methods is overly complex

| | |
|---|---|
| **Severity** | LOW |
| **Score impact** | -0.5 pts |
| **Effort to fix** | trivial |
| **Location** | `spec/01-app/05-enum-system.md` |

**Why it can fail** (mediocre-AI scenario):  
An AI creating a new enum is told to split predicate methods into separate files if there are '>6 predicates OR any single predicate exceeds 20 lines'. A mediocre AI may mis-count, or struggle with the compound boolean logic, leading to inconsistent file structures.

**Evidence**:  
> The rule in `05-enum-system.md`'s 'Predicate file-split rule' section requires counting and line-length estimation, which is brittle for automation.

**Fix**:  
Simplify the rule. A simpler, deterministic rule would be: 'Predicate methods for enums with more than 5 values must each be in their own file. Otherwise, they may be grouped.'

---

### 🔵 F08 — Deprecated `ExpectedLines` field in `GenericGherkins` creates ambiguity

| | |
|---|---|
| **Severity** | LOW |
| **Score impact** | -0.2 pts |
| **Effort to fix** | small |
| **Location** | `spec/06-testing-guidelines/02-test-case-types.md` |

**Why it can fail** (mediocre-AI scenario):  
A mediocre AI sees the `ExpectedLines` field in the `GenericGherkins` struct definition. Although the spec says it's deprecated, its presence is a signal. The AI might use it, opting for opaque `[]string` assertions instead of the preferred, self-documenting `Expected: args.Map{...}` pattern.

**Evidence**:  
> The 'Field Responsibility' table for `GenericGherkins` lists `ExpectedLines` with the note 'Deprecated for new tests'.

**Fix**:  
Remove the `ExpectedLines` field from the struct definition in the documentation entirely. State that older tests may use it, but it is not part of the modern API.

---

## Strengths (don't break these)

- ✅ The spec is self-healing; lessons from the `05-failing-tests/` log are explicitly incorporated back into the main architectural documents, preventing repeat mistakes.
- ✅ A clear reading order and comprehensive cross-linking across atomic, topic-oriented files make discoverability exceptionally high for an AI agent.
- ✅ The AAA (Arrange-Act-Assert) pattern and the strict separation of test data (`_testcases.go`) from test logic (`_test.go`) are enforced with zero ambiguity.
- ✅ The `spec/03-powershell-test-run/` documentation provides a robust, copy-paste ready 'Build Error Diagnosis & Fix Protocol' that enables excellent automated error recovery.
- ✅ The use of `args.Map` for both inputs and expectations, combined with diff-based assertions, creates a highly verifiable and debuggable testing environment.
- ✅ The `04-tooling/04-bootstrap-into-new-repo.md` guide makes the entire sophisticated PowerShell and CI toolchain portable to new projects.

---

## Roadmap to 100% — Ordered by ROI

Projected gain if **all** roadmap items shipped: **+12.5 pts** → projected score **98.8 / 100**.

| # | Step | Effort | Expected gain |
|---|---|---|---|
| 1 | Resolve the `internal/` testing contradiction in `06-branch-coverage.md` by aligning the rule with the observed practice. | trivial | +5 pts |
| 2 | Add a dedicated subsection to `00-llm-integration-guide.md` with 5-7 fully-worked examples of the method suffix combination rule. | small | +3 pts |
| 3 | Change the status of testing-API issues #03 and #04 from `wont-fix` to `resolved` with a clear stability statement for in-module use. | trivial | +2 pts |
| 4 | Globally replace ambiguous 'should' directives with 'must' for all hard rules. | medium | +1.5 pts |
| 5 | Document known test framework limitations (e.g., reflection with interfaces) in a new 'Known Limitations' section in the testing guidelines. | small | +1 pts |

### What stops 100% even after the roadmap?

- **1.2 pts** of irreducible gap from spec medium that no amount of editing can close (e.g., AI cannot fully verify rules without seeing the actual Go source code, which is outside this Lovable project).
- The Go module `core-v9` itself is **not in this Lovable project** — only the `spec/` deliverable. So the audit can only score the spec's *internal* qualities, not its *implementation fidelity*.
- For true 100%, the spec would need to be checked against the actual `core-v9` source tree in a Go workspace — out of scope here.

---

## Comparison with Self-Reported Audit (3 simulation cycles)

| Source | Score | Method |
|---|---|---|
| Cycle 1 (validator scenario) | 96.25% | Single-scenario simulation |
| Cycle 2 (converter scenario) | 97.75% | Single-scenario simulation |
| Cycle 3 (enum scenario) | 97.50% | Single-scenario simulation |
| **Self-reported avg** | **97.17%** | 3 disjoint simulations |
| **Independent AI audit** | **86.3%** | Whole-spec analysis, 8 axes, 8 findings |
| **Delta** | **+10.9 pp** | Self-report ran scenarios that *passed*; independent audit found contradictions and ambiguities not surfaced by passing scenarios |

**Interpretation**: Both numbers are correct in their own frame. The self-report measured what an AI *would produce*; this audit measured what an AI *might trip over*. The truth is between the two — call it ~91%.

---

## Action Items (immediate, this audit cycle)

Highest-ROI fixes (effort = trivial/small, total gain ≈ +11.0 pts):

- [ ] **+5 pts** (trivial): Resolve the `internal/` testing contradiction in `06-branch-coverage.md` by aligning the rule with the observed practice.
- [ ] **+3 pts** (small): Add a dedicated subsection to `00-llm-integration-guide.md` with 5-7 fully-worked examples of the method suffix combination rule.
- [ ] **+2 pts** (trivial): Change the status of testing-API issues #03 and #04 from `wont-fix` to `resolved` with a clear stability statement for in-module use.
- [ ] **+1 pts** (small): Document known test framework limitations (e.g., reflection with interfaces) in a new 'Known Limitations' section in the testing guidelines.

---

## Audit Methodology

1. **Bundling**: All 79 spec/*.md files concatenated into a single ~155K-token corpus.
2. **Prompting**: Gemini 2.5 Pro instructed to adopt a 'mediocre AI agent' perspective and audit across 8 weighted axes.
3. **Schema**: Structured JSON output enforced via tool-calling schema (no free-form drift).
4. **Independence**: This audit was run *after* the self-reported 3-cycle audit and had no shared state with it (cold-context AI).
5. **Reproducibility**: Re-run with `python3 /tmp/lovable_ai.py @/tmp/audit/full_prompt.txt --schema /tmp/audit/schema.json --model google/gemini-2.5-pro --output /tmp/audit/report.json`.

---

_Report generated by Lovable AI Gateway · Gemini 2.5 Pro · structured-output mode._
