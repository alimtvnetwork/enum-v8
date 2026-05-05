# Decision: Spec-Internal Consistency as an Audit Dimension

## Date

2026-05-05 (introduced in Cycle 13, `15-observability.md` audit)

## Context

Through Cycles 1–12, the audit protocol scored claims on a single axis: code-vs-spec match. Sections like §07 (`07-conditional-and-utilities.md`) and §09 (`09-converters.md`) ended up "N/A — no verifiable subset" because every claim was about an upstream-only API surface (❓), so there was nothing to score against.

But this misses a real verifiable property: **does the spec internally hang together?** Specifically:
- Do cross-references resolve to existing files?
- Does a sibling spec file contradict this one?
- Are the rules self-consistent across §X, §Y, §Z?

## Decision

Introduce **spec-internal consistency** as an explicit second verifiability dimension. A claim that is ❓ on code-vs-spec can still be ✅ on spec-internal consistency (e.g. "MUST use `errcore.HandleErr`" is ❓ on code probe but ✅ if `04-error-system.md` carries the same rule).

## Effect

- Cycle 13 (§15) closed at 100% verifiable with 14 ✅ / 13 ❓ — first cycle-on-first-pass to do so with a non-trivial verifiable subset.
- Cycle 14 (§16) closed at 100% verifiable with 17 ✅ / 13 ❓ — second cycle-on-first-pass to do so.
- Task AC opened: re-audit §07 and §09 under the new dimension; both probably have promotable claims (e.g. "decision matrix consistent with §06" → ✅).

## Backport candidates

Already identified for promotion (Task AC):
- §08 row 16 — `errcore.VarTwoNoType` (❓ → ✅; cross-referenced from `04-error-system.md:131`, `08-validators.md:240,307,329`, `15-observability.md`).

## Rule

When auditing a section with high ❓ density, run the spec-internal-consistency probe explicitly:
```
ls <every cross-referenced file>           # check resolution
rg '<recurring banned pattern>' <file>     # check banned-pattern absence
rg '<rule keyword>' spec/01-app/*.md       # check no-contradiction
```
