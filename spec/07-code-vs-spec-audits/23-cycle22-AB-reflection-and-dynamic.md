# Cycle 22 — AB Pass 4 — `spec/01-app/10-reflection-and-dynamic.md`

> **Date:** 2026-05-06
> **Auditor:** Lovable agent
> **Upstream pinned:** `github.com/alimtvnetwork/core-v9 v1.5.8` (clone `/tmp/core-v9-upstream`)
> **Method:** Promote 15 ❓ from Cycle 8 baseline against ground-truth source.

## Result

| Total claims | ✅ | ❌ | ❓ | Verifiable score |
|---|---|---|---|---|
| 19 | 5 | **8** | 6 | **38.5 %** |

> **Second-worst-drift section in the project (after §08 33.3 %).** The entire `coredynamic` package documented in §2 **does not exist** in upstream `core-v9 v1.5.8`. `reflectcore` (§3) is a thin re-export shim over `internal/reflectinternal`, not the predicate library described.

## Promotions

### ✅ Match (5)

| # | Claim | Evidence |
|---|-------|----------|
| 1 | `reflectcore` package exists | `/tmp/core-v9-upstream/reflectcore/vars.go` |
| 2 | `internal/reflectinternal` exists | `/tmp/core-v9-upstream/internal/reflectinternal/` |
| 3 | `reflectcore` re-exports `TypeName` | `vars.go:32 TypeName = reflectinternal.TypeName` |
| 4 | `reflectcore` re-exports `Is` predicate aggregate | `vars.go:31 Is = reflectinternal.Is` |
| 5 | `reflectinternal` has type-cache / GetFunc helpers | `internal/reflectinternal/getFunc.go` |

### ❌ Contradictions (8) — all NEW, severity mix 5× CRITICAL / 3× HIGH

| ID | Severity | Spec claim | Upstream reality |
|----|----------|------------|------------------|
| **C-CVS-29** | **CRITICAL** | §2 entire `coredynamic` package — `coredynamic.InvokeMethod`, `HasMethod`, `MethodNames`, `GetField`, `SetField`, `AllFields`, `TypeName`, `TypeFullName`, `IsNullOrUndefined` | **No `coredynamic/` directory exists** in `core-v9 v1.5.8`. `grep -rln coredynamic` upstream returns **zero source files** (only `.lovable/` notes). The entire §2 API surface is fabricated. |
| **C-CVS-30** | **CRITICAL** | §3.1 `reflectcore.IsPointer / IsStruct / IsSlice / IsMap / IsFunc / IsChannel / IsInterface` as top-level predicates | `reflectcore/vars.go` exposes only an aggregate var `Is = reflectinternal.Is` — predicates are reached via `reflectcore.Is.Pointer(v)` etc., not as bare functions. The bare-function form does not compile. |
| **C-CVS-31** | **CRITICAL** | §3.2 `reflectcore.WalkFields(target, func(name string, value any))` | Symbol does not exist in `reflectcore/`. Field walking lives in `reflectcore/reflectmodel/FieldProcessor.go` with a different shape (`*FieldProcessor` with `IsFieldType`/`IsFieldKind` methods). |
| **C-CVS-32** | **CRITICAL** | §3.2 `reflectcore.GetTag(target, "Name", "json")` | Symbol does not exist in upstream. |
| **C-CVS-33** | **CRITICAL** | §3.3 `reflectcore.DerefAll(ptr)` | Symbol does not exist in upstream. |
| **C-CVS-34** | HIGH | §4 "`internal/reflectinternal` is **off-limits to consumers** … reserved for framework internals" + §1 "consumer code … MUST NOT import `internal/reflectinternal`" | False as stated: `reflectcore/vars.go` **publicly re-exports 15 internal symbols** (`Converter, Utils, Looper, CodeStack, GetFunc, Is, TypeName, TypeNames, TypeNamesString, TypeNamesReferenceString, ReflectType, ReflectGetter, ReflectGetterUsingReflectValue, SliceConverter, MapConverter`). The "internal" boundary applies only to *direct* import; the reflectcore facade is the supported public form, but the spec's framing implies a behavioural firewall that is not present. |
| **C-CVS-35** | HIGH | §1 "panics are wrapped into `errcore` errors" via `coredynamic`/`reflectcore` | Unverifiable for `coredynamic` (does not exist); `reflectcore` re-exports raw `reflectinternal` symbols with no error-wrapping facade. |
| **C-CVS-36** | HIGH | §7 mistake row "Calling `InvokeMethod` without checking `HasMethod` first" | Both methods do not exist anywhere upstream — entire row is fabricated guidance. |

### ❓ Still unverifiable (6)

| # | Claim | Why |
|---|-------|-----|
| 1 | §1 motivation prose ("verbose, panics liberally, loses type info") | Subjective framing |
| 2 | §4 "unsafe pointer arithmetic for `corejson` fast-path (avoids 1 alloc per Marshal)" | No `corejson/` package exists upstream either; needs separate audit |
| 3 | §4 "type-cache keyed on `reflect.Type`" | Plausible but not a single discoverable symbol |
| 4 | §6 "10–100× slower than direct calls" | Benchmark claim, no benchmarks in repo |
| 5 | §6 "Lazy binding via `coreonce`" | `coreonce` not present in `core-v9 v1.5.8` root listing — pending §06 reaudit |
| 6 | §7 "use `isany.DeepEqual`" | `isany` exists but `DeepEqual` symbol not yet probed |

## New AJ items (BLOCKED on `spec/01-app/` 🧊 freeze)

| ID | File | Action |
|----|------|--------|
| **AJ-15** | `10-reflection-and-dynamic.md` §2 | Delete entire `coredynamic` section OR rewrite as "future API; not yet shipped" with explicit upstream-absence note |
| **AJ-16** | §3.1 | Rewrite predicates to `reflectcore.Is.Pointer(v)` form |
| **AJ-17** | §3.2 | Replace `WalkFields` / `GetTag` with `reflectcore/reflectmodel.FieldProcessor` API |
| **AJ-18** | §3.3 | Delete `DerefAll` snippet OR document upstream alternative |
| **AJ-19** | §1 + §4 | Rewrite "off-limits" framing — clarify `reflectcore` is the public facade re-exporting `reflectinternal`, not a separate behavioural layer |
| **AJ-20** | §7 | Delete the `InvokeMethod` row from common-mistakes table |

## Cumulative AB metrics

| Pass | Section | Score | NEW ❌ |
|------|---------|-------|--------|
| 1 | §09 converters | 66.7 % | 5 |
| 2 | §07 conditional & utilities | 70.6 % | 5 |
| 3 | §08 validators | 33.3 % | 8 |
| **4** | **§10 reflection & dynamic** | **38.5 %** | **8** |
| | **Cumulative ❌** | | **26** |

Fabrication rate is now **~45 %** of all AB-promoted claims (29 ❌ + 31 ✅).

## Recommendation

S-106 (`scripts/spec-api-check.psm1`) is now **mandatory** before any AJ rewrite. Without it, AJ-01..20 risk introducing the same fabrication pattern that produced 26 contradictions across 4 chapters.
