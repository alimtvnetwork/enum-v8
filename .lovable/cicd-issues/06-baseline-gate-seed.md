# 06 — golangci-lint Baseline Gate (Seed-Then-Gate)

## Symptom

`lint-baseline-diff` job in `.github/workflows/ci-guards.yml` either passes everything as warnings (when baseline is empty) or fails on NEW findings (when baseline is populated). Behaviour switches based on the contents of `.ci-baselines/golangci-lint.json`.

## Root Cause

This is the intentional **seed-then-gate** design:

1. **Seed phase:** Empty baseline → all current findings recorded as warnings, `update-baseline` step on `main` push regenerates the baseline.
2. **Gate phase:** Populated baseline → `scripts/ci/lint-baseline-diff.py` diffs current run against baseline; CI fails on findings not in the baseline.

Toolchain: `golangci-lint v1.64.8`.

## Fix / Workaround

Working as designed. Do not "simplify" by removing the seed-then-gate logic.

To intentionally accept a new lint finding into the baseline, push to `main` — the `update-baseline` step regenerates `.ci-baselines/golangci-lint.json`.

## Status

✅ Working as designed.

## Related

- `mem://features/ci-tooling.md`
- `.github/workflows/ci-guards.yml`
- `scripts/ci/lint-baseline-diff.py`
