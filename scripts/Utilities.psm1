# ─────────────────────────────────────────────────────────────────────────────
# Utilities.psm1 — Console output wrappers, test-log dir, line filtering/merging
#
# Usage:
#   Import-Module ./scripts/Utilities.psm1 -Force
#
# Dependencies: DashboardUI.psm1 (optional — gracefully degrades)
# ─────────────────────────────────────────────────────────────────────────────

$script:ESC    = [char]27
$script:cReset = "$script:ESC[0m"

if (-not $script:cLime)  { $script:cLime  = "$script:ESC[38;2;163;230;53m" }
if (-not $script:cRed)   { $script:cRed   = "$script:ESC[38;2;244;63;94m" }
if (-not $script:cCyan)  { $script:cCyan  = "$script:ESC[38;2;6;182;212m" }
if (-not $script:cWhite) { $script:cWhite = "$script:ESC[38;2;255;255;255m" }

function Write-Header {
    <#
    .SYNOPSIS
        Print a phase-start header line.
    #>
    [CmdletBinding()]
    param([string]$msg)

    if (Get-Command Write-DashboardHeader -ErrorAction SilentlyContinue) {
        Write-Host ""
        Write-PhaseStart -Name $msg
    } else {
        Write-Host "`n=== $msg ===" -ForegroundColor Cyan
    }
}

function Write-Success {
    <#
    .SYNOPSIS
        Print a green success line with a ✓ icon.
    #>
    [CmdletBinding()]
    param([string]$msg)
    Write-Host "  $($script:cLime)✓$($script:cReset) $($script:cLime)$msg$($script:cReset)"
}

function Write-Fail {
    <#
    .SYNOPSIS
        Print a red failure line with a ✗ icon.
    #>
    [CmdletBinding()]
    param([string]$msg)
    Write-Host "  $($script:cRed)✗$($script:cReset) $($script:cRed)$msg$($script:cReset)"
}

function Ensure-TestLogDir {
    <#
    .SYNOPSIS
        Create the test-log output directory if it doesn't exist.
    #>
    [CmdletBinding()]
    param()
    if (-not (Test-Path $global:TestLogDir)) {
        New-Item -ItemType Directory -Path $global:TestLogDir -Force | Out-Null
    }
}

function Get-CallerSource {
    <#
    .SYNOPSIS
        Return the calling module and function name from the call stack for error attribution.
    .DESCRIPTION
        Walks Get-PSCallStack to find the first caller outside Utilities.psm1.
        Returns a string like "CoverageRunner.psm1 → Invoke-TestCoverage" or
        "TestRunnerCore.psm1 → Invoke-BuildCheck" so error reports can trace
        which PowerShell module triggered the failure.
    .EXAMPLE
        $source = Get-CallerSource
        # Returns: "TestRunnerCore.psm1 → Invoke-BuildCheck"
    #>
    [CmdletBinding()]
    [OutputType([string])]
    param()

    $stack = Get-PSCallStack
    foreach ($frame in $stack) {
        $scriptName = if ($frame.ScriptName) { Split-Path $frame.ScriptName -Leaf } else { "" }
        $funcName = $frame.FunctionName
        # Skip this function, internal PS frames, and the Utilities module itself
        if ($funcName -eq "Get-CallerSource") { continue }
        if ($funcName -eq "<ScriptBlock>") { continue }
        if ($scriptName -eq "Utilities.psm1") { continue }
        if ($scriptName -and $funcName) {
            return "$scriptName → $funcName"
        }
        if ($scriptName) {
            return "$scriptName"
        }
        if ($funcName -and $funcName -ne "<ScriptBlock>") {
            return "$funcName"
        }
    }
    return "unknown"
}

function Filter-TestWarnings {
    <#
    .SYNOPSIS
        Remove "no packages being tested" warnings from Go test output.
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)

    return $lines | Where-Object {
        $_ -notmatch '^\s*warning: no packages being tested depend on matches for pattern'
    }
}

function Merge-UniqueOutputLines {
    <#
    .SYNOPSIS
        Merge two string arrays, deduplicating by exact content.
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$primary, [string[]]$secondary)

    $merged = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)

    foreach ($line in @($primary + $secondary)) {
        if ($null -eq $line) { continue }
        $normalized = $line.ToString().TrimEnd("`r")
        if (-not $normalized) { continue }
        if ($seen.Add($normalized)) { $merged.Add($normalized) | Out-Null }
    }

    return $merged.ToArray()
}

Export-ModuleMember -Function @(
    'Write-Header', 'Write-Success', 'Write-Fail',
    'Ensure-TestLogDir', 'Filter-TestWarnings', 'Merge-UniqueOutputLines',
    'Get-CallerSource'
)
