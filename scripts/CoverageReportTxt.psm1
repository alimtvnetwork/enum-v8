# ─────────────────────────────────────────────────────────────────────────────
# CoverageReportTxt.psm1 — TXT report generation + coverage parsing helpers
#
# Dependencies: None (standalone)
# ─────────────────────────────────────────────────────────────────────────────

function Build-SourcePackageCoverage {
    <#
    .SYNOPSIS
        Parse merged coverage lines into per-source-package coverage stats.
    .RETURNS
        Ordered hashtable: shortPkgName → @{ Stmts; Covered }
    #>
    [CmdletBinding()]
    param([string[]]$MergedLines)

    $srcPkgStmts = [ordered]@{}
    foreach ($covLine in $MergedLines) {
        if ($covLine -match "^mode:") { continue }
        if ($covLine -match "^(\S+?):(\d+)\.(\d+),(\d+)\.(\d+)\s+(\d+)\s+(\d+)") {
            $filePath = $Matches[1]
            $stmts = [int]$Matches[6]
            $count = [int]$Matches[7]
            $shortSrc = $filePath -replace '.*alimtvnetwork/core-v8/?', ''
            $shortSrc = $shortSrc -replace '/[^/]+$', ''
            if (-not $shortSrc) { $shortSrc = "(root)" }
            if (-not $srcPkgStmts.Contains($shortSrc)) {
                $srcPkgStmts[$shortSrc] = @{ Stmts = 0; Covered = 0 }
            }
            $srcPkgStmts[$shortSrc].Stmts += $stmts
            if ($count -gt 0) { $srcPkgStmts[$shortSrc].Covered += $stmts }
        }
    }
    return $srcPkgStmts
}

function Get-LowCoverageFunctions {
    <# .SYNOPSIS Extract functions with < 50% coverage from func output. #>
    [CmdletBinding()]
    param([string[]]$FuncOutput)

    $lowCovFuncs = [System.Collections.Generic.List[string]]::new()
    foreach ($line in $FuncOutput) {
        if ($line -match "(\d+\.\d+)%\s*$" -and $line -notmatch "^total:") {
            $pct = [double]$Matches[1]
            if ($pct -lt 50.0) { $lowCovFuncs.Add("  $line") }
        }
    }
    return $lowCovFuncs
}

function Write-CoverageSummaryReport {
    <# .SYNOPSIS Generate coverage-summary.txt file. #>
    [CmdletBinding()]
    param(
        [string]$CoverSummary, [string]$CoverProfile, [string]$CoverHtml,
        [string[]]$FuncOutput, [hashtable]$SrcPkgStmts
    )

    $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    $summaryLines = [System.Collections.Generic.List[string]]::new()
    $summaryLines.Add("# Coverage Summary — $timestamp"); $summaryLines.Add("")

    $totalLine = $FuncOutput | Where-Object { $_ -match "^total:" } | Select-Object -Last 1
    if ($totalLine) { $summaryLines.Add("## Total Coverage"); $summaryLines.Add("  $totalLine"); $summaryLines.Add("") }

    if ($SrcPkgStmts.Count -gt 0) {
        $summaryLines.Add("## Per-Package Coverage (Source)")
        $sorted = $SrcPkgStmts.GetEnumerator() | ForEach-Object {
            $pct = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
            [pscustomobject]@{ Name = $_.Key; Pct = $pct }
        } | Sort-Object Pct -Descending
        foreach ($entry in $sorted) { $summaryLines.Add("  $($entry.Pct)%`t$($entry.Name)") }
        $summaryLines.Add("")
    }

    $lowCovFuncs = Get-LowCoverageFunctions $FuncOutput
    if ($lowCovFuncs.Count -gt 0) {
        $summaryLines.Add("## Low Coverage Functions (< 50%)")
        $summaryLines.Add("  Count: $($lowCovFuncs.Count)"); $summaryLines.Add("")
        foreach ($f in $lowCovFuncs) { $summaryLines.Add($f) }; $summaryLines.Add("")
    }

    $summaryLines.Add("## Reports")
    $summaryLines.Add("  Profile:  $CoverProfile")
    $summaryLines.Add("  HTML:     $CoverHtml")
    $summaryLines.Add("  Summary:  $CoverSummary")
    Set-Content -Path $CoverSummary -Value ($summaryLines -join "`n") -Encoding UTF8
}

