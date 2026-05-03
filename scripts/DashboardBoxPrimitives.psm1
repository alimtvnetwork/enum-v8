# ─────────────────────────────────────────────────────────────────────────────
# DashboardBoxPrimitives.psm1 — Progress bar + box-drawing primitives
#
# Dependencies: DashboardTheme.psm1 (script-scope color variables)
# ─────────────────────────────────────────────────────────────────────────────

# Ambiguous-width overrides: code points that East Asian Width tables call
# "wide" but render as 1 column in many Western terminals (e.g. Windows
# Terminal with Cascadia Code, VS Code integrated terminal, iTerm2 default).
# Add/remove entries to match your environment.
$script:AmbiguousWidthOverrides = @{
    0x2713 = 1   # ✓ CHECK MARK
    0x2714 = 1   # ✔ HEAVY CHECK MARK
    0x2717 = 1   # ✗ BALLOT X
    0x2718 = 1   # ✘ HEAVY BALLOT X
    0x26A0 = 1   # ⚠ WARNING SIGN
    0x2716 = 1   # ✖ HEAVY MULTIPLICATION X
}

function Set-AmbiguousWidthOverride {
    <# .SYNOPSIS Add or update a code-point width override.
       .PARAMETER CodePoint Unicode code point (integer, e.g. 0x2713).
       .PARAMETER Width     Visual column width (1 or 2).
       .EXAMPLE Set-AmbiguousWidthOverride -CodePoint 0x26A0 -Width 2 #>
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)][int]$CodePoint,
        [Parameter(Mandatory)][ValidateSet(1,2)][int]$Width
    )
    $script:AmbiguousWidthOverrides[$CodePoint] = $Width
}

function Remove-AmbiguousWidthOverride {
    <# .SYNOPSIS Remove a code-point width override so it falls back to range tables.
       .PARAMETER CodePoint Unicode code point to remove. #>
    [CmdletBinding()]
    param([Parameter(Mandatory)][int]$CodePoint)
    if ($script:AmbiguousWidthOverrides.ContainsKey($CodePoint)) {
        $script:AmbiguousWidthOverrides.Remove($CodePoint)
    }
}

function Get-AmbiguousWidthOverrides {
    <# .SYNOPSIS Return the current override table (read-only copy). #>
    [CmdletBinding()]
    [OutputType([hashtable])]
    param()
    return $script:AmbiguousWidthOverrides.Clone()
}

function Test-IsWideVisualCodePoint {
    [CmdletBinding()]
    [OutputType([bool])]
    param([Parameter(Mandatory)][int]$CodePoint)

    # Check override table first — fast path for ambiguous characters
    if ($script:AmbiguousWidthOverrides.ContainsKey($CodePoint)) {
        return ($script:AmbiguousWidthOverrides[$CodePoint] -eq 2)
    }

    return (
        ($CodePoint -ge 0x1100 -and $CodePoint -le 0x115F) -or
        ($CodePoint -ge 0x2329 -and $CodePoint -le 0x232A) -or
        ($CodePoint -ge 0x2600 -and $CodePoint -le 0x26FF) -or   # Misc Symbols (⚠ ⚡ etc.)
        ($CodePoint -ge 0x2700 -and $CodePoint -le 0x27BF) -or   # Dingbats (✓ ✗ etc.)
        ($CodePoint -ge 0x2E80 -and $CodePoint -le 0xA4CF) -or
        ($CodePoint -ge 0xAC00 -and $CodePoint -le 0xD7A3) -or
        ($CodePoint -ge 0xF900 -and $CodePoint -le 0xFAFF) -or
        ($CodePoint -ge 0xFE10 -and $CodePoint -le 0xFE19) -or
        ($CodePoint -ge 0xFE30 -and $CodePoint -le 0xFE6F) -or
        ($CodePoint -ge 0xFF01 -and $CodePoint -le 0xFF60) -or
        ($CodePoint -ge 0xFFE0 -and $CodePoint -le 0xFFE6) -or
        ($CodePoint -ge 0x1F300 -and $CodePoint -le 0x1FAFF)
    )
}

function Get-AnsiVisualLength {
    <# .SYNOPSIS Calculate visual column width of a string, stripping ANSI escape codes and accounting for wide Unicode chars. #>
    [CmdletBinding()]
    [OutputType([int])]
    param([Parameter(Mandatory)][string]$Text)

    # Strip ANSI escape sequences: ESC[...m and ESC[...;...m etc.
    $stripped = $Text -replace "$([char]27)\[[0-9;]*m", ''
    $len = 0
    $enumerator = [System.Globalization.StringInfo]::GetTextElementEnumerator($stripped)

    while ($enumerator.MoveNext()) {
        $element = [string]$enumerator.GetTextElement()

        if ([string]::IsNullOrEmpty($element)) {
            continue
        }

        $cp = [char]::ConvertToUtf32($element, 0)

        if (Test-IsWideVisualCodePoint -CodePoint $cp) {
            $len += 2
        } else {
            $len += 1
        }
    }

    return $len
}

