# ─────────────────────────────────────────────────────────────────────────────
# CoverageRunner.psm1 — Thin orchestrator for full coverage pipeline (TC)
#
# Usage:
#   Import-Module ./scripts/CoverageRunner.psm1 -Force
#
# Dependencies:
#   CoveragePreChecks.psm1, CoverageCompileCheck.psm1,
#   CoverageProfileMerger.psm1, CoverageReport.psm1, PackageCoverage.psm1,
#   Utilities.psm1, TestLogWriter.psm1, TestRunner.psm1, DashboardUI.psm1
# ─────────────────────────────────────────────────────────────────────────────

function Invoke-TestCoverage {
    <#
    .SYNOPSIS
        Run full test coverage across all packages (TC command).
    .DESCRIPTION
        Orchestrates: pre-checks → compile check → split recovery →
        coverage run → profile merge → report generation → dashboard summary.
    #>
    Write-Header "Running tests with coverage"

    if (Get-Command Reset-Phases -ErrorAction SilentlyContinue) { Reset-Phases }

    Invoke-FetchLatest
    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
        Register-Phase "Git Pull" "pass" "pulled from remote"
        Register-Phase "Dependencies" "pass" "up to date"
    }

    # Clean data folder
    $dataDir = $global:DataDir
    if (Test-Path $dataDir) { Remove-Item -Recurse -Force $dataDir; Write-Host "  Cleaned data/ folder" -ForegroundColor Yellow }
    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Data Cleanup" "pass" "cleaned" }

    $coverDir = Join-Path $global:DataDir "coverage"
    $partialDir = Join-Path $coverDir "partial"
    New-Item -ItemType Directory -Path $partialDir -Force | Out-Null

    $coverProfile = Join-Path $coverDir "coverage.out"
    $coverHtml    = Join-Path $coverDir "coverage.html"
    $coverSummary = Join-Path $coverDir "coverage-summary.txt"

    # Repo build errors script
    $repoBuildErrorsFile = Join-Path $coverDir "repo-build-errors.txt"
    $repoBuildErrorsJsonFile = Join-Path $coverDir "repo-build-errors.json"
    $repoBuildErrorsScript = Join-Path $global:ProjectRoot "scripts" "coverage" "Export-RepoBuildErrors.ps1"
    if (Test-Path $repoBuildErrorsScript) {
        & $repoBuildErrorsScript -OutputTxt $repoBuildErrorsFile -OutputJson $repoBuildErrorsJsonFile
    }

    # Build package lists
    $allPkgs = go list ./... 2>&1 | ForEach-Object { $_.ToString() }
    $srcPkgs = $allPkgs | Where-Object { $_ -notmatch '/tests/' }
    $covPkgList = $srcPkgs -join ","

    $integrationTestPkgs = go list ./tests/integratedtests/... 2>&1 |
        ForEach-Object { $_.ToString() } | Where-Object { $_ -and $_ -notmatch '^warning:' }

    $inPkgTestPkgs = @()
    foreach ($srcPkg in $srcPkgs) {
        $relPath = $srcPkg -replace '^github\.com/alimtvnetwork/core-v8/', ''
        if ($relPath -and (Test-Path $relPath) -and (Get-ChildItem -Path $relPath -Filter '*_test.go' -File -ErrorAction SilentlyContinue)) {
            $inPkgTestPkgs += $srcPkg
        }
    }
    $allTestPkgs = @($integrationTestPkgs) + @($inPkgTestPkgs) | Sort-Object -Unique

    # ── Pre-checks ──
    $preCheckOk = Invoke-CoveragePreChecks -ScriptRoot $global:ProjectRoot -ExtraArgs $ExtraArgs -CoverDir $coverDir
    if (-not $preCheckOk) { exit 1 }

    # ── Compile check ──
    $isSyncMode = $ExtraArgs -and ($ExtraArgs -contains "--sync")
    $compileResult = Invoke-CoverageCompileCheck -AllTestPkgs $allTestPkgs -CovPkgList $covPkgList -IsSyncMode $isSyncMode

    # ── Split recovery ──
    $splitResult = Invoke-CoverageSplitRecovery -CompileResult $compileResult -AllTestPkgs $allTestPkgs -CovPkgList $covPkgList -IsSyncMode $isSyncMode

    $testPkgs = $compileResult.TestPkgs
    $blockedPkgs = $compileResult.BlockedPkgs
    $blockedErrors = $compileResult.BlockedErrors
    $buildErrorsByPackage = $compileResult.BuildErrorsByPackage
    $runtimeFailuresByPackage = $compileResult.RuntimeFailuresByPackage

    # Print blocked summary
    if ($blockedPkgs.Count -gt 0) {
        Write-Host ""
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Red
        Write-Host "  │ BLOCKED PACKAGES ($($blockedPkgs.Count) failed to compile)" -ForegroundColor Red
        Write-Host "  │" -ForegroundColor Red
        foreach ($bp in ($blockedPkgs | Sort-Object)) { Write-Host "  │   ✗ $bp" -ForegroundColor Red }
        Write-Host "  │" -ForegroundColor Red
        Write-Host "  │ These packages will be SKIPPED in coverage." -ForegroundColor Yellow
        Write-Host "  │ Fix their build errors to include them." -ForegroundColor Yellow
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Red
        Write-Host ""

        # Write blocked details file
        $blockedFile = Join-Path $coverDir "blocked-packages.txt"
        $sortedBlocked = $blockedPkgs | Sort-Object
        $callerSource = Get-CallerSource
        $blockedContent = @("# Blocked Packages — $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')", "# Count: $($blockedPkgs.Count)", "# Source: $callerSource (CoverageRunner.psm1 → Invoke-TestCoverage)", "")
        foreach ($bp in $sortedBlocked) {
            $blockedContent += "## $bp"
            if ($blockedErrors.ContainsKey($bp)) {
                $rawErrLines = $blockedErrors[$bp] -split "`n"
                $filtered = Resolve-BuildDiagnosticLines $rawErrLines
                if ($filtered.Count -gt 0) { $blockedContent += ($filtered -join "`n") } else { $blockedContent += "(no actionable compile errors captured)" }
            }
            $blockedContent += ""
        }
        Set-Content -Path $blockedFile -Value ($blockedContent -join "`n") -Encoding UTF8

        # Blocked JSON
        $blockedJsonFile = Join-Path $coverDir "blocked-packages.json"
        $blockedJsonItems = [System.Collections.Generic.List[object]]::new()
        foreach ($bp in $sortedBlocked) {
            $errText = ""; if ($blockedErrors.ContainsKey($bp)) { $errText = $blockedErrors[$bp] }
            $errLines = @(); if ($errText) { $errLines = Resolve-BuildDiagnosticLines ($errText -split "`n") }
            $blockedJsonItems.Add(@{ package = $bp; errorCount = $errLines.Count; errors = $errLines })
        }
        @{ timestamp = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"); blockedCount = $blockedPkgs.Count; compiledCount = $testPkgs.Count; totalCount = $allTestPkgs.Count; blockedPackages = $blockedJsonItems.ToArray(); missingProfiles = @() } |
            ConvertTo-Json -Depth 4 | Set-Content -Path $blockedJsonFile -Encoding UTF8
    } else {
        Write-Host ""; Write-Success "All $($testPkgs.Count) packages compiled successfully"
    }
    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
        if ($blockedPkgs.Count -gt 0) { Register-Phase "Compile Check" "warn" "$($testPkgs.Count)/$($allTestPkgs.Count) passed, $($blockedPkgs.Count) blocked" }
        else { Register-Phase "Compile Check" "pass" "$($testPkgs.Count)/$($allTestPkgs.Count) passed" }
    }

    if ($testPkgs.Count -eq 0) {
        $s = Get-CallerSource; Write-Fail "No packages compiled — aborting coverage run (source: $s)"
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Coverage Run" "fail" "no packages to run" }
        return
    }

    # ── Coverage run ──
    $allOutput = [System.Collections.Generic.List[string]]::new()
    $overallExit = 0
    $modeLabel = if ($isSyncMode) { "sync" } else { "parallel" }
    Write-Host ""; Write-Host "  Running $($testPkgs.Count) test packages ($modeLabel)..." -ForegroundColor Yellow

    if ($isSyncMode) {
        $pkgIndex = 0
        foreach ($testPkg in $testPkgs) {
            $pkgIndex++
            $shortName = $testPkg -replace '.*integratedtests/?', ''; if (-not $shortName) { $shortName = "(root)" }
            $partialProfile = Join-Path $partialDir "cover-$pkgIndex.out"
            $prevPref = $ErrorActionPreference; $ErrorActionPreference = "Continue"
            $output = & go test -count=1 "-coverprofile=$partialProfile" "-coverpkg=$covPkgList" "$testPkg" 2>&1 | ForEach-Object { $_.ToString() }
            $pkgExit = $LASTEXITCODE; $ErrorActionPreference = $prevPref
            if ($pkgExit -ne 0) {
                $overallExit = $pkgExit
                Add-BuildErrorsForPackage $buildErrorsByPackage $shortName $output
                Add-RuntimeFailuresForPackage $runtimeFailuresByPackage $shortName $output
            }
            if ($output) { foreach ($line in $output) { $allOutput.Add([string]$line) } }
        }
    } else {
        $throttle = [Math]::Min($testPkgs.Count, [Environment]::ProcessorCount * 2)
        $coverResults = $testPkgs | ForEach-Object -ThrottleLimit $throttle -Parallel {
            $pkg = $_; $covPkgs = $using:covPkgList; $pDir = $using:partialDir
            $safePkgName = $pkg -replace '[^a-zA-Z0-9\.-]', '_'
            $profile = Join-Path $pDir "cover-$safePkgName.out"
            $ErrorActionPreference = "Continue"
            $out = & go test -count=1 "-coverprofile=$profile" "-coverpkg=$covPkgs" "$pkg" 2>&1 | ForEach-Object { $_.ToString() }
            [pscustomobject]@{ Pkg = $pkg; Profile = $profile; ExitCode = $LASTEXITCODE; Output = $out }
        }
        foreach ($result in ($coverResults | Sort-Object Pkg)) {
            $shortName = $result.Pkg -replace '.*integratedtests/?', ''; if (-not $shortName) { $shortName = "(root)" }
            if ($result.ExitCode -ne 0) {
                $overallExit = $result.ExitCode
                Add-BuildErrorsForPackage $buildErrorsByPackage $shortName $result.Output
                Add-RuntimeFailuresForPackage $runtimeFailuresByPackage $shortName $result.Output
            }
            if ($result.Output) { foreach ($line in $result.Output) { $allOutput.Add([string]$line) } }
        }
    }

    # ── Missing profiles ──
    $missingProfiles = Find-MissingCoverageProfiles -TestPkgs $testPkgs -PartialDir $partialDir -IsSyncMode $isSyncMode

    # Backfill missing-profiles into blocked JSON
    $blockedJsonFile = Join-Path $coverDir "blocked-packages.json"
    if (Test-Path $blockedJsonFile) {
        $existingJson = Get-Content $blockedJsonFile -Raw | ConvertFrom-Json
        $existingJson | Add-Member -NotePropertyName "missingProfileCount" -NotePropertyValue $missingProfiles.Count -Force
        $existingJson | Add-Member -NotePropertyName "missingProfiles" -NotePropertyValue @($missingProfiles | ForEach-Object { $_ }) -Force
        $existingJson | ConvertTo-Json -Depth 4 | Set-Content -Path $blockedJsonFile -Encoding UTF8
    } elseif ($missingProfiles.Count -gt 0) {
        @{ timestamp = (Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"); blockedCount = 0; compiledCount = $testPkgs.Count; totalCount = $allTestPkgs.Count; blockedPackages = @(); missingProfileCount = $missingProfiles.Count; missingProfiles = @($missingProfiles | ForEach-Object { $_ }) } |
            ConvertTo-Json -Depth 4 | Set-Content -Path $blockedJsonFile -Encoding UTF8
    }

    # Write test logs
    Write-TestLogs $allOutput.ToArray()

    # Failing test summary
    $failedTestNames = [System.Collections.Generic.HashSet[string]]::new()
    foreach ($line in $allOutput) {
        if ($line -match "--- FAIL:\s+(.+?)\s+\(") { $failedTestNames.Add($Matches[1].Trim()) | Out-Null }
    }
    if ($failedTestNames.Count -gt 0) {
        Write-Host ""
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Red
        Write-Host "  │ FAILING TESTS ($($failedTestNames.Count) failed)" -ForegroundColor Red
        Write-Host "  │" -ForegroundColor Red
        foreach ($ft in ($failedTestNames | Sort-Object)) { Write-Host "  │   ✗ $ft" -ForegroundColor Red }
        Write-Host "  │" -ForegroundColor Red
        Write-Host "  │ See data/test-logs/failing-tests.txt for details." -ForegroundColor Yellow
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Red
    }

    # ── Merge profiles ──
    $mergedLines = Merge-CoverageProfiles -PartialDir $partialDir -CoverProfile $coverProfile

    if (Test-Path $coverProfile) {
        $funcOutput = & go tool cover "-func=$coverProfile" 2>&1 | ForEach-Object { $_.ToString() }

        $uncoveredJsonScript = Join-Path $global:ProjectRoot "scripts" "coverage" "Export-UncoveredMethodsJson.ps1"
        if (Test-Path $uncoveredJsonScript) {
            $uncoveredJsonFile = Join-Path $coverDir "uncovered-method-lines.json"
            & $uncoveredJsonScript -CoverProfile $coverProfile -FuncOutput $funcOutput -OutputFile $uncoveredJsonFile -ProjectRoot $global:ProjectRoot
        }

        $srcPkgStmts = Build-SourcePackageCoverage $mergedLines
        $totalLine = $funcOutput | Where-Object { $_ -match "^total:" } | Select-Object -Last 1

        # Generate all reports
        Write-CoverageHtmlWithAiButton -CoverHtml $coverHtml -CoverProfile $coverProfile -FuncOutput $funcOutput -SrcPkgStmts $srcPkgStmts
        Write-CoverageSummaryReport -CoverSummary $coverSummary -CoverProfile $coverProfile -CoverHtml $coverHtml -FuncOutput $funcOutput -SrcPkgStmts $srcPkgStmts
        $coverJsonFile = Join-Path $coverDir "coverage-summary.json"
        $totalPct = Write-CoverageJsonReport -CoverJsonFile $coverJsonFile -CoverProfile $coverProfile -CoverHtml $coverHtml -CoverSummary $coverSummary -FuncOutput $funcOutput -SrcPkgStmts $srcPkgStmts -CoverDir $coverDir
        Write-PerPackageCoverageReport -CoverDir $coverDir -SrcPkgStmts $srcPkgStmts -TotalPct $totalPct -TotalLine $totalLine
        Write-CoverageConsoleSummary -SrcPkgStmts $srcPkgStmts -FuncOutput $funcOutput
        Write-BuildErrorsReport -CoverDir $coverDir -BuildErrorsByPackage $buildErrorsByPackage -BlockedPkgs $blockedPkgs -BlockedErrors $blockedErrors
        Write-RuntimeFailuresReport -CoverDir $coverDir -RuntimeFailuresByPackage $runtimeFailuresByPackage -MissingProfiles $missingProfiles -BlockedPkgs $blockedPkgs

        # Written files summary
        Write-Host ""
        Write-Host "  ┌─────────────────────────────────────────────────" -ForegroundColor Gray
        Write-Host "  │ WRITTEN FILES" -ForegroundColor Gray
        foreach ($f in @($coverProfile, $coverHtml, $coverSummary, $coverJsonFile,
            (Join-Path $coverDir "per-package-coverage.txt"), (Join-Path $coverDir "per-package-coverage.json"),
            (Join-Path $coverDir "build-errors.txt"), (Join-Path $coverDir "build-errors.json"),
            (Join-Path $coverDir "runtime-failures.txt"), (Join-Path $coverDir "runtime-failures.json"))) {
            Write-Host "  │  $f" -ForegroundColor Gray
        }
        if (Test-Path $repoBuildErrorsFile) { Write-Host "  │  $repoBuildErrorsFile" -ForegroundColor Gray }
        if (Test-Path $repoBuildErrorsJsonFile) { Write-Host "  │  $repoBuildErrorsJsonFile" -ForegroundColor Gray }
        if ($blockedPkgs.Count -gt 0) {
            Write-Host "  │  $(Join-Path $coverDir 'blocked-packages.txt')" -ForegroundColor Gray
            Write-Host "  │  $(Join-Path $coverDir 'blocked-packages.json')" -ForegroundColor Gray
        }
        $syntaxIssuesFile = Join-Path $coverDir "syntax-issues.txt"
        if (Test-Path $syntaxIssuesFile) { Write-Host "  │  $syntaxIssuesFile" -ForegroundColor Gray }
        Write-Host "  └─────────────────────────────────────────────────" -ForegroundColor Gray

        # AI prompts
        $promptScript = Join-Path $global:ProjectRoot "scripts" "coverage" "Generate-CoveragePrompts.ps1"
        if (Test-Path $promptScript) {
            Write-Host ""; Write-Header "Generating coverage improvement prompts"
            & $promptScript -CoverProfile $coverProfile -FuncOutput $funcOutput -OutputDir (Join-Path $global:DataDir "prompts") -BatchSize 500 -ProjectRoot $global:ProjectRoot
        }

        # HTML auto-open
        if ($ExtraArgs -and $ExtraArgs[0] -eq "--open" -and (Test-Path $coverHtml)) {
            Write-Host ""; Write-Host "  Opening HTML coverage report in browser..." -ForegroundColor Yellow
            Start-Process $coverHtml
        }
    }
    Open-FailingTestsIfAny

    if (Get-Command Register-Phase -ErrorAction SilentlyContinue) {
        Register-Phase "Coverage Run" "pass" "$($testPkgs.Count) packages"
        Register-Phase "Coverage Report" "pass" "generated"
    }

    # Cleanup split subfolders
    $splitCleanupDirs = $splitResult.SplitCleanupDirs
    if ($splitCleanupDirs.Count -gt 0) {
        Write-Host ""; Write-Host "  Cleaning up $($splitCleanupDirs.Count) split subfolders..." -ForegroundColor Gray
        foreach ($cleanDir in $splitCleanupDirs) { if (Test-Path $cleanDir) { Remove-Item -LiteralPath $cleanDir -Recurse -Force } }
        Write-Host "  ✓ Split subfolders removed" -ForegroundColor Green
    }

    if (Get-Command Write-PhaseSummaryBox -ErrorAction SilentlyContinue) { Write-Host ""; Write-PhaseSummaryBox }
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Invoke-TestCoverage'
)
