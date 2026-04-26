# 02 — ReflectSetFromTo DraftToBytes Marshal-then-Unmarshal Bug

## Test
`Test_ReflectSetFromTo_DraftToBytes` — Case 0: "(otherType, *[]byte) -- try marshal, reflect"

## Root Cause (Two Issues)

### Issue A — Production bug: marshal then unmarshal into *[]byte
In `ReflectSetFromTo` (`coredata/coredynamic/ReflectSetFromTo.go`), the documented supported
case `(otherType, *[]byte)` marshals `From` (a struct) into JSON bytes, but then fell through
to `json.Unmarshal(rawBytes, toPointer)`. Since `toPointer` is `*[]byte`, `json.Unmarshal`
expects a **base64-encoded JSON string**, not a JSON object — causing an unmarshal error.

**Fix:** After successful `json.Marshal`, when the destination is `*[]byte`, directly assign
`*toPointer.(*[]byte) = rawBytes` instead of calling `json.Unmarshal`.

### Issue B — Test case type mismatch: ExpectedValue was []byte, To was *[]byte
After fixing Issue A, `err == nil` became `true` but `TypeSameStatus.IsSame` remained `false`.

`TypeSameStatus(tc.To, tc.ExpectedValue)` compares `reflect.TypeOf`:
- `tc.To` = `&[]byte{}` → type `*[]byte`
- `tc.ExpectedValue` = `DraftType.JsonBytesPtr()` → returns `[]byte` (value, NOT pointer)

Despite the misleading method name `JsonBytesPtr`, it returns `[]byte` not `*[]byte`.
So `*[]byte != []byte` → `IsSame = false`.

**Fix:** Wrapped the `ExpectedValue` in a pointer:
```go
ExpectedValue: func() *[]byte {
    b := ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr()
    return &b
}(),
```

## Iteration Details
1. **Attempt 1:** Changed test case `From` from `JsonBytesPtr()` to `&struct` — correctly
   set up the "otherType → *[]byte" scenario but exposed the production bug (both lines false).
2. **Attempt 2:** Fixed production code to short-circuit with direct byte assignment for
   `*[]byte` targets. Line 0 became `true` but line 1 remained `false`.
3. **Attempt 3:** Root-caused the `IsSame` failure to `JsonBytesPtr()` returning `[]byte`
   (not `*[]byte`). Wrapped `ExpectedValue` in a pointer to match `To`'s type.

## Learnings
- `json.Unmarshal` into `*[]byte` only works for base64-encoded JSON strings.
- Method names like `JsonBytesPtr` can be misleading — always verify the return type signature.
- `TypeSameStatus.IsSame` compares `reflect.TypeOf` exactly — `*[]byte != []byte`.

## What Not to Repeat
- Do not assume a method named `*Ptr` returns a pointer — check the signature.
- When setting up test wrappers with `any`-typed fields, verify the concrete types match
  what the assertion logic expects.
