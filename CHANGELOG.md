# Changelog

All notable changes to **enum-v4** are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

The release pipeline extracts the matching `## [vX.Y.Z]` section as the
GitHub Release body — keep entries small, sectioned, and human-readable.

---

## [Unreleased]

### Changed
- **Cycle 19 spec audit (Task AB pass 1)** — upstream `core-v9 v1.5.8`
  cloned to `/tmp/core-v9-upstream`; first ❓→ground-truth promotion pass on
  `spec/01-app/09-converters.md`. Result: **23 ❓ → 10 ✅ / 5 ❌ / 8 ❓**
  (verifiable score 66.7 %). Surfaced **5 NEW contradictions**
  (C-CVS-11..15): four HIGH (`StringTo.Integer64`, `StringTo.Float32`,
  `StringTo.Bool`, `PrettyJson.String`/`.FromAny`) and one CRITICAL
  (entire `typesconv` §2 + §4.3 example fabricated). All 5 are blocked
  pending a freeze waiver for `spec/01-app/`. Spawned action items
  AJ-01..03. Also corrected 2 stale Core-memory items in `mem://index.md`
  (M-CVS-01: `enum-v3`→`enum-v4` module name; M-CVS-02: upstream `go.mod`
  rename declared complete + `replace` bridge removal). See
  `spec/07-code-vs-spec-audits/20-cycle19-AB-converters-promotion.md`.
  Spec changelog → `spec-v0.34.0`.
- **Cycle 18 spec audit (Task AA + closes Task AH)** — closed
  `spec/02-app-issues/` (11 files, 402 lines) at **100 % verifiable** (21 ✅ /
  5 ❓ audit-history). Raised and resolved **5 LOW drifts (D-CVS-56 →
  D-CVS-60)**: 1 stale README index (5 open vs reality 9 resolved) + 4
  upstream-vs-`enum-v4` scope footnotes on the historical resolution files
  (`02-internal-package-coverage-policy.md`, `03-getassert-undocumented-api.md`,
  `04-testwrappers-public-surface.md`, `05-missing-params-go-files.md`).
  **🎉 Marks Task AH (cross-`spec/` cleanup sweep) Done** — every directory
  under `spec/` outside the immutable history folders is now baselined. See
  `spec/07-code-vs-spec-audits/19-cycle18-app-issues.md`. Spec changelog
  bumped to **spec-v0.33.0**.
- **Cycle 17 spec audit (Task AA + AH partial)** — closed `spec/04-tooling/`
  (10 files, 2 553 lines) at **100 % verifiable** (22 ✅ / 8 ❓ workflow-
  internal). Raised and resolved **7 LOW drifts (D-CVS-49 → D-CVS-55)** in the
  same cycle: 2 broken `cross-repo/core-v9/` paths in `00-overview.md`, 1
  missing-precedent in `04-bootstrap-into-new-repo.md` §7 (the AH-tracked
  occurrence), and 4 stale `enum-v2`/`cross-repo/core-v9` tokens in
  `06-cross-repo-sync.md` (line 80 template comment carried both stale tokens).
  Each fix includes a Core-memory clarification that `cross-repo/core-v8/`
  intentionally keeps its historical name even though the import path is
  `core-v9`. See `spec/07-code-vs-spec-audits/18-cycle17-tooling.md`. Spec
  changelog bumped to **spec-v0.32.0**.
- **Cycle 16 spec audit (Task AA + AH partial)** — closed
  `spec/03-powershell-test-run/` (9 files, 2 519 lines) at **100 % verifiable**
  (22 ✅ / 6 ❓; the 6 ❓ are runner-internal behaviours requiring a direct
  `scripts/*.psm1` probe). Raised and resolved **5 LOW drifts (D-CVS-44 →
  D-CVS-48)** in the same cycle via top-of-file consumer-coverage callouts
  (`01-overview.md`, `04-pre-commit-api-checker.md`, `08-generic-go-test-coverage-runner.md`,
  `09-ai-agent-complete-reference.md`) plus one inline rewrite
  (`06-coverage-prompt-generator.md` line 71). Folds in Task AH debt for this
  directory. See `spec/07-code-vs-spec-audits/17-cycle16-powershell-test-run.md`.
  Spec changelog bumped to **spec-v0.31.0**.
- **`spec/01-app/` DRIFT-FROZEN (Task AI)** — declared the directory closed for
  code-vs-spec drift work in `spec/CHANGELOG.md` as **spec-v0.30.0**. Allowed
  future edits: AB-driven ❓→✅ promotions, AC re-audit of §07/§09,
  upstream-API-change additions (paired with a new audit cycle row), typo/
  formatting fixes. Drift work moves to `spec/03-powershell-test-run/`,
  `spec/04-tooling/`, `spec/02-app-issues/` (Cycles 16+). Scoreboard top-line
  updated with 🧊 freeze marker.
