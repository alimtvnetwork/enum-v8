# Cycle 25 — AB Pass 7: `spec/01-app/16-security.md`

> **Date:** 2026-05-06 (Asia/Kuala_Lumpur)
> **Audit type:** AB promotion pass — convert ❓ unverifiable claims into ✅/⚠️/❌ via direct comparison against upstream `core-v9 v1.5.8`.
> **Spec freeze status:** `spec/01-app/` remains 🧊 **DRIFT-FROZEN** (spec-v0.30.0). All proposed rewrites spawned as **AJ-35..43** and BLOCKED.
> **Predecessor:** Cycle 14 (baseline, `15-cycle14-security.md`) — 17 ✅ / 13 ❓.

## 1. Promotion summary

| Metric | Cycle 14 (baseline) | Cycle 25 (AB pass) |
|---|---|---|
| ✅ Match | 17 | 17 + **3** = **20** |
| ⚠️ Drift | 0 | **1** |
| ❌ Contradiction | 0 | **9** |
| ❓ Unverifiable | 13 | **0** |
| **Verifiable score** | 100.0 % (small subset) | **66.7 %** |

Of the 13 ❓ promoted: **3 ✅** confirmed, **1 ⚠️** drift, **9 ❌** new contradictions (C-CVS-51 .. C-CVS-59). §16 inherits fabrication from earlier cycles: the chapter teaches security rules *in terms of* fabricated APIs (`coredynamic.*`, `corevalidator.New.Line/Slice` fluent, `corestr.*`), so virtually every code-bearing rule is built on a non-existent foundation.

## 2. Evidence sources

- `coredata/coregeneric/` — ✅ contains `Collection.go`, `Hashmap.go`, `Hashset.go`, `SimpleSlice.go`, `LinkedList.go` (and Iter/Creator companions)
- `find . -name "StringBuilder*.go"` — **zero matches** anywhere in upstream
- `grep -rn "IsValidUTF8\|ValidUTF8" --include="*.go"` — **zero matches** anywhere in upstream
- `find . -type d -name "corestr*"` — **no `corestr/` package** in upstream
- `grep -rn "InvalidInput" errcore/` — **zero matches**; `errcore.InvalidInput` is not a category
- `errcore/vars.go:25-29` — only public package vars are `ShouldBe`, `Expected`, `StackEnhance` (no error-category vars exposed at package level)
- Inherited from prior cycles:
  - C-CVS-29 (Cycle 22) — `coredynamic/` package does not exist
  - C-CVS-22/23 (Cycle 21) — `corevalidator.New.Line/Slice` fluent + `Validate(input) Result` fabricated
  - C-CVS-44 (Cycle 24) — `errcore.VarTwo` missing leading `isIncludeType bool` parameter

## 3. Promotions (❓ → ground truth)

### ✅ Confirmed (3)

