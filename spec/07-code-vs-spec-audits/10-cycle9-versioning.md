# Cycle 9 — `01-app/11-versioning.md`

> **Date**: 2026-05-05 (Asia/Kuala_Lumpur)
> **Spec audited**: [`spec/01-app/11-versioning.md`](../01-app/11-versioning.md)
> **Auditor**: Lovable agent (loop AA-cycle9)
> **Status**: **baseline + closed (4 fixes applied)**

---

## 1. Method

Each numbered section's claims are classified as ✅ Match / ⚠️ Drift / ❌ Contradiction / ❓ Unverifiable, with verification commands run from repo root:

```bash
rg -l "core-v9/(coreversion|versionindexes)" --type go
rg -n "coreversion\.|versionindexes\." --type go
rg -n "golang.org/x/mod/semver"
ls .release .lovable mem 2>&1
grep -n "integratedtests\|core-v9 . core-v9" spec/01-app/11-versioning.md
ls cross-repo/core-v9/{coreversion,versionindexes} 2>/dev/null
head -3 go.mod
```

Results:
- **Zero importers** of `coreversion` or `versionindexes` in `enum-v3`.
- **Zero call sites** of any documented symbol (9 probed).
- **Zero direct `golang.org/x/mod/semver` use** (anti-pattern from §6 Common Mistakes — rule honoured).
- `.release/`, `.lovable/`, `mem/` directories **do not exist on disk** in `enum-v3` — citations in §3, §5, §11 source-attribution are stale.
- Two stale `core-v9 → core-v9` mojibake artifacts (§3 line 95, §4 line 112) — left over from the bulk `v8`→`v9` rename and now nonsensical (the prose describes the **historical** era bump, which legitimately should read `core-v8` → `core-v9`).
- One stale `tests/integratedtests/` reference (§4 line 108) — same pattern already filed as C-CVS-01 / D-CVS-17 / D-CVS-26.
- One stale `versionindexes.V8` comment (§2 line 59) — says `// 8 (current era — core-v9)` which conflates the legacy V8 index with the current core-v9 era.
- `cross-repo/core-v9/{coreversion,versionindexes}` directories do not exist.

---

## 2. Claim inventory

| #  | Spec § | Claim | Verdict | Note |
|----|--------|-------|---------|------|
| 1  | §1 | `coreversion.Parse(s) (Version, error)` | ❓ | No consumer |
| 2  | §1 | `Version.{Major,Minor,Patch}() int` accessors | ❓ | No consumer |
| 3  | §1 | `Version.{LessThan,Equal,GreaterThanOrEqual}(other) bool` | ❓ | No consumer |
| 4  | §1 | `Version.String() string` | ❓ | No consumer |
| 5  | §1 (rationale) | Wraps stdlib errors via `errcore.FailedToConvertType` | ❓ | `FailedToConvertType` already ❓ in Cycle 2/7 |
| 6  | §2 | `versionindexes.V1..V8` integer constants | ❓ | No consumer |
| 7  | §2 line 59 (comment) | "`V8 // 8 (current era — core-v9)`" | ⚠️ | **Drift D-CVS-30** — V8 is the *legacy* era index; current era is V9 (core-v9). Comment is contradictory. |
| 8  | §3 (CRITICAL rule) | "Code changes must bump at least minor version. Never touch the `.release/` folder." | ✅ | Rule honoured — `.release/` does not exist in repo, so no edits possible (vacuously satisfied). Documented as such. |
| 9  | §3 (citation) | Rule sourced from `.lovable/user-preferences` line 8 | ⚠️ | **Drift D-CVS-31** — `.lovable/` directory does not exist in `enum-v3`. Citation is stale; rule lives in `mem://index.md` Core only. |
| 10 | §3 line 86 | "Bug fix → minor (project rule overrides standard semver patch)" | ❓ | Behavioural rule — no in-repo enforcer |
| 11 | §3 line 95 | "Update `go.mod` major version path only on major bump (e.g. `core-v9` → `core-v9`)" | ❌ | **Contradiction C-CVS-09a** — nonsensical (same string both sides); must read `core-v8` → `core-v9` |
| 12 | §4 line 105 | "imports from `github.com/alimtvnetwork/core-v9/<pkg>` will not break within an era" | ✅ | Verified — `enum-v3` imports use `core-v9` consistently (Cycle 1 + memory Core rule) |
| 13 | §4 line 108 | "Diagnostic message formats stable when consumed by tests in `tests/integratedtests/`" | ⚠️ | **Drift D-CVS-27** — repeats the C-CVS-01 / D-CVS-17 / D-CVS-26 pattern; actual root is `tests/creationtests/` |
| 14 | §4 line 112 | "Across eras: module path changes (`core-v9` → `core-v9`)" | ❌ | **Contradiction C-CVS-09b** — same mojibake as #11; must read `core-v8` → `core-v9` |
| 15 | §5 | `.release/` is OFF-LIMITS — never create / modify / delete | ✅ | Rule honoured — `.release/` does not exist in repo (vacuously satisfied; no violations possible). |
| 16 | §5 (citation) | "enforced via `.lovable/user-preferences` line 8 and `mem://index.md` Core" | ⚠️ | Same as #9 — `.lovable/` missing; only `mem://index.md` Core carries the rule |
| 17 | §6 (anti-pattern) | "Using `golang.org/x/mod/semver` directly" — Common Mistake | ✅ | Verified — `rg "golang.org/x/mod/semver"` → 0 hits in `enum-v3` |
| 18 | §6 (anti-pattern) | "Editing `.release/` to help" — Common Mistake | ✅ | Vacuously verified (folder absent) |
| 19 | §6 (anti-pattern) | "Relying on a specific patch number in tests" → use `versionindexes.V<N>` | ❓ | `versionindexes.V*` not consumed; no in-repo violation either |
| 20 | Source line 5 | "`coreversion` package + `.lovable/user-preferences` line 8" attribution | ⚠️ | Same `.lovable/` issue as #9 |

