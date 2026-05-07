# Test Failure RCA Patterns

## Pattern 9 — `Stringer` infinite recursion via `converters.AnyTo.ValueString(self)` (added 2026-05-07)

**Symptom:** A package's tests crash with `fatal error: stack overflow`, the trace alternating between `runtime/fmt/print.go` frames and one specific source line, e.g. `brackets/Pair.go:114` or `brackets/BothBrackets.go:100`. Go test runner reports it as a runtime failure for the package even though no `t.Error` was called.

**Root cause:** A `func (it T) String() string` implementation that returns `converters.AnyTo.ValueString(it)`. The helper falls through to `fmt.Sprintf("%v", it)`, which invokes the type's own `String()` method again — infinite recursion → stack overflow.

**Fix:** Never call `converters.AnyTo.ValueString` (or any `%v`/`fmt.Sprint` formatter) on the receiver inside a `String()` method. Always format the struct's fields explicitly:

```go
func (it Pair) String() string {
    return fmt.Sprintf("{Start:%s End:%s Category:%s}",
        it.Start.String(), it.End.String(), it.Category.String())
}
```

**Detection sweep:** `rg "converters\.AnyTo\.ValueString" --type go | rg -v _test.go` — review every hit; any one that passes `it`, `*it`, or the receiver's value to a `String()`/`MarshalJSON` method is a recursion bomb.

**Manifested as:** `brackets` package reported as RUNTIME FAILURE with stack overflow originating in `Pair.String` and `BothBrackets.String` (v1.14.0 fix).

---

## Pattern 8 — `-coverpkg` warning-only false-positive (added 2026-05-07)

**Symptom:** Packages reported as `Blocked` in pre-coverage compile check, or as `RUNTIME FAILURE` after the coverage run, but the captured diagnostic contains ONLY repeated lines like:

    warning: no packages being tested depend on matches for pattern github.com/alimtvnetwork/enum-v8/...

These warnings are emitted by Go when a `-coverpkg=` glob includes packages the test binary doesn't transitively import. They are harmless. Under parallel runspace load, the build cache contention sometimes causes `go test -c` to also exit non-zero transiently, masking the warning-only nature.

**Fix:** `scripts/Utilities.psm1` exports `Test-IsCoverpkgWarningOnlyOutput`. Call it as a final guard in:
- `CoverageCompileCheck.psm1` parallel runspace (set `$confirmed = $false` when warnings-only) and in the post-parallel re-confirmation pass.
- `CoverageRunner.psm1` sync + parallel coverage loops, so warning-only output never reaches `Add-RuntimeFailuresForPackage` / `Add-BuildErrorsForPackage`.

Manifested as: `licensetype` 58.9%, `onofftype` 32.1%, `rootcmdnames` 58.1% all blocked while passing tests cleanly elsewhere; `brackets` falsely listed under RUNTIME FAILURES despite "✓ No failing tests".

---

(prior patterns 1–7 retained — see git history)
