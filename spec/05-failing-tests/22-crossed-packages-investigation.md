# Investigation: Crossed Packages in Coverage Report

## Packages Investigated

### bytetype, codestack, corecmp (0/0 stmts)

**Symptom:** `matched=0/0 stmts=0/0` — no coverage data at all.

**Root Cause:** Test binaries crash at runtime, preventing partial coverage profiles from being written.

- **codestack:** Fixed in spec #14 — nil pointer panic in `newCreator.Create()` when `runtime.Caller` fails. The `codestacktests` binary would crash before writing its profile.
- **bytetype & corecmp:** No obvious panic source found via code review. Tests compile fine (not in blocked packages list). Likely a cascading effect — if these test packages indirectly depend on codestack initialization, the codestack panic could propagate. **Action:** Run `./run.ps1 TC` after syncing the codestack fix to verify if these resolve. If not, run `go test ./tests/integratedtests/bytetypetests/` and `go test ./tests/integratedtests/corecmptests/` separately to capture panic traces.

### chmodhelper (55.2% — below 80% threshold)

**Symptom:** Coverage exists but is below the 80% minimum.

**Fix Applied:** Added `Coverage2_test.go` with ~60 tests covering:
- `fileWriter` (All, AllLock, Remove, RemoveIf, ParentDir, Chmod, ChmodFile)
- `fileBytesWriter` (Default, WithDir, WithDirLock, WithDirChmod, WithDirChmodLock, Chmod)
- `fileStringWriter` (Default, DefaultLock, All, Chmod, ChmodLock)
- `fileReader` (Read, ReadBytes)
- `dirCreator` (If, IfMissing, IfMissingLock, Default, DefaultLock, Direct, DirectLock, ByChecking)
- `chmodVerifier` (GetRwxFull, GetRwx9, IsEqual, IsMismatch, GetExisting, PathIf, RwxFull)
- `chmodApplier` (ApplyIf, OnMismatchOption, PathsUsingFileModeConditions, Default)
- `SimpleFileReaderWriter` (InitializeDefault, IsExist, Write, Read, Expire, Clone, Json, etc.)
- `simpleFileWriter` Lock/Unlock
- `IsPathExistsPlusFileInfo`

### corecsv (95.5% — marked ✗)

**Symptom:** Above 80% overall but likely has individual functions below 80%.

**Fix Applied:** Added 4 tests to `Coverage_test.go` covering:
1. `AnyItemsToCsvString` empty-items branch
2. `StringsToCsvString` empty-items branch
3. `RangeNamesWithValuesIndexes` empty-items branch
4. `AnyToValuesTypeStrings` empty-string item branch (`finalString == ""`)

These 6 uncovered stmts (133-127=6) were the gap.
