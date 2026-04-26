# Zero Coverage Cascade: Build Failure Blocks All Test Packages

## Observed Symptoms

### 0% Coverage for `codestack` and `coreimpl/enumimpl/enumtype`
Despite having comprehensive test files in `tests/integratedtests/codestacktests/` and `tests/integratedtests/enumtypetests/`, these packages show **0% coverage** in the coverage summary.

### `Test_Variant_OnlySupportedErr` Failure (stale binary)
Error points to `Extended_test.go:463` but the file only has 302 lines. This indicates the Go test binary is stale and was compiled from a previous version of the file.

## Root Cause

### Build Failure Cascade
The `run.ps1 -tc` (test coverage) command performs a global build check at line 375:

```powershell
Push-Location tests
try { if (-not (Invoke-BuildCheck "./...")) { return } }
finally { Pop-Location }
```

This checks **ALL** test packages under `tests/`. If **any** package fails to build (e.g., `bytetypetests` had pointer receiver interface mismatch errors), the entire coverage run **aborts**. Result:
- No tests execute
- No coverage data is generated
- Previously-working packages show 0% or stale coverage
- New test packages never get their first coverage measurement

### Why Some Packages Still Show Coverage
The coverage summary may include data from **previous successful runs** that was cached or merged. Packages that had tests before the build-breaking change retain their old coverage numbers, while newly-added test packages (codestack, enumtype) show 0% because they never had a successful run.

### Stale Test Binary
Go's test binary cache may retain debug info from older source files, causing error messages to reference line numbers that no longer exist in the current source (e.g., `Extended_test.go:463` on a 302-line file).

## Fix

### Primary Fix
Resolve the `bytetypetests` build error (pointer receiver interface mismatch — see `15-bytetype-pointer-receiver-interface-mismatch.md`). Once **all** test packages compile, the build check passes and all tests run, generating coverage for every package.

### OnlySupportedErr Fix
Already applied in `Extended2_test.go:355-369`: uses `v.StringRanges()` (plain names like `"One"`) instead of `v.AllNameValues()` (formatted names like `"One(1)"`), matching the format `OnlySupportedErr` internally compares against.

## Prevention
When build-breaking changes are introduced in one test package, they block coverage measurement for **all** packages. Always run `go build ./tests/integratedtests/...` before a full coverage run to catch compilation issues early.
