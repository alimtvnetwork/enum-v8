# PowerShell Implementation Spec

> **Purpose**: Documents the implementation details of the `run.ps1` modular toolchain —
> module loading, error guarding, phase tracking, Go syntax validation, test patterns,
> and coverage generation. For AI agents working on the PowerShell tooling itself.

---

## Table of Contents

1. [Module Architecture](#1-module-architecture)
2. [Module Loading & Import Order](#2-module-loading--import-order)
3. [Error Guarding](#3-error-guarding)
4. [Phase Tracking System](#4-phase-tracking-system)
5. [Go Syntax Validation Pipeline](#5-go-syntax-validation-pipeline)
6. [Go Test Patterns](#6-go-test-patterns)
7. [Coverage Generation Workflow](#7-coverage-generation-workflow)
8. [Error Attribution System](#8-error-attribution-system)
9. [AI Agent Interaction Guide](#9-ai-agent-interaction-guide)

---

## 1. Module Architecture

`run.ps1` is a thin dispatcher (~167 lines) that imports `.psm1` modules from `scripts/`.

### Module Inventory

| Module | Lines | Exports | Description |
|--------|-------|---------|-------------|
| `DashboardUI.psm1` | ~1230 | 15+ functions | ANSI rendering, phase tracking, coverage tables, diff |
| `Utilities.psm1` | ~370 | 13 functions | Console output, error extraction, line filtering |
| `TestLogWriter.psm1` | ~200 | 1 function | Go test output parser → structured log files |
| `TestRunner.psm1` | ~250 | 7 functions | Test execution, build checks, git operations |
| `CoverageRunner.psm1` | ~1607 | 2 functions | TC + TCP coverage pipelines (largest module) |
| `BuildTools.psm1` | ~135 | 7 functions | Build, format, vet, tidy, clean |
| `PreCommitCheck.psm1` | ~310 | 1 function | PC pre-commit validation pipeline |
| `GoConvey.psm1` | ~60 | 1 function | GoConvey browser test runner launcher |
| `Help.psm1` | ~155 | 3 functions | Help display, fail log viewer, integrated tests |

### Dependency Graph

```
DashboardUI          (standalone)
Utilities            (standalone, optional DashboardUI fallback)
TestLogWriter        (→ Utilities)
TestRunner           (→ Utilities, TestLogWriter)
CoverageRunner       (→ Utilities, TestLogWriter, TestRunner, DashboardUI)
BuildTools           (→ Utilities)
PreCommitCheck       (→ Utilities, DashboardUI)
GoConvey             (standalone)
Help                 (→ Utilities, TestLogWriter, TestRunner)
```

**Rule**: No circular dependencies. Modules may only depend on modules above them in this graph.

---

## 2. Module Loading & Import Order

### Import Pattern in `run.ps1`

```powershell
$modulePath = Join-Path $PSScriptRoot "scripts" "ModuleName.psm1"
if (Test-Path $modulePath) {
    Import-Module $modulePath -Force -DisableNameChecking
}
```

### Import Order (must match dependency graph)

1. `DashboardUI.psm1` — standalone, provides ANSI + phase tracking
2. `Utilities.psm1` — standalone, provides console helpers
3. `TestLogWriter.psm1` — depends on Utilities
4. `TestRunner.psm1` — depends on Utilities + TestLogWriter
5. `CoverageRunner.psm1` — depends on all above
6. `BuildTools.psm1` — depends on Utilities
7. `GoConvey.psm1` — standalone
8. `PreCommitCheck.psm1` — depends on Utilities + DashboardUI
9. `Help.psm1` — depends on Utilities + TestLogWriter + TestRunner

### Flags

| Flag | Purpose |
|------|---------|
| `-Force` | Re-imports module even if already loaded (picks up changes during dev) |
| `-DisableNameChecking` | Suppresses warnings about non-standard verb names |
| `-ErrorAction SilentlyContinue` | Used on DashboardUI import only (optional module) |

### Shared Variables

`run.ps1` defines shared variables **before** module imports:

```powershell
$TestLogDir = Join-Path $PSScriptRoot "data" "test-logs"
```

Modules access these via the parent scope since `Import-Module` shares the session scope.

---

## 3. Error Guarding

### DashboardUI Guard Pattern

Every call to a DashboardUI function is wrapped:

```powershell
if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
    Register-Phase "Phase Name" "pass" "detail text"
}
```

### Why Not Try/Catch

1. No exception overhead
2. No partial error records in edge-case PowerShell hosts
3. Readable intent: "run only if available"
4. Granular per-call-site control

### Module Import Guard

```powershell
if (Test-Path $modulePath) {
    Import-Module $modulePath -Force -ErrorAction SilentlyContinue
}
```

### Design Principle

> **DashboardUI is always additive, never required.** Core functionality (testing, coverage, compilation) works identically with or without the UI module.

---

## 4. Phase Tracking System

### How It Works

1. `Reset-Phases` clears the `$script:phases` ordered dictionary at command start
2. `Register-Phase "Name" "status" "detail"` records each phase's result
3. `Write-PhaseSummaryBox` renders the bordered summary at command end

### Phase Statuses

| Status | Icon | Color | Meaning |
|--------|------|-------|---------|
| `pass` | `✓` | Lime | Phase completed successfully |
| `fail` | `✗` | Red | Phase failed — may block execution |
| `warn` | `⚠` | Yellow | Phase completed with warnings |
| `skip` | `○` | Muted | Phase was skipped (flag or not needed) |

### TC Phases (10)

`Git Pull → Dependencies → Data Cleanup → SafeTest Lint → Auto-Fixer → Syntax Check → Compile Check → Split Recovery → Coverage Run → Coverage Report`

### PC Phases (5)

`Regression Guard → SafeTest Lint → Auto-Fixer → Syntax Check → API Compile Check`

---

## 5. Go Syntax Validation Pipeline

### Tools

| Tool | Location | Purpose | Skippable |
|------|----------|---------|-----------|
| **autofix** | `scripts/autofix/main.go` | Auto-fixes common Go syntax issues | `--no-autofix` |
| **bracecheck** | `scripts/bracecheck/main.go` | Validates brace/bracket/paren balance | `--skip-bracecheck` |
| **safetest-boundaries** | `scripts/check-safetest-boundaries.ps1` | Lints safeTest boundary patterns + empty-if blocks | No |
| **regression-guard** | `scripts/check-integrated-regressions.ps1` | Detects known API-drift patterns | No |

### Execution Order (in TC and PC)

```
1. Regression guard (PC only)
2. SafeTest boundary lint
3. Go auto-fixer (unless --no-autofix or --skip-bracecheck)
4. Bracecheck syntax validation (unless --skip-bracecheck)
5. Compile check
```

### Auto-Fixer

Runs as a Go program: `go run ./scripts/autofix/`

- Scans `Coverage*.go` and `*_testcases.go` files
- Fixes: trailing commas, missing imports, formatting
- Supports `--dry-run` for preview mode
- Registers as `Auto-Fixer` phase
- **Important**: The Go tool output may already contain a `✓` prefix. Always strip leading `✓` from output before passing to `Write-Success` (which adds its own `✓`) to avoid double checkmarks: `$str = ($out | Out-String).Trim() -replace '^\s*✓\s*', ''`

### Bracecheck

Runs as a Go program: `go run ./scripts/bracecheck/`

- Scans all `.go` files in the project
- Validates balanced `{}`, `()`, `[]`
- Reports file:line for mismatches
- Writes results to `data/coverage/syntax-issues.txt`
- Registers as `Syntax Check` phase
- **Important**: Same `✓`-stripping rule as Auto-Fixer applies here

---

## 6. Go Test Patterns

### Test Execution (`Invoke-GoTestAndLog`)

```powershell
$ErrorActionPreference = "Continue"
$output = & go test -v -count=1 $targetPkg 2>&1 | ForEach-Object { $_.ToString() }
$exitCode = $LASTEXITCODE
$ErrorActionPreference = "Stop"
```

Key flags:
- `-v` — verbose output (required for log parsing)
- `-count=1` — disable test caching
- `2>&1` — merge stderr into stdout for unified parsing

### Build Check Gate (`Invoke-BuildCheck`)

Before running tests, compile-check the target:

```powershell
$out = & go build $targetPkg 2>&1 | ForEach-Object { $_.ToString() }
if ($LASTEXITCODE -ne 0) {
    # Write build errors to failing-tests.txt
    # Skip test execution entirely
}
```

### Test Log Parsing (`Write-TestLogs`)

Two-pass parser:

**Pass 1** — Classify each test:
- `=== RUN` → start tracking test
- `--- PASS:` → add to passing list
- `--- FAIL:` → add to failing list + capture diagnostics

**Pass 2** — Write structured files:
- `passing-tests.txt` — sorted names with count + timestamp
- `failing-tests.txt` — summary section + detailed diagnostics
- `raw-output.txt` — unprocessed output

---

## 7. Coverage Generation Workflow

### TC Pipeline (`Invoke-TestCoverage` in `CoverageRunner.psm1`)

```
Phase 1: Pre-checks (lint, autofix, bracecheck)
Phase 2: Package discovery (source pkgs, test pkgs)
Phase 3: Pre-coverage compile check (parallel go test -c)
Phase 4: Split recovery (per-file recheck for blocked pkgs)
Phase 5: Coverage run (parallel go test -coverprofile)
Phase 6: Profile merge (MAX-count dedup)
Phase 7: Report generation (HTML, JSON, TXT)
Phase 8: Coverage diff (compare against previous snapshot)
Phase 9: AI prompt generation (for coverage gaps)
```

### Parallel Execution

```powershell
$throttle = [Math]::Min($pkgCount, [Environment]::ProcessorCount * 2)
$results = $packages | ForEach-Object -ThrottleLimit $throttle -Parallel {
    # ... compile or test each package ...
    [pscustomobject]@{ Pkg = $pkg; ExitCode = $ec; Output = $out }
}
# CRITICAL: Sort results by Pkg for deterministic output
foreach ($r in ($results | Sort-Object Pkg)) { ... }
```

`--sync` flag falls back to sequential `foreach`.

### Profile Merge (MAX Count)

Multiple partial profiles merged using highest count per unique line:

```powershell
foreach ($line in $allPartialLines) {
    if ($line -match '^(\S+\.go:\d+\.\d+,\d+\.\d+\s+\d+)\s+(\d+)$') {
        $key = $Matches[1]
        $count = [int]$Matches[2]
        if (-not $map.ContainsKey($key) -or $count -gt $map[$key]) {
            $map[$key] = $count
        }
    }
}
```

### Coverage Diff

Compares current per-package coverage against `data/coverage/coverage-previous.json`:

| Indicator | Meaning |
|-----------|---------|
| `▲` | Coverage improved |
| `▼` | Coverage regressed |
| `★` | New package |
| `✗` | Removed package |
| `=` | No change |

After rendering, current data is saved as the new snapshot.

### TCP Pipeline (`Invoke-PackageTestCoverage`)

Same as TC but scoped to a single package. Uses the same diff/snapshot flow.

---

## 8. Error Attribution System

### Overview

Every error report — build failures, runtime failures, blocked packages — includes **source attribution** identifying the exact `.psm1` module and function that triggered the failure. This enables rapid root-cause analysis when reviewing logs.

### `Get-CallerSource` Function

Defined in `Utilities.psm1`. Walks `Get-PSCallStack` to find the first caller outside `Utilities.psm1` itself.

```powershell
$source = Get-CallerSource
# Returns: "CoverageRunner.psm1 → Invoke-TestCoverage"
```

**Behaviour:**
- Skips internal frames (`<ScriptBlock>`, `Get-CallerSource`, `Utilities.psm1`)
- Returns `"ModuleName.psm1 → FunctionName"` when both are available
- Falls back to script name or function name alone
- Returns `"unknown"` if no meaningful frame is found

### Where Attribution Appears

| Module | Context | Format |
|--------|---------|--------|
| `TestRunnerCore.psm1` | `Invoke-BuildCheck` build failure | `# Source:` header in `failing-tests.txt` + console `Write-Fail` |
| `TestRunnerCore.psm1` | `Invoke-GitPull`, `Invoke-FetchLatest` | Console `Write-Fail` with source |
| `TestRunner.psm1` | `Invoke-AllTests`, `Invoke-PackageTests` | Console `Write-Fail` with source |
| `TestLogWriter.psm1` | `Write-TestLogs` pass/fail logs | `# Source:` header in `passing-tests.txt` and `failing-tests.txt` |
| `CoverageReportJson.psm1` | JSON reports | `"source"` field in `build-errors.json` and `runtime-failures.json` |
| `CoverageReportJson.psm1` | Text reports | `# Source:` header in `build-errors.txt` and `runtime-failures.txt` |
| `CoverageRunner.psm1` | Blocked packages list + exit paths | `# Source:` header in `blocked-packages.txt` + console `Write-Fail` |
| `CoverageCompileCheck.psm1` | Compile-check failures | Console `Write-Fail` with source for both sync and parallel modes |
| `CoveragePreChecks.psm1` | Pre-check validation failures | Console `Write-Fail` with source |
| `CoverageSplitRecovery.psm1` | Subfolder recovery failures | Console `Write-Fail` with source |
| `CoverageReportHtml.psm1` | HTML report generation errors | Console `Write-Fail` with source |
| `BuildTools.psm1` | `Invoke-Build`, `Invoke-Vet` | Console `Write-Fail` with source |
| `PreCommitCheck.psm1` | Pre-commit regression failures | Console `Write-Fail` + `"source"` field in JSON report |
| `GoConvey.psm1` | GoConvey install failure | Console `Write-Fail` with source |
| `Help.psm1` | Help/utility error paths | Console `Write-Fail` with source |
| `PackageCoverage.psm1` | Package coverage failures | Console `Write-Fail` with source |

### Report Format Examples

**Text reports:**
```
# Source: CoverageRunner.psm1 → Invoke-TestCoverage
```

**JSON reports:**
```json
{
  "source": "CoverageRunner.psm1 → Invoke-TestCoverage",
  "generatedAt": "2026-04-03T12:00:00",
  ...
}
```

**Console output:**
```
  ✗ Build failed — skipping tests (source: TestRunnerCore.psm1 → Invoke-BuildCheck)
  ✗ Blocked: subpkg/foo (source: CoverageCompileCheck.psm1 → Invoke-CoverageCompileCheck)
```

### Design Rules

1. **Always use `Get-CallerSource`** in sync code paths where the call stack is available
2. **Hardcode the source string** in parallel (`ForEach-Object -Parallel`) blocks, since `Get-CallerSource` cannot cross thread boundaries
3. **Never omit attribution** — every error path must include a source reference

### Error Extraction Pipeline (4-Tier Fallback)

When a package fails (`[build failed]`, `[setup failed]`, or runtime crash), the toolchain extracts diagnostic lines using a 4-tier fallback chain. Each tier is tried in order; the first to return non-empty results wins.

| Tier | Function | Module | What it captures |
|------|----------|--------|-----------------|
| 1 | `Extract-BuildErrorLines` | `ErrorExtractor.psm1` | `.go:line:` errors, `[build failed]`, `[setup failed]`, `# pkg` headers |
| 2 | `Extract-ExecutionFailureLines` | `ErrorExtractor.psm1` | Tier 1 + `panic:`, `fatal error:`, `--- FAIL:`, `FAIL pkg`, `exit status` |
| 3 | `Extract-SetupFailedContext` | `ErrorExtractor.psm1` | Walks backward from `[setup failed]`/`[build failed]` FAIL lines, captures up to 10 preceding context lines |
| 4 | `Get-RawFallbackLines` | `ErrorExtractor.psm1` | All non-empty lines after noise removal (last resort) |

**Why 4 tiers?** Go's `[setup failed]` output includes plain-text error messages (e.g., missing fixtures, `init()` errors) that don't follow `.go:line:` or `panic:` patterns. Tiers 1–2 miss these. Tier 3 captures them by walking backward from the FAIL marker. Tier 4 ensures nothing is ever silently lost.

#### `Extract-SetupFailedContext`

```powershell
$context = Extract-SetupFailedContext $rawOutput -ContextLineCount 10
# Returns preceding context lines + the FAIL line for each [setup failed] occurrence
```

- Scans for lines matching `[setup failed]` or `[build failed]` with `FAIL`
- Walks backward up to N lines (default 10) to capture the error reason
- Deduplicates and strips noise lines (warnings, empty lines)

#### Where the Fallback Chain is Used

| Module | Function | Report |
|--------|----------|--------|
| `ErrorParser.psm1` | `Add-BuildErrorsForPackage` | Per-package build error accumulation |
| `ErrorParser.psm1` | `Add-RuntimeFailuresForPackage` | Per-package runtime failure accumulation |
| `CoverageReportJson.psm1` | `Write-BuildErrorsReport` | `build-errors.txt` / `build-errors.json` |
| `CoverageRunner.psm1` | Blocked packages writer | `blocked-packages.txt` / `blocked-packages.json` |

---

## 9. AI Agent Interaction Guide

### How to Modify the Toolchain

1. **Identify the module** — use `scripts/README.md` dependency graph
2. **Read the module** before editing — never infer function signatures
3. **Follow the export pattern** — `Export-ModuleMember -Function @('Name')`
4. **Add doc blocks** — `.SYNOPSIS`, `.PARAMETER`, `.EXAMPLE` on every function
5. **Guard DashboardUI calls** — always wrap with `Get-Command ... -ErrorAction SilentlyContinue`

### Adding a New Command

1. Create or extend a module in `scripts/`
2. Export the function via `Export-ModuleMember`
3. Add a switch case in `run.ps1`
4. Update `Show-Help` in `Help.psm1`
5. Update `scripts/README.md`

### Common Pitfalls

| Pitfall | Prevention |
|---------|------------|
| Circular module dependency | Check dependency graph before importing |
| Missing `$ErrorActionPreference = "Continue"` around `go` calls | `go test` returns non-zero on test failure; `"Stop"` would throw |
| Forgetting to restore `$ErrorActionPreference` | Always reset after the `go` call block |
| DashboardUI call without guard | Module may not be loaded — always guard |
| Non-deterministic parallel output | Always `Sort-Object Pkg` after parallel execution |

### Key Files for Context

| File | Purpose |
|------|---------|
| `run.ps1` | Thin dispatcher — read this first |
| `scripts/README.md` | Module map + dependency graph |
| `spec/04-tooling/02-powershell-dashboard-ui.md` | UI rendering spec |
| `spec/03-powershell-test-run/09-ai-agent-complete-reference.md` | Complete AI agent reference |
| `.lovable/memory/workflow/06-powershell-refactor-plan.md` | Refactoring roadmap |

---

## Version History

| Date | Change |
|------|--------|
| 2026-04-03 | Added §8 Error Extraction Pipeline (4-tier fallback) with `Extract-SetupFailedContext` |
| 2026-03-31 | Initial creation — documents modular architecture post-refactor |
