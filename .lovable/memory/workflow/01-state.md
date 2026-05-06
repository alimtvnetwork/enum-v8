# Workflow State

> Snapshot of where the project stands. Update at the end of every "Write memory" run.
> **Last updated:** 2026-05-06 (Task AM reported `creationtests` compile blocker fixed).

## ✅ Done

- **Cycles 1–14** of the `spec/01-app/` audit (all 14 numbered files processed).
  - 12 sections at 100% verifiable (§03, §04, §05, §06, §08, §10, §11, §12, §13, §14, §15, §16).
  - 2 sections baseline-only (§07, §09) — no verifiable subset until Task AB lands.
  - 10 contradiction findings resolved (C-CVS-01 … C-CVS-10).
  - 43 drift findings resolved (D-CVS-01 … D-CVS-43).
  - New audit dimension introduced: **spec-internal consistency** (Cycle 13).
  - `spec/01-app/` directory cleared of `tests/integratedtests/` (genuine clean as of Cycle 12).
- **Reliability risk report v2** produced → `/mnt/documents/02-reliability-risk-report-v2.md` (2026-05-06; supersedes v1)
- **Suggestions tracker** consolidated → `.lovable/memory/suggestions/01-suggestions-tracker.md`
- **Pending issues** consolidated → `.lovable/memory/pending-issues/01-all-pending-issues.md`
- **Plan.md** updated with phases, cycle plan, and next-task selection
- **W.** Upstream `core-v9` `go.mod` rename + `v1.5.8` tag — ✅ Done (2026-05-05)
- **AG.** Drop `replace` bridge and pin clean `core-v9 v1.5.8` — ✅ Done (2026-05-05)
- **core-v9 API investigation** — Mapped converter/coredynamic API changes (2026-05-06). See `.lovable/memory/06-core-v9-api-migration.md`.

## 🔄 In Progress

- **AA. Spec-audit cycles** — Cycle 15 target: `spec/06-testing-guidelines/`.
- **AH. Cross-`spec/` cleanup sweep** — folded into upcoming directory audits.
- **AM. core-v9 API migration** — Reported `tests/creationtests` compile blocker fixed in sandbox.
  - 2026-05-06: Applied confirmed renames (53 `TypeName`→`SafeTypeName`, 6 `AnyToValueString`→`AnyTo.ValueString`, 1 `Any.ToFullNameValueString`→`AnyTo.ToFullNameValueString`, remaining `StringTo*` → `StringTo.*` sites, plus follow-on `codestack`, `coreversion`, and `StringRangesPtr` API updates). Verified with `go test -mod=mod ./tests/creationtests -run '^$' -count=0`.

## ⏳ Pending

- **AB.** Fetch upstream `core-v9` source to resolve 148 ❓ claims.
- **AC.** Re-audit §07 and §09 against the spec-internal-consistency dimension.
- **AI.** Mark `spec/01-app/` as frozen for code-vs-spec drift in `spec/CHANGELOG.md`.
- **AJ.** Implement spec fixes from Cycle 15 findings.
- **AK.** New enum package creation (template validation).
- **AL.** Test coverage expansion.

## ⏭️ Manual user action (parked)

- **A.** Push `cross-repo/core-v8/` mirror to its upstream GitHub repo.

## Next logical step

1. **AA/Cycle 15** — Audit `spec/06-testing-guidelines/`.
2. **AI** — Mark `spec/01-app/` as frozen (quick win).
3. **AB** — Fetch upstream `core-v9` source for ❓ verification.
