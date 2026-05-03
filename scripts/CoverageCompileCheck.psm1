# ─────────────────────────────────────────────────────────────────────────────
# CoverageCompileCheck.psm1 — Pre-coverage compile checks (sync + parallel)
#
# Dependencies: Utilities.psm1, ErrorParser.psm1, ErrorExtractor.psm1
# ─────────────────────────────────────────────────────────────────────────────

function Invoke-CoverageCompileCheck {
    <#
    .SYNOPSIS
        Build-check each test package individually before coverage run.
    .RETURNS
        Hashtable with TestPkgs, BlockedPkgs, BlockedErrors, BuildErrorsByPackage, RuntimeFailuresByPackage.
    #>
    [CmdletBinding()]
    param([string[]]$AllTestPkgs, [string]$CovPkgList, [bool]$IsSyncMode)

    $blockedPkgs = [System.Collections.Generic.List[string]]::new()
    $blockedErrors = [System.Collections.Generic.Dictionary[string, string]]::new()
    $testPkgs = [System.Collections.Generic.List[string]]::new()
    $buildErrorsByPackage = @{}
    $runtimeFailuresByPackage = @{}

    $modeLabel = if ($IsSyncMode) { "sync" } else { "parallel" }
    Write-Host ""; Write-Header "Pre-coverage compile check ($($AllTestPkgs.Count) packages, $modeLabel mode)"

    if ($IsSyncMode) {
        foreach ($testPkg in $AllTestPkgs) {
            $shortName = $testPkg -replace '.*integratedtests/?', ''
            if (-not $shortName) { $shortName = "(root)" }

            $prevPref = $ErrorActionPreference
            $ErrorActionPreference = "Continue"
            $compileOut = & go test -count=1 -run '^$' -gcflags=all=-e "-coverpkg=$CovPkgList" "$testPkg" 2>&1 | ForEach-Object { $_.ToString() }
            $compileExit = $LASTEXITCODE
            $ErrorActionPreference = $prevPref

            if ($compileExit -eq 0) {
                $testPkgs.Add($testPkg)
            } else {
                $prevPref = $ErrorActionPreference
                $ErrorActionPreference = "Continue"
                $diagOut = & go test -count=1 -run '^$' -gcflags=all=-e "$testPkg" 2>&1 | ForEach-Object { $_.ToString() }
                $ErrorActionPreference = $prevPref

                $combinedOut = Merge-UniqueOutputLines $compileOut $diagOut
                $combinedOut = Resolve-BlockedPackageDiagnosticOutput -PackagePath $testPkg -Lines $combinedOut
                $callerSource = Get-CallerSource
                Write-Fail "Blocked: $shortName (source: $callerSource)"
                $blockedPkgs.Add($shortName)
                $blockedErrors[$shortName] = ($combinedOut -join "`n")
                Add-BuildErrorsForPackage $buildErrorsByPackage $shortName $combinedOut
            }
        }
    } else {
        $throttle = [Math]::Min($AllTestPkgs.Count, [Environment]::ProcessorCount * 2)
        Write-Host "  Launching $($AllTestPkgs.Count) compile checks ($throttle parallel)..." -ForegroundColor Gray

        $compileResults = $AllTestPkgs | ForEach-Object -ThrottleLimit $throttle -Parallel {
            $pkg = $_
            $covPkgs = $using:CovPkgList
            $ErrorActionPreference = "Continue"
            $rawOut = & go test -count=1 -run '^$' -gcflags=all=-e "-coverpkg=$covPkgs" "$pkg" 2>&1
            $ec = $LASTEXITCODE
            $out = @($rawOut | ForEach-Object { $_.ToString() })
            if ($ec -ne 0) {
                $diagRaw = & go test -count=1 -run '^$' -gcflags=all=-e "$pkg" 2>&1
                $diagOut = @($diagRaw | ForEach-Object { $_.ToString() })
                $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
                $merged = [System.Collections.Generic.List[string]]::new()
                foreach ($line in @($out + $diagOut)) {
                    if ($null -eq $line) { continue }
                    $normalized = $line.ToString().TrimEnd("`r")
                    if (-not $normalized) { continue }
                    if ($seen.Add($normalized)) { $merged.Add($normalized) | Out-Null }
                }
                $out = $merged.ToArray()
            }
            [pscustomobject]@{ Pkg = $pkg; ExitCode = $ec; Output = $out }
        }

        foreach ($result in ($compileResults | Sort-Object Pkg)) {
            $shortName = $result.Pkg -replace '.*integratedtests/?', ''
            if (-not $shortName) { $shortName = "(root)" }

            if ($result.ExitCode -eq 0) {
                $testPkgs.Add($result.Pkg)
            } else {
                $diagnosticOut = Resolve-BlockedPackageDiagnosticOutput -PackagePath $result.Pkg -Lines $result.Output
                $callerSource = "CoverageCompileCheck.psm1 → Invoke-CoverageCompileCheck (parallel)"
                Write-Fail "Blocked: $shortName (source: $callerSource)"
                $blockedPkgs.Add($shortName)
                $blockedErrors[$shortName] = ($diagnosticOut -join "`n")
                Add-BuildErrorsForPackage $buildErrorsByPackage $shortName $diagnosticOut
            }
        }
    }

    return @{
        TestPkgs = $testPkgs; BlockedPkgs = $blockedPkgs; BlockedErrors = $blockedErrors
        BuildErrorsByPackage = $buildErrorsByPackage; RuntimeFailuresByPackage = $runtimeFailuresByPackage
    }
}

Export-ModuleMember -Function @('Invoke-CoverageCompileCheck')
