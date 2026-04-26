# Issue 07 — `newCreator.go` filename casing not explicit

> **Status**: resolved (2026-04-23)
> **Severity**: low
> **Opened**: 2026-04-23 (Step 11 fresh-AI simulation)
> **Resolved**: 2026-04-23 — added explicit "Filename casing rule" block under Pillar 1 of `spec/01-app/02-design-philosophy.md` covering PascalCase (exported) and camelCase (unexported) cases.
> **Filed by**: spec audit, Step 11

## Summary

`spec/01-app/02-design-philosophy.md` mandates "one file per function" and
shows examples with **camelCase** filenames (e.g. `isSuccess.go`), but does
not state the rule explicitly. The Step 11 simulation correctly inferred
camelCase by analogy, but a fresh agent could equally pick `is_success.go`
(snake_case) — which is more idiomatic Go.

## Why it matters

Mixed casing within the same package would break grep-ability and the
struct-as-namespace visual rhythm.

## Proposed fix

Add a one-line rule to `spec/01-app/02-design-philosophy.md`:

> Filename matches the exported method name in **camelCase**, e.g.
> `IsSuccess` → `isSuccess.go`. Do not use snake_case.

## Impact if not fixed

New packages may mix casing styles, requiring a one-time cleanup pass later.
