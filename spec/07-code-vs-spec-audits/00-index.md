# 07 — Code-vs-Spec Audits

> **Purpose of this folder:** track every attempt to verify that the **actual Go code** in this repository matches what the **spec** claims about it. This folder is **about implementation drift, not spec gaps.** Spec gaps belong in [`99-audits/`](../99-audits/).
>
> **Why this folder exists:** the spec-only audit cycle reached its ceiling (97.5–98.0 / 100, 0 open findings) at spec-v0.17.1. The natural next phase is the inverse — does the code match the docs? See [`spec/99-audits/07-scoreboard.md`](../99-audits/07-scoreboard.md) "Decision: spec-only cycle DEFINITIVELY COMPLETE" for the rationale.

## Reading order

| # | File | Date | Type | Headline |
|---|------|------|------|----------|
| 00 | this file | living | Index + scope rules | — |
| 01 | [`01-scoreboard.md`](./01-scoreboard.md) | living | Aggregated scoreboard | Single source of truth for code-vs-spec drift count |
| 02 | [`02-cycle1-import-conventions.md`](./02-cycle1-import-conventions.md) | 2026-05-04 | Manual cycle | Verify §03 import-convention claims hold in the actual `.go` files |

## Scope rules (important)

- ✅ This folder records **code drift** — places where the code contradicts what the spec says it does.
- ❌ This folder does **NOT** record spec gaps (use `99-audits/`), failing tests (use `05-failing-tests/`), or general code defects unrelated to a spec claim (use `02-app-issues/`).
- ✅ Findings are versioned and immutable once published. Add resolution notes inline; never delete.
- ✅ Every finding must have: spec section that's contradicted (`file:line`), code location that contradicts it (`file:line`), severity, recommended fix, effort.
- ✅ A finding's **resolution path is one of two**:
  1. **Fix the code** to match the spec (preferred when the spec describes intended behavior).
  2. **Fix the spec** to match the code (when the code is correct and the spec drifted from reality).

## How to add a new cycle

1. Pick the next free numeric prefix (`02-`, `03-`, …).
2. Name it `NN-cycleN-<spec-section-slug>.md`.
3. State which spec file you're auditing (one per cycle keeps scope tractable).
4. For each claim in that spec file: quote the claim, point at the code that should implement it, mark ✅/⚠️/❌.
5. Add a row to the table above and to `01-scoreboard.md`.
6. Bump `spec/CHANGELOG.md`.

## Audit methodology

For each spec section being audited:

1. **Extract claims.** Bullet every "the code does X" / "function Y returns Z" assertion.
2. **Locate evidence.** Find the corresponding `.go` file(s) and function(s).
3. **Verify each claim:**
   - ✅ **Match** — code does what spec says.
   - ⚠️ **Drift** — code does *something close* but with a different signature, return type, or naming.
   - ❌ **Contradiction** — code does the opposite or doesn't exist.
4. **Tally** matches / drifts / contradictions.
5. **File findings** for every ⚠️ and ❌ with the resolution path.

A code-vs-spec score is then `matches / total_claims * 100`.
