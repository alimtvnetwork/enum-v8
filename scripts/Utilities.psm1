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

function Test-IsCoverpkgWarningOnlyOutput {
    <#
    .SYNOPSIS
        Returns $true when go test output contains ONLY harmless
        `warning: no packages being tested depend on matches for pattern …`
        lines (and blank lines / "ok"/"PASS" markers). These are emitted by
        `-coverpkg=` for transitive packages a test binary doesn't import,
        and must NEVER be classified as a build error or runtime failure.
    .NOTES
        Root cause (RCA, 2026-05-07): the parallel compile-check probe
        sometimes saw a non-zero exit even though the only diagnostic was
        a stream of these warnings — producing false-positive Blocked
        reports for licensetype / onofftype / rootcmdnames and a phantom
        "runtime failure" for brackets. Filtering at decision time fixes
        both surfaces.
    #>
    [CmdletBinding()]
    [OutputType([bool])]
    param([string[]]$lines)

    if (-not $lines -or $lines.Count -eq 0) { return $false }
    $sawWarning = $false
    foreach ($raw in $lines) {
        if ($null -eq $raw) { continue }
        $line = $raw.ToString().TrimEnd("`r").Trim()
        if (-not $line) { continue }
        if ($line -match '^warning: no packages being tested depend on matches for pattern') {
            $sawWarning = $true
            continue
        }
        # PASS/ok markers are noise we tolerate
        if ($line -match '^(PASS|ok\s|FAIL\s+\S+\s+\[(setup failed|build failed)\])') { continue }
        # any other content => not warnings-only
        return $false
    }
    return $sawWarning
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

function Resolve-TestSuiteRoot {
    <#
    .SYNOPSIS
        Resolve the test-suite root directory name (creationtests vs legacy integratedtests).
    .DESCRIPTION
        Per Core-memory rule: "Tests live under tests/creationtests/, NOT tests/integratedtests/.
        Tooling that probes test packages must accept either name (or read from disk) — never
        hard-code one." This helper returns the first existing root, preferring 'creationtests'.
    .PARAMETER ProjectRoot
        The repository root. Defaults to $global:ProjectRoot.
    .PARAMETER Package
        Optional package name. When supplied, only roots that actually contain
        tests/<root>/<Package> are considered a match.
    .OUTPUTS
        [string] One of 'creationtests' or 'integratedtests'. Returns 'creationtests' as
        the default fallback when neither exists, so the downstream `go test` error is the
        user-facing diagnostic ("no Go files in ...").
    .EXAMPLE
        $root = Resolve-TestSuiteRoot                     # → 'creationtests'
        $root = Resolve-TestSuiteRoot -Package 'osdetect' # → 'creationtests' if osdetect exists there
    #>
    [CmdletBinding()]
    param(
        [string]$ProjectRoot = $global:ProjectRoot,
        [string]$Package = ''
    )

    $candidates = @('creationtests', 'integratedtests')
    $testsDir   = Join-Path $ProjectRoot 'tests'

    foreach ($candidate in $candidates) {
        $rootPath = Join-Path $testsDir $candidate
        if (-not (Test-Path $rootPath)) { continue }
        if ($Package) {
            $pkgPath = Join-Path $rootPath $Package
            if (Test-Path $pkgPath) { return $candidate }
        } else {
            return $candidate
        }
    }

    return 'creationtests'
}

Export-ModuleMember -Function @(
    'Write-Header', 'Write-Success', 'Write-Fail',
    'Ensure-TestLogDir', 'Filter-TestWarnings', 'Test-IsCoverpkgWarningOnlyOutput',
    'Merge-UniqueOutputLines', 'Get-CallerSource', 'Resolve-TestSuiteRoot'
)
