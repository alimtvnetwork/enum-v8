# 03 — Args Reference (`coretests/args/`)

The `args` package provides typed input holders for test cases. All types are generic with `any`-based aliases for backward compatibility.

---

## args.Map

The **primary** input/output holder. A `map[string]any` with typed accessors.

### Declaration

```go
type Map map[string]any
```

### Creating

```go
// In ArrangeInput (test input):
ArrangeInput: args.Map{
    "name":    "Alice",
    "age":     30,
    "isAdmin": true,
}

// In ExpectedInput (test expectations):
ExpectedInput: args.Map{
    "isValid":  true,
    "fullName": "Alice Smith",
    "count":    1,
}

// In test runner (actual results):
actual := args.Map{
    "isValid":  result.IsValid(),
    "fullName": result.FullName(),
    "count":    result.Count(),
}
```

### Typed Getters (return `(T, bool)`)

| Method | Returns | Example |
|--------|---------|---------|
| `GetAsString("key")` | `(string, bool)` | `msg, ok := input.GetAsString("message")` |
| `GetAsInt("key")` | `(int, bool)` | `n, ok := input.GetAsInt("count")` |
| `GetAsBool("key")` | `(bool, bool)` | `flag, ok := input.GetAsBool("isEnabled")` |
| `GetAsStrings("key")` | `([]string, bool)` | `items, ok := input.GetAsStrings("tags")` |
| `GetAsAnyItems("key")` | `([]any, bool)` | `items, ok := input.GetAsAnyItems("data")` |
| `Get("key")` | `(any, bool)` | `val, ok := input.Get("anything")` |

### Default Getters

| Method | Returns | Example |
|--------|---------|---------|
| `GetAsStringDefault("key")` | `string` | Returns `""` if missing |
| `GetAsIntDefault("key", 0)` | `int` | Returns default if missing |
| `GetAsBoolDefault("key", false)` | `bool` | Returns default if missing |

### Presence Checks

| Method | Meaning |
|--------|---------|
| `Has("key")` | Key exists (value may be nil) |
| `HasDefined("key")` | Key exists AND value is non-nil |
| `HasDefinedAll("a", "b")` | All keys exist and are non-nil |
| `IsKeyMissing("key")` | Key does not exist |
| `IsKeyInvalid("key")` | Key missing OR value is nil |

### Positional Access (aliases for common keys)

| Method | Checks keys (in order) |
|--------|----------------------|
| `FirstItem()` | `"first"`, `"f1"`, `"p1"`, `"1"` |
| `SecondItem()` | `"second"`, `"f2"`, `"p2"`, `"2"` |
| `ThirdItem()` | `"third"`, `"f3"`, `"p3"`, `"3"` |
| `WorkFunc()` | `"func"`, `"work.func"`, `"workFunc"` |
| `Expected()` | `"expected"`, `"expects"`, `"expect"` |

### Function Invocation

```go
// Direct invocation via reflection:
results, err := input.Invoke(arg1, arg2)

// Invoke with all defined (non-func) values:
results, err := input.InvokeWithValidArgs()

// Invoke with specific named args:
results, err := input.InvokeArgs("first", "second")
```

### CompileToStrings (for assertions)

Converts map to sorted `"key : value"` lines using `%v` formatting:

```go
m := args.Map{"isZero": false, "value": 5}
m.CompileToStrings()
// Returns: []string{"isZero : false", "value : 5"}
```

This is what `ShouldBeEqualMap` uses internally to compare expected vs actual.

### Formatting Rules

```go
// ✅ GOOD — multi-line for 2+ entries:
ArrangeInput: args.Map{
    "name": "Alice",
    "age":  30,
}

// ✅ OK — inline for single entry:
ExpectedInput: args.Map{"isValid": true}

// ❌ BAD — inline for 2+ entries:
ExpectedInput: args.Map{"isValid": true, "count": 5}
```

---

## args.One[T1] through args.Six[T1...T6]

Positional argument holders for 1–6 typed parameters.

### Structure (args.One example)

```go
type One[T1 any] struct {
    First  T1
    Expect any
}
```

### Aliases

```go
type OneAny   = One[any]
type TwoAny   = Two[any, any]
type ThreeAny = Three[any, any, any]
// ... through SixAny
```

### Key Methods

| Method | Exists on | Purpose |
|--------|-----------|---------|
| `FirstItem()` | All | Returns `First` as `any` |
| `SecondItem()` | Two+ | Returns `Second` as `any` |
| `HasFirst()` | All | Non-nil check |
| `ValidArgs()` | All | Returns all defined args as `[]any` |
| `Args(upTo)` | All | Returns first N args |
| `Expected()` | All | Returns `Expect` field |

### When to Use

Use in `ArrangeInput` when test data is fundamentally positional (not key-value):

```go
// ✅ GOOD — positional data in ArrangeInput:
ArrangeInput: &args.TwoAny{
    First:  "hello",
    Second: 42,
    Expect: "hello-42",
}

// ❌ BAD — positional data in ExpectedInput (use args.Map instead):
ExpectedInput: &args.OneAny{First: "result"}
```

