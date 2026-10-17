# g — reference

The full reference for the `g` package: the mental model, every function
grouped by topic, the type constraints, and a set of practical recipes.

Ukrainian version: **[DOC.UK.md](DOC.UK.md)**.

## Contents

- [Mental model](#mental-model)
- [Type constraints](#type-constraints)
- [Logic and conditionals](#logic-and-conditionals)
- [Presence and value checks](#presence-and-value-checks)
- [Pointers](#pointers)
- [Numbers and math](#numbers-and-math)
- [Statistics and randomness](#statistics-and-randomness)
- [Collections](#collections)
- [Functional helpers](#functional-helpers)
- [Set operations](#set-operations)
- [Strings and parsing](#strings-and-parsing)
- [Date and time](#date-and-time)
- [Lookup and ranking](#lookup-and-ranking)
- [Ranges](#ranges)
- [Errors](#errors)
- [Recipes and tips](#recipes-and-tips)

## Mental model

`g` is a flat collection of small, generic helpers. There is no state, no
configuration and no hidden concurrency — every function is a pure, allocation
aware operation over its arguments.

Three ideas run through the whole package:

1. **Generics over duplication.** One `Map`, one `Sort`, one `Clamp` that works
   for every element type, constrained by a handful of interfaces
   (`Numerable`, `Verifiable`, `Ordered`).
2. **Delegate to the standard library on hot paths.** Sorting, min/max, set
   membership and randomness are backed by `slices`, `cmp` and `math/rand/v2`,
   so you get well-tested, `O(n log n)` behaviour and concurrency-safe
   randomness for free.
3. **Deterministic edge cases.** Empty input, `NaN`, overflow and degenerate
   ranges return a defined value instead of panicking. Where a result may be
   absent, an `…Ok` variant reports it with a boolean.

Import it under the short name:

```go
import g "github.com/goloop/g/v2"
```

## Type constraints

The helpers are constrained by a small set of interfaces:

| Constraint | Permits |
|------------|---------|
| `Integer`    | every signed and unsigned integer type |
| `Float`      | `float32`, `float64` |
| `Numerable`  | `Integer` ∪ `Float` — anything you can do arithmetic on |
| `Ordered`    | `Integer` ∪ `Float` ∪ `~string` — anything with `< <= >= >` |
| `Verifiable` | `Integer` ∪ `Float` ∪ `string` ∪ `rune` — comparable classic types |
| `comparable` | Go's built-in constraint, used by the set helpers |

`Pair[T, U]` is a generic struct (`First`, `Second`) used as the element type
of `Zip`.

## Logic and conditionals

```go
func If[L bool | trit.Trit, T any](e L, t, f T, u ...T) T
func IfFunc[T any](cond bool, t, f func() T) T
```

`If` is the ternary operator Go doesn't have. The condition may be a `bool` or
a `trit.Trit`; with a `Trit` you can pass an optional third value `u` for the
`Unknown` state.

```go
max := g.If(a > b, a, b)
label := g.If(t, "yes", "no", "maybe") // t is a trit.Trit
```

`If` is **eager** — both `t` and `f` are evaluated before the call. When a
branch is expensive or unsafe to evaluate unconditionally, use `IfFunc`, which
runs only the chosen branch:

```go
cfg := g.IfFunc(hasFile, func() Config { return load() }, func() Config { return Default() })
```

## Presence and value checks

```go
func All[T any](v ...T) bool          func AllList[T any](v []T) bool
func Any[T any](v ...T) bool          func AnyList[T any](v []T) bool
func Contains[T Verifiable](v T, vs []T) bool
func In[T Verifiable](v T, list ...T) bool
func Index[T comparable](vs []T, v T) int
func IsEmpty[T any](v T) bool
func IsTrue[T any](v T) bool           func IsFalse[T any](v T) bool
func Value[T any](v ...T) T
```

`All`/`Any` report whether every / any argument is a non-zero ("truthy") value.
The `…List` variants take a slice instead of a variadic list.

`In` and `Contains` are membership tests (`In` is variadic, `Contains` takes a
slice). `Index` returns the position of the first match, or `-1`.

`IsEmpty` is true for the zero value of any type. `IsTrue`/`IsFalse` are the
single-value forms. `Value` returns the first argument, or the zero value when
called with none — handy for "optional variadic" APIs.

```go
if g.In(status, 200, 201, 204) { /* ... */ }
name := g.Value(names...) // "" if names is empty
```

## Pointers

```go
func Ptr[T any](v ...T) *T
func PtrIf[T any](exp bool, v ...T) *T
```

`Ptr` returns a pointer to its argument (or to the zero value when called with
none) — useful for building structs whose fields are pointers. `PtrIf` returns
the pointer only when `exp` is true, otherwise `nil`.

```go
req := Request{Timeout: g.Ptr(30)}       // *int
req.Retry = g.PtrIf(enabled, 3)          // *int or nil
```

## Numbers and math

```go
func Abs[T Numerable](v T) T
func Clamp[T Ordered](v, lo, hi T) T
func Sum[T Numerable](v ...T) T
func SafeSum[T Numerable](v ...T) (T, error)
func Product[T Numerable](v ...T) T
func Min[T Verifiable](v ...T) T        func Max[T Verifiable](v ...T) T
func MinOk[T Verifiable](v ...T) (T, bool)
func MaxOk[T Verifiable](v ...T) (T, bool)
func MinList[T Verifiable](v []T, defaults ...T) T
func MaxList[T Verifiable](v []T, defaults ...T) T
func MinListOk[T Verifiable](v []T, defaults ...T) (T, bool)
func MaxListOk[T Verifiable](v []T, defaults ...T) (T, bool)
func IsEven[T Numerable](v T, f ...bool) bool
func IsOdd[T Numerable](v T, f ...bool) bool
func IsWhole[T Numerable](v T) bool
```

`Sum` wraps on overflow, exactly like the `+` operator. When you need to
**detect** overflow, use `SafeSum`, which returns an error wrapping
[`ErrOverflow`](#errors).

```go
total := g.Sum(1, 2, 3)                  // 6
_, err := g.SafeSum(math.MaxInt64, 1)    // errors.Is(err, g.ErrOverflow)
```

`Clamp` constrains a value to `[lo, hi]`:

```go
g.Clamp(15, 0, 10) // 10
g.Clamp(-3, 0, 10) // 0
```

`Min`/`Max` delegate to the standard library and propagate `NaN`
deterministically. They return the zero value on empty input; the `…Ok`
variants report whether any value was present. The `…List` forms take a slice
plus optional fallbacks used when the slice is empty.

`IsEven`/`IsOdd` compute parity in `float64`, so they stay correct for values
beyond the `int` range (e.g. `1e19`). `IsWhole` reports whether a float has no
fractional part.

## Statistics and randomness

```go
func Average[T Numerable](v ...T) float64
func Median[T Numerable](v ...T) float64
func Random[T Numerable](v ...T) T
func RandomList[T any](v []T) T
func RandomListPlural[T any](n int, v []T) []T
func RandomMap[K comparable, T any](m map[K]T) T
func RandomMapPlural[K comparable, T any](n int, m map[K]T) []T
```

`Average` and `Median` return `float64` for any numeric input. `Median` handles
even-length inputs by averaging the two middle values.

`Random` never panics on a degenerate range: `Random()` returns the zero value,
`Random(n)` returns a value in `[0, n)`, and `Random(a, b)` returns a value in
`[a, b)`; a non-positive or inverted range yields the lower bound. `RandomList`
/ `RandomMap` pick one element; the `…Plural` variants pick `n`. Randomness uses
the concurrency-safe `math/rand/v2` source, so these are safe to call from
several goroutines.

```go
n := g.Random(1, 10)              // int in [1, 10)
pick := g.RandomList(options)     // one element
three := g.RandomListPlural(3, options)
```

## Collections

```go
func Sort[T Verifiable](v []T, inverse ...bool)
func Reverse[T any](v []T)
func Shuffle[T any](v []T)
func Distinct[T comparable](v []T) []T
func Merge[T Verifiable](a []T, b []T, sort ...bool) []T
func Zip[T, U any](a []T, b []U) []Pair[T, U]
func CartesianProduct[T any](a []T, b []T) [][2]T
```

`Sort` mutates the slice in place using the standard-library introsort
(guaranteed `O(n log n)`); pass `true` for descending order. `Reverse` and
`Shuffle` also mutate in place. `Distinct` returns a new slice with duplicates
removed, preserving first-seen order.

`Merge` concatenates two slices (optionally sorting the result). `Zip` pairs two
slices element-wise into `[]Pair`; `CartesianProduct` produces every `[a, b]`
combination.

```go
g.Sort(nums)                 // ascending, in place
g.Sort(nums, true)           // descending
u := g.Distinct([]int{1,1,2}) // [1 2]
pairs := g.Zip(keys, values) // []Pair[K, V]
```

## Functional helpers

```go
func Map[T any, U any](vs []T, f func(T) U) []U
func Filter[T any](vs []T, f func(T) bool) []T
func Reduce[T any, U any](vs []T, f func(U, T) U, init U) U
```

The classic trio. `Map` transforms each element (the result type may differ),
`Filter` keeps elements matching a predicate, `Reduce` folds a slice into a
single accumulator.

```go
doubled := g.Map(nums, func(n int) int { return n * 2 })
evens := g.Filter(nums, func(n int) bool { return n%2 == 0 })
total := g.Reduce(nums, func(acc, n int) int { return acc + n }, 0)
```

## Set operations

```go
func Union[T comparable](a []T, b []T) []T
func Intersection[T comparable](a []T, b []T) []T
func Difference[T comparable](a []T, b []T) []T        // alias: Diff
func SymmetricDifference[T comparable](a []T, b []T) []T // alias: Sdiff
func Complement[T comparable](a []T, b []T) []T
```

Slice-based set algebra. `Union` returns the distinct elements of both slices;
`Intersection` the elements in both; `Difference` (`Diff`) the elements in `a`
but not `b`; `SymmetricDifference` (`Sdiff`) the elements in exactly one;
`Complement` the elements of `b` not present in `a`.

```go
g.Union([]int{1,2}, []int{2,3})        // [1 2 3]
g.Intersection([]int{1,2}, []int{2,3}) // [2]
g.Difference([]int{1,2}, []int{2,3})   // [1]
```

## Strings and parsing

```go
func StringToBool(v string, def ...bool) (bool, error)
func StringToInt(v string, def ...int) (int, error)
func StringToFloat(v string, def ...float64) (float64, error)
func BoolToString(b bool) string
func IntToString[T Integer](v T) string
func FloatToString[T ~float32 | ~float64](v T) string
func Weed(s string, patterns ...string) string
func Preserve(s string, patterns ...string) string
func Trim(s string, patterns ...string) string
```

The `StringToX` parsers accept an optional default returned when parsing fails
(they still return the error so you can decide what to do). The `XToString`
formatters are the inverse.

`Weed` removes the characters listed in `patterns`, `Preserve` keeps only those
characters, and `Trim` strips them from both ends. Ready-made pattern constants
are provided: `Numbers`, `Letters`, `Symbols`, `Whitespaces`. When no pattern
is given, whitespace is the default target.

```go
n, _ := g.StringToInt("123")             // 123
n, _ = g.StringToInt("x", 42)            // 42 (+ error)
digits := g.Preserve("+1-234-567", g.Numbers) // "1234567"
clean := g.Weed("Hello\t World\n")       // "Hello World"
```

## Date and time

```go
func StringToDate(s string, patterns ...string) (time.Time, error)
func DateToString(t time.Time, patterns ...string) (string, error)
func DateToStrings(t time.Time, patterns ...string) ([]string, error)
func ChangeTimeZone(t time.Time, timezone string) (time.Time, error)
func SetTimeZone(t time.Time, timezone string) (time.Time, error)
func MoveTimeZone(t time.Time, tz int) time.Time
```

`StringToDate` tries a list of layouts (or a built-in set of common formats
when none is given) and returns an error wrapping [`ErrParseDate`](#errors) when
nothing matches. `DateToString` formats a time with the first layout that
applies; `DateToStrings` returns one string per layout.

The time-zone helpers differ in intent:

- `ChangeTimeZone` — convert the instant to another zone (the wall clock
  changes, the instant is the same).
- `SetTimeZone` — reinterpret the same wall-clock reading in another zone (the
  instant changes).
- `MoveTimeZone` — shift by a fixed offset in hours.

```go
d, _ := g.StringToDate("2023-12-01")
s, _ := g.DateToString(d, "2006-01-02")
ny, _ := g.ChangeTimeZone(time.Now(), "America/New_York")
```

## Lookup and ranking

```go
func HLookup[T comparable, U any](v T, lookup []T, result []U, def U) U
func VLookup[T comparable, U any](v T, lookup []T, result []U, def U) U
func Rank[T Verifiable](number T, array []T, ascending ...bool) int
```

Spreadsheet-style helpers. `HLookup`/`VLookup` find `v` in `lookup` and return
the value at the same index in `result`, or `def` when there is no match.
`Rank` returns the 1-based rank of `number` within `array` (descending by
default; pass `true` for ascending).

```go
label := g.VLookup(code, codes, labels, "unknown")
r := g.Rank(7, []float64{1, 5, 2, 3, 7, 8}) // 2 (7 is the 2nd largest)
```

## Ranges

```go
func Range(a int, opt ...int) []int
func Rangef[T any](fn func(int) T, a int, opt ...int) []T
```

`Range` builds an integer slice, following the Python `range` convention:
`Range(n)` is `[0, n)`, `Range(a, b)` is `[a, b)`, and `Range(a, b, step)` adds
a step. `Rangef` maps a function over the same indices, producing `[]T`. The
result is capped by `MaxRangeSize` to guard against accidental huge allocations.

```go
g.Range(5)          // [0 1 2 3 4]
g.Range(2, 8, 2)    // [2 4 6]
g.Rangef(func(i int) string { return strconv.Itoa(i) }, 3) // ["0" "1" "2"]
```

## Errors

```go
var ErrOverflow  = errors.New("g: numeric overflow")
var ErrParseDate = errors.New("g: date and time could not be recognized")
```

Both are sentinels returned **wrapped**, so match them with `errors.Is`:

```go
if _, err := g.SafeSum(a, b); errors.Is(err, g.ErrOverflow) { /* ... */ }
if _, err := g.StringToDate(s); errors.Is(err, g.ErrParseDate) { /* ... */ }
```

## Recipes and tips

**Default a possibly-empty optional.** `Value` collapses a variadic argument to
one value with a zero-value fallback:

```go
func greet(name ...string) string {
    return "Hi, " + g.If(g.IsEmpty(g.Value(name...)), "friend", g.Value(name...))
}
```

**Avoid an expensive default.** Reach for `IfFunc`/`PtrIf` so the unused branch
never runs:

```go
conn := g.IfFunc(cached, func() *Conn { return pool.Get() }, func() *Conn { return dial() })
```

**Overflow-safe accumulation.** Prefer `SafeSum` in code that adds
externally-supplied numbers; keep `Sum` for known-small inputs where wrapping is
acceptable.

**Deterministic sampling in tests.** `Random*` use `math/rand/v2`; for
reproducible tests, sample from a fixed slice with `RandomListPlural` and assert
on membership rather than on a specific value.

**Clean user input.** Chain `Weed`/`Preserve`/`Trim` with the pattern constants
to normalise phone numbers, codes and identifiers without a regexp.

## Release v2.3.1

Handle named floating-point overflow and nil pointers, and simplify intersection tracking.
