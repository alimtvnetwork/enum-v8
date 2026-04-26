# 07 — Conditional & Utilities

> ✅ **Status**: filled in audit Step 5b (2026-04-23, Asia/Kuala_Lumpur).
> **Audience**: anyone using ternary helpers, type predicates, regex, sorting, or the misc utility packages.
> **Source**: extracted from [`/spec/00-llm-integration-guide.md`](../00-llm-integration-guide.md) §conditional + §Utility Packages.

---

## Table of Contents

1. [`conditional` — Ternary & Nil-Safe](#1-conditional--ternary--nil-safe)
2. [`isany` — Type Predicates](#2-isany--type-predicates)
3. [`issetter` — 6-Valued Boolean](#3-issetter--6-valued-boolean)
4. [`regexnew` — Lazy-Compiled Regex](#4-regexnew--lazy-compiled-regex)
5. [`coremath`, `corecmp`, `coresort`](#5-coremath-corecmp-coresort)
6. [`corefuncs` — Function Wrappers](#6-corefuncs--function-wrappers)
7. [`namevalue` — Name-Value Pairs](#7-namevalue--name-value-pairs)
8. [`keymk` — Key Compilation](#8-keymk--key-compilation)
9. [Decision Matrix](#9-decision-matrix)

---

## 1. `conditional` — Ternary & Nil-Safe

Go has no ternary operator and no built-in nil-safe deref. `conditional` fills both gaps.

### 1.1 Generic Base Functions

```go
import "github.com/alimtvnetwork/core-v8/conditional"

// Ternary
result := conditional.If[int](isReady, 200, 500)

// Lazy ternary — only evaluates the chosen branch
val := conditional.IfFunc[string](ok,
    func() string { return expensiveCall() },
    func() string { return "fallback" },
)

// Nil-safe deref with default
val := conditional.NilDef[int](ptr, 42)        // *ptr or 42
p   := conditional.NilDefPtr[string](ptr, "x") // ptr or &"x"

// Zero-value deref (no default needed)
active := conditional.ValueOrZero[bool](flagPtr) // *flagPtr or false
```

### 1.2 Typed Wrappers (15 primitive types)

Each primitive type gets 11 type-specific functions so callers don't need a type parameter:

```go
conditional.IfInt(cond, 2, 7)
conditional.IfFuncString(ok, trueFunc, falseFunc)
conditional.NilDefFloat64(ptr, 3.14)
conditional.ValueOrZeroBool(flagPtr)
```

Supported primitive types: `bool`, `byte`, `string`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`.

### 1.3 Batch Execution

```go
// Run error-returning functions, aggregate failures
err := conditional.ErrorFunc(fn1, fn2, fn3)

// Typed variant: collects results + first error
results, err := conditional.TypedErrorFunctionsExecuteResults[string](fn1, fn2)
```

### When to Use Lazy vs Eager

| Need | Use |
|---|---|
| Both branches are cheap values | `If` / `IfInt` (eager) |
| One branch is expensive (DB, network, allocation) | `IfFunc` / `IfFuncString` (lazy) |
| Pointer might be nil, want a fallback value | `NilDef*` |
| Pointer might be nil, want zero on nil | `ValueOrZero*` |

> **Rule**: Default to eager. Switch to `*Func` only when the un-chosen branch has measurable cost.

---

## 2. `isany` — Type Predicates

Reflection-based predicates over `any` values. Use when you need to make decisions about untyped data — most often inside generic-but-not-typed pipelines (e.g. JSON, dynamic dispatch, test framework).

```go
import "github.com/alimtvnetwork/core-v8/isany"

isany.Null(val)        // true if nil
isany.Defined(val)     // true if non-nil
isany.Zero(val)        // true if zero value of its type
isany.DeepEqual(a, b)  // reflect.DeepEqual wrapper
isany.JsonEqual(a, b)  // serialize-and-compare via corejson
```

### Why not just `val == nil`?

Because of Go's typed-nil trap: a `var p *MyStruct = nil` stored in an `any` is **not** equal to bare `nil` via `==`. `isany.Null` handles this correctly.

```go
var p *MyStruct          // nil pointer
var a any = p            // non-nil interface holding a nil pointer
a == nil                 // false (typed nil)
isany.Null(a)            // true
```

> **Rule**: When checking nil on an `any` parameter, always use `isany.Null` / `isany.Defined`. Never `== nil`.

---

## 3. `issetter` — 6-Valued Boolean

Go's `bool` has only true/false. `issetter.Value` is a `byte`-backed enum with **six** states, designed for "tri-state plus metadata" cases like config overrides, feature flags, and the `BaseTestCase.IsEnable` field (see [`/spec/06-testing-guidelines/02-test-case-types.md` §Style B](../06-testing-guidelines/02-test-case-types.md#style-b--basetestcase--testwrapper)).

| Value | Meaning |
|---|---|
| `Uninitialized` (0) | Never touched — distinct from "explicitly false" |
| `True` (1) | Set to true |
| `False` (2) | Set to false |
| `Unset` (3) | Explicitly cleared |
| `Set` (4) | Explicitly set, value unspecified (use with `True` semantics) |
| `Wildcard` (5) | Matches anything (used in test selectors) |

```go
import "github.com/alimtvnetwork/core-v8/issetter"

status := issetter.True
status.IsOn()           // true (True or Set)
status.IsOff()          // false (False or Unset)
status.HasInitialized() // true (any non-Uninitialized state)
```

### When to Use

- A struct field where "default" must be distinguishable from "explicitly false" (e.g. config layering).
- Test enabled-flags where `Uninitialized` should mean "skip silently" but `False` means "user explicitly disabled".
- Wildcard matching (`Wildcard` matches anything in the consumer's predicate).

> **Pitfall**: `issetter.Value` is **not** a drop-in for `bool`. Read the `IsOn` / `IsOff` semantics — they collapse 6 states into a 2-state predicate.

---

## 4. `regexnew` — Lazy-Compiled Regex

Wraps `regexp.Compile` with `sync.Once` so a regex is compiled at most once across all callers.

```go
import "github.com/alimtvnetwork/core-v8/regexnew"

// Package-level (init context — no lock needed)
var digitRegex = regexnew.New.Lazy(`\d+`)

// Method-local (compile on first IsMatch / FindAll, then cached)
lazy := regexnew.New.LazyLock(`^[a-z]+\d+$`)
lazy.IsMatch(input)
lazy.IsApplicable()  // true if pattern compiled successfully
lazy.IsDefined()     // true if pattern is non-empty
lazy.IsFailedMatch(input) // inverse
```

### `Lazy` vs `LazyLock`

| Constructor | When to use |
|---|---|
| `New.Lazy` | Package-level `var` declarations — initialization order guarantees safety |
| `New.LazyLock` | Inside methods or local scopes — uses `sync.Once` for thread-safe lazy compile |

### Why this exists

Compiling regex at package init slows startup. Compiling on every call wastes CPU. `LazyLock` defers the cost to first use, then caches.

> See the live test patterns in [`/spec/06-testing-guidelines/02-test-case-types.md` §MapGherkins](../06-testing-guidelines/02-test-case-types.md#mapgherkins--preferred-for-multi-field-tests) — `regexnew` is the canonical example for `MapGherkins`.

---

## 5. `coremath`, `corecmp`, `coresort`

Three small packages that fill stdlib gaps with typed, allocation-free helpers.

### `coremath`

```go
import "github.com/alimtvnetwork/core-v8/coremath"
// Min/Max for byte, int, int16, int32, int64, float32, float64
```

Type-specific to avoid the boxing/conversion costs of `math.Max` (which is `float64`-only).

### `corecmp`

```go
import "github.com/alimtvnetwork/core-v8/corecmp"
// Byte, Integer, Integer8/16/32/64, String, Time
// Plus pointer variants (e.g. CompareIntegerPtr)
```

Returns `constants.CompareEqual` / `CompareLess` / `CompareGreater` (`0` / `-1` / `1`). Use when implementing `sort.Interface.Less`, building diffs, or routing on three-way comparison.

### `coresort`

```go
import "github.com/alimtvnetwork/core-v8/coresort/strsort"

fruits := []string{"banana", "mango", "apple"}
strsort.Quick(&fruits)    // ascending
strsort.QuickDsc(&fruits) // descending
```

Pointer-receiver mutators — sort in place without taking a copy.

---

## 6. `corefuncs` — Function Wrappers

Reflection-based helpers for working with functions as values. Used heavily by the test framework's diagnostic output.

```go
import "github.com/alimtvnetwork/core-v8/corefuncs"

corefuncs.GetFuncName(fn)        // "MyFunc"
corefuncs.GetFuncFullName(fn)    // "github.com/.../pkg.MyFunc"

// Wrapper types for common function signatures
corefuncs.ActionReturnsErrorFuncWrapper
corefuncs.InOutErrFuncWrapper
```

### When to use

- Error messages that should name the failing function.
- Test diagnostic output (the `coretests.GetAssert` helpers use this internally).
- Building plugin / dispatcher systems where functions are first-class values.

---

## 7. `namevalue` — Name-Value Pairs

Two simple types: a single `Instance` and a `Collection` of instances.

```go
import "github.com/alimtvnetwork/core-v8/namevalue"

inst := namevalue.NewInstance("env", "production")
inst.Name()       // "env"
inst.ValueAny()   // "production"

col := namevalue.NewCollection()
col.Add(inst)
col.Add(namevalue.NewInstance("region", "us-east-1"))
col.ToMap()       // map[string]any{"env": "production", "region": "us-east-1"}
```

### When to use

- Configuration overrides where insertion order matters (a `map` would lose order).
- Building structured log lines.
- Anywhere you'd be tempted to use `[]struct{Name, Value string}`.

---

## 8. `keymk` — Key Compilation

Template-based key builder with named placeholders. Use for cache keys, audit-log keys, or any structured identifier.

```go
import "github.com/alimtvnetwork/core-v8/keymk"

// Template: "user/{userID}/post/{postID}"
key := keymk.New.Compile("user", userID, "post", postID)
// → "user/usr-123/post/pst-456"
```

> Avoids `fmt.Sprintf("user/%s/post/%s", ...)` drift when many call sites build the same key shape.

---

## 9. Decision Matrix

| You need to… | Use |
|---|---|
| Pick between two values based on a bool | `conditional.If*` |
| Lazy-evaluate one of two branches | `conditional.IfFunc*` |
| Default if a pointer is nil | `conditional.NilDef*` |
| Check `nil` on an `any` value | `isany.Null` (NEVER `== nil`) |
| Tri-state or 6-state boolean | `issetter.Value` |
| Compile a regex once, lazily | `regexnew.New.Lazy` (package-level) or `LazyLock` (method-local) |
| Min/Max on typed numerics | `coremath` (type-specific funcs) |
| Three-way compare | `corecmp.<Type>` |
| Sort a slice in-place | `coresort.<type>sort.Quick` |
| Reflect on a function value | `corefuncs.GetFuncName` |
| Insertion-ordered name→value | `namevalue.Collection` |
| Build a templated key string | `keymk.New.Compile` |

---

## See Also

- [`02-design-philosophy.md`](./02-design-philosophy.md) — `newCreator` pattern (§5), one-file-per-function (§1)
- [`04-error-system.md`](./04-error-system.md) — `conditional.ErrorFunc` plays well with `errcore.MergeErrors`
- [`06-data-structures.md`](./06-data-structures.md) — `coreonce` is the cousin of `regexnew.LazyLock` for non-regex values
- [`08-validators.md`](./08-validators.md) — Validators consume `regexnew.Lazy` patterns for string-format checks
- [`/spec/00-llm-integration-guide.md` §Utility Packages](../00-llm-integration-guide.md#utility-packages) — Quick reference
