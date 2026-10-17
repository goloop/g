[![deps.dev](https://img.shields.io/badge/deps.dev-insights-4c8dbc)](https://deps.dev/go/github.com%2Fgoloop%2Fg%2Fv2) [![Go Reference](https://pkg.go.dev/badge/github.com/goloop/g/v2.svg)](https://pkg.go.dev/github.com/goloop/g/v2) [![License](https://img.shields.io/badge/license-MIT-brightgreen?style=flat)](https://github.com/goloop/g/blob/v2/LICENSE) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)

# g

`g` is a comprehensive utility library for Go — a flat collection of small,
generic, type-safe helpers for the tasks you write over and over: conditionals,
collection and set operations, math, string cleaning, date parsing and more.

It favours correctness and predictable performance: hot paths delegate to the
standard library (`slices`, `maps`, `cmp`, `math/rand/v2`), edge cases (empty
input, `NaN`, overflow, degenerate ranges) are handled deterministically instead
of panicking, and there is no hidden concurrency.

## Features

- **A ternary you can read** — `If`/`IfFunc` (eager and lazy).
- **Collections** — `Sort`, `Reverse`, `Shuffle`, `Distinct`, `Map`, `Filter`,
  `Reduce`, `Zip`, `CartesianProduct`.
- **Set algebra** — `Union`, `Intersection`, `Difference`, `SymmetricDifference`.
- **Math** — `Clamp`, `Sum`/`SafeSum` (overflow-aware), `Average`, `Median`,
  `Min`/`Max` with `…Ok` and `…List` variants.
- **Randomness** — `Random`, `RandomList`, `RandomMap` on the concurrency-safe
  `math/rand/v2` source; never panics on degenerate ranges.
- **Strings** — `StringToInt`/`Float`/`Bool`, `Weed`/`Preserve`/`Trim` with
  ready-made character-set constants.
- **Date & time** — `StringToDate`, `DateToString`, time-zone helpers.
- **Spreadsheet helpers** — `HLookup`/`VLookup`, `Rank`, `Range`/`Rangef`.

## Installation

```bash
go get github.com/goloop/g/v2
```

```go
import g "github.com/goloop/g/v2"
```

Requires Go 1.24 or newer. The package has no third-party dependencies.

## Quick start

```go
package main

import (
    "fmt"

    g "github.com/goloop/g/v2"
)

func main() {
    // A readable ternary.
    max := g.If(3 > 5, 3, 5) // 5

    // Membership without a manual loop.
    ok := g.In("PATCH", "GET", "POST", "PATCH") // true

    // Functional pipeline.
    nums := []int{3, 1, 4, 1, 5, 9, 2}
    g.Sort(nums)                                        // in place
    evens := g.Filter(nums, func(n int) bool { return n%2 == 0 })
    doubled := g.Map(evens, func(n int) int { return n * 2 })

    // Overflow-aware math.
    sum, err := g.SafeSum(nums...)

    fmt.Println(max, ok, doubled, sum, err)
}
```

## The API at a glance

| Area | Functions |
|------|-----------|
| Logic | `If`, `IfFunc`, `All`/`Any`, `In`/`Contains`, `Index`, `IsEmpty`, `Value` |
| Pointers | `Ptr`, `PtrIf` |
| Math | `Abs`, `Clamp`, `Sum`/`SafeSum`, `Product`, `Min`/`Max` (+`Ok`/`List`), `IsEven`/`IsOdd`/`IsWhole` |
| Stats & random | `Average`, `Median`, `Random`, `RandomList`/`RandomMap` (+`Plural`) |
| Collections | `Sort`, `Reverse`, `Shuffle`, `Distinct`, `Merge`, `Zip`, `CartesianProduct` |
| Functional | `Map`, `Filter`, `Reduce` |
| Sets | `Union`, `Intersection`, `Difference`, `SymmetricDifference`, `Complement` |
| Strings | `StringToInt`/`Float`/`Bool`, `…ToString`, `Weed`/`Preserve`/`Trim` |
| Date & time | `StringToDate`, `DateToString(s)`, `ChangeTimeZone`/`SetTimeZone`/`MoveTimeZone` |
| Lookup | `HLookup`/`VLookup`, `Rank`, `Range`/`Rangef` |

See **[DOC.md](DOC.md)** for the full reference — every function, signatures,
examples and recipes (Ukrainian: **[DOC.UK.md](DOC.UK.md)**).

## Documentation

- Full reference and recipes: [DOC.md](DOC.md) · [DOC.UK.md](DOC.UK.md)
- Package API: [pkg.go.dev/github.com/goloop/g/v2](https://pkg.go.dev/github.com/goloop/g/v2)
- Changes between versions: [CHANGELOG.md](CHANGELOG.md)

## Contributing

Contributions are welcome. Please run `go test ./...`, `go vet ./...` and
`gofmt -l .` before submitting a pull request.

## License

`g` is released under the MIT License. See [LICENSE](LICENSE).

## Release v2.3.1

Handle named floating-point overflow and nil pointers, and simplify intersection tracking.
