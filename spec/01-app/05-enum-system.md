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
│  - consts.go (the type + iota)     │
│  - vars.go (BasicEnumImpl + Ranges)│
│  - <Type>.go (method set)          │
└────────────────────────────────────┘
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

This is the **canonical 3-file pattern** for a `byte`-backed enum named `Status`. Adapt only what's marked `// CHANGE`.

### Step 1 — `consts.go`

Define the type and the iota constants. The first constant **must be `Invalid`** (or equivalent zero-value) so an unset variable is detectable.

```go
package status                                  // CHANGE: package name

type Status byte                                // CHANGE: type name + backing type

const (
    Invalid Status = iota                       // CHANGE: keep "Invalid" first
    Pending
    Ready
    Failed
)
```

### Step 2 — `vars.go`

Define the human-readable name slice and instantiate the impl.

```go
package status

import (
    "github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
    "github.com/alimtvnetwork/core-v8/internal/reflectinternal"
)

var (
    Ranges = [...]string{                       // CHANGE: keys must match consts.go
        Invalid: "Invalid",
        Pending: "Pending",
        Ready:   "Ready",
        Failed:  "Failed",
    }

    BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
        reflectinternal.TypeName(Invalid),       // auto-derives "Status"
        Ranges[:],
    )
)
```

### Step 3 — `Status.go` (method set)

All methods below are required to satisfy `StandardEnumer` + the type-specific `Basic*Enumer`. They are mostly one-line delegations to `BasicEnumImpl`.

> 🧱 **F-NEW-07 — Why so many `Value<Type>()` methods?**
> Every method in the *Value accessors* block is **REQUIRED** to satisfy the generic `enuminf.BasicEnumValuer` interface, which downstream generic consumers (validators, converters, JSON encoders) rely on. They are intentional and **MUST NOT be deleted** even when they look like trivial conversions.

```go
package status

import "github.com/alimtvnetwork/core-v8/coreinterface/enuminf"

// --- Value accessors (BasicEnumValuer) — ALL required by enuminf.BasicEnumValuer ---
// DO NOT remove any of these even if they seem redundant; the interface needs them all.
func (it Status) Value() byte         { return byte(it) }
func (it Status) ValueByte() byte     { return byte(it) }
func (it Status) ValueInt() int       { return int(it) }
func (it Status) ValueInt8() int8     { return int8(it) }
func (it Status) ValueInt16() int16   { return int16(it) }
func (it Status) ValueUInt16() uint16 { return uint16(it) }
func (it Status) ValueInt32() int32   { return int32(it) }
func (it Status) ValueString() string { return BasicEnumImpl.ToNumberString(it.Value()) }

// --- Naming ---
func (it Status) Name() string           { return BasicEnumImpl.ToEnumString(it.Value()) }
func (it Status) String() string         { return BasicEnumImpl.ToEnumString(it.Value()) }
func (it Status) TypeName() string       { return BasicEnumImpl.TypeName() }
func (it Status) NameValue() string      { return BasicEnumImpl.NameWithValue(it.Value()) }
func (it Status) ToNumberString() string { return BasicEnumImpl.ToNumberString(it.Value()) }

// --- Equality ---
func (it Status) IsNameEqual(name string) bool { return it.Name() == name }
func (it Status) IsAnyNamesOf(names ...string) bool {
    n := it.Name()
    for _, name := range names {
        if name == n {
            return true
        }
    }
    return false
}

// --- Valid / Invalid ---
func (it Status) IsValid() bool   { return it != Invalid }
func (it Status) IsInvalid() bool { return it == Invalid }

// --- Range info (BasicEnumer) ---
func (it Status) RangeNamesCsv() string            { return BasicEnumImpl.RangeNamesCsv() }
func (it Status) MinMaxAny() (min, max any)        { return BasicEnumImpl.MinMaxAny() }
func (it Status) MinValueString() string           { return BasicEnumImpl.MinValueString() }
func (it Status) MaxValueString() string           { return BasicEnumImpl.MaxValueString() }
func (it Status) MaxInt() int                      { return BasicEnumImpl.MaxInt() }
func (it Status) MinInt() int                      { return BasicEnumImpl.MinInt() }
func (it Status) RangesDynamicMap() map[string]any { return BasicEnumImpl.RangesDynamicMap() }
func (it Status) AllNameValues() []string          { return BasicEnumImpl.AllNameValues() }
func (it Status) IntegerEnumRanges() []int         { return BasicEnumImpl.IntegerEnumRanges() }

// --- OnlySupportedNamesErrorer ---
func (it Status) OnlySupportedErr(names ...string) error {
    return BasicEnumImpl.OnlySupportedErr(names...)
}
func (it Status) OnlySupportedMsgErr(message string, names ...string) error {
    return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

// --- Format — placeholder keys: {type-name}, {name}, {value} ---
func (it Status) Format(format string) string {
    return BasicEnumImpl.Format(format, it.Value())
}

// --- Type-specific (BasicByteEnumer) ---
func (it Status) MaxByte() byte      { return BasicEnumImpl.Max() }
func (it Status) MinByte() byte      { return BasicEnumImpl.Min() }
func (it Status) RangesByte() []byte { return BasicEnumImpl.Ranges() }

// --- Range validation ---
func (it Status) IsValidRange() bool           { return BasicEnumImpl.IsValidRange(it.Value()) }
func (it Status) IsInvalidRange() bool         { return !it.IsValidRange() }
func (it Status) RangesInvalidMessage() string { return BasicEnumImpl.RangesInvalidMessage() }
func (it Status) RangesInvalidErr() error      { return BasicEnumImpl.RangesInvalidErr() }

// --- String ranges ---
func (it Status) StringRanges() []string    { return BasicEnumImpl.StringRanges() }
func (it Status) StringRangesPtr() []string { return BasicEnumImpl.StringRangesPtr() }

// --- JSON ---
func (it Status) MarshalJSON() ([]byte, error) {
    return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}
func (it *Status) UnmarshalJSON(data []byte) error {
    val, err := it.UnmarshallEnumToValue(data)
    if err == nil {
        *it = Status(val)
    }
    return err
}
func (it Status) UnmarshallEnumToValue(data []byte) (byte, error) {
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
func (it Status) EnumType() enuminf.EnumTyper {
    return BasicEnumImpl.EnumType()
}

// --- Domain-specific predicates (CHANGE: per enum) ---
func (it Status) IsPending() bool { return it == Pending }
func (it Status) IsReady() bool   { return it == Ready }
func (it Status) IsFailed() bool  { return it == Failed }
```

### Why so many methods?

This is the **explicit-API** pillar — every supported operation is a named method, so callers do not need to know which interface they are satisfying. IDE autocomplete on a `Status` value lists every available query.

### Predicate file-split rule

Predicate methods (`Is<Name>()`) **<20 lines each** may share `[Type].go`. Once you have **>6 predicates** OR any single predicate **exceeds 20 lines**, split each into its own `Is<Name>.go` file per [Pillar 1 of the design philosophy](./02-design-philosophy.md#pillar-1--one-file-per-function). This avoids file-explosion for small enums while preserving the one-file-per-function rule for non-trivial method sets.

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
