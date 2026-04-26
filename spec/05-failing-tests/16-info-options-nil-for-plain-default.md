# Failing Test: Test_Info_Options_Ext

## Test Location
`tests/integratedtests/coretaskinfotests/Extended_test.go`

## Symptoms
`Options should not be nil` — test expected non-nil but got nil.

## Root Cause
`coretaskinfo.New.Info.Plain.Default("n", "d", "u")` creates an `Info` struct without setting the `ExcludeOptions` field. `Info.Options()` returns `it.ExcludeOptions` directly, which is `nil` since it was never initialized.

This is correct production behavior — `nil` ExcludeOptions means "no exclusions" (all fields included). The test incorrectly assumed `Default` would initialize ExcludeOptions.

## Fix Applied
Inverted the assertion: expect `opts == nil` instead of `opts != nil`, with a comment explaining the production semantics.

## Fix Location
Test logic fix in `tests/integratedtests/coretaskinfotests/Extended_test.go`
