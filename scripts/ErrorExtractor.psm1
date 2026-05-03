# ─────────────────────────────────────────────────────────────────────────────
# ErrorExtractor.psm1 — Go build/runtime error line extraction
#
# Dependencies: None (standalone)
# ─────────────────────────────────────────────────────────────────────────────

function Filter-BlockedCompileLines {
    <# .SYNOPSIS Remove noisy/irrelevant lines from Go compile output. #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)
    $filtered = [System.Collections.Generic.List[string]]::new()
    foreach ($raw in $lines) {
        if ($null -eq $raw) { continue }; $line = $raw.ToString().TrimEnd("`r"); $trimmed = $line.Trim()
        if (-not $trimmed) { continue }
        if ($trimmed -match '^\s*warning:\s*no packages being tested depend on matches for pattern') { continue }
        if ($trimmed -match '^#\s+\S+' -and $trimmed -notmatch '\.go:\d+') { continue }
        if ($trimmed -match '^(github\.com|gitlab\.com)/\S+(\s+\[[^\]]+\])?$' -and $trimmed -notmatch '\.go:\d+') { continue }
        $filtered.Add($line) | Out-Null
    }
    return $filtered.ToArray()
}

function Extract-BuildErrorLines {
    <# .SYNOPSIS Extract compile-time error lines from Go build output. #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)
    $candidates = Filter-BlockedCompileLines $lines
    $errors = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
    foreach ($raw in $candidates) {
        if ($null -eq $raw) { continue }; $line = $raw.ToString().TrimEnd("`r"); $trimmed = $line.Trim()
        if (-not $trimmed) { continue }
        if ($trimmed -match '\.go:\d+(?::\d+)?:' -or $trimmed -match '^#\s+\S+') {
            if ($seen.Add($line)) { $errors.Add($line) | Out-Null }
        }
    }
    return $errors.ToArray()
}

function Extract-ExecutionFailureLines {
    <# .SYNOPSIS Extract execution failure lines (compile + runtime) from Go output. #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)
    $candidates = Filter-BlockedCompileLines $lines
    $errors = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
    foreach ($raw in $candidates) {
        if ($null -eq $raw) { continue }; $line = $raw.ToString().TrimEnd("`r"); $trimmed = $line.Trim()
        if (-not $trimmed) { continue }
        $isSetupOrBuildFail = $trimmed -match '\[setup failed\]' -or $trimmed -match '\[build failed\]' -or
            $trimmed -match '(?i)\bsetup failed\b' -or $trimmed -match '(?i)\bbuild failed\b'
        if ($trimmed -match '\.go:\d+(?::\d+)?:' -or $trimmed -match '^#\s+\S+' -or
            $trimmed -match '^(?i)panic:' -or $trimmed -match '^(?i)fatal error:' -or
            $trimmed -match '^--- FAIL:\s+' -or ($trimmed -match '^\s*FAIL\s+\S+' -and -not $isSetupOrBuildFail) -or
            $trimmed -match '^\s*exit status \d+\s*$') {
            if ($seen.Add($line)) { $errors.Add($line) | Out-Null }
        }
    }
    return $errors.ToArray()
}

function Extract-RuntimeFailureLines {
    <# .SYNOPSIS Extract ONLY runtime failure lines from Go output (no compile errors). #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)
    $candidates = Filter-BlockedCompileLines $lines
    $errors = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
    foreach ($raw in $candidates) {
        if ($null -eq $raw) { continue }; $line = $raw.ToString().TrimEnd("`r"); $trimmed = $line.Trim()
        if (-not $trimmed) { continue }
        $isSetupOrBuildFail = $trimmed -match '\[setup failed\]' -or $trimmed -match '\[build failed\]' -or
            $trimmed -match '(?i)\bsetup failed\b' -or $trimmed -match '(?i)\bbuild failed\b'
        if ($trimmed -match '^(?i)panic:' -or $trimmed -match '^(?i)fatal error:' -or
            $trimmed -match '^(?i)goroutine \d+' -or $trimmed -match '^--- FAIL:\s+' -or
            ($trimmed -match '^\s*FAIL\s+\S+' -and -not $isSetupOrBuildFail) -or
            $trimmed -match '^\s*exit status \d+\s*$' -or $trimmed -match '(?i)signal:\s+' -or
            $trimmed -match '(?i)runtime error:') {
            if ($seen.Add($line)) { $errors.Add($line) | Out-Null }
        }
    }
    return $errors.ToArray()
}

