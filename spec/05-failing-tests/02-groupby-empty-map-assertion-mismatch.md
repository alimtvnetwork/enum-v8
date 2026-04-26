# 01 — GroupBy Empty Map Assertion Mismatch

## Test
`Test_Collection_GroupBy_Verification` — Case 1: "GroupBy returns empty map -- empty input"

## Root Cause
The test case used `ExpectedInput: ""` (a string), which routes through the `else` branch
calling `ShouldBeEqual(t, caseIndex, actLines...)`. When `groups` is empty, `actLines` is
an empty slice, so the variadic expands to **0 arguments**. However, the assertion framework
interprets `ExpectedInput: ""` as expecting **1 line** (the empty string), causing a length
mismatch: `Expect [""] != [""] Actual`.

## Solution
Changed `ExpectedInput` from `""` to `args.Map{}`. This routes the empty case through the
`if _, isMap := ...` branch, which builds an empty `args.Map{}` actual and compares it against
an empty `args.Map{}` expected — both empty maps match.

## Learnings
- `ShouldBeEqual` with a string `ExpectedInput` always expects at least one line.
- For zero-output cases, use `args.Map{}` to go through the map-comparison path instead.
- The `GroupByCount` empty case (`ExpectedInput: "0"`) works because it explicitly formats
  `len(counts)` as a string argument (`fmt.Sprintf("%d", len(counts))`), always passing 1 line.

## What Not to Repeat
- Do not use empty string `""` as `ExpectedInput` to represent "no output" — it creates
  a single expected line. Use `args.Map{}` for truly empty expectations.
