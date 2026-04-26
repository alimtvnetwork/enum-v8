# Fix: codestack/newCreator.go nil pointer panic

## Root Cause
`Create(skipIndex)` calls `runtime.Caller(skipIndex + defaultInternalSkip)`. When the skip depth exceeds the actual call stack, `isOkay` returns `false` and `runtime.FuncForPC(pc)` returns `nil`. The subsequent `funcInfo.Name()` call panics with a nil pointer dereference.

This panic crashes the **entire test binary**, causing 0/0 coverage stmts for codestack AND potentially other packages (bytetype, corecmp) if they share the same test run or if the runner aborts.

## Fix
Added nil guards: return an empty `Trace{IsOkay: false}` when `runtime.Caller` fails, and check `funcInfo != nil` before calling `.Name()`.
