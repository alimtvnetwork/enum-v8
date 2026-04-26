# Parallel-Sync Mechanism Spec (Deterministic Parallel Coverage Output)

## Purpose
Run package compile/tests in parallel for speed while **always displaying results in a stable order**.

## Problem
Parallel jobs finish nondeterministically, so direct streaming output causes:
- jumbled package rows
- inconsistent blocked-package order
- confusing run-to-run diffs for humans and AI tools

## Core Pattern
Use a two-phase model:
1. **Plan phase (ordered queue):** build a pre-sorted package array with stable index.
2. **Execute phase (parallel):** run jobs async.
3. **Sync phase (deterministic emit):** map each async result back to its planned slot, then render by index.

## Data Contract
Each package work item must include:

```text
{
  Index: int          // stable order source of truth
  Package: string     // full package name
  ShortName: string   // display package name
}
```

Each async result must include:

```text
{
  Index: int
  Package: string
  ExitCode: int
  Output: []string
  ProfilePath: string // test/coverage phase only
}
```

## Algorithm

### A) Compile Check (Parallel + Ordered Display)
1. `allPackages = go list ... | sort`
2. `plan = allPackages.map((pkg, idx) => {Index, Package, ShortName})`
3. Run `plan` with parallel workers.
4. Collect result objects (do not print mid-job).
5. `ordered = results.sortBy(Index)`
6. Emit compile rows in `ordered` sequence.
7. Build blocked list from `ordered` failures.
8. Write `blocked-packages.txt` using same ordered list.

### B) Coverage Run (Parallel + Ordered Display)
1. Build `runPlan` from compilable packages in prior ordered sequence.
2. Run tests in parallel; each job writes its own partial profile.
3. Collect result objects.
4. Sort by `Index`.
5. Emit status table (`[i/N]`, icon, package, percent) in order.
6. Merge coverage profiles after all jobs complete.

## Ordering Rules
- Primary: `Index` from plan phase.
- Secondary fallback: `Package` lexical sort if index missing.
- Never display by completion order.

## File Naming Rules (Parallel Safe)
- Compile outputs and coverage partials must use package-derived safe names.
- Sanitize package name: replace non `[a-zA-Z0-9.-]` with `_`.
- Avoid shared filenames (`compile-1`, `cover-1`) across jobs.

## Error Handling
- A failed package must still produce a result object (`ExitCode != 0`).
- If output has no `.go:line` matches, preserve full stderr as fallback.
- If all compile jobs fail, stop coverage run and print explicit abort reason.

## Acceptance Criteria
- Same input package list => same output row order across repeated runs.
- `blocked-packages.txt` package order is deterministic.
- Parallel mode and sync mode produce equivalent package ordering.
- No file collision errors for parallel compile outputs.

## AI Implementation Checklist
- [ ] Build explicit `plan[]` with stable index.
- [ ] Pass `Index` through async jobs.
- [ ] Never print directly from worker threads.
- [ ] Sort collected results before any display/file write.
- [ ] Use unique, sanitized artifact filenames.
- [ ] Keep diagnostics complete for failed packages.

## Non-Goals
- Changing coverage math.
- Hiding failing packages.
- Reordering output by runtime completion speed.