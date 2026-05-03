# ─────────────────────────────────────────────────────────────────────────────
# TestLogWriter.psm1 — Go test output log parser and file writer
#
# Parses raw Go test output, classifies tests as passing/failing,
# captures diagnostic details, and writes structured log files.
#
# Usage:
#   Import-Module ./scripts/TestLogWriter.psm1 -Force
#   Write-TestLogs -rawOutput $goTestOutput
#
# Dependencies: Utilities.psm1 (Filter-TestWarnings, Write-Success, Write-Fail,
#               Ensure-TestLogDir)
# ─────────────────────────────────────────────────────────────────────────────

function Write-TestLogs {
    <#
    .SYNOPSIS
        Parse Go test output and write structured passing/failing log files.
    .DESCRIPTION
        Two-pass parser:
          Pass 1 — Identify which test names passed (--- PASS) vs failed (--- FAIL).
          Pass 2 — Collect diagnostic lines (assertions, panics, expected/actual)
                   for each failed test, skipping noisy coverage/summary lines.

        Writes three files to $TestLogDir (data/test-logs/):
          - passing-tests.txt  — timestamped list of passing test names
          - failing-tests.txt  — summary + per-test diagnostic blocks
          - raw-output.txt     — filtered raw output for debugging

        Also handles pure compilation failures (no === RUN lines) by
        extracting .go:line: errors and [build failed] markers.
    .PARAMETER rawOutput
        Array of raw lines from `go test` stdout+stderr.
    .EXAMPLE
        $output = & go test ./... 2>&1 | ForEach-Object { $_.ToString() }
        Write-TestLogs -rawOutput $output
    #>
    [CmdletBinding()]
    param([string[]]$rawOutput)

    Ensure-TestLogDir

    $passingFile = Join-Path $global:TestLogDir "passing-tests.txt"
    $failingFile = Join-Path $global:TestLogDir "failing-tests.txt"
    $rawFile     = Join-Path $global:TestLogDir "raw-output.txt"

    # Clear previous log files before writing new results
    @($passingFile, $failingFile, $rawFile) | ForEach-Object {
        if (Test-Path $_) { Remove-Item $_ -Force }
    }

    $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    $passing = [System.Collections.Generic.List[string]]::new()
    $failing = [System.Collections.Generic.List[string]]::new()

    # Remove noisy go-test coverpkg warnings from logs
    $filteredOutput = Filter-TestWarnings $rawOutput

    # Save filtered output for debugging
    Set-Content -Path $rawFile -Value ($filteredOutput -join "`n") -Encoding UTF8

    # ── Pass 1: Identify which tests passed and which failed ──
    $failedNames = [System.Collections.Generic.HashSet[string]]::new()
    $passedNames = [System.Collections.Generic.HashSet[string]]::new()

    foreach ($line in $filteredOutput) {
        if ($line -match "--- FAIL:\s+(.+?)\s+\(") {
            $failedNames.Add($Matches[1].Trim()) | Out-Null
        }
        elseif ($line -match "--- PASS:\s+(.+?)\s+\(") {
            $passedNames.Add($Matches[1].Trim()) | Out-Null
        }
    }

    # ── Pass 2: Collect diagnostic details for failed tests ──
    $currentTest = ""
    $currentBlock = [System.Collections.Generic.List[string]]::new()

    foreach ($line in $filteredOutput) {
        if ($line -match "=== RUN\s+(.+)$") {
            # Flush previous block if it was a failed test
            if ($currentTest -and $failedNames.Contains($currentTest)) {
                $failing.Add("FAIL: $currentTest")
                foreach ($detail in $currentBlock) {
                    $failing.Add("  $detail")
                }
                $failing.Add("")
            }

            $currentTest = $Matches[1].Trim()
            $currentBlock.Clear()
        }
        elseif ($line -match "--- PASS:\s+(.+?)\s+\(") {
            # Passing test — flush and reset
            $currentTest = ""
            $currentBlock.Clear()
        }
        elseif ($line -match "--- FAIL:\s+(.+?)\s+\(") {
            # Capture the --- FAIL line itself as part of diagnostics
            if (-not $currentTest) {
                $currentTest = $Matches[1].Trim()
            }
            if ($currentTest) {
                $currentBlock.Add($line)
            }
        }
        else {
            if ($currentTest) {
                # Skip noisy non-diagnostic lines (coverage/package summaries/progress markers)
                $lineForMatch = $line.TrimEnd("`r")
                if ($lineForMatch -match '^\s*coverage:\s+\d' -or
                    $lineForMatch -match '^\s*(ok|FAIL)\s+\S+\s+\d+(\.\d+)?s(\s+coverage:.*)?\s*$' -or
                    $lineForMatch -match '^\s*(ok|FAIL|PASS)\s*$' -or
                    $lineForMatch -match '^\s*\?\s+\S+\s+\[no test files\]\s*$' -or
                    $lineForMatch -match '^\s*===\s+(RUN|PAUSE|CONT)\s+' -or
                    $lineForMatch -match '^\s*\.+\s*(FAIL|ok)\s*$') {
                    continue
                }
                $currentBlock.Add($line)
            }
        }
    }

    # Flush last block
    if ($currentTest -and $failedNames.Contains($currentTest)) {
        $failing.Add("FAIL: $currentTest")
        foreach ($detail in $currentBlock) {
            $failing.Add("  $detail")
        }
        $failing.Add("")
    }

    # Collect passing test names
    foreach ($name in $passedNames) {
        $passing.Add($name)
    }

    # ── Write passing tests ──
    $callerSource = Get-CallerSource
    $passingContent = @("# Passing Tests — $timestamp", "# Count: $($passing.Count)", "# Source: $callerSource (TestLogWriter.psm1 → Write-TestLogs)", "")
    $passingContent += $passing
    Set-Content -Path $passingFile -Value ($passingContent -join "`n") -Encoding UTF8

    # ── Write failing tests ──
    $failCount = $failedNames.Count
    $failingContent = @("# Failing Tests — $timestamp", "# Count: $failCount", "# Source: $callerSource (TestLogWriter.psm1 → Write-TestLogs)", "")

    # Summary section: list failed test names first
    if ($failCount -gt 0) {
        $failingContent += "# ── Summary ──"
        $sortedFailed = $failedNames | Sort-Object
        foreach ($name in $sortedFailed) {
            $failingContent += "  - $name"
        }
        $failingContent += @("", "# ── Details ──", "")
    }
    $failingContent += $failing

    # Fallback diagnostics: if no per-test block was captured, include raw reasons
    if ($failCount -gt 0 -and $failing.Count -eq 0) {
        $failingContent += @("# Diagnostic Snippets:", "")

        $snippetLines = $filteredOutput | Where-Object {
            $_ -match "--- FAIL:\s+" -or
            $_ -match "_test\.go:\d+:" -or
            $_ -match "^\s*panic:" -or
            $_ -match "^\s*Expected:" -or
            $_ -match "^\s*Actual:"
        }

        if ($snippetLines) {
            $failingContent += $snippetLines
        }
        else {
            $failingContent += "No detailed failure lines were captured from raw output."
        }

        $failingContent += ""
    }

    # Also capture compilation errors (no === RUN lines at all)
    $hasAnyRun = $filteredOutput | Where-Object { $_ -match "^=== RUN" } | Select-Object -First 1

    if (-not $hasAnyRun) {
        $compileErrors = $filteredOutput | Where-Object {
            $_ -match "\.go:\d+:" -or $_ -match "^#\s+" -or $_ -match "\[build failed\]"
        }

        if ($compileErrors) {
            $failingContent += @("", "# Compilation Errors:", "")
            $failingContent += $compileErrors
            $failCount = $failCount + 1
        }
    }

    $failingContent[1] = "# Count: $failCount"

    Set-Content -Path $failingFile -Value ($failingContent -join "`n") -Encoding UTF8

    $passCount = $passing.Count

    Write-Host ""
    if ($passCount -gt 0) { Write-Success "$passCount passing test(s) → $passingFile" }
    if ($failCount -gt 0) { $s = Get-CallerSource; Write-Fail "$failCount failing test(s) → $failingFile (source: $s)" }
    elseif ($failCount -eq 0) { Write-Success "No failing tests" }
    Write-Host "  Raw output → $rawFile" -ForegroundColor Gray
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Write-TestLogs'
)
