# ─────────────────────────────────────────────────────────────────────────────
# Help.psm1 — Help display, fail log viewer, and integrated tests
#
# Usage:
#   Import-Module ./scripts/Help.psm1 -Force
#
# Dependencies:
#   - Utilities.psm1  (Write-Header, Write-Success, Write-Fail, Filter-TestWarnings)
#   - TestLogWriter.psm1 (Write-TestLogs)
#   - TestRunner.psm1 (Invoke-FetchLatest, Invoke-BuildCheck, Open-FailingTestsIfAny)
# ─────────────────────────────────────────────────────────────────────────────

function Invoke-IntegratedTests {
    <#
    .SYNOPSIS
        Run only the integrated test suite (tests/integratedtests/...).
    .EXAMPLE
        Invoke-IntegratedTests
    #>
    [CmdletBinding()]
    param()

    Write-Header "Running integrated tests only"
    Invoke-FetchLatest
    Push-Location tests
    try {
        if (-not (Invoke-BuildCheck "./integratedtests/...")) { return }

        $prevPref = $ErrorActionPreference
        $ErrorActionPreference = "Continue"
        $output = & go test -v -count=1 ./integratedtests/... 2>&1 | ForEach-Object { $_.ToString() }
        $exitCode = $LASTEXITCODE
        $ErrorActionPreference = $prevPref

        Filter-TestWarnings $output | ForEach-Object { Write-Host $_ }
        Write-TestLogs $output

        if ($exitCode -eq 0) { Write-Success "Integrated tests passed" }
        else { $s = Get-CallerSource; Write-Fail "Integrated tests failed (exit code: $exitCode) (source: $s)" }
    }
    finally { Pop-Location }
    Open-FailingTestsIfAny
}

function Invoke-ShowFailLog {
    <#
    .SYNOPSIS
        Display the contents of the last failing-tests log file.
    .EXAMPLE
        Invoke-ShowFailLog
    #>
    [CmdletBinding()]
    param()

    $failingFile = Join-Path $global:TestLogDir "failing-tests.txt"
    if (-not (Test-Path $failingFile)) {
        Write-Header "No failing tests log found"
        Write-Host "  Run tests first: ./run.ps1 T" -ForegroundColor Yellow
        return
    }

    Write-Header "Last Failing Tests"
    $content = Get-Content $failingFile -Raw
    if ($content -match '# Count: 0') {
        Write-Success "No failing tests in last run"
    }
    else {
        Write-Host $content
    }
    Write-Host ""
    Write-Host "  Log file: $failingFile" -ForegroundColor Gray
}

function Show-Help {
    <#
    .SYNOPSIS
        Display usage help for all run.ps1 commands.
    .EXAMPLE
        Show-Help
    #>
    [CmdletBinding()]
    param()

    Write-Host ""
    Write-Host "  Project Runner — ./run.ps1 <command> [options]" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "  Testing:" -ForegroundColor Yellow
    Write-Host "    T   | -t   | test          Run all tests (verbose)"
    Write-Host "    TP  | -tp  | test-pkg      Run tests for a specific package"
    Write-Host "    TC  | -tc  | test-cover    Run tests with coverage (HTML + summary)"
    Write-Host "    TCP | -tcp | test-cover-pkg Run coverage for a specific package"
    Write-Host "    TI  | -ti  | test-int      Run integrated tests only"
    Write-Host "    TF  | -tf  | test-fail     Show last failing tests log"
    Write-Host "    GC  | -gc  | goconvey      Launch GoConvey (browser test runner)"
    Write-Host ""
    Write-Host "  Build & Run:" -ForegroundColor Yellow
    Write-Host "    R   | -r   | run           Run the main application"
    Write-Host "    B   | -b   | build         Build the binary"
    Write-Host "    BR  | -br  | build-run     Build then run"
    Write-Host ""
    Write-Host "  Code Quality:" -ForegroundColor Yellow
    Write-Host "    F   | -f   | fmt           Format all Go files"
    Write-Host "    L   | -l   | lint          Run go vet"
    Write-Host "    V   | -v   | vet           Run go vet"
    Write-Host "    TY  | -ty  | tidy          Run go mod tidy"
    Write-Host "    PC  | -pc  | pre-commit    Check Coverage* files for API mismatches"
    Write-Host ""
    Write-Host "  Other:" -ForegroundColor Yellow
    Write-Host "    C   | -c   | clean         Clean build artifacts"
    Write-Host "    H   | -h   | help          Show this help"
    Write-Host ""
    Write-Host "  Mode Options (for TC/TCP/PC):" -ForegroundColor Yellow
    Write-Host "    --sync      Run precompile + tests sequentially (default: parallel)"
    Write-Host "    --open      Open HTML coverage report in browser after TC/TCP"
    Write-Host ""
    Write-Host "  Examples:" -ForegroundColor Gray
    Write-Host "    ./run.ps1 T"
    Write-Host "    ./run.ps1 -t"
    Write-Host "    ./run.ps1 TP regexnewtests"
    Write-Host "    ./run.ps1 -tp regexnewtests"
    Write-Host "    ./run.ps1 TCP regexnewtests  (package coverage)"
    Write-Host "    ./run.ps1 TC                 (parallel by default)"
    Write-Host "    ./run.ps1 TC --sync          (sequential mode)"
    Write-Host "    ./run.ps1 TC --sync --no-open"
    Write-Host "    ./run.ps1 PC                 (pre-commit check)"
    Write-Host "    ./run.ps1 PC corejsontests   (check single package)"
    Write-Host "    ./run.ps1 -gc"
    Write-Host "    ./run.ps1 -gc 9090          (custom port)"
    Write-Host ""
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Invoke-IntegratedTests',
    'Invoke-ShowFailLog',
    'Show-Help'
)
