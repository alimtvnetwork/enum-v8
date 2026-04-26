# Issue 09 — Enum predicate methods: file-split rule unclear

> **Status**: resolved (2026-04-23)
> **Severity**: low
> **Opened**: 2026-04-23 (Step 11 Cycle 3 fresh-AI simulation)
> **Resolved**: 2026-04-23 — added explicit "Predicate file-split rule" sentence to `spec/01-app/05-enum-system.md` §4.
> **Filed by**: spec audit, Step 11 Cycle 3

## Problem

`spec/01-app/05-enum-system.md` §4 shows the canonical 3-file enum recipe
(`consts.go` + `vars.go` + `[Type].go`) but does not explicitly say whether
**predicate methods** like `IsLow()`, `IsCritical()` should live in
`[Type].go` or in separate per-method files (`IsLow.go`, etc.).

`spec/01-app/02-design-philosophy.md` Pillar 1 says "one file per function"
with an exception for "tiny method sets on a value type" (<20 lines each).
A 4-value enum's predicates are each ~3 lines — squarely in the exception.

## Decision applied

**Predicate methods <20 lines each may live in `[Type].go` together; if the
method count exceeds 6 OR any single predicate exceeds 20 lines, split
each into its own file matching the method name.**

This aligns with Pillar 1's exception clause and avoids file-explosion for
small enums.

## Spec edit applied

Added to `spec/01-app/05-enum-system.md` §4 (recipe section):

> **Predicate file-split rule**: Predicate methods (`Is<Name>()`) <20 lines each
> may share `[Type].go`. Once you have >6 predicates OR any predicate exceeds
> 20 lines, split each into its own `Is<Name>.go` file (Pillar 1 of design
> philosophy).
