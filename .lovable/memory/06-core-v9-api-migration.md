---
name: core-v9 API migration mapping
description: Maps old core-v8 function-level API to core-v9 struct-namespace API discovered via user's PowerShell investigation
type: feature
---

# core-v9 API Migration Mapping

> Discovered 2026-05-06 by user running PowerShell `Select-String` commands against their local `core-v9` checkout (tag v1.5.8).

## Breaking change pattern

`core-v9` moved from **package-level functions** to **struct-based namespaces** exposed as package-level vars in `converters/vars.go`:

```go
var (
    StringsTo     = stringsTo{}
    AnyTo         = anyItemConverter{}
    Map           = convertinternal.Map
    StringTo      = stringTo{}
    PrettyJson    = jsoninternal.Pretty
    JsonString    = jsoninternal.String
    BytesTo       = bytesTo{}
    Integers      = convertinternal.Integers
    KeyValuesTo   = convertinternal.KeyValuesTo
    CodeFormatter = convertinternal.CodeFormatter
)
```

## Confirmed migration map

| Old call (enum-v3 code today) | New call (core-v9 v1.5.8) | Verified? |
|---|---|---|
| `converters.AnyToValueString(x)` | `converters.AnyTo.ValueString(x)` | ✅ User confirmed method exists |
| `converters.StringToInteger(s)` | `converters.StringTo.Integer(s)` | ⏳ Awaiting `stringTo` method list |
| `converters.StringToIntegerWithDefault(s, d)` | `converters.StringTo.IntegerWithDefault(s, d)` | ⏳ Awaiting `stringTo` method list |
| `converters.StringToByte(s)` | `converters.StringTo.Byte(s)` | ⏳ Awaiting `stringTo` method list |
| `converters.StringToIntegerDefault(s)` | `converters.StringTo.IntegerDefault(s)` | ⏳ Awaiting `stringTo` method list |
| `converters.Any.ToFullNameValueString(x)` | `converters.AnyTo.ToFullNameValueString(x)` | ⏳ Need confirmation |
| `coredynamic.TypeName(x)` | `coredynamic.SafeTypeName(x)` | ✅ User confirmed — no public `TypeName` exists, only `SafeTypeName` |

## `anyItemConverter` methods (confirmed from user output)

```
ToString(anyItem any) string
String(anyItem any) string
FullString(anyItem any) string
ValueString(anyItem any) string
Bytes(anyItem any) []byte
```

(Plus ~20 more methods; partial list from `Select-String` output.)

## `coredynamic` public functions (confirmed)

- `EmptyAnyCollection()`, `NewAnyCollection(capacity int)`
- `SafeTypeName(item any) string` — **this is the replacement for the missing `TypeName`**
- No public `TypeName` function exists in `coredata/coredynamic/`

## `stringTo` methods — PENDING

User was asked to run:
```powershell
Select-String -Path "converters/*.go" -Pattern "func \(.*stringTo\)" | Select-Object -First 20
```
Response not yet received at session end. This is the next piece needed to complete the mapping.

## Where broken calls live in enum-v3

Run this to find all affected files:
```bash
rg -n 'converters\.(AnyToValueString|StringToInteger|StringToByte|StringToIntegerWithDefault|StringToIntegerDefault)' --type go
rg -n 'coredynamic\.TypeName[^s]' --type go
rg -n 'converters\.Any\.' --type go
```

## Next steps

1. Get `stringTo` method list from user (pending command output)
2. Search enum-v3 codebase for all broken calls
3. Create a new task (AM?) to apply the migration fixes
4. Bump version after changes
