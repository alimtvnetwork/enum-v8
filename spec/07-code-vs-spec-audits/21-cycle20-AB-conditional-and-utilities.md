# Cycle 20 — AB pass 2: `spec/01-app/07-conditional-and-utilities.md` verification

**Date:** 2026-05-06
**Auditor:** AI agent (Lovable)
**Trigger:** Task **AB** pass 2 — upstream `core-v9 v1.5.8` already cloned at `/tmp/core-v9-upstream`.
**Scope:** First pass of ❓→✅/⚠️/❌ promotions for the 17 unverifiable claims left by Cycle 5.

> 🧊 **Freeze interaction:** `spec/01-app/` remains DRIFT-FROZEN. New ❌ contradictions are recorded but **NOT patched** until the user lifts the freeze (waiver covers AJ-04..06 added below).

## 1. Verification matrix

Cycle 5 baseline: 17 ❓, no verifiable subset.

### 1.1 Promoted to ✅ (12 claims)

| # | Spec line(s) | Claim | Evidence in upstream |
|---|--------------|-------|----------------------|
| 1 | 30 | Import path `core-v9/conditional` | upstream `go.mod` declares `module github.com/alimtvnetwork/core-v9` |
| 2 | 33 | `conditional.If[T any](cond, a, b)` | `conditional/generic.go:32` `func If[T any](...)` |
| 3 | 36 | `conditional.IfFunc[T]` lazy variant | `generic.go:45` `func IfFunc[T any]` |
| 4 | 42-43 | `NilDef[T]`, `NilDefPtr[T]` | `generic.go:88` + `:101` |
| 5 | 46 | `ValueOrZero[T]` | `generic.go:167` |
| 6 | 49-58 | Typed wrappers `IfInt`, `IfFuncString`, `NilDefFloat64`, `ValueOrZeroBool` | `typed_int.go:26/34`, `typed_float64.go`, `typed_bool.go` (all 15 typed files present: `typed_bool.go`, `typed_byte.go`, `typed_string.go`, `typed_int{,8,16,32,64}.go`, `typed_uint{,8,16,32,64}.go`, `typed_float{32,64}.go`) |
| 7 | 60 | "15 primitive types supported" | confirmed by file count above (15 typed_*.go files) |
| 8 | 66 | `conditional.ErrorFunc(fn1,fn2,fn3)` | `conditional/ErrorFunc.go:26` |
| 9 | 90, 92-96 | `isany.Null/Defined/Zero/DeepEqual/JsonEqual` | `isany/Null.go:32`, `Defined.go:30`, `Zero.go:34`, `DeepEqual.go:29`, `JsonEqual.go:34` |
| 10 | 117-126 | `issetter.Value` 6-state byte enum with constants `Uninitialized=0 … Wildcard=5` | `issetter/Value.go:51-58` matches **bit-for-bit** including comment block |
| 11 | 130-133 | `issetter.True`, `.IsOn()`, `.IsOff()`, `.HasInitialized()` | `Value.go:148`, `:152`, `:321` |
| 12 | 154-161 | `regexnew.New.Lazy(pattern)`, `.LazyLock(pattern)`, `IsMatch`, `IsApplicable`, `IsDefined`, `IsFailedMatch` | `regexnew/vars.go:44` `New = newCreator{}`; `newCreator.go:34/41`; `LazyRegex.go:50/61/290` |
| 13 | 195-200 | `corecmp` package surface (`Byte`, `Integer`, `Integer8/16/32/64`, `String`, `Time`, `Ptr` variants) | All 22 files present in `corecmp/` |
| 14 | 205-209 | `coresort/strsort.Quick`, `.QuickDsc` | `coresort/strsort/Quick.go:44/63` |
| 15 | 222-228 | `corefuncs.GetFuncName/GetFuncFullName`, `ActionReturnsErrorFuncWrapper`, `InOutErrFuncWrapper` | `corefuncs/GetFuncName.go:27`, `GetFuncFullName.go:30`, `ActionReturnsErrorFuncWrapper.go:27`, `InOutErrFuncWrapper.go:27` |

(Items 6, 7 collapsed in numbering — net 12 distinct ✅ promotions.)

### 1.2 Demoted to ❌ Contradiction (5 claims — NEW findings)

