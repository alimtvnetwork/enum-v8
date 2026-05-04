# 05 — Enum System

> ✅ **Status**: filled in audit Step 5 (2026-04-23, Asia/Kuala_Lumpur).
> **Audience**: anyone defining a new enum or extending the enum framework.
> **Source**: extracted from [`/spec/00-llm-integration-guide.md`](../00-llm-integration-guide.md) §Enum System and the canonical `BasicByte` recipe.

---

## Table of Contents

1. [Architecture](#1-architecture)
2. [`enuminf` — Interface Layer](#2-enuminf--interface-layer)
3. [`enumimpl` — Implementation Engine](#3-enumimpl--implementation-engine)
4. [Defining a New Enum (full recipe)](#4-defining-a-new-enum-full-recipe)
5. [Adapting for Other Backing Types](#5-adapting-for-other-backing-types)
6. [Factory Method Reference](#6-factory-method-reference)
7. [Serialization & Comparison](#7-serialization--comparison)
8. [Testing an Enum](#8-testing-an-enum)
9. [Common Mistakes](#9-common-mistakes)

---

## 1. Architecture

The enum system uses a three-layer split that satisfies several design pillars at once:

```
┌────────────────────────────────────┐
│  enuminf (coreinterface/enuminf/)  │  ← Layer 1: pure interface contracts
│  - BaseEnumer, BasicEnumer         │
│  - BasicByteEnumer, BasicInt8Enumer│
│  - StandardEnumer, etc.            │
└────────────────────────────────────┘
                 │ implements
                 ▼
┌────────────────────────────────────┐
│  enumimpl (coreimpl/enumimpl/)     │  ← Layer 2: reusable engine
│  - enumimpl.New.BasicByte          │
│  - enumimpl.New.BasicInt8          │
│  - factory + lookup tables         │
└────────────────────────────────────┘
                 │ embeds via package-level var
                 ▼
┌────────────────────────────────────┐
│  Your enum package                 │  ← Layer 3: domain-specific enum
│  - <TypeName>.go (type + iota +    │
│    method set + predicates)        │
│  - vars.go (Ranges + BasicEnumImpl)│
└────────────────────────────────────┘

> **Filename convention:** the method-set file is named after its type (e.g. `Variant.go`, `Bracket.go`, `Architecture.go`). Across this repo's 71 enum packages, the type is conventionally named `Variant` (64 / 71); the remaining 7 use a domain-specific name (`Bracket`, `Category`, `Precedence`, `Action`, `Quote`, `ExitCode`, `Architecture`). There is **no** `consts.go` — the type and its `iota` block live alongside the method set in `<TypeName>.go`.
```

### Why three layers?

- **`enuminf`** lets consumers code against a contract (e.g. `func Process(e enuminf.BasicEnumer)`) without depending on a specific enum type. This is the **interface-first** pillar.
- **`enumimpl`** is the **shared engine** so each new enum doesn't reimplement formatting, lookup, JSON marshalling, or aliasing.
- **Your enum package** is thin — mostly delegation to `enumimpl` plus domain-specific predicates (`IsPending()`, `IsReady()`).

---

## 2. `enuminf` — Interface Layer

Located at `coreinterface/enuminf/`. Defines the contracts.

### Key interfaces

| Interface | Purpose |
|---|---|
| `BaseEnumer` | The minimum: `Value()`, `Name()`, `String()`, `IsValid()`, `IsInvalid()` |
| `BasicEnumValuer` | **Comprehensive typed value accessors** required by generic consumers — `Value()`, `ValueByte()`, `ValueInt()`, `ValueInt8/16/32/64()`, `ValueUInt16()`, `ValueString()`. *Every enum MUST implement all of them; see §Step 3 (F-V12-09 / F-NEW-07).* |
| `BasicEnumer` | Adds: `RangeNamesCsv()`, `MinMaxAny()`, `IntegerEnumRanges()`, `Format()` |
| `StandardEnumer` | Composes `BasicEnumer` + JSON marshal/unmarshal contracts |
| `BasicByteEnumer` | `BasicEnumer` + `byte`-specific accessors (`MaxByte`, `MinByte`, `RangesByte`) |
| `BasicInt8Enumer`, `BasicInt16Enumer`, `BasicInt32Enumer`, `BasicUInt16Enumer`, `BasicStringEnumer` | One per backing type |
| `OnlySupportedNamesErrorer` | Standard "value must be one of …" error formatter |
| `EnumTyper` | Returns metadata about the enum's type (used by the JSON pipeline) |

### Composition pattern

Interfaces are small and combined by composition:

```go
type BasicByteEnumer interface {
    BasicEnumer
    Max() byte
    Min() byte
    Ranges() []byte
}
```

Your enum type implements `BasicByteEnumer` (or another `Basic*Enumer`) by delegating to `enumimpl`.

---

## 3. `enumimpl` — Implementation Engine

Located at `coreimpl/enumimpl/`. Exposes a `New` struct-as-namespace var with one factory per backing type:

```go
enumimpl.New.BasicByte    // returns a *BasicByteEnumImpl factory
enumimpl.New.BasicInt8
enumimpl.New.BasicInt16
enumimpl.New.BasicInt32
enumimpl.New.BasicUInt16
enumimpl.New.BasicString
```

Each factory exposes constructors (`UsingTypeSlice`, `Default`, `CreateUsingMap`, …) that return a configured enum-impl ready to be embedded as a package-level `var`. See [§6](#6-factory-method-reference) for the full method list.

### Why `enumimpl.New` and not `enumimpl.NewBasicByte(...)`?

This follows the **`newCreator` pattern** from [`02-design-philosophy.md` §5](./02-design-philosophy.md). The `New` namespace groups all factories so IDE autocomplete on `enumimpl.New.` lists every supported backing type.

---

## 4. Defining a New Enum (full recipe)

This is the **canonical 2-file pattern** for a `byte`-backed enum. The conventional type name across this repo is `Variant`, so the recipe below uses that name (your domain may use a more descriptive name — see §4.3). Adapt only what's marked `// CHANGE`.

### 4.1 Sentinel-first rule  *(F-NEW-08)*

The first iota constant **must occupy the zero value** of the backing type so an unset variable is detectable. The conventional name is `Invalid`, but several equivalent **sentinel** names are accepted across the codebase, depending on domain semantics:

| Sentinel name | Used by (examples) | When to prefer |
|---|---|---|
| `Invalid` | most enums (61 / 71) | the default — use it unless a domain term is clearer |
| `Unspecified` | `revokereason` | wire-protocol / RFC vocabulary |
| `Uninitialized` | `envtype` | distinguishes "never set" from "set to default" |
| `Default` | `compresslevels`, `scripttype`, `sqljointype`, `taskpriority` | the zero value is itself a meaningful default |
| (domain term) | `compressformats.Zip`, `logtype.Silent`, `strtype.<...>` | the domain has a natural zero-meaning value |

> **Signed-int exception:** `inttype.Variant` declares `InvalidIndex Variant = -1` because its backing type (`int`) uses `-1` as the "no index" sentinel. When the backing type is signed and `-1` is the conventional unset value, name the first constant `InvalidIndex` (or similar) and assign `= -1` explicitly. This is the **only** documented case where the sentinel does not occupy the Go zero value.

Whatever the name, the rule is: **the first declared iota constant is the sentinel**, and `IsValid()` / `IsInvalid()` test against it.

### 4.2 Step 1 — `Variant.go` (type + iota + method set)

Type declaration, iota constants, and method set all live in this single file:

```go
package status                                  // CHANGE: package name

import "github.com/alimtvnetwork/core-v9/coreinterface/enuminf"

// CHANGE: type name + backing type. Convention: name it `Variant`.
type Variant byte

const (
    Invalid Variant = iota                      // sentinel — see §4.1
    Pending
    Ready
    Failed
)

// --- Value accessors (BasicEnumValuer) — ALL required by enuminf.BasicEnumValuer ---
// DO NOT remove any of these even if they seem redundant; the interface needs them all.
func (it Variant) Value() byte         { return byte(it) }
func (it Variant) ValueByte() byte     { return byte(it) }
func (it Variant) ValueInt() int       { return int(it) }
func (it Variant) ValueInt8() int8     { return int8(it) }
func (it Variant) ValueInt16() int16   { return int16(it) }
func (it Variant) ValueUInt16() uint16 { return uint16(it) }
func (it Variant) ValueInt32() int32   { return int32(it) }
func (it Variant) ValueString() string { return BasicEnumImpl.ToNumberString(it.Value()) }

// --- Naming ---
func (it Variant) Name() string           { return BasicEnumImpl.ToEnumString(it.Value()) }
func (it Variant) String() string         { return BasicEnumImpl.ToEnumString(it.Value()) }
func (it Variant) TypeName() string       { return BasicEnumImpl.TypeName() }
func (it Variant) NameValue() string      { return BasicEnumImpl.NameWithValue(it.Value()) }
func (it Variant) ToNumberString() string { return BasicEnumImpl.ToNumberString(it.Value()) }

// --- Equality ---
func (it Variant) IsNameEqual(name string) bool { return it.Name() == name }
func (it Variant) IsAnyNamesOf(names ...string) bool {
    n := it.Name()
    for _, name := range names {
        if name == n {
            return true
        }
    }
    return false
}

// --- Valid / Invalid (test against sentinel — see §4.1) ---
func (it Variant) IsValid() bool   { return it != Invalid }
func (it Variant) IsInvalid() bool { return it == Invalid }

// --- Range info (BasicEnumer) ---
func (it Variant) RangeNamesCsv() string            { return BasicEnumImpl.RangeNamesCsv() }
func (it Variant) MinMaxAny() (min, max any)        { return BasicEnumImpl.MinMaxAny() }
func (it Variant) MinValueString() string           { return BasicEnumImpl.MinValueString() }
func (it Variant) MaxValueString() string           { return BasicEnumImpl.MaxValueString() }
func (it Variant) MaxInt() int                      { return BasicEnumImpl.MaxInt() }
func (it Variant) MinInt() int                      { return BasicEnumImpl.MinInt() }
func (it Variant) RangesDynamicMap() map[string]any { return BasicEnumImpl.RangesDynamicMap() }
func (it Variant) AllNameValues() []string          { return BasicEnumImpl.AllNameValues() }
func (it Variant) IntegerEnumRanges() []int         { return BasicEnumImpl.IntegerEnumRanges() }

// --- OnlySupportedNamesErrorer ---
func (it Variant) OnlySupportedErr(names ...string) error {
    return BasicEnumImpl.OnlySupportedErr(names...)
}
func (it Variant) OnlySupportedMsgErr(message string, names ...string) error {
    return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

// --- Format — placeholder keys: {type-name}, {name}, {value} ---
func (it Variant) Format(format string) string {
    return BasicEnumImpl.Format(format, it.Value())
}

// --- Type-specific (BasicByteEnumer) ---
func (it Variant) MaxByte() byte      { return BasicEnumImpl.Max() }
func (it Variant) MinByte() byte      { return BasicEnumImpl.Min() }
func (it Variant) RangesByte() []byte { return BasicEnumImpl.Ranges() }

// --- Range validation ---
func (it Variant) IsValidRange() bool           { return BasicEnumImpl.IsValidRange(it.Value()) }
func (it Variant) IsInvalidRange() bool         { return !it.IsValidRange() }
func (it Variant) RangesInvalidMessage() string { return BasicEnumImpl.RangesInvalidMessage() }
func (it Variant) RangesInvalidErr() error      { return BasicEnumImpl.RangesInvalidErr() }

// --- String ranges ---
func (it Variant) StringRanges() []string    { return BasicEnumImpl.StringRanges() }
func (it Variant) StringRangesPtr() []string { return BasicEnumImpl.StringRangesPtr() }

// --- JSON ---
func (it Variant) MarshalJSON() ([]byte, error) {
    return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}
func (it *Variant) UnmarshalJSON(data []byte) error {
    val, err := it.UnmarshallEnumToValue(data)
    if err == nil {
        *it = Variant(val)
    }
    return err
}
func (it Variant) UnmarshallEnumToValue(data []byte) (byte, error) {
    return BasicEnumImpl.UnmarshallToValue(true, data)
}
// ⚠️ Spelling note (F-V12-06 fix):
//   - Standard library / json.Unmarshaler interface methods use ONE 'l': UnmarshalJSON.
//   - The enum engine's typed converter helpers use TWO 'l's: UnmarshallEnumToValue,
//     UnmarshallToValue, UnmarshallEnumToValueInt8, UnmarshallEnumToValueInt16, etc.
// This is a deliberate, historical naming convention to distinguish enum-engine
// helpers from the stdlib JSON contract. DO NOT "correct" the double-l spelling
// — `enumimpl` exposes only the double-l names; renaming will not compile.

// --- EnumType ---
func (it Variant) EnumType() enuminf.EnumTyper {
    return BasicEnumImpl.EnumType()
}

// --- Domain-specific predicates (CHANGE: per enum) ---
func (it Variant) IsPending() bool { return it == Pending }
func (it Variant) IsReady() bool   { return it == Ready }
func (it Variant) IsFailed() bool  { return it == Failed }
```

> 🧱 **F-NEW-07 — Why so many `Value<Type>()` methods?**
> Every method in the *Value accessors* block is **REQUIRED** to satisfy the generic `enuminf.BasicEnumValuer` interface, which downstream generic consumers (validators, converters, JSON encoders) rely on. They are intentional and **MUST NOT be deleted** even when they look like trivial conversions.

### 4.3 Step 2 — `vars.go` (Ranges + BasicEnumImpl)

Define the human-readable name slice and instantiate the impl:

```go
package status

import "github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"

var (
    Ranges = [...]string{                       // CHANGE: keys must match consts above
        Invalid: "Invalid",
        Pending: "Pending",
        Ready:   "Ready",
        Failed:  "Failed",
    }

    // Recommended: DefaultAllCases derives the type name from the first item
    // via the engine itself (no internal-package import needed) AND registers
    // the "all cases" slice for downstream tooling.
    BasicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(Invalid, Ranges[:])
)
```

> ⚠️ **Do NOT import `core-v9/internal/reflectinternal`.** It lives behind Go's `internal/` boundary and is not importable from this module. Earlier drafts of this spec showed `reflectinternal.TypeName(Invalid)` — that example was unrunnable. Use one of the supported patterns instead:
>
> | Pattern | Use when |
> |---|---|
> | `enumimpl.New.BasicByte.DefaultAllCases(Invalid, Ranges[:])` | **Recommended.** Engine derives type name from `Invalid`'s reflected type. |
> | `enumimpl.New.BasicByte.UsingTypeSlice("Variant", Ranges[:])` | You'd rather pass the type name as a string literal. |

### 4.4 Why so many methods?

This is the **explicit-API** pillar — every supported operation is a named method, so callers do not need to know which interface they are satisfying. IDE autocomplete on a `Variant` value lists every available query.

### 4.5 Predicate file-split guideline

Predicate methods (`Is<Name>()`) generally stay in `<TypeName>.go` alongside the rest of the method set, regardless of count — the largest enum in the repo (`pathpatterntype`, 113 constants) keeps every predicate in `Variant.go`. If a single predicate body grows past ~20 lines (rare — usually it implies the predicate should be extracted to a small helper), split *that one* into its own `Is<Name>.go` file per [Pillar 1 of the design philosophy](./02-design-philosophy.md#pillar-1--one-file-per-function). Otherwise, prefer cohesion over file fan-out.

---

## 5. Adapting for Other Backing Types

The recipe above is `byte`-backed. For other backings, replace:

| Backing | `Value()` returns | Unmarshal method name | Type-specific accessors |
|---|---|---|---|
| `byte` | `byte` | `UnmarshallEnumToValue` | `MaxByte`, `MinByte`, `RangesByte` |
| `int8` | `int8` | `UnmarshallEnumToValueInt8` | `MaxInt8`, `MinInt8`, `RangesInt8` |
| `int16` | `int16` | `UnmarshallEnumToValueInt16` | `MaxInt16`, `MinInt16`, `RangesInt16` |
| `int32` | `int32` | `UnmarshallEnumToValueInt32` | `MaxInt32`, `MinInt32`, `RangesInt32` |
| `uint16` | `uint16` | `UnmarshallEnumToValueUInt16` | `MaxUInt16`, `MinUInt16`, `RangesUInt16` |
| `string` | `string` | `UnmarshallEnumToValueString` | `Ranges() []string` |

Also swap `enumimpl.New.BasicByte` for the matching factory (see next section).

---

## 6. Factory Method Reference

Each `enumimpl.New.Basic<Type>` factory exposes the same constructor surface:

| Method | Description |
|---|---|
| `UsingTypeSlice(typeName, names[])` | Contiguous iota from a string slice (most common) |
| `Default(firstItem, names[])` | Same, but infers `typeName` via reflection on `firstItem` |
| `DefaultWithAliasMap(firstItem, names[], aliasMap)` | Contiguous + alias name → canonical name lookup |
| `CreateUsingMap(typeName, map[T]string)` | Non-contiguous — explicit value-to-name pairs |
| `CreateUsingMapPlusAliasMap(typeName, map[T]string, aliasMap)` | Explicit values + aliases |

### When to pick which

- **`UsingTypeSlice`** — your enum is contiguous (`iota`-style) and you already have the type name string.
- **`Default`** — same but you'd rather have the impl reflect-derive `"Status"` from the first value.
- **`CreateUsingMap`** — sparse / explicit values (e.g. wire-protocol codes like `0x10, 0x20, 0xFF`).
- **`*WithAliasMap` / `*PlusAliasMap`** — wire format accepts multiple synonyms (e.g. `"ok"`, `"OK"`, `"success"` all map to `Ready`).

---

## 7. Serialization & Comparison

### JSON

`MarshalJSON` outputs the **enum name as a JSON string** (e.g. `"Pending"`), not the numeric value. This makes payloads readable and resistant to enum-value reordering.

`UnmarshalJSON` accepts:
- The canonical name (`"Pending"`)
- Any registered alias (if `*WithAliasMap` was used)
- The numeric string form (`"1"`) — the second `bool` arg to `UnmarshallToValue` controls whether numeric input is accepted

#### ⚠️ Serialization Asymmetry  *(F-NEW-05 fix — read carefully)*

The read and write contracts are **deliberately asymmetric**:

| Direction | Accepts / produces |
|---|---|
| `MarshalJSON` (write) | **MUST** emit the canonical enum name as a JSON string (e.g. `"Pending"`). Never numeric. |
| `UnmarshalJSON` (read) | **MAY** accept canonical name, registered aliases, OR numeric-string form (e.g. `"1"`). |

Implications for downstream services:
- **Consumers of JSON produced by this library** should expect string names only — never code branches that parse numeric enum values from our output.
- **This library** can ingest numeric-string enums from third-party systems for compatibility, but does not advertise them.
- Round-tripping `Marshal → Unmarshal` is lossless; round-tripping `Unmarshal("1") → Marshal` returns `"Pending"`, **not** `"1"`.

### Comparison

Enums are value types — compare with `==`:

```go
if status == status.Ready { ... }
if status.IsAnyNamesOf("Ready", "Pending") { ... }
```

Do **not** compare via `Value()` unless you need to interop with a numeric API.

---

## 8. Testing an Enum

A new enum needs three test groups under `tests/integratedtests/<pkg>tests/`:

| Test family | Style | Verifies |
|---|---|---|
| `<EnumType>_Verification_test.go` | A (`CaseV1` + `args.Map`) | Range names, min/max, format strings, JSON round-trip |
| `<EnumType>_NilReceiver_test.go` | `CaseNilSafe` | Pointer receiver methods (`UnmarshalJSON`) don't panic on nil |
| `<EnumType>_OnlySupportedErr_test.go` | A | Error formatting is stable (regex-tested) |

See [`13-testing-patterns.md`](./13-testing-patterns.md) and [`/spec/06-testing-guidelines/02-test-case-types.md`](../06-testing-guidelines/02-test-case-types.md).

---

## 9. Common Mistakes

| Mistake | Why bad | Fix |
|---|---|---|
| First constant is not `Invalid` | Zero value of the type isn't catchable as "unset" | Always start with `Invalid Status = iota` |
| Forgetting `Ranges[:]` (passing `Ranges` instead of `Ranges[:]`) | Array vs slice — compile error | Always pass `Ranges[:]` |
| Implementing only `Value()` and `Name()` | Doesn't satisfy `BasicEnumer` — won't pass to functions expecting the interface | Implement the full method set in §4 |
| Hand-writing JSON marshalling | Drifts from canonical format used elsewhere | Always delegate to `BasicEnumImpl.ToEnumJsonBytes` / `UnmarshallToValue` |
| Comparing via `Value()` | Loses type safety and IDE goto-definition | Use `==` directly on the enum type |
| Skipping nil-receiver tests for `*UnmarshalJSON` | Panics in production when JSON parser hands a nil receiver | Add a `CaseNilSafe` test |

---

## See Also

- [`02-design-philosophy.md`](./02-design-philosophy.md) §3 (struct-as-namespace), §5 (newCreator pattern), §7 (interface-first)
- [`03-import-conventions.md`](./03-import-conventions.md) — Why `enuminf` is L1 and `enumimpl` is L4
- [`13-testing-patterns.md`](./13-testing-patterns.md) — Test style matrix for enum verification
- [`/spec/06-testing-guidelines/02-test-case-types.md`](../06-testing-guidelines/02-test-case-types.md) — `CaseNilSafe` for nil-receiver tests
- [`/spec/00-llm-integration-guide.md` §Enum System](../00-llm-integration-guide.md#enum-system-enuminf--enumimpl) — Quick reference
