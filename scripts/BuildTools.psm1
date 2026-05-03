# ─────────────────────────────────────────────────────────────────────────────
# BuildTools.psm1 — Build, run, format, vet, tidy, and clean commands
#
# Usage:
#   Import-Module ./scripts/BuildTools.psm1 -Force
#
# Dependencies: Utilities.psm1 (Write-Header, Write-Success, Write-Fail)
# ─────────────────────────────────────────────────────────────────────────────

function Invoke-RunMain {
    <#
    .SYNOPSIS
        Run the main Go application via `go run`.
    .EXAMPLE
        Invoke-RunMain
    #>
    [CmdletBinding()]
    param()

    Write-Header "Running main application"
    go run ./cmd/main/*.go
}

function Invoke-Build {
    <#
    .SYNOPSIS
        Build the Go binary into the build/ directory.
    .EXAMPLE
        Invoke-Build
    #>
    [CmdletBinding()]
    param()

    Write-Header "Building binary"
    $buildDir = "build"
    if (-not (Test-Path $buildDir)) { New-Item -ItemType Directory -Path $buildDir | Out-Null }
    go build -o "$buildDir/cli" ./cmd/main/
    if ($LASTEXITCODE -eq 0) { Write-Success "Build complete: $buildDir/cli" }
    else { $s = Get-CallerSource; Write-Fail "Build failed (source: $s)" }
}

function Invoke-BuildRun {
    <#
    .SYNOPSIS
        Build the binary and then run it.
    .EXAMPLE
        Invoke-BuildRun
    #>
    [CmdletBinding()]
    param()

    Invoke-Build
    if ($LASTEXITCODE -eq 0) {
        Write-Header "Running built binary"
        & ./build/cli
    }
}

function Invoke-Format {
    <#
    .SYNOPSIS
        Format all Go source files using `gofmt -w -s`.
    .EXAMPLE
        Invoke-Format
    #>
    [CmdletBinding()]
    param()

    Write-Header "Formatting Go files"
    gofmt -w -s .
    Write-Success "Formatting complete"
}

function Invoke-Vet {
    <#
    .SYNOPSIS
        Run `go vet` on all packages.
    .EXAMPLE
        Invoke-Vet
    #>
    [CmdletBinding()]
    param()

    Write-Header "Running go vet"
    go vet ./...
    if ($LASTEXITCODE -eq 0) { Write-Success "No issues found" }
    else { $s = Get-CallerSource; Write-Fail "Issues found (source: $s)" }
}

function Invoke-Tidy {
    <#
    .SYNOPSIS
        Run `go mod tidy` to sync module dependencies.
    .EXAMPLE
        Invoke-Tidy
    #>
    [CmdletBinding()]
    param()

    Write-Header "Running go mod tidy"
    go mod tidy
    Write-Success "Tidy complete"
}

function Invoke-Clean {
    <#
    .SYNOPSIS
        Remove build artifacts, coverage reports, and precommit data.
    .EXAMPLE
        Invoke-Clean
    #>
    [CmdletBinding()]
    param()

    Write-Header "Cleaning build artifacts"
    if (Test-Path build) { Remove-Item -Recurse -Force build }
    if (Test-Path tests/coverage.out) { Remove-Item tests/coverage.out }
    $coverDir = Join-Path $global:DataDir "coverage"
    if (Test-Path $coverDir) { Remove-Item -Recurse -Force $coverDir; Write-Success "Removed coverage reports" }
    $precommitDir = Join-Path $global:DataDir "precommit"
    if (Test-Path $precommitDir) { Remove-Item -Recurse -Force $precommitDir; Write-Success "Removed precommit reports" }
    Write-Success "Clean complete"
}

# ═══════════════════════════════════════════════════════════════════════════════
# Module Export
# ═══════════════════════════════════════════════════════════════════════════════

Export-ModuleMember -Function @(
    'Invoke-RunMain',
    'Invoke-Build',
    'Invoke-BuildRun',
    'Invoke-Format',
    'Invoke-Vet',
    'Invoke-Tidy',
    'Invoke-Clean'
)
