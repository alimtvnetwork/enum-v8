# Deferred / Skipped Tasks

> Tasks the user has explicitly deferred, parked under their own responsibility, or skipped. The AI must NOT propose these unprompted.

## ⏭️ Manual user actions (parked)

### Task A — Push `cross-repo/core-v9/` mirror to upstream GitHub

- **Why parked:** Credential-bound. The AI cannot perform git pushes to external repos.
- **Trigger:** Whenever main-repo CI changes (`.github/workflows/`, `.golangci.yml`, `.github/dependabot.yml`, `.ci-baselines/`, `scripts/ci/test_*.py`, `CHANGELOG.md`).
- **AI's role:** Mirror the file changes into `cross-repo/core-v9/` and commit. NEVER attempt the push.

### Task W — Upstream `core-v9` `go.mod` rename + tag `v1.5.8`

- **Why parked:** Requires editing a different repository the AI does not have write access to.
- **AI's role:** Re-state the blocker every time it surfaces; never propose pseudo-version workarounds that need non-existent predecessor tags.

## 🚫 Suggestions the user has NOT accepted (do not re-propose unprompted)

### Pinning Go toolchain to 1.22 as a stopgap for W

- **Offered:** Multiple sessions including Cycle 13 turn.
- **Status:** User has neither accepted nor explicitly rejected — they keep saying "next" instead.
- **Rule:** Mention as a one-line option only when build failures from the dual-path bridge resurface; do NOT re-pitch on every "next".

## 🚫 Hard prohibitions (see `.lovable/strictly-avoid.md`)

These are not "skipped tasks" — they are forbidden categories. Listed here for completeness of the avoid surface:

- Email-based notifications of any kind.
- Reintroducing `core-v8` outside `cross-repo/core-v9/`.
- Writing `tests/integratedtests/` in new spec content.
- Renaming or modifying `cross-repo/core-v9/` in v1→v2 or v8→v9 sweeps.
