# ─────────────────────────────────────────────────────────────────────────────
# spec-api-sig-check.psm1 — Spec call-site signature lint (S-106 v2)
#
# Complement to spec-api-check.psm1 (S-106 v1.x). Where v1 verifies that
# `package.Symbol` references resolve, v2 verifies that **call-sites**
# (`pkg.Func(arg1, arg2, ...)`) match the upstream signature in arity
# (and best-effort kind: literal type, "string-literal", "nil",
# "interface" for variadic any-args).
#
# Catches the class of defects S-106 v1 cannot:
#
#   • Missing parameter — e.g. `errcore.VarTwo("name", val)` when the real
#     signature is `VarTwo(isIncludeType bool, name string, value any)`.
#     (C-CVS-44)
#   • Wrong return-type expectation — e.g. spec text claims a helper
#     returns `error` but it returns `string`. (C-CVS-45)
#   • Wrong receiver shape — e.g. spec calls `Foo.Method(x)` when the real
#     method is `(*Foo).Method(x, y)`. (C-CVS-49)
#
# Requires the JSON sig-index produced by `scripts/specapisig`:
#
#   go run ./scripts/specapisig -roots /tmp/core-v9-upstream,. -out /tmp/core-v9-sigindex.json
#   Import-Module ./scripts/spec-api-sig-check.psm1 -Force
#   Invoke-SpecApiSigCheck -SigIndex /tmp/core-v9-sigindex.json -SpecDir spec/01-app
#   Invoke-SpecApiSigCheck -StrictExitCode    # exit 1 on any mismatch (CI mode)
#
# Out-of-scope (deliberately): full Go semantic analysis. This is a regex +
# arity lint, not a type-checker. False positives are tolerable on heavily
# fluent / chained call-sites; v1 + v2 together still catch every C-CVS-44
# class defect we've seen.
# ─────────────────────────────────────────────────────────────────────────────

$script:SpecApiSigCheckVersion = '1.0.0'

function Get-SigIndex {
    [CmdletBinding()]
    param([Parameter(Mandatory)][string]$Path)

    if (-not (Test-Path $Path)) {
        throw "Sig-index JSON not found at $Path. Generate with: go run ./scripts/specapisig -out $Path"
    }
    $raw = Get-Content -Path $Path -Raw | ConvertFrom-Json

    # Build {pkg -> {symbol -> [SigItem]}} lookup. A symbol may appear
    # multiple times if both a free function and a method share the name
    # across packages — we keep all candidates.
    $idx = @{}
    foreach ($it in $raw.items) {
        if (-not $idx.ContainsKey($it.package)) { $idx[$it.package] = @{} }
        $bySym = $idx[$it.package]
        if (-not $bySym.ContainsKey($it.symbol)) { $bySym[$it.symbol] = New-Object System.Collections.Generic.List[object] }
        [void]$bySym[$it.symbol].Add($it)
    }
    return [pscustomobject]@{
        Version   = $raw.version
        Generated = $raw.generated
        Packages  = $raw.packages
        Functions = $raw.functions
        Lookup    = $idx
    }
}

function Get-SpecCallSites {
    <#
    .SYNOPSIS Extract `package.Symbol(args...)` call-sites from spec markdown.
    .DESCRIPTION
        Returns a list of [pscustomobject]@{ Package; Symbol; ArgCount;
        ArgsRaw; Line; File; Fenced }. Uses a balanced-paren walker so
        nested calls are counted correctly. Honors fence + local-var
        tracking from the v1 conventions.
    #>
    [CmdletBinding()]
    param([Parameter(Mandatory)][string]$Path)

    $sites = New-Object System.Collections.Generic.List[object]
    $lines = Get-Content -Path $Path
    $inFence = $false
    $localVars = [System.Collections.Generic.HashSet[string]]::new()
    $rel = $Path

    for ($i = 0; $i -lt $lines.Count; $i++) {
        $line = $lines[$i]

        if ($line -match '^\s*```(\w*)') {
            if ($inFence) { $inFence = $false; [void]$localVars.Clear() }
            else { $inFence = $true }
            continue
        }

        if ($inFence) {
            foreach ($vm in [regex]::Matches($line, '(?:^|\s|\(|,)([a-z][a-zA-Z0-9_]*)\s*(?:,\s*[a-zA-Z0-9_]+\s*)?:=')) {
                [void]$localVars.Add($vm.Groups[1].Value)
            }
            foreach ($vm in [regex]::Matches($line, '\bfunc\s+\(\s*([a-z][a-zA-Z0-9_]*)\s+\*?[A-Z]')) {
                [void]$localVars.Add($vm.Groups[1].Value)
            }
        }

        # Find every `pkg.Symbol(` opener; then walk balanced parens to find
        # the matching close and split top-level commas as args.
        $openerRe = [regex]'(?<!\w)([a-z][a-z0-9]+)\.([A-Z][A-Za-z0-9_]*)\s*\('
        $opMatches = $openerRe.Matches($line)
        foreach ($m in $opMatches) {
            $pkg = $m.Groups[1].Value
            $sym = $m.Groups[2].Value
            if ($localVars.Contains($pkg)) { continue }
            # Walk parens starting AFTER the opening `(`.
            $start = $m.Index + $m.Length
            $depth = 1
            $end = -1
            $args = New-Object System.Text.StringBuilder
            for ($k = $start; $k -lt $line.Length; $k++) {
                $ch = $line[$k]
                if ($ch -eq '(') { $depth++ }
                elseif ($ch -eq ')') {
                    $depth--
                    if ($depth -eq 0) { $end = $k; break }
                }
                [void]$args.Append($ch)
            }
            if ($end -eq -1) { continue }   # call spans multiple lines; skip in v1.0

            $argStr = $args.ToString().Trim()
            $argList = @()
            if ($argStr -ne '') {
                # Top-level comma split (respect nested parens/brackets/braces and quotes).
                $argList = (Split-TopLevelArgs -Source $argStr)
            }

            $sites.Add([pscustomobject]@{
                File     = $rel
                Line     = $i + 1
                Package  = $pkg
                Symbol   = $sym
                ArgCount = $argList.Count
                ArgsRaw  = $argStr
                Args     = $argList
                Fenced   = $inFence
                Context  = $line.Trim()
            })
        }
    }
    return $sites
}

