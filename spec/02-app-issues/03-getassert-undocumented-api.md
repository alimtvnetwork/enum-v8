# 03 — `GetAssert` Public-Api Stability

> **Status**: resolved (2026-04-23, spec-v0.8.0)
> **Severity**: low
> **Resolved via**: AI-driven audit finding F03 — explicit stability declaration replaces deferred maintainer call.

## Stability Declaration (authoritative)

`coretests.GetAssert` is **STABLE for any test code inside this module**.

- ✅ **MUST** be reused by all in-module test packages instead of writing
  duplicate assertion helpers.
- ✅ **MAY** be extended with new helpers following the rules in
  `spec/01-app/14-tests-folder-walkthrough.md` §3.
- ❌ **MUST NOT** be imported by external consumers of this Go module.
  External consumers should treat `coretests.*` as internal — no backward
  compatibility guarantees apply across module boundaries.

This declaration is sufficient for autonomous AI agents to use the API
with full confidence inside the module. The narrower question of
publishing `coretests` as a stable external package remains a separate
maintainer-only decision and is **out of scope** for this issue.

## Resolution Rationale (spec-v0.8.0)

The audit (F03) found that the previous `wont-fix` status caused cautious
AI agents to avoid `GetAssert` and write verbose ad-hoc assertions
instead. The fix is purely a **wording change** in this issue file plus
the explicit "stable for in-module use" rule above — no code changes
required.

The discoverability work (referenced below) was already complete; only
the stability ambiguity remained, and it is now removed.

## Discoverability (already in place — reference)

`spec/01-app/14-tests-folder-walkthrough.md` §3 inventories the methods,
usage pattern, and rules for adding helpers. Step 11 Cycles 1, 2, and 3
all reused `GetAssert` correctly.

## Reopening criteria

Reopen only when:
- A maintainer decides to publish `coretests` as a stable external
  package (would expand the stability scope beyond in-module use), OR
- A breaking change to `GetAssert` is proposed (would require deprecation
  notes here).

## Original problem statement (preserved)

`coretests.GetAssert` exposed a rich helper surface used by hundreds of
test cases, but **none** of it was documented in `spec/`. A fresh AI could
not discover or reuse it.
