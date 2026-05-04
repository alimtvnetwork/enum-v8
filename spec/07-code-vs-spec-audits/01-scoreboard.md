# 01 — Code-vs-Spec Drift Scoreboard (Living Document)

> **Single source of truth** for code-vs-spec drift. Updated after every cycle.

## Current MEASURED drift score: **83.3 / 100** *(Cycle 1 post-fix, n=12 claims, n=1 spec section)*

> **Caveat:** sample size is one spec section. Cycle 1 baseline was 41.7%. After applying the 5 LOW drift fixes (D-CVS-01..05) on 2026-05-04, projected match for §03 is 10/12 = **83.3%**. C-CVS-01 (the missing `tests/integratedtests/` folder question) and C-CVS-02 (`internal/reflectinternal` example) remain open; resolving both lifts §03 to 100%.

## Cycle history

| Date | Cycle | Spec audited | Claims | ✅ Match | ⚠️ Drift | ❌ Contradiction | Score |
|------|-------|--------------|--------|---------|---------|------------------|-------|
| 2026-05-04 | 1 (baseline) | `01-app/03-import-conventions.md` | 12 | 5 | 5 | 2 | **41.7%** |
| 2026-05-04 | 1 (post-fix) | `01-app/03-import-conventions.md` | 12 | 10 | 0 | 2 | **83.3%** |

## Open drift findings

| ID | Title | Severity | Spec ref | Code ref | Resolution path |
|----|-------|----------|----------|----------|-----------------|
| C-CVS-01 | Spec §03 line 129 references nonexistent `tests/integratedtests/` directory | **HIGH** | `spec/01-app/03-import-conventions.md:129` | `tests/` (only `creationtests/` exists) | Decide: fix spec OR restructure tests |
| C-CVS-02 | Spec §03 line 118 `internal/reflectinternal` example doesn't apply to this repo | MED | `spec/01-app/03-import-conventions.md:118` | (0 importing files) | Fix spec — move to "see core-v9 source" reference |

## Resolved drift findings

| ID | Title | Resolved at | Fix location | Path taken |
|----|-------|-------------|--------------|------------|
| D-CVS-01 | Spec §03 line 4 says "consumes `core-v8`" — stale | 2026-05-04 | `spec/01-app/03-import-conventions.md:4` | s/core-v8/core-v9/ |
| D-CVS-02 | Spec §03 line 88 says path "ends in `core-v8`" — stale | 2026-05-04 | `spec/01-app/03-import-conventions.md:88` | s/core-v8/core-v9/ + s/corev8/corev9/ |
| D-CVS-03 | Spec §03 line 94 prose/example mismatch (v8 vs v9) | 2026-05-04 | `spec/01-app/03-import-conventions.md:94` | s/core-v8/core-v9/ |
| D-CVS-04 | Spec §03 line 121 conflates "test module" with "core module" | 2026-05-04 | `spec/01-app/03-import-conventions.md:121` | Reworded to be module-generic |
| D-CVS-05 | `coregeneric` canonical-import listing not annotated as optional | 2026-05-04 | `spec/01-app/03-import-conventions.md:61,73` | Inline `// optional` comment + consumer-coverage note |

## Targets

| Milestone | Score | Status |
|-----------|-------|--------|
| ✅ First measured baseline (Cycle 1) | **41.7** | 2026-05-04 |
| ✅ Apply 5 LOW spec fixes from Cycle 1 (D-CVS-01..05) | **83.3** on §03 | 2026-05-04 |
| 🚧 Resolve C-CVS-01 + C-CVS-02 → §03 at 100% | 100% on §03 | Pending |
| 🎯 Audit all 16 sections of `01-app/` | 16/16 | 1/16 done |
| 🎯 Reach ≥95% aggregate match rate | ≥ 95 | Pending |
| 🎯 Zero ❌ contradictions | 0 (currently 2) | Pending |