**Total claims**: 20
**Verifiable subset**: 9 (claims #7, #8, #9, #11, #12, #13, #14, #15, #16, #17, #18 — counted as 9 distinct rule/path/symbol checks; #16 and #20 are duplicates of #9)
**Verifiable match rate (baseline)**: **4 ✅ / 2 ❌ / 3 ⚠️ / 0 ❓ on verifiable subset = 4 / 9 = 44.4%**

---

## 3. Score row

| Date       | Cycle | Spec audited                  | Claims | ✅ | ⚠️ | ❌ | ❓ | Score (verifiable) |
|------------|-------|-------------------------------|--------|----|-----|----|----|--------------------|
| 2026-05-05 | 9 (baseline) | `01-app/11-versioning.md` | 20 | 4 | 3 | 2 | 11 | **44.4%** *(verifiable)* |
| 2026-05-05 | 9 (closed)   | `01-app/11-versioning.md` | 20 | 9 | 0 | 0 | 11 | **100.0%** *(verifiable)* |

---

## 4. Findings & fixes

### C-CVS-09a / C-CVS-09b — `core-v9 → core-v9` mojibake (§3 line 95, §4 line 112) — **HIGH (contradiction)**

Both lines describe historical era bumps (the v8→v9 rename) but render as nonsensical `core-v9 → core-v9`. Artifacts of the bulk `v8`→`v9` rewrite that the user invoked across all spec docs.

**Fix**: rewrite to `core-v8` → `core-v9` (the legitimate historical reference). These are exactly the two spots where mentioning `core-v8` outside `cross-repo/core-v8/` is correct, since they describe a past migration boundary.

### D-CVS-27 — `tests/integratedtests/` (§4 line 108) — **LOW (path-string)**

Same drift pattern as C-CVS-01 / D-CVS-17 / D-CVS-26. **Fix**: rewrite to `tests/creationtests/`.

### D-CVS-30 — `versionindexes.V8` comment mojibake (§2 line 59) — **LOW**

Comment says `V8 // 8 (current era — core-v9)`. **Fix**: rewrite to `V8 // 8 (legacy era; the current core-v9 era is V9)` so the comment accurately distinguishes the index from the era name.

### D-CVS-31 — Stale `.lovable/user-preferences` citation (§3 source line, §3 line 78, §5 line 133) — **LOW (citation hygiene)**

`.lovable/` does not exist in `enum-v3`. Both rules (`bump at least minor`, `never touch .release/`) live in `mem://index.md` Core. **Fix**: rewrite citations to point only to `mem://index.md` Core, drop the `.lovable/user-preferences line 8` reference.

---

## 5. Next actions

1. Apply the 4 fixes above to `spec/01-app/11-versioning.md` (done in this cycle).
2. Update scoreboard with Cycle 9 baseline + closed rows; add resolved findings C-CVS-09a/b, D-CVS-27, D-CVS-30, D-CVS-31.
3. Continue to Cycle 10 → `12-cmd-entrypoints.md` on next `next`.
