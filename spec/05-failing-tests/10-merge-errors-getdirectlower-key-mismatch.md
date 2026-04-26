# 05 — MergeErrors GetDirectLower Key Mismatch

## Failing Test
- `Test_ErrCore_MergeErrors_Verification` (Case 1)

## Root Cause

The test used `input.GetDirectLower("hasError")` to retrieve a boolean from the test case map. `GetDirectLower` converts the lookup key to lowercase:

```go
func (it Map) GetDirectLower(name string) any {
    x, has := it[strings.ToLower(name)]  // looks for "haserror"
```

But the map key is camelCase `"hasError"`. The lookup for `"haserror"` found nothing, returning `nil`. The condition `if hasError == true` evaluated to `false`, so `primaryErr` was never set, `MergeErrors(nil, nil)` returned `nil`, and the test produced `"true"` instead of expected `"false"`.

## Solution

Changed to typed accessor with a reusable `params` struct:

```go
hasError, _ := input.GetAsBool(params.hasError)
```

A local `params.go` was added to `errcoretests/` to hold reusable key constants.

## File Changed
- `tests/integratedtests/errcoretests/ErrType_test.go` (line 36)

## Learnings
- `GetDirectLower` is designed for case-insensitive lookups but requires **all map keys to also be lowercase**.
- CamelCase keys + `GetDirectLower` = silent `nil` return, no compile error.

## What Not to Repeat
- Don't use `GetDirectLower` with camelCase map keys. Use direct map access or `GetAsBool` for typed retrieval.
