# Test Failure RCA Patterns

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
