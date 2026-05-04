# 01 — Code-vs-Spec Drift Scoreboard (Living Document)

> **Single source of truth** for code-vs-spec drift. Updated after every cycle.

## Current MEASURED drift score: **41.7 / 100** *(Cycle 1, n=12 claims, n=1 spec section)*

> **Caveat:** sample size is one spec section. Score is dominated by stale `core-v8` prose that didn't get rewritten during the v8→v9 rename. After applying the 5 LOW drift fixes (one-line spec edits), projected score for §03 is ~91.7 (11/12). After resolving C-CVS-01 (the missing `tests/integratedtests/` folder question), projected ~100% for §03.

## Cycle history

| Date | Cycle | Spec audited | Claims | ✅ Match | ⚠️ Drift | ❌ Contradiction | Score |
|------|-------|--------------|--------|---------|---------|------------------|-------|
| 2026-05-04 | 1 | `01-app/03-import-conventions.md` | 12 | 5 | 5 | 2 | **41.7%** |

## Open drift findings

| ID | Title | Severity | Spec ref | Code ref | Resolution path |
|----|-------|----------|----------|----------|-----------------|
| D-CVS-01 | Spec §03 line 4 says "consumes `core-v8`" — stale, should be `core-v9` | LOW | `spec/01-app/03-import-conventions.md:4` | `go.mod:1` | Fix spec |
| D-CVS-02 | Spec §03 line 88 says path "ends in `core-v8`" — stale | LOW | `spec/01-app/03-import-conventions.md:88` | `go.mod:1` | Fix spec |
| D-CVS-03 | Spec §03 line 94 prose/example mismatch (v8 vs v9) | LOW | `spec/01-app/03-import-conventions.md:94-98` | n/a (internal consistency) | Fix spec |
| D-CVS-04 | Spec §03 line 121 conflates "test module" with "core module" | MED | `spec/01-app/03-import-conventions.md:121` | n/a | Fix spec — generic wording |
| D-CVS-05 | `coregeneric` canonical-import listing not annotated as optional | LOW | `spec/01-app/03-import-conventions.md:61` | (0 importing files) | Fix spec |
| C-CVS-01 | Spec §03 line 129 references nonexistent `tests/integratedtests/` directory | **HIGH** | `spec/01-app/03-import-conventions.md:129` | `tests/` (only `creationtests/` exists) | Decide: fix spec OR restructure tests |
| C-CVS-02 | Spec §03 line 118 `internal/reflectinternal` example doesn't apply to this repo | MED | `spec/01-app/03-import-conventions.md:118` | (0 importing files) | Fix spec — move to "see core-v9 source" reference |

## Resolved drift findings

| ID | Title | Resolved at | Fix location | Path taken |
|----|-------|-------------|--------------|------------|
| — | None yet | — | — | — |

## Targets

| Milestone | Score | Status |
|-----------|-------|--------|
| ✅ First measured baseline (Cycle 1) | **41.7** | 2026-05-04 |
| 🚧 Apply 5 LOW + 2 MED/HIGH spec fixes from Cycle 1 → re-measure | ~100% on §03 | Pending |
| 🎯 Audit all 16 sections of `01-app/` | 16/16 | 1/16 done |
| 🎯 Reach ≥95% aggregate match rate | ≥ 95 | Pending |
| 🎯 Zero ❌ contradictions | 0 (currently 2) | Pending |
