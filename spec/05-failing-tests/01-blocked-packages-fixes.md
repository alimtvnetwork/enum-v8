# Blocked Packages Fixes — 2026-03-22

## Status: ✅ RESOLVED

### 1. coredynamictests — `ts.IsEqual()` missing argument
- **File**: `Coverage50_MapAnyItemDiff_LeftRight_test.go:419`
- **Error**: `not enough arguments in call to ts.IsEqual`
- **Root Cause**: `TypeStatus.IsEqual` requires `*TypeStatus` parameter: `func (it *TypeStatus) IsEqual(next *TypeStatus) bool`
- **Fix**: Added `sameTs := lr.TypeStatus()` and passed it: `ts.IsEqual(sameTs)`

### 2. corejsontests — `testStringer` redeclared
- **File**: `Coverage41_Serializer_Branches_test.go:97` (OLD file)
- **Root Cause**: Old `Coverage41_Serializer_Branches_test.go` was already deleted and replaced with proper AAA files. Blocked error was stale.
- **Fix**: No action needed — already resolved.

### 3. corestrtests — 8 API signature mismatches
- **File**: `Coverage41_Iteration8_test.go`
- **Errors & Fixes**:
  1. L142: `IsEqualsWithSensitive(b, false)` → `IsEqualsWithSensitive(false, b)` — args swapped
  2. L155: `AddWithWgLock("a", &wg)` → `AddWithWgLock(&wg, "a")` — args swapped
  3. L158: `AddStringsAsync([]string{...})` → `AddStringsAsync(&wg, []string{...})` — missing WaitGroup
  4. L159: `AddsAsync("d", "e")` → `AddsAsync(&wg, "d", "e")` — missing WaitGroup
  5. L257: `CsvLinesOptions(true, ", ")` → `CsvLinesOptions(true)` — takes 1 arg
  6. L326-327: `func(a any) bool` → `corestr.IsStringFilter(func(str string, index int) (string, bool, bool) {...})` — wrong type
  7. L344: `GetPagedCollection(1, 2)` → `GetPagedCollection(2)` — takes 1 arg
  8. L346: `GetSinglePageCollection(2)` → `GetSinglePageCollection(2, 1)` — takes 2 args
