# ─────────────────────────────────────────────────────────────────────────────
# DashboardCoverageTable.psm1 — Bordered per-package coverage table
#
# Dependencies: DashboardTheme.psm1, DashboardBoxPrimitives.psm1
# ─────────────────────────────────────────────────────────────────────────────

function Write-CoverageTable {
    <# .SYNOPSIS Render a bordered per-package coverage table with progress bars. #>
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)][array]$CoverageData,
        [string]$Title = "P A C K A G E   C O V E R A G E",
        [bool]$ShowTarget = $true, [int]$BarWidth = 12
    )

    if (-not $CoverageData -or $CoverageData.Count -eq 0) { return }

    $sorted = $CoverageData | Sort-Object { [double]$_.Coverage }
    $pkgCol = 24; $pctCol = 7; $testCol = 5
    $contentWidth = 1 + $pkgCol + 1 + $pctCol + 2 + $BarWidth + 2 + $testCol + 1
    $w = [math]::Max($script:BoxWidth, $contentWidth)

    Write-BoxTop -Width $w; Write-BoxLineCenter -Text $Title -Width $w
    Write-BoxDivider -Width $w; Write-BoxEmptyLine -Width $w

    $barHeader = ''.PadRight($BarWidth)
    Write-BoxLine -Content "$($script:cMuted)$("Package".PadRight($pkgCol)) $("Cov %".PadLeft($pctCol))  $barHeader  $("Tests".PadLeft($testCol))$($script:cReset)" -Width $w
    Write-BoxLine -Content "$($script:cMuted)$("─" * $pkgCol) $("─" * $pctCol)  $("─" * $BarWidth)  $("─" * $testCol)$($script:cReset)" -Width $w

    $totalCoverage = 0.0; $at100Count = 0; $below100Count = 0
    foreach ($entry in $sorted) {
        $pkg = "$($entry.Package)"; $cov = [double]$entry.Coverage
        $tests = if ($null -ne $entry.Tests) { [int]$entry.Tests } else { 0 }
        $totalCoverage += $cov
        if ($cov -ge 100.0) { $at100Count++ } else { $below100Count++ }
        if ($pkg.Length -gt $pkgCol) { $pkg = $pkg.Substring(0, $pkgCol - 2) + ".." }
        $pkgStr = $pkg.PadRight($pkgCol); $pctStr = ("{0:F1}%" -f $cov).PadLeft($pctCol)
        $rowColor = if ($cov -ge 100.0) { $script:cLime } elseif ($cov -ge 98.0) { $script:cWhite } elseif ($cov -ge 95.0) { $script:cYellow } else { $script:cRed }
        $bar = Get-ProgressBar -Score ([int][math]::Round($cov)) -BarWidth $BarWidth
        $testStr = "$tests".PadLeft($testCol)
        Write-BoxLine -Content "$rowColor$pkgStr$($script:cReset) $rowColor$pctStr$($script:cReset)  $bar  $($script:cMuted)$testStr$($script:cReset)" -Width $w
    }

    Write-BoxEmptyLine -Width $w; Write-BoxDivider -Width $w; Write-BoxEmptyLine -Width $w
    $avgCoverage = if ($sorted.Count -gt 0) { $totalCoverage / $sorted.Count } else { 0 }
    $summaryBar = Get-ProgressBar -Score ([int][math]::Round($avgCoverage)) -BarWidth $BarWidth
    Write-BoxLine -Content "$($script:cWhite)$($script:cBold)$("AVERAGE".PadRight($pkgCol))$($script:cReset) $($script:cWhite)$($script:cBold)$(("{0:F1}%" -f $avgCoverage).PadLeft($pctCol))$($script:cReset)  $summaryBar  $($script:cMuted)$("$($sorted.Count)".PadLeft($testCol))$($script:cReset)" -Width $w
    $countText = "$($script:cLime)$at100Count$($script:cReset)$($script:cMuted) at 100%$($script:cReset)  $($script:cYellow)$below100Count$($script:cReset)$($script:cMuted) below$($script:cReset)"
    $countPad = ''.PadRight($pkgCol)
    Write-BoxLine -Content "$countPad$countText" -Width $w
    if ($ShowTarget) {
        $targetText = "$($script:cLime)$($script:cBold)100.0%$($script:cReset)$($script:cMuted) (non-internal packages)$($script:cReset)"
        Write-BoxLine -Content "$($script:cWhite)$("TARGET".PadRight($pkgCol))$($script:cReset)$targetText" -Width $w
    }
    Write-BoxEmptyLine -Width $w; Write-BoxBottom -Width $w
}

Export-ModuleMember -Function @('Write-CoverageTable')
