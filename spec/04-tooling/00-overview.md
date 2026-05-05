# 04 — Tooling: Overview

> **Status**: spec-v0.7.0 (2026-05-04, Asia/Kuala_Lumpur).
> **Scope**: Index for the tooling subtree — CI/CD pipelines, PowerShell
> toolchain, vulnerability scanning, CI guards, branch protection, and
> cross-repo sync.

This page is the landing index for `spec/04-tooling/`. Each subsection
below points at a single, focused spec file. Read in order on first
visit; afterwards jump straight to the topic you need.

---

## Map

| # | Spec | Audience | Companion code |
|---|------|----------|----------------|
| 01 | [CI Pipeline](./01-ci-pipeline.md) | Repo maintainers | [`.github/workflows/ci.yml`](../../.github/workflows/ci.yml) |
| 02a | [Release Pipeline](./02-release-pipeline.md) | Release managers | [`.github/workflows/release.yml`](../../.github/workflows/release.yml) |
| 02b | [PowerShell Dashboard UI](./02-powershell-dashboard-ui.md) | UX of `run.ps1` | [`run.ps1`](../../run.ps1) |
| 03a | [Vulnerability Scanning](./03-vulnerability-scanning.md) | Security reviewers | [`.github/workflows/vulncheck.yml`](../../.github/workflows/vulncheck.yml) |
| 03b | [PowerShell Implementation](./03-powershell-implementation.md) | AI agents editing `run.ps1` | [`run.ps1`](../../run.ps1) |
| 04a | [Bootstrap Into a New Repo](./04-bootstrap-into-new-repo.md) | New-repo onboarding | [`run.ps1`](../../run.ps1), workflows |
| 04b | [CI Guards](./04-ci-guards.md) | All maintainers | [`.github/workflows/ci-guards.yml`](../../.github/workflows/ci-guards.yml), [`scripts/ci/`](../../scripts/ci/) |
| 05 | [Branch Protection](./05-branch-protection.md) | Repo admins | GitHub repo settings |
| 06 | [Cross-Repo Sync](./06-cross-repo-sync.md) | Cross-repo maintainers | [`cross-repo/core-v9/`](../../cross-repo/core-v9/) |

> The duplicate `02-` / `03-` / `04-` prefixes are historical: each pair
> covers a CI/release topic alongside its PowerShell or guard companion.
> Do not renumber without coordinating cross-references.

---

## Reading paths

**Setting up a new repository**
1. [04a — Bootstrap Into a New Repo](./04-bootstrap-into-new-repo.md)
2. [01 — CI Pipeline](./01-ci-pipeline.md)
3. [04b — CI Guards](./04-ci-guards.md)
4. [05 — Branch Protection](./05-branch-protection.md)
5. [06 — Cross-Repo Sync](./06-cross-repo-sync.md) *(if mirroring to upstream)*

**Cutting a release**
1. [02a — Release Pipeline](./02-release-pipeline.md)
2. [03a — Vulnerability Scanning](./03-vulnerability-scanning.md)

**Hacking on `run.ps1` / `run.sh`**
1. [02b — PowerShell Dashboard UI](./02-powershell-dashboard-ui.md)
2. [03b — PowerShell Implementation](./03-powershell-implementation.md)

**Investigating a CI failure**
1. [04b — CI Guards](./04-ci-guards.md) — for `collision-audit`,
   `lint-baseline-diff`, `test-summary` failures.
2. [03a — Vulnerability Scanning](./03-vulnerability-scanning.md) — for
   `govulncheck` failures.
3. [01 — CI Pipeline](./01-ci-pipeline.md) — for build/test failures.

---

## Cross-references

- Upstream rulebook:
  [`coding-guidelines-v20/spec/12-cicd-pipeline-workflows`](https://github.com/alimtvnetwork/coding-guidelines-v20/tree/main/spec/12-cicd-pipeline-workflows)
- Repo-level entry points: [`README.md`](../../README.md),
  [`CONTRIBUTING.md`](../../CONTRIBUTING.md),
  [`.github/PULL_REQUEST_TEMPLATE.md`](../../.github/PULL_REQUEST_TEMPLATE.md)
- Lint baselines: [`.ci-baselines/README.md`](../../.ci-baselines/README.md)
- Dependabot config: [`.github/dependabot.yml`](../../.github/dependabot.yml)

---

## Maintenance

- When you add a new spec file under `spec/04-tooling/`, add a row to
  the **Map** table above and slot it into the appropriate **Reading
  path**.
- When a workflow file is renamed, update the *Companion code* column.
- Per [06 — Cross-Repo Sync](./06-cross-repo-sync.md), any spec change
  that documents shared tooling should also be reflected in the staged
  `cross-repo/core-v9/` copy where applicable.