**Rule**: Positional types are **allowed in ArrangeInput** but **prohibited in ExpectedInput**. Always use `args.Map` for expectations.

---

## args.Dynamic[T]

A map-based argument holder with a typed `Expect` field. Combines the flexibility of `args.Map` with typed expectations.

### Structure

```go
type Dynamic[T any] struct {
    Params Map
    Expect T
}
```

### Alias

```go
type DynamicAny = Dynamic[any]
```

### When to Use

When the input is inherently dynamic (variable number of parameters) but the expected result has a known type:

```go
ArrangeInput: &args.Dynamic[bool]{
    Params: args.Map{
        "input": "test@example.com",
        "func":  validateEmail,
    },
    Expect: true,
}
```

### Key Methods

All `args.Map` methods are delegated through `Params`:
- `Get`, `GetAsString`, `GetAsInt`, `HasDefined`, etc.
- `Invoke`, `InvokeWithValidArgs` (function invocation)

---

## args.Holder[T]

A 6-slot positional holder with a typed `WorkFunc` field and a fallback `Hashmap`.

### Structure

```go
type Holder[T any] struct {
    First    any
    Second   any
    Third    any
    Fourth   any
    Fifth    any
    Sixth    any
    WorkFunc T
    Expect   any
    Hashmap  Map
}
```

### Alias

```go
type HolderAny = Holder[any]
```

### When to Use

When you need both positional parameters AND a function reference in the same test case:

```go
ArrangeInput: &args.HolderAny{
    First:    "input-data",
    Second:   42,
    WorkFunc: myProcessFunc,
    Hashmap:  args.Map{"extra": "config"},
}
```

---

## args.LeftRight[TLeft, TRight]

Semantic two-item holder for cases where left/right directionality matters.

### Structure

```go
type LeftRight[TLeft, TRight any] struct {
    Left   TLeft
    Right  TRight
    Expect any
}
```

### Alias

```go
type LeftRightAny = LeftRight[any, any]
```

### When to Use

Comparison tests, diff tests, migration tests — anywhere "left vs right" is meaningful:

```go
ArrangeInput: &args.LeftRightAny{
    Left:   originalConfig,
    Right:  modifiedConfig,
    Expect: "3 differences",
}
```

---

## args.ThreeFunc / args.ThreeFuncAny

Variant of `Three` that includes a `WorkFunc` field for function invocation tests:

```go
type ThreeFunc[T1, T2, T3 any] struct {
    First    T1
    Second   T2
    Third    T3
    WorkFunc any
    Expect   any
}
```

### Example

```go
ArrangeInput: args.ThreeFuncAny{
    First:    "arg1",
    Second:   "arg2",
    Third:    "arg3",
    WorkFunc: myFunction,
}
```

---

## Decision Matrix: Which Args Type?

| Scenario | Use |
|----------|-----|
| Key-value input (most common) | `args.Map` |
| Positional function arguments (2-6 params) | `args.Two` – `args.Six` |
| Variable number of params + typed expect | `args.Dynamic[T]` |
| Positional params + function reference | `args.Holder[T]` |
| Left/right comparison | `args.LeftRight[T1, T2]` |
| Function params + function ref (3 args) | `args.ThreeFuncAny` |
| Expected output (always) | `args.Map` ✅ |

---

## args.Map as a Self-Contained Assertion Target (Style C)

`args.Map` is not just an input/output container — it is also a **standalone assertion target**. This enables **Style C tests**: one-off micro-assertions where the test function is the case (no shared `_testcases.go` slice, no loop).

### When to Use Style C

✅ Use Style C when:
- The test has < 5 lines of arrange.
- Adding a case to a shared slice would be more boilerplate than the test itself.
- The function under test is small and exercised by exactly one scenario.

❌ Do **not** use Style C when:
- You have ≥ 2 scenarios for the same function — promote to Style A (`CaseV1` slice).
- The expected output is a `[]string` line list — use Style B with `ShouldBeEqual`.
- You need `VerifyTypeOf` or shared parameters — use Style A.

### Pattern

The **expected** `args.Map` is the receiver; the **actual** `args.Map` is the third positional argument. The second argument is a string title that replaces the `tc.Title` field of Style A.

```go
expected.ShouldBeEqual(t, caseIndex, title, actual)
```

### Complete Example

```go
// Map_Length_test.go
package argstests

import (
    "testing"

    "github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_Map_Length_FromMapLength(t *testing.T) {
    // Arrange
    m := args.Map{"a": 1, "b": 2}

    // Act
    actual := args.Map{"length": m.Length()}

    // Assert
    expected := args.Map{"length": 2}
    expected.ShouldBeEqual(t, 0, "Map.Length returns 2 -- two entries", actual)
}
```

### Style C Method Signatures on args.Map

