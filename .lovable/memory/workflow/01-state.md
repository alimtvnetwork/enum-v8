# Workflow State

> Snapshot of where the project stands. Update at the end of every "Write memory" run.
> **Last updated:** 2026-05-06 (Cycle 17 closed — `spec/04-tooling/` baselined at 100 %).

## ✅ Done

- **Cycles 1–14** of the `spec/01-app/` audit (all 14 numbered files processed).
  - 12 sections at 100% verifiable (§03, §04, §05, §06, §08, §10, §11, §12, §13, §14, §15, §16).
  - 2 sections baseline-only (§07, §09) — no verifiable subset until Task AB lands.
  - 10 contradiction findings resolved (C-CVS-01 … C-CVS-10).
  - 43 drift findings resolved (D-CVS-01 … D-CVS-43).
  - New audit dimension introduced: **spec-internal consistency** (Cycle 13).
  - `spec/01-app/` directory cleared of `tests/integratedtests/` (genuine clean as of Cycle 12).
- **Cycle 15** — `spec/06-testing-guidelines/` directory baselined and closed at **100 % verifiable** (32 claims sampled across 10 files; 22 ✅ / 10 ❓ pending AB). Resolved D-CVS-43 (LOW) via README + `01-folder-structure.md` consumer-coverage callouts (same pattern as Cycle 12). Spec changelog → **spec-v0.29.0**. Audit file: `spec/07-code-vs-spec-audits/16-cycle15-testing-guidelines.md`.
- **AI** — `spec/01-app/` declared **🧊 DRIFT-FROZEN** in `spec/CHANGELOG.md` as **spec-v0.30.0** (2026-05-06). Allowed future edits limited to: AB-driven ❓→✅ promotions, AC re-audit of §07/§09, upstream-API-change additions (paired with new audit cycle row), typo/formatting fixes. Scoreboard top-line carries the 🧊 marker.
- **Cycle 16** — `spec/03-powershell-test-run/` (9 files, 2 519 lines) baselined and closed at **100 % verifiable** (22 ✅ / 6 ❓ runner-internal). Resolved D-CVS-44 → D-CVS-48 (5 LOW) via consumer-coverage callouts. Folds in **AH** debt for this directory. Spec changelog → **spec-v0.31.0**. Audit file: `spec/07-code-vs-spec-audits/17-cycle16-powershell-test-run.md`.
- **Cycle 17** — `spec/04-tooling/` (10 files, 2 553 lines) baselined and closed at **100 % verifiable** (22 ✅ / 8 ❓ workflow-internal). Resolved **D-CVS-49 → D-CVS-55** (7 LOW): 2 broken `cross-repo/core-v9/` paths + 1 AH-tracked `tests/integratedtests/` row + 4 stale `enum-v2`/`cross-repo/core-v9` tokens (template comment line 80 carried both). Each fix carries a Core-memory clarification on the historical `cross-repo/core-v8/` directory name. Spec changelog → **spec-v0.32.0**. Audit file: `spec/07-code-vs-spec-audits/18-cycle17-tooling.md`.
- **Reliability risk report v2** produced → `/mnt/documents/02-reliability-risk-report-v2.md` (2026-05-06; supersedes v1)
- **Suggestions tracker** consolidated → `.lovable/memory/suggestions/01-suggestions-tracker.md`
- **Pending issues** consolidated → `.lovable/memory/pending-issues/01-all-pending-issues.md`
- **Plan.md** updated with phases, cycle plan, and next-task selection
- **W.** Upstream `core-v9` `go.mod` rename + `v1.5.8` tag — ✅ Done (2026-05-05)
- **AG.** Drop `replace` bridge and pin clean `core-v9 v1.5.8` — ✅ Done (2026-05-05)
- **core-v9 API investigation** — Mapped converter/coredynamic API changes (2026-05-06). See `.lovable/memory/06-core-v9-api-migration.md`.

## 🔄 In Progress

- **AA. Spec-audit cycles** — Next target: Cycle 18 → `spec/02-app-issues/` (carries final residual AH debt for `02-internal-package-coverage-policy.md`).
- **AH. Cross-`spec/` cleanup sweep** — `spec/03` cleared Cycle 16; `spec/04` cleared Cycle 17; final residual: `spec/02-app-issues/02-internal-package-coverage-policy.md` (folds into Cycle 18).
- **AM. core-v9 API migration** — Reported `tests/creationtests` compile blocker fixed in sandbox.
  - 2026-05-06: Applied confirmed renames (53 `TypeName`→`SafeTypeName`, 6 `AnyToValueString`→`AnyTo.ValueString`, 1 `Any.ToFullNameValueString`→`AnyTo.ToFullNameValueString`, remaining `StringTo*` → `StringTo.*` sites, plus follow-on `codestack`, `coreversion`, and `StringRangesPtr` API updates). Verified with `go test -mod=mod ./tests/creationtests -run '^$' -count=0`.

## ⏳ Pending

- **AB.** Fetch upstream `core-v9` source to resolve 158 ❓ claims (148 in `spec/01-app/` + 10 in `spec/06-testing-guidelines/`). Plus 14 workflow/script-internal ❓ (6 from Cycle 16 + 8 from Cycle 17) needing direct probes of `scripts/*.psm1` and `.github/workflows/*.yml` (separate sub-task).
- **AC.** Re-audit §07 and §09 against the spec-internal-consistency dimension.
- **AJ.** Implement any follow-on fixes from Cycle 15/16+ findings (currently none open).
- **AK.** New enum package creation (template validation).
- **AL.** Test coverage expansion.

## ⏭️ Manual user action (parked)

- **A.** Push `cross-repo/core-v8/` mirror to its upstream GitHub repo.

## Next logical step

1. **AA/Cycle 18** — Audit `spec/02-app-issues/` (carries final residual AH debt; closes the cross-`spec/` sweep entirely).
2. **AB** — Fetch upstream `core-v9` source for ❓ verification (unblocks 158 claims + Cycle 7/9 closure + 14 workflow/script-internal claims).
3. **AK** — New enum package creation / template validation.
