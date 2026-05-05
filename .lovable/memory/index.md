# Project Memory

## Core
This is a Go enum library (enum-v3, module path `github.com/alimtvnetwork/enum-v3` — renamed from `enum-v2`, originally `enum-v1`). The frontend React/Vite shell is incidental — work happens in Go packages and PowerShell/Python tooling.
Core dependency import path is `github.com/alimtvnetwork/core-v9` (was renamed from `core-v8`). All source imports use `core-v9` — never reintroduce `core-v8` outside `cross-repo/core-v8/`.
**`go.mod` bridge RESOLVED:** upstream `core-v9` now declares `module github.com/alimtvnetwork/core-v9` (tag `v1.5.8`). The `replace` directive has been removed. `go.mod` pins `require core-v9 v1.5.8` cleanly.
The `cross-repo/core-v8/` directory mirrors CI surface to a separate upstream repo and intentionally keeps its `core-v8` name. Do not rename it. Also do not rewrite `enum-v1` → `enum-v3` inside this directory — it tracks a different module.
**Tests live under `tests/creationtests/`, NOT `tests/integratedtests/`.** Tooling that probes test packages must accept either name (or read from disk) — never hard-code one. Spec docs that still say `integratedtests` are stale (see audit finding C-CVS-01).
**NEVER send emails or configure email-based notifications** (Dependabot recipients, SMTP, CI email alerts). User rejects all email flows.
User workflow: incremental "next" commands, expects the remaining-task list shown after every step.
**Any code changes must bump at least minor version everywhere except `.release/` folder — never modify `.release/`.**
**Suggestions go in `.lovable/memory/suggestions/01-suggestions-tracker.md` — single file, update in-place.**
**Pending issues go in `.lovable/memory/pending-issues/01-all-pending-issues.md` — single file, update in-place.**
**Plan lives at `.lovable/plan.md` — single file with phases, task letters, and next-task selection.**

## Memories
- [Workflow conventions](mem://preferences/workflow) — "next" loop, task-letter scheme, always list remaining tasks
- [CI tooling layout](mem://features/ci-tooling) — ci-guards.yml structure, baseline gate, python-tests job, run.ps1 pre-checks
- [No email notifications](mem://constraints/no-email-notifications) — never send or configure email/notification delivery of any kind
- [Suggestions tracker](mem://features/suggestions-tracker) — .lovable/memory/suggestions/01-suggestions-tracker.md convention
- [Reliability report](mem://features/reliability-report) — /mnt/documents/01-reliability-risk-report.md produced 2026-05-05
- [Go mod bridge — resolved](mem://02-go-mod-bridge) — Was blocking; now resolved. core-v9 v1.5.8 pinned cleanly.
