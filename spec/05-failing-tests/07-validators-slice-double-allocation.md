# 03 — Validators() Slice Double-Allocation

## Failing Tests
- `Test_RangeSegmentsValidator_VerifyUpto`
- `Test_RangeSegmentsValidator_VerifyUptoDefault`
- `Test_RangeSegmentsValidator_SetActualOnAll`

## Root Cause

`RangeSegmentsValidator.Validators()` in `corevalidator/RangeSegmentsValidator.go` used:

```go
validators := make([]HeaderSliceValidator, it.LengthOfVerifierSegments())
```

This creates a slice with **length = N** (zero-valued elements), then uses `append` to add N real validators, resulting in **2N total elements** (N zero + N real).

**Effects:**
- `len(validators)` returned 4 instead of 2 for a 2-segment validator
- Zero-valued validators at the front caused `VerifyUpto` and `VerifyUptoDefault` to return false errors on matching segments
- `IsMatch(true)` evaluated against zero-valued validators, producing incorrect results

## Solution

Changed to capacity-only allocation:

```go
validators := make([]HeaderSliceValidator, 0, it.LengthOfVerifierSegments())
```

## File Changed
- `corevalidator/RangeSegmentsValidator.go` (line 31)

## Learnings
- `make([]T, n)` + `append` = double allocation. Always use `make([]T, 0, n)` when building via append.
- This is a common Go gotcha that static analysis tools (`go vet`) don't catch.

## What Not to Repeat
- Never use `make([]T, length)` followed by `append`. Pick one: either pre-allocate and index-assign, or use `make([]T, 0, cap)` with append.
