# CI Baselines

Cached snapshots used by the CI guards in `.github/workflows/ci-guards.yml`.

## `golangci-lint.json`

Baseline of accepted `golangci-lint` findings. The lint gate
(`scripts/ci/lint-baseline-diff.py`) compares a fresh report against
this file and fails only on **new** findings.

### Lifecycle

- **Seeding** — when this file is empty (`{"Issues": []}`) the gate
  emits warnings and passes. Use this state to onboard a repo with
  pre-existing lint debt.
- **Updating** — pushes to `main` regenerate the baseline automatically
  via the `update-baseline` job in `ci-guards.yml`. Reviewers should
  inspect the diff: a sudden surge of new accepted issues likely means
  someone bypassed the gate and the change should be reverted or the
  underlying issues fixed.
- **Manual refresh** — to consciously accept legacy debt locally:

  ```bash
  golangci-lint run --out-format=json ./... > .ci-baselines/golangci-lint.json
  ```

### Identity

A finding is keyed by the 4-tuple `(file, line, linter, message)`.
Column and source snippet are intentionally excluded so cosmetic
re-formatting doesn't churn the baseline.
