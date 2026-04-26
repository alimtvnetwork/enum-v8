# CI Pipeline Specification

## Overview

Continuous Integration pipeline for the Go project, running automated quality gates on every merge to protected branches.

## Trigger Conditions

- **Push** to any branch
- **Pull requests** targeting any branch

> CI runs on all branches to catch issues early.

## Pipeline Stages

### 1. SHA Cache Check (Pre-flight)

Before running any jobs, the pipeline checks if the current commit SHA has already been successfully built.

- Uses GitHub Actions cache keyed by `ci-passed-<SHA>`
- If cache hit → all jobs exit with success immediately (no redundant work)
- If cache miss → full pipeline runs
- On successful completion of all jobs → cache entry is written for future runs

**Rationale**: Prevents re-running identical checks on merge commits or force-pushes to the same SHA.

### 2. Lint

- **Tool**: `golangci-lint` (latest, 5-minute timeout)
- **Config**: `.golangci.yml` (project root)
- Includes `govet`, `revive`, `gosec`, `errcheck`, `staticcheck`, and 20+ linters

### 3. Vet

- **Tool**: `go vet ./...`
- Explicit `go vet` step independent of golangci-lint for clarity and redundancy
- Catches issues like misuse of `printf` verbs, unreachable code, and struct tag errors

### 4. Test & Coverage

- **Command**: `go test -v -race -count=1 -parallel=4 -coverprofile=coverage.out -covermode=atomic ./...`
- Verbose output with race detection enabled
- All tests run to completion (`set +e`) — failures don't short-circuit
- **Minimum coverage threshold**: 60%
- Test results and coverage reports are posted as PR comments (updated in-place)
- Coverage artifacts uploaded for download

### 5. Security / Vulnerability Scan

- **Tool**: `govulncheck ./...`
- Scans all dependencies for known Go vulnerabilities
- Blocks the pipeline on any found vulnerability

### 6. Build

- **Command**: `go build ./...`
- Runs after lint, vet, test, and security pass
- Validates the entire module compiles cleanly

### 7. Release (conditional)

- **Trigger**: Push to `release/**` branches only
- Uses GoReleaser for cross-compilation and GitHub Release publishing
- Version tag derived from branch name (e.g., `release/v1.2.0` → tag `v1.2.0`)

## Notification Policy

**Strict no-notification policy**:

- ❌ No email notifications
- ❌ No Slack/Discord/webhook integrations
- ❌ No external notification services

**Allowed reporting**:

- ✅ Console logs (GitHub Actions output)
- ✅ PR comments (test results, coverage summary)
- ✅ GitHub Actions status checks (required for branch protection)

## SHA Cache Behavior

| Scenario | Behavior |
|----------|----------|
| New commit, never built | Full pipeline runs |
| Same SHA, previously passed | All jobs skip with success |
| Same SHA, previously failed | Full pipeline runs (no cache on failure) |
| Force-push same content | Cache hit → skip |

## Go Version

All jobs use Go `1.25` (pinned in workflow).

## Dependencies

- `actions/checkout@v4`
- `actions/setup-go@v5`
- `actions/cache@v4` (SHA dedup)
- `golangci/golangci-lint-action@v6`
- `actions/upload-artifact@v4`
- `actions/github-script@v7` (PR comments)
- `goreleaser/goreleaser-action@v6` (release only)
