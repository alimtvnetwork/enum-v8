# 99-audits — Spec Reproducibility Audits

> **Purpose of this folder:** track every attempt to measure whether *another AI* could reimplement this codebase using **only** the spec. This folder is **about the spec, not the implementation.** Issues raised here describe documentation gaps — not code defects. Code defects belong in `spec/02-app-issues/` or `spec/05-failing-tests/`.

## Reading order

| # | File | Date | Type | Headline |
|---|------|------|------|----------|
| 01 | [`01-original-11-step-plan.md`](./01-original-11-step-plan.md) | 2026-04-22 | Manual gap analysis | Original 11-step plan to reach ≥95% reproducibility |
| 02 | [`02-step11-simulation-cycle1.md`](./02-step11-simulation-cycle1.md) | 2026-04-22 | Fresh-AI simulation | Validator scenario — **96.25%** PASS |
| 03 | [`03-step11-simulation-cycle2.md`](./03-step11-simulation-cycle2.md) | 2026-04-22 | Fresh-AI simulation | Converter scenario — **97.75%** PASS |
| 04 | [`04-step11-simulation-cycle3.md`](./04-step11-simulation-cycle3.md) | 2026-04-22 | Fresh-AI simulation | Enum scenario — **97.50%** PASS |
| 05 | [`05-ai-audit-2026-04-23-gemini.md`](./05-ai-audit-2026-04-23-gemini.md) | 2026-04-23 | Independent AI audit | Gemini Flash — 5 findings F01–F05, all RESOLVED |
| 06 | [`06-ai-audit-2026-04-25-gemini-2.5-pro.md`](./06-ai-audit-2026-04-25-gemini-2.5-pro.md) | 2026-04-25 | Independent AI audit | Gemini 2.5 Pro — **85.5/100**, 7 findings F-NEW-01..07, all RESOLVED at v0.12.0 |
| 07 | [`07-scoreboard.md`](./07-scoreboard.md) | living | Aggregated scoreboard | Single source of truth for current spec quality score |
| 08 | [`08-ai-audit-2026-04-25-gpt5-confirm.md`](./08-ai-audit-2026-04-25-gpt5-confirm.md) | 2026-04-25 | 3rd-party confirmation audit | GPT-5 — **88.9/100** measured; 9 findings F-V12-01..09, all RESOLVED at v0.14.0 |
| 09 | [`09-ai-audit-2026-04-25-claude-sonnet-4.5.md`](./09-ai-audit-2026-04-25-claude-sonnet-4.5.md) | 2026-04-25 | 4th-party confirmation audit | Claude Sonnet 4.5 — **96.0/100** measured; F-V14-01..04 RESOLVED at v0.15.1, F-V14-05 RESOLVED at v0.16.0 |
| 10 | [`10-ai-audit-2026-04-25-mediocre-sim.md`](./10-ai-audit-2026-04-25-mediocre-sim.md) | 2026-04-25 | 5th-party mediocre-AI simulation | **97.5/100** measured; 1 optional finding F-V16-01 OPEN |

## Current spec quality (as of 2026-04-25, post spec-v0.17.1)

| Metric | Value |
|---|---|
| Latest INDEPENDENT MEASURED score | **97.5 / 100** (mediocre-AI sim, 5th-party) |
| Projected score post-v0.17.1 | **~98.0 / 100** (F-V16-01 closed) |
| Score progression (5 audits) | 85.5 → 88.9 → 96.0 → 97.5 (converged inside 97–98 ceiling) |
| Open spec findings | **0** |
| Resolved spec findings | **27** (F01–F05, F-NEW-01..07, F-V12-01..09, F-V14-01..05, F-V16-01) |
| Practical ceiling | **~97–98 / 100** (per GPT-5; corroborated by Claude × 2 audits) |

See [`07-scoreboard.md`](./07-scoreboard.md) for the breakdown.

## How to add a new audit

1. Pick the next free numeric prefix (`07-`, `08-`, …).
2. Use this naming pattern:
   - Manual reviews: `NN-<topic>.md`
   - Fresh-AI simulations: `NN-step11-simulation-cycle<N>.md`
   - Independent AI audits: `NN-ai-audit-YYYY-MM-DD-<model>.md`
3. Add a row to the table above and to `07-scoreboard.md`.
4. Bump `spec/CHANGELOG.md`.
5. **Never** rewrite history — append, don't edit prior audit conclusions.

## Scope rules (important)

- ✅ This folder records **spec gaps** (documentation that would mislead an implementer AI).
- ❌ This folder does **NOT** record code bugs, failing tests, or implementation TODOs.
- ✅ Findings are versioned and immutable once published. Add resolution notes inline; never delete.
- ✅ Every finding must have: severity, evidence (file:line), failure mode, recommended fix, effort, and score uplift.
