# Issue 06 — Validator error message has no canonical example

> **Status**: resolved (2026-04-23)
> **Severity**: low
> **Opened**: 2026-04-23 (Step 11 fresh-AI simulation)
> **Resolved**: 2026-04-23 — added "Canonical error message format" subsection with 4 worked examples to `spec/01-app/08-validators.md` §5.
> **Filed by**: spec audit, Step 11

## Summary

`spec/01-app/08-validators.md` describes the validator interface contract and
the `errcore.NewValidationErrorType` mapping, but does **not** give a verbatim
example of the error string a validator should produce.

## Why it matters

During the Step 11 simulation, the agent had to **infer** the message format
(e.g. `"value is empty"` vs `"NonEmptyLine: value must not be empty"`). Both
are valid under the spec, but only one matches existing repo style.

## Proposed fix

Add a "Canonical error message format" subsection to
`spec/01-app/08-validators.md` with 2-3 worked examples copied verbatim from
existing validators (e.g. `linevalidator`, `slicevalidator`).

## Impact if not fixed

Cosmetic drift in error strings across third-party validator packages — does
not affect functionality or test attribution.
