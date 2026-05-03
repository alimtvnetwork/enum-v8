# ─────────────────────────────────────────────────────────────────────────────
# CoverageReportJson.psm1 — JSON report generation + error/failure reports
#
# Dependencies: CoverageReportTxt.psm1 (Get-LowCoverageFunctions),
#               ErrorExtractor.psm1 (Resolve-BuildDiagnosticLines)
# ─────────────────────────────────────────────────────────────────────────────

function Write-CoverageJsonReport {
    <# .SYNOPSIS Generate coverage-summary.json with per-package and low-coverage data. #>
    [CmdletBinding()]
    param(
        [string]$CoverJsonFile, [string]$CoverProfile, [string]$CoverHtml, [string]$CoverSummary,
        [string[]]$FuncOutput, [hashtable]$SrcPkgStmts, [string]$CoverDir
    )

    $totalLine = $FuncOutput | Where-Object { $_ -match "^total:" } | Select-Object -Last 1
    $totalPct = 0.0
    if ($totalLine -match "(\d+\.\d+)%") { $totalPct = [double]$Matches[1] }

    $pkgJsonItems = [System.Collections.Generic.List[object]]::new()
    if ($SrcPkgStmts.Count -gt 0) {
        $sorted = $SrcPkgStmts.GetEnumerator() | ForEach-Object {
            $pct = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
            [pscustomobject]@{ Name = $_.Key; Pct = $pct; Stmts = $_.Value.Stmts; Covered = $_.Value.Covered }
        } | Sort-Object Pct
        foreach ($e in $sorted) {
            $pkgJsonItems.Add(@{ package = $e.Name; coverage = $e.Pct; statements = $e.Stmts; covered = $e.Covered; uncovered = $e.Stmts - $e.Covered })
        }
    }

    $lowCovJsonItems = [System.Collections.Generic.List[object]]::new()
    foreach ($line in $FuncOutput) {
        if ($line -match "(\d+\.\d+)%\s*$" -and $line -notmatch "^total:") {
            $pctF = [double]$Matches[1]
            if ($pctF -lt 50.0) {
                $funcName = ""; $funcFile = ""
                if ($line -match "^(\S+):\s+(\S+)\s+(\d+\.\d+)%") { $funcFile = $Matches[1]; $funcName = $Matches[2] }
                $lowCovJsonItems.Add(@{ file = $funcFile; function = $funcName; coverage = $pctF })
            }
        }
    }

    $blockedRef = @()
    $blockedJsonPath = Join-Path $CoverDir "blocked-packages.json"
    if (Test-Path $blockedJsonPath) {
        $blockedRef = @((Get-Content $blockedJsonPath -Raw | ConvertFrom-Json).blockedPackages | ForEach-Object { $_.package })
    }

    $coverJsonObj = @{
        timestamp = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"); totalCoverage = $totalPct
        packageCount = $pkgJsonItems.Count; packages = $pkgJsonItems.ToArray()
        lowCoverageFuncCount = $lowCovJsonItems.Count; lowCoverageFunctions = $lowCovJsonItems.ToArray()
        blockedPackages = $blockedRef
        reports = @{ profile = $CoverProfile; html = $CoverHtml; summary = $CoverSummary; json = $CoverJsonFile }
    }
    $coverJsonObj | ConvertTo-Json -Depth 4 | Set-Content -Path $CoverJsonFile -Encoding UTF8
    return $totalPct
}