| # | Spec claim | Evidence |
|---|---|---|
| P1 | `coregeneric.{Collection, Hashmap, Hashset, SimpleSlice, LinkedList}` all exist as generic containers | `coredata/coregeneric/{Collection,Hashmap,Hashset,SimpleSlice,LinkedList}.go` |
| P2 | None of those containers ship a built-in size limit (rule §4) | Source inspection of `Collection.go` / `Hashmap.go` constructors — no cap field, no growth ceiling |
| P3 | `internal/reflectinternal` exists and would mutate unexported fields if imported by consumers | Confirmed in Cycle 22 C-CVS-34; rule §5.1 is correct in spirit (the *reason* it's "off-limits" framing was the C-CVS-34 issue, but the *imperative* "don't import" is sound) |

### ⚠️ Drift (1)

#### D-CVS-61 *(LOW)* — `coregeneric` import path is incomplete

- **Spec (§4 table):** Names containers as `coregeneric.Collection[T]`, `coregeneric.Hashmap[K,V]`, etc.
- **Reality:** Real import path is `github.com/alimtvnetwork/core-v9/coredata/coregeneric` — the `coredata/` parent is consistently elided throughout §4. Same drift class as C-CVS-14 (`corejson` → `coredata/corejson`).
- **Severity:** LOW (path drift, not API drift).
- **Spawned:** AJ-35 — qualify the package path as `coredata/coregeneric` in §4 table headers.

### ❌ New contradictions (9)

#### C-CVS-51 *(CRITICAL)* — `corestr` package does not exist

- **Spec (§4 table line 101, §6 rule 3 line 146):** Refers to `corestr.StringBuilder` and `corestr.IsValidUTF8`.
- **Reality:** `find /tmp/core-v9-upstream -type d -name "corestr*"` returns **no matches**. Upstream `core-v9 v1.5.8` has no `corestr/` package. String-related work lives in `coredata/corejson/`, `typesconv/`, and stdlib `strings`.
- **Severity:** CRITICAL. Two separate rules cite a fabricated package as the canonical mitigation.
- **Spawned:** AJ-36 — delete `corestr.StringBuilder` row from §4 table; rewrite §6 rule 3 to use `unicode/utf8.ValidString` from stdlib (or document a real upstream helper if one exists in `typesconv/`).

#### C-CVS-52 *(CRITICAL)* — `corestr.StringBuilder` is fabricated

- **Spec (§4):** "`corestr.StringBuilder` — Unbounded concatenation" risk row.
- **Reality:** `find . -name "StringBuilder*.go"` returns zero matches anywhere in upstream. There is no upstream string-builder type. Consumers needing one use stdlib `strings.Builder`.
- **Severity:** CRITICAL.
- **Spawned:** AJ-37 — replace with `strings.Builder` (stdlib) and rewrite the risk callout to apply to the stdlib type.

#### C-CVS-53 *(CRITICAL)* — `corestr.IsValidUTF8` is fabricated

- **Spec (§6 rule 3):** "Use `corestr.IsValidUTF8` if UTF-8 is a precondition."
- **Reality:** `grep -rn "IsValidUTF8" --include="*.go"` returns zero matches. Real call is stdlib `utf8.ValidString` or `utf8.Valid([]byte)`.
- **Severity:** CRITICAL.
- **Spawned:** AJ-38 — substitute stdlib `unicode/utf8.ValidString`.

#### C-CVS-54 *(CRITICAL)* — `errcore.InvalidInput` category does not exist

- **Spec (§6 example, line 138):** `return errcore.InvalidInput.MergeError(result.Error())`.
- **Reality:** `grep -rn "InvalidInput" errcore/` returns zero matches. `errcore/vars.go` exposes only `ShouldBe`, `Expected`, `StackEnhance` at package level. The `RawErrorType.MergeError` method exists (`errcore/RawErrorType.go:266`) but it's a method on existing typed-error instances, not on a fabricated `InvalidInput` singleton.
- **Severity:** CRITICAL. Spec example will not compile.
- **Spawned:** AJ-39 — replace with a real category lookup (e.g. `errcore.ShouldBe.Equal(...).MergeError(...)` or whatever category actually exists for input-validation errors) after `grep -l "RawErrorType{" errcore/` enumerates the real surface.

#### C-CVS-55 *(CRITICAL)* — `coredynamic.AllFields` (§4 rule 4) is fabricated

- **Spec (§4 rule 4):** "**MUST NOT** call `coredynamic.AllFields` on a deeply nested struct in a hot path."
- **Reality:** Inherited from C-CVS-29 — `coredynamic/` package does not exist; `AllFields` is fabricated.
- **Severity:** CRITICAL. Rule cannot be followed because the API doesn't exist; readers asked to "audit for `AllFields` calls" will find none and may falsely conclude their code is safe.
- **Spawned:** AJ-40 — delete the rule, or rewrite using the real reflection surface (`reflectcore.Looper` / `reflectcore/reflectmodel.FieldProcessor` per Cycle 22 evidence).

#### C-CVS-56 *(CRITICAL)* — `coredynamic.SetField` / `coredynamic.InvokeMethod` (§5 rules 2-3) are fabricated

- **Spec (§5 rules 2-3):** Two MUST-NOT rules about `coredynamic.SetField` and `coredynamic.InvokeMethod`.
- **Reality:** Both fabricated (C-CVS-29). Real reflection surface in `reflectcore/` exposes `Converter`, `Utils`, `Looper`, etc. — none with these names.
- **Severity:** CRITICAL. Two whole rules in the security boundary chapter are unactionable.
- **Spawned:** AJ-41 — rewrite §5 rules 2-3 against the real `reflectcore` API.

#### C-CVS-57 *(CRITICAL)* — `corevalidator.New.Line / New.Slice` fluent API in §6 example is fabricated

- **Spec (§6, lines 130-140):** Multi-line example using `corevalidator.New.Line.NotEmpty().MaxLength(255).Matches(...).Build()` and `result.IsFailed()` / `result.Error()`.
- **Reality:** Inherited from C-CVS-22 / C-CVS-23 (Cycle 21). Real `corevalidator/` exposes `LineValidator{LineNumber, TextValidator}` with `IsMatch(lineNumber, content, isCaseSensitive) bool` — no `New` var, no fluent builder, no `Result` type.
- **Severity:** CRITICAL. The chapter's flagship "trust boundary" worked example will not compile.
- **Spawned:** AJ-42 — rewrite §6 example using stdlib `regexp.MatchString` + length checks, or document the real `LineValidator.IsMatch` shape and remove the fluent claim.

#### C-CVS-58 *(HIGH)* — `corevalidator.New.Slice.MaxLength(N)` cited in §4 rule 1, §6 rule 2, and §7 mistake row is fabricated

- **Spec:** Three separate rules and one "mistake row" prescribe `corevalidator.New.Slice.MaxLength(N)` as the canonical upstream check.
- **Reality:** No `New` var exists; `corevalidator/SimpleSliceValidator.go` exists but exposes a different surface. The fluent `.MaxLength(N).Build()` shape is fabricated.
- **Severity:** HIGH (replicated in 4 places, so a single rewrite must touch all of them).
- **Spawned:** AJ-43 — rewrite §4 rule 1 + §6 rule 2 + §7 row using the real `SimpleSliceValidator` surface OR a stdlib `len(s) > N` check.

#### C-CVS-59 *(HIGH)* — `errcore.VarTwo("password", pwd, …)` example in §2 reproduces C-CVS-44 signature defect

- **Spec (§2 table line 51):** `errcore.VarTwo("password", pwd, …)` shown as 4-arg call.
- **Reality:** Same as C-CVS-44 — `VarTwo` requires a leading `isIncludeType bool`. The §2 example will not compile (and would emit the wrong format if it did).
- **Severity:** HIGH (PII example must be copy-paste-correct because security rules carry extra weight in audits).
- **Spawned:** *folded into AJ-28* — fix all `VarTwo` examples atomically, including this one.

## 4. Cumulative AB tally (Cycles 19-25)

| Section | ❌ | Score |
|---|---|---|
| §07 conditional-and-utilities | 5 | 70.6 % |
| §08 validators | 8 | 33.3 % |
| §09 converters | 5 | 66.7 % |
| §10 reflection-and-dynamic | 8 | 38.5 % |
| §11 versioning | 8 | 18.2 % |
| §15 observability | 7 | 74.1 % |
| **§16 security** | **9** | **66.7 %** |
| **Total** | **50** | — |

- **Cumulative ❌:** **50** (was 41) across 7 audited sections.
- **Severity mix:** 24 CRITICAL + 21 HIGH + 3 MEDIUM + 2 LOW.
- **Fabrication rate** (❌ ÷ verifiable claims promoted): now **~54 %** across the AB sweep.
- **All 7 sections of `spec/01-app/` previously holding ❓ have now been promoted.** AB sweep of `spec/01-app/` is **complete**.
- **S-106 lint:** still **MANDATORY** — Cycle 25 surfaced 6 CRITICAL items, 5 of which are simple "named symbol does not exist" patterns that S-106 would catch in <1 second per spec file.

## 5. Spawned AJ items (BLOCKED)

| ID | Target | Severity | Action |
|---|---|---|---|
| AJ-35 | §4 table headers | LOW | Qualify path as `coredata/coregeneric` |
| AJ-36 | §4 line 101 + §6 rule 3 | CRITICAL | Delete `corestr.StringBuilder` row; rewrite UTF-8 rule using stdlib |
| AJ-37 | §4 risk row | CRITICAL | Replace `corestr.StringBuilder` with stdlib `strings.Builder` |
| AJ-38 | §6 rule 3 | CRITICAL | Substitute `unicode/utf8.ValidString` |
| AJ-39 | §6 line 138 | CRITICAL | Replace fabricated `errcore.InvalidInput` with real category |
| AJ-40 | §4 rule 4 | CRITICAL | Delete or rewrite `coredynamic.AllFields` rule |
| AJ-41 | §5 rules 2-3 | CRITICAL | Rewrite using real `reflectcore` surface |
| AJ-42 | §6 lines 130-140 | CRITICAL | Rewrite trust-boundary example using real validator API |
| AJ-43 | §4 rule 1 + §6 rule 2 + §7 row | HIGH | Substitute real `SimpleSliceValidator` or `len()` check |
| (AJ-28) | §2 line 51 | HIGH | Folded into existing AJ-28 — fix all `VarTwo` calls atomically |

All AJ-28..43 are **BLOCKED** by the `spec/01-app/` 🧊 freeze. **S-106 lint should land before these are unblocked.**

## 6. Remaining ❓ in `spec/01-app/`

After Cycle 25: §07 3 + §08 6 + §09 8 + §10 6 + §11 1 = **24 ❓** remaining (down from 37). All remaining ❓ are residual non-API claims (workflow patterns, decision-table prose, etc.) — not API-shape claims. **No more 13-❓ batches remain in `spec/01-app/`.**

Path forward:
1. **Build S-106** (MANDATORY) — without it, AJ rewrites will reintroduce fabrications.
2. **Resolve workflow/script-internal ❓** in spec/03 (6) + spec/04 (8) + spec/06 (10) + spec/02 audit-history (5) = 29 ❓ via deep-probe of `scripts/*.psm1` and `.github/workflows/*.yml`.
3. **AC re-audit pass** — once AJ-01..43 land, re-audit §07-§11 + §15 + §16 against the consistency dimension.