- **Cycle 15 spec audit (Task AA)** — closed `spec/06-testing-guidelines/`
  directory at **100 % of its verifiable subset** (32 claims sampled across 10
  files; 22 ✅ / 10 ❓ pending task AB). Resolved one LOW drift (D-CVS-43) by
  adding an `enum-v4` consumer-coverage callout to `spec/06-testing-guidelines/README.md`
  and a `⚠️ Scope` warning to `01-folder-structure.md`. See
  `spec/07-code-vs-spec-audits/16-cycle15-testing-guidelines.md`. Spec changelog
  bumped to **spec-v0.29.0**.
- **core-v9 API migration (Task AM)** — Applied all confirmed `core-v9 v1.5.8`
  namespace rewrites across `enum-v4` Go source: `coredynamic.TypeName(...)` →
  `coredynamic.SafeTypeName(...)`, `converters.AnyToValueString(x)` →
  `converters.AnyTo.ValueString(x)`, `converters.Any.ToFullNameValueString` →
  `converters.AnyTo.ToFullNameValueString`, and the remaining
  `converters.StringTo*` calls → `converters.StringTo.*` methods.

### Fixed
- **Task AM / `tests/creationtests` compile blocker** — removed the obsolete
  package-level converter calls that caused `undefined: converters.StringToInteger`,
  `undefined: converters.StringToIntegerWithDefault`,
  `undefined: converters.StringToIntegerDefault`, and `undefined: converters.StringToByte`
  after upgrading to `core-v9 v1.5.8`.

- **`scripts/CoverageRunner.psm1`, `scripts/CoverageCompileCheck.psm1`** — second
  pass at the `Blocked: (root) — : no such file or directory` failure that
  surfaced during parallel `./run.ps1 -tc` runs. Even with the prior multi-root
  probe, a failing `go list ./...` (e.g. when the upstream `core-v9`
  module-loader error fires) could leak stderr fragments past the regex filter
  as a single-package "phantom (root)" entry, which then aborted the whole run
  before any coverage % was reported. Hardening:
  1. Discovery now captures `$LASTEXITCODE` from `go list`, anchors the
     keep-filter to the actual `module` declared in `go.mod`, rejects lines
     containing whitespace / `...` / known stderr prefixes (`go:`, `warning:`,
     `matched no packages`, `package`, `can't load`, `cannot find`, `err:`,
     `# `), and aborts loudly with the raw `go list` output if nothing valid
     survives — instead of silently producing one bogus package path.
  2. `Get-PackageShortName` replaces the bare `'.*(integratedtests|creationtests)/?'`
     regex. It always returns a non-empty label (trailing test segment → last
     path segment → full path) so blocked-package summaries never collapse to
     the unhelpful `(root)`. Blocked log lines now include the full import path
     next to the short name.
  3. All three early-abort code paths in `Invoke-TestCoverage` now call
     `Write-PhaseSummaryBox` before returning, so the dashboard summary (and
     any `Coverage Run: fail` phase status) renders even when the pipeline
     aborts before the merge step. This is what was hiding the coverage
     percentage — the run never reached `Write-CoverageConsoleSummary` because
     the aborted summary box was suppressed.

### Previously
- **`scripts/CoverageRunner.psm1`** — pre-coverage compile check no longer
  aborts with `Blocked: : no such file or directory`. Two root causes:
  1. Test discovery was hard-coded to `./tests/integratedtests/...`, which
     does not exist in this repo (tests live under `./tests/creationtests/`).
     `go list` errors were being coerced into a single empty package path
     and handed to `go test`. The probe now tries both directory names,
     skips paths not present on disk, and filters `go list` warning/error
     lines so only valid import paths reach the compile check.
  2. The in-package-test scan hard-coded the `core-v9` module prefix when
     stripping import paths to filesystem-relative paths. It now reads the
     `module` line from `go.mod` so the same script works in `enum-v4`,
     `core-v9`, and any other module.
- **`scripts/CoverageRunner.psm1`, `scripts/CoverageCompileCheck.psm1`** —
  `shortName` regex updated from `'.*integratedtests/?'` to
  `'.*(integratedtests|creationtests)/?'` so blocked-package labels render
  the trailing path segment in either layout.

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
- **Spec docs**: `spec/04-tooling/00-overview.md` (tooling index/landing
  page), `02-release-pipeline.md`, `03-vulnerability-scanning.md`,
  `04-bootstrap-into-new-repo.md`, `04-ci-guards.md`,
  `05-branch-protection.md`, `06-cross-repo-sync.md`.
- **`CONTRIBUTING.md`** — local dev (`./run.sh`), commit conventions,
  release procedure.
- `.golangci.yml`, `CODEOWNERS`, `.github/PULL_REQUEST_TEMPLATE.md`
  (now with structured local / CI-guard / cross-repo checklists),
  `.gitattributes`.
- **Dependabot** (`.github/dependabot.yml`) — weekly `gomod` and
  `github-actions` updates, grouped minor/patch PRs, scheduled Mondays
  09:00 Asia/Kuala_Lumpur.
- **Cross-repo staging** under `cross-repo/core-v9/` — mirrored
  workflows, `.golangci.yml`, `dependabot.yml`, baselines, and a
  README install guide so the `core-v9` upstream can adopt the same
  CI surface (governed by `spec/04-tooling/06-cross-repo-sync.md`).
