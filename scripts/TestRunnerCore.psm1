# ─────────────────────────────────────────────────────────────────────────────
# TestRunnerCore.psm1 — Git ops, build check, test invocation primitives
#
# Dependencies: Utilities.psm1, ErrorParser.psm1, TestLogWriter.psm1
# ─────────────────────────────────────────────────────────────────────────────

function Invoke-GitPull {
    <#
    .SYNOPSIS
        Pull the latest changes from the remote git repository.
    .DESCRIPTION
        S-113: probes the configured `origin` remote with `git ls-remote --exit-code origin`
        BEFORE attempting `git pull`, so a missing/unreachable remote produces a clean
        `skip` result instead of a confusing "Repository not found" + soft-fail.

        S-112: returns a structured status object so callers can register a truthful
        dashboard phase row instead of hard-coding `pass`.
    .OUTPUTS
        [pscustomobject] @{ Status = 'pass' | 'warn' | 'skip'; Message = '...' }
    #>
    [CmdletBinding()]
    param()
    Write-Header "Pulling latest from remote"
    $prevPref = $ErrorActionPreference; $ErrorActionPreference = "Continue"

    # S-113 — early remote probe. If `origin` doesn't exist or is unreachable,
    # don't attempt `git pull` (which would emit a misleading "Repository not found").
    $hasOrigin = $false
    $originUrl = (git remote get-url origin 2>$null)
    if ($LASTEXITCODE -eq 0 -and $originUrl) { $hasOrigin = $true }
    if (-not $hasOrigin) {
        Write-Host "  No 'origin' remote configured — skipping pull" -ForegroundColor Yellow
        $ErrorActionPreference = $prevPref
        return [pscustomobject]@{ Status = 'skip'; Message = 'no origin remote' }
    }

    # Probe reachability without printing remote refs.
    $null = git ls-remote --exit-code origin HEAD 2>&1
    $reachable = ($LASTEXITCODE -eq 0)
    if (-not $reachable) {
        Write-Host "  Remote 'origin' ($originUrl) unreachable — skipping pull" -ForegroundColor Yellow
        $ErrorActionPreference = $prevPref
        return [pscustomobject]@{ Status = 'skip'; Message = 'remote unreachable' }
    }

    git pull 2>&1 | ForEach-Object { Write-Host "  $_" -ForegroundColor Gray }
    $pullExit = $LASTEXITCODE
    $ErrorActionPreference = $prevPref
    if ($pullExit -eq 0) {
        Write-Success "Git pull complete"
        return [pscustomobject]@{ Status = 'pass'; Message = 'pulled from remote' }
    }
    $s = Get-CallerSource
    Write-Fail "git pull failed (continuing anyway) (source: $s)"
    return [pscustomobject]@{ Status = 'warn'; Message = 'pull failed (continuing)' }
}

function Invoke-FetchLatest {
    <#
    .SYNOPSIS
        Pull git changes and run `go mod tidy` to sync dependencies.
    .OUTPUTS
        [pscustomobject] @{ GitPull = <Invoke-GitPull result>; Tidy = @{ Status; Message } }
    #>
    [CmdletBinding()]
    param()
    $gitResult = Invoke-GitPull
    Write-Header "Fetching latest dependencies"
    go mod tidy
    if ($LASTEXITCODE -eq 0) {
        Write-Success "Dependencies up to date"
        $tidy = [pscustomobject]@{ Status = 'pass'; Message = 'up to date' }
    } else {
        $s = Get-CallerSource
        Write-Fail "go mod tidy failed (source: $s)"
        $tidy = [pscustomobject]@{ Status = 'warn'; Message = 'tidy failed' }
    }
    return [pscustomobject]@{ GitPull = $gitResult; Tidy = $tidy }
}

function Invoke-BuildCheck {
    <# .SYNOPSIS Compile-check a Go package path before running tests. Returns $true/$false. #>
    [CmdletBinding()]
    param([string]$buildPath)
    Write-Header "Build check: $buildPath"
    $prevPref = $ErrorActionPreference; $ErrorActionPreference = "Continue"
    $output = & go build $buildPath 2>&1 | ForEach-Object { $_.ToString() }
    $exitCode = $LASTEXITCODE; $ErrorActionPreference = $prevPref

    if ($exitCode -ne 0) {
        $callerSource = Get-CallerSource
        Write-Fail "Build failed — skipping tests (source: $callerSource)"
        Ensure-TestLogDir
        $failingFile = Join-Path $global:TestLogDir "failing-tests.txt"
        $rawFile     = Join-Path $global:TestLogDir "raw-output.txt"
        $timestamp   = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
        $buildErrors = Extract-BuildErrorLines $output
        $errorCount = if ($buildErrors.Count -gt 0) { $buildErrors.Count } else { 1 }
        $failingContent = @("# Failing Tests — $timestamp", "# Count: $errorCount", "# Source: $callerSource (TestRunnerCore.psm1 → Invoke-BuildCheck)", "", "# Build Failed — tests were NOT run", "", "# ── Build Errors ──", "")
        if ($buildErrors.Count -gt 0) { $failingContent += $buildErrors } else { $failingContent += $output }
        Set-Content -Path $failingFile -Value ($failingContent -join "`n") -Encoding UTF8
        Set-Content -Path $rawFile -Value ($output -join "`n") -Encoding UTF8
        $output | ForEach-Object { Write-Host "  $_" -ForegroundColor Red }
        Open-FailingTestsIfAny
        return $false
    }
    Write-Success "Build OK"; return $true
}

function Invoke-GoTestAndLog {
    <# .SYNOPSIS Run `go test` with given args, print output, and write test logs. Returns exit code. #>
    [CmdletBinding()]
    param([string]$testArgs)
    $prevPref = $ErrorActionPreference; $ErrorActionPreference = "Continue"
    $output = & go test -v -count=1 $testArgs 2>&1 | ForEach-Object { $_.ToString() }
    $exitCode = $LASTEXITCODE; $ErrorActionPreference = $prevPref
    Filter-TestWarnings $output | ForEach-Object { Write-Host $_ }
    Write-TestLogs $output
    return $exitCode
}

function Open-FailingTestsIfAny {
    <# .SYNOPSIS Open the failing-tests log file if it contains failures. #>
    [CmdletBinding()]
    param()
    $failingFile = Join-Path $global:TestLogDir "failing-tests.txt"
    if ((Test-Path $failingFile)) {
        $content = Get-Content $failingFile -Raw
        if ($content -and $content -notmatch '# Count: 0') {
            Write-Host ""; Write-Host "  Opening failing tests log..." -ForegroundColor Yellow
            Invoke-Item $failingFile
        }
    }
}

Export-ModuleMember -Function @(
    'Invoke-GitPull', 'Invoke-FetchLatest', 'Invoke-BuildCheck',
    'Invoke-GoTestAndLog', 'Open-FailingTestsIfAny'
)
