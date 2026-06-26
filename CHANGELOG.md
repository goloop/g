# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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

### Fixed
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
- Removed `reflect` from hot paths (`Value`, `Random*`, `Min`/`Max`) in favour
  of plain generics.
