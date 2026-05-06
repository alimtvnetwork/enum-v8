# `spec/02-app-issues/` — Forward-Looking Issues Catalog

> **Status**: ✅ All 9 issues resolved (last update spec-v0.8.0). Originally scaffolded in Step 6 of the audit plan.
> **Source**: [`spec/99-audits/01-original-11-step-plan.md` §7.2](../99-audits/01-original-11-step-plan.md#72-proposed-spec02-app-issues-skeleton)
> **Date scaffolded**: 2026-04-23 (Asia/Kuala_Lumpur, UTC+8)

This folder is the **forward-looking** issues catalog: known limitations, planned migrations, and technical debt.

> **Contrast with `spec/05-failing-tests/`** — that folder is the **backward-looking** fix log
> (one file per past test failure with root-cause + resolution). This folder is the **forward-looking**
> roadmap of *open* questions the team has not yet resolved.

## Index

| # | Topic | Status | Severity |
|---|---|---|---|
| 01 | [Style B / Style A coexistence](./01-style-b-style-a-coexistence.md) | ✅ resolved | medium |
| 02 | [Internal-package coverage policy](./02-internal-package-coverage-policy.md) | ✅ resolved | low |
| 03 | [`GetAssert` undocumented API](./03-getassert-undocumented-api.md) | ✅ resolved | low |
| 04 | [`testwrappers` public surface](./04-testwrappers-public-surface.md) | ✅ resolved | low |
| 05 | [Missing `params.go` files](./05-missing-params-go-files.md) | ✅ resolved | low |
| 06 | [Validator error canonical example](./06-validator-error-canonical-example.md) | ✅ resolved | low |
| 07 | [`newCreator.go` filename casing](./07-newcreator-filename-casing.md) | ✅ resolved | low |
| 08 | [`errcore` type-boundary examples](./08-errcore-type-boundary-examples.md) | ✅ resolved | low |
| 09 | [Enum predicate file-split rule](./09-enum-predicate-file-split.md) | ✅ resolved | low |

See [`00-issues-index.md`](./00-issues-index.md) for the canonical machine-readable index.
