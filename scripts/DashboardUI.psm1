# ─────────────────────────────────────────────────────────────────────────────
# DashboardUI.psm1 — Thin facade: composite dashboard renderer + re-exports
#
# Usage:
#   Import-Module ./scripts/DashboardUI.psm1 -Force
#
# Dependencies: DashboardTheme, DashboardBoxes, DashboardPhases, DashboardCoverage
# ─────────────────────────────────────────────────────────────────────────────

function Write-Dashboard {
    <#
    .SYNOPSIS
        Render the full dashboard from a data hashtable.
    #>
    [CmdletBinding()]
    param([Parameter(Mandatory)][hashtable]$Data)

    Write-Host ""

    # Header
    Write-DashboardHeader -Data $Data
    Write-Host ""

    # Scan summary
    if ($null -ne $Data.IssuesFound) { Write-ScanSummary -Data $Data; Write-Host "" }

    # Score box
    if ($Data.Scores -and $Data.Scores.Count -gt 0) { Write-ScoreBox -Data $Data; Write-Host "" }

    # Phase summary
    $phases = if ($Data.Phases) { $Data.Phases } else { $script:Phases }
    if ($phases -and $phases.Count -gt 0) { Write-PhaseSummaryBox -Phases $phases; Write-Host "" }

    # Coverage table
    if ($Data.CoverageData -and $Data.CoverageData.Count -gt 0) { Write-CoverageTable -CoverageData $Data.CoverageData; Write-Host "" }

    # Blocked details
    if ($Data.BlockedDetails -and $Data.BlockedDetails.Count -gt 0) { Write-BlockedDetails -BlockedDetails $Data.BlockedDetails; Write-Host "" }

    # Resolution summary
    if ($null -ne $Data.IssuesFixed -or $null -ne $Data.ManualTodos) { Write-ResolutionSummary -Data $Data; Write-Host "" }

    # Footer
    $tagline = if ($Data.Tagline) { $Data.Tagline } else { "Ship it. One command. Production-ready." }
    Write-FooterTagline -Text $tagline
    Write-Host ""
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Exports — re-export everything from sub-modules for backward compat
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Write-Dashboard'
)
