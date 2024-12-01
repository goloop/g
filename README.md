[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/g)](https://goreportcard.com/report/github.com/goloop/g) [![License](https://img.shields.io/badge/godoc-A+-brightgreen?style=flat)](https://godoc.org/github.com/goloop/g) [![License](https://img.shields.io/badge/license-MIT-brightgreen?style=flat)](https://github.com/goloop/g/blob/master/LICENSE) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)

# G

The **g** package is a comprehensive utility library for Go that provides a rich collection of generic helper functions to streamline common programming tasks. Built with modern Go features, particularly generics, this package offers efficient, reliable, and easy-to-use solutions for everyday development challenges.

## Installation

To install the package, use `go get`:

```bash
go get github.com/goloop/g
```

**Note**: This package requires Go 1.20 or later due to its extensive use of generics.

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

The `g` package provides an efficient concurrent implementation:

```go
if g.In(a, someSlice...) {
    // do something
}
```

## Key Features

### Type Conversion & Validation
- String to various types (bool, int, float)
- Type checking and verification
- Safe numeric conversions with overflow protection

```go
// String to int conversion with default value
num, err := g.StringToInt("123", 0)

// Safe sum with overflow protection
sum, err := g.SafeSum(1, math.MaxInt64)
```

### Mathematical Operations
- Basic arithmetic with overflow protection
- Statistical functions (Average, Median)
- Random number generation
- Number properties (Even, Odd, Whole)

```go
avg := g.Average(1, 2, 3, 4, 5)
median := g.Median(1, 2, 3, 4, 5)
random := g.Random(1, 10)
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
- `If` - Ternary operator alternative
- `In` - Check if element exists in slice
- `All`/`Any` - Check conditions across values
- `IsEmpty`/`IsWhole`/`IsEven`/`IsOdd` - Value validation

### Mathematical
- `Min`/`Max` - Find extremes
- `Sum`/`SafeSum` - Addition with optional overflow protection
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

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)

