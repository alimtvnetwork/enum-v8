# 04 — Results Reference (`coretests/results/`)

The `results` package provides typed return-value containers with panic recovery and structured assertion support.

---

## results.Result[T]

Typed container for a single function invocation result.

### Structure

```go
type Result[T any] struct {
    Value       T       // Primary return value
    Error       error   // Error returned by the function (or wrapped panic)
    Panicked    bool    // True if invocation recovered from a panic
    PanicValue  any     // Raw value recovered from panic (nil if no panic)
    AllResults  []any   // All return values as []any (for multi-return)
    ReturnCount int     // Number of return values from the function
}
```

### Common Aliases

```go
type ResultAny    = Result[any]      // Most common — used by CaseNilSafe
type ResultBool   = Result[bool]
type ResultString = Result[string]
type ResultInt    = Result[int]
type ResultError  = Result[error]
```

### Key Methods

| Method | Returns | Purpose |
|--------|---------|---------|
| `IsSafe()` | `bool` | `!Panicked && Error == nil` |
| `HasError()` | `bool` | `Error != nil` |
| `HasPanicked()` | `bool` | `Panicked == true` |
| `ValueString()` | `string` | `fmt.Sprintf("%v", Value)` |
| `IsResult(expected)` | `bool` | String comparison of Value vs expected |
| `ResultAt(index)` | `any` | Access multi-return value by index |
| `ToMap()` | `args.Map` | Full map with value, panicked, isSafe, hasError, returnCount |
| `ToMapCompact()` | `args.Map` | Minimal map with only value and panicked |

### ToMap Output

```go
result.ToMap()
// Returns:
// args.Map{
//     "value":       "false",
//     "panicked":    false,
//     "isSafe":      true,
//     "hasError":    false,
//     "returnCount": 1,
// }
```

---

## results.Results[T1, T2]

Two-value typed result for functions returning `(T1, T2)`.

### Structure

```go
type Results[T1, T2 any] struct {
    Result[T1]           // Embeds primary result
    Result2 T2           // Second return value
}
```

### Common Aliases

```go
type ResultsAny         = Results[any, any]
type ResultsAnyError    = Results[any, error]
type ResultsStringError = Results[string, error]
type ResultsBoolError   = Results[bool, error]
```

### Converting from ResultAny

```go
// Convert untyped ResultAny to typed Results:
raw := results.InvokeWithPanicRecovery(funcRef, receiver)
typed := results.FromResultAny[string, error](raw)

fmt.Println(typed.Value)    // string
fmt.Println(typed.Result2)  // error
```

---

## results.ExpectAnyError

Sentinel value used in `Expected` to indicate "expect any non-nil error":

```go
var ExpectAnyError = fmt.Errorf("expect-any-error")
```

### Usage in CaseNilSafe

```go
{
    Title: "Deserialize on nil returns error",
    Func: func(d *MyStruct) error {
        return d.ValueNullErr()
    },
    Expected: results.ResultAny{
        Error:    results.ExpectAnyError,  // Any non-nil error passes
        Panicked: false,
    },
}
```

---

## InvokeWithPanicRecovery

The core invocation engine used by `CaseNilSafe`. Calls a function reference with a receiver and optional args, recovering from any panic.

### Signature

```go
func InvokeWithPanicRecovery(
    funcRef any,      // Method expression: (*Type).Method or func literal
    receiver any,     // The receiver value (nil for nil-safety tests)
    args ...any,      // Additional arguments
) ResultAny
```

### How It Works

1. Validates `funcRef` is not nil and is a function
2. Builds call args with typed nil receiver for nil-safety tests
3. Calls the function via `reflect.Value.Call`
4. Captures all return values into `ResultAny.AllResults`
5. Extracts the primary value and last-error if applicable
6. Recovers from panics, setting `Panicked = true` and `PanicValue`

### Edge Cases Handled

| Scenario | Result |
|----------|--------|
| `funcRef == nil` | `Panicked: true, PanicValue: "funcRef is nil"` |
| `funcRef` is not a function | `Panicked: true, PanicValue: "funcRef is not a function: int"` |
| Void method (no returns) | `ReturnCount: 0, Value: nil` |
| Multi-return `(string, error)` | `Value: first, Error: last if implements error` |
| Nil interface return | Properly detected via `IsNil()` |
| Value receiver with nil pointer | Natural Go panic recovered |

### Direct Usage (rare)

Usually you use `CaseNilSafe.InvokeNil()` which delegates to this. Direct usage:

```go
result := results.InvokeWithPanicRecovery(
    (*MyStruct).IsValid,  // method expression
    nil,                   // nil receiver
)

fmt.Println(result.Value)    // "false"
fmt.Println(result.Panicked) // false
```

---

## ShouldMatchResult

The assertion engine for `Result[T]`. Compares actual vs expected using auto-derived or explicit field selection.

### Auto-Derived Fields

| Condition | Field Compared |
|-----------|---------------|
| Always | `panicked` |
| `Expected.Value != nil` | `value` |
| `Expected.Error != nil` | `hasError` |
| `Expected.ReturnCount != 0` | `returnCount` |

### How It Works

1. Both actual and expected `Result` are converted to `args.Map` via `ToMap()`
2. Maps are filtered to only the compared fields
3. Filtered maps are compiled to sorted string lines
4. Lines are compared via `errcore.AssertDiffOnMismatch`

### Override Example

```go
result.ShouldMatchResult(
    t,
    caseIndex,
    "MyMethod on nil",
    expected,
    "panicked", "returnCount",  // Only compare these two fields
)
```

---

## MethodName

Extracts the method name from a function reference:

```go
results.MethodName((*MyStruct).IsValid) // → "IsValid"
results.MethodName((*MyStruct).String)  // → "String"
results.MethodName(nil)                  // → ""
```

Used internally by `CaseNilSafe.MethodName()` for auto-generating test titles.