function Write-PerPackageCoverageReport {
    <# .SYNOPSIS Generate per-package-coverage.txt and per-package-coverage.json. #>
    [CmdletBinding()]
    param([string]$CoverDir, [hashtable]$SrcPkgStmts, [double]$TotalPct, [string]$TotalLine)

    $perPkgTxtFile = Join-Path $CoverDir "per-package-coverage.txt"
    $perPkgJsonFile = Join-Path $CoverDir "per-package-coverage.json"
    $txtLines = [System.Collections.Generic.List[string]]::new()
    $txtLines.Add("# Per-Package Coverage Report — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')")
    $txtLines.Add("# Total: $TotalLine"); $txtLines.Add("")
    $txtLines.Add(("Package".PadRight(50)) + " " + ("Stmts".PadLeft(8)) + " " + ("Covered".PadLeft(8)) + " " + ("Uncovered".PadLeft(10)) + " " + ("Cov%".PadLeft(8)))
    $txtLines.Add(("─" * 50) + " " + ("─" * 8) + " " + ("─" * 8) + " " + ("─" * 10) + " " + ("─" * 8))

    $jsonItems = [System.Collections.Generic.List[object]]::new()
    if ($SrcPkgStmts.Count -gt 0) {
        $sorted = $SrcPkgStmts.GetEnumerator() | ForEach-Object {
            $pct = if ($_.Value.Stmts -gt 0) { [math]::Round(($_.Value.Covered / $_.Value.Stmts) * 100, 1) } else { 0 }
            [pscustomobject]@{ Name = $_.Key; Pct = $pct; Stmts = $_.Value.Stmts; Covered = $_.Value.Covered; Uncovered = $_.Value.Stmts - $_.Value.Covered }
        } | Sort-Object Pct
        foreach ($pp in $sorted) {
            $statusMark = if ($pp.Pct -ge 100) { "✓" } elseif ($pp.Pct -ge 80) { "○" } else { "✗" }
            $row = ("$statusMark $($pp.Name)").PadRight(50)
            $row += " " + $pp.Stmts.ToString().PadLeft(8) + " " + $pp.Covered.ToString().PadLeft(8)
            $row += " " + $pp.Uncovered.ToString().PadLeft(10)
            $row += " " + (([string]::Format([System.Globalization.CultureInfo]::InvariantCulture, "{0:0.0}", $pp.Pct)) + "%").PadLeft(8)
            $txtLines.Add($row)
            $jsonItems.Add(@{ package = $pp.Name; coverage = $pp.Pct; statements = $pp.Stmts; covered = $pp.Covered; uncovered = $pp.Uncovered; status = if ($pp.Pct -ge 100) { "full" } elseif ($pp.Pct -ge 80) { "good" } else { "low" } })
        }
        $totalStmts = ($sorted | Measure-Object -Property Stmts -Sum).Sum
        $totalCovered = ($sorted | Measure-Object -Property Covered -Sum).Sum
        $fullCount = ($sorted | Where-Object { $_.Pct -ge 100 }).Count
        $lowCount = ($sorted | Where-Object { $_.Pct -lt 80 }).Count
        $txtLines.Add(""); $txtLines.Add("# Summary")
        $txtLines.Add("#   Packages:  $($sorted.Count)"); $txtLines.Add("#   100%:      $fullCount")
        $txtLines.Add("#   < 80%:     $lowCount")
        $txtLines.Add("#   Total stmts: $totalStmts  covered: $totalCovered  uncovered: $($totalStmts - $totalCovered)")
    }
    Set-Content -Path $perPkgTxtFile -Value ($txtLines -join "`n") -Encoding UTF8
    @{ timestamp = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"); totalCoverage = $TotalPct; packageCount = $jsonItems.Count; packages = $jsonItems.ToArray() } |
        ConvertTo-Json -Depth 4 | Set-Content -Path $perPkgJsonFile -Encoding UTF8
}

Export-ModuleMember -Function @(
    'Build-SourcePackageCoverage', 'Get-LowCoverageFunctions',
    'Write-CoverageSummaryReport', 'Write-PerPackageCoverageReport'
)
