# Issue 08 — `errcore.FailedToConvertType` vs `ValidationFailedType` — boundary examples missing

> **Status**: resolved (2026-04-23)
> **Severity**: low
> **Opened**: 2026-04-23 (Step 11 Cycle 2 fresh-AI simulation)
> **Resolved**: 2026-04-23 — added "Boundary Cases" section with 6-row table and decision tree to `spec/01-app/04-error-system.md` before §"See Also".
> **Filed by**: spec audit, Step 11 Cycle 2

## Summary

`spec/01-app/04-error-system.md` lists `FailedToConvertType` and
`ValidationFailedType` as distinct error categories, but does not provide
**boundary case examples** for inputs that could plausibly fall under either.

## Why it matters

During Cycle 2, the simulation correctly chose `FailedToConvertType` for
"missing duration suffix", but the choice required reasoning about intent
(input format wrong vs. valid format with rejected value). A fresh agent
might pick `ValidationFailedType` for the same input.

## Cycle 2 evidence

The simulation explicitly noted: *"`FailedToConvertType` because the input
shape is wrong; `ValidationFailedType` would apply if the input were a valid
duration but outside an allowed range."* This reasoning is correct but
**inferred**, not spec-stated.

## Proposed fix

Add a "Boundary cases" subsection to `spec/01-app/04-error-system.md` with
3-4 worked rows:

| Input | Use type | Rationale |
|---|---|---|
| `"abc"` → `Integer` | `FailedToConvertType` | Unparseable shape |
| `"1500"` → `DurationMillis` | `FailedToConvertType` | Missing required unit suffix |
| `"-5"` → `PositiveInteger` validator | `ValidationFailedType` | Parses fine, fails range rule |
| `""` → any non-nullable | `ValidationFailedType` | Format-valid but business-rejected |

## Impact if not fixed

Mild category drift between converter packages and validator packages.
Does not break attribution (both types are recognised by the runner) but
makes log filtering less precise.
