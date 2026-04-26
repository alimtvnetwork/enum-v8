# Batch Fix: jsoninternal, corepayload, coreonce, enumimpl, corevalidator

## 1. converters/anyItemConverter.go — ToStrings returns len 1 instead of 2

**Root cause**: `ItemsToStringsSkipOnNil(anyItems)` passes the `[]any` slice as a single variadic argument instead of spreading it.
**Fix**: Changed to `ItemsToStringsSkipOnNil(anyItems...)`.

## 2. coredata/corepayload/SessionInfo.go — IsEmpty returns false for empty struct

**Root cause**: `IsEmpty()` only checked `it == nil`, so `SessionInfo{}` (non-nil pointer to zero-value struct) returned false.
**Fix**: Added field checks: `it.Id == "" && it.User == nil && it.SessionPath == ""`.

## 3. coredata/coreonce/StringsOnce.go & IntegersOnce.go — IsEqual order-sensitive

**Root cause**: `IsEqual` compared elements positionally (`currentItems[i] != comparingItems[i]`), but tests pass items in different order expecting set equality.
**Fix**: Changed to frequency-map comparison (count occurrences, decrement on match).

## 4. corevalidator/SliceValidator.go — IsUsedAlready returns false after SetActualVsExpected

**Root cause**: `IsUsedAlready` checked `comparingValidators != nil` (lazy-initialized), but tests expect it to be true after `SetActualVsExpected`.
**Fix**: Added `isUsed bool` field, set to `true` in `SetActual` and `SetActualVsExpected`.

## 5. coreimpl/enumimpl/newBasicStringCreator.go — CreateAliasMapOnly wrong min/max

**Root cause**: `min` initialized to `""`, and no string is lexicographically less than `""`, so min stayed empty forever.
**Fix**: Initialize min/max from the first element of the range.

## 6. coreimpl/enumimpl/newBasicStringCreator.go — CreateUsingAliasMap missing nameWithIndexMap

**Root cause**: `nameWithIndexMap` was built in the function but not assigned to the returned `BasicString` struct.
**Fix**: Added `nameWithIndexMap: nameWithIndexMap` to the struct literal.
