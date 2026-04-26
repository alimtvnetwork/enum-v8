# Failing Test: Test_Variant_OnlySupportedErr (stringcompareas)

## Test Location
`tests/integratedtests/stringcompareastests/Extended_test.go`

## Symptoms
`all names supported should not error` — but got an error.

## Root Cause
The test originally used `v.AllNameValues()` which returns formatted strings like
`"Equal(0)"`, but `OnlySupportedErr` internally compares against `StringRanges()`
which returns plain names like `"Equal"`. The formats never match, so all names
appear unsupported.

A subsequent fix attempted using `BasicEnumImpl.StringRanges()` directly, which
should theoretically return the exact same slice. However, to eliminate any
subtle reference/copy issues and match the proven pattern from `issettertests`,
the fix now uses explicit string literals matching the enum's `stringRanges` array.

## Fix Applied
Changed to use explicit name strings matching the enum's `vars.go` `stringRanges`
array values:
```go
allNames := []string{
    "Equal", "StartsWith", "EndsWith", "Anywhere",
    "IsContains", "AnyChars", "Regex",
    "NotEqual", "NotStartsWith", "NotEndsWith",
    "NotContains", "NotAnyChars", "NotMatchRegex",
    "Glob", "NonGlob", "Invalid",
}
```

## Fix Location
Test logic fix in `tests/integratedtests/stringcompareastests/Extended_test.go`
