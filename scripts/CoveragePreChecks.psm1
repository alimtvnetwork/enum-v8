# ─────────────────────────────────────────────────────────────────────────────
# CoveragePreChecks.psm1 — Pre-coverage validation (safetest, autofix, bracecheck)
#
# Usage:
#   Import-Module ./scripts/CoveragePreChecks.psm1 -Force
#
# Dependencies: Utilities.psm1 (Write-Header, Write-Success, Write-Fail)
#               DashboardUI.psm1 (Register-Phase — optional)
# ─────────────────────────────────────────────────────────────────────────────

function Invoke-CoveragePreChecks {
    <#
    .SYNOPSIS
        Run all pre-coverage checks: safetest boundaries, Go auto-fixer, bracecheck.
    .DESCRIPTION
        Returns $true if all checks pass, $false if any critical check fails.
        Registers dashboard phases for each check.
    .PARAMETER ScriptRoot
        The root directory of the project (typically $PSScriptRoot from caller).
    .PARAMETER ExtraArgs
        Optional extra arguments (--no-autofix, --skip-bracecheck, --dry-run).
    .PARAMETER CoverDir
        Directory for coverage output files (for syntax-issues.txt).
    #>
    [CmdletBinding()]
    param(
        [string]$ScriptRoot,
        [string[]]$ExtraArgs,
        [string]$CoverDir
    )

    # ── In-package import lint check ──────────────────────────────────
    $inpkgScript = Join-Path $ScriptRoot "scripts" "check-inpkg-imports.ps1"
    if (Test-Path $inpkgScript) {
        Write-Host ""
        Write-Host "  Running in-package import lint check..." -ForegroundColor Yellow
        & $inpkgScript
        if ($LASTEXITCODE -ne 0) {
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "InPkg Import Lint" "fail" "forbidden imports found" }
            $s = Get-CallerSource; Write-Fail "In-package import check failed. Move heavy imports to tests/integratedtests/. (source: $s)"
            return $false
        }
    }
    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "InPkg Import Lint" "pass" "all clean" }

    # ── safeTest boundary + empty-if lint check ────────────────────
    $boundaryScript = Join-Path $ScriptRoot "scripts" "check-safetest-boundaries.ps1"
    if (Test-Path $boundaryScript) {
        Write-Host ""
        Write-Host "  Running safeTest boundary + empty-if lint check..." -ForegroundColor Yellow
        & $boundaryScript
        if ($LASTEXITCODE -ne 0) {
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "SafeTest Lint" "fail" "boundary check failed" }
            $s = Get-CallerSource; Write-Fail "safeTest boundary check failed. Fix reported issues before TC. (source: $s)"
            return $false
        }
    }
    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "SafeTest Lint" "pass" "all clean" }

    # ── Go auto-fixer ─────────────────────────────────────────────────
    $skipAutofix = $ExtraArgs -and ($ExtraArgs -contains '--no-autofix')
    $skipBrace = $ExtraArgs -and ($ExtraArgs -contains '--skip-bracecheck')
    if ($skipBrace) {
        Write-Host "  Skipping Go auto-fixer and syntax pre-check (--skip-bracecheck)" -ForegroundColor DarkYellow
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Auto-Fixer" "skip" "skipped (--skip-bracecheck)" }
    } elseif ($skipAutofix) {
        Write-Host "  Skipping Go auto-fixer (--no-autofix)" -ForegroundColor DarkYellow
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Auto-Fixer" "skip" "skipped (--no-autofix)" }
    } else {
        $dryRunFlag = if ($ExtraArgs -and ($ExtraArgs -contains '--dry-run')) { '--dry-run' } else { $null }
        $dryLabel = if ($dryRunFlag) { " (dry-run)" } else { "" }
        Write-Host "  Running Go auto-fixer$dryLabel..." -ForegroundColor Yellow
        $fixArgs = @('./scripts/autofix/')
        if ($dryRunFlag) { $fixArgs += '--dry-run' }
        $fixOut = & go run @fixArgs 2>&1
        if ($LASTEXITCODE -ne 0) {
            Write-Host ($fixOut | Out-String) -ForegroundColor Red
            $s = Get-CallerSource; Write-Fail "Go auto-fixer encountered errors. (source: $s)"
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Auto-Fixer" "warn" "errors encountered" }
        } else {
            $fixStr = ($fixOut | Out-String).Trim() -replace '^\s*✓\s*', ''
            if ($fixStr) { Write-Success $fixStr }
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Auto-Fixer" "pass" "no fixable issues" }
        }
    }

    # ── Go syntax pre-check (bracecheck) ──────────────────────────────
    if ($skipBrace) {
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Syntax Check" "skip" "skipped (--skip-bracecheck)" }
    } else {
        Write-Host "  Running Go syntax pre-check (bracecheck)..." -ForegroundColor Yellow
        $braceOut = & go run ./scripts/bracecheck/ 2>&1
        if ($LASTEXITCODE -ne 0) {
            Write-Host ($braceOut | Out-String) -ForegroundColor Red
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Syntax Check" "fail" "bracecheck failed" }
            $s = Get-CallerSource; Write-Fail "Go syntax check failed. Fix reported issues before TC. (source: $s)"
            return $false
        } else {
            $braceStr2 = ($braceOut | Out-String).Trim() -replace '^\s*✓\s*', ''
            Write-Success $braceStr2
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Syntax Check" "pass" $braceStr2 }
        }

        # ── Write syntax-issues.txt report ────────────────────────────
        $syntaxReportDir = $CoverDir
        New-Item -ItemType Directory -Path $syntaxReportDir -Force | Out-Null
        $syntaxReportFile = Join-Path $syntaxReportDir "syntax-issues.txt"
        $braceStr = ($braceOut | Out-String).Trim()
        $syntaxTs = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
        $syntaxContent = @(
            "================================================================================"
            "  Syntax Issues Report — $syntaxTs"
            "  Generated by: autofix + bracecheck pipeline"
            "================================================================================"
            ""
        )
        if (Test-Path $syntaxReportFile) {
            $existing = Get-Content $syntaxReportFile -Raw
            $syntaxContent = @($existing.TrimEnd(), "", "")
        }
        $syntaxContent += @(
            "────────────────────────────────────────────────────────────────────────────────"
            " BRACECHECK RESULTS"
            "────────────────────────────────────────────────────────────────────────────────"
            ""
            "  $braceStr"
            ""
            "================================================================================"
        )
        Set-Content -Path $syntaxReportFile -Value ($syntaxContent -join "`n") -Encoding UTF8
    }

    return $true
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Invoke-CoveragePreChecks'
)
