# 08 — SliceValidator Diagnostic Format Drift

## Failing Tests
- `Test_SliceValidator`
- `Test_SliceValidator_FirstError`

## Root Cause

The `LineDiffToString` function in `errcore/LineDiff.go` was updated with two formatting changes that weren't reflected in the SliceValidator test expectations:

1. **Indentation change**: Mismatch lines shifted from 11/9 spaces (actual/expected) to 14/12 spaces to align colons at column 21 per the project's diagnostic formatting standard.

2. **Removed `=== End Diff ===` footer**: The trailing footer line was removed to minimize diagnostic noise per project conventions.

These changes caused a length mismatch (expected 34 lines including footer, actual 33 lines without it), which triggered the `initialVerifyError` length check before any content comparison could occur:

```
ActualLines, ExpectedLines Length is not equal. - Expect ["33"] != ["34"] Actual
```

## Solution

Updated both test case expectations in `slicevalidators_testCases.go`:
- Changed indentation from `"           actual"` (11 spaces) to `"              actual"` (14 spaces)
- Changed indentation from `"         expected"` (9 spaces) to `"            expected"` (12 spaces)
- Removed the `"=== End Diff ==="` line from expected output

## File Changed
- `tests/integratedtests/corevalidatortests/slicevalidators_testCases.go` (both test case blocks)

## Learnings
- Diagnostic format tests are fragile — any formatting change requires updating all snapshot-style expectations.
- The project's diagnostic formatting standards (column alignment, footer conventions) must be cross-referenced when modifying output generators.

## What Not to Repeat
- When changing diagnostic output formatting, always grep for test files that assert on the exact output format and update them in the same commit.
