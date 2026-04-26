# 07 — NilLazyRegex OnRequiredCompiled Single-Return Error

## Failing Test
- `Test_NilLazyRegex` (Case 8: OnRequiredCompiled on nil returns error)

## Root Cause

`OnRequiredCompiled()` returns a single `error`. The `CaseNilSafe` framework invokes methods via reflection and extracts results using `extractResult()`.

The original `extractResult` only populated `result.Error` from **multi-return** functions (the last return value). For **single-return error** functions, the error value was placed in `result.Value` only, leaving `result.Error = nil`.

The test case expected `Error: results.ExpectAnyError`, but `ShouldMatchResult` compared `result.Error != nil` (false) against `expected.Error != nil` (true), causing the mismatch:

```
actual : `hasError : false`
expected : `hasError : true`
```

## Solution

Updated `extractResult()` in `coretests/results/Invoke.go` to also check single-return values for error type:

```go
if len(returnValues) == 1 {
    result.Error = extractErrorFromValue(returnValues[0])
} else {
    last := returnValues[len(returnValues)-1]
    result.Error = extractErrorFromValue(last)
}
```

This correctly populates both `Value` and `Error` for single-return error methods, maintaining backward compatibility since `extractErrorFromValue` only extracts values implementing the `error` interface.

## File Changed
- `coretests/results/Invoke.go` (extractResult function)

## Learnings
- Single-return error methods are a valid pattern (e.g., `OnRequiredCompiled() error`). The reflection-based invocation framework must handle them.
- The `Error` field should always be populated when any return value is an error, regardless of return count.

## What Not to Repeat
- When building reflection-based test frameworks, handle all return-count scenarios (0, 1, N) explicitly. Don't assume error is always the "second" return.
