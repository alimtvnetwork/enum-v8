# 01 — Style B / Style A Coexistence

> **Status**: resolved (2026-04-23)
> **Severity**: medium
> **Resolved**: 2026-04-23 — audit Step 7 added a full "Style B" section to `spec/06-testing-guidelines/02-test-case-types.md` (lines 185-333), including structure, required boilerplate, when-to-use decision matrix vs Style A, and pattern-abuse warnings. The decision was: **keep both styles**; Style A is preferred when both fit, Style B is endorsed for typed-slice-input → stringified-lines-output cases shared across multiple `_test.go` files.

## Original Problem

The codebase has two equally-common test styles:

- **Style A** — `CaseV1` + `args.Map` (modern, well documented in `spec/06-testing-guidelines/`)
- **Style B** — `BaseTestCase` + `[]testWrapper` (legacy, used in 6+ packages, undocumented)

A fresh AI given only the spec would force everything into Style A and produce verbose tests when Style B is the natural fit.

## Resolution

Keep both, document Style B explicitly. Done in [`spec/06-testing-guidelines/02-test-case-types.md`](../06-testing-guidelines/02-test-case-types.md) §"Style B" — see the decision matrix at line 320 ("When to Use Style B vs Style A").

## Verification

Step 11 fresh-AI simulation (2026-04-23) did **not** re-trigger this issue; the spec correctly led the simulation to pick Style A for the validator scenario without confusion about Style B's existence.
