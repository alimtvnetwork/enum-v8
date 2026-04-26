# Blocked Packages Fixes — 2026-04-04

## Status: ✅ RESOLVED

### Summary: 28 blocked packages fixed

#### Category 1: Split-Recovery Helper Visibility (22 packages)
- **Root Cause**: Shared helpers (`caseV1Compat`, `testErr`, `covTempDir`, `covWriteFile`, `newTestRW`, `testUserCov23`, `makeCollectionCov23`, `makeTypedWrapperCov23`, `createNumberedUsers`) were defined in `*_test.go` files and became invisible when split-recovery isolated each test into its own subfolder.
- **Fix**: Moved all shared helpers to non-test `.go` files (`shared_compat_helpers.go`, `shared_coverage_helpers.go`, `shared_typed_helpers.go`) that get copied to every subfolder.
- **Affected packages**: corestrtests (Coverage43-54, SimpleSlice_S03, LinkedList_S12, SimpleSlice_S11a), chmodhelpertests (Coverage2, Coverage3, Coverage13_Remaining), corepayloadtests (Coverage24, Coverage25, TypedCollectionPagingEdge)

#### Category 2: API Signature Mismatches (6 packages)
1. **chmodhelpertests/Coverage18_Final_Gaps**: `GetRecursivePaths(string)` → `GetRecursivePaths(bool, string)`, `WriteBytes` → `Write`, `CreateDirWithFiles` extra arg removed
2. **coredynamictests/Coverage74_FinalGaps**: `MapAnyItems.ToKeyValCollection()` → `MapAnyItems.JsonMapResults()` (method doesn't exist on MapAnyItems)
3. **corejsontests/Coverage51_Gaps**: `GetPagedItems` → `GetPagedCollection`, `CastAny.Deserialize` → `CastAny.OrDeserializeTo`, `UsingResultsPtr` → `UsingResults`
4. **corerangetests/Coverage8_FinalGaps**: Removed tests for non-existent methods (`ClonePtr`, `IsEmpty`, `Length`, `String`, `Clone` on various types, `Within.StringRangeInt`), replaced with valid method tests
5. **corestrtests/Coverage27_FinalGaps**: `Collection.Strings("a","b")` → `Collection.Strings([]string{"a","b"})`, removed `IsEqualLinesInsensitive` tests
6. **corevalidatortests/Coverage18_Gaps**: `stringcompareas.EqualTextCompare` → `stringcompareas.Equal`
7. **enumimpltests/Coverage18_FinalGaps**: `BasicInt32/Int8/UInt16.Create` → `CreateUsingMap` with proper map args, `BasicString.Create` fixed arg count, `BasicByte.Create` fixed to 5 args, `ValueByName` → `GetIndexByName`
8. **corepayloadtests/Coverage25_Gaps**: `TypedPayloadWrapperDeserializeMust` → `TypedPayloadWrapperDeserialize` (non-Must version)
