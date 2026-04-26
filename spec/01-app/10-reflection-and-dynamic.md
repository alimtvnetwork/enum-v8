# 10 — Reflection & Dynamic

> ✅ **Status**: filled in audit Step 5c (2026-04-23, Asia/Kuala_Lumpur).
> **Audience**: anyone working on dynamic dispatch, JSON-shape inspection, or test infrastructure that needs reflection.
> **Source**: extracted from [`/spec/00-llm-integration-guide.md`](../00-llm-integration-guide.md), the `internal/reflectinternal` package, and the failing-test post-mortems in [`/spec/05-failing-tests/`](../05-failing-tests/).

---

## Table of Contents

1. [Why a Dedicated Reflection Layer?](#1-why-a-dedicated-reflection-layer)
2. [`coredynamic` API](#2-coredynamic-api)
3. [`reflectcore` Helpers](#3-reflectcore-helpers)
4. [`internal/reflectinternal`](#4-internalreflectinternal)
5. [When to Prefer Reflection vs Generics](#5-when-to-prefer-reflection-vs-generics)
6. [Performance Notes](#6-performance-notes)
7. [Common Mistakes](#7-common-mistakes)

---

## 1. Why a Dedicated Reflection Layer?

Go's stdlib `reflect` package is correct but verbose, panics liberally, and loses type information at compile time. The framework wraps reflection in three layers so most callers never see `reflect.Value`:

| Layer | Visibility | Purpose |
|---|---|---|
| `coredynamic` | public | Dynamic method invocation, `any`-typed property access |
| `reflectcore` | public | Type predicates, kind checks, struct walking |
| `internal/reflectinternal` | internal-only | Low-level `reflect.Value` mutation, used by `enumimpl` and JSON pipeline |

> **Rule (MUST)**: New consumer code **MUST** reach for `coredynamic` or `reflectcore` first and **MUST NOT** import `internal/reflectinternal` directly — that package is reserved for framework internals (see [`03-import-conventions.md`](./03-import-conventions.md)). Direct use of the `reflect` standard-library package in consumer code is also forbidden; route every reflection call through `coredynamic` / `reflectcore` so panics are wrapped into `errcore` errors.
>
> **Convention vs Hard Rule (project-wide)**: Throughout this spec, **MUST / MUST NOT / MAY** indicate non-negotiable rules — an AI agent that violates them produces incorrect code. **"should" / "rule of thumb" / "prefer"** indicate guidance — overriding them requires a documented reason but is not a defect. When in doubt, treat ambiguous wording as **MUST** for safety/correctness rules and as guidance for stylistic preferences.

---

## 2. `coredynamic` API

Located at `coredynamic/`. Provides safe dynamic dispatch over arbitrary `any` values.

### 2.1 Dynamic method invocation

```go
import "github.com/alimtvnetwork/core-v8/coredynamic"

// Call a method by name with positional args.
// Signature: InvokeMethod(target any, name string, args ...any) (any, error)
// There is NO single-return overload — both values MUST be received.
// (F-V14-02 fix.)
result, err := coredynamic.InvokeMethod(target, "Validate", arg1, arg2)

// Check method existence first; still receive both returns even when discarding.
if coredynamic.HasMethod(target, "Reset") {
    _, _ = coredynamic.InvokeMethod(target, "Reset")
}

// List all methods on a value
names := coredynamic.MethodNames(target) // []string
```

### 2.2 Property access

```go
v, ok := coredynamic.GetField(target, "Name")
err   := coredynamic.SetField(target, "Name", "new value")
m     := coredynamic.AllFields(target) // map[string]any
```

### 2.3 Diagnostic / introspection

```go
typeName := coredynamic.TypeName(target)         // e.g. "MyStruct"
fullName := coredynamic.TypeFullName(target)     // e.g. "github.com/.../pkg.MyStruct"
isNil    := coredynamic.IsNullOrUndefined(target) // typed-nil safe
```

> See [Failing Test #14 — Dynamic methods cases](../05-failing-tests/) for the canonical post-mortem on why `coredynamic.InvokeMethod` returns `(any, error)` rather than a slice of `reflect.Value`.

---

## 3. `reflectcore` Helpers

Located at `reflectcore/`. Stateless predicates and structural helpers — never mutates the input.

### 3.1 Kind predicates

```go
import "github.com/alimtvnetwork/core-v8/reflectcore"

reflectcore.IsPointer(val)
reflectcore.IsStruct(val)
reflectcore.IsSlice(val)
reflectcore.IsMap(val)
reflectcore.IsFunc(val)
reflectcore.IsChannel(val)
reflectcore.IsInterface(val)
```

### 3.2 Struct walking

```go
// Walk every exported field
reflectcore.WalkFields(target, func(name string, value any) {
    fmt.Printf("%s = %v\n", name, value)
})

// Get tag values (e.g. JSON tags)
tag := reflectcore.GetTag(target, "Name", "json") // e.g. "name,omitempty"
```

### 3.3 Pointer resolution

```go
// Dereference one or more pointer levels safely
underlying := reflectcore.DerefAll(ptrToPtrToStruct)
```

---

## 4. `internal/reflectinternal`

Off-limits to consumers. Documented here only for framework maintainers.

### Responsibilities

1. **Low-level `reflect.Value` setters** used by `enumimpl` to populate enum vars-table entries.
2. **Unsafe pointer arithmetic** for the `corejson` fast-path (avoids one allocation per Marshal).
3. **Type-cache** keyed on `reflect.Type` to amortise reflection cost across calls.

### Why internal

- Misuse can corrupt enum singletons (every enum value must be a single addressable instance).
- The `unsafe.Pointer` paths assume `GOARCH=amd64` / `arm64` alignment.
- API is unstable and tied to Go runtime version.

> See [`03-import-conventions.md` §2](./03-import-conventions.md) for the `internal/` boundary rules.

---

## 5. When to Prefer Reflection vs Generics

Decision matrix:

| Scenario | Use |
|---|---|
| Type is known at compile time | **Generics** ([`coregeneric`](./06-data-structures.md#2-coregeneric--generic-containers)) |
| Type is `any` and you need a method call | `coredynamic.InvokeMethod` |
| Type is `any` and you only need to check shape | `reflectcore.Is*` predicates |
| Need to walk arbitrary structs (logging, diff) | `reflectcore.WalkFields` |
| JSON serialization | [`corejson`](./06-data-structures.md#4-corejson--json-pipeline) (handles reflection internally) |
| Inside framework internals, need raw `reflect.Value` | `internal/reflectinternal` (maintainer only) |
| Test framework needs to attribute failures | `corefuncs.GetFuncName` (see [`07-conditional-and-utilities.md` §6](./07-conditional-and-utilities.md#6-corefuncs--function-wrappers)) |

> **Rule of thumb**: if a generic version exists, use it. Reflection should appear only when the compiler genuinely cannot know the type — typically at API boundaries, plugin systems, or test infrastructure.

---

## 6. Performance Notes

Reflection is ~10–100× slower than direct calls. The framework mitigates this in three ways:

1. **Type caching** in `internal/reflectinternal` — `reflect.Type → method-table` lookups are memoized.
2. **Lazy binding** via [`coreonce`](./06-data-structures.md#5-coreonce--compute-once-values) — expensive reflection results are computed once, reused everywhere.
3. **Generics-first defaults** — `coregeneric` containers use generics; reflection is opt-in.

### When to worry

- Hot paths processing >10k calls/sec
- Tight loops over large slices that touch reflection per element

### When not to worry

- Init-time wiring (config parsing, enum registration)
- Test infrastructure (already off the hot path)
- One-shot CLI commands

---

## 7. Common Mistakes

| Mistake | Why bad | Fix |
|---|---|---|
| Importing `reflect` directly in consumer code | Bypasses error categorisation and panic safety | Use `coredynamic` or `reflectcore` |
| Importing `internal/reflectinternal` from outside the module | Compile error (`internal/` rule) | Use the public layer |
| Using reflection where generics would compile | Slow + harder to read | Switch to `coregeneric` |
| Calling `InvokeMethod` without checking `HasMethod` first | Returns an error per call — wasteful in hot paths | Check existence once, cache the result |
| Forgetting `DerefAll` on a `**MyStruct` | Predicates report `IsPointer=true` not `IsStruct=true` | Always `DerefAll` first when you want the underlying kind |
| Using `reflect.DeepEqual` directly | Doesn't handle typed-nil quirks | Use `isany.DeepEqual` (see [`07-conditional-and-utilities.md` §2](./07-conditional-and-utilities.md#2-isany--type-predicates)) |

---

## See Also

- [`02-design-philosophy.md`](./02-design-philosophy.md) — Why the framework wraps stdlib instead of using it directly
- [`03-import-conventions.md`](./03-import-conventions.md) — `internal/` boundary
- [`05-enum-system.md`](./05-enum-system.md) — `enumimpl` is the heaviest internal consumer of `reflectinternal`
- [`06-data-structures.md`](./06-data-structures.md) — Generics-first containers
- [`07-conditional-and-utilities.md`](./07-conditional-and-utilities.md) — `isany`, `corefuncs` use reflection internally
- [`/spec/05-failing-tests/`](../05-failing-tests/) — Real-world reflection bugs and resolutions
