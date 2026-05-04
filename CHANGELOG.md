# Changelog

All notable changes to **enum-v1** are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

The release pipeline extracts the matching `## [vX.Y.Z]` section as the
GitHub Release body — keep entries small, sectioned, and human-readable.

---

## [Unreleased]

### Added
- CI/CD pipeline (`.github/workflows/ci.yml`) with SHA dedup, lint, vulncheck,
  test matrix (4 suites), test summary, and `go build ./...` gate.
- Standalone weekly vulnerability scan workflow.
- Release workflow for `release/**` branches and `v*` tags producing source
  archives + checksums.
- Reusable CI guards: cross-file collision audit and baseline-diff lint gate.

### Changed
- Module path migrated from `gitlab.com/auk-go/enum` to
  `github.com/alimtvnetwork/enum-v1`.
- Dependency `gitlab.com/auk-go/core` replaced with
  `github.com/alimtvnetwork/core-v8` (pinned to `feature/1.5.6`).
