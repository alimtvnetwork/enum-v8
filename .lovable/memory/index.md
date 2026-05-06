# Project Memory

## Core
This is a Go enum library (enum-v3, module path `github.com/alimtvnetwork/enum-v3` — renamed from `enum-v2`, originally `enum-v1`). The frontend React/Vite shell is incidental — work happens in Go packages and PowerShell/Python tooling.
Core dependency import path is `github.com/alimtvnetwork/core-v9` (was renamed from `core-v8`). All source imports use `core-v9` — never reintroduce `core-v8` outside `cross-repo/core-v8/`.
**`go.mod` bridge RESOLVED:** upstream `core-v9` now declares `module github.com/alimtvnetwork/core-v9` (tag `v1.5.8`). The `replace` directive has been removed. `go.mod` pins `require core-v9 v1.5.8` cleanly.
**API migration needed:** `core-v9` refactored converters from package-level functions to struct namespaces (`AnyTo`, `StringTo`). `coredynamic.TypeName` → `SafeTypeName`. See `06-core-v9-api-migration.md`.
The `cross-repo/core-v8/` directory mirrors CI surface to a separate upstream repo and intentionally keeps its `core-v8` name. Do not rename it. Also do not rewrite `enum-v1` → `enum-v3` inside this directory — it tracks a different module.
**Tests live under `tests/creationtests/`, NOT `tests/integratedtests/`.** Tooling that probes test packages must accept either name (or read from disk) — never hard-code one. Spec docs that still say `integratedtests` are stale (see audit finding C-CVS-01).
**NEVER send emails or configure email-based notifications** (Dependabot recipients, SMTP, CI email alerts). User rejects all email flows.
User workflow: incremental "next" commands, expects the remaining-task list shown after every step.
**Any code changes must bump at least minor version everywhere except `.release/` folder — never modify `.release/`.**
**Suggestions go in `.lovable/suggestions.md` — single file, update in-place.**
**Pending issues go in `.lovable/pending-issues/` — one file per issue.**
**Plan lives at `.lovable/plan.md` — single file with phases, task letters, and next-task selection.**

## Memories
- [Project identity](mem://01-project-identity) — Module path, core dependency, repo layout
- [Go mod bridge — resolved](mem://02-go-mod-bridge) — Was blocking; now resolved. core-v9 v1.5.8 pinned cleanly.
- [Spec audit protocol](mem://03-spec-audit-protocol) — How audit cycles work, scoring dimensions
- [Test layout](mem://04-test-layout) — tests/creationtests/ not integratedtests/
- [Cross-repo mirror](mem://05-cross-repo-mirror) — cross-repo/core-v8/ rules
- [core-v9 API migration](mem://06-core-v9-api-migration) — Converter/coredynamic function→struct namespace mapping
- [Workflow conventions](mem://workflow/01-state) — Current workflow state snapshot
- [Task letter scheme](mem://workflow/02-task-letter-scheme) — Letter IDs, "next" loop, remaining-task list
- [Deferred tasks](mem://avoid/01-deferred-tasks) — Tasks to skip or defer
- [Spec-internal consistency](mem://decisions/01-spec-internal-consistency) — Audit dimension added Cycle 13
- [cmd/main carve-out](mem://decisions/02-cmd-main-carve-out) — Smoke-test policy
- [Consumer-coverage callouts](mem://decisions/03-consumer-coverage-callouts) — Upstream-only API labeling
- [Cycle summary](mem://audits/01-cycle-summary) — Cycles 1-14 overview
- [Finding registry](mem://audits/02-finding-registry) — C-CVS and D-CVS findings
- [Suggestions tracker](mem://suggestions/01-suggestions-tracker) — Legacy tracker (now consolidated in .lovable/suggestions.md)
- [Pending issues](mem://pending-issues/01-all-pending-issues) — Legacy tracker (now individual files in .lovable/pending-issues/)
- [CI tooling layout](mem://features/ci-tooling) — ci-guards.yml structure, baseline gate, python-tests job, run.ps1 pre-checks (if exists)
- [No email notifications](mem://constraints/no-email-notifications) — never send or configure email/notification delivery of any kind (if exists)
- [Reliability report](mem://features/reliability-report) — latest: /mnt/documents/02-reliability-risk-report-v2.md (2026-05-06); prior: 01-reliability-risk-report.md (2026-05-05) (if exists)