function Get-ProgressBar {
    <# .SYNOPSIS Returns a colored progress bar string using ANSI + Unicode block chars. #>
    [CmdletBinding()]
    [OutputType([string])]
    param([Parameter(Mandatory)][int]$Score, [int]$MaxScore = 100, [int]$BarWidth = 15)

    if ($MaxScore -le 0) { $MaxScore = 100 }
    $percentage  = [math]::Min(1.0, [math]::Max(0.0, $Score / $MaxScore))
    $filledCount = [math]::Round($percentage * $BarWidth)
    $emptyCount  = $BarWidth - $filledCount
    $filled = if ($filledCount -gt 0) { [string]::new([char]0x2588, $filledCount) } else { "" }
    $empty  = if ($emptyCount  -gt 0) { [string]::new([char]0x2592, $emptyCount)  } else { "" }
    return "$($script:cLime)$filled$($script:cBarE)$empty$($script:cReset)"
}

function Resolve-DashboardBoxWidth {
    [CmdletBinding()]
    [OutputType([int])]
    param([int]$Width = 0)

    if ($Width -gt 0) {
        return $Width
    }

    if (Get-Command Get-DashboardBoxWidth -ErrorAction SilentlyContinue) {
        return [int](Get-DashboardBoxWidth)
    }

    if ($script:BoxWidth -is [int] -and $script:BoxWidth -gt 0) {
        return [int]$script:BoxWidth
    }

    return 48
}

function Write-BoxTop {
    [CmdletBinding()]
    param([int]$Width = 0)
    $Width = Resolve-DashboardBoxWidth -Width $Width
    Write-Host "$($script:cBorder)╔$("═" * $Width)╗$($script:cReset)"
}
 
function Write-BoxBottom {
    [CmdletBinding()]
    param([int]$Width = 0)
    $Width = Resolve-DashboardBoxWidth -Width $Width
    Write-Host "$($script:cBorder)╚$("═" * $Width)╝$($script:cReset)"
}
 
function Write-BoxDivider {
    [CmdletBinding()]
    param([int]$Width = 0)
    $Width = Resolve-DashboardBoxWidth -Width $Width
    Write-Host "$($script:cBorder)╠$("═" * $Width)╣$($script:cReset)"
}
 
function Write-BoxEmptyLine {
    [CmdletBinding()]
    param([int]$Width = 0)
    $Width = Resolve-DashboardBoxWidth -Width $Width
    Write-Host "$($script:cBorder)║$($script:cReset)$(" " * $Width)$($script:cBorder)║$($script:cReset)"
}
 
function Write-BoxLine {
    [CmdletBinding()]
    param([string]$Content, [int]$Width = 0, [int]$VisualLength = -1)
    $Width = Resolve-DashboardBoxWidth -Width $Width
    if ($VisualLength -lt 0) { $VisualLength = Get-AnsiVisualLength -Text $Content }
    $rightPad = [math]::Max(0, $Width - $VisualLength - 1)
    Write-Host "$($script:cBorder)║$($script:cReset) $Content$(" " * $rightPad)$($script:cBorder)║$($script:cReset)"
}
 
function Write-BoxLineCenter {
    [CmdletBinding()]
    param([string]$Text, [int]$Width = 0, [string]$Color = "")
    $Width = Resolve-DashboardBoxWidth -Width $Width
    if (-not $Color) { $Color = $script:cWhite }
    $textLen = Get-AnsiVisualLength -Text $Text
    $leftPad = [math]::Max(0, [math]::Floor(($Width - $textLen) / 2))
    $rightPad = [math]::Max(0, $Width - $textLen - $leftPad)
    $line = (" " * $leftPad) + $Text + (" " * $rightPad)
    Write-Host "$($script:cBorder)║$($script:cReset)$Color$($script:cBold)$line$($script:cReset)$($script:cBorder)║$($script:cReset)"
}
 
Export-ModuleMember -Function @(
    'Get-AnsiVisualLength',
    'Set-AmbiguousWidthOverride', 'Remove-AmbiguousWidthOverride', 'Get-AmbiguousWidthOverrides',
    'Get-ProgressBar', 'Write-BoxTop', 'Write-BoxBottom', 'Write-BoxDivider',
    'Write-BoxEmptyLine', 'Write-BoxLine', 'Write-BoxLineCenter'
)
