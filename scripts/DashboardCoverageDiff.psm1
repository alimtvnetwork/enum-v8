# ─────────────────────────────────────────────────────────────────────────────
# DashboardCoverageDiff.psm1 — Coverage diff rendering + snapshot management
#
# Dependencies: DashboardTheme.psm1, DashboardBoxPrimitives.psm1
# ─────────────────────────────────────────────────────────────────────────────

function Write-CoverageComparison {
    <# .SYNOPSIS Show coverage diff between current and previous run. #>
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)][array]$Current, [array]$Previous,
        [string]$PreviousJsonPath, [double]$Threshold = 0.0,
        [string]$Title = "C O V E R A G E   D I F F"
    )

    if ((-not $Previous -or $Previous.Count -eq 0) -and $PreviousJsonPath -and (Test-Path $PreviousJsonPath)) {
        try { $Previous = Get-Content $PreviousJsonPath -Raw | ConvertFrom-Json }
        catch { Write-Host "  $($script:cYellow)⚠$($script:cReset) $($script:cMuted)Could not load previous coverage$($script:cReset)"; return }
    }
    if (-not $Previous -or $Previous.Count -eq 0) {
        Write-Host "  $($script:cMuted)No previous coverage data available for comparison.$($script:cReset)"; return
    }

    $prevMap = @{}; foreach ($e in $Previous) { $prevMap[$e.Package] = [double]$e.Coverage }
    $currMap = @{}; foreach ($e in $Current)  { $currMap[$e.Package] = [double]$e.Coverage }
    $diffs = [System.Collections.Generic.List[hashtable]]::new()
    $allPackages = @($currMap.Keys) + @($prevMap.Keys) | Sort-Object -Unique

    foreach ($pkg in $allPackages) {
        $hasCurr = $currMap.ContainsKey($pkg); $hasPrev = $prevMap.ContainsKey($pkg)
        $curr = if ($hasCurr) { $currMap[$pkg] } else { $null }
        $prev = if ($hasPrev) { $prevMap[$pkg] } else { $null }
        if ($null -ne $curr -and $null -ne $prev) {
            $delta = $curr - $prev
            if ([math]::Abs($delta) -ge $Threshold) {
                $diffs.Add(@{ Package = $pkg; Current = $curr; Previous = $prev; Delta = $delta; Status = if ($delta -gt 0) { "up" } elseif ($delta -lt 0) { "down" } else { "same" } })
            }
        } elseif ($null -ne $curr -and $null -eq $prev) {
            $diffs.Add(@{ Package = $pkg; Current = $curr; Previous = $null; Delta = $null; Status = "new" })
        } elseif ($null -eq $curr -and $null -ne $prev) {
            $diffs.Add(@{ Package = $pkg; Current = $null; Previous = $prev; Delta = $null; Status = "removed" })
        }
    }

    if ($diffs.Count -eq 0) { Write-Host "  $($script:cLime)✓$($script:cReset) $($script:cMuted)No coverage changes detected.$($script:cReset)"; return }

    $sorted = $diffs | Sort-Object { if ($null -ne $_.Delta) { $_.Delta } else { -999 } }
    $pkgCol = 24; $prevCol = 7; $currCol = 7; $deltaCol = 8
    $w = [math]::Max($script:BoxWidth, $pkgCol + $prevCol + $currCol + $deltaCol + 8)
    Write-BoxTop -Width $w; Write-BoxLineCenter -Text $Title -Width $w; Write-BoxDivider -Width $w; Write-BoxEmptyLine -Width $w
    $hdrVisLen = $pkgCol + 1 + $prevCol + 1 + $currCol + 1 + $deltaCol
    Write-BoxLine -Content "$($script:cMuted)$("Package".PadRight($pkgCol)) $("Prev".PadLeft($prevCol)) $("Curr".PadLeft($currCol)) $("Delta".PadLeft($deltaCol))$($script:cReset)" -Width $w -VisualLength $hdrVisLen
    Write-BoxLine -Content "$($script:cMuted)$("─" * $pkgCol) $("─" * $prevCol) $("─" * $currCol) $("─" * $deltaCol)$($script:cReset)" -Width $w -VisualLength $hdrVisLen

    $regressions = 0; $improvements = 0; $newPkgs = 0; $lost100 = 0; $gained100 = 0
    foreach ($d in $sorted) {
        $pkg = "$($d.Package)"; if ($pkg.Length -gt $pkgCol) { $pkg = $pkg.Substring(0, $pkgCol - 2) + ".." }
        $pkgStr = $pkg.PadRight($pkgCol)
        $prevStr = if ($null -ne $d.Previous) { ("{0:F1}%" -f $d.Previous).PadLeft($prevCol) } else { "—".PadLeft($prevCol) }
        $currStr = if ($null -ne $d.Current) { ("{0:F1}%" -f $d.Current).PadLeft($currCol) } else { "—".PadLeft($currCol) }
        switch ($d.Status) {
            "up"      { $improvements++; $deltaStr = ("+{0:F1}%" -f $d.Delta).PadLeft($deltaCol); $icon = "▲"; $rowColor = $script:cLime; if ($d.Current -ge 100.0 -and $d.Previous -lt 100.0) { $gained100++ } }
            "down"    { $regressions++;  $deltaStr = ("{0:F1}%" -f $d.Delta).PadLeft($deltaCol);  $icon = "▼"; $rowColor = $script:cRed;  if ($d.Previous -ge 100.0 -and $d.Current -lt 100.0) { $lost100++ } }
            "same"    { $deltaStr = "0.0%".PadLeft($deltaCol); $icon = "─"; $rowColor = $script:cMuted }
            "new"     { $newPkgs++; $deltaStr = "NEW".PadLeft($deltaCol); $icon = "★"; $rowColor = $script:cCyan }
            "removed" { $deltaStr = "GONE".PadLeft($deltaCol); $icon = "✗"; $rowColor = $script:cYellow }
        }
        Write-BoxLine -Content "$rowColor$icon $pkgStr$($script:cReset) $($script:cMuted)$prevStr$($script:cReset) $rowColor$currStr$($script:cReset) $rowColor$deltaStr$($script:cReset)" -Width $w -VisualLength (2 + $pkgCol + 1 + $prevCol + 1 + $currCol + 1 + $deltaCol)
    }

    Write-BoxEmptyLine -Width $w; Write-BoxDivider -Width $w; Write-BoxEmptyLine -Width $w
    $line1 = " "
    if ($regressions -gt 0) { $line1 += "$($script:cRed)▼ $regressions regressions$($script:cReset)" } else { $line1 += "$($script:cLime)✓ 0 regressions$($script:cReset)" }
    $line1 += "  $($script:cLime)▲ $improvements improved$($script:cReset)"
    if ($newPkgs -gt 0) { $line1 += "  $($script:cCyan)★ $newPkgs new$($script:cReset)" }
    Write-BoxLine -Content $line1 -Width $w -VisualLength ($pkgCol + 42)
    if ($lost100 -gt 0) { Write-BoxLine -Content " $($script:cRed)$($script:cBold)⚠ $lost100 package(s) dropped from 100%$($script:cReset)" -Width $w -VisualLength ($pkgCol + 37) }
    if ($gained100 -gt 0) { Write-BoxLine -Content " $($script:cLime)★ $gained100 package(s) reached 100%$($script:cReset)" -Width $w -VisualLength ($pkgCol + 32) }
    Write-BoxEmptyLine -Width $w; Write-BoxBottom -Width $w
}

