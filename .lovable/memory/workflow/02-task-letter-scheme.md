# Task Letter Scheme

## How letter IDs work

- Tasks get a single letter (A, B, C, …) when first proposed.
- Letters are **stable across sessions and "next" turns**.
- Letters are NEVER reused — once a letter is retired (task done, withdrawn, or merged into another), the letter is dead.
- After Z, continue with AA, AB, …, AZ, then BA, BB, …

## Current letter map (snapshot)

| Letter | Task | Status |
|---|---|---|
| A | Manual `cross-repo/core-v8/` push | ⏭️ Manual |
| W | Upstream `core-v9` `go.mod` rename + `v1.5.8` tag | ✅ Done |
| AA | Spec-audit cycles (next: Cycle 15 `spec/06-testing-guidelines/`) | 🔄 In Progress |
| AB | Fetch upstream `core-v9` source for ❓ resolution | ⏳ Pending |
| AC | Re-audit §07/§09 against spec-internal consistency | ⏳ Pending |
| AG | Drop `replace` bridge after W | ✅ Done |
| AH | Cross-`spec/` cleanup sweep | 🔄 In Progress |
| AI | Mark `spec/01-app/` frozen in CHANGELOG | ⏳ Pending |
| AJ | Implement spec fixes from Cycle 15 findings | 📋 Planned |
| AK | New enum package creation (template validation) | 📋 Planned |
| AL | Test coverage expansion | 📋 Planned |
| AM | Fix broken `core-v9` API calls (converter + coredynamic migration) | 🔄 In Progress — patched, awaiting local validation |

(Letters B–V, X, Y, Z, AD, AE, AF — earlier tasks that were either completed across cycles 1–14 or merged into the surviving letters above. Per "never delete history" rule: those completions are recorded in `.lovable/plan.md` Completed section and in the per-cycle audit reports.)

## End-of-reply requirement

Every "next"-loop reply MUST end with the remaining-task list:
- Letter ID
- One-line description
- Current status marker
- Closing line: "Say **next** for **<recommended letter>**, or pick a letter."
- If no obvious next task is active, read `.lovable/plan.md` and this workflow memory, then suggest the highest-priority remaining task.

Skipping this list breaks the user's mental tracking and is a hard violation.
