# Suggestions Tracker

> Single file tracking all Lovable suggestions. Update in-place — do not create per-suggestion files.
> When a suggestion is completed, move it to the **Completed** section with date and notes.

---

## Convention

- **suggestionId**: `S-<NNN>` (sequential)
- **status**: `open` | `in-progress` | `done` | `rejected`
- When done: move to Completed section, add completion date and notes
- When rejected: move to Rejected section, add reason

---

## Open Suggestions

### S-001: Pin Go toolchain to 1.22 as stopgap for Task W

- **createdAt:** 2026-05-05
- **source:** Lovable (Cycle 13)
- **affectedProject:** enum-v4
- **description:** Go 1.25 rejects the dual-path `replace` bridge. Pinning to Go 1.22 would unblock builds.
- **rationale:** Allows development to continue while waiting for upstream `core-v9` `go.mod` rename.
- **proposed change:** Add `toolchain go1.22.0` to `go.mod`.
- **acceptance criteria:** `./run.ps1 -tc` passes with Go 1.22 toolchain.
- **status:** open
- **risk:** Masks the underlying issue; locks to older toolchain.

### S-002: Promote `errcore.VarTwoNoType` from ❓ to ✅ in Cycle 6

- **createdAt:** 2026-05-05
- **source:** Lovable (Cycle 13, §15 audit)
- **affectedProject:** enum-v4 spec
- **description:** `VarTwoNoType` was scored ❓ but IS cross-referenced from multiple spec files.
- **rationale:** Under spec-internal-consistency dimension it qualifies as ✅.
- **proposed change:** Backport promotion when Task AC runs.
- **acceptance criteria:** Cycle 6 audit report row 16 updated.
- **status:** open

### S-003: Fix `spec/06-testing-guidelines/01-folder-structure.md` stale path

- **createdAt:** 2026-05-05
- **source:** Lovable (reliability report)
- **affectedProject:** enum-v4 spec
- **description:** Line 13 references `tests/integratedtests/` which doesn't exist. Should be `tests/creationtests/`.
- **rationale:** This is the #1 failure risk for any AI following the spec to write tests.
- **proposed change:** Replace `integratedtests` with `creationtests` throughout the file.
- **acceptance criteria:** `rg integratedtests spec/06-testing-guidelines/01-folder-structure.md` returns 0 hits.
- **status:** open

### S-004: Fix `spec/00-llm-integration-guide.md` stale test path reference

- **createdAt:** 2026-05-05
- **source:** Lovable (reliability report)
- **affectedProject:** enum-v4 spec
- **description:** Line 36 references `tests/integratedtests/` in the decision matrix table.
- **rationale:** First file any AI reads; stale path causes immediate misdirection.
- **proposed change:** Replace stale reference with `tests/creationtests/`.
- **acceptance criteria:** `rg integratedtests spec/00-llm-integration-guide.md` returns only the anti-pattern callout lines.
- **status:** open

---

## Completed Suggestions

### S-100: Add `cmd/main/` smoke-test policy carve-out to spec §12

- **completedAt:** 2026-05-05 (Cycle 10, fix C-CVS-10)
- **notes:** Spec §12 rewritten as "library-first, smoke-test allowed" policy.

### S-101: Rewrite §06 around `SimpleSlice`/`Hashset`/`SimpleStringOnce`

- **completedAt:** 2026-05-04 (Cycle 4, fixes D-CVS-22/23/24)
- **notes:** Replaced fictional `coreonce.New.String` with actual constructors.

### S-102: Add consumer-coverage callouts for upstream-only API

- **completedAt:** 2026-05-04 → 2026-05-05 (D-CVS-25, D-CVS-38, D-CVS-42)
- **notes:** Three sections now carry explicit upstream-only callouts.

---

## Rejected Suggestions

_(none yet)_
