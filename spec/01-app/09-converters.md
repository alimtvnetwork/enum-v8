# 09 — Converters

> ✅ **Status**: filled in audit Step 5c (2026-04-23, Asia/Kuala_Lumpur).
> **Audience**: anyone converting between primitive types or normalising data at API boundaries.
> **Source**: extracted from [`/spec/00-llm-integration-guide.md`](../00-llm-integration-guide.md) §converters.

---

## Table of Contents

1. [`converters` Package](#1-converters-package)
2. [`typesconv` Package](#2-typesconv-package)
3. [Conversion Safety Contract](#3-conversion-safety-contract)
4. [Examples](#4-examples)
5. [Common Mistakes](#5-common-mistakes)

---

## 1. `converters` Package

Located at `converters/`. Provides struct-as-namespace converters for the most common primitive conversions. Each conversion family hangs off a noun-typed namespace (`StringTo`, `BytesTo`, `PrettyJson`).

### 1.1 `StringTo`

```go
import "github.com/alimtvnetwork/core-v8/converters"

// String → integer (returns error)
val, err := converters.StringTo.Integer("42")
val, err := converters.StringTo.Integer64("9223372036854775807")

// String → integer with fallback (no error, returns ok flag)
val, ok := converters.StringTo.IntegerWithDefault("abc", -1) // val=-1, ok=false

// String → float
f, err := converters.StringTo.Float64("3.14")
f, err := converters.StringTo.Float32("3.14")

// String → byte
b, err := converters.StringTo.Byte("255")

// String → bool. Accepts the same inputs as stdlib `strconv.ParseBool`
// (case-sensitive set: "1", "t", "T", "TRUE", "true", "True",
//  "0", "f", "F", "FALSE", "false", "False"). Leading/trailing whitespace
// is REJECTED — call `strings.TrimSpace` upstream if you need lenient parsing.
// "yes"/"no"/"on"/"off" are NOT accepted. (F-V14-03 fix.)
ok, err := converters.StringTo.Bool("true")
```

### 1.2 `BytesTo`

```go
s := converters.BytesTo.String([]byte("hello"))     // direct cast
s := converters.BytesTo.PrettyJsonString(jsonBytes) // formatted
```

### 1.3 `PrettyJson`

```go
prettyStr := converters.PrettyJson.String(jsonBytes)
prettyStr := converters.PrettyJson.FromAny(myStruct)
```

> The `PrettyJson` namespace duplicates a subset of [`corejson`](./06-data-structures.md#4-corejson--json-pipeline). Prefer `corejson.NewPtr(x).PrettyJsonString()` in new code; `PrettyJson` exists for legacy call sites.

---

## 2. `typesconv` Package

Located at `typesconv/`. Lower-level companion to `converters` — used when you need to convert **between two non-string types** (e.g. `int` → `int64`, `float64` → `int`) with explicit overflow / precision-loss handling.

```go
import "github.com/alimtvnetwork/core-v8/typesconv"

// Numeric widening (always safe)
i64 := typesconv.IntToInt64(myInt)

// Numeric narrowing (may lose precision — returns ok flag)
i32, ok := typesconv.Int64ToInt32(largeValue)

// Float → int (truncates, returns ok flag for NaN/Inf)
i, ok := typesconv.Float64ToInt(3.7) // i=3, ok=true
i, ok := typesconv.Float64ToInt(math.NaN()) // ok=false
```

### When to use `typesconv` vs `converters`

| Scenario | Use |
|---|---|
| Source value is a string (parsing) | `converters.StringTo.*` |
| Source value is `[]byte` | `converters.BytesTo.*` |
| Source and target are both numeric | `typesconv.<Source>To<Target>` |
| Pretty-printing JSON | `corejson.NewPtr(x).PrettyJsonString()` |

---

## 3. Conversion Safety Contract

All converters in both packages follow this contract:

1. **Two return modes**:
   - `(value, error)` — for parsing failures with a recoverable cause (use when caller wants to log the malformed input).
   - `(value, bool)` — for "best-effort" conversions where the caller has a sensible default (use when failure is non-exceptional).
2. **No panics on bad input** — invalid input returns the zero value of the target type plus the failure signal.
3. **Errors are categorised via `errcore`** — see [`04-error-system.md`](./04-error-system.md). Specifically `errcore.FailedToConvertType` wraps stdlib parsing errors so the test runner can attribute them.
4. **Truncation is silent** when using `*WithDefault` variants. Callers who care about precision must use the `(value, error)` form and check the result.
5. **Locale-independent** — number parsers always use `.` as the decimal separator and never apply thousand grouping.

### Choosing between the two return modes

```go
// Recoverable failure — log and reject
val, err := converters.StringTo.Integer(userInput)
if err != nil {
    return errcore.FailedToConvertType.Fmt("invalid integer: %s", userInput)
}

// Non-exceptional failure — fall back to default
limit, _ := converters.StringTo.IntegerWithDefault(queryParam, 25)
```

---

## 4. Examples

### 4.1 Parsing a query string

```go
func parsePagination(r *http.Request) (page, perPage int) {
    page, _ = converters.StringTo.IntegerWithDefault(r.URL.Query().Get("page"), 1)
    perPage, _ = converters.StringTo.IntegerWithDefault(r.URL.Query().Get("per_page"), 20)
    return
}
```

### 4.2 Validating a numeric field with explicit error

```go
func parseAge(s string) (int, error) {
    age, err := converters.StringTo.Integer(s)
    if err != nil {
        return 0, errcore.FailedToConvertType.Fmt(
            "age %q is not an integer", s,
        )
    }
    if age < 0 || age > 200 {
        return 0, errcore.ValidationFailedType.Fmt(
            "age %d is out of range [0..200]", age,
        )
    }
    return age, nil
}
```

### 4.3 Narrowing a counter for a 32-bit downstream

```go
total64 := computeTotal() // returns int64
total32, ok := typesconv.Int64ToInt32(total64)
if !ok {
    return errcore.OverflowType.Fmt("total %d exceeds int32", total64)
}
```

---

## 5. Common Mistakes

| Mistake | Why bad | Fix |
|---|---|---|
| Using `strconv.Atoi` directly | Loses `errcore` categorisation, breaks test attribution | Use `converters.StringTo.Integer` |
| Using `*WithDefault` then re-validating | Hides the malformed input from logs | Use `(value, error)` form when you want to log |
| Mixing `converters.PrettyJson` and `corejson.PrettyJsonString` in the same package | Inconsistent style | Pick `corejson` in new code |
| Catching float→int overflow with `recover` | `typesconv` never panics — `recover` is dead code | Check the `ok` return flag |
| Calling `BytesTo.String` and re-allocating | The conversion is already zero-copy | Just use the returned string |

---

## See Also

- [`04-error-system.md`](./04-error-system.md) — `errcore.FailedToConvertType` and related error categories
- [`06-data-structures.md`](./06-data-structures.md) — `corejson` is the canonical JSON pipeline; prefer it over `converters.PrettyJson`
- [`08-validators.md`](./08-validators.md) — Validators commonly chain with converters at API boundaries
- [`/spec/00-llm-integration-guide.md` §converters](../00-llm-integration-guide.md#converters--type-conversions) — Quick reference
