# Write Memory

> **Purpose:** After completing work or at the end of a session, the AI must persist everything it learned, did, and left undone — so the next AI session can pick up seamlessly with zero context loss.
>
> **Trigger phrases:** "Write memory", "End memory", "update memory".
>
> **When to run:** At the end of every session, after completing a task batch, or when explicitly asked.

---

## Table of Contents

1. [Core Principle](#core-principle)
2. [Phase 1 — Audit Current State](#phase-1--audit-current-state)
3. [Phase 2 — Update Memory Files](#phase-2--update-memory-files)
4. [Phase 3 — Update Plans & Suggestions](#phase-3--update-plans--suggestions)
5. [Phase 4 — Update Issues](#phase-4--update-issues)
6. [Phase 5 — Consistency Validation](#phase-5--consistency-validation)
7. [File Naming & Structure Rules](#file-naming--structure-rules)
8. [Anti-Corruption Rules](#anti-corruption-rules)
9. Write all Lovable suggestions in the appropriate folder with index for suggestions.
10. Write all CI/CD issues in `.lovable/cicd-issues/xx-issue-name.md` (xx = 01, 02, …) and maintain `.lovable/cicd-index.md` as the summary index. Collect every CI/CD issue we have so far. Do not duplicate.
11. Review current discussions; if a bigger spec was given to write to file system, also put a verbatim copy in memory so the next AI can understand it.

---

## Core Principle

> **The memory system is the project's brain.** If you did something and didn't write it down, it didn't happen. If something is pending and you didn't record it, it will be lost. Write memory as if the next AI has amnesia — because it does.

---

## Phase 1 — Audit Current State

Before writing anything, take inventory:

### What was done this session?
- [ ] List every task completed (features, fixes, refactors)
- [ ] List every file created, modified, or deleted
- [ ] List every decision made and why

### What is still pending?
- [ ] Tasks started but not finished
- [ ] Tasks discussed but not started
- [ ] Blockers or dependencies that prevented completion

### What was learned?
- [ ] New patterns or conventions discovered
- [ ] Gotchas or edge cases encountered
- [ ] User preferences expressed (explicitly or implicitly)

### What went wrong?
- [ ] Bugs encountered and their root causes
- [ ] Approaches that failed and why
- [ ] Things that should never be repeated

---

## Phase 2 — Update Memory Files

### Target: `.lovable/memory/`

#### Step 2.1 — Read the current index: `.lovable/memory/index.md`. Avoid duplicates.

#### Step 2.2 — Update existing memory files
For each file affected by this session: open, append, mark completed items `[x]` or ✅, never truncate unrelated entries.

#### Step 2.3 — Create new memory files (if needed)
1. Naming: `XX-descriptive-name.md` in `.lovable/memory/`.
2. **Immediately update** `.lovable/memory/index.md`.

#### Step 2.4 — Update workflow state in `.lovable/memory/workflow/`
Statuses:
| Status | Marker |
|--------|--------|
| Done | ✅ Done |
| In Progress | 🔄 In Progress |
| Pending | ⏳ Pending |
| Blocked | 🚫 Blocked — [reason] |
| Avoid or Skip | 🚫 Blocked — [avoid] |

---

## Phase 3 — Update Plans & Suggestions

### 3A — Plans (`.lovable/plan.md`)
- Update task statuses.
- Add new tasks discovered.
- Move fully-complete items to a `## Completed` section in the same file (do NOT delete).

### 3B — Suggestions (`.lovable/suggestions.md`) — single file
Structure:
```markdown
## Active Suggestions
### [Title]
- **Status:** Pending | In Review | Approved | Rejected
- **Priority:** High | Medium | Low
- **Description:** What and why
- **Added:** [date]

## Implemented Suggestions
### [Title]
- **Implemented:** [date]
- **Notes:** Details
```
When implemented: move from Active → Implemented, add notes, reference commit/file/task.

---

## Phase 4 — Update Issues

### 4A — Pending: `.lovable/pending-issues/XX-short-description.md`
```markdown
# [Issue Title]
## Description
## Root Cause
## Steps to Reproduce
## Attempted Solutions
## Priority
## Blocked By
```

### 4B — Solved: `.lovable/solved-issues/XX-…md` (move from pending) and append:
```markdown
## Solution
## Iteration Count
## Learning
## What NOT to Repeat
```

### 4C — Strictly Avoided Patterns: `.lovable/strictly-avoid.md`
```markdown
- **[Pattern]:** [Why forbidden]. See: `.lovable/solved-issues/XX-….md`
```

---

## Phase 5 — Consistency Validation

### 5.1 — Index Integrity
Every file in `.lovable/memory/` (incl. subfolders) must be in `index.md`.

### 5.2 — Cross-Reference Check
- Every `✅ Done` task in `plan.md` has corresponding evidence.
- Every `pending-issues/` item is reflected in `plan.md` or `suggestions.md` if actionable.
- No file in both `pending-issues/` and `solved-issues/`.

### 5.3 — Orphan Check
- No memory file without an index entry.
- No "Implemented" suggestion without codebase evidence.
- No solved issue without a `## Solution`.

### 5.4 — Final Confirmation
```
✅ Memory update complete.
Session Summary:
- Tasks completed: [X]
- Tasks pending: [Y]
- New memory files created: [Z]
- Issues resolved: [N]
- Issues opened: [M]
- Suggestions added: [S]
- Suggestions implemented: [T]
Files modified:
- [list every file touched]
Inconsistencies found and fixed:
- [list any, or "None"]
The next AI session can pick up from: [describe state and next logical step]
```

---

## File Naming & Structure Rules

| Rule | Example |
|------|---------|
| Numeric prefix | `01-auth-flow.md` |
| Lowercase, hyphen-separated | `03-error-handling.md` ✅ / `03_Error_Handling.md` ❌ |
| Plans → single file | `.lovable/plan.md` |
| Suggestions → single file | `.lovable/suggestions.md` |
| Pending issues → one per issue | `.lovable/pending-issues/01-login-crash.md` |
| Solved issues → one per issue | `.lovable/solved-issues/01-login-crash.md` |
| Memory → grouped by topic | `.lovable/memory/workflow/`, `.lovable/memory/decisions/` |
| Completed plans/suggestions → `## Completed` in same file | No separate `completed/` folders |

### Folder Structure
```
.lovable/
├── overview.md
├── strictly-avoid.md
├── user-preferences
├── plan.md
├── suggestions.md
├── prompt.md
├── cicd-index.md
├── memory/
│   ├── index.md
│   ├── workflow/
│   ├── decisions/
│   └── [topic]/
├── pending-issues/
├── solved-issues/
├── cicd-issues/
└── prompts/
```

> ⚠️ **NEVER** create `.lovable/memories/` (with trailing `s`). The correct path is `.lovable/memory/`.

---

## Anti-Corruption Rules

1. **Never delete history.**
2. **Never overwrite blindly** — read before write.
3. **Never leave orphans** — every file indexed, every reference resolves.
4. **Never split what should be unified** — plans/suggestions each live in ONE file.
5. **Never mix states** — pending vs solved are mutually exclusive.
6. **Never skip the index update** — same operation as creating the file.
7. **Never assume the next AI knows anything.**

Any task marked to skip or avoid → put into `.lovable/memory/avoid/` or `.lovable/strictly-avoid.md`.

---

## Important

Save this prompt as **"Write memory"** or **"End memory"**. When invoked, write the specific files and update memory carefully. Do not lose any conversation as much as possible. Restructure folder if it doesn't follow these rules. Write all md files in lowercase-hyphenated.
