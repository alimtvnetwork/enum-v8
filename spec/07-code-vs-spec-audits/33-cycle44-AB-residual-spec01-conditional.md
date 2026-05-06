# Cycle 44 — AB-residual deep-probe of `spec/01-app/07-conditional-and-utilities.md` (Cycle 20 carry-over)

**Date:** 2026-05-06
**Scope:** Settle the 3 ❓ items left by Cycle 20 (`21-cycle20-AB-conditional-and-utilities.md` §1.3) using direct evidence from upstream `core-v9 v1.5.8` clone at `/tmp/core-v9-upstream`.
**Result:** **2 ❓ → ✅** (1 with mechanism-name drift) + **1 ❓ retained as advisory/out-of-band**.
**Allowed under freeze:** read-only audit promotion (no spec rewrites).

## 1. Promotions

### 1.1 Row 51 — "`LazyLock` defers cost to first use, then caches" → ✅ (with NEW D-CVS-66 LOW)

- `regexnew/LazyRegex.go:34-50` defines `LazyRegex struct { mu sync.Mutex; isCompiled bool; isApplicable bool; pattern string; regex *regexp.Regexp; compiledErr error; compiler func(...) ... }`.
- `Compile()` body (lines 90-110) — locks `mu`, returns cached `(regex, compiledErr)` if `isCompiled`, otherwise calls `compiler`, stores result + sets `isCompiled = true`. Subsequent calls hit the cache branch under the lock.
- `IsApplicable()` (lines 61-88) — fast-path check under lock; if `isApplicable` already true, returns immediately; otherwise triggers the lazy compile.
- `regexnew/vars.go:33-34` — package-level `regexMutex = sync.Mutex{}` + `lazyRegexLock = sync.Mutex{}` (the "Lock" suffix referenced by spec's `LazyLock` naming).
- **Behavioural claim ✅:** "defers cost to first use, then caches" — confirmed verbatim (cost = `compiler(pattern)`; cache = `it.regex` + `it.compiledErr` + `isCompiled` guard).
- **NEW D-CVS-66 (LOW) — mechanism-name drift:** spec line 173 implies `sync.Once`; actual implementation uses `sync.Mutex` + boolean guard (`isCompiled`). Functionally equivalent but mechanically distinct (`sync.Once.Do` cannot return values whereas this design needs `(regex, err)` return + cached-error semantics). Recommend spec footnote: "implemented via `sync.Mutex` + `isCompiled` flag rather than `sync.Once` because the cached call must return `(*regexp.Regexp, error)`." → folded into existing AJ-04 as **AJ-04b**.

### 1.2 Row 52 — `corecmp` returns `constants.CompareEqual / Less / Greater` (`0 / -1 / 1`) → ✅ (verbatim)

- `constants/constants.go:336-338`:
  ```go
  CompareEqual                             = 0
  CompareLess                              = -1
  CompareGreater                           = 1
  ```
- Names match spec exactly. Values match spec exactly.
- `corecmp/` package exists (10+ files: `AnyItem.go`, `Byte.go`, `Integer.go`, `Integer16.go`, …`Integer64.go`, `*Ptr.go` variants) — confirms `corecmp` is the real package (not a fabricated namespace).
- `corecomparator/Compare.go:250` — `Compare` value type with `IsCompareEqualLogically(...)` method and friends; the comparator surface uses these constants downstream.
- **Verdict:** ✅ — spec line 200 is verbatim correct. No drift, no new findings.

## 2. Retained ❓ (out-of-band / advisory)

| Spec line | Claim | Why retained |
|-----------|-------|--------------|
| 142 | `issetter.Value` "Pitfall: not a drop-in for `bool`" | **Advisory / behavioural** — not directly verifiable from code. Spec is teaching a developer-experience pitfall, not asserting a code property. **Out-of-band** — Task AC dimension. |

## 3. Updated Cycle-20 scoreboard

| Pass | ✅ | ⚠️ | ❌ | ❓ | Verifiable score |
|------|----|---|----|---|---|
| Cycle 5 baseline | 0 | 0 | 0 | 17 | N/A |
| Cycle 20 (AB pass 1) | 12 | 0 | 5 | 3 | 70.6% |
| **Cycle 44 (this)** | **14** | 0 | 5 | **1** (advisory only) | **73.7%** (14/19) |

§07 verifiable subset grows by 2 (rows 51, 52). ❌ count unchanged. The remaining 1 ❓ is explicitly classified as advisory — not unknown.

## 4. Cumulative AB-residual ❓ pool

- Pre-Cycle-44: 20 ❓ in `spec/01-app/` (Cycle 19=4 OOB / Cycle 20=3 / Cycle 21=6 / Cycle 22=6 / Cycle 23=1).
- Post-Cycle-44: **20 − 2 = 18 ❓** in `spec/01-app/` (Cycle 19=4 OOB / Cycle 20=1 OOB / Cycle 21=6 / Cycle 22=6 / Cycle 23=1).

## 5. Action items spawned / amended

- **AJ-04b (NEW, BLOCKED by freeze):** add footnote at `07-conditional-and-utilities.md:173` clarifying `sync.Mutex` + `isCompiled` guard mechanism vs `sync.Once` (D-CVS-66 LOW).

## 6. Memory + bookkeeping

- Spec changelog → `spec-v0.49.0`.
- `package.json` → `0.14.0`.
- `.lovable/memory/workflow/01-state.md` updated.
- Scoreboard top-line updated.

---

_Audit file: `spec/07-code-vs-spec-audits/33-cycle44-AB-residual-spec01-conditional.md`_
_See also: `21-cycle20-AB-conditional-and-utilities.md` (the cycle whose ❓ are settled here)._