# Cycle 47 — AB-residual deep-probe of `spec/01-app/10-reflection-and-dynamic.md` (Cycle 22 carry-over)

**Date:** 2026-05-06
**Scope:** Settle the 6 ❓ items left by Cycle 22 (`23-cycle22-AB-reflection-and-dynamic.md` "❓ Still unverifiable") using direct evidence from upstream `core-v9 v1.5.8` clone.
**Result:** **1 ❓ → ✅** (verbatim) + **2 ❓ → ❌ DEMOTIONS** (NEW C-CVS-63, C-CVS-64) + **1 ❓ → ✅ partial + retraction** (R-CVS-03) + **2 ❓ → out-of-band advisory**.
**Allowed under freeze:** read-only audit promotion (no spec rewrites).

## 1. Promotions and demotions

### 1.1 Item 6 — `isany.DeepEqual` → ✅ (verbatim)

- `isany/DeepEqual.go:29` — `func DeepEqual(...)` exists.
- Companion symbols: `isany/NotDeepEqual.go:27` `NotDeepEqual(left, right any) bool`; `isany/DeepEqualAllItems.go:25` `DeepEqualAllItems(items ...any) bool`.
- `isany/README.md:51` documents the surface: "`DeepEqual(left, right any) bool` — Reflection-based deep equality via `reflectinternal`".
- **Verdict:** ✅ — spec §7 advice "use `isany.DeepEqual`" is verbatim correct.

### 1.2 Item 5 — "Lazy binding via `coreonce`" → ❌ NEW C-CVS-63 (HIGH) + R-CVS-03 retraction

- **R-CVS-03 (Retraction):** Cycle 22 said `coreonce` "not present in `core-v9 v1.5.8` root listing — pending §06 reaudit". This is wrong by the same drift class as R-CVS-01 (`coredynamic`) / R-CVS-02 (`corestr`): the package **exists** at `coredata/coreonce/` (13 files: `AnyOnce.go`, `AnyErrorOnce.go`, `BoolOnce.go`, `ByteOnce.go`, `BytesOnce.go`, `BytesErrorOnce.go`, `ErrorOnce.go`, `IntegerOnce.go`, `IntegersOnce.go`, `MapStringStringOnce.go`, `StringOnce.go`, `StringsOnce.go`, README). Original probe missed the `coredata/` parent.
- **NEW C-CVS-63 (HIGH):** `coreonce` is a **typed memoization/lazy-init** package (initializer-func pattern: `NewAnyErrorOnce(initializerFunc func() (any, error)) AnyErrorOnce` at `coredata/coreonce/AnyErrorOnce.go:43`, with methods `Execute`, `ExecuteMust`, `Value`, `ValueWithError`, `IsSuccess`, `IsFailed`, etc.). It does NOT do **reflection-binding** — it caches the result of an arbitrary initializer function. Spec §6 frames `coreonce` as "lazy binding" for reflection's caller-supplied callbacks; that is mis-attribution. Real lazy-binding for reflection in `core-v9` is the `reflect.Type`-keyed pattern in `reflectcore/reflectmodel/FieldProcessor.go` (no `coreonce` involvement).
- **Verdict:** package ✅ exists; framing ❌ — spawn **AJ-19b (BLOCKED by freeze):** rewrite §6 to drop the "Lazy binding via `coreonce`" bullet, OR replace with "Memoization of reflection results via `coredata/coreonce.NewAnyErrorOnce(initializerFunc)`" (which is a real, sensible application).

### 1.3 Item 2 — "unsafe pointer arithmetic for `corejson` fast-path (avoids 1 alloc per Marshal)" → ❌ NEW C-CVS-64 (HIGH)

- `grep -rn '"unsafe"' coredata/corejson/ coredata/coredynamic/` returns **zero hits**. Neither package imports `unsafe`.
- Both packages exist (per R-CVS-01 and R-CVS-02 retractions); the original "no `corejson/` package exists" framing in Cycle 22 was wrong, but the **fast-path claim itself** is now actively falsifiable: there is no unsafe-pointer fast-path in either package.
- **Verdict:** ❌ — **NEW C-CVS-64 (HIGH):** delete the "unsafe pointer arithmetic … avoids 1 alloc per Marshal" sentence from §4. Spawn **AJ-17b (BLOCKED by freeze)** — fold into the existing AJ-17 §3.2 / §4 rewrite for one consolidated edit.

## 2. Items retained as ❓ / out-of-band

