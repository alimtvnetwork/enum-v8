# ─────────────────────────────────────────────────────────────────────────────
# PackageCoverage.psm1 — Single-package coverage command (TCP)
#
# Usage:
#   Import-Module ./scripts/PackageCoverage.psm1 -Force
#
# Dependencies: Utilities.psm1, TestLogWriter.psm1, TestRunner.psm1,
#               CoveragePreChecks.psm1, CoverageReport.psm1, DashboardUI.psm1
# ─────────────────────────────────────────────────────────────────────────────

function Invoke-PackageTestCoverage {
    <#
    .SYNOPSIS
        Run coverage for a single test package (TCP command).
    .PARAMETER pkg
        The test package directory name under tests/integratedtests/.
    .EXAMPLE
        Invoke-PackageTestCoverage "regexnewtests"
    #>
    param([string]$pkg)

    if (-not $pkg) {
        $s = Get-CallerSource; Write-Fail "Usage: ./run.ps1 TCP <package-name> (source: $s)"
        Write-Host "  Example: ./run.ps1 TCP regexnewtests" -ForegroundColor Gray
        return
    }

    Write-Header "Running coverage for package: $pkg"
    Invoke-FetchLatest

    # Clean data folder
    $dataDir = $global:DataDir
    if (Test-Path $dataDir) {
        Remove-Item -Recurse -Force $dataDir
        Write-Host "  Cleaned data/ folder" -ForegroundColor Yellow
    }

    # Pre-checks
    $coverDir = Join-Path $global:DataDir "coverage"
    $preCheckOk = Invoke-CoveragePreChecks -ScriptRoot $global:ProjectRoot -ExtraArgs $ExtraArgs -CoverDir $coverDir
    if (-not $preCheckOk) { exit 1 }

    # Build check
    Push-Location tests
    try { if (-not (Invoke-BuildCheck "./integratedtests/$pkg/...")) { return } }
    finally { Pop-Location }

    New-Item -ItemType Directory -Path $coverDir -Force | Out-Null

    $coverProfile = Join-Path $coverDir "coverage-$pkg.out"
    $coverHtml    = Join-Path $coverDir "coverage-$pkg.html"
    $coverSummary = Join-Path $coverDir "coverage-$pkg-summary.txt"

    # Build coverpkg list
    $allPkgs = go list ./... 2>&1 | ForEach-Object { $_.ToString() }
    $srcPkgs = $allPkgs | Where-Object { $_ -notmatch '/tests/' }
    $covPkgList = $srcPkgs -join ","

    # Run coverage
    $prevPref = $ErrorActionPreference
    $ErrorActionPreference = "Continue"
    $output = & go test -v -count=1 "-coverprofile=$coverProfile" "-coverpkg=$covPkgList" "./tests/integratedtests/$pkg/..." 2>&1 | ForEach-Object { $_.ToString() }
    $exitCode = $LASTEXITCODE
    $ErrorActionPreference = $prevPref

    Filter-TestWarnings $output | ForEach-Object { Write-Host $_ }
    Write-TestLogs $output

    if (Test-Path $coverProfile) {
        $funcOutput = & go tool cover "-func=$coverProfile" 2>&1 | ForEach-Object { $_.ToString() }

        # HTML report
        $htmlArgs = @("-html=$coverProfile", "-o=$coverHtml")
        & go tool cover $htmlArgs 2>&1 | Out-Null

        # Summary report
        $totalLine = $funcOutput | Where-Object { $_ -match "^total:" } | Select-Object -Last 1
        $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
        $summaryLines = [System.Collections.Generic.List[string]]::new()
        $summaryLines.Add("# Coverage Summary ($pkg) — $timestamp")
        $summaryLines.Add("")

        if ($totalLine) { $summaryLines.Add("## Total Coverage"); $summaryLines.Add("  $totalLine"); $summaryLines.Add("") }

        $lowCovFuncs = Get-LowCoverageFunctions $funcOutput
        if ($lowCovFuncs.Count -gt 0) {
            $summaryLines.Add("## Low Coverage Functions (< 50%)")
            $summaryLines.Add("  Count: $($lowCovFuncs.Count)")
            $summaryLines.Add("")
            foreach ($f in $lowCovFuncs) { $summaryLines.Add($f) }
            $summaryLines.Add("")
        }

        $summaryLines.Add("## Reports")
        $summaryLines.Add("  Profile:  $coverProfile")
        $summaryLines.Add("  HTML:     $coverHtml")
        $summaryLines.Add("  Summary:  $coverSummary")
        Set-Content -Path $coverSummary -Value ($summaryLines -join "`n") -Encoding UTF8

        # Console output
        Write-Host ""
        if ($totalLine -and $totalLine -match "(\d+\.\d+)%") {
            Write-Host "  total: (statements)  $($Matches[1])%" -ForegroundColor Cyan
        }
        Write-Host ""
        Write-Success "Coverage profile:  $coverProfile"
        Write-Success "HTML report:       $coverHtml"
        Write-Success "Summary:           $coverSummary"

        if ($lowCovFuncs.Count -gt 0) {
            Write-Host ""
            Write-Host "  ⚠ $($lowCovFuncs.Count) function(s) below 50% coverage" -ForegroundColor Yellow
        }

        # Coverage comparison
        if (Get-Command Write-CoverageComparison -ErrorAction SilentlyContinue) {
            $srcPkgMap = @{}
            foreach ($fLine in $funcOutput) {
                if ($fLine -match "^(\S+):\s+(\S+)\s+(\d+\.\d+)%\s*$" -and $fLine -notmatch "^total:") {
                    $filePath = $Matches[1]
                    $fPct = [double]$Matches[3]
                    $pathParts = $filePath -split '/'
                    $srcPkg = $pathParts[-2]
                    if (-not $srcPkgMap.ContainsKey($srcPkg)) {
                        $srcPkgMap[$srcPkg] = [System.Collections.Generic.List[double]]::new()
                    }
                    $srcPkgMap[$srcPkg].Add($fPct)
                }
            }
            $currentCovData = @($srcPkgMap.GetEnumerator() | ForEach-Object {
                $avg = ($_.Value | Measure-Object -Average).Average
                @{ Package = $_.Key; Coverage = [math]::Round($avg, 1) }
            })

            if ($currentCovData.Count -gt 0) {
                $previousCovData = $null
                if (Get-Command Load-CoverageSnapshot -ErrorAction SilentlyContinue) { $previousCovData = Load-CoverageSnapshot }
                Write-Host ""
                Write-CoverageComparison -Current $currentCovData -Previous $previousCovData
                if (Get-Command Save-CoverageSnapshot -ErrorAction SilentlyContinue) { Save-CoverageSnapshot -CoverageData $currentCovData }
            }
        }

        # AI coverage prompts
        $promptScript = Join-Path $global:ProjectRoot "scripts" "coverage" "Generate-CoveragePrompts.ps1"
        if (Test-Path $promptScript) {
            Write-Host ""
            Write-Header "Generating coverage improvement prompts"
            $promptsDir = Join-Path $global:DataDir "prompts"
            & $promptScript -CoverProfile $coverProfile -FuncOutput $funcOutput -OutputDir $promptsDir -BatchSize 500 -ProjectRoot $global:ProjectRoot
        }

        # HTML auto-open
        $openHtml = $false
        if ($ExtraArgs -and $ExtraArgs[-1] -eq "--open") { $openHtml = $true }
        if ($openHtml -and (Test-Path $coverHtml)) {
            Write-Host ""
            Write-Host "  Opening HTML coverage report..." -ForegroundColor Yellow
            Start-Process $coverHtml
        }
    }
    Open-FailingTestsIfAny
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Invoke-PackageTestCoverage'
)
