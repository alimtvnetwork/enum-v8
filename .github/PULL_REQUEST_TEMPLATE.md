## Summary

<!-- What changed and why. Link related issues. -->

## Type of change

- [ ] Bug fix
- [ ] New feature / enum
- [ ] Refactor (no behavior change)
- [ ] Docs / spec update
- [ ] CI / tooling

## Checklist

### Local quality gates
- [ ] `./run.sh tc` (or `./run.ps1 TC`) passes locally with ≥ 60% coverage
- [ ] `golangci-lint run --timeout=5m` is clean
- [ ] `go vet ./...` is clean
- [ ] `govulncheck ./...` is clean (or known-acceptable findings noted below)

### CI guards (see `spec/04-tooling/04-ci-guards.md`)
- [ ] **collision-audit**: no new build-tag or accessor-pair collisions introduced
      (`python3 scripts/ci/check-collisions.py`)
- [ ] **lint-baseline-diff**: no NEW `golangci-lint` findings vs `.ci-baselines/golangci-lint.json`
      — if the baseline legitimately needs to grow, explain why in the PR description
- [ ] **test-summary**: PR test/coverage comment reviewed; no silently-skipped suites
- [ ] **vulncheck**: `govulncheck` workflow green on this SHA

### Cross-repo & docs
- [ ] If `.github/workflows/*.yml` changed → mirrored under `cross-repo/core-v9/.github/workflows/`
      (per `spec/04-tooling/06-cross-repo-sync.md`)
- [ ] `CHANGELOG.md` updated under `## [Unreleased]` if user-visible
- [ ] No `@latest` / `@main` pins added in workflows or scripts
- [ ] Dependabot scope unchanged, or `.github/dependabot.yml` updated if a new
      ecosystem/path was introduced
