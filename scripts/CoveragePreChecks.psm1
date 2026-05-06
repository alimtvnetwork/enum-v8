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

    # ── Spec-API fabrication lint (S-106) ─────────────────────────────
    # Soft gate: warn-only by default. Skipped if upstream clone or spec dir
    # is absent (typical on contributor machines). Use --strict-spec-api to
    # fail the run on any fabrication.
    $skipSpecApi = $ExtraArgs -and ($ExtraArgs -contains '--no-spec-api')
    if ($skipSpecApi) {
        Write-Host "  Skipping spec-API fabrication lint (--no-spec-api)" -ForegroundColor DarkYellow
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Spec-API Lint" "skip" "skipped (--no-spec-api)" }
    } else {
        $specApiModule = Join-Path $ScriptRoot "scripts" "spec-api-check.psm1"
        $specDir       = Join-Path $ScriptRoot "spec" "01-app"
        $upstreamDir   = "/tmp/core-v9-upstream"
        if (-not (Test-Path $specApiModule)) {
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Spec-API Lint" "skip" "spec-api-check.psm1 missing" }
        } elseif (-not (Test-Path $specDir)) {
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Spec-API Lint" "skip" "spec/01-app/ missing" }
        } else {
            # S-115: import early so Test-UpstreamClone is available for the sentinel check.
            try { Import-Module $specApiModule -Force -DisableNameChecking } catch {}
            $cloneStatus = $null
            if (Get-Command Test-UpstreamClone -ErrorAction SilentlyContinue) {
                $cloneStatus = Test-UpstreamClone -Path $upstreamDir
            } elseif (Test-Path $upstreamDir) {
                $cloneStatus = [pscustomobject]@{ Ok = $true; Reason = 'ok' }
            } else {
                $cloneStatus = [pscustomobject]@{ Ok = $false; Reason = 'missing' }
            }
            if (-not $cloneStatus.Ok) {
                Write-Host "  Skipping spec-API lint (upstream clone unavailable: $($cloneStatus.Reason))" -ForegroundColor DarkYellow
                Write-Host "    Hint: git clone --depth 1 --branch v1.5.8 https://github.com/alimtvnetwork/core-v9 $upstreamDir" -ForegroundColor DarkGray
                if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Spec-API Lint" "skip" "upstream clone $($cloneStatus.Reason)" }
            } else {
            Write-Host ""
            Write-Host "  Running spec-API fabrication lint (S-106)..." -ForegroundColor Yellow
            try {
                $strict = $ExtraArgs -and ($ExtraArgs -contains '--strict-spec-api')
                $specResult = Invoke-SpecApiCheck -SpecDir $specDir -UpstreamDir $upstreamDir
                $fabCount = 0
                if ($specResult) {
                    $fabCount = ($specResult.PkgFabrications.Count + $specResult.SymFabrications.Count)
                }
                if ($fabCount -eq 0) {
                    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Spec-API Lint" "pass" "no fabrications" }
                } elseif ($strict) {
                    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Spec-API Lint" "fail" "$fabCount fabrications (strict)" }
                    $s = Get-CallerSource; Write-Fail "Spec-API lint found $fabCount fabrications (strict mode). (source: $s)"
                    return $false
                } else {
                    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Spec-API Lint" "warn" "$fabCount fabrications (warn-only)" }
                }

                # ── S-106 v2 — signature/arity lint (depends on Go sig-index) ──
                $sigIndexer = Join-Path $ScriptRoot "scripts" "specapisig"
                $sigJson    = "/tmp/core-v9-sigindex.json"
                $sigDriver  = Join-Path $ScriptRoot "scripts" "spec-api-sig-check.psm1"
                if ((Test-Path $sigIndexer) -and (Test-Path $sigDriver)) {
                    Write-Host ""
                    Write-Host "  Running spec-API signature lint (S-106 v2)..." -ForegroundColor Yellow
                    # Always regenerate the sig-index to stay current with upstream.
                    $sigBuild = & go run "./scripts/specapisig" -roots "$upstreamDir,." -out $sigJson 2>&1
                    if ($LASTEXITCODE -ne 0) {
                        Write-Host ($sigBuild | Out-String) -ForegroundColor DarkYellow
                        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Spec-API Sig" "warn" "indexer failed" }
                    } else {
                        try {
                            Import-Module $sigDriver -Force -DisableNameChecking
                            $sigResult = Invoke-SpecApiSigCheck -SigIndex $sigJson -SpecDir $specDir
                            $sigMismatches = if ($sigResult) { $sigResult.ArityMismatches.Count } else { 0 }
                            if ($sigMismatches -eq 0) {
                                if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Spec-API Sig" "pass" "no arity mismatches" }
                            } elseif ($strict) {
                                if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Spec-API Sig" "fail" "$sigMismatches arity mismatches (strict)" }
                                $s = Get-CallerSource; Write-Fail "Spec-API sig lint found $sigMismatches arity mismatches (strict mode). (source: $s)"
                                return $false
                            } else {
                                if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Spec-API Sig" "warn" "$sigMismatches arity mismatches (warn-only)" }
                            }
                        } catch {
                            Write-Host "  Spec-API sig lint errored: $_" -ForegroundColor DarkYellow
                            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Spec-API Sig" "warn" "sig lint errored" }
                        }
                    }
                }
            } catch {
                Write-Host "  Spec-API lint errored: $_" -ForegroundColor DarkYellow
                if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Spec-API Lint" "warn" "lint errored" }
            }
            }
        }
    }

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
        $autofixDir = Join-Path $ScriptRoot "scripts" "autofix"
        if (-not (Test-Path $autofixDir)) {
            Write-Host "  Skipping Go auto-fixer (scripts/autofix/ not present in repo)" -ForegroundColor DarkYellow
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Auto-Fixer" "skip" "scripts/autofix/ missing" }
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
    }

    # ── Go syntax pre-check (bracecheck) ──────────────────────────────
    if ($skipBrace) {
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Syntax Check" "skip" "skipped (--skip-bracecheck)" }
    } else {
        $bracecheckDir = Join-Path $ScriptRoot "scripts" "bracecheck"
        if (-not (Test-Path $bracecheckDir)) {
            Write-Host "  Skipping Go syntax pre-check (scripts/bracecheck/ not present in repo)" -ForegroundColor DarkYellow
            if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Syntax Check" "skip" "scripts/bracecheck/ missing" }
            return $true
        }
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
