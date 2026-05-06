#!/usr/bin/env pwsh
# S-115 smoke test for Test-UpstreamClone helper.
#
# Verifies four cases:
#   1. Missing path           → Ok=$false, Reason='missing'
#   2. Path exists but no sentinel → Ok=$false, Reason starts 'sentinel-missing'
#   3. Real /tmp/core-v9-upstream  → Ok=$true, PackageCount > 100 (skipped if absent)
#   4. Get-UpstreamPackages emits sentinel warning when given a fake dir

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$repoRoot = Split-Path -Parent (Split-Path -Parent $PSScriptRoot)
Import-Module (Join-Path $repoRoot 'scripts/spec-api-check.psm1') -Force

$failed = 0
function Assert($cond, $msg) {
    if ($cond) { Write-Host "  ✓ $msg" -ForegroundColor Green }
    else { Write-Host "  ✗ $msg" -ForegroundColor Red; $script:failed++ }
}

Write-Host '── S-115 Test-UpstreamClone ──' -ForegroundColor Cyan

# Case 1: missing path
$missing = "/tmp/__nope_$(Get-Random)"
$r1 = Test-UpstreamClone -Path $missing
Assert (-not $r1.Ok) "Case 1: Ok=false for missing path"
Assert ($r1.Reason -eq 'missing') "Case 1: Reason='missing' (got '$($r1.Reason)')"

# Case 2: path exists but no sentinel
$fake = "/tmp/__fake_clone_$(Get-Random)"
New-Item -ItemType Directory -Path $fake | Out-Null
New-Item -ItemType File -Path (Join-Path $fake 'README.md') -Value '# fake' | Out-Null
$r2 = Test-UpstreamClone -Path $fake
Assert (-not $r2.Ok) "Case 2: Ok=false for path without sentinel"
Assert ($r2.Reason -like 'sentinel-missing*') "Case 2: Reason='sentinel-missing*' (got '$($r2.Reason)')"

# Case 4: Get-UpstreamPackages warns when sentinel missing
$warns = @()
Get-UpstreamPackages -UpstreamDir $fake -WarningVariable warns -WarningAction SilentlyContinue | Out-Null
Assert ($warns.Count -ge 1) "Case 4: Get-UpstreamPackages emits warning when sentinel absent"
Remove-Item -Recurse -Force $fake

# Case 3: real clone (skip if absent)
$real = '/tmp/core-v9-upstream'
if (Test-Path (Join-Path $real 'coredata/coregeneric')) {
    $r3 = Test-UpstreamClone -Path $real
    Assert ($r3.Ok) "Case 3: Ok=true for real clone"
    Assert ($r3.PackageCount -gt 100) "Case 3: PackageCount > 100 (got $($r3.PackageCount))"
} else {
    Write-Host "  ⊘ Case 3 skipped (no /tmp/core-v9-upstream)" -ForegroundColor Yellow
}

if ($failed -gt 0) { Write-Host "`n$failed test(s) failed" -ForegroundColor Red; exit 1 }
Write-Host "`nAll tests passed" -ForegroundColor Green