function Save-CoverageSnapshot {
    <# .SYNOPSIS Save current coverage data as JSON for future comparison. #>
    [CmdletBinding()]
    param([Parameter(Mandatory)][array]$CoverageData, [string]$Path)
    if (-not $Path) { $Path = Join-Path $global:DataDir "coverage" "coverage-previous.json" }
    $dir = Split-Path $Path -Parent
    if ($dir -and -not (Test-Path $dir)) { New-Item -ItemType Directory -Path $dir -Force | Out-Null }
    @{ timestamp = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"); packages = $CoverageData } |
        ConvertTo-Json -Depth 5 | Set-Content -Path $Path -Encoding UTF8
    Write-Host "  $($script:cMuted)Coverage snapshot saved → $Path$($script:cReset)"
}

function Load-CoverageSnapshot {
    <# .SYNOPSIS Load a previously saved coverage snapshot. #>
    [CmdletBinding()]
    [OutputType([array])]
    param([string]$Path)
    if (-not $Path) { $Path = Join-Path $global:DataDir "coverage" "coverage-previous.json" }
    if (-not (Test-Path $Path)) { return $null }
    try { $json = Get-Content $Path -Raw | ConvertFrom-Json; if ($json.packages) { return $json.packages }; return $json } catch { return $null }
}

Export-ModuleMember -Function @('Write-CoverageComparison', 'Save-CoverageSnapshot', 'Load-CoverageSnapshot')
