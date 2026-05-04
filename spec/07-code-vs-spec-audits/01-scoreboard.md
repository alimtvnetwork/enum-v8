# 01 — Code-vs-Spec Drift Scoreboard (Living Document)

> **Single source of truth** for code-vs-spec drift. Updated after every cycle.

## Current MEASURED drift score: **100.0 / 100** *(Cycle 1 fully closed, n=12 claims, n=1 spec section)*

> Cycle 1 baseline was 41.7%. After D-CVS-01..05 (LOW spec edits) the score moved to 83.3%. C-CVS-01 and C-CVS-02 are now resolved by rewriting §3 "Common `internal/` packages used by tests" and §4 "Test-Package Imports" to match the actual `tests/creationtests/` layout in `enum-v1`, with a cross-reference to the upstream `core-v9` repo's per-suite layout. §03 is now at 12/12.

## Cycle history

| Date | Cycle | Spec audited | Claims | ✅ Match | ⚠️ Drift | ❌ Contradiction | Score |
|------|-------|--------------|--------|---------|---------|------------------|-------|
| 2026-05-04 | 1 (baseline) | `01-app/03-import-conventions.md` | 12 | 5 | 5 | 2 | **41.7%** |
| 2026-05-04 | 1 (post-LOW) | `01-app/03-import-conventions.md` | 12 | 10 | 0 | 2 | **83.3%** |
| 2026-05-04 | 1 (closed) | `01-app/03-import-conventions.md` | 12 | 12 | 0 | 0 | **100.0%** |

## Open drift findings

| ID | Title | Severity | Spec ref | Code ref | Resolution path |
|----|-------|----------|----------|----------|-----------------|
| — | None | — | — | — | — |

## Resolved drift findings

| ID | Title | Resolved at | Fix location | Path taken |
|----|-------|-------------|--------------|------------|
| D-CVS-01 | Spec §03 line 4 says "consumes `core-v8`" — stale | 2026-05-04 | `spec/01-app/03-import-conventions.md:4` | s/core-v8/core-v9/ |
| D-CVS-02 | Spec §03 line 88 says path "ends in `core-v8`" — stale | 2026-05-04 | `spec/01-app/03-import-conventions.md:88` | s/core-v8/core-v9/ + s/corev8/corev9/ |
| D-CVS-03 | Spec §03 line 94 prose/example mismatch (v8 vs v9) | 2026-05-04 | `spec/01-app/03-import-conventions.md:94` | s/core-v8/core-v9/ |
| D-CVS-04 | Spec §03 line 121 conflates "test module" with "core module" | 2026-05-04 | `spec/01-app/03-import-conventions.md:121` | Reworded to be module-generic |
| D-CVS-05 | `coregeneric` canonical-import listing not annotated as optional | 2026-05-04 | `spec/01-app/03-import-conventions.md:61,73` | Inline `// optional` + consumer-coverage note |
| C-CVS-01 | Spec §03 line 129 references nonexistent `tests/integratedtests/` | 2026-05-04 | `spec/01-app/03-import-conventions.md:127-145` | Spec rewritten to `tests/creationtests/` layout (matches code) + cross-ref to upstream `core-v9` per-suite layout |
| C-CVS-02 | Spec §03 line 118 `internal/reflectinternal` example doesn't apply to this repo | 2026-05-04 | `spec/01-app/03-import-conventions.md:113-123` | Section reframed as "in upstream `core-v9` tests"; consumer-side note added explaining cross-module `internal/` is forbidden |

## Targets

| Milestone | Score | Status |
|-----------|-------|--------|
| ✅ First measured baseline (Cycle 1) | **41.7** | 2026-05-04 |
| ✅ Apply 5 LOW spec fixes from Cycle 1 (D-CVS-01..05) | **83.3** on §03 | 2026-05-04 |
| 🚧 Resolve C-CVS-01 + C-CVS-02 → §03 at 100% | 100% on §03 | Pending |
| 🎯 Audit all 16 sections of `01-app/` | 16/16 | 1/16 done |
| 🎯 Reach ≥95% aggregate match rate | ≥ 95 | Pending |
| 🎯 Zero ❌ contradictions | 0 (currently 2) | Pending |
