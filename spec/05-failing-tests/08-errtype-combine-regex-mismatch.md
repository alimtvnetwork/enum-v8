# 04 — ErrType Combine Regex Mismatch

## Failing Test
- `Test_ErrType_Combine_Verification` (Case 2)

## Root Cause

Test case expected regex `.*BytesAreNilOrEmpty.*` but `BytesAreNilOrEmptyType.Combine("", "")` returns the human-readable string from `RawErrorType.String()`:

```
"Bytes data either nil or empty."
```

The type is declared as:
```go
BytesAreNilOrEmptyType RawErrorType = "Bytes data either nil or empty."
```

`Combine` calls `CombineWithMsgTypeNoStack` which uses `genericMsg.String()` — returning the **string value**, not the Go identifier name.

## Solution

Updated test expectation to match the actual string output:

```go
ExpectedInput: ".*Bytes data either nil or empty.*",
```

## File Changed
- `tests/integratedtests/errcoretests/ErrType_testcases.go` (line 34)

## Learnings
- `RawErrorType` is a `string` typedef — `.String()` returns the string value, not the Go type/variable name.
- Regex-based assertions should be validated against actual runtime output before committing.

## What Not to Repeat
- Don't assume type identifier names appear in runtime output of string-backed types.
