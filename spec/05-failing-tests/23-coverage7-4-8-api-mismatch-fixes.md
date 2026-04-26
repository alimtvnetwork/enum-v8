# Fix: API Mismatches in Coverage7, Coverage4, Coverage8

## Date: 2026-03-15

## enumimpltests/Coverage7_test.go

| Issue | Root Cause | Fix |
|-------|-----------|-----|
| `enumimpl.Format.ToString` | `Format` doesn't exist as a var | Removed — non-existent API |
| `enumimpl.FormatUsingFmt.NameValue` | `FormatUsingFmt` is a function, not a struct | Removed — can't call without `formatter` interface impl |
| `enumimpl.PrependJoin.DotJoin` | Doesn't exist | Removed |
| `enumimpl.JoinPrependUsingDot` | Doesn't exist | Removed |
| `enumimpl.ConvAnyValToInteger` | Wrong name, wrong return | `ConvEnumAnyValToInteger` returns `int` only (not `int, bool`) |
| `enumimpl.NameWithValue{...}` | It's a function, not a struct | Changed to `NameWithValue(10)` function call |
| `enumimpl.AllNameValues(m)` | Wrong signature `(map)` | Removed — requires `([]string, any)` with reflect slice |
| `enumimpl.UnsupportedNames(s, a)` | Param order wrong, 2nd is variadic | Fixed to `(allNames, supported...)` |
| `KeyAnyVal{Value: ...}` | Field is `AnyValue` not `Value` | Fixed field name |

Added: `DiffLeftRight` method tests, `DefaultDiffCheckerImpl`, `LeftRightDiffCheckerImpl` tests.

## corejsontests/Coverage4_test.go

| Issue | Root Cause | Fix |
|-------|-----------|-----|
| `corejson.New(...).JsonString()` | `JsonString()` is pointer receiver | Use `NewPtr()` or access `r.Bytes` directly |
| `r.Err()` | Method doesn't exist | Changed to `r.MeaningfulError()` |
| `corejson.JsonString(...)` | Returns `(string, error)` not `string` | Capture both return values |
| `BytesCloneIf(false, ...)` samePtr | Returns `[]byte{}` not original | Fixed: expect `len: 0` |

## corestrtests/Coverage8_test.go

| Issue | Root Cause | Fix |
|-------|-----------|-----|
| `corestr.NewSimpleStringOnce("hello")` | Function doesn't exist | Use `corestr.New.SimpleStringOnce.Init("hello")` |
| `NewValidValue("")` → `IsValid: false` | `NewValidValue` always sets `IsValid: true` | Fixed expectation to `true` |
| `CloneSlice(nil)` → `isNil: true` | Returns `[]string{}` not nil | Fixed: check `len: 0` |
