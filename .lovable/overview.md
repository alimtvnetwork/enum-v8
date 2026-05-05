# Project Overview — enum-v3

## Identity

- **Module path:** `github.com/alimtvnetwork/enum-v3` (renamed from `enum-v1`)
- **Type:** Go enum library with PowerShell + Python tooling
- **Frontend:** A React/Vite shell exists but is incidental — all real work happens in Go packages, PowerShell scripts, and Python CI guards.

## Core dependency

- **Import path used in source:** `github.com/alimtvnetwork/core-v9` (renamed from `core-v8`)
- **Upstream `go.mod` reality:** the upstream repo still declares `module github.com/alimtvnetwork/core-v8`. A `replace` directive in `enum-v3/go.mod` bridges the path mismatch.
- **Bridge limitation:** Go's `internal/` rule is enforced against the cached module's declared path (`core-v8`), so any `core-v9` package that transitively imports an `internal/` package is rejected for consumers under `enum-v3/...`. Only fix is updating upstream `go.mod` and tagging a release (Task **W**).

## Directories the AI must respect

- `cross-repo/core-v8/` — mirror of a separate upstream repo. **Never** rename `core-v8` → `core-v9` here. **Never** rewrite `enum-v1` → `enum-v3` here.
- `tests/creationtests/` — actual test layout. The string `tests/integratedtests/` in any spec file is **stale** (audit finding C-CVS-01).

## Hard prohibitions

- **No emails.** Never configure Dependabot recipients, SMTP, or any email-based notification flow. User explicitly rejects all email flows.
- **No `core-v8` reintroduction** outside `cross-repo/core-v8/`.
- **No `tests/integratedtests/`** in new spec content.

## Workflow

- User drives via incremental "Next" commands.
- Tasks are tracked with letter IDs (A, B, C, …, AI, AJ, …) carried across turns.
- Every reply MUST end with the remaining-task list.
