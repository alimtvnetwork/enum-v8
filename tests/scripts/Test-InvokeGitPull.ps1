# Smoke test for S-112 + S-113 fixes to Invoke-GitPull.
# Runs in a temporary git repo with no remote, verifies skip-with-no-origin path.
# Usage: pwsh -File tests/scripts/Test-InvokeGitPull.ps1

$ErrorActionPreference = 'Stop'
$repoRoot = (Resolve-Path "$PSScriptRoot/../..").Path

# Stub out helper functions normally provided by Utilities.psm1 + ErrorParser.psm1
function Write-Header { param($m) Write-Host "  ▶ $m" }
function Write-Success { param($m) Write-Host "  ✓ $m" }
function Write-Fail { param($m) Write-Host "  ✗ $m" }
function Get-CallerSource { return 'test-stub' }

Import-Module "$repoRoot/scripts/TestRunnerCore.psm1" -Force -DisableNameChecking

$tmp = Join-Path ([System.IO.Path]::GetTempPath()) ("git-pull-test-" + [guid]::NewGuid().ToString('N').Substring(0,8))
New-Item -ItemType Directory -Path $tmp | Out-Null
Push-Location $tmp
try {
    git init -q
    git config user.email "test@example.com"
    git config user.name "test"
    "x" | Out-File -FilePath README.md
    git add README.md; git commit -q -m "init"

    Write-Host "`n--- Test 1: no origin remote ---"
    $r1 = Invoke-GitPull
    if ($r1.Status -ne 'skip') { throw "Expected Status=skip, got '$($r1.Status)'" }
    if ($r1.Message -notmatch 'no origin') { throw "Expected 'no origin' message, got '$($r1.Message)'" }
    Write-Host "PASS — Status=$($r1.Status) Message=$($r1.Message)"

    Write-Host "`n--- Test 2: bogus origin (unreachable) ---"
    git remote add origin https://github.com/this-org-does-not-exist-12345/nope.git
    $r2 = Invoke-GitPull
    if ($r2.Status -ne 'skip') { throw "Expected Status=skip for unreachable remote, got '$($r2.Status)'" }
    Write-Host "PASS — Status=$($r2.Status) Message=$($r2.Message)"

    Write-Host "`n--- All tests passed ---"
} finally {
    Pop-Location
    Remove-Item -Recurse -Force $tmp
}
