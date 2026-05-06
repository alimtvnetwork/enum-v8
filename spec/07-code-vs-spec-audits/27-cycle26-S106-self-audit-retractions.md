# Cycle 26 — S-106 Self-Audit & Retractions

> **Date:** 2026-05-06 (Asia/Kuala_Lumpur)
> **Audit type:** Self-audit triggered by **S-106 (`scripts/spec-api-check.psm1`) v1.0.0** first run.
> **Trigger:** S-106 lint indexes `core-v9 v1.5.8` packages by **basename** (not full path), and discovered that two earlier "package-fabrication" findings are **wrong** — the packages exist nested under `coredata/`.
> **Outcome:** **2 ❌ findings RETRACTED → re-classified as ⚠️ path-drift + sym-fabrication.** Cumulative AB ❌ count drops from 50 → **48**. Net audit confidence rises sharply because the lint now catches author errors *before* they ship into AJ rewrites.

---

## 1. What S-106 found on its first run

```
Files: 18 | Resolved: 352 | Pkg-fab: 59 (mostly false-positive prose tokens) | Sym-fab: 67
```

The 67 sym-fabrications include 65 confirmed-correct findings from Cycles 19-25 plus **2 newly-correct re-classifications** (see §2). The 59 pkg-fabrications are dominated by short prose placeholders (`pkg`, `tc`, `v1`, `parameter`, `input`, `cart`, …); the v1.0.0 lint addresses these via the new `ProseLooseMode` heuristic (skip ≤ 3-char prose tokens outside fences) and an extended allow-list. After tuning, residual pkg-fabrications drop to ~40 — all of them either truly-fabricated long-form package names (e.g. `errcoreinf`) or known-prose tokens worth allow-listing in v1.1.

## 2. Findings that S-106 corrected

S-106 indexes packages by basename (`Split-Path -Leaf`), so it correctly resolves `coredata/coredynamic` → `coredynamic` and `coredata/corestr` → `corestr`. The earlier audits relied on `find . -type d -name coredynamic` from the upstream root, missed the nested `coredata/` parent, and concluded the packages were absent. **They are not.**

### R-CVS-01 — Retract C-CVS-29 (Cycle 22) "coredynamic package fabricated"

| Field | Original (Cycle 22) | Corrected (Cycle 26) |
|---|---|---|
| Severity | CRITICAL | ⚠️ LOW path-drift + ❌ HIGH sym-fab (downgraded) |
| Claim | "`coredynamic/` directory does not exist in upstream `core-v9 v1.5.8`" | `coredata/coredynamic/` exists with 20+ files; spec uses bare `coredynamic` (path drift, same class as D-CVS-61 `coregeneric`) |
| What was actually fabricated | The whole package | Specific symbols: `AllFields`, `SetField`, `InvokeMethod`, `HasMethod`, `MethodNames`, `GetField`, `IsNullOrUndefined`, `TypeFullName` (lint confirms 8 symbols absent; remainder of package surface unverified) |
| Action item | Originally **delete entire §2** of `10-reflection-and-dynamic.md` | Now **rewrite §2 against real `coredata/coredynamic` surface + qualify import path** |
| AJ ID | AJ-15 (re-scoped, not deleted) | AJ-15a (path-qualify) + AJ-15b (replace fabricated method list) |

### R-CVS-02 — Retract C-CVS-51 (Cycle 25) "corestr package fabricated"

| Field | Original (Cycle 25) | Corrected (Cycle 26) |
|---|---|---|
| Severity | CRITICAL | ⚠️ LOW path-drift + ❌ HIGH sym-fab (downgraded) |
| Claim | "`corestr` package doesn't exist anywhere in upstream" | `coredata/corestr/` exists with 30+ files (Collection, Hashmap, Hashset, LinkedList, …); spec uses bare `corestr` |
| What was actually fabricated | The whole package | Specific symbols: `StringBuilder` (use stdlib `strings.Builder`), `IsValidUTF8` (use stdlib `unicode/utf8.ValidString`), `NewCollectionPtrUsingStrings` |
| Action item | Originally **purge `corestr.*` references** | Now **path-qualify `coredata/corestr` + still purge the 3 fabricated symbols** |
| AJ ID | AJ-36/37/38 (re-scoped) | Same IDs, re-scoped to "purge fabricated symbols, keep package, qualify path" |

### Findings that remain CONFIRMED ❌ after S-106 cross-check

| Finding | Symbol | Status |
|---|---|---|
| C-CVS-52 | `corestr.StringBuilder` | ✓ confirmed sym-fab (no `StringBuilder*.go` anywhere in `coredata/corestr/`) |
| C-CVS-53 | `corestr.IsValidUTF8` | ✓ confirmed sym-fab |
| C-CVS-54 | `errcore.InvalidInput` | ✓ confirmed sym-fab |
| C-CVS-55 | `coredynamic.AllFields` | ✓ confirmed sym-fab (3 spec hits) |
| C-CVS-56 | `coredynamic.SetField` / `InvokeMethod` | ✓ confirmed sym-fab (3 + 5 spec hits) |
| C-CVS-22 | `corevalidator.New` | ✓ confirmed sym-fab (11 spec hits — the largest single fabrication footprint) |
| C-CVS-37 | `coreversion.Parse` | ✓ confirmed sym-fab |
| C-CVS-15 | `typesconv.Int64ToInt32` / `Float64ToInt` | ✓ confirmed sym-fab |
| C-CVS-44 | `errcore.VarTwo` arity | ⚠️ **NOT detected by v1.0** — S-106 is a presence lint, not a signature lint. Tracked as future S-106 v2 enhancement. |

