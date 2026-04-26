# Fix: corepayload Attributes.Clone(deep=true) error on nil AnyKeyValuePairs

## Root Cause
`Attributes.deepClonePtr()` unconditionally calls `it.AnyKeyValuePairs.ClonePtr()`.
When `AnyKeyValuePairs` is nil (the common case for simple payloads),
`MapAnyItems.ClonePtr()` returns `nil, defaulterr.NilResult` — a non-nil error.
This error propagates up, causing `Clone(deep=true)` to fail with a
misleading "nil result" error even though the nil state is perfectly valid.

## Fix
Added a nil guard around `AnyKeyValuePairs` before calling `ClonePtr()`:
- If nil: skip clone, pass nil through to the new Attributes
- If non-nil: clone normally, propagate real errors

## Learning
- Nil-returning clone methods that signal nil via error (rather than returning `nil, nil`)
  create hidden coupling — callers must know which nil pattern each ClonePtr uses.
- Always nil-guard optional fields before delegating to their clone methods.

## What Not To Repeat
- Don't assume all ClonePtr methods treat nil receivers as `nil, nil`.
  Some return sentinel errors (e.g., `defaulterr.NilResult`).