function Write-BuildErrorsReport {
    <# .SYNOPSIS Generate build-errors.txt and build-errors.json reports. #>
    [CmdletBinding()]
    param(
        [string]$CoverDir, [hashtable]$BuildErrorsByPackage,
        [System.Collections.Generic.List[string]]$BlockedPkgs,
        [System.Collections.Generic.Dictionary[string, string]]$BlockedErrors
    )

    if ($BlockedPkgs.Count -gt 0) {
        foreach ($bp in ($BlockedPkgs | Sort-Object)) {
            if ($BlockedErrors.ContainsKey($bp)) {
                $rawErrLines = $BlockedErrors[$bp] -split "`n"
                $filteredErrLines = Resolve-BuildDiagnosticLines $rawErrLines
                if ($filteredErrLines.Count -gt 0) {
                    if (-not $BuildErrorsByPackage.ContainsKey($bp)) { $BuildErrorsByPackage[$bp] = [System.Collections.Generic.List[string]]::new() }
                    foreach ($errLine in $filteredErrLines) {
                        if (-not $BuildErrorsByPackage[$bp].Contains($errLine)) { $BuildErrorsByPackage[$bp].Add($errLine) | Out-Null }
                    }
                }
            }
        }
    }

    $buildErrorsFile = Join-Path $CoverDir "build-errors.txt"
    $buildErrorsJsonFile = Join-Path $CoverDir "build-errors.json"
    $buildErrorPkgs = @($BuildErrorsByPackage.Keys | Sort-Object)
    $callerSource = Get-CallerSource
    $lines = [System.Collections.Generic.List[string]]::new()
    $lines.Add("# Build Errors — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')"); $lines.Add("# Count: $($buildErrorPkgs.Count)")
    $lines.Add("# Source: $callerSource (CoverageReportJson.psm1 → Write-BuildErrorsReport)"); $lines.Add("")
    $jsonItems = [System.Collections.Generic.List[object]]::new()

    if ($buildErrorPkgs.Count -eq 0) { $lines.Add("No build errors captured.") }
    else {
        foreach ($pkgName in $buildErrorPkgs) {
            $pkgLines = @($BuildErrorsByPackage[$pkgName])
            $isBlocked = $BlockedPkgs.Contains($pkgName)
            $lines.Add($(if ($isBlocked) { "## $pkgName [BLOCKED — compile failure]" } else { "## $pkgName [coverage-run error]" }))
            if ($pkgLines.Count -gt 0) { $lines.AddRange([string[]]$pkgLines) } else { $lines.Add("(no actionable compile errors captured)") }
            $lines.Add("")
            $jsonItems.Add(@{ package = $pkgName; errorCount = $pkgLines.Count; errors = $pkgLines; source = if ($isBlocked) { "compile-check" } else { "coverage-run" } }) | Out-Null
        }
    }
    Set-Content -Path $buildErrorsFile -Value ($lines -join "`n") -Encoding UTF8
    @{ timestamp = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"); packageCount = $buildErrorPkgs.Count; blockedCount = $BlockedPkgs.Count; source = $callerSource; packages = $jsonItems.ToArray() } |
        ConvertTo-Json -Depth 5 | Set-Content -Path $buildErrorsJsonFile -Encoding UTF8
}

function Write-RuntimeFailuresReport {
    <# .SYNOPSIS Generate runtime-failures.txt and runtime-failures.json reports. #>
    [CmdletBinding()]
    param(
        [string]$CoverDir, [hashtable]$RuntimeFailuresByPackage,
        [System.Collections.Generic.List[string]]$MissingProfiles,
        [System.Collections.Generic.List[string]]$BlockedPkgs
    )

    if ($MissingProfiles.Count -gt 0) {
        foreach ($mp in $MissingProfiles) {
            if (-not $RuntimeFailuresByPackage.ContainsKey($mp)) { $RuntimeFailuresByPackage[$mp] = [System.Collections.Generic.List[string]]::new() }
            $crashMsg = "coverage profile missing — test binary likely crashed (panic/os.Exit)"
            if (-not $RuntimeFailuresByPackage[$mp].Contains($crashMsg)) { $RuntimeFailuresByPackage[$mp].Add($crashMsg) | Out-Null }
        }
    }

    $runtimeFailuresFile = Join-Path $CoverDir "runtime-failures.txt"
    $runtimeFailuresJsonFile = Join-Path $CoverDir "runtime-failures.json"
    $blockedLookup = @{}
    if ($BlockedPkgs) {
        foreach ($bp in $BlockedPkgs) { $blockedLookup[$bp] = $true }
    }
    $runtimeFailurePkgs = @($RuntimeFailuresByPackage.Keys | Where-Object { -not $blockedLookup.ContainsKey($_) } | Sort-Object)
    $callerSource = Get-CallerSource
    $rtLines = [System.Collections.Generic.List[string]]::new()
    $rtLines.Add("# Runtime Failures — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')")
    $rtLines.Add("# Panics, os.Exit, test binary crashes, fatal errors")
    $rtLines.Add("# Count: $($runtimeFailurePkgs.Count) package(s)")
    $rtLines.Add("# Source: $callerSource (CoverageReportJson.psm1 → Write-RuntimeFailuresReport)"); $rtLines.Add("")
    $rtJsonItems = [System.Collections.Generic.List[object]]::new()

    if ($runtimeFailurePkgs.Count -eq 0) { $rtLines.Add("No runtime failures captured.") }
    else {
        foreach ($pkgName in $runtimeFailurePkgs) {
            $pkgLines = @($RuntimeFailuresByPackage[$pkgName])
            $rtLines.Add("## $pkgName")
            if ($pkgLines.Count -gt 0) { $rtLines.AddRange([string[]]$pkgLines) } else { $rtLines.Add("(no failure details captured)") }
            $rtLines.Add("")
            $rtJsonItems.Add(@{ package = $pkgName; failureCount = $pkgLines.Count; failures = $pkgLines }) | Out-Null
        }
    }
    Set-Content -Path $runtimeFailuresFile -Value ($rtLines -join "`n") -Encoding UTF8
    @{ timestamp = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"); packageCount = $runtimeFailurePkgs.Count; source = $callerSource; packages = $rtJsonItems.ToArray() } |
        ConvertTo-Json -Depth 5 | Set-Content -Path $runtimeFailuresJsonFile -Encoding UTF8

    if ($runtimeFailurePkgs.Count -gt 0) {
        Write-Host ""
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Magenta
        Write-Host "  │ RUNTIME FAILURES ($($runtimeFailurePkgs.Count) package(s))" -ForegroundColor Magenta
        Write-Host "  │" -ForegroundColor Magenta
        foreach ($rp in $runtimeFailurePkgs) { Write-Host "  │   ⚠ $rp" -ForegroundColor Yellow }
        Write-Host "  │" -ForegroundColor Magenta
        Write-Host "  │ See data/coverage/runtime-failures.txt for details." -ForegroundColor Yellow
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Magenta
    }
}

Export-ModuleMember -Function @(
    'Write-CoverageJsonReport', 'Write-BuildErrorsReport', 'Write-RuntimeFailuresReport'
)
