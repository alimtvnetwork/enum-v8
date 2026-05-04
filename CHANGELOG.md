# Changelog

All notable changes to **enum-v1** are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

The release pipeline extracts the matching `## [vX.Y.Z]` section as the
GitHub Release body — keep entries small, sectioned, and human-readable.

---

## [Unreleased]

### Added
- **CI/CD pipeline** (`.github/workflows/ci.yml`) — SHA dedup, `golangci-lint`,
  `govulncheck`, 4-suite parallel test matrix, aggregated test summary,
  60% coverage gate, and `go build ./...` smoke check.
- **Weekly vulnerability scan** (`.github/workflows/vulncheck.yml`) — scheduled
  `govulncheck` run with two-tier classification (fail on third-party,
  warn on stdlib).
- **Release workflow** (`.github/workflows/release.yml`) — triggers on
  `release/**` branches and `v*` tags; produces source archives, SHA-256
  checksums, and a GitHub Release whose body is extracted from this file.
- **Reusable CI guards** (`.github/workflows/ci-guards.yml`):
  - `scripts/ci/check-collisions.py` — per-package identifier collision
    audit (cross-file, case-insensitive, intra-file). GOOS/GOARCH build-tag
    siblings and Exported/unexported accessor pairs are recognised and
    excluded from false positives.
  - `scripts/ci/lint-baseline-diff.py` — lint gate that fails only on
    **new** `golangci-lint` findings; baseline cached per `main` push and
    seeded from `.ci-baselines/golangci-lint.json` on cold cache.
- **Spec docs**: `spec/04-tooling/02-release-pipeline.md`,
  `03-vulnerability-scanning.md`, `04-ci-guards.md`.
- **`CONTRIBUTING.md`** — local dev (`./run.sh`), commit conventions,
  release procedure.
- `.golangci.yml`, `CODEOWNERS`, `.github/PULL_REQUEST_TEMPLATE.md`,
  `.gitattributes`.

### Changed
- Module path migrated from `gitlab.com/auk-go/enum` to
  `github.com/alimtvnetwork/enum-v1`.
- Dependency `gitlab.com/auk-go/core` replaced with
  `github.com/alimtvnetwork/core-v8`, pinned to pseudo-version
  `v1.5.6-0.20260423064907-72bcd64c06b5` (commit `72bcd64` on
  `feature/1.5.6`) so CI can resolve the module deterministically.

