# Workflow State

> Snapshot of where the project stands. Update at the end of every "Write memory" run.
> **Last updated:** 2026-05-06 (Cycle 23 — AB pass 5 on `11-versioning.md`: 8 more ❌ surfaced — `coreversion.Parse/Major/Minor/Patch/LessThan` and `versionindexes` "version eras" all fabricated; cumulative AB ❌ = 34; **WORST section in project at 18.2 %**).

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

- **AA. Spec-audit cycles** — All non-frozen / non-history `spec/` directories baselined. Next target options: AB pass 6 on `15-observability.md` (13 ❓), deep-probe `scripts/*.psm1` + `.github/workflows/*.yml` to resolve 14 workflow/script-internal ❓ from Cycles 16/17, or **build S-106 lint** (now MANDATORY).
- **AM. core-v9 API migration** — Reported `tests/creationtests` compile blocker fixed in sandbox.
  - 2026-05-06: Applied confirmed renames (53 `TypeName`→`SafeTypeName`, 6 `AnyToValueString`→`AnyTo.ValueString`, 1 `Any.ToFullNameValueString`→`AnyTo.ToFullNameValueString`, remaining `StringTo*` → `StringTo.*` sites, plus follow-on `codestack`, `coreversion`, and `StringRangesPtr` API updates). Verified with `go test -mod=mod ./tests/creationtests -run '^$' -count=0`.

## ⏳ Pending

- **AB. (in progress)** Upstream `core-v9 v1.5.8` cloned. **Done:** pass 1 §09 (66.7 %), pass 2 §07 (70.6 %), pass 3 §08 (33.3 %), pass 4 §10 (38.5 %), pass 5 §11 (**18.2 % — new worst**). **Pass-6 targets:** `15-observability.md` (13 ❓), `16-security.md` (13 ❓). Plus 14 workflow/script-internal ❓ and 5 audit-history ❓.
- **AC.** Re-audit §07 / §08 / §09 / §10 / §11 against consistency dimension. Now partially unblocked by Cycles 19+20+21+22+23 — re-run after AJ-01..27 land.
- **AJ.** **NEW open items: AJ-01..27** (all blocked by `spec/01-app/` freeze). Most-impactful: AJ-27 (rewrite entire `versionindexes` §2 — wrong purpose), AJ-15 (delete entire `coredynamic` §2 — package does not exist), and AJ-08..14 (rewrite almost all of `08-validators.md`). User decision needed on lifting freeze for an AB-fix waiver — but S-106 lint is now **MANDATORY** first so AJ rewrites don't introduce fresh fabrications (cumulative ❌ = 34, fabrication rate ~55 %, ~44 % CRITICAL).
- **AK.** New enum package creation (template validation).
- **AL.** Test coverage expansion.

## 🆕 New findings (Cycle 19)

- **C-CVS-11..15** — 5 ❌ contradictions in `spec/01-app/09-converters.md`. See `spec/07-code-vs-spec-audits/20-cycle19-AB-converters-promotion.md` for evidence and proposed rewrites. Severity: 4× HIGH, 1× CRITICAL (`typesconv` §2 fully fabricated).
- **M-CVS-01/02** — 2 stale Core-memory items, corrected in-turn (`enum-v3`→`enum-v4` module name; upstream `go.mod` rename declared complete with `replace` bridge removal).

## 🆕 New findings (Cycle 20)

- **C-CVS-16..20** — 5 ❌ contradictions in `spec/01-app/07-conditional-and-utilities.md`. See `spec/07-code-vs-spec-audits/21-cycle20-AB-conditional-and-utilities.md`. Severity: 3× HIGH, 2× CRITICAL (`namevalue.NewInstance` + `keymk.New.Compile` both fabricated). Pattern across cycles 19+20: ~25 % of all spec API claims authored against pre-rename `core-v8` are fabricated.

## 🆕 New findings (Cycle 21)

- **C-CVS-21..28** — 8 ❌ contradictions in `spec/01-app/08-validators.md` — almost the entire chapter is fabricated. Severity: 4× HIGH, 4× CRITICAL. Real `corevalidator/` exposes `LineValidator{LineNumber, TextValidator}` with `IsMatch(lineNumber, content, isCaseSensitive) bool`; the spec describes a fluent-builder + `Validate(input) Result` API that does not exist. **Cumulative fabrication rate is now 41 %** across 3 audited sections.
- **Recommendation:** S-106 lint should land **before any AJ rewrite** — without it, the same author pattern that produced 18 ❌ will likely produce more.

## 🆕 New findings (Cycle 22)

- **C-CVS-29..36** — 8 ❌ contradictions in `spec/01-app/10-reflection-and-dynamic.md`. Severity: 3× HIGH, 5× CRITICAL. The **entire `coredynamic` package documented in §2 does not exist** in upstream `core-v9 v1.5.8` — `coredynamic/` directory is absent and `grep -rln coredynamic` returns zero source files. `reflectcore` (§3) is a thin re-export shim over `internal/reflectinternal` (15 re-exported symbols), not the predicate library described — `IsPointer`/`WalkFields`/`GetTag`/`DerefAll` are all fabricated. The "internal off-limits" framing in §4 is also misleading because `reflectcore/vars.go` publicly re-exports the very symbols it claims are walled off. See `spec/07-code-vs-spec-audits/23-cycle22-AB-reflection-and-dynamic.md`. Spawned **AJ-15..20**. **Cumulative fabrication rate is now ~45 %** across 4 audited sections.

## 🆕 New findings (Cycle 23)

- **C-CVS-37..43** — 8 ❌ contradictions in `spec/01-app/11-versioning.md` — **new worst-drift section in the project** at 18.2 % verifiable. Severity: 4× CRITICAL, 3× HIGH, 1× LOW. Constructor (`Parse`), accessors (`Major()`/`Minor()`/`Patch()`), comparators (`LessThan`/`Equal`/`GreaterThanOrEqual`), error-wrapping rationale (`errcore.FailedToConvertType`), and package path (`versionindexes/`) all fabricated. **C-CVS-43 is conceptually unique:** `versionindexes` does not enumerate "version eras" (`V1=1, V2=2, V8=8`) — it enumerates **version-component positions** (`Major=0, Minor=1, Patch=2, Build=3, Invalid=4`), making the entire §2 framing wrong about the package's *purpose*. Real `coreversion.Version` is a public-field struct created via `coreversion.New.Default(s)` (no error) and compared via package-level `coreversion.Compare(left, right *Version)`. See `spec/07-code-vs-spec-audits/24-cycle23-AB-versioning.md`. Spawned **AJ-21..27**. **Cumulative fabrication rate is now ~55 %** across 5 audited sections.

## ⏭️ Manual user action (parked)

- **A.** Push `cross-repo/core-v8/` mirror to its upstream GitHub repo.

## Next logical step

1. **Build S-106** (`scripts/spec-api-check.psm1`) — lint to catch the fabrication pattern before AJ rewrites. **MANDATORY next given 34 ❌ accumulated and ~55 % fabrication rate.** OR
2. **AB pass 6** — Cycle 24 on `15-observability.md` (13 ❓). OR
3. **User decision: lift `spec/01-app/` freeze** for AJ-01..27 patches (highly risky without S-106). OR
4. **AK** — New enum package creation / template validation. OR
5. **AL** — Test-coverage expansion.
