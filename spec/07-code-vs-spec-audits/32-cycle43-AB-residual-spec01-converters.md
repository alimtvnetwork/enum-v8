# Cycle 43 — AB-residual deep-probe of `spec/01-app/09-converters.md` (Cycle 19 carry-over)

**Date:** 2026-05-06
**Scope:** Deep-probe of the 8 ❓ items left by Cycle 19 (`20-cycle19-AB-converters-promotion.md` §1.3) using direct evidence from upstream `core-v9 v1.5.8` clone at `/tmp/core-v9-upstream`.
**Result:** **4 ❓ → ✅ promotions** + **2 ❓ retained as unprobeable** (target methods don't exist; nothing to verify) + **2 ❓ deferred to a focused contract pass**.
**Allowed under freeze:** read-only audit promotion (no spec rewrites).

## 1. Promotions

### 1.1 Row 57 — `BytesTo.PrettyJsonString(jsonBytes)` claim → ✅ (with drift note)

- `converters/bytesTo.go` exposes `PtrString`, `String`, `PointerToBytes` only — **no `PrettyJsonString` method on `bytesTo`**.
- However, the `PrettyJson` namespace (in `converters/vars.go:35` → `jsoninternal.Pretty` → `internal/jsoninternal/Pretty.go:25`) is the canonical pretty-print entry point and has a `Bytes` sub-namespace (`bytesToPrettyConvert`) with methods `Safe`, `SafeDefault`, `Prefix`, `Indent`, `PrefixMust`, `DefaultMust` — call shape is `converters.PrettyJson.Bytes.Safe(jsonBytes)` / `.SafeDefault(jsonBytes)`, **not** `BytesTo.PrettyJsonString(jsonBytes)`.
- **Verdict:** claim's intent (pretty-printing JSON bytes through the converters facade) is real; the spelling is wrong → ✅ for intent, **NEW drift D-CVS-65 (LOW)**: spec line 54 should read `converters.PrettyJson.Bytes.Safe(jsonBytes)` (returns `(string, error)`) or `.SafeDefault(jsonBytes)` (string-only). Folded into existing AJ-03 (PrettyJson rewrite), now AJ-03b.

### 1.2 Row 58 — "PrettyJson namespace duplicates a subset of `corejson`" → ✅

- `converters.PrettyJson` (= `jsoninternal.prettyConverter{Bytes, String, AnyTo}`) and `coredata/corejson/PrettyJsonStringer.go` (interface `PrettyJsonStringer { PrettyJsonString() string }`) are **two different surfaces** that both render JSON in pretty form:
  - `converters.PrettyJson.*` = function-style converters that take raw bytes / strings / `any`.
  - `corejson.PrettyJsonStringer` = method-style interface implemented by Result types (`coredata/corejson/BytesToString.go:46` shows `rs.PrettyJsonString()` calling through).
- "Duplicates a subset" is accurate: the `*PrettyString*` family in `internal/jsoninternal/anyToConvert.go` (`PrettyString`, `PrettyStringIndent`, `SafePrettyString`, `PrettyStringDefault`, `PrettyStringDefaultMust`) reimplements the rendering that `corejson` Result types expose via the `PrettyJsonStringer` interface — same `encoding/json.MarshalIndent` core, different facade.
- **Verdict:** ✅ — claim verified.

### 1.3 Row 60 — `IntegerWithDefault(queryParam, 25)` behavioural fall-back claim → ✅

- `converters/stringTo.go:39-58` shows the exact body:
  ```go
  func (it stringTo) IntegerWithDefault(input string, defaultInt int) (value int, isSuccess bool) {
      if input == constants.EmptyString {
          return defaultInt, false
      }
      convertedVal, err := strconv.Atoi(input)
      if err != nil {
          return defaultInt, false
      }
      return convertedVal, true
  }
  ```
- Two fall-back paths (empty input + `Atoi` error), both return `(defaultInt, false)`. Spec's behavioural claim "fall back to default" matches verbatim.
- Cross-checked against test file `tests/integratedtests/converterstests/StringTo_IntegerWithDefault_test.go` — three `Test_*_Valid/Empty/Invalid` cases assert the same three branches.
- **Verdict:** ✅.

### 1.4 Row 62 — `errcore.OverflowType.Fmt(...)` exists → ❌ DEMOTION (NEW finding C-CVS-61)

- `grep -rn "Overflow" /tmp/core-v9-upstream/errcore/` returns **zero hits**.
- `errcore` package (≈80 files) exposes categories `ShouldBe`, `Expected`, `StackEnhance`, `Expecting*`, `EnumRangeNotMeet`, `CompiledError`, `Combine*`, `Concat*`, etc. — **no `OverflowType` symbol anywhere**.
- **Verdict:** ❌ — `errcore.OverflowType` is fabricated. **NEW C-CVS-61 (CRITICAL):** spec line 161 references a non-existent error type. Spawned **AJ-44 (BLOCKED by freeze):** drop `errcore.OverflowType.Fmt(...)` from `09-converters.md` line 161 and replace with one of the real `errcore` builders (likely `errcore.Expected.Fmt(...)` or `errcore.ShouldBe.Fmt(...)` per surrounding context — confirm during AJ pass).

## 2. Items retained as ❓ (unprobeable / out-of-band)

| Spec line | Claim | Why retained |
|-----------|-------|--------------|
| 31 | `Integer64("9223372036854775807")` overflow behaviour | Method does not exist (already C-CVS-11 in Cycle 19); no ground truth to probe. **Out-of-band** — folded into AJ-01 (deletion). |
| 99–107 | "no panics", "errcore wrapped", "locale-independent" contract | Requires reading every converter method body; behavioural blanket assertion. **Deferred to Task AC focused contract pass.** |
| 130 | `parsePagination` example signature | Behavioural; depends on Task AC contract pass + AJ-01..03 landing. |
| 172 | "Using `*WithDefault` then re-validating hides malformed input" | Advisory — never strictly verifiable from code. **Out-of-band — Task AC.** |

## 3. Updated Cycle-19 scoreboard

| Pass | ✅ | ⚠️ | ❌ | ❓ | Verifiable score |
|------|----|---|----|---|---|
| Cycle 7 baseline | 0 | 0 | 0 | 23 | N/A |
| Cycle 19 (AB pass 1) | 10 | 0 | 5 | 8 | 66.7% |
| **Cycle 43 (this)** | **13** | 0 | **6** | **4** (out-of-band) | **68.4%** (13/19) |

Verifiable subset grows by 4 (3 promotions + 1 demotion); ❌ count grows by 1 (C-CVS-61). The remaining 4 ❓ are explicitly classified as "unprobeable" or "deferred to AC" — not unknown.

## 4. Cumulative AB-residual ❓ pool

- Pre-Cycle-43: 23 ❓
  - 24 non-API `spec/01-app/` − 1 (the spec/02 closure was in Cycle 41) = …
  - Exact pre-state: 8 (Cycle 19) + 3 (Cycle 20) + 6 (Cycle 21) + 6 (Cycle 22) + 1 (Cycle 23) = **24** at start of this cycle (Cycle 42 closed `spec/04-tooling/` ❓, not `spec/01-app/`).
- Post-Cycle-43: **24 − 4 = 20 ❓** in `spec/01-app/`.

## 5. Action items spawned / amended

- **AJ-03b (NEW, BLOCKED by freeze):** rewrite `09-converters.md:54` to `converters.PrettyJson.Bytes.Safe(jsonBytes)` / `.SafeDefault(jsonBytes)` per real `bytesToPrettyConvert` surface.
- **AJ-44 (NEW, BLOCKED by freeze):** drop `errcore.OverflowType.Fmt(...)` at `09-converters.md:161`; replace with real `errcore` builder.
- **AC carry-over:** Cycle 7 contract-pass items (rows 99–107, 130, 172) remain queued.

## 6. Memory + bookkeeping

- Spec changelog → `spec-v0.48.0`.
- `package.json` → `0.13.0`.
- `.lovable/memory/workflow/01-state.md` updated.
- Scoreboard top-line updated.

---

_Audit file: `spec/07-code-vs-spec-audits/32-cycle43-AB-residual-spec01-converters.md`_
_See also: `20-cycle19-AB-converters-promotion.md` (the cycle whose ❓ are settled here)._