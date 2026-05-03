# ─────────────────────────────────────────────────────────────────────────────
# DashboardTheme.psm1 — Theme detection, ANSI color initialization
#
# Usage:
#   Import-Module ./scripts/DashboardTheme.psm1 -Force
#
# Dependencies: None (standalone)
# ─────────────────────────────────────────────────────────────────────────────

# ═══════════════════════════════════════════════════════════════════════════════
# Environment Setup
# ═══════════════════════════════════════════════════════════════════════════════

$script:ESC    = [char]27
$script:cReset = "$script:ESC[0m"
$script:cBold  = "$script:ESC[1m"
$script:cDim   = "$script:ESC[2m"
$script:BoxWidth = 48

# ═══════════════════════════════════════════════════════════════════════════════
# Theme Detection
# ═══════════════════════════════════════════════════════════════════════════════

function Get-TerminalTheme {
    [CmdletBinding()]
    [OutputType([string])]
    param()

    if ($env:DASHBOARD_THEME) { return $env:DASHBOARD_THEME.ToLower() }

    $isInteractiveConsole = $false
    try {
        $isInteractiveConsole =
            [Environment]::UserInteractive -and
            $null -ne [Console]::In -and
            $null -ne [Console]::Out -and
            -not [Console]::IsInputRedirected -and
            -not [Console]::IsOutputRedirected
    } catch { }

    if ($env:WT_SESSION -and $env:LOCALAPPDATA) {
        $wtSettings = Join-Path $env:LOCALAPPDATA "Packages\Microsoft.WindowsTerminal_8wekyb3d8bbwe\LocalState\settings.json"
        if (Test-Path $wtSettings) {
            try {
                $json = Get-Content $wtSettings -Raw | ConvertFrom-Json
                $schemeName = $json.profiles.defaults.colorScheme
                if (-not $schemeName -and $json.defaultProfile) {
                    $schemeName = $json.profiles.list |
                        Where-Object { $_.guid -eq $json.defaultProfile } |
                        Select-Object -ExpandProperty colorScheme -ErrorAction SilentlyContinue
                }
                if ($schemeName) {
                    $scheme = $json.schemes | Where-Object { $_.name -eq $schemeName }
                    if ($scheme -and $scheme.background) {
                        $bg = $scheme.background -replace '^#', ''
                        $r = [convert]::ToInt32($bg.Substring(0,2), 16)
                        $g = [convert]::ToInt32($bg.Substring(2,2), 16)
                        $b = [convert]::ToInt32($bg.Substring(4,2), 16)
                        $luminance = (0.2126 * $r + 0.7152 * $g + 0.0722 * $b) / 255
                        return $(if ($luminance -lt 0.5) { "dark" } else { "light" })
                    }
                }
            } catch { }
        }
    }

    if (($IsLinux -or $IsMacOS) -and $isInteractiveConsole) {
        $sttyOld = $null
        try {
            $sttyOld = (& stty -g 2>$null | Out-String).Trim()
            if (-not $sttyOld) { throw 'stty state unavailable' }

            & stty raw -echo min 0 time 1 2>$null | Out-Null
            [Console]::Write("$([char]27)]11;?$([char]27)\")
            [Console]::Out.Flush()
            Start-Sleep -Milliseconds 100

            $response = ""
            while ([Console]::KeyAvailable) { $response += [char][Console]::Read() }
            if ($response -match 'rgb:([0-9a-f]{2,4})/([0-9a-f]{2,4})/([0-9a-f]{2,4})') {
                $r = [convert]::ToInt32($Matches[1].Substring(0,2), 16)
                $g = [convert]::ToInt32($Matches[2].Substring(0,2), 16)
                $b = [convert]::ToInt32($Matches[3].Substring(0,2), 16)
                $luminance = (0.2126 * $r + 0.7152 * $g + 0.0722 * $b) / 255
                return $(if ($luminance -lt 0.5) { "dark" } else { "light" })
            }
        } catch { }
        finally {
            if ($sttyOld) {
                & stty $sttyOld 2>$null | Out-Null
            }
        }
    }

    try {
        $bg = $Host.UI.RawUI.BackgroundColor
        if ($bg -in @("White", "Gray", "Yellow", "Cyan")) { return "light" }
    } catch { }

    return "dark"
}

function Set-ThemeColors {
    [CmdletBinding()]
    param([string]$Theme = "dark")

    $e = $script:ESC
    if ($Theme -eq "light") {
        $script:cLime   = "$e[38;2;22;163;74m"
        $script:cRed    = "$e[38;2;185;28;28m"
        $script:cPurple = "$e[38;2;109;40;217m"
        $script:cCyan   = "$e[38;2;14;116;144m"
        $script:cYellow = "$e[38;2;161;98;7m"
        $script:cMuted  = "$e[38;2;107;114;128m"
        $script:cWhite  = "$e[38;2;15;23;42m"
        $script:cBarE   = "$e[38;2;209;213;219m"
        $script:cBorder = "$e[38;2;156;163;175m"
    } else {
        $script:cLime   = "$e[38;2;163;230;53m"
        $script:cRed    = "$e[38;2;244;63;94m"
        $script:cPurple = "$e[38;2;168;85;247m"
        $script:cCyan   = "$e[38;2;6;182;212m"
        $script:cYellow = "$e[38;2;250;204;21m"
        $script:cMuted  = "$e[38;2;156;163;175m"
        $script:cWhite  = "$e[38;2;255;255;255m"
        $script:cBarE   = "$e[38;2;100;100;100m"
        $script:cBorder = "$e[38;2;156;163;175m"
    }
}

function Get-DashboardBoxWidth {
    [CmdletBinding()]
    [OutputType([int])]
    param()

    if ($script:BoxWidth -is [int] -and $script:BoxWidth -gt 0) {
        return [int]$script:BoxWidth
    }

    return 48
}

function Initialize-DashboardUI {
    <#
    .SYNOPSIS
        Initialize the dashboard module: UTF-8 encoding + theme colors.
    .PARAMETER Theme
        Force "dark" or "light". Omit to auto-detect.
    #>
    [CmdletBinding()]
    param([string]$Theme)

    [console]::OutputEncoding = [System.Text.Encoding]::UTF8
    if (-not $Theme) { $Theme = Get-TerminalTheme }
    $script:CurrentTheme = $Theme
    Set-ThemeColors $Theme
}

function Test-DashboardTheme {
    <#
    .SYNOPSIS
        Render a color swatch for both themes to visually verify contrast.
    #>
    [CmdletBinding()]
    param()

    foreach ($theme in @("dark", "light")) {
        Set-ThemeColors $theme
        Write-Host ""
        Write-Host "$($script:cBold)=== Theme: $theme ===$($script:cReset)"
        Write-Host "  $($script:cLime)✓ Success / Lime$($script:cReset)"
        Write-Host "  $($script:cRed)✗ Error / Red$($script:cReset)"
        Write-Host "  $($script:cPurple)● Purple / Todo$($script:cReset)"
        Write-Host "  $($script:cCyan)▶ Cyan / Info$($script:cReset)"
        Write-Host "  $($script:cYellow)⚠ Yellow / Warning$($script:cReset)"
        Write-Host "  $($script:cMuted)Muted text$($script:cReset)"
        Write-Host "  $($script:cWhite)Primary text$($script:cReset)"
        Write-Host "  Bar: $(Get-ProgressBar -Score 73)"
        Write-Host ""
    }
    Set-ThemeColors $script:CurrentTheme
}

Export-ModuleMember -Function @(
    'Initialize-DashboardUI',
    'Get-TerminalTheme',
    'Set-ThemeColors',
    'Get-DashboardBoxWidth',
    'Test-DashboardTheme'
)
