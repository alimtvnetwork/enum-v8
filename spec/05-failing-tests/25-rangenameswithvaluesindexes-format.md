# Failing Test: RangeNamesWithValuesIndexes format mismatch

## Root Cause

`constants.EnumNameValueFormat` is `"%s(%d)"` (parentheses), but test expectations used
bracket format `name[index]`. The production format constant is authoritative.

## Fix

Updated test case data in `corecsvtests/testCases.go` to use `()` format matching the
current `EnumNameValueFormat` constant.
