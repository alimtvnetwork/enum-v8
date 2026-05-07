# CI/CD Issues — Index

> Summary of every CI/CD issue encountered across sessions. Each issue has its own file in `.lovable/cicd-issues/NN-name.md`.

| # | File | Title | Status |
|---|---|---|---|
| 01 | [`cicd-issues/01-go-mod-rename-bridge.md`](./cicd-issues/01-go-mod-rename-bridge.md) | `core-v9` ⇄ `core-v8` `go.mod` rename bridge breaks Go 1.25 builds | ✅ Resolved — Task W done, bridge removed |
| 02 | [`cicd-issues/02-internal-package-rule.md`](./cicd-issues/02-internal-package-rule.md) | Go's `internal/` rule rejects consumers under `enum-v7/...` due to declared module path mismatch | ✅ Resolved — root cause (#01) fixed |
| 03 | [`cicd-issues/03-pseudo-version-rejected.md`](./cicd-issues/03-pseudo-version-rejected.md) | Pseudo-version `v1.5.6-0.<date>-<sha>` rejected — needs non-existent `v1.5.5` predecessor tag | ✅ Documented; do not re-attempt |
| 04 | [`cicd-issues/04-no-email-notifications.md`](./cicd-issues/04-no-email-notifications.md) | Constraint: never configure email-based CI notifications | 🚫 Permanent constraint |
| 05 | [`cicd-issues/05-cross-repo-mirror-drift.md`](./cicd-issues/05-cross-repo-mirror-drift.md) | `cross-repo/core-v8/` can drift if main-repo CI changes aren't mirrored | 🔄 Procedural — handled per-change |
| 06 | [`cicd-issues/06-baseline-gate-seed.md`](./cicd-issues/06-baseline-gate-seed.md) | golangci-lint baseline gate is seed-then-gate; empty baseline = warnings only | ✅ Working as designed |
| 07 | [`cicd-issues/07-coverage-gate-60.md`](./cicd-issues/07-coverage-gate-60.md) | Coverage gate ≥ 60% enforced in `.github/workflows/ci.yml` | ✅ Working as designed |

## How to add a new CI/CD issue

1. Create `.lovable/cicd-issues/NN-short-name.md` (NN = next sequence number).
2. Use this structure:
   ```markdown
   # [Issue Title]

   ## Symptom
   ## Root Cause
   ## Fix / Workaround
   ## Status
   ## Related
   ```
3. Add a row to this index in the same operation (anti-corruption rule 3 + 6).
4. Do NOT duplicate an existing issue — extend the existing entry instead.
