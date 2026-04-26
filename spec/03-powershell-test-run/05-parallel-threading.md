# Parallel Threading Strategy

## Overview

The `run.ps1` script uses PowerShell 7's `ForEach-Object -Parallel` for concurrent execution across three phases:

1. **Pre-coverage compile check** — compile each test package in parallel
2. **Coverage test run** — run test packages with coverage profiles in parallel
3. **Pre-commit API check** — compile-check packages in parallel

## Thread Count Formula

```powershell
$throttle = [Math]::Min($packageCount, [Environment]::ProcessorCount * 2)
```

The throttle limit is set to **2× the logical processor count**, capped by the number of packages. This over-subscribes CPU cores to account for I/O-bound waits (disk reads, `go` compiler subprocess spawning).

### Rationale

| Factor | Detail |
|--------|--------|
| Go compilation is mixed CPU+I/O | File reads, linker invocation, temp file writes leave CPU idle gaps |
| `ForEach-Object -Parallel` uses runspaces | Lightweight — not full processes; low overhead per slot |
| Diminishing returns beyond 2× | At 3×+ cores, OS scheduling overhead and disk contention negate gains |
| Capped by package count | Never spawns more runspaces than there are work items |

### Example

On an 8-core machine with 30 test packages:

```
ProcessorCount = 8
Throttle = Min(30, 8 * 2) = 16
```

16 packages compile/test concurrently; remaining 14 queue behind.

## Sequential Fallback

Pass `--sync` to disable parallelism entirely:

```powershell
./run.ps1 TC --sync
```

In sync mode, packages run one-at-a-time in sorted order. Useful for:
- Debugging flaky tests with shared state
- Machines with limited RAM
- Reproducing deterministic output without the sync-phase sort

## Deterministic Output

Parallel execution produces nondeterministic completion order. The script collects all results into an array, then **sorts by package name** before displaying. See [03-parallel-sync-mechanism.md](./03-parallel-sync-mechanism.md) for the full pattern.

## Resource Impact

At 2× cores, expect:
- **CPU**: Sustained 60–90% utilization during compile/test phases
- **Memory**: ~200–300 MB for `pwsh` + child `go test` processes (varies with project size)
- **Disk**: Parallel temp files use unique sanitized names — no collisions

## Tuning

To adjust the multiplier, change the `* 2` factor in `run.ps1` at all three throttle calculation sites:
- Line ~453: compile check throttle
- Line ~648: coverage run throttle  
- Line ~1358: pre-commit check throttle

All three MUST use the same formula for consistent behavior.
