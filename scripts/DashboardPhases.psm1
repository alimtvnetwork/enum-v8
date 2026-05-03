# ─────────────────────────────────────────────────────────────────────────────
# DashboardPhases.psm1 — Phase registry, summary box rendering
#
# Usage:
#   Import-Module ./scripts/DashboardPhases.psm1 -Force
#
# Dependencies: DashboardTheme.psm1, DashboardBoxes.psm1
# ─────────────────────────────────────────────────────────────────────────────

$script:Phases = [ordered]@{}

function Register-Phase {
    <#
    .SYNOPSIS
        Record a phase result for the final dashboard summary.
    #>
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)][string]$Name,
        [Parameter(Mandatory)][ValidateSet("pass","fail","skip","warn")][string]$Status,
        [string]$Detail = ""
    )
    $script:Phases[$Name] = @{ Status = $Status; Detail = $Detail }
}

function Reset-Phases {
    <# .SYNOPSIS Clear all registered phases. #>
    [CmdletBinding()]
    param()
    $script:Phases = [ordered]@{}
}

function Get-IconVisualWidth {
    [CmdletBinding()]
    [OutputType([int])]
    param([string]$Status)
    if ($Status -eq "warn") { return 2 } else { return 1 }
}

function Get-PhaseIcon {
    [CmdletBinding()]
    [OutputType([string])]
    param([string]$Status)

    switch ($Status) {
        "pass" { return "$($script:cLime)✓$($script:cReset)" }
        "fail" { return "$($script:cRed)✗$($script:cReset)" }
        "skip" { return "$($script:cMuted)⊘$($script:cReset)" }
        "warn" { return "$($script:cYellow)⚠$($script:cReset)" }
        default { return "$($script:cMuted)?$($script:cReset)" }
    }
}

function Write-PhaseStart {
    [CmdletBinding()]
    param([Parameter(Mandatory)][string]$Name)
    Write-Host "  $($script:cCyan)▶$($script:cReset) $($script:cWhite)$Name$($script:cReset)$($script:cMuted)...$($script:cReset)"
}

function Write-PhaseEnd {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)][string]$Name,
        [Parameter(Mandatory)][ValidateSet("pass","fail","skip","warn")][string]$Status,
        [string]$Detail = ""
    )

    $icon = Get-PhaseIcon $Status
    $detailColor = switch ($Status) {
        "fail" { $script:cRed }
        "warn" { $script:cYellow }
        default { $script:cMuted }
    }
    $detailStr = if ($Detail) { "  $detailColor$Detail$($script:cReset)" } else { "" }
    Write-Host "  $icon $($script:cWhite)$Name$($script:cReset)$detailStr"
}

function Write-PhaseSummaryBox {
    <#
    .SYNOPSIS
        Render the bordered phase summary box from registered phases.
    #>
    [CmdletBinding()]
    param([System.Collections.Specialized.OrderedDictionary]$Phases)

    if (-not $Phases) { $Phases = $script:Phases }
    if ($Phases.Count -eq 0) { return }

    $w = $script:BoxWidth
    $phaseLabelWidth = 20

    Write-BoxTop -Width $w
    Write-BoxLineCenter -Text "P H A S E   S U M M A R Y" -Width $w
    Write-BoxDivider -Width $w
    Write-BoxEmptyLine -Width $w

    $passCount = 0; $failCount = 0; $warnCount = 0

    foreach ($key in $Phases.Keys) {
        $phase  = $Phases[$key]
        $status = $phase.Status
        $detail = if ($phase.Detail) { $phase.Detail } else { "" }
        $icon   = Get-PhaseIcon $status

        switch ($status) {
            "pass" { $passCount++ }
            "fail" { $failCount++ }
            "warn" { $warnCount++ }
        }

        $label = $key.PadRight($phaseLabelWidth)
        $line = "$icon $($script:cWhite)$label$($script:cReset)$($script:cMuted)$detail$($script:cReset)"
        Write-BoxLine -Content $line -Width $w
    }

    Write-BoxEmptyLine -Width $w; Write-BoxDivider -Width $w; Write-BoxEmptyLine -Width $w

    $total = $Phases.Count
    $phasesLabel = "PHASES".PadRight($phaseLabelWidth)
    $phasesVal = "$passCount/$total passed"
    $phasesLine = "$($script:cWhite)$($script:cBold)$phasesLabel$($script:cReset)$($script:cWhite)$phasesVal$($script:cReset)"
    Write-BoxLine -Content $phasesLine -Width $w

    $statusLabel = "STATUS".PadRight($phaseLabelWidth)
    if ($failCount -gt 0) {
        $statusIcon = "$($script:cRed)✗$($script:cReset)"; $statusText = "$($script:cRed)BLOCKED$($script:cReset)"
    } elseif ($warnCount -gt 0) {
        $statusIcon = "$($script:cYellow)⚠$($script:cReset)"; $statusText = "$($script:cYellow)REVIEW$($script:cReset)"
    } else {
        $statusIcon = "$($script:cLime)✓$($script:cReset)"; $statusText = "$($script:cLime)READY TO COMMIT$($script:cReset)"
    }
    $statusLine = "$($script:cWhite)$($script:cBold)$statusLabel$($script:cReset)$statusIcon $statusText"
    Write-BoxLine -Content $statusLine -Width $w

    Write-BoxEmptyLine -Width $w
    Write-BoxBottom -Width $w
}

Export-ModuleMember -Function @(
    'Register-Phase', 'Reset-Phases',
    'Get-PhaseIcon', 'Get-IconVisualWidth',
    'Write-PhaseStart', 'Write-PhaseEnd', 'Write-PhaseSummaryBox'
)