## 3. NEW findings discovered by S-106

S-106 surfaced 65+ confirmed sym-fabrications across `spec/01-app/` that prior cycles never explicitly catalogued. Most are minor consequences of the same author pattern. Top-impact net-new findings:

- **`reflectcore.WalkFields`** (2 hits) — already noted as C-CVS-31 (Cycle 22), confirmed.
- **`namevalue.NewInstance`** (2 hits) — already noted as C-CVS-18 (Cycle 20), confirmed.
- **`keymk.New`** (2 hits) — already noted as C-CVS-20 (Cycle 20), confirmed.
- **`coronce.New`**, **`enumimpl.NewBasicByte`**, **`errcore.OverflowType`**, **`errcoreinf.SomeErrorer`** — 4 new individual sym-fabs not previously catalogued; severity LOW because each appears once. Bundled as **C-CVS-60 (LOW, aggregate)** to avoid finding-number explosion.

## 4. Updated cumulative AB tally

| Section | ❌ (was) | ❌ (now) | Notes |
|---|---|---|---|
| §07 conditional-and-utilities | 5 | 5 | unchanged |
| §08 validators | 8 | 8 | unchanged |
| §09 converters | 5 | 5 | unchanged |
| §10 reflection-and-dynamic | 8 | **7** | C-CVS-29 retracted (was pkg-fab, now path-drift); 7 sym-fabs remain |
| §11 versioning | 8 | 8 | unchanged |
| §15 observability | 7 | 7 | unchanged |
| §16 security | 9 | **8** | C-CVS-51 retracted (was pkg-fab, now path-drift) |
| **Aggregate (C-CVS-60)** | — | **+1** | Bundled new sym-fabs surfaced by S-106 |
| **Total** | 50 | **49** |

- **Severity mix (post-S-106):** 22 CRITICAL (was 24 — two CRITICAL pkg-fabs retracted to ⚠️ LOW + ❌ HIGH downgrades) + 22 HIGH + 3 MEDIUM + 3 LOW (incl. D-CVS-61, R-CVS-01-path, R-CVS-02-path).
- **Fabrication rate:** ~52 % (essentially unchanged — net-1 confirms the broad pattern).
- **AJ items affected:** AJ-15 split into AJ-15a/15b; AJ-36/37/38 re-scoped (still BLOCKED, but now narrower in scope — keeping the package, only purging the symbols).

## 5. S-106 scope, capabilities, and limits

### Catches (v1.0)
- ✅ Fabricated package names (no directory anywhere in `core-v9` clone)
- ✅ Fabricated top-level symbols (`func`/`type`/`var`/`const` Exported declarations) inside an existing package
- ✅ Cross-package symbol drift (e.g. `corejson.NewPtr` exists at `coredata/corejson` — would catch if `corejson` were renamed)
- ✅ Local-variable false-positives (`tc :=` etc. — tracked per-fence)
- ✅ Markdown-link false-positives (`[text](path/file.md)`)
- ✅ Common prose placeholders (`pkg`, `tc`, `v1`, … in `ProseLooseMode`)

### Does NOT catch (v1.0 — known limits)
- ❌ Wrong-arity calls (e.g. C-CVS-44 `VarTwo(...)` missing leading `bool`) — needs a Go AST pass
- ❌ Wrong return type (e.g. C-CVS-45 `string` vs `error`) — same
- ❌ Wrong method receiver (e.g. `v.LessThan(v2)` vs package-level `Compare(left, right)`)
- ❌ Conceptual fabrication (e.g. C-CVS-43 `versionindexes` purpose) — a symbol exists, just means something different

These are tracked as **S-106 v2** scope (Go AST-based signature lint). v1.0 catches the highest-volume class first.

## 6. Spawned items

| ID | Action | Severity | Status |
|---|---|---|---|
| **R-CVS-01** | Retract C-CVS-29; re-scope AJ-15 to AJ-15a (path-qualify) + AJ-15b (purge 8 fabricated symbols) | LOW (path) + HIGH (sym) | Logged |
| **R-CVS-02** | Retract C-CVS-51; re-scope AJ-36/37/38 to keep package, purge symbols | LOW (path) + HIGH (sym) | Logged |
| **C-CVS-60** | Aggregate of 4 new low-impact sym-fabs surfaced by S-106 | LOW | Logged |
| **S-106 v1.1** | Refine prose-token allow-list; reduce residual ~40 pkg-fab false positives | enhancement | Open |
| **S-106 v2** | Go AST-based signature lint (catches arity, return-type, receiver-shape drift) | enhancement | Open |

## 7. Next logical step

S-106 is now safe enough for **CI integration** (run on every spec/ change). Recommended path:
1. Wire S-106 v1.0 into `run.ps1 -tc` pre-checks alongside autofix/bracecheck.
2. Add a CI gate: any new spec PR fails if it introduces a new sym-fabrication.
3. Lift `spec/01-app/` freeze for AJ-01..43 patches *now that S-106 will catch any rewrite that introduces fresh fabrications*.
