# ─────────────────────────────────────────────────────────────────────────────
# ErrorParser.psm1 — Error accumulation + compile error parser
#
# Dependencies: ErrorExtractor.psm1 (Extract-BuildErrorLines, Extract-RuntimeFailureLines,
#               Extract-SetupFailedContext, Get-RawFallbackLines)
# ─────────────────────────────────────────────────────────────────────────────

function Add-BuildErrorsForPackage {
    <# .SYNOPSIS Accumulate build errors into a per-package hashtable. Falls back to raw output when extractors return empty. #>
    [CmdletBinding()]
    param([hashtable]$BuildErrorMap, [string]$PackageName, [string[]]$Lines)
    if (-not $BuildErrorMap -or -not $PackageName) { return }
    $buildLines = Resolve-BuildDiagnosticLines $Lines
    if (-not $buildLines -or $buildLines.Count -eq 0) { return }
    if (-not $BuildErrorMap.ContainsKey($PackageName)) { $BuildErrorMap[$PackageName] = [System.Collections.Generic.List[string]]::new() }
    foreach ($line in $buildLines) {
        if (-not $BuildErrorMap[$PackageName].Contains($line)) { $BuildErrorMap[$PackageName].Add($line) | Out-Null }
    }
}

function Add-RuntimeFailuresForPackage {
    <# .SYNOPSIS Accumulate runtime failures into a per-package hashtable. Falls back to raw output when extractors return empty. #>
    [CmdletBinding()]
    param([hashtable]$FailureMap, [string]$PackageName, [string[]]$Lines)
    if (-not $FailureMap -or -not $PackageName) { return }
    $runtimeLines = Resolve-RuntimeDiagnosticLines $Lines
    if (-not $runtimeLines -or $runtimeLines.Count -eq 0) { return }
    if (-not $FailureMap.ContainsKey($PackageName)) { $FailureMap[$PackageName] = [System.Collections.Generic.List[string]]::new() }
    foreach ($line in $runtimeLines) {
        if (-not $FailureMap[$PackageName].Contains($line)) { $FailureMap[$PackageName].Add($line) | Out-Null }
    }
}

function ParseCompileErrors {
    <# .SYNOPSIS Parse Go compile error lines into structured objects. #>
    [CmdletBinding()]
    [OutputType([hashtable[]])]
    param([string[]]$output)
    $errors = [System.Collections.Generic.List[object]]::new()
    foreach ($line in $output) {
        if ($line -match '^(.+?\.go):(\d+)(?::\d+)?:\s*(.+)$') {
            $file = Split-Path $Matches[1] -Leaf; $lineNum = [int]$Matches[2]; $msg = $Matches[3].Trim()
            $category = "other"
            if ($msg -match 'too many arguments|not enough arguments') { $category = "arg-count" }
            elseif ($msg -match 'undefined:') { $category = "undefined" }
            elseif ($msg -match 'cannot use .* as') { $category = "type-mismatch" }
            elseif ($msg -match 'has no field or method') { $category = "missing-member" }
            elseif ($msg -match 'cannot call non-function') { $category = "field-vs-method" }
            $errors.Add(@{ file = $file; line = $lineNum; message = $msg; category = $category; raw = $line })
        }
    }
    return $errors.ToArray()
}

Export-ModuleMember -Function @(
    'Add-BuildErrorsForPackage', 'Add-RuntimeFailuresForPackage', 'ParseCompileErrors'
)