function Extract-SetupFailedContext {
    <#
    .SYNOPSIS
        Capture preceding context lines before a [setup failed] or [build failed] FAIL line.
    .DESCRIPTION
        Go outputs plain-text error messages before the final "FAIL pkg [setup failed]" line.
        Standard extractors miss these because they don't match .go:line: or panic: patterns.
        This function walks backward from each FAIL marker and captures up to N preceding
        non-empty lines as diagnostic context.
    .PARAMETER lines
        Raw Go test output lines.
    .PARAMETER ContextLineCount
        Max number of preceding lines to capture per FAIL marker (default 10).
    .EXAMPLE
        $context = Extract-SetupFailedContext $rawOutput
        # Returns context lines + the FAIL line for each [setup failed] occurrence
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param(
        [string[]]$lines,
        [int]$ContextLineCount = 10
    )

    if (-not $lines -or $lines.Count -eq 0) { return @() }

    $result = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)

    for ($i = 0; $i -lt $lines.Count; $i++) {
        $raw = $lines[$i]
        if ($null -eq $raw) { continue }
        $trimmed = $raw.ToString().TrimEnd("`r").Trim()
        if ($trimmed -match '\[setup failed\]' -or ($trimmed -match '\[build failed\]' -and $trimmed -match '^\s*FAIL\s+')) {
            # Walk backward to capture context
            $startIdx = [Math]::Max(0, $i - $ContextLineCount)
            for ($j = $startIdx; $j -le $i; $j++) {
                $ctxRaw = $lines[$j]
                if ($null -eq $ctxRaw) { continue }
                $ctxLine = $ctxRaw.ToString().TrimEnd("`r")
                if (-not $ctxLine.Trim()) { continue }
                # Skip noise lines
                if ($ctxLine.Trim() -match '^\s*warning:\s*no packages being tested') { continue }
                if ($seen.Add($ctxLine)) { $result.Add($ctxLine) | Out-Null }
            }
        }
    }

    return $result.ToArray()
}

function Get-RawFallbackLines {
    <#
    .SYNOPSIS
        Return all non-empty filtered lines as fallback when extractors find nothing actionable.
    .DESCRIPTION
        Used when Extract-BuildErrorLines and Extract-ExecutionFailureLines both return empty.
        Returns Filter-BlockedCompileLines output so plain-text error messages are preserved.
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)
    $candidates = Filter-BlockedCompileLines $lines
    $result = [System.Collections.Generic.List[string]]::new()
    foreach ($raw in $candidates) {
        if ($null -eq $raw) { continue }; $line = $raw.ToString().TrimEnd("`r")
        if ($line.Trim()) { $result.Add($line) | Out-Null }
    }
    return $result.ToArray()
}

function Resolve-BuildDiagnosticLines {
    <# .SYNOPSIS Return the most useful diagnostic lines for compile/setup failures. #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)

    $setupContext = Extract-SetupFailedContext $lines
    if ($setupContext -and $setupContext.Count -gt 0) { return $setupContext }

    $buildLines = Extract-BuildErrorLines $lines
    if ($buildLines -and $buildLines.Count -gt 0) { return $buildLines }

    $executionLines = Extract-ExecutionFailureLines $lines
    if ($executionLines -and $executionLines.Count -gt 0) { return $executionLines }

    return Get-RawFallbackLines $lines
}

function Resolve-RuntimeDiagnosticLines {
    <# .SYNOPSIS Return the most useful diagnostic lines for runtime/setup failures. #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string[]]$lines)

    $setupContext = Extract-SetupFailedContext $lines
    if ($setupContext -and $setupContext.Count -gt 0) { return $setupContext }

    $runtimeLines = Extract-RuntimeFailureLines $lines
    if ($runtimeLines -and $runtimeLines.Count -gt 0) { return $runtimeLines }

    return Get-RawFallbackLines $lines
}

function Test-IsGenericSetupFailureOutput {
    <# .SYNOPSIS Detect when diagnostics only contain package headers plus a setup/build FAIL marker. #>
    [CmdletBinding()]
    [OutputType([bool])]
    param([string[]]$lines)

    $diagnosticLines = Resolve-BuildDiagnosticLines $lines
    if (-not $diagnosticLines -or $diagnosticLines.Count -eq 0) { return $false }

    $hasSetupOrBuildFail = $false
    foreach ($raw in $diagnosticLines) {
        if ($null -eq $raw) { continue }
        $line = $raw.ToString().TrimEnd("`r")
        $trimmed = $line.Trim()
        if (-not $trimmed) { continue }
        if ($trimmed -match '^\s*FAIL\s+\S+\s+\[(setup failed|build failed)\]\s*$') {
            $hasSetupOrBuildFail = $true
            continue
        }
        if ($trimmed -match '^#\s+\S+(\s+\[[^\]]+\])?\s*$') { continue }
        return $false
    }

    return $hasSetupOrBuildFail
}

