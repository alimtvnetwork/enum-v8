# Cycle 6 — `01-app/08-validators.md`

> **Date**: 2026-05-05 (Asia/Kuala_Lumpur)
> **Spec audited**: [`spec/01-app/08-validators.md`](../01-app/08-validators.md)
> **Auditor**: Lovable agent (loop AA-cycle6)
> **Status**: **baseline recorded — section is 100 % upstream-only**

---

## 1. Method

For each numbered section in the spec, classify every concrete claim (import path, exported symbol, signature, behavioural rule, diagnostic-format contract) as:

- ✅ **Match** — claim verified against `enum-v3` source on disk.
- ⚠️ **Drift** — verifiable but inaccurate.
- ❌ **Contradiction** — verifiable and wrong.
- ❓ **Unverifiable** — package not consumed by `enum-v3` and no upstream `core-v9` source on disk; defer to task **AB**.

Verification commands run from repo root:

```bash
rg -l "core-v9/corevalidator" --type go
rg -n "corevalidator\.(New|LineValidator|SliceValidator|TextValidator|RangeValidator|StringCompareAs|Result)" --type go
rg -n "errcore\.VarTwoNoType|ValidationFailedType" --type go
rg -n "regexnew\.New\.Lazy" --type go
rg -n "coretestcases\.(CaseV1|CaseNilSafe)" --type go
ls cross-repo/core-v9/corevalidator 2>/dev/null
```

All commands returned **zero matches**: no `enum-v3` package imports `corevalidator` or any of its peers (`errcore.VarTwoNoType`, `ValidationFailedType`, `regexnew.New.Lazy`, `coretestcases.CaseV1/CaseNilSafe`), and the `cross-repo/core-v9/` mirror does not carry `corevalidator` source.

---

## 2. Claim inventory

| #  | Spec § | Claim                                                                                      | Verdict | Note |
|----|--------|--------------------------------------------------------------------------------------------|---------|------|
| 1  | §1     | `corevalidator` provides Line/Slice/Text/Range validators implementing `coreinterface/*`   | ❓ | No consumer; needs upstream source |
| 2  | §1     | Validator contract: `IsValid`, `IsSuccess`, `IsFailed`, `Message`, `Error`                 | ❓ | Contract surface unverifiable |
| 3  | §2.1   | `corevalidator.New.Line.NotEmpty().MaxLength(N).Matches(re).Build() *LineValidator`        | ❓ | No consumer |
| 4  | §2.1   | `Build()` returns pointer; value-typed storage breaks interface assertion (F-V14-01)       | ❓ | Behavioural rule — no consumer |
| 5  | §2.2   | `corevalidator.New.Slice.MinLength.MaxLength.EachItem(v).Build()`                          | ❓ | No consumer |
| 6  | §2.3   | `TextValidator` — multi-line whole-document rules                                          | ❓ | No consumer |
| 7  | §2.4   | `corevalidator.New.Range.Int.Min(N).MaxExclusive(N).Build()` semantics                     | ❓ | No consumer |
| 8  | §2.5   | `StringCompareAs` diagnostic format is regex-checked by framework tests                    | ❓ | No consumer |
| 9  | §3.1   | `errcore.MergeErrors(errs...)` aggregates validator errors                                 | ❓ | `MergeErrors` not invoked anywhere in `enum-v3` |
| 10 | §3.2   | Domain-type embedding pattern (`(r *T) Validate() error`)                                  | ❓ | Pattern not exercised in `enum-v3` |
| 11 | §3.3   | `conditional.IfFunc[error](...)` for branching validation                                  | ❓ | `conditional` already ❓ in Cycle 5 |
| 12 | §4.1–4.3 | Custom-validator template uses `regexnew.New.Lazy` + nil-receiver guard                  | ❓ | No consumer |
| 13 | §4.3   | `Result.Error()` returns `errcore.ValidationFailedType.Fmt(...)`                           | ❓ | `ValidationFailedType` symbol not imported |
| 14 | §4.4   | Tests use `coretestcases.CaseV1` + `CaseNilSafe` (Style A)                                 | ❓ | Neither symbol imported anywhere in `enum-v3` |
| 15 | §5     | Diagnostic message format: `<Label>: field=<n> value=<v> reason=<short>`                   | ❓ | Format contract has no in-repo emitter |
| 16 | §5     | `errcore.VarTwoNoType("field", label, "value", actual)` is the canonical builder          | ❓ | Symbol not imported anywhere in `enum-v3` |
| 17 | §6     | Three-test layout: `*_Verification_test.go` / `*_NilReceiver_test.go` / `*_Format_test.go` | ❓ | No validator tests under `tests/` |
| 18 | §6     | Path: `tests/integratedtests/<pkg>tests/`                                                  | ⚠️ | **Stale path** — repeats C-CVS-01 / D-CVS-17 pattern: actual layout is `tests/creationtests/`. Filed as **D-CVS-26**. |
| 19 | §7     | Common-mistakes table (7 rows)                                                             | ❓ | All rows depend on §1–§4 surface |

**Total claims**: 19
**Verifiable subset**: 1 (claim #18 — path-string drift, verified against on-disk `tests/` layout)
**Verifiable match rate (baseline)**: **0 / 1 = 0.0%** *(one ⚠️ Drift, no other measurable claims)*

---

## 3. Score row

| Date       | Cycle | Spec audited                      | Claims | ✅ | ⚠️ | ❌ | ❓ | Score (verifiable) |
|------------|-------|-----------------------------------|--------|----|-----|----|----|--------------------|
| 2026-05-05 | 6 (baseline) | `01-app/08-validators.md`   | 19     | 0  | 1   | 0  | 18 | **0.0%** *(1 of 19 measurable today)* |

---

## 4. Findings

### D-CVS-26 — Stale `tests/integratedtests/` reference (§6) — **LOW (path-string fix)**

Spec §6 line 347:
> Each validator package needs three test files in `tests/integratedtests/<pkg>tests/`:

Repository reality: `tests/integratedtests/` does not exist; tests live under `tests/creationtests/` (this is the same drift pattern already filed as **C-CVS-01** for §03 and **D-CVS-17** for §05, both resolved with identical s/integratedtests/creationtests/ rewrites).

**Fix**: rewrite §6 path to `tests/creationtests/<pkg>tests/` and cross-reference the established convention in `13-testing-patterns.md` (matching the C-CVS-01 / D-CVS-17 resolutions).

**No other contradictions found.** All 17 ❓ rows defer to task **AB** (fetch upstream `core-v9/corevalidator` source).

---

## 5. Next actions

1. Apply **D-CVS-26** as a one-line spec fix (closes §08's only verifiable drift).
2. Update the scoreboard (`01-scoreboard.md`) with the Cycle 6 baseline + closed rows.
3. Defer the remaining 17 ❓ to **AB** alongside Cycle 5's 17 ❓ — both cycles share the same blocker (no upstream `corevalidator` / `conditional` / `isany` / `regexnew` source on disk).
