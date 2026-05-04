# ─────────────────────────────────────────────────────────────────────────────
# CoverageCompileCheck.psm1 — Pre-coverage compile checks (sync + parallel)
#
# Dependencies: Utilities.psm1, ErrorParser.psm1, ErrorExtractor.psm1
# ─────────────────────────────────────────────────────────────────────────────

function Write-BlockedDiagnostic {
    <#
    .SYNOPSIS
        Echo the (filtered) go compile diagnostic for a blocked package inline,
        capped at a sensible length so the console stays readable.
    #>
    param([string[]]$Lines, [int]$MaxLines = 25)
    if (-not $Lines -or $Lines.Count -eq 0) {
        Write-Host "      (no diagnostic captured - see data/coverage/blocked-packages.txt)" -ForegroundColor DarkGray
        return
    }
    $shown = 0
    foreach ($line in $Lines) {
        if (-not $line) { continue }
        if ($shown -ge $MaxLines) {
            $remaining = $Lines.Count - $shown
            Write-Host "      ... +$remaining more line(s) - full diagnostic in data/coverage/blocked-packages.txt" -ForegroundColor DarkGray
            break
        }
        Write-Host "      $line" -ForegroundColor DarkYellow
        $shown++
    }
}

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

    # Helper: produce a short, ALWAYS-non-empty label for a package path.
    # Falls back through: trailing creationtests/integratedtests segment →
    # last path segment → full path → literal "(unknown)". Never returns "(root)"
    # for non-empty input — that label was masking the real failing import path.
    function Get-PackageShortName {
        param([string]$pkg)
        if (-not $pkg) { return "(unknown)" }
        $name = $pkg -replace '.*(integratedtests|creationtests)/', ''
        if ($name -and $name -ne $pkg) { return $name.TrimEnd('/') }
        $segments = $pkg.TrimEnd('/').Split('/')
        if ($segments.Count -ge 1 -and $segments[-1]) { return $segments[-1] }
        return $pkg
    }

    if ($IsSyncMode) {
        foreach ($testPkg in $AllTestPkgs) {
            $shortName = Get-PackageShortName $testPkg

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
                Write-Fail "Blocked: $shortName  ($testPkg) (source: $callerSource)"
                Write-BlockedDiagnostic $combinedOut
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
            $shortName = Get-PackageShortName $result.Pkg

            if ($result.ExitCode -eq 0) {
                $testPkgs.Add($result.Pkg)
            } else {
                $diagnosticOut = Resolve-BlockedPackageDiagnosticOutput -PackagePath $result.Pkg -Lines $result.Output
                $callerSource = "CoverageCompileCheck.psm1 → Invoke-CoverageCompileCheck (parallel)"
                Write-Fail "Blocked: $shortName  ($($result.Pkg)) (source: $callerSource)"
                Write-BlockedDiagnostic $diagnosticOut
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