function Get-PackageLoaderDiagnosticLines {
    <#
    .SYNOPSIS
        Query go list for package-loader diagnostics when go test returns only generic setup/build failure markers.
    .PARAMETER PackagePath
        Import path or relative package path to inspect.
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string]$PackagePath)

    if (-not $PackagePath) { return @() }

    $template = @'
{{if or .Error (gt (len .DepsErrors) 0)}}package-load: {{.ImportPath}}{{if .ForTest}} [for {{.ForTest}}]{{end}}
{{with .Error}}{{if .Pos}}position: {{.Pos}}
{{end}}{{.Err}}
{{end}}{{range .DepsErrors}}package-load: {{$.ImportPath}}{{if $.ForTest}} [for {{$.ForTest}}]{{end}}
{{if .ImportStack}}import-stack: {{range $index, $pkg := .ImportStack}}{{if $index}} -> {{end}}{{$pkg}}{{end}}
{{end}}{{if .Pos}}position: {{.Pos}}
{{end}}{{.Err}}
{{end}}
{{end}}
'@

    $prevPref = $ErrorActionPreference
    $ErrorActionPreference = "Continue"
    $raw = & go list -e -deps -test -f $template "$PackagePath" 2>&1 | ForEach-Object { $_.ToString() }
    $ErrorActionPreference = $prevPref

    return Get-RawFallbackLines $raw
}

function Resolve-BlockedPackageDiagnosticOutput {
    <#
    .SYNOPSIS
        Merge package-loader diagnostics into blocked package output when go test only reports generic setup/build failure markers.
    #>
    [CmdletBinding()]
    [OutputType([string[]])]
    param([string]$PackagePath, [string[]]$lines)

    if (-not $lines -or $lines.Count -eq 0) { return @() }
    if (-not $PackagePath) { return $lines }
    if (-not (Test-IsGenericSetupFailureOutput $lines)) { return $lines }

    $loaderLines = Get-PackageLoaderDiagnosticLines $PackagePath
    if (-not $loaderLines -or $loaderLines.Count -eq 0) { return $lines }

    $orderedLines = [System.Collections.Generic.List[string]]::new()
    $insertedLoaderLines = $false
    foreach ($raw in $lines) {
        if ($null -eq $raw) { continue }

        $line = $raw.ToString().TrimEnd("`r")
        $trimmed = $line.Trim()
        if (-not $insertedLoaderLines -and $trimmed -match '^\s*FAIL\s+\S+\s+\[(setup failed|build failed)\]\s*$') {
            foreach ($loaderRaw in $loaderLines) {
                if ($null -eq $loaderRaw) { continue }
                $orderedLines.Add($loaderRaw.ToString().TrimEnd("`r")) | Out-Null
            }
            $insertedLoaderLines = $true
        }

        $orderedLines.Add($line) | Out-Null
    }

    if (-not $insertedLoaderLines) {
        foreach ($loaderRaw in $loaderLines) {
            if ($null -eq $loaderRaw) { continue }
            $orderedLines.Add($loaderRaw.ToString().TrimEnd("`r")) | Out-Null
        }
    }

    $merged = [System.Collections.Generic.List[string]]::new()
    $seen = [System.Collections.Generic.HashSet[string]]::new([System.StringComparer]::Ordinal)
    foreach ($raw in $orderedLines) {
        if ($null -eq $raw) { continue }
        $line = $raw.ToString().TrimEnd("`r")
        if (-not $line.Trim()) { continue }
        if ($seen.Add($line)) { $merged.Add($line) | Out-Null }
    }

    return $merged.ToArray()
}

Export-ModuleMember -Function @(
    'Filter-BlockedCompileLines', 'Extract-BuildErrorLines',
    'Extract-ExecutionFailureLines', 'Extract-RuntimeFailureLines',
    'Extract-SetupFailedContext', 'Get-RawFallbackLines',
    'Resolve-BuildDiagnosticLines', 'Resolve-RuntimeDiagnosticLines',
    'Resolve-BlockedPackageDiagnosticOutput'
)
