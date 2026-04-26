# 06 — LazyRegex Compile/New Test Expectations

## Failing Tests
- `Test_LazyRegex_Compile_Verification` (Cases 1, 2)
- `Test_New_Lazy_Verification` (Case 4)

## Root Cause

Test expectations were written with incorrect assumptions about `LazyRegex` behavior:

### Compile Case 1 (invalid pattern `[bad`)
- **Expected**: `isCompiled = false` after failed compile
- **Actual**: `isCompiled = true` — the `Compile()` method sets `it.isCompiled = true` after **any** compile attempt (success or failure). `isCompiled` means "compile was attempted", not "compile succeeded".
- **Fix**: Changed expected `isCompiled` from `"false"` to `"true"`.

### Compile Case 2 (empty pattern `""`)
- **Expected**: `hasError = false`
- **Actual**: `hasError = true` — `Compile()` on an undefined lazy regex (empty pattern) returns `errors.New("lazy regex is undefined or nil")`.
- **Fix**: Changed expected `hasError` from `"false"` to `"true"`.

### New.Lazy Case 4 (empty pattern `""`)
- **Expected**: `isFailedMatch = false`
- **Actual**: `isFailedMatch = true` — `IsFailedMatch` on an undefined regex returns `true` because `Compile()` returns nil regex, triggering the `regEx == nil` guard.
- **Fix**: Changed expected `isFailedMatch` from `"false"` to `"true"`.

## Files Changed
- `tests/integratedtests/regexnewtests/LazyRegex_Methods_testcases.go` (Cases 1, 2)
- `tests/integratedtests/regexnewtests/LazyRegex_testcases.go` (Case 4)

## Learnings
- `isCompiled` is a "has been attempted" flag, not a "succeeded" flag. Use `isApplicable` to check success.
- Undefined/empty pattern lazies consistently return error/false/true for all operations — no silent success.

## What Not to Repeat
- Don't write test expectations based on assumed semantics. Always trace through the actual implementation path first.
