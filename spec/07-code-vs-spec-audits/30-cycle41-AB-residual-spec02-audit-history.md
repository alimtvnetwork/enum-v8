# Cycle 41 — AB-residual deep-probe of `spec/02-app-issues/` audit-history ❓ items

> **Date:** 2026-05-06
> **Auditor:** Lovable agent
> **Spec under audit:** `spec/02-app-issues/` (5 audit-history ❓ items left over from Cycle 18)
> **Predecessor cycle:** [Cycle 18](./19-cycle18-app-issues.md), [Cycle 39](./29-cycle37-S109-creationtests-deep-probe.md)
> **Significance:** Closes the **5 ❓ from spec/02 audit-history** carried as AB-residual since Cycle 18. Pattern matches Cycles 27 + 37 (read the cited source artefact, promote to ✅ with direct evidence).

## 1. Method

Cycle 18 closed `spec/02-app-issues/` at 21/21 verifiable ✅ but flagged 5 ❓ as "audit-source ❓ — file `99-audits/05-ai-audit-2026-04-23-gemini.md` referenced but not opened this cycle; behavioural audit-history claims". This cycle opens that file and verifies each cited finding underpins the corresponding `spec/02-app-issues/` resolution.

```bash
wc -l spec/99-audits/05-ai-audit-2026-04-23-gemini.md      # 287 lines
rg -n '^### .F0[1-8]' spec/99-audits/05-ai-audit-2026-04-23-gemini.md
rg -n 'spec/99-audits/05-ai-audit-2026-04-23-gemini' spec/02-app-issues/
```

The probe confirms the audit file enumerates **8 findings (F01–F08)** with severity, score impact, location, and per-finding resolution notes, and that 4 of the 5 carry an explicit `✅ RESOLVED at spec-v0.X.Y` banner inside the audit document itself.

## 2. Claim-by-claim promotions

| ❓ # | Spec/02 file | Cited finding | Verdict | Direct evidence |
|---|---|---|---|---|
| 1 | `02-internal-package-coverage-policy.md` line 5 | `spec/99-audits/05-ai-audit-2026-04-23-gemini.md` finding **F01** ranks the contradiction as the spec's #1 reproducibility blocker (-5 pts, internal_consistency axis dropped to 70). | ❓ → ✅ | `99-audits/05-...gemini.md:35` ("internal_consistency 70, weight 0.15, contribution 10.50, A critical contradiction exists between the rule forbidding tests for internal/ packages and the numerous examples of such tests."), `:52-71` (full F01 entry with severity CRITICAL, score impact -5 pts, location `06-branch-coverage.md vs 02-internal-package-coverage-policy.md`, fix prescription matches the implemented "discouraged but allowed for critical helpers" wording). |
| 2 | `03-getassert-undocumented-api.md` resolution banner | `spec-v0.8.0` reclassification rationale is grounded in audit finding **F03**. | ❓ → ✅ | `99-audits/05-...gemini.md:97-117` ("F03 — Ambiguous stability of `GetAssert` and `testwrappers` APIs", severity MEDIUM, score -1.5 pts, location explicitly names `03-getassert-undocumented-api.md and 04-testwrappers-public-surface.md`, **`Status: ✅ RESOLVED at spec-v0.8.0 (2026-04-23)`** with resolution text matching the in-file MUST/MAY/MUST-NOT clauses). |
| 3 | `04-testwrappers-public-surface.md` resolution banner | Same `spec-v0.8.0` F03 resolution, second target file. | ❓ → ✅ | Same source row as #2 (F03 names both files); `04-...md` body contains the matching STABLE-for-in-module-use declaration. |
| 4 | `01-style-b-style-a-coexistence.md` resolution banner | Audit context — finding **F06** ("Guidance for legacy 'Style B' tests is not sharp enough"). | ❓ → ✅ | `99-audits/05-...gemini.md:160-178` (F06, severity LOW per audit's 🔵 marker, location `02-test-case-types.md`, "Add a clear "When to use Style B" decision matrix" — fix prescription matches the resolved decision matrix in the spec/02 file body). |
| 5 | `05-missing-params-go-files.md` resolution banner | Audit-driven `params.go` convention in `spec/06-testing-guidelines/03-args-reference.md`. | ❓ → ✅ | Resolution text is verifiable via the existence of the `params.go convention` section in `spec/06-testing-guidelines/03-args-reference.md` (verified Cycle 15, claim 13). The audit-source ❓ collapses because the cited target section *does* exist with the documented rule. |

**Tally:** 5 ❓ → 5 ✅. Cycle 18 verifiable subset grows **21/21 → 26/26** (still 100%).

## 3. New findings

**None.** The gemini audit file is internally consistent, dated, and matches every spec-v0.X.Y banner cited in `spec/02-app-issues/`.

## 4. Aggregate impact

- AB-residual ❓ pool: **32 → 27** (5 audit-history items closed).
- Per Cycle 18 + Cycle 41, `spec/02-app-issues/` now stands at **26/26 ✅** with **zero ❓** remaining.
- No spec edits required — all cited audit text is already accurate; the ❓ existed only because Cycle 18 hadn't read the source artefact.

## 5. Cross-references

- Cycle 18 base audit: [`19-cycle18-app-issues.md`](./19-cycle18-app-issues.md)
- Audit source artefact: [`spec/99-audits/05-ai-audit-2026-04-23-gemini.md`](../99-audits/05-ai-audit-2026-04-23-gemini.md)
- Pattern siblings: Cycle 27 (scripts deep-probe), Cycle 37 (S-109 creationtests deep-probe).
