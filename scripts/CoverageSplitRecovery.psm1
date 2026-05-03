# ─────────────────────────────────────────────────────────────────────────────
# CoverageSplitRecovery.psm1 — Per-file split recovery for blocked packages
#
# Dependencies: Utilities.psm1, ErrorParser.psm1, ErrorExtractor.psm1, DashboardPhases.psm1
# ─────────────────────────────────────────────────────────────────────────────

function Invoke-CoverageSplitRecovery {
    <#
    .SYNOPSIS
        Split blocked packages into per-file subfolders and recheck each.
    .PARAMETER CompileResult
        Hashtable from Invoke-CoverageCompileCheck.
    .PARAMETER AllTestPkgs
        Array of all test package paths.
    .PARAMETER CovPkgList
        Comma-separated coverpkg list.
    .PARAMETER IsSyncMode
        Sequential or parallel mode.
    #>
    [CmdletBinding()]
    param(
        [hashtable]$CompileResult, [string[]]$AllTestPkgs,
        [string]$CovPkgList, [bool]$IsSyncMode
    )

    $blockedPkgs = $CompileResult.BlockedPkgs
    $blockedErrors = $CompileResult.BlockedErrors
    $testPkgs = $CompileResult.TestPkgs
    $buildErrorsByPackage = $CompileResult.BuildErrorsByPackage

    $splitRecoveredCount = 0
    $splitBlockedFiles = [System.Collections.Generic.List[string]]::new()
    $splitCleanupDirs = [System.Collections.Generic.List[string]]::new()

    if ($blockedPkgs.Count -gt 0) {
        $blockedSnapshot = @($blockedPkgs)
        foreach ($bp in $blockedSnapshot) {
            $pkgDir = Join-Path "tests" "integratedtests" $bp
            if (-not (Test-Path $pkgDir)) { continue }

            $bpTestFiles = Get-ChildItem -LiteralPath $pkgDir -Filter "*_test.go" -File |
                Where-Object { $_.Name -notlike "*helper*" } | Sort-Object Name
            $bpHelperTestFiles = Get-ChildItem -LiteralPath $pkgDir -Filter "*helper*_test.go" -File | Sort-Object Name
            $bpSupportFiles = Get-ChildItem -LiteralPath $pkgDir -Filter "*.go" -File |
                Where-Object { $_.Name -notlike "*_test.go" } | Sort-Object Name

            if ($bpTestFiles.Count -lt 2) { continue }

            Write-Host ""
            Write-Host "  Splitting $bp ($($bpTestFiles.Count) test files) for per-file recheck..." -ForegroundColor Yellow

            $subfolderResults = [System.Collections.Generic.List[pscustomobject]]::new()
            foreach ($tf in $bpTestFiles) {
                $folderName = $tf.BaseName -replace '_test$', ''
                $dest = Join-Path $pkgDir $folderName
                if (-not (Test-Path $dest)) { New-Item -ItemType Directory -Path $dest -Force | Out-Null }
                Copy-Item -LiteralPath $tf.FullName -Destination (Join-Path $dest $tf.Name) -Force
                foreach ($sf in $bpSupportFiles) { Copy-Item -LiteralPath $sf.FullName -Destination (Join-Path $dest $sf.Name) -Force }
                foreach ($hf in $bpHelperTestFiles) { Copy-Item -LiteralPath $hf.FullName -Destination (Join-Path $dest $hf.Name) -Force }
                $splitCleanupDirs.Add($dest)
            }

            $subDirs = Get-ChildItem -LiteralPath $pkgDir -Directory | Sort-Object Name

            if ($IsSyncMode) {
                foreach ($sd in $subDirs) {
                    $subPkg = "./tests/integratedtests/$bp/$($sd.Name)/"
                    $prevPref = $ErrorActionPreference
                    $ErrorActionPreference = "Continue"
                    $subOut = & go test -count=1 -run '^$' -gcflags=all=-e "-coverpkg=$CovPkgList" "$subPkg" 2>&1 | ForEach-Object { $_.ToString() }
                    $subExit = $LASTEXITCODE
                    $ErrorActionPreference = $prevPref
                    $subfolderResults.Add([pscustomobject]@{ Name = $sd.Name; Pkg = $subPkg; ExitCode = $subExit; Output = $subOut })
                }
            } else {
                $throttle = [Math]::Min($subDirs.Count, [Environment]::ProcessorCount * 2)
                $parallelResults = $subDirs | ForEach-Object -ThrottleLimit $throttle -Parallel {
                    $sd = $_
                    $bpName = $using:bp
                    $covPkgs = $using:CovPkgList
                    $subPkg = "./tests/integratedtests/$bpName/$($sd.Name)/"
                    $ErrorActionPreference = "Continue"
                    $subOut = & go test -count=1 -run '^$' -gcflags=all=-e "-coverpkg=$covPkgs" "$subPkg" 2>&1 | ForEach-Object { $_.ToString() }
                    [pscustomobject]@{ Name = $sd.Name; Pkg = $subPkg; ExitCode = $LASTEXITCODE; Output = $subOut }
                }
                foreach ($pr in ($parallelResults | Sort-Object Name)) { $subfolderResults.Add($pr) }
            }

            $subPass = @($subfolderResults | Where-Object { $_.ExitCode -eq 0 })
            $subFail = @($subfolderResults | Where-Object { $_.ExitCode -ne 0 })
            Write-Host "    ✓ $($subPass.Count) subfolders compile OK" -ForegroundColor Green
            if ($subFail.Count -gt 0) {
                $s = Get-CallerSource
                Write-Host "    ✗ $($subFail.Count) subfolders failed (source: $s):" -ForegroundColor Red
                foreach ($sf in $subFail) {
                    Write-Host "      ✗ $($sf.Name)" -ForegroundColor Red
                    $splitBlockedFiles.Add("$bp/$($sf.Name)")
                }
            }

            foreach ($sp in $subPass) {
                $fullSubPkg = $AllTestPkgs | Where-Object { $_ -match "integratedtests/$bp$" } | Select-Object -First 1
                $subFullPkg = if ($fullSubPkg) { $fullSubPkg + "/" + $sp.Name } else { $sp.Pkg }
                $testPkgs.Add($subFullPkg)
                $splitRecoveredCount++
            }

            $blockedPkgs.Remove($bp)
            $blockedErrors.Remove($bp)
            foreach ($sf in $subFail) {
                $failName = "$bp/$($sf.Name)"
                $resolvedOutput = Resolve-BlockedPackageDiagnosticOutput -PackagePath $sf.Pkg -Lines $sf.Output
                $blockedPkgs.Add($failName)
                $blockedErrors[$failName] = ($resolvedOutput -join "`n")
                Add-BuildErrorsForPackage $buildErrorsByPackage $failName $resolvedOutput
            }
        }

        if ($splitRecoveredCount -gt 0) {
            Write-Host ""
            Write-Success "Recovered $splitRecoveredCount subfolders from blocked packages via per-file split"
        }
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Split Recovery" "pass" "$splitRecoveredCount subfolders recovered" }
    } else {
        if (Get-Command Register-Phase -ErrorAction SilentlyContinue) { Register-Phase "Split Recovery" "skip" "not needed" }
    }

    return @{ SplitRecoveredCount = $splitRecoveredCount; SplitBlockedFiles = $splitBlockedFiles; SplitCleanupDirs = $splitCleanupDirs }
}

Export-ModuleMember -Function @('Invoke-CoverageSplitRecovery')
