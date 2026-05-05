# `spec/01-app/` — Per-App Architectural Documentation

> **Status**: 🚧 Skeleton — created in Step 2 of the audit plan.
> **Source**: [`spec/99-audits/01-original-11-step-plan.md` §7.1](../99-audits/01-original-11-step-plan.md#71-proposed-spec01-app-skeleton-to-be-created-in-a-later-step)
> **Date scaffolded**: 2026-04-23 (Asia/Kuala_Lumpur, UTC+8)

This folder contains the **authoritative architectural reference** for the `core-v9` Go module, split into atomic per-topic pages. Most content will be extracted from `spec/00-llm-integration-guide.md` in subsequent steps (3, 4, 5).

## Suggested AI Reading Order

1. [`00-repo-overview.md`](./00-repo-overview.md) — module identity, top-level layout
2. [`01-package-map.md`](./01-package-map.md) — every package, one-line purpose
3. [`02-design-philosophy.md`](./02-design-philosophy.md) — struct-as-namespace, one-file-per-function
4. [`03-import-conventions.md`](./03-import-conventions.md) — public vs `internal/`
5. [`04-error-system.md`](./04-error-system.md) — `errcore` + `errcoreinf`
6. [`05-enum-system.md`](./05-enum-system.md) — `enuminf` + `enumimpl`
7. [`06-data-structures.md`](./06-data-structures.md) — `coredata` umbrella
8. [`07-conditional-and-utilities.md`](./07-conditional-and-utilities.md)
9. [`08-validators.md`](./08-validators.md)
10. [`09-converters.md`](./09-converters.md)
11. [`10-reflection-and-dynamic.md`](./10-reflection-and-dynamic.md)
12. [`11-versioning.md`](./11-versioning.md)
13. [`12-cmd-entrypoints.md`](./12-cmd-entrypoints.md)
14. [`13-testing-patterns.md`](./13-testing-patterns.md) — forwards to `spec/06-testing-guidelines/`
15. [`14-tests-folder-walkthrough.md`](./14-tests-folder-walkthrough.md) — `tests/integratedtests/` + `tests/testwrappers/`
16. [`15-observability.md`](./15-observability.md) — Diagnostic primitives, logging boundaries, tracing scope *(spec-v0.16.0)*
17. [`16-security.md`](./16-security.md) — PII/secret handling, panic policy, allocation safety *(spec-v0.16.0)*

## Status Legend

- 🚧 **Skeleton** — file exists, content is a placeholder
- 📝 **Draft** — content present, awaiting review
- ✅ **Done** — content reviewed and stable

## Per-File Status

| File | Status | Filled in |
|---|---|---|
| `00-repo-overview.md` | ✅ Done | Step 3 (reviewed 2026-04-23) |
| `01-package-map.md` | ✅ Done | Step 3 (reviewed 2026-04-23) |
| `02-design-philosophy.md` | ✅ Done | Step 3 + Followup #1 filename-casing rule (reviewed 2026-04-23) |
| `03-import-conventions.md` | ✅ Done | Step 5a (reviewed 2026-04-23) |
| `04-error-system.md` | ✅ Done | Step 5a (reviewed 2026-04-23) |
| `05-enum-system.md` | ✅ Done | Step 5a (reviewed 2026-04-23) |
| `06-data-structures.md` | ✅ Done | Step 5b (reviewed 2026-04-23) |
| `07-conditional-and-utilities.md` | ✅ Done | Step 5b (reviewed 2026-04-23) |
| `08-validators.md` | ✅ Done | Step 5b + Followup #1 canonical error format (reviewed 2026-04-23) |
| `09-converters.md` | ✅ Done | Step 5c (reviewed 2026-04-23) |
| `10-reflection-and-dynamic.md` | ✅ Done | Step 5c (reviewed 2026-04-23) |
| `11-versioning.md` | ✅ Done | Step 5c (reviewed 2026-04-23) |
| `12-cmd-entrypoints.md` | ✅ Done | Step 5c (reviewed 2026-04-23) |
| `13-testing-patterns.md` | ✅ Done | Step 4 (reviewed 2026-04-23) |
| `14-tests-folder-walkthrough.md` | ✅ Done | Step 4 (reviewed 2026-04-23) |
| `15-observability.md` | ✅ Done | spec-v0.16.0 (2026-04-25) — closes F-V14-05 observability half |
| `16-security.md` | ✅ Done | spec-v0.16.0 (2026-04-25) — closes F-V14-05 security half |

> **Review pass complete (Followup #3, 2026-04-23)**: all 14 files validated against the Step 11 fresh-AI simulation (96.25% reproducibility) and promoted to ✅ Done. Total content: ~3,400 lines across 14 atomic topic files.
>
> **spec-v0.16.0 expansion (2026-04-25)**: 2 new files added (`15-observability.md`, `16-security.md`) to close F-V14-05. Total content now ~3,750 lines across 16 atomic topic files.