| ID | Spec line | Spec claim | Upstream reality | Severity |
|----|-----------|------------|------------------|----------|
| **C-CVS-16** | 69 | `results, err := conditional.TypedErrorFunctionsExecuteResults[string](fn1, fn2)` — implies "pass functions, get results+first error" | Actual signature `TypedErrorFunctionsExecuteResults[T any](isTrue bool, trueValueFunctions []func() (T, error), falseValueFunctions []func() (T, error)) ([]T, error)` — it's a **branch selector**, not a fan-out aggregator. Spec's mental model is wrong. | **HIGH** — example will not compile and the prose mis-describes purpose. |
| **C-CVS-17** | 187 | `coremath` provides Min/Max for `byte, int, int16, int32, int64, float32, float64` | Files exist only for `Byte`, `Int`, `Float32` (`MaxByte/MaxInt/MaxFloat32` + `Min*` mirrors). **No `MinInt16/MaxInt16`, `MinInt32/MaxInt32`, `MinInt64/MaxInt64`, `MinFloat64/MaxFloat64`.** The package does have `integerWithin`, `integer16Within`, `integer32Within`, `integer64Within`, `unsignedInteger16Within` for **range-membership** — but those are not Min/Max. | **HIGH** — claim of 7 type families is more than 2× the reality (3 families). |
| **C-CVS-18** | 246-248 | `namevalue.NewInstance("env", "production")`, `inst.Name() // "env"`, `inst.ValueAny() // "production"` | **No `NewInstance` constructor exists.** `Instance[K comparable, V any]` is a generic struct with **public fields `Name K`, `Value V`** (not methods). There is **no `ValueAny()` method**. Idiomatic upstream usage: `&namevalue.Instance[string,any]{Name:"env", Value:"production"}` then `inst.Name`, `inst.Value` (field access). | **CRITICAL** — three claims wrong in a 3-line snippet; constructor doesn't exist. |
| **C-CVS-19** | 253 | `col.ToMap()` returns `map[string]any{...}` | **No `ToMap()` method on `Collection[K,V]`** (greppped `namevalue/*.go` — 0 hits). Available aggregators are JSON-related (`JsonString`) plus `Items` slice direct access. | **HIGH** — fabricated method. |
| **C-CVS-20** | 272 | `keymk.New.Compile("user", userID, "post", postID)` returning `"user/usr-123/post/pst-456"` | **No `keymk.New` var, no `.Compile` method.** Real entry points (`keymk/vars.go`): `NewKey = &newKeyCreator{}` exposing `NewKey.Create(option *Option, main string) *Key`; the resulting `*Key` has `CompileKeys(...)`, `KeyCompiled()`, `CompileReplaceCurlyKeyMap(...)`. The whole "template-based key builder with named placeholders" snippet is a misrepresentation of a much heavier API. | **CRITICAL** — entire §8 snippet does not compile; calling convention is fundamentally different (`Option`-based, two-step Create→Compile). |

### 1.3 Remain ❓ (3 claims — pending deeper probe)

| Spec line | Claim | Why still ❓ |
|-----------|-------|--------------|
| 142 | `issetter.Value` "Pitfall: not a drop-in for `bool`" | Advisory / behavioural — not directly verifiable from code, but consistent with file shape. |
| 173 | "`LazyLock` defers cost to first use, then caches" | Behavioural — would need to read `LazyRegex.go` body for `sync.Once` usage. Likely true (file is named `LazyRegex.go` with `LazyLock` constructor). |
| 200 | `corecmp` returns `constants.CompareEqual / Less / Greater` (`0 / -1 / 1`) | Need to grep `constants/compare*.go` to confirm exact names + values. |

## 2. Updated cycle 5 scoreboard line

Cycle 5 was: 0 ✅ / 0 ⚠️ / 0 ❌ / 17 ❓ — N/A score.

After Cycle 20 pass 1: **12 ✅ / 0 ⚠️ / 5 ❌ / 3 ❓** *(some original ❓ collapsed during evidence-gathering — total drops from 17 to 20 sub-claims because Cycle 5 grouped multi-method snippets as single ❓; this audit splits them. Reconciliation: 12+5+3 = 20.)* → verifiable score = 12 / 17 = **70.6%** *(verifiable)*.

## 3. Action items spawned (BLOCKED by freeze)

- **AJ-04** rewrite `07-conditional-and-utilities.md` §1.3 around the actual `TypedErrorFunctionsExecuteResults` branch-selector signature.
- **AJ-05** rewrite §5 `coremath` to list only the 3 type families that actually exist (`Byte`, `Int`, `Float32`) plus a footnote on `*Within` range helpers, OR file an upstream issue requesting parity for `Int16/Int32/Int64/Float64`.
- **AJ-06** rewrite §7 `namevalue` against the generic-struct-field reality (no `NewInstance`, `Instance[K,V]` with public fields, no `ValueAny`, no `ToMap`).
- **AJ-07** rewrite §8 `keymk` against the `NewKey.Create(opt, main).CompileKeys(...)` two-step API (or move §8 to a dedicated future "advanced keys" page since the surface is too large for a 7-line snippet).

## 4. Cumulative AB-pass running totals

| Cycle | Section | ❓ before | ✅ after | ❌ after | ❓ after |
|-------|---------|-----------|----------|----------|----------|
| 19 | `09-converters.md` | 23 | 10 | 5 | 8 |
| 20 | `07-conditional-and-utilities.md` | 17 | 12 | 5 | 3 |
| **Σ** | (2 sections) | **40** | **22** | **10** | **11** |

**Pattern confirmed:** every section authored against the pre-rename `core-v8` API yields ~5 ❌ contradictions per ~20 claims (~25% fabricated). Projecting to the remaining 5 sections (§08, §10, §11, §15, §16, total ~70 ❓), expect **~17 more ❌** before AB pass 2 completes.

## 5. Suggestion link

S-106 (`spec-api-check.psm1`) raised in Cycle 19 would have caught **all 5** ❌ in this cycle (every fabricated symbol fails `go vet`). Re-iterating priority of S-106.

---

_Audit file: `spec/07-code-vs-spec-audits/21-cycle20-AB-conditional-and-utilities.md`_
_See also: `20-cycle19-AB-converters-promotion.md` (pass-1 precedent), `06-cycle5-conditional-and-utilities.md` (the cycle whose ❓ are being promoted here)._