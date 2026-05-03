#!/usr/bin/env pwsh
<#
.SYNOPSIS
    Project runner script with shorthands for common operations.

.DESCRIPTION
    Usage: ./run.ps1 <command> [options]

    Commands (uppercase shorthands OR hyphen-lowercase):
        T   | -t   | test          Run all tests (verbose)
        TP  | -tp  | test-pkg      Run tests for a specific package: ./run.ps1 TP regexnewtests
        TC  | -tc  | test-cover    Run tests with coverage (parallel by default)
        TCP | -tcp | test-cover-pkg Run coverage for a specific package: ./run.ps1 TCP regexnewtests
        TI  | -ti  | test-int      Run integrated tests only
        TF  | -tf  | test-fail     Show last failing tests log
        GC  | -gc  | goconvey      Launch GoConvey (browser test runner)
        R   | -r   | run           Run the main application
        B   | -b   | build         Build the binary
        BR  | -br  | build-run     Build then run
        F   | -f   | fmt           Format all Go files
        L   | -l   | lint          Run go vet on all packages
        V   | -v   | vet           Run go vet
        TY  | -ty  | tidy          Run go mod tidy
        PC  | -pc  | pre-commit    Check Coverage* files for API mismatches
        C   | -c   | clean         Clean build artifacts
        H   | -h   | help          Show this help

    Mode options (for TC/TCP):
        --sync             Run precompile + tests sequentially (default: parallel)
        --open             Open HTML coverage report in browser (default: don't open)
        --skip-bracecheck  Skip the Go syntax pre-check for faster runs
        --no-autofix       Skip the Go auto-fixer before bracecheck
        --dry-run          Run auto-fixer in preview mode (show fixes without applying)

.EXAMPLE
    ./run.ps1 T
    ./run.ps1 -t
    ./run.ps1 TP regexnewtests
    ./run.ps1 -tp regexnewtests
    ./run.ps1 TC --sync
    ./run.ps1 -gc
#>

param(
    [Parameter(Position = 0)]
    [string]$Command = "help",

    [Parameter(Position = 1, ValueFromRemainingArguments)]
    [string[]]$ExtraArgs
)

# Normalize: if $Command was swallowed by PowerShell as a switch
# (e.g. -gc parsed away), $Command will be "help" — detect via $PSBoundParameters.
if (-not $PSBoundParameters.ContainsKey('Command')) {
    # Check $MyInvocation.Line for the actual argument
    $rawLine = $MyInvocation.Line
    $match = [regex]::Match($rawLine, '(?i)run\.ps1\s+(-?\w[\w-]*)\s*(.*)')
    if ($match.Success) {
        $Command = $match.Groups[1].Value
        # Capture remaining args that PowerShell swallowed
        $trailing = $match.Groups[2].Value.Trim()
        if ($trailing -and (-not $ExtraArgs -or $ExtraArgs.Count -eq 0)) {
            $ExtraArgs = @($trailing -split '\s+')
        }
    }
}

# -- Encoding (fix garbled Unicode on Windows terminals) --
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8
$OutputEncoding            = [System.Text.Encoding]::UTF8

$ErrorActionPreference = "Stop"

# -- Module Loading --
# Load all modules in dependency order via a single loop
$moduleOrder = @(
    "DashboardTheme", "DashboardBoxPrimitives", "DashboardSections", "DashboardPhases",
    "DashboardCoverageTable", "DashboardCoverageDiff", "DashboardUI",
    "Utilities", "ErrorExtractor", "ErrorParser", "TestLogWriter",
    "TestRunnerCore", "TestRunner",
    "CoveragePreChecks", "CoverageCompileCheck", "CoverageSplitRecovery",
    "CoverageProfileMerger", "CoverageReportTxt", "CoverageReportJson", "CoverageReportHtml",
    "PackageCoverage", "CoverageRunner",
    "BuildTools", "GoConvey", "PreCommitCheck", "Help"
)
foreach ($mod in $moduleOrder) {
    $modPath = Join-Path $PSScriptRoot "scripts" "$mod.psm1"
    if (Test-Path $modPath) { Import-Module $modPath -Force -DisableNameChecking }
}

# Auto-detect theme, or honor --light / --dark flags
if (Get-Command Initialize-DashboardUI -ErrorAction SilentlyContinue) {
    $themeOverride = $null
    if ($ExtraArgs -contains '--light') { $themeOverride = "light" }
    if ($ExtraArgs -contains '--dark')  { $themeOverride = "dark" }
    if ($themeOverride) { Initialize-DashboardUI -Theme $themeOverride }
    else { Initialize-DashboardUI }
}

# -- Shared Variables --
$global:ProjectRoot = $PSScriptRoot
$global:DataDir     = Join-Path $PSScriptRoot "data"
$global:TestLogDir  = Join-Path $global:DataDir "test-logs"

# ═══════════════════════════════════════════════════════════════════════════════
# Command Dispatch
# ═══════════════════════════════════════════════════════════════════════════════

$firstExtraArg = if ($ExtraArgs -and $ExtraArgs.Count -gt 0) { $ExtraArgs[0] } else { $null }

switch ($Command.ToLower()) {
    { $_ -in "t", "-t", "test" }              { Invoke-AllTests }
    { $_ -in "tp", "-tp", "test-pkg" }        { Invoke-PackageTests $firstExtraArg }
    { $_ -in "tc", "-tc", "test-cover" }      { Invoke-TestCoverage }
    { $_ -in "tcp", "-tcp", "test-cover-pkg" } { Invoke-PackageTestCoverage $firstExtraArg }
    { $_ -in "ti", "-ti", "test-int" }        { Invoke-IntegratedTests }
    { $_ -in "tf", "-tf", "test-fail" }       { Invoke-ShowFailLog }
    { $_ -in "gc", "-gc", "goconvey" }        { Invoke-GoConvey -ExtraArgs $ExtraArgs }
    { $_ -in "r", "-r", "run" }               { Invoke-RunMain }
    { $_ -in "b", "-b", "build" }             { Invoke-Build }
    { $_ -in "br", "-br", "build-run" }       { Invoke-BuildRun }
    { $_ -in "f", "-f", "fmt" }               { Invoke-Format }
    { $_ -in "l", "-l", "lint", "v", "-v", "vet" } { Invoke-Vet }
    { $_ -in "ty", "-ty", "tidy" }            { Invoke-Tidy }
    { $_ -in "pc", "-pc", "pre-commit" }      { Invoke-PreCommitCheck $firstExtraArg }
    { $_ -in "c", "-c", "clean" }             { Invoke-Clean }
    { $_ -in "h", "-h", "help", "" }          { Show-Help }
    default {
        Write-Fail "Unknown command: '$Command'"
        Show-Help
    }
}