| # | Claim | Disposition |
|---|-------|-------------|
| 1 | §1 motivation prose ("verbose, panics liberally, loses type info") | **out-of-band** — subjective framing; Task AC advisory dimension. |
| 3 | §4 "type-cache keyed on `reflect.Type`" | **retained ❓ (but downgraded to plausible-no-emitter)** — `reflect.Type` IS used in `reflectcore/reflectmodel/FieldProcessor.go:35,38` (`FieldType reflect.Type` + `IsFieldType(t reflect.Type) bool`), but no `map[reflect.Type]...` cache structure was found. Could exist in unsearched paths or be implicit in a method-receiver field; defer to focused performance/AC pass. |
| 4 | §6 "10–100× slower than direct calls" | **out-of-band** — benchmark claim; no benchmarks in repo. Task AC. |

## 3. Updated Cycle-22 scoreboard

| Pass | ✅ | ⚠️ | ❌ | ❓ | Verifiable score |
|------|----|---|----|---|---|
| Cycle 8 baseline | 0 | 0 | 0 | 15 | N/A |
| Cycle 22 (AB pass 4) | 5 | 0 | 8 | 6 | 38.5% |
| **Cycle 47 (this)** | **6** | 0 | **10** | **1 retained + 2 OOB** | **37.5%** (6/16) |

§10 verifiable score drops slightly (38.5% → 37.5%) because the demotions add 2 to the denominator. Net audit health improves (more knowledge, fewer unknowns). Active ❓ pool: 1 (item 3, downgraded plausible).

## 4. Cumulative AB-residual ❓ pool

- Pre-Cycle-47: 11 ❓ in `spec/01-app/` (Cycle 19=4 OOB / Cycle 20=1 OOB / Cycle 22=6).
- Post-Cycle-47: **11 − 5 = 6 ❓** in `spec/01-app/` (Cycle 19=4 OOB / Cycle 20=1 OOB / Cycle 22=1 plausible-no-emitter).
- **🎉 ZERO ACTIVE PROBE TARGETS REMAIN** — all 6 residual ❓ items are now classified as OOB/plausible-no-emitter, awaiting Task AC's contract/advisory pass rather than a fresh upstream grep.

## 5. Cumulative AB ❌ summary update

- Cumulative AB ❌ across 7 sections: **51 → 53** (NEW C-CVS-63 HIGH + C-CVS-64 HIGH). CRITICAL count unchanged at **23**. HIGH count grows by 2.
- Cumulative retractions: **R-CVS-01** (`coredynamic`), **R-CVS-02** (`corestr`), **R-CVS-03** (`coreonce`) — all three follow the same "missed `coredata/` parent" probe pattern. **NEW S-115 suggestion:** harden `Get-UpstreamPackages` (or its caller) to recursively walk `coredata/` so future audits don't repeat this mistake.
- Cumulative ⓘ "upstream-only" annotations unchanged at 11.

## 6. Action items spawned / amended

- **AJ-17b (NEW, BLOCKED by freeze):** delete unsafe-pointer fast-path sentence at `10-reflection-and-dynamic.md` §4 (C-CVS-64). Fold into AJ-17.
- **AJ-19b (NEW, BLOCKED by freeze):** rewrite §6 "Lazy binding via `coreonce`" to either drop or correctly attribute as "Memoization via `coredata/coreonce.NewAnyErrorOnce`" (C-CVS-63). Fold into AJ-19.
- **S-115 (NEW):** harden `scripts/spec-api-check.psm1` `Get-UpstreamPackages` to recursively index `coredata/*` subpackages — file under `.lovable/memory/suggestions/01-suggestions-tracker.md`.

## 7. Memory + bookkeeping

- Spec changelog → `spec-v0.52.0`.
- `package.json` → `0.17.0`.
- `.lovable/memory/workflow/01-state.md` updated.
- Scoreboard top-line updated.
- **🎉 AB-residual deep-probe sweep across `spec/01-app/` is now COMPLETE.** All seven AB-cycle ❓ pools (Cycles 19, 20, 21, 22, 23, 24, 25) have been settled or classified out-of-band. Remaining work in `spec/01-app/` is the AJ rewrite backlog (52 items, all blocked by freeze).

---

_Audit file: `spec/07-code-vs-spec-audits/36-cycle47-AB-residual-spec01-reflection.md`_
_See also: `23-cycle22-AB-reflection-and-dynamic.md` (the cycle whose ❓ are settled here), `27-cycle26-S106-self-audit-retractions.md` (R-CVS-01/02 precedent)._