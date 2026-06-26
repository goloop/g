// Package g is a comprehensive utility library for Go 1.24+ that provides
// a rich set of generic, type-safe helper functions to streamline common
// programming tasks. Built with modern Go features, particularly generics,
// this package offers efficient, reliable, and easy-to-use solutions for
// everyday development challenges.
//
// This is the v2 module (github.com/goloop/g/v2). Import it as:
//
//	import g "github.com/goloop/g/v2"
//
// # Core Features
//
// Type Conversion & Handling:
//   - String-to-X conversions (bool, int, float)
//   - X-to-String conversions (bool, int, float)
//   - Pointer creation and conditional pointer handling
//   - Type checking and verification utilities
//
// Date & Time Management:
//   - Flexible date parsing with multiple format support
//   - Python-style date formatting compatibility
//   - Time zone manipulation and conversion
//   - Date-to-string formatting with various templates
//
// Collection Operations:
//   - Functional programming helpers (Map, Filter, Reduce)
//   - Set operations (Union, Intersection, Difference)
//   - List manipulation (Sort, Shuffle, Reverse)
//   - Array searching and filtering
//   - Element ranking and lookup
//
// Mathematical Operations:
//   - Basic arithmetic; Sum wraps like the + operator, SafeSum detects overflow
//   - Value clamping to a range (Clamp)
//   - Statistical functions (Average, Median)
//   - Random number generation
//   - Range generation with custom steps
//   - Number property checking (Even, Odd, Whole)
//
// String Processing:
//   - Character filtering and preservation
//   - String cleaning and normalization
//   - Pattern-based string manipulation
//   - Whitespace and special character handling
//
// Logical Operations:
//   - Conditional evaluation (If)
//   - Value presence checking (Any, All)
//   - Zero-value detection
//   - Type-specific comparisons
//
// Excel-like Functions:
//   - HLOOKUP/VLOOKUP implementations
//   - Range operations
//   - Value ranking
//
// # Key Types and Interfaces
//
// The package provides several key types and interfaces:
//   - Numerable: Interface for numeric types
//   - Verifiable: Interface for comparable types
//   - Pair: Generic struct for paired values
//
// # Performance and Reliability
//
// The package favours correctness and predictable performance:
//   - Hot paths delegate to the standard library (slices, maps, cmp) for
//     well-tested, O(n log n) algorithms instead of bespoke ones
//   - Randomness uses the concurrency-safe, auto-seeded math/rand/v2 source
//   - No hidden goroutines: the helpers are simple, allocation-aware loops
//   - Edge cases (empty input, NaN, overflow, degenerate ranges) are handled
//     deterministically rather than left to panic
//
// # Example Usage
//
// Type conversion:
//
//	num, err := g.StringToInt("123")
//	str := g.IntToString(456)
//
// Date handling:
//
//	date, err := g.StringToDate("2023-12-01")
//	formatted, err := g.DateToString(date, "2006-01-02")
//
// Collection operations:
//
//	numbers := []int{3, 1, 4, 1, 5}
//	g.Sort(numbers)
//	unique := g.Distinct(numbers)
//
// Mathematical operations:
//
//	avg := g.Average(1, 2, 3, 4, 5)
//	random := g.Random(1, 10)
//
// String processing:
//
//	cleaned := g.Weed("Hello\t World\n")
//	preserved := g.Preserve("+1-234-567-8900", g.Numbers)
//
// Logical operations:
//
//	result := g.If(condition, trueVal, falseVal)
//	exists := g.Any(slice...)
//
// # Focus and Design Principles
//
// The package is designed with emphasis on:
//   - Type safety through generics
//   - Consistent and intuitive API design
//   - Comprehensive error handling
//   - Efficient performance
//   - Developer-friendly documentation
//   - Practical utility in real-world applications
//
// This library serves as a robust foundation for Go applications,
// providing tested, efficient implementations of commonly needed
// functionality while maintaining type safety and performance.
package g