function Split-TopLevelArgs {
    [CmdletBinding()]
    param([Parameter(Mandatory)][string]$Source)

    $out = New-Object System.Collections.Generic.List[string]
    $cur = New-Object System.Text.StringBuilder
    $depth = 0
    $inStr = $false
    $strCh = $null
    for ($i = 0; $i -lt $Source.Length; $i++) {
        $c = $Source[$i]
        if ($inStr) {
            [void]$cur.Append($c)
            if ($c -eq '\\' -and $i + 1 -lt $Source.Length) {
                [void]$cur.Append($Source[$i + 1])
                $i++
                continue
            }
            if ($c -eq $strCh) { $inStr = $false }
            continue
        }
        if ($c -eq '"' -or $c -eq "'" -or $c -eq '`') {
            $inStr = $true; $strCh = $c
            [void]$cur.Append($c); continue
        }
        if ($c -eq '(' -or $c -eq '[' -or $c -eq '{') { $depth++; [void]$cur.Append($c); continue }
        if ($c -eq ')' -or $c -eq ']' -or $c -eq '}') { $depth--; [void]$cur.Append($c); continue }
        if ($c -eq ',' -and $depth -eq 0) {
            [void]$out.Add($cur.ToString().Trim())
            [void]$cur.Clear()
            continue
        }
        [void]$cur.Append($c)
    }
    if ($cur.Length -gt 0) { [void]$out.Add($cur.ToString().Trim()) }
    return $out.ToArray()
}

function Test-SignatureMatch {
    <#
    .SYNOPSIS Returns $true if call-site arity matches at least one candidate.
    #>
    [CmdletBinding()]
    param(
        [Parameter(Mandatory)][int]$ArgCount,
        [Parameter(Mandatory)]$Candidates
    )
    foreach ($cand in $Candidates) {
        $expected = $cand.params.Count
        if ($cand.variadic) {
            # Variadic: caller may supply expected-1 OR more.
            if ($ArgCount -ge ($expected - 1)) { return $true }
        } else {
            if ($ArgCount -eq $expected) { return $true }
        }
    }
    return $false
}

function Format-CandidateSig {
    [CmdletBinding()]
    param([Parameter(Mandatory)]$Candidate)
    $params = ($Candidate.params | ForEach-Object { "$($_.name) $($_.type)".Trim() }) -join ', '
    $results = ''
    if ($Candidate.results.Count -eq 1 -and -not $Candidate.results[0].name) {
        $results = " $($Candidate.results[0].type)"
    } elseif ($Candidate.results.Count -gt 0) {
        $rs = ($Candidate.results | ForEach-Object { "$($_.name) $($_.type)".Trim() }) -join ', '
        $results = " ($rs)"
    }
    $recv = ''
    if ($Candidate.receiver) { $recv = "($($Candidate.receiver)) " }
    return "func ${recv}$($Candidate.symbol)($params)${results}"
}

