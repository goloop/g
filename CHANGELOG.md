# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.2.0]

Numeric correctness fixes, a new checked absolute value and honest conversion
errors. Two small behaviour changes are noted below.

### Added
- `AbsOk` returns the absolute value and a bool reporting whether it is exact;
  it is false only for the most negative value of a signed type.
- `ErrEmptyString` and `ErrInvalidBool` sentinels, returned by the `StringTo*`
  helpers and matchable with `errors.Is`.

### Fixed
- `Random` no longer collapses full-width integer ranges to the lower bound
  (`Random[uint64](0, MaxUint64)`, `Random[int64](MinInt64, MaxInt64)`); the
  span is computed in `uint64`.
- `Random` on extreme floating-point bounds no longer returns `+Inf`/`NaN`;
  the value is drawn as a convex combination of the bounds.
- `Range`/`Rangef` no longer panic or return a wrong slice at extreme bounds
  (`Range(-1, MaxInt)`, `Range(MinInt, MaxInt)`); the step count is computed in
  `uint64` and out-of-range sizes return `nil`.
- `IsEven`/`IsOdd` are correct for integers past 2^53 (computed in integer
  arithmetic instead of `float64`).
- `FloatToString` formats a named `~float32` type at 32-bit width instead of
  64-bit.
- `HLookup`/`VLookup` return the default value instead of panicking when the
  match falls beyond a shorter `result` slice.

### Changed
- **Behaviour:** `IsEmpty` now treats a non-nil empty map (and channel) as
  empty, like an empty slice, and treats a zero-filled array as empty by its
  zero value. The three container kinds previously disagreed.
- **Behaviour:** the `StringTo*` helpers return `ErrEmptyString` for empty input
  instead of an `errors.New("empty string and no default value")` whose text was
  wrong when a default was supplied. `StringToBool` returns `ErrInvalidBool` for
  an unrecognised string.

## [2.0.0]

The module path changes to `github.com/goloop/g/v2` and the minimum Go version
is now 1.24. Import it as `import g "github.com/goloop/g/v2"`.

### Changed (breaking)
- `Sum` now wraps on integer overflow like the `+` operator instead of silently
  returning the zero value. Use `SafeSum` for overflow detection.
- Removed the `ParallelTasks` function and the implicit concurrency in
  `In`, `All`, `Any`, `Min`, and `Max`. These are now linear or
  standard-library-backed operations with predictable performance.
- `IsPointer` and `IsNumber` accept `any` instead of `interface{}` and return
  `false` for a `nil` argument instead of panicking.
- `Sort` delegates to the standard library, guaranteeing `O(n log n)` and
  removing the previous `O(n²)` worst case on sorted input.
- `Min`/`Max` propagate `NaN` deterministically (independent of position).
- `Rank` is now 1-based (the best value ranks `1`), matching the spreadsheet
  RANK function it documents. A missing value still returns `-1`.
- `DateToString`/`DateToStrings` no longer reject layouts that keep only a
  subset of time fields. They format with the standard library and return a
  nil error (an unrecognized layout echoes its literal text).

### Fixed
- `Random` with a wide range in a narrow type (e.g. `Random[int8](-128, 127)`)
  no longer collapses to the lower bound; the span is computed in `int64`.
- `Random` no longer panics on degenerate ranges such as `Random(0)` or
  `Random(-1)`; it returns the zero value / lower bound.
- `IsEven`/`IsOdd` are correct for values outside the `int` range (parity is
  computed in `float64`).
- `Shuffle` uses `math/rand/v2`, is concurrency-safe, and no longer reseeds the
  global generator on every call.
- `Rangef` returns `nil` (consistent with `Range`) when the range exceeds
  `MaxRangeSize`.

### Added
- `IfFunc` — lazy variant of `If` that evaluates only the selected branch.
- `Clamp` — constrain a value to an inclusive range.
- `MinOk`/`MaxOk` and `MinListOk`/`MaxListOk` — extremum helpers that report
  whether any value was present.
- `ErrOverflow` — sentinel error wrapped by `SafeSum`.
- `ErrParseDate` — sentinel error wrapped by `StringToDate`.

### Performance
- `Rank` runs in a single `O(n)` pass instead of `O(n²)`.
- Set operations and `Distinct` use `map[T]struct{}` with preallocation.
- Removed `reflect` from `Random*`, `Min`/`Max` in favour of plain generics.
  `IsEmpty`/`IsFalse` (used by `All`/`Any`/`Value`) now take a scalar fast path
  and only fall back to `reflect` for composite types.
