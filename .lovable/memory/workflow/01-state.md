# Workflow State

> Snapshot of where the project stands. Update at the end of every "Write memory" run.
> **Last updated:** 2026-05-05 (post reliability report, pre-implementation).

## ✅ Done

- **Cycles 1–14** of the `spec/01-app/` audit (all 14 numbered files processed).
- 12 sections at 100% verifiable (§03, §04, §05, §06, §08, §10, §11, §12, §13, §14, §15, §16).
- 2 sections baseline-only (§07, §09) — no verifiable subset until Task AB lands.
- 10 contradiction findings resolved (C-CVS-01 … C-CVS-10).
- 43 drift findings resolved (D-CVS-01 … D-CVS-43).
- New audit dimension introduced: **spec-internal consistency** (Cycle 13).
- `spec/01-app/` directory cleared of `tests/integratedtests/` (genuine clean as of Cycle 12).
- **Reliability risk report** produced → `/mnt/documents/01-reliability-risk-report.md`
- **Suggestions tracker** consolidated → `.lovable/memory/suggestions/01-suggestions-tracker.md`
- **Pending issues** consolidated → `.lovable/memory/pending-issues/01-all-pending-issues.md`
- **Plan.md** updated with phases, cycle plan, and next-task selection

## 🔄 In Progress

- **AA. Spec-audit cycles** — Cycle 15 target: `spec/06-testing-guidelines/`.
- **AH. Cross-`spec/` cleanup sweep** — folded into upcoming directory audits.

## ⏳ Pending

- **AG.** Drop the `replace` bridge once Task W lands.
- **AB.** Fetch upstream `core-v9` source to resolve 148 ❓ claims.
- **AC.** Re-audit §07 and §09 against the spec-internal-consistency dimension.
- **AI.** Mark `spec/01-app/` as frozen for code-vs-spec drift in `spec/CHANGELOG.md`.
- **AJ.** Implement spec fixes from Cycle 15 findings.
- **AK.** New enum package creation (template validation).
- **AL.** Test coverage expansion.

## 🚫 Blocked

- **W.** Upstream `core-v9` `go.mod` rename + `v1.5.8` tag — **manual upstream action required**.

## ⏭️ Manual user action (parked)

- **A.** Push `cross-repo/core-v8/` mirror to its upstream GitHub repo.

## Next logical step

User picks from: AA/Cycle 15, AI (freeze), S-003+S-004 (quick path fixes), or W (upstream fix).
