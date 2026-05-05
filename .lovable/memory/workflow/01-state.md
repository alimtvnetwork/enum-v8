# Workflow State

> Snapshot of where the project stands. Update at the end of every "Write memory" run.
> **Last updated:** 2026-05-05 (after Cycle 14).

## ✅ Done

- **Cycles 1–14** of the `spec/01-app/` audit (all 14 numbered files processed).
- 12 sections at 100% verifiable (§03, §04, §05, §06, §08, §10, §11, §12, §13, §14, §15, §16).
- 2 sections baseline-only (§07, §09) — no verifiable subset until Task AB lands.
- 10 contradiction findings resolved (C-CVS-01 … C-CVS-10).
- 43 drift findings resolved (D-CVS-01 … D-CVS-43).
- New audit dimension introduced: **spec-internal consistency** (Cycle 13).
- `spec/01-app/` directory cleared of `tests/integratedtests/` (genuine clean as of Cycle 12).

## 🔄 In Progress

- **AA. Spec-audit cycles** — pivoting from `spec/01-app/` (done) to next directory. Cycle 15 target candidate: `spec/06-testing-guidelines/`.
- **AH. Cross-`spec/` cleanup sweep** — to be folded into upcoming directory audits.

## ⏳ Pending

- **AG.** Drop the `replace` bridge once Task W lands.
- **AB.** Fetch upstream `core-v9` source to resolve 148 ❓ claims.
- **AC.** Re-audit §07 and §09 against the spec-internal-consistency dimension.
- **AI.** Mark `spec/01-app/` as frozen for code-vs-spec drift in `spec/CHANGELOG.md`.

## 🚫 Blocked

- **W.** Upstream `core-v9` `go.mod` rename + `v1.5.8` tag — **manual upstream action required**. Until this lands, Go 1.25 builds that touch any `core-v9` package transitively importing `internal/` will fail with `used for two different module paths`.

## ⏭️ Manual user action (parked)

- **A.** Push `cross-repo/core-v8/` mirror to its upstream GitHub repo. Credential-bound; AI mirrors files but never pushes.

## Next logical step

User says **next** → start Cycle 15 on `spec/06-testing-guidelines/` (recommended; combines audit + AH stale-path sweep). Alternative: pick `spec/02-app-issues/`, `spec/03-powershell-test-run/`, or `spec/04-tooling/`.
