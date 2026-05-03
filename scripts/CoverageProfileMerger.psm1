# ─────────────────────────────────────────────────────────────────────────────
# CoverageProfileMerger.psm1 — Merge partial coverage profiles & detect gaps
#
# Usage:
#   Import-Module ./scripts/CoverageProfileMerger.psm1 -Force
#
# Dependencies: None (standalone)
# ─────────────────────────────────────────────────────────────────────────────

function Merge-CoverageProfiles {
    <#
    .SYNOPSIS
        Merge all partial coverage profiles into one using MAX count per line.
    .PARAMETER PartialDir
        Directory containing cover-*.out partial profiles.
    .PARAMETER CoverProfile
        Output path for the merged coverage.out file.
    .RETURNS
        Array of merged lines (including "mode: set" header).
    #>
    [CmdletBinding()]
    param(
        [string]$PartialDir,
        [string]$CoverProfile
    )

    $partialFiles = Get-ChildItem -Path $PartialDir -Filter "cover-*.out" | Sort-Object Name
    $coverMap = [System.Collections.Generic.Dictionary[string, int]]::new()

    foreach ($pf in $partialFiles) {
        $lines = Get-Content $pf.FullName
        foreach ($line in $lines) {
            if (-not $line -or $line -match "^mode:") { continue }
            if ($line -match "^(\S+\.go:\d+\.\d+,\d+\.\d+\s+\d+)\s+(\d+)\s*$") {
                $key = $Matches[1]
                $count = [int]$Matches[2]
                if ($coverMap.ContainsKey($key)) {
                    if ($count -gt $coverMap[$key]) { $coverMap[$key] = $count }
                } else {
                    $coverMap[$key] = $count
                }
            }
        }
    }

    $mergedLines = [System.Collections.Generic.List[string]]::new()
    $mergedLines.Add("mode: set")
    foreach ($entry in $coverMap.GetEnumerator()) {
        $mergedLines.Add("$($entry.Key) $($entry.Value)")
    }

    Set-Content -Path $CoverProfile -Value ($mergedLines -join "`n") -Encoding UTF8
    return $mergedLines.ToArray()
}

function Find-MissingCoverageProfiles {
    <#
    .SYNOPSIS
        Detect test packages whose coverage profiles are missing (binary crash).
    .PARAMETER TestPkgs
        List of test packages that were run.
    .PARAMETER PartialDir
        Directory where partial profiles were written.
    .PARAMETER IsSyncMode
        Whether sync mode was used (affects profile naming).
    .RETURNS
        List of short package names with missing profiles.
    #>
    [CmdletBinding()]
    param(
        [System.Collections.Generic.List[string]]$TestPkgs,
        [string]$PartialDir,
        [bool]$IsSyncMode
    )

    $missingProfiles = [System.Collections.Generic.List[string]]::new()
    foreach ($testPkg in $TestPkgs) {
        if ($IsSyncMode) {
            $idx = [array]::IndexOf($TestPkgs, $testPkg) + 1
            $expectedProfile = Join-Path $PartialDir "cover-$idx.out"
        } else {
            $safeName = $testPkg -replace '[^a-zA-Z0-9\.-]', '_'
            $expectedProfile = Join-Path $PartialDir "cover-$safeName.out"
        }
        if (-not (Test-Path $expectedProfile)) {
            $shortPkg = $testPkg -replace '.*integratedtests/?', ''
            if (-not $shortPkg) { $shortPkg = $testPkg }
            $missingProfiles.Add($shortPkg)
        }
    }

    if ($missingProfiles.Count -gt 0) {
        Write-Host ""
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Magenta
        Write-Host "  │ ⚠ MISSING COVERAGE PROFILES ($($missingProfiles.Count) package(s))" -ForegroundColor Magenta
        Write-Host "  │" -ForegroundColor Magenta
        Write-Host "  │ These test binaries likely crashed (panic/os.Exit)" -ForegroundColor Magenta
        Write-Host "  │ before Go could write their coverage profile." -ForegroundColor Magenta
        Write-Host "  │ Their coverage is NOT included in the report." -ForegroundColor Magenta
        Write-Host "  │" -ForegroundColor Magenta
        foreach ($mp in $missingProfiles) {
            Write-Host "  │   ⚠ $mp" -ForegroundColor Yellow
        }
        Write-Host "  │" -ForegroundColor Magenta
        Write-Host "  │ Fix: ensure tests use recover() for expected panics" -ForegroundColor Magenta
        Write-Host "  │ and never call os.Exit() in test code." -ForegroundColor Magenta
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Magenta
    }

    return $missingProfiles
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Merge-CoverageProfiles',
    'Find-MissingCoverageProfiles'
)