| Method | Signature | Use When |
|--------|-----------|----------|
| `ShouldBeEqual(t, caseIndex, title, actual)` | `(t, int, string, Map)` | Standalone micro-assertion, expected is the receiver |
| `CompileToStrings()` | `() []string` | Manually convert to sorted `"key : value"` lines for custom assertions |

### Style C vs Style A Side-by-Side

❌ **Overkill — Style A for a single trivial case:**

```go
var mapLengthTestCases = []coretestcases.CaseV1{
    {
        Title:         "Map.Length returns 2",
        ArrangeInput:  args.Map{"a": 1, "b": 2},
        ExpectedInput: args.Map{"length": 2},
    },
}

func Test_Map_Length(t *testing.T) {
    for caseIndex, tc := range mapLengthTestCases {
        input := tc.ArrangeInput.(args.Map)
        actual := args.Map{"length": input.Length()}
        tc.ShouldBeEqualMap(t, caseIndex, actual)
    }
}
```

✅ **Right-sized — Style C:**

```go
func Test_Map_Length(t *testing.T) {
    m := args.Map{"a": 1, "b": 2}
    actual := args.Map{"length": m.Length()}
    args.Map{"length": 2}.ShouldBeEqual(t, 0, "Map.Length returns 2 -- two entries", actual)
}
```

> The full Style C narrative lives in [`/spec/01-app/13-testing-patterns.md` §4](../01-app/13-testing-patterns.md#4-style-c--standalone-argsmapshouldbeequal-micro-tests).

---

## Gotcha: Empty Expected Output — Use `args.Map{}`, Never `""`

This is the single most-stepped-on landmine in `args.Map` assertions. It is **fixed** in `spec/05-failing-tests/02-groupby-empty-map-assertion-mismatch.md` but worth restating here because the symptom is misleading.

### The Bug

For test cases that produce **zero output lines**, you might be tempted to write:

```go
// ❌ BAD — produces a confusing length mismatch
{
    Title:         "GroupBy returns empty map -- empty input",
    ArrangeInput:  args.Map{"items": []string{}},
    ExpectedInput: "",
}
```

When the runner builds `actLines = []string{}` (empty slice) and calls `tc.ShouldBeEqual(t, idx, actLines...)`, the variadic expands to **zero arguments**. But `ExpectedInput: ""` is normalized to **one expected line** (the empty string `""`). The assertion fails with:

```
Expect [""] != [] Actual
```

### The Fix

Use `args.Map{}` for "truly empty" expectations:

```go
// ✅ GOOD — both sides are empty maps, length matches
{
    Title:         "GroupBy returns empty map -- empty input",
    ArrangeInput:  args.Map{"items": []string{}},
    ExpectedInput: args.Map{},
}
```

This routes through the `if _, isMap := ExpectedInput.(args.Map); isMap` branch, which builds an empty `args.Map{}` actual and compares it against the empty `args.Map{}` expected — both are length 0, both pass.

### Decision Rule

| Expected output | Correct `ExpectedInput` |
|---|---|
| Zero lines / nothing produced | `args.Map{}` ✅ |
| One empty string line `""` (rare, intentional) | `[]string{""}` |
| One line `"0"` (e.g. count assertions) | `"0"` or `args.Map{"count": 0}` |
| One or more semantic lines | `args.Map{...}` (preferred) or `[]string{...}` |

### Why This Happens

`CaseV1.ExpectedLines()` normalizes:
- `string` → `[]string{s}` (always 1 line, even if `s == ""`)
- `args.Map` → `CompileToStrings()` (zero lines if map is empty)

The map path is the **only** way to express "expect zero lines". The string path always expects ≥ 1 line.

### Related Failing Test

See [`/spec/05-failing-tests/02-groupby-empty-map-assertion-mismatch.md`](../05-failing-tests/02-groupby-empty-map-assertion-mismatch.md) for the original bug report and the exact assertion-framework code paths involved.

---

## `params.go` Convention

A `params.go` file at the root of a test package centralises `args.Map` key constants used across multiple test cases, eliminating magic-string duplication.

### Rule

| Package size | Rule |
|---|---|
| **New package, > 3 test cases** sharing args keys | **Mandatory** — create `params.go` |
| **New package, ≤ 3 test cases** | **Optional** — inline string literals are equally readable |
| **Existing package** without `params.go` | **Grandfathered** — no back-fill required |

### Example

```go
// errcoretests/params.go
package errcoretests

const (
    keyInput    = "input"
    keyExpected = "expected"
    keyWantErr  = "wantErr"
)
```

Then in test files:

```go
ArrangeInput: args.Map{
    keyInput:    "hello",
    keyExpected: "HELLO",
    keyWantErr:  false,
}
```

### Rationale

- Centralising keys pays off only when reused across **3+ cases**; below that, inline literals are clearer and reduce indirection.
- A whole-tree back-fill across `tests/integratedtests/` would generate noise without behaviour change.
- This rule resolves [`/spec/02-app-issues/05-missing-params-go-files.md`](../02-app-issues/05-missing-params-go-files.md).