- **Python regression tests** for the CI guards:
  `scripts/ci/test_check_collisions.py` (22 cases covering build-tag
  collapsing, accessor pairs, decl parsing, string/comment skipping,
  per-package scoping) and `scripts/ci/test_lint_baseline_diff.py`
  (15 cases covering load/identity rules, seeding mode, gate mode,
  exit codes, summary counts).

### Changed
- Module path migrated from `gitlab.com/auk-go/enum` to
  `github.com/alimtvnetwork/enum-v4`.
- **Core dependency renamed** `github.com/alimtvnetwork/core-v9` →
  `github.com/alimtvnetwork/core-v9` across all 307 source files
  (`go.mod`, all package imports, spec docs, CI configs, coverage
  scripts, PR template). The `cross-repo/core-v9/` staging directory
  name is intentionally retained — it tracks the upstream repo name,
  not the module path. Pseudo-version pin
  `v1.5.6-0.20260423064907-72bcd64c06b5` carries over unchanged.
- Dependency `gitlab.com/auk-go/core` replaced with
  `github.com/alimtvnetwork/core-v9`, pinned to pseudo-version
  `v0.0.0-20260423064907-72bcd64c06b5` (commit `72bcd64` on
  `feature/1.5.6`) so CI can resolve the module deterministically.
- **`go.mod` pseudo-version downgraded** from
  `v1.5.6-0.<date>-<sha>` to `v0.0.0-<date>-<sha>`. The `v1.5.6-0.`
  form requires a preceding `v1.5.5` tag on the upstream `core-v9`
  repo, which doesn't exist (the v8 repo had v1.5.5; v9 was just
  renamed and has no tags yet). The `v0.0.0-` form has no predecessor
  requirement. Re-pin to a real `vX.Y.Z` tag once `core-v9` upstream
  ships its first tagged release.
- **`go.mod` rename-bridge `replace` directive** — the upstream
  `core-v9` repo's own `go.mod` still declares
  `module github.com/alimtvnetwork/core-v9`; Go enforces import-path
  / module-path equality so the v9 path can't load directly. Until
  upstream commits its `module github.com/alimtvnetwork/core-v9`
  line, `replace github.com/alimtvnetwork/core-v9 =>
  github.com/alimtvnetwork/core-v9 v0.0.0-<date>-<sha>` resolves the
  v9 import path to the v8 artifact at the same pinned commit. All
  source-code imports stay on `core-v9`; only the resolution target
  is bridged. Remove the `replace` once upstream's `go.mod` is fixed.

### CI
- `ci-guards.yml` gained a `python-tests` job that runs all
  `scripts/ci/test_*.py` via `unittest discover`. The existing
  `collision-audit` and `lint-baseline-diff` jobs now `needs:
  python-tests` so a broken gate script fails fast before the slower
  jobs spend CI minutes producing meaningless results.
- `scripts/CoveragePreChecks.psm1` — auto-fixer and bracecheck steps
  now skip gracefully (with `Register-Phase ... "skip"`) when
  `scripts/autofix/` or `scripts/bracecheck/` are absent from the
  repo, instead of hard-failing the entire `./run.ps1 -tc` run.
- **`scripts/bracecheck/`** (NEW Go tool, ~210 lines + README) —
  fast syntax pre-check. Lexical brace/bracket/paren balance
  validation (skips strings, runes, comments) plus a full
  `parser.AllErrors` pass over every `.go` file. Reports issues as
  `<relpath>:<line>:<col>: <message>`. Verified clean on 637 files.
- **`scripts/autofix/`** (NEW Go tool, ~165 lines + README) —
  conservative auto-fixer. Trims trailing whitespace, collapses 3+
  blank lines to 2, ensures one trailing newline, runs
  `format.Source`. Idempotent. Files that don't parse are skipped
  with a warning so bracecheck pinpoints the syntax issue. Supports
  `--dry-run`. With both tools restored, `./run.ps1 -tc` no longer
  prints the "scripts/autofix/ not present" skip notice.
- **`.github/workflows/python-tests.yml`** (NEW) — standalone runner
  for the CI-guard Python tests, triggered on `v*` tags, manual
  dispatch, and `scripts/ci/**` changes. Matrix tests across Python
  3.10/3.11/3.12. Complements the in-line `python-tests` job in
  `ci-guards.yml` by also catching releases and long-lived branches.

### Docs
- `CONTRIBUTING.md` — pre-push checklist rewritten as checkboxes
  mirroring `.github/PULL_REQUEST_TEMPLATE.md`; Spec References now
  links to `spec/04-tooling/00-overview.md` plus all six tooling
  spec files (00–06).
- `.ci-baselines/README.md` — fully documents the seed-then-gate
  workflow: seeding mode (warnings, never blocks), gating mode
  (NEW/FIXED/UNCHANGED diff with exit codes), mode-transition
  commands, and reviewer guidance for baseline drift.

