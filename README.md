[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/g)](https://goreportcard.com/report/github.com/goloop/g) [![License](https://img.shields.io/badge/godoc-A+-brightgreen?style=flat)](https://godoc.org/github.com/goloop/g) [![License](https://img.shields.io/badge/license-MIT-brightgreen?style=flat)](https://github.com/goloop/g/blob/master/LICENSE) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)

# G

The **g** package is a comprehensive utility library for Go that provides a rich collection of generic helper functions to streamline common programming tasks. Built with modern Go features, particularly generics, this package offers efficient, reliable, and easy-to-use solutions for everyday development challenges.

## Installation

To install the package, use `go get`:

```bash
go get github.com/goloop/g/v2
```

Import it as:

```go
import g "github.com/goloop/g/v2"
```

**Note**: This is the **v2** module and requires **Go 1.24 or later**. It uses
generics and delegates hot paths to the standard library (`slices`, `maps`,
`cmp`, `math/rand/v2`). See [What's new in v2](#whats-new-in-v2) for the
breaking changes from v1.

## Why This Package?

### Ternary Operator Alternative
In languages like C++ and Python, you can write concise conditional expressions:

```cpp
// C++
int max = (a > b) ? a : b;

// Python
max = a if a > b else b
```

Go doesn't have a ternary operator, leading to verbose code:

```go
max := a
if a < b {
    max = b
}
```

With the `g` package, you can write:

```go
max := g.If(a > b, a, b)
```

### Efficient List Operations
Python makes checking if an element is in a slice easy:

```python
if a in some_slice:
    # do something
```

The `g` package provides a clean, type-safe wrapper:

```go
if g.In(a, someSlice...) {
    // do something
}
```

## Key Features

### Type Conversion & Validation
- String to various types (bool, int, float)
- Type checking and verification
- Overflow-aware summation via `SafeSum`

```go
// String to int conversion with default value
num, err := g.StringToInt("123", 0)

// Overflow-aware sum: returns an error wrapping g.ErrOverflow on overflow
sum, err := g.SafeSum(math.MaxInt64, 1)
```

### Mathematical Operations
- `Sum` wraps like the `+` operator; `SafeSum` detects overflow
- Value clamping with `Clamp`
- Statistical functions (Average, Median)
- Random number generation (never panics on degenerate ranges)
- Number properties (Even, Odd, Whole)

```go
avg := g.Average(1, 2, 3, 4, 5)
median := g.Median(1, 2, 3, 4, 5)
random := g.Random(1, 10)      // a random int in [1, 10)
clamped := g.Clamp(15, 0, 10)  // 10
```

### Collection Operations
- Set operations (Union, Intersection, Difference)
- List manipulation (Sort, Shuffle, Reverse)
- Functional programming helpers (Map, Filter, Reduce)

```go
unique := g.Union(slice1, slice2)
g.Sort(numbers)
doubled := g.Map(numbers, func(n int) int { return n * 2 })
```

### String Processing
- Character filtering and preservation
- String cleaning and normalization
- Pattern-based manipulation

```go
// Remove unwanted characters
cleaned := g.Weed("Hello\t World\n")

// Keep only specific characters
numbers := g.Preserve("+1-234-567-8900", g.Numbers)
```

### Date & Time
- Flexible date parsing
- Time zone manipulation
- Python-style date formatting

```go
date, err := g.StringToDate("2023-12-01")
newTime, err := g.ChangeTimeZone(time.Now(), "America/New_York")
```

### Excel-like Functions
- HLOOKUP/VLOOKUP implementations
- Range operations
- Value ranking

```go
rank := g.Rank(7, []float64{1, 5, 2, 3, 7, 8})
value := g.HLookup("key", lookupSlice, resultSlice, defaultValue)
```

## Complete Function List

[View the complete function documentation](https://pkg.go.dev/github.com/goloop/g)

Here are some key function categories:

### Basic Operations
- `If` - Ternary operator alternative (eager)
- `IfFunc` - Lazy ternary: evaluates only the chosen branch
- `In` - Check if element exists in slice
- `All`/`Any` - Check conditions across values
- `IsEmpty`/`IsWhole`/`IsEven`/`IsOdd` - Value validation

### Mathematical
- `Min`/`Max` - Find extremes (`MinOk`/`MaxOk` report the empty case)
- `MinList`/`MaxList` - Extremes of a list with defaults (`*ListOk` variants)
- `Clamp` - Constrain a value to a range
- `Sum`/`SafeSum` - Wrapping addition / overflow-aware addition (`ErrOverflow`)
- `Average`/`Median` - Statistical calculations
- `Random`/`RandomList` - Random value generation

### Collection Operations
- `Union`/`Intersection`/`Difference`/`SymmetricDifference` - Set operations
- `Sort`/`Shuffle`/`Reverse` - List manipulation
- `Map`/`Filter`/`Reduce` - Functional programming
- `Zip`/`CartesianProduct` - List combinations

### String Operations
- `StringToInt`/`StringToFloat`/`StringToBool` - String parsing
- `Weed`/`Preserve`/`Trim` - String cleaning
- `IntToString`/`FloatToString`/`BoolToString` - Value formatting

### Date & Time
- `StringToDate`/`DateToString` - Date parsing and formatting
- `ChangeTimeZone`/`SetTimeZone`/`MoveTimeZone` - Time zone operations

### Excel-like Functions
- `HLookup`/`VLookup` - Value lookups
- `Rank` - Value ranking
- `Range`/`Rangef` - Range generation

## What's new in v2

v2 modernises the package for Go 1.24 and fixes a number of long-standing
correctness issues. The import path changes to `github.com/goloop/g/v2`.

**Breaking changes**

- `Sum` now wraps on integer overflow, exactly like the `+` operator, instead
  of silently returning `0`. Use `SafeSum` when you need overflow detection.
- The `ParallelTasks` function and the implicit "parallel by default" behaviour
  of `In`/`All`/`Any`/`Min`/`Max` were removed. These helpers are now simple
  linear or standard-library-backed operations with predictable performance.
- `IsPointer` and `IsNumber` take `any` instead of `interface{}` and no longer
  panic on a `nil` argument.

**Fixes**

- `Random` no longer panics on `Random(0)`, `Random(-1)`, or other degenerate
  ranges — it returns the zero value / lower bound instead.
- `IsEven`/`IsOdd` are correct for values outside the `int` range (e.g. `1e19`)
  because parity is computed in `float64` without an `int` conversion.
- `Min`/`Max` handle `NaN` deterministically (NaN propagates regardless of
  position) by delegating to the standard library.
- `Sort` is guaranteed `O(n log n)` (standard-library introsort) and no longer
  degrades to `O(n²)` on already-sorted input.
- `Shuffle` uses the concurrency-safe `math/rand/v2` source and is safe to call
  from multiple goroutines.

**New API**

- `IfFunc` — lazy `If`.
- `Clamp` — constrain a value to `[lo, hi]`.
- `MinOk`/`MaxOk`, `MinListOk`/`MaxListOk` — extremum with an explicit
  "was there any value?" boolean.
- `ErrOverflow` — sentinel error returned (wrapped) by `SafeSum`.
- `ErrParseDate` — sentinel error returned (wrapped) by `StringToDate`.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)

