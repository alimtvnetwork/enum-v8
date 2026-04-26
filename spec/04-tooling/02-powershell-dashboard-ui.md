# PowerShell Dashboard UI Spec

## Table of Contents

1. [Environment Setup](#1-environment-setup)
2. [Design Tokens](#2-design-tokens)
3. [Unicode Characters](#3-unicode-characters)
4. [Progress Bar Function](#4-progress-bar-function)
5. [Layout Engine](#5-layout-engine)
6. [Output Sections (Top-to-Bottom)](#6-output-sections-top-to-bottom)
7. [Indentation Rules](#7-indentation-rules)
8. [Data Contract](#8-data-contract)
9. [Composability](#9-composability)
10. [Terminal Compatibility Notes](#10-terminal-compatibility-notes)
11. [Complete Rendering Example](#11-complete-rendering-example)
12. [Adapting to run.ps1 Phases](#12-adapting-to-runps1-phases)
13. [Theme Variants (Dark / Light) with Auto-Detection](#13-theme-variants-dark--light-with-auto-detection)
14. [Per-Package Coverage Table](#14-per-package-coverage-table)
15. [Coverage Comparison (Regression Detection)](#15-coverage-comparison-regression-detection)
16. [Phase Tracking Integration (TC & PC)](#16-phase-tracking-integration-tc--pc)
17. [Error-Guarding Pattern (Module Availability)](#17-error-guarding-pattern-module-availability)
18. [Modular Architecture](#18-modular-architecture)

---

## Overview

Transform standard sequential `Write-Host` console output into a structured, bordered, color-coded dashboard UI using ANSI escape sequences and Unicode box-drawing characters. This spec is self-contained — any AI or developer can implement it for any PowerShell script.

---

## 1. Environment Setup

Before any output, configure the console for UTF-8 and ANSI support:

```powershell
[console]::OutputEncoding = [System.Text.Encoding]::UTF8
$ESC = [char]27
```

All color output uses ANSI 24-bit RGB sequences, **not** `Write-Host -ForegroundColor`.

---

## 2. Design Tokens

### 2.1 Color Palette

| Token Name       | Purpose                          | RGB             | Hex       | ANSI Foreground Code              |
|------------------|----------------------------------|-----------------|-----------|-----------------------------------|
| `$cLime`         | Success, checkmarks, bars, titles | `163, 230, 53`  | `#a3e635` | `$ESC[38;2;163;230;53m`           |
| `$cRed`          | Errors, failure counts           | `244, 63, 94`   | `#f43f5e` | `$ESC[38;2;244;63;94m`            |
| `$cPurple`       | Action items, todos              | `168, 85, 247`  | `#a855f7` | `$ESC[38;2;168;85;247m`           |
| `$cCyan`         | Sub-items, info labels           | `6, 182, 212`   | `#06b6d4` | `$ESC[38;2;6;182;212m`            |
| `$cYellow`       | Warnings, phase headers          | `250, 204, 21`  | `#facc15` | `$ESC[38;2;250;204;21m`           |
| `$cMuted`        | Borders, dim text                | `156, 163, 175` | `#9ca3af` | `$ESC[38;2;156;163;175m`          |
| `$cWhite`        | Headers, scores, emphasis        | `255, 255, 255` | `#ffffff` | `$ESC[38;2;255;255;255m`          |
| `$cBarEmpty`     | Empty portion of progress bars   | `100, 100, 100` | `#646464` | `$ESC[38;2;100;100;100m`          |
| `$cReset`        | Reset all formatting             | —               | —         | `$ESC[0m`                         |

### 2.2 Text Formatting

| Style     | ANSI Code       |
|-----------|-----------------|
| Bold      | `$ESC[1m`       |
| Dim       | `$ESC[2m`       |
| Italic    | `$ESC[3m`       |
| Reset     | `$ESC[0m`       |

### 2.3 Variable Definitions (Copy-Paste Block)

```powershell
$ESC    = [char]27
$cLime  = "$ESC[38;2;163;230;53m"
$cRed   = "$ESC[38;2;244;63;94m"
$cPurple= "$ESC[38;2;168;85;247m"
$cCyan  = "$ESC[38;2;6;182;212m"
$cYellow= "$ESC[38;2;250;204;21m"
$cMuted = "$ESC[38;2;156;163;175m"
$cWhite = "$ESC[38;2;255;255;255m"
$cBarE  = "$ESC[38;2;100;100;100m"
$cReset = "$ESC[0m"
$cBold  = "$ESC[1m"
$cDim   = "$ESC[2m"
```

---

## 3. Unicode Characters

### 3.1 Icons

| Symbol | Name              | Unicode  | Usage                    |
|--------|-------------------|----------|--------------------------|
| `⚡`   | Lightning Bolt    | `U+26A1` | Product name prefix      |
| `▶`    | Triangular Bullet | `U+25B6` | Action/step prefix       |
| `✓`    | Checkmark         | `U+2713` | Success indicator        |
| `●`    | Solid Dot         | `U+25CF` | Todo/pending indicator   |
| `✗`    | Ballot X          | `U+2717` | Failure indicator        |

### 3.2 Box-Drawing Characters (Double-Line Style)

| Symbol | Name              | Unicode  |
|--------|-------------------|----------|
| `╔`    | Top-Left Corner   | `U+2554` |
| `╗`    | Top-Right Corner  | `U+2557` |
| `╚`    | Bottom-Left Corner| `U+255A` |
| `╝`    | Bottom-Right Cnr  | `U+255D` |
| `║`    | Vertical Wall     | `U+2551` |
| `═`    | Horizontal Wall   | `U+2550` |
| `╠`    | Left T-Junction   | `U+2560` |
| `╣`    | Right T-Junction  | `U+2563` |

### 3.3 Progress Bar Characters

| Symbol | Name          | Unicode  | Usage              |
|--------|---------------|----------|--------------------|
| `█`    | Full Block    | `U+2588` | Filled portion     |
| `▒`    | Medium Shade  | `U+2592` | Empty portion      |

---

## 4. Progress Bar Function

A reusable function that returns a colored progress bar string.

```powershell
function Get-ProgressBar {
    param (
        [int]$Score,
        [int]$MaxScore = 100,
        [int]$BarWidth = 15
    )

    $percentage   = $Score / $MaxScore
    $filledCount  = [math]::Round($percentage * $BarWidth)
    $emptyCount   = $BarWidth - $filledCount

    $filled = if ($filledCount -gt 0) { "█" * $filledCount } else { "" }
    $empty  = if ($emptyCount  -gt 0) { "▒" * $emptyCount  } else { "" }

    return "${cLime}${filled}${cBarE}${empty}${cReset}"
}
```

**Rules:**
- Bar width is always fixed (default 15 chars).
- Filled portion uses `$cLime`, empty uses `$cBarEmpty`.
- For "PASS"/"FAIL" labels (e.g., Browser Test), skip the bar and print the label in `$cLime` or `$cRed`.

---

## 5. Layout Engine

### 5.1 Box Width

All boxes use a fixed internal content width. Recommended: **48 characters** (50 including `║ ` and ` ║`).

```powershell
$boxWidth = 48
```

### 5.2 Helper Functions

```powershell
function Write-BoxTop {
    param([int]$Width = 48)
    Write-Host "${cMuted}╔$("═" * $Width)╗${cReset}"
}

function Write-BoxBottom {
    param([int]$Width = 48)
    Write-Host "${cMuted}╚$("═" * $Width)╝${cReset}"
}

function Write-BoxDivider {
    param([int]$Width = 48)
    Write-Host "${cMuted}╠$("═" * $Width)╣${cReset}"
}

function Write-BoxLine {
    param(
        [string]$Content,
        [int]$Width = 48
    )
    # $Content may contain ANSI codes, so visible length != string length.
    # Caller must ensure visual content fits within $Width.
    Write-Host "${cMuted}║${cReset} ${Content}"
}

function Write-BoxLineCenter {
    param(
        [string]$Text,
        [int]$Width = 48,
        [string]$Color = $cWhite
    )
    $pad = [math]::Max(0, [math]::Floor(($Width - $Text.Length) / 2))
    $line = (" " * $pad) + $Text
    Write-Host "${cMuted}║${cReset}${Color}${cBold}${line}${cReset}"
}
```

### 5.3 Column Alignment

Use `.PadRight()` and `.PadLeft()` for strict column alignment:

```powershell
# Score row example
$label     = "SEO/GEO/AEO".PadRight(16)
$scoreText = "92/100".PadLeft(7)
$bar       = Get-ProgressBar -Score 92

Write-BoxLine "$cWhite$label $scoreText  $bar"
```

**Column layout for score grid:**

| Column        | Width   | Alignment |
|---------------|---------|-----------|
| Label         | 16 char | Left      |
| Score (N/100) | 7 char  | Right     |
| Gap           | 2 char  | —         |
| Progress Bar  | 15 char | Left      |

---

## 6. Output Sections (Top-to-Bottom)

The dashboard is rendered in sequential sections. Each section is a self-contained block.

### 6.1 Header Banner

```
  ⚡  PRODUCT_NAME v1.2.0
  ────────────────────────────────
```

- Lightning bolt in `$cLime`, product name in `$cWhite` + `$cBold`.
- Horizontal rule using `─` (`U+2500`) in `$cMuted`, width matches box width.

### 6.2 Scan Summary Block (No Box)

```
  ▶ Scanning...        47 issues found
  ▶ Auto-fixing...     12 resolved ✓
  ▶ 5 agents running   SEO · Perf · Security · Quality · Browser
```

- `▶` in `$cCyan`.
- Labels in `$cCyan`.
- Issue count in `$cRed`.
- Resolved count in `$cLime` with `✓`.
- Agent names in `$cMuted`, separated by ` · `.
- Use `.PadRight()` on labels to align the right column.

### 6.3 Score Dashboard Box

```
╔══════════════════════════════════════════════════╗
║       U L T R A S H I P   S C O R E             ║
╠══════════════════════════════════════════════════╣
║                                                  ║
║  SEO/GEO/AEO    92/100  ███████████████▒▒        ║
║  Performance    87/100  █████████████▒▒▒▒        ║
║  Security       95/100  ██████████████▒▒         ║
║  Code Quality   88/100  █████████████▒▒▒         ║
║  Browser Test   PASS    ███████████████           ║
║                                                  ║
╠══════════════════════════════════════════════════╣
║                                                  ║
║  OVERALL        90/100                           ║
║  STATUS         [?] READY TO SHIP                ║
║                                                  ║
╚══════════════════════════════════════════════════╝
```

**Rendering rules:**
- Title: spaced-out letters (`S P A C E D`), centered, in `$cWhite` + `$cBold`.
- Score rows: label in `$cWhite`, score in `$cWhite`, bar from `Get-ProgressBar`.
- "PASS" keyword: rendered in `$cLime` + `$cBold`, no bar.
- "FAIL" keyword: rendered in `$cRed` + `$cBold`, no bar.
- OVERALL score: `$cWhite` + `$cBold`.
- STATUS text: `$cLime` if passing, `$cRed` if failing.
- `[?]` prefix on status in `$cYellow`.

### 6.4 Resolution Summary Block (No Box)

```
  ✓ Fixed:  12 issues auto-resolved
  ● Todo:    2 manual items remaining
```

- `✓` in `$cLime`, "Fixed:" label in `$cLime`.
- `●` in `$cYellow`, "Todo:" label in `$cYellow`.
- Counts in `$cWhite`.
- Description text in `$cMuted`.

### 6.5 Footer Tagline

```
  Ship it. One command. Production-ready.
```

- Entire line in `$cLime` + `$cBold`.

---

## 7. Indentation Rules

- All content outside boxes: **2-space indent** from left edge.
- Content inside boxes: **1 space** after `║` and before closing `║`.
- Blank lines between sections: exactly **1 blank line**.

---

## 8. Data Contract

The rendering functions should accept a data object, not hardcoded values:

```powershell
$dashboardData = @{
    ProductName = "ULTRASHIP"
    Version     = "v1.2.0"
    IssuesFound = 47
    IssuesFixed = 12
    AgentCount  = 5
    Agents      = @("SEO", "Perf", "Security", "Quality", "Browser")
    Scores      = [ordered]@{
        "SEO/GEO/AEO"   = 92
        "Performance"    = 87
        "Security"       = 95
        "Code Quality"   = 88
        "Browser Test"   = "PASS"   # string = label, int = score
    }
    OverallScore = 90
    Status       = "READY TO SHIP"
    StatusReady  = $true
    ManualTodos  = 2
}
```

**Type rules:**
- If a score value is `[int]` → render numeric score + progress bar.
- If a score value is `[string]` → render as label (PASS/FAIL) with appropriate color, no bar.

---

## 9. Composability

The spec is designed so each section is a standalone function:

```powershell
function Write-DashboardHeader  ($data) { ... }
function Write-ScanSummary      ($data) { ... }
function Write-ScoreBox         ($data) { ... }
function Write-ResolutionSummary($data) { ... }
function Write-FooterTagline    ($data) { ... }

# Main render
function Write-Dashboard ($data) {
    Write-Host ""
    Write-DashboardHeader   $data
    Write-Host ""
    Write-ScanSummary       $data
    Write-Host ""
    Write-ScoreBox          $data
    Write-Host ""
    Write-ResolutionSummary $data
    Write-Host ""
    Write-FooterTagline     $data
    Write-Host ""
}
```

---

## 10. Terminal Compatibility Notes

| Requirement          | Minimum                          |
|----------------------|----------------------------------|
| PowerShell version   | 7.0+ (pwsh) recommended          |
| Windows Terminal     | Any version (ANSI native)        |
| Legacy `conhost.exe` | May not render ANSI; add VT check|
| Encoding             | UTF-8 required                   |

**Optional VT fallback check:**

```powershell
$vtSupported = $null -ne $env:WT_SESSION -or $PSVersionTable.PSVersion.Major -ge 7
if (-not $vtSupported) {
    Write-Warning "Terminal may not support ANSI colors. Use Windows Terminal or pwsh 7+."
}
```

---

## 11. Complete Rendering Example

For reference, a full rendering call:

```powershell
# 1. Setup
[console]::OutputEncoding = [System.Text.Encoding]::UTF8
# ... define color vars, functions ...

# 2. Collect results from your script logic
$data = @{
    ProductName  = "ULTRASHIP"
    Version      = "v1.2.0"
    IssuesFound  = 47
    IssuesFixed  = 12
    AgentCount   = 5
    Agents       = @("SEO", "Perf", "Security", "Quality", "Browser")
    Scores       = [ordered]@{
        "SEO/GEO/AEO" = 92
        "Performance"  = 87
        "Security"     = 95
        "Code Quality" = 88
        "Browser Test" = "PASS"
    }
    OverallScore = 90
    Status       = "READY TO SHIP"
    StatusReady  = $true
    ManualTodos  = 2
}

# 3. Render
Write-Dashboard $data
```

This produces the exact visual layout shown in the reference image.

---

## 12. Adapting to `run.ps1` Phases

`run.ps1` has two primary dashboard-producing commands: **TC** (test-cover) and **PC** (pre-commit). Each runs through a pipeline of phases. This section maps each phase to the dashboard UI components defined above.

### 12.1 Phase Registry

Each phase is tracked in a `$phases` ordered dictionary. As each phase completes, it records its status so the final dashboard can render them all.

```powershell
$phases = [ordered]@{}

function Register-Phase {
    param(
        [string]$Name,
        [string]$Status,   # "pass", "fail", "skip", "warn"
        [string]$Detail    # optional one-line summary
    )
    $phases[$Name] = @{ Status = $Status; Detail = $Detail }
}
```

### 12.2 Phase Definitions — TC (Test Coverage) Command

| #  | Phase Name              | Source Function / Code Block              | Success Condition                         | Dashboard Label          |
|----|-------------------------|-------------------------------------------|-------------------------------------------|--------------------------|
| 1  | Git Pull                | `Invoke-GitPull`                          | No merge conflicts                        | `Git Pull`               |
| 2  | Dependency Fetch        | `Invoke-FetchLatest` → `go mod tidy`      | `$LASTEXITCODE -eq 0`                     | `Dependencies`           |
| 3  | Data Cleanup            | `Remove-Item data/` + `Cleaned data/`     | Directory removed                         | `Data Cleanup`           |
| 4  | SafeTest Boundary Lint  | `check-safetest-boundaries.ps1`           | `$LASTEXITCODE -eq 0`                     | `SafeTest Lint`          |
| 5  | Go Auto-Fixer           | `go run ./scripts/autofix/`               | `$LASTEXITCODE -eq 0`                     | `Auto-Fixer`             |
| 6  | Syntax Pre-Check        | `go run ./scripts/bracecheck/`            | `$LASTEXITCODE -eq 0`, file count logged  | `Syntax Check`           |
| 7  | Pre-Coverage Compile    | Parallel `go test -run '^$'` per package  | 0 blocked packages                        | `Compile Check`          |
| 8  | Per-File Split Recovery | Split blocked pkgs, recheck per-file      | Recovered file count                      | `Split Recovery`         |
| 9  | Coverage Run            | `go test -coverprofile` per package       | All packages produce profiles             | `Coverage Run`           |
| 10 | Coverage Merge & Report | Merge profiles, generate HTML             | Report files generated                    | `Coverage Report`        |

### 12.3 Phase Definitions — PC (Pre-Commit) Command

| #  | Phase Name              | Source Function / Code Block              | Success Condition                         | Dashboard Label          |
|----|-------------------------|-------------------------------------------|-------------------------------------------|--------------------------|
| 1  | Regression Guard        | `check-integrated-regressions.ps1`        | `$LASTEXITCODE -eq 0`                     | `Regression Guard`       |
| 2  | SafeTest Boundary Lint  | `check-safetest-boundaries.ps1`           | `$LASTEXITCODE -eq 0`                     | `SafeTest Lint`          |
| 3  | Go Auto-Fixer           | `go run ./scripts/autofix/`               | `$LASTEXITCODE -eq 0`                     | `Auto-Fixer`             |
| 4  | Syntax Pre-Check        | `go run ./scripts/bracecheck/`            | `$LASTEXITCODE -eq 0`, file count logged  | `Syntax Check`           |
| 5  | API Compile Check       | `go test -c` per Coverage* package        | 0 failures                                | `API Compile Check`      |

### 12.4 Phase Status Mapping to UI

Each phase status maps to specific colors and icons:

| Status   | Icon | Color      | Token      |
|----------|------|------------|------------|
| `pass`   | `✓`  | Lime Green | `$cLime`   |
| `fail`   | `✗`  | Red        | `$cRed`    |
| `skip`   | `⊘`  | Muted Gray | `$cMuted`  |
| `warn`   | `⚠`  | Yellow     | `$cYellow` |

### 12.5 Phase Summary Box

After all phases complete, render a summary box using the score dashboard pattern:

```
╔══════════════════════════════════════════════════╗
║         P H A S E   S U M M A R Y               ║
╠══════════════════════════════════════════════════╣
║                                                  ║
║  ✓ Git Pull            pulled 6 files            ║
║  ✓ Dependencies        up to date                ║
║  ✓ Data Cleanup        cleaned                   ║
║  ✓ SafeTest Lint       all clean                 ║
║  ✓ Auto-Fixer          no fixable issues         ║
║  ✓ Syntax Check        209 files parsed OK       ║
║  ✓ Compile Check       90/90 passed              ║
║  ⊘ Split Recovery      not needed                ║
║  ✓ Coverage Run        88 packages               ║
║  ✓ Coverage Report     generated                 ║
║                                                  ║
╠══════════════════════════════════════════════════╣
║                                                  ║
║  PHASES      10/10 passed                        ║
║  STATUS      ✓ READY TO COMMIT                   ║
║                                                  ║
╚══════════════════════════════════════════════════╝
```

**Rendering rules:**
- Status icon and label use the color from §12.4.
- Detail text uses `$cMuted`.
- Label column: `.PadRight(20)`.
- If any phase is `fail`, STATUS becomes `✗ BLOCKED` in `$cRed`.
- If any phase is `warn` but none `fail`, STATUS becomes `⚠ REVIEW` in `$cYellow`.

### 12.6 Live Phase Progress (Inline)

During execution, each phase prints a single-line status as it starts and completes:

```powershell
# On phase start:
Write-Host "  ${cCyan}▶${cReset} ${cWhite}Syntax Check${cReset}${cMuted}...${cReset}"

# On phase complete (overwrite or append):
Write-Host "  ${cLime}✓${cReset} ${cWhite}Syntax Check${cReset}  ${cMuted}209 files parsed OK${cReset}"

# On phase fail:
Write-Host "  ${cRed}✗${cReset} ${cWhite}Compile Check${cReset}  ${cRed}3 packages blocked${cReset}"
```

### 12.7 Data Contract for `run.ps1`

Extend the generic data contract (§8) with phase-specific fields:

```powershell
$dashboardData = @{
    ProductName  = "run.ps1"
    Version      = "TC"                  # or "PC"
    Command      = "test-cover"          # or "pre-commit"
    
    # Phase results (ordered)
    Phases       = [ordered]@{
        "Git Pull"         = @{ Status = "pass"; Detail = "pulled 6 files" }
        "Dependencies"     = @{ Status = "pass"; Detail = "up to date" }
        "Data Cleanup"     = @{ Status = "pass"; Detail = "cleaned" }
        "SafeTest Lint"    = @{ Status = "pass"; Detail = "all clean" }
        "Auto-Fixer"       = @{ Status = "pass"; Detail = "no fixable issues" }
        "Syntax Check"     = @{ Status = "pass"; Detail = "209 files parsed OK" }
        "Compile Check"    = @{ Status = "pass"; Detail = "90/90 passed" }
        "Split Recovery"   = @{ Status = "skip"; Detail = "not needed" }
        "Coverage Run"     = @{ Status = "pass"; Detail = "88 packages" }
        "Coverage Report"  = @{ Status = "pass"; Detail = "generated" }
    }
    
    # Score metrics (TC only — from coverage results)
    Scores       = [ordered]@{
        "Overall Coverage" = 97    # percentage from merged profile
        "Package Pass"     = 88    # packages that passed / total
        "Compile Pass"     = 90    # packages that compiled / total
        "Lint Check"       = "PASS"
        "Syntax Check"     = "PASS"
    }
    
    # Numeric rollups
    TotalPackages   = 90
    PassedPackages  = 88
    BlockedPackages = 2
    OverallCoverage = 97.3
    
    # Status
    Status       = "READY TO COMMIT"
    StatusReady  = $true
    
    # Issue tracking
    IssuesFound  = 5
    IssuesFixed  = 3
    ManualTodos  = 2
    
    # Blocked package details (for error section)
    BlockedDetails = @(
        @{ Package = "mapdiffinternal"; Errors = @("undefined: someSym") }
        @{ Package = "corepayloadtests"; Errors = @("type mismatch in Coverage20") }
    )
}
```

### 12.8 Error Detail Section

When packages are blocked, render an error detail section below the dashboard box:

```
  ── Blocked Packages ──────────────────────────────

  ✗ mapdiffinternal
      Coverage12_Gaps_test.go:45 [undefined] undefined: someSym
      Coverage12_Gaps_test.go:52 [type-mismatch] cannot use x as y

  ✗ corepayloadtests
      Coverage20_test.go:18 [undefined] undefined: stringerImpl

  ─────────────────────────────────────────────────
```

**Rendering rules:**
- Package name in `$cRed` + `$cBold`.
- Error lines: file:line in `$cYellow`, category in `$cMuted` brackets, message in `$cWhite`.
- Uses single-line box drawing `─` (`U+2500`) for dividers.
- Error categories come from `ParseCompileErrors`: `arg-count`, `undefined`, `type-mismatch`, `missing-member`, `field-vs-method`, `other`.

### 12.9 Integration Points

Since `run.ps1` uses a modular architecture, `Register-Phase` calls live inside the individual modules (not in `run.ps1` itself):

```powershell
# Example: in scripts/TestRunner.psm1 — Invoke-FetchLatest
function Invoke-FetchLatest {
    Invoke-GitPull
    Register-Phase "Git Pull" "pass" "pulled from remote"
    
    Write-Header "Fetching latest dependencies"
    go mod tidy
    if ($LASTEXITCODE -eq 0) {
        Register-Phase "Dependencies" "pass" "up to date"
    } else {
        Register-Phase "Dependencies" "fail" "go mod tidy failed"
    }
}

# At end of TC (in scripts/CoverageRunner.psm1) or PC (in scripts/PreCommitCheck.psm1):
if (Get-Command Write-PhaseSummaryBox -ErrorAction SilentlyContinue) {
    Write-Host ""
    Write-PhaseSummaryBox
}
```

**Module locations for phase-producing commands:**

| Command | Module | Function |
|---------|--------|----------|
| TC | `scripts/CoverageRunner.psm1` | `Invoke-TestCoverage` |
| TCP | `scripts/CoverageRunner.psm1` | `Invoke-PackageTestCoverage` |
| PC | `scripts/PreCommitCheck.psm1` | `Invoke-PreCommitCheck` |

### 12.10 Shared vs Command-Specific Phases

| Phase              | TC | PC | Notes                                    |
|--------------------|----|----|------------------------------------------|
| Git Pull           | ✓  | ✗  | TC pulls; PC assumes already pulled      |
| Dependencies       | ✓  | ✗  | TC runs `go mod tidy`                    |
| Data Cleanup       | ✓  | ✗  | TC cleans `data/` dir                    |
| Regression Guard   | ✗  | ✓  | PC-only CaseV1/corejson regression scan  |
| SafeTest Lint      | ✓  | ✓  | Both commands run boundary check         |
| Auto-Fixer         | ✓  | ✓  | Both, skippable via `--no-autofix`       |
| Syntax Check       | ✓  | ✓  | Both, skippable via `--skip-bracecheck`  |
| Compile Check      | ✓  | ✓  | TC: all pkgs; PC: Coverage* pkgs only    |
| Split Recovery     | ✓  | ✗  | TC-only per-file split for blocked pkgs  |
| Coverage Run       | ✓  | ✗  | TC-only actual test execution            |
| Coverage Report    | ✓  | ✗  | TC-only merge + HTML generation          |


---

## 13. Theme Variants (Dark / Light) with Auto-Detection

### 13.1 Theme Detection Strategy

PowerShell has no native API to read terminal background color. Use a layered detection approach:

```powershell
function Get-TerminalTheme {
    # Priority 1: Explicit override via environment variable
    if ($env:DASHBOARD_THEME) {
        return $env:DASHBOARD_THEME.ToLower()  # "dark" or "light"
    }

    # Priority 2: Windows Terminal settings JSON
    if ($env:WT_SESSION) {
        $wtSettings = Join-Path $env:LOCALAPPDATA "Packages\Microsoft.WindowsTerminal_8wekyb3d8bbwe\LocalState\settings.json"
        if (Test-Path $wtSettings) {
            try {
                $json = Get-Content $wtSettings -Raw | ConvertFrom-Json
                $schemeName = $json.profiles.defaults.colorScheme
                if (-not $schemeName) {
                    $schemeName = $json.profiles.list |
                        Where-Object { $_.guid -eq $json.defaultProfile } |
                        Select-Object -ExpandProperty colorScheme -ErrorAction SilentlyContinue
                }
                if ($schemeName) {
                    $scheme = $json.schemes | Where-Object { $_.name -eq $schemeName }
                    if ($scheme -and $scheme.background) {
                        $bg = $scheme.background -replace '^#', ''
                        $r = [convert]::ToInt32($bg.Substring(0,2), 16)
                        $g = [convert]::ToInt32($bg.Substring(2,2), 16)
                        $b = [convert]::ToInt32($bg.Substring(4,2), 16)
                        # Relative luminance (ITU-R BT.709)
                        $luminance = (0.2126 * $r + 0.7152 * $g + 0.0722 * $b) / 255
                        return if ($luminance -lt 0.5) { "dark" } else { "light" }
                    }
                }
            } catch { }
        }
    }

    # Priority 3: macOS/Linux — query terminal via OSC 11 (background color query)
    if ($IsLinux -or $IsMacOS) {
        try {
            $sttyOld = stty -g 2>/dev/null
            stty raw -echo min 0 time 1 2>/dev/null
            [Console]::Write("$([char]27)]11;?$([char]27)\")
            Start-Sleep -Milliseconds 100
            $response = ""
            while ([Console]::KeyAvailable) {
                $response += [char][Console]::Read()
            }
            stty $sttyOld 2>/dev/null
            # Response format: ESC]11;rgb:RRRR/GGGG/BBBB ESC\
            if ($response -match 'rgb:([0-9a-f]{2,4})/([0-9a-f]{2,4})/([0-9a-f]{2,4})') {
                $r = [convert]::ToInt32($Matches[1].Substring(0,2), 16)
                $g = [convert]::ToInt32($Matches[2].Substring(0,2), 16)
                $b = [convert]::ToInt32($Matches[3].Substring(0,2), 16)
                $luminance = (0.2126 * $r + 0.7152 * $g + 0.0722 * $b) / 255
                return if ($luminance -lt 0.5) { "dark" } else { "light" }
            }
        } catch { }
    }

    # Priority 4: PowerShell host background heuristic
    try {
        $bg = $Host.UI.RawUI.BackgroundColor
        $lightBgs = @("White", "Gray", "Yellow", "Cyan")
        if ($bg -in $lightBgs) { return "light" }
    } catch { }

    # Default: dark (most developer terminals are dark)
    return "dark"
}
```

**Detection priority order:**
1. `$env:DASHBOARD_THEME` — explicit user override (`dark` or `light`)
2. Windows Terminal `settings.json` — parse active color scheme background
3. OSC 11 query — terminal-standard background color query (macOS/Linux)
4. `$Host.UI.RawUI.BackgroundColor` — PowerShell host fallback
5. Default to `dark`

### 13.2 Color Palettes by Theme

#### Dark Theme (default — designed for dark terminal backgrounds)

This is the palette defined in §2.1. No changes needed.

| Token        | RGB             | Hex       | Rationale                         |
|--------------|-----------------|-----------|-----------------------------------|
| `$cLime`     | `163, 230, 53`  | `#a3e635` | High contrast on dark bg          |
| `$cRed`      | `244, 63, 94`   | `#f43f5e` | Bright, urgent on dark bg         |
| `$cPurple`   | `168, 85, 247`  | `#a855f7` | Vibrant on dark bg                |
| `$cCyan`     | `6, 182, 212`   | `#06b6d4` | Clear info tone                   |
| `$cYellow`   | `250, 204, 21`  | `#facc15` | Bright warning                    |
| `$cMuted`    | `156, 163, 175` | `#9ca3af` | Subtle but readable on dark       |
| `$cWhite`    | `255, 255, 255` | `#ffffff` | Maximum contrast on dark          |
| `$cBarEmpty` | `100, 100, 100` | `#646464` | Dim but visible on dark           |
| `$cBorder`   | `156, 163, 175` | `#9ca3af` | Same as muted (box-drawing chars) |

#### Light Theme (designed for white/light terminal backgrounds)

| Token        | RGB             | Hex       | Rationale                         |
|--------------|-----------------|-----------|-----------------------------------|
| `$cLime`     | `22, 163, 74`   | `#16a34a` | Darker green, readable on white   |
| `$cRed`      | `185, 28, 28`   | `#b91c1c` | Deep red, not washed out          |
| `$cPurple`   | `109, 40, 217`  | `#6d28d9` | Saturated purple on light bg      |
| `$cCyan`     | `14, 116, 144`  | `#0e7490` | Teal, strong on white             |
| `$cYellow`   | `161, 98, 7`    | `#a16207` | Amber/brown — yellow is invisible on white |
| `$cMuted`    | `107, 114, 128` | `#6b7280` | Medium gray, visible on white     |
| `$cWhite`    | `15, 23, 42`    | `#0f172a` | Near-black for text on light bg   |
| `$cBarEmpty` | `209, 213, 219` | `#d1d5db` | Light gray, subtle on white       |
| `$cBorder`   | `156, 163, 175` | `#9ca3af` | Medium gray borders               |

**Key differences:**
- `$cWhite` becomes near-black (it's the "primary text" token, not literal white)
- `$cYellow` shifts to amber — pure yellow is invisible on white backgrounds
- `$cLime` darkens significantly — neon green washes out on light
- `$cBarEmpty` lightens — dark empty blocks are too prominent on light backgrounds

### 13.3 Theme Initialization (Copy-Paste Block)

```powershell
$ESC    = [char]27
$cReset = "$ESC[0m"
$cBold  = "$ESC[1m"
$cDim   = "$ESC[2m"

$global:DashboardTheme = Get-TerminalTheme

function Set-ThemeColors {
    param([string]$Theme)

    if ($Theme -eq "light") {
        $script:cLime   = "$ESC[38;2;22;163;74m"
        $script:cRed    = "$ESC[38;2;185;28;28m"
        $script:cPurple = "$ESC[38;2;109;40;217m"
        $script:cCyan   = "$ESC[38;2;14;116;144m"
        $script:cYellow = "$ESC[38;2;161;98;7m"
        $script:cMuted  = "$ESC[38;2;107;114;128m"
        $script:cWhite  = "$ESC[38;2;15;23;42m"
        $script:cBarE   = "$ESC[38;2;209;213;219m"
        $script:cBorder = "$ESC[38;2;156;163;175m"
    } else {
        $script:cLime   = "$ESC[38;2;163;230;53m"
        $script:cRed    = "$ESC[38;2;244;63;94m"
        $script:cPurple = "$ESC[38;2;168;85;247m"
        $script:cCyan   = "$ESC[38;2;6;182;212m"
        $script:cYellow = "$ESC[38;2;250;204;21m"
        $script:cMuted  = "$ESC[38;2;156;163;175m"
        $script:cWhite  = "$ESC[38;2;255;255;255m"
        $script:cBarE   = "$ESC[38;2;100;100;100m"
        $script:cBorder = "$ESC[38;2;156;163;175m"
    }
}

Set-ThemeColors $global:DashboardTheme
```

### 13.4 Overriding the Theme

Users can force a theme via environment variable before running:

```powershell
# Force light theme
$env:DASHBOARD_THEME = "light"
./run.ps1 TC

# Force dark theme
$env:DASHBOARD_THEME = "dark"
./run.ps1 PC

# Or inline (PowerShell 7+)
$env:DASHBOARD_THEME = "light"; ./run.ps1 TC
```

Or via a `--theme` flag if added to `run.ps1`:

```powershell
# In run.ps1 arg parsing
if ($ExtraArgs -contains '--light') { $env:DASHBOARD_THEME = "light" }
if ($ExtraArgs -contains '--dark')  { $env:DASHBOARD_THEME = "dark" }
```

### 13.5 Component Rendering Rules by Theme

All rendering functions (§5, §6, §12) use the `$cXxx` variables — they are **never hardcoded** to specific RGB values. This means all components automatically adapt when `Set-ThemeColors` is called.

Special per-theme rules:

| Component            | Dark Theme                    | Light Theme                      |
|----------------------|-------------------------------|----------------------------------|
| Box borders (`║═`)   | `$cBorder` (gray)             | `$cBorder` (gray) — same        |
| Progress bar filled  | `$cLime` (neon green `█`)     | `$cLime` (forest green `█`)     |
| Progress bar empty   | `$cBarE` (dark gray `▒`)      | `$cBarE` (light gray `▒`)       |
| Primary text         | `$cWhite` (white)             | `$cWhite` (near-black)          |
| Error text           | `$cRed` (bright pink-red)     | `$cRed` (deep red)              |
| Warning labels       | `$cYellow` (bright yellow)    | `$cYellow` (amber)              |
| Footer tagline       | `$cLime` + `$cBold`           | `$cLime` + `$cBold`             |

### 13.6 Testing Themes

To visually verify both themes, add a test function:

```powershell
function Test-DashboardTheme {
    foreach ($theme in @("dark", "light")) {
        Set-ThemeColors $theme
        Write-Host ""
        Write-Host "${cBold}=== Theme: $theme ===${cReset}"
        Write-Host "  ${cLime}✓ Success / Lime${cReset}"
        Write-Host "  ${cRed}✗ Error / Red${cReset}"
        Write-Host "  ${cPurple}● Purple / Todo${cReset}"
        Write-Host "  ${cCyan}▶ Cyan / Info${cReset}"
        Write-Host "  ${cYellow}⚠ Yellow / Warning${cReset}"
        Write-Host "  ${cMuted}Muted text${cReset}"
        Write-Host "  ${cWhite}Primary text${cReset}"
        Write-Host "  Bar: $(Get-ProgressBar -Score 73)"
        Write-Host ""
    }
    # Restore actual theme
    Set-ThemeColors $global:DashboardTheme
}
```

### 13.7 Luminance Formula Reference

Used in detection (§13.1) and available for custom threshold tuning:

```
Relative Luminance (ITU-R BT.709):
  L = (0.2126 × R + 0.7152 × G + 0.0722 × B) / 255

  L < 0.5 → dark background → use dark theme
  L ≥ 0.5 → light background → use light theme
```

Threshold can be adjusted. `0.5` works for most terminals. Lower values (e.g., `0.4`) bias toward dark detection.

---

## 14. Per-Package Coverage Table

### 14.1 Purpose

Render a bordered, color-coded table of per-package coverage results at the end of `TC` runs. Packages are sorted ascending by coverage (lowest first) to highlight gaps.

### 14.2 Function Signature

```powershell
function Write-CoverageTable {
    param(
        [Parameter(Mandatory)]
        [hashtable[]]$CoverageData    # Each: @{ Package="..."; Coverage=95.2; Tests=12 }
    )
}
```

### 14.3 Color Thresholds

| Coverage       | Color     | Token      |
|----------------|-----------|------------|
| ≥ 100%         | Lime      | `$cLime`   |
| ≥ 98%          | White     | `$cWhite`  |
| ≥ 95%          | Yellow    | `$cYellow` |
| < 95%          | Red       | `$cRed`    |

### 14.4 Progress Bar

Each row includes a 20-character progress bar using `█` (filled) and `░` (empty), colored per the threshold above.

### 14.5 Summary Footer

After the table, display:
- Average coverage across all packages
- Count of packages at 100% vs. below
- Target threshold line (default: 100%)

### 14.6 Integration

Called automatically by `Write-Dashboard` when `$Data.CoverageData` is populated.

---

## 15. Coverage Comparison (Regression Detection)

### 15.1 Purpose

Compare current coverage results against a previously saved snapshot to detect regressions, improvements, new packages, and removed packages. This enables run-over-run tracking without external tools.

### 15.2 Functions

#### `Write-CoverageComparison`

Renders a diff table between current and previous coverage arrays.

```powershell
function Write-CoverageComparison {
    param(
        [Parameter(Mandatory)]
        [hashtable[]]$Current,        # Current run coverage data

        [Parameter()]
        [hashtable[]]$Previous = @()  # Previous run snapshot (empty = first run)
    )
}
```

**Diff indicators:**

| Symbol | Meaning          | Color     |
|--------|------------------|-----------|
| `▲`    | Coverage improved | Lime      |
| `▼`    | Coverage regressed | Red      |
| `★`    | New package       | Cyan      |
| `✗`    | Removed package   | Yellow    |
| `=`    | No change         | Muted     |

**Special flags:**
- Packages dropping **from** 100% are flagged `⚠ LOST 100%` in red
- Packages reaching 100% are flagged `🎯 REACHED 100%` in lime

**Sort order:** Regressions first (largest drop), then improvements, then unchanged.

#### `Save-CoverageSnapshot`

Persists current coverage data to a JSON file for future comparison.

```powershell
function Save-CoverageSnapshot {
    param(
        [Parameter(Mandatory)]
        [hashtable[]]$CoverageData,

        [string]$Path = "data/coverage/coverage-previous.json"
    )
}
```

Creates the directory if missing. Writes an array of `{ Package, Coverage, Tests }` objects.

#### `Load-CoverageSnapshot`

Loads the last saved snapshot for comparison.

```powershell
function Load-CoverageSnapshot {
    param(
        [string]$Path = "data/coverage/coverage-previous.json"
    )
}
```

Returns `@()` if the file does not exist (first run).

### 15.3 Integration in TC and PC

Both `Invoke-TestCoverage` (in `scripts/CoverageRunner.psm1`) and `Invoke-PackageTestCoverage` (in `scripts/CoverageRunner.psm1`) wire up the comparison flow after coverage data is collected:

```powershell
if (Get-Command Write-CoverageComparison -ErrorAction SilentlyContinue) {
    $previousCovData = $null
    if (Get-Command Load-CoverageSnapshot -ErrorAction SilentlyContinue) {
        $previousCovData = Load-CoverageSnapshot
    }
    Write-CoverageComparison -Current $currentCovData -Previous $previousCovData
    if (Get-Command Save-CoverageSnapshot -ErrorAction SilentlyContinue) {
        Save-CoverageSnapshot -CoverageData $currentCovData
    }
}
```

All calls are guarded per the error-guarding pattern (§17).

- **TC** (`Invoke-TestCoverage` in `scripts/CoverageRunner.psm1`): builds `$currentCovData` from the `$srcPkgStmts` hashtable (statement-level aggregation).
- **TCP** (`Invoke-PackageTestCoverage` in `scripts/CoverageRunner.psm1`): aggregates per-source-package coverage from `go tool cover -func` output lines.
- **PC** (`Invoke-PreCommitCheck` in `scripts/PreCommitCheck.psm1`): does not produce coverage data.

> **Console output spec**: The rendered diff table is documented as **Section 4: Coverage Diff** in
> [`spec/03-powershell-test-run/07-tc-console-output.md`](../03-powershell-test-run/07-tc-console-output.md#section-4-coverage-diff).

This gives automatic regression detection on every `./run.ps1 TC` or `./run.ps1 TCP <pkg>` run.

### 15.4 Summary Footer

The comparison table ends with a summary line:

```
  ▲ 3 improved  ▼ 1 regressed  ★ 2 new  ✗ 0 removed  = 12 unchanged
```

---

## 16. Phase Tracking Integration (TC & PC)

Both `Invoke-TestCoverage` (in `scripts/CoverageRunner.psm1`) and `Invoke-PreCommitCheck` (in `scripts/PreCommitCheck.psm1`) integrate the phase-tracking system (`Register-Phase` / `Write-PhaseSummaryBox`) to render a bordered execution summary at the end of each run.

### 16.1 Lifecycle

Each command follows the same pattern:

```powershell
# At the start of the command
if (Get-Command Reset-Phases -ErrorAction SilentlyContinue) { Reset-Phases }

# After each logical phase completes
if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
    Register-Phase "<PhaseName>" "<status>" "<detail>"
}

# At the very end of the command
if (Get-Command Write-PhaseSummaryBox -ErrorAction SilentlyContinue) {
    Write-Host ""
    Write-PhaseSummaryBox
}
```

All calls are guarded with `Get-Command ... -ErrorAction SilentlyContinue` so `run.ps1` remains functional even if `DashboardUI.psm1` is not loaded.

### 16.2 TC Phases

`Invoke-TestCoverage` registers **10 phases** in execution order:

| #  | Phase Name       | Possible Statuses        | Detail Examples                              |
|----|------------------|--------------------------|----------------------------------------------|
| 1  | Git Pull         | `pass`                   | `pulled from remote`                         |
| 2  | Dependencies     | `pass`                   | `up to date`                                 |
| 3  | Data Cleanup     | `pass`                   | `cleaned`                                    |
| 4  | SafeTest Lint    | `pass`, `fail`           | `all clean` / `boundary check failed`        |
| 5  | Auto-Fixer       | `pass`, `warn`, `skip`   | `no fixable issues` / `errors encountered` / `skipped (--no-autofix)` |
| 6  | Syntax Check     | `pass`, `fail`, `skip`   | bracecheck output / `skipped (--skip-bracecheck)` |
| 7  | Compile Check    | `pass`, `warn`           | `42/42 passed` / `40/42 passed, 2 blocked`   |
| 8  | Split Recovery   | `pass`, `skip`           | `3 subfolders recovered` / `not needed`      |
| 9  | Coverage Run     | `pass`, `fail`           | `42 packages` / `no packages to run`         |
| 10 | Coverage Report  | `pass`                   | coverage summary stats                       |

### 16.3 PC Phases

`Invoke-PreCommitCheck` registers **5 phases** in execution order:

| #  | Phase Name        | Possible Statuses      | Detail Examples                              |
|----|-------------------|------------------------|----------------------------------------------|
| 1  | Regression Guard  | `pass`, `fail`         | `no regressions` / `regressions detected`    |
| 2  | SafeTest Lint     | `pass`, `fail`         | `all clean` / `boundary check failed`        |
| 3  | Auto-Fixer        | `pass`, `warn`, `skip` | `no fixable issues` / `skipped (--no-autofix)` |
| 4  | Syntax Check      | `pass`, `fail`, `skip` | bracecheck output / `skipped (--skip-bracecheck)` |
| 5  | API Compile Check | `pass`, `fail`         | `12/12 passed` / `3 failed, 9 passed`        |

### 16.4 Status-to-Color Mapping

Phase statuses map to design tokens as defined in §8:

| Status | Color    | Icon |
|--------|----------|------|
| `pass` | `$cLime` | `✓`  |
| `fail` | `$cRed`  | `✗`  |
| `warn` | `$cYellow` | `⚠` |
| `skip` | `$cMuted`  | `○` |

### 16.5 Example Output

```
╔══════════════════════════════════════════════════════╗
║  Phase Summary                                       ║
╠══════════════════════════════════════════════════════╣
║  ✓ Git Pull .............. pulled from remote        ║
║  ✓ Dependencies .......... up to date                ║
║  ✓ Data Cleanup .......... cleaned                   ║
║  ✓ SafeTest Lint ......... all clean                 ║
║  ✓ Auto-Fixer ............ no fixable issues         ║
║  ✓ Syntax Check .......... all clear                 ║
║  ✓ Compile Check ......... 42/42 passed              ║
║  ○ Split Recovery ........ not needed                ║
║  ✓ Coverage Run .......... 42 packages               ║
║  ✓ Coverage Report ....... 98.4% average             ║
╚══════════════════════════════════════════════════════╝
```

---

## 17. Error-Guarding Pattern (Module Availability)

All modules in `scripts/` are designed to function correctly even when `DashboardUI.psm1` is not loaded. Every call to a DashboardUI function is wrapped in a guard that silently skips the call if the function is unavailable. The `run.ps1` dispatcher imports all modules but none are mandatory — missing modules simply disable the commands that depend on them.

### 17.1 Guard Pattern

```powershell
if (Get-Command <FunctionName> -ErrorAction SilentlyContinue) {
    <FunctionName> [arguments]
}
```

`Get-Command` with `-ErrorAction SilentlyContinue` returns `$null` when the function is not defined, causing the `if` block to be skipped without throwing an error or producing output.

### 17.2 Guarded Functions

All exported `DashboardUI.psm1` functions used in `run.ps1` are guarded:

| Function                  | Usage Context                     |
|---------------------------|-----------------------------------|
| `Reset-Phases`            | Start of TC / PC commands         |
| `Register-Phase`          | After each phase boundary         |
| `Write-PhaseSummaryBox`   | End of TC / PC commands           |
| `Write-CoverageTable`     | After coverage data is collected  |
| `Write-CoverageComparison`| After coverage snapshot comparison|
| `Load-CoverageSnapshot`   | Before coverage comparison        |
| `Save-CoverageSnapshot`   | After coverage comparison         |

### 17.3 Why Not Try/Catch

The `Get-Command` guard is preferred over `try/catch` because:

1. **No exception overhead** — avoids the cost of throwing and catching `CommandNotFoundException`
2. **No noise** — `try/catch` with `-ErrorAction Stop` can still emit partial error records in some PowerShell hosts
3. **Readable intent** — the guard clearly communicates "run only if available" rather than "run and recover from failure"
4. **Granular control** — each call site independently decides whether to skip, unlike a module-level try/catch that would be all-or-nothing

### 17.4 Module Import Guard

The module import itself in `run.ps1` also uses a safe pattern:

```powershell
$dashboardModule = Join-Path $PSScriptRoot "scripts" "DashboardUI.psm1"
if (Test-Path $dashboardModule) {
    Import-Module $dashboardModule -Force -ErrorAction SilentlyContinue
}
```

This ensures the script does not fail at startup if the module file is missing or contains syntax errors.

### 17.5 Design Principle

> **DashboardUI is always additive, never required.** All core functionality (testing, coverage, compilation) must work identically with or without the module. The dashboard layer provides enhanced visual feedback but never gates execution.

---

## 18. Modular Architecture

As of the Task 10 refactor, `run.ps1` is a thin dispatcher (≤200 lines) that imports specialized `.psm1` modules from the `scripts/` directory. DashboardUI is one of these modules.

### 18.1 Module Map

| Module | Responsibility | DashboardUI Dependency |
|--------|---------------|----------------------|
| `DashboardUI.psm1` | ANSI rendering, phase tracking, coverage tables | — (this IS the module) |
| `Utilities.psm1` | Console helpers, error extraction | Optional (graceful fallback) |
| `TestLogWriter.psm1` | Go test output → log files | None |
| `TestRunner.psm1` | Test execution, build checks, git ops | None |
| `CoverageRunner.psm1` | TC + TCP coverage pipelines | Yes (phase tracking, coverage tables, diff) |
| `BuildTools.psm1` | Build, format, vet, tidy, clean | None |
| `PreCommitCheck.psm1` | PC pre-commit validation | Yes (phase tracking, summary box) |
| `GoConvey.psm1` | GoConvey launcher | None |
| `Help.psm1` | Help display, fail log, integrated tests | None |

### 18.2 How DashboardUI Is Consumed

Modules that use DashboardUI functions (primarily `CoverageRunner` and `PreCommitCheck`) guard every call:

```powershell
if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
    Register-Phase "Compile Check" "pass" "42/42 passed"
}
```

This means the dashboard UI layer can be removed, replaced, or broken without affecting test/coverage functionality.

### 18.3 Reference

Full module documentation: [`scripts/README.md`](../../scripts/README.md)
Refactoring roadmap: [`.lovable/memory/workflow/06-powershell-refactor-plan.md`](../../.lovable/memory/workflow/06-powershell-refactor-plan.md)
