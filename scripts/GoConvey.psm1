# ─────────────────────────────────────────────────────────────────────────────
# GoConvey.psm1 — Launch the GoConvey browser-based test runner
#
# Usage:
#   Import-Module ./scripts/GoConvey.psm1 -Force
#
# Dependencies: Utilities.psm1 (Write-Header, Write-Success, Write-Fail)
# ─────────────────────────────────────────────────────────────────────────────

function Invoke-GoConvey {
    <#
    .SYNOPSIS
        Launch GoConvey, a browser-based Go test runner with live reload.
    .DESCRIPTION
        Installs GoConvey if not present, then starts it on the specified
        port (default 8080) from the tests/ directory.
    .PARAMETER ExtraArgs
        Optional array; first element is the port number.
    .EXAMPLE
        Invoke-GoConvey               # default port 8080
        Invoke-GoConvey @("9090")     # custom port
    #>
    [CmdletBinding()]
    param([string[]]$ExtraArgs)

    Write-Header "Launching GoConvey"

    # Check if goconvey is installed
    $gcPath = Get-Command goconvey -ErrorAction SilentlyContinue
    if (-not $gcPath) {
        Write-Host "  GoConvey not found. Installing..." -ForegroundColor Yellow
        go install github.com/smartystreets/goconvey@latest
        if ($LASTEXITCODE -ne 0) {
            $s = Get-CallerSource; Write-Fail "Failed to install GoConvey (source: $s)"
            return
        }
        Write-Success "GoConvey installed"
    }

    $port = if ($ExtraArgs -and $ExtraArgs[0]) { $ExtraArgs[0] } else { "8080" }
    Write-Host "  Starting GoConvey on http://localhost:$port" -ForegroundColor Yellow
    Write-Host "  Press Ctrl+C to stop" -ForegroundColor Gray

    Push-Location tests
    try {
        goconvey -port $port
    }
    finally { Pop-Location }
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Invoke-GoConvey'
)