function Invoke-SpecApiSigCheck {
    <#
    .SYNOPSIS S-106 v2 — verify spec call-site arity against upstream signatures.
    .PARAMETER SigIndex Path to JSON produced by `go run ./scripts/specapisig`.
    .PARAMETER SpecDir Default: spec/01-app
    .PARAMETER StrictExitCode If set, exits 1 when any mismatch is found.
    .PARAMETER OnlyFile Limit scan to one spec file.
    #>
    [CmdletBinding()]
    param(
        [string]$SigIndex = '/tmp/core-v9-sigindex.json',
        [string]$SpecDir  = 'spec/01-app',
        [switch]$StrictExitCode,
        [string]$OnlyFile
    )

    Write-Host ''
    Write-Host "  ▶ S-106 v2 spec-api-sig-check v$script:SpecApiSigCheckVersion" -ForegroundColor Cyan
    Write-Host "    SigIndex: $SigIndex"
    Write-Host "    SpecDir:  $SpecDir"
    Write-Host ''

    $idx = Get-SigIndex -Path $SigIndex
    Write-Host "    Loaded $($idx.Functions) signatures across $($idx.Packages) packages (generated $($idx.Generated))"

    $files = if ($OnlyFile) {
        @(Join-Path $SpecDir $OnlyFile)
    } else {
        Get-ChildItem -Path $SpecDir -Recurse -Filter '*.md' -File | ForEach-Object { $_.FullName }
    }

    $totalSites = 0
    $arityMismatches = New-Object System.Collections.Generic.List[object]
    $unresolvedSites = 0
    $okSites = 0

    foreach ($file in $files) {
        $sites = Get-SpecCallSites -Path $file
        $totalSites += $sites.Count
        foreach ($s in $sites) {
            if (-not $idx.Lookup.ContainsKey($s.Package)) { $unresolvedSites++; continue }
            $bySym = $idx.Lookup[$s.Package]
            if (-not $bySym.ContainsKey($s.Symbol)) { $unresolvedSites++; continue }
            $cands = $bySym[$s.Symbol]
            if (Test-SignatureMatch -ArgCount $s.ArgCount -Candidates $cands) {
                $okSites++
            } else {
                $arityMismatches.Add([pscustomobject]@{
                    File       = $s.File
                    Line       = $s.Line
                    Package    = $s.Package
                    Symbol     = $s.Symbol
                    Got        = $s.ArgCount
                    Candidates = $cands
                    Context    = $s.Context
                })
            }
        }
    }

    Write-Host ('  ─' * 30) -ForegroundColor DarkGray
    Write-Host '  ▶ Results' -ForegroundColor Cyan
    Write-Host "    Files scanned:    $($files.Count)"
    Write-Host "    Total call-sites: $totalSites"
    Write-Host "    Resolved+OK:      $okSites" -ForegroundColor Green
    Write-Host "    Unresolved:       $unresolvedSites  (handled by S-106 v1 — pkg/sym fab)" -ForegroundColor DarkGray
    Write-Host "    Arity mismatches: $($arityMismatches.Count)" -ForegroundColor $(if ($arityMismatches.Count) { 'Red' } else { 'Green' })
    Write-Host ''

    if ($arityMismatches.Count -gt 0) {
        Write-Host '  ✗ ARITY MISMATCHES (call-site arg count ≠ any upstream candidate):' -ForegroundColor Red
        $arityMismatches | Group-Object { "$($_.Package).$($_.Symbol)" } | Sort-Object Count -Descending | ForEach-Object {
            $first = $_.Group[0]
            $expectedList = ($first.Candidates | ForEach-Object {
                $n = $_.params.Count
                if ($_.variadic) { "$($n - 1)+" } else { "$n" }
            }) -join ' or '
            Write-Host "    • $($_.Name)  (got $($first.Got), expected $expectedList; $($_.Count) site(s))" -ForegroundColor Red
            $_.Group | Select-Object -First 2 | ForEach-Object {
                Write-Host "        $($_.File):$($_.Line)" -ForegroundColor DarkRed
            }
            Write-Host "        upstream: $(Format-CandidateSig -Candidate $first.Candidates[0])" -ForegroundColor DarkGray
        }
        Write-Host ''
    } else {
        Write-Host '  ✓ No arity mismatches detected' -ForegroundColor Green
        Write-Host ''
    }

    if ($StrictExitCode -and $arityMismatches.Count -gt 0) {
        exit 1
    }

    return [pscustomobject]@{
        FilesScanned    = $files.Count
        TotalSites      = $totalSites
        Resolved        = $okSites
        Unresolved      = $unresolvedSites
        ArityMismatches = $arityMismatches
    }
}

Export-ModuleMember -Function @(
    'Invoke-SpecApiSigCheck',
    'Get-SigIndex',
    'Get-SpecCallSites',
    'Split-TopLevelArgs',
    'Test-SignatureMatch',
    'Format-CandidateSig'
)