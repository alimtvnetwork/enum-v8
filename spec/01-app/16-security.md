# 16 — Security

> ✅ **Status**: drafted at spec-v0.16.0 (2026-04-25, Asia/Kuala_Lumpur).
> **Audience**: anyone implementing security-sensitive consumers, or extending `core-v9` itself.
> **Source**: unifies guidance previously scattered across `04-error-system.md` (panic policy), `00-llm-integration-guide.md` (allocation rules), `02-design-philosophy.md` (no-side-effect rule).
> **Closes**: F-V14-05 *(security half)*.

---

## Table of Contents

1. [Scope — Threat Model](#1-scope--threat-model)
2. [PII & Secret Handling](#2-pii--secret-handling)
3. [Panic Policy](#3-panic-policy)
4. [Allocation & Resource Safety](#4-allocation--resource-safety)
5. [Reflection & Dynamic Dispatch Safety](#5-reflection--dynamic-dispatch-safety)
6. [Input Validation Boundary](#6-input-validation-boundary)
7. [Common Mistakes](#7-common-mistakes)

---

## 1. Scope — Threat Model

`core-v9` is a library. It does **not** itself:

- Parse network input.
- Open files, sockets, or external connections.
- Execute external commands.
- Hold long-lived state.

Its security surface is therefore narrow and consists of:

| Surface | Concern |
|---|---|
| Diagnostic output (`errcore`, `coretests`) | Leaking PII / secrets in logs |
| Panic helpers (`errcore.HandleErr`, `*Must` methods) | Crashing consumer processes |
| Reflection (`coredynamic`, `reflectcore`) | Bypassing type safety, accessing unexported fields |
| Generic containers (`Collection`, `Hashmap`, `Hashset`) | Unbounded allocation under attacker-controlled input |
| Validators (`corevalidator`) | False positives / negatives at trust boundaries |

Everything else (transport security, authn, authz, rate limiting) is the consumer's responsibility.

---

## 2. PII & Secret Handling

The library has **no PII awareness**. If you pass a password, API key, JWT, or PII into one of these helpers, it will appear verbatim in the diagnostic output:

| Helper | What it emits |
|---|---|
| `errcore.VarTwo("password", pwd, …)` | `(password [t:string], …) = (hunter2, …)` |
| `errcore.MessageVarMap(..., map{"token": jwt})` | `... token=eyJhbGciOi...` |
| `corejson.NewPtr(struct).PrettyJsonString()` | Full JSON dump including any secret fields |
| `coredynamic.AllFields(target)` | All exported fields, secrets included |

### Rules

1. **MUST NOT** pass secrets, tokens, passwords, or unredacted PII directly to any `errcore.Var*` or `MessageVarMap` helper.
2. **SHOULD** pass an opaque identifier (`userID`, `tokenID`) and let the secure store resolve it elsewhere.
3. **MUST** scrub or mask sensitive fields before serialising structs through `corejson` for logging:
   ```go
   safe := *original                  // shallow copy
   safe.PasswordHash = ""             // explicit zero
   safe.APIKey       = "<redacted>"   // explicit marker
   logEntry := corejson.NewPtr(&safe).PrettyJsonString()
   ```
4. **MUST** treat any `error` returned from this library as **potentially containing the values that were passed in**. Do not echo errors verbatim to untrusted clients — log internally, return a generic message externally.

---

## 3. Panic Policy

| Code path | Panic allowed? |
|---|---|
| Public functions returning `(T, error)` | ❌ Never. Return the error. |
| `*Must` variants (e.g. `ResultMust`, `DeserializeMust`) | ✅ Yes — that is their entire purpose. Must use `errcore.HandleErr(err)`, not bare `panic(err)`. |
| Init-time validation (`init()` functions) | ✅ Yes — failure here means the binary cannot start. Use `errcore.HandleErr` for the stack-enhanced trace. |
| Internal helpers with documented preconditions | ⚠️ Only when the precondition is impossible to violate from outside the package. Prefer returning an error. |
| `internal/` packages | ⚠️ Same as above — internal callers are framework-controlled, but still prefer errors. |

### Rules

1. **MUST** use `errcore.HandleErr(err)` — never bare `panic(err)`. See [`04-error-system.md` §1](./04-error-system.md) for the contract.
2. **MUST** name `*Must` methods with the exact `Must` suffix at the end of the canonical 8-slot suffix order *(see `00-llm-integration-guide.md` §Pattern 7)*. This makes panic-risk callers visible at the call site.
3. **MUST NOT** recover from panics inside `core-v9`. Recovery is the consumer's decision (e.g. an HTTP middleware).
4. **SHOULD** prefer the error-returning variant in library code; reserve `*Must` for tests, init, and CLI entrypoints in consuming applications.

---

## 4. Allocation & Resource Safety

The library exposes generic containers with **no built-in size limits**. An attacker-controlled input that drives one of these can exhaust memory.

| Container | Risk |
|---|---|
| `coregeneric.Collection[T]` | Unbounded `Adds` from input |
| `coregeneric.SimpleSlice[T]` | Same |
| `coregeneric.LinkedList[T]` | Same, plus per-node allocation overhead |
| `coregeneric.Hashmap[K,V]` | Unbounded `Set` from attacker-controlled keys (also: hash-collision DoS in extreme cases) |
| `coregeneric.Hashset[T]` | Same |
| `corestr.StringBuilder` | Unbounded concatenation |

### Rules

1. **MUST** validate length / cardinality of attacker-controlled input *before* feeding it into a container. Use `corevalidator.New.Slice.MaxLength(N)` upstream.
2. **SHOULD** prefer `SimpleSlice[T]` over `Collection[T]` when the lifetime is single-function-local and the mutex would never be contested *(see `00-llm-integration-guide.md` §Decision matrix)*.
3. **MUST NOT** rely on Go's runtime to "eventually" free containers — explicit `Clear()` calls are required when reusing a long-lived container in a hot path.
4. **MUST NOT** call `coredynamic.AllFields` on a deeply nested struct in a hot path — it builds a fresh `map[string]any` per call.

---

## 5. Reflection & Dynamic Dispatch Safety

`coredynamic` and `reflectcore` wrap stdlib `reflect` to make panics impossible *from the consumer's side*, but they do not prevent a malicious caller from probing internals.

### Rules

1. **MUST NOT** import `internal/reflectinternal` from consumer code *(see `03-import-conventions.md` §3 and `10-reflection-and-dynamic.md` §1)*. That package can mutate unexported fields and is reserved for `enumimpl` and the JSON pipeline.
2. **MUST NOT** call `coredynamic.SetField` on a value supplied by an untrusted caller — the field name is attacker-controlled and could overwrite invariants.
3. **MUST** validate the method name against an allow-list before calling `coredynamic.InvokeMethod` on a value reachable from network input.
4. **SHOULD** prefer compile-time generics over reflection wherever the type is statically known. Reflection is the escape hatch, not the default.

---

## 6. Input Validation Boundary

Trust boundaries are crossed where untrusted bytes become typed Go values. `core-v9` provides validators specifically for this:

```go
v := corevalidator.New.Line.
    NotEmpty().
    MaxLength(255).
    Matches(`^[a-zA-Z0-9_.-]+$`).
    Build()

result := v.Validate(userInput)
if result.IsFailed() {
    return errcore.InvalidInput.MergeError(result.Error())
}
```

### Rules

1. **MUST** validate every untrusted string with at least `NotEmpty()` + `MaxLength(N)` before passing it to a downstream helper. Unbounded strings can DoS converters and JSON serializers.
2. **MUST** use `corevalidator.New.Slice.MaxLength(N)` for slice inputs *before* iteration.
3. **MUST NOT** rely on Go's `string` being valid UTF-8 — an attacker may submit raw bytes through `[]byte` → `string` conversion. Use `corestr.IsValidUTF8` if UTF-8 is a precondition.
4. **SHOULD** centralise validation in a single `validate*` function per request type, not inline at every call site, so security review can audit one location.

---

## 7. Common Mistakes

| Mistake | Fix |
|---|---|
| `errcore.VarTwo("apiKey", key, …)` in production logs | Pass an opaque identifier; resolve the secret elsewhere. |
| `coredynamic.SetField(target, userInput, value)` | Validate `userInput` against an allow-list first. |
| Returning `errcore.<Cat>.MergeError(err)` directly to an HTTP client | Log internally, return a generic 4xx/5xx body to the client. |
| `Hashmap.Set` in a loop with no upstream length validation | Add `corevalidator.New.Slice.MaxLength(N)` before the loop. |
| Bare `panic(err)` inside a `*Must` method | Use `errcore.HandleErr(err)` — see `04-error-system.md` §1. |
| Recovering from a panic inside `core-v9` | Recovery is the consumer's job (e.g. middleware). |

---

## See Also

- [`04-error-system.md`](./04-error-system.md) — `errcore.HandleErr` and the `*Must` panic helper contract.
- [`15-observability.md`](./15-observability.md) — Companion file: diagnostic primitives, logging boundaries.
- [`08-validators.md`](./08-validators.md) — `corevalidator` API for trust-boundary checks.
- [`10-reflection-and-dynamic.md`](./10-reflection-and-dynamic.md) — Reflection-layer safety rules.
- [`02-design-philosophy.md`](./02-design-philosophy.md) — No-side-effect rule, struct-as-namespace boundary.
