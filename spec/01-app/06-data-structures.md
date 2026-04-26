# 06 — Data Structures

> ✅ **Status**: filled in audit Step 5b (2026-04-23, Asia/Kuala_Lumpur).
> **Audience**: anyone using or extending `coredata/*` packages.
> **Source**: extracted from [`/spec/00-llm-integration-guide.md`](../00-llm-integration-guide.md) §coredata + §coregeneric.

---

## Table of Contents

1. [`coredata` Umbrella](#1-coredata-umbrella)
2. [`coregeneric` — Generic Containers](#2-coregeneric--generic-containers)
3. [`corestr` — String Collections](#3-corestr--string-collections)
4. [`corejson` — JSON Pipeline](#4-corejson--json-pipeline)
5. [`coreonce` — Compute-Once Values](#5-coreonce--compute-once-values)
6. [`corepayload` — `PayloadWrapper`](#6-corepayload--payloadwrapper)
7. [Choosing the Right Container](#7-choosing-the-right-container)

---

## 1. `coredata` Umbrella

`coredata/` is a parent directory, not a package. It groups data-structure sub-packages that share the same design rules:

| Sub-package | Purpose |
|---|---|
| [`coregeneric`](#2-coregeneric--generic-containers) | Generic `Collection[T]`, `Hashset[T]`, `Hashmap[K,V]`, `SimpleSlice[T]`, `LinkedList[T]`, `Pair`, `Triple` |
| [`corestr`](#3-corestr--string-collections) | String-typed collections with thread-safe operations |
| [`corejson`](#4-corejson--json-pipeline) | JSON serialize / deserialize / pretty-print |
| [`coreonce`](#5-coreonce--compute-once-values) | Lazy compute-once cached values |
| [`corepayload`](#6-corepayload--payloadwrapper) | Wire-format payload envelopes |

### Shared rules

1. **Constructors return non-nil values.** An "empty" collection is still a usable receiver — see Zero-Nil Safety in [`02-design-philosophy.md`](./02-design-philosophy.md).
2. **All containers expose a `*Lock` family for thread-safe access.** Pick the lockless variant when you own the only goroutine touching the value.
3. **Mutation methods return the receiver** (fluent style) so calls can chain.
4. **Read methods never mutate, never lock unless `*Lock`.**

---

## 2. `coregeneric` — Generic Containers

Located at `coredata/coregeneric/`. Generic-first, built on Go 1.18+ type parameters.

### 2.1 `Collection[T any]`

Slice-backed collection with embedded `sync.Mutex`.

#### Construction

```go
import "github.com/alimtvnetwork/core-v8/coredata/coregeneric"

// Via the New struct-as-namespace
col := coregeneric.New.Collection.String.Cap(10)
col := coregeneric.New.Collection.Int.Items(1, 2, 3)
col := coregeneric.New.Collection.Float64.Empty()

// Via package-level functions
col := coregeneric.EmptyCollection[string]()
col := coregeneric.NewCollection[MyStruct](20)
col := coregeneric.CollectionFrom(existingSlice)   // no copy
col := coregeneric.CollectionClone(existingSlice)  // copies
```

#### Mutation (returns `*Collection[T]`)

| Method | Description |
|---|---|
| `Add(item)` / `AddLock(item)` | Append one |
| `Adds(items...)` / `AddsLock(items...)` | Append variadic |
| `AddSlice([]T)` | Append from slice |
| `AddIf(bool, item)` / `AddIfMany(bool, items...)` | Conditional append |
| `AddFunc(func() T)` | Append function result |
| `AddCollection(*Collection[T])` / `AddCollections(...*Collection[T])` | Merge |
| `RemoveAt(index) bool` | Remove by index |
| `SortFunc(less)` | In-place sort |
| `Reverse()` | In-place reverse |
| `ConcatNew(items...)` | New collection = this + items |
| `Clone()` | Deep copy |

#### Query

| Method | Returns | Description |
|---|---|---|
| `Length()` / `Count()` / `LengthLock()` | `int` | Size |
| `IsEmpty()` / `IsEmptyLock()` | `bool` | Empty check |
| `HasItems()` / `HasAnyItem()` | `bool` | Non-empty |
| `HasIndex(i)` | `bool` | Bounds check |
| `Items()` | `[]T` | Underlying slice |
| `First()` / `Last()` | `T` | **Panics** if empty |
| `FirstOrDefault()` / `LastOrDefault()` | `T` | Zero-value if empty |
| `SafeAt(i)` | `T` | Zero-value if OOB |
| `Skip(n)` / `Take(n)` | `[]T` | Slice ops |
| `Filter(pred)` | `*Collection[T]` | New filtered |
| `CountFunc(pred)` | `int` | Count matching |
| `ForEach(fn)` / `ForEachBreak(fn)` | — | Iterate |

#### Iterators (Go 1.23+)

```go
for i, item := range col.All()    { ... }   // iter.Seq2[int, T]
for item    := range col.Values() { ... }   // iter.Seq[T]
```

#### Package-level generic helpers

| Function | Constraint | Description |
|---|---|---|
| `MapCollection(src, fn)` | `T→U` | Transform `Collection[T]` → `Collection[U]` |
| `FlatMapCollection(src, fn)` | `T→[]U` | Flatten-transform |
| `ReduceCollection(src, init, fn)` | `T→U` | Fold |
| `GroupByCollection(src, keyFn)` | `K comparable` | `map[K]*Collection[T]` |
| `ContainsFunc` / `ContainsItem` | predicate / `comparable` | Search |
| `IndexOfFunc` / `IndexOfItem` | — | Find index |
| `Distinct(src)` | `comparable` | Deduplicate |
| `ContainsAll` / `ContainsAny` | `comparable` | Bulk membership |
| `RemoveItem` / `RemoveAllItems` | `comparable` | Remove |
| `ToHashset(src)` | `comparable` | Convert |

### 2.2 `Hashset[T comparable]`

Set backed by `map[T]bool` with embedded `sync.Mutex`.

```go
hs := coregeneric.New.Hashset.String.Empty()
hs := coregeneric.New.Hashset.Int.Cap(100)
hs := coregeneric.HashsetFrom([]string{"a", "b"})
hs := coregeneric.HashsetFromMap(existingMap)

hs.Add("x")
hs.Contains("x")
hs.RemoveItem("x")
hs.ToSlice()      // []T
hs.Length()
```

### 2.3 `Hashmap[K comparable, V any]`

Map wrapper with the same locking convention as `Collection`.

```go
hm := coregeneric.New.Hashmap.StringInt.Empty()
hm.Set("a", 1)
v, ok := hm.Get("a")
hm.Delete("a")
hm.Length()
hm.Keys()    // []K
hm.Values()  // []V
```

### 2.4 `SimpleSlice[T any]`

Lighter than `Collection[T]` — no mutex, no lock-suffixed methods. Use when you own the slice exclusively.

```go
ss := coregeneric.New.SimpleSlice.String.Cap(8)
ss.Append("a")
ss.AppendFmt("%d : %s", i, name)
ss.Strings()   // []string (for string-typed only)
```

> Used heavily inside Style B test runners — see [`13-testing-patterns.md` §3](./13-testing-patterns.md#3-style-b---basetestcase--testwrapper-typed-slice--hetero-input).

### 2.5 `LinkedList[T any]`, `Pair[L, R]`, `Triple[A, B, C]`

Lower-volume containers; same construction pattern via `coregeneric.New.LinkedList`, `coregeneric.New.Pair`, `coregeneric.New.Triple`. Use `Pair` / `Triple` when you need a typed multi-return without declaring a struct.

---

## 3. `corestr` — String Collections

Located at `coredata/corestr/`. String-typed collection with batteries included for the most common case: a thread-safe list of strings.

```go
import "github.com/alimtvnetwork/core-v8/coredata/corestr"

collection := corestr.NewCollectionPtrUsingStrings(&values, 0)
collection.AddsLock("new item")
fmt.Println(collection.Length())
```

When to prefer `corestr.Collection` over `coregeneric.Collection[string]`:

| Reason | `corestr` | `coregeneric` |
|---|---|---|
| Need string-specific helpers (joining, sorting) | ✅ | ❌ |
| Need to interoperate with `coregeneric` generic helpers | ❌ | ✅ |
| Default for new code | ⚪ | ✅ — generic-first is the framework standard |

> The framework historically used `corestr` heavily; new code should prefer `coregeneric.New.Collection.String` for consistency.

---

## 4. `corejson` — JSON Pipeline

Located at `coredata/corejson/`. The single canonical entry point for JSON in the framework.

```go
import "github.com/alimtvnetwork/core-v8/coredata/corejson"

// Serialize
jsonStr,   err := corejson.Serialize.ToString(myStruct)
jsonBytes, err := corejson.Serialize.Raw(myStruct)

// Deserialize
err := corejson.Deserialize.UsingBytes(jsonBytes, &target)
err := corejson.Deserialize.FromTo(source, &target)  // deep copy via JSON

// Pretty print
pretty := corejson.NewPtr(myStruct).PrettyJsonString()
```

### Why `corejson` and not `encoding/json`?

1. **Consistent error categories** — wraps stdlib errors in `errcore.UnMarshallingFailedType` / `FailedToConvertType` so log scanners can attribute failures.
2. **`FromTo` deep-copy** — common pattern (serialize then deserialize into another shape) is one call instead of two.
3. **Pretty-printing baked in** — no manual `json.MarshalIndent` formatting.

### Rule

Any package that touches JSON should import `corejson` and **never** `encoding/json` directly.

---

## 5. `coreonce` — Compute-Once Values

Located at `coredata/coreonce/`. Lazy-evaluated cached values for all common types — read-modify-once semantics built on `sync.Once`.

```go
import "github.com/alimtvnetwork/core-v8/coredata/coreonce"

// Construct with the producer
once := coreonce.New.String(func() string { return expensiveCall() })

// First Value() call invokes the producer; later calls return cached
v := once.Value()
v = once.Value()  // no recompute
```

Use for expensive package-level computations (regex compilation, file reads, network handshakes) that should run at most once.

---

## 6. `corepayload` — `PayloadWrapper`

Located at `coredata/corepayload/`. Wire-format envelope for payload-bearing messages.

```go
import "github.com/alimtvnetwork/core-v8/coredata/corepayload"

// Empty placeholder
payload := corepayload.New.PayloadWrapper.Empty()

// From an instruction
payload = corepayload.New.PayloadWrapper.UsingInstruction(&corepayload.PayloadCreateInstruction{
    Name:       "user-create",
    Identifier: "usr-123",
    Payloads:   myStruct,
})
```

### When to use

- Inter-service messages where the payload type is dynamic.
- Storing heterogeneous data in a single field that a downstream consumer will type-assert.
- Building command/event buses with structured metadata (`Name`, `Identifier`).

### When **not** to use

- Plain in-process function returns — use a typed struct.
- Cases where the consumer knows the concrete payload type — use that type directly.

---

## 7. Choosing the Right Container

Decision matrix:

| Need | Use |
|---|---|
| Generic, mutex-protected, slice-backed | `coregeneric.New.Collection.<Type>` |
| Generic, mutex-protected, set semantics | `coregeneric.New.Hashset.<Type>` |
| Generic, mutex-protected, key→value | `coregeneric.New.Hashmap.<KV>` |
| Lightweight slice, no concurrency | `coregeneric.New.SimpleSlice.<Type>` |
| String-only, legacy code | `corestr.Collection` |
| Two- or three-tuple return | `coregeneric.New.Pair` / `Triple` |
| Linked-list semantics (cheap insert) | `coregeneric.New.LinkedList` |
| Cached lazy value | `coreonce.New.<Type>` |
| Wire-format envelope | `corepayload.New.PayloadWrapper` |
| JSON serialize / deserialize | `corejson.Serialize` / `Deserialize` |

> Generic-first rule: when both a generic and a string-typed option exist, prefer the generic version unless a string-specific helper is essential.

---

## See Also

- [`02-design-philosophy.md`](./02-design-philosophy.md) — `newCreator` pattern (§5), zero-nil safety (§4)
- [`05-enum-system.md`](./05-enum-system.md) — `enumimpl` builds on the same `New` namespace pattern
- [`07-conditional-and-utilities.md`](./07-conditional-and-utilities.md) — `regexnew` and `coreonce` pair well for cached compiled regex
- [`13-testing-patterns.md`](./13-testing-patterns.md) — `SimpleSlice` is heavily used in Style B test runners
- [`/spec/00-llm-integration-guide.md` §coregeneric](../00-llm-integration-guide.md#coregeneric--generic-data-structures-api-reference) — Full API tables
