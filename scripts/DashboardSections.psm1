# ─────────────────────────────────────────────────────────────────────────────
# DashboardSections.psm1 — High-level section renderers (header, score, etc.)
#
# Dependencies: DashboardTheme.psm1, DashboardBoxPrimitives.psm1
# ─────────────────────────────────────────────────────────────────────────────

function Write-DashboardHeader {
    [CmdletBinding()]
    param([hashtable]$Data)
    $name    = if ($Data.ProductName) { $Data.ProductName } else { "Dashboard" }
    $version = if ($Data.Version)     { $Data.Version }     else { "" }
    Write-Host "  $($script:cLime)⚡$($script:cReset)  $($script:cWhite)$($script:cBold)$name $version$($script:cReset)"
    Write-Host "  $($script:cMuted)$("─" * ($script:BoxWidth - 2))$($script:cReset)"
}

function Write-ScanSummary {
    [CmdletBinding()]
    param([hashtable]$Data)
    $issuesFound = if ($null -ne $Data.IssuesFound) { $Data.IssuesFound } else { 0 }
    $issuesFixed = if ($null -ne $Data.IssuesFixed) { $Data.IssuesFixed } else { 0 }
    $agentCount  = if ($null -ne $Data.AgentCount)  { $Data.AgentCount }  else { 0 }
    $agents      = if ($Data.Agents) { $Data.Agents } else { @() }
    $labelWidth = 22
    Write-Host "  $($script:cCyan)▶$($script:cReset) $($script:cCyan)$("Scanning...".PadRight($labelWidth))$($script:cReset)$($script:cRed)$issuesFound issues found$($script:cReset)"
    Write-Host "  $($script:cCyan)▶$($script:cReset) $($script:cCyan)$("Auto-fixing...".PadRight($labelWidth))$($script:cReset)$($script:cLime)$issuesFixed resolved ✓$($script:cReset)"
    if ($agentCount -gt 0) {
        $agentList = ($agents -join " $($script:cMuted)·$($script:cReset) $($script:cMuted)")
        Write-Host "  $($script:cCyan)▶$($script:cReset) $($script:cCyan)$("$agentCount agents running".PadRight($labelWidth))$($script:cReset)$($script:cMuted)$agentList$($script:cReset)"
    }
}

function Write-ScoreBox {
    [CmdletBinding()]
    param([hashtable]$Data, [string]$Title = "")
    $w = $script:BoxWidth
    if (-not $Title -and $Data.ProductName) { $spaced = ($Data.ProductName.ToUpper().ToCharArray() -join " "); $Title = "$spaced   S C O R E" }
    elseif (-not $Title) { $Title = "S C O R E" }
    Write-BoxTop -Width $w; Write-BoxLineCenter -Text $Title -Width $w; Write-BoxDivider -Width $w; Write-BoxEmptyLine -Width $w
    $labelCol = 16; $scoreCol = 7; $barWidth = 15
    if ($Data.Scores) {
        foreach ($key in $Data.Scores.Keys) {
            $val = $Data.Scores[$key]; $label = $key.PadRight($labelCol)
            if ($val -is [int] -or $val -is [double] -or $val -is [decimal]) {
                $intVal = [int]$val; $scoreText = "$intVal/100".PadLeft($scoreCol)
                $bar = Get-ProgressBar -Score $intVal -BarWidth $barWidth
                Write-BoxLine -Content "$($script:cWhite)$label $scoreText  $bar" -Width $w
            } else {
                $valStr = "$val"
                $colored = if ($valStr -eq "PASS") { "$($script:cLime)$($script:cBold)$valStr$($script:cReset)" }
                    elseif ($valStr -eq "FAIL") { "$($script:cRed)$($script:cBold)$valStr$($script:cReset)" }
                    else { "$($script:cWhite)$valStr$($script:cReset)" }
                Write-BoxLine -Content "$($script:cWhite)$label $colored" -Width $w
            }
        }
    }
    Write-BoxEmptyLine -Width $w; Write-BoxDivider -Width $w; Write-BoxEmptyLine -Width $w
    $overallLabel = "OVERALL".PadRight($labelCol)
    $overallVal = if ($null -ne $Data.OverallScore) { "$($Data.OverallScore)/100" } else { "—" }
    Write-BoxLine -Content "$($script:cWhite)$($script:cBold)$overallLabel $($overallVal.PadLeft($scoreCol))$($script:cReset)" -Width $w
    $statusLabel = "STATUS".PadRight($labelCol)
    $statusText = if ($Data.Status) { $Data.Status } else { "UNKNOWN" }
    $statusReady = if ($null -ne $Data.StatusReady) { $Data.StatusReady } else { $false }
    $statusIcon = "$($script:cYellow)[?]$($script:cReset) "
    $statusColor = if ($statusReady) { $script:cLime } else { $script:cRed }
    Write-BoxLine -Content "$($script:cWhite)$statusLabel $statusIcon$statusColor$statusText$($script:cReset)" -Width $w
    Write-BoxEmptyLine -Width $w; Write-BoxBottom -Width $w
}

function Write-ResolutionSummary {
    [CmdletBinding()]
    param([hashtable]$Data)
    $fixed = if ($null -ne $Data.IssuesFixed) { $Data.IssuesFixed } else { 0 }
    $todos = if ($null -ne $Data.ManualTodos) { $Data.ManualTodos } else { 0 }
    Write-Host "  $($script:cLime)✓$($script:cReset) $($script:cLime)Fixed:$($script:cReset)  $($script:cWhite)$fixed$($script:cReset) $($script:cMuted)issues auto-resolved$($script:cReset)"
    if ($todos -gt 0) { Write-Host "  $($script:cYellow)●$($script:cReset) $($script:cYellow)Todo:$($script:cReset)   $($script:cWhite)$todos$($script:cReset) $($script:cMuted)manual items remaining$($script:cReset)" }
}

function Write-FooterTagline {
    [CmdletBinding()]
    param([string]$Text = "Ship it. One command. Production-ready.")
    Write-Host "  $($script:cLime)$($script:cBold)$Text$($script:cReset)"
}

function Write-BlockedDetails {
    [CmdletBinding()]
    param([array]$BlockedDetails)
    if (-not $BlockedDetails -or $BlockedDetails.Count -eq 0) { return }
    $dividerWidth = $script:BoxWidth
    Write-Host ""; Write-Host "  $($script:cMuted)── Blocked Packages $("─" * ($dividerWidth - 22))$($script:cReset)"; Write-Host ""
    foreach ($block in $BlockedDetails) {
        $pkg = if ($block.Package) { $block.Package } else { "unknown" }
        Write-Host "  $($script:cRed)$($script:cBold)✗ $pkg$($script:cReset)"
        if ($block.Errors) { foreach ($errLine in $block.Errors) { Write-Host "      $($script:cYellow)$errLine$($script:cReset)" } }
        Write-Host ""
    }
    Write-Host "  $($script:cMuted)$("─" * $dividerWidth)$($script:cReset)"
}

Export-ModuleMember -Function @(
    'Write-DashboardHeader', 'Write-ScanSummary', 'Write-ScoreBox',
    'Write-ResolutionSummary', 'Write-FooterTagline', 'Write-BlockedDetails'
)
