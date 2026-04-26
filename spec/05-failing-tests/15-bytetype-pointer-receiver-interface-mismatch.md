# Failing Test: bytetypetests Build Failure + Test_Variant_OnlySupportedErr

## Test Location
`tests/integratedtests/bytetypetests/Extended2_test.go`

## Symptoms

### Build Failure
```
cannot use otherVariant (variable of byte type bytetype.Variant) as enuminf.BasicEnumer value:
bytetype.Variant does not implement enuminf.BasicEnumer (method UnmarshalJSON has pointer receiver)
```
Affected tests: `Test_Variant_IsEnumEqual`, `Test_Variant_IsAnyEnumsEqual`

### Test_Variant_OnlySupportedErr
`all names supported should not error` — but got an error.

## Root Cause

### Build Failure
`bytetype.Variant` has `UnmarshalJSON` on a **pointer receiver** (`*Variant`), so the value type `Variant` does not satisfy `enuminf.BasicEnumer`. The test was passing value types directly to `IsEnumEqual(enuminf.BasicEnumer)` and `IsAnyEnumsEqual(...enuminf.BasicEnumer)`.

### OnlySupportedErr
The test called `v.AllNameValues()` which returns `"Name(Value)"` formatted strings (e.g., `"Zero(0)"`), but `OnlySupportedErr` internally compares against `StringRanges()` which returns plain names (e.g., `"Zero"`). The formats never match, so all names appear unsupported.

## Fix Applied

### Build Failure
Changed value arguments to pointer arguments:
- `v.IsEnumEqual(otherVariant)` → `v.IsEnumEqual(&otherVariant)`
- `v.IsAnyEnumsEqual(bytetype.Two, ...)` → assign to variables, pass `&two, ...`

### OnlySupportedErr
Changed `v.AllNameValues()` → `v.StringRanges()` to match the format expected by `OnlySupportedErr`.

## Fix Location
Test logic fix in `tests/integratedtests/bytetypetests/Extended2_test.go`
