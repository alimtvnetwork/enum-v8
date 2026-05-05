# Memory Index

> Institutional knowledge for `enum-v2`. Every file in `.lovable/memory/` (including subfolders) MUST be listed here.

## Top-level memory files

- [`01-project-identity.md`](./01-project-identity.md) — Module paths, rename history, hard project facts
- [`02-go-mod-bridge.md`](./02-go-mod-bridge.md) — The `core-v9` ⇄ `core-v8` `replace` bridge: why it exists, why it's broken, what unblocks it
- [`03-spec-audit-protocol.md`](./03-spec-audit-protocol.md) — How code-vs-spec audit cycles work; verifiability dimensions; scoring
- [`04-test-layout.md`](./04-test-layout.md) — `tests/creationtests/` is real, `tests/integratedtests/` is stale upstream-only
- [`05-cross-repo-mirror.md`](./05-cross-repo-mirror.md) — Rules for `cross-repo/core-v8/` (don't rename, don't rewrite v1→v2)

## Subfolders

### `workflow/`
- [`workflow/01-state.md`](./workflow/01-state.md) — Current pipeline state (which cycle, what's blocked, what's next)
- [`workflow/02-task-letter-scheme.md`](./workflow/02-task-letter-scheme.md) — How letter IDs work across "next" turns

### `decisions/`
- [`decisions/01-spec-internal-consistency.md`](./decisions/01-spec-internal-consistency.md) — Cycle 13 introduced spec-internal-consistency as an audit dimension
- [`decisions/02-cmd-main-carve-out.md`](./decisions/02-cmd-main-carve-out.md) — Cycle 10 carved out `cmd/main/` as single permitted smoke-test entrypoint
- [`decisions/03-consumer-coverage-callouts.md`](./decisions/03-consumer-coverage-callouts.md) — Why upstream-only API surfaces get explicit callouts in the spec

### `audits/`
- [`audits/01-cycle-summary.md`](./audits/01-cycle-summary.md) — Compact roll-up of every audit cycle (1–14) for fast lookup
- [`audits/02-finding-registry.md`](./audits/02-finding-registry.md) — All C-CVS-XX and D-CVS-XX findings with status

### `avoid/`
- [`avoid/01-deferred-tasks.md`](./avoid/01-deferred-tasks.md) — Tasks the user has explicitly deferred or skipped (with reasons)

### `suggestions/`
- [`suggestions/01-suggestions-tracker.md`](./suggestions/01-suggestions-tracker.md) — All Lovable suggestions with status tracking (open/done/rejected)

### `pending-issues/`
- [`pending-issues/01-all-pending-issues.md`](./pending-issues/01-all-pending-issues.md) — Consolidated tracker for all pending issues (PI-001 through PI-004)
