# TC Console Output Specification

## Overview

The `./run.ps1 TC` command produces a clean, structured console output with exactly **five sections**. No per-package compile/test rows are printed — only summaries.

## Section 1: Build Failure Packages

Appears only when one or more test packages fail `go test -c`.

```
  ┌─────────────────────────────────────────────────
  │ BLOCKED PACKAGES (2 failed to compile)
  │
  │   ✗ corestrtests
  │   ✗ corecmptests
  │
  │ These packages will be SKIPPED in coverage.
  │ Fix their build errors to include them.
  └─────────────────────────────────────────────────
```

If all packages compile: `✓ All N packages compiled successfully`

## Section 2: Failing Test Summary

Appears only when tests produce `--- FAIL:` output.

```
  ┌─────────────────────────────────────────────────
  │ FAILING TESTS (3 failed)
  │
  │   ✗ Test_Cov8_SomeMethod_NilInput
  │   ✗ Test_Cov8_OtherMethod_EmptySlice
  │   ✗ Test_Cov4_KeyMaker_Overflow
  │
  │ See data/test-logs/failing-tests.txt for details.
  └─────────────────────────────────────────────────
```

## Section 3: Coverage Summary

Per-source-package coverage table inside a box, sorted by percentage descending.

```
  ┌─────────────────────────────────────────────────
  │ COVERAGE SUMMARY
  │
  │  100.0%  chmodhelper
  │  100.0%  coreonce
  │  95.6%   keymk
  │  3.3%    corestr
  │
  │  total:  (statements) 62.5%
  │  ⚠ 42 function(s) below 50% coverage
  └─────────────────────────────────────────────────
```

Color coding: ≥100% green, ≥80% yellow, <80% red.

## Section 4: Coverage Diff

Appears after the Coverage Summary. Compares current per-package coverage against the previous run's snapshot (`data/coverage/coverage-previous.json`). Only shown when the `DashboardUI` module is loaded.

- **First run**: Prints `No previous coverage data available for comparison.` (no snapshot exists yet).
- **Subsequent runs**: Renders a boxed diff table showing per-package deltas.
- **No changes**: Prints `✓ No coverage changes detected.`

```
  ╔══════════════════════════════════════════════════╗
  ║  C O V E R A G E   D I F F                      ║
  ╠══════════════════════════════════════════════════╣
  ║  Package                  Prev    Curr    Delta  ║
  ║  ────────────────────────────────────────────────║
  ║  corestr                  3.3%    5.1%   +1.8%  ║
  ║  keymk                   95.6%   95.6%    0.0%  ║
  ║  coreonce               100.0%  100.0%    0.0%  ║
  ║  chmodhelper             100.0%  100.0%    0.0%  ║
  ╚══════════════════════════════════════════════════╝
  Coverage snapshot saved → data/coverage/coverage-previous.json
```

Color coding for Delta column: positive (improvement) green, negative (regression) red, zero dim/muted.

Special status indicators:
- `NEW` — package appears for the first time (no previous entry)
- `REMOVED` — package was in previous snapshot but absent in current run

After rendering the diff, the current coverage data is saved as the new snapshot for the next run.

### Applies to both TC and PC commands

The `./run.ps1 TCP <package>` command uses the same comparison flow. Coverage is aggregated per source package from `go tool cover -func` output, compared against the shared snapshot, and then saved.

## Section 5: Written Files Summary

Lists all generated report files.

```
  ┌─────────────────────────────────────────────────
  │ WRITTEN FILES
  │  data/coverage/coverage.out
  │  data/coverage/coverage.html
  │  data/coverage/coverage-summary.txt
  │  data/coverage/coverage-summary.json
  │  data/coverage/per-package-coverage.txt
  │  data/coverage/per-package-coverage.json
  │  data/coverage/blocked-packages.txt      (if blocked)
  │  data/coverage/blocked-packages.json     (if blocked)
  └─────────────────────────────────────────────────
```

## What Is NOT Printed

- Individual `✓ pkgname` lines during compile check
- Individual `✓ pkgname — X%` lines during coverage run
- Raw `go test` output (written to `data/test-logs/raw-output.txt` only)
- Debug lines (`[debug] ...`)
- Merge progress messages
