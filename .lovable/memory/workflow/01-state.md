# Workflow State

> Snapshot of where the project stands. Update at the end of every "Write memory" run.
> **Last updated:** 2026-05-06 (Cycle 19 — AB pass 1 on `09-converters.md`: first ❓→✅/❌ promotions, 5 NEW HIGH/CRITICAL contradictions surfaced).

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
- **Cycle 18** — `spec/02-app-issues/` (11 files, 402 lines) baselined and closed at **100 % verifiable** (21 ✅ / 5 ❓ audit-history). Resolved **D-CVS-56 → D-CVS-60** (5 LOW): stale README index + 4 upstream-vs-`enum-v4` scope footnotes. **🎉 Marks Task AH Done** — every directory under `spec/` outside immutable history folders is now baselined. Spec changelog → **spec-v0.33.0**. Audit file: `spec/07-code-vs-spec-audits/19-cycle18-app-issues.md`.
- **AH** — Cross-`spec/` cleanup sweep COMPLETE (Cycles 11/12 → 15 → 16 → 17 → 18).
- **Reliability risk report v2** produced → `/mnt/documents/02-reliability-risk-report-v2.md` (2026-05-06; supersedes v1)
- **Suggestions tracker** consolidated → `.lovable/memory/suggestions/01-suggestions-tracker.md`
- **Pending issues** consolidated → `.lovable/memory/pending-issues/01-all-pending-issues.md`
- **Plan.md** updated with phases, cycle plan, and next-task selection
- **W.** Upstream `core-v9` `go.mod` rename + `v1.5.8` tag — ✅ Done (2026-05-05)
- **AG.** Drop `replace` bridge and pin clean `core-v9 v1.5.8` — ✅ Done (2026-05-05)
- **core-v9 API investigation** — Mapped converter/coredynamic API changes (2026-05-06). See `.lovable/memory/06-core-v9-api-migration.md`.

## 🔄 In Progress

- **AA. Spec-audit cycles** — All non-frozen / non-history `spec/` directories baselined. Next target options: deep-probe `scripts/*.psm1` + `.github/workflows/*.yml` to resolve 14 workflow/script-internal ❓ from Cycles 16/17, or wait for AB.
- **AM. core-v9 API migration** — Reported `tests/creationtests` compile blocker fixed in sandbox.
  - 2026-05-06: Applied confirmed renames (53 `TypeName`→`SafeTypeName`, 6 `AnyToValueString`→`AnyTo.ValueString`, 1 `Any.ToFullNameValueString`→`AnyTo.ToFullNameValueString`, remaining `StringTo*` → `StringTo.*` sites, plus follow-on `codestack`, `coreversion`, and `StringRangesPtr` API updates). Verified with `go test -mod=mod ./tests/creationtests -run '^$' -count=0`.

## ⏳ Pending

- **AB. (in progress)** Upstream `core-v9 v1.5.8` cloned to `/tmp/core-v9-upstream` (35 MB, 3 172 .go files). Pass 1 done on `09-converters.md` (10 ✅ / 5 ❌ / 8 ❓ remaining). **Pass 2 targets:** `07-conditional-and-utilities.md` (17 ❓), `08-validators.md` (18 ❓), `10-reflection-and-dynamic.md` (15 ❓), `11-versioning.md` (11 ❓), `15-observability.md` (13 ❓), `16-security.md` (TBD ❓). Plus the 14 workflow/script-internal ❓ and 5 audit-history ❓.
- **AC.** Re-audit §07 and §09 against the spec-internal-consistency dimension. Now partially unblocked for §09 by Cycle 19 — re-run after AJ-01..03 land.
- **AJ.** **NEW open items: AJ-01, AJ-02, AJ-03** (all blocked by `spec/01-app/` freeze): rewrite `09-converters.md` §1.1, §2 + §4.3, and §1.3/§4 PrettyJson callsites against the verified upstream API. User decision needed on lifting freeze for an AB-fix waiver.
- **AK.** New enum package creation (template validation).
- **AL.** Test coverage expansion.

## 🆕 New findings (Cycle 19)

- **C-CVS-11..15** — 5 ❌ contradictions in `spec/01-app/09-converters.md`. See `spec/07-code-vs-spec-audits/20-cycle19-AB-converters-promotion.md` for evidence and proposed rewrites. Severity: 4× HIGH, 1× CRITICAL (`typesconv` §2 fully fabricated).
- **M-CVS-01/02** — 2 stale Core-memory items, corrected in-turn (`enum-v3`→`enum-v4` module name; upstream `go.mod` rename declared complete with `replace` bridge removal).

## ⏭️ Manual user action (parked)

- **A.** Push `cross-repo/core-v8/` mirror to its upstream GitHub repo.

## Next logical step

1. **AB pass 2** — apply same promotion-and-grep pass to Cycles 5, 6, 8, 9 (each is expected to surface a similar number of ❌ findings because all four were authored against the pre-rename `core-v8` surface). OR
2. **User decision: lift `spec/01-app/` freeze** for AJ-01..03 patches landing the corrected `09-converters.md` API. OR
3. **AK** — New enum package creation / template validation.
4. **AL** — Test-coverage expansion.
